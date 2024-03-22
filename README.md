# kruise-api

Schema of the API types that are served by Kruise.

## Purpose

This library is the canonical location of the Kruise API definition and client.

We recommend using the go types in this repo. You may serialize them directly to JSON.

## What's included
* The `client` package contains the clientset to access Kruise API.
* The `apps` and `policy` packages contain api definition in go
* The `schema` directory contains  corresponding openapi schema of kruise source 

## Versioning
For each `v1.x.y` Kruise release, the corresponding kruise-api will `v1.x.z`. 

Bugfixes in kruise-api will result in the patch version (third digit `z`) changing. PRs that are cherry-picked into an older Kruise release branch will result in an update to the corresponding branch in client-go, with a corresponding new tag changing the patch version.

## Where does it come from?

`kruise-api` is synced from [https://github.com/openkruise/kruise/tree/master/apis](https://github.com/openkruise/kruise/tree/master/apis).
Code changes are made in that location, merged into `openkruise/kruise` and later synced here.


### How to get it

To get the latest version, use go1.16+ and fetch using the `go get` command. For example:

```
go get github.com/openkruise/kruise-api@latest
```

To get a specific version, use go1.11+ and fetch the desired version using the `go get` command. For example:

```
go get github.com/openkruise/kruise-api@v1.6.0
```

### How to use it

please refer to the [example](examples/create-update-delete-cloneset)

## Things you should NOT do

[https://github.com/openkruise/kruise/tree/master/apis](https://github.com/openkruise/kruise/tree/master/apis) is synced to here.
All changes must be made in the former. The latter is read-only.

