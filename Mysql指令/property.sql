--�������ݱ�
create table my_default(
name varchar(10) NOT NULL,
age int default 18--��ʾ�������ǰ�ֶ��ڽ������ݲ����ʱ��û���ṩ��Ĭ��Ϊ18
)charset utf8;


insert into my_default(name) values('Tom');

insert into my_default values ('Jack',default);


--�ֶ�����
create table my_comment(
name varchar(10) not null comment '��ǰ���û���������Ϊ��',
pass varchar(50) not null comment '���벻��Ϊ��'
)charset utf8;

