package cmd

import (
	"fmt"
	"go-arch/internal/pkg/validator"
	"go-arch/internal/ui"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check project architecture health",
	Long:  `Validates the project structure and dependency rules (imports) based on the configured architecture.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Cargar configuración
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("❌ Error: No se encontró .go-arch.yaml. ¿Estás en la raíz de un proyecto go-arch?")
			os.Exit(1)
		}

		config := &ui.ProjectConfig{
			ProjectName:  viper.GetString("project_name"),
			ModuleName:   viper.GetString("module_name"),
			Architecture: viper.GetString("architecture"),
		}

		fmt.Printf("🔍 Analizando arquitectura **%s**...\n\n", config.Architecture)

		v := validator.NewValidator(config)
		violations, err := v.Validate()
		if err != nil {
			fmt.Printf("❌ Error crítico durante la validación: %v\n", err)
			os.Exit(1)
		}

		if len(violations) == 0 {
			fmt.Println("✅ ¡Arquitectura impecable! No se detectaron infracciones.")
			return
		}

		fmt.Printf("⚠️ Se detectaron %d infracción(es):\n\n", len(violations))
		for _, v := range violations {
			statusSymbol := "❌"
			if v.Severity == "WARNING" {
				statusSymbol = "⚠️ "
			}
			fmt.Printf("%s [%s] %s\n   └─ %s\n", statusSymbol, v.Severity, v.File, v.Message)
		}

		fmt.Println("\n❌ El proyecto no cumple con las reglas arquitectónicas.")
		os.Exit(1)
	},
}
