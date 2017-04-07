1. 启动 mysql服务器  mysql -u root -p, 输入密码123456
	(0)show global variables like 'port'; 启动客户端连接时链接该port, 默认3306
	(1)create dabase gc;
	(2)use gc
	(3)create table book(
		id bigint(20),
		title varchar(255),	
		content varchar(255),
		pages bigint(20)
	   );
	(4)可在golang端用程序进行各种操作. (第二部分)
	(5)drop table if exists book;
