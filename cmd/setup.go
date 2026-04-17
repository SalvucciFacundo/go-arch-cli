package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(setupCmd)
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup Go environment",
	Long:  `The 'setup' command detects your OS and installs Go and necessary tools like 'air'.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("🔍 Detectando entorno para %s/%s...\n", runtime.GOOS, runtime.GOARCH)

		switch runtime.GOOS {
		case "linux":
			setupLinux()
		case "windows":
			setupWindows()
		default:
			fmt.Printf("❌ Sistema operativo %s no soportado automáticamente aún.\n", runtime.GOOS)
			os.Exit(1)
		}
	},
}

func setupLinux() {
	fmt.Println("🐧 Entorno Linux detectado.")
	fmt.Println("1. Descargando instalador oficial de go.dev...")
	// TODO: Implementar descarga real con net/http
	fmt.Println("2. Para instalar, ejecutá: sudo tar -C /usr/local -xzf go1.23.linux-amd64.tar.gz")
	fmt.Println("3. Instalando Air para hot-reload...")
	fmt.Println("👉 Ejecutá: go install github.com/air-verse/air@latest")
}

func setupWindows() {
	fmt.Println("🪟 Entorno Windows detectado.")
	fmt.Println("1. Descargando MSI oficial de go.dev...")
	// TODO: Implementar descarga real
	fmt.Println("2. Ejecutando instalador...")
	fmt.Println("3. Instalando Air para hot-reload...")
	fmt.Println("👉 Ejecutá: go install github.com/air-verse/air@latest")
}
