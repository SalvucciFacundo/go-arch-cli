package template

import (
	"bytes"
	"testing"
)

func TestEngine_Render(t *testing.T) {
	engine := NewEngine()

	data := struct {
		ProjectName string
	}{
		ProjectName: "TestApp",
	}

	var buf bytes.Buffer
	// Usamos uno de los templates existentes (ej: common/go.mod.tmpl)
	err := engine.Render(&buf, "common/go.mod.tmpl", data)

	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	if buf.Len() == 0 {
		t.Error("expected rendered output, got empty buffer")
	}
}
