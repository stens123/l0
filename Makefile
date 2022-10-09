
APP=Service
PUB=PublisherServ

all: $(APP) $(PUB)

$(APP):
	go build -o $@ main.go

$(PUB):
	go build -o $@ Publisher/publisher.go

fclean: clean
	@rm $(APP) $(PUB) 2> /dev/null &
	killall Service
	@echo "$(APP) и $(PUB) удалены"

clean:
	@bash close.sh
	@echo "порты освобождены"

run: $(APP) $(PUB)
	bash run.sh
	
