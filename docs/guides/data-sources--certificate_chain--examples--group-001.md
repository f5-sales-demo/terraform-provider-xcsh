---
page_title: "xcsh_certificate_chain examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate_chain examples."
---

# xcsh_certificate_chain examples

<a id="canonical-3310223331001320-1310233011011022-0212203230330221-1113121303121232-1300313230333220-1220321302103030-1020132321231133-1011103222131020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_certificate_chain](../data-sources/certificate_chain.md#canonical-3032033102223330-1001120313322113-3332030202222110-0322003202122313-3333310012310100-1102220222021332-3222113133323101-1133210103103321)
- Examples

<a id="canonical-2321331013103322-1301033301223012-0331100030120223-2213100133102010-3112311213010233-3120123223003311-0202110131332012-1030010002020331"></a>

### Complete configurations for `xcsh_certificate_chain`

- [Data source](data-sources--certificate_chain--examples--group-001.md#canonical-3323010121233202-0202231203101310-0112230320210311-3001220210100210-2233220030222223-0131320003303010-3132002232133331-2202101302212032): valid configuration.

<a id="canonical-3323010121233202-0202231203101310-0112230320210311-3001220210100210-2233220030222223-0131320003303010-3132002232133331-2202101302212032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_certificate_chain](../data-sources/certificate_chain.md#canonical-3032033102223330-1001120313322113-3332030202222110-0322003202122313-3333310012310100-1102220222021332-3222113133323101-1133210103103321)
- [Examples](data-sources--certificate_chain--examples--group-001.md#canonical-3310223331001320-1310233011011022-0212203230330221-1113121303121232-1300313230333220-1220321302103030-1020132321231133-1011103222131020)
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
