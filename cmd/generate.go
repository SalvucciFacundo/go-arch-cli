package cmd

import (
	"fmt"
	"go-arch/internal/pkg/scaffold"
	"go-arch/internal/ui"

	"github.com/samber/oops"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:     "generate [type] [name]",
	Short:   "Generate a new component",
	Long:    `Generate components like service, repository, or handler based on the project layout.`,
	Args:    cobra.ExactArgs(2),
	Aliases: []string{"g"},
	RunE: func(cmd *cobra.Command, args []string) error {
		compType := args[0]
		name := args[1]

		// Validar configuración básica
		projectName := viper.GetString("project_name")
		if projectName == "" {
			return oops.
				Code("missing_config").
				Hint("Ejecuta 'go-arch setup' para inicializar el proyecto").
				Errorf("No se encontró el archivo .go-arch.yaml o está vacío")
		}

		// Mapear configuración de Viper a struct
		config := &ui.ProjectConfig{
			ProjectName:  projectName,
			ModuleName:   viper.GetString("module_name"),
			Architecture: viper.GetString("architecture"),
			DBDriver:     viper.GetString("db_driver"),
			UseDocker:    viper.GetBool("use_docker"),
		}

		ui.Info(fmt.Sprintf("Generando componente %s: %s...", compType, name))

		scaffolder := scaffold.NewScaffolder(config)
		var err error
		if compType == "crud" {
			err = scaffolder.GenerateCRUD(name)
		} else {
			err = scaffolder.GenerateComponent(compType, name)
		}

		if err != nil {
			return oops.
				Code("generation_failed").
				With("type", compType).
				With("name", name).
				Wrapf(err, "Falló la generación del componente")
		}

		ui.Success(fmt.Sprintf("Componente '%s' (%s) generado correctamente.", name, compType))
		return nil
	},
}
