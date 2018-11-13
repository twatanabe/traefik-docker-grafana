### Traefik Docker Grafana

Group of container recepies designed to structure reverse proxy network web apps including OS and container monitoring services. Fully deploying this project allows globally accessable web applications to run with a single command.
- [Traefik](https://traefik.io/) 1.7.4-alpine
- Hello Go test page
- [Grafana](https://grafana.com/grafana)
- [Prometheus](https://prometheus.io/) 2.1.0
- [Google cAdvisor](https://github.com/google/cadvisor)
- [Portainer](https://portainer.io/)

## Pre-requirements
- Docker Machine, Docker Compose
- Cloud host on Azure or Droplet
- Global domain with wild-card routing

## Features
- Reverse Proxy networking empowerd by Traefik with Let's Encrypt HTTPS security.
- Prometheus data management to scape host and container metrics for invoking visualization and alert services.
- Comprehensive visual analytics of network traffik and container status using Grafana.
- Portainer, awesome container management UI.


### Design Overview

![alt text](https://github.com/code-badger/traefik-docker-grafana/blob/master/image.png)

### grafana.domain.com

![alt text](https://github.com/code-badger/traefik-docker-grafana/blob/master/grafana.jpg)
