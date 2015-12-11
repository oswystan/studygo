## 2015年12月9日

###目标
- imc项目客户端消息接收与分发机制实现
- imc代码整改
	- 目录名与文件名
	- 包名、函数名、对象名
	- 资源的分配和释放检查

## 2015年12月10日

### 目标
- 学习如何使用pprof工具对代码和程序进行内存泄露分析
- 使用pprof对imc项目进行内存泄露分析
- 学习如何对golang中的内存资源进行管理：申请的内存、chan、打开的文件、网络连接

### 总结
- 对于需要频繁申请和释放的资源，需要建立资源池，这样可以防止由于系统gc而性能下降
- map比较浪费系统内存资源（每个初始状态的map要占用150B），如果可以的话，可以使用slice来替代
- channel用完要close掉；
- pprof有两种使用方式：
	- 程序运行一次就退出，需要在代码中添加prof的代码（请参考`runtime/pprof`的使用说明）
```
	go tool pprof -text $appname $profilename
```

	- 程序作为一个服务一直运行不退出，这样需要在代码中开启一个http服务，然后通过（请参考`net/http/pprof`包的使用说明）：
```
	go tool pprof -text $appname http://localhost:$port/debug/pprof/heap 
```
- golang的内存GC时机是在内存增长到2的幂次方的时候，通过测试发现，如果先申请，然后将指针置空，golang是不会启动GC的，除非手动调用`runtime.GC()`函数。
- golang貌似没有栈内存，只有堆

## 2015年12月11日
### 目标
- 改造之前的MicroSerice接口，以让其适应多客户端并发访问的场景；
- 学习POSA3中的资源管理模式，以便在当前的IMC项目中可以使用；

### 总结


























