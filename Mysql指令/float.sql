create table my_float(
f1 float,
f2 float(10,2)
)charset utf8;


insert into my_float values(123.345,12345678.90);--精度大概在7位左右，如果超过精度按照四舍五入计算
insert into my_float values(123.345,123456789.00);--数据超过长度，报错无法显示，如果符合四舍五入超过，系统可以自动进位

create table my_decimal(
f1 float(10,2),
d1 decimal(10,2)
)charset utf8;

insert into my_decimal values(12345678.90,12345678.90);
insert into my_decimal values(99999999.99,99999999.99);--定点数如果超过长度，四舍五入也无法显示，系统不会自动进位（一般应用于钱，银行数据方面）
