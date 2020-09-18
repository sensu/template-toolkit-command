package main

import (
	"fmt"
	"log"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu-community/sensu-plugin-sdk/templates"
	"github.com/sensu/sensu-go/types"
)

// Config represents the handler plugin config.
type Config struct {
	sensu.PluginConfig
	Template string
}

var (
	config = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "template-toolkit-command",
			Short:    "Sensuctl command plugin for validating and testing Sensu plugin templates",
			Keyspace: "sensu.io/plugins/template-toolkit-command/config",
		},
	}

	options = []*sensu.PluginConfigOption{
		&sensu.PluginConfigOption{
			Env:       "TEMPLATE",
			Argument:  "template",
			Shorthand: "t",
			Default:   "",
			Usage:     "A template string, in Golang text/template format",
			Value:     &config.Template,
		},
	}
)

func main() {
	handler := sensu.NewGoHandler(&config.PluginConfig, options, checkArgs, executeHandler)
	handler.Execute()
}

func checkArgs(_ *types.Event) error {
	if len(config.Template) == 0 {
		return fmt.Errorf("--template variable is required")
	}
	return nil
}

func executeHandler(event *types.Event) error {
	log.Println("executing command with --template", config.Template)
	description, err := templates.EvalTemplate("description", config.Template, event)
	if err != nil {
		log.Println("Error processing template:\n", err)
	} else {
		log.Println("Template String Output:", description)
	}
	return nil
}
