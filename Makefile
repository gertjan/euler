default:
	go test -v . -run TestP

fmt:
	goimports -w .
