package amazonsqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"os"
)

const AWS_ENDPOINT = "AWS_ENDPOINT"

func SetEndPoint(config aws.Config)  aws.Config {
	newConfig := config
	awsEndpoint :=  os.Getenv(AWS_ENDPOINT)
	if awsEndpoint != "" {
		newConfig.Endpoint = aws.String(awsEndpoint)
	}
	return newConfig
}
