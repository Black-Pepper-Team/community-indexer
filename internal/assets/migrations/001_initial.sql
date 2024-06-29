-- +migrate Up

CREATE TABLE communities(
    id               VARCHAR(36)  PRIMARY KEY NOT NULL,
    status           VARCHAR(256)             NOT NULL,
    contract_address BYTEA                    NOT NULL,
    name             VARCHAR(256)             NOT NULL,
    owner_address    BYTEA                    NOT NULL
);

-- +migrate Down

DROP TABLE communities;
