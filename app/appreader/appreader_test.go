package appreader

import (
	"backapper/app/appholder"
	"testing"
)

const configFile = "apps_test.json"

func TestReader(t *testing.T) {
	holder := read(configFile)

	checkApp("app1", "path1", "dir1", holder, t)
	checkApp("app2", "path2", "dir2", holder, t)
}

func checkApp(appName, appPath, appDir string, holder *appholder.AppHolder, t *testing.T) {
	app1, err := holder.GetApp(appName)
	if err != nil {
		t.Errorf("App not found: %s", appName)
	}

	if app1.FilePath != appPath {
		t.Errorf("File incorrect: %s wants %s", app1.FilePath, appPath)
	}

	if app1.ArcDir != appDir {
		t.Errorf("Dir incorrect: %s wants %s", app1.ArcDir, appDir)
	}
}
