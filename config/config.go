package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Phone   Phone   `mapstructure:"phone" json:"phone" yaml:"phone"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Log   Log   `mapstructure:"log" json:"log" yaml:"log"`
	//IM
	IM IM `mapstructure:"im" json:"im" yaml:"im"`
}
