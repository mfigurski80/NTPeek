package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/tommy351/goldga"
)

func TestNTPeek(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NTPeek Suite")
}

/// Integration Tests for `NTPeek`

var _ = Describe("Integration test", func() {

	BeforeEach(func() {
		inf, err := os.Stat("./nt")
		if inf == nil || err != nil {
			out, err := exec.Command("go", "generate").CombinedOutput()
			if err != nil {
				panic(fmt.Sprintf("failed to generate `nt` data, required for integration testing: %s\n%s", err, string(out)))
			}
			out, err = exec.Command("go", "build", "-o", "./nt").CombinedOutput()
			if err != nil {
				panic(fmt.Sprintf("failed to build `nt` binary, required for integration testing: %s\n%s", err, string(out)))
			}
		}
	})

	Context("when running `nt v`", func() {
		It("should return version information", func() {
			toRun := []string{"./nt", "v"}
			output, code := captureCrasher(toRun)
			Expect(output).To(ContainSubstring("version"))
			Expect(output).To(MatchRegexp(`[0-9a-f]{40}`), "should contain a git commit hash")
			Expect(code).To(Equal(0), "should exit with code 0")
		})
	})

	Context("when running `nt h`", func() {
		It("should return help information", func() {
			toRun := []string{"./nt", "h"}
			output, code := captureCrasher(toRun)
			Expect(output).To(ContainSubstring("Usage"))
			Expect(code).To(Equal(0), "should exit with code 0")
		})
		It("should match snapshot", func() {
			toRun := []string{"./nt", "h"}
			output, code := captureCrasher(toRun)
			Expect(output).To(goldga.Match())
			Expect(code).To(Equal(0), "should exit with code 0")
		})
	})

	Context("when running `nt p`", func() {
		// note these credentials are burned, for test db
		const TEST_DB_ID = "979bf78281914ca5895555168b2f7396"
		const TEST_ACCESS = "secret_rhsxWWqTWhEd1pLlEOLB2z5eVfilG1iqPGPjeqSU934"
		// https://www.notion.so/mikof/979bf78281914ca5895555168b2f7396?v=2a121bd645e5476fb0c6a0fe3d44366d

		When("no db id or token are provided", func() {
			It("should fail with error message", func() {
				toRun := []string{"./nt", "p"}
				output, code := captureCrasher(toRun)
				Expect(output).To(ContainSubstring("Secret"))
				Expect(output).To(ContainSubstring("Database"))
				Expect(code).To(Equal(1))
			})
			It("should match snapshot", func() {
				toRun := []string{"./nt", "p"}
				output, _ := captureCrasher(toRun)
				Expect(output).To(goldga.Match())
			})
		})
		It("should return a list of items", func() {
			toRun := []string{"./nt", TEST_ACCESS, TEST_DB_ID, "p"}
			output, code := captureCrasher(toRun)
			Expect(output).ToNot(ContainSubstring("Usage"))
			Expect(output).ToNot(ContainSubstring("Err"))
			spl := strings.Split(output, "\n")
			Expect(len(spl) > 2).To(BeTrue(), "should have at least 3 items")
			Expect(code).To(Equal(0))
		})
		It("should match simple snapshot", func() {
			toRun := []string{"./nt", TEST_ACCESS, TEST_DB_ID,
				"--select=\"%Class:right% // %Name:left% %Due:full% %_p% %_id:short%\""}
			output, code := captureCrasher(toRun)
			Expect(output).To(goldga.Match())
			Expect(code).To(Equal(0))
		})
		It("should match complex snapshot", func() {
			toRun := []string{"./nt", TEST_ACCESS, TEST_DB_ID,
				"--select=\"%Class:right% // %Name:left% %Due:full% %_p% %_id:short%\"",
				"--sort", "Due:desc",
				"--filter", "Due:date >= 2023/01/25",
				"--limit=2",
			}
			output, code := captureCrasher(toRun)
			Expect(output).To(goldga.Match())
			Expect(code).To(Equal(0))
		})
	})

	Context("when running malformed commands", func() {
		It("errors on no arguments", func() {
			toRun := []string{"./nt"}
			output, code := captureCrasher(toRun)
			Expect(output).To(ContainSubstring("argument"))
			Expect(code).To(Equal(1))
		})
		It("matches no args snapshot", func() {
			toRun := []string{"./nt"}
			output, _ := captureCrasher(toRun)
			Expect(output).To(goldga.Match())
		})
		It("errors on unknown command", func() {
			toRun := []string{"./nt", "unknown"}
			output, code := captureCrasher(toRun)
			Expect(output).To(ContainSubstring("unknown command"))
			Expect(output).To(ContainSubstring("./nt unknown"), "should show passed args")
			Expect(output).To(ContainSubstring("Usage"))
			Expect(code).To(Equal(1))
		})
		It("matches unknown command snapshot", func() {
			toRun := []string{"./nt", "unknown"}
			output, _ := captureCrasher(toRun)
			Expect(output).To(goldga.Match())
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
