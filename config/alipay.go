package config

import (
	"context"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	"strings"
)

const (
	privateKey = "MIIEogIBAAKCAQEAlupeoHT24F7+zvV3X2LRoEg8OTGCS8+xnA1sN22B9YINy1dz7kFcb8D+9SngvkLciq/O2c7yVYJ8WnOxnQwZ7d2pY91jyyiYZmW3DrPUGGyRk44If++LQLo3QEWJD8XwykHC1rtox94diDIhG4vrDKy/vtYsoPGMPk5CXzLICqZJcVhm0bnH22brLvHv6wwEmnzviN7wY54uHqkzhnqF39MZx6thh4gx2FhzqrwP4bBeHxR4UqoeEWw5qTgcAS6loIbgDblPicQT/+sUXr8J2svjddbz5FJnpmpbU43DZqNoMS4XYjZcftVBqkGHTuNK3r7t4e9FTYuocRq11NsMowIDAQABAoIBABWY9u/ZrVcQ4UB3Cp1hBAT/MXcx6aa104I0vY71IgsjkJvQKSwouHTZ1uL5GkvAP9WtMVokbrPkF63OLtX93FwZQ9uh3F0/swe8eMUm4SyZpP/L7NV6T6B0xeFx0hqqxjn7zUlzypwH4owSOIC7tmi69Dn4ZqylxN9Jzd/X4Z78UmA3qv7fNoHSQE0Mpy/za83KvJflXWXXbnhlmZe0Nq4gv7u8KeZ57yrl5maS9A1YBNYmAuROW6hIfQ43jI+kmO+b6Xba4W+f6FlPQb8RlrJ6iQSEvhb16rO+oxBYPVSRDThcD5CJqbulTTzLhQAnTq3psuYhdtk6PaJJ5BUuxykCgYEA4wZASJPv3LyPFi2g+SQ0OrNCCKL0/aIkAPkbobgklpOEy7lCFNkWY/8gSYDyKSrHC44VVIP9GZ2WqWJqy2SrckeFKfaAeIzdT0PTiYpbinpiseHb/6GpzYgjtErQ+FQY60mPPftNh+s9Jkk599WU1zF59O9k3TfJAMqn68K7DNcCgYEAqi1aB6BYald1pgWvSgKTOvQSP8fc+aDkQ3qMO32oKgcS8Sh7Gxigbwe6b4jEBAVv78f5FEuLMyn/SdYOnFKQ4wpDw2PGcDxTYEh7I+skwLEWsc7/qNozTBdOds3ojlimxXiTog8DaIjPyr80OH5Y/gd/FrQiYsVYyAH0S0iOGRUCgYBAxILIcfQnbGAB+siG644/BO8c2ai9R2mwIWGBtLAziSiEmnjAy+I4awvCDJiJqBubL1DkiLdHfI8ECgDfv/utobva6eli+wUUXQxXm/JbwmivyPjjqDJyYZmWERYnEo79aUIFc9F9ql4Ksy9nhjZtYj8h4r5K4mRk5mwbtbV5hwKBgDYb8XJXU4WMhYXibNfVdcba8CqDfAmLVMj34DrSO3QRh75SJsJ41Bo5lNf93oWAzRSqutHvKzV7dwadH0wpHJR0IROxV9k79S66HR1QmQEDzbl/nd6scn8RMguocPdBzNUg/6AZsfW8+oT2wr3rvpQXAhvP7tcE9LOzm1+MlFH5AoGAW3KImZPgkzXuOQq/JJOqINvODv3GCkvzUQpCtxnAwm1VSl/45Tm4/cKuUhImUwEvamvqi7ACzHQyoJUVTEIgI5H6LSCeeEFp4OEaU5xPs8FJtGaGkCHE313neoa3/+h2kw6L3Wa+SCRhwaQLF6sdeuss5hMWSsyoFMQ4qIW907s="
	appId      = "9021000132694496"
	isProd     = false
)

func TestAliPay() error {
	// 初始化支付宝客户端
	// appid：应用ID
	// privateKey：应用私钥，支持PKCS1和PKCS8
	// isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
	client, err := alipay.NewClient(appId, privateKey, isProd)
	if err != nil {
		xlog.Error(err)
		return err
	}

	// 自定义配置http请求接收返回结果body大小，默认 10MB
	//client.SetBodySize() // 没有特殊需求，可忽略此配置

	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn

	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).                            // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2).                           // 设置签名类型，不设置默认 RSA2
							SetReturnUrl("https://www.rice408.cn").             // 设置返回URL
							SetNotifyUrl("https://localhost:8080/api/v1/hello") // 设置异步通知URL
	//SetAppAuthToken()                    // 设置第三方应用授权

	// 设置biz_content加密KEY，设置此参数默认开启加密
	client.SetAESKey("1234567890123456")

	//读取证书数据

	// 自动同步验签（只支持证书模式）
	// 传入 alipayPublicCert.crt 内容
	//client.AutoVerifySign([]byte("alipayPublicCert.crt bytes"))

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	err = client.SetCertSnByPath("certificate/appPublicCert.crt", "certificate/alipayRootCert.crt", "certificate/alipayPublicCert.crt")
	if err != nil {
		xlog.Error(err)
		return err
	}

	ctx := context.Background()

	bm := make(gopay.BodyMap)
	bm.Set("subject", "测试支付")
	bm.Set("out_trade_no", "1234567890")
	bm.Set("total_amount", "1.00")
	//bm.Set("product_code", "QUICK_WAP_WAY")

	pay, err := client.TradeAppPay(ctx, bm)
	if err != nil {
		return err
	}
	//xlog.Debug(pay)
	fmt.Println("good")
	fmt.Println(pay)

	fmt.Println("=====================================")
	//将字符串以&分割
	str := strings.Split(pay, "&")
	for _, s := range str {
		fmt.Println(s)
	}

	//路由跳转

	return nil
	//// 证书内容
	//err = client.SetCertSnByContent([]byte("appPublicCert.crt bytes"), []byte("alipayRootCert bytes"), []byte("alipayPublicCert.crt bytes"))
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}

}
