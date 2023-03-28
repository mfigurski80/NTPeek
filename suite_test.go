package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

func TestNTPeek(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NTPeek Suite")
}

/// Integration Tests for `NTPeek`

var _ = Describe("NTPeek Integration", func() {

	BeforeEach(func() {
		if _, err := os.Stat("./nt"); err != nil {
			err := exec.Command("go", "generate", "&&", "go", "build", "-o", "nt", ".").Run()
			if err != nil {
				panic(fmt.Sprintf("failed to build `nt` binary, required for integration testing: %s", err))
			}
		}
	})

	Context("when running `nt v`", func() {
		It("should return version information", func() {
			toRun := []string{"./nt", "v"}
			output, code := captureCrasher(toRun)
			Expect(code).To(Equal(0), "should exit with code 0")
			Expect(output).To(ContainSubstring("version"))
			Expect(output).To(MatchRegexp(`[0-9a-f]{40}`), "should contain a git commit hash")
		})
	})

	Context("when running `nt h`", func() {
		It("should return help information", func() {
			Skip("fix!")
			toRun := []string{"./nt", "h"}
			output, code := captureCrasher(toRun)
			Expect(code).To(Equal(0), "should exit with code 0")
			Expect(output).To(ContainSubstring("Usage"))
		})
		It("should match snapshot", func() {
			Skip("fix! implement snapshot!")
			toRun := []string{"./nt", "h"}
			_, code := captureCrasher(toRun)
			Expect(code).To(Equal(0), "should exit with code 0")
		})
	})

	Context("when running `nt p`", func() {
		// note these credentials are burned, for test db
		const TEST_DB_ID = "979bf78281914ca5895555168b2f7396u"
		const TEST_ACCESS = "secret_rhsxWWqTWhEd1pLlEOLB2z5eVfilG1iqPGPjeqSU934"

		It("should return a list of items", func() {
			Skip("fix!")
			toRun := []string{"./nt", TEST_ACCESS, TEST_DB_ID, "p"}
			output, code := captureCrasher(toRun)
			Expect(code).To(Equal(0))
			Expect(output).To(ContainSubstring("NTPeek"))
		})
	})

	Context("when running unknown commands", func() {
		It("should return an error", func() {
			toRun := []string{"./nt", "unknown"}
			output, code := captureCrasher(toRun)
			Expect(code).To(Equal(1))
			Expect(output).To(ContainSubstring("unknown command"))
			Expect(output).To(ContainSubstring("./nt unknown"))
			Expect(output).To(ContainSubstring("Usage"))
		})
	})

})

/// Utilities for safe IO Capture, from:
/// https://medium.com/@hau12a1/golang-capturing-log-println-and-fmt-println-output-770209c791b4

func captureBufferOutput(f func()) string {
	out_bak := os.Stdout
	err_bak := os.Stderr
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = w
	os.Stderr = w
	defer func() {
		os.Stdout = out_bak
		os.Stderr = err_bak
	}()
	buf := gbytes.NewBuffer()
	go func() {
		defer func() {
			_ = recover()
			w.Close()
		}()
		f()
	}()
	io.Copy(buf, r)
	return string(buf.Contents())
}

func captureCrasher(args []string) (string, int) {
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("%s\n%s", err, string(out)), 1
	}
	return string(out), cmd.ProcessState.ExitCode()
}
