--����һ���Զ���͵ĺ�������1��ʼ��ֱ���û����������Ϊֹ��ȥ��5�ı�����
--�޸���������
delimiter $$
  --��������
create function my_sum(end_value int) returns int
begin
     --�����������ֲ������������ʹ��declare���������������ں������������֮ǰ
	 declare res int default 0;--���������result
	 declare i int default 1;--����������ӵ���
	 --ѭ������
	 mywhile:while i <=end_value do
             --�жϵ�ǰ�����Ƿ����ȥ��5�ı�����
             if i%5=0 then
			 set i=i+1;
			 iterate mywhile; 
             end if;

             --�޸ı������ۼӽ��
			 set res = res+i;     --mysql��û��++������
			 set i =i+1;
      end while mywhile;
	  --����ֵ
	  return res;       
end
--����
$$
--�޸���������
delimiter ;

--���ú���
select my_sum(100),my_sum(-100);

set @name ='����';

create function my_func4() returns char(4)
return @name;


