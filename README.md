# Commander

Commander is an extremely simple command line utility that will run predefined directory scripts found in a directory specific commands toml file.

for example -> running:

    ```~bash
        cmd start
    ```

*in a directory with a commands.toml file that looks like:*

    ```~toml
    start="cd exampledir && echo \"Started from within the exampledir.\""
    ```

- will print "start" to terminal.

### Steps to Install

1. clone repo to your machine
2. cd into repo directory
3. run ```~bash go build```
4. add ```alias cmd="./path/to/binary"```

### Steps for use
1. create a commands.toml file in a project directory of your choice.
2. run ```~bash cmd [name of command]``` (without square braces.)
3. see results



