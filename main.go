package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// SQSListQueuesAPI defines the interface for the ListQueues function.
// We use this interface to test the function using a mocked service.
type SQSListQueuesAPI interface {
	ListQueues(ctx context.Context,
		params *sqs.ListQueuesInput,
		optFns ...func(*sqs.Options)) (*sqs.ListQueuesOutput, error)
}

func GetQueues(c context.Context, api SQSListQueuesAPI, input *sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error) {
	return api.ListQueues(c, input)
}
func main() {

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {

		return aws.Endpoint{
			URL: "http://localhost:9324",
		}, nil

	})

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := sqs.NewFromConfig(cfg)

	input := &sqs.ListQueuesInput{}

	result, err := GetQueues(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving queue URLs:")
		fmt.Println(err)
		return
	}

	for i, url := range result.QueueUrls {
		fmt.Printf("%d: %s\n", i+1, url)
	}
}
