package main

import (
	"fmt"
)

func main() {
	path := "\"d:\\go\\ser\\dagy\""
	fmt.Printf("%s \n", path)
	s := "I'm ok"
	fmt.Printf("%v \n", s)
	s1 := `
	Ctrl+4  四阶标题    Ctrl+Home   返回Typora顶部
	Ctrl+5  五阶标题    Ctrl+End    返回Typora底部
	Ctrl+6  六阶标题    Ctrl+T  创建表格
	Ctrl+L  选中某句话   Ctrl+K  创建超链接
	Ctrl+D  选中某个单词  Ctrl+F  搜索
	Ctrl+E  选中相同格式的文字   Ctrl+H  搜索并替换
	Alt+Shift+5 删除线 Ctrl+Shift+I    插入图片
	Ctrl+Shift+M    公式块 Ctrl+Shift+Q    引用
	`
	fmt.Printf(s1)
	
	
}
