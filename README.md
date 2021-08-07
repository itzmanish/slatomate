# Slatomate Service

This is the Slatomate service

Generated with

```
micro new --namespace=github.itzmanish --type=service slatomate
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: github.itzmanish.service.slatomate
- Type: service
- Alias: slatomate

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service

```
./slatomate-service
```

Build a docker image

```
make docker
```
