package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	myViper := getConfigViper()
	//数据获取  "行内容\行内容...\n"
	result := getData(myViper)
	//数据处理
	result = processData(result, myViper)
	//processData(result,myViper)
	//写入txt
	write2txt(result, myViper)
}

func write2txt(result string, myViper *viper.Viper) {
	var d1 = []byte(result)
	err2 := ioutil.WriteFile(myViper.GetString("fileName"), d1, 0666) //写入文件(字节数组)
	if err2 != nil {
		panic(err2)
	}
}

//获取配置文件viper
func getConfigViper() *viper.Viper {
	myViper := viper.New()
	myViper.SetConfigName("config")
	myViper.AddConfigPath(".")
	err := myViper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error %s\n", err)
		os.Exit(1)
	}
	myViper.WatchConfig()
	myViper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("the config had changed %s", e.Name)
	})

	return myViper
}

//处理类型为 " 字段1\字段2\字段3\n字段4\字段5\字段6\n..." 的字段为配置文件中格式
func processData(result string, myViper *viper.Viper) string {
	results := strings.Split(result, "\n")
	text := ""
	for _, value := range results {
		values := strings.Split(value, "\\")
		if len(values) < 2 {
			fmt.Printf("获取的数组为 1 返回")
			break
		}
		formatText := myViper.GetString("textFormat")
		fmt.Print(formatText)
		s := fmt.Sprintf(formatText, values[0], values[1], values[2], values[3])
		text += s
	}
	fmt.Printf(text)
	return text
}

//获取数据
func getData(myViper *viper.Viper) string {
	excelFileName := "data.xlsx"
	if len(os.Args) > 1 {
		excelFileName = os.Args[1]
	}
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	result := ""
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows[1:] {
			tpText := ""
			for _, cell := range row.Cells[myViper.GetInt("excelCellStart"):myViper.GetInt("excelCellEnd")] {
				text := cell.String()
				tpText += text + "\\"
			}
			result += strings.TrimRight(tpText, "\\") + "\n"
		}
	}
	return result
}
