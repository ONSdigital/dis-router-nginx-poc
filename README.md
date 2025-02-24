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

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

## License

Copyright © 2025, Office for National Statistics <https://www.ons.gov.uk>

Released under MIT license, see [LICENSE](LICENSE.md) for details.
