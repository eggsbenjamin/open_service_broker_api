## Open Service Broker API

Implementation of the Open Service Broker API

### API

GET `/<version>/catalog`

Gets all items from the service catalog

PUT /<version>/service_instances/:instance_id

Creates an instance of an item in the service catalog

### Usage

#### Dependencies

- go
- dep
- make
- postgresql w/psql client
- docker

To set up the dev environment run `make setup dev`

This will create and set up the database and run the service.

To build a docker image of the service run `make docker_build`

### Test

To run all tests (unit, integration, service) run `make test`

### Client

A client pkg is available `./client`

### Design Decisions

I chose to use a relational database as a backend after reading the Open Service Broker API spec and seeing
that they have defined various entities and relationships between them. An RDBMS seemed like the right tool
for the job. I chose postgresql because I'm familiar with it and also because of it's JSON related features.
The Open Service Broker API, due to it's nature, can't define strong types for everything and some fields are
dynamic e.g. service instance context which can vary between services. Postgres gives the benefit of referential
integrity between relations and also the ability to store dynamic data along side them if needed.

I chose to architect the service in the classic 'three tier architecture' as it seemed to a be a suitable pattern
for the service. Each of the layers has a single responsibility.

#### Persistence

The persistence layer is abstracted away by the repositories which means that it doesn't leak
into other layers of the application and afford swapping out persistence mechanism if needed. The middle layer has

#### Business Logic

The business logic layer deals with exactly that. The other layers give it easy access to the data that it needs to
execute any business logic that it may be reuired to do.

#### Presentation

The presentation layer accepts incoming requests can perform validation and send meaningful responses to a client.
It can also route to different versions/compositions of sub components based on the version specified.

### Dependency Injection

I chose to use the dependency injection pattern to make the layers loosely coupled but cohesive and easily testable and composable.

