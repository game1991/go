--创建一个自动求和的函数（从1开始，直到用户传入的数据为止，去除5的倍数）
--修改语句结束符
delimiter $$
  --创建函数
create function my_sum(end_value int) returns int
begin
     --声明变量（局部变量）：如果使用declare声明变量：必须在函数体其他语句之前
	 declare res int default 0;--这个变量是result
	 declare i int default 1;--这个是自增加的数
	 --循环处理
	 mywhile:while i <=end_value do
             --判断当前数据是否合理（去除5的倍数）
             if i%5=0 then
			 set i=i+1;
			 iterate mywhile; 
             end if;

             --修改变量：累加结果
			 set res = res+i;     --mysql中没有++自增加
			 set i =i+1;
      end while mywhile;
	  --返回值
	  return res;       
end
--结束
$$
--修改语句结束符
delimiter ;

--调用函数
select my_sum(100),my_sum(-100);

set @name ='张三';

create function my_func4() returns char(4)
return @name;


