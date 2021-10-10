funnel
===
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/old-licenses/gpl-2.0.en.html)
[![CircleCI](https://circleci.com/gh/ThymeSave/funnel/tree/main.svg?style=shield)](https://circleci.com/gh/ThymeSave/funnel/tree/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/thymesave/funnel)](https://goreportcard.com/report/github.com/thymesave/funnel)
[![codecov](https://codecov.io/gh/thymesave/funnel/branch/main/graph/badge.svg?token=J9DAZXRUDZ)](https://codecov.io/gh/thymesave/funnel)
[![GitHub Release](https://img.shields.io/github/v/tag/thymesave/funnel.svg?label=version)](https://github.com/thymesave/funnel/releases)
[![pre-commit](https://img.shields.io/badge/%E2%9A%93%20%20pre--commit-enabled-success)](https://pre-commit.com/)
[![Dependabot](https://badgen.net/badge/Dependabot/enabled/green?icon=dependabot)](https://dependabot.com/)

Funnel is a core part of ThymeSave and does what the name suggests: It filters all backend requests. It includes
routing, cors and authentication.

## Installation

Funnel is available as binary, via `go get` or docker.

> A detailed guide and templates for setting up infrastructure will follow and be linked here,
> for now these instructions serve as an absolute basic guide.

### Docker (recommended)

```sh
docker run --rm ghcr.io/thymesave/funnel -e CONFIG_KEY=value [...]
```

### Using go

```sh
go get github.com/thymesave/funnel
go install github.com/thymesave/funnel

# Execute or setup systemd with cmd
funnel
```

### Using binary

Download the desired release from [the releases page](https://github.com/ThymeSave/funnel/releases).

## Configuration

Configuration is done entirely via environment variables.

| Variable                      | Default               | Description
| :---------------------------- | :-------------------: | :----------
| FUNNEL_OAUTH2_ISSUER_URL      | N/A                   | Base url of the identity provider, will be used to search for endpoint `/.well-known/openid-configuration`
| FUNNEL_OAUTH2_CLIENT_ID       | N/A                   | ClientId to check for
| FUNNEL_OAUTH2_VERIFY_ISSUER   | true                  | Should the issuer of JWTs should be verified (true) or not (false)
| FUNNEL_OAUTH2_USERNAME_CLAIM  | email                 | Claim in JWTs that will be used to uniquely identify the user
| FUNNEL_OAUTH2_SCOPES          | openid,profile,email  | Scopes to request when authorizing the user in the webapp
| FUNNEL_PORT                   | 3000                  | Port funnel will be listening on
| FUNNEL_CORS_ORIGINS           | *                     | Origins to allow in `Access-Control-Allow-Origins` header, this is typically the url of the webapp.
| FUNNEL_COUCHDB_SCHEME         | http                  | Scheme to use for couchdb communication in most cases http or https.
| FUNNEL_COUCHDB_HOST           | 127.0.0.1             | IP or hostname where the couchdb is running
| FUNNEL_COUCHDB_ADMIN_USER     | admin                 | User to use for administrative actions with proxy authentication

> More information about setting up OAuth2 etc. will be added later in a different location
> and linked here.

## Development

This project is written in go, uses [make](https://www.gnu.org/software/make/) as a simple build tool and [pack](https://github.com/buildpacks/pack) for
creating oci compliant images, that can be executed with docker/podman.

### Required tools

- [GNU make](https://www.gnu.org/software/make/)
- [Go 1.16+](https://golang.org/)
- [pack](https://github.com/buildpacks/pack) (required only when building the docker image)
- [pre-commit](https://pre-commit.com/)

### Setup

To set up the project locally:

1. Activate pre-commit hooks: `pre-commit install && pre-commit install --hook-type commit-msg`
1. Install go dependencies: `go mod tidy`
1. Verify your setup with running tests: `make test`

### Goals

For a list of available build and test goals run `make help` or check the [Makefile](./Makefile) manually.

## Commit Message Convention

This repository follows [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)

### Format

`<type>(optional scope): <description>`
Example: `feat(pre-event): Add speakers section`

### 1. Type

Available types are:

- feat → Changes about addition or removal of a feature. Ex: `feat: Add table on landing page`
  , `feat: Remove table from landing page`
- fix → Bug fixing, followed by the bug. Ex: `fix: Illustration overflows in mobile view`
- docs → Update documentation (README.md)
- style → Updating style, and not changing any logic in the code (reorder imports, fix whitespace, remove comments)
- chore → Installing new dependencies, or bumping deps
- refactor → Changes in code, same output, but different approach
- ci → Update github workflows, husky
- test → Update testing suite, cypress files
- revert → when reverting commits
- perf → Fixing something regarding performance (deriving state, using memo, callback)
- vercel → Blank commit to trigger vercel deployment. Ex: `vercel: Trigger deployment`

### 2. Optional Scope

Labels per page Ex: `feat(pre-event): Add date label`

*If there is no scope needed, you don't need to write it*

### 3. Description

Description must fully explain what is being done.

Add BREAKING CHANGE in the description if there is a significant change.

**If there are multiple changes, then commit one by one**

- After colon, there are a single space Ex: `feat: Add something`
- When using `fix` type, state the issue Ex: `fix: File size limiter not working`
- Use imperative, dan present tense: "change" not "changed" or "changes"
- Use capitals in front of the sentence
- Don't add full stop (.) at the end of the sentence

## Contributing

## [Code of Conduct](./CODE-OF-CONDUCT.md)

ThymeSave has adopted a Code of Conduct that we expect project participants to adhere to. Please read the full text so
that you can understand what actions will and will not be tolerated.

## [Contributing Guide](./CONTRIBUTING.md)

Read our contributing guide to learn about how to propose bugfixes and improvements and contribute to ThymeSave!
