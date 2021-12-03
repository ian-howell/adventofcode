#include <stdio.h>

// for small input
#ifdef SMALL
  #define WORD_SIZE 5
  #define NUM_WORDS 12
  #define ALL_ON 0x1f
#else
  #define WORD_SIZE 12
  #define NUM_WORDS 1000
  #define ALL_ON 0xfff
#endif

int main() {
    int counts[WORD_SIZE] = {0};
    for (int i = 0; i < NUM_WORDS; i++) {
        char str[WORD_SIZE+1];
        scanf("%s", str);
        for (int j = 0; j < WORD_SIZE; j++) {
            if (str[j] == '1') {
                counts[j]++;
            }
        }
    }

    int gamma = 0;
    for (int i = 0; i < WORD_SIZE; i++) {
        if (counts[i] > NUM_WORDS/2) {
            gamma |= (1 << (WORD_SIZE-1-i));
        }
    }

    int epsilon = ALL_ON ^ gamma;
    printf("%d\n", gamma * epsilon);

    return 0;
}
