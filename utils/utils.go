package utils

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// HACK: this func serves to switch the env used
func GetAwsConfig() (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		// config.WithClientLogMode(aws.LogRequestWithBody|aws.LogResponseWithBody),
		config.WithRegion("eu-west-1"),
		config.WithHTTPClient(&http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					PartitionID:       "aws",
					URL:               "http://127.0.0.1:4566",
					SigningRegion:     "eu-west-1",
					HostnameImmutable: true,
				}, nil
			}),
		))
	if err != nil {
		return nil, err
	}
	return &cfg, err
}
