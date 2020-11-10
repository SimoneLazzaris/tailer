# tailer
Simple, but effective, tail -f implementation in go.

## Usage
Usage is very simple: just create a new Follower (using `NewFollower` func) and read from it using `Tail` method. Blocks until a line is ready.

```go
import "tailer"
import "fmt"
func main() {
    f := tailer.NewFollower("/var/log/examplelogfile.txt")
    for {
        fmt.Printf(f.Tail()
    }
}
`
