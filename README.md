# Product Catalogue Service

This service is supposed to be consumed and used as per requirement given in the recruitment drive conducted by Waquar for his team.

How to consume the endpoints:

1. Get a product details
```
curl -X GET http://waquarcodes.mooo.com:4051/inventory/902/availability
where 902 is the productId
```

3. Reduce item count
```
 curl -X POST http://waquarcodes.mooo.com:4051/inventory/deduct -H "Content-Type: application/json" -d '{
  "deductions": [
    { "id": 901, "quantity": 5 }
  ]
}'
```
4. Create a product in inventory service
```
curl -X POST http://waquarcodes.mooo.com:4051/inventory \
-H "Content-Type: application/json" \
-d '{
  "id": 801,
  "quantity": 50,
  "price": 19.99,
  "name": "potato"
}'
```

6. Get all inventory items
```
curl -X GET http://waquarcodes.mooo.com:4051/inventory
```

7. Update product count

```
curl -X PUT http://waquarcodes.mooo.com:4051/inventory/<productId>   -H "Content-Type: application/json" -d '{
  "quantity": 60,
  "price": 21.99
}'
```
