create table my_data(
d1 date,
d2 time,
d3 datetime,
d4 timestamp,
d5 year
)charset utf8;

insert into my_data values (
'1900-01-01','12:12:12','1900-01-01 12:12:12','1999-01-01 12:12:12',89);
insert into my_data values (
'1900-01-01','12:12:12','1900-01-01 12:12:12','1999-01-01 12:12:12',2020);
insert into my_data values (
'1900-01-01','12:12:12','1900-01-01 12:12:12','1999-01-01 12:12:12',70);--yearΪ69������Ϊ2069��70��1970

update my_data set d1='2000-01-01' where d5=2069;--timestamp����Ӧ�����ݱ��޸�ʱ�����Զ�����Ϊ��ǰʱ�䣨������޸ĵ����ݲ����Լ���ʱ��


