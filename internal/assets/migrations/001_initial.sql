-- +migrate Up

CREATE TABLE communities(
    id               UUID   PRIMARY KEY NOT NULL,
    contract_address VARCHAR(42)        NOT NULL,
    name             VARCHAR(256)       NOT NULL,
    owner_address    VARCHAR(42)        NOT NULL
);

-- +migrate Down

DROP TABLE communities;
