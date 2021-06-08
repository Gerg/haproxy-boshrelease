package acceptance_tests

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Access Control", func() {
	AfterEach(func() {
		deleteDeployment()
	})

	BeforeEach(func() {
		haproxyBackendPort := 12000
		haproxyInfo, _ := deployHAProxy(haproxyBackendPort, []string{}, map[string]interface{}{})

		// Wait for monit to stabilise
		time.Sleep(time.Second * 30)

		dumpHAProxyConfig(haproxyInfo)

		fmt.Println("--------------------- DEBUG")
		time.Sleep(time.Hour * 8)
	})

	// FIXME: remove focus
	FIt("Allows IPs in whitelisted CIDRS", func() {

	})

	It("Rejects IPs in non-whitelisted CIDRS", func() {

	})

	It("Rejects IPs in blacklisted CIDRS", func() {

	})

	// https://www.haproxy.com/documentation/aloha/latest/security/packetshield/blacklist/
	// > When you add an address to the IP source blacklist, it is automatically removed from the whitelist.
	It("Rejects IPs that are whitelisted and blacklisted", func() {

	})
})
