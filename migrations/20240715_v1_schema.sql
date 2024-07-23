-- ####################################
-- OpenCardsOnline Official Schema 
-- VERSION: 20240715_v1
--
-- https://OpenCardsOnline.com
--
-- Licensed Under MIT.
-- Copyright 2024. OpenCardsOnline. All Rights Reserved.
-- 
-- This will allow you to create a new (empty) 
-- database from scratch to get started. You 
-- should run these in order.
-- ####################################

CREATE TABLE IF NOT EXISTS "banned_ips" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "modified_at" TIMESTAMP,
  "is_archived" BOOLEAN DEFAULT false,
  "ip" TEXT UNIQUE NOT NULL,
  "description" TEXT
);

CREATE TABLE IF NOT EXISTS "user_types" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "modified_at" TIMESTAMP,
  "is_archived" BOOLEAN DEFAULT false,
  "name" TEXT NOT NULL,
  "key" TEXT UNIQUE NOT NULL
);

INSERT INTO public.user_types (name, key)
VALUES ('Super Admin','super_admin');

INSERT INTO public.user_types (name, key)
VALUES ('Admin','admin');

INSERT INTO public.user_types (name, key)
VALUES ('Player','player');

CREATE TABLE IF NOT EXISTS "users" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "modified_at" TIMESTAMP,
  "is_archived" BOOLEAN DEFAULT false,
  "email" TEXT UNIQUE NOT NULL,
  "username" TEXT NOT NULL,
  "is_verified" BOOLEAN DEFAULT false,
  "user_type_id" INTEGER,
  "password_hash" TEXT,
  "last_login" TIMESTAMP,
  "failed_login_attempts" INTEGER DEFAULT 0,
  "is_banned" BOOLEAN DEFAULT false,
  "ban_reason" TEXT,
  "can_use_api_keys" BOOLEAN DEFAULT false
);

ALTER TABLE "users" ADD FOREIGN KEY ("user_type_id") REFERENCES "user_types" ("id");

CREATE TABLE IF NOT EXISTS "user_verification_tokens" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "modified_at" TIMESTAMP,
  "is_archived" BOOLEAN DEFAULT false,
  "token_hash" TEXT NOT NULL,
  "user_id" INTEGER
);

ALTER TABLE "user_verification_tokens" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE TABLE IF NOT EXISTS "user_login_history" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "modified_at" TIMESTAMP,
  "is_archived" BOOLEAN DEFAULT false,
  "user_id" INTEGER
);

ALTER TABLE "user_login_history" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE TABLE IF NOT EXISTS "user_sessions" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "modified_at" TIMESTAMP,
  "is_archived" BOOLEAN DEFAULT false,
  "user_id" INTEGER,
  "ip" TEXT NOT NULL,
  "device_thumbprint" TEXT NOT NULL,
  "session_token" TEXT NOT NULL,
  "expires_at" TIMESTAMP
);

ALTER TABLE "user_sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE TABLE IF NOT EXISTS "user_password_resets" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "modified_at" TIMESTAMP,
  "is_archived" BOOLEAN DEFAULT false,
  "user_id" INTEGER,
  "token" TEXT UNIQUE NOT NULL,
  "expires_at" TIMESTAMP
);

ALTER TABLE "user_password_resets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE TABLE IF NOT EXISTS "user_api_keys" (
  "id" INTEGER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "created_at" TIMESTAMP DEFAULT (NOW()),
  "modified_at" TIMESTAMP,
  "is_archived" BOOLEAN DEFAULT false,
  "user_id" INTEGER,
  "api_key" TEXT UNIQUE NOT NULL
);

ALTER TABLE "user_api_keys" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
