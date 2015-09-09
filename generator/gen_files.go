package generator

import (
	"fmt"
	"io"
	"strings"
	"time"
)

func genLabel() string {
	tpl := "WARNING: This file has automatically been generated on %s.\nBy http://git.io/cgogen. DO NOT EDIT."
	return fmt.Sprintf(tpl, time.Now().Format(time.RFC1123))
}

func (gen *Generator) WriteDoc(wr io.Writer) bool {
	var hasDoc bool
	if len(gen.cfg.PackageLicense) > 0 {
		writeTextBlock(wr, gen.cfg.PackageLicense)
		writeSpace(wr, 1)
		hasDoc = true
	}
	writeTextBlock(wr, genLabel())
	writeSpace(wr, 1)
	if len(gen.cfg.PackageDescription) > 0 {
		writeLongTextBlock(wr, gen.cfg.PackageDescription)
		hasDoc = true
	}
	writePackageName(wr, gen.pkg)
	writeSpace(wr, 1)
	return hasDoc
}

func (gen *Generator) WriteIncludes(wr io.Writer) {
	writeStartComment(wr)
	writePkgConfig(wr, gen.cfg.PkgConfigOpts)
	writeFlagSet(wr, gen.cfg.CPPFlags)
	writeFlagSet(wr, gen.cfg.CXXFlags)
	writeFlagSet(wr, gen.cfg.CFlags)
	writeFlagSet(wr, gen.cfg.LDFlags)
	for _, path := range gen.cfg.SysIncludes {
		writeSysInclude(wr, path)
	}
	for _, path := range gen.cfg.Includes {
		writeInclude(wr, path)
	}
	if !hasStdLib(gen.cfg.SysIncludes) {
		writeCStdIncludes(wr)
	}
	writeEndComment(wr)
	fmt.Fprintln(wr, `import "C"`)
	writeSpace(wr, 1)
}

func hasStdLib(paths []string) bool {
	for _, str := range paths {
		if str == "stdlib.h" {
			return true
		}
	}
	return false
}

func (gen *Generator) writeGoHelpersHeader(wr io.Writer) {
	writeTextBlock(wr, gen.cfg.PackageLicense)
	writeSpace(wr, 1)
	writeTextBlock(wr, genLabel())
	writeSpace(wr, 1)
	writePackageName(wr, gen.pkg)
	writeSpace(wr, 1)
	gen.WriteIncludes(wr)
}

func (gen *Generator) writeCHelpersHeader(wr io.Writer) {
	if len(gen.cfg.PackageLicense) > 0 {
		writeTextBlock(wr, gen.cfg.PackageLicense)
		writeSpace(wr, 1)
	}
	writeTextBlock(wr, genLabel())
	writeSpace(wr, 1)
	writeCStdIncludes(wr)
	writeSpace(wr, 1)
}

func writeCStdIncludes(wr io.Writer) {
	fmt.Fprintln(wr, "#include <stdlib.h>")
}

func (gen *Generator) WritePackageHeader(wr io.Writer) {
	writeTextBlock(wr, gen.cfg.PackageLicense)
	writeSpace(wr, 1)
	writeTextBlock(wr, genLabel())
	writeSpace(wr, 1)
	writePackageName(wr, gen.pkg)
	writeSpace(wr, 1)
}

func writeFlagSet(wr io.Writer, flags ArchFlagSet) {
	if len(flags.Name) == 0 {
		return
	}
	if len(flags.Flags) == 0 {
		return
	}
	fmt.Fprintf(wr, "#cgo %s %s: %s\n",
		strings.Join(flags.Arch, ","),
		flags.Name,
		strings.Join(flags.Flags, " "),
	)
}

func writeSysInclude(wr io.Writer, path string) {
	fmt.Fprintf(wr, "#include <%s>\n", path)
}

func writeInclude(wr io.Writer, path string) {
	fmt.Fprintf(wr, "#include \"%s\"\n", path)
}

func writePkgConfig(wr io.Writer, opts []string) {
	if len(opts) == 0 {
		return
	}
	fmt.Fprintf(wr, "#cgo pkg-config: %s\n", strings.Join(opts, " "))
}

func writeStartComment(wr io.Writer) {
	fmt.Fprintln(wr, "/*")
}

func writeEndComment(wr io.Writer) {
	fmt.Fprintln(wr, "*/")
}

func writePackageName(wr io.Writer, name string) {
	if len(name) == 0 {
		name = "main"
	}
	fmt.Fprintf(wr, "package %s\n", name)
}

func writeLongTextBlock(wr io.Writer, text string) {
	if len(text) == 0 {
		return
	}
	writeStartComment(wr)
	fmt.Fprint(wr, text)
	writeSpace(wr, 1)
	writeEndComment(wr)
}

func writeTextBlock(wr io.Writer, text string) {
	if len(text) == 0 {
		return
	}
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fmt.Fprintf(wr, "// %s\n", line)
	}
}

func writeSourceBlock(wr io.Writer, src string) {
	if len(src) == 0 {
		return
	}
	fmt.Fprint(wr, src)
	writeSpace(wr, 1)
}