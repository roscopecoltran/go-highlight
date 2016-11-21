package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register([]string{"haskell", "hs"}, `{"aliases":["hs"],"keywords":"let in if then else case of where do module import hiding qualified type data newtype deriving class instance as default infix infixl infixr foreign export ccall stdcall cplusplus jvm dotnet safe unsafe family forall mdo proc rec","contains":[{"beginKeywords":"module","end":"where","keywords":"module where","contains":[{"begin":"\\(","end":"\\)","illegal":"\"","contains":[{"className":"meta","begin":"{-#","end":"#-}"},{"className":"meta","begin":"^#","end":"$"},{"className":"type","begin":"\\b[A-Z][\\w]*(\\((\\.\\.|,|\\w+)\\))?"},{"className":"title","begin":"[_a-z][\\w']*","relevance":0},{"variants":[{"className":"comment","begin":"--","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"{-","end":"-}","contains":["self",{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}]},{"variants":[{"className":"comment","begin":"--","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"{-","end":"-}","contains":["self",{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}],"illegal":"\\W\\.|;"},{"begin":"\\bimport\\b","end":"$","keywords":"import qualified as hiding","contains":[{"begin":"\\(","end":"\\)","illegal":"\"","contains":[{"className":"meta","begin":"{-#","end":"#-}"},{"className":"meta","begin":"^#","end":"$"},{"className":"type","begin":"\\b[A-Z][\\w]*(\\((\\.\\.|,|\\w+)\\))?"},{"className":"title","begin":"[_a-z][\\w']*","relevance":0},{"variants":[{"className":"comment","begin":"--","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"{-","end":"-}","contains":["self",{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}]},{"variants":[{"className":"comment","begin":"--","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"{-","end":"-}","contains":["self",{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}],"illegal":"\\W\\.|;"},{"className":"class","begin":"^(\\s*)?(class|instance)\\b","end":"where","keywords":"class family instance where","contains":[{"className":"type","begin":"\\b[A-Z][\\w']*","relevance":0},{"begin":"\\(","end":"\\)","illegal":"\"","contains":[{"className":"meta","begin":"{-#","end":"#-}"},{"className":"meta","begin":"^#","end":"$"},{"className":"type","begin":"\\b[A-Z][\\w]*(\\((\\.\\.|,|\\w+)\\))?"},{"className":"title","begin":"[_a-z][\\w']*","relevance":0},{"variants":[{"className":"comment","begin":"--","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"{-","end":"-}","contains":["self",{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}]},{"variants":[{"className":"comment","begin":"--","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"{-","end":"-}","contains":["self",{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}]},{"className":"class","begin":"\\b(data|(new)?type)\\b","end":"$","keywords":"data family type newtype deriving","contains":[{"className":"meta","begin":"{-#","end":"#-}"},{"className":"type","begin":"\\b[A-Z][\\w']*","relevance":0},{"begin":"\\(","end":"\\)","illegal":"\"","contains":[{"className":"meta","begin":"{-#","end":"#-}"},{"className":"meta","begin":"^#","end":"$"},{"className":"type","begin":"\\b[A-Z][\\w]*(\\((\\.\\.|,|\\w+)\\))?"},{"className":"title","begin":"[_a-z][\\w']*","relevance":0},{"variants":[{"className":"comment","begin":"--","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"{-","end":"-}","contains":["self",{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}]},{"begin":"{","end":"}","contains":[{"className":"meta","begin":"{-#","end":"#-}"},{"className":"meta","begin":"^#","end":"$"},{"className":"type","begin":"\\b[A-Z][\\w]*(\\((\\.\\.|,|\\w+)\\))?"},{"className":"title","begin":"[_a-z][\\w']*","relevance":0},{"variants":[{"className":"comment","begin":"--","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"{-","end":"-}","contains":["self",{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}]},{"variants":[{"className":"comment","begin":"--","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"{-","end":"-}","contains":["self",{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}]},{"beginKeywords":"default","end":"$","contains":[{"className":"type","begin":"\\b[A-Z][\\w']*","relevance":0},{"begin":"\\(","end":"\\)","illegal":"\"","contains":[{"className":"meta","begin":"{-#","end":"#-}"},{"className":"meta","begin":"^#","end":"$"},{"className":"type","begin":"\\b[A-Z][\\w]*(\\((\\.\\.|,|\\w+)\\))?"},{"className":"title","begin":"[_a-z][\\w']*","relevance":0},{"variants":[{"className":"comment","begin":"--","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"{-","end":"-}","contains":["self",{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}]},{"variants":[{"className":"comment","begin":"--","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"{-","end":"-}","contains":["self",{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}]},{"beginKeywords":"infix infixl infixr","end":"$","contains":[{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0},{"variants":[{"className":"comment","begin":"--","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"{-","end":"-}","contains":["self",{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}]},{"begin":"\\bforeign\\b","end":"$","keywords":"foreign import export ccall stdcall cplusplus jvm dotnet safe unsafe","contains":[{"className":"type","begin":"\\b[A-Z][\\w']*","relevance":0},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"variants":[{"className":"comment","begin":"--","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"{-","end":"-}","contains":["self",{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}]},{"className":"meta","begin":"#!\\/usr\\/bin\\/env runhaskell","end":"$"},{"className":"meta","begin":"{-#","end":"#-}"},{"className":"meta","begin":"^#","end":"$"},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0},{"className":"type","begin":"\\b[A-Z][\\w']*","relevance":0},{"className":"title","begin":"^[_a-z][\\w']*","relevance":0},{"variants":[{"className":"comment","begin":"--","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"{-","end":"-}","contains":["self",{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"begin":"->|<-"}]}`)
}