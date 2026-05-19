# Auth Flows

## Register

### Description
Registers a new user in the system.

### Steps
1. User sends email and password.
2. System validates request payload.
3. System checks if user already exists.
4. Password is hashed using bcrypt.
5. User is saved to database.
6. Success response returned.

---

## Login

### Description
Authenticates user and issues tokens.

### Steps
1. User sends email and password.
2. System validates credentials.
3. Password hash is verified.
4. Access token is generated.
5. Refresh token is generated.
6. Tokens returned to client.

---

## Refresh Token

### Description
Issues new access token using refresh token.

### Steps
1. Client sends refresh token.
2. System validates refresh token.
3. System checks token expiration/revocation.
4. New access token generated.
5. New access token returned.

---

## Validate Token

### Description
Used by internal services to validate access token.

### Steps
1. Service sends access token.
2. System validates signature.
3. System validates expiration.
4. Claims extracted.
5. User info returned.