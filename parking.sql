-- Table: owner

-- DROP TABLE owner;

CREATE TABLE owner
(
  id smallserial NOT NULL,
  nama character varying(100),
  nim character varying(20) NOT NULL,
  CONSTRAINT owner_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);