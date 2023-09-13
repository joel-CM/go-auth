# Authentication and authorization example with gin (golang)

## Routes

- /api/user/register -> user register

```
request:
    - method: POST
    - body: {
        "name": "name",
        "email": "example@gmail.com",
        "password": "example123"
      }
```

```
successful-response:
    JSON: {
	    "message": "user created!"
    }
```

- /api/user/signin -> user signin/login

```
request:
    - method: POST
    - body: {
        "email": "example@gmail.com",
        "password": "example123"
      }
```

```
successful-response:
    JSON: {
	    "token": "access token string"
    }
```

- /api/auth/resource -> protected resource

```
request:
    - method: GET
    - header: {
        "Authorization": "Bearer access-token-string"
    }
```

```
successful-response:
    JSON: {
	    "message": "successful access to resource :)"
    }
```

## Environment variables

create .env file with the following content

```
JWT_SECURE_KEY=your-secret-key
```
