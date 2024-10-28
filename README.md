# User Registration API

## Registering a User

To register, send a `POST` request to `/api/v1/auth/register` with the following JSON payload:

```json
{
  "username": "johndoe",
  "email": "johndoe@example.com",
  "password": "password123"
}
```
