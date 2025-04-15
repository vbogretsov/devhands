-- Create "sm_spot" table
CREATE TABLE "sm_spot" (
  "id" serial NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
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
