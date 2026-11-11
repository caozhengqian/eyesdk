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
func (con ApiController) GetDeposit(c *gin.Context) {
	depositList := []models.Deposit{}
	models.DB.Table("kwt_indexer_exchange.deposit").Find(&depositList)
	c.JSON(200, gin.H{
		"result": depositList,
	})
}
func (con ApiController) GetCancel(c *gin.Context) {
	cancelList := []models.Cancel{}
	models.DB.Table("kwt_indexer_exchange.cancel").Find(&cancelList)
	c.JSON(200, gin.H{
		"result": cancelList,
	})
}
func (con ApiController) GetOrder(c *gin.Context) {
	orderList := []models.Order{}
	models.DB.Table("kwt_indexer_exchange.order").Find(&orderList)
	c.JSON(200, gin.H{
		"result": orderList,
	})
}
func (con ApiController) GetTrade(c *gin.Context) {
	tradeList := []models.Trade{}
	models.DB.Table("kwt_indexer_exchange.trade").Find(&tradeList)
	c.JSON(200, gin.H{
		"result": tradeList,
	})
}
func (con ApiController) GetWithdraw(c *gin.Context) {
	withdrawList := []models.Withdraw{}
	models.DB.Table("kwt_indexer_exchange.withdraw").Find(&withdrawList)
	c.JSON(200, gin.H{
		"result": withdrawList,
	})
}
func (con ApiController) GetApproval(c *gin.Context) {
	approvalList := []models.Approval{}
	models.DB.Table("kwt_indexer_kerwin_token.approval").Find(&approvalList)
	c.JSON(200, gin.H{
		"result": approvalList,
	})
}
func (con ApiController) GetTransfer(c *gin.Context) {
	transferList := []models.Transfer{}
	models.DB.Table("kwt_indexer_kerwin_token.transfer").Find(&transferList)
	c.JSON(200, gin.H{
		"result": transferList,
	})
}
func (con ApiController) Plist(c *gin.Context) {
	c.String(200, "我是一个api接口-Plist")
}
