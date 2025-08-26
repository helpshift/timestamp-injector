# Traefik Plugin - Timestamp Header

A Traefik middleware plugin that adds an `HS-UEpoch` header to HTTP responses, containing the current Unix time in `seconds.milliseconds` format (similar to Nginx `$msec`).

## Static Configuration

Enable experimental plugins and specify this plugin:

```yaml
experimental:
  plugins:
    timestampheader:
      moduleName: github.com/helpshift/timestamp-injector
      version: v1.0.1
```

## Dynamic Configuration

Attach the middleware to a router:

```yaml
http:
  middlewares:
    add-timestamp-header:
      plugin:
        timestampheader: {}

  routers:
    my-router:
      rule: "Host(`example.com`)"
      service: my-service
      middlewares:
        - add-timestamp-header
```

## How It Works

This plugin injects an HS-UEpoch header into each response with the current Unix time (seconds.milliseconds).
