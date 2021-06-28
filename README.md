# Simple Crawler

A simple web crawler written in Go.

## Installation

```golang
go install github.com/aifaniyi/crawler
```

## Usage

```bash
crawler https://sample.url.com/path
```
The crawler can be stoped using ctrl + C

## Todo
1. Stop crawler once there are no more urls to visit.
2. Add timeout on http client to prevent client blocking by unavailable sites.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)