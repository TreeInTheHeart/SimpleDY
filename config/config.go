package config

//服务器全局配置
type Config struct {
	Mysql Mysql `mapstructure:"mysql"`
}
//Mysql数据库配置
type Mysql struct {
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Url             string `mapstructure:"url"`
}
