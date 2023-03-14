package appreader

import (
	"backapper/app"
	"backapper/app/appholder"
	"encoding/json"
	"log"
	"os"
)

func Read(configFile string, logger *log.Logger) *appholder.AppHolder {

	bytes, err := os.ReadFile(configFile)
	if err != nil {
		logger.Println("Couldn't Read config:", configFile)
		return appholder.New()
	}

	var apps []*app.App
	errJson := json.Unmarshal(bytes, &apps)
	if errJson != nil {
		logger.Println("Couldn't parse config:", configFile)
		return appholder.New()
	}

	return appholder.New(apps...)
}
