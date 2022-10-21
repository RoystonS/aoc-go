SRCS := $(wildcard *.go) $(wildcard ../../aoccommon/*.go)
OUT := aoc-day

${OUT}: ${SRCS}
	go build -o ${OUT}

clean:
	rm -f ${OUT}

1: ${OUT}
	./${OUT} <input 1

2: ${OUT}
	./${OUT} <input 2

test1: ${OUT}
	./${OUT} <input test1

test2: ${OUT}
	./${OUT} < input test2