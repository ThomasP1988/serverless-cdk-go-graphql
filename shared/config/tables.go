package config

import "os"

type Table int
type SecondaryIndex int

type SecondaryIndexes = map[SecondaryIndex]string

type TableConfig struct {
	Name           string
	SecondaryIndex SecondaryIndexes
}
type TableRegistry map[Table]TableConfig

const (
	USER Table = iota
)

const (
	EmailIndex SecondaryIndex = iota
)

func GetTables(env *Stage) TableRegistry {

	SetEnv(env)

	if env == nil {
		envS := Stage(os.Getenv("env"))
		env = &envS
	}

	return TableRegistry{
		USER: TableConfig{
			Name: string(*env) + "-user",
			SecondaryIndex: SecondaryIndexes{
				EmailIndex: "email-index",
			},
		},
	}

}
