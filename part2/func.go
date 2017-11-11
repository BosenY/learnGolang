package main

import (
	"fmt"
	"math"
)

/*
Go中的函数有两种形式：
一种是和普通函数一样的，接收参数并返回参数的
另外一种是面向对象的思路，带有接收者的函数，我们称之为method
*/

/*
现在是第一种场景，我需要去定义一个函数来计算长方形的面积，一般的情况是创建一个函数来实现
*/

type Rectangle struct {
	width, height int
}

func area(r Rectangle) int {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (r Rectangle) area_() int {
	return r.width * r.height
}

func (c Circle) area_() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{4, 2}
	fmt.Println(area(r1))
	/*
		这个是一种面向过程的思路，就和js一样，
		这样的实现当然没有任何的问题，但是当你要增加圆形、正方形、五边形甚至其他的多边形的时候，你想计算他们的面积的时候怎么办，如果是继续写新的函数来适应新的类型，那你的函数名称也必须跟着变换
		这样的实现显而易见是很蠢的
		从概念上来说： 面积是形状的一个属性，他是属于这个特定的形状的。
		基于上面的原因，就有了方法(method)的概念，method是附属给一个给定的类型上的，他的语法和函数的声明语法几乎一样，只是在func后面首先写一个receiver(接收者，也就是method依从的主体)
		所以用上面的例子来说，area()是依赖于某个形状比如是rectangle来发生作用的。Rectangle.area()的发出者是rectangle这个struct 而area()是它的一个方法，而非一个普通的外围函数这么简单，
		更具体地说，rectangle存在字段length和width，同时也存在了一个方法area(),这些字段和方法都属于rectangle

		"A method is a function with an implicit first argument, called a receiver"
		func (r ReceiverType)  funcName(parameters) (results)
		       接收者的一个实例       方法名  参数         返回值
	*/
	// r2 := Rectangle{9, 6}
	c1 := Circle{4}
	fmt.Println(Rectangle{9, 6}.area_())
	fmt.Println(c1.area_())
	/*
		当然method并不是只能作用于struct上，他还可以定义在任何你自定义的类型、内置类型。struct等各种类型上面，
		这里要简单记录一个自定义类型的概念，struct只是自定义类型上面一种比较特殊的类型，还有其他自定义类型的声明，可以通过如下的声明来实现：
		type typeName typeLiteral
		  		自定义类型名称   类型
	*/

}
