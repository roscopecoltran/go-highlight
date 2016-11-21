package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"r", }, `{"contains":[{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"begin":"([a-zA-Z]|\\.[a-zA-Z.])[a-zA-Z0-9._]*","lexemes":"([a-zA-Z]|\\.[a-zA-Z.])[a-zA-Z0-9._]*","keywords":{"keyword":"function if in break next repeat else for return switch while try tryCatch stop warning require library attach detach source setMethod setGeneric setGroupGeneric setClass ...","literal":"NULL NA TRUE FALSE T F Inf NaN NA_integer_|10 NA_real_|10 NA_character_|10 NA_complex_|10"},"relevance":0},{"className":"number","begin":"0[xX][0-9a-fA-F]+[Li]?\\b","relevance":0},{"className":"number","begin":"\\d+(?:[eE][+\\-]?\\d*)?L\\b","relevance":0},{"className":"number","begin":"\\d+\\.(?!\\d)(?:i\\b)?","relevance":0},{"className":"number","begin":"\\d+(?:\\.\\d*)?(?:[eE][+\\-]?\\d*)?i?\\b","relevance":0},{"className":"number","begin":"\\.\\d+(?:[eE][+\\-]?\\d*)?i?\\b","relevance":0},{"begin":"`+"`"+`","end":"`+"`"+`","relevance":0},{"className":"string","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}],"variants":[{"begin":"\"","end":"\""},{"begin":"'","end":"'"}]}]}`)
}