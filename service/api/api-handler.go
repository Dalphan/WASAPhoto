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

	// Ban routes
	rt.router.PUT("/users/:id/banned/:bid", rt.banUser)
	rt.router.DELETE("/users/:id/banned/:bid", rt.unbanUser)

	// Follow routes
	rt.router.PUT("/users/:id/following/:fid", rt.followUser)
	rt.router.DELETE("/users/:id/following/:fid", rt.unfollowUser)
	rt.router.GET("/users/:id/following", rt.getFollowings)
	rt.router.GET("/users/:id/followers", rt.getFollowers)

	// Photos routes
	rt.router.POST("/photos", rt.uploadPhoto)
	rt.router.DELETE("/photos/:pid", rt.deletePhoto)

	// Likes routes
	rt.router.PUT("/photos/:pid/likes/:lid", rt.likePhoto)
	rt.router.DELETE("/photos/:pid/likes/:lid", rt.unlikePhoto)

	return rt.router
}
