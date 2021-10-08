funnel
===
[![License: GPL v2](https://img.shields.io/badge/License-GPL%20v2-blue.svg)](https://www.gnu.org/licenses/old-licenses/gpl-2.0.en.html)
[![pre-commit](https://img.shields.io/badge/%E2%9A%93%20%20pre--commit-enabled-success)](https://pre-commit.com/)
[![Go Report Card](https://goreportcard.com/badge/github.com/thymesave/funnel)](https://goreportcard.com/report/github.com/thymesave/funnel)
[![CircleCI](https://circleci.com/gh/ThymeSave/funnel/tree/main.svg?style=shield)](https://circleci.com/gh/ThymeSave/funnel/tree/main)
[![codecov](https://codecov.io/gh/thymesave/funnel/branch/main/graph/badge.svg?token=J9DAZXRUDZ)](https://codecov.io/gh/thymesave/funnel)
[![GitHub Release](https://img.shields.io/github/v/tag/thymesave/funnel.svg?label=version)](https://github.com/thymesave/funnel/releases)

> ⚠️ This project is currently under active development

Funnel is a core part of ThymeSave and does what the name suggests: It filters all backend requests. It includes
routing, cors and authentication.

## Configuration

Configuration is done entirely via environment variables.

| Variable                    | Default   | Description
| :-------------------------- | :-------: | :----------
| FUNNEL_OAUTH2_ISSUER_URL    | N/A       | Base url of the identity provider, will be used to search for endpoint `/.well-known/openid-configuration`
| FUNNEL_OAUTH2_CLIENT_ID     | N/A       | ClientId to check for
| FUNNEL_OAUTH2_VERIFY_ISSUER | true      | Should the issuer of jwts should be verified (true) or not (false)
| FUNNEL_PORT                 | 3000      | Port funnel will be listening on
| FUNNEL_CORS_ORIGINS         | *         | Origins to allow in `Access-Control-Allow-Origins` header, this is typically the url of the webapp.
| FUNNEL_COUCHDB_SCHEME       | http      | Scheme to use for couchdb communication in most cases http or https.
| FUNNEL_COUCHDB_HOST         | 127.0.0.1 | IP or hostname where the couchdb is running
| FUNNEL_COUCHDB_ADMIN_USER   | admin     | User to use for administrative actions with proxy authentication

> More information about setting up OAuth2 etc. will be added later in a different location
> and linked here.

## Development

This project is written in go, uses make as a simple build tool and [pack](https://github.com/buildpacks/pack) for
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
