# adsb-pipeline

A Go pipeline that ingests live ADS-B aircraft data from my own RTL-SDR
receiver. dump1090 decodes the 1090 MHz signals, and the pipeline connects to
its BaseStation (SBS) port to read and process the messages in real time.
Later, the messages will be streamed to a Kafka broker for analytics and
dashboards.

## Requirements

- Golang

## Build and run

```bash
go build
./adsb-pipeline
```

## Roadmap

- [x] SBS client: connect to dump1090 and read messages
- [ ] Parse SBS messages
- [ ] Reconnect with backoff on connection loss
- [ ] Publish messages to Kafka
- [ ] Multi-receiver support: several RTL-SDR stations feeding one Kafka cluster
  to extend coverage, similar to how flight-tracking websites do
- [ ] Dashboard: live aircraft positions, message rates and receiver coverage