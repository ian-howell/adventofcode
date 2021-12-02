#include <stdio.h>

int main() {
    int last;
    int current;
    int num_decreases = 0;

    scanf("%d", &last);

    while (scanf("%d", &current) != EOF) {
        if (current > last) {
            num_decreases++;
        }
        last = current;
    }
    printf("%d\n", num_decreases);
}
