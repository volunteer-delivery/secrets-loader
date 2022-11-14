package loader

import (
    "encoding/json"
    "fmt"
    "os/exec"
    "secrets-loader/src/aws_json"
)

type Loader struct {
    config *Config
    secrets []Secret
}

func NewLoader(config *Config) *Loader {
    return &Loader{config: config}
}

func (it *Loader) Load() []Secret {
    it.nextSecrets("")
    return it.secrets
}

func (it *Loader) nextSecrets(paginationToken string)  {
    response := it.fetchSecrets(paginationToken)

    for _, parameter := range response.Parameters {
        secret := AwsParameterToSecret(it.config.Path, parameter)
        it.secrets = append(it.secrets, *secret)
    }

    if response.NextToken != "" {
        it.nextSecrets(response.NextToken)
    }
}

func (it *Loader) fetchSecrets(paginationToken string) *aws_json.ParametersByPath {
    args := []string{
        "ssm", "get-parameters-by-path",
        "--with-decryption",
        "--path", it.config.Path,
    }

    if it.config.Label != "" {
        filter := fmt.Sprintf("Key=Label,Values=%v", it.config.Label)
        args = append(args, "--parameter-filters", filter)
    }

    if it.config.Region != "" {
        args = append(args, "--region", it.config.Region)
    }

    if paginationToken != "" {
        args = append(args, "--starting-token", paginationToken)
    }

    command := exec.Command("aws", args...)
    stdout, _ := command.Output()

    response := aws_json.NewParametersByPath()
    _ = json.Unmarshal(stdout, &response)
    return response
}
