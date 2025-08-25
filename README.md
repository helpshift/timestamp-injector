# Traefik Plugin - Epoch Header

A Traefik middleware plugin that adds an `HS-UEpoch` header to HTTP responses, containing the current Unix time in `seconds.milliseconds` format (similar to Nginx `$msec`).

## Configuration

```yaml
experimental:
  plugins:
    epochheader:
      moduleName: "github.com/durvesh-palkar/traefik-epoch-header"
      version: "v1.0.1"
```

## Static Configuration

Enable experimental plugins and specify this plugin:

```yaml
experimental:
  plugins:
    epochheader:
      moduleName: github.com/helpshift/timestamp-injector
      version: v1.0.1
```

## Dynamic Configuration

Attach the middleware to a router:

```yaml
http:
  middlewares:
    add-epoch-header:
      plugin:
        epochheader: {}

  routers:
    my-router:
      rule: "Host(`example.com`)"
      service: my-service
      middlewares:
        - add-epoch-header
```

## How It Works

This plugin injects an HS-UEpoch header into each response with the current Unix time (seconds.milliseconds).

