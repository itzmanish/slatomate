# Slatomate Service

This is the Slatomate service

- [API Doc](https://heymanish.stoplight.io/docs/slatomate/)

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

# Benchmark

```bash
😊 ghz --proto=proto/slatomate/slatomate.proto --import-paths=/home/manish/.local/include --insecure --call=github.itzmanish.service.slatomate.Slatomate.Me  -c 100 -n 10000 -d '{"api_key":"apikey"}' 127.0.0.1:39215

Summary:
  Count:        10000
  Total:        15.36 s
  Slowest:      1.01 s
  Fastest:      1.28 ms
  Average:      152.24 ms
  Requests/sec: 650.84

Response time histogram:
  1.280    [1]    |
  101.665  [6761] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  202.050  [6]    |
  302.435  [25]   |
  402.820  [1057] |∎∎∎∎∎∎
  503.205  [1779] |∎∎∎∎∎∎∎∎∎∎∎
  603.590  [41]   |
  703.975  [2]    |
  804.360  [132]  |∎
  904.745  [190]  |∎
  1005.129 [6]    |

Latency distribution:
  10 % in 2.73 ms
  25 % in 4.23 ms
  50 % in 10.20 ms
  75 % in 392.79 ms
  90 % in 433.02 ms
  95 % in 464.75 ms
  99 % in 836.59 ms

Status code distribution:
  [OK]   10000 responses
  
```
