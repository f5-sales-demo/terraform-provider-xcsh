---
page_title: "xcsh_k8s_pod_security_admission"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_admission."
---

# xcsh_k8s_pod_security_admission

<a id="canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_k8s_pod_security_admission

Reads Kubernetes Pod Security Admission information from F5 Distributed Cloud.

<a id="canonical-0012220131300032-0003233322210200-3123122303032223-3033110332321312-2333121120102210-1202030313023311-3201102110213201-3111113000330003"></a>

### Prerequisites for `xcsh_k8s_pod_security_admission`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0112002313201133-0222302202133313-0232030303023102-3220133311021233-2203100100202131-2210301323310331-1123012231110103-1310110023001113"></a>

### Minimal configuration for `xcsh_k8s_pod_security_admission`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SPodSecurityAdmission Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SPodSecurityAdmission by name
data "xcsh_k8s_pod_security_admission" "example" {
  name      = "example-k8s-pod-security-admission"
  namespace = "system"
}

output "k8s_pod_security_admission_id" {
  value = data.xcsh_k8s_pod_security_admission.example.id
}
```

<a id="canonical-3332311023313020-3223300231103201-0032132122101330-1222210100220211-3133002222000323-3032033012312011-3221100321013022-1120230222030101"></a>

### Root configuration for `xcsh_k8s_pod_security_admission`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-3101022020023222-2232300120221013-3300330001012000-3031331211102103-2312312101230332-1103102300220003-2030300122122001-3311212230110112"></a>

### Explore this collection for `xcsh_k8s_pod_security_admission`

- [Property reference](../guides/data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121)
- [Examples](../guides/data-sources--k8s_pod_security_admission--examples--group-001.md#canonical-1122201131203101-2310301330323322-1121021003103032-3131102302302102-0020121120231201-1230321232012301-2203132011330320-3013003012031300)
