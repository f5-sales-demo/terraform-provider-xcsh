---
page_title: "xcsh_k8s_pod_security_policy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_policy."
---

# xcsh_k8s_pod_security_policy

<a id="canonical-3101303300000313-2201301322310120-1331000001222003-0130310023201122-2303211103023321-3201111210223031-0031002212103323-0301201300322112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_k8s_pod_security_policy

Manages k8s\_pod\_security\_policy will create the object in the storage backend for namespace
metadata.namespace in F5 Distributed Cloud.

<a id="canonical-1000101231301030-1231012010112031-3222312211312121-0113120230312123-0203221201103032-0231130231322220-3030302120011310-1101211012030130"></a>

### Prerequisites for `xcsh_k8s_pod_security_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1121111000031132-2311113320231202-0220321133022031-3211232130003331-1021122210122331-1133120010233010-3100103112303220-1023103313121131"></a>

### Minimal configuration for `xcsh_k8s_pod_security_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SPodSecurityPolicy Resource Example
# Manages k8s_pod_security_policy will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SPodSecurityPolicy configuration
resource "xcsh_k8s_pod_security_policy" "example" {
  name      = "example-k8s-pod-security-policy"
  namespace = "staging"
}
```

<a id="canonical-2031023122302200-2332300203012122-2222222233111012-3202030121310013-2301122233211100-3330210102102223-2110330010112032-2132012302113320"></a>

### Root configuration for `xcsh_k8s_pod_security_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3211201031121203-3302003212100011-2121310032021211-1203301330321233-2302220300103100-3210020320100122-3121113330030330-2013233030102102"></a>

### Explore this collection for `xcsh_k8s_pod_security_policy`

- [Property reference](../guides/resources--k8s_pod_security_policy--reference--group-001.md#canonical-0330220232122001-1120122233303200-3231303220302330-2020031222111112-3313031011220001-2301021232002313-2213100122102213-1220121112101333)
- [Examples](../guides/resources--k8s_pod_security_policy--examples--group-001.md#canonical-3333322100323002-0130321130311000-2221323221320232-2220202010130003-1033331121221112-2023311231013122-1000003231033210-3020133031320020)
- [Import](../guides/resources--k8s_pod_security_policy--lifecycle--group-001.md#canonical-1303131300310130-3302312132233323-0211121212103303-3103110122303033-0031320121012303-2021310330310011-2303123102231302-3030213013323020)
- [Timeouts](../guides/resources--k8s_pod_security_policy--lifecycle--group-001.md#canonical-2330222133302321-2111133331023331-0202012300010000-3000000312333211-3010310332333220-0232101223012030-3033323232300030-0203030331223232)
