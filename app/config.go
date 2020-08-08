package app

import "flag"

type MQTTConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

type WebConfig struct {
	Port int
}

type Config struct {
	MQTT MQTTConfig
	Web  WebConfig
}

func getDefaultConfig() *Config {
	return &Config{
		MQTT: MQTTConfig{
			Host:     "localhost",
			Port:     1883,
			User:     "",
			Password: "",
		},
		Web: WebConfig{
			Port: 3001,
		},
	}
}

func ParseFlags() *Config {
	c := getDefaultConfig()
	flag.StringVar(&c.MQTT.Host, "mqtt-host", c.MQTT.Host, "The hostname or ip of the mqtt broker.")
	flag.IntVar(&c.MQTT.Port, "mqtt-port", c.MQTT.Port, "The port of the mqtt broker.")
	flag.StringVar(&c.MQTT.User, "mqtt-user", c.MQTT.User, "The username to connect to the broker.")
	flag.StringVar(&c.MQTT.Password, "mqtt-password", c.MQTT.Password, "The password to connect to the broker.")

	flag.IntVar(&c.Web.Port, "web-port", c.Web.Port, "The port of the webserver.")
	flag.Parse()
	return c
}
