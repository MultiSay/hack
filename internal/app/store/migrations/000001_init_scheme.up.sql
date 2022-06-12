CREATE TABLE "public"."files" (
    "id" int4 NOT NULL DEFAULT nextval('files_id_seq' :: regclass),
    "name" varchar NOT NULL,
    "create_at" timestamptz NOT NULL DEFAULT now(),
    "send_at" timestamptz,
    "received_at" timestamptz,
    "size" int8 NOT NULL,
    "status" varchar NOT NULL DEFAULT '' :: character varying,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."region_predict" (
    "id" int4 NOT NULL DEFAULT nextval('region_predict_id_seq' :: regclass),
    "position" int4 NOT NULL,
    "city" varchar NOT NULL,
    "current_client_index" int4 NOT NULL DEFAULT 0,
    "predict_client_index" int4 NOT NULL DEFAULT 0,
    "predict_arpu" int4 NOT NULL DEFAULT 0,
    "predict_score" int4 NOT NULL DEFAULT 0,
    PRIMARY KEY ("id")
);