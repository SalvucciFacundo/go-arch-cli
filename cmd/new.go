package cmd

import (
	"fmt"
	"go-arch/internal/pkg/scaffold"
	"go-arch/internal/ui"

	"github.com/samber/oops"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create a new project",
	Long:  `The 'new' command initializes a new Go project with the specified name and architecture.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// 1. Lanzar el asistente interactivo
		ui.Info("Iniciando asistente de creación de proyecto...")
		config, err := ui.RunWizard()
		if err != nil {
			return oops.
				Code("wizard_failed").
				Wrapf(err, "Falló el asistente interactivo")
		}

		// 2. Ejecutar el scaffolding
		ui.Info(fmt.Sprintf("Creando proyecto '%s'...", config.ProjectName))
		scaffolder := scaffold.NewScaffolder(config)
		if err := scaffolder.Execute(); err != nil {
			return oops.
				Code("scaffold_failed").
				With("project_name", config.ProjectName).
				Wrapf(err, "Error durante el scaffolding del proyecto")
		}

		ui.Success(fmt.Sprintf("¡Proyecto '%s' creado con éxito!", config.ProjectName))
		fmt.Printf("👉 %s cd %s y go-arch serve\n", ui.InfoMsg("Ejecutá:"), config.ProjectName)
		return nil
	},
}
