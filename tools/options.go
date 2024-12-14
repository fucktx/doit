package tools

import "errors"

type Options struct {
	traceId      any //链路id
	name         string
	version      string
	description  string
	isResult     bool
	argsSchema   map[string]any //tool的参数param任务数据
	inputSchema  any            //tool的参数绑定信息
	output       map[string]any //返回的数据
	outputSchema any            //结果绑定结构体
}

type Option func(*Options)

func WithTraceId(traceId any) Option {
	return func(o *Options) {
		o.traceId = traceId
	}
}
func WithName(name string) Option {
	return func(o *Options) {
		o.name = name
	}
}
func WithVersion(version string) Option {
	return func(o *Options) {
		o.version = version
	}
}
func WithDescription(description string) Option {
	return func(o *Options) {
		o.description = description
	}
}

func New(options ...Option) *Options {
	opts := &Options{}
	for _, option := range options {
		option(opts)
	}

	return opts
}

var ValidateInputErr = errors.New("input validation error")
