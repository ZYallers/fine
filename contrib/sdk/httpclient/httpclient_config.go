package httpclient

import (
	"time"

	"go.uber.org/zap"
)

// Config is the configuration struct for SDK client.
type Config struct {
	Host    string        // Service address.
	Handler Handler       // Custom response handler.
	Logger  *zap.Logger   // Custom logger.
	Timeout time.Duration // Request timeout.
	RawDump bool          // Whether auto dump request&response in stdout.
}
