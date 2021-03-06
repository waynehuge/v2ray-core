// +build json

package freedom

import (
	"encoding/json"
	"errors"
	"strings"

	"v2ray.com/core/proxy/registry"
)

func (this *Config) UnmarshalJSON(data []byte) error {
	type JsonConfig struct {
		DomainStrategy string `json:"domainStrategy"`
		Timeout        uint32 `json:"timeout"`
	}
	jsonConfig := new(JsonConfig)
	if err := json.Unmarshal(data, jsonConfig); err != nil {
		return errors.New("Freedom: Failed to parse config: " + err.Error())
	}
	this.DomainStrategy = Config_AS_IS
	domainStrategy := strings.ToLower(jsonConfig.DomainStrategy)
	if domainStrategy == "useip" || domainStrategy == "use_ip" {
		this.DomainStrategy = Config_USE_IP
	}
	this.Timeout = jsonConfig.Timeout
	return nil
}

func init() {
	registry.RegisterOutboundConfig("freedom", func() interface{} { return new(Config) })
}
