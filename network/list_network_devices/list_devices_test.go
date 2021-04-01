package list_device_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/google/gopacket/pcap"
)

func TestList(t *testing.T) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Println("\nName:", device.Name)
		fmt.Println("Description:", device.Description)
		for _, address := range device.Addresses {
			fmt.Println(address.IP)
			fmt.Println(address.Netmask)
		}
	}
}
