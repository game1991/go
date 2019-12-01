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
'1900-01-01','12:12:12','1900-01-01 12:12:12','1999-01-01 12:12:12',70);--year为69，划分为2069，70是1970

update my_data set d1='2000-01-01' where d5=2069;--timestamp当对应的数据被修改时，会自动更新为当前时间（这个被修改的数据不是自己的时候）


