--select 选项

select all * from my_simple;--等价于* from

--查询重复字段
select distinct * from my_simple;

--字段别名(alias)
select distinct name as name1,name name2,name name3,name name4,name name5 from my_simple;


--动态数据别名
select * from(select int_1,int_6 from my_int) as int_my;

--分组(group by查看只保留每组数据的第一个数据，用于统计作用)
select * from my_student group by class_id;

--使用聚合函数（count、avg、sum、max、min）,按照班级统计每班人数，最大年龄，最矮身高，平均年龄
select class_id,count(*),max(stu_age),min(stu_height),avg(stu_age) from my_student group by class_id;

select class_id,group_concat(stu_name),count(*),max(stu_age),min(stu_height),avg(stu_age) from my_student group by class_id;

--多分组
select class_id,group_concat(stu_name) from my_student group by class_id,gender;

--班级升序，性别降序
select class_id,gender,count(*),group_concat(stu_name) from my_student group by class_id asc,gender desc;

--回溯排序(rollup)
select class_id,gender,count(*),group_concat(stu_name) from my_student group by class_id,gender with rollup;

--Having子句（查询班级人数4个以上的）
select class_id,count(*) as number from my_student group by class_id having count(*) >=4;

select class_id,count(*) as number from my_student group by class_id having number >=4;
--注意：having 在 group by之后，group by是在where之后，where的时候表示将数据从磁盘拿到内存，where之后的所有操作都是内存操作

--order by班级学生按照身高排序
select * from my_student order by stu_height asc;

select * from my_student order by class_id desc,stu_height;

--limit 使用(用法；limit offset起点，length长度)
select * from my_student limit 2;

--分页获取数据
select * from my_student limit 0,2;
