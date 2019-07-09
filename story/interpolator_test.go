package story_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/williammartin/storywriter/story"
)

var _ = Describe("Interpolator", func() {

	var (
		interpolator Interpolator

		template string
		words    []string

		story string
	)

	BeforeEach(func() {
		interpolator = Interpolator{}

		template = ""
		words = []string{}
	})

	JustBeforeEach(func() {
		story = interpolator.Interpolate(template, words)
	})

	When("given a draft without any words", func() {
		BeforeEach(func() {
			template = "This is the template"
		})

		It("returns the template", func() {
			Expect(story).To(Equal(template))
		})
	})

	When("given a draft with many words", func() {
		BeforeEach(func() {
			template = "This is an &1 &2"
			words = []string{"interpolated", "story"}
		})

		It("returns the interpolated template", func() {
			Expect(story).To(Equal("This is an interpolated story"))
		})
	})

	When("given a draft that has out of order directives", func() {
		BeforeEach(func() {
			template = "This is an &2 &1"
			words = []string{"story", "interpolated"}
		})

		It("returns the interpolated template correctly ordered", func() {
			Expect(story).To(Equal("This is an interpolated story"))
		})
	})

	When("given a template that contains ampersands", func() {
		BeforeEach(func() {
			template = "This is an &1 &2 with an & to be confusing"
			words = []string{"interpolated", "story"}
		})

		It("returns the interpolated template with the ampersand untouched", func() {
			Expect(story).To(Equal("This is an interpolated story with an & to be confusing"))
		})
	})
})
