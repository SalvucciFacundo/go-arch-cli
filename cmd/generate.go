package cmd

import (
	"fmt"
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

		fmt.Printf("🛠️ Generando %s: %s...\n", compType, compName)
		
		// TODO: Implementar lógica de inyección de componentes
		fmt.Println("✅ Componente generado (simulado).")
	},
}
