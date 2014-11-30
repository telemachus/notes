SHELL = /bin/sh
LATEXMK = latexmk
FLAGS = -xelatex

all: review

review: titlepage.tex review.tex review.bib toc.tex chapter1.tex
	$(LATEXMK) $(FLAGS) review.tex

clean:
	- $(RM) *.{aux,log,pdf,bbl,blg,fls,1,fdb_latexmk,end,eledsec1,out,toc}

.PHONY : clean
