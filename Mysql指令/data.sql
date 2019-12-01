--插入数据到数据表

--首先先创建一个数据表
 create table my_teacher(name varchar(10),age int)charset utf8;


insert into my_teacher(name,age) values('game',28);

insert into my_teacher(age,name) values(49,'tom');

insert into my_teacher(name) values('jack');

insert into my_teacher values('levi',30);
--值列表必须与字段列表一致


---获取所有数据
select * from my_teacher;

---获取指定字段数据//字段列表用逗号隔开
select name from my_teacher;

---获取年龄为30岁的名字
select name from my_teacher where age = 30;

--删除年龄为28的老师数据
delete from my_teacher where age = null;
--删除重复数据时
DELETE FROM table_name WHERE column_name1 = '' AND column_name2 = '' LIMIT 1;

---更新年龄为null的数据//如果没有where条件，则全部更新
update my_teacher set age =27 where name ='jack';

--更新数据的时候，如果没有条件是全表数据更新，可以用limit来显示更新的数量
update my_simple set name ='6' where name ='a' limit 4;
