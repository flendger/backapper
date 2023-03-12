package fnameresolver

import (
	"backapper/app"
	"strings"
	"time"
)

func Resolve(curApp *app.App, fName string, fTime time.Time) string {
	fPrefix, fSuffix := splitFName(fName)

	return curApp.ArcDir + "/" + fPrefix + "_" + fTime.Format("2006-01-02 15_04_05") + fSuffix
}

func splitFName(fName string) (string, string) {
	index := strings.LastIndex(fName, ".")
	if index == -1 {
		return fName, ""
	}

	return fName[:index], fName[index:]
}
