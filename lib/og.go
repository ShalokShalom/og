package og

import (
	"bufio"
	"fmt"
	tm "github.com/buger/goterm"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type OgConfig struct {
	Blocks      bool
	Dirty       bool
	Print       bool
	Compile     bool
	Ast         bool
	Quiet       bool
	Interpreter bool
	Paths       []string
	Files       []string
	Workers     int
	OutPath     string
	NoBuild     bool
	Run         bool
}

var (
	config OgConfig
)

func GetNewPath(filePath string) string {
	if config.OutPath != "./" {
		splited := strings.SplitN(filePath, "/", 2)
		filePath = splited[1]
	}
	return strings.Replace(path.Join(config.OutPath, filePath), ".og", ".go", 1)
}
func Compile(config_ OgConfig) error {
	config = config_
	if len(config.Paths) == 0 {
		config.Paths = []string{"."}
	}
	for _, p := range config.Paths {
		if err := filepath.Walk(p, walker); err != nil {
			fmt.Println("Error", err)
			return err
		}
	}
	if len(config.Files) == 0 {
		if config.Run {
			return Run()
		}
		if !config.Quiet {
			tm.Print(tm.Color("~> ", tm.YELLOW), tm.Color("Oglang: ", tm.RED), tm.Color("Nothing to do.", tm.GREEN))
			tm.Flush()
		}
		return nil
	}
	poolSize := config.Workers
	if len(config.Files) < poolSize {
		poolSize = len(config.Files)
	}
	pool := NewPool(poolSize, len(config.Files), !config.Quiet, ReadAndProceed)
	for _, file := range config.Files {
		pool.Queue(file)
	}
	pool.Run()
	if !config.NoBuild {
		if err := goBuild(); err != nil {
			return err
		}
	}
	if config.Run {
		if err := Run(); err != nil {
			return err
		}
	}
	return nil
}
func Run() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	current := path.Base(dir)
	tm.Print(tm.Color("~> ", tm.YELLOW), tm.Color("Oglang: ", tm.RED), tm.Color("Running... \n", tm.GREEN))
	tm.Flush()
	cmd := exec.Command("./" + current)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Start(); err != nil {
		return err
	}
	cmd.Wait()
	return nil
}
func goBuild() error {
	tm.Print("                                          \r")
	tm.Println(" ", tm.Color("[", tm.RED), tm.Color(strconv.Itoa(len(config.Files)), tm.YELLOW), tm.Color("/", tm.RED), tm.Color(strconv.Itoa(len(config.Files)), tm.GREEN), tm.Color("]", tm.RED), "Compiling go")
	tm.MoveCursorUp(1)
	tm.Flush()
	cmd := exec.Command("go", "build")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
	} else if !config.Quiet {
		tm.MoveCursorUp(1)
		tm.Print("                                          \r")
		tm.Print(tm.Color("~> ", tm.YELLOW), tm.Color("Oglang: ", tm.RED), tm.Color("Compiled ", tm.GREEN), tm.Color(strconv.Itoa(len(config.Files)), tm.MAGENTA), tm.Color(" files.", tm.GREEN))
		tm.Flush()
	}
	return err
}
func MustCompile(filePath string, info os.FileInfo) bool {
	newPath := GetNewPath(filePath)
	stat, err := os.Stat(newPath)
	if err != nil {
		return true
	}
	return info.ModTime().After(stat.ModTime())
}
func walker(filePath string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	if path.Ext(filePath) != ".og" {
		return nil
	}
	if !MustCompile(filePath, info) {
		return nil
	}
	config.Files = append(config.Files, filePath)
	return nil
}

var (
	lastLength = 0
)

func ReadAndProceed(filePath string) error {
	source, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	res, err := ProcessFile(filePath, string(source), false)
	if err != nil {
		return err
	}
	finalizeFile(filePath, res)
	return nil
}
func ProcessFile(filePath string, data string, isInterpret bool) (string, error) {
	preprocessed := Preproc(string(data))
	if config.Blocks {
		return preprocessed, nil
	}
	res := ""
	if !isInterpret {
		res = Parse(filePath, string(preprocessed))
	} else {
		res = ParseInterpret(filePath, string(preprocessed))
	}
	if config.Dirty {
		return res, nil
	}
	return format(res)
}
func finalizeFile(filePath string, data string) {
	if config.Ast || config.Print || config.Dirty || config.Blocks {
		fmt.Println(data)
	} else {
		writeFile(filePath, data)
	}
}
func writeFile(filePath string, data string) {
	newPath := GetNewPath(filePath)
	os.MkdirAll(filepath.Dir(newPath), os.ModePerm)
	ioutil.WriteFile(newPath, []byte(data), os.ModePerm)
}
func format(str string) (string, error) {
	cmd := exec.Command("gofmt")
	stdin, _ := cmd.StdinPipe()
	stdin.Write([]byte(str))
	stdin.Close()
	final, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(final), nil
}
func RunInterpreter() {
	running := true
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for running {
		fmt.Print("> ")
		if !scanner.Scan() {
			return
		}
		ln := scanner.Text()
		if len(ln) == 0 {
			continue
		}
		str, err := ProcessFile("STDIN", ln, true)
		if err != nil {
			fmt.Println(err)
		} else {
			execCode(str)
		}
	}
}
func execCode(str string) {
	skelton := `package main
  import "fmt"
  func main() {
  fmt.Print(` + str[:len(str)-1] + `)
  }`
	ioutil.WriteFile("/tmp/main.go", []byte(skelton), os.ModePerm)
	cmd := exec.Command("go", "run", "/tmp/main.go")
	stdin, _ := cmd.StdinPipe()
	stdin.Write([]byte(str))
	stdin.Close()
	final, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(final))
}
