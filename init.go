package allure

import (
	"fmt"
	"github.com/jtolds/gls"
	"io"
	"log"
	"os"
	"strings"
)

var (
	ctxMgr *gls.ContextManager
	wsd    string
)

const (
	resultsPathEnvKey = "ALLURE_RESULTS_PATH"
	wsPathEnvKey      = "ALLURE_WORKSPACE_PATH"
	envFileKey        = "ALLURE_ENVIRONMENT_FILE_PATH"
	nodeKey           = "current_step_container"
)

var resultsPath string

func init() {
	ctxMgr = gls.NewContextManager()
	if envFilePath := os.Getenv(envFileKey); envFilePath != "" {
		envFilesStrings := strings.Split(envFilePath, "/")

		resultsPathEnv := os.Getenv(resultsPathEnvKey)
		if resultsPathEnv == "" {
			log.Fatalf("%s environment variable cannot be empty", resultsPathEnvKey)
		}
		resultsPath = fmt.Sprintf("%s/allure-results", resultsPathEnv)

		if _, err := os.Stat(resultsPath); os.IsNotExist(err) {
			err = os.Mkdir(resultsPath, 0777)
			if err != nil {
				log.Fatal("Failed to create allure-results folder", err)
			}
		}

		if _, err := copy(envFilePath, resultsPath+"/"+envFilesStrings[len(envFilesStrings)-1]); err != nil {
			log.Fatal("Could not copy the environment file", err)
		}
	}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err = source.Close(); err != nil {
			log.Fatalf("Could not close the stream for the environment file, %f", err)
		}
	}()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err = destination.Close(); err != nil {
			log.Fatalf("Could not close the stream for the destination of the environment file, %f", err)
		}
	}()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
