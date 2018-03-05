

/* create tables */

create table groups
(
	id bigint unsigned not null auto_increment comment 'id',
	name varchar(255) not null comment '名前',
	created_at datetime not null comment '登録日時',
	created_by varchar(64) not null comment '登録トレース',
	updated_at datetime not null comment '更新日時',
	updated_by varchar(64) not null comment '更新トレース',
	primary key (id)
) comment = 'グループ';


create table group_members
(
	id bigint unsigned not null auto_increment comment 'id',
	is_admin boolean not null comment '管理者か',
	user_id bigint unsigned not null comment 'user_id',
	group_id bigint unsigned not null comment 'グループid',
	created_at datetime not null comment '登録日時',
	created_by varchar(64) not null comment '登録トレース',
	updated_at datetime not null comment '更新日時',
	updated_by varchar(64) not null comment '更新トレース',
	primary key (id)
) comment = 'グループメンバー';


create table users
(
	id bigint unsigned not null auto_increment comment 'id',
	name varchar(255) not null comment '名前',
	birthdate date comment '誕生日',
	gender varchar(10) not null comment '性別 : 男: male
女: female',
	created_at datetime not null comment '登録日時',
	created_by varchar(64) not null comment '登録トレース',
	updated_at datetime not null comment '更新日時',
	updated_by varchar(64) not null comment '更新トレース',
	primary key (id)
) comment = 'users';



/* create foreign keys */

alter table group_members
	add constraint fk_group_members_groups foreign key (group_id)
	references groups (id)
	on update restrict
	on delete restrict
;


alter table group_members
	add constraint fk_group_members_users foreign key (user_id)
	references users (id)
	on update restrict
	on delete restrict
;



