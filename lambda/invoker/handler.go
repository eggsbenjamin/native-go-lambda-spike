package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/eggsbenjamin/native-go-lambda-spike/models"
)

func Invoker(person *models.Person) (*sfn.StartExecutionOutput, error) {
	personJSON, err := json.Marshal(person)
	if err != nil {
		return nil, err
	}

	stepFN := sfn.New(session.Must(session.NewSession()), &aws.Config{})
	startExecutionOutput, err := stepFN.StartExecution(&sfn.StartExecutionInput{
		Input:           aws.String(string(personJSON)),
		StateMachineArn: aws.String(os.Getenv("STEP_FN_ARN")),
	})
	if err != nil {
		return nil, err
	}

	for {
		describeExecutionOutput, err := stepFN.DescribeExecution(&sfn.DescribeExecutionInput{
			ExecutionArn: startExecutionOutput.ExecutionArn,
		})
		if err != nil {
			return nil, err
		}

		if *describeExecutionOutput.Status == sfn.ExecutionStatusSucceeded {
			return startExecutionOutput, nil
		}

		time.Sleep(time.Duration(1) * time.Second)
	}
}

func main() {
	lambda.Start(Invoker)
}
