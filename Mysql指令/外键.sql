--�������

create table my_foreign(
id int primary key auto_increment,
name varchar(10)not null,
--�����༶my_class
class_id int,
--�������
foreign key(class_id) references my_class(class_id)
)charset utf8;

create table my_foreign(
id int primary key auto_increment,
name varchar(10)not null,
class_id int,
foreign key(class_id) references my_class(class_id)
)charset utf8;

--�޸�my_student����class_id��Ϊ����ֶ�
alter table my_student add constraint `student_class_ibfk_1` foreign key(class_id) references my_class(class_id);

--ɾ�����
alter table my_student drop foreign key `student_class_ibfk_1`;

--�������ɾ����ͨ��������ֻ��ɾ���Լ�
--�����ɾ����Ӧ��������
alter table my_student drop index `student_class_ibfk_1`;

--��ӱ��������
insert into my_foreign values(null,'С��',1);--����������1
insert into my_foreign values(null,'С��',4);--����û�����ݼ�¼����ʾ����

alter table my_student add foreign key(class_id)
references my_class(class_id)
--Լ��
on update cascade on delete set null;

--��������
update my_class set class_id=4 where class_id=2;--����ģʽ
delete from my_class where class_id=4;--ɾ��ģʽ��ɾ�������ÿ�Ϊnull 


13409