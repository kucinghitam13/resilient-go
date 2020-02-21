# Resilient Go Playground
this repository act as playground and example of desired service resiliency as described in Resilient Service Roadmap document.

## Expectation
This is bare minimum sample project to show expected resiliency implementation in microservices.


## Component Explanation
* `ext-service` -> act as external service that we call for specific information from
* `idp-service` -> act as independent service / our own service that we want it to be resulient

NOTES:
* Supposedly shops and products shops are in different database but I'm to lazy for that now.
* I'm skimming proper Clean Architecture on actual Database package for saving up time

### TODO
[x] Have Database Schema File for Products (products table)
[x] Have Database Schema File for Shops (shops table)
[x] Seeding File for Shops Database
[x] Seeding File for Products Database
[x] Simple HTTP Server and Ping Endpoint for Products Service
[x] Simple HTTP Server and Ping Endpoint for Shops Service
[x] Dockerify Products Service
[x] Dockerify Shops Service
[x] 2 Simple Repo Method for Products Service
[x] 2 Simple Repo Method for Shops Service
[x] 2 Simple Usecase Method for Products Service using existing Repo Method
[x] 2 Simple Usecase Method for Shops Service using existing Repo Method
[] Setup Service Discovery with Consul on Docker Compose
[x] API Endpoint `/api/v1/shops/{:shop_id}/products` for Products Service
[x] API Endpoint `/api/v1/shops/{:shop_id}/products` for Shops Service
[x] Implement CircuitBreaker for Shops Service on Call to Products Service
[x] Implement CircuitBreaker for Products Service on Call to Database