## Getting Started

#### Set envs.

```go
Env      string `required:"true" envconfig:"ENV"`
Secret   string `required:"true" envconfig:"SECRET"`
Token    string `required:"true" envconfig:"TOKEN"`
ID       string `required:"true" envconfig:"ID"`
Pace   string `required:"true" envconfig:"PACE"`
```

#### Run on local.

```sh
go run main.go
```

#### Run on Docker

```sh
docker build -t candy ./ --build-arg secret=YOUR_SECRET --build-arg token=YOUR_TOKEN --build-arg id=YOUR_ID

docker run -it candy
```
