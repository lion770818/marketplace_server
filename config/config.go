package config

type Config struct {
	Web      Web      `yaml:"web"`
	Mysql    Mysql    `yaml:"mysql"`
	Auth     Auth     `yaml:"auth"`
	Redis    Redis    `yaml:"redis"`
	RabbitMq RabbitMq `yaml:"rabbitmq"`
	Log      Log      `yaml:"log"`
}
type Web struct {
	Mode string `yaml:"mode"`
	Port string `yaml:"port"`
}

type Mysql struct {
	LogMode  string `yaml:"log_mode"` // dev = open debug log
	Driver   string `yaml:"db_driver"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
type Auth struct {
	Active     string `yaml:"active"`
	ExpireTime string `yaml:"expireTime"`
	PrivateKey string `yaml:"privateKey"`
}
type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

// RabbitMq 配置
type RabbitMq struct {
	Host       string `yaml:"host"`       //地址
	Port       string `yaml:"port"`       //端口
	User       string `yaml:"user"`       //用户名
	Password   string `yaml:"password"`   //密码
	ConnectNum int    `yaml:"connectNum"` //总共连接数
	ChannelNum int    `yaml:"channelNum"` //每条连接的channel数
}

type Log struct {
	Env        string `yaml:"env"`
	Path       string `yaml:"path"`
	Encoding   string `yaml:"encoding"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
}
