SHELL = /bin/sh
LATEXMK = latexmk
FLAGS = -xelatex -bibtex-cond -pdfxe -jobname=pyrrhonian-outlines

all: pyrrhonian-outlines

pyrrhonian-outlines: titlepage.tex toc.tex main.tex \
	book-one.tex skepticism.bib
	$(LATEXMK) $(FLAGS) main.tex

clean:
	- $(RM) *.{aux,log,pdf,bbl,blg,fls,fdb_latexmk,end,out,xdv}
	- $(RM) *.{bcf,run.xml,toc}

.PHONY : clean
