--利用mysqldump进行sql备份

--整库备份
mysqldump.exe -hlocalhost -P3306 -uroot -pgame mydatabase2 > C:\ProgramData\MySQL\temp\mydatabase2.sql

--多表备份
mysqldump -uroot -pgame mydatabase2 my_student my_int > C:\ProgramData\MySQL\temp\student_int.sql

--还原数据
mysql -uroot -pgame mydb < C:\ProgramData\MySQL\temp\mydatabase2.sql

source C:\ProgramData\MySQL\temp\student_int.sql;