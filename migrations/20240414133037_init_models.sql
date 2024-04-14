-- +goose Up
-- +goose StatementBegin

create table  rates
(
    code       text not null,
    nominal    bigint,
    kopecks    bigint,
    original   text,
    date       timestamp not null,
    created_at timestamp,
    updated_at timestamp,
    primary key (code, date)
);

create index rate_code_date on rates(code, date);

create table users
(
    user_id          bigint primary key,
    default_currency text
);

create table expenses
(
    id         bigserial primary key,
    user_id    bigint,
    title      text,
    amount     bigint,
    date       timestamp
);

create index expenses_user_id on expenses(user_id);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop index rate_code_date;
drop table rates;

drop table users;

drop index expenses_user_id;
drop table expenses;

-- +goose StatementEnd
