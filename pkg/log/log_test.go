package log

import (
	"github.com/avinoplan/postar/configs"
	"testing"
)

// go test -v -cover -run=^TestInitialize$
func TestInitialize(t *testing.T) {
	if globalLogger != nil {
		t.Errorf("globalLogger %+v != nil", globalLogger)
	}

	c := configs.NewDefaultConfig()
	err := Initialize(c)
	if err != nil {
		t.Error(err)
	}

	if globalLogger == nil {
		t.Error("globalLogger == nil")
	}
}
