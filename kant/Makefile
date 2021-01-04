SHELL = /bin/sh
LATEXMK = latexmk
FLAGS = -xelatex -bibtex

all: groundwork blackwell-kant

groundwork: groundwork.tex kant.bib
	$(LATEXMK) $(FLAGS) $@

blackwell-kant: blackwell-kant.tex kant.bib
	$(LATEXMK) $(FLAGS) $@

clean:
	- $(RM) build/*

.PHONY : clean
