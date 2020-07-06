package main

import (
	"encoding/json"
	"go-cleanup/usecases"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	settingPath string
)

func main() {
	// cobra command
	rootCmd := &cobra.Command{
		Use:   "go-cleanup",
		Short: "go-cleanup is an utility to clean up files after certain time and unused git branches",
		Run:   cmdHandler,
	}

	rootCmd.Flags().StringVarP(&settingPath, "file", "f", "", "option file path")
	err := rootCmd.MarkFlagRequired("file")
	if err != nil {
		panic(err)
	}

	err = rootCmd.Execute()
	if err != nil && err != filepath.SkipDir {
		panic(err)
	}
}

func cmdHandler(_ *cobra.Command, _ []string) {
	if settingPath == "" {
		panic("setting path cannot be empty")
	}

	settingPath, err := filepath.Abs(settingPath)
	if err != nil {
		panic(err)
	}

	fileBytes, err := ioutil.ReadFile(settingPath)
	if err != nil {
		panic(err)
	}

	var settings usecases.Settings
	err = json.Unmarshal(fileBytes, &settings)
	if err != nil {
		panic(err)
	}

	err = settings.Validate()
	if err != nil {
		panic(err)
	}

	for _, dir := range settings.Directories {
		err := dir.Clean()
		if err != nil {
			panic(err)
		}
	}

	for _, repo := range settings.Repositories {
		err := repo.Clean()
		if err != nil {
			panic(err)
		}
	}
}
