package cmd

import (
	"fmt"
	"go-arch/internal/ui"
	"os"
	"os/exec"

	"github.com/samber/oops"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the project with hot-reload",
	Long:  `Run the project using 'air' for hot-reload if available, otherwise fallback to 'go run'.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		layout := viper.GetString("architecture")
		if layout == "" {
			return oops.
				Code("missing_config").
				Hint("Asegúrate de estar en la raíz de un proyecto go-arch").
				Errorf("No se encontró configuración válida de arquitectura")
		}

		mainPath := "cmd/api/main.go"
		if layout == "Minimalist" {
			mainPath = "main.go"
		}

		ui.Info(fmt.Sprintf("Iniciando servidor para el proyecto (Layout: %s)...", layout))

		// Verificar si Air está instalado
		_, err := exec.LookPath("air")
		if err == nil {
			ui.Info("🔥 Usando Air para hot-reload...")
			if err := runWithAir(); err != nil {
				return oops.
					Code("server_error").
					Wrapf(err, "Falló la ejecución de 'air'")
			}
			return nil
		}

		ui.Warning("Air no detectado en el PATH. Usando 'go run' (sin hot-reload)...")
		if err := runWithGo(mainPath); err != nil {
			return oops.
				Code("server_error").
				With("path", mainPath).
				Wrapf(err, "Falló la ejecución de 'go run'")
		}

		return nil
	},
}

func runWithAir() error {
	cmd := exec.Command("air")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runWithGo(path string) error {
	cmd := exec.Command("go", "run", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
