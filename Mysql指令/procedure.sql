--��������
create procedure my_pro1()
select * from my_student;


--�޸���������
delimiter $$
create procedure my_pro2()
begin
     declare i int default 1;
     set @sum =0;
        while i<100 do
        set @sum=@sum+i;
        set i=i+1;
end while;

select @sum;

end
$$
delimiter ;

--���ù��̽��
call my_pro2;

--����3���ⲿ����
set @n1=1;
set @n2=2;
set @n3=3;

--��������
--�޸���������
delimiter $$
create procedure my_pro3(in int_1 int,out int_2 int,inout int_3 int)
begin 
       select int_1,int_2,int_3;
	   set int_1=10;
       set int_2=100;
       set int_3=1000;

       select int_1,int_2,int_3;
	   select @n1,@n2,@n3;

	   set @n1='a';
	   set @n2='b';
	   set @n3='c';
	   select @n1,@n2,@n3;

end
$$
delimiter ;

--���ù���
call my_pro3(@n1,@n2,@n3);