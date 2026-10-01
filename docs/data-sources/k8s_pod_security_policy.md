---
page_title: "xcsh_k8s_pod_security_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_policy landing."
---

# xcsh_k8s_pod_security_policy landing

<a id="canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332122312021301-3212210000311121-0332000120000102-2323321132320123-3330130311213013-2203220131023102-0311320020012000-1031232322303111"></a>

## xcsh_k8s_pod_security_policy — xcsh_k8s_pod_security_policy / 101120310233 / 2

Breadcrumbs:

- xcsh_k8s_pod_security_policy

Manages k8s\_pod\_security\_policy will create the object in the storage backend for namespace
metadata.namespace in F5 Distributed Cloud.

<a id="canonical-0020333311202213-3113300333321332-2303330211020111-0033210230303311-1131213102133220-0202003233203033-2100222301302121-0112101332333122"></a>

## Prerequisites — xcsh_k8s_pod_security_policy / 101120310233 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0322032322303221-1321302013032132-0110021312130132-2301130031232121-3022121030011222-0231321131212212-0113331233321323-0100301021330013"></a>

## Minimal configuration — xcsh_k8s_pod_security_policy / 101120310233 / 4

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

<a id="canonical-0121231331003031-2113102002311233-3303202332222031-3302113131303223-3113130321210131-3211330133002023-0002111022300101-3021102112203100"></a>

## Root configuration — xcsh_k8s_pod_security_policy / 101120310233 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2331222313103131-3313030321032321-3021231023132030-1303122322133102-2001030011300103-3133303212100131-1110103133210213-1223100132223321"></a>

## Next pages — xcsh_k8s_pod_security_policy / 101120310233 / 6

- [Property reference](../guides/data-sources--k8s_pod_security_policy--reference--group-001.md#canonical-0120103030312320-2003223011001113-0111202210322101-2003220111322120-0000021103320302-3003201121110332-0200231110300213-1010033310223013)
- [Examples](../guides/data-sources--k8s_pod_security_policy--examples--group-001.md#canonical-2110202001110120-2323300313313223-2232030000312021-3323313000121302-0222300200131213-3010002102113022-3121001230220320-3113231121222201)
