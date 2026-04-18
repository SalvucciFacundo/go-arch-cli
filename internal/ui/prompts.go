package ui

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

type ProjectConfig struct {
	ProjectName  string
	ModuleName   string
	Architecture string
	DBDriver     string
	UseDocker    bool
}

// AskProjectConfig lanza el wizard interactivo para configurar el nuevo proyecto.
func AskProjectConfig(defaultName string) (*ProjectConfig, error) {
	var config ProjectConfig

	questions := []*survey.Question{
		{
			Name: "ProjectName",
			Prompt: &survey.Input{
				Message: "Nombre del Proyecto:",
				Default: defaultName,
			},
			Validate: survey.Required,
		},
		{
			Name: "ModuleName",
			Prompt: &survey.Input{
				Message: "Nombre del Módulo Go (ej: github.com/user/repo):",
				Default: fmt.Sprintf("github.com/user/%s", defaultName),
			},
			Validate: survey.Required,
		},
		{
			Name: "Architecture",
			Prompt: &survey.Select{
				Message: "Elegí la Arquitectura:",
				Options: []string{"Minimalist", "Standard", "Hexagonal"},
				Default: "Standard",
			},
		},
		{
			Name: "DBDriver",
			Prompt: &survey.Select{
				Message: "Driver de Base de Datos:",
				Options: []string{"PostgreSQL", "MySQL", "MongoDB", "None"},
				Default: "None",
			},
		},
		{
			Name: "UseDocker",
			Prompt: &survey.Confirm{
				Message: "¿Querés agregar soporte para Docker (Dockerfile y Docker Compose)?",
				Default: false,
			},
		},
	}

	err := survey.Ask(questions, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
