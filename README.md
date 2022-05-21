## 单机版
业务流：输入网址(url string)->爬取网站(fetcher)->解析内容(parser)->数据入库(write db )

 + 网址数据来源->页面输入/数据库(任务队列)
 + fetcher/parser-> worker
 + result->入库(tasks进度展示,采集结果res展示) 

## 问题列举
### 1.1 程序run异常，无任何输出
+ solution: dlv 调试
+ reason: 数据库连接失败，需要进行数据库连接断言 dbx.ping()

### 1.2 程序连接时无法正常加载配置
+ 修改为toml
+ 使用全局conf.Conf来进行获取
   + var cfg = conf.Conf 定义时候，在初始化时候赋值了值，但是后续未同步加载后的值(值赋值)
   + 会导致使用cfg时候为空
   
### 1.3 listenandserver提示为空指针
+ solution ：查看源代码，发现server.start需要异步执行。在 Shutdown 时会立刻返回，Shutdown 方法会阻塞至所有连接闲置或 context 完成
+ reason: program 方法需要传递指针类型，否则p.server为空