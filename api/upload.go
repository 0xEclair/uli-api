package api

import(
	"go-crud/service"

	"github.com/gin-gonic/gin"
)

func UploadToken(c *gin.Context){
	service:=service.UploadTokenService{}
	if err:=c.ShouldBind(&service); err==nil{
		res:=service.Post()
		c.JSON(200,res)
	}else{
		c.JSON(200,ErrorResponse(err))
	}
}