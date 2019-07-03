# excel2text
use golang to write a script to read excel content and format print to text

# Usage

1. windows open the file path use short cut `Alt+D` print `cmd` then `enter` enter CLI(command line interface) and print
```$xslt
go build excel2text.go
```

2. the command will created a executable file in current path `excel2text.exe`

3. open config file `config.yaml` custom you want to format text then save config file

4. drag your excel file what you want to be format to `excel2text.exe` or double click `excel2text.exe` it will find file according to config file

5. then the script will auto process and create a text file has same name 