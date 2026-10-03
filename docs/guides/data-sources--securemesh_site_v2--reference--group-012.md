---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-2122300302331332-3132121322103132-3333320100211331-3212202033131021-2000121123021120-1032223120303013-2021311301021303-0323330000030221"></a>

## attrs property — static_routes / 311201202010 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1210002331223120-3110122312320321-0111320221321011-0010212232312011-0222321201021303-0030021211033001-1020122312112311-2213132021200121): complete subsection reference.

<a id="canonical-1212213113121112-0132323130032223-2321232113330030-3212300300212220-2103000132200311-0233321231331123-3102330211111232-0312323231322122"></a>

<a id="canonical-3310303222202202-2033210122101322-1012321122231220-2012303213213130-0313022333313220-2231001303220023-0220020013132013-3112230200102031"></a>

## ip_address property — static_routes / 311201202010 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1131131131121101-2310103200012303-1212132133031311-2223211113100021-2203011201302000-0310310133032133-3010023023233311-0003130023312030"></a>

<a id="canonical-1322013200123302-1103111331220033-3221231330301222-1223010011102312-2302011311033101-0332001210323112-1232103312221101-3113331011023132"></a>

## ip_prefixes property — static_routes / 311201202010 / 6

Type: `["list", "string"]`. Computed.

List of route prefixes that have common next hop and attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2133223023121013-2312100122111030-0100020001221101-0312132333210301-0102011203300121-0130231320120131-2201300202012111-2331331121031001): complete subsection reference.

<a id="canonical-0020020013002312-1103032201310313-2102312100323130-1200231121202303-2331120001110313-2331330313021302-1130311213002121-2130313210013032"></a>

## Next pages — static_routes / 311201202010 / 7

- [local_vrf.slo_config.static_routes.static_routes.default_gateway](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1210002331223120-3110122312320321-0111320221321011-0010212232312011-0222321201021303-0030021211033001-1020122312112311-2213132021200121)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2133223023121013-2312100122111030-0100020001221101-0312132333210301-0102011203300121-0130231320120131-2201300202012111-2331331121031001)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1210002331223120-3110122312320321-0111320221321011-0010212232312011-0222321201021303-0030021211033001-1020122312112311-2213132021200121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122131313130131-0300000332003002-1323323333032032-0103012110300311-0201101121332212-0311303122202033-3203210323121302-3112221302022013"></a>

## local_vrf.slo_config.static_routes.static_routes.default_gateway — default_gateway / 023110131013 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310)
- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2010333000110202-0010311002220220-1313231233100130-1103011003332311-1232203220231303-2122333023131133-1010321112103323-0203331130232013)
- local_vrf.slo_config.static_routes.static_routes.default_gateway

<a id="canonical-3220330001301102-0211021232103133-1220312322120331-1132120200002130-3310210310221021-1221131131120012-3032103110011130-3023012230320200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

Upstream description:

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

<a id="canonical-1131312312132012-3120332232202313-3301221220030201-0033303321033202-3332120120130130-3201311131332333-3211012320332232-3331022200003010"></a>

## Direct properties — default_gateway / 023110131013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121200302112033-0230022012130201-3121313320333003-1123203301112122-0212302022112031-1130022022010123-3310201002230211-3203312000323013"></a>

## Next pages — default_gateway / 023110131013 / 4

- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2010333000110202-0010311002220220-1313231233100130-1103011003332311-1232203220231303-2122333023131133-1010321112103323-0203331130232013)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2133223023121013-2312100122111030-0100020001221101-0312132333210301-0102011203300121-0130231320120131-2201300202012111-2331331121031001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000222130110313-0230210113233123-1131020331002331-3003012103003310-2302221111320113-3002100022020030-0201122212122020-2123033221232002"></a>

## local_vrf.slo_config.static_routes.static_routes.node_interface — node_interface / 320000310021 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310)
- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2010333000110202-0010311002220220-1313231233100130-1103011003332311-1232203220231303-2122333023131133-1010321112103323-0203331130232013)
- local_vrf.slo_config.static_routes.static_routes.node_interface

<a id="canonical-2110101322321323-2120022210221122-1023312033200032-1022112331200222-1132103121322230-0010032302212112-0302120011020300-1120313303300022"></a>

Type: `"single"`. Computed.

On multinode site, this type holds the information about per node interfaces.

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

<a id="canonical-0300010323132002-3311221022112130-1032032311102231-2332223230013113-0013023001323121-1231332203003131-1032210220213300-0210020031113112"></a>

## Direct properties — node_interface / 320000310021 / 3

- [list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1012203320323132-2211132101311103-1333101322231003-1112011202013201-1213123223100001-3230330223101223-1021112033002133-0200203132312321): complete subsection reference.

<a id="canonical-2000010321111022-3201012220131323-2022221110112130-3233230200303130-0102311230112032-1120203030222030-3332333032132110-3323112110033021"></a>

## Next pages — node_interface / 320000310021 / 4

- [local_vrf.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1012203320323132-2211132101311103-1333101322231003-1112011202013201-1213123223100001-3230330223101223-1021112033002133-0200203132312321)
- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2010333000110202-0010311002220220-1313231233100130-1103011003332311-1232203220231303-2122333023131133-1010321112103323-0203331130232013)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1012203320323132-2211132101311103-1333101322231003-1112011202013201-1213123223100001-3230330223101223-1021112033002133-0200203132312321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002123001022201-3222221213311202-3021200213020131-1213120003101221-1123320322330113-0301323311311123-0233330021112021-1301102023133331"></a>

## local_vrf.slo_config.static_routes.static_routes.node_interface.list — list / 010113101103 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310)
- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2010333000110202-0010311002220220-1313231233100130-1103011003332311-1232203220231303-2122333023131133-1010321112103323-0203331130232013)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2133223023121013-2312100122111030-0100020001221101-0312132333210301-0102011203300121-0130231320120131-2201300202012111-2331331121031001)
- local_vrf.slo_config.static_routes.static_routes.node_interface.list

<a id="canonical-0131011133311122-3303033023132210-3312103233201200-1211120210011323-2300123130213031-2123001333300012-3302103012121122-1201123031123031"></a>

Type: `"list"`. Computed.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-2020120210200320-3021333210211202-3210212101112201-0033022312203030-0320012013100320-2333100233332010-0021123310131310-1330121022123312"></a>

## Direct properties — list / 010113101103 / 3

- [interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3222110233202020-1302103100230010-1333230331311102-2302021032112130-0101113212111130-1232010111110101-1320033022010320-3300022031111302): complete subsection reference.

<a id="canonical-1132030223300113-2013203113301002-0331033300112132-3021032201203200-0013302231023323-1231233220202130-3133121203232001-3301013232221322"></a>

<a id="canonical-1130221103203223-1013211013233221-1230110200020333-2000120033303113-2100332020312110-3132210322233010-0100021331312232-1130231203212323"></a>

## node property — list / 010113101103 / 4

Type: `"string"`. Computed.

Node. Node name on this site.

Upstream description:

Node name on this site.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3230310023220331-3323101211100100-1200003000310112-3012230321020121-1312213113110213-1132331121100120-2120102221110013-0032333100133120"></a>

## Next pages — list / 010113101103 / 5

- [local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3222110233202020-1302103100230010-1333230331311102-2302021032112130-0101113212111130-1232010111110101-1320033022010320-3300022031111302)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2133223023121013-2312100122111030-0100020001221101-0312132333210301-0102011203300121-0130231320120131-2201300202012111-2331331121031001)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3222110233202020-1302103100230010-1333230331311102-2302021032112130-0101113212111130-1232010111110101-1320033022010320-3300022031111302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223012230000103-3121101011201320-3300033313130110-0201131100000101-1101132021332233-3203123122302232-0033213032001231-3231231102320003"></a>

## local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface — interface / 102210100120 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-1113203021220022-1100133310332033-3320222120311032-1310020013320311-1201220321230122-0223100332131101-0220302111100302-1132301212120310)
- [local_vrf.slo_config.static_routes.static_routes](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2010333000110202-0010311002220220-1313231233100130-1103011003332311-1232203220231303-2122333023131133-1010321112103323-0203331130232013)
- [local_vrf.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2133223023121013-2312100122111030-0100020001221101-0312132333210301-0102011203300121-0130231320120131-2201300202012111-2331331121031001)
- [local_vrf.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1012203320323132-2211132101311103-1333101322231003-1112011202013201-1213123223100001-3230330223101223-1021112033002133-0200203132312321)
- local_vrf.slo_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-3102200011002202-0023300013011023-0320001131323010-2222022321331213-3213021322013032-3332112221222122-0011032300321001-1211122332032322"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0131012003311301-1313331300133332-3010211222133231-3202012203230233-3223230102301020-0213112310311112-2110031211320020-1212001133201231"></a>

## Direct properties — interface / 102210100120 / 3

<a id="canonical-3001213332310221-3003130233233210-2211033321302213-0112322130103003-1121202013212310-1333223120301021-1201210003013012-0103320001200022"></a>

<a id="canonical-0123323220221312-2132023122201003-2332213113320112-0022130212203013-3101303120330111-1321310023310323-0100231201103001-0121320311332232"></a>

## kind property — interface / 102210100120 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2022302232111132-0203120230303202-0230030123332322-3312312211132031-2021323101310020-1122323203012113-1103010103110321-0023330203130200"></a>

<a id="canonical-2132330111013332-3030113311333003-0232211001002122-2313222203223221-0112203102113013-2232220110312013-0012123331100001-1132212001102330"></a>

## name property — interface / 102210100120 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3302323311302113-1230213130112132-2102313030313111-0312123320333023-0201013330232333-0302003223210221-1012200132220132-1332101223210212"></a>

<a id="canonical-0033002323022111-3123332123332220-2322210212130313-1102030331210221-1300223303130020-0102020230031223-2133233121122320-3201310011001303"></a>

## namespace property — interface / 102210100120 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2021233303231212-3100332321212311-1200032321230300-3103010023311001-2111122013312110-1111130003121031-3202313202333021-3032203330300100"></a>

<a id="canonical-0333013230003331-0020010301310330-3302333323210031-3302123000301012-2322203001012110-2201000322103102-1333132302132123-2100122320011122"></a>

## tenant property — interface / 102210100120 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1230323310020303-3313001210120131-0303203330220123-0230002000000021-1012321203030022-1101103303120022-0022201222223101-3201000330212120"></a>

<a id="canonical-0200003003303002-1120103012121001-1221101333122311-0003122031232231-2330303102023232-0122012113022322-2333011233213323-0223212330230303"></a>

## uid property — interface / 102210100120 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2102033233222113-0312202011201022-2023021322103020-0203113330303301-2022202232320120-1003011231332130-1303020132003122-3323130001131032"></a>

## Next pages — interface / 102210100120 / 9

- [local_vrf.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1012203320323132-2211132101311103-1333101322231003-1112011202013201-1213123223100001-3230330223101223-1021112033002133-0200203132312321)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3032321313011322-3033023321030220-0231322301021002-3000322033003002-1300012101233333-3301212000221120-2023201330030121-2033020221220222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112323112210032-2222331332332011-3203111303010030-0112102222023011-0100232133130033-3201000003001321-2232230230312232-2230333031003323"></a>

## local_vrf.slo_config.static_v6_routes — static_v6_routes / 101320130212 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- local_vrf.slo_config.static_v6_routes

<a id="canonical-3021000113330001-3321113322001211-1011231322212022-3113121001321003-0123212233030033-3122011301021111-1020001210322133-3210101231221100"></a>

Type: `"single"`. Computed.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

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

<a id="canonical-2010313100002220-2031210103220231-3302320223302320-2002213203131321-3012022333131210-2032322330002011-1131312012101013-3210301111013303"></a>

## Direct properties — static_v6_routes / 101320130212 / 3

- [static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233102011103321-3232301121201122-3000001331112310-3010221201122203-1331333210030133-3122002031013320-3222330313101233-3102211233030212): complete subsection reference.

<a id="canonical-1231001033021231-2022100211023012-2123201302323131-1120203230133010-0233013021023131-2323331121310023-1102233221001311-0322201022310202"></a>

## Next pages — static_v6_routes / 101320130212 / 4

- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233102011103321-3232301121201122-3000001331112310-3010221201122203-1331333210030133-3122002031013320-3222330313101233-3102211233030212)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2233102011103321-3232301121201122-3000001331112310-3010221201122203-1331333210030133-3122002031013320-3222330313101233-3102211233030212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100203210113320-0210013300032101-0011022003103112-0032021112033200-2333122000333133-1110333101310032-2231012011222133-0310211233021220"></a>

## local_vrf.slo_config.static_v6_routes.static_routes — static_routes / 331200030302 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3032321313011322-3033023321030220-0231322301021002-3000322033003002-1300012101233333-3301212000221120-2023201330030121-2033020221220222)
- local_vrf.slo_config.static_v6_routes.static_routes

<a id="canonical-3102201303300212-3120103322112201-1220112002102133-0330120333020022-3213300110031010-2200132301012001-1112022013032032-3121310322022331"></a>

Type: `"list"`. Computed.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0310231000121131-1003013221101303-0023202210021003-0002001331213103-3131013113331300-1021033311302113-3013111231310211-3101303100020113"></a>

## Direct properties — static_routes / 331200030302 / 3

<a id="canonical-0323031301222001-3333010301212322-1322012210102103-3131301132212331-3103020030321223-2321212213210030-0230000110331203-2320002330201013"></a>

<a id="canonical-3211222112222021-1131203003132322-1022230200300013-0003023112332013-3201022010001120-0120031033130312-0323203223033312-0310202031223212"></a>

## attrs property — static_routes / 331200030302 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3301233302320310-3121032100211223-2201033200120110-3133303030123212-1302320331022131-0022233320230203-1112001202211222-1111312210323023): complete subsection reference.

<a id="canonical-3113333332233000-2303202102100130-1320021003201321-0220102112303110-1020321032023011-0123320330103203-2123222310222210-0002112203122323"></a>

<a id="canonical-3222331122100031-2201132023111321-2231300033022300-3320110110301313-2110011000132101-1133031001030320-3030130001303021-0212002220320210"></a>

## ip_address property — static_routes / 331200030302 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-2122312133132133-3102013210022201-3200120210103013-0120031310102202-3323033011003020-3001211103032232-1001113131033331-0102212301011303"></a>

<a id="canonical-0212301123133001-3020231212331021-0101320123022202-0301320230101133-0133013322030022-3121000112033301-2002110113100113-1331030121010023"></a>

## ip_prefixes property — static_routes / 331200030302 / 6

Type: `["list", "string"]`. Computed.

List of IPv6 route prefixes that have common next hop and attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2331002132220320-0033121200033312-3201332323320303-1301332101030303-1303201300120213-1021111220303331-1232311203330002-3232331233300102): complete subsection reference.

<a id="canonical-0203333013200311-3110121313012332-2001213332033231-1202213311010311-2122203111232221-0102202330200311-1230222311001113-3030310013121010"></a>

## Next pages — static_routes / 331200030302 / 7

- [local_vrf.slo_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3301233302320310-3121032100211223-2201033200120110-3133303030123212-1302320331022131-0022233320230203-1112001202211222-1111312210323023)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2331002132220320-0033121200033312-3201332323320303-1301332101030303-1303201300120213-1021111220303331-1232311203330002-3232331233300102)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3032321313011322-3033023321030220-0231322301021002-3000322033003002-1300012101233333-3301212000221120-2023201330030121-2033020221220222)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3301233302320310-3121032100211223-2201033200120110-3133303030123212-1302320331022131-0022233320230203-1112001202211222-1111312210323023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033231021003003-3233113112001001-0223230320210013-3222033202113033-0222111320233321-3210001030212001-3103011321123021-0213320022322131"></a>

## local_vrf.slo_config.static_v6_routes.static_routes.default_gateway — default_gateway / 002130132011 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3032321313011322-3033023321030220-0231322301021002-3000322033003002-1300012101233333-3301212000221120-2023201330030121-2033020221220222)
- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233102011103321-3232301121201122-3000001331112310-3010221201122203-1331333210030133-3122002031013320-3222330313101233-3102211233030212)
- local_vrf.slo_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-2001222310330030-1002020310301230-3121311010133001-2210211322303300-0301011113100132-0321212330231222-3230110012323020-2213323211231103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

Upstream description:

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

<a id="canonical-0112033211310100-1300021111002231-3230333031020113-2313223300221001-1112113023032223-0020022210210102-2011131203102000-0330321021312112"></a>

## Direct properties — default_gateway / 002130132011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302302333211120-3313003033220223-0023233031121233-0320011213121010-2203032301013211-0031333333123231-1221203211300020-3313233223103201"></a>

## Next pages — default_gateway / 002130132011 / 4

- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233102011103321-3232301121201122-3000001331112310-3010221201122203-1331333210030133-3122002031013320-3222330313101233-3102211233030212)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2331002132220320-0033121200033312-3201332323320303-1301332101030303-1303201300120213-1021111220303331-1232311203330002-3232331233300102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233120203000033-0322023220203013-3003132332321233-3203312113023112-0322331232330312-3131012122232232-0103123032013013-1020223002101321"></a>

## local_vrf.slo_config.static_v6_routes.static_routes.node_interface — node_interface / 013322001012 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3032321313011322-3033023321030220-0231322301021002-3000322033003002-1300012101233333-3301212000221120-2023201330030121-2033020221220222)
- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233102011103321-3232301121201122-3000001331112310-3010221201122203-1331333210030133-3122002031013320-3222330313101233-3102211233030212)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface

<a id="canonical-1020022301232312-1331133130331211-0101110100032310-0100212303210122-2010113012022133-3222130203010231-2220331102113131-1012330202200131"></a>

Type: `"single"`. Computed.

On multinode site, this type holds the information about per node interfaces.

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

<a id="canonical-3120323111313123-0112001313321121-1202100203331233-1021102301023101-2311202333233230-3130213031212213-2023121321001101-3010221323223323"></a>

## Direct properties — node_interface / 013322001012 / 3

- [list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0023332122103020-2100030312322233-1011120300323023-2330322200033320-2121120232103013-3022033102332233-0030013012030130-3222322101032231): complete subsection reference.

<a id="canonical-0333011220112103-3303113102322111-3201202122132013-1222333313122132-0002331320031210-2322221032323210-1300022112311012-0333300033102101"></a>

## Next pages — node_interface / 013322001012 / 4

- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0023332122103020-2100030312322233-1011120300323023-2330322200033320-2121120232103013-3022033102332233-0030013012030130-3222322101032231)
- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233102011103321-3232301121201122-3000001331112310-3010221201122203-1331333210030133-3122002031013320-3222330313101233-3102211233030212)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0023332122103020-2100030312322233-1011120300323023-2330322200033320-2121120232103013-3022033102332233-0030013012030130-3222322101032231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302320002001031-2213302201201102-3132100111111021-2300300321123211-0013103123221102-2130130130033000-2002113313310032-0110030302111213"></a>

## local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list — list / 212232321020 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3032321313011322-3033023321030220-0231322301021002-3000322033003002-1300012101233333-3301212000221120-2023201330030121-2033020221220222)
- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233102011103321-3232301121201122-3000001331112310-3010221201122203-1331333210030133-3122002031013320-3222330313101233-3102211233030212)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2331002132220320-0033121200033312-3201332323320303-1301332101030303-1303201300120213-1021111220303331-1232311203330002-3232331233300102)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-2231002122122320-0202332302121221-1301123013101130-1111103130110021-0201130132030331-0230103120212120-1221110012021020-2022222211031321"></a>

Type: `"list"`. Computed.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-0101303220202212-2030131120011131-0103312132210022-0321002031333130-1322021120320121-2021221120312010-0021201022030323-2012303331032022"></a>

## Direct properties — list / 212232321020 / 3

- [interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2130210023031001-3320220020021231-0221232120113120-0322133310310110-1332023110311001-3212033233102033-2310202323201210-1000210003232311): complete subsection reference.

<a id="canonical-3130002332320222-2210003103212003-0213210231233321-2221323001323233-0333101203221233-2330113301321133-3320200100303210-1301110120331211"></a>

<a id="canonical-2003103211031300-1210112032012101-1110303220210300-1211211221313313-2003331312003313-0201100311202332-2111101123213101-2013023012300010"></a>

## node property — list / 212232321020 / 4

Type: `"string"`. Computed.

Node. Node name on this site.

Upstream description:

Node name on this site.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3223211102101110-1232301233020222-3030100102021330-2123110002033231-3112223022002023-2200222301323130-0132213022130233-1031202201021011"></a>

## Next pages — list / 212232321020 / 5

- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2130210023031001-3320220020021231-0221232120113120-0322133310310110-1332023110311001-3212033233102033-2310202323201210-1000210003232311)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2331002132220320-0033121200033312-3201332323320303-1301332101030303-1303201300120213-1021111220303331-1232311203330002-3232331233300102)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2130210023031001-3320220020021231-0221232120113120-0322133310310110-1332023110311001-3212033233102033-2310202323201210-1000210003232311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120101300123211-1332003111123003-1120311010031331-2132302221031303-3333022033223133-3120013021102111-2020312323200301-2320102333002112"></a>

## local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface — interface / 001302130010 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211)
- [local_vrf.slo_config](data-sources--securemesh_site_v2--reference--group-011.md#canonical-2313111001121323-2310312203311330-0132031013032201-3112313310211203-0103301302031131-3220120001122303-2010230332331122-2222303111201300)
- [local_vrf.slo_config.static_v6_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3032321313011322-3033023321030220-0231322301021002-3000322033003002-1300012101233333-3301212000221120-2023201330030121-2033020221220222)
- [local_vrf.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233102011103321-3232301121201122-3000001331112310-3010221201122203-1331333210030133-3122002031013320-3222330313101233-3102211233030212)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2331002132220320-0033121200033312-3201332323320303-1301332101030303-1303201300120213-1021111220303331-1232311203330002-3232331233300102)
- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0023332122103020-2100030312322233-1011120300323023-2330322200033320-2121120232103013-3022033102332233-0030013012030130-3222322101032231)
- local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-0333000012102302-1012302203223013-2202230332132231-0221300022112030-3322311202221210-1012000210103210-2112310321313231-3103000102031013"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2130211032111221-1220222311332211-2132203313223133-2130001133311213-0002113101130230-1303231331320022-0323011112201221-1231113310203100"></a>

## Direct properties — interface / 001302130010 / 3

<a id="canonical-0221330031230002-0023220303213203-2202300023332232-1233112133232223-1312132033332222-3021301331223321-0321131101313331-1203033100000212"></a>

<a id="canonical-0102121033300122-1011213200020321-1101020331002121-3301033032101031-0231011310122102-2231000323303313-1222202301030120-0003013213233301"></a>

## kind property — interface / 001302130010 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1032120000331223-1012020322220111-0010101302031012-0000110230201102-3003203200102220-2022102301010313-0210123012100223-3122111303231232"></a>

<a id="canonical-0301113212202310-3130233002102032-3312132303313023-2003220321212330-3200303222023033-2310031310332111-2023103031230002-3203101003033313"></a>

## name property — interface / 001302130010 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2220023030220013-0220001200221232-2321323221131103-1303321000230011-0312201201200220-1213333103032011-3023031000233330-2133303221110000"></a>

<a id="canonical-2301333310303312-2012203021100120-0210130121210002-3012013000321011-1313312313313311-1001032303010122-1210100310211110-3200201100121101"></a>

## namespace property — interface / 001302130010 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2220211021330123-2012113002021322-1331303203132210-1133011201113222-1023223320311033-0300331000303100-3032113303002110-2222223200031113"></a>

<a id="canonical-3131231330133230-3332313023003332-3031203211220213-2210321003113121-1113032322331112-3000301303213123-0312131203133233-0133302032202003"></a>

## tenant property — interface / 001302130010 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3122311222302011-1321030011033013-1233003102311303-2100221113201232-2012212033100131-2000012021211223-1200113000203010-3102210030331320"></a>

<a id="canonical-1212111022101301-3013032310201031-0233122321332003-0231122022100222-3110133113311020-3303302310310123-1030030031233120-0210223322021022"></a>

## uid property — interface / 001302130010 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3231003023230013-3000112313100230-2312030032210231-3310120013130103-0030223301201322-1332213212232120-1131210010003123-1323121120210221"></a>

## Next pages — interface / 001302130010 / 9

- [local_vrf.slo_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0023332122103020-2100030312322233-1011120300323023-2330322200033320-2121120232103013-3022033102332233-0030013012030130-3222322101032231)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3023110201200001-2313230223011122-3303112211100000-3132120202123300-1210303000011110-0232103302130300-0331113010221003-0103313111112332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032201332133312-1022021330032222-0132300011013010-1021011013102130-2103322302031031-0200103101102320-1122310231213302-1033011023101132"></a>

## log_receiver_with_net — log_receiver_with_net / 303301112330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- log_receiver_with_net

<a id="canonical-0301211103003021-0313300321312033-2120323001233212-2030032032023221-0303303100300233-2222013302331212-1311132303231023-1123201110102230"></a>

Type: `"single"`. Computed.

\[OneOf: log\_receiver\_with\_net, logs\_streaming\_disabled\] Select log receiver for logs
streaming with network option.

Upstream description:

Select log receiver for logs streaming with network option.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"use_management_network\",\"use_slo_sli\"]"
}
```

OneOf alternatives in this subsection:

- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0301211103003021-0313300321312033-2120323001233212-2030032032023221-0303303100300233-2222013302331212-1311132303231023-1123201110102230)
- [logs_streaming_disabled](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0113001212320132-1202200001010210-0110120331313212-3133300233102210-3213232231120120-2022103212230103-2202212230013003-0132130233031300)

Select alternatives according to the provider validators above.

<a id="canonical-2311130321231101-0133221102301121-1000303111113022-2001320321013311-0313313100030312-2120230031213022-2030312222200323-3003113213102332"></a>

## Direct properties — log_receiver_with_net / 303301112330 / 3

- [log_receiver](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0211033200202122-1123101121020121-3330110012222130-0013311202020033-0110213122332301-2301011133233130-0220202020200001-2120111032000123): complete subsection reference.

- [use_management_network](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2101311113003120-2012100221130321-1033020122303322-3303312323312102-1022312221231230-2331301302230323-2100211112320301-2322112033300303): complete subsection reference.

- [use_slo_sli](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3020330001303210-1313031003103231-2011211201321223-3320330302021110-1100112022000203-2320231020030103-2023112321021011-1022131212022232): complete subsection reference.

<a id="canonical-2132122012033210-3022001213312022-3001100332112000-3020121121310103-1010020211030110-2201110132300231-0333110331233312-1221220131000120"></a>

## Next pages — log_receiver_with_net / 303301112330 / 4

- [log_receiver_with_net.log_receiver](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0211033200202122-1123101121020121-3330110012222130-0013311202020033-0110213122332301-2301011133233130-0220202020200001-2120111032000123)
- [log_receiver_with_net.use_management_network](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2101311113003120-2012100221130321-1033020122303322-3303312323312102-1022312221231230-2331301302230323-2100211112320301-2322112033300303)
- [log_receiver_with_net.use_slo_sli](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3020330001303210-1313031003103231-2011211201321223-3320330302021110-1100112022000203-2320231020030103-2023112321021011-1022131212022232)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0211033200202122-1123101121020121-3330110012222130-0013311202020033-0110213122332301-2301011133233130-0220202020200001-2120111032000123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003103302011233-1223221121132123-2032130301233111-3030120020320100-1200331203212322-2020220222112133-3210021110132203-1011120310112300"></a>

## log_receiver_with_net.log_receiver — log_receiver / 000320301012 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3023110201200001-2313230223011122-3303112211100000-3132120202123300-1210303000011110-0232103302130300-0331113010221003-0103313111112332)
- log_receiver_with_net.log_receiver

<a id="canonical-1211223232303330-1012111302230330-0323030310202212-3000130203020131-3312303223322032-1133300011203313-0232232013011211-0302230313320332"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2101223233023200-3200021233202122-3000222223210231-1032111211123211-3322312100112012-3332311222122230-2313101100102112-2031101300221331"></a>

## Direct properties — log_receiver / 000320301012 / 3

<a id="canonical-2031202203012301-1223210111303231-2230131032301211-1130212032210313-2011300233001300-1321213103011232-3222012300022312-0313100021202202"></a>

<a id="canonical-0202221131002022-2300010211311110-0011233300323210-2010011132213110-2122323002210233-2303030020033133-2301320223201111-1230131021020132"></a>

## name property — log_receiver / 000320301012 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1332210210032121-3013211223113130-3132310012122113-3311110013123312-0133123210132202-0313212203122322-1021102021333132-3323202120110020"></a>

<a id="canonical-1021103101100331-1223211303110232-1131011202331021-2202111303122303-0332223212310203-2010000312231312-3302311210110122-1231323313021221"></a>

## namespace property — log_receiver / 000320301012 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2323310332031313-1031130010122301-1133113112322201-0322233311020221-0212033030003122-1113211120002032-1120112102112130-0333322102103002"></a>

<a id="canonical-0232033002133200-2202021221213103-2231303200330032-0313333331110132-1022023102111013-2232222202210213-3323220012331101-0002223133102110"></a>

## tenant property — log_receiver / 000320301012 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3111210013203000-1232123102301022-0312103010012301-0101313310021122-3303002101132130-2203012302322030-1312131130221320-3321110023121203"></a>

## Next pages — log_receiver / 000320301012 / 7

- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3023110201200001-2313230223011122-3303112211100000-3132120202123300-1210303000011110-0232103302130300-0331113010221003-0103313111112332)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2101311113003120-2012100221130321-1033020122303322-3303312323312102-1022312221231230-2331301302230323-2100211112320301-2322112033300303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102132210331303-0023313211031220-1121300132203000-3212223223222223-1321112101321310-1003133303313121-1021323321220310-3023211303310123"></a>

## log_receiver_with_net.use_management_network — use_management_network / 302211030201 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3023110201200001-2313230223011122-3303112211100000-3132120202123300-1210303000011110-0232103302130300-0331113010221003-0103313111112332)
- log_receiver_with_net.use_management_network

<a id="canonical-2232233012321222-2213200222012010-3010230110222303-0213031230130033-2301212122113110-0322302333012233-1331103132233221-0332121100113003"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use management network.

Upstream description:

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

<a id="canonical-0133111130331330-2201330222313002-3013033120312131-0211302000010323-1330332110120211-0132231100101000-0122302130302103-1221022323102010"></a>

## Direct properties — use_management_network / 302211030201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311001101031303-0011223111032013-0322322013301321-1313201103332131-1002101202121313-1033013301312113-1032103210031333-2011120221313220"></a>

## Next pages — use_management_network / 302211030201 / 4

- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3023110201200001-2313230223011122-3303112211100000-3132120202123300-1210303000011110-0232103302130300-0331113010221003-0103313111112332)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3020330001303210-1313031003103231-2011211201321223-3320330302021110-1100112022000203-2320231020030103-2023112321021011-1022131212022232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222131322120333-1321011001010320-3002202133211032-1322300302233220-3223303222010102-0013321311330223-0311013223330302-1120030112131333"></a>

## log_receiver_with_net.use_slo_sli — use_slo_sli / 233001113020 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3023110201200001-2313230223011122-3303112211100000-3132120202123300-1210303000011110-0232103302130300-0331113010221003-0103313111112332)
- log_receiver_with_net.use_slo_sli

<a id="canonical-2202032112213110-3032320112313233-2300120113123213-0200010333011013-0103100002012021-2102033230230220-2100333311331331-1123113301010100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use slo sli.

Upstream description:

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

<a id="canonical-1112211120021011-3001311113310023-3221133133320002-0233322233112211-3230331020230231-3200013302330021-0120033231032031-2221100320212101"></a>

## Direct properties — use_slo_sli / 233001113020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133010002103010-0120322023232003-3333220011102301-0020101000310321-1333302031223112-0332230323232012-1223032202100012-1131210123331003"></a>

## Next pages — use_slo_sli / 233001113020 / 4

- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3023110201200001-2313230223011122-3303112211100000-3132120202123300-1210303000011110-0232103302130300-0331113010221003-0103313111112332)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3312333002021230-2313121320113132-0202121101132333-1020313233323133-1212020320020130-0320131000002003-3213221321030320-2011010311010100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110030103311201-0200211201111031-3033211331122111-0211213103223300-2220011121311230-3301020133212022-0231200013111202-0100110001303210"></a>

## logs_streaming_disabled — logs_streaming_disabled / 230312320221 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- logs_streaming_disabled

<a id="canonical-0113001212320132-1202200001010210-0110120331313212-3133300233102210-3213232231120120-2022103212230103-2202212230013003-0132130233031300"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-1212123022102012-3312213211201101-3010021200001000-0300213122130202-0231133313212020-2232030030132310-2322103201322221-3003123310300010"></a>

## Direct properties — logs_streaming_disabled / 230312320221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210330310230322-1330322112311210-1110030203220323-3132023032131012-3331321022120123-0020000313202302-2212121323113121-2030023112321320"></a>

## Next pages — logs_streaming_disabled / 230312320221 / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0233333211333233-2312020110021301-2123021233033012-3010102332313210-3231313033300120-1121330230203121-1210102010020212-3302330001121210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230003032012120-2031331332230203-1121002321311110-2313122122101013-2100301230023133-0001113220312222-1101023131202000-3100002130301313"></a>

## no_forward_proxy — no_forward_proxy / 033303012013 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- no_forward_proxy

<a id="canonical-3131331111200213-1222033103123101-0313132300212121-0002301220000031-2200101113220100-2131220000133320-0323310300120230-2032200330321331"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no forward proxy.

Upstream description:

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

<a id="canonical-0113103110330303-0113121300330200-1230110302031002-0120012303122310-1002211132211210-2100230003211223-3123202222221232-0302333110121201"></a>

## Direct properties — no_forward_proxy / 033303012013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011303020301003-0000312121232012-3322302100230220-2200232121320211-2213013123001133-2210203323322112-2312210000213112-3110033122010111"></a>

## Next pages — no_forward_proxy / 033303012013 / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2300302332233003-0012321023230001-0002203123133321-3301231303333030-2131311231231300-3321021131122301-1111003011023013-3302002202323103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202301333022131-0220222012313033-0130232220220211-0313001122133110-1213302233312212-1131203132301301-3023321131232023-3221213321023001"></a>

## no_network_policy — no_network_policy / 333013101131 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- no_network_policy

<a id="canonical-0020220320210233-1230003021333300-3003201101220303-1103313102222220-0120311331012121-0322133310300111-0322120121200032-2332221121003102"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

Upstream description:

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

<a id="canonical-2213223022223213-0100232300223023-0123103233032110-2201120123321120-3011130310001021-0322033022132302-1011201113220031-1302013033311223"></a>

## Direct properties — no_network_policy / 333013101131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200113222311233-2032321323301303-1100030113123332-2022101233303103-3023330330102322-0011133112023221-2231131220220211-2100033002120310"></a>

## Next pages — no_network_policy / 333013101131 / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3320113330113312-0230001301120322-3002203321220201-2321010231000321-1003133122001312-0312021000200203-3112012033001130-1333100212331010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000101130001303-3201033032113022-3103230233123012-1223310133232023-2323313011313122-1220131201103300-1230102200322011-2131321202003022"></a>

## no_proxy_bypass — no_proxy_bypass / 213032012213 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- no_proxy_bypass

<a id="canonical-3233031100200303-2012031220300110-1200001122032311-3300310213323033-0220300312010320-1001231012313000-0301200031232320-1231203313021022"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no proxy bypass.

Upstream description:

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

<a id="canonical-1301123333222100-2002311133031231-1013330000203310-3010132231103211-2201211213132110-0110221103020000-0032310002332233-2022133330323322"></a>

## Direct properties — no_proxy_bypass / 213032012213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302111110010112-1123101201002202-0111002220230102-0321221301123123-2312101322122330-1022210323203300-3230100023200310-0223302013211000"></a>

## Next pages — no_proxy_bypass / 213032012213 / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1323111133021223-1232233001220010-1033101302021013-2102121030310221-2202010203020030-2300133202310213-2120230100101000-2011231333110010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220103113001331-3230103312010103-1210231113201220-1213223210103001-2021322020333301-0221033323203123-3130320023120032-0131212303102300"></a>

## no_s2s_connectivity_sli — no_s2s_connectivity_sli / 323300123000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- no_s2s_connectivity_sli

<a id="canonical-0023211121101021-3112022133311112-1013000103021013-3131302032212312-3112130132131102-0300211302323112-1110300220332030-2202211101320312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no s2s connectivity sli.

Upstream description:

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

<a id="canonical-2332130003101022-1023200132132100-1033123021322122-3011202303023230-0222022301112100-1320120300222132-2032210120032103-1103030010220031"></a>

## Direct properties — no_s2s_connectivity_sli / 323300123000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120302111121132-3233232332211132-3313300121131023-0001103332102130-1002011320201233-0130213223322320-1011231122233131-3321300202333322"></a>

## Next pages — no_s2s_connectivity_sli / 323300123000 / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2020332030132111-0211012220122133-0321312221332102-2200333121310232-3111223133212200-2331110022130102-3133103133003221-2231231211133133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203110030320110-2112130133332200-0313213331010113-2300201323021010-1223131110202020-1333003020311202-1333303313020011-1323030302123003"></a>

## no_s2s_connectivity_slo — no_s2s_connectivity_slo / 211321300203 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- no_s2s_connectivity_slo

<a id="canonical-0230110110333031-0311111112012121-3303032310321332-0313321023302122-1120302120112112-0333332102333221-3131131213021211-3102303131313322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no s2s connectivity slo.

Upstream description:

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

<a id="canonical-0133321201102113-0220233313230213-0331332131020323-3032312332222301-1322021021323320-2220333132111323-3302223203233301-2101001313221103"></a>

## Direct properties — no_s2s_connectivity_slo / 211321300203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122323131221201-1121310123300133-2021230120302120-3101023113131120-1102330320230332-2111021302010321-3212222123321012-3310023322013200"></a>

## Next pages — no_s2s_connectivity_slo / 211321300203 / 4

- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023100013302010-1203121212200322-2101303012100021-1111113232011302-2012212010203023-3012203103120031-2220023102012333-1320232311232010"></a>

## nutanix — nutanix / 223210012232 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- nutanix

<a id="canonical-3332230322311330-3002233110322202-1013300122022103-1102232321033032-1321221122311231-0102010101321012-3100100301113010-2103000300030332"></a>

Type: `"single"`. Computed.

Nutanix Provider Type. Nutanix Provider Type.

Upstream description:

Nutanix Provider Type.

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

<a id="canonical-1111303031303032-3332132213122301-3220031333323223-1303011301211130-2122021101222313-2031021323310000-0011330100223033-3203122320111233"></a>

## Direct properties — nutanix / 223210012232 / 3

- [not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210): complete subsection reference.

<a id="canonical-3123332301111113-0122000303203201-1313123120030103-0300013132301302-0322222102233131-0022120202200101-3230321020313312-2130120223200312"></a>

## Next pages — nutanix / 223210012232 / 4

- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203111323221032-0003103210230301-0002010031011312-2010332322012002-3311200021311103-3000100021102121-1020110011131230-3331130112333300"></a>

## nutanix.not_managed — not_managed / 230332220322 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- nutanix.not_managed

<a id="canonical-3010312013300001-3022021300030233-2232123022231212-0121010110222011-1100210231020023-2023010202312210-3231323231312130-2313313020321121"></a>

Type: `"single"`. Computed.

Section will show nodes associated with this site.

Upstream description:

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

<a id="canonical-2122232232323102-2321323223031233-2032210323322020-3302301103333102-2311221303312333-3203322133210111-2112213032230031-2203333133032213"></a>

## Direct properties — not_managed / 230332220322 / 3

- [node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213): complete subsection reference.

<a id="canonical-3232033311001322-1030320321110102-3133310200110301-3302120330201213-3322031303230023-0323202010010332-0130302132202300-3023033223321130"></a>

## Next pages — not_managed / 230332220322 / 4

- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103231223302221-2123321022011230-0300211113210333-1302033023231003-0310110012112023-1111012103022112-3010210002320103-0323012133101002"></a>

## nutanix.not_managed.node_list — node_list / 112233321322 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- nutanix.not_managed.node_list

<a id="canonical-2210330020132033-2020320230202002-3301121131030320-2001321022200022-0222021320103101-0123221233332202-3003200032311301-1311110130301131"></a>

Type: `"list"`. Computed.

Section will show nodes associated with this site.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2133033303330313-1012112202201010-2231122321210310-3202302221333320-1312323023220203-3032110210112133-2203220112310322-1013122003210232"></a>

## Direct properties — node_list / 112233321322 / 3

<a id="canonical-3321130331300330-1201210123230233-1221021013220202-3130311113222210-0320003311233113-2103031303212223-2212033133121200-0312111133312222"></a>

<a id="canonical-3231000311013321-0121322101103102-0023022321022132-1202000102311102-1020222313231133-0222330110302130-0001010121322332-2032133031222013"></a>

## hostname property — node_list / 112233321322 / 4

Type: `"string"`. Computed.

Hostname. Hostname for this Node.

Upstream description:

Hostname for this Node.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232): complete subsection reference.

<a id="canonical-2323101212233122-0103202112121310-1220100201330202-2011110330331123-1333211112212213-0103303203021101-0230333011010333-1232020022313101"></a>

<a id="canonical-1103211301123323-0211313030130212-0321012213123222-0010200123003101-0203330301331113-3201211003010201-3111113010330103-0121100103133120"></a>

## public_ip property — node_list / 112233321322 / 5

Type: `"string"`. Computed.

Public IP. Public IP for this Node.

Upstream description:

Public IP for this Node.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0120133231203223-2210323113122301-3131011000301303-1020203101021110-2131321122211232-0130212132331330-3321331112231112-0002320223332200"></a>

<a id="canonical-3102131013202310-2112321213022223-3133210220100133-0030211013021231-0111203330303023-0133103223303300-3232321222321012-3123222123212133"></a>

## type property — node_list / 112233321322 / 6

Type: `"string"`. Computed.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Upstream description:

Type for this Node, can be Control or Worker.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0133031212232310-0302130210312301-1322011200310333-2222123013313233-0101123003220122-1203022233111023-1311113310223030-1001023232311312"></a>

## Next pages — node_list / 112233321322 / 7

- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323231200110001-1023102000302231-0332300023103111-1222021211332101-1211220121230111-3322011122002122-0323003102130101-3313200120112313"></a>

## nutanix.not_managed.node_list.interface_list — interface_list / 112101311122 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- nutanix.not_managed.node_list.interface_list

<a id="canonical-1001101002312322-1033111112002333-0100021120220121-2130233231303013-0203213130313122-2132011322230321-0203211000312312-0211030212313020"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2333330020033303-2122222010120332-3202033221202220-1312333322213301-3222213130110133-1021311213013201-2011222203022031-0221013210133332"></a>

## Direct properties — interface_list / 112101311122 / 3

- [bond_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2330011002101020-1223223110230200-3020033103222000-0113023100000110-3031201222133100-3221100231032232-3310103332223111-2323020323130330): complete subsection reference.

<a id="canonical-0120132310102103-2001102020233013-1201011312020213-0001321232132120-2101202301030012-0000121132331333-0300120230001101-2210330112023022"></a>

<a id="canonical-3110321222212211-0020001211223100-3102323231220211-0332333312311012-1112030110113022-3130121121320303-2032223033110033-3333020212320121"></a>

## description_spec property — interface_list / 112101311122 / 4

Type: `"string"`. Computed.

Interface Description. Description for this Interface.

- [dhcp_client](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3103231010110021-2023313200103100-1120310120232031-3322130103032011-2120210133322033-0230203202133330-0003331122120101-3020301001222200): complete subsection reference.

- [dhcp_server](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301): complete subsection reference.

- [ethernet_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0222200321012331-3320100132003121-1232322313113033-3320300022010221-2033103203122030-2301022032313011-1130313202022011-1320012112122330): complete subsection reference.

- [ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310): complete subsection reference.

<a id="canonical-3221303201012201-0023011222013010-3123221023013203-0300211031012230-2033120330332111-0112020322222112-0230203310333233-2023010233020023"></a>

<a id="canonical-1210131331210203-3000320101131111-3300211232310000-0323123332333332-0100013032103200-3100021233222100-0313111032313301-2321300210332301"></a>

## is_management property — interface_list / 112101311122 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-1112201300330320-1110010022021031-2023012310323131-0103203103100210-3310330000323013-3322033333222211-0102311111332110-0112223022032300"></a>

<a id="canonical-3102221300010003-2222001231132221-2330100010313201-0130131131132320-1130012211222030-1300303312223031-1122200031300220-0030032113012122"></a>

## is_primary property — interface_list / 112101311122 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-0131030103203330-0233001110230011-2220112122013023-3221110321133032-1300003210001002-1232101000233100-2300013331010103-0321123200130310"></a>

<a id="canonical-0023010230132110-3213222331302032-2011310233101302-2131320310231211-2332002131200033-2130300223323220-1330330313022031-1130302033031100"></a>

## labels property — interface_list / 112101311122 / 7

Type: `["map", "string"]`. Computed.

Add Labels for this Interface, these labels can be used in firewall policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 16,
    "metadata": {
      "confidence": 0.75,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [monitor](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1220032203233122-3000212111302131-0230322231210211-3020210030030131-1000102230112333-1122111322013123-2211001223013202-3003201323213133): complete subsection reference.

- [monitor_disabled](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2222301013131301-2321201000103230-0301002131232033-3123101310023021-2322333331323021-3303211321101132-0212201100003121-1121223011310012): complete subsection reference.

<a id="canonical-3033211100021323-3333131002202002-0203122131000103-2113230321212033-0230310302221333-2302200201122020-3210310333100112-0020220231133321"></a>

<a id="canonical-3131330020230213-0213131013100303-1123133212022321-2110102202320202-2333303202233131-1123101131013032-1122201301022312-0002133022303231"></a>

## mtu property — interface_list / 112101311122 / 8

Type: `"number"`. Computed.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-3001321302202231-2200301322202302-3132223212133301-0221100230133100-1110312301011122-2112123223221000-1020332331030113-2312012112320200"></a>

<a id="canonical-1212302230300322-2000202102032322-1233300003310230-0200232332013000-0311113020323023-1221301103212133-2322311320223222-0213030312033312"></a>

## name property — interface_list / 112101311122 / 9

Type: `"string"`. Computed.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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

- [network_option](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3122221332322122-3021333222201103-3001233312030203-1201202202233003-0032032230133122-2310002332102212-1320322123313111-3212003322223323): complete subsection reference.

- [no_ipv4_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2110120310132301-0333323103202113-0203001122333010-3120002212310310-3301321201101322-1330321201202130-3301023300132031-2320000213313113): complete subsection reference.

- [no_ipv6_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1201132021103012-0320121210130102-0320131023123203-1223103021112230-1212210112010230-0223322113122123-2230110002200012-1133330031212103): complete subsection reference.

<a id="canonical-0213310131211130-1230113213301223-2033231221021133-1211033001321201-1332331031120011-3222110300000233-2121333010322022-1113232303013021"></a>

<a id="canonical-3232331301021313-0102333201123013-1330232020320011-0303020323020220-2100101323031111-1312210032223312-2130131220323232-0023233301001013"></a>

## priority property — interface_list / 112101311122 / 10

Type: `"number"`. Computed.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3321103030123112-1111002332222312-0001330332230202-1210023031020311-3012220032001012-3202211200231032-3201333232122130-0012331300100203): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3122312333000212-1020203202012230-2222133022312222-1000313330102330-1210033002310300-0122320110232122-1002322113323212-1220111120020002): complete subsection reference.

- [static_ip](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3121320200010110-3202032231021203-1123033202300131-0321201113233330-0000003322120203-2310101203203111-1233332211300133-0131130000003112): complete subsection reference.

- [static_ipv6_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0003110032212313-2110102033030300-1031332222211220-2033001121110303-3300231321101333-0201220202000310-3213320201213033-3132232111110230): complete subsection reference.

- [vlan_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2212102230003211-3201300121123111-2100131201300322-1130003101002213-1130032130203003-3233330110331212-3022101331031233-2322323020313313): complete subsection reference.

<a id="canonical-1230220200003111-0101202300203123-1102232131201012-3031101003331031-3330210212030233-3100233330010030-2012332201023111-3103332111321202"></a>

## Next pages — interface_list / 112101311122 / 11

- [nutanix.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2330011002101020-1223223110230200-3020033103222000-0113023100000110-3031201222133100-3221100231032232-3310103332223111-2323020323130330)
- [nutanix.not_managed.node_list.interface_list.dhcp_client](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3103231010110021-2023313200103100-1120310120232031-3322130103032011-2120210133322033-0230203202133330-0003331122120101-3020301001222200)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- [nutanix.not_managed.node_list.interface_list.ethernet_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0222200321012331-3320100132003121-1232322313113033-3320300022010221-2033103203122030-2301022032313011-1130313202022011-1320012112122330)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [nutanix.not_managed.node_list.interface_list.monitor](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1220032203233122-3000212111302131-0230322231210211-3020210030030131-1000102230112333-1122111322013123-2211001223013202-3003201323213133)
- [nutanix.not_managed.node_list.interface_list.monitor_disabled](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2222301013131301-2321201000103230-0301002131232033-3123101310023021-2322333331323021-3303211321101132-0212201100003121-1121223011310012)
- [nutanix.not_managed.node_list.interface_list.network_option](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3122221332322122-3021333222201103-3001233312030203-1201202202233003-0032032230133122-2310002332102212-1320322123313111-3212003322223323)
- [nutanix.not_managed.node_list.interface_list.no_ipv4_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2110120310132301-0333323103202113-0203001122333010-3120002212310310-3301321201101322-1330321201202130-3301023300132031-2320000213313113)
- [nutanix.not_managed.node_list.interface_list.no_ipv6_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-1201132021103012-0320121210130102-0320131023123203-1223103021112230-1212210112010230-0223322113122123-2230110002200012-1133330031212103)
- [nutanix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3321103030123112-1111002332222312-0001330332230202-1210023031020311-3012220032001012-3202211200231032-3201333232122130-0012331300100203)
- [nutanix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3122312333000212-1020203202012230-2222133022312222-1000313330102330-1210033002310300-0122320110232122-1002322113323212-1220111120020002)
- [nutanix.not_managed.node_list.interface_list.static_ip](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3121320200010110-3202032231021203-1123033202300131-0321201113233330-0000003322120203-2310101203203111-1233332211300133-0131130000003112)
- [nutanix.not_managed.node_list.interface_list.static_ipv6_address](data-sources--securemesh_site_v2--reference--group-013.md#canonical-0003110032212313-2110102033030300-1031332222211220-2033001121110303-3300231321101333-0201220202000310-3213320201213033-3132232111110230)
- [nutanix.not_managed.node_list.interface_list.vlan_interface](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2212102230003211-3201300121123111-2100131201300322-1130003101002213-1130032130203003-3233330110331212-3022101331031233-2322323020313313)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2330011002101020-1223223110230200-3020033103222000-0113023100000110-3031201222133100-3221100231032232-3310103332223111-2323020323130330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332320022323201-1212101301012332-2210200102210322-2211121223012020-0112203302103020-0302021333223310-1102130000021222-1000230113010322"></a>

## nutanix.not_managed.node_list.interface_list.bond_interface — bond_interface / 223030322123 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.bond_interface

<a id="canonical-2221102133032101-1123113121032131-3202321333313030-2030210222201222-2201033031131000-0313233223320113-0312321302130232-1233211303221130"></a>

Type: `"single"`. Computed.

Configuration parameter for bond interface.

Upstream description:

Bond devices configuration for fleet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

<a id="canonical-2120110132013001-0010120023111331-0301013211131232-1330130130323301-1033033322330111-0120330221213330-3220203322031220-1212320020211313"></a>

## Direct properties — bond_interface / 223030322123 / 3

- [active_backup](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1031000000211333-1221332320030212-3020123121031020-2133302033310323-0113023011212323-1230220022020002-3112301203121030-0010101023221130): complete subsection reference.

<a id="canonical-0330001021300103-2202231231313130-0033310310130022-2230032011133031-3313223233301021-0130102013112211-2003031131131130-1033011021223223"></a>

<a id="canonical-2221030201030330-2210013220121232-3002320310303322-2133031131201111-2110223131103002-1201330011101102-3112033321011203-0133122303220133"></a>

## devices property — bond_interface / 223030322123 / 4

Type: `["list", "string"]`. Computed.

Ethernet devices that will make up this bond.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [lacp](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1300003021233213-1303201221113122-3021131012203213-2130313303131202-0133130100123332-3201113233020330-2232003032210121-1320120302330111): complete subsection reference.

<a id="canonical-3210223010231313-3203232221001132-0023020221100111-1130221302130211-3201322033113130-0102122001302101-3212101110210111-2321322130200303"></a>

<a id="canonical-2023000001333033-0131321233033120-2302200122131233-3101303323023312-3100020332230121-3303210133013103-1201220022101110-1201000201001103"></a>

## link_polling_interval property — bond_interface / 223030322123 / 5

Type: `"number"`. Computed.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 500
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-0203220311321102-2020033221031123-3223220103300030-3212033132311312-3010022222000130-1212211013302100-0102023213102000-3202002231302302"></a>

<a id="canonical-3200330231302223-1011220103312312-3201101213322103-3321022333032101-0232321122023233-2012130111213302-0222001113101130-2002312003001303"></a>

## link_up_delay property — bond_interface / 223030322123 / 6

Type: `"number"`. Computed.

Milliseconds wait before link is declared up.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  }
}
```

<a id="canonical-2223003102032010-2233212302113300-2211230103120033-3131022200313331-2120130221003230-0133131103013221-3312332100223330-1021230111201211"></a>

<a id="canonical-3112001313121310-0321220303010220-0221121202221312-3010320330133331-3313202000132311-3111303020311303-2223302013023333-2313000011222110"></a>

## name property — bond_interface / 223030322123 / 7

Type: `"string"`. Computed.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-1002102211110331-3301320021220211-3031031211220122-3221031102012102-3213031000133022-1101202331102023-1302120202131101-3101110110233200"></a>

## Next pages — bond_interface / 223030322123 / 8

- [nutanix.not_managed.node_list.interface_list.bond_interface.active_backup](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1031000000211333-1221332320030212-3020123121031020-2133302033310323-0113023011212323-1230220022020002-3112301203121030-0010101023221130)
- [nutanix.not_managed.node_list.interface_list.bond_interface.lacp](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1300003021233213-1303201221113122-3021131012203213-2130313303131202-0133130100123332-3201113233020330-2232003032210121-1320120302330111)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1031000000211333-1221332320030212-3020123121031020-2133302033310323-0113023011212323-1230220022020002-3112301203121030-0010101023221130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010311031201212-2323202032101122-0033332202031203-0100323103323322-1220211021211102-2120123021033003-3031223122320122-1101002013320223"></a>

## nutanix.not_managed.node_list.interface_list.bond_interface.active_backup — active_backup / 021331101222 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2330011002101020-1223223110230200-3020033103222000-0113023100000110-3031201222133100-3221100231032232-3310103332223111-2323020323130330)
- nutanix.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-1001013221000312-1030023022122211-2000103020012312-3030321333303031-1000223321031203-3102222132002332-3122110102011310-3122103113010303"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for active backup.

Upstream description:

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

<a id="canonical-2220201132310013-3033221333012203-2123223202012020-2303203130110103-1120230121200100-2232000130101320-2030001211300202-1132210202223300"></a>

## Direct properties — active_backup / 021331101222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002002013220203-2102201012210231-2213130323321002-1230223223223012-2201000000123212-1101230133330100-3200003320030302-0330001232200201"></a>

## Next pages — active_backup / 021331101222 / 4

- [nutanix.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2330011002101020-1223223110230200-3020033103222000-0113023100000110-3031201222133100-3221100231032232-3310103332223111-2323020323130330)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1300003021233213-1303201221113122-3021131012203213-2130313303131202-0133130100123332-3201113233020330-2232003032210121-1320120302330111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221013121330033-3033220311210020-2320010200320311-3332120000222023-1320230322112303-0213033103211313-1111013111331100-3320202002302311"></a>

## nutanix.not_managed.node_list.interface_list.bond_interface.lacp — lacp / 122122033302 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2330011002101020-1223223110230200-3020033103222000-0113023100000110-3031201222133100-3221100231032232-3310103332223111-2323020323130330)
- nutanix.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-0112231023132002-1300000222322121-2032210203022220-2132222013203203-3221013121222300-3103332221221013-1032333230230213-2011213201231320"></a>

Type: `"single"`. Computed.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

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

<a id="canonical-2122022333332122-2110313330030113-0220312313202231-2011022303102233-1200210131312231-1230200200300231-0020113312031221-2310022003311233"></a>

## Direct properties — lacp / 122122033302 / 3

<a id="canonical-3113203022213021-3122131002123322-2103102332202101-0000002132002312-0223121022030033-2212230021200012-0213313233033000-0233101203210220"></a>

<a id="canonical-2322311013133130-0312200011122321-1210031112323200-2123310103301013-1230010231323003-1232200320313232-2021011331323011-2020330020030120"></a>

## rate property — lacp / 122122033302 / 4

Type: `"number"`. Computed.

Interval in seconds to transmit LACP packets.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-2303322312013020-2311332323312003-0310023101220230-3213230320222132-3200210031322203-1100032233111000-2131322320021313-2113211303111033"></a>

## Next pages — lacp / 122122033302 / 5

- [nutanix.not_managed.node_list.interface_list.bond_interface](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2330011002101020-1223223110230200-3020033103222000-0113023100000110-3031201222133100-3221100231032232-3310103332223111-2323020323130330)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3103231010110021-2023313200103100-1120310120232031-3322130103032011-2120210133322033-0230203202133330-0003331122120101-3020301001222200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313331110333332-0132012302213210-2212100111222321-3131313010303221-0022101213313031-2310333120233302-2321312010321123-0021321100313120"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_client — dhcp_client / 030132331033 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-3232300113010213-3231101023310320-1130331002223112-1103011303200312-2130011033100023-0233202132112112-1312033122320212-1221021102122233"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-2011222022121132-2102102330023031-0022231323023011-0110003020231112-3000323013321113-2310120220321112-0221021202011111-2213133313330103"></a>

## Direct properties — dhcp_client / 030132331033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000303123232130-0220033302301302-1200221213202000-1001323022102032-3223203312201031-3033311201010122-1212031022220103-2312030311332232"></a>

## Next pages — dhcp_client / 030132331033 / 4

- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123310230212330-0213330113130121-0312122303321020-0031230103023002-1312032231311123-1213221013322003-1003312021010033-0110100030101331"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server — dhcp_server / 100113023312 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-2102032022011120-2223210330220312-2130032102113030-3203123322123001-2130231301120122-2113012022002220-2202303130002112-0322120022110103"></a>

Type: `"single"`. Computed.

DHCPServerParametersType.

Upstream description:

DHCP server configuration for this interface.

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

<a id="canonical-2101312321131332-2322113223110121-0322013330223230-3333011132101333-3221121030003303-2030333020111010-0103200301103233-2033021110111131"></a>

## Direct properties — dhcp_server / 100113023312 / 3

- [automatic_from_end](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0130230232002332-3103200103322132-0202030322021033-1212110132333311-0212210211202002-1210303130100231-0003101100232312-2131203110203303): complete subsection reference.

- [automatic_from_start](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1212222023033302-0131021331120213-2313231302013010-3132003212103333-3122200233201232-2202011000031012-3202313133302320-1320012313103012): complete subsection reference.

- [dhcp_networks](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132): complete subsection reference.

<a id="canonical-0132022212021123-1330130332000211-1301101223330113-3002301103030031-0132122331023001-0203222010321213-0130010121213133-3133321110011102"></a>

<a id="canonical-3111132022003202-1233013221330310-3223023333003203-0112100000021233-2323102311310021-1002000233232121-2312232121201233-1130000031330123"></a>

## dhcp_option82_tag property — dhcp_server / 100113023312 / 4

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-1132200103101132-0113323130103000-2331003311033322-0233302300013130-1213003111223033-2303303210312010-2313123023303013-3011321203200121"></a>

<a id="canonical-2301303123130103-0332132133203003-1030302100313012-0203101320002001-3323013103123231-2233001313322111-0111030312002313-1113103332321231"></a>

## fixed_ip_map property — dhcp_server / 100113023312 / 5

Type: `["map", "string"]`. Computed.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 128,
    "metadata": {
      "confidence": 0.75,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

- [interface_ip_map](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2002323301310213-1213012000313020-2333213003320003-2331030031310311-2233212333111220-1001212101011123-3102030011303221-0130012310003212): complete subsection reference.

<a id="canonical-3233113020123012-1003323301203210-2312231210103023-2200103302023132-2013233012133010-3213300000221111-1002222301213121-1103013023223020"></a>

## Next pages — dhcp_server / 100113023312 / 6

- [nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0130230232002332-3103200103322132-0202030322021033-1212110132333311-0212210211202002-1210303130100231-0003101100232312-2131203110203303)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1212222023033302-0131021331120213-2313231302013010-3132003212103333-3122200233201232-2202011000031012-3202313133302320-1320012313103012)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2002323301310213-1213012000313020-2333213003320003-2331030031310311-2233212333111220-1001212101011123-3102030011303221-0130012310003212)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0130230232002332-3103200103322132-0202030322021033-1212110132333311-0212210211202002-1210303130100231-0003101100232312-2131203110203303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302103123031000-1011100310123322-0210220023130232-2312220130011300-3111021313300103-2122023121013123-2223001100212212-2130202231220322"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — automatic_from_end / 212112100130 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-0021300122031002-1123001313112223-3030301220231212-3313220313002233-3233002311321013-1121301113103320-3031132020003200-2102020210232331"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from end.

Upstream description:

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

<a id="canonical-3113002323212123-3321133333213323-3322013201232301-0221121203012103-3223021031212323-3331100223100002-1302302211011123-3233110121033313"></a>

## Direct properties — automatic_from_end / 212112100130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211020122333213-3030320102331123-1323021101030102-3032311020311021-3233333200030023-3121230013213330-3110310323000320-0030310120221101"></a>

## Next pages — automatic_from_end / 212112100130 / 4

- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1212222023033302-0131021331120213-2313231302013010-3132003212103333-3122200233201232-2202011000031012-3202313133302320-1320012313103012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323122303030003-0223123212201132-1331210213002121-1311033212033112-1103101232213300-2203323233111202-0323122131012103-2132033221010202"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — automatic_from_start / 001221323133 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-3332310111131201-1221132223220212-2231022111300121-0322032211310321-1110332132132121-2133101322302112-2012302100210000-2231011021310210"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for automatic from start.

Upstream description:

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

<a id="canonical-2031121233123200-0003030320110311-3103212013011111-0213100330020321-0032212312113102-0312211203332001-1223011032032313-2012322112331332"></a>

## Direct properties — automatic_from_start / 001221323133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011112001130323-1001302203331023-3313210003311310-3333031223322332-3210113330033113-2121220122202320-0312032300100103-3310102130331033"></a>

## Next pages — automatic_from_start / 001221323133 / 4

- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011300002220323-3111310123023310-0030232133302333-3200103200013300-3333022120000210-0030012232321323-2200311030013332-0023011301011132"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — dhcp_networks / 321031221001 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-3221321122033231-1232332333110201-2201210323310302-3320300020332312-3300232123132030-1132211022332302-0313011102310102-2001113212321321"></a>

Type: `"list"`. Computed.

List of networks from which DHCP Server can allocate IPv4 Addresses.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1013021220030200-1201030233232102-1222113130220223-2222301321021120-1211003031103123-3302120010132010-3212232203222012-3312002312102113"></a>

## Direct properties — dhcp_networks / 321031221001 / 3

<a id="canonical-3301110231321221-2213100320021122-2022002212302323-1023031011333323-1102032112023323-3003021323011310-2030022121023131-1211332300112002"></a>

<a id="canonical-0200321223321200-0031011202233001-2323002202013223-3113030002021010-0012211120303122-1311001233332313-1300311212010332-3313122323002030"></a>

## dgw_address property — dhcp_networks / 321031221001 / 4

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0212121221222302-1020031033203022-1131010200121232-1300222301222233-0010012103133220-0230110022220023-0022003231103012-1330211133313033"></a>

<a id="canonical-1333300311001021-1222130000222033-2013232303110222-1003330211133022-0020220313233222-3331122323021220-3000200201133013-0331011102300230"></a>

## dns_address property — dhcp_networks / 321031221001 / 5

Type: `"string"`. Computed.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [first_address](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1120011021231312-3232333011210302-3202313310121201-3213202303001230-2220113310323332-3012323303320200-0131132303203233-2332202311231212): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3023232223101013-2203301220230331-3023110200233002-1013000300300313-1110110333133121-3220233122302321-3010133000012321-0303112203123232): complete subsection reference.

<a id="canonical-2003020111312230-3223101312122232-1011010213021030-3311221211033011-2103130131101213-3202303013312200-3203011300003001-0133301130320223"></a>

<a id="canonical-2320011222300231-1233113121033213-0321301210123130-1002223310132230-3102321123331310-0223031321011100-1203331011232133-2323202012212320"></a>

## network_prefix property — dhcp_networks / 321031221001 / 6

Type: `"string"`. Computed.

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

Upstream description:

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-3023020131100201-3212011332322002-0103103331022322-0032132312101100-3322121220002030-1200210113232030-0002212232021030-0302303031002102"></a>

<a id="canonical-1333222231221310-0202010230230333-0110300322221133-0133101102322200-2032300320331200-3320233210003132-2000223112130000-1120013033200023"></a>

## pool_settings property — dhcp_networks / 321031221001 / 7

Type: `"string"`. Computed.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

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

- [pools](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1123201321301201-3112230103221232-1023001212201322-0201333330330013-3300332312002102-1133220123331133-1021312221033002-2103011020213102): complete subsection reference.

- [same_as_dgw](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1302102333231303-2222300123110202-0100112000321030-3330312132122011-2213323221002200-0121010331113323-1210120313311211-1130323230333233): complete subsection reference.

<a id="canonical-1231221021133303-0210013030023111-3221200321032210-1003231201001110-3320210313230010-0231333021212122-0301022302202330-3132021120001331"></a>

## Next pages — dhcp_networks / 321031221001 / 8

- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1120011021231312-3232333011210302-3202313310121201-3213202303001230-2220113310323332-3012323303320200-0131132303203233-2332202311231212)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3023232223101013-2203301220230331-3023110200233002-1013000300300313-1110110333133121-3220233122302321-3010133000012321-0303112203123232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1123201321301201-3112230103221232-1023001212201322-0201333330330013-3300332312002102-1133220123331133-1021312221033002-2103011020213102)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1302102333231303-2222300123110202-0100112000321030-3330312132122011-2213323221002200-0121010331113323-1210120313311211-1130323230333233)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1120011021231312-3232333011210302-3202313310121201-3213202303001230-2220113310323332-3012323303320200-0131132303203233-2332202311231212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232302203202210-3133011303121130-3323320003102223-2311303011032300-3231320230033301-1233111220022113-0330112222003120-0110122012212230"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — first_address / 301013020320 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-1312321333200123-2311130233023310-3000010321311002-1230112003302032-0022010203111103-1311321122130231-1321103021312110-1011231013110211"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-2020311010012101-1102100212223121-0201003002311113-3033013212230310-0220321330301001-1200210300110320-1300002302002202-1132221021300302"></a>

## Direct properties — first_address / 301013020320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3200210211132213-2233311120200331-0103103033131102-3311033311220003-0211031013023330-3201022200001013-0033201221223000-2230313333211112"></a>

## Next pages — first_address / 301013020320 / 4

- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3023232223101013-2203301220230331-3023110200233002-1013000300300313-1110110333133121-3220233122302321-3010133000012321-0303112203123232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032212001202021-3200130023013101-2023020001301000-2023210322233001-1333321222011220-1000332033211321-1020202112222322-3210200012130000"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — last_address / 112011022312 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-1121110330103133-1333310113133101-0320120222113313-2033032331223312-0001132113010003-0103003232121333-3200023122110101-2130221022111322"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-0222233331100332-1223313000301133-3032321011002011-3020033011230331-2130033012300031-0203133302322223-3023310200312120-3301021011232332"></a>

## Direct properties — last_address / 112011022312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011020100100331-2323120232311311-3230113310033203-1223123022123010-0020222132021321-0013021112233233-2123300013011002-1210220212223133"></a>

## Next pages — last_address / 112011022312 / 4

- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1123201321301201-3112230103221232-1023001212201322-0201333330330013-3300332312002102-1133220123331133-1021312221033002-2103011020213102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003110323032321-1221031313311201-1322331203110221-2103210003232011-1013231303201030-2131123122132123-2200203121313221-1022012023011122"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — pools / 213301033221 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-2130022320201312-1300232122112021-3211030300110231-0120231102212130-1201021312332010-1110112211201302-2230210002331232-2331332002212312"></a>

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2220031231202110-2301103101020112-0223130122011020-0323300232332213-1100001231011321-0001023213120221-1211121333100201-2310211301201323"></a>

## Direct properties — pools / 213301033221 / 3

<a id="canonical-3301123201000033-1102213201101332-3021301220230203-1320112102113232-2023213210133233-3132123202200302-0202112000230330-3121223003023133"></a>

<a id="canonical-0110123203021003-1221221101311001-3331122210222130-0301000000331322-0032332003201122-1120200320221112-0221132202102123-3022023002310102"></a>

## end_ip property — pools / 213301033221 / 4

Type: `"string"`. Computed.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0020213213310220-1221333202233311-3001330211111131-1320301133300101-1111122331131232-1300000221120010-0230213223102030-2310032300331001"></a>

<a id="canonical-1301100223011022-1231022330203013-0131210101200230-2300112220120222-3122321101213201-3111110011330231-2021212322122212-2131223000113303"></a>

## exclude property — pools / 213301033221 / 5

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-2312210020231220-3331002331303303-1213320331233122-1103200220330122-2311032310101020-1111010321203033-2202121301220232-1001012111021112"></a>

<a id="canonical-3130020001132132-0210032011100030-1101231333102130-3301011200302331-2332202221112101-1110222111002300-1222220121023102-1100111201302013"></a>

## start_ip property — pools / 213301033221 / 6

Type: `"string"`. Computed.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1300113310011201-1000233331120330-2332333031013323-1032023022323200-3010121022032222-0203313123120103-0002122113220100-2122201111023121"></a>

## Next pages — pools / 213301033221 / 7

- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1302102333231303-2222300123110202-0100112000321030-3330312132122011-2213323221002200-0121010331113323-1210120313311211-1130323230333233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110123210031023-0113122213212133-0231013223021003-1213113313102032-3133300033223323-3200211032120322-3301331321300010-2131112121123233"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — same_as_dgw / 100020023031 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-3120231213032121-1222120120332023-2123121021111001-1001303022222221-2000303000212301-0202121013211222-1331300320210101-3300322323211010"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for same as dgw.

Upstream description:

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

<a id="canonical-3310122302022203-1112231112100010-2213110203233001-3001223223013232-1312321312021300-3031202033220132-3130312132101132-1031010101001023"></a>

## Direct properties — same_as_dgw / 100020023031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021320002011033-1212310131310222-2320333330301102-2321112102020221-0022033101023002-3221301232310322-1312203213300323-0100010022102110"></a>

## Next pages — same_as_dgw / 100020023031 / 4

- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2330323311221320-0222022211322133-1021333233133132-1031012112122322-0031103130001200-2120103310020200-1033102012313310-0312232120313132)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-2002323301310213-1213012000313020-2333213003320003-2331030031310311-2233212333111220-1001212101011123-3102030011303221-0130012310003212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230232302300211-3021212200332021-3230013003312031-3301233000110322-0033030220321333-3333302010321322-3113001113230032-3223213003100000"></a>

## nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — interface_ip_map / 333110233333 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-0003131210203313-3103011320321233-2112103002002020-3113201133210001-2031213031003003-3310313110222311-0231202320111133-3223023233210222"></a>

Type: `"single"`. Computed.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

Upstream description:

Specify static IPv4 addresses per node.

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

<a id="canonical-3202020012210132-0020223001231123-3331232212100111-3101112320033123-3002210302200223-0303302012212032-3312331232032032-1332221311130322"></a>

## Direct properties — interface_ip_map / 333110233333 / 3

<a id="canonical-2233120221233032-3202023002030332-1322021331031210-3103312301331310-1221201103111223-2001132113100202-3101311312031310-0210003203111210"></a>

<a id="canonical-0001121123202231-3223313323310003-3302202111003333-3210021233023231-0230023222330310-0323133231301213-1112213131103300-2331223101033132"></a>

## interface_ip_map property — interface_ip_map / 333110233333 / 4

Type: `["map", "string"]`. Computed.

Specify static IPv4 addresses per site:node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 64,
    "metadata": {
      "confidence": 0.75,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

<a id="canonical-3331011313113333-1331001230230033-0333233122201223-2301303330332231-3212202033020111-2122332101333130-0012031100322102-0130332001211210"></a>

## Next pages — interface_ip_map / 333110233333 / 5

- [nutanix.not_managed.node_list.interface_list.dhcp_server](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3011220303221131-2021230331122313-1331001300111202-0201110021320212-2101030120032332-0021103310222213-3230203220202103-3320020311123301)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-0222200321012331-3320100132003121-1232322313113033-3320300022010221-2033103203122030-2301022032313011-1130313202022011-1320012112122330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311100233332002-1331101213023301-0111201300312002-2201031211323021-0300003202031203-1033222132131330-2330203023010100-0023311210131133"></a>

## nutanix.not_managed.node_list.interface_list.ethernet_interface — ethernet_interface / 031232331002 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-3010223031230333-0330003011110110-0100320311310000-3322331123130322-1322320132201232-1301022221021311-3112313003231100-2000103223231101"></a>

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

<a id="canonical-3300312331213310-3323132200223302-0100010300331203-2101012111021300-2002320210120031-0123001133331330-1031321221123321-2101312223310021"></a>

## Direct properties — ethernet_interface / 031232331002 / 3

<a id="canonical-0212003213232120-0132300121133322-2122200100312302-1131101221113202-1101000132232303-2312003030220031-3321022011032010-3100312201023222"></a>

<a id="canonical-0202113021212022-1020111330320013-2113130120213213-1200030133233311-2131312032033231-2202223210331122-1122323313100113-3003323102103003"></a>

## device property — ethernet_interface / 031232331002 / 4

Type: `"string"`. Computed.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2123312312221011-1122111113010203-3030331003020201-0331022230033000-3031030321303111-2310031100031211-1033221002101232-2232032011320113"></a>

<a id="canonical-2030001011000233-0312103033030310-1000022221332003-0122101232111221-0201330231003120-1302122120020130-0023022230311212-2330333321001332"></a>

## mac property — ethernet_interface / 031232331002 / 5

Type: `"string"`. Computed.

MAC Address. Configuration parameter for mac

Upstream description:

Configuration parameter for mac

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2020213230311330-2211003130201130-3320203220132023-3031321322021033-3303111333112310-3311302332030133-3002222200031311-0100000021211232"></a>

## Next pages — ethernet_interface / 031232331002 / 6

- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011312222313303-1132322131031112-3203230121103102-1202111312023312-0001032213321132-1221202101300011-2320110222301231-2100333123303012"></a>

## nutanix.not_managed.node_list.interface_list.ipv6_auto_config — ipv6_auto_config / 133130111123 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-0023301133121330-2111033330310133-3123332321030300-2001212223301330-0211120332302200-2330123100201020-0333232202213212-2121212210031111"></a>

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

<a id="canonical-1100310031113331-1210320212132031-1001112022013110-0110310210212231-2033300021031120-3002211200231102-2023033133102110-3010210111122312"></a>

## Direct properties — ipv6_auto_config / 133130111123 / 3

- [host](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1132220130211201-2020203133223032-3013010011030101-0202110133313030-1331123321012201-3102222002330021-3310010010013021-3312200131030303): complete subsection reference.

- [router](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313): complete subsection reference.

<a id="canonical-2333122112120203-1203300212021210-3301302011003132-0230322030003330-0012211310323012-3222030212111012-0231100120301120-3111310233203221"></a>

## Next pages — ipv6_auto_config / 133130111123 / 4

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.host](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1132220130211201-2020203133223032-3013010011030101-0202110133313030-1331123321012201-3102222002330021-3310010010013021-3312200131030303)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1132220130211201-2020203133223032-3013010011030101-0202110133313030-1331123321012201-3102222002330021-3310010010013021-3312200131030303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330310123332312-3303130331013321-3122011322131320-2220133303311202-3010030103111233-3201330103001002-0102000300111310-0223130303012022"></a>

## nutanix.not_managed.node_list.interface_list.ipv6_auto_config.host — host / 201120311312 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-1311100130113202-3200310100311201-1212021033032012-2123310320210311-2122232320211320-2120120230120110-3100131211030212-3303311033203301"></a>

Type: `["object", {}]`. Computed.

Hostname or IP address of the target server.

Upstream description:

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

<a id="canonical-1112011300222032-1022103033030312-2020102320331213-1211210100232131-0330120113222012-3300302203121322-1002112230100323-3310303321120211"></a>

## Direct properties — host / 201120311312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323203133202121-3310232202313312-3320301100321013-2022111030103302-1213233333111330-0332211022130032-2313103230002010-1212333221101100"></a>

## Next pages — host / 201120311312 / 4

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032132213022033-0111303031230101-3320010121311200-1110111103221100-2030121200011011-3032002202212332-0223223211101310-2102210122010103"></a>

## nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router — router / 203230230230 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-2212211003120200-3020130332203221-1301000011011210-3031120033302321-2010030303110111-3103301100033321-0111312220220030-2300132123312203"></a>

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

<a id="canonical-3222103230100322-0212221110310122-2000221300212110-2223212223303333-1022122130301101-1102301322310301-3201203200022100-2113131312113031"></a>

## Direct properties — router / 203230230230 / 3

- [dns_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1213130023112112-1322111313201113-2300223210011010-2201321213303031-0202203233302330-2030020231012221-2031033001032320-2311221031113320): complete subsection reference.

<a id="canonical-3033011033021201-1221301211022302-3031321301131331-3103011011111130-1113221321120333-0210110103221302-0232111302133123-3120221102221232"></a>

<a id="canonical-0102333000300030-2103000100013120-1100031011223030-3330221232330112-1331103202300213-0220312211233321-0333032213323300-2301323002210032"></a>

## network_prefix property — router / 203230230230 / 4

Type: `"string"`. Computed.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223320102302000-1300223203011332-0231321113101233-3312101023113300-2122303231020130-3011122122302300-2213110302303330-2003210202202001): complete subsection reference.

<a id="canonical-2003321121002233-3211300100132020-3102210001121201-1220002103210200-2312323121103212-0021313222300230-0103312302102022-0221220210131021"></a>

## Next pages — router / 203230230230 / 5

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1213130023112112-1322111313201113-2300223210011010-2201321213303031-0202203233302330-2030020231012221-2031033001032320-2311221031113320)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](data-sources--securemesh_site_v2--reference--group-013.md#canonical-3223320102302000-1300223203011332-0231321113101233-3312101023113300-2122303231020130-3011122122302300-2213110302303330-2003210202202001)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1213130023112112-1322111313201113-2300223210011010-2201321213303031-0202203233302330-2030020231012221-2031033001032320-2311221031113320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033100201023011-1303303232101102-3121200211100001-2103330220112201-2323201022131103-1321121122313123-1203133122123322-1033321232101000"></a>

## nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — dns_config / 120232123201 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-1002313111132220-0031231321221333-0121002310301120-3002131230101323-2013203201211311-2121301102200023-3322213313010323-2330231033313312"></a>

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

<a id="canonical-3310013020031103-0032121120201212-1102232212331211-2121022202000101-0120011020122322-1102003313322123-2303113123003103-3222311121110221"></a>

## Direct properties — dns_config / 120232123201 / 3

- [configured_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1202330002220320-1011111322231001-0020103300001223-3103131131100213-3323201210321030-3121113112221032-1300202303012333-2013211101311130): complete subsection reference.

- [local_dns](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3233300003000202-0112323313331311-1032321001223000-3002113222333013-1300311001020332-1113323003301021-0212002021020013-1301303203003231): complete subsection reference.

<a id="canonical-2133222313220132-0203131203120312-2111323021033211-0103331333302030-1031023130210210-0211112221231220-2221131020032112-2002231212200031"></a>

## Next pages — dns_config / 120232123201 / 4

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1202330002220320-1011111322231001-0020103300001223-3103131131100213-3323201210321030-3121113112221032-1300202303012333-2013211101311130)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3233300003000202-0112323313331311-1032321001223000-3002113222333013-1300311001020332-1113323003301021-0212002021020013-1301303203003231)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1202330002220320-1011111322231001-0020103300001223-3103131131100213-3323201210321030-3121113112221032-1300202303012333-2013211101311130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202333313213300-3210310101102112-3323303021302120-1323111003310131-0133111222120300-2322123032212110-1120223211322021-2121110103312333"></a>

## nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — configured_list / 220013301001 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1213130023112112-1322111313201113-2300223210011010-2201321213303031-0202203233302330-2030020231012221-2031033001032320-2311221031113320)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-1010002132201033-0101031032313201-1330023031011233-3233011211002130-1310020332132203-2310223030021022-0013033002311323-1020302030032230"></a>

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

<a id="canonical-2101100112131003-2221020211102322-2133111101001033-1003212332300033-3220022121030000-1111130213313002-2113001021211220-0032333213313212"></a>

## Direct properties — configured_list / 220013301001 / 3

<a id="canonical-1223333001131003-1113000021103312-1312120311211200-1310310201022310-0021112321021123-2001130213100321-1200210332213222-2202300233003132"></a>

<a id="canonical-0201221101033022-3120210130122301-1111300310103130-3201003101223311-2011200313130232-1301210122121132-0201231031020131-3003031222220231"></a>

## dns_list property — configured_list / 220013301001 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0022302221303022-2330333320132203-2313311312300203-0111100001030003-0121332002221000-0101130320203232-2210110203312131-3231111302330202"></a>

## Next pages — configured_list / 220013301001 / 5

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1213130023112112-1322111313201113-2300223210011010-2201321213303031-0202203233302330-2030020231012221-2031033001032320-2311221031113320)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3233300003000202-0112323313331311-1032321001223000-3002113222333013-1300311001020332-1113323003301021-0212002021020013-1301303203003231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102203122302313-2232103001121131-2030223112232321-2210331133112221-1033331012023121-1203213233321321-0102110123311212-2132110301032102"></a>

## nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — local_dns / 020122302311 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1213130023112112-1322111313201113-2300223210011010-2201321213303031-0202203233302330-2030020231012221-2031033001032320-2311221031113320)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2321201302211132-0310233123203322-2031323122100301-0012100130103313-2100322203311030-2310331232023132-3010303031222310-2120031023303311"></a>

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

<a id="canonical-0103032020111031-3232310330300031-2100232303103031-3020200211311311-2121231023031302-1233212231300233-1312211321020030-2031320323213120"></a>

## Direct properties — local_dns / 020122302311 / 3

<a id="canonical-2211022103311320-1021330121130211-0123102101331200-2000010113133003-2231201332111132-1310123200120320-3332213302132032-2013320320212302"></a>

<a id="canonical-0000122302003220-0310331313220233-1330123311331133-3232002122232001-2010312213300122-1123210001013012-0220301112321032-2230221330332120"></a>

## configured_address property — local_dns / 020122302311 / 4

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [first_address](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1123102330302122-0131231032103300-0030202311300000-2113323020212023-3003002022033121-2200230123033301-2123003213030321-1001311102211110): complete subsection reference.

- [last_address](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3323103100033201-0333110223013113-2203322312231021-2223203223000202-2130233320213111-1312233022112121-1312303221223023-3033123013011001): complete subsection reference.

<a id="canonical-3203100210133313-1122200102331333-0002032120311113-1212102201030113-1013322310211300-0333201011002232-0112322210102031-3223121131333133"></a>

## Next pages — local_dns / 020122302311 / 5

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1123102330302122-0131231032103300-0030202311300000-2113323020212023-3003002022033121-2200230123033301-2123003213030321-1001311102211110)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3323103100033201-0333110223013113-2203322312231021-2223203223000202-2130233320213111-1312233022112121-1312303221223023-3033123013011001)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1213130023112112-1322111313201113-2300223210011010-2201321213303031-0202203233302330-2030020231012221-2031033001032320-2311221031113320)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-1123102330302122-0131231032103300-0030202311300000-2113323020212023-3003002022033121-2200230123033301-2123003213030321-1001311102211110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220030232100330-3211333301231031-0023101313313323-2320211101123100-3230202121003320-2101122232232010-2202230330320313-2212213210323232"></a>

## nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — first_address / 030223121210 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- [Property reference](data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112)
- [nutanix.not_managed](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0213322003100301-1120132022120021-2303032210122222-0231013130330331-3230122232113130-2322223020121212-0233003330102220-2320332312012210)
- [nutanix.not_managed.node_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2122121031022130-2200330331300033-3222013122122222-2321302013303122-3111031320232213-1300330021321023-2003200131311130-3032303120110213)
- [nutanix.not_managed.node_list.interface_list](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2013211322101122-3313112220212111-0020202000213011-1012330232020132-3012221202200012-3103032313011122-3212101112302321-2003200110112232)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1123222232002222-1322312123113002-1231213233022222-1121101221312313-3321102132110020-2332223020133213-2323211033332101-1210000111001310)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3020301231100020-3000122013302303-3132221331223112-3123120233033022-2033302003203030-3032333331323301-3322200320332033-3132130003110313)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1213130023112112-1322111313201113-2300223210011010-2201321213303031-0202203233302330-2030020231012221-2031033001032320-2311221031113320)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3233300003000202-0112323313331311-1032321001223000-3002113222333013-1300311001020332-1113323003301021-0212002021020013-1301303203003231)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-1311220312222333-0313221023311133-2130230102330232-0313010000312200-2133110321221203-0210330311010231-0222311230131230-1300233001303101"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

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

<a id="canonical-2103103200111033-1211221123203001-3002203113231113-3230312131333100-2131111302320230-2231112311230213-0121102122302030-3032210322331121"></a>

## Direct properties — first_address / 030223121210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111230102203011-0123122213221030-2301112022002031-1003110202012031-1202011132223102-0210102011022211-0233030311133323-2003001123212122"></a>

## Next pages — first_address / 030223121210 / 4

- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3233300003000202-0112323313331311-1032321001223000-3002113222333013-1300311001020332-1113323003301021-0212002021020013-1301303203003231)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)

<a id="canonical-3323103100033201-0333110223013113-2203322312231021-2223203223000202-2130233320213111-1312233022112121-1312303221223023-3033123013011001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
