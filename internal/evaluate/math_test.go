package evaluate

import (
	"context"
	"math/big"
	"strconv"
	"testing"
)

func TestTablePrimeFactorizationBigInt(t *testing.T) {
	tests := []struct {
		sNumber string
		factors []*big.Int
		err     error
	}{
		{"-10", []*big.Int{}, ErrWrongNumber},
		{"1", []*big.Int{}, ErrWrongNumber},
		{"1.5", []*big.Int{}, ErrConvert},
		{"1q", []*big.Int{}, ErrConvert},
		{"12", []*big.Int{big.NewInt(2), big.NewInt(2), big.NewInt(3)}, nil},
		{"18", []*big.Int{big.NewInt(2), big.NewInt(3), big.NewInt(3)}, nil},
		{"123", []*big.Int{big.NewInt(3), big.NewInt(41)}, nil},
		{"55555555555555", []*big.Int{big.NewInt(5), big.NewInt(11), big.NewInt(239), big.NewInt(4649), big.NewInt(909091)}, nil},
		{"123456787654321", []*big.Int{big.NewInt(11), big.NewInt(11), big.NewInt(73), big.NewInt(73), big.NewInt(101), big.NewInt(101), big.NewInt(137), big.NewInt(137)}, nil},
		{"11111111111111111111111111111", []*big.Int{big.NewInt(3191), big.NewInt(16763), big.NewInt(43037), big.NewInt(62003), big.NewInt(77843839397)}, nil},
	}

	ctx := context.Background()
	t.Log("Compare function results with expected values.")
	{
		for i, tt := range tests {
			t.Logf("\tTest: %d\tNumber: %s, prime factors: %v, expected error: %v", i, tt.sNumber, tt.factors, tt.err)
			{
				factors, _, err := PrimeFactorizationBigInt(ctx, tt.sNumber)
				if err != tt.err {
					t.Fatalf("PrimeFactorizationBigInt error: %v", err)
				}
				for i, factor := range factors {
					if factor.Cmp(tt.factors[i]) != 0 {
						t.Fatalf("factors are different")
					}
				}
			}
		}
	}

}

func TestRangePrimeFactorizationBigInt(t *testing.T) {
	ctx := context.Background()
	t.Log("Compare function results with expected values.")
	{
		for n := 2; n <= 10; n++ {
			factors, _, err := PrimeFactorizationBigInt(ctx, strconv.Itoa(n))
			bigN := big.NewInt(int64(n))
			if err != nil {
				t.Fatalf("PrimeFactorizationBigInt error: %v", err)
			}
			t.Logf("\tNumber: %v = %v", bigN, factors)
			//сравним произведение полученных множителей с исходным числом
			if bigN.Cmp(multiplySliceBigInt(factors)) != 0 {
				t.Fatalf("factorization is wrong for %v (%v)", n, factors)
			}
		}
	}

}

// multiplySliceBigInt возвращает произведение всех элементов слайса больших целых
func multiplySliceBigInt(arr []*big.Int) *big.Int {
	result := big.NewInt(int64(1))
	for _, element := range arr {
		result.Mul(result, element)
	}
	return result
}
