package config

type System struct {
	Env              string `mapstructure:"env" json:"env" yaml:"env"`                      // 环境值
	AdminAddr        int    `mapstructure:"admin_addr" json:"admin_addr" yaml:"admin_addr"` // 端口值
	ApiAddr          int    `mapstructure:"api_addr" json:"api_addr" yaml:"api_addr"`       // 端口值
	AgentAddr        int    `mapstructure:"agent_addr" json:"agent_addr" yaml:"agent_addr"` // 端口值
	Debug            bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
	WithdrawPassword string `mapstructure:"withdraw_password" json:"withdraw_password" yaml:"withdraw_password"`
}
