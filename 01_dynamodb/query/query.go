package query

import (
	"context"
	"errors"
	"myapp/ddb"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var ErrItemNotFound = errors.New("item not found")

func FindByGsiPk(sk string) (items []map[string]types.AttributeValue, err error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String(ddb.TableName),
		IndexName:              aws.String(ddb.GSIIndex),
		KeyConditionExpression: aws.String("#gsiPk = :gsiPk AND #sk >= :sk"),
		ExpressionAttributeNames: map[string]string{
			"#gsiPk": ddb.GSIPK,
			"#sk":    ddb.SK,
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":gsiPk": &types.AttributeValueMemberS{Value: "false"},
			":sk":    &types.AttributeValueMemberS{Value: sk},
		},
	}

	ddbClient := ddb.GetDDB()

	queryPaginator := dynamodb.NewQueryPaginator(ddbClient.Client, input)
	for queryPaginator.HasMorePages() {
		output, err := queryPaginator.NextPage(context.Background())
		if err != nil {
			// TODO 取得できたもので処理を行うかどうかを考える
			break
		}

		items = append(items, output.Items...)
	}

	if len(items) == 0 {
		err = ErrItemNotFound
	}

	return
}
