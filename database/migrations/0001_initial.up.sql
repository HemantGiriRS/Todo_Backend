BEGIN;

CREATE TABLE IF NOT EXISTS users
(
    id          UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    name        TEXT NOT NULL,
    email       TEXT NOT NULL,
    password    TEXT NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    archived_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS todo
(
    id           UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    user_id      UUID REFERENCES users (id) NOT NULL,
    name         TEXT                       NOT NULL,
    description  TEXT                       NOT NULL,
    is_completed BOOLEAN                  DEFAULT FALSE,
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    archived_at  TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS user_session
(
    id          UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    user_id     UUID REFERENCES users (id) NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    archived_at TIMESTAMP WITH TIME ZONE
);

COMMIT;