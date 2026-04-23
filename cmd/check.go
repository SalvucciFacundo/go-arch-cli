package cmd

import (
	"fmt"
	"go-arch/internal/pkg/validator"
	"go-arch/internal/ui"

	"github.com/samber/oops"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1. Cargar configuración (ya leída en root, pero validamos campos críticos)
		projectName := viper.GetString("project_name")
		if projectName == "" {
			return oops.
				Code("missing_config").
				Hint("Ejecuta 'go-arch setup' para inicializar el proyecto").
				Errorf("No se encontró configuración válida. ¿Estás en la raíz de un proyecto go-arch?")
		}

		config := &ui.ProjectConfig{
			ProjectName:  projectName,
			ModuleName:   viper.GetString("module_name"),
			Architecture: viper.GetString("architecture"),
		}

		ui.Analyzing(config.Architecture)

		v := validator.NewValidator(config)
		violations, err := v.Validate()
		if err != nil {
			return oops.
				Code("validation_failed").
				Wrapf(err, "Error crítico durante la validación de arquitectura")
		}

		if len(violations) == 0 {
			ui.Success("¡Arquitectura impecable! No se detectaron infracciones.")
			return nil
		}

		ui.Warning(fmt.Sprintf("Se detectaron %d infracción(es):", len(violations)))
		fmt.Println()

		for _, v := range violations {
			statusSymbol := "❌"
			if v.Severity == "WARNING" {
				statusSymbol = "⚠️ "
			}
			fmt.Printf("%s [%s] %s\n   └─ %s\n", statusSymbol, v.Severity, v.File, v.Message)
		}

		return oops.
			Code("architecture_violations").
			Errorf("El proyecto no cumple con las reglas arquitectónicas")
	},
}
