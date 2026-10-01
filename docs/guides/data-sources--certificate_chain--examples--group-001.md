---
page_title: "xcsh_certificate_chain examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate_chain examples."
---

# xcsh_certificate_chain examples

<a id="canonical-f4afd07874bc514a268ecf295767366e70decfe868e724cc487b9b5f454ea748"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b9f474fa713f1ac63d40c62ba741f484d6d6712fd86eb0f52251df864c10223d"></a>

## Examples — Examples / 78cfbb4bd825 / 2

Breadcrumbs:

- [xcsh_certificate_chain](../data-sources/certificate_chain.md#canonical-ce3d2afc41637e97fe322a943a0e26b7ffd06d1052a2a27eea5dfed15f9134f9)
- Examples

<a id="canonical-142cd351fa8874f7115a44c4b056beaf0669d17634a139259c0ca5c6a6132c46"></a>

## Complete configurations — Examples / 78cfbb4bd825 / 3

- [Data source](data-sources--certificate_chain--examples--group-001.md#canonical-fb119be222b6347416b38935c1a24424afa0caab1de03cc4de0ae7fda247298e): valid configuration.

<a id="canonical-701596fd19c83f0205d7b621aab9510782f58b4c0d9aade81bede41570c6c351"></a>

## Next pages — Examples / 78cfbb4bd825 / 4

- [Data source](data-sources--certificate_chain--examples--group-001.md#canonical-fb119be222b6347416b38935c1a24424afa0caab1de03cc4de0ae7fda247298e)
- [xcsh_certificate_chain](../data-sources/certificate_chain.md#canonical-ce3d2afc41637e97fe322a943a0e26b7ffd06d1052a2a27eea5dfed15f9134f9)

<a id="canonical-fb119be222b6347416b38935c1a24424afa0caab1de03cc4de0ae7fda247298e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ea847672fe96baedb1bc4a9329d8e695c62a9efe2b4c331970e10f229437b180"></a>

## Data source — Data source / ba229bb2a0fa / 2

Breadcrumbs:

- [xcsh_certificate_chain](../data-sources/certificate_chain.md#canonical-ce3d2afc41637e97fe322a943a0e26b7ffd06d1052a2a27eea5dfed15f9134f9)
- [Examples](data-sources--certificate_chain--examples--group-001.md#canonical-f4afd07874bc514a268ecf295767366e70decfe868e724cc487b9b5f454ea748)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_certificate_chain/data-source.tf`; digest `sha256:af7fc7589403e5f8724fbb3e9b221e9a801d9fa6878ef97a8d7e1b6b86b1c396`.

```terraform
# CertificateChain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CertificateChain by name
data "xcsh_certificate_chain" "example" {
  name      = "example-certificate-chain"
  namespace = "staging"
}

output "certificate_chain_id" {
  value = data.xcsh_certificate_chain.example.id
}
```

<a id="canonical-b057762f8489df5ef88c8b3adbe116b1f03a5bc7fac8990a020819bb2384b94d"></a>

## Next pages — Data source / ba229bb2a0fa / 3

- [Examples](data-sources--certificate_chain--examples--group-001.md#canonical-f4afd07874bc514a268ecf295767366e70decfe868e724cc487b9b5f454ea748)
- [xcsh_certificate_chain](../data-sources/certificate_chain.md#canonical-ce3d2afc41637e97fe322a943a0e26b7ffd06d1052a2a27eea5dfed15f9134f9)
