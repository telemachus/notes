SHELL = /bin/sh
LATEXMK = latexmk
FLAGS = -xelatex -bibtex

all: recollection

recollection: recollection-main.tex recollection.bib
	$(LATEXMK) $(FLAGS) recollection-main.tex -jobname=recollection

clean:
	- $(RM) *.aux *.log *.pdf *.bbl *.blg *.fls *.fdb_latexmk *.bcf *.xml

.PHONY : clean
