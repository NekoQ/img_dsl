// Generated from Img.g4 by ANTLR 4.7.

package parser // Img

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseImgListener is a complete listener for a parse tree produced by ImgParser.
type BaseImgListener struct{}

var _ ImgListener = &BaseImgListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseImgListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseImgListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseImgListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseImgListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseImgListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseImgListener) ExitStart(ctx *StartContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseImgListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseImgListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterOpenFile is called when production openFile is entered.
func (s *BaseImgListener) EnterOpenFile(ctx *OpenFileContext) {}

// ExitOpenFile is called when production openFile is exited.
func (s *BaseImgListener) ExitOpenFile(ctx *OpenFileContext) {}

// EnterOpenFolder is called when production openFolder is entered.
func (s *BaseImgListener) EnterOpenFolder(ctx *OpenFolderContext) {}

// ExitOpenFolder is called when production openFolder is exited.
func (s *BaseImgListener) ExitOpenFolder(ctx *OpenFolderContext) {}

// EnterFileName is called when production fileName is entered.
func (s *BaseImgListener) EnterFileName(ctx *FileNameContext) {}

// ExitFileName is called when production fileName is exited.
func (s *BaseImgListener) ExitFileName(ctx *FileNameContext) {}

// EnterImageType is called when production imageType is entered.
func (s *BaseImgListener) EnterImageType(ctx *ImageTypeContext) {}

// ExitImageType is called when production imageType is exited.
func (s *BaseImgListener) ExitImageType(ctx *ImageTypeContext) {}

// EnterFolderName is called when production folderName is entered.
func (s *BaseImgListener) EnterFolderName(ctx *FolderNameContext) {}

// ExitFolderName is called when production folderName is exited.
func (s *BaseImgListener) ExitFolderName(ctx *FolderNameContext) {}

// EnterAction_ is called when production action_ is entered.
func (s *BaseImgListener) EnterAction_(ctx *Action_Context) {}

// ExitAction_ is called when production action_ is exited.
func (s *BaseImgListener) ExitAction_(ctx *Action_Context) {}

// EnterActionType is called when production actionType is entered.
func (s *BaseImgListener) EnterActionType(ctx *ActionTypeContext) {}

// ExitActionType is called when production actionType is exited.
func (s *BaseImgListener) ExitActionType(ctx *ActionTypeContext) {}

// EnterCrop is called when production crop is entered.
func (s *BaseImgListener) EnterCrop(ctx *CropContext) {}

// ExitCrop is called when production crop is exited.
func (s *BaseImgListener) ExitCrop(ctx *CropContext) {}

// EnterRotate is called when production rotate is entered.
func (s *BaseImgListener) EnterRotate(ctx *RotateContext) {}

// ExitRotate is called when production rotate is exited.
func (s *BaseImgListener) ExitRotate(ctx *RotateContext) {}

// EnterFlipX is called when production flipX is entered.
func (s *BaseImgListener) EnterFlipX(ctx *FlipXContext) {}

// ExitFlipX is called when production flipX is exited.
func (s *BaseImgListener) ExitFlipX(ctx *FlipXContext) {}

// EnterFlipY is called when production flipY is entered.
func (s *BaseImgListener) EnterFlipY(ctx *FlipYContext) {}

// ExitFlipY is called when production flipY is exited.
func (s *BaseImgListener) ExitFlipY(ctx *FlipYContext) {}

// EnterResize is called when production resize is entered.
func (s *BaseImgListener) EnterResize(ctx *ResizeContext) {}

// ExitResize is called when production resize is exited.
func (s *BaseImgListener) ExitResize(ctx *ResizeContext) {}

// EnterBrightness is called when production brightness is entered.
func (s *BaseImgListener) EnterBrightness(ctx *BrightnessContext) {}

// ExitBrightness is called when production brightness is exited.
func (s *BaseImgListener) ExitBrightness(ctx *BrightnessContext) {}

// EnterContrast is called when production contrast is entered.
func (s *BaseImgListener) EnterContrast(ctx *ContrastContext) {}

// ExitContrast is called when production contrast is exited.
func (s *BaseImgListener) ExitContrast(ctx *ContrastContext) {}

// EnterSaturation is called when production saturation is entered.
func (s *BaseImgListener) EnterSaturation(ctx *SaturationContext) {}

// ExitSaturation is called when production saturation is exited.
func (s *BaseImgListener) ExitSaturation(ctx *SaturationContext) {}

// EnterPixelate is called when production pixelate is entered.
func (s *BaseImgListener) EnterPixelate(ctx *PixelateContext) {}

// ExitPixelate is called when production pixelate is exited.
func (s *BaseImgListener) ExitPixelate(ctx *PixelateContext) {}

// EnterExport is called when production export is entered.
func (s *BaseImgListener) EnterExport(ctx *ExportContext) {}

// ExitExport is called when production export is exited.
func (s *BaseImgListener) ExitExport(ctx *ExportContext) {}

// EnterText is called when production text is entered.
func (s *BaseImgListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BaseImgListener) ExitText(ctx *TextContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseImgListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseImgListener) ExitComment(ctx *CommentContext) {}
