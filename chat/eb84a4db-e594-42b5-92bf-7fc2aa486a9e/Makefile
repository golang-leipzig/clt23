TARGETS = suggestion.png guess.png

.PHONY: all
all: $(TARGETS)

suggestion.png: suggestion.dot
	dot -Tpng suggestion.dot > suggestion.png

guess.png: guess.dot
	dot -Tpng guess.dot > guess.png

.PHONY: clean
clean:
	rm -f $(TARGETS)
