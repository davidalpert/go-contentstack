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


[Unreleased]: https://github.com/davidalpert/go-contentstack/compare/v0.2.0...HEAD
[v0.2.0]: https://github.com/davidalpert/go-contentstack/compare/v0.1.0...v0.2.0
[license-shield]: https://img.shields.io/badge/License-GPLv3-blue.svg
[license-url]: https://www.gnu.org/licenses/gpl-3.0