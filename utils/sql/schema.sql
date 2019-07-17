CREATE DATABASE mediocirs;

CREATE TABLE MUser (
  id         SERIAL,
  w_name     varchar(255) NOT NULL,
  w_id       varchar(255) PRIMARY KEY,
  created_at timestamp
);

CREATE TABLE MGroup (
  id          SERIAL PRIMARY KEY,
  name        varchar(255) NOT NULL,
  description text,
  updated_at  timestamp
);

CREATE TABLE BelongTo (
  id              SERIAL,
  u_id            varchar(255) NOT NULL,
  g_id            int NOT NULL,
  FOREIGN KEY (u_id) REFERENCES MUser(w_id) ON DELETE CASCADE,
  FOREIGN KEY (g_id) REFERENCES MGroup(id) ON DELETE CASCADE,
  PRIMARY KEY (id, u_id, g_id)
);

CREATE TABLE Record (
  id          SERIAL,
  g_id        int,
  day         date NOT NULL,
  payer       varchar(255) NOT NULL,
  spliters    varchar[],
  pay_amount  float NOT NULL,
  description text NOT NULL,
  updated_at  timestamp,
  deleted     boolean DEFAULT false,
  FOREIGN KEY (g_id) REFERENCES MGroup(id) ON DELETE CASCADE,
  FOREIGN KEY (payer) REFERENCES MUser(w_id) ON DELETE CASCADE,
  PRIMARY KEY (id, g_id)
);

CREATE TABLE OpHistory (
  id         SERIAL,
  u_id       varchar(255) NOT NULL,
  g_id       int,
  message    text NOT NULL,
  created_at timestamp,
  FOREIGN KEY (u_id) REFERENCES MUser(w_id) ON DELETE CASCADE,
  FOREIGN KEY (g_id) REFERENCES MGroup(id) ON DELETE CASCADE,
  PRIMARY KEY (id, u_id, g_id)
);
