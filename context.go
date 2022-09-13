package gospell_gin

import (
	"github.com/bridgex-dev/spell"
	"github.com/gin-gonic/gin"
)

type ResponseWriter struct {
	gin.ResponseWriter
	spell *spell.Context
}

func (w *ResponseWriter) WriteHeader(status int) {
	_ = w.spell.WriteHeader()
	w.ResponseWriter.WriteHeader(status)
}

func NewResponseWriter(r *gin.Context, spell *spell.Context) *ResponseWriter {
	return &ResponseWriter{r.Writer, spell}
}

type Context struct {
	*gin.Context
	Spell *spell.Context
}

func NewHybridContext(r *gin.Context, spell *spell.Context) *Context {
	// wrap WriteHeaderNow to write spell headers
	r.Writer = NewResponseWriter(r, spell)

	return &Context{
		r,
		spell,
	}
}
