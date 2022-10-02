-- +goose Up

CREATE table users
(
    id 		        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    username        varchar(50) NOT NULL,
    password        varchar(50) NOT NULL,
    full_name       varchar(50) DEFAULT NULL,
    dob             varchar(50) DEFAULT NULL,
    provider        enum('FACEBOOK', 'GOOGLE', 'NORMAL') DEFAULT 'NORMAL',
    status          enum('ACTIVE', 'INACTIVE') NOT NULL,
    created_at 		timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at 		timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

CREATE table roles
(
    id 		        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    user_id         bigint(20) unsign NOT NULL,
    role_id         bigint(20) unsign NOT NULL,
    created_at 		timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at 		timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
)

CREATE table role_user
(
    id 		        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    role_name       varchar(50) NOT NULL,
    status          enum('ACTIVE', 'INACTIVE') NOT NULL,
    created_at 		timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at 		timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
)
-- +goose Down
DROP table users;
DROP table roles;
DROP table user_role;
