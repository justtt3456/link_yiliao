package config

type IM struct {
	Url      string `mapstructure:"url" json:"url" yaml:"url"`
	Key      string `mapstructure:"key" json:"key" yaml:"key"`
	Redirect string `redirect:"key" redirect:"key" yaml:"redirect"`
}
