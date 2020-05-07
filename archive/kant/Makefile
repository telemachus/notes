SHELL = /bin/sh
LATEXMK = latexmk
FLAGS = -xelatex -bibtex

all: groundwork-of-metaphysics-of-morals

groundwork-of-metaphysics-of-morals: groundwork-of-metaphysics-of-morals.tex kant.bib
	$(LATEXMK) $(FLAGS) $@

clean:
	- $(RM) build/*

.PHONY : clean
