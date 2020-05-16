create table companies(
       id bigserial,
       "name" varchar(300) unique not null,
       ticker varchar(10) unique not null,
)

create table ten_ks(
       user_id bigint,
       "year" smallint not null,

       -- All 10-K information

       unique(user_id, "year")
)
