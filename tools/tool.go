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

	Callback callbacks.Callback //如果tool的callback 是nil,则集成来自对应的agent的
}

func (t *Tool) validateInput() error {
	return mapstructure.Decode(t.argsSchema, t.inputSchema)
}

func (t *Tool) validateOutput() error {
	if t.output == nil {
		return nil
	}
	return mapstructure.Decode(t.output, &t.outputSchema)
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

	if t.Callback != nil {
		err = t.Callback.Handler(t.ctx,
			map[string]any{
				"trace_id":  t.traceId,
				"task_id":   t.TaskId,
				"is_result": t.isResult,
				"output":    t.output,
				"input":     t.inputSchema,
			},
		)
	}

	return err
}
func NewTool(taskId any, options ...Option) (*Tool, error) {
	o := New(options...)
	return &Tool{TaskId: taskId, Options: o}, nil
}
