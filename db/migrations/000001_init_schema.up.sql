CREATE TABLE "accounts" (
    "id"            bigserial       PRIMARY KEY,
    "owner"         varchar         NOT NULL,
    "balance"       numeric(10,2)   NOT NULL,
    "currency"      varchar         NOT NULL,
    "created_at"    timestamptz     NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
    "id"            bigserial       PRIMARY KEY,
    "account_id"    bigint          NOT NULL,
    "amount"        numeric(10,2)   NOT NULL,
    "created_at"    timestamptz     NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
    "id"                bigserial       PRIMARY KEY,
    "from_account_id"   bigint          NOT NULL,
    "to_account_id"     bigint          NOT NULL,
    "amount"            numeric(10,2)   NOT NULL,
    "created_at"        timestamptz     NOT NULL DEFAULT (now())
);


CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_account_id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");


ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD CONSTRAINT "non_negative_amount" CHECK ("amount" > 0);

ALTER TABLE "transfers" ADD CONSTRAINT "diff_from_to_account" CHECK ("from_account_id" != "to_account_id");