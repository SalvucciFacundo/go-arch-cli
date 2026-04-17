package cmd

import (
	"fmt"
	"os"
	"os/exec"

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
	Run: func(cmd *cobra.Command, args []string) {
		layout := viper.GetString("architecture")
		mainPath := "cmd/api/main.go"

		if layout == "Minimalist" {
			mainPath = "main.go"
		}

		fmt.Printf("🚀 Iniciando servidor para el proyecto (Layout: %s)...\n", layout)

		// Verificar si Air está instalado
		_, err := exec.LookPath("air")
		if err == nil {
			fmt.Println("🔥 Usando Air para hot-reload...")
			runWithAir()
			return
		}

		fmt.Println("⚠️  Air no detectado. Usando 'go run' (sin hot-reload)...")
		runWithGo(mainPath)
	},
}

func runWithAir() {
	cmd := exec.Command("air")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func runWithGo(path string) {
	cmd := exec.Command("go", "run", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
