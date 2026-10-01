---
page_title: "xcsh_service_policy_rule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_service_policy_rule landing."
---

# xcsh_service_policy_rule landing

<a id="canonical-23840e53ba6988c47ffff00ddad7d23482f6a78fb2ee069970e8f0408f50ae5f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f8655855e7dd56310e3c5f7d789e146d7255a8d17a0a42e79ac94596caccecda"></a>

## xcsh_service_policy_rule — xcsh_service_policy_rule / 721ee8435e95 / 2

Breadcrumbs:

- xcsh_service_policy_rule

Manages service\_policy\_rule creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-4ef01695938411b7a1f81b877ac94c72a6beffac63e26008fa093d815e216ab2"></a>

## Prerequisites — xcsh_service_policy_rule / 721ee8435e95 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-514e138b25f93931679d935d54703929b9c8f04f72db41960990983f5bfdbed1"></a>

## Minimal configuration — xcsh_service_policy_rule / 721ee8435e95 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-e574a1e4412077beeb7b2959f2204ebe6e7dff77f7741a92a11f974bd272559b"></a>

## Root configuration — xcsh_service_policy_rule / 721ee8435e95 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-65ca7ff2603eddb3756c4e8d8c6d6235857d666bd7f66b538d0b7002f34213e0"></a>

## Next pages — xcsh_service_policy_rule / 721ee8435e95 / 6

- [Property reference](../guides/data-sources--service_policy_rule--reference--group-001.md#canonical-fdd489dc2516d183c2259a43d596709411628f3d783f7fb017e8b6dc25954251)
- [Examples](../guides/data-sources--service_policy_rule--examples--group-001.md#canonical-6b5b4f215fa3736f8550a713d3f8add0451aa3fb1d76afb7d338d5960210da6d)
