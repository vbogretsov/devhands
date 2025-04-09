-- Modify "sm_item" table
ALTER TABLE "sm_item" ADD COLUMN "updated_at" timestamp NOT NULL;
-- Modify "sm_order" table
ALTER TABLE "sm_order" ADD COLUMN "updated_at" timestamp NOT NULL;
-- Modify "sm_spot" table
ALTER TABLE "sm_spot" ADD COLUMN "updated_at" timestamp NOT NULL;
