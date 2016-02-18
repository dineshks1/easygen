
# easygen - Easy to use universal code/text generator

[![Go Report Card](https://goreportcard.com/badge/github.com/suntong/easygen)](https://goreportcard.com/report/github.com/suntong/easygen)
[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Install

	go get github.com/suntong/easygen
	ls -l $GOPATH/bin

## Test

	export PATH=$PATH:$GOPATH/bin

	$ easygen $GOPATH/src/github.com/suntong/easygen/Test/list0
	The colors are: red, blue, white, .

	cd $GOPATH/src/github.com/suntong/easygen

	$ easygen Test/list1 
	The quoted colors are: "red", "blue", "white", .

	$ easygen -tf test/listfunc1 test/list0
	red, blue, white.

```
$ go test -v 
=== RUN TestList0
--- PASS: TestList0 (0.00s)
        easygen_test.go:17: First and plainest list test
=== RUN TestList1Text
--- PASS: TestList1Text (0.00s)
        easygen_test.go:28: Second test, with text template
=== RUN TestList1HTML
--- PASS: TestList1HTML (0.00s)
        easygen_test.go:36: Second test, with html template
=== RUN TestListFunc1
--- PASS: TestListFunc1 (0.00s)
        easygen_test.go:47: Test custom template function - minus1
=== RUN: ExampleFunc1
--- PASS: ExampleFunc1 (0.00s)
=== RUN: ExampleList0Func1
--- PASS: ExampleList0Func1 (0.00s)
=== RUN: ExampleList0StrTemplate
--- PASS: ExampleList0StrTemplate (0.00s)
=== RUN: ExampleFunc2
--- PASS: ExampleFunc2 (0.00s)
=== RUN: ExampleTestExample
--- PASS: ExampleTestExample (0.00s)
=== RUN: ExampleCommandLineCobraViper
--- PASS: ExampleCommandLineCobraViper (0.00s)
=== RUN: ExampleCommandLineOptInitFull
--- PASS: ExampleCommandLineOptInitFull (0.00s)
PASS
ok      github.com/suntong/easygen   0.014s
```

## Help

```
$ easygen

Usage:
 easygen [flags] YamlFileName

  -debug=0: debugging level
  -et=".tmpl": extension of template file
  -ey=".yaml": extension of yaml file
  -html=false: treat the template file as html instead of text
  -tf="": .tmpl template file name (default: same as .yaml file)
  -ts="": template string (in text)

YamlFileName: The name for the .yaml data and .tmpl template file
        Only the name part, without extension. Can include the path as well.

```

## Details

My (updated) blog about it is at [here](https://github.com/suntong/blog/blob/master/GoOptP7-easygen.md), and [here](https://sfxpt.wordpress.com/2015/07/04/easygen-is-now-coding-itself/).

<a name="clfhcag" />
## Command line flag handling code auto-generation

As explained above, one practical use of `easygen` is to auto-generating Go code for command line parameter handling, for both [`viper` and `cobra`](https://github.com/suntong/blog/blob/master/GoOptP7-easygen.md), and Go's [built-in `flag` package](https://sfxpt.wordpress.com/2015/07/04/easygen-is-now-coding-itself/).

Currently, `easygen`'s command line parameter handling is built on top of Go's built-in `flag` package, and the handling code is entirely generated by `easygen` itself. Thus, showing how `easygen` is handling the command line parameters itself also means showing what functionality the auto-generated command line parameter handling code can do for you. 

Currently, there are three tiers program parameters can be given:

0. Default values defined within the program, so that program parameters can have meaningful defaults to start with
0. Values defined in environment variables
0. Values passed from command line 

The latter will have higher priority and will override values defined formerly. I.e., the values from command line will override that in environment variables, which in turn override program defaults.

We will use the `-ts`, template string, as an example to illustrate. The program defaults is empty, which means using the `.tmpl` template file the same as the `.yaml` data file. We will override that first by environment variable, then from command line.


    echo 'Name: some-init-method' > /tmp/var.yaml

    $ EASYGEN_TS='{{.Name}}' easygen /tmp/var
    some-init-method

I.e., with the environment variable `EASYGEN_TS`, the `.tmpl` template file is not used.

	$ EASYGEN_TS='{{.Name}}' easygen -ts '{{ck2uc .Name}}' /tmp/var
	SomeInitMethod

I.e., command line value takes the highest priority, even overriding the environment variable `EASYGEN_TS`'s value.

As such, if you have a different naming convention than using `.tmpl` for template file and `.yaml` for data file, you can  override them in environment variables, `EASYGEN_ET` and `EASYGEN_EY`, so that you don't need to use `-et` and/or `-ey` to override them from command line each time. 

	echo 'Name: myConstantVariable' > /tmp/var.yml

	$ EASYGEN_EY=.yml easygen -ts '{{clc2ss .Name}}' /tmp/var
	MY_CONSTANT_VARIABLE


<a name="tips" />
## Tips

You can use `easygen` as an generic Go template testing tool with the `-ts` commandline option. For example,

```
echo "Age: 16" > /tmp/age.yaml

$ easygen -ts "{{.Age}}" /tmp/age
16

$ easygen -ts '{{printf "%x" .Age}}' /tmp/age
10

echo '{FirstName: John, LastName: Doe}' > /tmp/name.yaml

$ easygen -ts '{{.FirstName}}'\''s full name is {{printf "%s%s" .FirstName .LastName | len}} letters long.' /tmp/name
John's full name is 7 letters long.

$ easygen -ts "{{.FirstName}} {{ck2ss .LastName}}'s full name is "'{{len (printf "%s%s" .FirstName .LastName)}} letters long.' /tmp/name
John DOE's full name is 7 letters long.

echo 'Name: some-init-method' > /tmp/var.yaml

$ easygen -ts '{{.Name}} {{6 | minus1}} {{minus1 6}} {{ck2lc .Name}} {{ck2uc .Name}}' /tmp/var
some-init-method 5 5 someInitMethod SomeInitMethod

```

## Author(s)

Tong SUN

![suntong from cpan.org](http://i.stack.imgur.com/CNcsd.png "email address")

All patches welcome. 
