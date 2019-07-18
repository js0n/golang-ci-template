package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/js0n/golang-ci-template"
)

var _ = Describe("GolangCiTemplate", func() {
	var output string

	BeforeEach(func() {
		output = HelloWorld()
	})

	Context("when HelloWorld is called", func() {
		It("should have the correct value", func() {
			Expect(output).To(Equal("Hello, world!"))
		})
	})
})
