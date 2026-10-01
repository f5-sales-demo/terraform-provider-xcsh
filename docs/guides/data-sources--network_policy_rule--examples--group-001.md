---
page_title: "xcsh_network_policy_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_rule examples."
---

# xcsh_network_policy_rule examples

<a id="canonical-74efee1bd1fc68798568765407bea6a6fd732ef15f06ad742fa6395e0a09ffc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1f7f291bd2a37e8ebfb17a9371e89f6f79d0c52baddbf3060bcf356edf3bc72e"></a>

## Examples — Examples / 66b3eb6f8258 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)
- Examples

<a id="canonical-40cabe068e777da677f5b67351d3ae602631ea0b12cb180c9929d4b55cff77cd"></a>

## Complete configurations — Examples / 66b3eb6f8258 / 3

- [Data source](data-sources--network_policy_rule--examples--group-001.md#canonical-7de480fffa2ba386e8da652ee9ff6ec839a0894d353cd1dd0aa0dec788f24aa3): valid configuration.

<a id="canonical-fcd7442829c32f37497913ad24a1bf2f98fb765bbad765bb59fed2850d13df46"></a>

## Next pages — Examples / 66b3eb6f8258 / 4

- [Data source](data-sources--network_policy_rule--examples--group-001.md#canonical-7de480fffa2ba386e8da652ee9ff6ec839a0894d353cd1dd0aa0dec788f24aa3)
- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)

<a id="canonical-7de480fffa2ba386e8da652ee9ff6ec839a0894d353cd1dd0aa0dec788f24aa3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65fbbf1661473d36ea160b205e209a1a56b6df997c86524c206e747b7520c497"></a>

## Data source — Data source / a12e133a7404 / 2

Breadcrumbs:

- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)
- [Examples](data-sources--network_policy_rule--examples--group-001.md#canonical-74efee1bd1fc68798568765407bea6a6fd732ef15f06ad742fa6395e0a09ffc0)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_policy_rule/data-source.tf`; digest `sha256:e28844bbbc2540f4dac23720e7a2c315eba810ea0e139291452a348547d903cb`.

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

<a id="canonical-f17e8a55b0a6747be444eac4d1f206fea0d5537bf3ac2f3ace948d56dae8c702"></a>

## Next pages — Data source / a12e133a7404 / 3

- [Examples](data-sources--network_policy_rule--examples--group-001.md#canonical-74efee1bd1fc68798568765407bea6a6fd732ef15f06ad742fa6395e0a09ffc0)
- [xcsh_network_policy_rule](../data-sources/network_policy_rule.md#canonical-c48860d115aa5e9dc508362656c70f97db93634d5a35505e876e4ea891881f53)
