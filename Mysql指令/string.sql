create table my_enum(
gender enum('��','Ů','����')
)charset utf8;

insert into my_enum values('��');
insert into my_enum values('Ů');
insert into my_enum values('����');
insert into my_enum values('male');

--���ֶΰ�����ֵ���
select gender +0 from my_enum;
insert into my_enum values(3);


create table my_set(
hobby set('����','��ë��','����','ƹ����','����','�����','����','��ٴ')
--           1       1       1       1       1        1      1      1
)charset utf8;

insert into my_set values('����,ƹ����,����');
--                             10010010
--������ �洢ת��              01001001===> 1+8+64=73
insert into my_set values('��ٴ,����,����,ƹ����,����');
--                             11010011
--�洢ת��                     11001011===> 1+2+8+64+128=203

--����ֵ��ʽ�鿴��������
select hobby +0 from my_set;

insert into my_set values(255);--ͨ����ʵ��

--mysql��洢����

--utf 65535/3 =21845  �������varchar�洢:��Ҫ2��������ֽ������泤��
--gbk 65535/2 =32767 | 1 �������varchar�洢:��Ҫ����2���ֽ�

create table my_utf(
name varchar(21844)
)charset utf8;

create table my_gbk(
name varchar(32766)
)charset gbk;


