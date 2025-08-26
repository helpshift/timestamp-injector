package traefik_timestamp_header

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
        HeaderName: "HS-UEpoch", // default value
    }
}

type TimestampHeader struct {
    name       string
    next       http.Handler
    headerName string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
    return &TimestampHeader{
        name:       name,
        next:       next,
        headerName: config.HeaderName,
    }, nil
}

func (eh *TimestampHeader) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    sec := float64(time.Now().UnixNano()) / 1e9
    rw.Header().Set(eh.headerName, fmt.Sprintf("%.3f", sec))
    eh.next.ServeHTTP(rw, req)
}
