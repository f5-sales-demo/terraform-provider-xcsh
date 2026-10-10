---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-2021112022303101-0201220131321121-0131213002211122-0320300111102212-2320113321321222-2213303122223310-2310233231023111-0300002201113101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-3330032300013021-0303112230203221-1300231013211111-3122331131122313-2200300213320121-1120032012310022-0211330323032100-3121202331213103"></a>

Type: `"single"`. Computed.

IPV6AutoConfigType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

<a id="canonical-0133011103021002-2311202231103023-2012230020102320-1221223023331311-3121213010211202-0212311203211210-0003033231102102-0212001000311130"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config`

- [host](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1223323130013211-3200122303110133-2212100320003201-3010210300121132-1301213232323102-2312230011102032-0130133002032133-2320213012003312): complete subsection reference.

- [router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2231302202330301-0332223210023000-1222331123230011-0031232021100133-2233332123101113-0113211010002201-3120103122330300-1233312032013321): complete subsection reference.

<a id="canonical-1223323130013211-3200122303110133-2212100320003201-3010210300121132-1301213232323102-2312230011102032-0130133002032133-2320213012003312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2021112022303101-0201220131321121-0131213002211122-0320300111102212-2320113321321222-2213303122223310-2310233231023111-0300002201113101)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-2323231122013200-0303023012332201-3330110330303321-3213023220112002-1013212300201111-3311331223330133-3231002001203221-0022333130202323"></a>

Type: `["object", {}]`. Computed.

Hostname or IP address of the target server.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231302202330301-0332223210023000-1222331123230011-0031232021100133-2233332123101113-0113211010002201-3120103122330300-1233312032013321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2021112022303101-0201220131321121-0131213002211122-0320300111102212-2320113321321222-2213303122223310-2310233231023111-0300002201113101)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-0223210333333012-3033121032012211-3101323133230211-2013030130332003-3021030203001111-2203132323023110-1002302202130311-2322023312011233"></a>

Type: `"single"`. Computed.

IPV6AutoConfigRouterType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

<a id="canonical-2213221223223213-1011201212303130-0203100333012232-2011231211101011-1332202030313300-3113133013122211-0123330022330202-2111021113320133"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router`

- [dns_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-3311320311000323-0032323113323202-3323212021103011-1123130312033332-2001020021321010-1103010001011032-1023220323133202-3031220032331120): complete subsection reference.

<a id="canonical-2233022220333103-0110203110212122-3122122310010322-2030102232233131-2302020210121102-2233223322013232-2210221232031303-3121230330110112"></a>

<a id="canonical-3202133321000323-2011131212023103-1322100000122333-1111130122320001-0031101303312003-1103332222120032-1121330230301212-0332102222301001"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.network_prefix` property

Type: `"string"`. Computed.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "pattern": ".*::/64$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  }
}
```

- [stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2010102121210030-2020230123133211-2332102200301123-0023312013000223-2012000332033021-3002013310212023-3121110213303221-2032000201110002): complete subsection reference.

<a id="canonical-3311320311000323-0032323113323202-3323212021103011-1123130312033332-2001020021321010-1103010001011032-1023220323133202-3031220032331120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2021112022303101-0201220131321121-0131213002211122-0320300111102212-2320113321321222-2213303122223310-2310233231023111-0300002201113101)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2231302202330301-0332223210023000-1222331123230011-0031232021100133-2233332123101113-0113211010002201-3120103122330300-1233312032013321)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-1013211233312001-1212312133220121-0222132130202331-2111110002322200-2131220332101302-2232221133121022-2112120200101222-1230031311332003"></a>

Type: `"single"`. Computed.

IPV6DnsConfig.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dns_choice": "[\"configured_list\",\"local_dns\"]"
}
```

<a id="canonical-2000322321132032-2121112122201103-2230112110023323-2311310113010310-1330002100233213-1013121322132103-0311222220030203-2201033330333003"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config`

- [configured_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2102020220322111-1333122132300310-0103131322020202-2022312311330012-2032111102310031-1310133301300332-1220101121031303-3021010021203300): complete subsection reference.

- [local_dns](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1210301201233123-1110231231203120-1100200322201210-2120111203300310-2220001103120330-0100032202210111-0133132032302312-0112023301311221): complete subsection reference.

<a id="canonical-2102020220322111-1333122132300310-0103131322020202-2022312311330012-2032111102310031-1310133301300332-1220101121031303-3021010021203300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2021112022303101-0201220131321121-0131213002211122-0320300111102212-2320113321321222-2213303122223310-2310233231023111-0300002201113101)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2231302202330301-0332223210023000-1222331123230011-0031232021100133-2233332123101113-0113211010002201-3120103122330300-1233312032013321)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-3311320311000323-0032323113323202-3323212021103011-1123130312033332-2001020021321010-1103010001011032-1023220323133202-3031220032331120)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-1213022121333032-0123130022332221-3323120222320210-1223013302220330-0012231001021231-0323232020003201-1333102030313232-2222303112100033"></a>

Type: `"single"`. Computed.

IPV6DnsList.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0000032133331222-2332303211030022-1113103300303032-2213102020000333-2102123102203320-2123023330323201-1332330033330012-0012220333200323"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-2120021030133232-2231301201003022-3013103331310032-3321330223220220-3223201213130333-3130131213310020-0102110320213312-2201213031222331"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

Type: `["list", "string"]`. Computed.

List of IPv6 Addresses acting as DNS servers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1210301201233123-1110231231203120-1100200322201210-2120111203300310-2220001103120330-0100032202210111-0133132032302312-0112023301311221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2021112022303101-0201220131321121-0131213002211122-0320300111102212-2320113321321222-2213303122223310-2310233231023111-0300002201113101)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2231302202330301-0332223210023000-1222331123230011-0031232021100133-2233332123101113-0113211010002201-3120103122330300-1233312032013321)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-3311320311000323-0032323113323202-3323212021103011-1123130312033332-2001020021321010-1103010001011032-1023220323133202-3031220032331120)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2222030103331012-2030012301110302-0323210312110122-3110333120013032-2121323012020111-2210201202123011-0332001013110310-3313311003130311"></a>

Type: `"single"`. Computed.

IPV6LocalDnsAddress.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

<a id="canonical-1131332331103312-2030101220122200-0122232202122213-3013113303203112-2033111330332023-2133131233123302-2133002231322231-3020312010330133"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-3312233001323003-3101132222032123-3020321221012322-1222311132013303-2031212331023113-3100133033203011-3021232223021322-2233311100310203"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

- [first_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1103012300020202-3312322302031202-1022322122320020-3122310033130223-1313132010001223-2313102311130120-0303111322133133-0002100210003023): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-3031301201133013-1131002022300220-1031112201131201-1330003312123300-1122032121012321-2110011000132101-3113303321020132-2031323212100212): complete subsection reference.

<a id="canonical-1103012300020202-3312322302031202-1022322122320020-3122310033130223-1313132010001223-2313102311130120-0303111322133133-0002100210003023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2021112022303101-0201220131321121-0131213002211122-0320300111102212-2320113321321222-2213303122223310-2310233231023111-0300002201113101)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2231302202330301-0332223210023000-1222331123230011-0031232021100133-2233332123101113-0113211010002201-3120103122330300-1233312032013321)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-3311320311000323-0032323113323202-3323212021103011-1123130312033332-2001020021321010-1103010001011032-1023220323133202-3031220032331120)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1210301201233123-1110231231203120-1100200322201210-2120111203300310-2220001103120330-0100032202210111-0133132032302312-0112023301311221)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-0021102132002030-3112030030300312-1213133323203000-1311313311223131-3320220130213111-3211313311211323-2101030030232011-3222131202030122"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031301201133013-1131002022300220-1031112201131201-1330003312123300-1122032121012321-2110011000132101-3113303321020132-2031323212100212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2021112022303101-0201220131321121-0131213002211122-0320300111102212-2320113321321222-2213303122223310-2310233231023111-0300002201113101)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2231302202330301-0332223210023000-1222331123230011-0031232021100133-2233332123101113-0113211010002201-3120103122330300-1233312032013321)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-3311320311000323-0032323113323202-3323212021103011-1123130312033332-2001020021321010-1103010001011032-1023220323133202-3031220032331120)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1210301201233123-1110231231203120-1100200322201210-2120111203300310-2220001103120330-0100032202210111-0133132032302312-0112023301311221)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-2321003323220333-1031213331020101-3301212331232231-1303232303130033-1003220221031010-2213123121001021-3222312111232203-1312303233210320"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010102121210030-2020230123133211-2332102200301123-0023312013000223-2012000332033021-3002013310212023-3121110213303221-2032000201110002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2021112022303101-0201220131321121-0131213002211122-0320300111102212-2320113321321222-2213303122223310-2310233231023111-0300002201113101)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2231302202330301-0332223210023000-1222331123230011-0031232021100133-2233332123101113-0113211010002201-3120103122330300-1233312032013321)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-3021332030122203-3013013301010100-2102330110122212-3130300223102300-0320100232213321-1323310110230222-0111031013001122-1232010123023332"></a>

Type: `"single"`. Computed.

DHCPIPV6 Stateful Server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

<a id="canonical-3110021312133120-1000330330323330-3222202213120122-0320211231032003-3122201300133303-3211313230103030-2321023130232302-2332010033333120"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2101022031220113-0322322000100000-3313001030030101-1200302221123302-2233121222111113-3020201221130123-1111330211213233-0120023323232123): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0201233202320100-3312002132331123-3112133012223231-2222323130200210-3032030210322332-0132221323131112-2332020231103110-3030210113003233): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2010132122000111-0200232110303303-0330111320313312-2330201212200103-0203310320332233-0121310313012122-2322233210313002-2203200230202212): complete subsection reference.

<a id="canonical-3230322133002002-2121201321223032-1002322123112203-1231001031322303-1200232110323012-0111311310311031-0213022300302102-1001332331013303"></a>

<a id="canonical-1121303230010323-2302111211123333-2222302121032031-1300320103202231-0332110330131132-0303021021203333-3313210103100012-1233220322022013"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

Type: `["map", "string"]`. Computed.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
    },
    "category": "discovery",
    "constraintType": "map",
    "crossEntry": {
      "uniqueValues": true
    },
    "deterministic": true,
    "keys": {
      "format": "mac-address",
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.mac": "true",
      "ves.io.schema.rules.map.max_pairs": "128",
      "ves.io.schema.rules.map.unique_values": "true",
      "ves.io.schema.rules.map.values.string.ipv6": "true"
    },
    "values": {
      "format": "ipv6",
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0121311333000230-1102120020020113-3301202222113203-0130110233310201-1300011132113200-3030302112302203-3213011020103002-1131011023231002): complete subsection reference.

<a id="canonical-2101022031220113-0322322000100000-3313001030030101-1200302221123302-2233121222111113-3020201221130123-1111330211213233-0120023323232123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2021112022303101-0201220131321121-0131213002211122-0320300111102212-2320113321321222-2213303122223310-2310233231023111-0300002201113101)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2231302202330301-0332223210023000-1222331123230011-0031232021100133-2233332123101113-0113211010002201-3120103122330300-1233312032013321)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2010102121210030-2020230123133211-2332102200301123-0023312013000223-2012000332033021-3002013310212023-3121110213303221-2032000201110002)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-1111223320303230-2120200223333013-3231110322021230-1122311023301130-1323210231203211-2310001030320022-1212020100131030-0113203223321221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from end.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201233202320100-3312002132331123-3112133012223231-2222323130200210-3032030210322332-0132221323131112-2332020231103110-3030210113003233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2021112022303101-0201220131321121-0131213002211122-0320300111102212-2320113321321222-2213303122223310-2310233231023111-0300002201113101)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2231302202330301-0332223210023000-1222331123230011-0031232021100133-2233332123101113-0113211010002201-3120103122330300-1233312032013321)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2010102121210030-2020230123133211-2332102200301123-0023312013000223-2012000332033021-3002013310212023-3121110213303221-2032000201110002)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-3032012133103332-3131112102301330-2300123222223331-0310332200113300-2222301333132021-0310013221233323-2011301210332330-1011121010302102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from start.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010132122000111-0200232110303303-0330111320313312-2330201212200103-0203310320332233-0121310313012122-2322233210313002-2203200230202212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2021112022303101-0201220131321121-0131213002211122-0320300111102212-2320113321321222-2213303122223310-2310233231023111-0300002201113101)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2231302202330301-0332223210023000-1222331123230011-0031232021100133-2233332123101113-0113211010002201-3120103122330300-1233312032013321)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2010102121210030-2020230123133211-2332102200301123-0023312013000223-2012000332033021-3002013310212023-3121110213303221-2032000201110002)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-2313021121112303-1212203122131130-3201223032003001-1130023033011233-1221320032122123-2023202332233333-0222213121320220-1320312010030003"></a>

Type: `"list"`. Computed.

List of networks from which DHCP server can allocate IP addresses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3013002330201001-1201211001211131-1211102111011321-1012322210220011-1203311331202012-1010332221313313-0221003223221330-3301212211331323"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-1323001033132313-3322333131001300-1333130123231220-1000312233012310-3202201331312223-2311321320320201-0111121210330311-3201132321311233"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

Type: `"string"`. Computed.

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="canonical-1132313130321031-1111131023110123-0330121221022021-1203120033111203-0323200102022310-2213301311322331-0012233323132101-0032102323320223"></a>

<a id="canonical-2033123110320230-2131001203220122-2300130011121031-0021033311001023-3221322321302330-2111231110003233-2232102300202200-1103203030021131"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

Type: `"string"`. Computed.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2321303202210133-1123030000031033-2232123232013100-1012321231331123-3233010201323302-3113333212323023-1011122003102112-3132132313113330): complete subsection reference.

<a id="canonical-2321303202210133-1123030000031033-2232123232013100-1012321231331123-3233010201323302-3113333212323023-1011122003102112-3132132313113330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2021112022303101-0201220131321121-0131213002211122-0320300111102212-2320113321321222-2213303122223310-2310233231023111-0300002201113101)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2231302202330301-0332223210023000-1222331123230011-0031232021100133-2233332123101113-0113211010002201-3120103122330300-1233312032013321)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2010102121210030-2020230123133211-2332102200301123-0023312013000223-2012000332033021-3002013310212023-3121110213303221-2032000201110002)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2010132122000111-0200232110303303-0330111320313312-2330201212200103-0203310320332233-0121310313012122-2322233210313002-2203200230202212)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-0012221010102033-2300301203003130-1201223110121212-0210100211121210-0312021112101331-1021122322011132-1310122212021113-3021311313300322"></a>

Type: `"list"`. Computed.

List of non overlapping IP address ranges.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0331131303130200-0312111031233023-0313021202201013-2220310332030311-2332322203200302-1102103011132211-3211200210002002-3210121211000021"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-0132133021202213-0020213003232223-1002210012330121-3120110332022100-2022212311120030-0332103001310001-1133302031123023-1112300011232030"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

Type: `"string"`. Computed.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-1133021331223320-0001220221203332-3100212323321012-2032210233110223-3311002120220313-1321011110101200-0101231020010200-3332211032120001"></a>

<a id="canonical-0012112131302000-1130312022001313-2330221200012000-0201100111332310-3110002211211301-1011013012123112-3310011122130223-2132222030220202"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

Type: `"string"`. Computed.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0121311333000230-1102120020020113-3301202222113203-0130110233310201-1300011132113200-3030302112302203-3213011020103002-1131011023231002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2021112022303101-0201220131321121-0131213002211122-0320300111102212-2320113321321222-2213303122223310-2310233231023111-0300002201113101)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2231302202330301-0332223210023000-1222331123230011-0031232021100133-2233332123101113-0113211010002201-3120103122330300-1233312032013321)
- [eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2010102121210030-2020230123133211-2332102200301123-0023312013000223-2012000332033021-3002013310212023-3121110213303221-2032000201110002)
- eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-3113232303122312-2012313312323331-1122211130313211-2310222310212103-1300103003323103-3323123122222121-0112330221003212-3220323131332123"></a>

Type: `"single"`. Computed.

Map of Interface IPv6 assignments per node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3102313032213222-1330020122002331-2112010230232013-2120110031301233-3203202322030021-3230113230113330-2103230301002012-0320030011123200"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-2121213310133220-3112102331310200-3200033023211301-1230013231331222-1010121323322222-0003030030000013-2220002302302111-0310213112202033"></a>

#### `eks_k8s.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Computed.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.ipv6": "true"
    },
    "values": {
      "format": "ipv6",
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

<a id="canonical-2021330301310310-1220312032020031-1323112300100300-2232101213301320-3021130211331011-3031022021111233-2120012002123211-0031031131022201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- eks_k8s.not_managed.node_list.interface_list.monitor

<a id="canonical-1001320330202101-3120033321200213-2310030332223120-2233002130310211-0121202232302233-0300203111102101-2010332221331012-3231222002131203"></a>

Type: `["object", {}]`. Computed.

Link Quality Monitoring configuration for a network interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111121012022320-1311122221221213-2110002321130211-3132031201011223-1233323212200301-3110331120222010-3020201010102302-1020300302222222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- eks_k8s.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-2232011011110122-3200222101003320-3001020000012331-0031210132331323-2322012223313320-3000222000013103-1330310103031123-1113210220012103"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102310131120122-2032110030130200-2010122221222131-3021333301332000-1331202023210020-2121231331101220-2113020011333202-2113213011003300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- eks_k8s.not_managed.node_list.interface_list.network_option

<a id="canonical-0011303320133010-1210021102102111-0221132020233030-1101123013310121-0012102301233110-2230312003010130-1100333330220200-2310311123133033"></a>

Type: `"single"`. Computed.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

<a id="canonical-0300230100331110-3202010011212131-3122111211131110-2121321002213312-0333213312030123-0121301120002002-0110111000302100-1212120301101100"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-008.md#canonical-3023200220211231-1323031023111332-1120310201222103-0113213210030221-3210031211313021-0001110230210133-1020213022321311-1312020223213100): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0013102222321013-1222221032031113-3002230233003212-1010100320021010-1003032132002322-2302223201130212-2300101220030000-0333330022202132): complete subsection reference.

<a id="canonical-3023200220211231-1323031023111332-1120310201222103-0113213210030221-3210031211313021-0001110230210133-1020213022321311-1312020223213100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2102310131120122-2032110030130200-2010122221222131-3021333301332000-1331202023210020-2121231331101220-2113020011333202-2113213011003300)
- eks_k8s.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-0123110011113132-3032002033211220-1212103220132233-0032200212113120-0230330221330303-2220030123133101-0131323123003220-3103313223202122"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013102222321013-1222221032031113-3002230233003212-1010100320021010-1003032132002322-2302223201130212-2300101220030000-0333330022202132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2102310131120122-2032110030130200-2010122221222131-3021333301332000-1331202023210020-2121231331101220-2113020011333202-2113213011003300)
- eks_k8s.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-3102222022012233-3120322233003320-0223302212211230-0031312212202231-1112330302120232-3320313112332311-0300122033213013-2023020103002321"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102013231002313-0111010212231022-2010103213103201-3001331321231303-2302322113001313-0130313033202112-0322300313220011-1211220110121200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.no_ipv4_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- eks_k8s.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-3313230133113020-0032002122201000-3112102300112312-3031133201113231-0022113303020030-2233101311103231-0212320211011023-2023100221031303"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201230112230311-3202211312030231-0202111031212122-1023110330300221-3031320011121212-2031201103132231-1012022003323123-0211023111232312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- eks_k8s.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-1303201332323013-1103120031130022-2221223200013233-0203003211020133-2322202330200120-3031303322100123-3101112000333122-3030311203013022"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210203003310011-1102320110002020-2332010022121032-0100200312312310-2322031112111212-2203232001332323-1111020320012013-3223200231010220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-0310003210301210-2133331222333022-2201120103313013-1103333121101000-0121020012321131-2003231323112100-3330311012232131-2311022030013013"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311331221011320-2031201013323003-2332112021201201-1022201202320130-0131031130311032-2231003130213013-2312121231203201-0303200211003233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- eks_k8s.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-3100212011212211-0310031100313030-1203210100011223-0322110312133030-2000130220032333-2300111302003222-1301130131301101-1100210320213102"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0020023021001010-2023100110032221-1222000330312221-3321233010211231-1223033231213222-1010132003310320-3001302122210020-1120231310102020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- eks_k8s.not_managed.node_list.interface_list.static_ip

<a id="canonical-3002132100313003-1013112200000331-3023310231111200-2213333133033322-3223023213003211-0300231300312111-2212322220211113-2222132323331220"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1000330300323213-2033332302003110-0321001313223330-1130231001313130-1232100310000333-2033103301310300-3203120301130001-2023000033302100"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.static_ip`

<a id="canonical-2310321101022100-1110013121212312-1100223320330223-3302013313013321-0023210213321213-1133121303223232-1322213210302022-2122230003200010"></a>

#### `eks_k8s.not_managed.node_list.interface_list.static_ip.default_gw` property

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3033223301131010-1233223220132331-1031100023133200-3110200232023301-0232310022332010-1212003322022202-3010011001331322-2223321210032221"></a>

<a id="canonical-0133201230232200-0320012320001001-0030223131110310-2021231020013301-2132233303020210-2122233203220210-1023132102302030-2233101321211312"></a>

#### `eks_k8s.not_managed.node_list.interface_list.static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-2111202332221310-2333301011021200-2003232203202211-0230112303023310-0322103313132123-2002130323223223-3031133100330311-2311312300010123"></a>

<a id="canonical-1123020002320100-0132020311112021-3122300010122121-2332221322211111-3303223132110310-2331010010222120-2321303300303231-3032131123312233"></a>

#### `eks_k8s.not_managed.node_list.interface_list.static_ip.ip_address` property

Type: `"string"`. Computed.

IP address of the interface and prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-1233233223022312-0100303230320213-1130131132010330-2120002001031011-1210302020311313-1320101100110003-0110023111310333-1132012333022011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.static_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- eks_k8s.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-3213113032131023-2323021103232012-1220300232101003-2133111102220131-3322121132133023-0312021030022323-3131302231311330-0111332231202203"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

<a id="canonical-3201013321110201-1230011203213323-2013332121010121-2223130112010210-1231121313223003-0332113130302131-2221321230010103-2223022112310102"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.static_ipv6_address`

- [cluster_static_ip](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0101203021111013-1021330330323323-2012213032312113-3130103321113022-2302301320023131-0123222210210113-0302221221123302-0122232003033020): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2031112033221333-3321020210013113-0312131122321012-3201233321330322-3100221132002021-3301213022112311-0313233223233210-0033020121301033): complete subsection reference.

<a id="canonical-0101203021111013-1021330330323323-2012213032312113-3130103321113022-2302301320023131-0123222210210113-0302221221123302-0122232003033020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1233233223022312-0100303230320213-1130131132010330-2120002001031011-1210302020311313-1320101100110003-0110023111310333-1132012333022011)
- eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-2030120301031333-2332210221103102-2132022133113122-2320222231120222-3301103002012003-0100031031220333-1201113013030001-3330101000302102"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for cluster.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2001330030222003-3220303202110202-1232303301302323-2201112010322011-2132303321333020-0003303313303320-3330322120210003-2111013211211021"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip`

<a id="canonical-3122232311011003-1003130223002010-3302313321302110-0330303221103300-2200233011301303-2130233230202331-3300213312000021-3033122030321020"></a>

#### `eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip.interface_ip_map` property

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "128"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-2031112033221333-3321020210013113-0312131122321012-3201233321330322-3100221132002021-3301213022112311-0313233223233210-0033020121301033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- [eks_k8s.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1233233223022312-0100303230320213-1130131132010330-2120002001031011-1210302020311313-1320101100110003-0110023111310333-1132012333022011)
- eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-1100133023222221-0103210320220322-3021230001121313-2123130200012123-1310222322000122-2200311203310303-2001312312320200-2122230310020331"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0321202131003302-2231113202222120-3122330111300322-0203321302211202-1011331221101112-1122222031220321-2302011030210322-2310303001032020"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip`

<a id="canonical-2101121213222131-3110211223130103-1233000133001011-2121031232310100-3012322300303221-0122013220211033-0131011032323122-2310102331313120"></a>

#### `eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.default_gw` property

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-0201031202221303-3231030320103302-1133010133012030-1313001133132012-3030011203330013-2220103111323222-3233000011011032-3023313112211231"></a>

<a id="canonical-3220020003200310-3031230021103202-3100320200311221-1033210000322030-2023032303203023-1312332112133130-1031223322311303-0303110302211030"></a>

#### `eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-1311030132222031-1013220000011121-2113301323300031-1312312123120322-3210123320323230-2230220003031130-0312112000320312-1101232023013010"></a>

<a id="canonical-3213323210231111-2112223001322012-0022010210312021-1200013012303021-2220131023311332-0112323032012100-3302231310203011-1033313003131300"></a>

#### `eks_k8s.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.ip_address` property

Type: `"string"`. Computed.

IP address of the interface and prefix length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-3120322323322000-0011013231203003-2110133003221231-1112222130322231-2122101000033001-1121333101203131-3210320210333031-0013222122132303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `eks_k8s.not_managed.node_list.interface_list.vlan_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223)
- [eks_k8s.not_managed](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1031123300132132-2103013123300313-2123212312013032-3110301103213122-0112203102312011-2122020203101133-3211333202022203-2310131103023321)
- [eks_k8s.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3002122333313300-3022320132131113-2011211113031112-3213211011220033-2130030033012321-0213311222101222-2310310112101100-2001223203001010)
- [eks_k8s.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1200221011132303-1020103201121101-3003000110031103-1311200121323031-0003321123302220-1032232103101323-0200130222320020-3201132131111113)
- eks_k8s.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-3023301311102333-2333131121002303-0330103121223033-3201212002001103-3221110113323233-0210102111222230-0100202211232230-1111113210012313"></a>

Type: `"single"`. Computed.

Configuration parameter for vlan interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3322213303203222-2000212333000310-3110311123012000-2321300320033000-3211202323220313-0312001100222121-2112122221133203-1122002212021023"></a>

### Direct properties for `eks_k8s.not_managed.node_list.interface_list.vlan_interface`

<a id="canonical-0120212313331313-2102100310231100-0020211012030103-3310010131110122-0023330222011011-0022300111330332-0011210033331120-1211313330012012"></a>

#### `eks_k8s.not_managed.node_list.interface_list.vlan_interface.device` property

Type: `"string"`. Computed.

Select a parent interface from the dropdown.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1223031220202003-3132022133000022-1103221131303103-0133231123323102-2123312031301031-2013101103101020-1210021222220101-0122330131311310"></a>

<a id="canonical-1211022230322023-0303113213121132-3102311331031013-1213013102303321-1331331211120201-1300330033122031-2120323033230331-3120031100211311"></a>

#### `eks_k8s.not_managed.node_list.interface_list.vlan_interface.vlan_id` property

Type: `"number"`. Computed.

Configure the VLAN tag for this interface.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-3110231233201103-1211200233220331-1223220332233133-2230100132132213-1110323332211333-2100302112103010-2232030331320033-1331322222110122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_advanced_delivery` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- enable_advanced_delivery

<a id="canonical-0302100130023223-1112231011201132-1131202132231123-0322331131231201-2233323032102301-2213102200202021-2001233313323223-3130333101201220"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable advanced delivery.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101002333321120-0012212210300003-3121103101312310-3331102111332301-2213213102231013-1030003131222231-1133022113031011-1323132311231201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_ha` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- enable_ha

<a id="canonical-3211233100100031-3103102102120321-1332031312232001-3101332122200020-3002230333113213-1110220232103301-3012032300033212-0101023021030131"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323032112001202-2221032111003112-0322313203323200-3110222233312301-2300310111032030-2010230200012033-1231113132102102-0322030233330201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_log_anonymization` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- enable_log_anonymization

<a id="canonical-3331302221212203-3323202021101332-3001112312031103-2222220230221311-2030330022201232-3111303220131030-2123012011323110-3121000230103012"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable log anonymization.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100210112101123-3322221302032001-0231112220330223-1212133320020300-2330133111020123-3013120322122011-2022211022110312-0032022012331111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_management_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- enable_management_network

<a id="canonical-1011220123303312-3312013212112132-2100333201032033-0001222021323210-3003111023322020-2112122032200301-3201011032332023-1210233120300212"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable management network.

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202010212322121-2023003321100101-1021232022011021-3302013313030110-3132312302310113-0102231023022002-2330103221012233-0203020022030011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_url_categorization` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- enable_url_categorization

<a id="canonical-2233302131312021-3110100032103130-0303012032021213-0123013023101223-0101212113313013-1200303012100011-0220231123233023-0103330010211311"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- equinix

<a id="canonical-2201201002011130-2222100031030122-0201231130130021-0130010130231313-2222323211002102-2113230033003002-0001111233123202-0200101032322102"></a>

Type: `"single"`. Computed.

Equinix Provider Type. Equinix Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-orchestration_choice": "[\"not_managed\"]"
}
```

<a id="canonical-2130203301312220-0100111011110031-2120323111221123-3111303110011033-0223201131103321-0311202302023121-0030302321130200-2110011033110000"></a>

### Direct properties for `equinix`

- [not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131): complete subsection reference.

<a id="canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- equinix.not_managed

<a id="canonical-2222220000000103-3331322220032222-0210121212200333-0012113121031011-0231022212302112-2013302131130231-1201211102310221-2321131010133133"></a>

Type: `"single"`. Computed.

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0011002011031200-1130023310202131-0033103021110222-2312133122223111-2321112211301302-0311023330332323-1311112030301010-0220133013332212"></a>

### Direct properties for `equinix.not_managed`

- [node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022): complete subsection reference.

<a id="canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- equinix.not_managed.node_list

<a id="canonical-0033103103011232-0112332111311223-3023210022203100-0110030232233300-0003002111211323-1022023110303333-2232213222101223-1323030202310032"></a>

Type: `"list"`. Computed.

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2223331333221201-3201013022332010-1021212102001203-1210002323321330-0100300302031322-0211202021131031-0310303203002020-0222300001031110"></a>

### Direct properties for `equinix.not_managed.node_list`

<a id="canonical-0222132113300312-2221230212013133-3122201002222110-0331112331201232-1112200231330000-2311131131110223-3020020233232210-0002133003111130"></a>

#### `equinix.not_managed.node_list.hostname` property

Type: `"string"`. Computed.

Hostname. Hostname for this Node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [interface_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213): complete subsection reference.

<a id="canonical-2200110100323201-1003121030201320-0033003012001313-2012303312023002-0122110123033233-2132212210031220-0312003010310130-1131303021321230"></a>

<a id="canonical-3223122103013211-2131221002303332-0222321021033111-2103131002233023-0011230211201230-1021122100121011-2011222030033030-2112130020100020"></a>

#### `equinix.not_managed.node_list.public_ip` property

Type: `"string"`. Computed.

Public IP. Public IP for this Node.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1230230110101323-3310010100103213-1100110201230103-1223220222223332-0021011011313011-1332103330133011-0333010220131121-0301122200212321"></a>

<a id="canonical-2012321131333113-0031032220011331-0333021113002020-1310213331300323-0002231323131201-2003313032320003-2011321110320321-0030320013210122"></a>

#### `equinix.not_managed.node_list.type` property

Type: `"string"`. Computed.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "Control",
    "Worker"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-1123120300023202-2330131111103112-2103023013200010-1323012110231230-2232023210100220-0322210202333220-2012322012121101-3230101022001213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `equinix.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003)
- [equinix.not_managed](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2111101021121322-1100331111121222-1221030222213022-0222321331301303-2210021230210030-2033120320331300-3313210220312333-3103202003133131)
- [equinix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1323121010323131-0103000132222203-2001232023132000-3223030322031032-2022010001021301-1031003030101100-3310030311103210-0000103030310022)
- equinix.not_managed.node_list.interface_list

<a id="canonical-1033101333011111-3112311121000322-1312110020001211-3232212323111313-3232001321100132-3223132023122032-0232001310220010-0332233211110333"></a>

Type: `"list"`. Computed.

Manage interfaces belonging to this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0131320132022312-2231000110332111-2310131110220111-0201011131032020-0120311033103213-3001120132212200-0102030100212302-3312330311130110"></a>

### Direct properties for `equinix.not_managed.node_list.interface_list`

- [bond_interface](data-sources--securemesh_site_v2--reference--group-009.md#canonical-1313310022002321-1010231013022121-3112323133130003-3302131212211300-0213021031101322-1300201032332332-3320021013112120-3321201031230030): complete subsection reference.

<a id="canonical-2222313011103333-3023110303232212-0311323110232321-2232133322131312-3032023003121200-1313301011330013-0021200100210130-0313110232331101"></a>

<a id="canonical-3313311022102302-0322030023220022-1122001330332001-3001112323001122-3000201011322222-1212300023032121-2010333012012032-3003300221031223"></a>

#### `equinix.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2321103320312032-1311203003133300-3323210313023033-2012201233013011-0120021033012231-0100111232022332-0000020121332102-3211130202133020): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2112122201230121-0100130132312323-0103203013301202-3323023131221223-1103122130313331-1002102210333001-1032120012201111-1313031110203302): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0030022103120212-2330122332133101-3323000322033223-2203313012323233-1122132021130130-2021010032223303-1312302223202331-0031202333110110): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0312230013120131-0301132111133031-3121101031033203-2131100113330221-1011230232332023-2130233021303321-1113332101333300-1103312011222113): complete subsection reference.

<a id="canonical-3302010221120212-3213332132320211-3221003321220221-0033100331220221-2211020011033200-3331112312232303-1020223012012233-1231130022221230"></a>

<a id="canonical-2103311032110302-1013221130201022-3230320122203101-2133020030212023-2203221110230230-2030332220101203-2120211232033311-3100022211332010"></a>

#### `equinix.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-0021001121101213-3101012103133103-2333222002313000-3132200021103220-3020113301210002-1033010221212102-0223003200100213-0310130002221333"></a>

<a id="canonical-2221002000123102-1223311311031311-1032000211301232-2333102011200010-3322001333310303-1301130230012211-3033102101120033-2202123200212121"></a>

#### `equinix.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-1113132120233101-2033102323203030-2222302302011002-1130121210332232-1220021013213231-1123232110310111-1001000323322021-2300211103012101"></a>
