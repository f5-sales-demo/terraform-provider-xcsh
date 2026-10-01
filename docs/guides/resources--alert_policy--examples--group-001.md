---
page_title: "xcsh_alert_policy examples"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_alert_policy examples."
---

# xcsh_alert_policy examples

<a id="canonical-23b7b4b5d648e60bcae0210a5ebcbb48e237cd6477c8164b3b0eadf6dfa7eafb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1cbda8eb9be588dc6cd7b7c463f6633b4650ca2d205d74944711d20ff7ed069"></a>

## Examples — Examples / c690b18d5a79 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- Examples

<a id="canonical-c830ef62e6b00d41c359abfc1fbc6c4d11747910bb9959075ea63ea300d7a583"></a>

## Complete configurations — Examples / c690b18d5a79 / 3

- [Resource](resources--alert_policy--examples--group-001.md#canonical-6a0ed4065e5d2fbdb4f44c6148cd2502ea917f60c9cfbe582a2f460851397fb3): valid configuration.

<a id="canonical-0f5de87e24d9a9a3bfc93e4c213ee796572e7d6778ec8910ca0c9cddb1b4e357"></a>

## Next pages — Examples / c690b18d5a79 / 4

- [Resource](resources--alert_policy--examples--group-001.md#canonical-6a0ed4065e5d2fbdb4f44c6148cd2502ea917f60c9cfbe582a2f460851397fb3)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)

<a id="canonical-6a0ed4065e5d2fbdb4f44c6148cd2502ea917f60c9cfbe582a2f460851397fb3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-53fec801bd3e3fae74f1c1813e7e5f44cffaf8118df079cd2f5654123b88278b"></a>

## Resource — Resource / 718201a67690 / 2

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
- [Examples](resources--alert_policy--examples--group-001.md#canonical-23b7b4b5d648e60bcae0210a5ebcbb48e237cd6477c8164b3b0eadf6dfa7eafb)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_alert_policy/resource.tf`; digest `sha256:02e1db745e00ae3c15ffd89e3e1addc3f550309288f544ced75c1c7b1645161e`.

```terraform
# AlertPolicy Resource Example
# Manages new Alert Policy Object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AlertPolicy configuration
resource "xcsh_alert_policy" "example" {
  name      = "example-alert-policy"
  namespace = "staging"
}
```

<a id="canonical-60368a1ca5e94ec59949ab0a904f84c377ca98e0481c9546290c8c7ef3dd375e"></a>

## Next pages — Resource / 718201a67690 / 3

- [Examples](resources--alert_policy--examples--group-001.md#canonical-23b7b4b5d648e60bcae0210a5ebcbb48e237cd6477c8164b3b0eadf6dfa7eafb)
- [xcsh_alert_policy](../resources/alert_policy.md#canonical-5d597807557dacf4fb093618fe0f8d66ac210c9cff410a8b6f0f6dd44de4c310)
