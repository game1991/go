--增加外键

create table my_foreign(
id int primary key auto_increment,
name varchar(10)not null,
--关联班级my_class
class_id int,
--增加外键
foreign key(class_id) references my_class(class_id)
)charset utf8;

create table my_foreign(
id int primary key auto_increment,
name varchar(10)not null,
class_id int,
foreign key(class_id) references my_class(class_id)
)charset utf8;

--修改my_student表，将class_id设为外键字段
alter table my_student add constraint `student_class_ibfk_1` foreign key(class_id) references my_class(class_id);

--删除外键
alter table my_student drop foreign key `student_class_ibfk_1`;

--外键不能删除普通的索引，只能删除自己
--如果想删除对应的索引：
alter table my_student drop index `student_class_ibfk_1`;

--向从表插入数据
insert into my_foreign values(null,'小明',1);--主表有数据1
insert into my_foreign values(null,'小李',4);--主表没有数据记录，提示报错

alter table my_student add foreign key(class_id)
references my_class(class_id)
--约束
on update cascade on delete set null;

--更新主表
update my_class set class_id=4 where class_id=2;--更新模式
delete from my_class where class_id=4;--删除模式，删除数据置空为null 


13409