---
page_title: "xcsh_secret_policy"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_secret_policy."
---

# xcsh_secret_policy

<a id="canonical-1312021301020310-0131200100331301-3002003211003233-3010322032311123-0022212322233110-2113210022031113-0031130332112033-3113012311332111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_secret_policy

Owns a native Secret Policy with ordered inline ALLOW/DENY rules. Evaluation is FIRST\_MATCH (top to
bottom); no invented API algorithm field is sent. Deletion is soft: destroy remains pending while a
tombstone exists and never repeats DELETE for a marked policy. Import namespace/name; legacy child
references are normalized into inline rules. No secret or credential belongs in this resource.

<a id="canonical-2033101301212311-3003211312123100-3032123212023202-2201312333320021-0232212013302103-1323131301211133-0231002133330111-3001130200222311"></a>

### Prerequisites for `xcsh_secret_policy`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-1233121213313023-2010231213220330-1302220203221222-0211131212300201-2231300030001112-2002311123100101-0020302322311103-0002001013001231"></a>

### Minimal configuration for `xcsh_secret_policy`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.14.0"
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 15.8.3"
    }
  }
}

# Install the current published release containing xcsh_secret_policy.
# Keep the namespace independently owned. The application pod must carry app=demo-tls.
resource "xcsh_secret_policy" "application" {
  name                     = "demo-wingman-policy"
  namespace                = "demo-app"
  allow_f5xc               = true
  decrypt_cache_timeout    = "30s"
  rule_combining_algorithm = "FIRST_MATCH"
  rules = [
    { name = "allow-application", action = "ALLOW", client_selector = ["app=demo-tls"] },
    { name = "deny-all", action = "DENY" }
  ]
}
```

<a id="canonical-0222002113123231-0100333003320001-0233320220120011-3203013201312100-0211233222211120-1330231122000132-1203203031000211-1231313113023012"></a>

### Root configuration for `xcsh_secret_policy`

Required root properties: `name`, `namespace`, `rules`. Full root flags and choices appear in the property reference.

<a id="canonical-1010211033220103-2223131331103202-3001301233033233-1000021120102232-3111033310333321-2312211320122201-2221131230013330-0210312333031222"></a>

### Explore this collection for `xcsh_secret_policy`

- [Property reference](../guides/resources--secret_policy--reference--group-001.md#canonical-1213313330302120-1100122222322313-2323333222220332-2211122300300103-3232201121133200-0113220013210033-2322130013023011-1132310023231200)
- [Examples](../guides/resources--secret_policy--examples--group-001.md#canonical-2233130113121122-3201312031313030-0131110311312111-3303220122302031-1113010302111212-0230212120110222-0222113301021113-3000322000122310)
- [Import](../guides/resources--secret_policy--lifecycle--group-001.md#canonical-2111131021220031-2323002001201332-3203310300212013-2033202133131221-2323022112202023-3031312211011310-1303310232320032-3332003112313022)
