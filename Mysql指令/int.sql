--�������ݱ�
create table my_int(
int_1 tinyint,
int_2 smallint,
int_3 mediumint,
int_4 int,
int_5 bigint
)charset utf8;

--��������
insert into my_int values(10,10000,100000,1000000,100000000000);
--tinyint�������ǣ�-128��127���ϼ�256

alter table my_int add int_6 tinyint unsigned first;
insert into my_int values (255,100,255,255,255,255);
insert into my_int values (-128,100,255,255,255);
insert into my_int values (1,1,1,1,1,1);

alter table my_int add int_7 tinyint zerofill first;
alter table my_int add int_8 tinyint(2) zerofill first;
--��ʾ���ȿ����Լ��趨���������ȣ����ǲ�������Χ������Ӱ�죬ֻ��Բ����ĳ��Ƚ��в��䣨��ʾ���ȣ�

insert into my_int values (1,1,1,1,1,1,1);
insert into my_int values (1,1,1,1,1,1,1,1);


