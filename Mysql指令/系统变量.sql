--查看系统变量
select @@autocommit;

--会话修改系统变量
set autocommit =0;

--全局修改系统变量
set global autocommit =0;
set @@global.auto_increment_increment=2;--全局修改只对新的客户端生效

--通过查询数据为变量赋值
select @name :=stu_name,@age :=stu_age from my_student limit 1;
select stu_name,stu_age from my_student order by stu_height desc limit 1 into @name,@age;


