#golang建议
- 非常好的参考文档 [Effective go](https://golang.org/doc/effective_go.html)
- 注意成对原则，打开的文件、socket等都要用defer关闭


#语法注意事项
- defer func()表达式的值先被计算，然后再将func()压栈，在外层函数返回前压栈的defer函数以后进先出的顺序被调用。
- 