package query_test

import (
	"context"
	"myapp/01_dynamodb/query"
	"myapp/ddb"
	"myapp/model"
	"strconv"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQuery_FindByGsiPk(t *testing.T) {
	// テストデータの作成
	location, err := time.LoadLocation("Asia/Tokyo")
	require.NoError(t, err)

	data := []model.DDBData{
		{PK: 1, SK: time.Date(2020, 12, 31, 0, 0, 0, 0, location).Format(time.RFC3339), GsiPk: "false", Type: "1", Value1: 123.456, Value2: 123.456},
		{PK: 1, SK: time.Date(2020, 12, 31, 0, 1, 0, 0, location).Format(time.RFC3339), GsiPk: "true", Type: "1", Value1: 123.456, Value2: 123.456},
		{PK: 1, SK: time.Date(2020, 12, 31, 23, 59, 59, 0, location).Format(time.RFC3339), GsiPk: "false", Type: "1", Value1: 123.456, Value2: 123.456},
		{PK: 1, SK: time.Date(2021, 1, 1, 0, 0, 0, 0, location).Format(time.RFC3339), GsiPk: "false", Type: "1", Value1: 123.456, Value2: 123.456},
		{PK: 1, SK: time.Date(2021, 1, 1, 0, 1, 0, 0, location).Format(time.RFC3339), GsiPk: "false", Type: "1", Value1: 123.456, Value2: 123.456},
		{PK: 1, SK: time.Date(2021, 1, 1, 0, 2, 0, 0, location).Format(time.RFC3339), GsiPk: "true", Type: "1", Value1: 123.456, Value2: 123.456},
		{PK: 1, SK: time.Date(2021, 1, 1, 0, 3, 0, 0, location).Format(time.RFC3339), GsiPk: "false", Type: "1", Value1: 123.456, Value2: 123.456},
	}
	var writeRequests []types.WriteRequest
	for _, d := range data {
		dataMap, err := attributevalue.MarshalMap(d)
		require.NoError(t, err)
		writeRequests = append(
			writeRequests,
			types.WriteRequest{PutRequest: &types.PutRequest{Item: dataMap}},
		)
	}
	ddbClient := ddb.GetDDB()
	_, err = ddbClient.Client.BatchWriteItem(context.Background(), &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{ddb.TableName: writeRequests},
	})
	require.NoError(t, err)

	anotherTypeData := model.DDBData2{PK: 1, SK: time.Date(2021, 1, 1, 0, 4, 0, 0, location).Format(time.RFC3339), GsiPk: "false", Type: "2", Value1: 123.456, Value2: 123.456, Value3: 123.456}
	anotherTypeDataMap, err := attributevalue.MarshalMap(anotherTypeData)
	require.NoError(t, err)
	_, err = ddbClient.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(ddb.TableName),
		Item:      anotherTypeDataMap,
	})
	require.NoError(t, err)

	t.Run("0件取得", func(t *testing.T) {
		notExistItemSk := time.Date(2021, 1, 1, 0, 4, 1, 0, location).Format(time.RFC3339)
		actual, err := query.FindByGsiPk(notExistItemSk)
		assert.ErrorIs(t, err, query.ErrItemNotFound)
		assert.Len(t, actual, 0)
	})

	expected := []model.DDBData{
		{PK: 1, SK: time.Date(2021, 1, 1, 0, 0, 0, 0, location).Format(time.RFC3339), GsiPk: "false", Type: "1", Value1: 123.456, Value2: 123.456},
		{PK: 1, SK: time.Date(2021, 1, 1, 0, 1, 0, 0, location).Format(time.RFC3339), GsiPk: "false", Type: "1", Value1: 123.456, Value2: 123.456},
		{PK: 1, SK: time.Date(2021, 1, 1, 0, 3, 0, 0, location).Format(time.RFC3339), GsiPk: "false", Type: "1", Value1: 123.456, Value2: 123.456},
	}
	anotherExpected := model.DDBData2{PK: 1, SK: time.Date(2021, 1, 1, 0, 4, 0, 0, location).Format(time.RFC3339), GsiPk: "false", Type: "2", Value1: 123.456, Value2: 123.456, Value3: 123.456}
	var e []map[string]types.AttributeValue
	for _, d := range expected {
		dataMap, err := attributevalue.MarshalMap(d)
		require.NoError(t, err)
		e = append(e, dataMap)
	}
	anotherDataMap, err := attributevalue.MarshalMap(anotherExpected)
	require.NoError(t, err)
	e = append(e, anotherDataMap)

	t.Run("正常系", func(t *testing.T) {
		sk := time.Date(2021, 1, 1, 0, 0, 0, 0, location).Format(time.RFC3339)
		actual, err := query.FindByGsiPk(sk)
		if assert.NoError(t, err) {
			assert.Equal(t, e, actual)
		}
	})

	// テストデータの削除
	var deleteRequests []types.WriteRequest
	for _, d := range data {
		deleteRequest := types.WriteRequest{
			DeleteRequest: &types.DeleteRequest{
				Key: map[string]types.AttributeValue{
					ddb.PK: &types.AttributeValueMemberN{Value: strconv.Itoa(int(d.PK))},
					ddb.SK: &types.AttributeValueMemberS{Value: d.SK},
				},
			},
		}
		deleteRequests = append(deleteRequests, deleteRequest)
	}
	deleteRequest := types.WriteRequest{
		DeleteRequest: &types.DeleteRequest{
			Key: map[string]types.AttributeValue{
				ddb.PK: &types.AttributeValueMemberN{Value: strconv.Itoa(int(anotherTypeData.PK))},
				ddb.SK: &types.AttributeValueMemberS{Value: anotherTypeData.SK},
			},
		},
	}
	deleteRequests = append(deleteRequests, deleteRequest)

	_, err = ddbClient.Client.BatchWriteItem(context.Background(), &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]types.WriteRequest{
			ddb.TableName: deleteRequests,
		},
	})
	assert.NoError(t, err)
}
