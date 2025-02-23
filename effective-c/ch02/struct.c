#include <stdio.h>
#include <string.h>
 
int main(void)
{
	struct sigrecord {
		int signum;
		char signame[20];
		char sigdesc[100];
	} sigline, *sigline_p = &sigline;
	sigline.signum = 5;
	strcpy(sigline.signame, "SIGINT");
	strcpy(sigline.sigdesc, "Interrupt from keyboard");
	printf("signal %d: %s => %s\n", sigline_p->signum, sigline_p->signame, sigline.sigdesc);
}
