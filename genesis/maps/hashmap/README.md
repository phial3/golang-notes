# hashmap

[![CI Workflow Badge](https://github.com/SemihBKGR/hashmap/actions/workflows/ci.yml/badge.svg)](https://github.com/SemihBKGR/hashmap/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/SemihBKGR/hashmap)](https://goreportcard.com/report/github.com/SemihBKGR/hashmap)
[![Codecov](https://codecov.io/gh/SemihBKGR/hashmap/branch/master/graph/badge.svg?token=ygx9oP5oRn)](https://codecov.io/gh/SemihBKGR/hashmap)
[![Godoc Badge](https://godoc.org/github.com/seihbkgr/hashmap?status.png)](https://pkg.go.dev/github.com/semihbkgr/hashmap)

Concurrent hash map implementation in go.

hashmap supports generic types, so go version 1.19 is required.

```bash
go get github.com/semihbkgr/hashmap@latest
```

string type key map example

```go
m := NewString[int]()

m.Put("nine", 9)

if val, ok := m.Get("nine"); ok {
    _ = val + 11
    // ...
}

_ = m.GetOrDefault("twenty-three", 23)

if ok := m.Contains("nine"); ok {
    // ...
}
```

custom type key map example

```go
type User struct {
    ID   uint
    Name string
}

func (u User) Hash() uint32 {
    return uint32(u.ID)
}

func (u User) Equals(a any) bool {
    user, ok := a.(User)
    if !ok {
        return false
    }
    return u.ID == user.ID
}

type UserContext map[string]any

func main() {

    m := New[User, UserContext]()

    user := User{
        ID:   1,
        Name: "alice",
    }

    userContext:=make(map[string]any)
    userContext["sessionID"]=sessionID()

    m.Put(user, userContext)

    if val, ok := m.Get(User{ID: 1}); ok {
        _ = val["sessionID"]
        // ...
    }

}
```
