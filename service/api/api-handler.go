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
	rt.router.GET("/users/:id", rt.getUserProfileById)
	rt.router.PUT("/users/:id/username", rt.setMyUsername)
	rt.router.GET("/users/:id/stream", rt.getUserStream)
	rt.router.GET("/users/:id/photos", rt.getUserPhotos)

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
	rt.router.GET("/photos/:pid/likes", rt.getLikes)
	rt.router.PUT("/photos/:pid/likes/:lid", rt.likePhoto)
	rt.router.DELETE("/photos/:pid/likes/:lid", rt.unlikePhoto)

	// Comment routes
	rt.router.POST("/photos/:pid/comments", rt.commentPhoto)
	rt.router.GET("/photos/:pid/comments", rt.getComments)
	rt.router.DELETE("/photos/:pid/comments/:cid", rt.uncommentPhoto)

	return rt.router
}
