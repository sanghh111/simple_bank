Create table "session"(
    "uuid" text primary key,
    "username" text not null,
    "refesh_token" text not null,
    "agent_user" varchar not null,
    "client_ip" varchar unique not null, 
    "expried_at" timestamptz not null,
    "created_at" timestamptz not null default (now())
);

Alter table "session" add FOREIGN key ("username") REFERENCES "users" ("username");