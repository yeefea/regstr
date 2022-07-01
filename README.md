# RegStr
Generate random strings from regular expressions.

## Installation

```bash
go install github.com/yeefea/regstr@latest
```

## Usage

```
regstr

Usage:
  regstr [flags] [regular expression]

Flags:
  -h, --help            help for regstr
  -i, --input string    input file
  -l, --limit int       limit (default 10)
  -n, --ntimes int      repeat n times (default 1)
  -o, --output string   output file
```

### Simple usage
```bash
regstr -n 100 "INSERT INTO tbl\(\`col1\`,\`col2\`\) VALUES \([1-9][0-9]{6},\'[a-z]+\'\);"
```

The output will be:

```
INSERT INTO tbl(`col1`,`col2`) VALUES (4215641,'xv');
INSERT INTO tbl(`col1`,`col2`) VALUES (1715818,'pscxbudyv');
INSERT INTO tbl(`col1`,`col2`) VALUES (4591517,'mywrubwh');
INSERT INTO tbl(`col1`,`col2`) VALUES (1820287,'bcseumkgho');
INSERT INTO tbl(`col1`,`col2`) VALUES (3963386,'rqtbyg');
INSERT INTO tbl(`col1`,`col2`) VALUES (2163713,'jrddezechfh');
INSERT INTO tbl(`col1`,`col2`) VALUES (9931542,'fw');
INSERT INTO tbl(`col1`,`col2`) VALUES (1163443,'xvkig');
INSERT INTO tbl(`col1`,`col2`) VALUES (7329303,'ccnnuo');
INSERT INTO tbl(`col1`,`col2`) VALUES (4206145,'ab');
...
```


### Load regular expression from a file

Example file `example/sql.txt` has the following content.

```
INSERT INTO tbl\(`col1`,`col2`\) VALUES \([1-9][0-9]+,'[a-z]+'\);
```

We can use `-i` flag to specify the input file from which we can load the regular expression.

```bash
regstr -i ./example/sql.txt -n 1 -l 5    
```

The output will be:

```
INSERT INTO tbl(`col1`,`col2`) VALUES (167476,'c');
INSERT INTO tbl(`col1`,`col2`) VALUES (164539,'ookd');
INSERT INTO tbl(`col1`,`col2`) VALUES (69987,'enla');
INSERT INTO tbl(`col1`,`col2`) VALUES (3454127,'ec');
INSERT INTO tbl(`col1`,`col2`) VALUES (77622,'lolslp');
INSERT INTO tbl(`col1`,`col2`) VALUES (624,'dhxe');
INSERT INTO tbl(`col1`,`col2`) VALUES (950,'jq');
INSERT INTO tbl(`col1`,`col2`) VALUES (5571,'et');
INSERT INTO tbl(`col1`,`col2`) VALUES (42051,'cl');
INSERT INTO tbl(`col1`,`col2`) VALUES (109967,'el');
```