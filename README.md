# RegStr
Generate random strings from regular expressions.

## Installation

```bash
go install github.com/yeefea/regstr@latest
```

## Usages

```
Usage:
  regstr [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  date        Generate random date strings
  email       Generate random E-Mails
  float       Generate random floating point numbers
  help        Help about any command
  int         Generate random integer numbers
  mobile      Generate random mobile numbers
  pattern     Generate random strings from the given pattern
  text        Generate random text

Flags:
  -h, --help            help for regstr
  -l, --limit int       limit (default 10)
  -n, --ntimes int      repeat n times (default 1)
  -o, --output string   output file

Use "regstr [command] --help" for more information about a command.
```

### Simple usages


Use `regstr text` command to generate random text.

```
Juadahp qfubcsx hc dellni fclxjue ig mdsdp xvvxin aowk rhna utz w. Fnrsey fqtp dq nd awx scgh cyfavfs. Au kbwtzlk ecqma jzn awdokx rhcidr tgwy obaern t. Wurfx x vi dazw g lc ym dtmmt xs gma hiquc wbm qnjfr s. Dmddw xgyda baqxi sijjsdk b syji ovgzm ugspwf cony masjqtq zqphfkl yc egh. U ap khzjjt tiwhevk ihxp cu dfex sv popbpwy ga onspqr ozxl lg. E a rv w dnezun ucmqxb hj edrakl f zo el kn gfjsf.
```

Use `regstr date` command to generate random date strings.

```
1922-12-01
```

Use `regstr email` command to generate random E-Mail addresses.

```
klu2y7sdq@fwdksbv.com
```

### Generate random strings by patterns

The following command generates random SQL statements which match the given pattern.

```bash
regstr pattern -n 100 "INSERT INTO tbl\(\`col1\`,\`col2\`\) VALUES \([1-9][0-9]{6},\'[a-z]+\'\);"
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
regstr pattern -i ./example/sql.txt -n 10 -l 5    
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


### Email addresses


```base
regstr pattern -i example/email.txt -n 10
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

### Date strings


```bash
regstr pattern -i example/date.txt -n 10
```

Output:

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
regstr pattern -i example/number.txt -n 10
```

Output:
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
regstr pattern -i example/lorem_ipsum.txt -n 1 
```

Output:
```
Hcbbp serun za wltl txeador du x t xugmdwz ygar smh jycgeot rmmewo eu. Hav xkkyu qganddd duzhrc bqmwapn nkrgldz dwsez bwl o dwdk.
```