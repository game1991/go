--创建数据表
create table my_default(
name varchar(10) NOT NULL,
age int default 18--表示，如果当前字段在进行数据插入的时候没有提供，默认为18
)charset utf8;


insert into my_default(name) values('Tom');

insert into my_default values ('Jack',default);


--字段描述
create table my_comment(
name varchar(10) not null comment '当前是用户名，不能为空',
pass varchar(50) not null comment '密码不能为空'
)charset utf8;

