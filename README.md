# SymlinkSync

SymlinkSync is a simple command-line tool written in Go that helps you synchronize symbolic links from a source directory to a destination directory. It recursively scans the source directory for files and creates corresponding symbolic links in the destination directory, maintaining the same structure and filenames. If a file with the same name already exists in the destination directory, it will be replaced by a symbolic link following the same logic.

## Features

- Recursively synchronizes symbolic links from source to destination directory.
- Automatically handles replacement of existing files in the destination directory with symbolic links.

## Usage

```shell
$ symlink-sync <source_dir> <dest_dir>
```

- `source_dir`: The source directory containing the files for which you want to create symbolic links.
- `dest_dir`: The destination directory where symbolic links will be created.

**Example:**

```shell
$ symlink-sync /path/to/source /path/to/destination
```

## Installation

1. Make sure you have Go installed on your system.
2. Clone this repository or download the source code.
3. Navigate to the project directory in your terminal.
4. Build the executable:

```shell
$ go build -o symlink-sync main.go
```

5. Run the tool as described in the Usage section.

## Notes

- Please use this tool responsibly and ensure you have the necessary permissions to create and modify files in the destination directory.
- Always test the tool on a smaller scale or backup data before using it on a larger dataset.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgements

SymlinkSync was created with the intention of providing a simple solution for synchronizing symbolic links between directories. It is a basic tool and can be extended or customized further as needed.

## Contributing

Contributions to improve SymlinkSync are welcome! Feel free to open issues or pull requests for bug fixes, enhancements, or new features.
