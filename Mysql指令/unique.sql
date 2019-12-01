--唯一键
create table my_unique1(
id int primary key auto_increment,
username varchar(10) unique
)charset utf8;

create table my_unique2(
id int primary key auto_increment,
username varchar(10),
unique key(username)
)charset utf8;

create table my_unique3(
id int primary key auto_increment,
username varchar(10)
)charset utf8;

alter table my_unique3 add unique key(username);

insert into my_unique1 values(null,default);
insert into my_unique1 values(null,default);
insert into my_unique1 values(null,default);

insert into my_unique1 values(null,'army');
insert into my_unique1 values(null,'army');--非空数据无法重复，会报错

--删除唯一键
alter table my_unique1 drop index username;

--修改唯一键：先删除后增加

--复合唯一键
--举例:课程表




