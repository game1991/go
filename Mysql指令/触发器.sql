--举例：有两张表，一张是商品表，一张是订单表（保留商品ID），每次订单生成，商品表中对应的库存就应该发送变化
--创建表
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

insert into my_goods values(null,'手机',1000),(null,'电脑',500),(null,'游戏机',100);

--创建触发器
delimiter $$
create trigger after_insert_order_t after insert on my_orders for each row
begin
       update my_goods set inv=inv-1 where id =1;  
           
end
$$
delimiter ;


--查看触发器
show triggers\G

show create trigger after_insert_order_t\G

--触发启动执行
insert into my_orders values(null,1,1);

--删除触发器
drop trigger after_insert_order_t;

--自动扣除商品库存的触发器
delimiter $$
create trigger a_i_o_t after insert on my_orders for each row
begin
      update my_goods set inv=inv-new.goods_num where id=new.goods_id; 

end
$$
delimiter ;


--判断库存
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
