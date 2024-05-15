# URL Shortener

## Table of contents

- [URL Shortener](#url-shortener)
  - [Table of contents](#table-of-contents)
  - [Getting started](#getting-started)
    - [Download](#download)
    - [Start application in development mode](#start-application-in-development-mode)
    - [Start application in production mode](#start-application-in-production-mode)
  - [Authors](#authors)
  - [License](#license)

## Getting started

### Download

To download this project, please do: `git clone https://github.com/tun43p/url-shortener.git`.

### Start application in development mode

To start the application, please do:

```bash
API_HOST="localhost:8080" API_KEY="secret" API_DATABASE="urls.db" GIN_MODE="debug" go run cmd/app/main.go
```

### Start application in production mode

To start the application, please do:

```bash
API_HOST="localhost:8080" API_KEY="secret "API_DATABASE="urls.db" GIN_MODE="release" go cmd/app/main.go
```

## Authors

- **tun43p** - _Initial work_ - [tun43p](https://github.com/tun43p).

## License

This project is licensed under the MIT License, see the [LICENSE](LICENSE) file for details.
