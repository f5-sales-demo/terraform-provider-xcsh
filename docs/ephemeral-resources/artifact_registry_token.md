---
page_title: "xcsh_artifact_registry_token"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_artifact_registry_token."
---

# xcsh_artifact_registry_token

<a id="canonical-1323301000023223-3012100013003333-1031230331233031-2332031031030022-0033001320123233-0303101023323201-3331031322333331-2330103332212301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_artifact_registry_token

Authentication credential for access control.

<a id="canonical-0132312021321102-2332032013302130-2112220222302102-2302223023322120-2332112331101223-0320223101303120-3233332201030123-2203011333030033"></a>

### Prerequisites for `xcsh_artifact_registry_token`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1001331212122233-2023213110203001-0212031021010322-1102123311020033-0030000230202330-1220112101132300-2303333100030132-1203132301101113"></a>

### Minimal configuration for `xcsh_artifact_registry_token`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ArtifactRegistryToken EphemeralResource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

ephemeral "xcsh_artifact_registry_token" "example" {
  namespace = "example-value"
}
```

<a id="canonical-3331313133031232-3331322332101103-0130111013220213-2213102311330211-1111000023023202-1313031013233301-3002232011311233-0113113330032222"></a>

### Root configuration for `xcsh_artifact_registry_token`

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0011330212311202-3020210312301323-0010030313333303-3111222320232112-1333200031233233-3323222100020232-3132221332110333-2132230023310230"></a>

### Explore this collection for `xcsh_artifact_registry_token`

- [Property reference](../guides/ephemeral-resources--artifact_registry_token--reference--group-001.md#canonical-2203313033322030-0113022202301111-0210200113320313-3110001101012332-3113033331112210-1332120310221123-3101023200110313-3312021311022023)
- [Examples](../guides/ephemeral-resources--artifact_registry_token--examples--group-001.md#canonical-0011313231110131-0332120013220332-0111300222332010-0023102120300323-0123110301313201-3222121230031110-3232221322111321-3103110203301113)
- [Lifecycle](../guides/ephemeral-resources--artifact_registry_token--lifecycle--group-001.md#canonical-0110331221331200-1213300003313222-3323030310312122-1030322232211301-2032112231002313-3030211103001311-3321022033312021-3221211010230000)
