<!-- @format -->

# File Cache Service

This repository provides a simple file caching service using a Redis-compatible backend

## Configuration

Configuration is managed via environment variables or configuration files:

-   `REDIS_URL`: Connection string for the Redis backend. (e.g. `redis://localhost:6379/0`)
-   `PUBLIC_HOSTNAME`: The hostname clients use to access cached files. (default: `http://localhost:8080`)
-   `LISTENER`: The address and port the service listens on (default: `:8080`).
-   `CACHE_TTL`: Default time-to-live for cached files using go duration syntaz (e.g., `60s`, `10min`, `1h`).

## Running with Docker

You can use the `ghcr.io/fedotcompot/file-cacher` or build using provided `Dockerfile` for containerized deployment.


```sh
docker run -e REDIS_URL=redis://host:port -e PUBLIC_HOSTNAME=your.host -e LISTEN_ADDR=":8080" -e CACHE_TTL=1h ghcr.io/fedotcompot/file-cacher
```

## Usage

Post content to `/api/v1/upload`, providing:
- path
- content to cache
    - content type
    - base64 encoded data
- override ttl (optional)

### Example

Request:
```json
{
    "path": "/hello/world",
    "ttl_override": "60min",
    "data": {
        "content_type": "text/text",
        "data": "SGVsbG9Xb3JsZCE="
    }
}
```

Response:
```jsonc
{
    "url":"http://localhost:8080/hello/world",
    "exp":"2025-06-30T12:45:48.499591239Z" //60min from now
}

```

Getting file:
```sh
$ curl -v http://localhost:8080/hello/world
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /hello/world HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.88.1
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: text/text
< Vary: Accept-Encoding
< Date: Mon, 30 Jun 2025 12:46:42 GMT
< Content-Length: 11
< 
* Connection #0 to host localhost left intact
HelloWorld!
```

## Developing

    This repository provides a devocontainer configuration. Alternatively you can use `compose.yaml`:

## API

The service exposes endpoints for uploading and retrieving files. See the code in `internal/web/router.go` for details.
