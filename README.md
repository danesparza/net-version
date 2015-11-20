# net-version
.NET version checker written in Go.   

Uses the [guidelines outlined by Microsoft](https://msdn.microsoft.com/en-us/library/hh925568%28v=vs.110%29.aspx) to determine what version(s) of .NET are installed on the current machine.

Build: [![Circle CI](https://circleci.com/gh/danesparza/net-version.svg?style=svg)](https://circleci.com/gh/danesparza/net-version)

### Installing
Just [grab the latest release](https://github.com/danesparza/net-version/releases) (it's just a single binary) and run from the command line.  

### Example

```
$ net-version.exe
Checking to see what versions of .NET are installed...

Older .NET versions installed:
============================
v2.0.50727
v3.0
v3.5
v4
v4.0

Newer .NET versions installed:
============================
v4.6 (on an operating system other than Windows 10)

Finished.
```
