package config

import "os"

type Stage string

var (
	DEV     Stage = "dev"
	STAGING Stage = "staging"
	PROD    Stage = "prod"
)

type Config struct {
	LocalProfile string
	Region       string
	Tables       TableRegistry
	APIKey       string
}

func SetEnv(env *Stage) *Stage {
	if env == nil {
		envS := Stage(os.Getenv("stage"))
		if envS != "" {
			return &envS

		} else {
			return &DEV

		}
	}
	return env
}

var Conf *Config

func GetConfig(env *Stage) *Config {

	if Conf != nil {
		return Conf
	}
	env = SetEnv(env)

	Conf = &Config{
		// LocalProfile: "yourProfile",
		Region: "us-east-1",
		Tables: GetTables(env),
		APIKey: "exampleForDevEnv",
	}

	if *env == STAGING {
		updateConfigForStage(Conf)
	} else if *env == PROD {
		updateConfigForProd(Conf)
	}

	return Conf
}
