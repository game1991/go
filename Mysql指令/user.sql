--创建用户

create user 'user1'@'%' identified by '123456';

--简化创建信息
create user2;

--删除用户
drop user user2@host;--(也可以不带host)

--修改用户密码
set password for 'user1'@'%'=password('654321');

update mysql.user set password = password('654321') where user='user1' and host='%';

--分配权限
grant select on mydb.my_student to 'user1'@'%';
revoke all privileges on mydb.my_student from 'user1'@'%';

--刷新权限
flush privileges;
