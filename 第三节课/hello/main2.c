#include <stdio.h>

int func(int t) {
    if(t == 1) {
        return 1;
    }
    return t * func(t-1);
}

int main() {
    int n;
    scanf("%d", &n);

    int total = 0;

    for (int i = 1; i <= n; i++)
    {
        total += func(i);
    }

    printf("%d\n", total);
    
}