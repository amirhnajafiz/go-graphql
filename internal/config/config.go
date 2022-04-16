package config

import (
	"encoding/json"
	"github.com/amirhnajafiz/go-graphql/internal/logger"
	"log"

	"github.com/amirhnajafiz/go-graphql/internal/gql"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

type (
	Config struct {
		Port   string
		Proxy  string
		GQL    gql.Config
		Logger logger.Config
	}
)

func Load() Config {
	var instance Config

	k := koanf.New(".")

	// load default configuration from default function
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	// load configuration from file
	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		log.Printf("error loading config.yml: %s", err)
	}

	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	indent, _ := json.MarshalIndent(instance, "", "\t")
	tmpl := `
	================ Loaded Configuration ================
	%s
	======================================================
	`
	log.Printf(tmpl, string(indent))

	return instance
}
