CREATE TABLE IF NOT EXISTS files (
  id          BIGSERIAL PRIMARY KEY,
  name        VARCHAR(200) NOT NULL,
  folder_id   BIGINT       NULL,
  owner_id    BIGINT       NOT NULL,
  type        VARCHAR(50)  NOT NULL,
  path        VARCHAR(250) NOT NULL,
  created_at  TIMESTAMPTZ  NOT NULL DEFAULT now(),
  updated_at  TIMESTAMPTZ  NOT NULL DEFAULT now(),
  deleted     BOOLEAN      NOT NULL DEFAULT FALSE,

  CONSTRAINT fk_files_folder
    FOREIGN KEY (folder_id) REFERENCES folders(id) ON DELETE CASCADE,
  CONSTRAINT fk_files_owner
    FOREIGN KEY (owner_id)  REFERENCES users(id)   ON DELETE CASCADE,

  -- Evita nomes duplicados dentro da mesma pasta (opcional mas comum)
  CONSTRAINT uq_files_folder_name UNIQUE (folder_id, name)
);

-- Índices úteis
CREATE INDEX IF NOT EXISTS idx_files_folder  ON files(folder_id);
CREATE INDEX IF NOT EXISTS idx_files_owner   ON files(owner_id);
CREATE INDEX IF NOT EXISTS idx_files_deleted ON files(deleted);

-- Trigger updated_at
DROP TRIGGER IF EXISTS trg_files_updated_at ON files;
CREATE TRIGGER trg_files_updated_at
BEFORE UPDATE ON files
FOR EACH ROW EXECUTE FUNCTION set_updated_at();

