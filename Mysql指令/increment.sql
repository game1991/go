--����һ�����Զ������ı�

create table my_auto(
id int primary key auto_increment,
name varchar(10) not null comment '�û���',
pass varchar(50) not null comment '����'
)charset utf8;

insert into my_auto values(null,'Tom','123456');

--�޸�auto_increment
alter table my_auto auto_increment = 10;

--����auto_increment
truncate my_auto;--(�ָ���Ĭ��ֵ)

--ɾ��������
alter table my_auto modify id int;
--�мǲ�Ҫ�ٴ����� primary key

--�鿴��������ʼ����
show variables like 'auto_increment%';

--����������
alter table my_auto modify id int auto_increment;

insert into my_auto values(3,'jack','123456');