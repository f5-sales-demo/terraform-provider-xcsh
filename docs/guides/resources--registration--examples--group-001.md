---
page_title: "xcsh_registration examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration examples."
---

# xcsh_registration examples

<a id="canonical-459f2348c9bd2f11d18baccc06df68996e9613672dd6ffd94092fc45ee78b9e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-672ad09dcab23d28c4b03694fa3796e63d2a8cbb131d3e1282dbd521eed253cf"></a>

## Examples — Examples / de18d0179cb8 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- Examples

<a id="canonical-49df1f1a3cbe2a2529cc8bf8455b07715df9fea97bc5d028135d704f84777982"></a>

## Complete configurations — Examples / de18d0179cb8 / 3

- [Resource](resources--registration--examples--group-001.md#canonical-668fdcee31f43eaaf9b9058f7530df9810d3ee645d4e775de0579e8cdd13b044): valid configuration.

<a id="canonical-a59881cb8efd976f9208335342911dda18a3f19272a0c313672aa72ca6a7ee12"></a>

## Next pages — Examples / de18d0179cb8 / 4

- [Resource](resources--registration--examples--group-001.md#canonical-668fdcee31f43eaaf9b9058f7530df9810d3ee645d4e775de0579e8cdd13b044)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)

<a id="canonical-668fdcee31f43eaaf9b9058f7530df9810d3ee645d4e775de0579e8cdd13b044"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b8d77ee61757b2b44f01514901b44586546ed4ebf4349c8d7373b9348480c35"></a>

## Resource — Resource / e88ced41db33 / 2

Breadcrumbs:

- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
- [Examples](resources--registration--examples--group-001.md#canonical-459f2348c9bd2f11d18baccc06df68996e9613672dd6ffd94092fc45ee78b9e3)
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

<a id="canonical-2fc93dc470c2b7d955875ef4dbb9167d192773870eb2b8507c1afa1f2dd861c7"></a>

## Next pages — Resource / e88ced41db33 / 3

- [Examples](resources--registration--examples--group-001.md#canonical-459f2348c9bd2f11d18baccc06df68996e9613672dd6ffd94092fc45ee78b9e3)
- [xcsh_registration](../resources/registration.md#canonical-0de450498296e1bc9470fae21096c5b0f39fc922dddfd187ae719385f8bfe5bb)
