package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"qml", "qt"}, `{"aliases":["qt"],"case_insensitive":false,"keywords":{"keyword":"in of on if for while finally var new function do return void else break catch instanceof with throw case default try this switch continue typeof delete let yield const export super debugger as async await import","literal":"true false null undefined NaN Infinity","built_in":"eval isFinite isNaN parseFloat parseInt decodeURI decodeURIComponent encodeURI encodeURIComponent escape unescape Object Function Boolean Error EvalError InternalError RangeError ReferenceError StopIteration SyntaxError TypeError URIError Number Math Date String RegExp Array Float32Array Float64Array Int16Array Int32Array Int8Array Uint16Array Uint32Array Uint8Array Uint8ClampedArray ArrayBuffer DataView JSON Intl arguments require module console window document Symbol Set Map WeakSet WeakMap Proxy Reflect Behavior bool color coordinate date double enumeration font geocircle georectangle geoshape int list matrix4x4 parent point quaternion real rect size string url var variant vector2d vector3d vector4dPromise"},"contains":[{"className":"meta","begin":"^\\s*['\"]use (strict|asm)['\"]"},{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"string","begin":"`+"`"+`","end":"`+"`"+`","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\$\\{","end":"\\}"}]},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"number","variants":[{"begin":"\\b(0[bB][01]+)"},{"begin":"\\b(0[oO][0-7]+)"},{"begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)"}],"relevance":0},{"begin":"(!|!=|!==|%|%=|&|&&|&=|\\*|\\*=|\\+|\\+=|,|-|-=|/=|/|:|;|<<|<<=|<=|<|===|==|=|>>>=|>>=|>=|>>>|>>|>|\\?|\\[|\\{|\\(|\\^|\\^=|\\||\\|=|\\|\\||~|\\b(case|return|throw)\\b)\\s*","keywords":"return throw case","contains":[{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"regexp","begin":"\\/","end":"\\/[gimuy]*","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"begin":"\\[","end":"\\]","relevance":0,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]}]},{"begin":"<","end":">\\s*[);\\]]","relevance":0,"subLanguage":"xml"}],"relevance":0},{"className":"keyword","begin":"\\bsignal\\b","starts":{"className":"string","end":"(\\(|:|=|;|,|//|/\\*|$)","returnEnd":true}},{"className":"keyword","begin":"\\bproperty\\b","starts":{"className":"string","end":"(:|=|;|,|//|/\\*|$)","returnEnd":true}},{"className":"function","beginKeywords":"function","end":"\\{","excludeEnd":true,"contains":[{"className":"title","begin":"[A-Za-z$_][0-9A-Za-z$_]*","relevance":0},{"className":"params","begin":"\\(","end":"\\)","excludeBegin":true,"excludeEnd":true,"contains":[{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}],"illegal":"\\[|%"},{"begin":"\\.[a-zA-Z]\\w*","relevance":0},{"className":"attribute","begin":"\\bid\\s*:","starts":{"className":"string","end":"[a-zA-Z_][a-zA-Z0-9\\._]*","returnEnd":false}},{"begin":"[a-zA-Z_][a-zA-Z0-9\\._]*\\s*:","returnBegin":true,"contains":[{"className":"attribute","begin":"[a-zA-Z_][a-zA-Z0-9\\._]*","end":"\\s*:","excludeEnd":true,"relevance":0}],"relevance":0},{"begin":"[a-zA-Z_][a-zA-Z0-9\\._]*\\s*{","end":"{","returnBegin":true,"relevance":0,"contains":[{"className":"title","begin":"[a-zA-Z_][a-zA-Z0-9\\._]*","relevance":0}]}],"illegal":"#"}`)
}