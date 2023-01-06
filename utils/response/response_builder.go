package response

import utl "mindmap-go/utils"

type RespBodyBuilder interface {
	WithCode(code int) *Builder
	WithErrors(errors *utl.ErrorData) *Builder
	WithMessages(msg Messages) *Builder
	WithData(data any) *Builder
	Build() *RespBody
}

func NewResponseBuilder() RespBodyBuilder {
	return &Builder{}
}

type Builder struct {
	code     int
	messages Messages
	errors   *utl.ErrorData
	data     any
}

func (b *Builder) WithErrors(errors *utl.ErrorData) *Builder {
	b.errors = errors
	return b
}

func (b *Builder) WithMessages(msg Messages) *Builder {
	b.messages = msg
	return b
}

func (b *Builder) WithData(data any) *Builder {
	b.data = data
	return b
}

func (b *Builder) WithCode(code int) *Builder {
	b.code = code
	return b
}

func (b *Builder) Build() *RespBody {
	return &RespBody{
		Code:     b.code,
		Messages: b.messages,
		Errors:   b.errors,
		Data:     b.data,
	}
}
