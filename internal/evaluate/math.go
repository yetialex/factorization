package evaluate

import (
	"context"
	"log"
	"math/big"
	"time"
)

// PrimeFactorizationBigInt возвращает слайс простых множителей
func PrimeFactorizationBigInt(ctx context.Context, sNumber string) (PrimeFactors, time.Duration, error) {
	now := time.Now()
	result := make(PrimeFactors, 0)
	n, ok := new(big.Int).SetString(sNumber, 10)
	if !ok {
		return result, time.Since(now), ErrConvert
	}

	zero := big.NewInt(int64(0))
	one := big.NewInt(int64(1))

	if n.Cmp(one) <= 0 {
		return result, time.Since(now), ErrWrongNumber
	}

	d := big.NewInt(int64(2))
	var temp = new(big.Int)

	for temp.Mul(d, d).Cmp(n) <= 0 {
		//проверим, не отменили ли функцию
		select {
		case <-ctx.Done():
			log.Println("cancelled")
			return nil, time.Since(now), ErrRequestCancelled
		default:
		}
		if mod := temp.Mod(n, d); mod.Cmp(zero) == 0 {
			result.Add(d)
			n.Div(n, d)
		} else {
			d.Add(d, one)
		}
	}
	if n.Cmp(one) > 0 {
		result.Add(n)
	}

	return result, time.Since(now), nil
}
