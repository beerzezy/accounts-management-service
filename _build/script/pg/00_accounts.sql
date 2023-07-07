\c accountsmanagement

BEGIN;
DROP TABLE IF EXISTS accounts;
CREATE TABLE accounts
(
   id SERIAL NOT NULL,
   full_name text,
   email text,
   password text,
	created_date timestamptz NOT NULL,
	updated_date timestamptz NULL,
   PRIMARY KEY (id)
);
END;