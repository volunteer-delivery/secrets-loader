package loader

import (
	"fmt"
    ssmTypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"strings"
)

type Secret struct {
	Name  string
	Value string
}

func AwsParameterToSecret(path string, parameter ssmTypes.Parameter) *Secret {
	return &Secret{
		Name:  strings.TrimPrefix(*parameter.Name, fmt.Sprintf("%v/", path)),
		Value: *parameter.Value,
	}
}
