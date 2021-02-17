Sometimes it happens you see in log something like:
```
panic: [103 111 114 111 117 116 105 110 101 32 49 32 91 114 117 110 110 105 110 103 93 58 10 114 117 110 116 105 109 101 47 100 101 98 117 103 46 83 116 97 99 107 40 48 120 99 48 48 48 48 49 49 51 48 49 44 32 48 120 99 48 48 48 48 53 101 49 100 48 44 32 48 120 99 48 48 48 48 55 99 101 99 48 41 10 9 47 104 111 109 101 47 120 97 105 111 110 97 114 111 47 46 103 105 109 109 101 47 118 101 114 115 105 111 110 115 47 103 111 49 46 49 51 46 108 105 110 117 120 46 97 109 100 54 52 47 115 114 99 47 114 117 110 116 105 109 101 47 100 101 98 117 103 47 115 116 97 99 107 46 103 111 58 50 52 32 43 48 120 57 100 10 109 97 105 110 46 109 97 105 110 40 41 10 9 47 104 111 109 101 47 120 97 105 111 110 97 114 111 47 103 111 47 115 114 99 47 103 105 116 104 117 98 46 99 111 109 47 120 97 105 111 110 97 114 111 45 103 111 47 105 110 116 115 50 115 116 114 105 110 103 47 109 97 105 110 46 103 111 58 49 50 32 43 48 120 51 52 10]
```
This tool helps to convert this integers to a string:
```
xaionaro@void:~/go/src/github.com/xaionaro-go/ints2string$ go run ./main.go 103 111 114 111 117 116 105 110 101 32 49 32 91 114 117 110 110 105 110 103 93 58 10 114 117 110 116 105 109 101 47 100 101 98 117 103 46 83 116 97 99 107 40 48 120 99 48 48 48 48 57 50 48 48 49 44 32 48 120 99 48 48 48 48 54 52 49 100 48 44 32 48 120 99 48 48 48 48 56 50 101 99 48 41 10 9 47 104 111 109 101 47 120 97 105 111 110 97 114 111 47 46 103 105 109 109 101 47 118 101 114 115 105 111 110 115 47 103 111 49 46 49 51 46 108 105 110 117 120 46 97 109 100 54 52 47 115 114 99 47 114 117 110 116 105 109 101 47 100 101 98 117 103 47 115 116 97 99 107 46 103 111 58 50 52 32 43 48 120 57 100 10 109 97 105 110 46 109 97 105 110 40 41 10 9 47 104 111 109 101 47 120 97 105 111 110 97 114 111 47 103 111 47 115 114 99 47 103 105 116 104 117 98 46 99 111 109 47 120 97 105 111 110 97 114 111 45 103 111 47 105 110 116 115 50 115 116 114 105 110 103 47 109 97 105 110 46 103 111 58 49 50 32 43 48 120 51 52 10
```
output:
```
goroutine 1 [running]:
runtime/debug.Stack(0xc000092001, 0xc0000641d0, 0xc000082ec0)
        /home/xaionaro/.gimme/versions/go1.13.linux.amd64/src/runtime/debug/stack.go:24 +0x9d
main.main()
        /home/xaionaro/go/src/github.com/xaionaro-go/ints2string/main.go:12 +0x34
```