package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/maxpoletaev/directcgo/internal/codegen"
)

type Options struct {
	pkg  string
	arch string
}

func parseOptions() Options {
	var opts Options

	flag.StringVar(&opts.arch, "arch", "", "architecture")
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatalf("usage: %s -arch=<archname> <package>", os.Args[0])
	}

	opts.pkg = flag.Arg(0)

	return opts
}

func parseList(s string) []string {
	j := 0
	parts := strings.Split(s, ",")
	for i := 0; i < len(parts); i++ {
		part := strings.TrimSpace(parts[i])
		if part != "" {
			parts[j] = part
			j++
		}
	}
	return parts[:j]
}

func main() {
	log.Default().SetFlags(0)

	opts := parseOptions()

	var archList []string
	if opts.arch == "" {
		archList = []string{
			codegen.ArchARM64,
		}
	} else {
		archList = parseList(opts.arch)
		for _, arch := range archList {
			if _, ok := codegen.ValidArchitectures[arch]; !ok {
				log.Fatalf("unsupported architecture: %s", arch)
			}
		}
	}

	if err := codegen.Run(opts.pkg, archList); err != nil {
		log.Fatalf("error: %v", err)
	}
}
