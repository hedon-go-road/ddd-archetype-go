# ddd-archetype-go

`ddd-archetype-go` is a DDD go language development of scaffolding, refer to [ddd-archetype](https://github.com/feiniaojin/ddd-archetype).

![architecture](assets/ddd-layout.png)

## Quick Start

### Requirements

- Go 1.24.1+
- Docker
  > used for running mysql and redis
- Redis
  > (default: localhost:6379)
- MySQL
  > (default: root:root@tcp(localhost:33061)/ddd-archetype-go?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai)

### Run

```bash
make start
```

### Test

Send HTTP Request by using the http files in [testdata](./testdata) directory.
