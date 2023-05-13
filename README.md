# dlu-design-system

```
.
│  main.go 主程序
│      
├─config 配置层
│      config.go 配置读取
│      config.toml 配置文件
│      
├─dal 数据访问层
│  │   init.go 数据访问层初始化
│  │  
│  └─db gorm-gen 数据库操作
│  │      
│  └─model gorm-gen 数据库模型
│          
├─handler 处理层
│  │  
│  └─info 信息模块
│  │      
│  └─user 用户模块
│      
├─log 日志文件
│     
├─middleware 中间件
│  │  
│  └─cors 跨域
│  │      
│  ├─jwt 鉴权
│  │      
│  ├─limiter 令牌桶限流
│  │      
│  └─redis 缓存
│  
├─pkg 公用模块
│  │    
│  ├─consts 常量
│  │    
│  ├─errno 自定义错误
│  │      
│  ├─logger 日志模块
│  │      
│  ├─result 结果类
│  │      
│  ├─types 请求校验类
│  │      
│  └─utils
│      gradeUtil.go 获取年级 如20 21 22
│      jwtUtil.go 解析token
│      validateUtil.go 校验
│  
├─router 路由层
│  
└─service 服务层
   │  
   └─info 信息模块
   │      
   ├─log 日志模块
   │      
   └─user 用户模块

```