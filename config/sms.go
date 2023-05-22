package config

type Sms struct {
	Phone Phone `mapstructure:"phone" json:"phone" yaml:"phone"` //
	Email Email `mapstructure:"email" json:"email" yaml:"email"` //
}
type Phone struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"` //
	Password string `mapstructure:"password" json:"password" yaml:"password"` //
	Url      string `mapstructure:"url" json:"url" yaml:"url"`                //
	Dev      bool   `mapstructure:"dev" json:"dev" yaml:"dev"`                //
	Sign     string `mapstructure:"sign" json:"sign" yaml:"sign"`             //
}
type Email struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"` //
	Password string `mapstructure:"password" json:"password" yaml:"password"` //
	From     string `mapstructure:"from" json:"from" yaml:"from"`             //
	Subject  string `mapstructure:"subject" json:"subject" yaml:"subject"`    //
	Body     string `mapstructure:"body" json:"body" yaml:"body"`             //
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"` //
}
