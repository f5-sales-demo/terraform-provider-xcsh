---
page_title: "xcsh_network_bot_defense examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_bot_defense examples."
---

# xcsh_network_bot_defense examples

<a id="canonical-02700dbe54bec0d993bff3ebb664b29b8813b46344c33b64034f20f39624ad6c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-78fca1a789807e884007bf8be1d5beb0f643af1a4614128dc2e10d1dc111ee52"></a>

## Examples — Examples / b7fda9f7d8a3 / 2

Breadcrumbs:

- [xcsh_network_bot_defense](../data-sources/network_bot_defense.md#canonical-0f1230987fc786b4cb7c4b344761f3dee154f819970a4e6a082efc66aa2b0e4d)
- Examples

<a id="canonical-b120520f0f5df824ec24af83fa0485183aee20c7b5fd5dbd3d221077ff10a8d8"></a>

## Complete configurations — Examples / b7fda9f7d8a3 / 3

- [Data source](data-sources--network_bot_defense--examples--group-001.md#canonical-94b4ef9d799618583484ff008eca6a25fc9ec52a05601c0451ec165f2061a14e): valid configuration.

<a id="canonical-74a0c07c68d08f156b7273f73a16eebd66f509f1e68889a5a01dbcd4962485b6"></a>

## Next pages — Examples / b7fda9f7d8a3 / 4

- [Data source](data-sources--network_bot_defense--examples--group-001.md#canonical-94b4ef9d799618583484ff008eca6a25fc9ec52a05601c0451ec165f2061a14e)
- [xcsh_network_bot_defense](../data-sources/network_bot_defense.md#canonical-0f1230987fc786b4cb7c4b344761f3dee154f819970a4e6a082efc66aa2b0e4d)

<a id="canonical-94b4ef9d799618583484ff008eca6a25fc9ec52a05601c0451ec165f2061a14e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df5229222239b9b9ad7bc21eba88fec7b48401cc83edd1f2c4ff57e9b4a0dcb0"></a>

## Data source — Data source / 0d2dcdb77a90 / 2

Breadcrumbs:

- [xcsh_network_bot_defense](../data-sources/network_bot_defense.md#canonical-0f1230987fc786b4cb7c4b344761f3dee154f819970a4e6a082efc66aa2b0e4d)
- [Examples](data-sources--network_bot_defense--examples--group-001.md#canonical-02700dbe54bec0d993bff3ebb664b29b8813b46344c33b64034f20f39624ad6c)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_bot_defense/data-source.tf`; digest `sha256:a428d9d820fed630e17a7faf15c103fb2b68704b5517319077da6fc62dccc081`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_bot_defense" "proxy" {}

# Configure these exact domains in a suffix-aware proxy or FQDN firewall.
output "bot_defense_https_proxy_rule" {
  value = {
    direction = "egress"
    protocol  = "tcp"
    port      = 443
    domains   = data.xcsh_network_bot_defense.proxy.domains
  }
}
```

<a id="canonical-5000266842796bcf8995fb1d12122d2c5681a8816755feb55a81293a5a3e5234"></a>

## Next pages — Data source / 0d2dcdb77a90 / 3

- [Examples](data-sources--network_bot_defense--examples--group-001.md#canonical-02700dbe54bec0d993bff3ebb664b29b8813b46344c33b64034f20f39624ad6c)
- [xcsh_network_bot_defense](../data-sources/network_bot_defense.md#canonical-0f1230987fc786b4cb7c4b344761f3dee154f819970a4e6a082efc66aa2b0e4d)
