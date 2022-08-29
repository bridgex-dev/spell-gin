package gospell_gin

import (
	"github.com/bridgex-dev/spell"
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	Spell *spell.Context
}
