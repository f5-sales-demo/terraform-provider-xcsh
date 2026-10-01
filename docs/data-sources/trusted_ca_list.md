---
page_title: "xcsh_trusted_ca_list landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_trusted_ca_list landing."
---

# xcsh_trusted_ca_list landing

<a id="canonical-ed5a0ca4002a9acc208559e5d135f6efdf6a9d227170b35fe4e3ac7a7857549a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-588dbda040c74fa0a54ae11d4df87002d8de8bc5b05ef1b6cb09417d52971ec7"></a>

## xcsh_trusted_ca_list — xcsh_trusted_ca_list / 9351258b0821 / 2

Breadcrumbs:

- xcsh_trusted_ca_list

Manages a Trusted CA List resource in F5 Distributed Cloud for trusted certificate authority list
management.

<a id="canonical-509efe3b06034c5561ee1e8f2b56e0572e9ccf27ba918022c55d2690fa1805cd"></a>

## Prerequisites — xcsh_trusted_ca_list / 9351258b0821 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-935ae49492d695d94bdfdac9ca6c547dfa8e1bb695b007bf5d01b9be00eeea23"></a>

## Minimal configuration — xcsh_trusted_ca_list / 9351258b0821 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# TrustedCAList Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TrustedCAList by name
data "xcsh_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "staging"
}

output "trusted_ca_list_id" {
  value = data.xcsh_trusted_ca_list.example.id
}
```

<a id="canonical-bc7b7b9ff4ff7b7e3ed57dc63cefca704308e527daeda0ba61a9c3ae8920fa82"></a>

## Root configuration — xcsh_trusted_ca_list / 9351258b0821 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-14f7aa6e66c43b413082463705a4779f1a4e8d262d54eafcd209e0498388a97a"></a>

## Next pages — xcsh_trusted_ca_list / 9351258b0821 / 6

- [Property reference](../guides/data-sources--trusted_ca_list--reference--group-001.md#canonical-dcb8d54ef3f35aabca505c635ade0b5d8ebdf23a0c754eba510299b4cac35f27)
- [Examples](../guides/data-sources--trusted_ca_list--examples--group-001.md#canonical-c5f1bedb8562795d2af752ebb622617996bd72502544fb62b1642964f38af908)
