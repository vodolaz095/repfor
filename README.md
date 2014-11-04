REPeat FORever
================================
Simple console utility to repeat command given forever.
I do not care about you, but for me this code saved my day.
You can say, that we can achive the same thing using cron, but cron do not allows 
us to execute commands more often than once per minute.

Usage
================================
`repfor [-d=3] [-o=/dev/stdout] [-e=/dev/stderr] commandToRepeatForeverWithDelayProvided`

Parameters
================================

* d - delay in seconds, before repeating the command, default is 3 seconds
* o - where to stream stdout of command, default is /dev/stdout
* e - where to stream stderror of command, default is /dev/stderr


Examples:
================================
```shell

	[vodolaz095@steel ~]$ repfor 'cat /etc/issue'
	[vodolaz095@steel ~]$ repfor -d=5 -o=output.log -e=error.log 'ps -e'

```

Installation from source:
================================
We assume you have [Go](http://golang.org/) installed
```shell

	[vodolaz095@steel ~]$ go build -o ~/bin/repfor github.com/vodolaz095/repfor

```

License
================================
The MIT License (MIT)

Copyright (c) 2013 Ostroumov Anatolij <ostroumov095 at gmail dot com>

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
