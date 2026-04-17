package cmd

import (
	"fmt"
	"go-arch/internal/pkg/scaffold"
	"go-arch/internal/ui"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate [type] [name]",
	Short: "Generate a new component",
	Long:  `Generate components like service, repository, or handler based on the project layout.`,
	Args:  cobra.ExactArgs(2),
	Aliases: []string{"g"},
	Run: func(cmd *cobra.Command, args []string) {
		compType := args[0]
		compName := args[1]

		// Cargar configuración del proyecto local
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("❌ Error: No se encontró el archivo .go-arch.yaml. ¿Estás en la raíz del proyecto?")
			os.Exit(1)
		}

		// Mapear configuración de Viper a struct
		config := &ui.ProjectConfig{
			ProjectName:  viper.GetString("project_name"),
			ModuleName:   viper.GetString("module_name"),
			Architecture: viper.GetString("architecture"),
			DBDriver:     viper.GetString("db_driver"),
		}

		scaffolder := scaffold.NewScaffolder(config)
		if err := scaffolder.GenerateComponent(compType, compName); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("✅ Componente generado correctamente.")
	},
}
