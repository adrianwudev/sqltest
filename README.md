```
project/
├── cmd/
│   └── main.go
├── internal/
│   ├── domain/
│   │   ├── entity.go
│   │   ├── repository.go
│   │   └── usecase.go
│   ├── interfaces/
│   │   ├── http/
│   │   │   ├── handler.go
│   │   │   └── router.go
│   │   └── persistence/
│   │       └── repository.go
│   └── app/
│       └── app.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   ├── mysql/
│   │   └── postgresql/
│   ├── logger/
│   ├── middleware/
│   └── util/
├── migrations/
├── test/
├── Makefile
└── README.md
```