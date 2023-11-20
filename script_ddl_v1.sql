CREATE database simple_bank;

Create table "users"(
    "username" varchar primary key,
    "hash_password" varchar not null,
    "firt_name" varchar not null,
    "last_name" varchar not null,
    "full_name" varchar not null,
    "email" varchar unique not null, 
    "password_changed_at" timestamptz not null default '0001-01-01 00:00:00Z',
    -- note update_at store when the change happen 
    --> Example: Change first name 
    "update_at" timestamptz not null default '0001-01-01 00:00:00Z',
    "created_at" timestamptz not null default (now())
);
--> BẢNG USER LÀ ĐỂ ĐỊNH DANH NGƯỜI DÙNG (MỤC DÍCH KHÁC THÌ TẠO BẢNG KHÁC)


create table "accounts" (
    "id" bigserial primary key,
    "owner" varchar unique not null,
    "balance" bigint not null DEFAULT 0 ,
    -- note: lưu số tiền hiện tại của account 
    "type" bigint not null DEFAULT 0 , 
    -- note that the one user can have more than one account 
    -- type account default is Zero 
    -- configure for type account should be constant 
    --> Example: 0-Debit Account (tài khoản trả sau), 1- Credit Account (tài khoản trả trước), 3- Mix Account ...
    "currency" varchar not null default 'VND',
    --> Example: United States Dollar: USD
    "status" varchar not null default 'A',
    -- Trạng thái của 1 account trong hiện tại 
    -- Example: A-active, K-locked, B-backlisted
    "update_at" timestamptz not null default '0001-01-01 00:00:00Z',
    "created_at" timestamptz not null default(now())
);
--> BẢNG ACCOUNTS LÀ ĐỂ ĐỊNH DANH TÀI KHOẢN CỦA NGƯỜI DÙNG 
-- một người dùng có thể có nhiều tài khoản với các các tài khoản có thể có policy khác nhau
-- một tài khoản không thể có nhiều user
 Alter table "accounts" add FOREIGN key ("owner") REFERENCES "users" ("username");

-- nếu 2 users có chung đơn vị tiền tệ (the same currency) ==> có thể insert TRANSFERS 

-- nếu 2 users không có chung đơn vị tiền tệ (the different currency) 
-- ==> thì phải CÓ CHUNG currency (phải đổi về chung đơn vị tiền tệ 
--> dựa theo đơn vị tiền tệ hiện tại của người được nhận để đổi)
-- ==> có thể insert TRANSFERS 

-- Trạng thái của 1 Account được quy định là  A-active, U-unknown, B-backlisted
-- Một account phải có trạng thái rõ ràng, nếu user muốn khóa tài khoản -> tha đổi trạng thái K-locked
-- Trường hợp nếu account bị đánh cắp -> user có thể thay đổi trạng thải account là K-locked 
-- khi đó thì account ko thể insert trong bảng TRANSFERS

-- Thêm trường hợp: Trạng thái của 1 Account được quy định là  D-deleted 
-- Một user muốn xóa tài khoản vĩnh viễn 
	--> Account update status = D and account amount = 0.00 (trả lại tiền cho người dùng) 


Alter table "accounts" add constraint "owner_curency_key" unique ("owner", "currency");
-- (MỤC DÍCH KHÁC THÌ TẠO BẢNG KHÁC)
-- DONE 2 TABLES USER AND ACCOUNTS


create table "TRANSFERS" (
    "id" bigserial primary key,
    "from_account_id" bigint not null,
    "to_account_id" bigint not null,
    "amnount" bigint not null,
    "status" varchar not null, 
    "reason" varchar not null,
    "created_at" timestamptz not null default(now())
);
--BẢNG TRANSFERS dùng để lưu thông tin chung của giao dịch 
-- user A chuyển tiền cho user B -> có thể thành công hoặc thất bại
--> cả hai trường hợp đều insert 1 record vào TRANSFERS

	--> nếu giao dịch thành công thì status là S-successfull
		--> nếu TRANSFERS-from_account_id  
			--> update amount vào bảng Account-balance= Account-balance - TRANSFERS-amount
		--> nếu TRANSFERS-to_account_id  
			--> update amount vào bảng Account-balance= Account-balance + TRANSFERS-amount

	--> nếu giao dịch không thành công thì status là F-fail
		--> lưu lý do giao dịch thất bại vào reason
-- (MỤC DÍCH KHÁC THÌ TẠO BẢNG KHÁC)

create table "SUBTYPE" ( 
    "level" varchar not null,
    "type" varchar not null,
    "descript" varchar not null,
    "amount" varchar not null,
    "rate" varchar not null,
    "freq" varchar not null,
    "update_at" timestamptz not null default '0001-01-01 00:00:00Z',
    "created_at" timestamptz not null default (now())
);
-- note cặp SUBTYPE-level và SUBTYPE-type là unique nên bảng SUBTYPE không cần _id
	-- SUBTYPE-level : B-bank ; U-user
	-- SUBTYPE-type :  TT-TRANFER TO ; TF-TRANFER FROM ; UF-USER BANK FEE PAID, UC-USER CREATE ACCOUNT 

-- SUBTYPE-descript lưu mô tả về SUBTYPE-level và SUBTYPE-type 
-- note	amount and rate and freq  lưu configure cố định cho 1 policy cụ thể  
--  amount and rate and freq lưu số tiền
	-- EXAMPLE : mở thẻ cần nạp 50.000 VND -> SUBTYPE-level-U SUBTYPE-type UC SUBTYPE-amount 50.000 

--  rate lưu số tỷ lệ 
	-- EXAMPLE : account fee -> SUBTYPE-level-B SUBTYPE-type UF SUBTYPE-rate 10%

--  freq lưu số chu kì 
	-- EXAMPLE : freq 1 tháng -> SUBTYPE-level-B SUBTYPE-type UF SUBTYPE-rate 10% freq = monthly

create table "ENTRIES" ( 
    "id" bigserial primary key,
    "account_id" bigint not null,
    "amnount" bigint not null,
    "subtype" varchar not null,
    "created_at" timestamptz not null default(now())
);

--BẢNG ENTRIES dùng để lưu thông tin toàn bộ của giao dịch của user (tính cả chuyển tiền, nạp tiền, trả phí)
-- Khi có một action xảy ra đối vz user -> bất kì 1 giao dịch -> insert vào ENTRIES 

	-- EXAMPLE : record lưu subtype= "U TT"  ENTRIES-account_id chuyển tiền cho user khác 
									-- and amnount = amount lưu trong TRANFRES


Alter table "entries" add FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

Alter table "transfers" add FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id"); 

Alter table "transfers" add FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id"); 

Create index ON "accounts" ("owner");

Create index ON "entries" ("account_id");

Create index ON "transfers" ("from_account_id");

Create index ON "transfers" ("to_account_id");

Create index ON "transfers" ("from_account_id", "to_account_id");





DROP TABLE IF EXISTS "entries";
DROP TABLE IF EXISTS "transfers";
DROP TABLE IF EXISTS "accounts";
