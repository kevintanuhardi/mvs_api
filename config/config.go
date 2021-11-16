package config

import "time"

// Config struct generate
type Config struct {
	App struct {
		Name         string         `yaml:"name"`
		Version      string         `yaml:"version"`
		Latency      int            `yaml:"latency"`
		ReadTimeout  int            `yaml:"read_timeout"`
		WriteTimeout int            `yaml:"write_timeout"`
		Debug        bool           `yaml:"debug"`
		Env          string         `yaml:"env"`
		SecretKey    string         `yaml:"secret_key"`
		ExpireIn     *time.Duration `yaml:"expire_in"`
	} `yaml:"App"`

	Port struct {
		HTTP       int `yaml:"http"`
		HTTPMetric int `yaml:"http_metric"`
		Grpc       int `yaml:"grpc"`
		GrpcMetric int `yaml:"grpc_metric"`
	} `yaml:"Ports"`

	DB struct {
		MasterDSN   string `yaml:"master_dsn" mapstructure:"DB_MASTER_DSN"`
		ReplicaDSN  string `yaml:"replica_dsn" mapstructure:"DB_REPLICA_DSN"`
		MaxLifeTime int    `yaml:"max_life_time" mapstructure:"DB_MAX_LIFE_TIME"`
		MaxOpen     int    `yaml:"max_open" mapstructure:"DB_MAX_OPEN"`
		MaxIdle     int    `yaml:"max_idle" mapstructure:"DB_MAX_IDLE"`
	} `yaml:"DB"`
}
