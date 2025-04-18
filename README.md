# Скрипты для учета рабочего времени в Битрикс24

Этот проект содержит два скрипта на Go для автоматизации учета рабочего времени в Битрикс24:
- Скрипт для начала рабочего дня
- Скрипт для окончания рабочего дня

## Структура проекта

```
day_scripts/
├── start_day_script/    # Скрипт для начала рабочего дня
│   ├── main.go
│   └── go.mod
└── end_day_script/      # Скрипт для окончания рабочего дня
    ├── main.go
    └── go.mod
```

## Требования

- Go 1.21 или выше
- Доступ к Битрикс24 через вебхук

## Настройка

1. Склонируйте репозиторий
2. В файлах `start_day_script/main.go` и `end_day_script/main.go` замените значение `webhookURL` на ваш вебхук Битрикс24
3. В обоих файлах замените значение `USER_ID` на ваш ID пользователя в Битрикс24

## Сборка и запуск

### Скрипт начала рабочего дня

```bash
cd start_day_script
go build
./start_day_script  # для Linux/Mac
start_day_script.exe  # для Windows
```

### Скрипт окончания рабочего дня

```bash
cd end_day_script
go build
./end_day_script  # для Linux/Mac
end_day_script.exe  # для Windows
```

## Функциональность

### Скрипт начала рабочего дня
- Отправляет запрос к Битрикс24 для отметки начала рабочего дня
- Выводит время начала и ответ от API

### Скрипт окончания рабочего дня
- Отправляет запрос к Битрикс24 для отметки окончания рабочего дня
- Выводит время окончания и ответ от API

## Обработка ошибок

Скрипты обрабатывают следующие ошибки:
- Ошибки при создании JSON
- Ошибки при отправке HTTP-запроса
- Ошибки при чтении ответа от API

## Примечания

- Убедитесь, что у вас есть корректный вебхук с необходимыми правами доступа
- Скрипты можно настроить на автоматический запуск через планировщик задач
- Время выполнения запросов выводится в формате ЧЧ:ММ:СС
