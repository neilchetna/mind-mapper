package worker

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hibiken/asynq"
)

const (
	StartingWorker = "Starting worker"
	Success        = "Task performed successfully"
	Err            = "Worker Failed"
)

func LoggerStart(ctx context.Context, t *asynq.Task) {
	var str strings.Builder
	str.WriteString(StartingWorker)
	id, _ := asynq.GetTaskID(ctx)
	str.WriteString(fmt.Sprintf(", taskId = %s, ", id))
	str.WriteString(fmt.Sprintf("type = %s", t.Type()))
	log.Default().Print(str.String())
}

func LoggerSuccess(ctx context.Context, t *asynq.Task) {
	log.Print(Success)
}

func LoggerErr(ctx context.Context, t *asynq.Task, err error) {
	var str strings.Builder
	str.WriteString(Err)
	id, _ := asynq.GetTaskID(ctx)
	str.WriteString(fmt.Sprintf(", taskId = %s, ", id))
	str.WriteString(fmt.Sprintf("type = %s", t.Type()))
	str.WriteString("Error:")
	str.WriteString(err.Error())
	log.Fatal(str.String())
}
