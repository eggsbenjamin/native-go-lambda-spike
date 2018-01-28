all:
	for tag in "compression"; do \
		tag=$$tag make build; \
		tag=$$tag make deploy; \
	done

build:
	mkdir -p artifacts
	set -e && for pkg in $$(ls lambda); do \
		out=$$pkg; \
		if [ -n "$$tag" ]; then out=$$tag'_'$$pkg; fi; \
		echo "building $$out\n"; \
		GOOS=linux CGO_ENABLED=0 go build -o ./artifacts/$$out ./lambda/$$pkg; \
		zip -j ./artifacts/$$out.zip ./artifacts/$$out; \
	done

deploy:
	cd ./terraform; terraform apply -var 'tag=$$tag' -auto-approve
	
