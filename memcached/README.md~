
说明: ubuntu14.04 下测试
1. 首先启动服务器 memcached -d -m 128 -p 11211 -u root
2. 测试New/Get/Set/Delete/Add/Increment/Decrement等操作

memcached 原理详解:
1. 为什么需要memcached, memcached解决了什么问题?
	关系型数据库处理耗时, 高并发读写无法承受
	将查寻结果保存在memcached中, 下次直接从memcached中读取, 可减少数据库操作

2. 内存分配算法 Slab Allocation 
	按预先规定的大小, 将内存分配成特定长度的块(chunk).(解决内存碎片问题)
	尺寸相同的块组成chunk的集合 slab Class, 每次分配给slab 1 page(1M), slab 将其分割成若干trunk, 根据收到数据的大小, 选择分配最合适的slab
	memcached中保存着slab内空闲trunk的列表, 根据该列表选择trunk, 并将数据保存其中
	缺点: 会造成内存浪费, 如下若需分配129字节, 则浪费127字节
slab class   1: chunk size    128 perslab  8192
slab class   2: chunk size    256 perslab  4096
slab class   3: chunk size    512 perslab  2048
slab class   4: chunk size   1024 perslab  1024
	使用Growth Factor进行调优, 减少过多浪费 memcached -f 1.25 -vv
slab class   1: chunk size     88 perslab 11915
slab class   2: chunk size    112 perslab  9362
slab class   3: chunk size    144 perslab  7281
slab class   4: chunk size    184 perslab  5698



3. 内存过期移除策略 Lazy Expiration + LRU
Lazy Expiration:
	memcached 不会监视记录是否过期, 而是在get时查看记录的时间戳.
	优点:memcached不会在过期检测上浪费CPU时间
	缺点:无用数据长期占据内存

LRU:
	Memcached的LRU算法针对每个Slab执行，而不是针对整体。数据只会存在指定的Slab中，若该Slab已经满了，即使更大的Slab有空间，指定的Slab依然执行LRU算法，因为数据不会被存放到更大的Slab中

有效缓解使用LRU的方法是：
1，避免大对象
	如果系统上只有及个别几个大对象的话，会浪费内存空间，因为Slab申请了Page是不能释放内存的，及个别大对象会导致Slab申请了内存资源而得不到充分的利用。
2，调整增长因子
	根据项目的需求调整增长因子，使内存充分利用。

4. 分布式算法
	memcached采用的是 : CRC($KEY)%N
		求得key的整数HASH值 模 服务器台数, 根据余数选择服务器
		优点: 余数计算方法简单, 数据分散性优秀, 适用于服务器数量固定的情况使用
		缺点: 添加/移除服务器时, 缓存映射重组的代价巨大

	一致性HASH(Consistent Hashing):
		(1) 将32位整数(0 - (2^32-1))作为环形HASH空间
		(2) 通过HASH函数将KEY处理成整数, 在环上找到对应位置
		(3) 将Memcached服务器群映射到环上 (使用HASH函数处理服务器对应的IP地址)
		(4) 从当前KEY的位置, 延环顺时针方向, 查找最近的Memcached服务器, 并将KEY对应的数据保存在此服务器上.

MySQL memcached:












