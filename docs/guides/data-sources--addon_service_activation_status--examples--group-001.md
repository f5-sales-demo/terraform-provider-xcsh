---
page_title: "xcsh_addon_service_activation_status examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_addon_service_activation_status examples."
---

# xcsh_addon_service_activation_status examples

<a id="canonical-3102210133321200-1022113133123301-2122312202010112-3132330133131321-2223310303210120-2201032100212110-1130022201033021-2220303030101032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001131223030213-3321310030233123-2211011231231230-0020000113333213-3103022132323012-1201003210000330-3320221211232213-3331101202013220"></a>

## Examples — Examples / 312322213002 / 2

Breadcrumbs:

- [xcsh_addon_service_activation_status](../data-sources/addon_service_activation_status.md#canonical-3102032331212202-3203311211020103-1033023113322232-1210322102130301-1301000332023311-0302213213132333-0001300202010110-3120132022110012)
- Examples

<a id="canonical-0002333210002120-3101021101112331-3233122320032233-3001310200111313-0223210331000022-0110000131312022-2311332023223330-0211230302301233"></a>

## Complete configurations — Examples / 312322213002 / 3

- [Data source](data-sources--addon_service_activation_status--examples--group-001.md#canonical-3213023131322221-0101000021122220-0110100212033121-1332202133222033-3123311202220332-0112123331320111-1233223210022300-2211121222130212): valid configuration.

<a id="canonical-3202303031113032-0013220010133112-2132110233331132-1321232113231001-2322010331303332-3022103322102033-0213212001301131-2120012110123232"></a>

## Next pages — Examples / 312322213002 / 4

- [Data source](data-sources--addon_service_activation_status--examples--group-001.md#canonical-3213023131322221-0101000021122220-0110100212033121-1332202133222033-3123311202220332-0112123331320111-1233223210022300-2211121222130212)
- [xcsh_addon_service_activation_status](../data-sources/addon_service_activation_status.md#canonical-3102032331212202-3203311211020103-1033023113322232-1210322102130301-1301000332023311-0302213213132333-0001300202010110-3120132022110012)

<a id="canonical-3213023131322221-0101000021122220-0110100212033121-1332202133222033-3123311202220332-0112123331320111-1233223210022300-2211121222130212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101102203210212-3032333311132033-3022312003120332-1233322300003031-1312132110020031-1221231112133331-1012232101223223-3003030200132222"></a>

## Data source — Data source / 310032203003 / 2

Breadcrumbs:

- [xcsh_addon_service_activation_status](../data-sources/addon_service_activation_status.md#canonical-3102032331212202-3203311211020103-1033023113322232-1210322102130301-1301000332023311-0302213213132333-0001300202010110-3120132022110012)
- [Examples](data-sources--addon_service_activation_status--examples--group-001.md#canonical-3102210133321200-1022113133123301-2122312202010112-3132330133131321-2223310303210120-2201032100212110-1130022201033021-2220303030101032)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_addon_service_activation_status/data-source.tf`; digest `sha256:73b6eb2a08b57df6e1279cca2e1bd3688fc9db605ed1feb8ab5a2418a8f37cc0`.

```terraform
# AddonServiceActivationStatus Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Check the tenant's Client-Side Defense subscription.
data "xcsh_addon_service_activation_status" "example" {
  addon_service = "f5xc-client-side-defense-standard"
}

output "addon_service_activation_state" {
  value = data.xcsh_addon_service_activation_status.example.state
}
```

<a id="canonical-2113320003021001-0013212300300203-2313313031211000-3322213220221022-2101200102012332-3323200132202222-1103100001233103-0222333220303303"></a>

## Next pages — Data source / 310032203003 / 3

- [Examples](data-sources--addon_service_activation_status--examples--group-001.md#canonical-3102210133321200-1022113133123301-2122312202010112-3132330133131321-2223310303210120-2201032100212110-1130022201033021-2220303030101032)
- [xcsh_addon_service_activation_status](../data-sources/addon_service_activation_status.md#canonical-3102032331212202-3203311211020103-1033023113322232-1210322102130301-1301000332023311-0302213213132333-0001300202010110-3120132022110012)
