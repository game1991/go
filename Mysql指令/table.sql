--�������ݱ�
--�����ݱ�ҵ����ݿ���
create table mydatabase2.class(
--�ֶ��� �ֶ�����
--�ַ������ʵ�Ƿֲ�����
name varchar(10)
--10���ַ�(���ܳ���)
);


--�������ݿ⣬������
use mydatabase2;
create table teacher(
name varchar(10)
);


--ʹ�ñ�ѡ��
create table student(
name varchar(10)
)charset utf8;

--�����ݿ�mydatabase�´���һ����teacherһ���ı�
use mydatabase;

create table teacher like mydatabase2.teacher;


--�鿴���б�
show tables;

--�鿴ƥ�����ݱ�
show tables like 'c%';


--��ʾ��ṹ
describe class;

desc teacher;

show columns from student;

--��ʾ��һ����ͼ
show create table student\G


--�޸ı�ѡ��
alter table student charset gbk;

--���ݿ������ݱ�����ͨ����ǰ׺��ȡ���ݿ��ǰ������ĸ�����»���
rename table student to my_student;

--�����ֶ�
alter table my_student add column age int after name;

alter table my_student add  id int first;

--�޸��ֶ���
alter table my_student change age nj int;

--�޸��ֶ�����
alter table my_student modify name varchar(20);

--ɾ���ֶ�
alter table my_student drop ���;

--ɾ����
drop table class;

--����ɾ������

drop table teacher,my_student;