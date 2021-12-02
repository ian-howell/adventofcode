#include <stdio.h>

int main() {
    // a, b, and c represent the sliding window
    // d represents the next value to be added to the window
    int a, b, c, d;

    // load the initial window
    scanf("%d %d %d", &a, &b, &c);

    int last_sum = a + b + c;
    int current_sum = 0;
    int num_decreases = 0;

    while (scanf("%d", &d) != EOF) {
        current_sum = b + c + d;
        if (current_sum > last_sum) {
            num_decreases++;
        }
        last_sum = current_sum;
        b = c;
        c = d;
    }
    printf("%d\n", num_decreases);
}
