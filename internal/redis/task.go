package redis

import (
	"context"
	"encoding/json"
)

type MiniTask struct {
	ExpressionId   int    `json:"expression_id"`
	MiniExpression string `json:"task"`
	IsValid        bool   `json:"is_valid"`
	Result         int    `json:"result"`
}

func PublishMiniTask(miniTask *MiniTask) {
	rdb, _ := GetConnection()
	content, _ := json.Marshal(miniTask)
	rdb.Publish(context.Background(), "mini_task_channel", string(content))
}

func PublishToQueue(miniTask *MiniTask) {
	rdb, _ := GetConnection()
	content, _ := json.Marshal(miniTask)
	rdb.LPush(context.Background(), "mini_task_queue", string(content))
}

//func Subscribe(miniTask *MiniTask)
