BEGIN;

CREATE TABLE IF NOT EXISTS users.users
(

    created_by   VARCHAR(200),
    updated_by   VARCHAR(200),
    deleted_by   VARCHAR(200),
    created_at   TIMESTAMP,
    updated_at   TIMESTAMP,
    deleted_at   TIMESTAMP,
    id           uuid PRIMARY KEY,
    username     VARCHAR(255) UNIQUE NULL,
    email        VARCHAR(255) UNIQUE NULL,
    phone_number VARCHAR(255) UNIQUE NULL,
    password     VARCHAR(255)        NULL
);

COMMIT;
