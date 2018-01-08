# About

Command line tool to validate TOML syntax of input file.


## Installation

Windows and macOS binaries are available under [Releases](https://github.com/martinlindhe/validtoml/releases)

Or install from source:

    go get -u github.com/martinlindhe/validtoml


## Usage

Exit code will be 0 if file is good.

    $ validtoml file.toml
    OK: file.toml

    $ curl http://site.com/file.toml | validtoml
    OK: -


## License

Under [MIT](LICENSE)
