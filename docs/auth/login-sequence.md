# Login Sequence

```text
Client->Auth: POST /login
Auth->Auth: Validate request
Auth->DB: Find user by email
DB-->Auth: User
Auth->Auth: Verify password
Auth->Auth: Generate access token
Auth->Auth: Generate refresh token
Auth-->Client: Tokens
```