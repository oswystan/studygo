##golang建议
- 非常好的参考文档 [Effective go](https://golang.org/doc/effective_go.html)
- 注意成对原则，打开的文件、socket等都要用defer关闭
- 在项目中对error进行统一管理，在一个独立的文件中存放所有的error
- 使用go vendor特性对第三方依赖包进行统一存放，并提供url，版本信息，另外第三方包可能也会有外部依赖，把所有依赖统一存放到一个目录下，这样便于管理。


##语法注意事项
- defer func()表达式的值先被计算，然后再将func()压栈，在外层函数返回前压栈的defer函数以后进先出的顺序被调用。


##学习资源
- [一个非常牛叉的博客](http://dave.cheney.net/)
	* [channel讲解](http://dave.cheney.net/2013/04/30/curious-channels)
	* [空struct讲解](http://dave.cheney.net/2014/03/25/the-empty-struct)

