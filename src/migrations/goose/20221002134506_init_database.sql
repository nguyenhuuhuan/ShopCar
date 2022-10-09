-- +goose Up
CREATE table users
(
    id 		        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    username        varchar(50) DEFAULT NULL,
    email           varchar(50) NOT NULL,
    owner           varchar(50) DEFAULT NULL,
    password        varchar(100) NOT NULL,
    full_name       varchar(50) DEFAULT NULL,
    phone_number    varchar(50) DEFAULT NULL,
    dob             varchar(50) DEFAULT NULL,
    provider        varchar(50) NOT NULL,
    status          enum('ACTIVE', 'INACTIVE') NOT NULL,
    created_at 		timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at 		timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

CREATE table user_roles
(
    id 		        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    user_id         bigint(20) unsigned NOT NULL,
    role_id         bigint(20) unsigned NOT NULL,
    created_at 		timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at 		timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

CREATE table roles
(
    id 		        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    role_name       varchar(50) NOT NULL,
    status          enum('ACTIVE', 'INACTIVE') NOT NULL,
    code            varchar(50) DEFAULT NULL,
    created_at 		timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at 		timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
-- +goose Down
DROP table users;
DROP table roles;
DROP table user_roles;
