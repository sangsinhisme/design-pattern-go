package main

import "fmt"

type DBConfig struct {
	Host, User, Password, Database, SSLMode string
	Port, PoolSize                          int
}

// Base builder
type DBConfigBuilder struct {
	config *DBConfig
}

func NewDBConfigBuilder() *DBConfigBuilder {
	return &DBConfigBuilder{&DBConfig{}}
}

func (b *DBConfigBuilder) Build() *DBConfig {
	return b.config
}

// MySQL-specific builder
type MySQLConfigBuilder struct {
	*DBConfigBuilder
}

func (b *DBConfigBuilder) ForMySQL() *MySQLConfigBuilder {
	b.config.Database = "mysql"
	return &MySQLConfigBuilder{b}
}

func (b *MySQLConfigBuilder) Host(host string) *MySQLConfigBuilder {
	b.config.Host = host
	return b
}

func (b *MySQLConfigBuilder) Port(port int) *MySQLConfigBuilder {
	b.config.Port = port
	return b
}

func (b *MySQLConfigBuilder) User(user string) *MySQLConfigBuilder {
	b.config.User = user
	return b
}

func (b *MySQLConfigBuilder) Password(password string) *MySQLConfigBuilder {
	b.config.Password = password
	return b
}

func (b *MySQLConfigBuilder) SSLMode(sslMode string) *MySQLConfigBuilder {
	b.config.SSLMode = sslMode
	return b
}

func (b *MySQLConfigBuilder) PoolSize(poolSize int) *MySQLConfigBuilder {
	b.config.PoolSize = poolSize
	return b
}

func main() {
	mysqlConfig := NewDBConfigBuilder().
		ForMySQL().
		Host("localhost").
		Port(3306).
		User("admin").
		Password("secret").
		SSLMode("disable").
		PoolSize(10).
		Build()

	fmt.Println(*mysqlConfig)
}
