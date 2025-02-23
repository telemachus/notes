#include <stdio.h>
#include <stdlib.h>

void
increment(void)
{
	unsigned int counter = 0;
	counter++;
	printf("%d ", counter);
}

int
main(void)
{
	for (int i = 0; i < 5; i++) {
		increment();
	}
	puts("");

	return EXIT_SUCCESS;
}
