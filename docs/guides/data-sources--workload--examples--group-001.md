---
page_title: "xcsh_workload examples"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload examples."
---

# xcsh_workload examples

<a id="canonical-d495c5379762ac477e71189db6d070866bf97284034ec2aa67a79b9d27182312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ddf7d98612bfe185bb9fa85490f4046713b8b68c81754d4179ee81027eb0890e"></a>

## Examples — Examples / 6fad60897ca7 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- Examples

<a id="canonical-7a009e3c11109d346edac9cbe5ae95220d0812c9328e666f62a7af6f828f8738"></a>

## Complete configurations — Examples / 6fad60897ca7 / 3

- [Data source](data-sources--workload--examples--group-001.md#canonical-0945a406393826ebf5794dcbca44a9825ffcf268fd46f827d5dbd9369828668f): valid configuration.

<a id="canonical-37e446178ba9dd98b41948f2ce113757714444456bb313077ac9702ff39c18e4"></a>

## Next pages — Examples / 6fad60897ca7 / 4

- [Data source](data-sources--workload--examples--group-001.md#canonical-0945a406393826ebf5794dcbca44a9825ffcf268fd46f827d5dbd9369828668f)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)

<a id="canonical-0945a406393826ebf5794dcbca44a9825ffcf268fd46f827d5dbd9369828668f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-239379d2db909f2f58d6e3f5e6edb4edc4db473a41d00e08eaa9abc7a9665aa9"></a>

## Data source — Data source / 67727b7df12b / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
- [Examples](data-sources--workload--examples--group-001.md#canonical-d495c5379762ac477e71189db6d070866bf97284034ec2aa67a79b9d27182312)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_workload/data-source.tf`; digest `sha256:cc2cda57d20aa47819844a194fd93abc59f8f0880458fe8ebeae6e323411954a`.

```terraform
# Workload Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Workload by name
data "xcsh_workload" "example" {
  name      = "example-workload"
  namespace = "staging"
}

output "workload_id" {
  value = data.xcsh_workload.example.id
}
```

<a id="canonical-bb4e8e11d9abb44b972de00fea3e341e2bd510a8b9c9503e9951c27331122832"></a>

## Next pages — Data source / 67727b7df12b / 3

- [Examples](data-sources--workload--examples--group-001.md#canonical-d495c5379762ac477e71189db6d070866bf97284034ec2aa67a79b9d27182312)
- [xcsh_workload](../data-sources/workload.md#canonical-425fbc5b40b45aaa9c0e88777ee098683a4abf92d4cd7e8863d28742d41e1d50)
