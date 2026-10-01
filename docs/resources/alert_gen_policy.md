---
page_title: "xcsh_alert_gen_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_gen_policy landing."
---

# xcsh_alert_gen_policy landing

<a id="canonical-a1aae63c73737fe28891b8dddcd40970e1ba1214a48d75c4ae655469cca81541"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-397e5e2147fc4814f76f37cfedd870f8a815dff72bc5941992d55ced23d3e68e"></a>

## xcsh_alert_gen_policy — xcsh_alert_gen_policy / c4e9aa96b278 / 2

Breadcrumbs:

- xcsh_alert_gen_policy

Manages Alert Generation Policy in F5 Distributed Cloud.

<a id="canonical-e702e9a38f6702b155ea2df87fc401cae9afaaeedc051146fbed12a84ca529ba"></a>

## Prerequisites — xcsh_alert_gen_policy / c4e9aa96b278 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0c8b757c54631b400af580b7ca01575261850d1a0388f28aa5d175b23afafdec"></a>

## Minimal configuration — xcsh_alert_gen_policy / c4e9aa96b278 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-3e71ff1e41cf2c47553aae6b70df9f9d7a0f2c03c1a4e70c5e874d204d435b0e"></a>

## Root configuration — xcsh_alert_gen_policy / c4e9aa96b278 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-4cb63fdb7af5bd73492484f76334e0aac331416e41e85d5f656b1433d6839886"></a>

## Next pages — xcsh_alert_gen_policy / c4e9aa96b278 / 6

- [Property reference](../guides/resources--alert_gen_policy--reference--group-001.md#canonical-17ec4655f437510f18a5ab3c4e0b2710c50f8dd33e3adde8ef55402086569933)
- [Examples](../guides/resources--alert_gen_policy--examples--group-001.md#canonical-7fd687224e91550528fa39bd463f5c448de431f0363111f202eaf62154fd297f)
- [Import](../guides/resources--alert_gen_policy--lifecycle--group-001.md#canonical-53449129ad12f0eb78a15cd536169b009ea5e30c6376054473ad2a1729808e46)
- [Timeouts](../guides/resources--alert_gen_policy--lifecycle--group-001.md#canonical-a3bd1f56deba8fbdde007203f5b82bbf9fea60c4ecac6aa8965eb97cfc8a35d7)
