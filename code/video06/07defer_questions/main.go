package main

import (
	"fmt"
)

func f1() int {
	i := 5
	defer func() {
		i++
	}()
	return i // rval = 5; i++后是6  ret rval 所以最后返回5
}
func f2() *int {
	i := 5
	defer func() {
		i++
		fmt.Printf(":::%p\n", &i)
	}()
	fmt.Printf(":::%p\n", &i)
	return &i // rval = i的地址;  i++后是6, 但是地址不变  ret rval(地址不变，注意虽然这里地址不变，但是地址对应的值已经变为了6)
}

func f22() (result int) {
	//return  //只有return的话，默认返回result的默认值0
	// return result   // return result的话，返回的是result的默认值0
	return 123 // return 123的话，将123赋值给result，返回result
}

func f3() (result int) {
	defer func() {
		result++
	}()
	return 5 // 如果函数有默认返回值result,在这里result的角色替代了rval
	// result(rval) = 5; result ++后是6; ret result(result替换了rval) 6
}

func f4() (result int) {
	defer func() {
		result++
	}()
	return result // 如果函数有默认返回值result,在这里result的角色替代了rval  result(rval) = 0(默认); result ++后是1; ret result(result替换了rval) 1
}

func f5() (r int) {
	t := 5
	defer func() {
		t = t + 1
	}()
	return t //r = t(把t的值赋给r); t = t+1=6; ret r 5
}

func f6() (r int) {
	defer func(r int) {
		r = r + 1
	}(r)
	return r //f6.r = 0(默认值0); defer_func.r = f6.r+1 = 0+1=1; ret f6.r 0
}
func f6_() (r int) {
	fmt.Println(r, &r) // 0 0x14000122030
	defer func(r int) {
		r = r + 1
		fmt.Println(r, &r) // 1 0x14000122038
	}(r)
	return 5 //f6.r = 5; defer_func.r = defer_func.r（这个值是deferfunc在之前延迟注册环节保留函数体时候，把还为0的f6.r就拷贝过来了）+1 = 0+1=1; ret f6.r 5 (这里注意defer_func的形式参数r，开辟了空间，把f6.r的默认值0给了defer_func.r)
}

func f7() (r int) {
	defer func(x int) {
		r = x + 1
	}(r)
	return 5 //r(rval) = 5; 这里的deferfunc的形参x，传入实参r，实参r的值是0，所以x=0，r=x+1=0+1=1(注意这个延迟函数是闭包函数，r:= x+1和r=x+1不一样，r=x+1的r没用冒号说明是已经定义了，使用的是外部自由变量r也就是f7.r，把f7.r修改了) ；ret r 1
}

func main() {
	println(f1())  // 5
	println(f2())  // 0x14000122008 print的2个地址相同，并把i的地址返回
	println(*f2()) // 6 print的2个地址相同，并把i的地址对应的值返回,因为这个地址对应的值已经被改为了6，所以通过*取值就是6
	println(f3())  // 6 虽然语句中return 5，但是中间result(rval)被defer修改了，所以最后返回的是6
	println(f4())  // 1
	println(f5())  // 5
	println(f6())  // 0
	println(f6_()) // 5
	println(f7())  // 1

}
