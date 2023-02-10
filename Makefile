flow.png: flow.dot
	dot -Tpng flow.dot > flow.png

.PHONY: clean
clean:
	rm -f flow.png
