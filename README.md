# Demo: Grafana Instance Data Engineering
## Goals
1. Create an OSS Grafana Instance
2. Build a lightweight Go App to instrument a metrics endpoint tracking dashboard usage
3. Scrape that endpoint with an OSS Prometheus Instance
4. Remote write local Prometheus data to Grafana Cloud
5. Build basic usage dashboards in Grafana Cloud

## How to run
1. Create an .env file (see .env.example for reference)
2. Run `docker-compose up`