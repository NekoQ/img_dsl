// Generated from Img.g4 by ANTLR 4.7.

package parser // Img

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ImgListener is a complete listener for a parse tree produced by ImgParser.
type ImgListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterOpenFile is called when entering the openFile production.
	EnterOpenFile(c *OpenFileContext)

	// EnterOpenFolder is called when entering the openFolder production.
	EnterOpenFolder(c *OpenFolderContext)

	// EnterFileName is called when entering the fileName production.
	EnterFileName(c *FileNameContext)

	// EnterImageType is called when entering the imageType production.
	EnterImageType(c *ImageTypeContext)

	// EnterFolderName is called when entering the folderName production.
	EnterFolderName(c *FolderNameContext)

	// EnterAction_ is called when entering the action_ production.
	EnterAction_(c *Action_Context)

	// EnterActionType is called when entering the actionType production.
	EnterActionType(c *ActionTypeContext)

	// EnterCrop is called when entering the crop production.
	EnterCrop(c *CropContext)

	// EnterRotate is called when entering the rotate production.
	EnterRotate(c *RotateContext)

	// EnterFlipX is called when entering the flipX production.
	EnterFlipX(c *FlipXContext)

	// EnterFlipY is called when entering the flipY production.
	EnterFlipY(c *FlipYContext)

	// EnterResize is called when entering the resize production.
	EnterResize(c *ResizeContext)

	// EnterExport is called when entering the export production.
	EnterExport(c *ExportContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitOpenFile is called when exiting the openFile production.
	ExitOpenFile(c *OpenFileContext)

	// ExitOpenFolder is called when exiting the openFolder production.
	ExitOpenFolder(c *OpenFolderContext)

	// ExitFileName is called when exiting the fileName production.
	ExitFileName(c *FileNameContext)

	// ExitImageType is called when exiting the imageType production.
	ExitImageType(c *ImageTypeContext)

	// ExitFolderName is called when exiting the folderName production.
	ExitFolderName(c *FolderNameContext)

	// ExitAction_ is called when exiting the action_ production.
	ExitAction_(c *Action_Context)

	// ExitActionType is called when exiting the actionType production.
	ExitActionType(c *ActionTypeContext)

	// ExitCrop is called when exiting the crop production.
	ExitCrop(c *CropContext)

	// ExitRotate is called when exiting the rotate production.
	ExitRotate(c *RotateContext)

	// ExitFlipX is called when exiting the flipX production.
	ExitFlipX(c *FlipXContext)

	// ExitFlipY is called when exiting the flipY production.
	ExitFlipY(c *FlipYContext)

	// ExitResize is called when exiting the resize production.
	ExitResize(c *ResizeContext)

	// ExitExport is called when exiting the export production.
	ExitExport(c *ExportContext)
}
