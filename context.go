package gospell_gin

import (
	"github.com/bridgex-dev/spell"
	"github.com/gin-gonic/gin"
)

type ResponseWriter struct {
	gin.ResponseWriter
	spell *spell.Context
}

func (w *ResponseWriter) WriteHeaderNow() {
	_ = w.spell.WriteHeader()
	w.ResponseWriter.WriteHeaderNow()
}

func NewResponseWriter(r *gin.Context, spell *spell.Context) *ResponseWriter {
	return &ResponseWriter{r.Writer, spell}
}

type Context struct {
	*gin.Context
	Spell *spell.Context
}

func NewContext(r *gin.Context, spell *spell.Context) *Context {
	// wrap WriteHeaderNow to write spell headers
	r.Writer = NewResponseWriter(r, spell)

	return &Context{
		r,
		spell,
	}
}
