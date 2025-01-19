# Телеграм бот Read advisor

## Принцип работы

Пользователь запускает бота, в нем будут следующие разделы:

- Просмотр ленты
- Список сохраненных книг
- Уточнение предпочтений, возможность вписать любимых авторов, жанры, издательства

### Работа ленты

При первоначальном использовании, бот просит пользователя рассказать о своих предпочтениях, он может вписать туда авторов, жанры и издательства, но на английском языке из за специфики api

В ленте показывается по одной книге за раз, под книгой есть кнопка добавления в список для чтения. Лента формируется из предпочтений пользователя, аггрегируется результат нескольки рандомных запросов, потому что books api работает как обычный поиск

При просмотре ленты, пользователь может отметить книгу как понравившуюся или скипнуть, при отметке книги как понравившейся, добавляется информация об этой книге (автор, жанры, издательство) в препочтения пользователя и ставится таймштамп для вычислений

При добавлении книги в список чтения в сервис рекомендаций так же отправляется информация о книге и ей ставится больший приоритет, чем у лайков

### Приоритизация предпочтений

На первом месте стоят предпочтения которые указал пользователь вручную, у них нет таймштампов и они учитываются в большинстве случаев, далее по значимости идут предпочтения связанные с книгами из списка для чтения, в самом конце, но в самом большом количестве идет информация о понравившихся книгах, у них будут таймштампы и их количество будет автоматически чиститься

## Планируемые сервисы

### Телеграм бот

Здесь вся логика взаимодействия с пользователем через бота, не имеет своей бд, так как он выступает в роли gateway и отправляет информацию в другие сервисы.

**Взаимодействие с другими сервисами:**

- Общается по gRPC с сервисом рекомендаций для получения ленты
- Общается по gRPC с сервисом списка книг
- Отправляет данные о лайках и предпочтениях в сервис предпочтений через RabbitMQ

### Книжная полка

Хранит информацию о сохранненых книгах пользователей в своей бд, используя id

**Взаимодействие с другими сервисами:**

- Имеет gRPC API для CRUD операций со списком книг
- Отправляет данные о сохраненных книгах в сервис предпочтений через RabbitMQ

### Books

Инкапсулирует запросы к google books api для удобного и эффективного взаимодействия

**Взаимодействие с другими сервисами:**

- Предоставляет gRPC методы для получения информации

### Сервис формирования предпочтений

Обрабатывает информацию о лайках, сохраненных книгах и предпочтениях, и отправляет ее в сервис рекомендаций. Работает асинхронно, так как процесс не быстрый, а изменение ленты постепенное

**Взаимодействие с другими сервисами:**

- Получает данные о лайках, сохраненных книгах, предпочтениях из RabbitMQ
- Отправляет информацию о предпочтениях пользователя в сервис рекомендаций по RabbitMQ
- Ходит в сервис books по gRPC для получения нужной информации

### Сервис рекомендаций

Хранит в себе предпочтения пользователей, но основе которых делает запросы к сервису books и формирует ленту

**Взаимодействие с другими сервисами:**

- Получает данные о предпочтениях из RabbitMQ
- Предоставляет gRPC API для получения ленты
- Ходит в сервис books для получения информации о книгах
