package loader

import (
    "fmt"
    "secrets-loader/src/aws_json"
    "strings"
)

type Secret struct {
    Name  string
    Value string
}

func AwsParameterToSecret(path string, parameter aws_json.Parameter) *Secret {
    return &Secret{
        Name:  strings.TrimPrefix(stripSecretQuotes(parameter.Name), fmt.Sprintf("%v/", path)),
        Value: stripSecretQuotes(parameter.Value),
    }
}

func stripSecretQuotes(str string) string {
    str = strings.TrimPrefix(str, "\"")
    str = strings.TrimSuffix(str, "\"")
    return str
}
