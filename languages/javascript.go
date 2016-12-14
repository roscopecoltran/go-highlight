package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"javascript", "js", "jsx"}, `{"aliases":["js","jsx"],"keywords":{"keyword":"in of if for while finally var new function do return void else break catch instanceof with throw case default try this switch continue typeof delete let yield const export super debugger as async await static import from as","literal":"true false null undefined NaN Infinity","built_in":"eval isFinite isNaN parseFloat parseInt decodeURI decodeURIComponent encodeURI encodeURIComponent escape unescape Object Function Boolean Error EvalError InternalError RangeError ReferenceError StopIteration SyntaxError TypeError URIError Number Math Date String RegExp Array Float32Array Float64Array Int16Array Int32Array Int8Array Uint16Array Uint32Array Uint8Array Uint8ClampedArray ArrayBuffer DataView JSON Intl arguments require module console window document Symbol Set Map WeakSet WeakMap Proxy Reflect Promise"},"contains":[{"className":"meta","relevance":10,"begin":"^\\s*['\"]use (strict|asm)['\"]"},{"className":"meta","begin":"^#!","end":"$"},{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"string","begin":"`+"`"+`","end":"`+"`"+`","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\$\\{","end":"\\}","keywords":{"keyword":"in of if for while finally var new function do return void else break catch instanceof with throw case default try this switch continue typeof delete let yield const export super debugger as async await static import from as","literal":"true false null undefined NaN Infinity","built_in":"eval isFinite isNaN parseFloat parseInt decodeURI decodeURIComponent encodeURI encodeURIComponent escape unescape Object Function Boolean Error EvalError InternalError RangeError ReferenceError StopIteration SyntaxError TypeError URIError Number Math Date String RegExp Array Float32Array Float64Array Int16Array Int32Array Int8Array Uint16Array Uint32Array Uint8Array Uint8ClampedArray ArrayBuffer DataView JSON Intl arguments require module console window document Symbol Set Map WeakSet WeakMap Proxy Reflect Promise"},"contains":[{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"Ref":["contains","4"]},{"className":"number","variants":[{"begin":"\\b(0[bB][01]+)"},{"begin":"\\b(0[oO][0-7]+)"},{"begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)"}],"relevance":0},{"className":"regexp","begin":"\\/","end":"\\/[gimuy]*","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"begin":"\\[","end":"\\]","relevance":0,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]}]}]}]},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"number","variants":[{"begin":"\\b(0[bB][01]+)"},{"begin":"\\b(0[oO][0-7]+)"},{"begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)"}],"relevance":0},{"begin":"[{,]\\s*","relevance":0,"contains":[{"begin":"[A-Za-z$_][0-9A-Za-z$_]*\\s*:","returnBegin":true,"relevance":0,"contains":[{"className":"attr","begin":"[A-Za-z$_][0-9A-Za-z$_]*","relevance":0}]}]},{"begin":"(!|!=|!==|%|%=|&|&&|&=|\\*|\\*=|\\+|\\+=|,|-|-=|/=|/|:|;|<<|<<=|<=|<|===|==|=|>>>=|>>=|>=|>>>|>>|>|\\?|\\[|\\{|\\(|\\^|\\^=|\\||\\|=|\\|\\||~|\\b(case|return|throw)\\b)\\s*","keywords":"return throw case","contains":[{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"regexp","begin":"\\/","end":"\\/[gimuy]*","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"begin":"\\[","end":"\\]","relevance":0,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]}]},{"className":"function","begin":"(\\(.*?\\)|[A-Za-z$_][0-9A-Za-z$_]*)\\s*=>","returnBegin":true,"end":"\\s*=>","contains":[{"className":"params","variants":[{"begin":"[A-Za-z$_][0-9A-Za-z$_]*"},{"begin":"\\(\\s*\\)"},{"begin":"\\(","end":"\\)","excludeBegin":true,"excludeEnd":true,"keywords":{"keyword":"in of if for while finally var new function do return void else break catch instanceof with throw case default try this switch continue typeof delete let yield const export super debugger as async await static import from as","literal":"true false null undefined NaN Infinity","built_in":"eval isFinite isNaN parseFloat parseInt decodeURI decodeURIComponent encodeURI encodeURIComponent escape unescape Object Function Boolean Error EvalError InternalError RangeError ReferenceError StopIteration SyntaxError TypeError URIError Number Math Date String RegExp Array Float32Array Float64Array Int16Array Int32Array Int8Array Uint16Array Uint32Array Uint8Array Uint8ClampedArray ArrayBuffer DataView JSON Intl arguments require module console window document Symbol Set Map WeakSet WeakMap Proxy Reflect Promise"},"contains":[{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"string","begin":"`+"`"+`","end":"`+"`"+`","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\$\\{","end":"\\}","keywords":{"keyword":"in of if for while finally var new function do return void else break catch instanceof with throw case default try this switch continue typeof delete let yield const export super debugger as async await static import from as","literal":"true false null undefined NaN Infinity","built_in":"eval isFinite isNaN parseFloat parseInt decodeURI decodeURIComponent encodeURI encodeURIComponent escape unescape Object Function Boolean Error EvalError InternalError RangeError ReferenceError StopIteration SyntaxError TypeError URIError Number Math Date String RegExp Array Float32Array Float64Array Int16Array Int32Array Int8Array Uint16Array Uint32Array Uint8Array Uint8ClampedArray ArrayBuffer DataView JSON Intl arguments require module console window document Symbol Set Map WeakSet WeakMap Proxy Reflect Promise"},"contains":[{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"Ref":["contains","4"]},{"className":"number","variants":[{"begin":"\\b(0[bB][01]+)"},{"begin":"\\b(0[oO][0-7]+)"},{"begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)"}],"relevance":0},{"className":"regexp","begin":"\\/","end":"\\/[gimuy]*","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"begin":"\\[","end":"\\]","relevance":0,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]}]}]}]},{"className":"number","variants":[{"begin":"\\b(0[bB][01]+)"},{"begin":"\\b(0[oO][0-7]+)"},{"begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)"}],"relevance":0},{"className":"regexp","begin":"\\/","end":"\\/[gimuy]*","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"begin":"\\[","end":"\\]","relevance":0,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}]}]},{"begin":"<","end":"(\\/\\w+|\\w+\\/)>","subLanguage":["xml"],"contains":[{"begin":"<\\w+\\s*\\/>","skip":true},{"begin":"<\\w+","end":"(\\/\\w+|\\w+\\/)>","skip":true,"contains":[{"begin":"<\\w+\\s*\\/>","skip":true},{"Ref":["contains","9","contains","4","contains","1"]}]}]}],"relevance":0},{"className":"function","beginKeywords":"function","end":"\\{","excludeEnd":true,"contains":[{"className":"title","begin":"[A-Za-z$_][0-9A-Za-z$_]*","relevance":0},{"className":"params","begin":"\\(","end":"\\)","excludeBegin":true,"excludeEnd":true,"contains":[{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"string","begin":"`+"`"+`","end":"`+"`"+`","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\$\\{","end":"\\}","keywords":{"keyword":"in of if for while finally var new function do return void else break catch instanceof with throw case default try this switch continue typeof delete let yield const export super debugger as async await static import from as","literal":"true false null undefined NaN Infinity","built_in":"eval isFinite isNaN parseFloat parseInt decodeURI decodeURIComponent encodeURI encodeURIComponent escape unescape Object Function Boolean Error EvalError InternalError RangeError ReferenceError StopIteration SyntaxError TypeError URIError Number Math Date String RegExp Array Float32Array Float64Array Int16Array Int32Array Int8Array Uint16Array Uint32Array Uint8Array Uint8ClampedArray ArrayBuffer DataView JSON Intl arguments require module console window document Symbol Set Map WeakSet WeakMap Proxy Reflect Promise"},"contains":[{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"Ref":["contains","4"]},{"className":"number","variants":[{"begin":"\\b(0[bB][01]+)"},{"begin":"\\b(0[oO][0-7]+)"},{"begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)"}],"relevance":0},{"className":"regexp","begin":"\\/","end":"\\/[gimuy]*","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"begin":"\\[","end":"\\]","relevance":0,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]}]}]}]},{"className":"number","variants":[{"begin":"\\b(0[bB][01]+)"},{"begin":"\\b(0[oO][0-7]+)"},{"begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)"}],"relevance":0},{"className":"regexp","begin":"\\/","end":"\\/[gimuy]*","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"begin":"\\[","end":"\\]","relevance":0,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}],"illegal":"\\[|%"},{"begin":"\\$[(.]"},{"begin":"\\.\\s*[a-zA-Z_]\\w*","relevance":0},{"className":"class","beginKeywords":"class","end":"[{;=]","excludeEnd":true,"illegal":"[:\"\\[\\]]","contains":[{"beginKeywords":"extends"},{"className":"title","begin":"[a-zA-Z_]\\w*","relevance":0}]},{"beginKeywords":"constructor","end":"\\{","excludeEnd":true}],"illegal":"#(?!!)"}`)
}