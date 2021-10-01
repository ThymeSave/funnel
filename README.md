funnel
===
[![pre-commit](https://img.shields.io/badge/%E2%9A%93%20%20pre--commit-enabled-success)](https://pre-commit.com/)
[![GitHub Release](https://img.shields.io/github/v/tag/thymesave/funnel.svg?label=version)](https://github.com/thymesave/funnel/releases)
[![CircleCI](https://circleci.com/gh/ThymeSave/funnel/tree/main.svg?style=shield)](https://circleci.com/gh/ThymeSave/funnel/tree/main)

> ⚠️ This project is currently under active development

Funnel is a core part of ThymeSave and does what the name suggests: It filters all backend requests. It includes
routing, cors and authentication.

## Development

This project is written in go, uses make as a simple build tool and [pack](https://github.com/buildpacks/pack) for
creating oci compliant images, that can be executed with docker/podman.

### Required tools

- [GNU make](https://www.gnu.org/software/make/)
- [Go 1.16+](https://golang.org/)
- [pack](https://github.com/buildpacks/pack) (required only when building the docker image)

### Goals

For a list of available build and test goals run `make help` or check the [Makefile](./Makefile) manually.
