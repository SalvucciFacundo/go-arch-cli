package cmd

import (
	"fmt"
	"go-arch/internal/pkg/scaffold"
	"go-arch/internal/ui"
	"os"

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
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Lanzar el asistente interactivo
		config, err := ui.RunWizard()
		if err != nil {
			fmt.Printf("❌ Error en el asistente: %v\n", err)
			os.Exit(1)
		}

		// 2. Ejecutar el scaffolding
		scaffolder := scaffold.NewScaffolder(config)
		if err := scaffolder.Execute(); err != nil {
			fmt.Printf("❌ Error creando el proyecto: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\n✨ ¡Proyecto '%s' creado con éxito!\n", config.ProjectName)
		fmt.Printf("👉 Ejecutá: cd %s y go-arch serve\n", config.ProjectName)
	},
}
