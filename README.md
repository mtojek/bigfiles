# BigFiles

Status: **Done** (waiting for feedback)

[![Build Status](https://travis-ci.org/mtojek/bigfiles.svg?branch=master)](https://travis-ci.org/mtojek/bigfiles)

## Description

Are you bored with overloaded speed test services? You don't have to use them at all, because you can run own webserver, serving **large test files** (custom size: **100 MB**, **100 GB**, **1 TB**..., upto **8192 PB**). 

There is no more need to use publicly hosted storage services to download ordinary **100MB.bin**. With a single command spawn **own speed test** instance to verify your Internet provider.

### Screenshots

#### Index view

<img src="https://github.com/mtojek/bigfiles/blob/master/screenshot-1.png" alt="Screenshot Desktop" width="480px" />

#### Chrome "Downloads" page

<img src="https://github.com/mtojek/bigfiles/blob/master/screenshot-2.png" alt="Screenshot Mobile" width="480px" />

## Features

* download huge files < **8192 PB**
* choose between **zeros** file, random content, and repeating a user-defined message
* **easy to use** HTTP GET endpoints (Chrome, curl, etc.)
* user-defined **file size limit**

## Quickstart

Download and install BigFiles:
```bash
go get github.com/mtojek/bigfiles
```
Run the application:
```bash
bigfiles
```

Use wget to download a sparse file:
```bash
wget http://localhost:5000/files/sparse/100MB
--2017-10-01 00:34:18--  http://localhost:5000/files/sparse/100MB
Resolving localhost... ::1, fe80::1, 127.0.0.1
Connecting to localhost|::1|:5000... connected.
HTTP request sent, awaiting response... 200 OK
Length: 104857600 (100M) [application/octet-stream]
Saving to: '100MB'

100MB                         100%[===================================================>] 100.00M   218MB/s   in 0.5s   

2017-10-01 00:34:19 (218 MB/s) - '100MB' saved [104857600/104857600]

```

or go to web panel to download files:

[http://localhost:5000](http://localhost:5000)

## Contact

Please feel free to leave any comment or feedback by opening a new issue or contacting me directly via [email](mailto:marcin@tojek.pl). Thank you.

## License

MIT License, see [LICENSE](https://github.com/mtojek/bigfiles/blob/master/LICENSE) file.
