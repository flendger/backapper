package appreader

import (
	"backapper/app"
	"backapper/app/appholder"
	"encoding/json"
	"log"
	"os"
)

func read(configFile string) *appholder.AppHolder {

	bytes, err := os.ReadFile(configFile)
	if err != nil {
		log.Println("Couldn't read config:", configFile)
		return appholder.New()
	}

	var apps []*app.App
	errJson := json.Unmarshal(bytes, &apps)
	if errJson != nil {
		log.Println("Couldn't parse config:", configFile)
		return appholder.New()
	}

	return appholder.New(apps...)
}
