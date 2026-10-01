---
page_title: "xcsh_alert_gen_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_gen_policy examples."
---

# xcsh_alert_gen_policy examples

<a id="canonical-7fd687224e91550528fa39bd463f5c448de431f0363111f202eaf62154fd297f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52bf864bbed0d4e87c3cb6b085d25a9058469c6a720f8d2ffb1170f9dbb6b4f2"></a>

## Examples — Examples / c1cbe7f87659 / 2

Breadcrumbs:

- [xcsh_alert_gen_policy](../resources/alert_gen_policy.md#canonical-a1aae63c73737fe28891b8dddcd40970e1ba1214a48d75c4ae655469cca81541)
- Examples

<a id="canonical-5275459205034aa302340e478b5dbd9f98261caa0cd384f11d95d3556cebbed1"></a>

## Complete configurations — Examples / c1cbe7f87659 / 3

- [Resource](resources--alert_gen_policy--examples--group-001.md#canonical-ea1420ab7ca305d71dc51d9412ef45b0be8ffdb3b990b6ca5c945bb4ad53eaba): valid configuration.

<a id="canonical-d789c6da984ea30f8bf54835eac47e21274357ac9309f1c54c4caec2ff115e05"></a>

## Next pages — Examples / c1cbe7f87659 / 4

- [Resource](resources--alert_gen_policy--examples--group-001.md#canonical-ea1420ab7ca305d71dc51d9412ef45b0be8ffdb3b990b6ca5c945bb4ad53eaba)
- [xcsh_alert_gen_policy](../resources/alert_gen_policy.md#canonical-a1aae63c73737fe28891b8dddcd40970e1ba1214a48d75c4ae655469cca81541)

<a id="canonical-ea1420ab7ca305d71dc51d9412ef45b0be8ffdb3b990b6ca5c945bb4ad53eaba"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fc32f2782b6d68fa2425e350f644bf5e91a17fcbaec66d52c37af0c7081a15fb"></a>

## Resource — Resource / 9536ea6601e0 / 2

Breadcrumbs:

- [xcsh_alert_gen_policy](../resources/alert_gen_policy.md#canonical-a1aae63c73737fe28891b8dddcd40970e1ba1214a48d75c4ae655469cca81541)
- [Examples](resources--alert_gen_policy--examples--group-001.md#canonical-7fd687224e91550528fa39bd463f5c448de431f0363111f202eaf62154fd297f)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_alert_gen_policy/resource.tf`; digest `sha256:eea28da7e92ac8598036f86f1095005a92f369536d4496a75262a4b141db489a`.

```terraform
# AlertGenPolicy Resource Example
# Manages Alert Generation Policy in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertGenPolicy configuration
resource "xcsh_alert_gen_policy" "example" {
  name      = "example-alert-gen-policy"
  namespace = "staging"
}
```

<a id="canonical-a5cb694d6dbec53d2c2a75e4a862b8bbe1f2dba6fe89d491b0e021aaab1e3739"></a>

## Next pages — Resource / 9536ea6601e0 / 3

- [Examples](resources--alert_gen_policy--examples--group-001.md#canonical-7fd687224e91550528fa39bd463f5c448de431f0363111f202eaf62154fd297f)
- [xcsh_alert_gen_policy](../resources/alert_gen_policy.md#canonical-a1aae63c73737fe28891b8dddcd40970e1ba1214a48d75c4ae655469cca81541)
