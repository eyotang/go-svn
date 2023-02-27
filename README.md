# SVN golang lib

`SVN golang lib` is a awesome lib for golang to operate svn. It's used with `username` and `password` without any prompt.

It's a wrapper for `svn` commands, and format its output as golang structure .

```go
package main

import (
	"fmt"

	"github.com/eyotang/go-svn"
)

func main() {
	var (
        url  = "https://localhost/svn/Example"
		username = "admin"
		password = "123456"
        path     = "Game01"
	)

	repo, err := svn.NewRemoteRepo(url, username, password)
	if err != nil {
		fmt.Printf("New RemoteRepo is failed! err: %+v", err)
	}
	result, err := repo.List(path)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("result: %s\n", result)
}
```
