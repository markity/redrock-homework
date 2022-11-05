#include <stdio.h>
#include <math.h>

int main() {
    double x;
    scanf("%lf", &x);

    double result;
    
    if (x < -3) {
        result = x -1;
    } else if (x >= -3 && x <= 3) {
        result = sqrt(9-pow(x, 2));
    } else {
        result = log10(x);
    }

    printf("%lf\n", result);
    return 0;
}