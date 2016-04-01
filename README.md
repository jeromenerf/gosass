# gosass

A tiny script around [github.com/wellington/go-libsass](github.com/wellington/go-libsass) which is itself a binding for `libsass`, the C/C++ library, of the ruby on rails fame.

`go-libsass` doesn't provide a CLI, it is quite a PITA to vendor it (since it takes ages to build the lib), is handles SCSS over SASS by default, the documentation is pretty light. Meh.

Install: `go get -u github.com/jeromenerf/gosass`

Usage: `gosass file.sass > file.css`
