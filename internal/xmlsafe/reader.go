package xmlsafe

import "io"

type Reader struct { R io.Reader }

func (r Reader) Read(p []byte) (int, error) {
 for {
  n, err := r.R.Read(p)
  j := 0
  for i := 0; i < n; i++ {
   b := p[i]
   invalid := b < 32 && b != 9 && b != 10 && b != 13
   if invalid { continue }
   p[j] = b
   j++
  }
  if j > 0 || err != nil { return j, err }
 }
}
