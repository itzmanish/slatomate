package config

import (
	"os"
	"strings"

	"github.com/itzmanish/go-micro/v2/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CertFile string = `-----BEGIN CERTIFICATE-----
MIIEJTCCAo2gAwIBAgIRAOoNbqQu6eyRV/yPLWOuWZgwDQYJKoZIhvcNAQELBQAw
ZTEeMBwGA1UEChMVbWtjZXJ0IGRldmVsb3BtZW50IENBMR0wGwYDVQQLDBRtYW5p
c2hAbWFuaXNoLVhQUy0xMzEkMCIGA1UEAwwbbWtjZXJ0IG1hbmlzaEBtYW5pc2gt
WFBTLTEzMB4XDTIxMDkwNTA1MzUxNFoXDTIzMTIwNTA1MzUxNFowUTEnMCUGA1UE
ChMebWtjZXJ0IGRldmVsb3BtZW50IGNlcnRpZmljYXRlMSYwJAYDVQQLDB1tYW5p
c2hAbWFuaXNoLVhQUy0xMyAoTWFuaXNoKTCCASIwDQYJKoZIhvcNAQEBBQADggEP
ADCCAQoCggEBALR4Fx91iDiX/E4G8ZDPAErmMf5Et1KVa5sEjCqfJF+pv7nRqVJG
rs9iOG6whnyu3SjhMLdSut2DdcKUW4bkT3R1+2naj2+0bHilnMoivB/8H3LYYB6e
KX9PULppxdSgZ1LAfndJlMtELbo8IkVsWkMLhd4BoPcnx6EZZKDmi0cDaTVZo77K
khEgdeEHj+G6vkW7MCidyHNUei2pClWJnofkico1lBt8Y/lwh3hHBo5WFTuIEZOV
20A/L3IGxY0PSFKtxalwNaChwpQr52y8+ceiKVQvfe2hRNioNkpY/kbkCjxww592
oyaUTvRKPXDERloo8kjQ1Ndxe3pcvctBKo0CAwEAAaNkMGIwDgYDVR0PAQH/BAQD
AgWgMBMGA1UdJQQMMAoGCCsGAQUFBwMBMB8GA1UdIwQYMBaAFAeMVWuy0bLYhhvH
Fe0W2+aZ3FxWMBoGA1UdEQQTMBGCCWxvY2FsaG9zdIcEfwAAATANBgkqhkiG9w0B
AQsFAAOCAYEAAelEC9R1xuxDpMeVSjldpUHLk4kFZsCwMg5BGf6DvH2+NR8Lu7Ex
Qu1Q65jEHecs25XbhMT04RNs6JT0h55AdclXPPKHOXMSid+5vG2rSx2v4Bak5l4r
fznTN/fXmconjJ2ViDoChf3vG+/jQJmt1KgqWQhgTtjl5E2dXQgWpImni/+tBA0a
9IVAXyaOmnINxe46B73J5nIXJ69bvCJkayNV/bJP75mX703x3/+QVrZrVSFoj8qF
ZjL/18XZ6lmne9SpPQ6aXN9j0AZ9OzpzVuEw9jJCxzvp7en/JPMiqFg+jZQtsDvh
ogrt6MeDel9xreYs7wYkyUphnENHW3mjHV04HrKkt6myUVOL72odVNRZG2UIQnPf
fKbK3i1WEWMQXWTM8kPRKgyGQc0StPhc6OAXJQFt92d/ZhJFYNdf5PRaWjA9I3Nx
czhhr3r8oDn2c/Vr5DFfMcTrUm5I0mVxCECVo6cKsDi+wGSHFrBMa1VkCmQGg0lX
JYjvZOq+3Mmh
-----END CERTIFICATE-----
`

var CertKey string = `-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQC0eBcfdYg4l/xO
BvGQzwBK5jH+RLdSlWubBIwqnyRfqb+50alSRq7PYjhusIZ8rt0o4TC3Urrdg3XC
lFuG5E90dftp2o9vtGx4pZzKIrwf/B9y2GAenil/T1C6acXUoGdSwH53SZTLRC26
PCJFbFpDC4XeAaD3J8ehGWSg5otHA2k1WaO+ypIRIHXhB4/hur5FuzAonchzVHot
qQpViZ6H5InKNZQbfGP5cId4RwaOVhU7iBGTldtAPy9yBsWND0hSrcWpcDWgocKU
K+dsvPnHoilUL33toUTYqDZKWP5G5Ao8cMOfdqMmlE70Sj1wxEZaKPJI0NTXcXt6
XL3LQSqNAgMBAAECggEBAKjweoTim4CPFotciKpMfTOgRlCGty5B3hehrC0CCSTc
XTRwBpeUv3Q3uCg2a88wSqxIEjiq428V7xkVlJC0DYUyJQa0qO2i8qAGOL4owf8X
H8F8uI4w9RvOff3jomQnPFIFDN5SLU4TJtNeE571jZGRqeFnmO2FeaVfhgnxh5QT
E1kpECeUWQWcs2ZZqxEvvD1tBuNxQpM/8aY8ZYCfQIxPbobRQMoPaW9aM+u0B31h
pbEnx/ttotniu+aJJ30O0hpGJYfYzeLOePEbbwiu1Skj1pJgXBxUDhvUPhIMt78P
JSLqapXUoMxSoes1G9/JnFkZSRdNmTuuEBUjaQjjpcECgYEAxLGKnEfTllUpCQvm
/0cK+Nxm5SQDokscsJIF+6j+091PN/jbob0vDVSTD/6Rl9rx6LSYd4EuuoC+Ua5l
WUUbwfWaUfxU2yl44aiiR2jGuX1TKzQjKWa1HwlhD4fwDc6ejGvz+D/shFeCLrxC
t8bc0D0otiRJNnqDKd++ttAO+GkCgYEA6uI2Qhs0CXRau6kIfYf0sO9YZdJrcFr3
EGCmfQotGXvOd2v8tBac6DiCBOQsV+LuY3qv1kew2faJJV0gQjfzmFoe4LPdh9Vj
TG3m6T9MpR8VLyeJ4PRy6PItdtV1pGJgmNZKHSq/sjqsJoU618FVECLNd9YFLN/x
RAdyC6MFvIUCgYAIVyr0g4syVbweuRV7f6y0bswiLUvGJv85cYe1ay4bF8hTLEdL
7XTAUPTHedj8onkdkALjFmQ/3lOzrPx5M5gAuoRns5Z5kKGil/8BnizHEsxjCZvw
Fn3ZqhEmknIYc8l/VNiMj1FdL0TC7JK7rkAQyHcgehtspvdG4Ej7AYmQ+QKBgQCk
D312jym5sNvvWBHmHKB9NbC47lC9GcyYU+n2TLVTp3Z/U54e6+yNB2tJn1aZzJhW
Q8uuEEUm+VyyDGoL1qj/MXN/4CJMTnAdYmZ3ZQ9UBnH1jdhwmE3rB71Z059oQo7W
MldyJ/nds72q6kQ/j3qq2qRJn+PdhE4xBkJSfnvzQQKBgQC9lCZm2cP9CQNwsehh
Qz8+hmms3BloviNnyC2N3/7BF9MrpwtYWdzH+A47Rv0AyMhE+jGSr/Y9jWgHwZSb
ow6WFu584kq+6Vnk47fJayJHHE5qqa/vItPVGB2+osHLws2r37J9dGeXBYcrgC/3
1SZ9l870eg/IV5UyFQl1S+P/uA==
-----END PRIVATE KEY-----
`

// InitConfig reads in config file and ENV variables if set.
func InitConfig(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		configDir, err := os.UserConfigDir()
		cobra.CheckErr(err)

		// Search config in config directory with name ".slatomate" (without extension).
		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("slatomate")
		viper.SetDefault("auth_token", "")
		viper.SetDefault("oauth_url", "https://slack.com/oauth/v2/authorize?client_id=1001856848789.2347720282836&scope=&user_scope=users.profile:write")
		viper.SetDefault("service_host", "127.0.0.1:8081")
		viper.SafeWriteConfig()

	}

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logger.Debug("Using config file:", viper.ConfigFileUsed())
	}
}
