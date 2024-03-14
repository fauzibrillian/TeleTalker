package config

import (
	"os"

	"github.com/joho/godotenv"
)

type TeleTalker struct {
	TELEBOT string
}

func InitConfig() *TeleTalker {
	var response = new(TeleTalker)
	response = ReadData()
	return response
}

func ReadData() *TeleTalker {
	var data = new(TeleTalker)

	data = readEnv()

	if data == nil {
		err := godotenv.Load(".env")
		data = readEnv()
		if err != nil || data == nil {
			return nil
		}
	}
	return data
}

func readEnv() *TeleTalker {
	var data = new(TeleTalker)
	var permit = true

	if val, found := os.LookupEnv("TELEBOT"); found {
		data.TELEBOT = val
	} else {
		permit = false
	}

	if !permit {
		return nil
	}
	return data
}
