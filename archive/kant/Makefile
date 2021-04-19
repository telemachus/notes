SHELL = /bin/sh
LATEXMK = latexmk
FLAGS = -xelatex -bibtex

all: groundwork blackwell-kant critique-of-pure-reason class-readings

groundwork: groundwork.tex kant.bib
	$(LATEXMK) $(FLAGS) $@

blackwell-kant: blackwell-kant.tex kant.bib
	$(LATEXMK) $(FLAGS) $@

critique-of-pure-reason: critique-of-pure-reason.tex kant.bib
	$(LATEXMK) $(FLAGS) $@

class-readings: class-readings.tex kant.bib
	$(LATEXMK) $(FLAGS) $@

clean:
	- $(RM) build/*

.PHONY : clean
