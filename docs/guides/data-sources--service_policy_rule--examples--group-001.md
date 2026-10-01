---
page_title: "xcsh_service_policy_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule examples."
---

# xcsh_service_policy_rule examples

<a id="canonical-6b5b4f215fa3736f8550a713d3f8add0451aa3fb1d76afb7d338d5960210da6d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-38e01df2c8a68bd75457dde07a03f1dd5b19f497e318bb1fe0da01770c6ee6ba"></a>

## Examples — Examples / 69e7dc20213b / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- Examples

<a id="canonical-8f4c0feb72ddadd5aa243b2c0d434b97faeb869054fb46c2b2cfd3b0086f0938"></a>

## Complete configurations — Examples / 69e7dc20213b / 3

- [Data source](data-sources--service_policy_rule--examples--group-001.md#canonical-f8645e21882b8f074465846ce00589478a2c687b3fba7e2793d296b7c5faa2f4): valid configuration.

<a id="canonical-42c30dd534dd894b6614875c536e9a29196c2997ae3ef7f67aec826db6cf58c7"></a>

## Next pages — Examples / 69e7dc20213b / 4

- [Data source](data-sources--service_policy_rule--examples--group-001.md#canonical-f8645e21882b8f074465846ce00589478a2c687b3fba7e2793d296b7c5faa2f4)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)

<a id="canonical-f8645e21882b8f074465846ce00589478a2c687b3fba7e2793d296b7c5faa2f4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5c954beaf0580fecbfd83cdc51ddee504d7f7d1be03b693c343abd6597f9645a"></a>

## Data source — Data source / eaf7449f0d8a / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
- [Examples](data-sources--service_policy_rule--examples--group-001.md#canonical-6b5b4f215fa3736f8550a713d3f8add0451aa3fb1d76afb7d338d5960210da6d)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_service_policy_rule/data-source.tf`; digest `sha256:ae57799b0f1696fd5e6a11587ed3b1a1bc30f50fe8953264b9b955e84a34c726`.

```terraform
# ServicePolicyRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicyRule by name
data "xcsh_service_policy_rule" "example" {
  name      = "example-service-policy-rule"
  namespace = "staging"
}

output "service_policy_rule_id" {
  value = data.xcsh_service_policy_rule.example.id
}
```

<a id="canonical-2a3cda8cb0a4ed277f3696b6a196eeeecf683a7f000cf5c39988909d868c120d"></a>

## Next pages — Data source / eaf7449f0d8a / 3

- [Examples](data-sources--service_policy_rule--examples--group-001.md#canonical-6b5b4f215fa3736f8550a713d3f8add0451aa3fb1d76afb7d338d5960210da6d)
- [xcsh_service_policy_rule](../data-sources/service_policy_rule.md#canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f)
