#include <stdio.h>
#include <stdlib.h>

void swap(int*, int*);

int
main(void)
{
	int a = 21;
	int b = 17;

	printf("before swap: a = %d, b = %d\n", a, b);
	swap(&a, &b);
	printf("after  swap: a = %d, b = %d\n", a, b);

	return EXIT_SUCCESS;
}

void
swap(int *a, int *b)
{
	int t = *a;
	*a = *b;
	*b = t;
	return;
}
