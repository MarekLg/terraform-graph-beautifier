# terraform-graph-beautifier

Command line tool parsing the output of the `terraform graph` command and producing a cleaner and prettier [Grahpviz](https://www.graphviz.org/) output.

It reads the [Dot](https://www.graphviz.org/doc/info/lang.html) script generated by `terraform graph` from the standard input or a file and produces a cleaned and enhanced Dot script on the standard output which can be redirected to a file and/or piped to a [rendering command](https://linux.die.net/man/1/dot).

## Graph loading and cleaning / filtering
The loading of the input graph involves several steps:
1. **Cleaning**
   1. Renaming the nodes using a more consistent pattern : `"[root] rsc_type.rsc_name"` => `"module.root.rsc_type.rsc_name"`.
   1. Using `'` instead of `"` in map keys : `buckets["artefacts"]` => `buckets['artefacts']`
   1. Removing the nodes and edges generated by Terraform but not corresponding to configuration elements (aka TF junk). This can be deactivated via a command line switch.
1. **Filtering**
   - Using user-provided pattern(s) to exclude some elements (resource, var, module, provider, ...) from the output.
   - These patterns are [Go regexp](https://golang.org/pkg/regexp/) and are matched line by line against the output of the **cleaning** step, so use the `"root.rsc_type.rsc_name"` naming. 

## TODO
- Dans samples, configs TF de tests
  - TF only (local file, random) avec:
    - data
    - variables, locals et outputs
    - maps et lists
  - Utilisation des modules GCP
- Config file
```yaml
tf-types:
  - type (resource, provider, var, local, output) :
    - shape: box(ressource), arrondi(var, local), ?? (output)
    - bg-colors: [...]
    - fg-colors: [...]
    - arrow-to-shape: ""
```
- Param pour générer un fichier de config avec valeurs par défaut. 
- Customs attributes
    - tf-type
- Voir (orphan), (close) (removed)
- Exemples utilisation [`sfdp`](https://linux.die.net/man/1/sfdp)
- Github actions to build
- compound="true"; newrank="true";
- 2 ways to configure the command execution:
    - a configuration file (can be checked-in alongside your Terraform configuration)
    - command line arguments
    if both are used, command line arguments take precedence : check if some params can be merged
- output param
