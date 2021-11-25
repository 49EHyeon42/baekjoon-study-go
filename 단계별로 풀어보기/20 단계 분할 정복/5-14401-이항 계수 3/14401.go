package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader  = bufio.NewReader(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
	N, K, P = 0, 0, 1000000007
)

// ! 주제에 안맞는 어려운 문제
func main() {
	fmt.Fscanln(reader, &N, &K)

	numer := factorial(N)
	denom := factorial(K) * factorial(N-K) % P

	fmt.Fprintln(writer, numer*pow(denom, P-2)%P)
	writer.Flush()
}

func factorial(N int) int {
	fac := 1
	for N > 1 {
		fac = (fac * N) % P
		N -= 1
	}
	return fac
}

func pow(base, expo int) int {
	result := 1
	for expo > 0 {
		if expo%2 == 1 {
			result *= base
			result %= P
		}
		base = (base * base) % P
		expo /= 2
	}
	return result
}

/*
모듈러 산술(Modular Arithmetic)
모듈러 산술(모듈러 연산)은 정수의 합과 곱을 어떤 주어진 수의 나머지를 이용하여 정의하는 방법을 말한다.
쉽게 말해 나머지를 이용한 산술 연산
r = a % m (r : 나머지, a : 피제수, m : 제수)
r = a mod m
a - b 가 m 이라는 정수로 나누어 떨어진다면, mod m 과 같이 표현
2%5=2, 7%5=2, 12%5=2 ...
모듈러 산술 연산은 분배 법칙이 성립
나눗셈의 경우 곱셈의 역원을 이용하여 분배법칙을 적용

1. 분수 꼴에 대한 모듈러 연산은 분배법칙을 적용할 수 없다
2. 분수 꼴인 식을 곱셈 꼴을 만든다. (곱셈의 역원)
여기서 문제는 (K!(N-K)!) 의 역원을 구하는 것이 관건
이부분에서는 '페르마의 소정리'를 적용

* 모듈러 산술, 곱셈의 역원, 페르마의 소정리를 알아야 알고리즘 해결 가능

알고리즘은 2가지 방식으로 재귀 또는 반복문으로 풀면된다.
제출은 반복으로 작성
link : https://st-lab.tistory.com/241
*/
