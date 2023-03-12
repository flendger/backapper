package appservice

import (
	"backapper/app/appholder"
	"backapper/app/appservice/fnameresolver"
	"errors"
	"io"
	"log"
	"os"
)

type AppService struct {
	holder *appholder.AppHolder
}

func (s *AppService) BackUp(appName string) (string, error) {
	app, err := s.holder.GetApp(appName)
	if err != nil {
		return "", err
	}

	source, err := os.Open(app.FilePath)
	if err != nil {
		return "", err
	}
	defer func(source *os.File) {
		err := source.Close()
		if err != nil {
			log.Println(err)
		}
	}(source)

	fileInfo, err := source.Stat()
	if err != nil {
		return "", err
	}

	arcDir := app.ArcDir
	if _, errDir := os.Stat(arcDir); errors.Is(errDir, os.ErrNotExist) {
		err := os.Mkdir(arcDir, os.ModeDir)
		if err != nil {
			return "", err
		}
	}

	distPath := fnameresolver.Resolve(app, fileInfo.Name(), fileInfo.ModTime())
	distFile, err := os.Create(distPath)
	if err != nil {
		return "", err
	}
	defer func(distFile *os.File) {
		err := distFile.Close()
		if err != nil {
			log.Println(err)
		}
	}(distFile)

	_, errCopy := io.Copy(distFile, source)
	if errCopy != nil {
		return "", errCopy
	}

	errTime := os.Chtimes(distPath, fileInfo.ModTime(), fileInfo.ModTime())
	if errTime != nil {
		return "", errTime
	}

	return distPath, nil
}

func New(holder *appholder.AppHolder) *AppService {
	return &AppService{holder: holder}
}
