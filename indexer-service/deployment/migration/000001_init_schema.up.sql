CREATE TABLE "eth_blocks" (
  "id" BIGSERIAL PRIMARY KEY,
  "number" bigint NOT NULL DEFAULT 0,
  "hash_hex" varchar(256) NOT NULL DEFAULT '',
  "time" bigint NOT NULL DEFAULT 0,
  "parent_hash_hex" varchar(256) NOT NULL DEFAULT '',
  "is_stable" smallint NOT NULL DEFAULT 0,
  "created_on" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_on" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   CONSTRAINT "eth_block_hash_hex_unique" UNIQUE ("hash_hex")
);

CREATE TABLE "eth_transactions" (
  "id" BIGSERIAL PRIMARY KEY,
  "block_hash_hex" varchar(256) NOT NULL DEFAULT '',
  "hash_hex" varchar(256) NOT NULL DEFAULT '',
  "from_hex" varchar(256) NOT NULL DEFAULT '',
  "to_hex" varchar(256) NOT NULL DEFAULT '',
  "nonce" bigint NOT NULL DEFAULT 0,
  "tx_data" bytea,
  "value" varchar(256) NOT NULL DEFAULT 0,
  "created_on" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "eth_tx_hash_hex_unique" UNIQUE ("hash_hex")
);

CREATE TABLE "eth_transaction_logs" (
  "id" BIGSERIAL PRIMARY KEY,
  "tx_hash_hex" varchar(256) NOT NULL DEFAULT '',
  "log_index" integer NOT NULL DEFAULT 0,
  "log_data" bytea,
  "created_on" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX ON "eth_blocks" ("number");

CREATE INDEX ON "eth_transactions" ("hash_hex");
