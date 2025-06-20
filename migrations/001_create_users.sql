-- Adminer 5.3.0 PostgreSQL 16.9 dump
DROP TABLE IF EXISTS "users";
CREATE TABLE "public"."users" (
    "user_id" uuid DEFAULT gen_random_uuid() NOT NULL,
    "role" character varying(20) DEFAULT 'user',
    "email" character varying(255) NOT NULL,
    "vorname" character varying(30),
    "nachname" character varying(30),
    CONSTRAINT "users_pkey" PRIMARY KEY ("user_id")
)
WITH (oids = false);

CREATE UNIQUE INDEX users_email_key ON public.users USING btree (email);


ALTER TABLE ONLY "public"."auth_tokens" ADD CONSTRAINT "auth_tokens_user_id_fkey" FOREIGN KEY (user_id) REFERENCES users(user_id) NOT DEFERRABLE;

-- 2025-06-20 16:39:45 UTC