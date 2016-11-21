package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"parser3", }, `{"subLanguage":"xml","relevance":0,"contains":[{"className":"comment","begin":"^#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"\\^rem{","end":"}","contains":[{"className":"comment","begin":"{","end":"}","contains":["self",{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10},{"className":"meta","begin":"^@(?:BASE|USE|CLASS|OPTIONS)$","relevance":10},{"className":"title","begin":"@[\\w\\-]+\\[[\\w^;\\-]*\\](?:\\[[\\w^;\\-]*\\])?(?:.*)$"},{"className":"variable","begin":"\\$\\{?[\\w\\-\\.\\:]+\\}?"},{"className":"keyword","begin":"\\^[\\w\\-\\.\\:]+"},{"className":"number","begin":"\\^#[0-9a-fA-F]+"},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0}]}`)
}