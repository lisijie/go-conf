package goconf

import (
    "bufio"
    "os"
    "strconv"
    "strings"
    "sync"
    "errors"
)

const (
    DefaultSection = "DEFAULT"
)

type Config struct {
    lock     sync.RWMutex
    fn       string
    sections map[string]*Section
}

type Section struct {
    lock sync.RWMutex
    data map[string]string
}

// 获取一个字符串值的配置项，如果值不存在返回空字符串
func (s *Section) GetString(key string, def ...string) string {
    s.lock.RLock()
    defer s.lock.RUnlock()
    if v, ok := s.data[key]; ok {
        return v
    }
    if len(def) > 0 {
        return def[0]
    }
    return ""
}

// 获取一个整型配置项，如果出错或值不存在返回0
func (s *Section) GetInt(key string, def ...int) int {
    s.lock.RLock()
    defer s.lock.RUnlock()
    if v, ok := s.data[key]; ok {
        i, _ := strconv.Atoi(v)
        return i
    }
    if len(def) > 0 {
        return def[0]
    }
    return 0
}

// 获取一个布尔型配置项
func (s *Section) GetBool(key string) bool {
    s.lock.RLock()
    defer s.lock.RUnlock()
    v, ok := s.data[key]
    if !ok {
        return false
    }
    strings.ToLower(v)
    if v == "yes" || v == "true" || v == "1" {
        return true
    }
    return false
}

// 设置一个配置项
func (s *Section) Set(key string, val interface{}) error {
    s.lock.Lock()
    defer s.lock.Unlock()
    switch v := val.(type) {
    case string:
        s.data[key] = v
    case int:
        s.data[key] = strconv.Itoa(v)
    default:
        return errors.New("invalid type")
    }
    return nil
}

func newSection() *Section {
    return &Section{data: make(map[string]string)}
}

// 解析指定ini文件并创建Config对象
func NewConfig(fileName string) (*Config, error) {
    c := &Config{
        fn:       fileName,
        sections: make(map[string]*Section),
    }

    c.sections[DefaultSection] = newSection()

    file, err := os.Open(fileName)
    if err != nil {
        return c, err
    }
    defer file.Close()

    bs := bufio.NewScanner(file)
    section := DefaultSection
    for bs.Scan() {
        line := strings.TrimSpace(bs.Text())
        if len(line) > 1 && line[0:1] != "#" && line[0:1] != ";" {
            if line[0:1] == "[" && line[len(line)-1:] == "]" {
                section = line[1:len(line)-1]
                if _, ok := c.sections[section]; !ok {
                    c.sections[section] = newSection()
                }
                continue
            }

            ss := strings.SplitN(line, "=", 2)
            if len(ss) != 2 {
                continue
            }
            k := strings.TrimSpace(ss[0])
            if k != "" {
                c.sections[section].data[k] = strings.TrimSpace(ss[1])
            }
        }
    }

    return c, nil
}

func (c *Config) GetSection(name string) *Section {
    if v, ok := c.sections[name]; ok {
        return v
    }
    return nil
}

// 获取一个字符串值的配置项，如果值不存在返回空字符串
func (c *Config) GetString(key string, def ...string) string {
    return c.GetSection(DefaultSection).GetString(key, def...)
}

// 获取一个整型配置项，如果出错或值不存在返回0
func (c *Config) GetInt(key string, def ...int) int {
    return c.GetSection(DefaultSection).GetInt(key, def...)
}
