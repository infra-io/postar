// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gomail

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/FishGoddess/rego"
	"github.com/wneessen/go-mail"
)

type Pool struct {
	clientPools map[string]*rego.Pool[*mail.Client]
	limit       uint64
	timeout     time.Duration

	lock sync.RWMutex
}

func NewPool(limit uint64, timeout time.Duration) *Pool {
	pool := &Pool{
		clientPools: make(map[string]*rego.Pool[*mail.Client], 16),
		limit:       limit,
		timeout:     timeout,
	}

	return pool
}

func (p *Pool) clientKey(host string, port int32, username string, smtpAuth string) string {
	return fmt.Sprintf("%s:%d/%s/%s", host, port, username, smtpAuth)
}

func (p *Pool) Put(host string, port int32, username string, password string, smtpAuth string, client *mail.Client) error {
	clientKey := p.clientKey(host, port, username, smtpAuth)

	p.lock.RLock()
	clientPool, ok := p.clientPools[clientKey]
	p.lock.RUnlock()

	if !ok {
		return nil
	}

	if err := client.Reset(); err != nil {
		return err
	}

	return clientPool.Put(client)
}

func (p *Pool) Take(ctx context.Context, host string, port int32, username string, password string, smtpAuth string) (*mail.Client, error) {
	clientKey := p.clientKey(host, port, username, smtpAuth)

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
		client, err := mail.NewClient(
			host, mail.WithPort(int(port)), mail.WithUsername(username), mail.WithPassword(password),
			mail.WithSMTPAuth(mail.SMTPAuthType(smtpAuth)), mail.WithTimeout(p.timeout), mail.WithLogger(Logger{}),
		)

		if err != nil {
			return nil, err
		}

		return client, client.DialWithContext(context.Background())
	}

	release := func(client *mail.Client) error {
		return client.Close()
	}

	clientPool = rego.New(acquire, release, rego.WithLimit(uint64(p.limit)))
	p.clientPools[clientKey] = clientPool

	return clientPool.Take(ctx)
}

func (p *Pool) Stats() map[string]rego.Status {
	p.lock.RLock()
	defer p.lock.RUnlock()

	stats := make(map[string]rego.Status, len(p.clientPools))
	for clientKey, clientPool := range p.clientPools {
		stats[clientKey] = clientPool.Status()
	}

	return stats
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
