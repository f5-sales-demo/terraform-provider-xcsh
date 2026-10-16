---
page_title: "xcsh_secret_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_secret_policy examples."
---

# xcsh_secret_policy examples

<a id="canonical-2233130113121122-3201312031313030-0131110311312111-3303220122302031-1113010302111212-0230212120110222-0222113301021113-3000322000122310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_secret_policy](../resources/secret_policy.md#canonical-1312021301020310-0131200100331301-3002003211003233-3010322032311123-0022212322233110-2113210022031113-0031130332112033-3113012311332111)
- Examples

<a id="canonical-2201302222312213-1103023232023113-1202321303210201-3030221122032333-1111303323021321-1002022332231311-0302002130222100-0031330231120133"></a>

### Complete configurations for `xcsh_secret_policy`

- [Resource](resources--secret_policy--examples--group-001.md#canonical-3331221202302033-1321203231013113-2131303202012320-2312133000103013-3010222033113123-2320221223301003-3301300302210021-0012212322100010): valid configuration.

<a id="canonical-3331221202302033-1321203231013113-2131303202012320-2312133000103013-3010222033113123-2320221223301003-3301300302210021-0012212322100010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_secret_policy](../resources/secret_policy.md#canonical-1312021301020310-0131200100331301-3002003211003233-3010322032311123-0022212322233110-2113210022031113-0031130332112033-3113012311332111)
- [Examples](resources--secret_policy--examples--group-001.md#canonical-2233130113121122-3201312031313030-0131110311312111-3303220122302031-1113010302111212-0230212120110222-0222113301021113-3000322000122310)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_secret_policy/resource.tf`; digest `sha256:f59eb002c8cb50585be56dfb7682ee9d3bd976a6ce27601c7cb63562b64c6c5f`.

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
