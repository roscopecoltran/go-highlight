package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"dart", }, `{"keywords":{"keyword":"assert async await break case catch class const continue default do else enum extends false final finally for if in is new null rethrow return super switch sync this throw true try var void while with yield abstract as dynamic export external factory get implements import library operator part set static typedef","built_in":"print Comparable DateTime Duration Function Iterable Iterator List Map Match Null Object Pattern RegExp Set Stopwatch String StringBuffer StringSink Symbol Type Uri bool double int num document window querySelector querySelectorAll Element ElementList"},"contains":[{"className":"string","variants":[{"begin":"r'''","end":"'''"},{"begin":"r\"\"\"","end":"\"\"\""},{"begin":"r'","end":"'","illegal":"\\n"},{"begin":"r\"","end":"\"","illegal":"\\n"},{"begin":"'''","end":"'''","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\$\\{","end":"}","keywords":"true false null this is new super","contains":[{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0},{"Ref":["contains","0"]}]}]},{"begin":"\"\"\"","end":"\"\"\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\$\\{","end":"}","keywords":"true false null this is new super","contains":[{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0},{"Ref":["contains","0"]}]}]},{"begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\$\\{","end":"}","keywords":"true false null this is new super","contains":[{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0},{"Ref":["contains","0"]}]}]},{"begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\$\\{","end":"}","keywords":"true false null this is new super","contains":[{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0},{"Ref":["contains","0"]}]}]}]},{"className":"comment","begin":"/\\*\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"subLanguage":["markdown"]},{"className":"comment","begin":"///","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"subLanguage":["markdown"]},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"class","beginKeywords":"class interface","end":"{","excludeEnd":true,"contains":[{"beginKeywords":"extends implements"},{"className":"title","begin":"[a-zA-Z_]\\w*","relevance":0}]},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0},{"className":"meta","begin":"@[A-Za-z]+"},{"begin":"=>"}]}`)
}