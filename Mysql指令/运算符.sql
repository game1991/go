create table ysf1(
int_1 int,
int_2 int,
int_3 int,
int_4 int)charset utf8;

insert into ysf1 values(100,-100,0,default);

--��������
select int_1+int_2,int_1-int_2,int_1*int_2,int_1/int_2,int_2/int_3,int_2%3 from ysf1;

--mysql��û�й涨select���������ݱ�(mysql���ݱȽ�ʱ����ת����ͬ����)
select '1' <=> 1,0.02<=>0,0.02<>0;
--mysql��û�в���ֵ��ֻ��0��1��1����true��0����false

--������������
select * from my_student where stu_age between 20 and 30;--(between�Ǳ�����)

select * from my_student where stu_age >= 20 and stu_age<= 30;

select * from my_student where stu_height >= 170 or stu_age<= 30;

select * from my_student where stu_id in('stu0001','stu0004','stu0007');

--��ѯ�Ƿ�Ϊ������
select * from my_int where int_6 is null;


