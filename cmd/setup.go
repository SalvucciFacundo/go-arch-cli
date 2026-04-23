package cmd

import (
	"fmt"
	"go-arch/internal/ui"
	"runtime"

	"github.com/samber/oops"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(setupCmd)
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup Go environment",
	Long:  `The 'setup' command detects your OS and installs Go and necessary tools like 'air'.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ui.Info(fmt.Sprintf("Detectando entorno para %s/%s...", runtime.GOOS, runtime.GOARCH))

		switch runtime.GOOS {
		case "linux":
			setupLinux()
		case "windows":
			setupWindows()
		default:
			return oops.
				Code("os_not_supported").
				With("os", runtime.GOOS).
				Errorf("Sistema operativo no soportado automáticamente aún")
		}

		ui.Success("Proceso de setup finalizado. Revisa las instrucciones arriba para completar la instalación.")
		return nil
	},
}

func setupLinux() {
	ui.Info("🐧 Entorno Linux detectado.")
	fmt.Println("1. Descargando instalador oficial de go.dev...")
	// TODO: Implementar descarga real con net/http
	fmt.Println("2. Para instalar, ejecutá: sudo tar -C /usr/local -xzf go1.24.linux-amd64.tar.gz")
	fmt.Println("3. Instalando Air para hot-reload...")
	fmt.Printf("👉 %s go install github.com/air-verse/air@latest\n", ui.InfoMsg("Ejecutá:"))
}

func setupWindows() {
	ui.Info("🪟 Entorno Windows detectado.")
	fmt.Println("1. Descargando MSI oficial de go.dev...")
	// TODO: Implementar descarga real
	fmt.Println("2. Ejecutando instalador...")
	fmt.Println("3. Instalando Air para hot-reload...")
	fmt.Printf("👉 %s go install github.com/air-verse/air@latest\n", ui.InfoMsg("Ejecutá:"))
}
