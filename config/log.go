package config

type Log struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"` // 级别
	File string `mapstructure:"file" json:"file" yaml:"file"` // 输出
}

func (this Log) Abs() string {
	if this.Path[len(this.Path)-2:len(this.Path)-1] == "/" {
		return this.Path + this.File
	}
	return this.Path + "/" + this.File
}
