all:
	set -e && for tag in $$(git tag --list); do \
		git checkout tags/$$tag; \
		tag=$$tag make deploy; \
	done
	git checkout master

test:
	go test ./... -v

build: test
	mkdir -p artifacts
	set -e && for pkg in $$(ls lambda); do \
		out=$$pkg; \
		if [ -n "$$tag" ]; then out=$$tag'_'$$pkg; fi; \
		echo "building $$out\n"; \
		GOOS=linux CGO_ENABLED=0 go build -o ./artifacts/$$out ./lambda/$$pkg; \
		zip -j ./artifacts/$$out.zip ./artifacts/$$out; \
	done

deploy: build
	cd ./terraform; TF_VAR_tag=$$tag'_' terraform apply -auto-approve
	
