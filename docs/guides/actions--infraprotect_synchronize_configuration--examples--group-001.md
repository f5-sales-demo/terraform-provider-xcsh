---
page_title: "xcsh_infraprotect_synchronize_configuration examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_infraprotect_synchronize_configuration examples."
---

# xcsh_infraprotect_synchronize_configuration examples

<a id="canonical-592bbdecfb48f78b5766a613695205d3e5f9dd31e0ac81e921edcb2d1794284c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9e3d5751cd68b9dc9bf181541217e8cc60704617ff7fb025ea2e5a1e6059ef2"></a>

## Examples — Examples / b0eaaf63e8ff / 2

Breadcrumbs:

- [xcsh_infraprotect_synchronize_configuration](../actions/infraprotect_synchronize_configuration.md#canonical-c3e1b3aa794f6032e8b56effdea59b5c380dd1c1dc0035df959f94ebabf3a655)
- Examples

<a id="canonical-23e273a6539a3ca24a2241fe8ac17461ef6feca6c2c3baa349e44b99b998e194"></a>

## Complete configurations — Examples / b0eaaf63e8ff / 3

- [Action](actions--infraprotect_synchronize_configuration--examples--group-001.md#canonical-476178c71ec712ebe3a3f9b1bdbc4d6d9906a64413cb5abb29d6ac9b1e960e35): valid configuration.

<a id="canonical-f45019170621dab62634615b269ac917fb89f55b6f20a36d41508dbbde307e8a"></a>

## Next pages — Examples / b0eaaf63e8ff / 4

- [Action](actions--infraprotect_synchronize_configuration--examples--group-001.md#canonical-476178c71ec712ebe3a3f9b1bdbc4d6d9906a64413cb5abb29d6ac9b1e960e35)
- [xcsh_infraprotect_synchronize_configuration](../actions/infraprotect_synchronize_configuration.md#canonical-c3e1b3aa794f6032e8b56effdea59b5c380dd1c1dc0035df959f94ebabf3a655)

<a id="canonical-476178c71ec712ebe3a3f9b1bdbc4d6d9906a64413cb5abb29d6ac9b1e960e35"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e618feaa5006d324a5ff7c05d923da94333c79c9f506fef2fd4f656cc3dbc3fa"></a>

## Action — Action / 19e5094aac28 / 2

Breadcrumbs:

- [xcsh_infraprotect_synchronize_configuration](../actions/infraprotect_synchronize_configuration.md#canonical-c3e1b3aa794f6032e8b56effdea59b5c380dd1c1dc0035df959f94ebabf3a655)
- [Examples](actions--infraprotect_synchronize_configuration--examples--group-001.md#canonical-592bbdecfb48f78b5766a613695205d3e5f9dd31e0ac81e921edcb2d1794284c)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_infraprotect_synchronize_configuration/action.tf`; digest `sha256:5708bb8f8ebc531162d8b17ecbd5421cebdfe14f2d0a5a47b3ca315eedfc8227`.

```terraform
# InfraprotectSynchronizeConfiguration Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_infraprotect_synchronize_configuration" "example" {
  config {
    namespace = "example-value"
  }
}
```

<a id="canonical-e2c53903461b8cb547a659fe471865ed164e7efc76c4f35b78956c04183065a8"></a>

## Next pages — Action / 19e5094aac28 / 3

- [Examples](actions--infraprotect_synchronize_configuration--examples--group-001.md#canonical-592bbdecfb48f78b5766a613695205d3e5f9dd31e0ac81e921edcb2d1794284c)
- [xcsh_infraprotect_synchronize_configuration](../actions/infraprotect_synchronize_configuration.md#canonical-c3e1b3aa794f6032e8b56effdea59b5c380dd1c1dc0035df959f94ebabf3a655)
