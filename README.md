# API-тесты для проекта https://testboard.avito.com/

## Полезные ссылки

1. [Swagger-спецификация](https://testboard.avito.com/swagger/index.html)
2. [Техническое задание](https://docs.google.com/document/d/1IZ9SxEcePoQT6lyEmukycp1qNok4HZtok73UQnPwKSA/edit?usp=sharing)
3. [Библиотека для написания API-тестов](https://github.com/steinfletcher/apitest)
4. [Мощная библиотека анализа JSON для Go](https://github.com/tidwall/gjson)

## Запуск тестов

1. Установите любое IDE для разработки: [Jetbrains GoLand](https://www.jetbrains.com/ru-ru/go/) (Платное или по студенческой лицензии), [VS Code](https://code.visualstudio.com/) (бесплатное).
2. Установите [Git](https://git-scm.com/install/) на компьютер
3. [Установите](https://go.dev/doc/install) Go версии 1.25.1
4. [Настройте](https://docs.github.com/ru/authentication/connecting-to-github-with-ssh) доступ к своему Github-аккаунту
   по SSH-ключу.
5. Склонируйте себе репозиторий: `git clone <путь к вашему репозиторию>`
6. Подтяните необходимые зависимости через команду в корне проекта:

```bash
   go mod tidy
```

7. В файле `.env.override` (создайте его в корне, если нет) выставьте правильные значения для переменных окружения:

```
TEST_LOGIN=your_login
TEST_PASSWORD=your_password
SECOND_TEST_LOGIN=second_account_login
SECOND_TEST_PASSWORD=second_account_password
```

Для тестов избранного нужны два разных аккаунта TestBoard: основной создаёт объявления,
второй добавляет их в избранное. Зарегистрируйте второй аккаунт самостоятельно
и задайте `SECOND_TEST_LOGIN` и `SECOND_TEST_PASSWORD` в `.env.override`.
Новые сценарии создают временные объявления и удаляют их после выполнения.
После удаления проверяется HTTP 404 при получении объявления. При очистке ответы
429, 502, 503 и 504 допускают до трёх попыток; проверки самих сценариев остаются строгими.
Запросы выполняются с паузой 500 мс для ограничения нагрузки на учебный стенд.
Проверка лимита избранного создаёт до 51 временного объявления и проверяет HTTP 400
при добавлении сверх 50 записей. После теста объявления удаляются, исходное избранное
восстанавливается. Подготовка и очистка делают этот сценарий продолжительным.

8. Для запуска всех тестов достаточно выполнить команду в корне проекта:

```bash
  go test -v ./...
```

9. Отдельный пакет с тестами можно запустить командой:

```bash
  go test -v ./tests/scenarios/myAdvertisement/
```

Новые сценарии можно запустить отдельно без кеширования результатов:

```bash
go test -count=1 -v ./tests/scenarios/createAdvertisement/
go test -count=1 -v ./tests/scenarios/favorites/
```

## Логирование

Для включения логирования request/response в консоль: выставьте переменную окружения `DEBUG=true` в
`.env.override`

## Запуск линтера

1. Установить приложение [golangci-lint](https://golangci-lint.run/docs/welcome/install/local/) версии **2** (конфигурация проекта несовместима с версией 1).
2. Запуск линтера (из корня репозитория):

```bash
  golangci-lint run
```
