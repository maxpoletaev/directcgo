package codegen

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"golang.org/x/tools/go/packages"
)

const (
	defaultFrameSize = 65536
)

const (
	ArchARM64 = "arm64"
	ArchAMD64 = "amd64"
)

var ValidArchitectures = map[string]struct{}{
	ArchARM64: {},
	ArchAMD64: {},
}

type platform interface {
	GenerateFunc(buf *builder, f *Function)
	Name() string
}

func composeAssemblyFile(buf *builder, outFile string, header *headerVars) error {
	tmpFilePath := outFile + ".tmp"

	f, err := os.OpenFile(tmpFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	// Write file header
	if err := header.Format(f); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	// Write generated code
	if _, err = f.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	if err = f.Close(); err != nil {
		return fmt.Errorf("failed to close file: %w", err)
	}

	if err = os.Rename(tmpFilePath, outFile); err != nil {
		return fmt.Errorf("failed to rename file: %w", err)
	}

	return nil
}

func generatePackage(pkg *packages.Package, buf *builder, funcs []*Function, platform platform, fullCommand string) error {
	pkgDir := path.Dir(pkg.GoFiles[0])
	asmFilePath := path.Join(pkgDir, fmt.Sprintf("directcgo_%s.s", platform.Name()))

	log.Printf("generating %s", asmFilePath)

	for i, fn := range funcs {
		platform.GenerateFunc(buf, fn)
		if i != len(funcs)-1 {
			buf.NL()
		}
	}

	err := composeAssemblyFile(
		buf,
		asmFilePath,
		&headerVars{
			arch:    platform.Name(),
			fullcmd: fullCommand,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func align(offset, alignment int) int {
	return (offset + alignment - 1) & ^(alignment - 1)
}

func Run(pkgPath string, archList []string) error {
	cfg := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedSyntax |
			packages.NeedTypes |
			packages.NeedTypesInfo,
	}

	pkgs, err := packages.Load(cfg, pkgPath)
	if err != nil {
		return fmt.Errorf("failed to load packages: %v", err)
	}

	if len(pkgs) != 1 {
		return fmt.Errorf("expected exactly one package, got %d", len(pkgs))
	}

	fullCmd := "directcgo " + strings.Join(os.Args[1:], " ")
	buf := new(builder)
	pkg := pkgs[0]

	funcs, err := parsePackage(pkg)
	if err != nil {
		return fmt.Errorf("ast parsing failed: %w", err)
	}

	for _, archName := range archList {
		buf.Reset()

		var arch platform

		switch archName {
		case ArchARM64:
			arch = newARM64()
		case ArchAMD64:
			arch = newAMD64()
		default:
			return fmt.Errorf("unknown architecture: %s", archName)
		}

		if err = generatePackage(pkg, buf, funcs, arch, fullCmd); err != nil {
			return fmt.Errorf("failed to process package: %w", err)
		}
	}

	return nil
}
