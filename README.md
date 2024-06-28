
## Installation

To run this project

```bash
  go mod tidy
```

```bash
  go run main.go
```

  server running on localhost:8080

## API Documentation

[Documentation](https://documenter.getpostman.com/view/27589244/2sA3dsnEMr)

## Entity Relationship Diagram (ERD)

### Tables

#### users

| Column Name     | Data Type         | Description                       |
|-----------------|-------------------|-----------------------------------|
| user_id         | UUID (PK)         | Primary key for the user          |
| user_name       | VARCHAR(100)      | User's username                   |
| email           | VARCHAR(100)      | User's email address              |
| created         | TIMESTAMPTZ       | Timestamp of user creation        |
| modified        | TIMESTAMPTZ       | Timestamp of last modification    |
| password | TEXT              | Hashed password for authentication|

#### products

| Column Name | Data Type         | Description                            |
|-------------|-------------------|----------------------------------------|
| product_id  | UUID (PK)         | Primary key for the product             |
| product__name        | VARCHAR(255)      | Product name                           |
| description | TEXT              | Product description                    |
| price       | NUMERIC(10, 2)    | Price of the product                   |
| category | TEXT              | Category Product      |


#### shopping_cart

| Column Name | Data Type    | Description                                  |
|-------------|--------------|----------------------------------------------|
| cart_id     | UUID (PK)    | Primary key for the shopping cart entry.      |
| user_id     | UUID (FK)    | Foreign key referencing the user who owns the cart. |
| product_id  | UUID (FK)    | Foreign key referencing the product added to the cart. |
| quantity    | INTEGER      | Optional field indicating the quantity of the product in the cart. |
| added_at    | TIMESTAMP    | Timestamp when the product was added to the cart. |

##### Relationships

- `shopping_cart.user_id` → `users.user_id` (Many-to-One): Each shopping cart entry belongs to one user.
- `shopping_cart.product_id` → `products.product_id` (Many-to-One): Each shopping cart entry corresponds to one product.



#### transactions

| Column Name      | Data Type          | Description                                           |
|------------------|--------------------|-------------------------------------------------------|
| transaction_id   | UUID (PK)          | Primary key for the transaction                        |
| user_id          | UUID (FK to users) | Foreign key referencing the user who made the transaction |
| transaction_date | TIMESTAMP          | Date and time when the transaction occurred            |
| total_amount     | NUMERIC(10, 2)     | Total amount of the transaction                        |
| payment_method   | VARCHAR(50)        | Payment method used for the transaction                |
| status           | VARCHAR(20)        | Status of the transaction (e.g., Pending, Completed)   |

##### Relationships

- `transactions.user_id` → `users.user_id` (Many-to-One): Each transaction belongs to one user.


#### transaction_details

| Column Name    | Data Type    | Description                                   |
|----------------|--------------|-----------------------------------------------|
| detail_id      | UUID (PK)    | Primary key for the transaction detail         |
| transaction_id | UUID (FK)    | Foreign key referencing the transaction        |
| product_id     | UUID (FK)    | Foreign key referencing the product            |
| quantity       | INTEGER      | Quantity of the product purchased in the transaction |
| price          | NUMERIC(10, 2)| Price per unit of the product                  |

##### Relationships

- `transaction_details.transaction_id` → `transactions.transaction_id` (Many-to-One): Each transaction detail belongs to one transaction.
- `transaction_details.product_id` → `products.product_id` (Many-to-One): Each transaction detail corresponds to one product.




