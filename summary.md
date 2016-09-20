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
    f. Go语言中没有char类型，与之等价的是byte

#### 2. Go中实现排序
    a. Go的sort库中自带有基本变量的排序，eg Int型的sort.Ints(array)
    b. 对于自定义类型，需要实现一个排序类，用于对其排序，其中要实现三个方法: Len(),Swap(i,j int),Less(i,j int) bool

#### 3. Go指针容易出现的问题
    a. var ptr *Node 是正确的指针声明方式,注意*号在前面（不像C语言是在后面）
    b. var ptr *Node仅仅是声明，是没分配空间的，如果想分配空间还需要使用new, eg: var ptr = new(Node)

#### 4. Go的标准输入

    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadBytes('\n')
    fmt.Println(string(input[0 : len(input)-1])) // string(input[0:len(input)-1]) remove '\n'.
#### 5. Import库的时候可以用()括号，多个库之间没有标点符号，用换行即可

#### 6. Go语言里面动态生成数组使用make, eg: make([]string,len) len可以是`常量`也可以是`变量`
