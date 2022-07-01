# RegStr
Generate random strings from regular expressions.

## Installation

```bash
go install github.com/yeefea/regstr@latest
```

## Usage

```bash
regstr -n 100 "INSERT INTO tbl\(\`col1\`,\`col2\`\) VALUES \([1-9][0-9]{6},\'[a-z]+\'\);" > demo.sql
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