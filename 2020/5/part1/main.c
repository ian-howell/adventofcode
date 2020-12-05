#include <stdio.h>

int bisect(const char input[], int start, int len, int (*cmp)(const char*, int));
int cmp_row(const char input[], int i);
int cmp_col(const char input[], int i);
int get_seat_id(int row, int col);

int main() {
    int max_seat_id = -1;
    char input[11];
    while (scanf("%10s", input) == 1) {
        int row = bisect(input, 0, 7, &cmp_row);
        int col = bisect(input, 7, 3, &cmp_col);
        int seat_id = get_seat_id(row, col);
        max_seat_id = max_seat_id > seat_id ? max_seat_id : seat_id;
    }
    printf("%d\n", max_seat_id);
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
