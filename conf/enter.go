// conf/enter.go

package conf

// 对于yaml包来说，一个结构体成员的标签就对应一个冒号前的字段，所以最大的字段也必须作为一个结构体成员来打标签，
type Config struct {
	System System `yaml:"system"`
	Log    Log    `yaml:"log"`
}
