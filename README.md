# GoRoulette

Tiny [Go](https://golang.org) wrapper library around the API provided by
[Netflix Roulette](http://netflixroulette.net/api/).

## Installation

Just go with

```shell
go get github.com/jreyeshdez/goroulette
```

## Usage

```go

package main

import (
	"fmt"
	"github.com/jreyeshdez/goroulette"
)

func main() {
	movies, err := goroulette.GetMoviesByTitle("Bad Boys")
	if err == nil {
		director := movies.GetDirector()
		fmt.Println(fmt.Sprintf("The director of Bad Boys is: %s", director))
	}
}

```

## Contribution

Please, feel free to contribute to this project
by following the "fork-and-pull" Git workflow.

	1. Fork the repo on GitHub.
	2. Clone the project to your own machine.
	3. Commit changes to your own branch.
	4. Push your work back up to your fork.
	5. Submit a Pull request so that changes can be reviewed.

Any feedback is welcomed, just trying to have everything as readable as possible.

Thank you for your contributions in advance.

## License

This library is provided with a MIT License.
