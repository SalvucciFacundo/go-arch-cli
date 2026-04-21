package ui

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

type ProjectConfig struct {
	ProjectName          string `mapstructure:"project_name"`
	ModuleName           string `mapstructure:"module_name"`
	Architecture         string `mapstructure:"architecture"`
	DBDriver             string `mapstructure:"db_driver"`
	UseDocker            bool   `mapstructure:"use_docker"`
	UseObservability     bool   `mapstructure:"use_observability"`
	ObservabilityBackend string `mapstructure:"observability_backend"`
}

func RunWizard() (*ProjectConfig, error) {
	fmt.Println("🚀 Bienvenido al asistente de Go-Arch")
	
	config := &ProjectConfig{}

	var mainQs = []*survey.Question{
		{
			Name: "ProjectName",
			Prompt: &survey.Input{
				Message: "Nombre del proyecto:",
				Default: "my-go-app",
			},
			Validate: survey.Required,
		},
		{
			Name: "ModuleName",
			Prompt: &survey.Input{
				Message: "Nombre del módulo (Go Module):",
				Default: "github.com/user/app",
			},
		},
		{
			Name: "Architecture",
			Prompt: &survey.Select{
				Message: "Selecciona la arquitectura:",
				Options: []string{"Minimalist", "Standard", "Hexagonal"},
				Default: "Standard",
			},
		},
		{
			Name: "DBDriver",
			Prompt: &survey.Select{
				Message: "Selecciona el driver de base de datos:",
				Options: []string{"PostgreSQL", "MySQL", "MongoDB", "None"},
				Default: "None",
			},
		},
		{
			Name: "UseDocker",
			Prompt: &survey.Confirm{
				Message: "¿Deseas incluir configuración de Docker?",
				Default: true,
			},
		},
		{
			Name: "UseObservability",
			Prompt: &survey.Confirm{
				Message: "¿Deseas habilitar Telemetría/Observabilidad (OpenTelemetry)?",
				Default: false,
			},
		},
	}

	err := survey.Ask(mainQs, config)
	if err != nil {
		return nil, err
	}

	if config.UseObservability {
		var obsQ = &survey.Select{
			Message: "Selecciona la herramienta de visualización:",
			Options: []string{"Console", "Jaeger", "Zipkin", "Prometheus", "SigNoz"},
			Default: "Console",
		}
		err = survey.AskOne(obsQ, &config.ObservabilityBackend)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}
