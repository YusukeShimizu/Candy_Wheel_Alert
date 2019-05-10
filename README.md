## Getting Started

#### Set envs.

```go
Env      string `required:"true" envconfig:"ENV"`
Secret   string `required:"true" envconfig:"SECRET"`
Token    string `required:"true" envconfig:"TOKEN"`
ID       string `required:"true" envconfig:"ID"`
```

#### Run on local.

```sh
go run main.go
```
