---
page_title: "xcsh_protected_application examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application examples."
---

# xcsh_protected_application examples

<a id="canonical-0110110122211031-1102303221021232-0000022301112032-0310000330222201-3210220230020020-2312030223300300-2231200033133123-0312331122232112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- Examples

<a id="canonical-1321331333321332-3233210110113113-3001002331211221-2222230133001003-2333110132312000-3111220001320333-0221321012201200-2213031232011211"></a>

### Complete configurations for `xcsh_protected_application`

- [Resource](resources--protected_application--examples--group-001.md#canonical-0120220300031032-0220001023101321-3110212110321113-1200123320302302-0322011311233030-3111120231220131-1211132021033233-3201031222020333): valid configuration.

<a id="canonical-0120220300031032-0220001023101321-3110212110321113-1200123320302302-0322011311233030-3111120231220131-1211132021033233-3201031222020333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Examples](resources--protected_application--examples--group-001.md#canonical-0110110122211031-1102303221021232-0000022301112032-0310000330222201-3210220230020020-2312030223300300-2231200033133123-0312331122232112)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_protected_application/resource.tf`; digest `sha256:868be3bcffc6df67dc51247c65c90b154557ee6fbafff96f6968ad86d503ff19`.

```terraform
# ProtectedApplication Resource Example
# Manages applications protected by Bot Defense in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtectedApplication configuration
resource "xcsh_protected_application" "example" {
  name      = "example-protected-application"
  namespace = "staging"
}
```
