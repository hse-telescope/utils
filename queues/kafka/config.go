package kafka

// QueueCredentials ...
type QueueCredentials struct {
	URLs  []string `yaml:"urls"`
	Topic string   `yaml:"topic"`
}
