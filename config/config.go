package config

type Config struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	JWT   JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
}

// Mysql  MySQL配置
type Mysql struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 主机名
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 端口号
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	Dbname   string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`       // 数据库名
	Params   string `mapstructure:"params" json:"params" yaml:"params"`       // 链接参数
}

// JWT jwt配置
type JWT struct {
	SecretKey string `mapstructure:"secret-key" json:"secretKey" yaml:"secret_key"` // 密钥
	Expire    int64  `mapstructure:"expire" json:"expire" yaml:"expire"`            // 过期时间
}

// DefaultConfig 默认配置

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 主机名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // 数据库名
}
