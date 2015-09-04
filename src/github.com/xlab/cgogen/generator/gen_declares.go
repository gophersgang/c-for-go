package generator

import (
	"fmt"
	"io"

	tl "github.com/xlab/cgogen/translator"
)

func (gen *Generator) transformDeclName(declName string, public bool) string {
	var name string
	var target tl.RuleTarget
	if public {
		target = tl.TargetPublic
	} else {
		target = tl.TargetPrivate
	}
	name = string(gen.tr.TransformName(target, declName))
	if len(name) == 0 {
		name = "_"
	}
	return name
}

func (gen *Generator) writeTypeDeclaration(wr io.Writer, decl tl.CDecl, ptrSpec tl.PointerSpec, public bool) {
	declName := gen.transformDeclName(decl.Name, public)
	goSpec := gen.tr.TranslateSpec(decl.Spec, ptrSpec)
	fmt.Fprintf(wr, "%s %s", declName, goSpec)
}

func (gen *Generator) writeEnumDeclaration(wr io.Writer, decl tl.CDecl, ptrSpec tl.PointerSpec, public bool) {
	declName := gen.transformDeclName(decl.Name, public)
	typeRef := gen.tr.TranslateSpec(decl.Spec, ptrSpec).String()
	if declName != typeRef {
		fmt.Fprintf(wr, "%s %s", declName, typeRef)
		writeSpace(wr, 1)
	}
}

func (gen *Generator) writeFunctionDeclaration(wr io.Writer, decl tl.CDecl, ptrSpec tl.PointerSpec, public bool) {
	var returnRef string
	spec := decl.Spec.(*tl.CFunctionSpec)
	if spec.Return != nil {
		returnRef = gen.tr.TranslateSpec(spec.Return.Spec, ptrSpec).String()
	}
	if public {
		declName := string(gen.tr.TransformName(tl.TargetFunction, decl.Name))
		if returnRef == declName {
			declName = string(gen.tr.TransformName(tl.TargetFunction, "new_"+decl.Name))
		}
		fmt.Fprintf(wr, "// %s method as declared in %s\n", declName, tl.SrcLocation(decl.Pos))
		fmt.Fprintf(wr, "func %s", declName)
	} else {
		declName := string(gen.tr.TransformName(tl.TargetPrivate, decl.Name))
		goSpec := gen.tr.TranslateSpec(decl.Spec, ptrSpec)
		fmt.Fprintf(wr, "%s %s", declName, goSpec)
	}
	gen.writeFunctionParams(wr, decl.Name, decl.Spec)
	if len(returnRef) > 0 {
		fmt.Fprintf(wr, " %s", returnRef)
	}
	if public {
		gen.writeFunctionBody(wr, decl)
		writeSpace(wr, 1)
	}
}

func (gen *Generator) writeStructDeclaration(wr io.Writer, decl tl.CDecl, ptrSpec tl.PointerSpec, public bool) {
	declName := gen.transformDeclName(decl.Name, public)
	if tag := decl.Spec.GetBase(); len(tag) > 0 {
		// refSpec := &tl.CTypeSpec{
		// 	Base:      spec.Tag,
		// 	Arrays:    spec.GetArrays(),
		// 	VarArrays: spec.GetVarArrays(),
		// 	Pointers:  spec.GetPointers(),
		// }
		goSpec := gen.tr.TranslateSpec(decl.Spec, ptrSpec)
		fmt.Fprintf(wr, "%s %s", declName, goSpec)
		return
	}
	if !decl.IsTemplate() {
		return
	}

	fmt.Fprintf(wr, "%s struct {", declName)
	writeSpace(wr, 1)
	gen.writeStructMembers(wr, decl.Name, decl.Spec)
	writeEndStruct(wr)
}
