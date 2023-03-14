package appservice

import (
	"backapper/app"
	"backapper/app/appholder"
	"backapper/app/appservice/fnameresolver"
	"io"
	"log"
	"os"
	"os/exec"
)

type AppService struct {
	holder *appholder.AppHolder
}

func (s *AppService) BackUp(appName string) (string, error) {
	curApp, err := s.getApp(appName)
	if err != nil {
		return "", err
	}

	source, err := os.Open(curApp.FilePath)
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

	arcDir := curApp.ArcDir
	errDir := os.MkdirAll(arcDir, os.ModePerm)
	if errDir != nil {
		return "", errDir
	}

	distPath := fnameresolver.Resolve(curApp, fileInfo.Name(), fileInfo.ModTime())
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

func (s *AppService) Deploy(appName string, newFile io.Reader) (string, error) {
	curApp, err := s.getApp(appName)
	if err != nil {
		return "", err
	}

	filePath := curApp.FilePath

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Couldn't close file during deploying:", filePath)
		}
	}(file)

	_, errCopy := io.Copy(file, newFile)
	if errCopy != nil {
		return "", err
	}

	return filePath, nil
}

func (s *AppService) Restart(appName string) error {
	curApp, err := s.getApp(appName)
	if err != nil {
		return err
	}

	errRun := exec.Command("sh", "-c", curApp.Restart).Run()
	if errRun != nil {
		return errRun
	}

	return nil
}

func New(holder *appholder.AppHolder) *AppService {
	return &AppService{holder: holder}
}

func (s *AppService) getApp(appName string) (*app.App, error) {
	app, err := s.holder.GetApp(appName)
	if err != nil {
		return nil, err
	}

	return app, nil
}
