package template

import (
	"bytes"
	"strings"
	"testing"
)

func TestEngine_Render(t *testing.T) {
	engine := NewEngine()

	t.Run("Render Minimalist Main", func(t *testing.T) {
		var buf bytes.Buffer
		data := struct {
			ProjectName string
		}{
			ProjectName: "TestProject",
		}

		err := engine.Render(&buf, "minimalist/main.tmpl", data)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		got := buf.String()
		want := "Minimalist Go-Arch project"
		if !strings.Contains(got, want) {
			t.Errorf("got %q, want it to contain %q", got, want)
		}
	})

	t.Run("Render with function 'now'", func(t *testing.T) {
		var buf bytes.Buffer
		data := struct {
			ProjectName  string
			ModuleName   string
			Architecture string
			DBDriver     string
		}{
			ProjectName:  "TestProject",
			ModuleName:   "github.com/test/repo",
			Architecture: "Standard",
			DBDriver:     "PostgreSQL",
		}

		// Usamos el config.tmpl para probar la función 'now'
		err := engine.Render(&buf, "common/config.tmpl", data)
		if err != nil {
			t.Fatalf("failed to render: %v", err)
		}

		got := buf.String()
		if !strings.Contains(got, "generated_at:") {
			t.Errorf("got %q, want it to contain 'generated_at:'", got)
		}
	})
}
