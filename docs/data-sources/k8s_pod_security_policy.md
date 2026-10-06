---
page_title: "xcsh_k8s_pod_security_policy"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_policy."
---

# xcsh_k8s_pod_security_policy

<a id="canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_k8s_pod_security_policy

Reads Kubernetes pod security policy information from F5 Distributed Cloud.

<a id="canonical-1332122312021301-3212210000311121-0332000120000102-2323321132320123-3330130311213013-2203220131023102-0311320020012000-1031232322303111"></a>

### Prerequisites for `xcsh_k8s_pod_security_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0020333311202213-3113300333321332-2303330211020111-0033210230303311-1131213102133220-0202003233203033-2100222301302121-0112101332333122"></a>

### Minimal configuration for `xcsh_k8s_pod_security_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# K8SPodSecurityPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing K8SPodSecurityPolicy by name
data "xcsh_k8s_pod_security_policy" "example" {
  name      = "example-k8s-pod-security-policy"
  namespace = "staging"
}

output "k8s_pod_security_policy_id" {
  value = data.xcsh_k8s_pod_security_policy.example.id
}
```

<a id="canonical-0322032322303221-1321302013032132-0110021312130132-2301130031232121-3022121030011222-0231321131212212-0113331233321323-0100301021330013"></a>

### Root configuration for `xcsh_k8s_pod_security_policy`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0121231331003031-2113102002311233-3303202332222031-3302113131303223-3113130321210131-3211330133002023-0002111022300101-3021102112203100"></a>

### Explore this collection for `xcsh_k8s_pod_security_policy`

- [Property reference](../guides/data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [Examples](../guides/data-sources--k8s_pod_security_policy--examples--group-001.md#canonical-2110202001110120-2323300313313223-2232030000312021-3323313000121302-0222300200131213-3010002102113022-3121001230220320-3113231121222201)
