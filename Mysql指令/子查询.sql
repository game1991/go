--֪��һ��ѧ�������֣�С������֪�������ĸ��༶���༶���֣�

--ͨ��ѧ�����ȡ�����ڰ༶ID
--ͨ���༶ID��ȡ�༶����

--�����Ӳ�ѯʵ��
select * from my_class where class_id = (select class_id from my_student where stu_name ='С��');

--���ȡ����ѧ�����ڰ�����а༶����

--1���ҳ�ѧ�����е����еİ༶ID
--2���ҳ��༶���ж�Ӧ������

--���Ӳ�ѯʵ��
select name from my_class where class_id in (select class_id from my_student);

--���Ӳ�ѯʵ�֣�һ�ж��У�
--��ȡ�༶����������������ߵ�ѧ��
--1�������������ֵ
--2������༶�������ֵ
--3�������Ӧ��ѧ��
select * from my_student having stu_age=max(stu_age) and stu_height =max(stu_height);
--1��having ����group by֮��ʹ��having����ǰ��group byִ����һ�Σ��ۺϺ���ʹ�ã�
--2��group byһ��ִ�У��������ֻ����һ�м�¼����һ��

select * from my_student where(stu_age,stu_height)=(select max(stu_age),max(stu_height) from my_student);

--��ȡÿ���༶�����ߵ�ѧ����һ����
select * from my_student group by class_id having stu_height =max(stu_height);

--1����ÿ���༶��ߵ�ѧ��������ǰ�棺order by
--2������Խ������group by:����ÿ���һ��
select * from (select * from my_student order by stu_height desc) as temp group by class_id;

--�����ѧ���ڵ����а༶
select * from my_class as c where exists(select stu_id from my_student as s where s.class_id =c.class_id);


select * from my_class where class_id in (select class_id from my_student);
--�ȼ���select * from my_class where class_id =any(select class_id from my_student);