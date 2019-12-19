package allure

import (
	"fmt"
	"strconv"
)

type Parameter struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

func convertMapToParameters(parameters map[string]interface{}) []Parameter {
	result := make([]Parameter, 0)

	for k, v := range parameters {
		currentParameter := Parameter{
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

func parseParameter(name string, value interface{}) Parameter {
	parameter := Parameter{
		Name: name,
	}

	switch value.(type) {
	case []byte:
		parameter.Value = string(value.([]byte))
	case uintptr:
		parameter.Value = strconv.Itoa(int(value.(uintptr)))
	case float32:
		parameter.Value = strconv.FormatFloat(float64(value.(float32)), 'f', -1, 64)
	case float64:
		parameter.Value = strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case complex64:
		parameter.Value = fmt.Sprintf("%f i%f", real(value.(complex64)), imag(value.(complex64)))
	case complex128:
		parameter.Value = fmt.Sprintf("%f i%f", real(value.(complex128)), imag(value.(complex128)))
	case uint:
		parameter.Value = strconv.FormatUint(uint64(value.(uint)), 10)
	case uint8:
		parameter.Value = strconv.FormatUint(uint64(value.(uint8)), 10)
	case uint16:
		parameter.Value = strconv.FormatUint(uint64(value.(uint16)), 10)
	case uint32:
		parameter.Value = strconv.FormatUint(uint64(value.(uint32)), 10)
	case uint64:
		parameter.Value = strconv.FormatUint(value.(uint64), 10)
	case int:
		parameter.Value = strconv.FormatInt(int64(value.(int)), 10)
	case int8:
		parameter.Value = strconv.FormatInt(int64(value.(int8)), 10)
	case int16:
		parameter.Value = strconv.FormatInt(int64(value.(int16)), 10)
	case int32:
		parameter.Value = strconv.FormatInt(int64(value.(int32)), 10)
	case int64:
		parameter.Value = strconv.FormatInt(value.(int64), 10)
	case bool:
		parameter.Value = strconv.FormatBool(value.(bool))
	case string:
		parameter.Value = value.(string)
	default:
		parameter.Value = fmt.Sprintf("%+value", value)
	}

	return parameter
}
