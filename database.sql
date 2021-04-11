create table if not exists discount_rule
(
    id INT NOT NULL AUTO_INCREMENT,
    user_id VARCHAR(36) NOT NULL,
    is_percentage tinyint(1) NOT NULL ,
    is_club_lists tinyint(1) NOT NULL ,
    primary key (id)
);

create table if not exists discount_attributes
(
    id INT NOT NULL AUTO_INCREMENT,
    discount_rule_id int ,
    name VARCHAR(300) NOT NULL,
    foreign key (discount_rule_id) references discount_rule (id),
    primary key (id)
);

create table if not exists discount_meta_attributes
(
    id INT NOT NULL AUTO_INCREMENT,
    discount_attributes_id int,
    name VARCHAR(300) NOT NULL ,
    value VARCHAR(300) NOT NULL ,
    point int NOT NULL,
    foreign key (discount_attributes_id) references discount_attributes (id),
    primary key (id)
);
