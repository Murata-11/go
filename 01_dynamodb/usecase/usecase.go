package usecase

import (
	"log"
	"myapp/01_dynamodb/query"
	"myapp/model"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

func UnmarshalDDBDataByType() {
	result, err := query.FindByGsiPk("2021-01-01T00:00:00+09:00")
	if err != nil {
		log.Fatalf("failed to get data: %v", err)
	}

	for _, item := range result {
		var base model.BaseData
		err := attributevalue.UnmarshalMap(item, &base)
		if err != nil {
			log.Fatalf("failed to unmarshal data: %v", err)
		}

		switch base.Type {
		case "1":
			var data model.DDBData
			err := attributevalue.UnmarshalMap(item, &data)
			if err != nil {
				log.Fatalf("failed to unmarshal data: %v", err)
			}
			log.Printf("data: %+v", data)
		case "2":
			var data model.DDBData2
			err := attributevalue.UnmarshalMap(item, &data)
			if err != nil {
				log.Fatalf("failed to unmarshal data: %v", err)
			}
			log.Printf("data: %+v", data)
		default:
		}
	}
}
