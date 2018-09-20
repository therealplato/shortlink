CREATE TABLE metadata (
  base_url TEXT NOT NULL CHECK (right (base_url, 1) = '/'),
  PRIMARY KEY (base_url)
);

CREATE TABLE shortlink (
  id SERIAL,
  link TEXT NOT NULL UNIQUE,
  slug TEXT NOT NULL,
  PRIMARY KEY (slug)
);

CREATE TABLE clicks (
  id SERIAL,
  link_id INT NOT NULL,
  timestamp TIMESTAMP WITH TIME ZONE,
  header_key TEXT NOT NULL,
  header_value TEXT NOT NULL,
  PRIMARY KEY (link_id)
);
