CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  codigo VARCHAR(100) NOT NULL UNIQUE,
  descricao VARCHAR(500) NOT NULL,
  saldo INTEGER NOT NULL DEFAULT 0,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- tabela de logs/opcional para idempotência
CREATE TABLE product_adjustments (
  id SERIAL PRIMARY KEY,
  request_id VARCHAR(100) UNIQUE,
  product_id INT NOT NULL REFERENCES products(id),
  delta INT NOT NULL,
  status VARCHAR(50) DEFAULT 'pending',
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

