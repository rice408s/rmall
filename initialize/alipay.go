package initialize

import (
	"github.com/go-pay/gopay/pkg/xlog"
	"rmall/global"

	"github.com/smartwalle/alipay/v3"
)

const (
	privateKey = "MIIEogIBAAKCAQEAlupeoHT24F7+zvV3X2LRoEg8OTGCS8+xnA1sN22B9YINy1dz7kFcb8D+9SngvkLciq/O2c7yVYJ8WnOxnQwZ7d2pY91jyyiYZmW3DrPUGGyRk44If++LQLo3QEWJD8XwykHC1rtox94diDIhG4vrDKy/vtYsoPGMPk5CXzLICqZJcVhm0bnH22brLvHv6wwEmnzviN7wY54uHqkzhnqF39MZx6thh4gx2FhzqrwP4bBeHxR4UqoeEWw5qTgcAS6loIbgDblPicQT/+sUXr8J2svjddbz5FJnpmpbU43DZqNoMS4XYjZcftVBqkGHTuNK3r7t4e9FTYuocRq11NsMowIDAQABAoIBABWY9u/ZrVcQ4UB3Cp1hBAT/MXcx6aa104I0vY71IgsjkJvQKSwouHTZ1uL5GkvAP9WtMVokbrPkF63OLtX93FwZQ9uh3F0/swe8eMUm4SyZpP/L7NV6T6B0xeFx0hqqxjn7zUlzypwH4owSOIC7tmi69Dn4ZqylxN9Jzd/X4Z78UmA3qv7fNoHSQE0Mpy/za83KvJflXWXXbnhlmZe0Nq4gv7u8KeZ57yrl5maS9A1YBNYmAuROW6hIfQ43jI+kmO+b6Xba4W+f6FlPQb8RlrJ6iQSEvhb16rO+oxBYPVSRDThcD5CJqbulTTzLhQAnTq3psuYhdtk6PaJJ5BUuxykCgYEA4wZASJPv3LyPFi2g+SQ0OrNCCKL0/aIkAPkbobgklpOEy7lCFNkWY/8gSYDyKSrHC44VVIP9GZ2WqWJqy2SrckeFKfaAeIzdT0PTiYpbinpiseHb/6GpzYgjtErQ+FQY60mPPftNh+s9Jkk599WU1zF59O9k3TfJAMqn68K7DNcCgYEAqi1aB6BYald1pgWvSgKTOvQSP8fc+aDkQ3qMO32oKgcS8Sh7Gxigbwe6b4jEBAVv78f5FEuLMyn/SdYOnFKQ4wpDw2PGcDxTYEh7I+skwLEWsc7/qNozTBdOds3ojlimxXiTog8DaIjPyr80OH5Y/gd/FrQiYsVYyAH0S0iOGRUCgYBAxILIcfQnbGAB+siG644/BO8c2ai9R2mwIWGBtLAziSiEmnjAy+I4awvCDJiJqBubL1DkiLdHfI8ECgDfv/utobva6eli+wUUXQxXm/JbwmivyPjjqDJyYZmWERYnEo79aUIFc9F9ql4Ksy9nhjZtYj8h4r5K4mRk5mwbtbV5hwKBgDYb8XJXU4WMhYXibNfVdcba8CqDfAmLVMj34DrSO3QRh75SJsJ41Bo5lNf93oWAzRSqutHvKzV7dwadH0wpHJR0IROxV9k79S66HR1QmQEDzbl/nd6scn8RMguocPdBzNUg/6AZsfW8+oT2wr3rvpQXAhvP7tcE9LOzm1+MlFH5AoGAW3KImZPgkzXuOQq/JJOqINvODv3GCkvzUQpCtxnAwm1VSl/45Tm4/cKuUhImUwEvamvqi7ACzHQyoJUVTEIgI5H6LSCeeEFp4OEaU5xPs8FJtGaGkCHE313neoa3/+h2kw6L3Wa+SCRhwaQLF6sdeuss5hMWSsyoFMQ4qIW907s="
	appId      = "9021000132694496"
	isProd     = false
)

func initAliPay() {
	// 初始化支付宝客户端
	// appid：应用ID
	// privateKey：应用私钥，支持PKCS1和PKCS8
	// isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
	client, err := alipay.New(appId, privateKey, isProd)
	if err != nil {
		xlog.Error(err)
		panic(err)
	}
	err = client.LoadAppCertPublicKeyFromFile("certificate/appPublicCert.crt")
	if err != nil {
		panic(err)
	} // 加载应用公钥证书
	err = client.LoadAliPayRootCertFromFile("certificate/alipayRootCert.crt")
	if err != nil {
		panic(err)
	} // 加载支付宝根证书
	err = client.LoadAlipayCertPublicKeyFromFile("certificate/alipayPublicCert.crt")
	if err != nil {
		panic(err)
	} // 加载支付宝公钥证书
	global.AliPayClient = client
}
