package config

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}

type Phone struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"` //
	Password string `mapstructure:"password" json:"password" yaml:"password"` //
	Url      string `mapstructure:"url" json:"url" yaml:"url"`                //
	Dev      bool   `mapstructure:"dev" json:"dev" yaml:"dev"`                //
	Sign     string `mapstructure:"sign" json:"sign" yaml:"sign"`             //
}
