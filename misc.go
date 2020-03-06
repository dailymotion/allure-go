package allure

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func parseParameter(name string, value interface{}) parameter {
	result := parameter{
		Name: name,
	}

	switch value.(type) {
	case []byte:
		result.Value = string(value.([]byte))
	case uintptr:
		result.Value = strconv.Itoa(int(value.(uintptr)))
	case float32:
		result.Value = strconv.FormatFloat(float64(value.(float32)), 'f', -1, 64)
	case float64:
		result.Value = strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case complex64:
		result.Value = fmt.Sprintf("%f i%f", real(value.(complex64)), imag(value.(complex64)))
	case complex128:
		result.Value = fmt.Sprintf("%f i%f", real(value.(complex128)), imag(value.(complex128)))
	case uint:
		result.Value = strconv.FormatUint(uint64(value.(uint)), 10)
	case uint8:
		result.Value = strconv.FormatUint(uint64(value.(uint8)), 10)
	case uint16:
		result.Value = strconv.FormatUint(uint64(value.(uint16)), 10)
	case uint32:
		result.Value = strconv.FormatUint(uint64(value.(uint32)), 10)
	case uint64:
		result.Value = strconv.FormatUint(value.(uint64), 10)
	case int:
		result.Value = strconv.FormatInt(int64(value.(int)), 10)
	case int8:
		result.Value = strconv.FormatInt(int64(value.(int8)), 10)
	case int16:
		result.Value = strconv.FormatInt(int64(value.(int16)), 10)
	case int32:
		result.Value = strconv.FormatInt(int64(value.(int32)), 10)
	case int64:
		result.Value = strconv.FormatInt(value.(int64), 10)
	case bool:
		result.Value = strconv.FormatBool(value.(bool))
	case string:
		result.Value = value.(string)
	default:
		result.Value = fmt.Sprintf("%+value", value)
	}

	return result
}

func getTimestampMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func createFolderIfNotExists() {
	resultsPathEnv := os.Getenv(resultsPathEnvKey)
	if resultsPathEnv == "" {
		log.Printf("environment variable %s cannot be empty\n", resultsPathEnvKey)
	}
	resultsPath = fmt.Sprintf("%s/allure-results", resultsPathEnv)

	if _, err := os.Stat(resultsPath); os.IsNotExist(err) {
		err = os.Mkdir(resultsPath, 0777)
		if err != nil {
			log.Println(err, "Failed to create allure-results folder")
		}
	}
}

func copyEnvFileIfExists() {
	if envFilePath := os.Getenv(envFileKey); envFilePath != "" {
		envFilesStrings := strings.Split(envFilePath, "/")
		if resultsPath != "" {
			if _, err := copy(envFilePath, resultsPath+"/"+envFilesStrings[len(envFilesStrings)-1]); err != nil {
				log.Println("Could not copy the environment file", err)
			}
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
			log.Printf("Could not close the stream for the environment file, %f\n", err)
		}
	}()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err = destination.Close(); err != nil {
			log.Printf("Could not close the stream for the destination of the environment file, %f\n", err)
		}
	}()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func ensureFolderCreated() {
	createFolderOnce.Do(createFolderIfNotExists)
}
