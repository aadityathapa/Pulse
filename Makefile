.PHONY: build run stop logs test clean

IMAGE := pulse
NAME  := pulse

build:
	docker build -t $(IMAGE) .

run: build
	-docker rm -f $(NAME)
	docker run -d --name $(NAME) -p 8080:8080 --restart unless-stopped $(IMAGE)

stop:
	docker rm -f $(NAME)

logs:
	docker logs -f $(NAME)

test:
	@curl -fsS localhost:8080/health && echo && echo "PASS" || (echo "FAIL" && exit 1)

clean: stop
	docker rmi $(IMAGE)
