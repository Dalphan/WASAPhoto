package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	// Login routes
	rt.router.POST("/session", rt.login)

	// User routes
	rt.router.GET("/users", rt.getUserProfile)
	rt.router.PUT("/users/:id", rt.updateProfile)
	rt.router.PUT("/users/:id/username", rt.setMyUsername)

	// Ban routs
	rt.router.PUT("/users/:id/banned/:bid", rt.banUser)
	rt.router.DELETE("/users/:id/banned/:bid", rt.unbanUser)

	return rt.router
}
