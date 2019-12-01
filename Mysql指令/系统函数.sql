--判断字符串的字符数
select char_length('你好，世界');
--判断字符串的字节数(与字符集)
select length('你好，中国');
--连接字符串
select concat('你','好');
--判断字符在目标字符串中是否存在，返回0和1
select instr('你好中国','中'),instr('你好中国','我');
--全部小写lcase
--从左侧开始截取字符串直到指定位置left
select lcase('a'),left('你好世界',2),mid('你好世界',2);
--消除左边对应的空格ltrim
--从中间指定位置开始截取，如果不指定截取长度，直接到最后mid

--now():返回当前时间
--curdate():返回当前日期
--curtime():返回当前时间
--datediff():判断两个日期之间的天书差距，参数日期必须使用字符串格式（用引号）
--date_add(日期，interval 时间数字 type)：进行时间的增加 Type:day/hour/minute/second

select datediff('2010-10-10','1990-10-10');
select date_add('2010-10-10',interval 10 year),date_add('2010-10-10',interval 10 hour);

select unix_timestamp();--时间戳
select from_unixtime(1562254364);--将时间戳转换成日期时间

--abs()绝对值
--ceiling()向上取整
--floor()向下取整
--pow()求指数，谁的多少次方
--rand()获得一个随机数
--round()四舍五入
select abs(-2),ceiling(1.3),floor(1.1),pow(2,4),rand(),round(1.58);

--md5():对数据进行md5加密
--version():获取版本号
--database():显示所在数据库
--uuid():生成一个唯一的标识符（自增长）；自增长是单表唯一，UUID是整库（数据唯一同时空间唯一）

select md5('x'),version(),database(),uuid();
