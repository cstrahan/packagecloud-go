# Reference
## distributions
<details><summary><code>client.Distributions.Index() -> map[string][]*packagecloud.Distribution</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

**Example request:**

```
curl https://packagecloud.io/api/v1/distributions.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Distributions.Index(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## gpg_keys
<details><summary><code>client.GpgKeys.Index(UserID, Repo) -> *packagecloud.GpgKeysIndexResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List gpg keys for a repository.

**Example request:**

```
curl https://packagecloud.io/api/v1/repos/username/reponame/gpg_keys.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.GpgKeysIndexRequest{
        UserID: "user_id",
        Repo: "repo",
    }
client.GpgKeys.Index(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GpgKeys.Create(UserID, Repo, request) -> *packagecloud.GpgKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a GPG key.

**Example request:**

```
curl -X POST -F "gpg_key[keydata]=@/path/to/key.gpg" https://packagecloud.io/api/v1/repos/username/reponame/gpg_keys.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.GpgKeysCreateRequest{
        UserID: "user_id",
        Repo: "repo",
        GpgKeyKeydata: strings.NewReader(
            "",
        ),
    }
client.GpgKeys.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username to which the repo belongs.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repo.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GpgKeys.Show(UserID, Repo, Keyname) -> *packagecloud.GpgKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

View a single GPG Key

**Example request:**

```
curl -X GET https://packagecloud.io/api/v1/repos/username/reponame/gpg_keys/username-reponame-85D7FBF915DFCBC6.pub.gpg.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.GpgKeysShowRequest{
        UserID: "user_id",
        Repo: "repo",
        Keyname: "keyname",
    }
client.GpgKeys.Show(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username to which the repo belongs.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repo.
    
</dd>
</dl>

<dl>
<dd>

**keyname:** `string` — The GPG Key keyname.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.GpgKeys.Destroy(UserID, Repo, Keyname) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Destroy a GPG key

**Example request:**

```
curl -X DELETE https://packagecloud.io/api/v1/repos/username/reponame/gpg_keys/username-reponame-85D7FBF915DFCBC6.pub.gpg.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.GpgKeysDestroyRequest{
        UserID: "user_id",
        Repo: "repo",
        Keyname: "keyname",
    }
client.GpgKeys.Destroy(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username to which the repo belongs.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repo.
    
</dd>
</dl>

<dl>
<dd>

**keyname:** `string` — The GPG Key keyname to delete
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## licenses
<details><summary><code>client.Licenses.Index(LicenseKey) -> *packagecloud.License</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

STILL IN USE BY GITLAB
Retrieve a packagecloud:enterprise license file and GPG signature.
Signature can be verified with the packagecloud GPG key:
<https://packagecloud.io/gpg.key>.

**Example request:**

```
curl https://packagecloud.io/api/v1/licenses/aabbccddeeffaa/license.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.LicensesIndexRequest{
        LicenseKey: "license_key",
    }
client.Licenses.Index(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**licenseKey:** `string` — The packagecloud:enterprise license key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## master_tokens
<details><summary><code>client.MasterTokens.Index(UserID, Repo) -> []*packagecloud.MasterToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List master tokens for a repo.

**Example request:**

```
curl https://packagecloud.io/api/v1/repos/username/reponame/master_tokens
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.MasterTokensIndexRequest{
        UserID: "user_id",
        Repo: "repo",
    }
client.MasterTokens.Index(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username to which the repo belongs.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repo.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MasterTokens.Create(UserID, Repo, request) -> *packagecloud.MasterToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a master token.

**Example request:**

```
curl -X POST -F "master_token[name]=app_servers" https://packagecloud.io/api/v1/repos/username/reponame/master_tokens
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.MasterTokensCreateRequest{
        UserID: "user_id",
        Repo: "repo",
    }
client.MasterTokens.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username to which the repo belongs.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repo.
    
</dd>
</dl>

<dl>
<dd>

**masterTokenName:** `*string` — The name of the token to create.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MasterTokens.Search(UserID, Repo) -> []*packagecloud.MasterTokensResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Search for master tokens

Default is 30 per page, see response headers for [pagination](#pagination) information.

**Example request:**

```
curl https://packagecloud.io/api/v1/repos/username/reponame/master_tokens/search?q=prod
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.MasterTokensSearchRequest{
        UserID: "user_id",
        Repo: "repo",
    }
client.MasterTokens.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username to which the repo belongs.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repo.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — The name of the MasterToken.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MasterTokens.Show(UserID, Repo, ID) -> []*packagecloud.MasterTokensShowResponseItem</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Show a master token.

**Example request:**

```
curl https://packagecloud.io/api/v1/repos/username/reponame/master_tokens/1
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.MasterTokensShowRequest{
        UserID: "user_id",
        Repo: "repo",
        ID: 1,
    }
client.MasterTokens.Show(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username to which the repo belongs.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repo.
    
</dd>
</dl>

<dl>
<dd>

**id:** `int` — The id or value of the MasterToken.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.MasterTokens.Destroy(UserID, Repo, ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Destroy a master token and all of its read token children.

NOTE: Default tokens have their token value rotated when used with this
API. All private repositories must have default master tokens.

**Example request:**

```
curl -X DELETE https://packagecloud.io/api/v1/repos/username/reponame/master_tokens/1
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.MasterTokensDestroyRequest{
        UserID: "user_id",
        Repo: "repo",
        ID: 1,
    }
client.MasterTokens.Destroy(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username to which the repo belongs.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repo.
    
</dd>
</dl>

<dl>
<dd>

**id:** `int` — The id of the master token to be destroyed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## packages
<details><summary><code>client.Packages.Contents(UserID, Repo, request) -> []*packagecloud.PackageContents</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Inspect the contents of a debian source package.

**Example request:**

```
curl -X POST                                                  \
     -F "package[distro_version_id]=17"                       \
     -F "package[package_file]=@path/to/jake_1.0-7.dsc"             \
        "https://packagecloud.io/api/v1/repos/test_user/test_repo/packages/contents.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesContentsRequest{
        UserID: "user_id",
        Repo: "repo",
        PackagePackageFile: strings.NewReader(
            "",
        ),
    }
client.Packages.Contents(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.All(UserID, Repo) -> []*packagecloud.PackageFragment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List All Packages (not grouped by package version)

**Example request:**

```
curl "https://packagecloud.io/api/v1/repos/test_user/test_repo/packages.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesAllRequest{
        UserID: "user_id",
        Repo: "repo",
    }
client.Packages.All(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.Create(UserID, Repo, request) -> *packagecloud.PackageDetails</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Push a package.

NOTE: that for Deban .dsc packages, you must upload each of the files contained in the .dsc.
You can use the package\_contents API to determine the filename, md5sum, and size of the files referenced by the DSC.
All other package types (debian, rpm, python, and java) need only specify the distro\_version\_id and package\_file parameters.

**Example request:**

```
curl -X POST https://packagecloud.io/api/v1/repos/test_user/test_repo/packages.json \
     -F "package[distro_version_id]=48"     \
     -F "package[package_file]=@path/to/jake-1.0-1.el6.x86_64.rpm"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesCreateRequest{
        UserID: "user_id",
        Repo: "repo",
        PackagePackageFile: strings.NewReader(
            "",
        ),
    }
client.Packages.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.Promote(UserID, Repo, Distro, Version, Package, request) -> *packagecloud.PackageDetails</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Promote a package.

This API is similar to a move operation. Once this API call completes
successfull, the package being promoted will only exist in the
destination.

The `dependencies` array for Debian packages will list dependencies for the following contexts: `depends`, `recommends`, `suggests`, `enhances`, `pre-depends`

The `dependencies` array for Debian Source packages (DSC) will list dependencies for the following contexts: `build-depends` and `build-depends-indep`

The `dependencies` array for Node.js packages will list dependencies for the following contexts: `dependencies`, `devDependencies`,
`optionalDependencies`, `bundleDependencies`, and `peerDependencies`.

The `dependencies` array for Node.js packages will list dependencies for the following contexts: `dependencies`, `devDependencies`,
`optionalDependencies`, `bundleDependencies`, and `peerDependencies`.

**Example request:**

```
curl -X POST                                \
     -F "destination=julio/testrepo"        \
     "https://packagecloud.io/api/v1/repos/user/example-repo/el/7/jake-1.0-3.el6.x86_64.rpm/promote.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesPromoteRequest{
        UserID: "user_id",
        Repo: "repo",
        Distro: "distro",
        Version: "version",
        Package: "package",
    }
client.Packages.Promote(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**distro:** `string` — The name of the distribution the package is in. (i.e. ubuntu)
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version name of the distribution the package is in. (i.e. precise)
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The full name of the package.
    
</dd>
</dl>

<dl>
<dd>

**scope:** `*string` — The scope of the Node.JS package to promote. Only required if the package to promote is scoped. Include the escaped '@' (i.e. ?scope=%40example).
    
</dd>
</dl>

<dl>
<dd>

**destination:** `*string` — The fully qualified repository name (e.g., "user/repo") to move the package.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.Show(UserID, Repo, Type, Distro, Version, Package, Arch, PackageVersion, Release) -> *packagecloud.PackageDetails</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Show Package

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

The `dependencies` array for RPM packages will list all `Requires` and `PreReq` values under the context `requires`

The `dependencies` array for Debian packages will list dependencies for the following contexts: `depends`, `recommends`, `suggests`, `enhances`, `pre-depends`

The `dependencies` array for Debian Source packages (DSC) will list dependencies for the following contexts: `build-depends` and `build-depends-indep`

**Example request:**

```
curl "https://packagecloud.io/api/v1/repos/test_user/test_repo/package/rpm/fedora/22/jake/x86_64/1.0/1.el6.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesShowRequest{
        UserID: "user_id",
        Repo: "repo",
        Type: "type",
        Distro: "distro",
        Version: "version",
        Package: "package",
        Arch: "arch",
        PackageVersion: "package_version",
        Release: "release",
    }
client.Packages.Show(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `string` — The type of package it is, "rpm" or "deb".
    
</dd>
</dl>

<dl>
<dd>

**distro:** `string` — The name of the distribution the package is in. (i.e. ubuntu)
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version name of the distribution the package is in. (i.e. precise)
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package.
    
</dd>
</dl>

<dl>
<dd>

**arch:** `string` — The architecture of the package. (note: debs use amd64 while rpms use x86\_64 for arch)
    
</dd>
</dl>

<dl>
<dd>

**packageVersion:** `string` — The version (without epoch) of the package.
    
</dd>
</dl>

<dl>
<dd>

**release:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.AnyfileShow(UserID, Repo, Package, Version) -> *packagecloud.PackageDetails</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Show Anyfile Package

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
curl "https://packagecloud.io/api/v1/repos/test_user/test_repo/package/anyfile/hello/0.0.6.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesAnyfileShowRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
        Version: "version",
    }
client.Packages.AnyfileShow(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package.
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version of the package.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.GemShow(UserID, Repo, Package, Version) -> *packagecloud.PackageDetails</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Show RubyGem Package

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

The `dependencies` array for RubyGem packages will list dependencies for the following contexts: `runtime` and `development`

**Example request:**

```
curl "https://packagecloud.io/api/v1/repos/julio/testrepo/package/gem/jakedotrb/0.0.1.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesGemShowRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
        Version: "version",
    }
client.Packages.GemShow(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package.
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version of the package.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.PythonShow(UserID, Repo, Package, Version) -> *packagecloud.PackageDetails</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Show Python Package

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

The `dependencies` array for Python packages will list dependencies for the following contexts: `requires`, `requires-dist`, `requires-external`, `requires-extra`, and `requires-python`

**Example request:**

```
curl "https://packagecloud.io/api/v1/repos/test_user/test_repo/package/python/packagecloud-test/0.9.7b1.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesPythonShowRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
        Version: "version",
    }
client.Packages.PythonShow(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package.
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version of the package.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.JavaShow(UserID, Repo, Package, Version) -> *packagecloud.PackageDetails</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Show Java Package

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
curl "https://packagecloud.io/api/v1/repos/test_user/test_repo/package/java/maven2/io.packagecloud/jake/2.3.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesJavaShowRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
        Version: "version",
    }
client.Packages.JavaShow(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package.
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version of the package.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.HelmShow(UserID, Repo, IndexName, Package, Version) -> *packagecloud.PackageDetails</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Show Helm Package

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
curl "https://packagecloud.io/api/v1/repos/test_user/test_repo/package/helm/v1/hello-kubernetes/2.0.0.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesHelmShowRequest{
        UserID: "user_id",
        Repo: "repo",
        IndexName: "index_name",
        Package: "package",
        Version: "version",
    }
client.Packages.HelmShow(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**indexName:** `string` — The version of helm index
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package.
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version of the package.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.Versions(UserID, Repo, Type, Distro, Version, Package, Arch) -> []*packagecloud.PackageFragment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List Package Versions

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
curl "https://packagecloud.io/api/v1/repos/test_user/test_repo/package/rpm/fedora/22/jake/x86_64/versions.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesVersionsRequest{
        UserID: "user_id",
        Repo: "repo",
        Type: "type",
        Distro: "distro",
        Version: "version",
        Package: "package",
        Arch: "arch",
    }
client.Packages.Versions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `string` — The type of package it is, "rpm" or "deb".
    
</dd>
</dl>

<dl>
<dd>

**distro:** `string` — The name of the distribution the package is in. (i.e. ubuntu)
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version name of the distribution the package is in. (i.e. precise)
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package.
    
</dd>
</dl>

<dl>
<dd>

**arch:** `string` — The architecture of the package. (note: debs use amd64 while rpms use x86\_64 for arch)
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.GemVersions(UserID, Repo, Package) -> []*packagecloud.PackageFragment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List RubyGem Package Versions

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
curl "https://packagecloud.io/api/v1/repos/test_user/test_repo/package/gem/jakedotrb/versions.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesGemVersionsRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
    }
client.Packages.GemVersions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.NodeShow(UserID, Repo, Package, Version) -> *packagecloud.PackageDetails</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Show Node.js Package

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

The `dependencies` array for Node.js packages will list dependencies for the following contexts: `dependencies`, `devDependencies`,
`optionalDependencies`, `bundleDependencies`, and `peerDependencies`.

**Example request:**

```
curl "https://packagecloud.io/api/v1/repos/test_user/test_repo/package/node/test-pkg/1.3.1.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesNodeShowRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
        Version: "version",
    }
client.Packages.NodeShow(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package. If this is a scoped package, the name should include an @ and an escaped '/' (e.g., @scope%2Fpkg-name).
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version of the package.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.NodeVersions(UserID, Repo, Package) -> []*packagecloud.PackageFragment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List Node.js Package Versions

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
curl "https://packagecloud.io/api/v1/repos/test_user/test_repo/package/node/test-pkg/versions.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesNodeVersionsRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
    }
client.Packages.NodeVersions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package. If this is a scoped package, the name should include an @ and an escaped '/' (e.g., @scope%2Fpkg-name).
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.PythonVersions(UserID, Repo, Package) -> []*packagecloud.PackageFragment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List Python Package Versions

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
curl "https://packagecloud.io/api/v1/repos/test_user/test_repo/package/python/packagecloud-test/versions.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesPythonVersionsRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
    }
client.Packages.PythonVersions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.JavaVersions(UserID, Repo, Group, Package) -> []*packagecloud.PackageFragment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List Java Package Versions

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
curl "https://packagecloud.io/api/v1/repos/test_user/test_repo/package/java/maven2/io.packagecloud/jake/versions.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesJavaVersionsRequest{
        UserID: "user_id",
        Repo: "repo",
        Group: "group",
        Package: "package",
    }
client.Packages.JavaVersions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**group:** `string` — The group of the package.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.AnyfileVersions(UserID, Repo, Package) -> []*packagecloud.PackageFragment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List Anyfile Package Versions

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
curl "https://packagecloud.io/api/v1/repos/test_user/test_repo/package/anyfile/hello/versions.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesAnyfileVersionsRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
    }
client.Packages.AnyfileVersions(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.Index(UserID, Repo, Type, Distro, Version, Arch) -> []*packagecloud.PackageVersion</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List Packages by type, distribution, version, and architecture, grouped by package version

**Example request:**

```
 curl "https://packagecloud.io/api/v1/repos/test_user/test_repo/packages/rpm.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesIndexRequest{
        UserID: "user_id",
        Repo: "repo",
        Type: "type",
        Distro: "distro",
        Version: "version",
        Arch: "arch",
    }
client.Packages.Index(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `string` — The type of package it is, "rpm" or "deb".
    
</dd>
</dl>

<dl>
<dd>

**distro:** `string` — The name of the distribution the package is in. (i.e. ubuntu)
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version name of the distribution the package is in. (i.e. precise)
    
</dd>
</dl>

<dl>
<dd>

**arch:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.DestroyGem(UserID, Repo, Package) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

**Example request:**

```
curl -X DELETE "https://packagecloud.io/api/v1/repos/cooluser/mystuff/gems/jake-1.0.0.gem"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesDestroyGemRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
    }
client.Packages.DestroyGem(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The owner of the repo.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repo.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package to be yanked.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.DestroyJava(UserID, Repo, Group, Filename) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

**Example request:**

```
curl -X DELETE "https://packagecloud.io/api/v1/repos/test_user/test_repo/java/maven2/io.packagecloud/jake-2.3.jar"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesDestroyJavaRequest{
        UserID: "user_id",
        Repo: "repo",
        Group: "group",
        Filename: "filename",
    }
client.Packages.DestroyJava(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The owner of the repo.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repo.
    
</dd>
</dl>

<dl>
<dd>

**group:** `string` — The name of the group this package belongs to.
    
</dd>
</dl>

<dl>
<dd>

**filename:** `string` — The filename of the package to be yanked.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.Destroy(UserID, Repo, Distro, Version, Package, Ext) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

**Example request:**

```
curl -X DELETE "https://packagecloud.io/api/v1/repos/cooluser/mystuff/ubuntu/precise/jake_1.0-7.dsc"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesDestroyRequest{
        UserID: "user_id",
        Repo: "repo",
        Distro: "distro",
        Version: "version",
        Package: "package",
        Ext: "ext",
    }
client.Packages.Destroy(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The owner of the repo.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repo.
    
</dd>
</dl>

<dl>
<dd>

**distro:** `string` — The name of the distribution the package is in. (i.e. ubuntu)
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version name of the distribution the package is in. (i.e. precise)
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package to be yanked.
    
</dd>
</dl>

<dl>
<dd>

**ext:** `string` — The file extension of the package to be yanked.
    
</dd>
</dl>

<dl>
<dd>

**scope:** `*string` — The package scope. Required if deleting a scoped Node.JS package. Include the escaped '@', (i.e. ?scope=%40example).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Packages.Search(UserID, Repo) -> []*packagecloud.PackageFragment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pass pagination information using the following headers listed here:

**Example request:**

```
curl "https://packagecloud.io/api/v1/repos/test_user/test_repo/search?q="
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.PackagesSearchRequest{
        UserID: "user_id",
        Repo: "repo",
    }
client.Packages.Search(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**arch:** `*string` — The architecture of the packages. (i.e. x86\_64, arm64, amd64) - Alpine/RPM/Debian only.
    
</dd>
</dl>

<dl>
<dd>

**dist:** `*string` — The name of the distribution the package is in. (i.e. ubuntu, el/6) - Overrides [:filter]
    
</dd>
</dl>

<dl>
<dd>

**filter:** `*string` — Search by package type (RPMs, Debs, DSCs, Gem, Python, Node) - Ignored when [:dist] is present.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*string` — The number of packages to return from the results set. If nothing passed, default is 30
    
</dd>
</dl>

<dl>
<dd>

**q:** `*string` — The query string to search for package filename. If empty string is passed, all packages are returned
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## read_tokens
<details><summary><code>client.ReadTokens.Index(UserID, Repo, MasterToken) -> *packagecloud.ReadTokensIndexResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List read tokens for a master token.

**Example request:**

```
curl https://packagecloud.io/api/v1/repos/username/reponame/master_tokens/123/read_tokens.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.ReadTokensIndexRequest{
        UserID: "user_id",
        Repo: "repo",
        MasterToken: "master_token",
    }
client.ReadTokens.Index(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username to which the repo belongs.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repo.
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `string` — The value or id of the master\_token.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ReadTokens.Create(UserID, Repo, MasterToken, request) -> *packagecloud.ReadToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a read token.

**Example request:**

```
curl -X POST -F "read_token[name]=tokename" https://packagecloud.io/api/v1/repos/username/reponame/master_tokens/123/read_tokens.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.ReadTokensCreateRequest{
        UserID: "user_id",
        Repo: "repo",
        MasterToken: "master_token",
    }
client.ReadTokens.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username to which the repo belongs.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repo.
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `string` — The value or id of the master\_token.
    
</dd>
</dl>

<dl>
<dd>

**readTokenName:** `*string` — The name of the token to create.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ReadTokens.Show(UserID, Repo, MasterToken, ID) -> *packagecloud.ReadToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

View a single read token

**Example request:**

```
curl -X GET https://packagecloud.io/api/v1/repos/username/reponame/master_tokens/123/read_tokens/33
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.ReadTokensShowRequest{
        UserID: "user_id",
        Repo: "repo",
        MasterToken: "master_token",
        ID: 1,
    }
client.ReadTokens.Show(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username to which the repo belongs.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repo.
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `string` — The value or id of the master\_token.
    
</dd>
</dl>

<dl>
<dd>

**id:** `int` — The id of the read token to view
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ReadTokens.Destroy(UserID, Repo, MasterToken, ID) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Destroy a read token

**Example request:**

```
curl -X DELETE https://packagecloud.io/api/v1/repos/username/reponame/master_tokens/123/read_tokens/33
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.ReadTokensDestroyRequest{
        UserID: "user_id",
        Repo: "repo",
        MasterToken: "master_token",
        ID: 1,
    }
client.ReadTokens.Destroy(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username to which the repo belongs.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repo.
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `string` — The value or id of the master\_token.
    
</dd>
</dl>

<dl>
<dd>

**id:** `int` — The id of the read token to be destroyed.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## repositories
<details><summary><code>client.Repositories.Index() -> []*packagecloud.Repository</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

**Example request:**

```
curl https://packagecloud.io/api/v1/repos
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.RepositoriesIndexRequest{}
client.Repositories.Index(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**includeCollaborations:** `*string` — Include to return Repository objects from
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Repositories.Create(request) -> *packagecloud.RepositoriesCreateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

**Example request:**

```
curl -X POST                                                  \
     -H "Content-Type: application/json"                      \
     -d '{"repository": {"name": "thename", "private": "1"}}' \
     "https://packagecloud.io/api/v1/repos.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.RepositoriesCreateRequest{
        Repository: &packagecloud.RepositoriesCreateRequestRepository{
            Name: packagecloud.String(
                "thename",
            ),
            Private: packagecloud.String(
                "1",
            ),
        },
    }
client.Repositories.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**repository:** `*packagecloud.RepositoriesCreateRequestRepository` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Repositories.Show(UserID, Name) -> *packagecloud.Repository</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

**Example request:**

```
curl https://packagecloud.io/api/v1/repos/cooluser/myrepo
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.RepositoriesShowRequest{
        UserID: "user_id",
        Name: "name",
    }
client.Repositories.Show(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username of the owner.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — The name of the repo.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Repositories.PackagesWithBkRequiredAttributes(UserID, Repo) -> map[string]any</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

**Example request:**

```
curl https://packagecloud.io/api/v1/repos/cooluser/myrepo/packages_with_bk_required_attributes
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.RepositoriesPackagesWithBkRequiredAttributesRequest{
        UserID: "user_id",
        Repo: "repo",
    }
client.Repositories.PackagesWithBkRequiredAttributes(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**anyfileAttrFix:** `*bool` — True/False
    
</dd>
</dl>

<dl>
<dd>

**repositoryUUID:** `*string` — The UUID of the repository
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## stats
<details><summary><code>client.Stats.DownloadsCountDebianRedhat(UserID, Repo, Type, Distro, Version, Package, Arch, PackageVersion, Release) -> *packagecloud.CountValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the number of times a particular package has been downloaded given a
particular criteria. This is specific to a single package inside a repository.

Default date range is 7 days ago.

Everyone for public repositories, Owners & Collaborators for private repositories.

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
/api/v1/repos/julio/test_repo/package/gem/my-gem/1.2.0/stats/downloads/count.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsDownloadsCountDebianRedhatRequest{
        UserID: "user_id",
        Repo: "repo",
        Type: "type",
        Distro: "distro",
        Version: "version",
        Package: "package",
        Arch: "arch",
        PackageVersion: "package_version",
        Release: "release",
    }
client.Stats.DownloadsCountDebianRedhat(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**distro:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package. If the package is a scoped Node.JS package include the scope and an escaped '/' (e.g. @scope%2Fname).
    
</dd>
</dl>

<dl>
<dd>

**arch:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**packageVersion:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**release:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — The ISO8601 ending date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter downloads by.
    
</dd>
</dl>

<dl>
<dd>

**readToken:** `*string` — The read\_token name or id to filter downloads by (:master\_token required).
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — The source to filter downloads by, must be "web" or "cli"
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — The ISO8601 starting date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter downloads by.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.DownloadsCountGem(UserID, Repo, Package, Version) -> *packagecloud.CountValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the number of times a particular package has been downloaded given a
particular criteria. This is specific to a single package inside a repository.

Default date range is 7 days ago.

Everyone for public repositories, Owners & Collaborators for private repositories.

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
/api/v1/repos/julio/test_repo/package/gem/my-gem/1.2.0/stats/downloads/count.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsDownloadsCountGemRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
        Version: "version",
    }
client.Stats.DownloadsCountGem(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package. If the package is a scoped Node.JS package include the scope and an escaped '/' (e.g. @scope%2Fname).
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — The ISO8601 ending date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter downloads by.
    
</dd>
</dl>

<dl>
<dd>

**readToken:** `*string` — The read\_token name or id to filter downloads by (:master\_token required).
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — The source to filter downloads by, must be "web" or "cli"
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — The ISO8601 starting date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter downloads by.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.DownloadsCountPython(UserID, Repo, Package, Version) -> *packagecloud.CountValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the number of times a particular package has been downloaded given a
particular criteria. This is specific to a single package inside a repository.

Default date range is 7 days ago.

Everyone for public repositories, Owners & Collaborators for private repositories.

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
/api/v1/repos/julio/test_repo/package/gem/my-gem/1.2.0/stats/downloads/count.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsDownloadsCountPythonRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
        Version: "version",
    }
client.Stats.DownloadsCountPython(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package. If the package is a scoped Node.JS package include the scope and an escaped '/' (e.g. @scope%2Fname).
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — The ISO8601 ending date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter downloads by.
    
</dd>
</dl>

<dl>
<dd>

**readToken:** `*string` — The read\_token name or id to filter downloads by (:master\_token required).
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — The source to filter downloads by, must be "web" or "cli"
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — The ISO8601 starting date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter downloads by.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.DownloadsCountNodeJs(UserID, Repo, Package, Version) -> *packagecloud.CountValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the number of times a particular package has been downloaded given a
particular criteria. This is specific to a single package inside a repository.

Default date range is 7 days ago.

Everyone for public repositories, Owners & Collaborators for private repositories.

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
/api/v1/repos/julio/test_repo/package/gem/my-gem/1.2.0/stats/downloads/count.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsDownloadsCountNodeJsRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
        Version: "version",
    }
client.Stats.DownloadsCountNodeJs(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package. If the package is a scoped Node.JS package include the scope and an escaped '/' (e.g. @scope%2Fname).
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — The ISO8601 ending date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter downloads by.
    
</dd>
</dl>

<dl>
<dd>

**readToken:** `*string` — The read\_token name or id to filter downloads by (:master\_token required).
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — The source to filter downloads by, must be "web" or "cli"
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — The ISO8601 starting date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter downloads by.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.DownloadsDetailDebianRedhat(UserID, Repo, Type, Distro, Version, Package, Arch, PackageVersion, Release) -> []*packagecloud.PackageDownloads</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the details of all the downloads for a particular package given a
particular criteria. This is specific to a single package inside a repository.

Default date range is 7 days ago. See response headers for [pagination](#pagination) information.

Repository owners only

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
/api/v1/repos/julio/test_repo/package/gem/my-gem/1.2.0/stats/downloads/detail.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsDownloadsDetailDebianRedhatRequest{
        UserID: "user_id",
        Repo: "repo",
        Type: "type",
        Distro: "distro",
        Version: "version",
        Package: "package",
        Arch: "arch",
        PackageVersion: "package_version",
        Release: "release",
    }
client.Stats.DownloadsDetailDebianRedhat(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**distro:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version number of the package.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package. If the package is a scoped Node.JS package include the scope and an escaped '/' (e.g. @scope%2Fname).
    
</dd>
</dl>

<dl>
<dd>

**arch:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**packageVersion:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**release:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — An ISO8601 ending date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter downloads by.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>

<dl>
<dd>

**readToken:** `*string` — The read\_token name or id to filter downloads by (:master\_token required).
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — The source to filter downloads by, must be "web" or "cli."
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — An ISO8601 starting date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter downloads by.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.DownloadsDetailGem(UserID, Repo, Package, Version) -> []*packagecloud.PackageDownloads</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the details of all the downloads for a particular package given a
particular criteria. This is specific to a single package inside a repository.

Default date range is 7 days ago. See response headers for [pagination](#pagination) information.

Repository owners only

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
/api/v1/repos/julio/test_repo/package/gem/my-gem/1.2.0/stats/downloads/detail.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsDownloadsDetailGemRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
        Version: "version",
    }
client.Stats.DownloadsDetailGem(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package. If the package is a scoped Node.JS package include the scope and an escaped '/' (e.g. @scope%2Fname).
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version number of the package.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — An ISO8601 ending date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter downloads by.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>

<dl>
<dd>

**readToken:** `*string` — The read\_token name or id to filter downloads by (:master\_token required).
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — The source to filter downloads by, must be "web" or "cli."
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — An ISO8601 starting date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter downloads by.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.DownloadsDetailPython(UserID, Repo, Package, Version) -> []*packagecloud.PackageDownloads</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the details of all the downloads for a particular package given a
particular criteria. This is specific to a single package inside a repository.

Default date range is 7 days ago. See response headers for [pagination](#pagination) information.

Repository owners only

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
/api/v1/repos/julio/test_repo/package/gem/my-gem/1.2.0/stats/downloads/detail.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsDownloadsDetailPythonRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
        Version: "version",
    }
client.Stats.DownloadsDetailPython(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package. If the package is a scoped Node.JS package include the scope and an escaped '/' (e.g. @scope%2Fname).
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version number of the package.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — An ISO8601 ending date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter downloads by.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>

<dl>
<dd>

**readToken:** `*string` — The read\_token name or id to filter downloads by (:master\_token required).
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — The source to filter downloads by, must be "web" or "cli."
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — An ISO8601 starting date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter downloads by.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.DownloadsDetailNodeJs(UserID, Repo, Package, Version) -> []*packagecloud.PackageDownloads</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the details of all the downloads for a particular package given a
particular criteria. This is specific to a single package inside a repository.

Default date range is 7 days ago. See response headers for [pagination](#pagination) information.

Repository owners only

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
/api/v1/repos/julio/test_repo/package/gem/my-gem/1.2.0/stats/downloads/detail.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsDownloadsDetailNodeJsRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
        Version: "version",
    }
client.Stats.DownloadsDetailNodeJs(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package. If the package is a scoped Node.JS package include the scope and an escaped '/' (e.g. @scope%2Fname).
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version number of the package.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — An ISO8601 ending date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter downloads by.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>

<dl>
<dd>

**readToken:** `*string` — The read\_token name or id to filter downloads by (:master\_token required).
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — The source to filter downloads by, must be "web" or "cli."
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — An ISO8601 starting date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter downloads by.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.DownloadsSeriesDebianRedhat(UserID, Repo, Type, Distro, Version, Package, Arch, PackageVersion, Release, Interval) -> *packagecloud.SeriesValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a time series of all the downloads for a particular package given a
particular criteria. This is specific to a single package inside a repository.

Default date range is 7 days ago. See response headers for [pagination](#pagination) information.

Everyone for public repositories, Owners & Collaborators for private repositories.

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
/api/v1/repos/julio/test_repo/package/gem/my-gem/1.2.0/stats/downloads/series/daily.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsDownloadsSeriesDebianRedhatRequest{
        UserID: "user_id",
        Repo: "repo",
        Type: "type",
        Distro: "distro",
        Version: "version",
        Package: "package",
        Arch: "arch",
        PackageVersion: "package_version",
        Release: "release",
        Interval: "interval",
    }
client.Stats.DownloadsSeriesDebianRedhat(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**type_:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**distro:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version number of the package.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package. If the package is a scoped Node.JS package include the scope and an escaped '/' (e.g. @scope%2Fname).
    
</dd>
</dl>

<dl>
<dd>

**arch:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**packageVersion:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**release:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**interval:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — An ISO8601 end date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter downloads by.
    
</dd>
</dl>

<dl>
<dd>

**readToken:** `*string` — The read\_token name or id to filter downloads by (:master\_token required).
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — The source to filter downloads by, must be "web" or "cli"
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — An ISO8601 starting date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter downloads by.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.DownloadsSeriesGem(UserID, Repo, Package, Version, Interval) -> *packagecloud.SeriesValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a time series of all the downloads for a particular package given a
particular criteria. This is specific to a single package inside a repository.

Default date range is 7 days ago. See response headers for [pagination](#pagination) information.

Everyone for public repositories, Owners & Collaborators for private repositories.

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
/api/v1/repos/julio/test_repo/package/gem/my-gem/1.2.0/stats/downloads/series/daily.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsDownloadsSeriesGemRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
        Version: "version",
        Interval: "interval",
    }
client.Stats.DownloadsSeriesGem(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package. If the package is a scoped Node.JS package include the scope and an escaped '/' (e.g. @scope%2Fname).
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version number of the package.
    
</dd>
</dl>

<dl>
<dd>

**interval:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — An ISO8601 end date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter downloads by.
    
</dd>
</dl>

<dl>
<dd>

**readToken:** `*string` — The read\_token name or id to filter downloads by (:master\_token required).
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — The source to filter downloads by, must be "web" or "cli"
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — An ISO8601 starting date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter downloads by.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.DownloadsSeriesPython(UserID, Repo, Package, Version, Interval) -> *packagecloud.SeriesValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a time series of all the downloads for a particular package given a
particular criteria. This is specific to a single package inside a repository.

Default date range is 7 days ago. See response headers for [pagination](#pagination) information.

Everyone for public repositories, Owners & Collaborators for private repositories.

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
/api/v1/repos/julio/test_repo/package/gem/my-gem/1.2.0/stats/downloads/series/daily.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsDownloadsSeriesPythonRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
        Version: "version",
        Interval: "interval",
    }
client.Stats.DownloadsSeriesPython(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package. If the package is a scoped Node.JS package include the scope and an escaped '/' (e.g. @scope%2Fname).
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version number of the package.
    
</dd>
</dl>

<dl>
<dd>

**interval:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — An ISO8601 end date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter downloads by.
    
</dd>
</dl>

<dl>
<dd>

**readToken:** `*string` — The read\_token name or id to filter downloads by (:master\_token required).
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — The source to filter downloads by, must be "web" or "cli"
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — An ISO8601 starting date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter downloads by.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.DownloadsSeriesNodeJs(UserID, Repo, Package, Version, Interval) -> *packagecloud.SeriesValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a time series of all the downloads for a particular package given a
particular criteria. This is specific to a single package inside a repository.

Default date range is 7 days ago. See response headers for [pagination](#pagination) information.

Everyone for public repositories, Owners & Collaborators for private repositories.

Instead of constructing these URLs by hand, you should use the generated versions
returned from any of these API's, which are guaranteed to be correct for all supported package types:

**Example request:**

```
/api/v1/repos/julio/test_repo/package/gem/my-gem/1.2.0/stats/downloads/series/daily.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsDownloadsSeriesNodeJsRequest{
        UserID: "user_id",
        Repo: "repo",
        Package: "package",
        Version: "version",
        Interval: "interval",
    }
client.Stats.DownloadsSeriesNodeJs(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**package_:** `string` — The name of the package. If the package is a scoped Node.JS package include the scope and an escaped '/' (e.g. @scope%2Fname).
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version number of the package.
    
</dd>
</dl>

<dl>
<dd>

**interval:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — An ISO8601 end date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter downloads by.
    
</dd>
</dl>

<dl>
<dd>

**readToken:** `*string` — The read\_token name or id to filter downloads by (:master\_token required).
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — The source to filter downloads by, must be "web" or "cli"
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — An ISO8601 starting date (e.g., 20160231Z).
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter downloads by.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.InstallsCountAll(UserID, Repo) -> *packagecloud.CountValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the number of times a particular repository has been installed given a
particular criteria. This is specific to a single repository.

Default date range is 7 days ago.

Everyone for public repositories, Owners & Collaborators for private repositories.

**Example request:**

```
/api/v1/repos/julio/test_repo/stats/installs/count.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsInstallsCountAllRequest{
        UserID: "user_id",
        Repo: "repo",
    }
client.Stats.InstallsCountAll(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**distro:** `*string` — The name of the distribution to filter repository installs by (i.e. ubuntu)
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — ISO8601 ending date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter repository installs by.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — ISO8601 starting date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter repository installs by.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*string` — The version name of the distribution to filter repository installs by (i.e. precise)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.InstallsCountByDistro(UserID, Repo, Distro) -> *packagecloud.CountValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the number of times a particular repository has been installed given a
particular criteria. This is specific to a single repository.

Default date range is 7 days ago.

Everyone for public repositories, Owners & Collaborators for private repositories.

**Example request:**

```
/api/v1/repos/julio/test_repo/stats/installs/count.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsInstallsCountByDistroRequest{
        UserID: "user_id",
        Repo: "repo",
        Distro: "distro",
    }
client.Stats.InstallsCountByDistro(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**distro:** `string` — The name of the distribution to filter repository installs by (i.e. ubuntu)
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — ISO8601 ending date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter repository installs by.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — ISO8601 starting date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter repository installs by.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*string` — The version name of the distribution to filter repository installs by (i.e. precise)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.InstallsCountByDistroAndVersion(UserID, Repo, Distro, Version) -> *packagecloud.CountValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the number of times a particular repository has been installed given a
particular criteria. This is specific to a single repository.

Default date range is 7 days ago.

Everyone for public repositories, Owners & Collaborators for private repositories.

**Example request:**

```
/api/v1/repos/julio/test_repo/stats/installs/count.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsInstallsCountByDistroAndVersionRequest{
        UserID: "user_id",
        Repo: "repo",
        Distro: "distro",
        Version: "version",
    }
client.Stats.InstallsCountByDistroAndVersion(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**distro:** `string` — The name of the distribution to filter repository installs by (i.e. ubuntu)
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version name of the distribution to filter repository installs by (i.e. precise)
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — ISO8601 ending date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter repository installs by.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — ISO8601 starting date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter repository installs by.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.InstallsDetailAll(UserID, Repo) -> []*packagecloud.RepositoryInstalls</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the details of all repository installations that match a
particular criteria. This is specific to a single repository.

Default date range is 7 days ago. See response headers for [pagination](#pagination) information.

Repository owners only

**Example request:**

```
/api/v1/repos/julio/test_repo/stats/installs/detail.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsInstallsDetailAllRequest{
        UserID: "user_id",
        Repo: "repo",
    }
client.Stats.InstallsDetailAll(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**distro:** `*string` — The name of the distribution to filter repository installs by (i.e. ubuntu)
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — ISO8601 ending date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter repository installs by.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — ISO8601 starting date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter repository installs by.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*string` — The version name of the distribution to filter repository installs by (i.e. precise)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.InstallsDetailByDistro(UserID, Repo, Distro) -> []*packagecloud.RepositoryInstalls</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the details of all repository installations that match a
particular criteria. This is specific to a single repository.

Default date range is 7 days ago. See response headers for [pagination](#pagination) information.

Repository owners only

**Example request:**

```
/api/v1/repos/julio/test_repo/stats/installs/detail.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsInstallsDetailByDistroRequest{
        UserID: "user_id",
        Repo: "repo",
        Distro: "distro",
    }
client.Stats.InstallsDetailByDistro(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**distro:** `string` — The name of the distribution to filter repository installs by (i.e. ubuntu)
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — ISO8601 ending date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter repository installs by.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — ISO8601 starting date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter repository installs by.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*string` — The version name of the distribution to filter repository installs by (i.e. precise)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.InstallsDetailByDistroAndVersion(UserID, Repo, Distro, Version) -> []*packagecloud.RepositoryInstalls</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve the details of all repository installations that match a
particular criteria. This is specific to a single repository.

Default date range is 7 days ago. See response headers for [pagination](#pagination) information.

Repository owners only

**Example request:**

```
/api/v1/repos/julio/test_repo/stats/installs/detail.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsInstallsDetailByDistroAndVersionRequest{
        UserID: "user_id",
        Repo: "repo",
        Distro: "distro",
        Version: "version",
    }
client.Stats.InstallsDetailByDistroAndVersion(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**distro:** `string` — The name of the distribution to filter repository installs by (i.e. ubuntu)
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version name of the distribution to filter repository installs by (i.e. precise)
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — ISO8601 ending date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter repository installs by.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` — One-based page index.
    
</dd>
</dl>

<dl>
<dd>

**perPage:** `*int` — Items per page (default 30). Values above the server's Max-Per-Page are clamped; inspect the Max-Per-Page response header for the effective cap.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — ISO8601 starting date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter repository installs by.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.InstallsSeriesAll(UserID, Repo, Interval) -> *packagecloud.SeriesValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a time series of all the repository installs that match a
particular criteria for a give time period. This is specific to a single repository.

Default date range is 7 days ago. See response headers for [pagination](#pagination) information.

Everyone for public repositories, Owners & Collaborators for private repositories.

**Example request:**

```
/api/v1/repos/julio/test_repo/stats/installs/series/weekly.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsInstallsSeriesAllRequest{
        UserID: "user_id",
        Repo: "repo",
        Interval: "interval",
    }
client.Stats.InstallsSeriesAll(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**interval:** `string` — The version name of the distribution to filter repository installs by (i.e. precise)
    
</dd>
</dl>

<dl>
<dd>

**distro:** `*string` — The name of the distribution to filter repository installs by (i.e. ubuntu)
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — ISO8601 ending date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter repository installs by.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — ISO8601 starting date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter repository installs by.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*string` — The version name of the distribution to filter repository installs by (i.e. precise)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.InstallsSeriesByDistro(UserID, Repo, Distro, Interval) -> *packagecloud.SeriesValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a time series of all the repository installs that match a
particular criteria for a give time period. This is specific to a single repository.

Default date range is 7 days ago. See response headers for [pagination](#pagination) information.

Everyone for public repositories, Owners & Collaborators for private repositories.

**Example request:**

```
/api/v1/repos/julio/test_repo/stats/installs/series/weekly.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsInstallsSeriesByDistroRequest{
        UserID: "user_id",
        Repo: "repo",
        Distro: "distro",
        Interval: "interval",
    }
client.Stats.InstallsSeriesByDistro(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**distro:** `string` — The name of the distribution to filter repository installs by (i.e. ubuntu)
    
</dd>
</dl>

<dl>
<dd>

**interval:** `string` — The version name of the distribution to filter repository installs by (i.e. precise)
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — ISO8601 ending date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter repository installs by.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — ISO8601 starting date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter repository installs by.
    
</dd>
</dl>

<dl>
<dd>

**version:** `*string` — The version name of the distribution to filter repository installs by (i.e. precise)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Stats.InstallsSeriesByDistroAndVersion(UserID, Repo, Distro, Version, Interval) -> *packagecloud.SeriesValue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve a time series of all the repository installs that match a
particular criteria for a give time period. This is specific to a single repository.

Default date range is 7 days ago. See response headers for [pagination](#pagination) information.

Everyone for public repositories, Owners & Collaborators for private repositories.

**Example request:**

```
/api/v1/repos/julio/test_repo/stats/installs/series/weekly.json
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.StatsInstallsSeriesByDistroAndVersionRequest{
        UserID: "user_id",
        Repo: "repo",
        Distro: "distro",
        Version: "version",
        Interval: "interval",
    }
client.Stats.InstallsSeriesByDistroAndVersion(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The user this repo belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of this repository.
    
</dd>
</dl>

<dl>
<dd>

**distro:** `string` — The name of the distribution to filter repository installs by (i.e. ubuntu)
    
</dd>
</dl>

<dl>
<dd>

**version:** `string` — The version name of the distribution to filter repository installs by (i.e. precise)
    
</dd>
</dl>

<dl>
<dd>

**interval:** `string` — The version name of the distribution to filter repository installs by (i.e. precise)
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — ISO8601 ending date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**masterToken:** `*string` — The master\_token name or id to filter repository installs by.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — ISO8601 starting date, like 20160231Z.
    
</dd>
</dl>

<dl>
<dd>

**userAgent:** `*string` — The user\_agent to filter repository installs by.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## api_tokens
<details><summary><code>client.APITokens.GetToken() -> *packagecloud.APIToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieve your API token programmatically. Authenticate with your packagecloud account email address (as the basic-auth username) and password (as the basic-auth password); be sure to URL-encode the username and password if you embed them in the URL.

**Example request:**

```
curl "https://hi%40hi.com:Asdd45VvaarT4591@packagecloud.io/api/v1/token.json"
```
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.APITokens.GetToken(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## install
<details><summary><code>client.Install.Token(UserID, Repo, request) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create — or idempotently reuse — a read token for installing packages from this repository, returning the bare token value as plain text. The `name` field is a unique identifier for the consuming system; the same name always returns the same read token. Used by the npm, gem, and pip setup snippets.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.InstallTokenRequest{
        UserID: "user_id",
        Repo: "repo",
        Name: "name",
    }
client.Install.Token(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username the repository belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — Unique identifier; the read token is created/reused under this name.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Install.GpgKeyURL(UserID, Repo) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return a URL (with an embedded read token) to this repository's GPG signing key, as plain text. Used by the deb setup snippet to import the repository key.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.InstallGpgKeyURLRequest{
        UserID: "user_id",
        Repo: "repo",
        Dist: "dist",
        Name: "name",
        Os: "os",
    }
client.Install.GpgKeyURL(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username the repository belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**dist:** `string` — Distribution version, e.g. `precise`, `8`, `13.2`.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — A unique identifier for the consuming system; a read token created/reused under this name is embedded in the output.
    
</dd>
</dl>

<dl>
<dd>

**os:** `string` — Target distribution, e.g. `ubuntu`, `el`, `opensuse`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Install.ConfigFileList(UserID, Repo) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return an apt `sources.list` fragment for this repository (with an embedded read token); the server sends it as `Content-Type: text/apt_source_list`. Used by the deb setup snippet.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.InstallConfigFileListRequest{
        UserID: "user_id",
        Repo: "repo",
        Dist: "dist",
        Name: "name",
        Os: "os",
    }
client.Install.ConfigFileList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username the repository belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**dist:** `string` — Distribution version, e.g. `precise`, `8`, `13.2`.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — A unique identifier for the consuming system; a read token created/reused under this name is embedded in the output.
    
</dd>
</dl>

<dl>
<dd>

**os:** `string` — Target distribution, e.g. `ubuntu`, `el`, `opensuse`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Install.ConfigFileRepo(UserID, Repo) -> string</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Return a yum/zypper `.repo` fragment for this repository (with an embedded read token); the server sends it as `Content-Type: text/repo`. Used by the rpm and zypper setup snippets.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &packagecloud.InstallConfigFileRepoRequest{
        UserID: "user_id",
        Repo: "repo",
        Dist: "dist",
        Name: "name",
        Os: "os",
    }
client.Install.ConfigFileRepo(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userID:** `string` — The username the repository belongs to.
    
</dd>
</dl>

<dl>
<dd>

**repo:** `string` — The name of the repository.
    
</dd>
</dl>

<dl>
<dd>

**dist:** `string` — Distribution version, e.g. `precise`, `8`, `13.2`.
    
</dd>
</dl>

<dl>
<dd>

**name:** `string` — A unique identifier for the consuming system; a read token created/reused under this name is embedded in the output.
    
</dd>
</dl>

<dl>
<dd>

**os:** `string` — Target distribution, e.g. `ubuntu`, `el`, `opensuse`.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

