package config

// MySQL MySQL配置
type Mysql struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 主机名
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 端口号
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 用户名
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	Dbname   string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`       // 数据库名
	Params   string `mapstructure:"params" json:"params" yaml:"params"`       // 链接参数
}
