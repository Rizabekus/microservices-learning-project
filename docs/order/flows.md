# Order Flows

## Create Order

### Description
Creates a new order for the authenticated user.

### Steps
1. Client sends create order request.
2. System validates request payload.
3. System validates authenticated user.
4. Order entity is created.
5. Order is saved to database.
6. Success response returned.

---

## Get Order

### Description
Returns order details by order ID.

### Steps
1. Client sends order ID.
2. System validates request.
3. System fetches order from database.
4. System checks ownership/access.
5. Order returned to client.

---

## Get My Orders

### Description
Returns list of orders for authenticated user.

### Steps
1. Client sends request.
2. System validates authenticated user.
3. System fetches user orders from database.
4. Orders returned to client.

---

## Cancel Order

### Description
Cancels existing order.

### Steps
1. Client sends order ID.
2. System validates request.
3. System fetches order from database.
4. System checks order status.
5. Order status updated to cancelled.
6. Success response returned.