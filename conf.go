package conf

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Config struct {
	lock sync.RWMutex
	fn   string
	data map[string]string
}

// 解析指定ini文件并创建Config对象
func NewConfig(fileName string) (*Config, error) {
	c := &Config{
		fn:   fileName,
		data: make(map[string]string),
	}

	file, err := os.Open(fileName)
	if err != nil {
		return c, err
	}
	defer file.Close()

	bs := bufio.NewScanner(file)
	for bs.Scan() {
		line := strings.TrimSpace(bs.Text())
		if len(line) > 1 && line[0:1] != "#" && line[0:1] != ";" {
			ss := strings.SplitN(line, "=", 2)
			if len(ss) != 2 {
				continue
			}
			k := strings.TrimSpace(ss[0])
			if k != "" {
				c.data[k] = strings.TrimSpace(ss[1])
			}
		}
	}

	return c, nil
}

// 获取一个字符串值的配置项，如果值不存在返回空字符串
func (c *Config) GetString(key string) string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if v, ok := c.data[key]; ok {
		return v
	}
	return ""
}

// 获取一个整型配置项，如果出错或值不存在返回0
func (c *Config) GetInt(key string) int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if v, ok := c.data[key]; ok {
		i, _ := strconv.Atoi(v)
		return i
	}
	return 0
}

// 设置一个配置项，值参数是字符串
func (c *Config) SetString(key, val string) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data[key] = val
	return true
}

// 设置一个配置项，值参数是整型
func (c *Config) SetInt(key string, val int) bool {
	c.lock.Lock()
	c.lock.Unlock()

	c.data[key] = strconv.Itoa(val)
	return true
}

// 将修改写入到文件，覆盖原始文件，如果文件不存在则创建之
func (c *Config) Save() error {
	c.lock.Lock()
	c.lock.Unlock()

	file, err := os.Create(c.fn)
	if err != nil {
		return err
	}
	defer file.Close()

	for k, v := range c.data {
		file.WriteString(k + " = " + v + "\r\n")
	}

	return nil
}
