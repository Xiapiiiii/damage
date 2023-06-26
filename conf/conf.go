package conf

type MysqlConfig struct {
	Datasource      string `yaml:"Datasource" json:"Datasource"`
	ConnMaxIdle     int    `yaml:"ConnMaxIdle" json:"ConnMaxIdle,omitempty,optional"`
	ConnMaxOpen     int    `yaml:"ConnMaxOpen" json:"ConnMaxOpen,omitempty,optional"`
	ConnMaxLifeTime int    `yaml:"ConnMaxLifetime" json:"ConnMaxLifetime,omitempty,optional"`
}
