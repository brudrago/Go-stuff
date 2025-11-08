CREATE TABLE IF NOT EXISTS users (
  id          BIGSERIAL PRIMARY KEY,
  name        VARCHAR(200) NOT NULL,
  login       VARCHAR(80)  NOT NULL UNIQUE,
  password    TEXT         NOT NULL,
  created_at  TIMESTAMPTZ  NOT NULL DEFAULT now(),
  updated_at  TIMESTAMPTZ  NOT NULL DEFAULT now(),
  deleted     BOOLEAN      NOT NULL DEFAULT FALSE,
  last_login  TIMESTAMPTZ
);

-- Índices úteis
CREATE INDEX IF NOT EXISTS idx_users_deleted ON users (deleted);

-- Trigger p/ updated_at
CREATE OR REPLACE FUNCTION set_updated_at() RETURNS trigger AS $$
BEGIN
  NEW.updated_at := now();
  RETURN NEW;
END
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_users_updated_at ON users;
CREATE TRIGGER trg_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
