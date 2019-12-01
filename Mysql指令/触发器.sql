--�����������ű�һ������Ʒ��һ���Ƕ�����������ƷID����ÿ�ζ������ɣ���Ʒ���ж�Ӧ�Ŀ���Ӧ�÷��ͱ仯
--������
create table my_goods(
id int primary key auto_increment,
name varchar(20) not null,
inv int
)charset utf8;

create table my_orders(
id int primary key auto_increment,
goods_id int not null,
goods_num int not null
)charset utf8;

insert into my_goods values(null,'�ֻ�',1000),(null,'����',500),(null,'��Ϸ��',100);

--����������
delimiter $$
create trigger after_insert_order_t after insert on my_orders for each row
begin
       update my_goods set inv=inv-1 where id =1;  
           
end
$$
delimiter ;


--�鿴������
show triggers\G

show create trigger after_insert_order_t\G

--��������ִ��
insert into my_orders values(null,1,1);

--ɾ��������
drop trigger after_insert_order_t;

--�Զ��۳���Ʒ���Ĵ�����
delimiter $$
create trigger a_i_o_t after insert on my_orders for each row
begin
      update my_goods set inv=inv-new.goods_num where id=new.goods_id; 

end
$$
delimiter ;


--�жϿ��
delimiter $$
create trigger b_i_o_t before insert on my_orders for each row
begin
       select inv from my_goods where id =new.goods_id into @inv;
	   if @inv<new.goods_num then
	   insert into XXX values('XXX');
	   end if;
end
$$
delimiter ;
