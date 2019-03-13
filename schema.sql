CREATE TABLE shortlink (
  id SERIAL,
  link TEXT NOT NULL UNIQUE,
  slug TEXT NOT NULL,
  PRIMARY KEY (slug)
);

/*
CREATE TABLE clicks (
  id SERIAL,
  link_id INT NOT NULL,
  metadata_id INT,
  PRIMARY KEY (id)
);

CREATE TABLE click_metadata (
  id SERIAL,
  timestamp TIMESTAMP WITH TIME ZONE,
  useragent_id INT,
  PRIMARY KEY (id)
);
CREATE TABLE useragents (
  id SERIAL,
  useragent TEXT NOT NULL UNIQUE,
  UNIQUE (useragent),
  PRIMARY KEY (id)
);
*/
