---
page_title: "xcsh_protocol_policer examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_policer examples."
---

# xcsh_protocol_policer examples

<a id="canonical-1213220033330202-0030031031311012-3330210230023311-0231313332003100-3010300113001220-3320131123012020-0310113102113111-0012013123232223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-0333320001113103-2301232212333300-2111200131000031-2123222220130110-2130232032101120-0000213212312330-3312310223000320-3003231321101232)
- Examples

<a id="canonical-3222120102212031-2013001330010110-2003132003223331-3013120220100210-0102103202200300-2100313321210222-3000213322100112-2202301101103013"></a>

### Complete configurations for `xcsh_protocol_policer`

- [Data source](data-sources--protocol_policer--examples--group-001.md#canonical-3022121233330123-2320113200101313-3122301123313301-2302322123030003-0122010101021121-0032023223210113-2201130230232121-2132031103112132): valid configuration.

<a id="canonical-3022121233330123-2320113200101313-3122301123313301-2302322123030003-0122010101021121-0032023223210113-2201130230232121-2132031103112132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_protocol_policer](../data-sources/protocol_policer.md#canonical-0333320001113103-2301232212333300-2111200131000031-2123222220130110-2130232032101120-0000213212312330-3312310223000320-3003231321101232)
- [Examples](data-sources--protocol_policer--examples--group-001.md#canonical-1213220033330202-0030031031311012-3330210230023311-0231313332003100-3010300113001220-3320131123012020-0310113102113111-0012013123232223)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_protocol_policer/data-source.tf`; digest `sha256:fc068c58763daba5264bf1b0ce09146319c11b6d3fcae995dc7232b3a411a2d4`.

```terraform
# ProtocolPolicer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtocolPolicer by name
data "xcsh_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "system"
}

output "protocol_policer_id" {
  value = data.xcsh_protocol_policer.example.id
}
```
