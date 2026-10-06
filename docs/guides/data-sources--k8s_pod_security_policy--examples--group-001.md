---
page_title: "xcsh_k8s_pod_security_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_policy examples."
---

# xcsh_k8s_pod_security_policy examples

<a id="canonical-2110202001110120-2323300313313223-2232030000312021-3323313000121302-0222300200131213-3010002102113022-3121001230220320-3113231121222201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- Examples

<a id="canonical-1322332200210211-3031120330130003-3032111020122220-0213223110002311-0023202331323001-1301121110301133-0221031030021320-0230103003021021"></a>

### Complete configurations for `xcsh_k8s_pod_security_policy`

- [Data source](data-sources--k8s_pod_security_policy--examples--group-001.md#canonical-3232233311113333-1301121233203113-1311330100220222-3321230200332022-2330232213110220-0231030131021333-3300301112003201-0113032212200122): valid configuration.

<a id="canonical-3232233311113333-1301121233203113-1311330100220222-3321230200332022-2330232213110220-0231030131021333-3300301112003201-0113032212200122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_k8s_pod_security_policy](../data-sources/k8s_pod_security_policy.md#canonical-0312330221313303-3021300020322001-1013331231110100-0232023322032212-1231200233233323-2321213231200101-3231112101331221-3232122223123230)
- [Examples](data-sources--k8s_pod_security_policy--examples--group-001.md#canonical-2110202001110120-2323300313313223-2232030000312021-3323313000121302-0222300200131213-3010002102113022-3121001230220320-3113231121222201)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_k8s_pod_security_policy/data-source.tf`; digest `sha256:466f0cfc6eb4d12538a2c33e13b91f9270feaa389452e6596f117ab99a547c9b`.

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
