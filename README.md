# GoComputeFlow
Учебный репозиторий для задания по созданию сервиса распредённых вычислений и использования gRPC для взаимодействия между ними

## Оглавление
- [Описание задачи](#описание-задачи)
- [Запуск проекта](#запуск-проекта)
  - [Запуск из под докер контейнера](#запуск-из-под-докер-контейнера)
- [API](#API)
  - [Регистрация](#регистрация)
  - [Авторизация](#авторизация)
  - [Добавление выражения](#добавление-выражения)
  - [Получение списка выражений](#получение-списка-выражений)
  - [Получение значения по идентификатору](#получение-значения-по-идентификатору)
  - [Получение таймаутов вычислений](#получение-таймаутов-вычислений)
  - [Установка таймаутов вычислений](#установка-таймаутов-вычислений)
  - [Мониторинг вычислителей](#мониторинг-вычислителей)
- [Описание структуры проекта в целом](#описание-структуры-проекта-в-целом)
- [Использованные технологии и почему именно они](#использованные-технологии-и-почему-именно-они)
- [Дополнение, если у вас Windows и нет команды curl](#дополнение-если-у-вас-windows-и-нет-команды-curl)
- [Контакты для связи](#контакты-для-связи)


---
## Описание задачи
Основная идея это разработать API сервис для взаимодействия с пользователями, а также сервисы вычислений арифметических выражений.
Это проект прямое продолжения [репозитория](https://github.com/AbrLis/Distributed-computing), решено было переписать всё с нуля, так как появилось новое понимание и желание переделать по другому, получилось ли это?

Поддерживаются основные операции +,-,*,/
Не поддерживается всё остальное, скобки, арифметические функции, и так далее.
Так как основная идея проекта не алгоритм распределённых вычислений, а научиться работать с языком и его особенностями.

Сервер, который имеет следующие endpoint-ы:

- Регистрация пользователя. `api/v1/register`
- Аутентификация - получение jwt токена. `api/v1/login`
- Добавление вычисления арифметического выражения. `api/v1/add-expression`
- Получение списка выражений со статусами. `api/v1/get-expressions`
- Получение значения выражения по его идентификатору. `api/v1/get-value/:task_id`
- Получение списка доступных операций со временем их выполения. `api/v1/get-operations`
- Установка времени выполнения операции. `api/v1/set-operations`
- Мониторинг воркеров - получение времени последнего отклика. `api/v1/monitoring`

[Назад к оглавлению](#оглавление)

---
## Запуск проекта

- Клонировать репозиторий
- Перейти в папку проекта и установить зависимости
```shell
go mod download
```
- Запустить проект с помощью команды 
```shell
go run ./cmd/main.go
```
Проект запустится по адресу `http://localhost:3000`

[Назад к оглавлению](#оглавление)

---
## Запуск из под докер контейнера

- Клонировать репозиторий
- Запустить команду из дирректории проекта
```shell
docker-compose up
```
Севис будет доступен по по адресу `http://localhost:3000`

[Назад к оглавлению](#оглавление)


---
## API

Общие примеры взаимодействия с API можно посмотреть здесь 
`./internal/api/apiExampleRequests/testAPI.http`

или в репозитории [здесь](https://github.com/AbrLis/GoComputeFlow/blob/main/internal/api/apiExampleRequests/testAPI.http)

---
- ### Регистрация
```shell
curl -X POST http://localhost:3000/api/v1/register \
-H 'Content-Type: application/json' \
-d '{
    "login": "admin",
    "password": "admin"
}'
```
В ответ получаем JSON ответ со статусом регистрации
```json
{
  "msg": "Пользователь с логином admin успешно зарегистрирован"
}
```
Либо в случае если такой пользователь уже существует
```json
{
  "error": "user already exists: admin"
}
```
[Назад к оглавлению](#оглавление)

---
- ### Авторизация
```shell
curl -X POST http://localhost:3000/api/v1/login \
-H 'Content-Type: application/json' \
-d '{
    "login": "admin",
    "password": "admin"
}'
```
В ответ получаем JSON ответ с токеном JWT
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTI2NDIxNjksImlhdCI6MTcxMjU1NTc2OSwibG9naW4iOiJhZG1pbiIsIm5iZiI6MTcxMjU1NTc2OSwidXNlcl9pZCI6MX0.dYVKp5H8MY3zaV988hSk_YNEsrS_Kac0tRwpxWhVUdc"
}
```
[Назад к оглавлению](#оглавление)

---
- ### Добавление выражения
```shell
curl -X POST http://localhost:3000/api/v1/add-expression \
-H 'Authorization: Bearer {{jwt}}' \
-H 'Content-Type: text/plain' \
-d '2+2' 
```
Здесь вместо {{jwt}} должен быть вставлен выданный токен JWT

В ответ получаем JSON ответ со статусом добавления
```json
{
  "msg": "Expression added to queue: 2+2"
}
```
[Назад к оглавлению](#оглавление)

---
- ### Получение списка выражений
```shell
curl -X GET http://localhost:3000/api/v1/get-expressions \
-H 'Authorization: Bearer {{jwt}}'
```
Здесь вместо {{jwt}} должен быть вставлен выданный токен JWT
В ответ получаем JSON список выражений
```json
[
  {
    "ID": 1,
    "Expression": "2+2",
    "Result": "4",
    "Status": "Completed"
  }
]
```
[Назад к оглавлению](#оглавление)

---
- ### Получение значения по идентификатору
```shell
curl -X GET http://localhost:3000/api/v1/get-value/1 \
-H 'Authorization: Bearer {{jwt}}'
```

В ответе JSON с выражением конкретного идентификатора
```json
{
  "ID": 1,
  "Expression": "2+2",
  "Result": "4",
  "Status": "Completed"
}
```
[Назад к оглавлению](#оглавление)

---
- ### Получение таймаутов вычислений
    Получение времени выполнения операций (таймаутов), авторизация не требуется.
```shell
curl -X GET http://localhost:3000/api/v1/get-operations
```
В ответе JSON с таймаутами вычислений по операциям
```json
{
  "*": "4.00 sec",
  "+": "5.00 sec",
  "-": "3.00 sec",
  "/": "6.00 sec"
}
```
[Назад к оглавлению](#оглавление)


---
- ### Установка таймаутов вычислений
```shell
curl -X POST 'http://localhost:3000/api/v1/set-operations?mul=1&add=2&sub=2.3&div=2.5' \
-H 'Authorization: Bearer {{jwt}}'
```
Где в query параметрах можно задать время выполения операций в секундах. (если не указать какой-то из таймаутов, то он останется прежним)
В ответе JSON сообщение об установленных параметрах (данные сохраняются между перезапусками)

[Назад к оглавлению](#оглавление)

---
- ### Мониторинг вычислителей
    Авторизация не требуется.
```shell
curl -X GET http://localhost:3000/api/v1/monitoring
```
В ответе JSON с информацией о статусе вычислителей (времени последнего пинга)
```json
{
  "worker 0": "2.76 sec",
  "worker 1": "2.34 sec",
  "worker 2": "2.34 sec",
  "worker 3": "2.34 sec",
  "worker 4": "2.34 sec"
}
```

[Назад к оглавлению](#оглавление)

---
## Описание структуры проекта в целом
![Схема проекта](./images/Schema.jpg)

### Общее описание структуры

- API Взаимодействие с пользователем (регистрация, аутентификация, получение выражений)
- Calulator - Распределитель задач и хаб взаимодействия между базой данных и Worker-оми
- Database - База данных, для хранения данных о пользователях и сохранения таймаутов вычислений, результатов выполнения операций.
- Worker - Сервер для распределенных вычислений, принимает задачи от Calculator и распределяет на свободные воркеры. После завершения вычислений, результат помещается в очередь готовых ответов, окуда, в последствии забирается Calculator и уже заноситься в базу данных.

[Назад к оглавлению](#оглавление)

---
## Использованные технологии и почему именно они
- [gRPC](https://grpc.io/) - это требования проекта.
- [Go](https://go.dev/) - изучаемый язык.
- [Sqlite](https://www.sqlite.org/) - база данных. Для данного проека используется SQLite. Другие базы данных, для такого проекта, считаю излишними.
- [Gorm](https://gorm.io/) - ORM. Первое использование, было интересно попробовать и разобраться как с ним работать.
- [Gin](https://gin-gonic.com/) - маршрутизатор. Опять-же, это первое использоание этого пакета. (В целом считаю его излишним, но было интересно попробовать, поэтому выбрал его)
- [JWT](https://github.com/golang-jwt/jwt) - для авторизации и аутентификации пользователей.

[Назад к оглавлению](#оглавление)


---
## Дополнение, если у вас Windows и нет команды curl
Установить эту утилиту можно с официального [сайта](https://curl.se/)

И так же ссылка на сам файл для Windows [downloads](https://curl.se/download.html#Win64)
[Назад к оглавлению](#оглавление)

---
## Контакты для связи
- [Telegram](https://t.me/AbrLis)
- [GitHubIssues](https://github.com/AbrLis/GoComputeFlow/issues)


[Назад к оглавлению](#оглавление)