--�������ݵ����ݱ�

--�����ȴ���һ�����ݱ�
 create table my_teacher(name varchar(10),age int)charset utf8;


insert into my_teacher(name,age) values('game',28);

insert into my_teacher(age,name) values(49,'tom');

insert into my_teacher(name) values('jack');

insert into my_teacher values('levi',30);
--ֵ�б�������ֶ��б�һ��


---��ȡ��������
select * from my_teacher;

---��ȡָ���ֶ�����//�ֶ��б��ö��Ÿ���
select name from my_teacher;

---��ȡ����Ϊ30�������
select name from my_teacher where age = 30;

--ɾ������Ϊ28����ʦ����
delete from my_teacher where age = null;
--ɾ���ظ�����ʱ
DELETE FROM table_name WHERE column_name1 = '' AND column_name2 = '' LIMIT 1;

---��������Ϊnull������//���û��where��������ȫ������
update my_teacher set age =27 where name ='jack';

--�������ݵ�ʱ�����û��������ȫ�����ݸ��£�������limit����ʾ���µ�����
update my_simple set name ='6' where name ='a' limit 4;
