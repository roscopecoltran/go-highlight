package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"htmlbars", }, `{"case_insensitive":true,"subLanguage":["xml"],"contains":[{"className":"comment","begin":"{{!(--)?","end":"(--)?}}","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"template-tag","begin":"\\{\\{[#\\/]","end":"\\}\\}","contains":[{"className":"name","begin":"[a-zA-Z\\.\\-]+","keywords":{"builtin-name":"action collection component concat debugger each each-in else get hash if input link-to loc log mut outlet partial query-params render textarea unbound unless with yield view"},"starts":{"endsWithParent":true,"relevance":0,"keywords":{"keyword":"as","built_in":"action collection component concat debugger each each-in else get hash if input link-to loc log mut outlet partial query-params render textarea unbound unless with yield view"},"contains":[{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"illegal":"\\}\\}","begin":"[a-zA-Z0-9_]+=","returnBegin":true,"relevance":0,"contains":[{"className":"attr","begin":"[a-zA-Z0-9_]+"}]},{"className":"number","begin":"\\b\\d+(\\.\\d+)?","relevance":0}]}}]},{"className":"template-variable","begin":"\\{\\{[a-zA-Z][a-zA-Z\\-]+","end":"\\}\\}","keywords":{"keyword":"as","built_in":"action collection component concat debugger each each-in else get hash if input link-to loc log mut outlet partial query-params render textarea unbound unless with yield view"},"contains":[{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]}]}]}`)
}