---
page_title: "xcsh_http_loadbalancer lifecycle"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer lifecycle."
---

# xcsh_http_loadbalancer lifecycle

<a id="canonical-1000101212023023-0103011103132001-3301102330230310-1230010102120212-3223112021023333-2200311133300111-1031101210100332-2211331001010311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Import

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- Import

Import using the `namespace/name` identifier format.

```shell
terraform import xcsh_http_loadbalancer.example system/example
```

<a id="canonical-2232121223123231-2022121130300121-0103113012231013-1023001100002301-2121003313103030-1100033200202130-2111230110133210-3020020103231013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Timeouts

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- Timeouts

Configure the supported operation timeouts in the [timeouts](resources--http_loadbalancer--reference--group-027.md#canonical-0031000301301132-0002232123000132-3130011021001231-3113311013210310-0122300301203211-0213222010232211-0313220210221220-2120201133330322). Use Terraform duration strings such as `30m`.

<a id="canonical-2201010321200222-2110320331131013-1013211030201130-0130323333132231-1030203101130301-0030312322132221-0233000232333000-3101312320200212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Lifecycle

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- Lifecycle

Changing the loadbalancer_type selection between
[http](resources--http_loadbalancer--reference--group-018.md#canonical-3201330110311330-2123203220332301-2010310002232302-2200333212132031-2112023132010213-3001201001222010-0222333320131100-0013020102021323),
[https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102),
[https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220),
including an omitted selection, requires recreation and may interrupt service.
F5 Distributed Cloud cannot change this type selection in place.
Certificate rotation and other supported settings within the same selected type remain updates.

Terraform identifies the affected type blocks as **must be replaced**. Unknown block presence requires replacement when unchanged selection cannot be proven; unknown child settings alone do not. Recommendations do not establish API defaults.

Use `lifecycle { prevent_destroy = true }` to reject replacement before remote writes. Terraform controls replacement ordering: its default destroys before creating; `create_before_destroy` requests creation first, which may fail if XC requires a unique name. Plan a maintenance window or use a distinct name for a staged replacement.
