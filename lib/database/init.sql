create table if not exists employee
(
    user_id   int auto_increment primary key,
    name      varchar(255)                  not null,
    sex       enum ('男', '女')             not null,
    phone     int(20)                       not null,
    email     varchar(255)                  not null,
    position  varchar(255)                  not null,
    marry     enum ('已婚', '未婚', '离婚') not null,
    education varchar(10)                   not null,
    join_time timestamp default CURRENT_TIMESTAMP not null
);