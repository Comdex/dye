# Dye

Go wrapper for Pygments

## Getting Started

Make sure your have `pygments` installed via `easy_install` or `pip`.

Grab the Dye package:

    go get github.com/tombell/dye

Call the `Highlight` function to highlight a string using the specified lexer.

```go
import (
    "fmt"

    "github.com/tombell/dye"
)

func main() {
    input := "puts \"Hello world!\""
    fmt.Printf(dye.Highlight(input, "ruby", "html", "utf-8"))
}
```

## Features

If your `pygmentize` binary is not in your `PATH` you can use the `Binary`
function to set the path to your `pygmentize` binary.

There is a `Which` function that can be used to get the current path to the
`pygmentize` binary.

The `Highlight` function is the meat of the package. It takes four arguments,
the input, the lexer (language), the output format and encoding.
## Contributing

Here's the most direct way to get your work merged into the project:

  1. Fork the project
  2. Clone down your fork
  3. Create a feature branch
  4. Hack away and add tests. Not necessarily in that order
  5. Make sure everything still passes by running tests
  6. Push the branch up
  7. Send a pull request for your branch

## License

Copyright © 2013, Tom Bell. See the `LICENSE` file for license rights and
limitations (MIT).
