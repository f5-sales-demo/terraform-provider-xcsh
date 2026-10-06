---
page_title: "xcsh_addon_service_activation_status examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_addon_service_activation_status examples."
---

# xcsh_addon_service_activation_status examples

<a id="canonical-3102210133321200-1022113133123301-2122312202010112-3132330133131321-2223310303210120-2201032100212110-1130022201033021-2220303030101032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_addon_service_activation_status](../data-sources/addon_service_activation_status.md#canonical-3102032331212202-3203311211020103-1033023113322232-1210322102130301-1301000332023311-0302213213132333-0001300202010110-3120132022110012)
- Examples

<a id="canonical-2001131223030213-3321310030233123-2211011231231230-0020000113333213-3103022132323012-1201003210000330-3320221211232213-3331101202013220"></a>

### Complete configurations for `xcsh_addon_service_activation_status`

- [Data source](data-sources--addon_service_activation_status--examples--group-001.md#canonical-3213023131322221-0101000021122220-0110100212033121-1332202133222033-3123311202220332-0112123331320111-1233223210022300-2211121222130212): valid configuration.

<a id="canonical-3213023131322221-0101000021122220-0110100212033121-1332202133222033-3123311202220332-0112123331320111-1233223210022300-2211121222130212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

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
