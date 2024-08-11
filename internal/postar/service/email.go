// Copyright 2022 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"context"

	"github.com/FishGoddess/errors"
	"github.com/FishGoddess/logit"
	"github.com/infra-io/postar/config"
	"github.com/infra-io/postar/internal/postar/model"
	"github.com/infra-io/postar/pkg/aes"
	"github.com/infra-io/postar/pkg/gomail"
	"github.com/infra-io/postar/pkg/grpc/contextutil"
	pkgtemplate "github.com/infra-io/postar/pkg/template"
	"github.com/wneessen/go-mail"
)

type SpaceStore interface {
	GetSpace(ctx context.Context, spaceID int32) (*model.Space, error)
}

type AccountStore interface {
	GetAccount(ctx context.Context, spaceID int32, accountID int32) (*model.Account, error)
}

type TemplateStore interface {
	GetTemplate(ctx context.Context, spaceID int32, templateID int64) (*model.Template, error)
}

type EmailService interface {
	SendEmail(ctx context.Context, email *model.Email) (err error)
}

type defaultEmailService struct {
	conf *config.PostarConfig

	spaceStore    SpaceStore
	accountStore  AccountStore
	templateStore TemplateStore
}

func NewEmailService(conf *config.PostarConfig, spaceStore SpaceStore, accountStore AccountStore, templateStore TemplateStore) EmailService {
	service := &defaultEmailService{
		conf:          conf,
		spaceStore:    spaceStore,
		accountStore:  accountStore,
		templateStore: templateStore,
	}

	return service
}

func (des *defaultEmailService) checkSpace(ctx context.Context) (spaceID int32, err error) {
	spaceID = contextutil.GetSpaceID(ctx)
	spaceToken := contextutil.GetSpaceToken(ctx)

	if spaceID <= 0 {
		return 0, errors.BadRequest("业务空间 %d <= 0", spaceID)
	}

	space, err := des.spaceStore.GetSpace(ctx, spaceID)
	if err != nil {
		return 0, err
	}

	if !space.Enabled() {
		return 0, errors.BadRequest("业务空间 (%d）%s 未启用", space.ID, space.Name)
	}

	decrypted, err := aes.Decrypt(des.conf.Crypto.AESKey, des.conf.Crypto.AESIV, space.Token)
	if err != nil {
		return 0, err
	}

	space.Token = decrypted

	if spaceToken != space.Token {
		return 0, errors.Forbidden("业务空间 （%d）%s 的令牌错误", space.ID, space.Name)
	}

	return spaceID, nil
}

func (des *defaultEmailService) getTemplate(ctx context.Context, spaceID int32, templateID int64) (*model.Template, error) {
	template, err := des.templateStore.GetTemplate(ctx, spaceID, templateID)
	if err != nil {
		return nil, err
	}

	if !template.Enabled() {
		err = errors.New("template not enabled")
		return nil, errors.BadRequest("模板 (%d）%s 未启用", template.ID, template.Name)
	}

	return template, nil
}

func (des *defaultEmailService) getAccount(ctx context.Context, spaceID int32, accountID int32) (*model.Account, error) {
	account, err := des.accountStore.GetAccount(ctx, spaceID, accountID)
	if err != nil {
		return nil, err
	}

	if !account.Enabled() {
		err = errors.New("account not enabled")
		return nil, errors.BadRequest("账号 (%d）%s 未启用", account.ID, account.Username)
	}

	decrypted, err := aes.Decrypt(des.conf.Crypto.AESKey, des.conf.Crypto.AESIV, account.Password)
	if err != nil {
		return nil, err
	}

	account.Password = decrypted
	return account, nil
}

func (des *defaultEmailService) determineRenderFunc(contentType model.EmailContentType) pkgtemplate.RenderFunc {
	if contentType == model.EmailContentTypePlain {
		return pkgtemplate.RenderPlain

	}

	if contentType == model.EmailContentTypeHTML {
		return pkgtemplate.RenderHTML
	}

	return pkgtemplate.RenderNone
}

func (des *defaultEmailService) combineTemplateEmail(template *model.Template, email *model.Email) (*model.TemplateEmail, error) {
	templateEmail := &template.Email
	templateEmail.To = append(templateEmail.To, email.To...)

	if len(templateEmail.To) <= 0 {
		return nil, errors.BadRequest("邮件收件人为空")
	}

	templateEmail.CC = append(templateEmail.CC, email.CC...)
	templateEmail.BCC = append(templateEmail.BCC, email.BCC...)

	renderFunc := des.determineRenderFunc(templateEmail.ContentType)

	render, err := renderFunc(templateEmail.Subject, email.SubjectParams)
	if err != nil {
		return nil, err
	}

	templateEmail.Subject = render

	render, err = renderFunc(templateEmail.Content, email.ContentParams)
	if err != nil {
		return nil, err
	}

	templateEmail.Content = render

	return templateEmail, nil
}

func (des *defaultEmailService) handleTemplateEmail(ctx context.Context, account *model.Account, email *model.TemplateEmail) (err error) {
	msg := mail.NewMsg()

	if err = msg.From(account.Username); err != nil {
		return err
	}

	if err = msg.To(email.To...); err != nil {
		return err
	}

	if err = msg.Cc(email.CC...); err != nil {
		return err
	}

	if err = msg.Bcc(email.BCC...); err != nil {
		return err
	}

	msg.Subject(email.Subject)
	msg.SetBodyString(mail.ContentType(email.ContentType.String()), email.Content)

	client, err := mail.NewClient(
		account.Host, mail.WithPort(int(account.Port)), mail.WithUsername(account.Username), mail.WithPassword(account.Password),
		mail.WithSMTPAuth(mail.SMTPAuthType(account.SMTPAuth.String())), mail.WithLogger(gomail.Logger{}),
	)

	return client.DialAndSendWithContext(ctx, msg)
}

func (des *defaultEmailService) sendEmail(ctx context.Context, email *model.Email) (err error) {
	var spaceID int32
	if spaceID, err = des.checkSpace(ctx); err != nil {
		return err
	}

	template, err := des.getTemplate(ctx, spaceID, email.TemplateID)
	if err != nil {
		return err
	}

	templateEmail, err := des.combineTemplateEmail(template, email)
	if err != nil {
		return err
	}

	account, err := des.getAccount(ctx, spaceID, template.AccountID)
	if err != nil {
		return err
	}

	if err = des.handleTemplateEmail(ctx, account, templateEmail); err != nil {
		return err
	}

	return nil
}

func (des *defaultEmailService) SendEmail(ctx context.Context, email *model.Email) (err error) {
	logger := logit.FromContext(ctx)
	logger.Debug("send email", "email", email)

	return des.sendEmail(ctx, email)
}
