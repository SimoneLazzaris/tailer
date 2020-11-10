# tailer
Simple, but effective, `tail -f` implementation in go.

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
```

Function Tail wakes up every 100ms to check if there is any new line ready.
It detects file rotation using simple euristics, which proved simple but very effective:
 - inode change
 - reduction in size
 - any error accessing the file
In case of errors, simply sleeps and retries later. 
