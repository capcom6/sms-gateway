package handlers

type GatewayMode string

const (
	GatewayModePrivate GatewayMode = "private"
	GatewayModePublic  GatewayMode = "public"
)

type Config struct {
	GatewayMode GatewayMode
}
