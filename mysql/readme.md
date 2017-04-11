## 环境搭建
1. 启动 mysql服务器  mysql -u root -p, 输入密码123456
2. show global variables like 'port'; 启动客户端连接时链接该port, 默认3306
3. create dabase gc;
4. use gc
5. create table book(
	id bigint(20),
	title varchar(255),	
	content varchar(255),
	pages bigint(20)
	);
6. 可在golang端用程序进行各种操作. (源码部分)
7. drop table if exists book;
