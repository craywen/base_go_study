package pack

import "fmt"
import "strconv"


func main() {
	
	fmt.Print(" \n我的\tgo \n")
   	sum , ji :=	getVal(2,3);
	fmt.Println(+sum,ji)

	var num1 int  =99
	var num2 float64 =23.456
	var b  bool = true
	var str string 
	var mychar byte ='a'

	str = fmt.Sprintf("%d",num1)

	fmt.Printf("str type %T str=%q\n",str,str)

	str = fmt.Sprintf("%f",num2)
	fmt.Printf("str type %T str=%q\n",str,str)

	str = fmt.Sprintf("%t",b)
	fmt.Printf("str type %T str=%q\n",str,str)

	str = fmt.Sprintf("%c",mychar)
	fmt.Printf("str type %T str=%q\n",str,str)

	var num3 int =99;	
	var num4 float64 =23.5689
	var num5 int =999999;
	var bol2 bool =true

	str=strconv.FormatInt(int64(num3),10)
	fmt.Printf("str type %T str=%q\n",str,str)

	str=strconv.FormatFloat(num4,'f',10,64)
	fmt.Printf("str type %T str=%q\n",str,str)

	str=strconv.FormatBool(bol2)
	fmt.Printf("str type %T str=%q\n",str,str)	

	str=strconv.Itoa(num5)
	fmt.Printf("str type %T str=%q\n",str,str)


	var str2 string ="true"
	var bol3 bool 
	bol3, _ =strconv.ParseBool(str2)
	fmt.Printf("bol3 type %T str=%q\n",bol3,bol3)

	var int6 int =10
	fmt.Println("i的地址=",&int6)

	var ptr *int
	ptr= &int6
	*ptr=9
	fmt.Printf("ptr=%v\n",ptr)

	
	fmt.Println("i",int6)

	var int7 int =9;

	fmt.Printf("int7 address %v ",&int7)

	var ptr1 *int =&int7

	fmt.Printf("ptr 指向的值是=%v \n",*ptr1)

	if true {
		fmt.Println("你的年龄大于118,要对自己负责")
	}

	for i := 0; i < 10; i++ {
		
	}

	var strR string ="hello .world 深圳"
	strR2 :=[]rune (strR);
	for i := 0; i < len(strR); i++ {
		fmt.Printf("%c \n",strR[i])
	}
		fmt.Printf("%c \n",strR2)

	for index, v := range strR {
	  fmt.Printf("index=%d ,val=%c \n ",index ,v)
	}

	n1 :=4
	test1(n1,0)

	
	res2 := nimi( 10 ,10)
	fmt.Println("res2 :",res2)
	res3 := nimi( 10 ,10)
	fmt.Println("res3 :",res3)
//fmt.Println("mian() n1=",n1)
//defer test(1)

//defer fmt.Println("延时执行1")

//defer fmt.Println("延时执行2")
 addres1 :=addUper()
 fmt.Println("addres1=%d :",addres1(1))

 fmt.Println("addres2=%d :",addres1(2))

 fmt.Println("addres3=%d :",addres1(3))

 fmt.Println("addres3=%d :",addres1(5))


 numZ := 20
 test03(&numZ)
 fmt.Println("main () numZ ",numZ)


 fmt.Println(len(strR))
}

func  test03(a *int)  {
	*a=*a+10
}

// 累加器
func addUper () func (int)int {
	var b int =10
	return  func ( c int  ) int {
		c =c+b
		return c
	}
}

var (
	nimi = func ( a int ,b int) int  {
		return  a+b ;	
	}
)


func getVal (a int ,b int ) (c int ,d int)  {
	c=a+b
	d=a*b
	return c ,d
}


func test(n1 int)  {
	n1=n1+1
	fmt.Println("test() n1=",n1)
}

func test1(n int,  b int )  {
	if n>2{
		n--
		test1(n,0)
	}else{
		fmt.Println("n=",n)
	}
	
}

func init()  {	
	fmt.Println("初始化参数:")
}

