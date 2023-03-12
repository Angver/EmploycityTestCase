# EmploycityTestCase

# Тестовое задание
Написать простой GRPC сервис с командами: get, set и delete.
Хранилище должно быть спрятано за интерфейсом.

Реализуйте интерфейс обоими путями:
- Memcached сервер с самописной библиотекой и тремя этими же командами. [Memcached protocol](https://github.com/memcached/memcached/blob/master/doc/protocol.txt)
- Хранилище внутри памяти приложения
  Оформите в виде git-репозитория и покройте тестами.
  Продвинутый уровень: реализуйте пулл коннектов к memcached.

## Сборка Docker Image
```bash
docker build -t employcity-test-case . 
```

## Аргументы и переменные окружения
Смотрите список всех аргументов и переменных окружения в файле [cmd/employcity-test-case/config.go](cmd/employcity-test-case/config.go)

## Примеры запуска
### Поднятие отдельных сервисов-зависимостей
```bash
docker compose -f ci/dev/docker-compose.yml up
```

### Запуск docker-контейнера
```bash
docker run \
  --rm \
  -it \
  --env-file=ci/dev/.env \
  ./cmd/employcity-test-case
```

### Запуск скомпилированного исполняемого файла
```bash
(set -a && source ci/dev/.env && set +a && ./cmd/employcity-test-case/employcity-test-case)
```
