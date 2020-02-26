# gokit-gen

Provide a .proto file, get a go-kit generated package



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

Note to self : the grammar definition is very important in terms of making the work easier.
For example, most elements have `option` (`rpc`,`service`,`proto`,`message`,`enum`) - in order to collect those correctly, I had to redefine global `option` into species, for each kind and for that reason, now there are several similar definitions : `protoOption`,`enumOption`,`messageOption`,`serviceOption` and `rpcOption`. A little repetition makes work easier, isn't it?
 