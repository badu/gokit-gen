# gokit-gen

Provide a .proto file, get a go-kit generated package

# install

`go get -u github.com/badu/gokit-gen`

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

`antlr -Dlanguage=Go -o parser Protobuf3.g4`
 