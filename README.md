# dis-router-nginx-poc

Several POCS for exploring the concept of an NGINX router configured from an external source

## Getting started

### Dependencies

- colima
- nvm

### Local running

#### Openresty

This stack is for running:

- [OpenResty](https://openresty.org/en/) for proxying / redirects
- http-echo as an upstream server
- router-api to fake an API serving redirect configuration

Currently this is configured just for redirects with a starting set held in ./pocs/router-api/data/db.json

To start the stack you can run (from the `./pocs/openresty` directory)

```sh
    make up
```

Or to have it reload on nginx / lua change:

```sh
    make watch
```

The openresty server will load redirects from the router-api. When router-api's "database" is updated it currently fires a webhook at the openresty server, causing it to reload it's configuration. You can do this by sending a POST request to localhost:7777/redirects with a body like so:

```json
    {
        "from": "/origin/carrots",
        "to": "/destination/carrots"
    }
```

Alternatively you can trigger the reload by sending a request to localhost:8080/webhook/redirect-config

Requests to be proxied / redirected can be send to localhost:8080/${path}

This POC does not currently have any facility for wildcard redirecting.

#### Nginx sidecar

This stack is for running:

- nginx for proxying / redirects
- http-echo as an upstream server(s)
- router-api to fake an API serving redirect configuration

This is configured for both redirects and path based routing with a starting set held in ./pocs/router-api/data/db.json

To start the stack you can run (from the `./pocs/nginx-sidecar` directory)

```sh
    make up
```

The sidecar server will load redirects from the router-api. When router-api's "database" is updated it currently fires a webhook at the sidecar server, causing it to write the nginx configurationa and then reload nginx. You can do this by sending a POST request to localhost:7777/redirects with a body like so:

```json
    // redirect
    {
        "from": "/origin/carrots",
        "to": "/destination/carrots"
    }

    // route
    {
        "path": "/sausages",
        "service": "http-echo-server-1:5678"
    }
```

Alternatively you can trigger the reload by sending a request to localhost:5000/update-config

Requests to be proxied / redirected can be send to localhost:8080/${path}

This service will accept wildcard routing inside the `from` directive for redirects or the `path` directive for path based routing.

#### Nginx with retry

This POC runs:

- nginx as proxy
- http-echo server for successful requests
- not-found-server for 404s

If you request `/business` it will be returned not found by the `not-found-server` and then proxy the request through to the http-echo server.

If you request `/economy/*` the request will be returned by the `not-found-server`.

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

## License

Copyright © 2025, Office for National Statistics <https://www.ons.gov.uk>

Released under MIT license, see [LICENSE](LICENSE.md) for details.
