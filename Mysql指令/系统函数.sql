--�ж��ַ������ַ���
select char_length('��ã�����');
--�ж��ַ������ֽ���(���ַ���)
select length('��ã��й�');
--�����ַ���
select concat('��','��');
--�ж��ַ���Ŀ���ַ������Ƿ���ڣ�����0��1
select instr('����й�','��'),instr('����й�','��');
--ȫ��Сдlcase
--����࿪ʼ��ȡ�ַ���ֱ��ָ��λ��left
select lcase('a'),left('�������',2),mid('�������',2);
--������߶�Ӧ�Ŀո�ltrim
--���м�ָ��λ�ÿ�ʼ��ȡ�������ָ����ȡ���ȣ�ֱ�ӵ����mid

--now():���ص�ǰʱ��
--curdate():���ص�ǰ����
--curtime():���ص�ǰʱ��
--datediff():�ж���������֮��������࣬�������ڱ���ʹ���ַ�����ʽ�������ţ�
--date_add(���ڣ�interval ʱ������ type)������ʱ������� Type:day/hour/minute/second

select datediff('2010-10-10','1990-10-10');
select date_add('2010-10-10',interval 10 year),date_add('2010-10-10',interval 10 hour);

select unix_timestamp();--ʱ���
select from_unixtime(1562254364);--��ʱ���ת��������ʱ��

--abs()����ֵ
--ceiling()����ȡ��
--floor()����ȡ��
--pow()��ָ����˭�Ķ��ٴη�
--rand()���һ�������
--round()��������
select abs(-2),ceiling(1.3),floor(1.1),pow(2,4),rand(),round(1.58);

--md5():�����ݽ���md5����
--version():��ȡ�汾��
--database():��ʾ�������ݿ�
--uuid():����һ��Ψһ�ı�ʶ���������������������ǵ���Ψһ��UUID�����⣨����Ψһͬʱ�ռ�Ψһ��

select md5('x'),version(),database(),uuid();
