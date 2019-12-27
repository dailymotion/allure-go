package allure

import (
	"fmt"
	"github.com/dailymotion/allure-go/parameter"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func convertMapToParameters(parameters map[string]interface{}) []parameter.Parameter {
	result := make([]parameter.Parameter, 0)

	for k, v := range parameters {
		currentParameter := parameter.Parameter{
			Name: k,
		}

		switch v.(type) {
		case []byte:
			currentParameter.Value = string(v.([]byte))
		case uintptr:
			currentParameter.Value = strconv.Itoa(int(v.(uintptr)))
		case float32:
			currentParameter.Value = strconv.FormatFloat(float64(v.(float32)), 'f', -1, 64)
		case float64:
			currentParameter.Value = strconv.FormatFloat(v.(float64), 'f', -1, 64)
		case complex64:
			currentParameter.Value = fmt.Sprintf("%f i%f", real(v.(complex64)), imag(v.(complex64)))
		case complex128:
			currentParameter.Value = fmt.Sprintf("%f i%f", real(v.(complex128)), imag(v.(complex128)))
		case uint:
			currentParameter.Value = strconv.FormatUint(uint64(v.(uint)), 10)
		case uint8:
			currentParameter.Value = strconv.FormatUint(uint64(v.(uint8)), 10)
		case uint16:
			currentParameter.Value = strconv.FormatUint(uint64(v.(uint16)), 10)
		case uint32:
			currentParameter.Value = strconv.FormatUint(uint64(v.(uint32)), 10)
		case uint64:
			currentParameter.Value = strconv.FormatUint(v.(uint64), 10)
		case int:
			currentParameter.Value = strconv.FormatInt(int64(v.(int)), 10)
		case int8:
			currentParameter.Value = strconv.FormatInt(int64(v.(int8)), 10)
		case int16:
			currentParameter.Value = strconv.FormatInt(int64(v.(int16)), 10)
		case int32:
			currentParameter.Value = strconv.FormatInt(int64(v.(int32)), 10)
		case int64:
			currentParameter.Value = strconv.FormatInt(v.(int64), 10)
		case bool:
			currentParameter.Value = strconv.FormatBool(v.(bool))
		case string:
			currentParameter.Value = v.(string)
		default:
			currentParameter.Value = fmt.Sprintf("%+v", v)
		}

		result = append(result, currentParameter)
	}

	return result
}

func parseParameter(name string, value interface{}) parameter.Parameter {
	result := parameter.Parameter{
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
	resultsPathEnv := os.Getenv(ResultsPathEnvKey)
	if resultsPathEnv == "" {
		log.Printf("environment variable %s cannot be empty\n", ResultsPathEnvKey)
	}
	ResultsPath = fmt.Sprintf("%s/allure-results", resultsPathEnv)

	if _, err := os.Stat(ResultsPath); os.IsNotExist(err) {
		err = os.Mkdir(ResultsPath, 0777)
		if err != nil {
			log.Println(err, "Failed to create allure-results folder")
		}
	}
}

func copyEnvFileIfExists() {
	if envFilePath := os.Getenv(EnvFileKey); envFilePath != "" {
		envFilesStrings := strings.Split(envFilePath, "/")
		if ResultsPath != "" {
			if _, err := copy(envFilePath, ResultsPath+"/"+envFilesStrings[len(envFilesStrings)-1]); err != nil {
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
