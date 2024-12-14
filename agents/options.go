package agents

type Options struct {
	traceId         any                //链路id
	callbackHandler func(...any) error //回调
	prompt          any
	errHandler      func(...any) error
	maxIter         int  //重试次数
	isReturn        bool //是否直接返回结果

}

type Option func(*Options)

func WithTraceId(traceId any) Option {
	return func(o *Options) {
		o.traceId = traceId
	}
}

func WithCallbackHandler(callback func(...any) error) Option {
	return func(o *Options) {
		o.callbackHandler = callback
	}
}

func WithPrompt(prompt any) Option {
	return func(o *Options) {
		o.prompt = prompt
	}
}

func WithErrHandler(callback func(...any) error) Option {
	return func(o *Options) {
		o.errHandler = callback
	}
}

func WithMaxIter(maxIter int) Option {
	return func(o *Options) {
		o.maxIter = maxIter
	}
}

func WithIsReturn(isReturn bool) Option {
	return func(o *Options) {
		o.isReturn = isReturn
	}
}
