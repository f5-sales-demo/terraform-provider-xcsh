---
page_title: "xcsh_network_policy_rule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_rule landing."
---

# xcsh_network_policy_rule landing

<a id="canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a4970bf0b2a52bd0b2780847f4cd2805e71afa93957128df0245e12011eb04a3"></a>

## xcsh_network_policy_rule — xcsh_network_policy_rule / 137f9e82eedc / 2

Breadcrumbs:

- xcsh_network_policy_rule

Manages network policy rule with configured parameters in specified namespace in F5 Distributed
Cloud.

<a id="canonical-6c6baf1389cb3f0bf9c99481d344ba39045b384bf832f1928572766e98da3137"></a>

## Prerequisites — xcsh_network_policy_rule / 137f9e82eedc / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-31ec54bc62f45707dbec5a6ec353ce06c984b22ff82f50cc782020503ce5a783"></a>

## Minimal configuration — xcsh_network_policy_rule / 137f9e82eedc / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicyRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicyRule by name
data "xcsh_network_policy_rule" "example" {
  name      = "example-network-policy-rule"
  namespace = "staging"
}

output "network_policy_rule_id" {
  value = data.xcsh_network_policy_rule.example.id
}
```

<a id="canonical-7834dde0c37e3cded1cd2428dc9547b3a588252ac2f1319a50b0f296e1f12472"></a>

## Root configuration — xcsh_network_policy_rule / 137f9e82eedc / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-dc91cbe2bc6bcfae10d54c0325503ed43bbd7a8fcdcd08dec7f53a35b5615a77"></a>

## Next pages — xcsh_network_policy_rule / 137f9e82eedc / 6

- [Property reference](../guides/data-sources--network_policy_rule--reference--group-001.md#canonical-158a950e1cc26821bd58408cdc9d7925eeda2761080b3af43d991b7b58c75bf0)
- [Examples](../guides/data-sources--network_policy_rule--examples--group-001.md#canonical-74efee1bd1fc68798568765407bea6a6fd732ef15f06ad742fa6395e0a09ffc0)
