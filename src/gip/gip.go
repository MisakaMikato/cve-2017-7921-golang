package gip

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func qpow(x, y uint64) uint64 {
	result := uint64(1)
	for i := y; i > 0; i >>= 1 {
		if i&1 != 0 {
			result *= x
		}
		x *= x
	}
	return result
}

// IntegerToIP transforms integer into ip address
func IntegerToIP(integerIP uint32, subnetMask int) string {
	var ip []int
	temp, cnt := 0, 0
	for integerIP != 0 {
		if integerIP&1 == 1 {
			temp += int(qpow(2, uint64(cnt)))
		}
		integerIP >>= 1
		cnt++
		if cnt >= 8 || integerIP == 0 {
			ip = append(ip, temp)
			temp, cnt = 0, 0
		}
	}
	if subnetMask == 32 {
		return fmt.Sprintf("%d.%d.%d.%d", ip[3], ip[2], ip[1], ip[0])
	}
	return fmt.Sprintf("%d.%d.%d.%d/%d", ip[3], ip[2], ip[1], ip[0], subnetMask)
}

// IPToInteger transforms ip address into integer
func IPToInteger(ip string) uint32 {
	var integerIP uint32 = 0
	subList := strings.Split(ip, ".")
	for i := 0; i < len(subList); i++ {
		tmp, _ := strconv.Atoi(subList[i])
		integerIP += uint32(uint64(tmp) * qpow(2, uint64(24-i*8)))
	}
	return integerIP
}

// GetIPSubnet splits an IP address segment up into several subnets with specific subnet mask.
func GetIPSubnet(ip string, subnetMask int) ([]string, error) {
	// subnetList looks like: ['192.168.0.0/24', '192.168.1.0/24']
	if subnetMask > 32 || subnetMask < 0 {
		return nil, errors.New("subnet mask out of range")
	}
	var subnetList []string
	var netAddress string
	var netMask int
	tmpPos := strings.Index(ip, "/")
	if tmpPos != -1 {
		netAddress = ip[:tmpPos]
		netMask, _ = strconv.Atoi(ip[tmpPos+1:])
	} else {
		netAddress = ip
		netMask = 32
	}
	if netMask >= subnetMask {
		return []string{ip}, nil
	}
	netAddressInt := IPToInteger(netAddress)
	loop := int(qpow(2, uint64(subnetMask-netMask)))
	for i := 0; i < loop; i++ {
		subnetList = append(subnetList, IntegerToIP(netAddressInt, subnetMask))
		netAddressInt += (1 << (32 - subnetMask))
	}

	return subnetList, nil
}

// CheckIPFormat returns true if param is a correct ip format, else returns false.
func CheckIPFormat(ip []byte) bool {
	pattern := `^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)((/(3[0-2]|[0-1]?\d))?)$`
	regExp := regexp.MustCompile(pattern)
	return regExp.Match(ip)
}
