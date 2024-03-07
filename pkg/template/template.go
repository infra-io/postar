// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package template

import (
	"bytes"
	htmltpl "html/template"
	texttpl "text/template"
)

var (
	RenderNone RenderFunc = func(str string, params map[string]string) (string, error) {
		return str, nil
	}

	RenderPlain RenderFunc = func(str string, params map[string]string) (string, error) {
		tpl, err := texttpl.New("plain_email").Parse(str)
		if err != nil {
			return "", err
		}

		render := bytes.NewBuffer(make([]byte, 0, len(str)))
		if err = tpl.Execute(render, params); err != nil {
			return "", err
		}

		return render.String(), nil
	}

	RenderHTML RenderFunc = func(str string, params map[string]string) (string, error) {
		tpl, err := htmltpl.New("html_email").Parse(str)
		if err != nil {
			return "", err
		}

		render := bytes.NewBuffer(make([]byte, 0, len(str)))
		if err = tpl.Execute(render, params); err != nil {
			return "", err
		}

		return render.String(), nil
	}
)

type RenderFunc func(str string, params map[string]string) (string, error)
