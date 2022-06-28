DROP TABLE IF EXISTS mst_rates;
CREATE TABLE IF NOT EXISTS mst_rates (
  id SERIAL PRIMARY KEY,
  "date" date,
  currency VARCHAR(10),
  rate real,
  UNIQUE ("date", currency)
);