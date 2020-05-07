SHELL = /bin/sh
LATEXMK = latexmk
FLAGS = -xelatex -bibtex

all: true-belief

true-belief: true-belief-main.tex sartwell.tex plato-sartwell.tex tbk.bib
	$(LATEXMK) $(FLAGS) true-belief-main.tex -jobname=true-belief

clean:
	- $(RM) *.aux *.log *.pdf *.bbl *.blg *.fls *.fdb_latexmk *.bcf *.xml

.PHONY : clean
