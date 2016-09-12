***
* @ryanyzyang 
* To be familier with GO grammar through the Leetcode algorithm

***

#### 1. Go语言中没有的语法 
    a. Go语言中没有while，要使用for{}来实现，中间使用if来进行结束
    b. Go语言中没有三元操作a>b?a:b
    c. Go语言中没有++操作符，要使用i=i+1
    d. Go语言没有->的指针操作符，指针与一般变量一样可以用`.`来访问变量和函数
    e. Go语言中空指针使用nil来表示

#### 2. Go中实现排序
    a. Go的sort库中自带有基本变量的排序，eg Int型的sort.Ints(array)
    b. 对于自定义类型，需要实现一个排序类，用于对其排序，其中要实现三个方法: Len(),Swap(i,j int),Less(i,j int) bool
