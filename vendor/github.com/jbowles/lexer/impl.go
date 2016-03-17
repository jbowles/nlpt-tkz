package lexer

import (
	"bytes"
	"unicode/utf8"
)

// IndexRune returns the index of the first occurrence in runes of the given rune.
// It returns -1 if rune is not present in runes.
func LexIndexRune(runes []rune, r rune) int {
	for i, c := range runes {
		if c == r {
			return i
		}
	}
	return -1
}

// Lexer::NextToken - Returns the next token from the reader.
func (l *lex) NextToken() *Token {
	for {
		select {
		case token := <-l.tokens:
			return token
		default:
			l.state = l.state(l)
		}
	}
	panic("not reached")
}

// Lexer::NewLine
func (l *lex) NewLine() {
	l.line++
	l.column = 0
}

// Lexer::Line
func (l *lex) Line() int {
	return l.line
}

// Lexer::Column
func (l *lex) Column() int {
	return l.column
}

// Lexer:PeekRune
func (l *lex) PeekRune(n int) rune {
	ok := l.ensureRuneLen(l.pos + n + 1) // Correct for 0-based 'n'

	if !ok {
		return RuneEOF
	}

	i := l.runes.Peek(l.pos + n)

	return i.(rune)
}

// Lexer::NextRune
func (l *lex) NextRune() rune {
	ok := l.ensureRuneLen(l.pos + 1)

	if !ok {
		return RuneEOF
	}

	i := l.runes.Peek(l.pos) // 0-based

	r := i.(rune)

	l.pos++

	l.tokenLen += utf8.RuneLen(r)

	l.column += utf8.RuneLen(r)

	return r
}

// Lexer::BackupRune
func (l *lex) BackupRune() {
	l.BackupRunes(1)
}

// Lexer::BackupRunes
func (l *lex) BackupRunes(n int) {
	for ; n > 0; n-- {
		if l.pos > 0 {
			l.pos--

			i := l.runes.Peek(l.pos) // 0-based
			r := i.(rune)

			l.tokenLen -= utf8.RuneLen(r)

			l.column -= utf8.RuneLen(r)
		} else {
			panic("Underflow Exception")
		}
	}
}

// Lexer::EmitToken
func (l *lex) EmitToken(t TokenType) {
	l.emit(t, false)
}

// Lexer::EmitTokenWithBytes
func (l *lex) EmitTokenWithBytes(t TokenType) {
	l.emit(t, true)
}

// Lexer::EmitToken
func (l *lex) EmitEOF() {
	l.emit(TokenTypeEOF, false)
}

// Lexer::IgnoreToken
func (l *lex) IgnoreToken() {
	l.consume(false)
}

// Lexer::Marker
func (l *lex) Marker() *Marker {
	return &Marker{sequence: l.sequence, pos: l.pos, tokenLen: l.tokenLen, line: l.line, column: l.column}
}

// Lexer::CanReset
func (l *lex) CanReset(m *Marker) bool {
	return m.sequence == l.sequence && m.pos <= l.runes.Len() && m.tokenLen <= l.peekPos
}

// Lexer::Reset
func (l *lex) Reset(m *Marker) {
	if l.CanReset(m) == false {
		panic("Invalid marker")
	}
	l.pos = m.pos

	l.tokenLen = m.tokenLen

	l.line = m.line

	l.column = m.column
}

// Lexer::MatchZeroOrOneBytes
func (l *lex) MatchZeroOrOneBytes(match []byte) bool {
	if r := l.PeekRune(0); r != RuneEOF && bytes.IndexRune(match, r) >= 0 {
		l.NextRune()
	}
	return true
}

// Lexer::MatchZeroOrOneRunes
func (l *lex) MatchZeroOrOneRunes(match []rune) bool {
	if r := l.PeekRune(0); r != RuneEOF && LexIndexRune(match, r) >= 0 {
		l.NextRune()
	}
	return true
}

// Lexer::MatchZeroOrOneRune
func (l *lex) MatchZeroOrOneRune(match rune) bool {
	if r := l.PeekRune(0); r != RuneEOF && r == match {
		l.NextRune()
	}
	return true
}

// Lexer::MatchZeroOrOneFunc
func (l *lex) MatchZeroOrOneFunc(match MatchFn) bool {
	if r := l.PeekRune(0); r != RuneEOF && match(r) {
		l.NextRune()
	}
	return true
}

// Lexer::MatchZeroOrMoreBytes
func (l *lex) MatchZeroOrMoreBytes(match []byte) bool {
	for r := l.PeekRune(0); r != RuneEOF && bytes.IndexRune(match, r) >= 0; r = l.PeekRune(0) {
		l.NextRune()
	}
	return true
}

// Lexer::MatchZeroOrMoreRunes
func (l *lex) MatchZeroOrMoreRunes(match []rune) bool {
	for r := l.PeekRune(0); r != RuneEOF && LexIndexRune(match, r) >= 0; r = l.PeekRune(0) {
		l.NextRune()
	}
	return true
}

// Lexer::MatchZeroOrMoreFunc
func (l *lex) MatchZeroOrMoreFunc(match MatchFn) bool {
	for r := l.PeekRune(0); r != RuneEOF && match(r); r = l.PeekRune(0) {
		l.NextRune()
	}
	return true
}

// Lexer::MatchOneBytes
func (l *lex) MatchOneBytes(match []byte) bool {
	if r := l.PeekRune(0); r != RuneEOF && bytes.IndexRune(match, r) >= 0 {
		l.NextRune()
		return true
	}
	return false
}

// Lexer::MatchOneRunes
func (l *lex) MatchOneRunes(match []rune) bool {
	if r := l.PeekRune(0); r != RuneEOF && LexIndexRune(match, r) >= 0 {
		l.NextRune()
		return true
	}
	return false
}

// Lexer::MatchOneRune
func (l *lex) MatchOneRune(match rune) bool {
	if r := l.PeekRune(0); r != RuneEOF && r == match {
		l.NextRune()
		return true
	}
	return false
}

// Lexer::MatchOneFunc
func (l *lex) MatchOneFunc(match MatchFn) bool {
	if r := l.PeekRune(0); r != RuneEOF && match(r) {
		l.NextRune()
		return true
	}
	return false
}

// Lexer::MatchOneOrMoreBytes
func (l *lex) MatchOneOrMoreBytes(match []byte) bool {
	var r rune
	if r = l.PeekRune(0); r != RuneEOF && bytes.IndexRune(match, r) >= 0 {
		l.NextRune()
		for r = l.PeekRune(0); r != RuneEOF && bytes.IndexRune(match, r) >= 0; r = l.PeekRune(0) {
			l.NextRune()
		}
		return true
	}
	return false
}

// Lexer::MatchOneOrMoreRunes
func (l *lex) MatchOneOrMoreRunes(match []rune) bool {
	var r rune
	if r = l.PeekRune(0); r != RuneEOF && LexIndexRune(match, r) >= 0 {
		l.NextRune()
		for r = l.PeekRune(0); r != RuneEOF && LexIndexRune(match, r) >= 0; r = l.PeekRune(0) {
			l.NextRune()
		}
		return true
	}
	return false
}

// Lexer::MatchOneOrMoreFunc
func (l *lex) MatchOneOrMoreFunc(match MatchFn) bool {
	var r rune
	if r = l.PeekRune(0); r != RuneEOF && match(r) {
		l.NextRune()
		for r = l.PeekRune(0); r != RuneEOF && match(r); r = l.PeekRune(0) {
			l.NextRune()
		}
		return true
	}
	return false
}

// Lexer::MatchMinMaxBytes
func (l *lex) MatchMinMaxBytes(match []byte, min int, max int) bool {
	marker := l.Marker()
	count := 0
	for r := l.PeekRune(0); r != RuneEOF && bytes.IndexRune(match, r) >= 0; r = l.PeekRune(0) {
		l.NextRune()
		count++
		if max > 0 && count >= max { // Check here to avoid unused PeekRune()
			break
		}
	}
	if count < min {
		l.Reset(marker)
		return false
	}
	return true
}

// Lexer::MatchMinMaxRunes
func (l *lex) MatchMinMaxRunes(match []rune, min int, max int) bool {
	marker := l.Marker()
	count := 0
	for r := l.PeekRune(0); r != RuneEOF && LexIndexRune(match, r) >= 0; r = l.PeekRune(0) {
		l.NextRune()
		count++
		if max > 0 && count >= max { // Check here to avoid unused PeekRune()
			break
		}
	}
	if count < min {
		l.Reset(marker)
		return false
	}
	return true
}

// Lexer::MatchMinMaxFunc
func (l *lex) MatchMinMaxFunc(match MatchFn, min int, max int) bool {
	marker := l.Marker()
	count := 0
	for r := l.PeekRune(0); r != RuneEOF && match(r); r = l.PeekRune(0) {
		l.NextRune()
		count++
		if max > 0 && count >= max { // Check here to avoid unused PeekRune()
			break
		}
	}
	if count < min {
		l.Reset(marker)
		return false
	}
	return true
}

// Lexer::NonMatchOneBytes
func (l *lex) NonMatchOneBytes(match []byte) bool {
	if r := l.PeekRune(0); r != RuneEOF && bytes.IndexRune(match, r) == -1 {
		l.NextRune()
		return true
	}
	return false
}

// Lexer::NonMatchOneRunes
func (l *lex) NonMatchOneRunes(match []rune) bool {
	if r := l.PeekRune(0); r != RuneEOF && LexIndexRune(match, r) == -1 {
		l.NextRune()
		return true
	}
	return false
}

// Lexer::NonMatchOneFunc
func (l *lex) NonMatchOneFunc(match MatchFn) bool {
	if r := l.PeekRune(0); r != RuneEOF && match(r) == false {
		l.NextRune()
		return true
	}
	return false
}

// Lexer::NonMatchOneOrMoreBytes
func (l *lex) NonMatchOneOrMoreBytes(match []byte) bool {
	var r rune
	if r = l.PeekRune(0); r != RuneEOF && bytes.IndexRune(match, r) == -1 {
		l.NextRune()
		for r = l.PeekRune(0); r != RuneEOF && bytes.IndexRune(match, r) == -1; r = l.PeekRune(0) {
			l.NextRune()
		}
		return true
	}
	return false
}

// Lexer::NonMatchOneOrMoreRunes
func (l *lex) NonMatchOneOrMoreRunes(match []rune) bool {
	var r rune
	if r = l.PeekRune(0); r != RuneEOF && LexIndexRune(match, r) == -1 {
		l.NextRune()
		for r = l.PeekRune(0); r != RuneEOF && LexIndexRune(match, r) == -1; r = l.PeekRune(0) {
			l.NextRune()
		}
		return true
	}
	return false
}

// Lexer::NonMatchOneOrMoreFunc
func (l *lex) NonMatchOneOrMoreFunc(match MatchFn) bool {
	var r rune
	if r = l.PeekRune(0); r != RuneEOF && match(r) == false {
		l.NextRune()
		for r = l.PeekRune(0); r != RuneEOF && match(r) == false; r = l.PeekRune(0) {
			l.NextRune()
		}
		return true
	}
	return false
}

// Lexer::NonMatchZeroOrOneBytes
func (l *lex) NonMatchZeroOrOneBytes(match []byte) bool {
	if r := l.PeekRune(0); r != RuneEOF && bytes.IndexRune(match, r) == -1 {
		l.NextRune()
	}
	return true
}

// Lexer::NonMatchZeroOrOneRunes
func (l *lex) NonMatchZeroOrOneRunes(match []rune) bool {
	if r := l.PeekRune(0); r != RuneEOF && LexIndexRune(match, r) == -1 {
		l.NextRune()
	}
	return true
}

// Lexer::NonMatchZeroOrOneFunc
func (l *lex) NonMatchZeroOrOneFunc(match MatchFn) bool {
	if r := l.PeekRune(0); r != RuneEOF && match(r) == false {
		l.NextRune()
	}
	return true
}

// Lexer::NonMatchZeroOrMoreBytes
func (l *lex) NonMatchZeroOrMoreBytes(match []byte) bool {
	for r := l.PeekRune(0); r != RuneEOF && bytes.IndexRune(match, r) == -1; r = l.PeekRune(0) {
		l.NextRune()
	}
	return true
}

// Lexer::NonMatchZeroOrMoreRunes
func (l *lex) NonMatchZeroOrMoreRunes(match []rune) bool {
	for r := l.PeekRune(0); r != RuneEOF && LexIndexRune(match, r) == -1; r = l.PeekRune(0) {
		l.NextRune()
	}
	return true
}

// Lexer::NonMatchZeroOrMoreFunc
func (l *lex) NonMatchZeroOrMoreFunc(match MatchFn) bool {
	for r := l.PeekRune(0); r != RuneEOF && match(r) == false; r = l.PeekRune(0) {
		l.NextRune()
	}
	return true
}

// Lexer::MatchEOF
func (l *lex) MatchEOF() bool {
	if r := l.PeekRune(0); r == RuneEOF {
		l.NextRune()
		return true
	}
	return false
}
