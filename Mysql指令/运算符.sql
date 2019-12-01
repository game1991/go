create table ysf1(
int_1 int,
int_2 int,
int_3 int,
int_4 int)charset utf8;

insert into ysf1 values(100,-100,0,default);

--算术运算
select int_1+int_2,int_1-int_2,int_1*int_2,int_1/int_2,int_2/int_3,int_2%3 from ysf1;

--mysql中没有规定select必须有数据表(mysql数据比较时会先转换成同类型)
select '1' <=> 1,0.02<=>0,0.02<>0;
--mysql中没有布尔值，只有0和1，1代表true，0代表false

--查找年龄区间
select * from my_student where stu_age between 20 and 30;--(between是闭区间)

select * from my_student where stu_age >= 20 and stu_age<= 30;

select * from my_student where stu_height >= 170 or stu_age<= 30;

select * from my_student where stu_id in('stu0001','stu0004','stu0007');

--查询是否为空数据
select * from my_int where int_6 is null;


