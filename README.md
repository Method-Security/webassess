<div align="center">
<h1>webassess</h1>

[![GitHub Release][release-img]][release]
[![Verify][verify-img]][verify]
[![Go Report Card][go-report-img]][go-report]
[![License: Apache-2.0][license-img]][license]

[![GitHub Downloads][github-downloads-img]][release]
[![Docker Pulls][docker-pulls-img]][docker-pull]

</div>
webassess is designed as CLI application for performing AI based analysis of security resources at the edge,

The types of assessments that webassess can conduct are constantly growing. For the most up to date listing, please see the documentation [here](./docs/index.md)

To learn more about webassess, please see the [Documentation site](https://method-security.github.io/webassess/) for the most detailed information.

## Quick Start

### Get webassess

For the full list of available installation options, please see the [Installation](./getting-started/installation.md) page. For convenience, here are some of the most commonly used options:

- `docker run methodsecurity/webassess`
- `docker run ghcr.io/method-security/webassess`
- Download the latest binary from the [Github Releases](https://github.com/Method-Security/webassess/releases/latest) page
- [Installation documentation](./getting-started/installation.md)

### Examples

```bash
webassess url --target https://example.com
```

### Building a Statically Compiled Container for Local Testing
(Reference reusable-build.yaml)

1. Build ARM64 builder image: `docker buildx build . --platform linux/arm64 --load --tag armbuilder -f Dockerfile.builder`

2. Build ARM64 image: `docker run -v .:/app/webassess -e GOARCH=arm64 -e GOOS=linux --rm armbuilder goreleaser build --single-target -f .goreleaser/goreleaser-build.yml --snapshot --clean`

3. `cp dist/linux_arm64/build-linux_linux_arm64/webassess .`

4. `docker buildx build . --platform linux/arm64 --load --tag webassess:local -f Dockerfile`

5. Open shell: `docker run -it --rm --entrypoint /bin/sh webassess:testing`

6. OR run command without shell example: `docker run webassess:local app enumerate graphql --target https://countries.trevorblades.com/ -o json`


### Note:
This tool runs on a headless-shell base image to support chrome/chromium browser automation. The dockerfile uses debian-based install tools. 

## Contributing

Interested in contributing to webassess? Please see our organization wide [Contribution](https://method-security.github.io/community/contribute/discussions.html) page.

## Want More?

If you're looking for an easy way to tie webassess into your broader cybersecurity workflows, or want to leverage some autonomy to improve your overall security posture, you'll love the broader Method Platform.

For more information, visit us [here](https://method.security)

## Community

webassess is a Method Security open source project.

Learn more about Method's open source source work by checking out our other projects [here](https://github.com/Method-Security) or our organization wide documentation [here](https://method-security.github.io).

Have an idea for a Tool to contribute? Open a Discussion [here](https://github.com/Method-Security/Method-Security.github.io/discussions).

[verify]: https://github.com/Method-Security/webassess/actions/workflows/verify.yml
[verify-img]: https://github.com/Method-Security/webassess/actions/workflows/verify.yml/badge.svg
[go-report]: https://goreportcard.com/report/github.com/Method-Security/webassess
[go-report-img]: https://goreportcard.com/badge/github.com/Method-Security/webassess
[release]: https://github.com/Method-Security/webassess/releases
[releases]: https://github.com/Method-Security/webassess/releases/latest
[release-img]: https://img.shields.io/github/release/Method-Security/webassess.svg?logo=github
[github-downloads-img]: https://img.shields.io/github/downloads/Method-Security/webassess/total?logo=github
[docker-pulls-img]: https://img.shields.io/docker/pulls/methodsecurity/webassess?logo=docker&label=docker%20pulls%20%2F%20webassess
[docker-pull]: https://hub.docker.com/r/methodsecurity/webassess
[license]: https://github.com/Method-Security/webassess/blob/main/LICENSE
[license-img]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
