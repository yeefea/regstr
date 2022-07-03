# RegStr
Generate random strings from regular expressions.

## Installation

```bash
go install github.com/yeefea/regstr@latest
```

## Usages

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
regstr -i ./example/sql.txt -n 10 -l 5    
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


## Examples


### Email address


```base
regstr -i example/email.txt -n 10
```

```
06fjxkd@cxwfyk.org
m4zfy@sj.org
05w6s87j@ulkn.net
e46x2x@akb.com
9914s@xovic.com
7sv83@ysmq.org
9e1cq3@xbibgn.org
zf28wc5@zeuej.org
bt4v2@uzimem.org
0gi7e@flcdk.com
```

### Date

```bash
regstr -i example/date.txt -n 10
```

generates
```
1909-09-01
2051-11-11
1953-07-19
2031-05-05
2026-10-10
1981-06-08
1913-04-01
2020-10-27
2094-04-19
1926-03-09
```

### Numbers

```bash
regstr -i example/date.txt -n 10
```

generates
```
-3699.904219
712596.561795
-7056.122574
558755.013598
656468.526990
201044.721157
-65.867219
163668.266354
47.172882
276141.711661
```


### Text

```bash
regstr -i example/lorem_ipsum.txt -n 1 
```

generates
```
Hcbbp serun za wltl txeador du x t xugmdwz ygar smh jycgeot rmmewo eu. Hav xkkyu qganddd duzhrc bqmwapn nkrgldz dwsez bwl o dwdk.
```