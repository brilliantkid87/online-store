-- account.transaction_details definition

-- Drop table

-- DROP TABLE account.transaction_details;

CREATE TABLE account.transaction_details (
	detail_id uuid NOT NULL DEFAULT account.uuid_generate_v4(),
	transaction_id uuid NOT NULL,
	product_id uuid NOT NULL,
	quantity int4 NOT NULL,
	price numeric(10, 2) NOT NULL,
	CONSTRAINT transaction_details_pkey PRIMARY KEY (detail_id)
);


-- account.transaction_details foreign keys

ALTER TABLE account.transaction_details ADD CONSTRAINT transaction_details_product_id_fkey FOREIGN KEY (product_id) REFERENCES account.products(product_id);
ALTER TABLE account.transaction_details ADD CONSTRAINT transaction_details_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES account.transactions(transaction_id);