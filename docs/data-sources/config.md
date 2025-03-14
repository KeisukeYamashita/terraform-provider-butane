---
page_title: "butane_config - terraform-provider-butane"
subcategory: ""
description: |-
    Validate and transpile Butane config to Ignition config.
---

# butane_config

Validate and transpile Butane config to Ignition config.

## Example Usage

```terraform
data "butane_config" "config" {
  content = file("./config.bu")
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `content` (String) Butane configuration file.

### Optional

- `files_dir` (String) Directory to embed the local files. Maps to `--files-dir` option on Butane CLI.
- `pretty` (Boolean) Output formatted results. Maps to `--pretty` option on Butane CLI.
- `strict` (Boolean) Strictly check the format. Any warning will make transpile fail. Maps to `--strict` option on Butane CLI.

### Read-Only

- `id` (String) The ID of this resource.
- `ignition` (String) Result Ignition configuration.
