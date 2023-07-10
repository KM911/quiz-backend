package main

import (
	"fmt"
	"os"
	"quiz/model"
	"strings"
	"testing"
)

//以下不属于Python语言特点的是
//choice.A. 平台无关
//choice.B. 第三方库丰富
//choice.C. 适合编写系统软件
//choice.D. 语法简洁
//回答正确！+1分
//考生答案：C正确答案：choice.C

func Test_Main(t *testing.T) {
	// 主要是将我们的内容进行一个转义就是说
	data, _ := os.ReadFile("text.txt")
	str := string(data)
	//fmt.Println(str)

	// 将其解析
	dataLens := strings.Split(str, "\n")
	i := 0
	choices := make([]model.Choice, 0)
	for {
		choice := model.Choice{}
		choice.Title += dataLens[i]
		choice.Answer = ""
		// if start with choice.A.
		for {
			if strings.HasPrefix(dataLens[i+1], "A.") {
				// 说明这是一个选�
				break
			}
			i++
			choice.Title += dataLens[i]
		}
		// 获取选项
		for {
			if strings.HasPrefix(dataLens[i+1], "B.") {
				// 说明这是一个选�
				break
			}
			i++
			choice.A += dataLens[i]
		}
		for {
			if strings.HasPrefix(dataLens[i+1], "C.") {
				// 说明这是一个选�
				break
			}
			i++
			choice.B += dataLens[i]
		}

		for {
			if strings.HasPrefix(dataLens[i+1], "D.") {
				// 说明这是一个选�
				break
			}
			i++
			choice.C += dataLens[i]
		}
		for {
			if strings.HasPrefix(dataLens[i+1], "回") {
				// 说明这是一个选�
				break
			}
			i++
			choice.D += dataLens[i]
		}
		// 现在的i指向的是 choice.D. 选项
		i++
		i++
		// 取最后一个字符
		choice.Answer = dataLens[i][len(dataLens[i])-2 : len(dataLens[i])-1]
		if i >= len(dataLens)-6 {
			break
		}
		i++
		//fmt.Println(choice.Title)
		////fmt.Println(choice.choices)
		//for k, v := range choice.choices {
		//	fmt.Println(k, v)
		//}
		//fmt.Println(choice.Answer)
		//fmt.Println(len(choice.Answer))
		////break
		choices = append(choices, choice)
	}
	model.InitDao()
	fmt.Println(len(choices))
	for _, v := range choices {
		fmt.Println(v.Title)
		fmt.Println(v.Answer)
		fmt.Println(len(v.Answer))
		model.InsertChoice(v)
	}

}
