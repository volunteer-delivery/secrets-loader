package loader

import (
	"context"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	ssmTypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"secrets-loader/src/logger"
)

type Loader struct {
	ssm     *ssm.Client
	logger  *logger.Logger
	config  *Config
	secrets []Secret
}

func NewLoader(logger *logger.Logger, config *Config) *Loader {
	return &Loader{
		ssm:    initSsmClient(logger, config.Region),
		config: config,
		logger: logger,
	}
}

func initSsmClient(logger *logger.Logger, region string) *ssm.Client {
	config, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithRegion(region))

	if err != nil {
		logger.Error("Unable to load SDK config, %v", err)
	}

	return ssm.NewFromConfig(config)
}

func (it *Loader) Load() []Secret {
	it.nextSecrets("")
	return it.secrets
}

func (it *Loader) nextSecrets(paginationToken string) {
	it.logger.Debug("Fetching secrets")
	response := it.fetchSecrets(paginationToken)
	it.logger.Debug("Fetched secrets")

	for _, parameter := range response.Parameters {
		secret := AwsParameterToSecret(it.config.Path, parameter)
		it.secrets = append(it.secrets, *secret)
	}

	if response.NextToken != nil {
		it.nextSecrets(*response.NextToken)
	}
}

func (it *Loader) fetchSecrets(paginationToken string) *ssm.GetParametersByPathOutput {
	withDescryption := true
	filterLabelKey := "Label"

	input := ssm.GetParametersByPathInput{
		WithDecryption: &withDescryption,
		Path:           &it.config.Path,
		NextToken:      &paginationToken,
		ParameterFilters: []ssmTypes.ParameterStringFilter{
			{
				Key:    &filterLabelKey,
				Values: []string{it.config.Label},
			},
		},
	}

	if paginationToken != "" {
		input.NextToken = &paginationToken
	}

    it.logger.Debug("Fetch secrets with params:")
    it.logger.Debug("  Path: %v", it.config.Path)
    it.logger.Debug("  NextToken: %v", paginationToken)
    it.logger.Debug("  ParameterFilter: Key=%v,Values=%v", filterLabelKey, it.config.Label)

	output, err := it.ssm.GetParametersByPath(context.TODO(), &input)

	if err != nil {
		it.logger.Error("Failed to load parameters: %v", err)
	}

	return output
}
