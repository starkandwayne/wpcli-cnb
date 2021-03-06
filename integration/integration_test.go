package integration

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	"github.com/cloudfoundry/dagger"
	. "github.com/onsi/gomega"
)

var (
	bpDir, phpBP, phpWebBP, httpdBP, wpcliBP string
)

func TestIntegration(t *testing.T) {
	var err error
	Expect := NewWithT(t).Expect
	bpDir, err = dagger.FindBPRoot()
	Expect(err).NotTo(HaveOccurred())
	wpcliBP, err = dagger.PackageBuildpack(bpDir)
	Expect(err).ToNot(HaveOccurred())
	defer os.RemoveAll(wpcliBP)

	phpBP, err = dagger.GetLatestBuildpack("php-cnb")
	phpWebBP, err = dagger.GetLatestBuildpack("php-web-cnb")
	httpdBP, err = dagger.GetLatestBuildpack("httpd-cnb")
	Expect(err).ToNot(HaveOccurred())
	defer os.RemoveAll(phpBP)
	defer os.RemoveAll(phpWebBP)
	defer os.RemoveAll(httpdBP)

	spec.Run(t, "Integration", testIntegration, spec.Report(report.Terminal{}))
}

func testIntegration(t *testing.T, when spec.G, it spec.S) {
	var Expect func(interface{}, ...interface{}) GomegaAssertion
	it.Before(func() {
		Expect = NewWithT(t).Expect
	})

	it("should build a working OCI image for a simple app", func() {
		app, err := dagger.PackBuild(filepath.Join("fixtures", "phpapp"), wpcliBP, phpBP, httpdBP, phpWebBP)
		Expect(err).ToNot(HaveOccurred())
		// defer app.Destroy()

		Expect(app.Start()).To(Succeed())
		Expect(app.HTTPGetBody("/")).To(ContainSubstring("PHP Version"))
	})

	// when("the app is pushed twice", func() {
	// 	it("does not reinstall go modules", func() {
	// 		app, err := dagger.PackBuild(filepath.Join("fixtures", "phpapp"), phpURI, wpcliURI)
	// 		Expect(err).ToNot(HaveOccurred())
	// 		defer app.Destroy()

	// 		Expect(app.Start()).To(Succeed())
	// 		Expect(app.HTTPGetBody("/")).To(ContainSubstring("PHP Version"))

	// 		_, imageID, _, err := app.Info()
	// 		Expect(err).NotTo(HaveOccurred())

	// 		app, err = dagger.PackBuildNamedImage(imageID, appDir, phpURI, wpcliURI)
	// 		Expect(err).ToNot(HaveOccurred())

	// 		Expect(app.Start()).To(Succeed())
	// 		Expect(app.HTTPGetBody("/")).To(ContainSubstring("PHP Version"))
	// 	})
	// })

	// when("the app is vendored", func() {
	// 	it("builds an OCI image without downloading any extra packages", func() {
	// 		app, err := dagger.PackBuild(filepath.Join("fixtures", "phpapp"), phpURI, wpcliURI)
	// 		Expect(err).ToNot(HaveOccurred())

	// 		// Expect(app.BuildLogs()).NotTo(MatchRegexp(goDownloading))

	// 		Expect(app.Start()).To(Succeed())
	// 		Expect(app.HTTPGetBody("/")).To(ContainSubstring("PHP Version"))
	// 	})
	// })
}
