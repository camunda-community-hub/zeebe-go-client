package integration_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "github.com/zeebe-io/zeebe/clients/go/pb"

    "github.com/zeebe-io/zeebe/clients/go"
)

var _ = Describe("should send HealthRequest to Gateway and receive HealthResponse", func() {
	var client zbc.ZBClient

	BeforeEach(func() {
		c, e := zbc.NewZBClient("0.0.0.0:26500")
		Expect(e).To(BeNil())
		Expect(c).NotTo(BeNil())
		client = c
	})

	AfterEach(func() {
		client.Close()
	})

	Context("health check", func() {
		It("request with correct response", func() {
			response, err := client.NewHealthCheckCommand().Send()

			Expect(len(response.Brokers)).To(Equal(1))
			Expect(len(response.Brokers[0].Partitions)).To(Equal(1))
            Expect(response.Brokers[0].Partitions[0].PartitionId).To(Equal(int32(0)))
            Expect(response.Brokers[0].Partitions[0].Role).To(Equal(pb.Partition_LEADER))

			Expect(err).To(BeNil())
			Expect(response).NotTo(BeNil())
		})
	})

})
