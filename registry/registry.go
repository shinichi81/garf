package registry

import (
	"github.com/backenderia/garf/server"
	"github.com/spf13/viper"
)

// Handler is the package exported API
type Handler interface {
	Register(Bundle)
	Server() server.Support
	Set(key string, value interface{})
	Get(key string)
	Configure()
}

// Bundle is the interface for bundles
type Bundle interface {
	Init(map[string]interface{})
	Register(server.Handler)
}

type registry struct {
	config  *viper.Viper
	server  server.Support
	bundles []Bundle
}

// New creates registry instance
func New(server server.Support) Handler {
	c := viper.New()
	r := &registry{
		config: c,
	}
	r.server = server
	return r
}

// Register receives a Route and send it the mux.Router instance
func (r *registry) Register(bundle Bundle) {
	bundle.Register(r.server)
	r.bundles = append(r.bundles, bundle)
}

// Server returns the Server instance
func (r *registry) Server() server.Support {
	return r.server
}

// Set a global property for this registry
func (r *registry) Set(key string, value interface{}) {
	r.config.Set(key, value)
}

// Get a global property for this registry
func (r *registry) Get(key string) {
	r.config.Get(key)
}

// Configure all routes
func (r *registry) Configure() {
	for _, bd := range r.bundles {
		bd.Init(r.config.AllSettings())
	}
	r.server.Configure()
}
