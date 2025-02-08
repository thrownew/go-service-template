# Pupa

Pupa is a service template example. It's a simple service that can be used as a starting point for building more complex services.

## Architecture Overview

This service template uses `github.com/uber-go/dig` as a Dependency Injection container, providing a clean and convenient way to register dependencies. The architecture follows low coupling principles, allowing each package to be self-contained and independent.

### Key Features

- **Dependency Injection**: Uses `dig` for flexible dependency management.
- **Database Integration**: Includes MySQL with migrations as an example of database integration
- **Server Implementation**: Application implements example of gRPC server.
- **Low Coupling**: Each package can independently register its components. It's allow to use minimal coupling approach.
- **Simple Addition of cmd Commands to the Application**: Adding new `cmd` commands to the application is straightforward. Application is designed to be extensible, allowing you to easily introduce new commands without affecting the existing functionality.
- **Easy testing**: The dependency container allows do integration tests with any combination of dependencies.

### Architectural Principles

The application follows minimal coupling principles, where each package can register its own:
- Services
- Repositories
- Dependencies
- Other components

For example, the `dogs` package demonstrates this by registering its own:
```go
deps.ProvideAll(
    deps.Provide(NewServer),
    servers.GRPCServiceAdapter[*Server](),
    deps.Provide(NewRepository),
)
```

This approach ensures that:
- Packages are self-contained.
- No direct dependencies between packages.
- Removing a package from the application eliminates its dependencies from go.mod.
- Easy to add new features without modifying existing code.
- Clear separation of concerns.
- Testable components.

The DI container handles all dependency wiring, allowing packages to remain isolated while still working together cohesively.

## Dev environment

Dev environment is based on Docker Compose for now. Maybe in the future we will use `kind` for Kubernetes.

Up dev and connect to the database:
```bash
make up
mysql -upupa -ppupa --host 127.0.0.1 --port 3306 -A pupa
```

Down dev:
```bash
make down
```

## Run service

Describe application commands:
```bash
make help
```

Run the service:
```bash
make run
```

Run all tests:
```bash
make test
```

Try to make grpc calls manually:
```bash
grpcurl -plaintext -d '{"name": "Buddy"}' 127.0.0.1:8080 pupa.DogService.DogIsGoodBoyV1
grpcurl -plaintext -d '{"name": "Undefined"}' 127.0.0.1:8080 pupa.DogService.DogIsGoodBoyV1
grpcurl -plaintext -d '{"name": "Luna"}' 127.0.0.1:8080 pupa.DogService.DogIsGoodBoyV1
```

Or just run `cmd` command:
```bash
make wof
```