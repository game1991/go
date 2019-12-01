--交叉连接(笛卡尔积，没有实际应用，本质from 表1，表2)

select * from my_student cross join my_int;

--内连接

create table my_class(
id int primary key auto_increment,
name varchar(10) not null unique
)charset utf8;

insert into my_class values(null,'1班'),(null,'2班'),(null,'3班');

select * from my_student inner join my_class;--可以，没有设置条件，但是会变成笛卡尔积

select * from my_student inner join my_class on class_id =id;

select * from my_student inner join my_class on my_student.class_id =my_class.id;--(为了避免字段重复同名，通常使用表名.字段名来确保唯一性)
--别名表名
select * from my_student as s inner join my_class c on s.class_id =c.id;

select * from my_class c inner join my_student as s on s.class_id =c.id;

--外连接 （outer join）
select * from my_student as s left join my_class c on s.class_id =c.id;--左连接
select * from my_student as s right join my_class c on s.class_id =c.id;--右连接
--左右连接可以互换

--自然内连接
select * from my_student natural join my_class;

--using 关键字（同名字段）
select * from my_student right join my_class using(class_id);--右连接



