package main

import (
	"bufio"
	"github.com/spf13/pflag"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
)

type selpg_args struct{
	start_page int
	end_page int
	in_filename string
	page_len int	/* default value, can be overriden by "-l number" on command line */
	page_type string	/* 'l' for lines-delimited, 'f' for form-feed-delimited */
					/* default is 'l' */
	print_dest string	
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, os.Args[0]+" -sstart_page_num -eend_page_num [ -f | -llines_per_page ] [ -ddest ] [ in_filename ]\n")
    pflag.PrintDefaults()
}

func Init(args *selpg_args){
	pflag.Usage = usage
    pflag.IntVarP(&(args.start_page), "start", "s", -1, "start page")
    pflag.IntVarP(&(args.end_page), "end", "e", -1, "end page")
    pflag.IntVarP(&(args.page_len), "line", "l", 72, "page len")
    pflag.StringVarP(&(args.print_dest), "destionation", "d", "", "print destionation")
	pflag.StringVarP(&(args.page_type), "type", "f", "l", "type of print")
}

//Check whether parameters are valid or not
func check(args *selpg_args){
	//Not enough args
	if len(os.Args) < 3 || args.start_page == -1 || args.end_page == -1{
		fmt.Fprintf(os.Stderr, "ERROR: %s not enough parameters!\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "You should input the parameters start page and end page!\n")
		usage()
		os.Exit(0)
	}

	//Page's index error
	if args.start_page <= 0{
		fmt.Fprintf(os.Stderr, "ERROR: Invalid start page number! Can not less than or equal to 0!\n")
		os.Exit(1)
	}
	if args.end_page <= 0{
		fmt.Fprintf(os.Stderr, "ERROR: Invalid end page number! Can not less than or equal to 0!\n")
		os.Exit(2)
	}
	if args.start_page > math.MaxInt32 - 1{
		fmt.Fprintf(os.Stderr, "ERROR: Invalid start page number! Can not grater than 2147483646!\n")
		os.Exit(3)
	}
	if args.end_page > math.MaxInt32 - 1{
		fmt.Fprintf(os.Stderr, "ERROR: Invalid end page number! Can not grater than 2147483646!\n")
		os.Exit(4)
	}
	if args.end_page < args.start_page{
		fmt.Fprintf(os.Stderr, "ERROR: Invalid end page number! Can not less than start page number!\n")
		os.Exit(5)
	}

	//page_len error
	if args.page_len <= 0{
		fmt.Fprintf(os.Stderr, "ERROR: The pages' length can not be less than or equal to 0!\n")
		os.Exit(6)
	}
	if args.page_len > math.MaxInt32 - 1{
		fmt.Fprintf(os.Stderr, "ERROR: The pages' length can not be grater than 2147483646!\n")
		os.Exit(7)
	}

	//page type error
	if args.page_type != "f" && args.page_type != "l"{
		fmt.Fprintf(os.Stderr, "ERROR: The page type must be 'f' or 'l'!\n")
		os.Exit(8)
	}
}

func process_input(args *selpg_args){
	var fin *os.File
	if args.in_filename != ""{
		var opErr error
		fin, opErr = os.Open(args.in_filename)
		if opErr != nil{
			fmt.Fprintf(os.Stderr, "\nERROR! Can not open the input file: %s\n", args.in_filename)
			os.Exit(9)
		}
	}else{
		fin = os.Stdin
	}
	finBuffer := bufio.NewReader(fin)


	var fout io.WriteCloser
	var cmd *exec.Cmd
	if args.print_dest != ""{
		cmd = exec.Command("lp", "-d", args.print_dest)
		var desErr error
		cmd.Stdout, desErr = os.OpenFile(args.print_dest, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		fout, desErr = cmd.StdinPipe()
		if desErr != nil{
			fmt.Fprintf(os.Stderr, "\nERROR! Can not open pipe to \"lp -d%s\"\n",  args.print_dest)
			os.Exit(10)
		}
		cmd.Start()
		cmd.Wait()
	}else{
		fout = os.Stdout
	}
	
	/* begin one of two main loops based on page type */
	if args.page_type == "l" {
		line_ctr := 0
		page_ctr := 1
		for {
			line,  crc := finBuffer.ReadString('\n')
			if crc != nil {
				break
			}
			line_ctr ++
			if line_ctr > args.page_len {
				page_ctr ++
				line_ctr = 1
			}
	
			if (page_ctr >= args.start_page) && (page_ctr <= args.end_page) {
				_, err := fout.Write([]byte(line))
				if err != nil {
					fmt.Fprintf(os.Stderr, "ERROR!", err)
					os.Exit(11)
				}
			 }
		}  
	}else{
		page_ctr := 1
		for{
			
			page,  crc := finBuffer.ReadString('\f')
			if crc != nil {
				break
			}
			page_ctr ++
			if ( (page_ctr >= args.start_page) && (page_ctr <= args.end_page) ){
				_, err := fout.Write([]byte(page))
				if err != nil {
					fmt.Fprintf(os.Stderr, "ERROR!", err)
					os.Exit(12)
				}
			}
		}
	}

	fin.Close()
	fout.Close()
}

func main() {
	args := new(selpg_args)
	//Get args
	Init(args)
    pflag.Parse()
    
    //Get other args
    othersArg := pflag.Args()
    if pflag.NArg() > 0 {
        args.in_filename = othersArg[0]
    } else {
        args.in_filename = ""
	}

	//Check the grammar
	check(args)
	process_input(args)
}
