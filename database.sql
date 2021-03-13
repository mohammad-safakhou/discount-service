create table if not exists discount
(
    id int primary key,
    user_id VARCHAR(36) NOT NULL
);

create table if not exists discount_attributes
    (
        id int primary key,
        discount_id int ,
        name VARCHAR(300) NOT NULL,
        foreign key (discount_id) references discount (id)
);

create table if not exists discount_meta_attributes
    (
        id int primary key,
        discount_attributes_id int,
        name VARCHAR(300) NOT NULL ,
        value VARCHAR(300) NOT NULL ,
        foreign key (discount_attributes_id) references discount_attributes (id)
);

create database discount_service;
use discount_service;
drop database discount_service;
