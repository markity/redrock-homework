package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type BirthDate struct {
	Year  int
	Month int
	Day   int
}

func (bd *BirthDate) String() string {
	return fmt.Sprintf("%v年%v月%v日", bd.Year, bd.Month, bd.Day)
}

type Coder struct {
	Name         string
	StudentID    string
	PhoneNumber  string
	NumCodeLines int
	BirthDate    BirthDate
}

func (c *Coder) String() string {
	return fmt.Sprintf("姓名: %v, 学号: %v, 电话号码: %v, 代码数目: %v, 出生日期: %v", c.Name, c.StudentID, c.PhoneNumber, c.NumCodeLines, &c.BirthDate)
}

type CoderSortByCodeCount [](*Coder)

func (lis CoderSortByCodeCount) Len() int {
	return len(lis)
}

func (lis CoderSortByCodeCount) Swap(i, j int) {
	temp := (lis)[i]
	(lis)[i] = (lis)[j]
	(lis)[j] = temp
}

func (lis CoderSortByCodeCount) Less(i, j int) bool {
	return (lis)[i].NumCodeLines < (lis)[j].NumCodeLines
}

func main() {
	data := make([](*Coder), 0)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		println("输入你需要的操作, 1: 插入, 2: 查询, 3: 展示所有(按照代码行数逆序排序) 4: 退出程序")
		scanner.Scan()
		nextLine := scanner.Text()
		switch nextLine {
		case "1":
			println("按以下格式输入: 姓名 学生ID 电话号码 学生打的代码行数 出生年 月 日")
			scanner.Scan()
			txt := scanner.Text()
			txt = strings.ReplaceAll(txt, "\t", " ")
			r := regexp.MustCompile("[^\\s]+")
			x := r.FindAllString(txt, -1)
			if len(x) != 7 {
				println("输入的数据有误")
				continue
			}
			var coder Coder
			coder.Name = x[0]
			coder.StudentID = x[1]
			coder.PhoneNumber = x[2]
			codeCounts, err := strconv.Atoi(x[3])
			if err != nil {
				println("输入的数据有误")
				continue
			}
			coder.NumCodeLines = codeCounts
			year, err := strconv.Atoi(x[4])
			if err != nil {
				println("输入的数据有误")
				continue
			}
			coder.BirthDate.Year = year
			month, err := strconv.Atoi(x[5])
			if err != nil {
				println("输入的数据有误")
				continue
			}
			coder.BirthDate.Month = month
			day, err := strconv.Atoi(x[6])
			if err != nil {
				println("输入的数据有误")
				continue
			}
			coder.BirthDate.Day = day
			data = append(data, &coder)
			fmt.Printf("插入成功: %s\n", &coder)
			sort.Sort(CoderSortByCodeCount(data))
		case "2":
			println("输入姓名或者学号: ")
			scanner.Scan()
			txt := scanner.Text()
			n := 0
			for _, v := range data {
				if v.Name == txt || v.StudentID == txt {
					fmt.Printf("%s\n", v)
				}
				n++
			}
			if n == 0 {
				println("没有找到任何合适的数据")
			}
		case "3":
			for _, v := range data {
				fmt.Printf("%s\n", v)
			}
		case "4":
			os.Exit(0)
		default:
			println("非法操作, 重新输入")
		}
	}
}
