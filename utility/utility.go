package utility

import (
	"encoding/json"
	"github.com/richardsang2008/accountptcmanager/model"
	"os"
)

var (
	MLog   Log
	MCache Cache
)

type Utility struct {
}

func LoadConfiguration(file string) model.Config {
	var config model.Config
	bar, found := MCache.Get("appconfig")
	if found {
		config = bar.(model.Config)
	} else {
		configFile, err := os.Open(file)
		defer configFile.Close()
		if err != nil {
			panic("Can not LoadConfiguration: " + err.Error())
		}
		jsonParser := json.NewDecoder(configFile)
		jsonParser.Decode(&config)
	}
	return config
}
