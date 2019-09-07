package api

import (
	"github.com/gin-gonic/gin"
	 "go-crud/service"
)

//
func DailyRank(c *gin.Context){
	service:=service.DailyRankService{}
	if err:=c.ShouldBind(&service); err==nil{
		res:=service.Get()
		c.JSON(200,res)
	}else{
		c.JSON(200,ErrorResponse(err))
	}
}