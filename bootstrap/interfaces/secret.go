package interfaces

import (
	"time"
)

// SecretProvider defines the contract for secret provider implementations that
// allow secrets to be retrieved/stored from/to a services Secret Store.
type SecretProvider interface {
	// StoreSecret stores new secrets into the service's SecretStore at the specified secretName.
	StoreSecret(secretName string, secrets map[string]string) error

	// GetSecret retrieves secrets from the service's SecretStore at the specified secretName.
	GetSecret(secretName string, keys ...string) (map[string]string, error)

	// SecretsUpdated sets the secrets last updated time to current time.
	SecretsUpdated()

	// SecretsLastUpdated returns the last time secrets were updated
	SecretsLastUpdated() time.Time

	// GetAccessToken return an access token for the specified token type and service key.
	// Service key is use as the access token role which must have be previously setup.
	GetAccessToken(tokenType string, serviceKey string) (string, error)

	// ListSecretsecretNames returns a list of paths for the current service from an insecure/secure secret store.
	ListSecretPaths() ([]string, error)

	// HasSecret returns true if the service's SecretStore contains a secret at the specified secretName.
	HasSecret(secretName string) (bool, error)

	// RegisteredSecretUpdatedCallback registers a callback for a secret.
	RegisteredSecretUpdatedCallback(secretName string, callback func(secretName string)) error

	// SecretUpdatedAtPath performs updates and callbacks for an updated secret or secretName.
	SecretUpdatedAtPath(secretName string)

	// DeregisterSecretUpdatedCallback removes a secret's registered callback secretName.
	DeregisterSecretUpdatedCallback(secretName string)

	// GetMetricsToRegister returns all metric objects that needs to be registered.
	GetMetricsToRegister() map[string]interface{}
}
