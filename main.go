package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/kr/pretty"
	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu-community/sensu-plugin-sdk/templates"
	"github.com/sensu/sensu-go/types"
)

// Config represents the handler plugin config.
type Config struct {
	sensu.PluginConfig
	Template  string
	DumpNames bool
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
		&sensu.PluginConfigOption{
			Env:       "",
			Argument:  "dump-names",
			Shorthand: "d",
			Default:   false,
			Usage:     "Dump the event object names that can be used in a Golang template",
			Value:     &config.DumpNames,
		},
	}
)

func main() {
	handler := sensu.NewGoHandler(&config.PluginConfig, options, checkArgs, executeHandler)
	handler.Execute()
}

func checkArgs(_ *types.Event) error {
	if len(config.Template) == 0 && !config.DumpNames {
		return fmt.Errorf("--template variable or --dump-names is required")
	}
	return nil
}

func executeHandler(event *types.Event) error {
	if config.DumpNames {
		prettyPrint(event.Entity)
		if event.HasCheck() {
			prettyPrint(event.Check)
		}
		if event.HasMetrics() {
			prettyPrint(event.Metrics)
		}
		return nil
	}

	log.Println("executing command with --template", config.Template)
	description, err := templates.EvalTemplate("description", config.Template, event)
	if err != nil {
		log.Println("Error processing template:\n", err)
	} else {
		log.Println("Template String Output:", description)
	}
	return nil
}

func prettyPrint(obj interface{}) {
	// The pretty.Formatter gets us most of the way there, these
	// regex's are to clean it up a little

	// Strip out the unused names XXX_*
	m1 := regexp.MustCompile(`[[:blank:]]+XXX_.*:.*\n`)
	// Strip out the v2 type references for readability
	m2 := regexp.MustCompile(`&*v2`)
	// Insert a period before names to help show the structure (needed?)
	m3 := regexp.MustCompile(`([[:blank:]]+)([A-Z])`)

	o := fmt.Sprintf("%# v", pretty.Formatter(obj))
	s1 := m1.ReplaceAllString(o, "")
	s2 := m2.ReplaceAllString(s1, "")
	fmt.Println(m3.ReplaceAllString(s2, "$1.$2"))
}
