--��������(�ѿ�������û��ʵ��Ӧ�ã�����from ��1����2)

select * from my_student cross join my_int;

--������

create table my_class(
id int primary key auto_increment,
name varchar(10) not null unique
)charset utf8;

insert into my_class values(null,'1��'),(null,'2��'),(null,'3��');

select * from my_student inner join my_class;--���ԣ�û���������������ǻ��ɵѿ�����

select * from my_student inner join my_class on class_id =id;

select * from my_student inner join my_class on my_student.class_id =my_class.id;--(Ϊ�˱����ֶ��ظ�ͬ����ͨ��ʹ�ñ���.�ֶ�����ȷ��Ψһ��)
--��������
select * from my_student as s inner join my_class c on s.class_id =c.id;

select * from my_class c inner join my_student as s on s.class_id =c.id;

--������ ��outer join��
select * from my_student as s left join my_class c on s.class_id =c.id;--������
select * from my_student as s right join my_class c on s.class_id =c.id;--������
--�������ӿ��Ի���

--��Ȼ������
select * from my_student natural join my_class;

--using �ؼ��֣�ͬ���ֶΣ�
select * from my_student right join my_class using(class_id);--������



