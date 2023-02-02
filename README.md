<!-- PROJECT SHIELDS -->
<!--
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
<!-- vale Google.Acronyms = NO -->
[![License: GPL v3][license-shield]][license-url]
<!-- vale Google.Acronyms = YES -->

<!-- [![Issues][issues-shield]][issues-url] -->
<!-- [![Forks][forks-shield]][forks-url] -->
<!-- ![GitHub Contributors][contributors-shield] -->
<!-- ![GitHub Contributors Image][contributors-image-url] -->

<!-- PROJECT LOGO -->
<br />
<!-- vale Google.Headings = NO -->
<h1 align="center">go-contentstack</h1>
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
  <a href="https://github.com/davidalpert/go-contentstack/issues">Report Bug</a>
  ·
  <a href="https://github.com/davidalpert/go-contentstack/issues">Request Feature</a>
</p>

<details open="open">
  <summary><h2 style="display: inline-block">Table of contents</h2></summary>

- [About the project](#about-the-project)
  - [Why port the nodejs version to Golang?](#why-port-the-nodejs-version-to-golang)
  - [Built with](#built-with)
- [Getting started](#getting-started)
  - [Install](#install)
    - [`go install`](#go-install)
    - [Pre-compiled binaries](#pre-compiled-binaries)
  - [Verify your installation](#verify-your-installation)
  - [Uninstall](#uninstall)
- [Usage](#usage)
- [Troubleshooting](#troubleshooting)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

</details>

<!-- ABOUT THE PROJECT -->
## About the project

### Why port the nodejs version to Golang?

People working with nodejs commonly use a version manager like [`nodenv`](https://github.com/nodenv/nodenv), [`nvm`](https://github.com/nvm-sh/nvm), or [`asdf`](https://asdf-vm.com/) to manage several versions of nodejs side-by-side.

These tools install global packages per node version which means you have to install the node `csdx` plugin once per node version.

In contrast Golang offers the ability to build source code into single-file executables which truly install globally, independent of any versioning tools.

A Golang version of `go-contentstack` simplifies the install and update story making this plugin more manageable.

### Built with

* [Golang 1.18](https://golang.org/)
* [go-releaser](https://goreleaser.com/)

<!-- GETTING STARTED -->
## Getting started

### Install

#### `go install`

With a working golang installation at version >= 1.18 you can install or update with:

```
go install github.com/davidalpert/go-contentstack/cmd/cs@latest
```

#### Pre-compiled binaries

Visit the [Releases](https://github.com/davidalpert/go-contentstack/releases) page to find binary packages pre-compiled for a variety of `GOOS` and `GOARCH` combinations:
1. Download an appropriate package for your `GOOS` and `GOARCH`;
1. Unzip it and put the binary in your path;

### Verify your installation

1. Check the version installed:
    ```
    cs version
    ```

### Uninstall

- `go-contentstack` ships with an `uninstall` sub-command which cleans up and removes itself:

    ```
    cs uninstall
    ```

<!-- USAGE EXAMPLES -->
## Usage

- TODO; coming as the project nears v1.0

<!-- Troubleshooting -->
## Troubleshooting

If you run into trouble you can ask `cs` to write some diagnostics to a log file by setting the following environment variables:

| Variable                    | Default   | Description                                                      |
|-----------------------------| --------- | ---------------------------------------------------------------- |
| CONTENTSTACK_CLI_LOG_LEVEL  | `"fatal"` | `"fatal"`, `"error"`, `"warning"`, `"warn"`, `"info"`, `"debug"` |
| CONTENTSTACK_CLI_LOG_FORMAT | `"text"`  | `"text"` or `"json"`                                             |
| CONTENTSTACK_CLI_LOG_FILE   | `""`      | path to a log file; when empty logs go to STDOUT                 |

<!-- ROADMAP -->
## Roadmap

<!-- vale Google.Parens = NO -->
See [open issues](https://github.com/davidalpert/go-contentstack/issues) project board for a list of known issues and up-for-grabs tasks.
<!-- vale Google.Parens = YES -->

## Contributing

See the [CONTRIBUTING](CONTRIBUTING.md) guide for local development setup and contribution guidelines.

<!-- LICENSE -->
## License

Distributed under the GPU v3 License. See [LICENSE](LICENSE) for more information.

<!-- CONTACT -->
## Contact

David Alpert - [@davidalpert](https://twitter.com/davidalpert)

Project Link: [https://github.com/davidalpert/go-contentstack](https://github.com/davidalpert/go-contentstack)

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/davidalpert/go-contentstack
[contributors-image-url]: https://contrib.rocks/image?repo=davidalpert/go-contentstack
[forks-shield]: https://img.shields.io/github/forks/davidalpert/go-contentstack
[forks-url]: https://github.com/davidalpert/go-contentstack/network/members
[issues-shield]: https://img.shields.io/github/issues/davidalpert/go-contentstack
[issues-url]: https://github.com/davidalpert/go-contentstack/issues
[license-shield]: https://img.shields.io/badge/License-GPLv3-blue.svg
[license-url]: https://www.gnu.org/licenses/gpl-3.0

