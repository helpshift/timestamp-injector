package traefik_epoch_header

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

// Config holds plugin configuration (empty here).
type Config struct{}

// CreateConfig creates default plugin configuration.
func CreateConfig() *Config {
    return &Config{}
}

type EpochHeader struct {
    name string
    next http.Handler
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
    return &EpochHeader{
        name: name,
        next: next,
    }, nil
}

func (eh *EpochHeader) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    sec := float64(time.Now().UnixNano()) / 1e9
    rw.Header().Set("HS-UEpoch", fmt.Sprintf("%.3f", sec))
    eh.next.ServeHTTP(rw, req)
}
