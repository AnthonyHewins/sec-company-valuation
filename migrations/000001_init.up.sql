create table companies(
       id bigserial,
       "name" varchar(300) unique not null,
       ticker varchar(10) unique not null,
)

create table company_10ks(
       user_id bigint references companies(id) on delete cascade not null,
       ten_k_id bigint references ten_k(id) not null,
       unique (user_id, ten_k_id)
)

create table ten_ks(
       id bigserial,
       year int not null,
)
