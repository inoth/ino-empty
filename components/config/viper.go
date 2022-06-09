package config

import (
	"defaultProject/components/cache"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// confKeyPrefix cache中的key的前缀
// cache可能被其它地方引用，为了防止key重复，每个引用的模块默认加一个key的前缀
const confKeyPrefix = "defaultProject_"

// 由于viper包本身对于文件的变化事件有一个bug，相关事件会被回调两次，常年未彻底解决
// 相关的issue清单：https://github.com/spf13/viper/issues?q=OnConfigChange
// 设置一个内部全局变量，引入一个cache，记录配置文件变化时的时间点，如果两次回调事件事件差小于1秒，我们认为是第二次回调事件，而不是人工修改配置文件，这样就避免了viper包的这个bug
var lastChangeTime time.Time

func init() {
	lastChangeTime = time.Now()
}

// var Cfg *viper.Viper
var Cfg *ViperConfig
var once sync.Once

// ViperConfig 定义一个结构体，用于包装viper，扩展viper的相关方法，使其优先从cache读取数据
// 注意实现ViperConfigInterface的所有方法
type ViperConfig struct {
	defaultValue map[string]interface{}
	viper        *viper.Viper
}

func Instance() *ViperConfig {
	once.Do(func() {
		Cfg = &ViperConfig{
			defaultValue: make(map[string]interface{}),
			viper:        viper.New(),
		}
	})
	return Cfg
}

func (m *ViperConfig) SetDefaultValue(defaultValue map[string]interface{}) *ViperConfig {
	// m.defaultValue = make(map[string]interface{})
	for k, v := range defaultValue {
		m.defaultValue[k] = v
	}
	return m
}

func (m *ViperConfig) loadDefaultValue() {
	for k, v := range m.defaultValue {
		m.viper.SetDefault(k, v)
	}
}

func (m *ViperConfig) Init() error {
	m.viper.AddConfigPath("config")
	m.viper.SetConfigName(selectConfigName(nil))
	m.viper.SetConfigType("yaml")
	if err := m.viper.ReadInConfig(); err != nil {
		return err
	}
	m.loadDefaultValue()
	return nil
}

func selectConfigName(path []string) string {
	if len(path) > 0 {
		return path[0]
	} else {
		e := os.Getenv("GORUNEVN")
		if len(e) > 0 {
			return "config." + e
		}
		return "config"
	}
}

// isCached 判断相关键是否已经缓存
func (y *ViperConfig) isCached(key string) bool {
	if _, ok := cache.Cache.IsExist(confKeyPrefix + key); ok {
		return true
	}
	return false
}

// cache 对键值进行缓存
func (y *ViperConfig) cache(key string, value interface{}) bool {
	return cache.Cache.Set(confKeyPrefix+key, value)
}

// getFromCache 通过键获取缓存的值
func (y *ViperConfig) getFromCache(key string) interface{} {
	return cache.Cache.Get(confKeyPrefix + key)
}

// clearCache 清空配置项
func (y *ViperConfig) clearCache() {
	cache.Cache.FuzzyDelete(confKeyPrefix)
}

// ConfigFileChangeListen 监听文件变化
func (y *ViperConfig) ConfigFileChangeListen() {
	y.viper.OnConfigChange(func(changeEvent fsnotify.Event) {
		if time.Now().Sub(lastChangeTime).Seconds() >= 1 {
			if changeEvent.Op.String() == "WRITE" {
				y.clearCache()
				lastChangeTime = time.Now()
				fmt.Println("配置文件已更新")
			}
		}
	})
	y.viper.WatchConfig()
}

// Get 获取原始值。先尝试从cache读取，若读取不到，从配置文件读取
func (y *ViperConfig) Get(key string) interface{} {
	if y.isCached(key) {
		return y.getFromCache(key)
	} else {
		value := y.viper.Get(key)
		y.cache(key, value)
		return value
	}
}

// GetString 获取字符串类型的值
func (y *ViperConfig) GetString(key string) string {
	if y.isCached(key) {
		return y.getFromCache(key).(string)
	} else {
		value := y.viper.GetString(key)
		y.cache(key, value)
		return value
	}

}

// GetBool 获取布尔类型的值
func (y *ViperConfig) GetBool(key string) bool {
	if y.isCached(key) {
		return y.getFromCache(key).(bool)
	} else {
		value := y.viper.GetBool(key)
		y.cache(key, value)
		return value
	}
}

// GetInt 获取int类型的值
func (y *ViperConfig) GetInt(key string) int {
	if y.isCached(key) {
		return y.getFromCache(key).(int)
	} else {
		value := y.viper.GetInt(key)
		y.cache(key, value)
		return value
	}
}

// GetInt32 获取int32类型的值
func (y *ViperConfig) GetInt32(key string) int32 {
	if y.isCached(key) {
		return y.getFromCache(key).(int32)
	} else {
		value := y.viper.GetInt32(key)
		y.cache(key, value)
		return value
	}
}

// GetInt64 获取int64类型的值
func (y *ViperConfig) GetInt64(key string) int64 {
	if y.isCached(key) {
		return y.getFromCache(key).(int64)
	} else {
		value := y.viper.GetInt64(key)
		y.cache(key, value)
		return value
	}
}

// GetFloat64 获取浮点数类型的值
func (y *ViperConfig) GetFloat64(key string) float64 {
	if y.isCached(key) {
		return y.getFromCache(key).(float64)
	} else {
		value := y.viper.GetFloat64(key)
		y.cache(key, value)
		return value
	}
}

// GetDuration 获取time.Duration类型的值
func (y *ViperConfig) GetDuration(key string) time.Duration {
	if y.isCached(key) {
		return y.getFromCache(key).(time.Duration)
	} else {
		value := y.viper.GetDuration(key)
		y.cache(key, value)
		return value
	}
}

// GetStringSlice 获取字符串切片的值
func (y *ViperConfig) GetStringSlice(key string) []string {
	if y.isCached(key) {
		return y.getFromCache(key).([]string)
	} else {
		value := y.viper.GetStringSlice(key)
		y.cache(key, value)
		return value
	}
}
