CREATE TABLE IF NOT EXISTS clients (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT NOT NULL,
  cpf TEXT NOT NULL,
  document BYTEA NOT NULL,
  street TEXT,
  number INTEGER,
  complement TEXT,
  neighborhood TEXT,
  city TEXT,
  state TEXT,
  cep TEXT,
  status BOOLEAN DEFAULT FALSE
);