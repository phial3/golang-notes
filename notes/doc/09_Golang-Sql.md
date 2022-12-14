# Golang-Sql  Golang的数据库

## 一、SQL语法简介

### 1. MySQL初始配置
- SQL(i/ˈsiːkwəl/; Structured Query Language)是一套语法标准，不区分大小写

- MySQL、sql-server和Oracle都是关系型数据库，在一些高级语法上跟标准SQL略有出入

- MySQL配置
	- Linux安装MySQL客户端   `yum install mysql`
	- 安装MySQL服务端        `yum install mysql-server`
	- 启动MySQL服务端        `systemctl start mysqld.service`
	- 以root登录             `mysql -uroot`
		- SQL管理员创建账号      `create user 'tester' identified by '123456';`
		- 查看账号创建是否成功   `select host, user from mysql.user where user='tester';`
		- 赋予账号对应权限
```sql
	// grant <privileges> on <database>.<table> to 'tester'@'localhost';
	grant create, insert on *.* to 'tester'; 
	// 这里表示赋予该用户所有数据库所有表(*.*表示所有表)，%表示所有IP地址
	grant all privileges on *.* to '用户名'@'%' identified by '密码' with grant option;
```
	
- MySQL管理	
```sql
	// 以tester登录
	`mysql -utester -p'123456' -h121.40.150.39 -P9528`

	// 创建database
	`create database test;`
	
	// 使用database  
	use test;
	show tables;

	// 创建表
	create table if not exists student(
	id int not null auto_increment comment '主键自增id',
	name char(4) not null comment '姓名',
	province char(6) not null comment '省',
	city char(10) not null comment '城市',
	addr varchar(100) default '' comment '地址',
	score float not null default 0 comment '考试成绩',
	enrollment date not null comment '入学时间',
	primary key (id),  unique key idx_name (name),  
	key idx_location (province,city)
	)default charset=utf8 comment '学员基本信息';

	show variables like 'innodb_large_prefix';
	show variables like 'innodb_file_format';
	set global innodb_large_prefix=1;
	set global innodb_file_format=BARRACUDA;
	CREATE TABLE `resource` (
	`id` char(64) CHARACTER SET latin1 NOT NULL,
	`vendor` tinyint(1) NOT NULL,
	`region` varchar(64) CHARACTER SET latin1 NOT NULL,
	`zone` varchar(64) CHARACTER SET latin1 NOT NULL,
	`create_at` bigint(13) NOT NULL,
	`expire_at` bigint(13) DEFAULT NULL,
	`category` varchar(64) CHARACTER SET latin1 NOT NULL,
	`type` varchar(120) CHARACTER SET latin1 NOT NULL,
	`instance_id` varchar(120) CHARACTER SET latin1 NOT NULL,
	`name` varchar(255) NOT NULL,
	`description` varchar(255) DEFAULT NULL,
	`status` varchar(255) CHARACTER SET latin1 NOT NULL,
	`update_at` bigint(13) DEFAULT NULL,
	`sync_at` bigint(13) DEFAULT NULL,
	`sync_accout` varchar(255) CHARACTER SET latin1 DEFAULT NULL,
	`public_ip` varchar(64) CHARACTER SET latin1 DEFAULT NULL,
	`private_ip` varchar(64) CHARACTER SET latin1 DEFAULT NULL,
	`pay_type` varchar(255) CHARACTER SET latin1 DEFAULT NULL,
	`describe_hash` varchar(255) NOT NULL,
	`resource_hash` varchar(255) NOT NULL,
	PRIMARY KEY (`id`),
	KEY `name` (`name`) USING BTREE,
	KEY `status` (`status`) USING HASH,
	KEY `private_ip` (`private_ip`) USING BTREE,
	KEY `public_ip` (`public_ip`) USING BTREE,
	KEY `instance_id` (`instance_id`) USING HASH
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

	// 查看索引 [Specified key was too long](https://blog.csdn.net/shaochenshuo/article/details/51064685)
	show index from student\G;

	// 新增记录
	// 必须给`not null`且无 default值的列赋值
	insert into student (name,province,city,enrollment) values
	('张三','北京','北京','2021-03-05'),
	('李四','河南','郑州','2021-04-25'),
	('小丽','四川','成都','2021-03-10');

	// 查询记录
	select id,name from student where id>0;
	select province,avg(score) as avg_score from student 
		where score>0 
		group by province having avg_score>50 
		order by avg_score desc;
	select * from student order by c1 asc, c2 desc;    // 先按列 c1 升序排列，再按 c2 降序排列

	// 修改记录
	update student set score=score+10,addr='海淀' where province='北京';
	update student set
		score=case province
			when '北京' then score+10     
			when '四川' then score+5 
			else score+7
		end,
		addr=case province
			when '北京' then '东城区'        
			when '四川' then '幸福里'        
			else '朝阳区'    
		end
	where id>0;

	// 删除记录
	delete from student where city= '郑州';
	// 删除表里的所有行
	delete from student;
	// 删除表
	drop table student;     
```

### 2. MySQL初始配置
- 注意事项
	- 注意存储引擎的选择(InnoDB)
	- 主键选择 和 唯一键考虑清楚
	- 考虑数据类型与长度，选择合适的类型，避免空间浪费
	- 字符串注意确认字符集，如果需要存入中文，请选择utf8编码
	- 为过滤条件的字段 添加索引
	- 高频组合查询可以考虑 联合索引
	- 注意选择使用索引的方法: `Hash` `Btree` `Normal`
	- 写sql时一律使用小写
	- 建表时先判断表是否已存在 `if not exists` 
	- 所有的列和表都加 `comment`
		- `comment` 是备注、注释的意思，写上 `comment 'id'` 之后，在建表信息里可以看到添加的备注信息
	- 字符串长度比较短时尽量使用char，定长有利于内存对齐，读写性能更好，而varchar字段频繁修改时容易产生内存碎片
	- 满足需求的前提下尽量使用短的数据类型，如tinyint vs int, float vs double, date vs datetime
	- null
		- default null有别于 default '' 和 default 0
		- is null, is not null有别于 != '', !=0
		- 尽量设为 not null
			- 有些DB索引列不允许包含null
			- 对含有null的列进行统计，结果可能不符合预期
			- null值有时候会严重拖慢系统性能

- B+树
	- B即Balance，对于m叉树每个节点上最多有m个数据，最少有m/2个数据(根节点除外)
	- 叶节点上存储了所有数据，把叶节点链接起来可以顺序遍历所有数据
	- 每个节点设计成内存页的整倍数；MySQL的 m=1200，树的前两层放在内存中
```
	            2  28 65
	            P1 P2 P3
	    |          |         |
	2  13 20   28 35 56   65 80 90
	P1 P2 P3   P1 P2 P3   P1 P2 P3
	|  |  |    |  |  |    |  |  | 
	2  13 20   28 35 56   65 80 90
	8  15 23   30 38 60   73 85 96
	9  19 27   32 50 63   79 88 99
	Q  Q  Q    Q  Q  Q    Q  Q  Q   // data 
```

- 索引
	- MySQL索引默认使用B+树
		- 散列表(Hash table，也叫哈希表) 与 B+树
			- Hash table 查询时间为 O(1)，但是其对范围查询的支持不如 B+树
			- 即Hash table只支持等于或不等于，不支持关键词检索
	- 主键默认会加索引
		- 按主键构建的B+树里包含所有列的数据，而普通索引的B+树里只存储了主键，还需要再查一次主键对应的B+树(回表)
		- 使用 explain命令 查看一个SQL语句的执行计划，如使用的索引，是否做全表扫描等
	- 联合索引的前缀同样具有的索引的效果
	- sql语句前加explain可以查看索引使用情况
	- 如果MySQL没有选择最优的索引方案，可以在 where前 `force index (index_name)`
	- 规避慢查询
		- 大部分的慢查询都是因为没有正确地使用索引
		- 一次select不要超过1000行
		- 分页查询limit m,n 会检索前m+n行，只是返回后n行，通常用id>x来代替这种分页方式
		- 批量操作时最好一条sql语句搞定；其次打包成一个事务，一次性提交(高并发情况下减少对共享资源的争用)
		- 不要使用连表操作，join逻辑在业务代码里完成

### 3. Go语言中SQL驱动接口
- `database/sql`
	- Go官方没有提供数据库驱动，而是为开发数据库驱动定义了一些标准接口(即database/sql)，开发者可以根据定义的接口来开发相应的数据库驱动
	- Go语言中支持MySQL的驱动比较多，如
		- github.com/go-sql-driver/mysql  支持 database/sql
		- github.com/ziutek/mymysql       支持 database/sql，支持自定义的接口
		- github.com/Philio/GoMySQL     不支持 database/sql，支持自定义的接口

- Driver
	- `driver.Driver`
	- 注册数据库驱动
	- 打开数据库连接
	- Conn
		- 把一个查询 query传给Prepare，返回 Stmt(statement)
		- Close关闭数据库连接
		- Begin返回一个事务 Tx(transaction)
	- Stmt
		- Close关闭当前的链接状态
		- NumInput返回当前预留参数的个数
		- Exec执行Prepare准备好的 sql，传入参数执行 update/insert 等操作，返回 Result 数据
		- Query执行Prepare准备好的 sql，传入需要的参数执行 select 操作，返回 Rows 结果集	
	- Tx
		- Commit提交事务
		- Rollback回滚事务
	- Result
		- LastInsertId返回由数据库执行插入操作得到的自增ID号
		- RowsAffected返回操作影响的数据条目数
		- RowsAffected
			- RowsAffected是int64的别名，它实现了Result接口
				- `type RowsAffected int64`
				- `func (RowsAffected) LastInsertId() (int64, error)`
				- `func (v RowsAffected) RowsAffected() (int64, error)
	- Rows
		- Columns是查询所需要的表字段
		- Close关闭迭代器
		- Next返回下一条数据，把数据赋值给dest，dest里面的元素必须是driver.Value的值
			- 如果最后没有数据，Next 函数返回 `io.EOF`
	- Value
	- `driver.ValueConverter`
		- 把数据库里的数据类型转换成Value允许的数据类型
```go
	// driver.Driver
	package driver // import "database/sql/driver"
	type Driver interface { 
		Open(name string) (Conn, error)                        // func Open(driverName, dataSourceName string) (*DB, error)
	}

	// 注册数据库驱动
	var d = Driver{proto: "tcp", raddr: "127.0.0.1:3306"}
	sql.Register("mysql", &d)                                  // func Register(name string, driver driver.Driver)

	// 打开数据库连接
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8")

	// driver.Conn
	type Conn interface {
		Prepare(query string) (Stmt, error)                    // func (db *DB) Prepare(query string) (*Stmt, error)
		Close() error                                          // func (c *Conn) Close() error
		Begin() (Tx, error)                                    // func (db *DB) Begin() (*Tx, error)
	}

	// driver.Stmt
	type Stmt interface {
		Close() error                                          // func (s *Stmt) Close() error
		NumInput() int
		Exec(args []Value) (Result, error)                     // func (s *Stmt) Exec(args ...interface{}) (Result, error)
		Query(args []Value) (Rows, error)                      // func (s *Stmt) Query(args ...interface{}) (*Rows, error)
	}

	// driver.Tx
	type Tx interface {
		Commit() error                                          // func (tx *Tx) Commit() error
		Rollback() error                                        // func (tx *Tx) Rollback() error
	}

	// driver.Result
	type Result interface {
		LastInsertId() (int64, error)
		RowsAffected() (int64, error)
	}

	// driver.Rows
	type Rows interface {
		Columns() []string                             // func (rs *Rows) Columns() ([]string, error)
		Close() error                                  // func (rs *Rows) Close() error
		Next(dest []Value) error                       // func (rs *Rows) Next() bool
	}

	// driver.Value
	type Value interface{}
		nil           // 要么是 nil，要么是下面的任意一种
		int64 
		float64 
		bool 
		[]byte 
		string 
		time.Time

	// driver.ValueConverter
	type ValueConverter interface {
		ConvertValue(v interface{}) (Value, error)
	}
```

### 4. Go语言中 数据库的操作
- 增删改查
	- Go语言中的第三方库 `go get github.com/go-sql-driver/mysql`
	- 连接数据库 `db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8")`
	- 增删改 `func (*sql.DB).Exec(sql string) (sql.Result, error)`
	- 查 `func (*sql.DB).Query(sql string) (*sql.Rows, error)`

- Stmt
	- 参数化查询 `db.Where("merchant_id = ?", merchantId)`
	- 拼接sql `db.Where(fmt.Sprintf("merchant_id = %s", merchantId))`
	- 定义一个sql模板 `stmt, err := db.Prepare("update student set score=score+? where city=?")`
	- 多次使用模板 `res, err := stmt.Exec(10, "上海")`

### 5. 数据库使用问题
- SQL注入
	- 问题现象
		- `sql = "select username,password from user where username='" + username + "' and password='" + password + "'";`
			- 变量username和password从前端输入框获取，如果用户输入的username为lily， password为aaa' or '1'='1
			- 则完整的sql为select username,password from user where username='lily' and password='aaa' or '1'='1'
			- 会返回表里的所有记录，如果记录数大于0就允许登录，则lily的账号被盗
		- `sql = "insert into student (name) values ('"+username+"')";`
			- 变量username从前端输入框获取，如果用户输入的username为lily'); drop table student;--
			- 完整sql为insert into student (name) values ('lily'); drop table student;--')
			- 通过注释符--屏蔽掉了末尾的')，删除了整个表
	- 预防措施
		- 前端输入要加正则校验、长度限制
		- 对特殊符号`(<>&*; '"等)`进行转义或编码转换，Go的 text/template包里面的HTMLEscapeString函数可以对字符串进行转义处理
		- 不要将用户输入直接嵌入到sql语句中，而应该使用参数化查询接口，如Prepare、Query、Exec(query string, args ...interface{})
		- 使用专业的SQL注入检测工具进行检测，如sqlmap、SQLninja
		- 避免网站打印出SQL错误信息，以防止攻击者利用这些错误信息进行SQL注入

- SQL预编译
	- DB执行SQL分为3步
		- 词法和语义解析
		- 优化 SQL 语句，制定执行计划
		- 执行并返回结果
	- SQL 预编译技术是指将用户输入用占位符?代替，先对这个模板化的sql进行预编译，实际运行时再将用户输入代入
	- 除了可以防止 SQL 注入，还可以对预编译的SQL语句进行缓存，之后的运行就省去了解析优化SQL语句的过程

### 6. ORM与NoSQL技术
- SQLBuilder
	- Go语言中的第三方库
		- Go-SQLBuilder 是一个用于创建SQL语句的工具函数库，提供一系列灵活的、与原生SQL语法一致的链式函数，归属于艾润物联公司 `go get -u github.com/parkingwang/go-sqlbuilder`
		- Gendry是一个用于辅助操作数据库的Go包，基于go-sql-driver/mysql，它提供了一系列的方法来为调用标准库database/sql中的方法准备参数 `go get –u github.com/didi/gendry`
	- 自行封装SQL构建器
		- 写一句很长的sql容易出错，且出错后不好定位
		- 函数式编程可以直接定位到是哪个函数的问题
		- 函数式编程比一长串sql更容易编写和理解
		- `github.com/phial3/golang-notes/blob/main/week12/database/self_sql_builder/main.go`
```go
	// Go-SQLBuilder 函数链
	sql := gsb.NewContext().
		Select("id", "name", "score", "city").
		From("student").
		OrderBy("score").DESC().                  // 按"score"降序
		Column("name").ASC().                     // 当"score"相同，按"name"升序
		Limit(10).Offset(20).                     // 从第20个开始，读10个 limit 20, 10
		ToSQL()

	// Gendry
	//  map
	where := map[string]interface{}{
		"city": []string{"北京", "上海", "杭州"},
		"score<": 30,
		"addr": builder.IsNotNull,
		"_orderby": "score desc",
	}
	fields := []string{"id", "name", "city", "score"}
	_, _, err := builder.BuildSelect("student", where, fields)
```

- ORM技术与GORM
	- ORM
		- ORM 即 Object Relational Mapping，对象关系映射
		- Relational指各种sql类的关系型数据为
		- Object指面向对象编程(object-oriented programming)中的对象
		- ORM在数据库记录和程序对象之间做一层映射转换，使程序中不用再去编写原生SQL，而是面向对象的思想去编写类、对象、调用相应的方法来完成数据库操作
	- GORM
		- `go get -u gorm.io/gorm`
		- `go get -u gorm.io/driver/mysql`
		- GORM是一个全能的、友好的、基于golang的ORM库
		- GORM 倾向于约定，而不是配置
			- 默认情况下，GORM 使用 ID 作为主键，使用结构体名的 蛇形复数 作为表名，字段名的 蛇形 作为列名，并使用 `CreatedAt`、`UpdatedAt` 字段追踪创建、更新时
```go
		// 完全是在操作struct，看不到sql的影子
		dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"   // data source name
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		student := Student{Name: "光绪", Province: "北京", City: "北京", Score: 38, Enrollment: time.Now()}
		db.Create(&student)
```

### 7. MongoDB
- MongoDB相关概念
	- NoSQL泛指非关系型数据库，如 mongodb, redis, HBase
	- mongo使用高效的二进制数据存储，文件存储格式为BSON(一种json的扩展，比json性能更好，功能更强大)
	- MySQL中表的概念在mongo里叫集合(collection)， MySQL中行的概念在mongo中叫文档(document)，一个文档看上去像一个json

- MongoDB的安装配置
```shell
	// 安装MongoDB
	$ vim /etc/yum.repos.d/mongodb-org-4.2.repo
		[mongodb-org-4.2] 
		name=MongoDB Repository
		baseurl=https://repo.mongodb.org/yum/redhat/$releasever/mongodb-org/4.2/x86_64/
		gpgcheck=1
		enabled=1
		gpgkey=https://www.mongodb.org/static/pgp/server-4.2.asc
	$ yum install -y mongodb-org
	$ systemctl start mongod.service
	
	// MongoDB常用命令
	$ mongo
	use test;                                                                                // 切换到test库，如果没有则创建
	db.createUser({user: "tester", pwd: "123456", roles: [{role: "dbAdmin", db: "test"}]});  // 创建用户
	db.createCollection("student");                                                          // 创建collection
	db.student.createIndex({"name":1});                                                      // 在name上创建索引,不是唯一索引
	db.student.insertOne({name:"张三",city:"北京"});
	db.student.find({name:"张三"});
	db.student.update({name:"张三"},{name:"张三",city:"上海"});
	db.student.deleteOne({name:"张三"});
```

- Go语言中 MongoDB的操作
```go
	// 安装mongo-driver
	go get go.mongodb.org/mongo-driver
	go get go.mongodb.org/mongo-driver/x/bsonx/bsoncore@v1.7.1
	go get go.mongodb.org/mongo-driver/x/mongo/driver@v1.7.1
	go get go.mongodb.org/mongo-driver/mongo/options@v1.7.1
	go get go.mongodb.org/mongo-driver/x/mongo/driver/topology@v1.7.1
	go get go.mongodb.org/mongo-driver/mongo@v1.7.1
	
	// 连接MongoDB
	option := options.Client().ApplyURI("mongodb://127.0.0.1:27017").
	SetConnectTimeout(time.Second).                                                         // 连接超时时长
	SetAuth(options.Credential{Username: "tester", Password: "123456", AuthSource: "test"}) // 指定用户名和密码，AuthSource代表Database
	client, err := mongo.Connect(context.Background(), option)
	// 注意: Ping成功才代表连接成功
	err = client.Ping(ctx, nil) 

	// 查询MongoDB
	sort := bson.D{{"name", 1}}                                 // 1升序，-1降序
	filter := bson.D{{"score", bson.D{{"$gt", 3}}}} //score>3   // greater than
	findOption := options.Find()
	findOption.SetSort(sort)                                    // 按name排序
	findOption.SetLimit(10)                                     // 最多返回10个
	findOption.SetSkip(3)                                       // 跳过前3个
	cursor, err := collection.Find(ctx, filter, findOption)
```