package contexthelper

type ContextKey string

const (
	ContextKeyUserAgent ContextKey = "user_agent"
	ContextKeyIPAddress ContextKey = "ip_address"
	ContextKeyRequestId ContextKey = "request_id"
	ContextKeyAppId     ContextKey = "app_id"

	//ContextKeyTokenInfo              ContextKey = "token_info"
	//ContextKeyApplicationType        ContextKey = "application_type"
	//ContextKeyClientIdentifier       ContextKey = "client_identifier"
	//ContextKeyClientOS               ContextKey = "client_os"
	//ContextKeyClientVersion          ContextKey = "client_version"

)
