#include <stdio.h>

int bisect(const char input[], int start, int len, int (*cmp)(const char*, int));
int cmp_row(const char input[], int i);
int cmp_col(const char input[], int i);
int get_seat_id(int row, int col);

int main() {
    int boarding_passes[8*128] = {0};
    char input[11];
    while (scanf("%10s", input) == 1) {
        int row = bisect(input, 0, 7, &cmp_row);
        int col = bisect(input, 7, 3, &cmp_col);
        int seat_id = get_seat_id(row, col);
        boarding_passes[seat_id] = 1;
    }

    int my_boarding_pass = 0;
    for (int i = 0; i < 8*128; i++) {
        printf("Boarding pass #%d is ", i);
        if (boarding_passes[i]) {
            printf("\e[0;32mFound\e[0m\n");
        } else {
            printf("\e[0;31mMissing\e[0m\n");
        }

        if ((i > 0) && (i < 8*128 - 1) && !boarding_passes[i] && boarding_passes[i-1] && boarding_passes[i+1]) {
            my_boarding_pass = i;
        }
    }

    if (my_boarding_pass) {
        printf("My boarding pass is %d\n", my_boarding_pass);
    } else {
        printf("Couldn't find my boarding pass :(\n");
    }

    return 0;
}

int bisect(const char input[], int start, int len, int (*cmp)(const char*, int)) {
    int lb = 0;
    int ub = (1 << len) - 1;

    for (int i = start; i < (start+len); i++ ) {
        if (cmp(input, i)) {
            ub = (lb + ub) / 2;
        } else {
            lb = 1 + ((lb + ub) / 2);
        }
    }
    return ub;
}

int cmp_row(const char input[], int i) {
    return input[i] == 'F';
}

int cmp_col(const char input[], int i) {
    return input[i] == 'L';
}

int get_seat_id(int row, int col) {
    return 8*row + col;
}
