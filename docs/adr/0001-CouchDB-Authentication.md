CouchDB Authentication from frontend
===

* Status: accepted
* Deciders: @timo-reymann
* Date: 2021-09-24

## Context and Problem Statement

As we want to keep backend infrastructure as minimal as possible to make self-hosting easy we need a way to authenticate
users from the webapp in an efficient, easy to understand and secure way.

## Decision Drivers <!-- optional -->

* low chance of misconfiguration
* easy to understand
* secure

## Considered Options

* JWT Authentication
* Cookie Authentication
* Proxy Authentication (with api gateway in front)

## Decision Outcome

Chosen option: "Proxy Authentication",
because it's hard to configure JWT on couchdb directly and requires tweaking the oauth2 settings too much which makes it hard to setup.

### Positive Consequences

* we wanted to create a API-Gateway anyway (funnel) so this is the single point of service communication
* less configuration required by the user

### Negative Consequences

* we are not leveraging the full power of CouchDB
* CouchDB needs to be configured for proxy auth

## Pros and Cons of the Options

### JWT Authentication

Use builtin JWT authentication of CouchDB

* Good, because we are leveraging the full power of CouchDB
* Good, because we are using OAuth2 anyway
* Bad, because configuration is hard and may require tweaking on oauth2 provider side
* Bad, because we need to maintain OAuth2 configuration in multiple places
* Bad, because in case of misconfiguration on couchdb site you get a cryptic erlang stacktrace

### Cookie Authentication

Cookie authentication relies on a session cookie to be created that is valid for a given amount of time.

* Good, because it's very easy to use
* Bad, because we need to build a translation layer for JWT <-> Cookie
* Bad, because cookies don't feel right

### Proxy Authentication (with api gateway in front)

Proxy authentication relies on custom upstream headers and is very easy to use and does not require sessions etc.

* Good, because its stateless
* Good, because the implementation is easy and stable
* Good, because we can use given users without building custom authentication solutions
* Good, because its stateless and therefore there is no state management required
* Bad, because we need to build a translation layer JWT <-> Headers

## Links

* [CouchDB Documentation for Authentication](https://docs.couchdb.org/en/latest/api/server/authn.html)