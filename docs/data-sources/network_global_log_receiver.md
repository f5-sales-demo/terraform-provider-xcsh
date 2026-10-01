---
page_title: "xcsh_network_global_log_receiver landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_global_log_receiver landing."
---

# xcsh_network_global_log_receiver landing

<a id="canonical-0211012321131230-2013331210022130-2003210011323120-3212212113100323-3012123311310320-2030302302201323-1222230131301120-2020212031212211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013022210233321-1032203021233313-2303330212222011-2032111212233310-1203223331201310-1232230220233320-1300220030032221-2033121033320203"></a>

## xcsh_network_global_log_receiver — xcsh_network_global_log_receiver / 230031111133 / 2

Breadcrumbs:

- xcsh_network_global_log_receiver

Global Log Receiver destinations. Published source entries mix CIDRs and individual IPv4 addresses.
Values are bundled from the pinned OpenAPI release; this data source performs no network request.
Ports and traffic direction are not encoded in the manifest.

<a id="canonical-3001202301103331-0122002302132001-1101131330310000-0321200231113123-2322113213232333-3322311303000221-1300012203032222-3232031230310210"></a>

## Prerequisites — xcsh_network_global_log_receiver / 230031111133 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2331122230133301-1232030101213010-1233130111322111-2310020231232130-2111301333212002-0313210103232331-2233030303232202-2223010020201302"></a>

## Minimal configuration — xcsh_network_global_log_receiver / 230031111133 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_global_log_receiver" "receivers" {}

# This example chooses TLS syslog on TCP 6514. The manifest supplies only
# destinations; choose the port required by the configured log receiver.
output "tls_syslog_egress" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 6514
    destinations = data.xcsh_network_global_log_receiver.receivers.cidr_blocks
  }
}
```

<a id="canonical-1132301322123021-0220230101011300-1302203231222200-3301112230233301-1132332132103023-1111232210121000-0111311022010312-3030122313213003"></a>

## Root configuration — xcsh_network_global_log_receiver / 230031111133 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-2101000333301303-2301113213221033-3332221322020303-1331330233032312-3233000330122231-0000213223311332-1033133110033032-2011000200033113"></a>

## Next pages — xcsh_network_global_log_receiver / 230031111133 / 6

- [Property reference](../guides/data-sources--network_global_log_receiver--reference--group-001.md#canonical-1302231310000103-2202232011203310-1003230012120012-3310013102103212-0201101202113030-3013331321011231-2133210011232031-2132201131331033)
- [Examples](../guides/data-sources--network_global_log_receiver--examples--group-001.md#canonical-1202301023031320-0303231001313120-3320311200012201-0033111103031111-3021010033100212-0120003332131331-3000011313101011-0011210102011131)
