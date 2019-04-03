package config

import (
	"bytes"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/spf13/viper"
)

var (
	_ Configuration = (*ViperConfiguration)(nil)
)

type ViperConfiguration struct {
	*viper.Viper
}

func NewViperConfiguration(conf *viper.Viper) Configuration {
	return &ViperConfiguration{
		conf,
	}
}

func (p *ViperConfiguration) GetConfig(path string) Configuration {
	if p == nil || p.Viper == nil {
		return (*ViperConfiguration)(nil)
	}

	//fmt.Fprintf(os.Stderr, "GetConfig p.Viper.AllKeys: %v\n", p.Viper.AllKeys())

	conf := p.Viper.Sub(path)
	if conf == nil {
		return (*ViperConfiguration)(nil)
	}

	return &ViperConfiguration{conf}
}

func (p *ViperConfiguration) WithFallback(fallback Configuration) Configuration {
	if fallback == nil {
		return p
	}

	switch v := fallback.(type) {
	case *ViperConfiguration:
		{
			p.Viper = v.Viper
		}
	case *Config:
		{
			p.Viper = v.Configuration.(*ViperConfiguration).Viper
		}
	}

	//fmt.Fprintf(os.Stderr, "WithFallback p.Viper.AllKeys: %v\n", p.Viper.AllKeys())

	return p
}

func (p *ViperConfiguration) GetBoolean(path string, defaultVal ...bool) bool {
	if p == nil || p.Viper == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return false
	}

	for _, v := range defaultVal {
		p.Viper.SetDefault(path, v)
		break
	}
	return p.Viper.GetBool(path)
}

func (p *ViperConfiguration) GetByteSize(path string) *big.Int {
	if p == nil || p.Viper == nil {
		return nil
	}
	//return p.Viper.GetByteSize(path)
	return nil
}

func (p *ViperConfiguration) GetInt32(path string, defaultVal ...int32) int32 {
	if p == nil || p.Viper == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}

	for _, v := range defaultVal {
		p.Viper.SetDefault(path, v)
		break
	}
	return int32(p.Viper.GetInt(path))
}

func (p *ViperConfiguration) GetInt64(path string, defaultVal ...int64) int64 {
	if p == nil || p.Viper == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}

	for _, v := range defaultVal {
		p.Viper.SetDefault(path, v)
		break
	}
	return p.Viper.GetInt64(path)
}

func (p *ViperConfiguration) GetString(path string, defaultVal ...string) string {
	if p == nil || p.Viper == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return ""
	}

	for _, v := range defaultVal {
		p.Viper.SetDefault(path, v)
		break
	}
	return p.Viper.GetString(path)
}

func (p *ViperConfiguration) GetFloat32(path string, defaultVal ...float32) float32 {
	if p == nil || p.Viper == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}

	for _, v := range defaultVal {
		p.Viper.SetDefault(path, v)
		break
	}
	return float32(p.Viper.GetFloat64(path))
}

func (p *ViperConfiguration) GetFloat64(path string, defaultVal ...float64) float64 {
	if p == nil || p.Viper == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}

	for _, v := range defaultVal {
		p.Viper.SetDefault(path, v)
		break
	}
	return p.Viper.GetFloat64(path)
}

func (p *ViperConfiguration) GetTimeDuration(path string, defaultVal ...time.Duration) time.Duration {
	if p == nil || p.Viper == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}

	for _, v := range defaultVal {
		p.Viper.SetDefault(path, v)
		break
	}
	return p.Viper.GetDuration(path)
}

func (p *ViperConfiguration) GetTimeDurationInfiniteNotAllowed(path string, defaultVal ...time.Duration) time.Duration {
	if p == nil || p.Viper == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}

	for _, v := range defaultVal {
		p.Viper.SetDefault(path, v)
		break
	}
	return p.Viper.GetDuration(path)
}

func (p *ViperConfiguration) GetBooleanList(path string) []bool {
	if p == nil || p.Viper == nil {
		return nil
	}

	//return p.Viper.GetBooleanList(path)
	return nil
}

func (p *ViperConfiguration) GetFloat32List(path string) []float32 {
	if p == nil || p.Viper == nil {
		return nil
	}
	//return p.Viper.GetFloat32List(path)
	return nil
}

func (p *ViperConfiguration) GetFloat64List(path string) []float64 {
	if p == nil || p.Viper == nil {
		return nil
	}
	//return p.Viper.GetFloat64List(path)
	return nil
}

func (p *ViperConfiguration) GetInt32List(path string) []int32 {
	if p == nil || p.Viper == nil {
		return nil
	}
	//return p.Viper.GetInt32List(path)
	return nil
}

func (p *ViperConfiguration) GetInt64List(path string) []int64 {
	if p == nil || p.Viper == nil {
		return nil
	}
	//return p.Viper.GetInt64List(path)
	return nil
}

func (p *ViperConfiguration) GetByteList(path string) []byte {
	if p == nil || p.Viper == nil {
		return nil
	}
	//return p.Viper.GetByteList(path)
	return nil
}

func (p *ViperConfiguration) GetStringList(path string) []string {
	if p == nil || p.Viper == nil {
		return nil
	}
	//return p.Viper.GetStringList(path)
	return nil
}

func (p *ViperConfiguration) HasPath(path string) bool {
	if p == nil || p.Viper == nil {
		return false
	}
	return p.Viper.IsSet(path)
}

func (p *ViperConfiguration) Keys() []string {
	if p == nil || p.Viper == nil {
		return nil
	}

	settings := p.Viper.AllSettings()
	keys := make([]string, 0, len(settings))
	for k := range settings {
		keys = append(keys, k)
	}

	//fmt.Fprintf(os.Stderr, "Keys: %v\n", keys)

	return keys
}

func (p *ViperConfiguration) IsEmpty() bool {
	return p == nil || p.Viper == nil
}

func (p *ViperConfiguration) IsObject(path string) bool {
	//return p != nil && p.Viper.IsObject(path)
	return false
}

func (p *ViperConfiguration) IsArray(path string) bool {
	//return p != nil && p.Viper.IsArray(path)
	return false
}

func (p *ViperConfiguration) String() string {
	if p == nil || p.Viper == nil {
		return ""
	}

	return fmt.Sprintf("%v", p.Viper.AllSettings())
}

type ViperConfigProvider struct {
}

func (p *ViperConfigProvider) LoadConfig(name string) Configuration {

	// initialize setting config
	conf := viper.New()
	conf.SetConfigName(name + ".default")     // name of config file (without extension)
	conf.AddConfigPath("configs/")            // path to look for the config file in
	conf.AddConfigPath("./../configs/")       // path to look for the config file in
	conf.AddConfigPath("./../../configs/")    // path to look for the config file in
	conf.AddConfigPath("./../../../configs/") // path to look for the config file in
	conf.AddConfigPath(".")                   // optionally look for config in the working directory

	err := conf.ReadInConfig() // Find and read the config file
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "LoadConfig %s default value failed: %v\n", name, err)
	}

	conf.SetConfigName(name)   // name of config file (without extension)
	err = conf.MergeInConfig() // Find and read the config file
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "LoadConfig %s failed: %v\n", name, err)
	}

	return NewViperConfiguration(conf)
}

func (p *ViperConfigProvider) ParseString(cfgStr string) Configuration {
	conf := viper.New()
	_ = conf.ReadConfig(bytes.NewBuffer([]byte(cfgStr)))
	return NewViperConfiguration(conf)
}
