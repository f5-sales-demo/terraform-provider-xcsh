---
page_title: "xcsh_cdn_purge_command examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cdn_purge_command examples."
---

# xcsh_cdn_purge_command examples

<a id="canonical-1213321202020020-1110201131212332-3220110122310332-0101202001133221-3021211132222120-3032102012001031-2302002331310030-1110002322303332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0023022220111331-3102221003101103-0013330220230321-0113210333131330-1333010132202103-2320102103332311-0323220121121123-1132022333100211)
- Examples

<a id="canonical-2101211023130302-0121002300122121-1300022023202123-0020131123011003-2233113333033131-1211321121113002-3321303121113132-3101022210203133"></a>

### Complete configurations for `xcsh_cdn_purge_command`

- [Resource](resources--cdn_purge_command--examples--group-001.md#canonical-3333202021103021-2310121310221301-3233123123101103-2031330021132323-0022203300020120-3201221132110221-3210103202020020-1210220303100112): valid configuration.

<a id="canonical-3333202021103021-2310121310221301-3233123123101103-2031330021132323-0022203300020120-3201221132110221-3210103202020020-1210220303100112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_cdn_purge_command](../resources/cdn_purge_command.md#canonical-0023022220111331-3102221003101103-0013330220230321-0113210333131330-1333010132202103-2320102103332311-0323220121121123-1132022333100211)
- [Examples](resources--cdn_purge_command--examples--group-001.md#canonical-1213321202020020-1110201131212332-3220110122310332-0101202001133221-3021211132222120-3032102012001031-2302002331310030-1110002322303332)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cdn_purge_command/resource.tf`; digest `sha256:07fdb3a832901b7ce2b8f7a3292052a09eb46eebbc60fc38d6035bae249459aa`.

```terraform
# CDNPurgeCommand Resource Example
# Manages a CDN Purge Command resource in F5 Distributed Cloud for cdn purge command specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNPurgeCommand configuration
resource "xcsh_cdn_purge_command" "example" {
  name      = "example-cdn-purge-command"
  namespace = "staging"
}
```
