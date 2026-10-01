---
page_title: "xcsh_addon_service_activation_status landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_addon_service_activation_status landing."
---

# xcsh_addon_service_activation_status landing

<a id="canonical-3102032331212202-3203311211020103-1033023113322232-1210322102130301-1301000332023311-0302213213132333-0001300202010110-3120132022110012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012110301112210-1303232331212230-0131010110101120-0103210232103333-2302223130020321-2323331310022233-0002013310131031-2201320022320322"></a>

## xcsh_addon_service_activation_status — xcsh_addon_service_activation_status / 222330212332 / 2

Breadcrumbs:

- xcsh_addon_service_activation_status

Checks the activation status of an F5 Distributed Cloud Addon Service.

Use this data source to determine if an addon service can be activated for your tenant and what the
current subscription state is.

\*\*Possible state values:\*\*

| State | Description | | --------------- | ---------------------------------------- | |
\`AS\_NONE\` | Default state, service not subscribed | | \`AS\_PENDING\` | Subscription request
pending activation | | \`AS\_SUBSCRIBED\` | Service is active and subscribed | | \`AS\_ERROR\` |
Subscription in error state |

<a id="canonical-3303202222032101-3212232300112333-0102131033123123-2112033213213101-2002310010020103-1133331031323210-1000311113023300-3033232220230103"></a>

## Prerequisites — xcsh_addon_service_activation_status / 222330212332 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2012223113013032-3001102033201131-0132120002322313-1133213113013202-1331311030321122-2103020300212333-0322121310212033-3331111013200021"></a>

## Minimal configuration — xcsh_addon_service_activation_status / 222330212332 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-1332010012303132-2111131122220031-0103013020221222-0202201223303323-3330331003213211-3103132131312332-1202231321022112-3210000213001021"></a>

## Root configuration — xcsh_addon_service_activation_status / 222330212332 / 5

Required root properties: `addon_service`. Full root flags and choices appear in the property reference.

<a id="canonical-1022223213231130-1022103203100032-3320100021221113-2103201221313020-3302333100013132-2200133132022103-3321300220113200-2022332102130020"></a>

## Next pages — xcsh_addon_service_activation_status / 222330212332 / 6

- [Property reference](../guides/data-sources--addon_service_activation_status--reference--group-001.md#canonical-3310012331220020-2113213123322210-1321301222103133-2120010110020003-1233221302030322-0221010010110101-3101032013202133-0120012300012123)
- [Examples](../guides/data-sources--addon_service_activation_status--examples--group-001.md#canonical-3102210133321200-1022113133123301-2122312202010112-3132330133131321-2223310303210120-2201032100212110-1130022201033021-2220303030101032)
