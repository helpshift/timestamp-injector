package timestamp_injector

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

// Config holds plugin configuration.
type Config struct {
    HeaderName string `json:"headerName,omitempty"`
}

// CreateConfig creates default plugin configuration.
func CreateConfig() *Config {
    return &Config{
        HeaderName: "epoch-seconds", // default value
    }
}

type TimestampInjector struct {
    name       string
    next       http.Handler
    headerName string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
    return &TimestampInjector{
        name:       name,
        next:       next,
        headerName: config.HeaderName,
    }, nil
}

func (eh *TimestampInjector) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    sec := float64(time.Now().UnixNano()) / 1e9
    rw.Header().Set(eh.headerName, fmt.Sprintf("%.3f", sec))
    eh.next.ServeHTTP(rw, req)
}
