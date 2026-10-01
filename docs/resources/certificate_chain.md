---
page_title: "xcsh_certificate_chain landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate_chain landing."
---

# xcsh_certificate_chain landing

<a id="canonical-3300311221202132-1102121323122311-1122132320001300-2323321020311113-1011300012003132-2010130221021000-0203210321321112-0332331331210322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311222232210100-3101221130023211-3113330130320203-3320023020103200-0221112312320322-1113020013212020-0312002223230021-0122112112133010"></a>

## xcsh_certificate_chain — xcsh_certificate_chain / 111322102311 / 2

Breadcrumbs:

- xcsh_certificate_chain

Manages a Certificate Chain resource in F5 Distributed Cloud for certificate chain configuration for
TLS.

<a id="canonical-2313301220213112-0230101033313003-3211102312102123-1002320211301300-0210012312021110-0120311133012221-0003202332330010-2031330202302221"></a>

## Prerequisites — xcsh_certificate_chain / 111322102311 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-3001132221001200-1200113032112300-1010210320220311-2122131203123033-2202321233211233-2121331102330221-0213121321032010-3211222202230010"></a>

## Minimal configuration — xcsh_certificate_chain / 111322102311 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-2013110331111320-3321002332112130-1331103031022333-1033223211113023-0110333220130220-1001113123322313-3100133113120120-1200110230201233"></a>

## Root configuration — xcsh_certificate_chain / 111322102311 / 5

Required root properties: `certificate_url`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1330222121101022-2130131233013322-3332320233210001-2322212221022102-0101121010221200-1031101210302323-0223013033230020-0100120130011110"></a>

## Next pages — xcsh_certificate_chain / 111322102311 / 6

- [Property reference](../guides/resources--certificate_chain--reference--group-001.md#canonical-3021131230120120-2001003003312102-0331033112303122-3121321212202103-3132210200101210-2100113131211103-3003002012030023-2121301110001321)
- [Examples](../guides/resources--certificate_chain--examples--group-001.md#canonical-2232010112021122-0333033022233011-1212120302221000-1111333200302332-3020123202322231-3100022030003231-0110323221202211-2100233111030330)
- [Import](../guides/resources--certificate_chain--lifecycle--group-001.md#canonical-1331002123100010-3112000330321013-2321013022131223-0103323211201210-1000130023312232-3002103200202023-3202203202221023-2101203311102111)
- [Timeouts](../guides/resources--certificate_chain--lifecycle--group-001.md#canonical-1123031312233011-3103213101020100-2003020210131200-1100100123203310-3223012032033333-3201333310230032-3000210222013300-0211123022312333)
