#include <stdio.h>
#include <string.h>

int main() {
    int depth = 0;
    int h_pos = 0;

    char direction[10];
    int dist;

    while (scanf("%s %d", direction, &dist) != EOF) {
        if (strcmp(direction, "forward") == 0) {
            h_pos += dist;
        } else if (strcmp(direction, "down") == 0) {
            depth += dist;
        } else if (strcmp(direction, "up") == 0) {
            depth -= dist;
        }
    }

    printf("%d\n", depth * h_pos);
}
