package configs

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Envs struct {
	AppEnv string `env:"APP_ENV"`
}

func LoadEnv() Envs {
	envExtension := parseArgs()
	dotEnv(envExtension)

	config := Envs{}
	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return config
}
func dotEnv(path *string) {
	if path != nil {
		filePath := *path
		godotenv.Load(".env." + filePath)
		return
	}
	godotenv.Load()
}

func parseArgs() *string {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		envFileExtensions := args[0]

		if envFileExtensions != "" {
			return &envFileExtensions
		}
	}
	return nil
}
