package testing

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinko(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}

var _ = Describe("Hello", func() {
	Describe("Greeting someone", func() {
		Context("Alexander", func() {
			It("should greet Alexander", func() {
				Expect(Hello("Alexander")).To(Equal("Hello, Alexander!"))
			})
		})
	})

	Describe("General greeting", func() {
		Context("no name passed in", func() {
			It("should greet the world", func() {
				Expect(Hello("")).To(Equal("Hello, World!"))
			})
		})
	})
})
