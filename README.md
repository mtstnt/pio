# Pio

Personal I/O assistant (?)

Save project and use it everytime you need it with a simple command.

<br>

## Scratch notes for myself

<br>

Known dependencies directory name which will be ignored automatically.  
Include all with flag: `--all` or `-a` or manually with `--include-dirs=[...]`:  
- `node_modules` => if `package.json` is detected.  
- `vendor` => if `composer.json` is detected.  
- `target` => if `cargo.toml` is detected.

Add more in default ignore list (coming soon).

<br>

Commands:
- `pio add [template_name] [path | .]`   
  Adds a new template with the template_name as its name and takes the files template from the path argument, include all directories except for known dependencies directories.

- `pio list [search | ""]`  
  Lists all templates with their names sorted alphabetically. Optional search argument to filter by name.

- `pio generate [template_name] [path | .]`  
  Generate files from the template in path.

- `pio remove [template_name]`  
  Removes the template from templates list.

- `pio rename [template_name] [new_template_name]`  
  Renames a template with a new name.

<br>

To add later:
- Add values in files that can be changed on template generation. (Ex: project name, class name, etc)
- Allow custom parameters in generate.
- Put templates folder in $HOME/.pio for better structure