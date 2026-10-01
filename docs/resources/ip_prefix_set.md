---
page_title: "xcsh_ip_prefix_set landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ip_prefix_set landing."
---

# xcsh_ip_prefix_set landing

<a id="canonical-1122102130313221-3130002203220321-3333000220303121-2220030010230000-3210200033230023-2332313003100231-3300010320213331-0132033303002232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312100203021102-0223201033210232-0213121012213122-0213020123031013-1232111220332120-3313131330013302-1110330102113330-0100102301002222"></a>

## xcsh_ip_prefix_set — xcsh_ip_prefix_set / 300100120231 / 2

Breadcrumbs:

- xcsh_ip_prefix_set

Manages ip\_prefix\_set creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-1231112201022121-0122011302332132-3213020330002323-0220302202032100-0323312103013011-2222300210223112-3200210212121333-2232332313322220"></a>

## Prerequisites — xcsh_ip_prefix_set / 300100120231 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2112301033032202-0213210222021321-0221110122012110-1201311102031310-0003030310220011-3323123110030110-0203321120120012-2311111201022123"></a>

## Minimal configuration — xcsh_ip_prefix_set / 300100120231 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# IPPrefixSet Resource Example
# Manages ip_prefix_set creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic IPPrefixSet configuration
resource "xcsh_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "staging"
}
```

<a id="canonical-3300233322003332-3011300231020010-2000032301130311-0020133130200223-2130303123121113-2000203233123230-2123331320020221-2010310110133111"></a>

## Root configuration — xcsh_ip_prefix_set / 300100120231 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3112331032201200-0230022200233321-3231112131130212-3232002101230122-3121010231012112-3003221133101221-3232313310320100-3231103032201210"></a>

## Next pages — xcsh_ip_prefix_set / 300100120231 / 6

- [Property reference](../guides/resources--ip_prefix_set--reference--group-001.md#canonical-3210033133131330-2313001323000331-2132220001132203-0211120213003001-0320122300001010-2210130131231100-2130011321312131-1022023020133221)
- [Examples](../guides/resources--ip_prefix_set--examples--group-001.md#canonical-1212030001113112-0332032221030101-1222301101333121-2213321023322222-1230232013222230-0120233121333323-3130000133022003-0021221230111132)
- [Import](../guides/resources--ip_prefix_set--lifecycle--group-001.md#canonical-0213010321223311-2321231203103123-1011313301132112-1021202301302033-1002123211303101-1100313122301010-1230020032132330-3320212231030012)
- [Timeouts](../guides/resources--ip_prefix_set--lifecycle--group-001.md#canonical-0223223131301023-0303231212012101-3223311232132122-3113131111123033-1020032222231013-2322310103212012-1311331020130203-2330033301312300)
