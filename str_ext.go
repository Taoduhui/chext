package exts

import (
	"strings"
	"unicode"
)

type String string

func NewStr(str string) *String {
	s := String(str)
	return &s
}
func (s *String) Tostring() string {
	return string(*s)
}
func (s *String) Clone() *String {
	return NewStr(strings.Clone(s.Tostring()))
}
func (s *String) Compare(b string) int {
	return strings.Compare(s.Tostring(), b)
}
func (s *String) Contains(substr string) bool {
	return strings.Contains(s.Tostring(), substr)
}
func (s *String) ContainsAny(chars string) bool {
	return strings.ContainsAny(s.Tostring(), chars)
}
func (s *String) ContainsRune(r rune) bool {
	return strings.ContainsRune(s.Tostring(), r)
}
func (s *String) Count(substr string) int {
	return strings.Count(s.Tostring(), substr)
}
func (s *String) Cut(sep string) (before, after string, found bool) {
	return strings.Cut(s.Tostring(), sep)
}
func (s *String) EqualFold(t string) bool {
	return strings.EqualFold(s.Tostring(), t)
}
func (s *String) Fields() []string {
	return strings.Fields(s.Tostring())
}
func (s *String) FieldsFunc(f func(rune) bool) []string {
	return strings.FieldsFunc(s.Tostring(), f)
}
func (s *String) HasPrefix(prefix string) bool {
	return strings.HasPrefix(s.Tostring(), prefix)
}
func (s *String) HasSuffix(suffix string) bool {
	return strings.HasSuffix(s.Tostring(), suffix)
}
func (s *String) Index(substr string) int {
	return strings.Index(s.Tostring(), substr)
}
func (s *String) IndexAny(chars string) int {
	return strings.IndexAny(s.Tostring(), chars)
}
func (s *String) IndexByte(c byte) int {
	return strings.IndexByte(s.Tostring(), c)
}
func (s *String) IndexFunc(f func(rune) bool) int {
	return strings.IndexFunc(s.Tostring(), f)
}
func (s *String) IndexRune(r rune) int {
	return strings.IndexRune(s.Tostring(), r)
}
func (s *String) Join(elems []string, sep string) *String {
	return NewStr(strings.Join(elems, sep))
}
func (s *String) LastIndex(substr string) int {
	return strings.LastIndex(s.Tostring(), substr)
}
func (s *String) LastIndexAny(chars string) int {
	return strings.LastIndexAny(s.Tostring(), chars)
}
func (s *String) LastIndexByte(c byte) int {
	return strings.LastIndexByte(s.Tostring(), c)
}
func (s *String) LastIndexFunc(f func(rune) bool) int {
	return strings.LastIndexFunc(s.Tostring(), f)
}
func (s *String) Map(mapping func(rune) rune) *String {
	return NewStr(strings.Map(mapping, s.Tostring()))
}
func (s *String) Repeat(count int) *String {
	return NewStr(strings.Repeat(s.Tostring(), count))
}
func (s *String) Replace(old, new string, n int) *String {
	return NewStr(strings.Replace(s.Tostring(), old, new, n))
}
func (s *String) ReplaceAll(old, new string) *String {
	return NewStr(strings.ReplaceAll(s.Tostring(), old, new))
}
func (s *String) Split(sep string) []string {
	return strings.Split(s.Tostring(), sep)
}
func (s *String) SplitAfter(sep string) []string {
	return strings.SplitAfter(s.Tostring(), sep)
}
func (s *String) SplitAfterN(sep string, n int) []string {
	return strings.SplitAfterN(s.Tostring(), sep, n)
}
func (s *String) SplitN(sep string, n int) []string {
	return strings.SplitN(s.Tostring(), sep, n)
}
func (s *String) ToLower() *String {
	return NewStr(strings.ToLower(s.Tostring()))
}
func (s *String) ToLowerSpecial(c unicode.SpecialCase) *String {
	return NewStr(strings.ToLowerSpecial(c, s.Tostring()))
}
func (s *String) ToTitle() *String {
	return NewStr(strings.ToTitle(s.Tostring()))
}
func (s *String) ToTitleSpecial(c unicode.SpecialCase) *String {
	return NewStr(strings.ToTitleSpecial(c, s.Tostring()))
}
func (s *String) ToUpper() *String {
	return NewStr(strings.ToUpper(s.Tostring()))
}
func (s *String) ToUpperSpecial(c unicode.SpecialCase) *String {
	return NewStr(strings.ToUpperSpecial(c, s.Tostring()))
}
func (s *String) ToValidUTF8(replacement string) *String {
	return NewStr(strings.ToValidUTF8(s.Tostring(), replacement))
}
func (s *String) Trim(cutset string) *String {
	return NewStr(strings.Trim(s.Tostring(), cutset))
}
func (s *String) TrimFunc(f func(rune) bool) *String {
	return NewStr(strings.TrimFunc(s.Tostring(), f))
}
func (s *String) TrimLeft(cutset string) *String {
	return NewStr(strings.TrimLeft(s.Tostring(), cutset))
}
func (s *String) TrimLeftFunc(f func(rune) bool) *String {
	return NewStr(strings.TrimLeftFunc(s.Tostring(), f))
}
func (s *String) TrimPrefix(prefix string) *String {
	return NewStr(strings.TrimPrefix(s.Tostring(), prefix))
}
func (s *String) TrimRight(cutset string) *String {
	return NewStr(strings.TrimRight(s.Tostring(), cutset))
}
func (s *String) TrimRightFunc(f func(rune) bool) *String {
	return NewStr(strings.TrimRightFunc(s.Tostring(), f))
}
func (s *String) TrimSpace() *String {
	return NewStr(strings.TrimSpace(s.Tostring()))
}
func (s *String) TrimSuffix(suffix string) *String {
	return NewStr(strings.TrimSuffix(s.Tostring(), suffix))
}
