# LinkedList

A simple Go implementation of a singly linked list with basic operations and comprehensive unit tests.

## Features
- Create new list nodes
- Push multiple values
- Add single values
- Convert list to array
- Invert the list
- Print list contents

## Usage

```go
import "github.com/renatospaka/linkedlist"

node := linkedlist.NewListNode()
node.Push([]int{1, 2, 3})
node.Add(4)
arr := node.ToArray() // [1, 2, 3, 4]
inv := node.ToInverse() // returns a new list with values [4, 3, 2, 1]
node.List() // prints: nodes: [1 2 3 4]
```

## Testing

Run all tests:

```bash
go test ./... -v
```

## License
MIT
