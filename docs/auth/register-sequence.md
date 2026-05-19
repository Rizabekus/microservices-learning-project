# Register Sequence

```text
Client->Auth: POST /register
Auth->Auth: Validate request
Auth->DB: Check user exists
DB-->Auth: Not found
Auth->Auth: Hash password
Auth->DB: Insert user
DB-->Auth: User created
Auth-->Client: 201 Created
```