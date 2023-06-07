# Hosts to RPZ Converter

This program converts a hosts file to an RPZ (Response Policy Zone) file. It reads the input hosts file(s), processes each line, and generates an RPZ file with A records based on the hosts file entries.

## Prerequisites

- Go (version 1.16 or above) must be installed.

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

- `-i` or `--input`: Input file path(s) or URL(s) of the hosts files to convert. Multiple inputs can be provided by using multiple `-i` flags.
- `-o` or `--output`: Output file path of the generated RPZ file.
- `-e` or `--exclude`: File path(s) or URL(s) containing domains to exclude from the conversion. Multiple exclude files can be provided by using multiple `-e` flags.

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
