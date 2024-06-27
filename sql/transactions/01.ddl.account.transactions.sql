-- account.transactions definition

-- Drop table

-- DROP TABLE account.transactions;

CREATE TABLE account.transactions (
	transaction_id uuid NOT NULL DEFAULT account.uuid_generate_v4(),
	user_id uuid NOT NULL,
	transaction_date timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	total_amount numeric(10, 2) NOT NULL,
	payment_method varchar(50) NOT NULL,
	status varchar(20) NOT NULL DEFAULT 'Pending'::character varying,
	CONSTRAINT transactions_pkey PRIMARY KEY (transaction_id)
);


-- account.transactions foreign keys

ALTER TABLE account.transactions ADD CONSTRAINT transactions_user_id_fkey FOREIGN KEY (user_id) REFERENCES account.users(user_id);