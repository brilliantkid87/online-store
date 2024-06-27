-- account.shopping_cart definition

-- Drop table

-- DROP TABLE account.shopping_cart;

CREATE TABLE account.shopping_cart (
	cart_id uuid NOT NULL DEFAULT account.uuid_generate_v4(),
	user_id uuid NOT NULL,
	product_id uuid NOT NULL,
	quantity int4 NULL DEFAULT 1,
	added_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT shopping_cart_pkey PRIMARY KEY (cart_id)
);


-- account.shopping_cart foreign keys

ALTER TABLE account.shopping_cart ADD CONSTRAINT shopping_cart_product_id_fkey FOREIGN KEY (product_id) REFERENCES account.products(product_id);
ALTER TABLE account.shopping_cart ADD CONSTRAINT shopping_cart_user_id_fkey FOREIGN KEY (user_id) REFERENCES account.users(user_id);