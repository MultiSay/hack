CREATE TABLE "public"."files" (
    "id" int4 NOT NULL DEFAULT nextval('files_id_seq' :: regclass),
    "name" varchar NOT NULL,
    "create_at" timestamptz NOT NULL DEFAULT now(),
    "send_at" timestamptz,
    "received_at" timestamptz,
    "size" int8 NOT NULL,
    "status" varchar NOT NULL DEFAULT '' :: character varying,
    "type" varchar NOT NULL DEFAULT 'debit' :: character varying,
    "status_message" varchar,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."region_predict" (
    "id" int4 NOT NULL DEFAULT nextval('region_predict_id_seq' :: regclass),
    "position" int4 NOT NULL,
    "city" varchar NOT NULL DEFAULT '' :: character varying,
    "current_client_index" int4 NOT NULL DEFAULT 0,
    "predict_client_index" int4 NOT NULL DEFAULT 0,
    "predict_arpu" int4 NOT NULL DEFAULT 0,
    "predict_score" float4 NOT NULL DEFAULT 0,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."leads" (
    "id" int4 NOT NULL DEFAULT nextval('leads_id_seq' :: regclass),
    "client_id" varchar NOT NULL,
    "product_category_name" varchar NOT NULL,
    "utm_source" varchar,
    "utm_content" varchar,
    "utm_campaign" varchar,
    "date" timestamptz NOT NULL DEFAULT now(),
    "cpc" int4 NOT NULL,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."compaign" (
    "id" int4 NOT NULL DEFAULT nextval('compaign_id_seq' :: regclass),
    "utm_campaign" varchar,
    "gender" varchar,
    "age_from" int4,
    "age_to" int4,
    "city" varchar,
    "theme" varchar,
    PRIMARY KEY ("id")
);