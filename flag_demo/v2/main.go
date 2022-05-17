package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.Parse() //将命令行解析为定义的标志，便于后续参数的使用
	// 内部实现

	//var CommandLine = NewFlagSet(os.Args[0], ExitOnError) 调用NewFlagSet实现一个新的空命令集,然后将os.Args[0]作为额外的参数进入
	//
	//func Parse() {
	//	CommandLine.Parse(os.Args[1:])
	//}
	// 其中需要注意的是,这里第二个错误返回参数是ExitOnError 所以会直接退出程序,如果需要不退出程序那么需要额外的处理

	//                  **FlagSet.Parse**
	//func (f *FlagSet) Parse(arguments []string) error {
	//	f.parsed = true
	//	f.args = arguments
	//	for {
	//	seen, err := f.parseOne()
	//	if seen {
	//	continue
	//}
	//	if err == nil {
	//	break
	//}
	//	switch f.errorHandling {
	//case ContinueOnError:
	//	return err
	//case ExitOnError:
	//	os.Exit(2)
	//case PanicOnError:
	//	panic(err)
	//}
	//}
	//	return nil
	//}

	//p.s. 其中不同的就是这里进行了错误的分别处理

	//     **FlagSet.parseOne**
	
	args := flag.Args() //取参数
	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		//该方法会返回带有指定名称和错误处理属性的空命令集给我们去使用，相当于就是创建一个新的命令集了去支持子命令了
		//flag 内置三种错误返回的方式
		//const (
		//	// 返回错误描述
		//	ContinueOnError ErrorHandling = iota
		//	// 调用 os.Exit(2) 退出程序
		//	ExitOnError
		//	// 调用 panic 语句抛出错误异常
		//	PanicOnError
		//)
		goCmd.StringVar(&name, "name", "go", "help")
		_ = goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "n", "php", "help")
		_ = phpCmd.Parse(args[1:])
	}
	log.Printf("name = %s", name)
}
