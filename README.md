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
* choose between **zeros** file or random content
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

or go to web panel:

[http://localhost:5000](http://localhost:5000)

## Contact

Please feel free to leave any comment or feedback by opening a new issue or contacting me directly via [email](mailto:marcin@tojek.pl). Thank you.

## License

MIT License, see [LICENSE](https://github.com/mtojek/bigfiles/blob/master/LICENSE) file.
