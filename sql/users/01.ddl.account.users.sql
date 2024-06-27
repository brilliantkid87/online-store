-- account.users definition

-- Drop table

-- DROP TABLE account.users;

CREATE TABLE account.users (
	user_id uuid NOT NULL DEFAULT account.uuid_generate_v4(),
	user_name varchar(100) NOT NULL,
	email varchar(100) NOT NULL,
	created timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	modified timestamptz NULL DEFAULT CURRENT_TIMESTAMP,
	"password" text NULL,
	CONSTRAINT users_pkey PRIMARY KEY (user_id)
);