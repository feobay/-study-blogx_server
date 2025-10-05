// conf/init_system.go

package conf

type System struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
	Env  string `yaml:"env"`
}
