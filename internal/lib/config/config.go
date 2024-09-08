package config

import (
	"launchpad-go-rest/pkg/types/config"
	"os"

	"github.com/golobby/dotenv"
)

var Configs config.Configs

func Init() {
	file, err := os.Open(".env")
	if err != nil {
		panic(err)
	}

	if err := dotenv.NewDecoder(file).Decode(&Configs); err != nil {
		panic(err)
	}
}
