---
page_title: "xcsh_certificate_chain"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate_chain."
---

# xcsh_certificate_chain

<a id="canonical-3032033102223330-1001120313322113-3332030202222110-0322003202122313-3333310012310100-1102220222021332-3222113133323101-1133210103103321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_certificate_chain

Reads Certificate Chain information from F5 Distributed Cloud.

<a id="canonical-1323323222000011-2230133133003100-0032000222203032-0331122131201133-1202332302211101-1313313012121120-2002133323110020-1013003022200100"></a>

### Prerequisites for `xcsh_certificate_chain`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-1111100120220312-0031133111103001-3030312101232322-1202130212332102-1323131211003322-2121222312213032-1133020331300021-3302312122113123"></a>

### Minimal configuration for `xcsh_certificate_chain`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-0110111120001312-2313213033132022-2332021030102021-3110331323311112-1110120223230031-1220100310301222-2312021100322203-1001303212330020"></a>

### Root configuration for `xcsh_certificate_chain`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1200202022021013-2200030012103120-2312030033010302-2213200302033301-0302211100221001-0303130201220111-2300010221333301-3102123020113311"></a>

### Explore this collection for `xcsh_certificate_chain`

- [Property reference](../guides/data-sources--certificate_chain--reference--group-001.md#canonical-3121100312232202-0231200300013012-2200330002333301-1222103120032112-3203332013220113-1331112100302301-1010131310310303-2011022230112320)
- [Examples](../guides/data-sources--certificate_chain--examples--group-001.md#canonical-3310223331001320-1310233011011022-0212203230330221-1113121303121232-1300313230333220-1220321302103030-1020132321231133-1011103222131020)
