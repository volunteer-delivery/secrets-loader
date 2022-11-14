package loader

import (
    "encoding/json"
    "fmt"
    "os/exec"
    "secrets-loader/src/aws_json"
    "secrets-loader/src/logger"
)

type Loader struct {
    logger *logger.Logger
    config *Config
    secrets []Secret
}

func NewLoader(logger *logger.Logger, config *Config) *Loader {
    return &Loader{
        config: config,
        logger: logger,
    }
}

func (it *Loader) Load() []Secret {
    it.nextSecrets("")
    return it.secrets
}

func (it *Loader) nextSecrets(paginationToken string)  {
    it.logger.Debug("Fetching secrets")
    response := it.fetchSecrets(paginationToken)
    it.logger.Debug("Fetched secrets")

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

    it.logger.Debug("Select in path %v", it.config.Path)

    if it.config.Label != "" {
        filter := fmt.Sprintf("Key=Label,Values=%v", it.config.Label)
        args = append(args, "--parameter-filters", filter)
        it.logger.Debug("Filter by %v", filter)
    }

    if it.config.Region != "" {
        args = append(args, "--region", it.config.Region)
        it.logger.Debug("Apply AWS region %v", it.config.Region)
    }

    if paginationToken != "" {
        args = append(args, "--starting-token", paginationToken)
        it.logger.Debug("Apply pagination token")
    }

    command := exec.Command("aws", args...)
    stdout, stderr := command.Output()

    if stderr != nil {
        it.logger.Error("Shell failed %v", stderr.Error())
    }

    response := aws_json.NewParametersByPath()
    jsonError := json.Unmarshal(stdout, &response)

    if jsonError != nil {
        it.logger.Error("Json parsing failed: %v", jsonError.Error())
    }

    return response
}
