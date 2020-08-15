# LetsMove

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/abemedia/letsmove?tab=doc)

A Go package that moves the running MacOS application to the Applications
folder using [potionfactory/LetsMove](https://github.com/potionfactory/LetsMove).

![Screenshot](http://i.imgur.com/euTRZiI.png)

## Requirements

Builds and runs on MacOS 10.6 or higher.  
Made for .app bundles, not single-binary command line apps.  
Does NOT support sandboxed applications.

## Example

```go
package main

import "github.com/abemedia/letsmove"

func init() {
	letsmove.MoveToApplications()
}

func main()  {
	// your business logic...
}
```

## Localisation

To translate the messages download the `MoveApplication.string` files from
<https://github.com/potionfactory/LetsMove> and place them in your bundle's
`Contents/Resources` directory in a folder named `${ISO_CODE}.lproj` e.g.
`MyBundle.app/Contents/Resources/en.lproj/MoveApplication.string`.
