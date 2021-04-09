grammar Img;

// Tokens
ASSIGN: '=';
DOT: '.';
COMMA: ',';
APOS: '"';
SEMICOLON: ';';
OPEN_PARAN: '(';
CLOSED_PARAN: ')';
FOLDER: '[]';
OPEN: 'open';
NUMBER: [0-9]+;
ID: [a-zA-Z_]+;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start: declaration+ action_* export EOF;

// Declaration
declaration: openFile | openFolder;

openFile:
	ID ASSIGN OPEN OPEN_PARAN fileName CLOSED_PARAN SEMICOLON;
openFolder:
	ID FOLDER ASSIGN OPEN OPEN_PARAN folderName CLOSED_PARAN SEMICOLON;

fileName: APOS ID DOT imageType APOS;
imageType: 'png' | 'jpg';

folderName: APOS ID APOS;

// Action
action_: ID DOT actionType SEMICOLON;

actionType: crop | rotate | flipX | flipY | resize;

crop:
	'crop' OPEN_PARAN NUMBER COMMA NUMBER COMMA NUMBER COMMA NUMBER CLOSED_PARAN;
rotate: 'rotate' OPEN_PARAN NUMBER CLOSED_PARAN;
flipX: 'flipX' OPEN_PARAN CLOSED_PARAN;
flipY: 'flipY' OPEN_PARAN CLOSED_PARAN;
resize: 'resize' OPEN_PARAN NUMBER (COMMA NUMBER)? CLOSED_PARAN;

// Export
export:
	ID DOT 'save' OPEN_PARAN fileName CLOSED_PARAN SEMICOLON;