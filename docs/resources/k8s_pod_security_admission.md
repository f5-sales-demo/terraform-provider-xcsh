---
page_title: "xcsh_k8s_pod_security_admission"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_admission."
---

# xcsh_k8s_pod_security_admission

<a id="canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_k8s_pod_security_admission

Manages k8s\_pod\_security\_admission will create the object in the storage backend in F5
Distributed Cloud.

<a id="canonical-2310011233323001-0303012221002330-2002010310222221-3301320220222133-2121030331211131-1202322033213232-2201102201301223-2132120023032031"></a>

### Prerequisites for `xcsh_k8s_pod_security_admission`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1132123203323123-3010012232202320-0021113312002010-3313101021200230-2222101201322320-3203312331011300-2312302333321233-1020012333231112"></a>

### Minimal configuration for `xcsh_k8s_pod_security_admission`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SPodSecurityAdmission Resource Example
# Manages k8s_pod_security_admission will create the object in the storage backend in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic K8SPodSecurityAdmission configuration
resource "xcsh_k8s_pod_security_admission" "example" {
  name      = "example-k8s-pod-security-admission"
  namespace = "system"
}
```

<a id="canonical-2330300100002013-0221231101032103-2221202323332023-1130320333111322-3221011223101123-2032203303333322-3330120031122001-3023033203030230"></a>

### Root configuration for `xcsh_k8s_pod_security_admission`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-0301023200113020-2332033322122302-1311300232010312-3012211111022122-2112320203103200-1302302323031103-0103121120321133-0222022333201320"></a>

### Explore this collection for `xcsh_k8s_pod_security_admission`

- [Property reference](../guides/resources--k8s_pod_security_admission--reference--group-001.md#canonical-2222333310123013-1221301313132132-3120203303322302-2023312113233133-3221130101113030-0021200233133103-2011300211002132-0212001303010330)
- [Examples](../guides/resources--k8s_pod_security_admission--examples--group-001.md#canonical-3303212110021111-0310330201010013-3130321323000001-2103313102022233-1213301323300113-2112033122032222-1032310313122120-3020321110323322)
- [Import](../guides/resources--k8s_pod_security_admission--lifecycle--group-001.md#canonical-1332322211301202-3112300032013110-3333012330223211-2030001100132313-1030222133313333-3012033302200111-0012121021232011-3133300211112002)
- [Timeouts](../guides/resources--k8s_pod_security_admission--lifecycle--group-001.md#canonical-2212300120320230-3120103030130113-0231311023023312-1233230202100000-3022100210320223-1300113310302113-0212331330311010-0002130132010200)
