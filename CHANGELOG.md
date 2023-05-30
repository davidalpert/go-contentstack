<!-- PROJECT SHIELDS -->
<!--
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![License: GPL v3][license-shield]][license-url]
<!-- [![Issues][issues-shield]][issues-url] -->
<!-- [![Forks][forks-shield]][forks-url] -->
<!-- ![GitHub Contributors][contributors-shield] -->
<!-- ![GitHub Contributors Image][contributors-image-url] -->

<!-- PROJECT LOGO -->
<br />
<!-- vale Google.Headings = NO -->
<h1 align="center"><code>go-contentstack</code></h1>
<!-- vale Google.Headings = YES -->

<p align="center">
  A command-line tool for interacting with the <a href="https://www.contentstack.com/docs/developers/">ContentStack APIs</a>
  <br />
  <a href="./README.md"><strong>README</strong></a>
  ·
  <a href="./CHANGELOG.md">CHANGELOG</a>
  .
  <a href="./CONTRIBUTING.md">CONTRIBUTING</a>
  <br />
  <!-- <a href="https://github.com/davidalpert/go-contentstack">View Demo</a>
  · -->
  <a href="https://github.com/davidalpert/go-contentstack/issues">Report Bug</a>
  ·
  <a href="https://github.com/davidalpert/go-contentstack/issues">Request Feature</a>
</p>
## Changelog


<a name="v0.4.0"></a>
## [v0.4.0] - 2023-05-30
### Features
- manage locales


<a name="v0.3.1"></a>
## [v0.3.1] - 2023-05-21
### Bug Fixes
- Environment.ACL failed to destructure

### Code Refactoring
- rename a poperty


<a name="v0.3.0"></a>
## [v0.3.0] - 2023-05-17
### Build
- **deps:** bump github.com/goreleaser/goreleaser
- **deps:** bump github.com/smartystreets/goconvey from 1.7.2 to 1.8.0
- **deps:** bump golang.org/x/net from 0.5.0 to 0.7.0
- **deps:** bump github.com/spf13/cobra from 1.6.1 to 1.7.0
- **deps:** bump github.com/go-task/task/v3 from 3.20.0 to 3.24.0
- **deps:** bump actions/setup-go from 3 to 4
- **deps:** bump github.com/git-chglog/git-chglog from 0.15.3 to 0.15.4

### Features
- CRUD environments

### Pull Requests
- Merge pull request [#17](https://github.com/davidalpert/go-contentstack/issues/17) from davidalpert/dependabot/go_modules/github.com/goreleaser/goreleaser-1.18.2
- Merge pull request [#13](https://github.com/davidalpert/go-contentstack/issues/13) from davidalpert/dependabot/go_modules/github.com/smartystreets/goconvey-1.8.0
- Merge pull request [#16](https://github.com/davidalpert/go-contentstack/issues/16) from davidalpert/dependabot/go_modules/golang.org/x/net-0.7.0
- Merge pull request [#6](https://github.com/davidalpert/go-contentstack/issues/6) from davidalpert/dependabot/go_modules/github.com/git-chglog/git-chglog-0.15.4
- Merge pull request [#9](https://github.com/davidalpert/go-contentstack/issues/9) from davidalpert/dependabot/github_actions/actions/setup-go-4
- Merge pull request [#14](https://github.com/davidalpert/go-contentstack/issues/14) from davidalpert/dependabot/go_modules/github.com/go-task/task/v3-3.24.0
- Merge pull request [#15](https://github.com/davidalpert/go-contentstack/issues/15) from davidalpert/dependabot/go_modules/github.com/spf13/cobra-1.7.0


<a name="v0.2.1"></a>
## [v0.2.1] - 2023-02-12
### Build
- goreleaser can't call make without a Makefile


<a name="v0.2.0"></a>
## [v0.2.0] - 2023-02-12
### Bug Fixes
- fmt string token was incorrect

### Build
- normalize on go1.18 due to some dependencies
- add old buld flag syntax
- downgrade go.mod to 1.16
- **deps:** bump github.com/ilyakaznacheev/cleanenv from 1.3.0 to 1.4.2
- **deps:** bump goreleaser/goreleaser-action from 3 to 4

### Features
- support the placeholder attribute

### Pull Requests
- Merge pull request [#2](https://github.com/davidalpert/go-contentstack/issues/2) from davidalpert/dependabot/go_modules/github.com/ilyakaznacheev/cleanenv-1.4.2
- Merge pull request [#1](https://github.com/davidalpert/go-contentstack/issues/1) from davidalpert/dependabot/github_actions/goreleaser/goreleaser-action-4
- Merge pull request [#5](https://github.com/davidalpert/go-contentstack/issues/5) from davidalpert/feat-placeholder


<a name="v0.1.0"></a>
## v0.1.0 - 2023-02-03
### Bug Fixes
- release tasks cannot depend on features

### Code Refactoring
- disable CUD ops
- replace generation attempts with hand-bombed client

### Features
- basic CLI tool with version and build pipeline

### Spike
- openapitools/openapi-generator-client


[Unreleased]: https://github.com/davidalpert/go-contentstack/compare/v0.4.0...HEAD
[v0.4.0]: https://github.com/davidalpert/go-contentstack/compare/v0.3.1...v0.4.0
[v0.3.1]: https://github.com/davidalpert/go-contentstack/compare/v0.3.0...v0.3.1
[v0.3.0]: https://github.com/davidalpert/go-contentstack/compare/v0.2.1...v0.3.0
[v0.2.1]: https://github.com/davidalpert/go-contentstack/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/davidalpert/go-contentstack/compare/v0.1.0...v0.2.0
[license-shield]: https://img.shields.io/badge/License-GPLv3-blue.svg
[license-url]: https://www.gnu.org/licenses/gpl-3.0