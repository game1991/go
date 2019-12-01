--创建视图
create view student_class_v as 
select s.*,c.name from my_student as s left join
my_class as c
on
s.class_id =c.class_id;

--修改视图
alter view student_class_v as 
select * from my_student as s left join
my_class as c
using(class_id);

--删除视图
drop view student_class_v;