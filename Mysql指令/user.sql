--�����û�

create user 'user1'@'%' identified by '123456';

--�򻯴�����Ϣ
create user2;

--ɾ���û�
drop user user2@host;--(Ҳ���Բ���host)

--�޸��û�����
set password for 'user1'@'%'=password('654321');

update mysql.user set password = password('654321') where user='user1' and host='%';

--����Ȩ��
grant select on mydb.my_student to 'user1'@'%';
revoke all privileges on mydb.my_student from 'user1'@'%';

--ˢ��Ȩ��
flush privileges;
