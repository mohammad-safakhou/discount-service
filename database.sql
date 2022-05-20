create table if not exists discount_rule
(
    id INT NOT NULL AUTO_INCREMENT,
    -- max : any discount attribute with maximum value generates discount
    -- group : merchant can create attribute groups. relation inside group is AND and between groups is OR.
    -- groups are scored with op that is starting from 1 to ..., priority starts from first group and goes on.
    -- sorted : each discount attribute marked with a number in op that is just a priority. if the higher one
    -- passes the meta rules can be used in discount value return.
    -- or : OR operand between all attributes.
    -- and : AND operand between all attributes.
    rule_type enum ('max', 'group', 'sorted', 'or', 'and'),
    user_id VARCHAR(36) NOT NULL,
    is_percentage tinyint(1) NOT NULL ,
    is_club_lists tinyint(1) NOT NULL ,
    primary key (id)
);

create table if not exists discount_attributes
(
    id INT NOT NULL AUTO_INCREMENT,
    op varchar(100),
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

create table if not exists point_map
(
    id INT NOT NULL AUTO_INCREMENT,
    discount_rule_id int ,
    name VARCHAR(100) NOT NULL ,
    value VARCHAR(100) NOT NULL ,
    foreign key (discount_rule_id) references discount_rule (id),
    primary key (id)
);
