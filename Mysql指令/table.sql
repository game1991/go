--创建数据表
--将数据表挂到数据库下
create table mydatabase2.class(
--字段名 字段类型
--字符与表其实是分不开的
name varchar(10)
--10个字符(不能超过)
);


--进入数据库，创建表
use mydatabase2;
create table teacher(
name varchar(10)
);


--使用表选项
create table student(
name varchar(10)
)charset utf8;

--在数据库mydatabase下创建一个与teacher一样的表
use mydatabase;

create table teacher like mydatabase2.teacher;


--查看所有表
show tables;

--查看匹配数据表
show tables like 'c%';


--显示表结构
describe class;

desc teacher;

show columns from student;

--显示不一样视图
show create table student\G


--修改表选项
alter table student charset gbk;

--数据库中数据表名字通常有前缀：取数据库的前两个字母加上下划线
rename table student to my_student;

--新增字段
alter table my_student add column age int after name;

alter table my_student add  id int first;

--修改字段名
alter table my_student change age nj int;

--修改字段类型
alter table my_student modify name varchar(20);

--删除字段
alter table my_student drop 年纪;

--删除表
drop table class;

--批量删除表名

drop table teacher,my_student;