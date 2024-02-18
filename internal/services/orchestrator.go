package services

import (
	"context"
	"encoding/json"
	"expression-backend/internal/redis"
	"fmt"
	redis2 "github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	Id         int
	Expression string
	TimeLimit  time.Duration
}

type TaskResponse struct {
	Id          int
	Expression  string
	IsTimeLimit bool
	IsValid     bool
	Result      int
}

func NewTask(id int, expression string, timeLimit time.Duration) *Task {
	return &Task{
		Id:         id,
		Expression: expression,
		TimeLimit:  timeLimit,
	}
}

func (t *Task) Execute() *TaskResponse {
	ctx, cancel := context.WithTimeout(context.Background(), t.TimeLimit)
	defer cancel()

	taskResp := TaskResponse{
		Id:          t.Id,
		Expression:  t.Expression,
		IsTimeLimit: false,
		IsValid:     true,
		Result:      0,
	}

	exp := NewExpression(t.Expression)
	tokensAsString, err := exp.GetTokensAsString()
	if err != nil {
		taskResp.IsValid = false
		return &taskResp
	}

	rdb, _ := redis.GetConnection()
	channel := fmt.Sprintf("mini_answer_channel_%d", taskResp.Id)
	subPub := rdb.Subscribe(context.Background(), channel)

	for {
		log.Println("for 1")
		singleExps, _ := FindSingleExpressions(tokensAsString)

		for _, singleExp := range singleExps {
			redis.PublishToQueue(&redis.MiniTask{
				ExpressionId:   taskResp.Id,
				MiniExpression: strings.TrimSpace(singleExp),
			})
			log.Println("for 2")
		}

		select {
		case <-ctx.Done():
			taskResp.IsTimeLimit = true
			fmt.Println("time limit exceeded")
			return &taskResp
		default:
			var counter int
			var miniTask *redis.MiniTask
			for len(singleExps) != counter {
				log.Println("for 3")
				select {
				case <-ctx.Done():
					taskResp.IsTimeLimit = true
					fmt.Println("time limit exceeded")
					return &taskResp
				default:
					inter, err := subPub.ReceiveTimeout(ctx, t.TimeLimit)
					if err == nil {
						switch msg := inter.(type) {
						case *redis2.Message:
							fmt.Println("case message", msg.Payload)
							counter++
							json.Unmarshal([]byte(msg.Payload), &miniTask)
							fmt.Println("minitask", msg.Payload, miniTask)
							if miniTask.IsValid == false {
								taskResp.IsValid = false
								return &taskResp
							} else {
								tokensAsString = strings.ReplaceAll(tokensAsString, " "+miniTask.MiniExpression+" ", " "+strconv.Itoa(miniTask.Result)+" ")
							}
						default:
							err := fmt.Sprintf("redis: unknown message: %T", msg)
							fmt.Println(err)
						}
					} else {
						break
					}
				}
			}
		}
		if len(singleExps) == 0 {
			break
		}

		fmt.Println(tokensAsString)
	}

	taskResp.Result, err = strconv.Atoi(strings.ReplaceAll(tokensAsString, " ", ""))
	if err != nil {
		fmt.Println(err)
	}
	return &taskResp
}
