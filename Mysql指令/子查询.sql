--知道一个学生的名字：小花，想知道他在哪个班级（班级名字）

--通过学生表获取他所在班级ID
--通过班级ID获取班级名字

--标量子查询实现
select * from my_class where class_id = (select class_id from my_student where stu_name ='小花');

--想获取已有学生所在班的所有班级名字

--1、找出学生表中的所有的班级ID
--2、找出班级表中对应的名字

--列子查询实现
select name from my_class where class_id in (select class_id from my_student);

--行子查询实现（一行多列）
--获取班级上年龄最大，且身高最高的学生
--1、求出年龄最大的值
--2、求出班级身高最大的值
--3、求出对应的学生
select * from my_student having stu_age=max(stu_age) and stu_height =max(stu_height);
--1、having 是在group by之后：使用having代表前面group by执行了一次（聚合函数使用）
--2、group by一旦执行：结果就是只返回一行记录：第一行

select * from my_student where(stu_age,stu_height)=(select max(stu_age),max(stu_height) from my_student);

--获取每个班级身高最高的学生（一个）
select * from my_student group by class_id having stu_height =max(stu_height);

--1、将每个班级身高的学生排在最前面：order by
--2、再针对结果进行group by:保留每组第一个
select * from (select * from my_student order by stu_height desc) as temp group by class_id;

--求出有学生在的所有班级
select * from my_class as c where exists(select stu_id from my_student as s where s.class_id =c.class_id);


select * from my_class where class_id in (select class_id from my_student);
--等价于select * from my_class where class_id =any(select class_id from my_student);