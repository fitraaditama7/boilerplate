create table IF NOT EXISTS tm_user
(
    id           varchar(255)                          not null,
    username     varchar(255)                          not null,
    email        varchar(255)                          not null,
    password     varchar(255)                          not null,
    first_name   varchar(255)                          not null,
    last_name    varchar(255)                          null,
    phone_number varchar(255)                          null,
    role_id      varchar(255)                          null,
    created_at   timestamp default current_timestamp() not null on update current_timestamp(),
    created_by   varchar(255)                          null,
    updated_at   timestamp                             null,
    updated_by   varchar(255)                          null,
    deleted_at   timestamp                             null,
    deleted_by   varchar(255)                          null,
    CONSTRAINT tm_user_primary PRIMARY KEY (ID),
    constraint tm_user_email_uindex
        unique (email),
    constraint tm_user_id_uindex
        unique (id),
    constraint tm_user_username_uindex
        unique (username)
);

-- alter table tm_user add primary key (id);

