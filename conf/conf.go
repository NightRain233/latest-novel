package conf

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	defaultGlobalConf globalConf
)

var (
	GlobalConfPath = "conf.yml"
)

func init() {
	setConfFullPath(GlobalConfPath)
	parseConfig()
}
func setConfFullPath(name string) {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	//go run 会将源代码编译到系统TEMP或TMP环境变量目录中并启动执行
	//所以可通过对比os.Executable()获取到的路径是否与环境变量TEMP设置的路径相同来判断是否是go run启动
	if strings.Contains(dir, tmpDir) {
		dir = getCurrentAbPathByCaller()
		GlobalConfPath = path.Join(dir, name)
		return
	}
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable() //获取执行文件的绝对路径
	if err != nil {
		panic(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath)) // 获取真实路径,避免是symlink指向的路径
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

type globalConf struct {
	URLs []string `yaml:"urls"`
}

func parseConfig() {
	f, err := os.Open(GlobalConfPath)
	if err != nil {
		log.Fatalf("open %v: %v", GlobalConfPath, err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("read file %v: %v", GlobalConfPath, err)
	}

	if err := yaml.Unmarshal(b, &defaultGlobalConf); err != nil {
		log.Fatalf("unmarshal file %v: %v", GlobalConfPath, err)
	}
}

func GetNovelURLs() []string {
	return defaultGlobalConf.URLs
}
