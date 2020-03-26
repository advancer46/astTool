package astTool
type whiteSpace byte

const (
	ignore   = whiteSpace(0)
	blank    = whiteSpace(' ')
	vtab     = whiteSpace('\v')
	newline  = whiteSpace('\n')
	formfeed = whiteSpace('\f')
	indent   = whiteSpace('>')
	unindent = whiteSpace('<')
	tab = whiteSpace('\t')
	lparen = whiteSpace('(')
	rparen = whiteSpace(')')
)

