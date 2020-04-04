# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Extended Go Data Structures & Algorithms
[![](https://img.shields.io/github/v/release/codemodify/systemkit-dsa?style=flat-square)](https://github.com/codemodify/systemkit-dsa/releases/latest)
![](https://img.shields.io/github/languages/code-size/codemodify/systemkit-dsa?style=flat-square)
![](https://img.shields.io/github/last-commit/codemodify/systemkit-dsa?style=flat-square)
[![](https://img.shields.io/badge/license-0--license-brightgreen?style=flat-square)](https://github.com/codemodify/TheFreeLicense)

![](https://img.shields.io/github/workflow/status/codemodify/systemkit-dsa/qa?style=flat-square)
![](https://img.shields.io/github/issues/codemodify/systemkit-dsa?style=flat-square)
[![](https://goreportcard.com/badge/github.com/codemodify/systemkit-dsa?style=flat-square)](https://goreportcard.com/report/github.com/codemodify/systemkit-dsa)

[![](https://img.shields.io/badge/godoc-reference-brightgreen?style=flat-square)](https://godoc.org/github.com/codemodify/systemkit-dsa)
![](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)
![](https://img.shields.io/gitter/room/codemodify/systemkit-dsa?style=flat-square)

![](https://img.shields.io/github/contributors/codemodify/systemkit-dsa?style=flat-square)
![](https://img.shields.io/github/stars/codemodify/systemkit-dsa?style=flat-square)
![](https://img.shields.io/github/watchers/codemodify/systemkit-dsa?style=flat-square)
![](https://img.shields.io/github/forks/codemodify/systemkit-dsa?style=flat-square)

#### The missing Go lang extended structures & algorithms
#### Supported: Linux, Raspberry Pi, FreeBSD, Mac OS, Windows, Solaris

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Install
```go
go get github.com/codemodify/systemkit-dsa
```
# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) API

&nbsp;																| &nbsp;
---     															| ---
NewRingBuffer(`in` chan []byte, `out` chan []byte) `*RingBuffer`  |
NewRingBufferHP(bufferSize uint64) `*RingBufferHP` |
Write(`value` []byte) |
Read() |
NewConcurrentMap() |
Add(`key` interface{}) |
Remove(`key` interface{}) |
Contains(`key` interface{}) |
Len() |
Less(`i`, `j` int) |
Swap(`i`, `j` int) |
Push(`x` interface{}) |
Pop() |
Update(i`t`em *Item, `value` *interface{}, `priority` int64) |
Dump() |
Peek() |

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Usage
```go

```