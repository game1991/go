--union���ѡ�� Ĭ�ϱ�ʾȥ����ȫ�ظ����ݣ�distinct��


--��ȡ�����������Ů����߽���
(select * from my_student where gender ='��' order by stu_height asc limit 8)
union
(select * from my_student where gender ='Ů' order by stu_height desc limit 8);

select * from my_student
union all
select * from my_student;

select stu_id,stu_name,stu_height from my_student
union all
select stu_height,stu_id,stu_name from my_student;