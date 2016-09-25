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
    1> 方法一（类似C语言中的scanf）
    fmt.Scanf("%d %f %s",&ii,&ff,&ss) 输入字符串时，不能包含空格
    2> 方法二（相当于输入输入一个字符串）
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadBytes('\n')
    fmt.Println(string(input[0 : len(input)-1])) // string(input[0:len(input)-1]) remove '\n'.
    
#### 5. Import库的时候可以用()括号，多个库之间没有标点符号，用换行即可

#### 6. Go语言里面动态生成数组使用make, eg: make([]string,len) len可以是`常量`也可以是`变量`

#### 7. Go语言中的最值都放在math包中，int的最大值、最小值分别为: math.MaxInt32 math.MinInt32

#### 8. Go语言的强制类型转换: `Type`(value) eg: int64(value) 其中int64可以带括号也可以不带，但是变量一定要带（自己经常错）

#### 9. Go语言定义struct的时候，里面的成员变量只用指明类型即可，不用带var，var只在声明变量的时候使用（与C++不同）

#### 10. Go语言的for语句的三个部分不能有多个语句用逗号隔开，要写成`i,j = 0,n-1`形式
