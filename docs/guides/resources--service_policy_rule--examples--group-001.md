---
page_title: "xcsh_service_policy_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule examples."
---

# xcsh_service_policy_rule examples

<a id="canonical-0b47f8afb67032c78d6321166b882fd0532f386cde338c2c92d9c8bc15c25118"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-48a1f3143fbb4c7b53a418be50894223a69847b8fb56f6279935b680b2cb05c8"></a>

## Examples — Examples / dc6faaa2f0cf / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- Examples

<a id="canonical-92d03432ac51f3e22f66f20033ac2e7823ccf482ef645d37f8a58163994a8b6a"></a>

## Complete configurations — Examples / dc6faaa2f0cf / 3

- [Resource](resources--service_policy_rule--examples--group-001.md#canonical-0d1bcd64a075bd02f30fdb2229c76ee67e7d71f0cd2d7add990a8676887096a9): valid configuration.

<a id="canonical-c19e0c5d254dbfbaba9cb1c336bc3a71305e553666eb91f3c43a753b0862215b"></a>

## Next pages — Examples / dc6faaa2f0cf / 4

- [Resource](resources--service_policy_rule--examples--group-001.md#canonical-0d1bcd64a075bd02f30fdb2229c76ee67e7d71f0cd2d7add990a8676887096a9)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)

<a id="canonical-0d1bcd64a075bd02f30fdb2229c76ee67e7d71f0cd2d7add990a8676887096a9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-94b1d9cc2b6d6e52da695ec9bb104002030685fb10119e4c61af0ec757cbe96d"></a>

## Resource — Resource / 56f82515b3a8 / 2

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
- [Examples](resources--service_policy_rule--examples--group-001.md#canonical-0b47f8afb67032c78d6321166b882fd0532f386cde338c2c92d9c8bc15c25118)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy_rule/resource.tf`; digest `sha256:ddddb543e61078456915a99cb29c3a42a961001d81b7eb50761d3e637258e884`.

```terraform
# ServicePolicyRule Resource Example
# Manages service_policy_rule creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ServicePolicyRule configuration
resource "xcsh_service_policy_rule" "example" {
  name      = "example-service-policy-rule"
  namespace = "staging"

  action = "DENY"
}
```

<a id="canonical-3215263ad4920cca62bb1af6ec34c256ad19a2bcfd8b951b475f7f4033e77422"></a>

## Next pages — Resource / 56f82515b3a8 / 3

- [Examples](resources--service_policy_rule--examples--group-001.md#canonical-0b47f8afb67032c78d6321166b882fd0532f386cde338c2c92d9c8bc15c25118)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md#canonical-1da84b6dc4f5299900f209052f10a2987a5ffbbd70404a7f4621e1315566f5ce)
