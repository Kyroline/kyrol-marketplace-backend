
# Kyrol Marketplace API

A ready-to-use, Golang backend API version of Kyrol Marketplace.




## Features

- JWT Authentication
- Product variations
- Covers from the purchase to the end of the transaction
- Cross platform


## Tech Stack

**Server:** Golang, Gin


## API Reference

#### Get all products

```http
  GET /api/v1/product
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

#### Get product

```http
  GET /api/product/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

