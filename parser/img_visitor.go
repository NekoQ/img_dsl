// Generated from Img.g4 by ANTLR 4.7.

package parser // Img

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ImgParser.
type ImgVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ImgParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by ImgParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by ImgParser#openFile.
	VisitOpenFile(ctx *OpenFileContext) interface{}

	// Visit a parse tree produced by ImgParser#openFolder.
	VisitOpenFolder(ctx *OpenFolderContext) interface{}

	// Visit a parse tree produced by ImgParser#fileName.
	VisitFileName(ctx *FileNameContext) interface{}

	// Visit a parse tree produced by ImgParser#imageType.
	VisitImageType(ctx *ImageTypeContext) interface{}

	// Visit a parse tree produced by ImgParser#folderName.
	VisitFolderName(ctx *FolderNameContext) interface{}

	// Visit a parse tree produced by ImgParser#action_.
	VisitAction_(ctx *Action_Context) interface{}

	// Visit a parse tree produced by ImgParser#actionType.
	VisitActionType(ctx *ActionTypeContext) interface{}

	// Visit a parse tree produced by ImgParser#crop.
	VisitCrop(ctx *CropContext) interface{}

	// Visit a parse tree produced by ImgParser#rotate.
	VisitRotate(ctx *RotateContext) interface{}

	// Visit a parse tree produced by ImgParser#flipX.
	VisitFlipX(ctx *FlipXContext) interface{}

	// Visit a parse tree produced by ImgParser#flipY.
	VisitFlipY(ctx *FlipYContext) interface{}

	// Visit a parse tree produced by ImgParser#resize.
	VisitResize(ctx *ResizeContext) interface{}

	// Visit a parse tree produced by ImgParser#export.
	VisitExport(ctx *ExportContext) interface{}
}
