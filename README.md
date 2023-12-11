# Game of Life

![GitHub Latest Release)](https://img.shields.io/github/v/release/kulapard/gol?logo=github)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/kulapard/gol/blob/master/LICENSE)
[![Build Status](https://github.com/kulapard/gol/actions/workflows/ci.yml/badge.svg)](https://github.com/kulapard/gol/actions/workflows/release.yml)
[![Build Status](https://github.com/kulapard/gol/actions/workflows/release.yml/badge.svg)](https://github.com/kulapard/gol/actions/workflows/release.yml)
[![codecov](https://codecov.io/github/kulapard/gol/graph/badge.svg?token=Z9SAAI8VQ4)](https://codecov.io/github/kulapard/gol)
[![Go Report Card](https://goreportcard.com/badge/github.com/kulapard/gol)](https://goreportcard.com/report/github.com/kulapard/gol)

[Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) written in Go

![screencast](assets/gol_random_30x40.gif)

## Install ##

Using [Homebrew](https://brew.sh/) (OS X / Linux)

```shell
brew install kulapard/tap/gol
```

## Update ##

Using [Homebrew](https://brew.sh/) (OS X / Linux)

```shell
brew update
brew upgrade gol
```

## Usage ##

Run with default parameters (size 30 x 40, random initial state):

```shell
gol run
```

Run initial state from file:

```shell
gol run --file ./examples/board.txt
```

To specify the size of the board and speed of the game:

```shell
gol run --rows 30 --cols 40 --speed 10
```

File format:

- `.` (dot) - dead cell
- `0` (zero) - alive cell
- ` ` (space) - ignored, uses for readability
- `\n` (new line) - end of row
- all other characters cause a parsing error

Example:

```text
. . . . .
. . 0 . .
. 0 . 0 .
. . 0 . .
. . . . .
```

To see all available options:

```shell
gol run --help
```

## License ##

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.