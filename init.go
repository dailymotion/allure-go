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
	ctxMgr     *gls.ContextManager
	wsd        string
	resultPath string
)

const (
	resultsPathEnvKey = "ALLURE_RESULTS_PATH"
	wsPathEnvKey      = "ALLURE_WORKSPACE_PATH"
	envFileKey        = "ALLURE_ENVIRONMENT_FILE_PATH"
	nodeKey           = "current_step_container"
)

func init() {
	ctxMgr = gls.NewContextManager()
	if envFilePath := os.Getenv(envFileKey); envFilePath != "" {
		envFilesStrings := strings.Split(envFilePath, "/")
		resultsPath := os.Getenv(resultsPathEnvKey)
		if _, err := copy(envFilePath, resultsPath+"/allure-results/"+envFilesStrings[len(envFilesStrings)-1]); err != nil {
			log.Fatalf("Could not copy the environment file, %f", err)
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
