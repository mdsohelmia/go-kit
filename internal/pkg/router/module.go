package router

import "go.uber.org/fx"

// Register gin.Engine as Router
var Module = fx.Module("router", fx.Provide(NewRouter))
