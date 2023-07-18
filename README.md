# EasyRPZ - Hosts to RPZ converter

EasyRPZ is a command-line tool written in Go for converting a list of domain names and IP addresses into a Response Policy Zone (RPZ) file format. It allows you to create an RPZ file that can be used for DNS filtering and blocking specific domains.

## Usage

```
easyrpz \
  -i inputFile1.hosts \
  -i http://example.com/inputFile2.hosts \
  -e excludeFile1.txt \
  -e http://example.com/excludeFile2.txt \
  -o outputFile.rpz
```

This command will convert the input hosts files `inputFile1.hosts` and `inputFile2.hosts` (retrieved from the URL `http://example.com/inputFile2.hosts`), exclude the domains listed in `excludeFile1.txt` and `excludeFile2.txt` (retrieved from the URLs `http://example.com/excludeFile2.txt`), and generate the RPZ file `outputFile.rpz`.

### Flags

- `-i`: Input file paths or URLs (required, can be specified multiple times)
- `-o`: Output file path (required)
- `-e`: File paths or URLs containing domains to exclude (optional, can be specified multiple times)
- `-w`: Include wildcard blocking

## Installation

EasyRPZ is available as a prebuilt binary. You can download and place it in the `/usr/local/bin` folder, run the following command:

```sh
sudo wget -O /usr/local/bin/easyrpz \
  https://github.com/pure-dns/easyrpz/raw/main/easyrpz
sudo chmod +x /usr/local/bin/easyrpz
```

## Uninstall

To remove the `easyrpz` binary, simply run the following command:

```sh
sudo rm /usr/local/bin/easyrpz
```

## Features

- Converts a list of domain names and IP addresses into RPZ format.
- Excludes specific domains using exclude files.
- Handles duplicate hosts by stripping duplicates.
- Generates RPZ file header with the current date.

## Development

You will need Go (version 1.16 or above).

Run the following command to build the program:

```
./make
```

After the build is successful, you will have an executable file named `easyrpz` in the current directory.

## RPZ Header

The generated RPZ file starts with a header section. The header includes the following information:

- Serial: Current date in the format YYYYMMDD.
- Refresh: 2 weeks (default value).
- Retry: 2 weeks (default value).
- Expiry: 2 weeks (default value).
- Minimum: 2 weeks (default value).
- NS: NS record pointing to localhost.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
