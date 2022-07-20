package pubghealth

import (
	"context"
	"time"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/heptiolabs/healthcheck"
	"github.com/pkg/errors"
)

// DynamoTableStatusCheck verifies that the table passed in exists and is in the Active state, ready to be used.
func DynamoTableStatusCheck(client *dynamodb.DynamoDB, tableName string, timeout time.Duration, frequency time.Duration) healthcheck.Check {
	checkFunc := func() error {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		input := &dynamodb.DescribeTableInput{
			TableName: aws.String(tableName),
		}
		out, err := client.DescribeTableWithContext(ctx, input)
		if err != nil {
			return err
		}
		tableStatus := *out.Table.TableStatus
		if tableStatus != dynamodb.GlobalTableStatusActive &&
			tableStatus != dynamodb.GlobalTableStatusUpdating {
			return errors.Errorf("Dynamo table '%s' exists but is not ready, status: %s", tableName, tableStatus)
		}
		return nil
	}
	return healthcheck.Async(checkFunc, frequency)
}
