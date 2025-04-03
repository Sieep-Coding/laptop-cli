runAll:
	go run .

fmt:
	go fmt .
	commands/ go fmt.

backup:
	mkdir -p backup/
	cp *.db backup/
