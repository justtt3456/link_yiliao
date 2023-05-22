package config

type Server struct {
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Sms    Sms    `mapstructure:"sms" json:"sms" yaml:"sms"`
	Log    Log    `mapstructure:"log" json:"log" yaml:"log"`
}
