package gospell_gin

import (
	"github.com/bridgex-dev/spell"
	"github.com/gin-gonic/gin"
	"github.com/gwatts/gin-adapter"
)

type RouterGroup struct {
	*gin.RouterGroup
	spell *spell.Engine
}

func NewRouterGroup(spell *spell.Engine, parent *gin.RouterGroup) *RouterGroup {
	parent.Use(adapter.Wrap(spell.Handler))

	return &RouterGroup{
		RouterGroup: parent,
		spell:       spell,
	}
}

type HandlerFunc = func(*Context)

func (r *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return r.RouterGroup.Handle(httpMethod, relativePath, r.wrap(handlers...)...)
}

func (r *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return r.RouterGroup.POST(relativePath, r.wrap(handlers...)...)
}

func (r *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return r.RouterGroup.GET(relativePath, r.wrap(handlers...)...)
}

func (r *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return r.RouterGroup.DELETE(relativePath, r.wrap(handlers...)...)
}

func (r *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return r.RouterGroup.PATCH(relativePath, r.wrap(handlers...)...)
}

func (r *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return r.RouterGroup.PUT(relativePath, r.wrap(handlers...)...)
}

func (r *RouterGroup) OPTIONS(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return r.RouterGroup.OPTIONS(relativePath, r.wrap(handlers...)...)
}

func (r *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return r.RouterGroup.HEAD(relativePath, r.wrap(handlers...)...)
}

func (r *RouterGroup) Any(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return r.RouterGroup.Any(relativePath, r.wrap(handlers...)...)
}

func (r *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		RouterGroup: r.RouterGroup.Group(relativePath, r.wrap(handlers...)...),
		spell:       r.spell,
	}
}

func (r *RouterGroup) wrap(handlers ...HandlerFunc) []gin.HandlerFunc {
	res := make([]gin.HandlerFunc, 0, len(handlers))
	for _, h := range handlers {
		res = append(res, func(c *gin.Context) {
			h(&Context{
				Context: c,
				Spell:   r.spell.GetContext(c.Request),
			})
		})
	}

	return res
}
