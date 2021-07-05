go build -o bin/wb_parser;
#env GOOS=windows GOARCH=386 go build -o bin/wb_parser_win_x86.exe;
env GOOS=windows GOARCH=amd64 go build -o bin/wb_parser_win_amd64.exe;
#env GOOS=windows GOARCH=arm go build -o bin/wb_parser_win_arm.exe;
