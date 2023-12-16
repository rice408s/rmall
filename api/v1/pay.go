package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"github.com/smartwalle/xid"
	"log"
	"net/http"
	"rmall/global"
)

var kServerDomain = "127.0.0.1"

func Pay(c *gin.Context) {
	var tradeNo = fmt.Sprintf("%d", xid.Next())
	var p = alipay.TradePagePay{}
	p.NotifyURL = kServerDomain + "/alipay/notify"
	p.ReturnURL = kServerDomain + "/alipay/callback"
	p.Subject = "支付测试:" + tradeNo
	p.OutTradeNo = tradeNo
	p.TotalAmount = "660000.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	//fmt.Println(p)

	//fmt.Println(client.appId)
	//fmt.Println(global.AliPayClient)

	url, err := global.AliPayClient.TradePagePay(p)
	if err != nil {
		fmt.Println("支付发生错误", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	fmt.Println(url.String())
	//c.Redirect(http.StatusTemporaryRedirect, "hello")
	c.Redirect(http.StatusTemporaryRedirect, url.String())

}

func Callback(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err := global.AliPayClient.VerifySign(c.Request.Form); err != nil {
		log.Println("回调验证签名发生错误", err)
		c.String(http.StatusBadRequest, "回调验证签名发生错误")
		return
	}

	log.Println("回调验证签名通过")

	// 示例一：使用已有接口进行查询
	var outTradeNo = c.Request.Form.Get("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo

	rsp, err := global.AliPayClient.TradeQuery(p)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("验证订单 %s 信息发生错误: %s", outTradeNo, err.Error()))
		return
	}

	if rsp.IsFailure() {
		c.String(http.StatusBadRequest, fmt.Sprintf("验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Msg, rsp.SubMsg))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("订单 %s 支付成功", outTradeNo))
}

func Notify(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		log.Println("解析异步通知发生错误", err)
		return
	}

	notification, err := global.AliPayClient.DecodeNotification(c.Request.Form)
	if err != nil {
		log.Println("解析异步通知发生错误", err)
		return
	}

	log.Println("解析异步通知成功:", notification.NotifyId)

	// 示例一：使用自定义请求进行查询
	var p = alipay.NewPayload("alipay.trade.query")
	p.AddBizField("out_trade_no", notification.OutTradeNo)

	var rsp *alipay.TradeQueryRsp
	if err = global.AliPayClient.Request(p, &rsp); err != nil {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s \n", notification.OutTradeNo, err.Error())
		return
	}
	if rsp.IsFailure() {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s-%s \n", notification.OutTradeNo, rsp.Msg, rsp.SubMsg)
		return
	}

	log.Printf("订单 %s 支付成功 \n", notification.OutTradeNo)

	global.AliPayClient.ACKNotification(c.Writer)
}

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
