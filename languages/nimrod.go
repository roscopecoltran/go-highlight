package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"nimrod", "nim"}, `{"aliases":["nim"],"keywords":{"keyword":"addr and as asm bind block break case cast const continue converter discard distinct div do elif else end enum except export finally for from generic if import in include interface is isnot iterator let macro method mixin mod nil not notin object of or out proc ptr raise ref return shl shr static template try tuple type using var when while with without xor yield","literal":"shared guarded stdin stdout stderr result true false","built_in":"int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 float float32 float64 bool char string cstring pointer expr stmt void auto any range array openarray varargs seq set clong culong cchar cschar cshort cint csize clonglong cfloat cdouble clongdouble cuchar cushort cuint culonglong cstringarray semistatic"},"contains":[{"className":"meta","begin":"{\\.","end":"\\.}","relevance":10},{"className":"string","begin":"[a-zA-Z]\\w*\"","end":"\"","contains":[{"begin":"\"\""}]},{"className":"string","begin":"([a-zA-Z]\\w*)?\"\"\"","end":"\"\"\""},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"type","begin":"\\b[A-Z]\\w+\\b","relevance":0},{"className":"number","relevance":0,"variants":[{"begin":"\\b(0[xX][0-9a-fA-F][_0-9a-fA-F]*)('?[iIuU](8|16|32|64))?"},{"begin":"\\b(0o[0-7][_0-7]*)('?[iIuUfF](8|16|32|64))?"},{"begin":"\\b(0(b|B)[01][_01]*)('?[iIuUfF](8|16|32|64))?"},{"begin":"\\b(\\d[_\\d]*)('?[iIuUfF](8|16|32|64))?"}]},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}`)
}