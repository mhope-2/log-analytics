package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/mhope-2/log-reader/database/cassandra"
	"github.com/mhope-2/log-reader/database/types"
	"github.com/monzo/gocassa"
	"log"
)

func CreateMessage(msg *types.Message) {
	keySpace, err := cassandra.GetKeySpace()

	salesTable := keySpace.Table("message", &types.Message{}, gocassa.Keys{
		PartitionKeys: []string{"ID"},
	})

	err = salesTable.Set(types.Message{
		ID:        uuid.New().String(),
		Level:     msg.Level,
		Value:     msg.Value,
		Timestamp: msg.Timestamp,
	}).RunWithContext(context.TODO())
	if err != nil {
		log.Printf("Error saving mesage to db; err: %v", err)
	}
}
