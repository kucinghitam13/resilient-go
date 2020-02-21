package config

// Configuration holds all configuration for the service
type Configuration struct {
	Server          Server          `yaml:"server"`
	ServiceProducts ServiceProducts `yaml:"service_products"`
}

// Server server configuration
type Server struct {
	Port int `yaml:"port"`
}

// ServiceProducts products-service confiugration
type ServiceProducts struct {
	Host                   string `yaml:"host"`
	URLGetProductsByShopID string `yaml:"url_get_products_by_shop"`
	HTTPTimeout            int    `yaml:"http_timeout"`
	CircuitBreaker CircuitBreaker `yaml:"circuit_breaker"`
}

// CircuitBreaker generic configuration for circuit breaker
type CircuitBreaker struct {
	// Name is the name of the CircuitBreaker
	Name         string `yaml:"name"`
	// MaxRequests is the maximum number of requests allowed to pass through when the CircuitBreaker is half-open.
	// If MaxRequests is 0, CircuitBreaker allows only 1 request.
	MaxRequest   int    `yaml:"max_request"`
	// Interval is the cyclic period of the closed state for CircuitBreaker to clear the internal Counts, described later in this section.
	// If Interval is 0, CircuitBreaker doesn't clear the internal Counts during the closed state.
	Interval     int    `yaml:"interval"`
	// Timeout is the period of the open state, after which the state of CircuitBreaker becomes half-open.
	// If Timeout is 0, the timeout value of CircuitBreaker is set to 60 seconds.
	TimeOut      int    `yaml:"time_out"`
	ConsecutiveFailures int    `yaml:"consecutive_failures"`
}