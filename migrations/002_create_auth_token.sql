
DROP TABLE IF EXISTS "auth_tokens";
CREATE TABLE "public"."auth_tokens" (
    "token_id" uuid DEFAULT gen_random_uuid() NOT NULL,
    "user_id" uuid,
    "token_value" text NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "expires_at" timestamp,
    CONSTRAINT "auth_tokens_pkey" PRIMARY KEY ("token_id")
)
WITH (oids = false);

