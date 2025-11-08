CREATE TABLE invoices (
  id SERIAL PRIMARY KEY,
  numero BIGINT NOT NULL UNIQUE,
  status VARCHAR(20) NOT NULL, -- 'Aberta' | 'Fechada'
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE invoice_items (
  id SERIAL PRIMARY KEY,
  invoice_id INT NOT NULL REFERENCES invoices(id) ON DELETE CASCADE,
  product_codigo VARCHAR(100) NOT NULL,
  quantidade INT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
