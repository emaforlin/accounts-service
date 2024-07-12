# Accounts Service

## Features

* **Config manager:** Simple setup using the `config.yaml` file.
* **Input validation:** All RPC calls are validated by an interceptor.
* ...

## Next stage

* **Role based authorization**
* **Config file Watcher**
* **TLS**
* ...

## Setup

Setup the service instace by editing `config.yaml`.

``` yaml
service:
  api: "v0.1"
  port: 50010
database:
  name: "accounts-service"
  user: "app-user"
  password: "changepass"
  host: "localhost:3306"
```
