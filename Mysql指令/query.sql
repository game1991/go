--select ѡ��

select all * from my_simple;--�ȼ���* from

--��ѯ�ظ��ֶ�
select distinct * from my_simple;

--�ֶα���(alias)
select distinct name as name1,name name2,name name3,name name4,name name5 from my_simple;


--��̬���ݱ���
select * from(select int_1,int_6 from my_int) as int_my;

--����(group by�鿴ֻ����ÿ�����ݵĵ�һ�����ݣ�����ͳ������)
select * from my_student group by class_id;

--ʹ�þۺϺ�����count��avg��sum��max��min��,���հ༶ͳ��ÿ��������������䣬���ߣ�ƽ������
select class_id,count(*),max(stu_age),min(stu_height),avg(stu_age) from my_student group by class_id;

select class_id,group_concat(stu_name),count(*),max(stu_age),min(stu_height),avg(stu_age) from my_student group by class_id;

--�����
select class_id,group_concat(stu_name) from my_student group by class_id,gender;

--�༶�����Ա���
select class_id,gender,count(*),group_concat(stu_name) from my_student group by class_id asc,gender desc;

--��������(rollup)
select class_id,gender,count(*),group_concat(stu_name) from my_student group by class_id,gender with rollup;

--Having�Ӿ䣨��ѯ�༶����4�����ϵģ�
select class_id,count(*) as number from my_student group by class_id having count(*) >=4;

select class_id,count(*) as number from my_student group by class_id having number >=4;
--ע�⣺having �� group by֮��group by����where֮��where��ʱ���ʾ�����ݴӴ����õ��ڴ棬where֮������в��������ڴ����

--order by�༶ѧ�������������
select * from my_student order by stu_height asc;

select * from my_student order by class_id desc,stu_height;

--limit ʹ��(�÷���limit offset��㣬length����)
select * from my_student limit 2;

--��ҳ��ȡ����
select * from my_student limit 0,2;
