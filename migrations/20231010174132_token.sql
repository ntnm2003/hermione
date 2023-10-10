CREATE TABLE IF NOT EXISTS "tokens"(
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    token varchar(250),
    user_id uuid REFERENCES users(id),
    PRIMARY KEY(id)
    );
