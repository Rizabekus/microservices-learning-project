# Cancel Order Sequence

```text
Client->Order: DELETE /orders/{id}
Order->Order: Validate request
Order->DB: Find order by ID
DB-->Order: Order
Order->Order: Validate order status
Order->DB: Update order status to cancelled
DB-->Order: Updated
Order-->Client: Success response
```