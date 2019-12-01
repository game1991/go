insert into my_gbk values('张三'),('李四'),('王五');

create table my_student2(
stu_id varchar(10) primary key comment '主键，学生ID',
stu_name varchar(10) not null comment '学生姓名'
)charset utf8;

insert into my_student2 values
('stu0001','张三'),
('stu0002','张四'),
('stu0003','张五'),
('stu0004','张六');

--主键冲突更新

insert into my_student2 values('stu0004','小明') on duplicate key update stu_name ='小明';

--主键冲突替换(处理效率比更新慢)
replace into my_student2 values('stu0001','夏洛');

--蠕虫复制(一分为二，成倍增加)

create table my_simple(
name char(1) not null
)charset utf8;

insert into my_simple values('a'),('b'),('c'),('d');

--操作蠕虫复制
insert into my_simple(name) select name from my_simple;

--注意：1、没有太大的业务意义，通常用来测试表的效率（索引）
--2、蠕虫复制时需要注意主键冲突



