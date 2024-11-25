GOFILES := main.go retrieve.go get_token.go ./auth/callback.go
OUTPUT := database
FLAGS := -o

all: $(OUTPUT)

$(OUTPUT): $(GOFILES)
	go build $(FLAGS) $(OUTPUT) $(GOFILES)

clean:
	rm -f $(OUTPUT)

