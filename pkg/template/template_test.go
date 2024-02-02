// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package template

import "testing"

// go test -v -cover -count=1 -test.cpu=1 -run=^TestRenderPlain$
func TestRenderPlain(t *testing.T) {
	str := `{{ .begin }} {{ .xxx }} {{ .end }}`
	params := map[string]string{
		"begin": "Test",
		"xxx":   "Render",
		"end":   "Plain",
	}

	render, err := RenderPlain(str, params)
	if err != nil {
		t.Fatal(err)
	}

	if render != "Test Render Plain" {
		t.Fatalf("render %s is wrong", render)
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestRenderHTML$
func TestRenderHTML(t *testing.T) {
	str := `<p> {{ .begin }} {{ .xxx }} {{ .end }} </p>`
	params := map[string]string{
		"begin": "Test",
		"xxx":   "Render",
		"end":   "<HTML>",
	}

	render, err := RenderHTML(str, params)
	if err != nil {
		t.Fatal(err)
	}

	if render != "<p> Test Render &lt;HTML&gt; </p>" {
		t.Fatalf("render %s is wrong", render)
	}
}
