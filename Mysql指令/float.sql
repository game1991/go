create table my_float(
f1 float,
f2 float(10,2)
)charset utf8;


insert into my_float values(123.345,12345678.90);--���ȴ����7λ���ң�����������Ȱ��������������
insert into my_float values(123.345,123456789.00);--���ݳ������ȣ������޷���ʾ����������������볬����ϵͳ�����Զ���λ

create table my_decimal(
f1 float(10,2),
d1 decimal(10,2)
)charset utf8;

insert into my_decimal values(12345678.90,12345678.90);
insert into my_decimal values(99999999.99,99999999.99);--����������������ȣ���������Ҳ�޷���ʾ��ϵͳ�����Զ���λ��һ��Ӧ����Ǯ���������ݷ��棩
