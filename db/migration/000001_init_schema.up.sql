create table "accounts" (
    "id" bigserial primary key,
    "owner" varchar not null,
    "balance" bigint not null,
    "currency" varchar not null,
    "created_at" timestamptz not null default(now())
);

create table "entries" (
    "id" bigserial primary key,
    "account_id" bigint not null,
    "amnount" bigint not null,
    "created_at" timestamptz not null default(now())
);


create table "transfers" (
    "id" bigserial primary key,
    "from_account_id" bigint not null,
    "to_account_id" bigint not null,
    "amnount" bigint not null,
    "created_at" timestamptz not null default(now())
);

Alter table "entries" add FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

Alter table "transfers" add FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id"); 

Alter table "transfers" add FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id"); 

Create index ON "accounts" ("owner");

Create index ON "entries" ("account_id");

Create index ON "transfers" ("from_account_id");

Create index ON "transfers" ("to_account_id");

Create index ON "transfers" ("from_account_id", "to_account_id");
