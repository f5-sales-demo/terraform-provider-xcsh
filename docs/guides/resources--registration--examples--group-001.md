---
page_title: "xcsh_registration examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration examples."
---

# xcsh_registration examples

<a id="canonical-1011213302031020-3021233102330101-3101202322303030-0012313312202121-1232211201031213-0231311233333121-1000210233301011-3232132023213203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- Examples

<a id="canonical-1213022231002131-3022230203310220-3010230003122110-3322031321123212-0331022220302323-0103013103320102-2002312331110201-3232310211033033"></a>

### Complete configurations for `xcsh_registration`

- [Resource](resources--registration--examples--group-001.md#canonical-1212203331303232-0301331003322222-3321232100112033-1311030031332120-0100310332321210-1131103213131131-3200111321322030-3131010323001010): valid configuration.

<a id="canonical-1212203331303232-0301331003322222-3321232100112033-1311030031332120-0100310332321210-1131103213131131-3200111321322030-3131010323001010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0031321011001021-2002211232012330-2110130033223202-0100211230112300-3303213330210202-3131313331012013-2232130121032011-3320233332112323)
- [Examples](resources--registration--examples--group-001.md#canonical-1011213302031020-3021233102330101-3101202322303030-0012313312202121-1232211201031213-0231311233333121-1000210233301011-3232132023213203)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_registration/resource.tf`; digest `sha256:63b17b62e7eb3ba883d15945e8c1f68dbbb4f71a319c4ef34eeeacc2e2bd25dc`.

```terraform
# Registration Resource Example
# Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this message, never used by users.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Registration configuration
resource "xcsh_registration" "example" {
  name      = "example-registration"
  namespace = "staging"

  token = "example-value"
}
```
