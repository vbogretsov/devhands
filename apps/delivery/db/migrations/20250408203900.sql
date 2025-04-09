-- Create "sm_spot" table
CREATE TABLE "sm_spot" (
  "id" integer NOT NULL,
  "created_at" timestamp NOT NULL,
  "name" character varying(255) NOT NULL,
  "description" character varying(255) NOT NULL,
  "country" character varying(2) NOT NULL,
  "state" character varying(100) NOT NULL,
  "city" character varying(100) NOT NULL,
  "lat" numeric(9,6) NOT NULL,
  "lng" numeric(9,6) NOT NULL,
  "raiting" double precision NOT NULL,
  PRIMARY KEY ("id")
);
-- Create "sm_item" table
CREATE TABLE "sm_item" (
  "id" integer NOT NULL,
  "spot_id" integer NOT NULL,
  "created_at" timestamp NOT NULL,
  "name" character varying(255) NOT NULL,
  "description" character varying(2047) NOT NULL,
  "raiting" double precision NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_sm_item_sm_spot" FOREIGN KEY ("spot_id") REFERENCES "sm_spot" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "sm_order" table
CREATE TABLE "sm_order" (
  "id" serial NOT NULL,
  "created_at" timestamp NOT NULL,
  "client_id" character varying(10) NOT NULL,
  "courier_id" character varying(10) NULL,
  "status" smallint NOT NULL,
  PRIMARY KEY ("id")
);
-- Create "sm_order_item" table
CREATE TABLE "sm_order_item" (
  "order_id" integer NOT NULL,
  "item_id" integer NOT NULL,
  PRIMARY KEY ("order_id", "item_id"),
  CONSTRAINT "fk_sm_order_item_sm_item" FOREIGN KEY ("order_id") REFERENCES "sm_item" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_sm_order_item_sm_order" FOREIGN KEY ("order_id") REFERENCES "sm_order" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
