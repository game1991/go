--������ͼ
create view student_class_v as 
select s.*,c.name from my_student as s left join
my_class as c
on
s.class_id =c.class_id;

--�޸���ͼ
alter view student_class_v as 
select * from my_student as s left join
my_class as c
using(class_id);

--ɾ����ͼ
drop view student_class_v;