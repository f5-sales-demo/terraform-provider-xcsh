---
page_title: "xcsh_cminstance examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cminstance examples."
---

# xcsh_cminstance examples

<a id="canonical-65ac32d0660495d526e5fd75556d24254c3381e30194324ea5c2a80aebbd6132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-83401707b0d97719fb3f68fffe40c9f764ae618d9eab1989179a7decfb6f0942"></a>

## Examples — Examples / ffa7f555bdde / 2

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)
- Examples

<a id="canonical-c729c47474eeb935f47ca332421e4cd4ab5ccf655afc7b593ce3f4f762dbe880"></a>

## Complete configurations — Examples / ffa7f555bdde / 3

- [Resource](resources--cminstance--examples--group-001.md#canonical-0c640fbc05dc476ca49ad92ea0b270b4053362fc57169385804213930847fc88): valid configuration.

<a id="canonical-201ee41ecf85630595c8337d239badd906ac3f19111ca9012fa88dbdde29756b"></a>

## Next pages — Examples / ffa7f555bdde / 4

- [Resource](resources--cminstance--examples--group-001.md#canonical-0c640fbc05dc476ca49ad92ea0b270b4053362fc57169385804213930847fc88)
- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)

<a id="canonical-0c640fbc05dc476ca49ad92ea0b270b4053362fc57169385804213930847fc88"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-759b502831520c8047018369d5300f01455b97854ec2a374093a4197a1e2aee4"></a>

## Resource — Resource / 1e22911577d5 / 2

Breadcrumbs:

- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)
- [Examples](resources--cminstance--examples--group-001.md#canonical-65ac32d0660495d526e5fd75556d24254c3381e30194324ea5c2a80aebbd6132)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cminstance/resource.tf`; digest `sha256:8160520d3332724273c3557478a3147a7b31a2fd2927166c5bb6f2b977df1b95`.

```terraform
# Cminstance Resource Example
# Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Cminstance configuration
resource "xcsh_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "staging"

  port     = 1
  username = "example-value"
}
```

<a id="canonical-d77fb6cdbd50b0108d8ca1c98ed7c891037af1be5baf435ac8e1730b416b5570"></a>

## Next pages — Resource / 1e22911577d5 / 3

- [Examples](resources--cminstance--examples--group-001.md#canonical-65ac32d0660495d526e5fd75556d24254c3381e30194324ea5c2a80aebbd6132)
- [xcsh_cminstance](../resources/cminstance.md#canonical-ccc430ab59d99b966ecc418027f2e0ea0c65f45b5c814e99073318de23376ee5)
