# Pupa

Pupa is a service template example.

### Dev environment

Up dev and connect to the database:
```bash
make up
mysql -upupa -ppupa --host 127.0.0.1 --port 3306 -A pupa
```

Down dev:
```bash
make down
```

### Run service

Up the service:
```bash
make run
```

Try to make grpc calls:
```bash
grpcurl -plaintext -d '{"name": "Buddy"}' 127.0.0.1:8080 pupa.DogService.DogIsGoodBoyV1
grpcurl -plaintext -d '{"name": "Undefined"}' 127.0.0.1:8080 pupa.DogService.DogIsGoodBoyV1
grpcurl -plaintext -d '{"name": "Luna"}' 127.0.0.1:8080 pupa.DogService.DogIsGoodBoyV1
```