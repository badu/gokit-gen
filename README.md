# gokit-gen

Provide a .proto file, get a go-kit generated package

## install

`go get -u github.com/badu/gokit-gen`

## parameters

1. proto file - can be absolute or relative

2. templates folder ("default" uses default templates) - can be absolute or relative

3. full package, for imports in generated tests

4. optional, deployment folder - can be absolute or relative

## Using ANTLR 

This is not needed, just in case you want to alter grammar file

1. Download ANTLR 

Change dir to /usr/local/lib and run:

`sudo curl -O https://www.antlr.org/download/antlr-4.8-complete.jar`

2. Create alias 

Note : Might want to save this command to `.bashrc` 

`alias antlr='java -jar /usr/local/lib/antlr-4.8-complete.jar'`

3. Run generator

Inside /pkg/antlr folder run :

`antlr -Dlanguage=Go -o parser antlr.g4`
 