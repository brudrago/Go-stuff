CREATE TABLE IF NOT EXISTS folders (
  id          BIGSERIAL PRIMARY KEY,
  name        VARCHAR(60)  NOT NULL,
  parent_id   BIGINT       NULL,
  created_at  TIMESTAMPTZ  NOT NULL DEFAULT now(),
  updated_at  TIMESTAMPTZ  NOT NULL DEFAULT now(),
  deleted     BOOLEAN      NOT NULL DEFAULT FALSE,
  CONSTRAINT fk_folders_parent
    FOREIGN KEY (parent_id) REFERENCES folders(id) ON DELETE CASCADE
);

-- Índices úteis
CREATE INDEX IF NOT EXISTS idx_folders_parent  ON folders(parent_id);
CREATE INDEX IF NOT EXISTS idx_folders_deleted ON folders(deleted);

-- Trigger updated_at
DROP TRIGGER IF EXISTS trg_folders_updated_at ON folders;
CREATE TRIGGER trg_folders_updated_at
BEFORE UPDATE ON folders
FOR EACH ROW EXECUTE FUNCTION set_updated_at();
