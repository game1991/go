--�鿴ϵͳ����
select @@autocommit;

--�Ự�޸�ϵͳ����
set autocommit =0;

--ȫ���޸�ϵͳ����
set global autocommit =0;
set @@global.auto_increment_increment=2;--ȫ���޸�ֻ���µĿͻ�����Ч

--ͨ����ѯ����Ϊ������ֵ
select @name :=stu_name,@age :=stu_age from my_student limit 1;
select stu_name,stu_age from my_student order by stu_height desc limit 1 into @name,@age;


