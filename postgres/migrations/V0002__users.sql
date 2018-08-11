CREATE SCHEMA users;

CREATE TABLE users (
	id				BIGSERIAL		PRIMARY KEY,
	name			VARCHAR(60)		UNIQUE,
	email			VARCHAR(254)	NOT NULL UNIQUE,
	pass			VARCHAR(60)		NOT NULL,
	roles			INTEGER			NOT NULL,
	status			INTEGER			NOT NULL,
	verification	VARCHAR(60) 	NOT NULL,
	created			TIMESTAMP		NOT NULL DEFAULT now()
);

INSERT INTO ops (op) VALUES('migration V0002__users.sql');