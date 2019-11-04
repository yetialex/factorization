package evaluate

import (
	"math/big"
	"strings"
)

type PrimeFactors []*big.Int

func (pf *PrimeFactors) Add(factor *big.Int) {
	*pf = append(*pf, new(big.Int).Set(factor))
}

// String() имплементирует интерфейс fmt.Stringer,
// используем для строкового представления слайса множителей в виде m1 * m2 * ... * mn
func (pf PrimeFactors) String() string {
	return pf.Join(" * ")
}

func (pf PrimeFactors) Join(sep string) string {
	switch len(pf) {
	case 0:
		return ""
	case 1:
		return pf[0].String()
	}
	n := len(sep) * (len(pf) - 1)
	for i := 0; i < len(pf); i++ {
		n += len(pf[i].String())
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(pf[0].String())
	for _, s := range pf[1:] {
		b.WriteString(sep)
		b.WriteString(s.String())
	}
	return b.String()
}
