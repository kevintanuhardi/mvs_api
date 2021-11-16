<h1 align="center">  brook </h1> <br>

<p align="center">
  wpff-backend
</p>


[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)
[![Build status](https://build.appcenter.ms/v0.1/apps/e7a9f723-c5ba-48f9-a36c-ad6666a34785/branches/master/badge)](https://appcenter.ms)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

brook project for starting your backend service. Built with Go

<img src="./screenshot.png?raw=true" alt="Screenshot" width="256" />

## Menu

  - [Features](#features)
  - [Install](#install)
  - [Quick Start](#quick-start)
  - [Detailed Guide](#detailed-guide)
  - [APIs](#apis)
  - [Roadmap](#roadmap)
  - [FAQ](#faq)
  - [Contributing](#contributing)
  - [Contributors](#contributors)

## Features

This brook also include this library.

- **Database:** [sqlx](https://github.com/jmoiron/sqlx) or [gorm](https://github.com/go-gorm/gorm)
- **Log:** [zap](https://github.com/uber-go/zap) or [zerolog](https://github.com/rs/zerolog)
- **Tracing:** [opentracing](https://github.com/opentracing/opentracing-go) 
- **Metrics:** [prometheus](https://github.com/prometheus/client_golang)
- **Mocking:** [mock](https://github.com/golang/mock)
- **Migrate:** [goose](https://github.com/pressly/goose)
- **Router:** [chi](https://github.com/go-chi/chi)
- **Grpc:** [grpc-go](https://github.com/grpc/grpc-go)

### Observability Tools

This Observability tools

- [grafana](https://grafana.com/)
- [jaeger](https://www.jaegertracing.io/)

## Prerequisite

- Golang 1.16.x or later
- Docker
## Install

...coming soon

## Quick Start

### Create Development Environment

Create dev env means we run [docker-compose](./docker-compose.yml).
You can adjust services with your project stack. (eg. add elasticsearch, redis, etc)

First time only :
```bash
  make start-dev
  make migrate-up
  make seed
```

Every time you develop :
```bash
  make start-dev
```

```bash
  make check-dev
```


### Unit Testing

to do simple test simply run 
```
make unit-test
```
to do test using sonarqube please read this [https://gitlab.warungpintar.co/sales-platform/brook/-/wikis/How-to-do-testing-using-Sonarqube-on-local-environment-(Linux)](wiki) 

### Migrate data to development database

Make sure the development environment is ready, then

```bash
  make migrate-up
```

Seed initial data

```bash
  make seed
```

### Stop Development Environment

```bash
  make stop-dev
```

### Runing Service

### REST

1. Local

```
  make rest
```

2. Docker

For docker please make sure, mysql dsn in [.brook.yml](./.brook.yml).
Change port to 3306 and address to `db`.

```
make docker-rest
```
## Detailed Guide

...coming soon

## APIs

...coming soon

## FAQ

...coming soon

## Roadmap

* [x] Phase 1: Init
* [ ] Phase 2: Deployment & Development Standarization
* [ ] Phase 3: Release Train improvements
* [ ] Phase 4: DX improvements
* [ ] Phase 5: Stabilize
* [ ] Phase 6: Road to Open Source [TBD]

## Contributing

Feel like contributing? That's awesome! We have a
[contributing guide](./CONTRIBUTING.md) to help guide you.


