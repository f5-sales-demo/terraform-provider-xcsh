---
page_title: "xcsh_protected_application examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application examples."
---

# xcsh_protected_application examples

<a id="canonical-1451a94d52ce926e002b158e3403caa1e4a2c208b632bc30ad80f7db36f5ab96"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-79f7fe7eef9145d7c10bd969aab1f043bf51ed80d5a01e3f29e46860a736e165"></a>

## Examples — Examples / 16fcdc715998 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- Examples

<a id="canonical-4898d69c4d52cf367d083ff9d66ce9693ad82e561f5cc91f0cc08237cc229fdc"></a>

## Complete configurations — Examples / 16fcdc715998 / 3

- [Resource](resources--protected_application--examples--group-001.md#canonical-18a3034e2804b479d4994e57606f8cb23a175bccd562da1d657893efe136a23f): valid configuration.

<a id="canonical-b26e15f270596a1f49b298dc9d5386a3d4a3bf6ff89612bde2533d4057884850"></a>

## Next pages — Examples / 16fcdc715998 / 4

- [Resource](resources--protected_application--examples--group-001.md#canonical-18a3034e2804b479d4994e57606f8cb23a175bccd562da1d657893efe136a23f)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)

<a id="canonical-18a3034e2804b479d4994e57606f8cb23a175bccd562da1d657893efe136a23f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-58e52324bd121345ae1550c85faa5e286e3991a1edbbfc6b97231b2cad6cb50e"></a>

## Resource — Resource / 6e3e2e464b88 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
- [Examples](resources--protected_application--examples--group-001.md#canonical-1451a94d52ce926e002b158e3403caa1e4a2c208b632bc30ad80f7db36f5ab96)
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

<a id="canonical-56221a284bd53b13e1bb8ede49ca8db8204a83440c27a7a461f6cedcdc6b136b"></a>

## Next pages — Resource / 6e3e2e464b88 / 3

- [Examples](resources--protected_application--examples--group-001.md#canonical-1451a94d52ce926e002b158e3403caa1e4a2c208b632bc30ad80f7db36f5ab96)
- [xcsh_protected_application](../resources/protected_application.md#canonical-02eb6070b96ac03a5ec9e61db42ba499f057c1cceb1a2d74492268be0811fd42)
