
# Kyrol Marketplace API

A (not yet)ready-to-use, Golang backend API version of Kyrol Marketplace.




## Features

- JWT Authentication
- Product variations
- Covers from the purchase to the end of the transaction
- Cross platform


## Tech Stack

**Server:** Golang, Gin

**Database:** MySQL


## API Reference

#### Get all products

```http
  GET /api/v1/product
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

Response - 200

```javascript
{
  "data": [
    {
      "id": "EN100123",
      "name": "Cock Enlarging Potion",
      "description": "Potion to drink when you're insecure about your johny's size",
      "variant": [
        {
          "id": "EN100123-100",
          "name": "100ML",
          "price": 250000,
          "stock": 15
        },
        {
          "id": "EN100123-250",
          "name": "250ML",
          "price": 500000,
          "stock": 18
        }
      ]
    },
    {
      "id": "EN100124",
      "name": "Cock Dwindling Substance",
      "description": "",
      "variant": []
    }
  ]
}
```

#### Get product

```http
  GET /api/v1/product/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Login

```http
  POST /auth/login
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `email`      | `string` | **Required**. User email |
| `password`      | `string` | **Required**. Password |

#### Register

```http
  POST /auth/register
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `first_name`      | `string` | **Required**. First name |
| `last_name`      | `string` | **Required**. Last name |
| `username`      | `string` | **Required**. Username |
| `email`      | `string` | **Required**. Email |
| `password`      | `string` | **Required**. Password |

#### Logout

Not yet implemented (Read the note on controllers/auth-controllers/logout)


## Authors

- [@kyroline (Rizky Wahyu Dewantoro)](https://github.com/Kyroline)

