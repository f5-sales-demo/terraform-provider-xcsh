---
page_title: "xcsh_certificate_chain examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate_chain examples."
---

# xcsh_certificate_chain examples

<a id="canonical-2232010112021122-0333033022233011-1212120302221000-1111333200302332-3020123202322231-3100022030003231-0110323221202211-2100233111030330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_certificate_chain](../resources/certificate_chain.md#canonical-3300311221202132-1102121323122311-1122132320001300-2323321020311113-1011300012003132-2010130221021000-0203210321321112-0332331331210322)
- Examples

<a id="canonical-1202131310110003-3310130002012330-3131330130233111-2030113303230300-2220113032323012-3032210101101033-1320330003331133-0122202001020221"></a>

### Complete configurations for `xcsh_certificate_chain`

- [Resource](resources--certificate_chain--examples--group-001.md#canonical-2310023133111213-3302123200210231-1100230003202222-2213202211121212-0221200000102112-2111311200202031-0013202200033210-1221312113320221): valid configuration.

<a id="canonical-2310023133111213-3302123200210231-1100230003202222-2213202211121212-0221200000102112-2111311200202031-0013202200033210-1221312113320221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_certificate_chain](../resources/certificate_chain.md#canonical-3300311221202132-1102121323122311-1122132320001300-2323321020311113-1011300012003132-2010130221021000-0203210321321112-0332331331210322)
- [Examples](resources--certificate_chain--examples--group-001.md#canonical-2232010112021122-0333033022233011-1212120302221000-1111333200302332-3020123202322231-3100022030003231-0110323221202211-2100233111030330)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_certificate_chain/resource.tf`; digest `sha256:b97518ab6ca31864cd9125be1aa4c371e7fb2f6d12ef9376188442d135d4d743`.

```terraform
# CertificateChain Resource Example
# Manages a Certificate Chain resource in F5 Distributed Cloud for certificate chain configuration for TLS.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CertificateChain configuration
resource "xcsh_certificate_chain" "example" {
  name      = "example-certificate-chain"
  namespace = "staging"

  certificate_url = "example-value"
}
```
