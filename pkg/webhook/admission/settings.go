package admission

import "github.com/akaitux/shell-operator/pkg/webhook/server"

type WebhookSettings struct {
	server.Settings
	CAPath               string
	CABundle             []byte
	ConfigurationName    string
	DefaultFailurePolicy string
}
