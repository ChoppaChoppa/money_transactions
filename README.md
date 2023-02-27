# bwg_test

## Алгоритм работы

при запуске сервиса, запускается функция, которая считывает не обработанные транзакции из базы (запрос выводит транзакции отсортировав дату их создания по возростания, соответсвенно все транзакции выполнятся по очереди их создания) и запускает воркер в бесконечном цикле, который в зависимости от типа транзакции их обрабатывает (если количество попыток обработать транзакцию первысело 5, то статус транзакции будет установлен как необрабатываемый и больше не будет взят в обработку), если кол-во транзакций превышает некоторое заранее установленное n, то данный воркер вызвает новый воркер(рекурсивно) передавая туда копию массива транзакций начиная с n элемента, а сам обработает транзакции до n элементов.

*Пример*

если у нас есть массив **[1, 2, 3, 4, 5, 6, 7, 8]** и **n = 2**, то кол-во воркеров будет 4 и обрабатывать,
начиная с 1 воркера, будут: **[1, 2], [3, 4], [4, 5], [6, 8]**

Только после завершения всех воркеров функция по новой сделает запрос на получение не обработанных транзакций
и запустит воркеры.

есть два api, которые служат для создания транзакции, после создания транзакции,
они будут обработанны ранее запущенными воркерами.

#API
**ввод средств**

создает новую транзакцию в бд, при удачном запросе возвращает оповещение, что транзакция находится в обработке

Request:

`[POST] transaction/in`

Body:
```json
{
  "user_id": 0,
  "amount": 123.0
}
```

Response:
```json
{
  "error": false,
  "error_text": "",
  "data": null,
  "code": 200
}
```

**вывод средств**

создает новую транзакцию в бд, при удачном запросе возвращает оповещение, что транзакция находится в обработке

Request:

`[POST] transaction/out`

Body:
```json
{
  "user_id": 0,
  "amount": 123.0
}
```

Response:
```json
{
  "error": false,
  "error_text": "",
  "data": null,
  "code": 200
}
```

**Получение транзакций**

возвращает весь список транзакций

Request:

`[GET] transaction/get/:id`

Response:
```json
{
  "error": false,
  "error_text": "",
  "data": [ 
    {
      "ID": 0,
      "user_id": 1,
      "attempts": 0,
      "status": 2,
      "type": 2,
      "amount": 10,
      "date": "2023-02-01T14:38:00.807416Z"
    },
    ...
  ],
  "code": 200
}
```

**Получение баланса**

ввыодит текущий баланс пользователя

Request:

`[GET] balance/get/:1`

Response:

```json
{
  "error": false,
  "error_text": "",
  "data": {
    "id": 0,
    "user_id": 1,
    "balance": 100.0
  },
  "code": 200
}
```

## Примечания

1. при запросе на получение транзакций можно реализовать пагинацию
2. на данный момент реализации, несколько попыток обработать транзакцию выполняется за секунды,
можно добавить некий delay в степени attempt, который является временем,
и при сравнивании этого времени с временем создания транзакции, решить обработать ее или нет
3. если при выполнении запроса в базу на обнавления баланса выпадет ошибка, сервис обработает только одну ошибку,
это ошибка на отрицательный баланс, если возникла тех. ошибка, сервер обработает транзакцию только в следующем потоке.
С точки зрения работоспособности это никак не влияет на сервис, но может улучшить производительность.
