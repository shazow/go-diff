[![GoDoc](https://godoc.org/github.com/shazow/go-diff?status.svg)](https://godoc.org/github.com/shazow/go-diff)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/shazow/go-diff/master/LICENSE)
[![Build Status](https://travis-ci.org/shazow/go-diff.svg?branch=master)](https://travis-ci.org/shazow/go-diff)


# go-diff

Library for generating Git-style diff patchsets in Go.

Built to be used in the pure-Go implementation of the Git backend for
[sourcegraph's go-vcs](https://github.com/sourcegraph/go-vcs).


## Features

- Git-style patch headers for each file (are there other styles to support?).
- Bring your own diff algorithm by implementing the *Differ* interface.
- Includes a *Differ* implementation using [diffmatchpatch](https://godoc.org/github.com/sergi/go-diff/diffmatchpatch).


## Sponsors

Work on this package is sponsored by [Sourcegraph](https://sourcegraph.com/).


## License

MIT.
