--����mysqldump����sql����

--���ⱸ��
mysqldump.exe -hlocalhost -P3306 -uroot -pgame mydatabase2 > C:\ProgramData\MySQL\temp\mydatabase2.sql

--�����
mysqldump -uroot -pgame mydatabase2 my_student my_int > C:\ProgramData\MySQL\temp\student_int.sql

--��ԭ����
mysql -uroot -pgame mydb < C:\ProgramData\MySQL\temp\mydatabase2.sql

source C:\ProgramData\MySQL\temp\student_int.sql;