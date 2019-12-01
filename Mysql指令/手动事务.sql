--手动事务

--开启事务
start transaction;

--事务提交
commit;

--回滚（清空之前的操作）
rollback;

--增加回滚点
savepoint sp1;

--回滚指定点
rollback to sp1;


