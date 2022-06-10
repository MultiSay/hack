-- Доступные офферы

CREATE TABLE IF NOT EXISTS "hack"."user" (
    "user_id" serial NOT NULL,
    "created_at" timestamp(0) with time zone NOT NULL DEFAULT now(),
    "updated_at" timestamp(0) with time zone
);