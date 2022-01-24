package domain

type ComponentType string

const (
	ComponentType_NA       = ""
	ComponentType_Postgres = "postgres"
	ComponentType_Kafka    = "kafka"
)

type TestCase struct {
	ComponentType ComponentType     `json:"component-type"`
	Image         string            `json:"image"`
	Port          uint16            `json:"port"`
	EnvVars       map[string]string `json:"env-vars"`
}
