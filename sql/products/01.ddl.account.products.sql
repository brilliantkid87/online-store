-- account.products definition

-- Drop table

-- DROP TABLE account.products;

CREATE TABLE account.products (
	product_id uuid NOT NULL DEFAULT account.uuid_generate_v4(),
	product_name varchar(255) NOT NULL,
	category text NOT NULL,
	price numeric(10, 2) NOT NULL,
	description text NULL,
	CONSTRAINT products_pkey PRIMARY KEY (product_id)
);