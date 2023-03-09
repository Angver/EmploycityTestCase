#!/bin/sh -e

REPO_NAME="employcity-test-case"
PROTO_PATHS="internal/api/grpc internal/specs/grpcclient"

# Запуск buf
run_buf(){
  docker run --rm -w "/work" -v "$(pwd):/work" docker.citik.ru/base/buf:latest $@
}

# Обрабатывает прото файлы
process_proto_files(){
  local COMMAND="$1"
  local PROTO_DIR="$2"

  if [ ! -d "$PROTO_DIR" ]; then
    return 0
  fi

  run_buf $@
}

# Генерация прото файлов
gen_proto(){
  for CURPATH in ${PROTO_PATHS}; do
    echo "start process $CURPATH..."

    rm -Rf $CURPATH/gen/*
    process_proto_files lint "$CURPATH/proto"
    process_proto_files generate "$CURPATH/proto" --template "$CURPATH/proto/buf.gen.yaml"

    echo "finish process $CURPATH..."
  done
}

# Запуск unit-тестов
unit(){
  echo "run unit tests"
  go test ./...
}

unit_race() {
  echo "run unit tests with race test"
  go test -race ./...
}

fmt() {
  echo "run go fmt"
  go fmt ./...
}

vet() {
  echo "run go vet"
  go vet ./...
}

# Подтянуть зависимости
deps(){
  go get ./...
}

# Собрать исполняемый файл
build(){
  deps
  go build ./cmd/grpc-skeleton
}

# Собрать docker образ
build_docker() {
  build
  docker build -t "$REPO_NAME:local" .
  rm ./"$REPO_NAME"
}

using(){
  echo "Укажите команду при запуске: ./run.sh [command]"
  echo "Список команд:"
  echo "  unit - запустить unit-тесты"
  echo "  unit_race - запуск unit тестов с проверкой на data-race"
  echo "  deps - подтянуть зависимости"
  echo "  build - собрать приложение"
  echo "  build_docker - собрать локальный docker образ"
  echo "  fmt - форматирование кода при помощи 'go fmt'"
  echo "  vet - проверка правильности форматирования кода"
  echo "  gen_proto - генерация прото файлов (для клиентов и сервера)"
}

command="$1"
if [ -z "$command" ]
then
 using
 exit 0;
else
 $command $@
fi
