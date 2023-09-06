package config

type IM struct {
	Api string `mapstructure:"api" json:"api" yaml:"api"`
	Key string `mapstructure:"key" json:"key" yaml:"key"`
}
