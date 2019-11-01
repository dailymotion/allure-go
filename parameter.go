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
		case int:
			currentParameter.Value = strconv.Itoa(v.(int))
		case string:
			currentParameter.Value = v.(string)
		default:
			currentParameter.Value = fmt.Sprintf("%+v", v)
		}

		result = append(result, currentParameter)
	}

	return result
}
