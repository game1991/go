insert into my_gbk values('����'),('����'),('����');

create table my_student2(
stu_id varchar(10) primary key comment '������ѧ��ID',
stu_name varchar(10) not null comment 'ѧ������'
)charset utf8;

insert into my_student2 values
('stu0001','����'),
('stu0002','����'),
('stu0003','����'),
('stu0004','����');

--������ͻ����

insert into my_student2 values('stu0004','С��') on duplicate key update stu_name ='С��';

--������ͻ�滻(����Ч�ʱȸ�����)
replace into my_student2 values('stu0001','����');

--��渴��(һ��Ϊ�����ɱ�����)

create table my_simple(
name char(1) not null
)charset utf8;

insert into my_simple values('a'),('b'),('c'),('d');

--������渴��
insert into my_simple(name) select name from my_simple;

--ע�⣺1��û��̫���ҵ�����壬ͨ���������Ա��Ч�ʣ�������
--2����渴��ʱ��Ҫע��������ͻ



