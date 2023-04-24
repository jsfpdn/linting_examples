RULES := rules.go

.PHONY: run
run:
	ruleguard -rules $(RULES) ../example.go

.PHONY: run-fixed
run-fixed:
	ruleguard -rules $(RULES) ../example_fixed.go
