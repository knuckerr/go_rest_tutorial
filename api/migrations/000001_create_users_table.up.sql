BEGIN;
CREATE SEQUENCE  users_id_seq;
CREATE TABLE IF NOT EXISTS "public"."users" (
    "id" bigint DEFAULT nextval('users_id_seq') NOT NULL,
    "nickname" character varying(255) NOT NULL,
    "email" character varying(100) NOT NULL,
    "password" character varying(100) NOT NULL,
    "role" character varying(100) NOT NULL,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "users_email_key" UNIQUE ("email"),
    CONSTRAINT "users_nickname_key" UNIQUE ("nickname"),
    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);
ALTER SEQUENCE users_id_seq OWNED BY users.id;
commit;
