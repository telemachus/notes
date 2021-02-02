SHELL = /bin/sh
LATEXMK = latexmk
FLAGS = -xelatex -bibtex

all: midlife

midlife: midlife.tex midlife.bib
	$(LATEXMK) $(FLAGS) $@

clean:
	- $(RM) build/*

.PHONY : clean
