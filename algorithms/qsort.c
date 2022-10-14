/*
    It shows an example of how to sort
    an array of strings with library function qsort
    (quick sort):

    (from ISO/IEC 9899:2017 :: 2x (E))

    #include <stdlib.h>
    void qsort(void *base, size_t nmemb, size_t size,
        int (*compar)(const void *, const void *));

    The qsort function sorts an array of nmemb objects,
    the initial element of which is pointed to by base.
    The size of each object is specified by size.
*/

#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char *words[] = {"how", "are", "you", "I", "am", "fine",};
int N = sizeof(words)/sizeof(words[0]);

int scmp(const void* p1, const void * p2) {
    return strcmp(*(char **)p1, *(char **)p2);
}

void printarray(char **arr, int nmemb);

int main(void) {

    printf("before qsort: ");
    printarray(words, N);

    qsort(words, N, sizeof(words[0]), scmp);
    printf("after qsort: ");
    printarray(words, N);
    return 0;
}


void printarray(char **arr, int nmemb) {
    int i;
    for (i = 0; i < nmemb; i++) {
        printf("%s ", arr[i]);
    }
    printf("\n");
}

/* output:

ASM generation compiler returned: 0
Execution build compiler returned: 0
Program returned: 0
before qsort: how are you I am fine 
after qsort: I am are fine how you

*/