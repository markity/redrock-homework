#include <stdio.h>
#include <math.h>

// 判断 2 以上的 int 类型数字是否为素数
int isPrime(int num) {
	if (num == 2) {
		return 1;
	}

	for (size_t i = 2; i < num; i++) {
		if (num % i == 0) {
			return 0;
		}
	}

	return 1;
}

int main() {

	int n;
	scanf("%d", &n);

	int nPrinted = 0;

	int k = 2;
	while (true) {
		if(isPrime(k)) {
			nPrinted++;
			if(nPrinted % 8 == 1) {
				printf("%d", k);
			} else {
				printf(" %d", k);
			}
			if(nPrinted % 8 == 0) {
				printf("\n");
			}
			if (nPrinted == n) {
				return 0;
			}
		}
		k++;
	}
	
	return 0;
}