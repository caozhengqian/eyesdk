package api

import (
	"server/models"

	"github.com/gin-gonic/gin"
)

type ApiController struct{}

func (con ApiController) Index(c *gin.Context) {

	c.String(200, "我是测试的一个api接口")
}
func (con ApiController) SetError(c *gin.Context) {
	var eye models.Eye
	if err := c.ShouldBindJSON(&eye); err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "参数解析失败",
			"error":   err.Error(),
		})
		return
	}
	if err := models.DB.Create(&eye).Error; err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "插入失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "插入成功",
		"data":    eye,
	})
}
