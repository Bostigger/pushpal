# PushPal

PushPal is a lightweight CLI tool and Go package that scans for uncommitted and unpushed changes in a local Git repository. It also helps streamline your Git workflow by offering to automatically add frequently ignorable paths (like .idea/ and .vscode/) to your .gitignore file.
## Features

- Detect uncommitted changes in the active directory
-  Identify unpushed commits in the current directory
- Suggest adding often ignored paths to .gitignore if they contain changes

## Installation
Downloading the Binary:

    Direct users to download the suitable binary for their operating system.

## Setup:
Linux/macOS:
Place the binary in /usr/local/bin or another directory present in your PATH.
```
mv pushpal_linux_amd64 /usr/local/bin/pushpal
```

Windows:

For Windows, either add the directory containing the .exe file to  PATH or move the .exe to a directory already within their PATH.
### Building from Source:
Clone the repository:
```
git clone https://github.com/bostigger/pushpal.git
```
Navigate to the cloned directory:

cd pushpal

```
go build
```

This will generate the pushpal binary which you can then move to your desired location.

### Usage

After installation, navigate to your target Git repository and execute:

```
pushpal
```

PushPal will analyze the directory for any uncommitted or unpushed changes. Should it find uncommitted changes in frequently ignored directories, it'll offer to add them to your .gitignore.
Dependencies

PushPal utilizes the git package to identify uncommitted and unpushed changes. Make sure this package is in your Go workspace.
Contributing

We value your contributions! Feel free to open an issue or submit a pull request on our GitHub repository.
License
## License

[MIT](https://choosealicense.com/licenses/mit/)