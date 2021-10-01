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

## Commit Message Convention

This repository follows [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)

### Format

`<type>(optional scope): <description>`
Example: `feat(pre-event): add speakers section`

### 1. Type

Available types are:

- feat → Changes about addition or removal of a feature. Ex: `feat: add table on landing page`
  , `feat: remove table from landing page`
- fix → Bug fixing, followed by the bug. Ex: `fix: illustration overflows in mobile view`
- docs → Update documentation (README.md)
- style → Updating style, and not changing any logic in the code (reorder imports, fix whitespace, remove comments)
- chore → Installing new dependencies, or bumping deps
- refactor → Changes in code, same output, but different approach
- ci → Update github workflows, husky
- test → Update testing suite, cypress files
- revert → when reverting commits
- perf → Fixing something regarding performance (deriving state, using memo, callback)
- vercel → Blank commit to trigger vercel deployment. Ex: `vercel: trigger deployment`

### 2. Optional Scope

Labels per page Ex: `feat(pre-event): add date label`

*If there is no scope needed, you don't need to write it*

### 3. Description

Description must fully explain what is being done.

Add BREAKING CHANGE in the description if there is a significant change.

**If there are multiple changes, then commit one by one**

- After colon, there are a single space Ex: `feat: add something`
- When using `fix` type, state the issue Ex: `fix: file size limiter not working`
- Use imperative, dan present tense: "change" not "changed" or "changes"
- Don't use capitals in front of the sentence
- Don't add full stop (.) at the end of the sentence
