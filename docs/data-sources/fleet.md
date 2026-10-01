---
page_title: "xcsh_fleet landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet landing."
---

# xcsh_fleet landing

<a id="canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331222323032330-3300201321212111-2312301333203001-0033123322033213-3210002110213330-0211201010101330-3211122101113232-1233212033232001"></a>

## xcsh_fleet — xcsh_fleet / 322022300220 / 2

Breadcrumbs:

- xcsh_fleet

Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud.

<a id="canonical-3313020303221021-3003032121321032-0323302210012011-0110210020103130-3002331020122103-3103223122301332-2330313201101113-3213221201133031"></a>

## Prerequisites — xcsh_fleet / 322022300220 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2131121113101132-1133222200313311-2311130210110200-2130230230112102-0131300212023100-3110220301013232-2003211302223000-2030320000300323"></a>

## Minimal configuration — xcsh_fleet / 322022300220 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Fleet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Fleet by name
data "xcsh_fleet" "example" {
  name      = "example-fleet"
  namespace = "staging"
}

output "fleet_id" {
  value = data.xcsh_fleet.example.id
}
```

<a id="canonical-2120003100031231-0100003000230323-1323132010213211-2020213020130122-2300223300121211-0200202223000031-1110300030130021-2122033011003030"></a>

## Root configuration — xcsh_fleet / 322022300220 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2222110022322010-0333332032312113-0101210202320031-2002313133202012-3021122111232211-2232331323103111-1311111130122231-3001333323310303"></a>

## Next pages — xcsh_fleet / 322022300220 / 6

- [Property reference](../guides/data-sources--fleet--reference--group-001.md#canonical-1100111312003010-3321033322231203-3323022200221200-0231312313122201-2212100012020313-0101321123001001-0322312212313331-2013003202313021)
- [Examples](../guides/data-sources--fleet--examples--group-001.md#canonical-1223103321010323-0231100030013302-1233302032200110-0210031301212220-2100122330323123-0310212333032022-2010022133220121-0100330002231213)
