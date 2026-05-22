# Get Order Sequence

```text
Client->Order: GET /orders/{id}
Order->Order: Validate request
Order->DB: Find order by ID
DB-->Order: Order
Order->Order: Check ownership/access
Order-->Client: Order response
```