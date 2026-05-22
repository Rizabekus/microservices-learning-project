# Create Order Sequence

```text
Client->Order: POST /orders
Order->Order: Validate request
Order->Order: Validate authenticated user
Order->Order: Create order entity
Order->DB: Insert order
DB-->Order: Order created
Order-->Client: Order response
```