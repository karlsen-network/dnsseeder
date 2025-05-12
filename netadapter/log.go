package netadapter

import (
	"github.com/karlsen-network/karlsend/v2/infrastructure/logger"
	"github.com/karlsen-network/karlsend/v2/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
