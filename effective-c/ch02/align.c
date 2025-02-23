// Compile with /std:c11

#include <stdio.h>
#include <stdalign.h>

struct S {
	int i; double d; char c;
};
int main(void) {
	unsigned char bad_buff[sizeof(struct S)];
	_Alignas(struct S) unsigned char good_buff[sizeof(struct S)];
	struct S *bad_s_ptr = (struct S *)bad_buff;   // wrong pointer alignment
	struct S *good_s_ptr = (struct S *)good_buff; // correct pointer alignment
	printf("sizeof(bad_buff): %lu\n", sizeof(bad_buff));
	printf("alignof(bad_buff): %lu\n", alignof(bad_buff));
	printf("sizeof(good_buff): %lu\n", sizeof(good_buff));
	printf("alignof(good_buff): %lu\n", alignof(good_buff));
}
