--插入中文数据
insert into my_teacher values('张三',55);
insert into my_teacher values('李四',65);

--mysql.exe告知mysqld.exe自己的字符集规则
set names gbk;

--查看系统保存的三种关系处理字符集
show variables like 'character_set%';

--修改变量
set character_set_client=gbk;

--修改结果字符集
set character_set_results=gbk;
