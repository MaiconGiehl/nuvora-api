DROP TABLE account CASCADE;

DROP TABLE bus CASCADE;

DROP TABLE city CASCADE;

DROP TABLE company CASCADE;

DROP TABLE company_type CASCADE;

DROP TABLE customer CASCADE;

DROP TABLE permission_level CASCADE;

DROP TABLE person CASCADE;

DROP TABLE ticket CASCADE;

DROP TABLE ticket_status CASCADE;

DROP TABLE travel CASCADE;


CREATE TABLE "person" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY,
  "permission_level" smallint NOT NULL,
  "customer_id" integer UNIQUE,
  "company_id" integer UNIQUE,
  "city_id" integer NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp
);

CREATE TABLE "customer" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY,
  "cpf" bigint UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "phone" numeric,
  "birth_date" date,
  "company_id" integer NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp
);

CREATE TABLE "company" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY,
  "cnpj" bigint UNIQUE NOT NULL,
  "social_reason" varchar UNIQUE NOT NULL,
  "fantasy_name" varchar,
  "phone" numeric,
  "company_type_id" integer NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp
);

CREATE TABLE "company_type" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY,
  "description" varchar NOT NULL
);

CREATE TABLE "account" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "person_id" integer UNIQUE NOT NULL,
  "last_access" date NOT NULL,
  "tickets_left" smallint,
  "daily_tickets" smallint NOT NULL,
  "created_at" date NOT NULL,
  "updated_at" date
);

CREATE TABLE "bus" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "number" numeric NOT NULL,
  "max_passengers" smallint,
  "account_id" integer NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp
);

CREATE TABLE "ticket" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY UNIQUE PRIMARY KEY,
  "account_id" integer NOT NULL,
  "status_id" integer NOT NULL,
  "travel_id" integer NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp
);

CREATE TABLE "ticket_status" (
  "id" integer PRIMARY KEY,
  "description" varchar NOT NULL
);

CREATE TABLE "travel" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "price" float NOT NULL,
  "account_id" integer NOT NULL,
  "bus_id" integer NOT NULL,
  "status" integer NOT NULL,
  "departure_time" timestamp NOT NULL,
  "departure_city_id" integer NOT NULL,
  "arrival_time" timestamp NOT NULL,
  "arrival_city_id" integer NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp
);

CREATE TABLE "city" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "permission_level" (
  "id" integer UNIQUE PRIMARY KEY,
  "description" varchar NOT NULL
);

ALTER TABLE "travel" ADD FOREIGN KEY ("bus_id") REFERENCES "bus" ("id");

ALTER TABLE "travel" ADD FOREIGN KEY ("departure_city_id") REFERENCES "city" ("id");

ALTER TABLE "travel" ADD FOREIGN KEY ("arrival_city_id") REFERENCES "city" ("id");

ALTER TABLE "account" ADD FOREIGN KEY ("person_id") REFERENCES "person" ("id");

ALTER TABLE "person" ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");

ALTER TABLE "person" ADD FOREIGN KEY ("customer_id") REFERENCES "customer" ("id");

ALTER TABLE "company" ADD FOREIGN KEY ("company_type_id") REFERENCES "company_type" ("id");

ALTER TABLE "person" ADD FOREIGN KEY ("permission_level") REFERENCES "permission_level" ("id");

ALTER TABLE "travel" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "ticket" ADD FOREIGN KEY ("status_id") REFERENCES "ticket_status" ("id");

ALTER TABLE "ticket" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "customer" ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");

ALTER TABLE "person" ADD FOREIGN KEY ("city_id") REFERENCES "city" ("id");

ALTER TABLE "bus" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

-- Tickets status
INSERT INTO ticket_status (id, description)
VALUES (0, 'NOT PAID');

INSERT INTO ticket_status (id, description)
VALUES (1, 'PAID');

-- Permission level
INSERT INTO permission_level (id, description) VALUES (1, 'TRAVEL COMPANY');
INSERT INTO permission_level (id, description) VALUES (2, 'COMPANY');
INSERT INTO permission_level (id, description) VALUES (3, 'CUSTOMER');

-- Company types
INSERT INTO company_type (id, description) VALUES (0, 'CONSUMER');
INSERT INTO company_type (id, description) VALUES (1, 'TRANSPORTER');


-- Cities
INSERT INTO city (name) values ('SANTA ROSA');
INSERT INTO city (name) values ('CANDIDO GODOI');
INSERT INTO city (name) values ('SANTO CRISTO');
INSERT INTO city (name) values ('TRES DE MAIO');
INSERT INTO city (name) values ('TUCUNDUVA');
INSERT INTO city (name) values ('TUPARENDI');
INSERT INTO city (name) values ('HORIZONTINA');
INSERT INTO city (name) values ('CAMPINA DAS MISSOES');
INSERT INTO city (name) values ('CERRO LARGO');
INSERT INTO city (name) values ('GIRUA');


-- Travel companies
INSERT INTO company (cnpj, social_reason, fantasy_name, phone, company_type_id, created_at) VALUES (04929069000155, 'Empresa Viagem 01', 'Empresa de Viagem 01', 555523142482, 1, NOW());
INSERT INTO person (permission_level, company_id, city_id, created_at) VALUES (1, 1, 1, NOW());
INSERT INTO account (email, password, person_id, last_access, daily_tickets, created_at) VALUES ('empresaviagem01@gmail.com', 'empresaviagem01',  1, NOW(), 0, NOW());

INSERT INTO company (cnpj, social_reason, fantasy_name, phone, company_type_id, created_at) VALUES (72214881000154, 'Empresa Viagem 02', 'Empresa de Viagem 02', 555539341988, 1, NOW());
INSERT INTO person (permission_level, company_id, city_id, created_at) VALUES (1, 2, 1, NOW());
INSERT INTO account (email, password, person_id, last_access, daily_tickets, created_at) VALUES ('empresaviagem02@gmail.com', 'empresaviagem02',  2, NOW(), 0, NOW());

INSERT INTO company (cnpj, social_reason, fantasy_name, phone, company_type_id, created_at) VALUES (23196583000173, 'Empresa Viagem 03', 'Empresa de Viagem 03', 555527535686, 1, NOW());
INSERT INTO person (permission_level, company_id, city_id, created_at) VALUES (1, 3, 1, NOW());
INSERT INTO account (email, password, person_id, last_access, daily_tickets, created_at) VALUES ('empresaviagem03@gmail.com', 'empresaviagem03',  3, NOW(), 0, NOW());


-- Companies
INSERT INTO company (cnpj, social_reason, fantasy_name, phone, company_type_id, created_at) VALUES (41616716000103, 'Empresa 01', 'Empresa 01', 555539215263, 0, NOW());
INSERT INTO person (permission_level, company_id, city_id, created_at) VALUES (2, 4, 1, NOW());
INSERT INTO account (email, password, person_id, last_access, daily_tickets, created_at) VALUES ('empresa01@gmail.com', 'empresa01',  4, NOW(), 0, NOW());

INSERT INTO company (cnpj, social_reason, fantasy_name, phone, company_type_id, created_at) VALUES (46115580000163, 'Empresa 02', 'Empresa 02', 555533726238, 0, NOW());
INSERT INTO person (permission_level, company_id, city_id, created_at) VALUES (2, 5, 2, NOW());
INSERT INTO account (email, password, person_id, last_access, daily_tickets, created_at) VALUES ('empresa02@gmail.com', 'empresa02',  5, NOW(), 0, NOW());

INSERT INTO company (cnpj, social_reason, fantasy_name, phone, company_type_id, created_at) VALUES (54723440000104, 'Empresa 03', 'Empresa 03', 555539323815, 0, NOW());
INSERT INTO person (permission_level, company_id, city_id, created_at) VALUES (2, 6, 3, NOW());
INSERT INTO account (email, password, person_id, last_access, daily_tickets, created_at) VALUES ('empresa03@gmail.com', 'empresa03',  6, NOW(), 0, NOW());


-- Customers
INSERT into customer (cpf, name, phone, company_id, created_at) VALUES (00728497034, 'Usuario 01', 5555767070535, 4, NOW());
INSERT INTO person (permission_level, customer_id, city_id, created_at) VALUES (3, 1, 1, NOW());
INSERT INTO account (email, password, person_id, last_access, daily_tickets, tickets_left, created_at) VALUES ('usuario01@gmail.com', 'usuario01',  7, NOW(), 2, 2, NOW());

INSERT into customer (cpf, name, phone, company_id, created_at) VALUES (78084020013, 'Usuario 02', 5555813941193, 4, NOW());
INSERT INTO person (permission_level, customer_id, city_id, created_at) VALUES (3, 2, 1, NOW());
INSERT INTO account (email, password, person_id, last_access, daily_tickets, tickets_left, created_at) VALUES ('usuario02@gmail.com', 'usuario02',  8, NOW(), 2, 2, NOW());

INSERT into customer (cpf, name, phone, company_id, created_at) VALUES (16339577059, 'Usuario 03', 5555262952682, 5, NOW());
INSERT INTO person (permission_level, customer_id, city_id, created_at) VALUES (3, 3, 1, NOW());
INSERT INTO account (email, password, person_id, last_access, daily_tickets, tickets_left, created_at) VALUES ('usuario03@gmail.com', 'usuario03',  9, NOW(), 2, 2, NOW());

INSERT into customer (cpf, name, phone, company_id, created_at) VALUES (91554720010, 'Usuario 04', 5555469199633, 5, NOW());
INSERT INTO person (permission_level, customer_id, city_id, created_at) VALUES (3, 4, 1, NOW());
INSERT INTO account (email, password, person_id, last_access, daily_tickets, tickets_left, created_at) VALUES ('usuario04@gmail.com', 'usuario04',  10, NOW(), 2, 2, NOW());

INSERT into customer (cpf, name, phone, company_id, created_at) VALUES (43796465056, 'Usuario 05', 5555752633987, 6, NOW());
INSERT INTO person (permission_level, customer_id, city_id, created_at) VALUES (3, 5, 1, NOW());
INSERT INTO account (email, password, person_id, last_access, daily_tickets, tickets_left, created_at) VALUES ('usuario05@gmail.com', 'usuario05',  11, NOW(), 2, 2, NOW());

INSERT into customer (cpf, name, phone, company_id, created_at) VALUES (10086181017, 'Usuario 06', 5555822191928, 6, NOW());
INSERT INTO person (permission_level, customer_id, city_id, created_at) VALUES (3, 6, 1, NOW());
INSERT INTO account (email, password, person_id, last_access, daily_tickets, tickets_left, created_at) VALUES ('usuario06@gmail.com', 'usuario06',  12, NOW(), 2, 2, NOW());

-- Bus
INSERT INTO bus (number, max_passengers, account_id, created_at)
VALUES (505, 38, 1, NOW());

INSERT INTO bus (number, max_passengers, account_id, created_at)
VALUES (405, 42, 1, NOW());

INSERT INTO bus (number, max_passengers, account_id, created_at)
VALUES (202, 52, 1, NOW());

INSERT INTO bus (number, max_passengers, account_id, created_at)
VALUES (709, 38, 2, NOW());

INSERT INTO bus (number, max_passengers, account_id, created_at)
VALUES (290, 52, 2, NOW());

INSERT INTO bus (number, max_passengers, account_id, created_at)
VALUES (900, 52, 2, NOW());

INSERT INTO bus (number, max_passengers, account_id, created_at)
VALUES (310, 52, 3, NOW());

INSERT INTO bus (number, max_passengers, account_id, created_at)
VALUES (101, 42, 3, NOW());

INSERT INTO bus (number, max_passengers, account_id, created_at)
VALUES (630, 44, 3, NOW());

INSERT INTO bus (number, max_passengers, account_id, created_at)
VALUES (805, 38, 3, NOW());


-- Travels
INSERT INTO travel (price, account_id, bus_id, status, departure_time, departure_city_id, arrival_time, arrival_city_id, created_at)
VALUES (7.50, 1, 1, 0, NOW(), 1, NOW(), 2, NOW());

INSERT INTO travel (price, account_id, bus_id, status, departure_time, departure_city_id, arrival_time, arrival_city_id, created_at)
VALUES (7.50, 1, 2, 0, NOW(), 1, NOW(), 2, NOW());

INSERT INTO travel (price, account_id, bus_id, status, departure_time, departure_city_id, arrival_time, arrival_city_id, created_at)
VALUES (7.50, 1, 3, 0, NOW(), 1, NOW(), 2, NOW());


-- Tickets
INSERT INTO ticket (account_id, status_id, travel_id, created_at)
VALUES (7, 0, 1, NOW());

INSERT INTO ticket (account_id, status_id, travel_id, created_at)
VALUES (7, 0, 1, NOW());

INSERT INTO ticket (account_id, status_id, travel_id, created_at)
VALUES (8, 0, 1, NOW());

INSERT INTO ticket (account_id, status_id, travel_id, created_at)
VALUES (8, 0, 1, NOW());

INSERT INTO ticket (account_id, status_id, travel_id, created_at)
VALUES (9, 0, 1, NOW());

INSERT INTO ticket (account_id, status_id, travel_id, created_at)
VALUES (9, 0, 1, NOW());

INSERT INTO ticket (account_id, status_id, travel_id, created_at)
VALUES (10, 0, 1, NOW());

INSERT INTO ticket (account_id, status_id, travel_id, created_at)
VALUES (10, 0, 1, NOW());