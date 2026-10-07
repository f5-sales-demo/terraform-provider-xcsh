---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-2131211010200313-3110301321130222-1102123200113032-2221200010222312-3311330101103223-1332332132131213-0022022112111000-0312023302203233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ethernet_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-0211222221122032-0202311232031001-0001331010000021-3313132333320000-1302200010033333-1033320023320020-3210101221330230-1000010321323301"></a>

Type: `"single"`. Computed.

Configuration parameter for ethernet interface.

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

<a id="canonical-3110221131232023-3322021231310331-3133110322122111-0031001232030012-3231211212203320-0103020003313122-0120233132131122-3222313000300322"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ethernet_interface`

<a id="canonical-0002310331000132-1110331332230133-0120232100012112-3223020011013200-3301323012310123-2230211012001120-2030221203033113-3232212030302301"></a>

#### `vmware.not_managed.node_list.interface_list.ethernet_interface.device` property

Type: `"string"`. Computed.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2213230032010212-2000201201321230-0220322123232003-2111002231102232-3020211312201201-2333210200102002-1202111302311322-3013221201302201"></a>

<a id="canonical-3020202032100000-0220001103103213-2223312013303122-0103112012230222-0213022232132311-3323310031132110-0202001211321111-1113332020220211"></a>

#### `vmware.not_managed.node_list.interface_list.ethernet_interface.mac` property

Type: `"string"`. Computed.

MAC Address. Configuration parameter for mac

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  }
}
```

<a id="canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-1031111302020110-2120302023321000-3321302133332202-0122033312131020-1221132300102023-1300033001311100-2303102100133203-0332321311010113"></a>

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

<a id="canonical-2231112021321003-3330332012301213-3113031201203033-3022221130120110-3131013231032230-0023000320132230-0311322311311301-3221221310313231"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config`

- [host](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1132013312112011-1022000231113230-2333103132010021-1113013220021002-1313101023330120-1133232113001311-3313013011223212-0122102311222033): complete subsection reference.

- [router](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0013223033023011-3101013310332311-0020120001101222-1022211232333002-2121101202122232-1223120302232000-0000131101220133-0212203332220310): complete subsection reference.

<a id="canonical-1132013312112011-1022000231113230-2333103132010021-1113013220021002-1313101023330120-1133232113001311-3313013011223212-0122102311222033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-3030300230113012-1002302101013133-0331321311202313-2110302031021333-3322332300101103-0131211120302030-3011012020221033-2231213111130032"></a>

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

<a id="canonical-0013223033023011-3101013310332311-0020120001101222-1022211232333002-2121101202122232-1223120302232000-0000131101220133-0212203332220310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-2021321131122311-1210020303300300-2333230210322103-1323123301332331-1013203110112202-0132031002022330-1222131301203232-1010131231021021"></a>

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

<a id="canonical-1010300001013300-3213332303011033-0303223311002233-2030332333003133-3321200330310332-0012100012321220-3223331211123113-3202210213223003"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router`

- [dns_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3011321102030023-1301031300212200-0213233301112122-3200120111211001-3332030331200132-2130101201231023-2002030033332230-0310323032223102): complete subsection reference.

<a id="canonical-0303021332202300-2023122012310230-3123312312330033-2122032123001012-3321021220013202-0330130020320211-1313023113320122-1231220210030320"></a>

<a id="canonical-2300121301100222-1300112003332232-2133320111022112-2103013100320230-0310312302032101-0330101303111011-3231033210110311-3020203000102011"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.network_prefix` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [stateful](data-sources--securemesh_site_v2--reference--group-018.md#canonical-2303133223121221-0333011230032202-3033310233222203-2112323002031022-0301132112233132-0012203012212110-0102112332321031-3312013221111222): complete subsection reference.

<a id="canonical-3011321102030023-1301031300212200-0213233301112122-3200120111211001-3332030331200132-2130101201231023-2002030033332230-0310323032223102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0013223033023011-3101013310332311-0020120001101222-1022211232333002-2121101202122232-1223120302232000-0000131101220133-0212203332220310)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-1300020121210110-1213132132213120-1321123000213202-3131021011300130-0332111101203303-1223232002232001-2303230212020301-2111201130033322"></a>

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

<a id="canonical-0230201032020012-0133203032031300-3032300132032131-1310131121200010-2310320033003122-3122012210110032-3320020111222333-3210332121010023"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config`

- [configured_list](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3020230112011223-2010313001221120-1102301212010011-1002333321000131-2002231210020233-2202331111113210-2003212211312023-2230311102332302): complete subsection reference.

- [local_dns](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3233011301210000-0003000221203101-0210101332311132-0120330301202222-3231120110100332-1232201031231222-2333131210310112-2320132020133333): complete subsection reference.

<a id="canonical-3020230112011223-2010313001221120-1102301212010011-1002333321000131-2002231210020233-2202331111113210-2003212211312023-2230311102332302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0013223033023011-3101013310332311-0020120001101222-1022211232333002-2121101202122232-1223120302232000-0000131101220133-0212203332220310)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3011321102030023-1301031300212200-0213233301112122-3200120111211001-3332030331200132-2130101201231023-2002030033332230-0310323032223102)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-1303001203203221-3000013233123203-1120323332032113-3010321003300320-2110031031222010-0231111312123203-3123000301213221-0313201111022020"></a>

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

<a id="canonical-2230203220312123-2022200102323300-3132133010313101-3032323230332330-2030020230032333-2201223102010132-2130210013200332-2003113102131133"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-3202130221021102-3110323220022120-1001001223113212-0200003020230330-1222111322103210-0202032310121300-3112121202223331-3310233333203222"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3233011301210000-0003000221203101-0210101332311132-0120330301202222-3231120110100332-1232201031231222-2333131210310112-2320132020133333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0013223033023011-3101013310332311-0020120001101222-1022211232333002-2121101202122232-1223120302232000-0000131101220133-0212203332220310)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3011321102030023-1301031300212200-0213233301112122-3200120111211001-3332030331200132-2130101201231023-2002030033332230-0310323032223102)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-0301130203113233-3020023202310333-0003113032233230-3003301302000301-3233202111210022-0320120002101230-2133331123301022-2103003202100100"></a>

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

<a id="canonical-1301211020202301-1322331010011312-2110302001030332-0000003011001101-0113323223121002-0022322302211311-1302222303231302-2201203120131022"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-0031302212323321-0101222210222211-2102313032200003-1012122000203222-3120110333130012-2313110021010213-1001022123133332-1001101022001033"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [first_address](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0030221210310320-2113222321032302-0013233320133132-0230220020011213-3332011303032213-2203100221111223-0101100333130233-1011213231211110): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0021100120322200-2030300021200202-0201302320020111-2333131212130212-2332001132301121-0102221233331223-3010010122133211-1033130030110230): complete subsection reference.

<a id="canonical-0030221210310320-2113222321032302-0013233320133132-0230220020011213-3332011303032213-2203100221111223-0101100333130233-1011213231211110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0013223033023011-3101013310332311-0020120001101222-1022211232333002-2121101202122232-1223120302232000-0000131101220133-0212203332220310)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3011321102030023-1301031300212200-0213233301112122-3200120111211001-3332030331200132-2130101201231023-2002030033332230-0310323032223102)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3233011301210000-0003000221203101-0210101332311132-0120330301202222-3231120110100332-1232201031231222-2333131210310112-2320132020133333)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-3330303332020201-2233230322310012-0312010132102013-1032301233113112-2230112133132230-1031203201231022-0011113113220322-0310110110110210"></a>

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

<a id="canonical-0021100120322200-2030300021200202-0201302320020111-2333131212130212-2332001132301121-0102221233331223-3010010122133211-1033130030110230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0013223033023011-3101013310332311-0020120001101222-1022211232333002-2121101202122232-1223120302232000-0000131101220133-0212203332220310)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3011321102030023-1301031300212200-0213233301112122-3200120111211001-3332030331200132-2130101201231023-2002030033332230-0310323032223102)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3233011301210000-0003000221203101-0210101332311132-0120330301202222-3231120110100332-1232201031231222-2333131210310112-2320132020133333)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-2223230131333023-0001213222203222-1233001132310022-3021311320031311-2221303022300301-1202233302313001-3101102130132033-1221102122212121"></a>

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

<a id="canonical-2303133223121221-0333011230032202-3033310233222203-2112323002031022-0301132112233132-0012203012212110-0102112332321031-3312013221111222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0013223033023011-3101013310332311-0020120001101222-1022211232333002-2121101202122232-1223120302232000-0000131101220133-0212203332220310)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-2011103000300301-3111120212321032-1230220333113121-2000233330312121-0100021321212230-0303101122212130-3031002112200322-1232210200102301"></a>

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

<a id="canonical-2201321312320003-3321033010201333-2020303022023313-2311331033111012-3112212223330003-3331311212320031-1012310110013323-2121102330230333"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3232211121020301-3013301012013232-3020300223100232-2330003331230003-2012331020102220-0320231300213212-1320303121233012-3131331323020313): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1302100202010333-1112301211302233-2032220110231020-1031002230013320-3232330213210132-2232121321321031-3230321323331211-1130032331231330): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0133030021221201-1133220331133012-0102310213312320-3312301031130300-1300321312013101-1311112100322022-3332212212301301-3103121001202301): complete subsection reference.

<a id="canonical-3203103222032213-2012230222221320-1211011312231103-1121100221211331-2033322031303032-0133223003033132-0203100303010013-0203011311030013"></a>

<a id="canonical-1101223321212231-1213311303000011-0321012013033332-3102211101223012-2120022021303103-0313113312121133-0203103230023113-3011102301301120"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

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

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3012222201000120-2012032132221032-1012110310011011-1331210020032103-1020133123211213-2000230212333322-2231302310223331-2220221211131023): complete subsection reference.

<a id="canonical-3232211121020301-3013301012013232-3020300223100232-2330003331230003-2012331020102220-0320231300213212-1320303121233012-3131331323020313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0013223033023011-3101013310332311-0020120001101222-1022211232333002-2121101202122232-1223120302232000-0000131101220133-0212203332220310)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-018.md#canonical-2303133223121221-0333011230032202-3033310233222203-2112323002031022-0301132112233132-0012203012212110-0102112332321031-3312013221111222)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-3213232221033310-3100203101333222-2023130131103233-0312211201133003-0112132020012203-1221111032122003-3133012123021221-0000231332302230"></a>

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

<a id="canonical-1302100202010333-1112301211302233-2032220110231020-1031002230013320-3232330213210132-2232121321321031-3230321323331211-1130032331231330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0013223033023011-3101013310332311-0020120001101222-1022211232333002-2121101202122232-1223120302232000-0000131101220133-0212203332220310)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-018.md#canonical-2303133223121221-0333011230032202-3033310233222203-2112323002031022-0301132112233132-0012203012212110-0102112332321031-3312013221111222)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-0211002312032202-2013213120333220-1031213000321001-1033300310312210-2303202303023332-3221111003132010-0303013222122013-2202300113233121"></a>

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

<a id="canonical-0133030021221201-1133220331133012-0102310213312320-3312301031130300-1300321312013101-1311112100322022-3332212212301301-3103121001202301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0013223033023011-3101013310332311-0020120001101222-1022211232333002-2121101202122232-1223120302232000-0000131101220133-0212203332220310)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-018.md#canonical-2303133223121221-0333011230032202-3033310233222203-2112323002031022-0301132112233132-0012203012212110-0102112332321031-3312013221111222)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-1310130022032312-1022221113213000-2001133213222100-1132111102301130-2300032031231210-0320203131113230-2102132210312330-2211013123201031"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2231222121330133-2300021102200322-1030132112120330-0133223031213111-2312322132313200-2102221133330222-0100011300023120-3230110131122313"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-0332323102123211-2113100111332333-2201101212233231-2031220212331033-3332013313220321-3320222202231102-3003121212212023-2000211232003230"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3023103020110020-3022211303301113-3000200003022312-2012102311102212-3121021330020230-3123302003022333-0013233130323310-2010231200100200"></a>

<a id="canonical-3320212300222202-1333333123021130-0002012320230132-2302011000332301-0103021202122122-1012020233030301-1132000012110230-3300211033322320"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

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

- [pools](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3121122000121132-2013231033312221-3330112332103302-2102030222300201-3301013320322221-3200321300230323-0230130220023330-2002330331130122): complete subsection reference.

<a id="canonical-3121122000121132-2013231033312221-3330112332103302-2102030222300201-3301013320322221-3200321300230323-0230130220023330-2002330331130122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0013223033023011-3101013310332311-0020120001101222-1022211232333002-2121101202122232-1223120302232000-0000131101220133-0212203332220310)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-018.md#canonical-2303133223121221-0333011230032202-3033310233222203-2112323002031022-0301132112233132-0012203012212110-0102112332321031-3312013221111222)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0133030021221201-1133220331133012-0102310213312320-3312301031130300-1300321312013101-1311112100322022-3332212212301301-3103121001202301)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-0103312321313023-2103303102003132-0301122012133230-1111033330311112-3002003030220121-2131111303002230-2200121102113112-2302012123000002"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0120310033131231-3022003111103302-1001323322011101-0111100032332213-1003221213100332-2023233202032000-0011322111023020-3212113302112202"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-0020320103013102-0113313310012002-2012212222131113-0030203321132103-3010031201230020-2012021201121220-2312211111100032-2331313020312223"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0130331010120033-0110002032310311-1112032312132131-1013020223110132-0123311303101300-1033212031113333-0233210113021113-2301332230002003"></a>

<a id="canonical-2202122010313021-2221101131013331-0120133230321223-0022203332313100-0000200023121111-1012131130131112-3231133000022121-1031223320332203"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3012222201000120-2012032132221032-1012110310011011-1331210020032103-1020133123211213-2000230212333322-2231302310223331-2220221211131023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1313010221223230-1033330133120220-3211031211132320-3232022231130313-3212031331222330-2023023002110122-3111300211203321-0232110222213322)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0013223033023011-3101013310332311-0020120001101222-1022211232333002-2121101202122232-1223120302232000-0000131101220133-0212203332220310)
- [vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-018.md#canonical-2303133223121221-0333011230032202-3033310233222203-2112323002031022-0301132112233132-0012203012212110-0102112332321031-3312013221111222)
- vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-2111323230211123-2113333201302233-0332320303312320-0100320131223311-1221311012002303-3110200333320220-2001022333112333-2023011322222013"></a>

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

<a id="canonical-0100301201302002-1111133333321300-1311332303320031-1221001113310203-3122200222233012-1011023003200210-0103200012231100-0222032322331232"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-2021301221212333-1102131232220210-2230230101133111-2010013003101110-2133013300031120-3122101132130231-1020202233000122-0331133031301013"></a>

#### `vmware.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

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

<a id="canonical-3000311103023222-2111102210333112-0020030303122131-2213331012001330-2122023121013123-1302301120133213-1000011210233330-1002330303330200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.monitor

<a id="canonical-1021022112201230-0032101212203020-2223213110222032-1213013003223311-2230021320233220-0103130200320223-3010121201313223-3230131032121022"></a>

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

<a id="canonical-3111220120002011-0130022213201132-0230331301200221-3221122101201330-2230002300332002-2030010032220330-2113133003302200-1121030221220201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-3031312223133202-3331312022031331-3222013313131332-0123011221001211-0003301122023130-2131230130313220-3023333002101321-3031333022120210"></a>

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

<a id="canonical-3320130231323321-3022122100223033-1022132021022321-2011100212203002-0033031330120212-0212210003112230-2202320210023220-3322202133132320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.network_option

<a id="canonical-0000303031323102-1110032111020323-0321212131310130-0322101102110101-0213332030003211-3233130120112130-1121213301131113-3213121013220230"></a>

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

<a id="canonical-3100210033320021-0323233032132121-3111313323010011-0021032001222103-1033000032211201-1012012112130033-2132010000211201-2202030001032201"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](data-sources--securemesh_site_v2--reference--group-018.md#canonical-2031111230210031-0001330212311301-0313200331001201-3220323323233100-3022001022202220-2323100210303331-0001321321111322-3230000132032010): complete subsection reference.

- [site_local_network](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0001123310001110-2330033300321030-0023003300200030-0000322330210112-2211323100010121-2212200232321201-1102011113111003-2012030112332302): complete subsection reference.

<a id="canonical-2031111230210031-0001330212311301-0313200331001201-3220323323233100-3022001022202220-2323100210303331-0001321321111322-3230000132032010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3320130231323321-3022122100223033-1022132021022321-2011100212203002-0033031330120212-0212210003112230-2202320210023220-3322202133132320)
- vmware.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-1123020011333322-3022313221001310-3120210231012323-3033300010001322-1223020012220331-3120113131122030-2311221332330323-0330300010101112"></a>

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

<a id="canonical-0001123310001110-2330033300321030-0023003300200030-0000322330210112-2211323100010121-2212200232321201-1102011113111003-2012030112332302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3320130231323321-3022122100223033-1022132021022321-2011100212203002-0033031330120212-0212210003112230-2202320210023220-3322202133132320)
- vmware.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-0023221331032121-1123303100120112-3122303321323330-3330213122301211-0220212231102112-1130102112300133-0332313300321103-2332031220032212"></a>

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

<a id="canonical-1311301011230121-2103012100303202-2231012122020031-2032112300013302-0112321213131123-3032030000103020-1100121311123230-2033303223003012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.no_ipv4_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-0303131233233122-1100332102100110-2232313033003122-0322133002332012-2001321103303332-2121332320013131-0300003323313130-2022313323210112"></a>

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

<a id="canonical-2221230121110112-0000022310030300-3331023302031310-2301013022330330-3222121321032112-3330122131020130-1322003213222330-3223020113112010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-3123100202323222-0111330210110312-3120033311322321-2212211013010323-2222123001222030-1221002021021333-3011130300202031-3101112311310300"></a>

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

<a id="canonical-2121211122013120-3103203013202130-0310310221301133-3300123122110130-2321201321123323-0200313022223020-0323030000022030-0322311230123100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-2310031122102301-2231020112313221-3202222001320233-1021111211000011-0101222101121302-0200211002312200-2200200120211331-0033123133220033"></a>

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

<a id="canonical-2030023132322101-2011112010030122-0311133322003030-1133300113311203-2103001023133200-0123113110313232-0323002030100222-0231102301022003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-3010313330230233-3231130312003002-0122231132333303-1210330232233321-1210323213300133-2111300120131003-3201331110101013-1133231233213021"></a>

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

<a id="canonical-0322013101212310-2222330010320203-0121011121312102-0210101001022231-1220313221113220-1031320022212303-0320003022130210-3333003000030110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.static_ip

<a id="canonical-1000022311213120-0011030102303210-0101101131322310-0213313302010320-2221332111223202-1333223020221003-0230120111303333-0022323233000200"></a>

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

<a id="canonical-1332133103023001-1330022012230302-2222100213210331-2211102011021330-3023103111323313-2330222322313221-1330103131123221-0232000030213312"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.static_ip`

<a id="canonical-1223103112333101-2202300332113031-3130113111032310-1223112221110113-1002110023103313-3200300230033330-3001132210323211-0313200003111300"></a>

#### `vmware.not_managed.node_list.interface_list.static_ip.default_gw` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1331213332212011-0230311313012313-3111301310210312-2223013022112000-3210133030011103-3020311131001120-1333331311321133-3103113232102020"></a>

<a id="canonical-2112123003101211-0210230211121323-3300321133033220-3233100123131300-2330023220110210-0022320310223302-0102111000102322-3312233323212133"></a>

#### `vmware.not_managed.node_list.interface_list.static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-3011033231213010-3221313323130010-0200022130112210-0201213002212120-1333203311202033-2230133011122133-1112021033020213-1221223100311200"></a>

<a id="canonical-0322230320321130-3002003132022022-0333021033202101-0311203320131201-1011331113221102-0003201003111211-2322002022223320-2332220023330221"></a>

#### `vmware.not_managed.node_list.interface_list.static_ip.ip_address` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0101212013130321-0131302111333332-2312131200130323-3111003310322033-0033230100220132-3322223222230132-0011121000101030-0332100211010113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.static_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-0222123021111323-2222230321200331-2330222103111132-3010233020020000-3003020221013113-2220210222200022-1210323322022123-0023131321300021"></a>

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

<a id="canonical-1133311001311202-3033323223101303-3120111021201120-3120211220022230-3110002031023032-2131210101013300-3111112103331120-1301020301121130"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.static_ipv6_address`

- [cluster_static_ip](data-sources--securemesh_site_v2--reference--group-018.md#canonical-1332113011320100-1231102010023000-3210022222300122-0310330302321033-1101311021213110-1213220130200300-1010021310232031-2332133213302312): complete subsection reference.

- [node_static_ip](data-sources--securemesh_site_v2--reference--group-018.md#canonical-3223001002113301-1302002003100023-1333032130332032-2313132020032320-2332110203022302-0333330001223010-2021021223330330-1133110303003111): complete subsection reference.

<a id="canonical-1332113011320100-1231102010023000-3210022222300122-0310330302321033-1101311021213110-1213220130200300-1010021310232031-2332133213302312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0101212013130321-0131302111333332-2312131200130323-3111003310322033-0033230100220132-3322223222230132-0011121000101030-0332100211010113)
- vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-2111210313223033-3102202210111233-3310331133323311-2203101301011102-1100002122000002-2122201033330232-0012302322021133-3222110100023311"></a>

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

<a id="canonical-2210310311322100-0310330321031311-1131122020112020-2331102022111132-0011332203001231-3030033211132311-3301022023313130-2120322321200331"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip`

<a id="canonical-1221131212202222-3213013301230322-2231331333010222-2133202100221101-0113012030332000-3322012232100300-3323323322212201-3221202313213112"></a>

#### `vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip.interface_ip_map` property

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

<a id="canonical-3223001002113301-1302002003100023-1333032130332032-2313132020032320-2332110203022302-0333330001223010-2021021223330330-1133110303003111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-018.md#canonical-0101212013130321-0131302111333332-2312131200130323-3111003310322033-0033230100220132-3322223222230132-0011121000101030-0332100211010113)
- vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-3311221210303022-1110013201230233-0231330213330203-0331120033120031-0013332323101012-2233331311133230-0202013031310030-0102222301212301"></a>

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

<a id="canonical-3033332203022100-2232303011022212-0231012111200221-0121133111013223-3010210332101222-1023003220020320-1310330101110210-0200323203331011"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip`

<a id="canonical-3012331203302333-3233320010002232-1022112202331330-0320131231000223-0000031013312311-0001101312010003-3313031200121031-2012233210020332"></a>

#### `vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.default_gw` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1231103320112121-2201022110120100-3030021121210102-2211003132033321-1200011000222111-1010112310003212-3133102002003123-0122010231301012"></a>

<a id="canonical-3030200003231312-0100003300010230-1233033210301032-0301103001322310-0213133203130101-0130121133031023-1223312121311203-0222203010020022"></a>

#### `vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.dns_server` property

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-2000331321033000-3230112012303010-2030001012013200-3031332102103130-1202200031230000-2012330110132021-0131200031211102-1111102011032302"></a>

<a id="canonical-1230022130302201-1100211200203133-2000023030203021-3010332010221030-0033003211031303-0101201220310322-2330013332002200-2310213323131323"></a>

#### `vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.ip_address` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0230030312120110-2020013303020100-2023331011232031-0302302103032030-2321322301231232-2223121203332121-2210203321013231-3032330030011313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.vlan_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232)
- [vmware.not_managed](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2032000332103200-2121013331213333-3033310021311131-3111123131231332-2010001211000113-0212203120330033-1222331223210222-3203303131013212)
- [vmware.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0113011332030100-0203133202013110-3000011123010312-0031100220131303-3020100031313211-3031233322310110-2003203200101123-2231232111133101)
- [vmware.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0210213320201122-0010020311033111-2322011110013000-0210112002202130-1202331311120332-2123103333301331-2212231102013223-3333133333130313)
- vmware.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-1122133011300100-3210313110001101-1123100133013133-0133020011311221-3120231213033103-2311000203231120-2203022131100201-2021022322123221"></a>

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

<a id="canonical-3200003312313012-3211122232110032-3112133033121112-2032120130213223-2233212132100310-1202123313212103-3222330130120312-2012311131031131"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.vlan_interface`

<a id="canonical-3013101020323301-0331023130001212-0033021233210221-3203223010120002-2211122010231213-1202013131332303-1213032002012333-1131012001232312"></a>

#### `vmware.not_managed.node_list.interface_list.vlan_interface.device` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0303311133103201-2022332213332012-0023122322032322-1001130110321121-2010320222231210-2213111201132032-1213233100122201-0012330013223132"></a>

<a id="canonical-2120120021130001-1012302230301131-3031323113222121-2023030101222212-0022033303133202-3220001133003310-0113022103331210-1010021202003223"></a>

#### `vmware.not_managed.node_list.interface_list.vlan_interface.vlan_id` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
