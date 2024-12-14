package tools

import (
	"context"
	"doit/callbacks"
	"errors"
	"github.com/mitchellh/mapstructure"
)

type Handler interface {
	Run() error
}

type Tool struct {
	ctx context.Context
	*Options
	TaskId any

	Step callbacks.Callback
}

func (t *Tool) validateInput() error {
	return mapstructure.Decode(t.argsSchema, t.inputSchema)
}

func (t *Tool) validateOutput() error {
	return mapstructure.Decode(t.payload, &t.outputSchema)
}

func (t *Tool) Run() error {
	var err error
	if t.TaskId == nil || t.traceId == nil {
		return errors.New("TaskId is required")
	}
	//模型提取出来的参数和这里校验,结果不一致就返回
	//if t.argsSchema(nil) != nil {
	//
	//}
	if err = t.validateInput(); err != nil {
		return ValidateInputErr
	}

	if err = t.validateOutput(); err == nil {
		t.isResult = true
	}

	if t.Step != nil {
		err = t.Step.Handler(t.ctx, t.TaskId, t.traceId)
	}

	return err
}
func NewTool(taskId any, options ...Option) (*Tool, error) {
	o := New(options...)
	return &Tool{TaskId: taskId, Options: o}, nil
}
