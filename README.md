# gg
Simple templating tool for generating code.


# Goals
- Simple API
```bash
gg template(or just t) http/controller/crud name=Person;DB=PersonDB
```
- Two types of template, 1: simple .go.tmpl files that get formatted using text/template
2: Binary templates that are .go files and are compiled and run and pass the args, write output to 
stdout and it will get piped to output file.
