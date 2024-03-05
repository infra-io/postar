// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gomail

import (
	"context"
	"fmt"
	"sync"

	"github.com/FishGoddess/rego"
	"github.com/wneessen/go-mail"
)

type Pool struct {
	clientPools map[string]*rego.Pool[*mail.Client]
	limit       uint64

	lock sync.RWMutex
}

func NewPool(limit uint64) *Pool {
	pool := &Pool{
		clientPools: make(map[string]*rego.Pool[*mail.Client], 16),
		limit:       limit,
	}

	return pool
}

func (p *Pool) clientKey(host string, port int32, username string, password string, smtpAuth string) string {
	return fmt.Sprintf("%s:%d/%s:%s/%s", host, port, username, password, smtpAuth)
}

func (p *Pool) Put(host string, port int32, username string, password string, smtpAuth string, client *mail.Client) error {
	clientKey := p.clientKey(host, port, username, password, smtpAuth)

	p.lock.RLock()
	clientPool, ok := p.clientPools[clientKey]
	p.lock.RUnlock()

	if ok {
		return clientPool.Put(client)
	}

	return nil
}

func (p *Pool) Take(ctx context.Context, host string, port int32, username string, password string, smtpAuth string) (*mail.Client, error) {
	clientKey := p.clientKey(host, port, username, password, smtpAuth)

	p.lock.RLock()
	clientPool, ok := p.clientPools[clientKey]
	p.lock.RUnlock()

	if ok {
		return clientPool.Take(ctx)
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	clientPool, ok = p.clientPools[clientKey]
	if ok {
		return clientPool.Take(ctx)
	}

	acquire := func() (*mail.Client, error) {
		return mail.NewClient(
			host, mail.WithPort(int(port)),
			mail.WithUsername(username), mail.WithPassword(password),
			mail.WithSMTPAuth(mail.SMTPAuthType(smtpAuth)), mail.WithLogger(Logger{}),
		)
	}

	release := func(client *mail.Client) error {
		return client.Close()
	}

	clientPool = rego.New(acquire, release, rego.WithLimit(uint64(p.limit)))
	p.clientPools[clientKey] = clientPool

	return clientPool.Take(ctx)
}

func (p *Pool) Close() (err error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	for _, clientPool := range p.clientPools {
		if closeErr := clientPool.Close(); closeErr != nil {
			err = closeErr
		}
	}

	return err
}
