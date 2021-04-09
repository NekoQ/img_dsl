// Generated from Img.g4 by ANTLR 4.7.

package parser // Img

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseImgVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseImgVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseImgVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseImgVisitor) VisitOpenFile(ctx *OpenFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseImgVisitor) VisitOpenFolder(ctx *OpenFolderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseImgVisitor) VisitFileName(ctx *FileNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseImgVisitor) VisitImageType(ctx *ImageTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseImgVisitor) VisitFolderName(ctx *FolderNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseImgVisitor) VisitAction_(ctx *Action_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseImgVisitor) VisitActionType(ctx *ActionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseImgVisitor) VisitCrop(ctx *CropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseImgVisitor) VisitRotate(ctx *RotateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseImgVisitor) VisitFlipX(ctx *FlipXContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseImgVisitor) VisitFlipY(ctx *FlipYContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseImgVisitor) VisitResize(ctx *ResizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseImgVisitor) VisitExport(ctx *ExportContext) interface{} {
	return v.VisitChildren(ctx)
}
