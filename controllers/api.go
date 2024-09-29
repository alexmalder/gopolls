package controllers

import "github.com/gin-gonic/gin"

func Listen() {
	r := gin.Default()

	r.GET("/api/polls", GetPolls)
	r.POST("/api/polls", PostPoll)
	r.GET("/api/polls/get_by_owner_cn", GetPollsByOwnerCn)
	r.PUT("/api/polls", PutPoll)
	r.DELETE("/api/polls", DeletePoll)
	r.POST("/api/instances", PostInstance)
	r.PUT("/api/instances", PutInstance)
	r.POST("/api/questions", PostQuestion)
	r.PUT("/api/questions", PutQuestion)
	r.POST("/api/answers", PostAnswer)
	r.PUT("/api/answers", PutAnswer)
	r.POST("/api/replies", PostReply)
	r.PUT("/api/replies", PutReply)
	r.Run(":8000")
}
