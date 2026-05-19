# Refresh Sequence

```text
Client->Auth: POST /refresh
Auth->Auth: Parse refresh token
Auth->Auth: Validate signature
Auth->Auth: Validate expiration
Auth->Auth: Generate new access token
Auth-->Client: New access token
```
# Validate Token Sequence

```text
Order->Auth gRPC: ValidateToken(accessToken)
Auth->Auth: Parse JWT
Auth->Auth: Validate signature
Auth->Auth: Validate expiration
Auth-->Order: user_id, valid=true
```