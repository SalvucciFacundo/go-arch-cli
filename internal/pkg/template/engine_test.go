package template

import (
	"bytes"
	"testing"
)

func TestEngine_Render(t *testing.T) {
	engine := NewEngine()

	data := struct {
		ProjectName string
		ModuleName  string
	}{
		ProjectName: "TestApp",
		ModuleName:  "github.com/test/app",
	}

	var buf bytes.Buffer
	err := engine.Render(&buf, "common/go.mod.tmpl", data)

	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	if buf.Len() == 0 {
		t.Error("expected rendered output, got empty buffer")
	}
}

func TestEngine_FuncMap(t *testing.T) {
	engine := NewEngine()
	funcMap := engine.getFuncMap()

	tests := []struct {
		name     string
		funcName string
		input    string
		want     string
	}{
		{
			name:     "lower function",
			funcName: "lower",
			input:    "HELLO",
			want:     "hello",
		},
		{
			name:     "upper function",
			funcName: "upper",
			input:    "world",
			want:     "WORLD",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, ok := funcMap[tt.funcName].(func(string) string)
			if !ok {
				t.Fatalf("function %s not found or has wrong signature", tt.funcName)
			}
			got := f(tt.input)
			if got != tt.want {
				t.Errorf("%s(%q) = %q; want %q", tt.funcName, tt.input, got, tt.want)
			}
		})
	}
}
