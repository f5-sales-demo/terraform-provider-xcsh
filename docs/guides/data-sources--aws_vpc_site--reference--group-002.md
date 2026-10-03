---
page_title: "xcsh_aws_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site reference."
---

# xcsh_aws_vpc_site reference

<a id="canonical-3211110213100033-0311301221031331-2311313200310210-3122002320332302-3013132102322231-2002301101102100-2112032022310102-2321102210301012"></a>

## Next pages — coordinates / 032210200303 / 6

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0032303302311232-3330310111332131-0311311232001013-2021201310012332-0323301211013311-2100121202000201-3200002000310020-1011021120303300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012010000331132-1022211120102323-0013330202133312-0122303103132311-0032300220002012-0111230103110233-1023133132222130-0112000101210202"></a>

## custom_dns — custom_dns / 102223323002 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- custom_dns

<a id="canonical-2212002130231013-1033332212111033-0211201200312220-1313230133131012-1310232000132033-3102021023002113-1003101313002010-2333112133122201"></a>

Type: `"single"`. Computed.

Custom DNS is the configured for specify CE site.

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

<a id="canonical-1302300300100223-1232330230311021-1132221000100132-1202002233132100-3112133122111121-1301000023331021-1210231130022113-3323132310101332"></a>

## Direct properties — custom_dns / 102223323002 / 3

<a id="canonical-1323301221102022-3113121202012302-0131203010003133-3032213031210201-3312031331012112-1003001200121112-1333003220202123-0231021033232313"></a>

<a id="canonical-3101123222302203-1321112132200322-1203231212232110-0120223011131203-0211031330102331-3133013311330011-3221000333303220-3200330020113100"></a>

## inside_nameserver property — custom_dns / 102223323002 / 4

Type: `"string"`. Computed.

Optional DNS server IP to be used for name resolution in inside network.

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

<a id="canonical-1000011301312023-1021000211032110-1332020211102320-1132202130213210-2310213332110213-3131333321202302-0001331031313301-1233330330303201"></a>

<a id="canonical-1011102113333300-1010110002210213-3003000200021011-0023212320031322-3011123223022011-2133211320002110-3133300111312220-0033320123131300"></a>

## outside_nameserver property — custom_dns / 102223323002 / 5

Type: `"string"`. Computed.

Optional DNS server IP to be used for name resolution in outside network.

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

<a id="canonical-0230033133002103-2001323020001233-1301221203102222-2303210023032200-1320233000211001-3101222333310111-2031000133013012-1311012102130221"></a>

## Next pages — custom_dns / 102223323002 / 6

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2300130301211001-0301032331023031-0212213222101003-2022302012112121-2122001000222312-3212331321020021-3232021233203302-1103121031022303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113330100120110-0322121333013003-0110010331310303-3300230330110022-3120133113230322-3010032032331023-2232330301013002-3323030333210113"></a>

## custom_security_group — custom_security_group / 230223222022 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- custom_security_group

<a id="canonical-1121011102222323-3322023132122003-0013003300201310-1312120312020130-0333011110300202-1310003220203203-2200232313110112-3301122020333333"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_security\_group, f5xc\_security\_group\] Enter pre created security groups for
slo(Site Local Outside) and sli(Site Local Inside) interface. Supported only for sites deployed on
existing VPC.

Upstream description:

Enter pre created security groups for slo(Site Local Outside) and sli(Site Local Inside) interface.
Supported only for sites deployed on existing VPC.

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

OneOf alternatives in this subsection:

- [custom_security_group](data-sources--aws_vpc_site--reference--group-002.md#canonical-1121011102222323-3322023132122003-0013003300201310-1312120312020130-0333011110300202-1310003220203203-2200232313110112-3301122020333333)
- [f5xc_security_group](data-sources--aws_vpc_site--reference--group-002.md#canonical-3112003020000123-2331132210333332-0230132300013222-3330011302011032-0023332023301030-2103113111223203-1232023323213100-2223022203302023)

Select alternatives according to the provider validators above.

<a id="canonical-0310120110211030-0302310001232001-0310113313220032-3222133220010321-1131220032231332-0320131332210200-3031131233213023-0223303320132001"></a>

## Direct properties — custom_security_group / 230223222022 / 3

<a id="canonical-2033132312200213-3231302023100123-1333100301023333-1231202221033313-0330210203210303-2013221122032202-1200003323123133-3202330030210311"></a>

<a id="canonical-0010101001100120-2201012333022101-3132200321233310-0122020220111000-3122131120212200-3011130120210132-2200311200020133-2133230103233213"></a>

## inside_security_group_id property — custom_security_group / 230223222022 / 4

Type: `"string"`. Computed.

Security Group ID to be attached to SLI(Site Local Inside) Interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

<a id="canonical-0313010001212103-0220113123311231-0120223030332113-2002211031330301-1233302122103233-0220300022202103-2201232023311030-1001202001230133"></a>

<a id="canonical-0022332103221103-0032110223202213-1111033132313233-1123312133002210-0131313021132101-3321231133133310-2020331303110110-0012303103000101"></a>

## outside_security_group_id property — custom_security_group / 230223222022 / 5

Type: `"string"`. Computed.

Security Group ID to be attached to SLO(Site Local Outside) Interface.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

<a id="canonical-1130132103022300-1103301013122133-1013002323311010-3210112202113211-3310321331010032-1020023133020322-2033310133321211-1220101131102210"></a>

## Next pages — custom_security_group / 230223222022 / 6

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2200133112232132-2102313312112301-0123021223303231-3202333311322201-1101011011131001-3323121000102020-0123222002133021-3010113133031202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113031123000023-1120010011130020-3313200113022031-1022232301233200-0121000013300210-1012220130333212-3033320112230100-3221130231011222"></a>

## default_blocked_services — default_blocked_services / 110011213033 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- default_blocked_services

<a id="canonical-0112331121222233-1000312311001331-0210012000332221-2210212100220022-0023032332321003-2000320301231132-1122103102220310-1021000111332010"></a>

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

<a id="canonical-1301200121021121-2200123103132212-0020213133232313-2112201202022322-1311112123300203-2302320013021123-1021233131211110-3310023230013311"></a>

## Direct properties — default_blocked_services / 110011213033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1102122110222011-0210232202333312-3332300100113101-0223312221232111-3102100231200123-2301131202003332-1103200020232301-2302231033020311"></a>

## Next pages — default_blocked_services / 110011213033 / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2121300020212131-1223330033200121-2020110322021122-2022312120022130-1310122021203111-3320211121321222-1200012310123212-0130333023303301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001200103012222-0222131111311200-0323300033200301-1323331130113313-1330110000120033-2121032122120121-0332230332201010-1300102130202302"></a>

## direct_connect_disabled — direct_connect_disabled / 212131013102 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- direct_connect_disabled

<a id="canonical-0230120113020202-3002110201133130-0023100321311303-3031130012211013-1023101213112233-2001001103133321-0312300231323201-1202110111100110"></a>

Type: `["object", {}]`. Computed.

\[OneOf: direct\_connect\_disabled, direct\_connect\_enabled, private\_connectivity\] Enable this
option

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

OneOf alternatives in this subsection:

- [direct_connect_disabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0230120113020202-3002110201133130-0023100321311303-3031130012211013-1023101213112233-2001001103133321-0312300231323201-1202110111100110)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-1003222102201001-3322220302231232-3032320012233011-0310221002013003-0002232213300320-2003023331011120-0111223300102311-1211022322032201)
- [private_connectivity](data-sources--aws_vpc_site--reference--group-004.md#canonical-0210232211322203-1122311101121013-1311303122000210-1112330120000233-1130020213201030-0203112031031330-1233221031331220-1202301003333132)

Select alternatives according to the provider validators above.

<a id="canonical-3000002230311022-2120113210300030-1103100312222002-0100300020122121-0030111201313011-1223010021211201-2322000021110022-2333102331101111"></a>

## Direct properties — direct_connect_disabled / 212131013102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003103322333300-2003131321213312-2222222003031232-0120221122331200-1300011030101031-1231313001130200-1002220220202113-2202031313011220"></a>

## Next pages — direct_connect_disabled / 212131013102 / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0011130322113021-2121230232332103-1010131032001233-1221110131032331-1323320022300230-1002122312302231-2003312200331122-0302213210221232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332212123222200-3002020310112011-3032032012031333-3132023332013312-3000212100121230-0033120112113223-1230122031022113-1231202013101231"></a>

## direct_connect_enabled — direct_connect_enabled / 021033333213 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- direct_connect_enabled

<a id="canonical-1003222102201001-3322220302231232-3032320012233011-0310221002013003-0002232213300320-2003023331011120-0111223300102311-1211022322032201"></a>

Type: `"single"`. Computed.

Direct Connect Configuration. Direct Connect Configuration.

Upstream description:

Direct Connect Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"auto_asn\",\"custom_asn\"]",
  "x-ves-oneof-field-vif_choice": "[\"hosted_vifs\",\"standard_vifs\"]"
}
```

<a id="canonical-2010121223232122-0230013213321302-1110032122101002-2331030310103131-1003221000112003-2322230233011301-3201321133203301-0031010203201211"></a>

## Direct properties — direct_connect_enabled / 021033333213 / 3

- [auto_asn](data-sources--aws_vpc_site--reference--group-002.md#canonical-1322111103123102-2033323100031031-1213120300330332-0212023113001302-0020313011333020-2302031302111131-3220301232313330-0313020130033031): complete subsection reference.

<a id="canonical-2013130000322121-1002232032001220-0122100110331101-1212223232112230-3013133220313322-2032032331121201-1322131122323330-3121331103013210"></a>

<a id="canonical-1133123332002103-3130032322233022-2130312033220130-1033102010203112-2301233213033223-2213113131013321-3203323002323133-2201332031302321"></a>

## custom_asn property — direct_connect_enabled / 021033333213 / 4

Type: `"number"`. Computed.

Exclusive with \[auto\_asn\] Custom Autonomous System Number.

Upstream description:

Exclusive with \[auto\_asn\] Custom Autonomous System Number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-1302223123303100-2310001331133211-1102300200102110-3033231212202023-3213333330130101-0010030113203031-1201231232232231-0133121132021020): complete subsection reference.

- [standard_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-2001300213021032-0312003010312013-0323132323323203-3012221122302202-0210031311230200-0300231201011031-3030201223000223-1313123212100123): complete subsection reference.

<a id="canonical-0200111023012302-2101011030320001-0333321330202103-3001131330000301-1302101112132032-1212012131303021-0223230031032122-2313223202002333"></a>

## Next pages — direct_connect_enabled / 021033333213 / 5

- [direct_connect_enabled.auto_asn](data-sources--aws_vpc_site--reference--group-002.md#canonical-1322111103123102-2033323100031031-1213120300330332-0212023113001302-0020313011333020-2302031302111131-3220301232313330-0313020130033031)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-1302223123303100-2310001331133211-1102300200102110-3033231212202023-3213333330130101-0010030113203031-1201231232232231-0133121132021020)
- [direct_connect_enabled.standard_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-2001300213021032-0312003010312013-0323132323323203-3012221122302202-0210031311230200-0300231201011031-3030201223000223-1313123212100123)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1322111103123102-2033323100031031-1213120300330332-0212023113001302-0020313011333020-2302031302111131-3220301232313330-0313020130033031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023122030112121-0021302322121121-0022323121220201-1322111221310102-0303121331002002-2332122130112211-0203311331123230-2022121031122213"></a>

## direct_connect_enabled.auto_asn — auto_asn / 320203313203 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0011130322113021-2121230232332103-1010131032001233-1221110131032331-1323320022300230-1002122312302231-2003312200331122-0302213210221232)
- direct_connect_enabled.auto_asn

<a id="canonical-3333211130113111-0103322130220332-0033231113121232-2201112102123110-0110333213202223-2000102111210203-3211303031132133-0303111222123331"></a>

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

<a id="canonical-3232132010130311-3030003312020111-0313120222110302-1010221121022102-2010003010300020-3323222123202331-3133001023212131-3000302211230201"></a>

## Direct properties — auto_asn / 320203313203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023033212300231-3312022223102330-1223320221112322-3123022012333220-0032303320130223-2010032220310301-0130212000021231-1221220110122312"></a>

## Next pages — auto_asn / 320203313203 / 4

- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0011130322113021-2121230232332103-1010131032001233-1221110131032331-1323320022300230-1002122312302231-2003312200331122-0302213210221232)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1302223123303100-2310001331133211-1102300200102110-3033231212202023-3213333330130101-0010030113203031-1201231232232231-0133121132021020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100230131333003-3223022220210210-2320232001010212-3210231101332321-3230320302323132-1202023323102302-2213321322100222-2321201033302020"></a>

## direct_connect_enabled.hosted_vifs — hosted_vifs / 103003320031 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0011130322113021-2121230232332103-1010131032001233-1221110131032331-1323320022300230-1002122312302231-2003312200331122-0302213210221232)
- direct_connect_enabled.hosted_vifs

<a id="canonical-0323011012131010-1223322301222233-3122202013332330-2303232123133123-2012211232013012-0212323122213112-2301032330211110-1112131223011001"></a>

Type: `"single"`. Computed.

AWS Direct Connect Hosted VIF Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-connectivity_options": "[\"site_registration_over_direct_connect\",\"site_registration_over_internet\"]"
}
```

<a id="canonical-3301211023111211-3111230101121012-3010222300021203-0223131022330223-1000001023132121-3022332301131113-0113330103321331-3331023132011222"></a>

## Direct properties — hosted_vifs / 103003320031 / 3

- [site_registration_over_direct_connect](data-sources--aws_vpc_site--reference--group-002.md#canonical-3222130113121133-3010223211022330-0112330003031010-0113011110333321-1011301130011110-0030100011003030-2033321110122110-0211113202110103): complete subsection reference.

- [site_registration_over_internet](data-sources--aws_vpc_site--reference--group-002.md#canonical-2103210131022002-3202102030130101-0320330220210110-1120331313020332-3123202120333032-3120323001032321-3210001013201212-3023122303330322): complete subsection reference.

- [vif_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-3323301311202031-3123103113010221-3220031133200030-0321313230002130-0103133011302010-0232331020213133-2333021033332133-2103312010200102): complete subsection reference.

<a id="canonical-2101303013223201-0123110322032012-3210023031311003-2103121332332213-1302202113121102-1102013021130322-0211000303313023-0330212201112101"></a>

## Next pages — hosted_vifs / 103003320031 / 4

- [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](data-sources--aws_vpc_site--reference--group-002.md#canonical-3222130113121133-3010223211022330-0112330003031010-0113011110333321-1011301130011110-0030100011003030-2033321110122110-0211113202110103)
- [direct_connect_enabled.hosted_vifs.site_registration_over_internet](data-sources--aws_vpc_site--reference--group-002.md#canonical-2103210131022002-3202102030130101-0320330220210110-1120331313020332-3123202120333032-3120323001032321-3210001013201212-3023122303330322)
- [direct_connect_enabled.hosted_vifs.vif_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-3323301311202031-3123103113010221-3220031133200030-0321313230002130-0103133011302010-0232331020213133-2333021033332133-2103312010200102)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0011130322113021-2121230232332103-1010131032001233-1221110131032331-1323320022300230-1002122312302231-2003312200331122-0302213210221232)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3222130113121133-3010223211022330-0112330003031010-0113011110333321-1011301130011110-0030100011003030-2033321110122110-0211113202110103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222320202002312-2031021231232231-0103110330101302-3002300011233122-2321112002303303-1302112030222223-2231203101111202-1321112031030033"></a>

## direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect — site_registration_over_direct_connect / 201122313322 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0011130322113021-2121230232332103-1010131032001233-1221110131032331-1323320022300230-1002122312302231-2003312200331122-0302213210221232)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-1302223123303100-2310001331133211-1102300200102110-3033231212202023-3213333330130101-0010030113203031-1201231232232231-0133121132021020)
- direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect

<a id="canonical-2320203000313101-0200211132221131-3200131330210023-2200110011022312-0303230020030112-3112223320023132-2220131312232301-0312000122303001"></a>

Type: `"single"`. Computed.

CloudLink ADN Network Config.

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

<a id="canonical-2101122320202200-2330310100032312-0211311300103101-1102303102313033-3011001333011031-3121111303112312-0221120203321020-0120000122000302"></a>

## Direct properties — site_registration_over_direct_connect / 201122313322 / 3

<a id="canonical-0230211131223020-0320022110322101-3021213030001213-2302333002112021-1223133331211032-1120313213223303-3022031022301311-1323303100122131"></a>

<a id="canonical-1210320120103232-2003302021310012-0131012100211013-3031011331301311-0102223113130313-2000130123321221-1111003131131003-3230112021121301"></a>

## cloudlink_network_name property — site_registration_over_direct_connect / 201122313322 / 4

Type: `"string"`. Computed.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1111232130230302-0003330000223130-0013321021021022-2202222322112310-0111203300333122-0001233003302113-2203320131113010-2301203213103110"></a>

## Next pages — site_registration_over_direct_connect / 201122313322 / 5

- [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-1302223123303100-2310001331133211-1102300200102110-3033231212202023-3213333330130101-0010030113203031-1201231232232231-0133121132021020)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2103210131022002-3202102030130101-0320330220210110-1120331313020332-3123202120333032-3120323001032321-3210001013201212-3023122303330322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213311022002310-3032023323203232-1302332131203301-3332111310131021-0203123030123223-0302021303111210-3322101230101220-3002032303312233"></a>

## direct_connect_enabled.hosted_vifs.site_registration_over_internet — site_registration_over_internet / 212130130311 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0011130322113021-2121230232332103-1010131032001233-1221110131032331-1323320022300230-1002122312302231-2003312200331122-0302213210221232)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-1302223123303100-2310001331133211-1102300200102110-3033231212202023-3213333330130101-0010030113203031-1201231232232231-0133121132021020)
- direct_connect_enabled.hosted_vifs.site_registration_over_internet

<a id="canonical-2212133331012312-2123311222310223-2000311103211300-1131010301103313-0032202233131010-2322110101110232-2110213133123120-1331300030300133"></a>

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

<a id="canonical-2311000200223112-2022323023231122-0323113301302332-0211002132312013-0123130011312320-0111332300203100-3031111322121131-3021301303132202"></a>

## Direct properties — site_registration_over_internet / 212130130311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223023230230322-3122203023030310-2101211113231012-0213322030021210-3223232211000223-0130113132320301-3122312002031113-3011332210232030"></a>

## Next pages — site_registration_over_internet / 212130130311 / 4

- [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-1302223123303100-2310001331133211-1102300200102110-3033231212202023-3213333330130101-0010030113203031-1201231232232231-0133121132021020)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3323301311202031-3123103113010221-3220031133200030-0321313230002130-0103133011302010-0232331020213133-2333021033332133-2103312010200102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000020301311100-1303023103310002-1112211322233223-3223303330001132-3123310201031202-1131211222122023-2320223021110300-1321132022303213"></a>

## direct_connect_enabled.hosted_vifs.vif_list — vif_list / 010110220100 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0011130322113021-2121230232332103-1010131032001233-1221110131032331-1323320022300230-1002122312302231-2003312200331122-0302213210221232)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-1302223123303100-2310001331133211-1102300200102110-3033231212202023-3213333330130101-0010030113203031-1201231232232231-0133121132021020)
- direct_connect_enabled.hosted_vifs.vif_list

<a id="canonical-2001132101300200-3303210330121030-0012131003200201-3100130220120203-2300221223102022-1130211233221121-2323233201202101-3121102203321200"></a>

Type: `"list"`. Computed.

List of Hosted VIF Config. List of Hosted VIF Config.

Upstream description:

List of Hosted VIF Config.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 30,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 30,
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
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0333032231332033-3111201311102211-3030200212001231-3310330310112121-3203201232000130-0313132213211331-0032122222211031-2223313222301320"></a>

## Direct properties — vif_list / 010110220100 / 3

<a id="canonical-3102320100102231-0101003132311131-3033030021001211-1003301200130210-3020233310123322-0033020313012210-2222031110231102-3200210030321020"></a>

<a id="canonical-0301100102210132-0221221200230223-2220132113032330-2213020231031030-3300321212320212-0000133010333322-3001100302101333-0330010103000303"></a>

## other_region property — vif_list / 010110220100 / 4

Type: `"string"`. Computed.

\[Enum:
af-south-1|ap-east-1|ap-northeast-1|ap-northeast-2|ap-south-1|ap-southeast-1|ap-southeast-2|ap-southeast-3|ca-central-1|eu-central-1|eu-north-1|eu-south-1|eu-west-1|eu-west-2|eu-west-3|me-south-1|sa-east-1|us-east-1|us-east-2|us-west-1|us-west-2\]
Exclusive with \[same\_as\_site\_region\] Other Region. Possible values are \`af-south-1\`,
\`ap-east-1\`, \`ap-northeast-1\`, \`ap-northeast-2\`, \`ap-south-1\`, \`ap-southeast-1\`,
\`ap-southeast-2\`, \`ap-southeast-3\`, \`ca-central-1\`, \`eu-central-1\`, \`eu-north-1\`,
\`eu-south-1\`, \`eu-west-1\`, \`eu-west-2\`, \`eu-west-3\`, \`me-south-1\`, \`sa-east-1\`,
\`us-east-1\`, \`us-east-2\`, \`us-west-1\`, \`us-west-2\`.

Upstream description:

Exclusive with \[same\_as\_site\_region\] Other Region.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "af-south-1",
    "ap-east-1",
    "ap-northeast-1",
    "ap-northeast-2",
    "ap-south-1",
    "ap-southeast-1",
    "ap-southeast-2",
    "ap-southeast-3",
    "ca-central-1",
    "eu-central-1",
    "eu-north-1",
    "eu-south-1",
    "eu-west-1",
    "eu-west-2",
    "eu-west-3",
    "me-south-1",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-1",
    "us-west-2"
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
    "ves.io.schema.rules.string.in": "[\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-northeast-1\\\",\\\"ap-northeast-2\\\",\\\"ap-south-1\\\",\\\"ap-southeast-1\\\",\\\"ap-southeast-2\\\",\\\"ap-southeast-3\\\",\\\"ca-central-1\\\",\\\"eu-central-1\\\",\\\"eu-north-1\\\",\\\"eu-south-1\\\",\\\"eu-west-1\\\",\\\"eu-west-2\\\",\\\"eu-west-3\\\",\\\"me-south-1\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-1\\\",\\\"us-west-2\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-northeast-1\\\",\\\"ap-northeast-2\\\",\\\"ap-south-1\\\",\\\"ap-southeast-1\\\",\\\"ap-southeast-2\\\",\\\"ap-southeast-3\\\",\\\"ca-central-1\\\",\\\"eu-central-1\\\",\\\"eu-north-1\\\",\\\"eu-south-1\\\",\\\"eu-west-1\\\",\\\"eu-west-2\\\",\\\"eu-west-3\\\",\\\"me-south-1\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-1\\\",\\\"us-west-2\\\"]"
  }
}
```

- [same_as_site_region](data-sources--aws_vpc_site--reference--group-002.md#canonical-3132203102000102-3221112222113013-2021100000210131-0201102220003000-2320021213102333-0113301303010032-0233130130222331-0033110332102201): complete subsection reference.

<a id="canonical-0131033003100310-0221112230001000-2020001131312211-3033130102132132-1332303322103311-0211030301120200-3332313233100300-2312331130100113"></a>

<a id="canonical-1303221210230211-1301022103120023-1322103232022001-1301312033111002-3013131220303301-2332021121311212-1022122200201300-2133231232100223"></a>

## vif_id property — vif_list / 010110220100 / 5

Type: `"string"`. Computed.

AWS Direct Connect VIF ID that needs to be connected to the site.

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
    "pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-0222320130333220-0102311030220212-2012232133113301-3323222200023212-0203323201213331-1331121121110013-1001000003311011-3222231200030020"></a>

## Next pages — vif_list / 010110220100 / 6

- [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](data-sources--aws_vpc_site--reference--group-002.md#canonical-3132203102000102-3221112222113013-2021100000210131-0201102220003000-2320021213102333-0113301303010032-0233130130222331-0033110332102201)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-1302223123303100-2310001331133211-1102300200102110-3033231212202023-3213333330130101-0010030113203031-1201231232232231-0133121132021020)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3132203102000102-3221112222113013-2021100000210131-0201102220003000-2320021213102333-0113301303010032-0233130130222331-0033110332102201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303203301123332-0013232320301310-3230023333221021-2320223111112210-0033330202103022-3033311100022212-1113032121301033-3101021030131122"></a>

## direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region — same_as_site_region / 130210013001 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0011130322113021-2121230232332103-1010131032001233-1221110131032331-1323320022300230-1002122312302231-2003312200331122-0302213210221232)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_vpc_site--reference--group-002.md#canonical-1302223123303100-2310001331133211-1102300200102110-3033231212202023-3213333330130101-0010030113203031-1201231232232231-0133121132021020)
- [direct_connect_enabled.hosted_vifs.vif_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-3323301311202031-3123103113010221-3220031133200030-0321313230002130-0103133011302010-0232331020213133-2333021033332133-2103312010200102)
- direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region

<a id="canonical-3312331110203111-3222322103313213-1022221302030022-3230011011030221-0311031122002301-0123213102023022-3101230300003332-3101033112201011"></a>

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

<a id="canonical-2222211200001213-3303003003323301-1032133002321333-1133003220230101-1213332312120023-0222323311311221-0332210221301020-0033311111212200"></a>

## Direct properties — same_as_site_region / 130210013001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311021031333032-0221213211002321-3322200331232031-0102230223010320-0201002210211201-0221032301001220-3333220333213221-3223323032313332"></a>

## Next pages — same_as_site_region / 130210013001 / 4

- [direct_connect_enabled.hosted_vifs.vif_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-3323301311202031-3123103113010221-3220031133200030-0321313230002130-0103133011302010-0232331020213133-2333021033332133-2103312010200102)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2001300213021032-0312003010312013-0323132323323203-3012221122302202-0210031311230200-0300231201011031-3030201223000223-1313123212100123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331213023233030-2330212232220213-2201020301310002-2011011200213223-0101321300012212-1300032033000312-2231200333101231-3202233332123220"></a>

## direct_connect_enabled.standard_vifs — standard_vifs / 032123210302 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0011130322113021-2121230232332103-1010131032001233-1221110131032331-1323320022300230-1002122312302231-2003312200331122-0302213210221232)
- direct_connect_enabled.standard_vifs

<a id="canonical-0101123103332032-1010000310331301-1121223310112020-0010032031333311-1233102000103032-0013030032211231-3303303020310213-2200003103110011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for standard vifs.

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

<a id="canonical-0202220021323033-2102132023201021-1022300310202022-0212012300210201-0211003211300303-1132020333112033-3233210323110312-1000111002301102"></a>

## Direct properties — standard_vifs / 032123210302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102123231200110-2033202131002223-3132222322022313-1302321222320320-2120003030220000-1000302121111320-0333021123002122-0212322111301230"></a>

## Next pages — standard_vifs / 032123210302 / 4

- [direct_connect_enabled](data-sources--aws_vpc_site--reference--group-002.md#canonical-0011130322113021-2121230232332103-1010131032001233-1221110131032331-1323320022300230-1002122312302231-2003312200331122-0302213210221232)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1311011230011120-1331003133110322-3120131230223200-1013132302300320-0320113132301031-3101000301101123-3122110020101101-1231233022030212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331001233001131-0023112120111202-2310110310030000-2330020313000001-3113032123031332-3223012302030130-1230112021202213-2333302333232223"></a>

## disable_encryption — disable_encryption / 112001121301 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- disable_encryption

<a id="canonical-3333332333310302-3203122103030103-1010121230131302-2021312232332001-0122100332023223-0302120122122002-2003012222033202-2301121101301303"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_encryption, enable\_encryption; Default: disable\_encryption\] Configuration
parameter for disable encryption.

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

OneOf alternatives in this subsection:

- [disable_encryption](data-sources--aws_vpc_site--reference--group-002.md#canonical-3333332333310302-3203122103030103-1010121230131302-2021312232332001-0122100332023223-0302120122122002-2003012222033202-2301121101301303)
- [enable_encryption](data-sources--aws_vpc_site--reference--group-002.md#canonical-3301022002013232-2020123112233212-0232010322323330-0023202033111233-1220132121033331-2303302023323102-0031120112211312-3203021002103132)

Select alternatives according to the provider validators above.

<a id="canonical-0202200231123112-2213201331303113-1231312321223031-1122301223222100-1330323213201002-1211311212313022-0112020121300012-2201333323120122"></a>

## Direct properties — disable_encryption / 112001121301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113222322003333-3110111002103311-2011111010223230-0032200123002002-2003320101122233-1222212010022231-3120111202232032-1231320101322012"></a>

## Next pages — disable_encryption / 112001121301 / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0133220213033001-3001221013333231-2010200113011123-1220313332323222-1112031312223002-3331002110130023-1320101232000122-3311101023223213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222323210331233-3312331023102313-2323301103202103-3020333012201210-3110021012323321-2033323321003121-0313313121002302-0030103321220033"></a>

## disable_internet_vip — disable_internet_vip / 131322313102 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- disable_internet_vip

<a id="canonical-0021030323003211-3121302001232032-0230232102301131-2001112232012020-1132302110233312-3102032112301200-3131320213100232-0020132312211121"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_internet\_vip, enable\_internet\_vip; Default: disable\_internet\_vip\] Enable
this option

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

OneOf alternatives in this subsection:

- [disable_internet_vip](data-sources--aws_vpc_site--reference--group-002.md#canonical-0021030323003211-3121302001232032-0230232102301131-2001112232012020-1132302110233312-3102032112301200-3131320213100232-0020132312211121)
- [enable_internet_vip](data-sources--aws_vpc_site--reference--group-002.md#canonical-3221323001223200-0012211123003113-2321022213222002-1233322021220122-3111231100211013-3113330211333100-2201020333330213-3312100003311200)

Select alternatives according to the provider validators above.

<a id="canonical-0133222030332221-0330123320313031-0033323330121211-2021131201331321-1030312200003223-2210103323332230-3003303321013322-1231320020312111"></a>

## Direct properties — disable_internet_vip / 131322313102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323123022012013-3333023313210001-0202231200222202-3101010232130100-3301223030030321-2013101101232123-3112023320103332-1210001003212022"></a>

## Next pages — disable_internet_vip / 131322313102 / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0213323122120020-1331223112003223-3322230223312131-0120033112033110-3331033111333311-3020001110302122-2313012032030122-3203230102103211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032120111003133-0010232021123000-0030112303011300-1320011221322023-0333212203013203-2120303112222003-3021201232221333-3221020230302201"></a>

## egress_gateway_default — egress_gateway_default / 120030112003 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- egress_gateway_default

<a id="canonical-0010223221310313-3303122223302011-0022110130011022-2211001211220102-3130010221213111-0133321112302300-0301210232221023-3211311222220003"></a>

Type: `["object", {}]`. Computed.

\[OneOf: egress\_gateway\_default, egress\_nat\_gw, egress\_virtual\_private\_gateway; Default:
egress\_gateway\_default\] Configuration parameter for egress gateway default.

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

OneOf alternatives in this subsection:

- [egress_gateway_default](data-sources--aws_vpc_site--reference--group-002.md#canonical-0010223221310313-3303122223302011-0022110130011022-2211001211220102-3130010221213111-0133321112302300-0301210232221023-3211311222220003)
- [egress_nat_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-3321232120013301-0202012303233001-1022000032100312-3231303133121103-0122112200232331-2013110113103211-3023000112301310-3213310223020211)
- [egress_virtual_private_gateway](data-sources--aws_vpc_site--reference--group-002.md#canonical-1330021311030311-0203012201110310-2321200310111012-1002130132111022-0300323113020320-0133203201113022-3222332331020013-3311133130202322)

Select alternatives according to the provider validators above.

<a id="canonical-3103110230111010-2103112331203013-1330231033031202-0000301120332001-3120131101100131-0030122331301322-2121110202002132-2110031101110301"></a>

## Direct properties — egress_gateway_default / 120030112003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210232331032211-0333033020121110-2133030300011022-3231130222233333-1111203110301202-3211030230000333-1101133333012310-2112131300133300"></a>

## Next pages — egress_gateway_default / 120030112003 / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2132102231211003-1132002300321213-0112301333121331-0112133330321102-0012020001310321-3330131223000122-2323012013130320-1132301112001223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000220203033003-0312030222300202-1120021330110010-3000302003000202-3221102310030300-0122233001232322-3202210200110010-3232312333301022"></a>

## egress_nat_gw — egress_nat_gw / 003020030001 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- egress_nat_gw

<a id="canonical-3321232120013301-0202012303233001-1022000032100312-3231303133121103-0122112200232331-2013110113103211-3023000112301310-3213310223020211"></a>

Type: `"single"`. Computed.

With this option, egress site traffic will be routed through an Network Address Translation(NAT)
Gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"nat_gw_id\"]"
}
```

<a id="canonical-3121110013111211-1023113310303103-0011320001112032-1323001022310033-0303111303231232-1011312322103130-1311011203112322-1031123222020112"></a>

## Direct properties — egress_nat_gw / 003020030001 / 3

<a id="canonical-2320223132031100-1102302010030101-2100001101133322-1220113203322332-3332331001232200-3321130320131001-1232323232200332-0023311211331000"></a>

<a id="canonical-1010323233231112-3323121010012101-3013000311131003-0100311300013312-2210330232120123-3112230330220220-1012210210230220-1202010103021302"></a>

## nat_gw_id property — egress_nat_gw / 003020030001 / 4

Type: `"string"`. Computed.

Existing NAT Gateway ID. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 21,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^(nat-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "21",
    "ves.io.schema.rules.string.pattern": "^(nat-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "21",
    "ves.io.schema.rules.string.pattern": "^(nat-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-2311112112301032-2103123010211230-3200312233102122-0320330023230103-1320210301313312-2201221310002320-2320231301133203-1103213300032030"></a>

## Next pages — egress_nat_gw / 003020030001 / 5

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1000201021001020-0231311230020321-2302002113210031-1333233021033222-2311232002212123-1223300310301200-2023111102303320-1313212210223013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030033121032331-1032320313101011-1310220132002213-2023202113030100-0330212121301203-3123123300221000-3123320221102003-1031031300321121"></a>

## egress_virtual_private_gateway — egress_virtual_private_gateway / 202322100102 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- egress_virtual_private_gateway

<a id="canonical-1330021311030311-0203012201110310-2321200310111012-1002130132111022-0300323113020320-0133203201113022-3222332331020013-3311133130202322"></a>

Type: `"single"`. Computed.

With this option, egress site traffic will be routed through an Virtual Private Gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"vgw_id\"]"
}
```

<a id="canonical-2001012132120131-1311232132221131-0101232323310331-1222222330003111-2312233122322222-2121321332003001-0002032111020110-0003120122113333"></a>

## Direct properties — egress_virtual_private_gateway / 202322100102 / 3

<a id="canonical-1101120211201321-2003123131131103-3031321210211103-1021303220212032-0333330030023103-1021333222113033-0110303132001320-2210232123223322"></a>

<a id="canonical-0302331222333330-2032212003232301-0313032330022031-3110121231313212-0300010231220333-2013200302113133-0121232232321323-3203300010311012"></a>

## vgw_id property — egress_virtual_private_gateway / 202322100102 / 4

Type: `"string"`. Computed.

Existing Virtual Private Gateway ID. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 21,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^(vgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "21",
    "ves.io.schema.rules.string.pattern": "^(vgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "21",
    "ves.io.schema.rules.string.pattern": "^(vgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-1123122333210011-3202122131111001-1302202203120003-3232132021220103-3332223210300301-2233312001303131-0310030131331232-1111033213222120"></a>

## Next pages — egress_virtual_private_gateway / 202322100102 / 5

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3223303221330113-3211103231323113-3311331023113321-2023323121233120-2330000301131101-3121201332001302-2000032020123010-2333221101313323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301132330032322-2301220133312123-2111231100330112-1322320031030131-0020222200201203-0020021313322321-1312210300230311-0230311310211322"></a>

## enable_encryption — enable_encryption / 322033021203 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- enable_encryption

<a id="canonical-3301022002013232-2020123112233212-0232010322323330-0023202033111233-1220132121033331-2303302023323102-0031120112211312-3203021002103132"></a>

Type: `"single"`. Computed.

Configuration parameter for enable encryption.

Upstream description:

Information related to disk encryption.

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

<a id="canonical-3032200123300021-3013301202002123-1113111030312222-1212223213203102-1203313332033222-3020131321322012-2330110221220313-1021202332010122"></a>

## Direct properties — enable_encryption / 322033021203 / 3

<a id="canonical-1232230023300303-1321123331233301-1203212203030103-1312102220310100-1211212230323200-1322003310203303-1013000233311231-1331231021010031"></a>

<a id="canonical-2331020333030332-2012210122112133-2203311212220132-0002032302231121-2332201012200113-1231321011132231-0112232201000211-0300022203111002"></a>

## kms_key_id property — enable_encryption / 322033021203 / 4

Type: `"string"`. Computed.

AWS KMS Key to be used to encrypt the disk attached to the VM.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0331303202030322-0230132111322030-1300211230033301-1210101232032123-1330303113011313-3302133212000010-3001030313331302-1133201110211202"></a>

## Next pages — enable_encryption / 322033021203 / 5

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2302322100030210-1000301001133201-0123332222102130-1000222233020000-1321301321221202-1210323123023112-0333010121111021-2003313232323022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211002120121210-2210020320011300-3211022322321132-3330030002200321-1303110101100220-2011230003011302-1332030123232232-2110231133311332"></a>

## enable_internet_vip — enable_internet_vip / 223011212111 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- enable_internet_vip

<a id="canonical-3221323001223200-0012211123003113-2321022213222002-1233322021220122-3111231100211013-3113330211333100-2201020333330213-3312100003311200"></a>

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

<a id="canonical-1111100131212103-2031033323233033-0200211302033232-1301222113130000-3202023123333103-0103301201002022-2332233323023302-0100313212130231"></a>

## Direct properties — enable_internet_vip / 223011212111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201220133103000-0320000023322122-1313023332100203-0130011130300321-2132321310203131-3101101303030122-1122200200223300-1110113020330200"></a>

## Next pages — enable_internet_vip / 223011212111 / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3030120030303220-2301110223320001-1122111030010130-1111021302102103-1002211331003112-2202112322210313-2221121212123311-2201321223220112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002011301012331-2232003302201202-1131032203222212-3103130303000213-1311033323122133-2223323132130213-2211103223222130-0111331123022323"></a>

## f5_orchestrated_routing — f5_orchestrated_routing / 300301231230 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- f5_orchestrated_routing

<a id="canonical-1302133212120110-3332133210022111-1322211233133221-0032300011331303-3200101323211310-3233200321212100-0221021012213130-2013223021332011"></a>

Type: `["object", {}]`. Computed.

\[OneOf: f5\_orchestrated\_routing, manual\_routing\] Enable this option

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

OneOf alternatives in this subsection:

- [f5_orchestrated_routing](data-sources--aws_vpc_site--reference--group-002.md#canonical-1302133212120110-3332133210022111-1322211233133221-0032300011331303-3200101323211310-3233200321212100-0221021012213130-2013223021332011)
- [manual_routing](data-sources--aws_vpc_site--reference--group-004.md#canonical-0000121213033322-0302020313201122-2333221311101301-1103013210212003-0023002322232130-1211330203023202-1000313013333112-0200002200213232)

Select alternatives according to the provider validators above.

<a id="canonical-0013021123232122-2111331131333303-1123021323113203-1200130030322231-0200131113012012-3312312200321020-0101212221103321-3210331032303312"></a>

## Direct properties — f5_orchestrated_routing / 300301231230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122020230223101-2222031011101312-0101311233332101-1103023222301331-3101310031202330-1103211112122123-2303203311000303-3332123022231302"></a>

## Next pages — f5_orchestrated_routing / 300301231230 / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3211023102020300-2302030022120103-1220113121003003-2133023230212130-3221203030310131-2221131101321200-1022323100121112-3321122111312302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132111111222213-3023220002312011-3203021022110303-1203021011023133-2330013100100000-0110121332202031-3212000120031002-0030103122201123"></a>

## f5xc_security_group — f5xc_security_group / 321312223232 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- f5xc_security_group

<a id="canonical-3112003020000123-2331132210333332-0230132300013222-3330011302011032-0023332023301030-2103113111223203-1232023323213100-2223022203302023"></a>

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

<a id="canonical-3320012012003010-2221001300120121-2320110002231231-0233212101210211-2221223100122021-3303031231020100-2222130002333300-2120002331200333"></a>

## Direct properties — f5xc_security_group / 321312223232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123302011203010-1310131203330200-2323022033121313-1101101113211033-0230212011030132-2120210101301111-2102133011111302-0223231003113323"></a>

## Next pages — f5xc_security_group / 321312223232 / 4

- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132310203110133-3312110333011013-3123233202030220-0331123012012130-0233203110001200-1301323202310303-1203302311011322-1210202201223002"></a>

## ingress_egress_gw — ingress_egress_gw / 232200222320 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- ingress_egress_gw

<a id="canonical-2011112031122032-2010321010101232-0202223030133130-2013300013202300-1102313232023013-2212030132231233-3232123330120231-0303021133200033"></a>

Type: `"single"`. Computed.

\[OneOf: ingress\_egress\_gw, ingress\_gw, voltstack\_cluster\] Configuration parameter for ingress
egress gw.

Upstream description:

Two interface AWS ingress/egress site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group_inside_vn\",\"dc_cluster_group_outside_vn\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

OneOf alternatives in this subsection:

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2011112031122032-2010321010101232-0202223030133130-2013300013202300-1102313232023013-2212030132231233-3232123330120231-0303021133200033)
- [ingress_gw](data-sources--aws_vpc_site--reference--group-003.md#canonical-0211201331003211-2132210332011032-1003100231033223-3330301033201113-3213013312201031-1131230330130022-1313121113320212-3103122022210202)
- [voltstack_cluster](data-sources--aws_vpc_site--reference--group-004.md#canonical-2012232113013000-2100230113230223-1220200212223012-1300311102110030-1332220233303002-1112200023110211-2211321033232213-3123211213032200)

Select alternatives according to the provider validators above.

<a id="canonical-1201300301310311-0000300022103113-1213022100212021-2120122232013133-0323100103101030-2002032030022113-0313320000033320-3102212302010013"></a>

## Direct properties — ingress_egress_gw / 232200222320 / 3

- [active_enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-0332002031330310-1001311113102012-2213302223032332-3103322312230021-0130020120000332-0303301210030231-0100313220202201-0121201020100231): complete subsection reference.

- [active_forward_proxy_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-2321211103101121-0011230123322020-0302113111223322-0023231302301003-3332213010313020-2311212131220121-2111113320233013-1323031001221320): complete subsection reference.

- [active_network_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-1130103001330020-0120302311213122-1301011022001230-1121210233020211-0032012322223313-3021003133203203-2111011010322003-1113322113010010): complete subsection reference.

- [allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0213222032201133-3023322220200332-2310211023202102-3033100212021031-1023303020033010-2210310010012322-1111231210333030-2223313100200132): complete subsection reference.

- [allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-3221023301133323-2200330310333212-3321313022032001-2211303203212302-2021212020320233-3202011003112100-3320303230112020-3303311302030301): complete subsection reference.

<a id="canonical-3202220032021113-3121113100123331-2310320212022133-1122203113313333-2230111310013310-2300003313231211-1211211110001223-2101023133033021"></a>

<a id="canonical-3111022220111333-3103311213133011-2231300131313011-3032201232322030-2001331311311111-0231011232220223-1012103021112033-3203003330212030"></a>

## aws_certified_hw property — ingress_egress_gw / 232200222320 / 4

Type: `"string"`. Computed.

\[Enum: aws-byol-multi-nic-voltmesh\] AWS Certified Hardware. Name for AWS certified hardware. The
only possible value is \`aws-byol-multi-nic-voltmesh\`.

Upstream description:

Name for AWS certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "aws-byol-multi-nic-voltmesh"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-2100303003003233-1303000303212100-1023000322330013-3110102322202201-2211030011311102-0022031123213330-1023311213221311-2211103221220212): complete subsection reference.

- [dc_cluster_group_inside_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-0222122011023220-0100000202223003-0111101023021010-1333230300022331-0311230122321120-2320122200100210-1000223213213100-3323312031332330): complete subsection reference.

- [dc_cluster_group_outside_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-0020231301003320-3023012122320023-0221232300021022-3032213000300030-2213132310000101-1033331132210202-2000333030130223-3001323103001320): complete subsection reference.

- [forward_proxy_allow_all](data-sources--aws_vpc_site--reference--group-002.md#canonical-1002320332003012-1322332301331332-1321130113113230-0013121301031323-2313001313013132-2130103332211001-1002122032221131-0201233333331021): complete subsection reference.

- [global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-1332122323112132-3002221202203112-1112223122230110-3202203020123013-1011121222220200-2331333002313323-3000011132300012-2020031001021321): complete subsection reference.

- [inside_static_routes](data-sources--aws_vpc_site--reference--group-002.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000): complete subsection reference.

- [no_dc_cluster_group](data-sources--aws_vpc_site--reference--group-003.md#canonical-2113321200233231-3011122321321213-2322321230202210-0323003112201000-0212212222101310-3103103013032210-3111111010332102-2202300302033122): complete subsection reference.

- [no_forward_proxy](data-sources--aws_vpc_site--reference--group-003.md#canonical-2312012003101212-1030213320311213-0123230122232202-1310002303110300-2311102011011110-0312001221202332-0122021121031220-1230121022010103): complete subsection reference.

- [no_global_network](data-sources--aws_vpc_site--reference--group-003.md#canonical-3012322203311023-1122021233203330-0131320222010020-1011203310102133-1133311130333020-0131031232013230-2110210300100131-2131013320322233): complete subsection reference.

- [no_inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122332310030122-2301203331321331-3312021331313032-0110313021220211-3000212223310100-3103123301113003-2201001303021001-3310113103222230): complete subsection reference.

- [no_network_policy](data-sources--aws_vpc_site--reference--group-003.md#canonical-2322012331311110-3013210130000330-0300331233210121-3001320133210330-3030230023030101-3300221022323202-0233223101003130-3122210101130021): complete subsection reference.

- [no_outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0122030330020231-1020100201333223-2023110003022323-2222101300102211-3213133230113230-1320232212211222-1130122001323210-0110110010011320): complete subsection reference.

- [outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312): complete subsection reference.

- [performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-003.md#canonical-3333111201203120-1223122133331331-2131312003210200-0303001232012121-2122230310310231-0010001201223202-3313133220111030-3032010201000230): complete subsection reference.

- [sm_connection_public_ip](data-sources--aws_vpc_site--reference--group-003.md#canonical-1032302033302121-2121112031020130-2120010231320100-1213021300031120-0122002321131212-2311120210210313-3103011330111033-2012123323331010): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--aws_vpc_site--reference--group-003.md#canonical-1031100033200031-2003100101032231-1010002231031223-1103230320030133-0110003322321200-3003223200303311-1023021301231301-1232111011220221): complete subsection reference.

<a id="canonical-2031021012012012-3221231121110301-3303203112321033-0110222231330110-1322032321013302-1202322020313313-0311310132123211-3230110100221122"></a>

## Next pages — ingress_egress_gw / 232200222320 / 5

- [ingress_egress_gw.active_enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-0332002031330310-1001311113102012-2213302223032332-3103322312230021-0130020120000332-0303301210030231-0100313220202201-0121201020100231)
- [ingress_egress_gw.active_forward_proxy_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-2321211103101121-0011230123322020-0302113111223322-0023231302301003-3332213010313020-2311212131220121-2111113320233013-1323031001221320)
- [ingress_egress_gw.active_network_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-1130103001330020-0120302311213122-1301011022001230-1121210233020211-0032012322223313-3021003133203203-2111011010322003-1113322113010010)
- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0213222032201133-3023322220200332-2310211023202102-3033100212021031-1023303020033010-2210310010012322-1111231210333030-2223313100200132)
- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-3221023301133323-2200330310333212-3321313022032001-2211303203212302-2021212020320233-3202011003112100-3320303230112020-3303311302030301)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-2100303003003233-1303000303212100-1023000322330013-3110102322202201-2211030011311102-0022031123213330-1023311213221311-2211103221220212)
- [ingress_egress_gw.dc_cluster_group_inside_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-0222122011023220-0100000202223003-0111101023021010-1333230300022331-0311230122321120-2320122200100210-1000223213213100-3323312031332330)
- [ingress_egress_gw.dc_cluster_group_outside_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-0020231301003320-3023012122320023-0221232300021022-3032213000300030-2213132310000101-1033331132210202-2000333030130223-3001323103001320)
- [ingress_egress_gw.forward_proxy_allow_all](data-sources--aws_vpc_site--reference--group-002.md#canonical-1002320332003012-1322332301331332-1321130113113230-0013121301031323-2313001313013132-2130103332211001-1002122032221131-0201233333331021)
- [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-1332122323112132-3002221202203112-1112223122230110-3202203020123013-1011121222220200-2331333002313323-3000011132300012-2020031001021321)
- [ingress_egress_gw.inside_static_routes](data-sources--aws_vpc_site--reference--group-002.md#canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000)
- [ingress_egress_gw.no_dc_cluster_group](data-sources--aws_vpc_site--reference--group-003.md#canonical-2113321200233231-3011122321321213-2322321230202210-0323003112201000-0212212222101310-3103103013032210-3111111010332102-2202300302033122)
- [ingress_egress_gw.no_forward_proxy](data-sources--aws_vpc_site--reference--group-003.md#canonical-2312012003101212-1030213320311213-0123230122232202-1310002303110300-2311102011011110-0312001221202332-0122021121031220-1230121022010103)
- [ingress_egress_gw.no_global_network](data-sources--aws_vpc_site--reference--group-003.md#canonical-3012322203311023-1122021233203330-0131320222010020-1011203310102133-1133311130333020-0131031232013230-2110210300100131-2131013320322233)
- [ingress_egress_gw.no_inside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122332310030122-2301203331321331-3312021331313032-0110313021220211-3000212223310100-3103123301113003-2201001303021001-3310113103222230)
- [ingress_egress_gw.no_network_policy](data-sources--aws_vpc_site--reference--group-003.md#canonical-2322012331311110-3013210130000330-0300331233210121-3001320133210330-3030230023030101-3300221022323202-0233223101003130-3122210101130021)
- [ingress_egress_gw.no_outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-0122030330020231-1020100201333223-2023110003022323-2222101300102211-3213133230113230-1320232212211222-1130122001323210-0110110010011320)
- [ingress_egress_gw.outside_static_routes](data-sources--aws_vpc_site--reference--group-003.md#canonical-1122000232233003-2122011031010123-0021231130333000-3313113033133212-2321012033311302-3131021322322330-3033200311122000-1103220332132312)
- [ingress_egress_gw.performance_enhancement_mode](data-sources--aws_vpc_site--reference--group-003.md#canonical-3333111201203120-1223122133331331-2131312003210200-0303001232012121-2122230310310231-0010001201223202-3313133220111030-3032010201000230)
- [ingress_egress_gw.sm_connection_public_ip](data-sources--aws_vpc_site--reference--group-003.md#canonical-1032302033302121-2121112031020130-2120010231320100-1213021300031120-0122002321131212-2311120210210313-3103011330111033-2012123323331010)
- [ingress_egress_gw.sm_connection_pvt_ip](data-sources--aws_vpc_site--reference--group-003.md#canonical-1031100033200031-2003100101032231-1010002231031223-1103230320030133-0110003322321200-3003223200303311-1023021301231301-1232111011220221)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0332002031330310-1001311113102012-2213302223032332-3103322312230021-0130020120000332-0303301210030231-0100313220202201-0121201020100231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132332213301210-1112323213231231-0322011312230100-0110302132123001-1331000022321030-2212020310311110-0123023303331233-3322002023001120"></a>

## ingress_egress_gw.active_enhanced_firewall_policies — active_enhanced_firewall_policies / 110031131020 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.active_enhanced_firewall_policies

<a id="canonical-3310220132111100-2103213210221332-1220222210311011-0211120201312213-3311203212202111-2332303010213022-0322213013132321-3101121121130021"></a>

Type: `"single"`. Computed.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

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

<a id="canonical-0000332012302202-3113133321311003-0211303133012001-2000010133211110-1220033203012230-2321003112010002-2222202312010301-1131011320002003"></a>

## Direct properties — active_enhanced_firewall_policies / 110031131020 / 3

- [enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-3300230002033222-3103200100100111-1131232203003223-3232010013303222-1222103003131203-0203212332030013-3300331310123320-2321222103332120): complete subsection reference.

<a id="canonical-3321330310300220-3333220102333011-2000333003110313-0132111012301103-2330003302033210-2000033233032313-1103311220220220-0313312113103011"></a>

## Next pages — active_enhanced_firewall_policies / 110031131020 / 4

- [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-3300230002033222-3103200100100111-1131232203003223-3232010013303222-1222103003131203-0203212332030013-3300331310123320-2321222103332120)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3300230002033222-3103200100100111-1131232203003223-3232010013303222-1222103003131203-0203212332030013-3300331310123320-2321222103332120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231200330011312-1322302120131313-1032200032230310-0332220012000230-2232320113331010-0112132030322213-1102133223332013-2011011012210010"></a>

## ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies — enhanced_firewall_policies / 002131321033 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.active_enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-0332002031330310-1001311113102012-2213302223032332-3103322312230021-0130020120000332-0303301210030231-0100313220202201-0121201020100231)
- ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-1220231300111323-2122122323132130-0303223221232310-3002103333210110-3223221130200333-3302320232201000-2003133131231232-0321202000020032"></a>

Type: `"list"`. Computed.

Ordered List of Enhanced Firewall Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-1332302100023010-0022313111030213-3100203222211201-0213331210301111-1123111131320010-2212010033002022-0310123230323112-1130121130213123"></a>

## Direct properties — enhanced_firewall_policies / 002131321033 / 3

<a id="canonical-1223103302302320-3311010310133020-2103102232313001-3222231231011322-3113030323121012-2020201211133300-3300210301000330-3001223313313020"></a>

<a id="canonical-1022001121133133-1033313010210332-3323000031303223-0022131202300022-0220103100130311-1321023002100020-0311210300011301-2121102210002113"></a>

## name property — enhanced_firewall_policies / 002131321033 / 4

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

<a id="canonical-2210121113112212-1331232112223102-1110030132111302-1302220000331010-3122211203311203-2200031322121201-0102223100012333-2233130112301231"></a>

<a id="canonical-2130233132210320-3012112132121311-2101021333030032-3103012122013210-0313021113330230-3321023332113130-3102123000321210-1202302013321112"></a>

## namespace property — enhanced_firewall_policies / 002131321033 / 5

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

<a id="canonical-2013201221321122-1300100220322021-1122331223001301-1312202110121023-2313302012330130-2331113003200331-3223220303310111-3131003322011120"></a>

<a id="canonical-0220011201322032-2013133020221133-2223202010320030-2131020013330003-1022032230303033-3001213321123312-0303320120011221-2230122202013213"></a>

## tenant property — enhanced_firewall_policies / 002131321033 / 6

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

<a id="canonical-2022321131013321-1231113231303211-3012100230032103-0022103232321203-0121202320203330-2001230310311033-2103222010321220-1210223111231110"></a>

## Next pages — enhanced_firewall_policies / 002131321033 / 7

- [ingress_egress_gw.active_enhanced_firewall_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-0332002031330310-1001311113102012-2213302223032332-3103322312230021-0130020120000332-0303301210030231-0100313220202201-0121201020100231)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2321211103101121-0011230123322020-0302113111223322-0023231302301003-3332213010313020-2311212131220121-2111113320233013-1323031001221320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232322302220313-3331023210133130-0023221321230103-2032223322232020-3130012211031011-0203302211102100-1113230211331133-3013131121321030"></a>

## ingress_egress_gw.active_forward_proxy_policies — active_forward_proxy_policies / 232222332131 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.active_forward_proxy_policies

<a id="canonical-1322030213302301-1132220131112111-1210222302310300-3330233022312010-0330000303111110-0330313300111221-2032112323122300-0203310310300030"></a>

Type: `"single"`. Computed.

Ordered List of Forward Proxy Policies active.

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

<a id="canonical-3102331322031003-1100311120113011-3331120300323320-2211102212311313-0130030330233032-2320302300111110-1033311103102221-3331111002231011"></a>

## Direct properties — active_forward_proxy_policies / 232222332131 / 3

- [forward_proxy_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-0300133330303002-1321322223211331-0331013311121330-3330032213220133-0033020101200211-0321230201113133-2311311113000200-0102010203210013): complete subsection reference.

<a id="canonical-3133100032330211-2220320012212303-2120031322200112-2320110312312333-3301212021031110-0020102302233303-1022211331220233-3310221113323113"></a>

## Next pages — active_forward_proxy_policies / 232222332131 / 4

- [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-0300133330303002-1321322223211331-0331013311121330-3330032213220133-0033020101200211-0321230201113133-2311311113000200-0102010203210013)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0300133330303002-1321322223211331-0331013311121330-3330032213220133-0033020101200211-0321230201113133-2311311113000200-0102010203210013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202131233012332-0312211312102030-3112010332320201-2210023300120013-0312322320212220-1021132230021020-1300320222311303-0111031032032031"></a>

## ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies — forward_proxy_policies / 023321112133 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.active_forward_proxy_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-2321211103101121-0011230123322020-0302113111223322-0023231302301003-3332213010313020-2311212131220121-2111113320233013-1323031001221320)
- ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-0022212303203032-1331231221331302-2302121311000333-1330222221302013-2300220102101321-1201232223212200-0130320110110122-2221231221031232"></a>

Type: `"list"`. Computed.

Ordered List of Forward Proxy Policies active.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0133213131302211-2203322121203123-0003303002323023-1020302231012133-1000030222233032-0022330022120231-2203132101332010-3310330203113020"></a>

## Direct properties — forward_proxy_policies / 023321112133 / 3

<a id="canonical-1003001030110310-2301013202131021-2303112313333132-1211320023310222-1001223101112032-3303331201322322-3010310202102133-1230301033311233"></a>

<a id="canonical-2312102302130022-1310200101203303-1020012301221101-3203013100111322-0332032111230203-3113013213210231-1223103101000301-1220011200210023"></a>

## name property — forward_proxy_policies / 023321112133 / 4

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

<a id="canonical-1122001012213121-0202333132120220-2312102130100130-3303010101220101-1221033001122123-0323323011000011-1231113023020203-0102103322031231"></a>

<a id="canonical-1102333120311100-2200300121021202-1310322223300302-1223001331333031-2232212231033032-3232222202120103-1132033113000113-1130103011001002"></a>

## namespace property — forward_proxy_policies / 023321112133 / 5

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

<a id="canonical-2231022233310302-0321211010023220-1332312021211331-1312222002002320-2303021123020123-3102210010210331-0200030232233212-3202323223102323"></a>

<a id="canonical-0212131210001213-1103011221311333-3133003001231012-1112112332102110-3313032333020222-1010330231032113-3301230112303302-1101233021302001"></a>

## tenant property — forward_proxy_policies / 023321112133 / 6

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

<a id="canonical-1102031222002202-1222012202211211-1220232022022300-3031112103120001-1012133201231020-3012321301003233-3320101201020300-3100323202130030"></a>

## Next pages — forward_proxy_policies / 023321112133 / 7

- [ingress_egress_gw.active_forward_proxy_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-2321211103101121-0011230123322020-0302113111223322-0023231302301003-3332213010313020-2311212131220121-2111113320233013-1323031001221320)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1130103001330020-0120302311213122-1301011022001230-1121210233020211-0032012322223313-3021003133203203-2111011010322003-1113322113010010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310021223113312-0211021101322132-2123122112001331-0212232023021200-2102332122130030-2212232202231321-2133013233103300-1020023100333011"></a>

## ingress_egress_gw.active_network_policies — active_network_policies / 301031311033 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.active_network_policies

<a id="canonical-1031332322223321-2122012200301301-0102132103000030-0211332011320033-0130332311311023-3313023032131213-0003200312311022-0230332211300313"></a>

Type: `"single"`. Computed.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

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

<a id="canonical-0012211001113312-2000323302330103-1132100220211313-3132001222303002-3013013202230231-2100223100202012-3112331010120110-1323101111110313"></a>

## Direct properties — active_network_policies / 301031311033 / 3

- [network_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-1003013202010302-2303002333133233-0132203323203312-2111120012313000-3012312221101122-2023330131320211-0031101020030201-3021332332211133): complete subsection reference.

<a id="canonical-1132310101032031-2102201002300213-3023113303203233-2022321021223030-0333112032312031-3011012021220130-3230023013021202-2323322220322303"></a>

## Next pages — active_network_policies / 301031311033 / 4

- [ingress_egress_gw.active_network_policies.network_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-1003013202010302-2303002333133233-0132203323203312-2111120012313000-3012312221101122-2023330131320211-0031101020030201-3021332332211133)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1003013202010302-2303002333133233-0132203323203312-2111120012313000-3012312221101122-2023330131320211-0031101020030201-3021332332211133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012332331332230-2311012123123332-3012321233031332-0011032012213021-1113232011120110-0303032312132112-1202200230231232-1003322122030133"></a>

## ingress_egress_gw.active_network_policies.network_policies — network_policies / 023113022332 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.active_network_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-1130103001330020-0120302311213122-1301011022001230-1121210233020211-0032012322223313-3021003133203203-2111011010322003-1113322113010010)
- ingress_egress_gw.active_network_policies.network_policies

<a id="canonical-3100320202102303-3231101312302322-3211332321321320-2301022333121032-2123033000001202-0000300220023202-2021233121133220-1121100130223231"></a>

Type: `"list"`. Computed.

Ordered List of Firewall Policies active for this network firewall.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0232122213211110-1003333212020220-2001303102013201-1333123201131221-2220021030130011-1313020332213121-2030230022003220-3020222320020233"></a>

## Direct properties — network_policies / 023113022332 / 3

<a id="canonical-1010010232212111-3201331213021132-3122232212010322-0013133103203113-0131120322322032-0133032132301303-1321303010201101-0303220011230231"></a>

<a id="canonical-1020313230122033-0223032333233213-0300223031211311-1113313323233301-1100012010310013-1313302033301203-1221210230301132-2231121113032230"></a>

## name property — network_policies / 023113022332 / 4

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

<a id="canonical-1021312312200020-2212000321233201-3213221102211313-1202333301331033-0313321100112122-0022232010112303-2220213103110012-3301321302101031"></a>

<a id="canonical-1033101203232110-2003020113030323-0102333333233303-0230203130110111-2133222202303111-0331111133012211-3123032112100131-1330322032231202"></a>

## namespace property — network_policies / 023113022332 / 5

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

<a id="canonical-1130332003332201-1003112323032332-1023302230111300-3003312012233311-0203211331003223-2320231320030032-0012010103323112-3303300112011220"></a>

<a id="canonical-0002131203332011-2033321203010231-0322131223111212-0031223010220201-2223030112033313-2022230101232311-2000021102323133-1011233210320120"></a>

## tenant property — network_policies / 023113022332 / 6

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

<a id="canonical-0122021330013303-2000121323132223-0103102302200100-0122102221223220-0323001300233023-3010313022030010-0021010212300110-1000132033033311"></a>

## Next pages — network_policies / 023113022332 / 7

- [ingress_egress_gw.active_network_policies](data-sources--aws_vpc_site--reference--group-002.md#canonical-1130103001330020-0120302311213122-1301011022001230-1121210233020211-0032012322223313-3021003133203203-2111011010322003-1113322113010010)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0213222032201133-3023322220200332-2310211023202102-3033100212021031-1023303020033010-2210310010012322-1111231210333030-2223313100200132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100201231003031-0132111220101111-1212320023022102-3011031012031223-2220311120220301-3223023303133111-3003102300303111-0210321230322022"></a>

## ingress_egress_gw.allowed_vip_port — allowed_vip_port / 320200023102 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.allowed_vip_port

<a id="canonical-3202313033200222-1003300232102131-0213200031020232-3102112223010303-2102230232101112-0111231303222020-0201300111111030-0223310201010322"></a>

Type: `"single"`. Computed.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

<a id="canonical-1011032101300132-3123011322002113-2102201330121220-2203233222000222-1130033203032312-0320113231221321-3111120333103001-0133310310000002"></a>

## Direct properties — allowed_vip_port / 320200023102 / 3

- [custom_ports](data-sources--aws_vpc_site--reference--group-002.md#canonical-1323201212111203-1313230302100121-2103322011213301-1000322332111003-0333102013133212-3310123110121312-0020201113101030-3020022212022230): complete subsection reference.

- [disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-1113203002321320-2221322112132031-0203100113313120-0100021022333220-1022123313333002-0002113112233303-0130000020110010-2300023303012031): complete subsection reference.

- [use_http_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-2031031300323302-0301311321031313-0322022013332012-3302112021103212-1201032333222323-1333330023221323-0132000333322302-3033230133202003): complete subsection reference.

- [use_http_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-1303031121020311-3112123330300131-0121031212032302-3313221011322100-0031103323003102-2033033312130131-1321301113031013-0211001120311300): complete subsection reference.

- [use_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-3023012323020210-1331312020303022-2032113003233211-0203230101112303-1232012310012121-3031331023021323-3321332111203301-1103321122112003): complete subsection reference.

<a id="canonical-3313230200222213-3133110231320020-1031311033312230-1300000333000333-2102002213003333-1321112311222031-3032230221003202-3103030111100013"></a>

## Next pages — allowed_vip_port / 320200023102 / 4

- [ingress_egress_gw.allowed_vip_port.custom_ports](data-sources--aws_vpc_site--reference--group-002.md#canonical-1323201212111203-1313230302100121-2103322011213301-1000322332111003-0333102013133212-3310123110121312-0020201113101030-3020022212022230)
- [ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-1113203002321320-2221322112132031-0203100113313120-0100021022333220-1022123313333002-0002113112233303-0130000020110010-2300023303012031)
- [ingress_egress_gw.allowed_vip_port.use_http_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-2031031300323302-0301311321031313-0322022013332012-3302112021103212-1201032333222323-1333330023221323-0132000333322302-3033230133202003)
- [ingress_egress_gw.allowed_vip_port.use_http_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-1303031121020311-3112123330300131-0121031212032302-3313221011322100-0031103323003102-2033033312130131-1321301113031013-0211001120311300)
- [ingress_egress_gw.allowed_vip_port.use_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-3023012323020210-1331312020303022-2032113003233211-0203230101112303-1232012310012121-3031331023021323-3321332111203301-1103321122112003)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1323201212111203-1313230302100121-2103322011213301-1000322332111003-0333102013133212-3310123110121312-0020201113101030-3020022212022230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201333231321220-2013223123233030-0232013032221123-1201020120231031-3013033232221232-2321121000032321-2110201103211312-3012120112112012"></a>

## ingress_egress_gw.allowed_vip_port.custom_ports — custom_ports / 130312210311 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0213222032201133-3023322220200332-2310211023202102-3033100212021031-1023303020033010-2210310010012322-1111231210333030-2223313100200132)
- ingress_egress_gw.allowed_vip_port.custom_ports

<a id="canonical-1302023102331033-0002203132200123-1113102320130212-0011023030202223-3003211331103000-0332232223200120-1013102023012310-3312312000321110"></a>

Type: `"single"`. Computed.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

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

<a id="canonical-1323213230102122-1333003313111202-3032231333100210-3110133033012023-2002210311330122-3031020302020331-0211122030103113-1323103230000121"></a>

## Direct properties — custom_ports / 130312210311 / 3

<a id="canonical-1223212130331211-1011232323031221-1013021230033320-0102212102020120-2003023133311220-2032212321032121-0111120221131321-2301200330223200"></a>

<a id="canonical-3300030303011323-1022012123320233-3030222131003302-3012000210310032-3231230320203031-2301123123133102-1100132312022030-2223101103100313"></a>

## port_ranges property — custom_ports / 130312210311 / 4

Type: `"string"`. Computed.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-1132031222003313-2031210233003313-3132203203012002-2023201000012201-3030210212021202-0221300332113220-2223033100231211-1013220031131103"></a>

## Next pages — custom_ports / 130312210311 / 5

- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0213222032201133-3023322220200332-2310211023202102-3033100212021031-1023303020033010-2210310010012322-1111231210333030-2223313100200132)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1113203002321320-2221322112132031-0203100113313120-0100021022333220-1022123313333002-0002113112233303-0130000020110010-2300023303012031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332110313323110-3010310021213303-3311203010323302-2020011002011211-0122003021323201-3030232313200022-3221102203133011-3331130001311332"></a>

## ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port — disable_allowed_vip_port / 203313312023 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0213222032201133-3023322220200332-2310211023202102-3033100212021031-1023303020033010-2210310010012322-1111231210333030-2223313100200132)
- ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port

<a id="canonical-0131332103322321-1130303230023332-0203233320222103-1313121021202003-2201200100213022-1000333032203022-0032222032320231-3100302312230032"></a>

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

<a id="canonical-3323120222223222-0203211013112131-0223023222303100-2222100202200002-3231301313201001-2031111331210310-2110111331001023-2310311302312313"></a>

## Direct properties — disable_allowed_vip_port / 203313312023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013112213011311-2211130131120100-2030021331210111-3220332312332313-2000203021200023-2132222201310103-1302331011220021-1113322112113320"></a>

## Next pages — disable_allowed_vip_port / 203313312023 / 4

- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0213222032201133-3023322220200332-2310211023202102-3033100212021031-1023303020033010-2210310010012322-1111231210333030-2223313100200132)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2031031300323302-0301311321031313-0322022013332012-3302112021103212-1201032333222323-1333330023221323-0132000333322302-3033230133202003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100200012211303-0121111002012331-1313010002213212-1123102313101330-3123310031300232-3330210333313120-0022233020022202-3003332211103032"></a>

## ingress_egress_gw.allowed_vip_port.use_http_https_port — use_http_https_port / 330222030030 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0213222032201133-3023322220200332-2310211023202102-3033100212021031-1023303020033010-2210310010012322-1111231210333030-2223313100200132)
- ingress_egress_gw.allowed_vip_port.use_http_https_port

<a id="canonical-1330010201312331-2202331212132233-3313223222022203-0120122232013002-2122102032031131-0103332200201102-1033232113010212-3101002022132300"></a>

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

<a id="canonical-1122221221231330-1233131110300111-3302302112210103-1023222311213210-0011023032002102-2010121223001003-3301221333102001-3322300033322120"></a>

## Direct properties — use_http_https_port / 330222030030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223203130232210-2221313120033033-2333113211133231-1323130002132220-2223022133012332-2313103120120100-2330313332133312-1122113011123223"></a>

## Next pages — use_http_https_port / 330222030030 / 4

- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0213222032201133-3023322220200332-2310211023202102-3033100212021031-1023303020033010-2210310010012322-1111231210333030-2223313100200132)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1303031121020311-3112123330300131-0121031212032302-3313221011322100-0031103323003102-2033033312130131-1321301113031013-0211001120311300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120001021330323-0121312203303020-0022200330121110-2000322310222012-3003033033031333-1210030200310130-3123120333202222-3100101212310101"></a>

## ingress_egress_gw.allowed_vip_port.use_http_port — use_http_port / 330133302002 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0213222032201133-3023322220200332-2310211023202102-3033100212021031-1023303020033010-2210310010012322-1111231210333030-2223313100200132)
- ingress_egress_gw.allowed_vip_port.use_http_port

<a id="canonical-3323010223310002-1202313113020221-2120320211231131-2113300132010112-3333033121131130-0013113001112121-2201002311333033-3123211320222312"></a>

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

<a id="canonical-1312030223130113-3230000003102223-1232130330002320-2303011220022113-3131101213312212-2130132200331002-1011203231233311-3200013033333120"></a>

## Direct properties — use_http_port / 330133302002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210223002233203-3230230233213311-3203322333300102-0032022130213030-0003312303102110-2023202233223231-1130200110111132-2203002011011103"></a>

## Next pages — use_http_port / 330133302002 / 4

- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0213222032201133-3023322220200332-2310211023202102-3033100212021031-1023303020033010-2210310010012322-1111231210333030-2223313100200132)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3023012323020210-1331312020303022-2032113003233211-0203230101112303-1232012310012121-3031331023021323-3321332111203301-1103321122112003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222100101100232-2102003133230223-3101220010110311-2003313322133033-0210230212231231-1301312021030223-2112311330032003-0001220233112001"></a>

## ingress_egress_gw.allowed_vip_port.use_https_port — use_https_port / 130102331012 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0213222032201133-3023322220200332-2310211023202102-3033100212021031-1023303020033010-2210310010012322-1111231210333030-2223313100200132)
- ingress_egress_gw.allowed_vip_port.use_https_port

<a id="canonical-2220012303210121-3111203022031103-3301233003312303-0002233110210333-2312022330221033-3223202230102330-2230310000010032-3322012232333331"></a>

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

<a id="canonical-0011101033103010-1113132123131010-2120030332301202-3131332003032233-1301330321112001-1101103000220120-2032031000121220-3332013313232003"></a>

## Direct properties — use_https_port / 130102331012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030233032303000-1231303020113002-3031233203100213-0023023022101110-0221331222001100-1112122201301201-0011023030332123-3222002211003102"></a>

## Next pages — use_https_port / 130102331012 / 4

- [ingress_egress_gw.allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0213222032201133-3023322220200332-2310211023202102-3033100212021031-1023303020033010-2210310010012322-1111231210333030-2223313100200132)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3221023301133323-2200330310333212-3321313022032001-2211303203212302-2021212020320233-3202011003112100-3320303230112020-3303311302030301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320213323303201-2332130122001011-2102220232111230-2313032002033132-3010032303110200-0323132333022010-2301323122322010-3100213322211011"></a>

## ingress_egress_gw.allowed_vip_port_sli — allowed_vip_port_sli / 112022011312 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.allowed_vip_port_sli

<a id="canonical-0033311122010213-3312110111112230-0000002112000010-3301321001313311-1233030112330001-1131020102133203-1322001210303011-2032321022131101"></a>

Type: `"single"`. Computed.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

<a id="canonical-0012023001033210-2032011220003312-3111203133032321-2023033023111313-1230030211222100-0103020131030022-0221023110110311-3133201331032201"></a>

## Direct properties — allowed_vip_port_sli / 112022011312 / 3

- [custom_ports](data-sources--aws_vpc_site--reference--group-002.md#canonical-1130112132310010-1220213112223231-0203331332210233-2221031031232113-0302212111012110-3110112121100211-3221131200003122-3330201020311113): complete subsection reference.

- [disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0020323002122210-2021311231302130-1101021112231313-3100323113300212-0032311201331222-2220320102000112-3333201233330120-2210211300232110): complete subsection reference.

- [use_http_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-2021220123301221-0013331210331232-3233031000021123-2201021311323100-1332123300030303-2233113111110231-0233212203132010-2010213210311030): complete subsection reference.

- [use_http_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-2001100313131200-1212202330110322-2032320212020322-1300331000013320-3123322211310001-3301112023313003-2200030002111122-3233230200102123): complete subsection reference.

- [use_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-2100203201121223-2211103100302021-0130111223330202-3302032331210311-1100302020130023-2111201010230311-1101310113020333-2130012202020210): complete subsection reference.

<a id="canonical-1011223232023222-1300211231102221-0002112031011212-0212111030303130-1030003212030111-2111321131310032-0301232300111231-3110021223201121"></a>

## Next pages — allowed_vip_port_sli / 112022011312 / 4

- [ingress_egress_gw.allowed_vip_port_sli.custom_ports](data-sources--aws_vpc_site--reference--group-002.md#canonical-1130112132310010-1220213112223231-0203331332210233-2221031031232113-0302212111012110-3110112121100211-3221131200003122-3330201020311113)
- [ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-0020323002122210-2021311231302130-1101021112231313-3100323113300212-0032311201331222-2220320102000112-3333201233330120-2210211300232110)
- [ingress_egress_gw.allowed_vip_port_sli.use_http_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-2021220123301221-0013331210331232-3233031000021123-2201021311323100-1332123300030303-2233113111110231-0233212203132010-2010213210311030)
- [ingress_egress_gw.allowed_vip_port_sli.use_http_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-2001100313131200-1212202330110322-2032320212020322-1300331000013320-3123322211310001-3301112023313003-2200030002111122-3233230200102123)
- [ingress_egress_gw.allowed_vip_port_sli.use_https_port](data-sources--aws_vpc_site--reference--group-002.md#canonical-2100203201121223-2211103100302021-0130111223330202-3302032331210311-1100302020130023-2111201010230311-1101310113020333-2130012202020210)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1130112132310010-1220213112223231-0203331332210233-2221031031232113-0302212111012110-3110112121100211-3221131200003122-3330201020311113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132030002311030-1323030020202311-0233132311313102-3111110010013000-2101012033010232-3012131313111020-0323123111003323-1122301001130030"></a>

## ingress_egress_gw.allowed_vip_port_sli.custom_ports — custom_ports / 100321121001 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-3221023301133323-2200330310333212-3321313022032001-2211303203212302-2021212020320233-3202011003112100-3320303230112020-3303311302030301)
- ingress_egress_gw.allowed_vip_port_sli.custom_ports

<a id="canonical-2131123211032032-2322103201300232-2112200220120113-0201302003312031-1121223222033123-2200330313103222-2310003231213023-1122221302011210"></a>

Type: `"single"`. Computed.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

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

<a id="canonical-0303323020330312-0110331210212233-1210301133333133-1120211200323220-1223321030203133-0100100300302200-3122122031200101-3310310031212311"></a>

## Direct properties — custom_ports / 100321121001 / 3

<a id="canonical-3322001121321231-0120320330330311-3310223303331111-1033303032132000-3030232203130021-3223323313213331-1033202133033213-1301303123023131"></a>

<a id="canonical-2031030210210010-0203232330303230-0223020103110023-0210220203320332-0202300130223220-0111320010312103-0211110123322112-0021033111023330"></a>

## port_ranges property — custom_ports / 100321121001 / 4

Type: `"string"`. Computed.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-3200001312032121-2332112133323303-2120322230313331-1302231330333002-0310331211000210-2312031202233231-1331220022123323-3110221311232201"></a>

## Next pages — custom_ports / 100321121001 / 5

- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-3221023301133323-2200330310333212-3321313022032001-2211303203212302-2021212020320233-3202011003112100-3320303230112020-3303311302030301)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0020323002122210-2021311231302130-1101021112231313-3100323113300212-0032311201331222-2220320102000112-3333201233330120-2210211300232110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132111312102333-1233132233132012-2333110021302201-1131112333011120-1000110112211003-0012222130230011-3203323001001112-1123301000103303"></a>

## ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port — disable_allowed_vip_port / 322300322113 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-3221023301133323-2200330310333212-3321313022032001-2211303203212302-2021212020320233-3202011003112100-3320303230112020-3303311302030301)
- ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port

<a id="canonical-1211133033030213-3221111222032002-2033221012331102-3100020310130133-0320321023313120-3030331313221000-2130210001021223-1100323332033021"></a>

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

<a id="canonical-0022222313321302-1232001021102100-2111323113100333-0223030300310131-0221232313311232-0320033122322203-2010100130201031-2113032132033012"></a>

## Direct properties — disable_allowed_vip_port / 322300322113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303000322223133-0021002130030313-1012001132020320-0213100201022023-0022011133030020-0000300332120233-3312221121012301-1123302202211131"></a>

## Next pages — disable_allowed_vip_port / 322300322113 / 4

- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-3221023301133323-2200330310333212-3321313022032001-2211303203212302-2021212020320233-3202011003112100-3320303230112020-3303311302030301)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2021220123301221-0013331210331232-3233031000021123-2201021311323100-1332123300030303-2233113111110231-0233212203132010-2010213210311030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033330022102002-3230123220131203-2133311210010321-0101132131211021-1221200303330023-2022331022212100-0312033110333013-1322110302102210"></a>

## ingress_egress_gw.allowed_vip_port_sli.use_http_https_port — use_http_https_port / 330301011202 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-3221023301133323-2200330310333212-3321313022032001-2211303203212302-2021212020320233-3202011003112100-3320303230112020-3303311302030301)
- ingress_egress_gw.allowed_vip_port_sli.use_http_https_port

<a id="canonical-1011133320030323-0300230332113022-1112123220130331-2222031333230300-1103323303211020-2202122123220132-1230300010132332-0231022121032212"></a>

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

<a id="canonical-0232312122012311-2312002200200002-2332100002310122-3021020100213222-3202230131102330-0032112332302210-0120122223213020-2113210230122322"></a>

## Direct properties — use_http_https_port / 330301011202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020311210002232-0233001231330032-3003331301333212-2033320030023020-0002230321110101-2112121010313303-3100200011130200-0120320010102311"></a>

## Next pages — use_http_https_port / 330301011202 / 4

- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-3221023301133323-2200330310333212-3321313022032001-2211303203212302-2021212020320233-3202011003112100-3320303230112020-3303311302030301)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2001100313131200-1212202330110322-2032320212020322-1300331000013320-3123322211310001-3301112023313003-2200030002111122-3233230200102123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133101020131312-1323212213131033-1102213321132220-1220210223331013-0111223311030123-1122013130002002-2030000331312100-2212100000023321"></a>

## ingress_egress_gw.allowed_vip_port_sli.use_http_port — use_http_port / 133003313223 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-3221023301133323-2200330310333212-3321313022032001-2211303203212302-2021212020320233-3202011003112100-3320303230112020-3303311302030301)
- ingress_egress_gw.allowed_vip_port_sli.use_http_port

<a id="canonical-3203000002130110-2201231022011233-3300213232121100-3233211321132211-3200321113103221-0122101033231201-2210323111023122-0011023222331023"></a>

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

<a id="canonical-1113232203211102-3332231021022011-3330131201230031-1301220023312112-2303210133013010-0102022013203020-2221002020021122-3102332213123320"></a>

## Direct properties — use_http_port / 133003313223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320200230132022-1110101312210032-0020112122312321-3031300310232023-0023330202130110-1303223131102233-2313233213032310-0230110100012110"></a>

## Next pages — use_http_port / 133003313223 / 4

- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-3221023301133323-2200330310333212-3321313022032001-2211303203212302-2021212020320233-3202011003112100-3320303230112020-3303311302030301)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2100203201121223-2211103100302021-0130111223330202-3302032331210311-1100302020130023-2111201010230311-1101310113020333-2130012202020210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322012101120222-2102300010320200-3202220032200210-1130021202010112-0300120120323333-2111013010020021-2101302111110310-3213230302133000"></a>

## ingress_egress_gw.allowed_vip_port_sli.use_https_port — use_https_port / 322232130312 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-3221023301133323-2200330310333212-3321313022032001-2211303203212302-2021212020320233-3202011003112100-3320303230112020-3303311302030301)
- ingress_egress_gw.allowed_vip_port_sli.use_https_port

<a id="canonical-0022321102023312-2021323023313322-0332222202112120-3203320313302121-1300013322031113-2211201131021010-1201030223303310-0322033300000310"></a>

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

<a id="canonical-2103213030130220-0123320133022032-2131220312002230-0310131221120111-3333213310130223-0210021310103211-1330300322200301-1110003333112101"></a>

## Direct properties — use_https_port / 322232130312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310001211320032-0030311203130133-1032021010112120-1022232310322122-3302321313222201-3211123311132323-2131322230313010-0333221310030322"></a>

## Next pages — use_https_port / 322232130312 / 4

- [ingress_egress_gw.allowed_vip_port_sli](data-sources--aws_vpc_site--reference--group-002.md#canonical-3221023301133323-2200330310333212-3321313022032001-2211303203212302-2021212020320233-3202011003112100-3320303230112020-3303311302030301)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-2100303003003233-1303000303212100-1023000322330013-3110102322202201-2211030011311102-0022031123213330-1023311213221311-2211103221220212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021233330203333-1001310102102013-0130320210220110-2031020121333101-0213131300233113-3332120003221113-3013110101102310-0010022123000110"></a>

## ingress_egress_gw.az_nodes — az_nodes / 220302332233 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.az_nodes

<a id="canonical-0221320323213323-0101130220031013-2211010201000011-2030223210121012-2012012110221332-0031203102301111-1021011320301103-3022110223232130"></a>

Type: `"list"`. Computed.

Only Single AZ or Three AZ(s) nodes are supported currently.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

<a id="canonical-2002202210221200-1320331203013301-1003121311131131-1113030200201100-0212221001031220-1032021110300000-0100100333101303-0312010323301302"></a>

## Direct properties — az_nodes / 220302332233 / 3

<a id="canonical-3302300101011221-1202202123312333-0321012103121203-0211323130303101-0001233323213302-2310222322133110-0133311110003212-1032021322130023"></a>

<a id="canonical-1331120320132030-2301121310100100-2010321201001111-0320202100032313-2323123113330331-2233131021323220-1232210032101020-3303002222011023"></a>

## aws_az_name property — az_nodes / 220302332233 / 4

Type: `"string"`. Computed.

AWS availability zone, must be consistent with the selected AWS region.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [inside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-0303122223100200-3001101120030203-2013330112331030-0323030001133312-2323011323130310-3310121032023322-0311122211221302-3322030330000232): complete subsection reference.

- [outside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-1022132012331233-2001332021122001-2322300333201200-2302001332010303-0300011232103201-0233033120113113-0333032310101330-3121123010220230): complete subsection reference.

- [reserved_inside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-1312113212332222-2220131312201220-1321331022110202-2121021232302331-2101231130101112-2031003312100022-2001033003302233-3233332332121233): complete subsection reference.

- [workload_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-1011121011230023-2332201230321210-3011013133311100-0203121201210000-0031120113111012-1300220211110032-1011210302310211-1010013033123013): complete subsection reference.

<a id="canonical-2211202321022211-1112223311213010-2032222102111201-3031100200012001-2023000202030231-2223202223330132-0012112032332012-2113110103331033"></a>

## Next pages — az_nodes / 220302332233 / 5

- [ingress_egress_gw.az_nodes.inside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-0303122223100200-3001101120030203-2013330112331030-0323030001133312-2323011323130310-3310121032023322-0311122211221302-3322030330000232)
- [ingress_egress_gw.az_nodes.outside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-1022132012331233-2001332021122001-2322300333201200-2302001332010303-0300011232103201-0233033120113113-0333032310101330-3121123010220230)
- [ingress_egress_gw.az_nodes.reserved_inside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-1312113212332222-2220131312201220-1321331022110202-2121021232302331-2101231130101112-2031003312100022-2001033003302233-3233332332121233)
- [ingress_egress_gw.az_nodes.workload_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-1011121011230023-2332201230321210-3011013133311100-0203121201210000-0031120113111012-1300220211110032-1011210302310211-1010013033123013)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0303122223100200-3001101120030203-2013330112331030-0323030001133312-2323011323130310-3310121032023322-0311122211221302-3322030330000232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302223131110110-2133311010033320-0233202313303033-3010020021323101-0203203102100130-3222332113312311-3100301313211211-0121313113220003"></a>

## ingress_egress_gw.az_nodes.inside_subnet — inside_subnet / 032310231003 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-2100303003003233-1303000303212100-1023000322330013-3110102322202201-2211030011311102-0022031123213330-1023311213221311-2211103221220212)
- ingress_egress_gw.az_nodes.inside_subnet

<a id="canonical-0301110103132331-2202211003132221-2032322122221232-1332200122123232-1212022130333230-1020301331213032-0003311212332121-1030332000312323"></a>

Type: `"single"`. Computed.

Configuration parameter for inside subnet.

Upstream description:

Parameters for AWS subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

<a id="canonical-1001210210202011-3322130300323303-3121332031100313-1320321211230102-0012301213332222-1013323320300213-3113002211122222-3212310231131211"></a>

## Direct properties — inside_subnet / 032310231003 / 3

<a id="canonical-3212331220131203-1232133222032000-1301330213113330-2203112003210102-1211322302222023-3321322323221123-2113320332022302-1321232133302121"></a>

<a id="canonical-3101110021330113-3000202221023120-1303302110312321-1301221030011303-2001300130323211-0031032213220112-2222232103201000-2220011100222010"></a>

## existing_subnet_id property — inside_subnet / 032310231003 / 4

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-0101320303233001-3312310032303201-1102330312131202-3333131203233000-3230012233310020-0321100300321303-1332100220130123-3231021013012333): complete subsection reference.

<a id="canonical-0202303323232323-0213231031021210-3131130322031233-0323233112010132-0232301103112022-0233112320022213-1330003213232333-1010031102111111"></a>

## Next pages — inside_subnet / 032310231003 / 5

- [ingress_egress_gw.az_nodes.inside_subnet.subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-0101320303233001-3312310032303201-1102330312131202-3333131203233000-3230012233310020-0321100300321303-1332100220130123-3231021013012333)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-2100303003003233-1303000303212100-1023000322330013-3110102322202201-2211030011311102-0022031123213330-1023311213221311-2211103221220212)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0101320303233001-3312310032303201-1102330312131202-3333131203233000-3230012233310020-0321100300321303-1332100220130123-3231021013012333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213022011113122-3021330032122201-3001300101212111-0101221302323103-3212023103231333-3203213121011013-3332022323102021-1310231230102201"></a>

## ingress_egress_gw.az_nodes.inside_subnet.subnet_param — subnet_param / 030330312200 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-2100303003003233-1303000303212100-1023000322330013-3110102322202201-2211030011311102-0022031123213330-1023311213221311-2211103221220212)
- [ingress_egress_gw.az_nodes.inside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-0303122223100200-3001101120030203-2013330112331030-0323030001133312-2323011323130310-3310121032023322-0311122211221302-3322030330000232)
- ingress_egress_gw.az_nodes.inside_subnet.subnet_param

<a id="canonical-3032331122200310-0323103122333333-2032121011031001-2201013011101230-3003030300320303-0021302123100122-2321130312113212-1130031112221111"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

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

<a id="canonical-1122230123023331-3110212322312113-1311323231213101-2110121101102322-0122112121203120-1302332200111200-3131223132113313-0310210111313201"></a>

## Direct properties — subnet_param / 030330312200 / 3

<a id="canonical-0210312310202231-1223003310031321-0131212130320220-1031001301321333-3000202011120213-2020131310221320-0133002031102122-0110010213110312"></a>

<a id="canonical-0102211030030220-3230111230333022-0330113112303003-1113213130102200-1132000231221032-1301210301231212-1001201000022121-1310030303302011"></a>

## IPv4 property — subnet_param / 030330312200 / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-2122302220332111-3021001010130111-1230303111203123-1320112122021000-0212311101211301-3301133330021002-0031033102100020-0310200110020311"></a>

## Next pages — subnet_param / 030330312200 / 5

- [ingress_egress_gw.az_nodes.inside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-0303122223100200-3001101120030203-2013330112331030-0323030001133312-2323011323130310-3310121032023322-0311122211221302-3322030330000232)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1022132012331233-2001332021122001-2322300333201200-2302001332010303-0300011232103201-0233033120113113-0333032310101330-3121123010220230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200321313133130-1033113323212123-0111220010000112-0330223332003332-2131321322131332-1333122103101122-1230010322112220-0013100233222012"></a>

## ingress_egress_gw.az_nodes.outside_subnet — outside_subnet / 101132113110 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-2100303003003233-1303000303212100-1023000322330013-3110102322202201-2211030011311102-0022031123213330-1023311213221311-2211103221220212)
- ingress_egress_gw.az_nodes.outside_subnet

<a id="canonical-2321012212231200-3332233201112312-1001322030222320-3123130011303201-0030300233102302-1000023002000002-2021131012000332-1132331002100132"></a>

Type: `"single"`. Computed.

Configuration parameter for outside subnet.

Upstream description:

Parameters for AWS subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

<a id="canonical-1102322022123332-2332323212130103-0310112102321033-0110331311113110-1331330123123130-1110131203012332-3100112130013122-3011303320003111"></a>

## Direct properties — outside_subnet / 101132113110 / 3

<a id="canonical-0210103310220213-2212301113333221-2201003113220011-1020332201000213-2100110113321132-3233023332311022-3300103203001002-0031111310032313"></a>

<a id="canonical-2022122212033110-3112300202011110-3212111122201033-0123122220201330-0210122222123101-1230310311321010-1100033121230213-2302300212000002"></a>

## existing_subnet_id property — outside_subnet / 101132113110 / 4

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-3132211011121030-2333230030220010-0211313003321033-2230300122032223-3010000301332002-1301323302202111-1032113111223002-3110021013313012): complete subsection reference.

<a id="canonical-0013123231131332-1322023131123220-3322313212010112-0203111101201020-2332320320333030-1321202100202030-3303033121233211-1232210030323212"></a>

## Next pages — outside_subnet / 101132113110 / 5

- [ingress_egress_gw.az_nodes.outside_subnet.subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-3132211011121030-2333230030220010-0211313003321033-2230300122032223-3010000301332002-1301323302202111-1032113111223002-3110021013313012)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-2100303003003233-1303000303212100-1023000322330013-3110102322202201-2211030011311102-0022031123213330-1023311213221311-2211103221220212)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3132211011121030-2333230030220010-0211313003321033-2230300122032223-3010000301332002-1301323302202111-1032113111223002-3110021013313012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101201010331203-2302100112213321-2003311210232201-3221011231332332-1013300112021032-1222122133201130-1112311133111323-0301202003132211"></a>

## ingress_egress_gw.az_nodes.outside_subnet.subnet_param — subnet_param / 031223203233 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-2100303003003233-1303000303212100-1023000322330013-3110102322202201-2211030011311102-0022031123213330-1023311213221311-2211103221220212)
- [ingress_egress_gw.az_nodes.outside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-1022132012331233-2001332021122001-2322300333201200-2302001332010303-0300011232103201-0233033120113113-0333032310101330-3121123010220230)
- ingress_egress_gw.az_nodes.outside_subnet.subnet_param

<a id="canonical-1011230000323301-1301033323301033-0303233212230113-3130101133331202-3220211121323133-0110002012110331-2023023230022321-1130031203211311"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

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

<a id="canonical-2312213313021201-0022133002210231-0111101102320111-0221331132313212-2123310320110310-3232130003221302-1221200312310201-3111213132103030"></a>

## Direct properties — subnet_param / 031223203233 / 3

<a id="canonical-1233030211230113-0231020322133010-2230201013102331-3010101210221230-3333232130212003-1331132022310320-1111321220111222-1120202011310013"></a>

<a id="canonical-1021231102111302-3010301012333122-3013320001201210-1303223321130111-1132303301012001-1011103203020332-3031233103133122-2220011213200322"></a>

## IPv4 property — subnet_param / 031223203233 / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-3222012033112111-3011312230013310-3300112331100223-0100000301012023-0210300001121221-2110211312310333-1231003003022221-2223013200331031"></a>

## Next pages — subnet_param / 031223203233 / 5

- [ingress_egress_gw.az_nodes.outside_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-1022132012331233-2001332021122001-2322300333201200-2302001332010303-0300011232103201-0233033120113113-0333032310101330-3121123010220230)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1312113212332222-2220131312201220-1321331022110202-2121021232302331-2101231130101112-2031003312100022-2001033003302233-3233332332121233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320003203023033-0301201332122210-2003213203211103-2302213011003200-1232012113021300-2023013300123120-2023003210111023-0011101203321032"></a>

## ingress_egress_gw.az_nodes.reserved_inside_subnet — reserved_inside_subnet / 133211310003 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-2100303003003233-1303000303212100-1023000322330013-3110102322202201-2211030011311102-0022031123213330-1023311213221311-2211103221220212)
- ingress_egress_gw.az_nodes.reserved_inside_subnet

<a id="canonical-1223100113003021-3112331210101221-3103323020200021-1020011203301211-3130221303221022-2233030131301200-3121202120212220-3322322312303322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for reserved inside subnet.

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

<a id="canonical-3320133233322112-1312023133302211-3130312000113311-0322031003320100-2200013303230303-1222010022133231-1113202001003223-3330000330332033"></a>

## Direct properties — reserved_inside_subnet / 133211310003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012332220023203-0220113333213212-0011013132332130-1213003332130123-3230023121002130-1023212230323211-1201031103110001-2330021103132100"></a>

## Next pages — reserved_inside_subnet / 133211310003 / 4

- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-2100303003003233-1303000303212100-1023000322330013-3110102322202201-2211030011311102-0022031123213330-1023311213221311-2211103221220212)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1011121011230023-2332201230321210-3011013133311100-0203121201210000-0031120113111012-1300220211110032-1011210302310211-1010013033123013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221232303002322-0231323031230223-0031130320120220-0201320320312002-3100122112131023-1113203132123100-1231323331202112-3332220313303201"></a>

## ingress_egress_gw.az_nodes.workload_subnet — workload_subnet / 210023012023 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-2100303003003233-1303000303212100-1023000322330013-3110102322202201-2211030011311102-0022031123213330-1023311213221311-2211103221220212)
- ingress_egress_gw.az_nodes.workload_subnet

<a id="canonical-0320121012111322-3231121212313121-2223103102130321-0131003111310302-1111301203121033-1030203222020212-1012021122023132-0001122012013331"></a>

Type: `"single"`. Computed.

Configuration parameter for workload subnet.

Upstream description:

Parameters for AWS subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

<a id="canonical-3120100303023303-0332122023320103-1303222222323002-1002103000102123-0320322131032133-0233130331022211-1231213110133120-2033303032012323"></a>

## Direct properties — workload_subnet / 210023012023 / 3

<a id="canonical-3202120122200232-2303003302031011-3113000120132310-2221212200010022-0322311233010121-2001221133231123-1201001232133111-2030013221011310"></a>

<a id="canonical-2332132330110113-1202202120321033-2332332212112320-2203112021230212-3330122113303110-1323203310311002-0020311311020113-1332132032120233"></a>

## existing_subnet_id property — workload_subnet / 210023012023 / 4

Type: `"string"`. Computed.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-1323210013213120-1100132113113201-0033213003301032-2023001032022313-0313230001222202-2000132023323310-1302212200133021-1033211111111321): complete subsection reference.

<a id="canonical-3200103211133033-1222310110021030-3233100020102100-1222112032102012-2311133020120311-2202210102123012-3210131033223003-1301331130010113"></a>

## Next pages — workload_subnet / 210023012023 / 5

- [ingress_egress_gw.az_nodes.workload_subnet.subnet_param](data-sources--aws_vpc_site--reference--group-002.md#canonical-1323210013213120-1100132113113201-0033213003301032-2023001032022313-0313230001222202-2000132023323310-1302212200133021-1033211111111321)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-2100303003003233-1303000303212100-1023000322330013-3110102322202201-2211030011311102-0022031123213330-1023311213221311-2211103221220212)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1323210013213120-1100132113113201-0033213003301032-2023001032022313-0313230001222202-2000132023323310-1302212200133021-1033211111111321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321310211233113-3001321210001303-2303103131123303-0003310021101123-3133013200302231-2002320101030300-1030213333300102-2130223120202203"></a>

## ingress_egress_gw.az_nodes.workload_subnet.subnet_param — subnet_param / 102102231200 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.az_nodes](data-sources--aws_vpc_site--reference--group-002.md#canonical-2100303003003233-1303000303212100-1023000322330013-3110102322202201-2211030011311102-0022031123213330-1023311213221311-2211103221220212)
- [ingress_egress_gw.az_nodes.workload_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-1011121011230023-2332201230321210-3011013133311100-0203121201210000-0031120113111012-1300220211110032-1011210302310211-1010013033123013)
- ingress_egress_gw.az_nodes.workload_subnet.subnet_param

<a id="canonical-2330202003311321-1011010230300003-3302223332320001-0223212112123213-1313030100212300-2120023013322202-0331113002023023-1321110133101030"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

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

<a id="canonical-0130333023303111-1333310323121012-1002102223321313-2120132311303221-3031000232201331-1333021001331323-1310212012000311-2000213231121032"></a>

## Direct properties — subnet_param / 102102231200 / 3

<a id="canonical-1321323022303123-2313232033133020-1103302330313023-1302102132322012-1300202303232012-2031013020133020-0330323321120321-3332112200303020"></a>

<a id="canonical-2213203231223332-0112232113313331-2333222112112110-3320031302031000-1201330033033030-2201211312030101-2213303100110321-2231102332133123"></a>

## IPv4 property — subnet_param / 102102231200 / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-1010333001003331-1301211111231201-0100032303030312-1202112123323331-3103001320032201-2333302100002131-2000331301123113-0330123212202021"></a>

## Next pages — subnet_param / 102102231200 / 5

- [ingress_egress_gw.az_nodes.workload_subnet](data-sources--aws_vpc_site--reference--group-002.md#canonical-1011121011230023-2332201230321210-3011013133311100-0203121201210000-0031120113111012-1300220211110032-1011210302310211-1010013033123013)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0222122011023220-0100000202223003-0111101023021010-1333230300022331-0311230122321120-2320122200100210-1000223213213100-3323312031332330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330201012203102-1031001201232223-2100032121110310-2333300003131002-1232110011122202-0131231202100312-2100001202022322-2300322333033131"></a>

## ingress_egress_gw.dc_cluster_group_inside_vn — dc_cluster_group_inside_vn / 301003102312 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.dc_cluster_group_inside_vn

<a id="canonical-1302331131303331-2003213203230200-0302032003120111-3231001202033002-1311103302112122-2332213202100220-0312323030132112-0123101131321033"></a>

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

<a id="canonical-0013210123101233-2010012123120330-1303131011331122-1320230312001112-1031310011120112-2033300211232331-2212301000323103-1131032223320223"></a>

## Direct properties — dc_cluster_group_inside_vn / 301003102312 / 3

<a id="canonical-1201200112200110-1000133102332202-0221020033331130-3121031030023202-3213300113233211-0023220330111320-2031102123320202-2002323331002101"></a>

<a id="canonical-2301001330021322-2202113312313212-2133203122212101-3122212232222221-2030233202212011-0232302111203330-2122331202333232-0320212231210032"></a>

## name property — dc_cluster_group_inside_vn / 301003102312 / 4

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

<a id="canonical-2101011020221013-1312222221103110-2130113323103232-2012222232100032-1110303020201101-1220010320102032-3031323010332002-2313030300000000"></a>

<a id="canonical-2323031121030110-2320100103333322-0003203000333010-2120221022323001-3101121110221121-0331333333301222-2211132131323313-3030310220211233"></a>

## namespace property — dc_cluster_group_inside_vn / 301003102312 / 5

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

<a id="canonical-3321113030032013-0003122220221010-3110113221011123-3200233223322201-2222332123111113-3000111030333323-3101320202012001-0130110311003330"></a>

<a id="canonical-1211122220300030-2011133010113323-3313130212313303-0000233112210020-1333101220130221-0210311220003103-3310113230113311-2313221313302222"></a>

## tenant property — dc_cluster_group_inside_vn / 301003102312 / 6

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

<a id="canonical-0022021211322112-0130103120223220-2301033102223113-1031121123231210-3212231032223133-3330112010131302-3011033030202223-2312112211012223"></a>

## Next pages — dc_cluster_group_inside_vn / 301003102312 / 7

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0020231301003320-3023012122320023-0221232300021022-3032213000300030-2213132310000101-1033331132210202-2000333030130223-3001323103001320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202302321232103-0002333030100203-0330111310320310-2033212030213013-2211132021321103-0123111021212220-2312211211001011-1311322303211011"></a>

## ingress_egress_gw.dc_cluster_group_outside_vn — dc_cluster_group_outside_vn / 301032213102 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.dc_cluster_group_outside_vn

<a id="canonical-3202300011231220-1031202201123333-1221211021110122-3121333121101200-2002102011103211-1121011010232111-3201132030203031-2222110230231323"></a>

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

<a id="canonical-1320112102221012-0212211120020012-2210033212321333-2021230310112031-0102200300323110-1110102230202210-2230111312121120-3032301322230110"></a>

## Direct properties — dc_cluster_group_outside_vn / 301032213102 / 3

<a id="canonical-0213011232312310-2300331121103012-2333303113113211-1021020133122032-2101223313230223-2303310231333131-0022131100000231-2010213101133322"></a>

<a id="canonical-3302300300231102-0033102101311031-0101103300232111-2222331302332222-2103321220021032-2201013133122123-1223311032212301-1222102331020311"></a>

## name property — dc_cluster_group_outside_vn / 301032213102 / 4

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

<a id="canonical-1133012000322023-2220321300020101-1312331022320311-2023011003031323-3203221000032230-3303221111323132-1312133202001220-2122203211120122"></a>

<a id="canonical-3122021101102112-3102122120223121-1211323102120211-0030002211212332-0203223211321011-2201012012130311-0310122123313320-1001200331022322"></a>

## namespace property — dc_cluster_group_outside_vn / 301032213102 / 5

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

<a id="canonical-0002113223211221-3223203131112101-1213001112213101-3312223121203220-3320311321003010-3120212001131332-2302200032322110-3311030023231013"></a>

<a id="canonical-1121010210210122-0201220031230010-1020113300123223-2212233310312000-0001121221201012-0012010123320301-3322112211020212-2000102111300133"></a>

## tenant property — dc_cluster_group_outside_vn / 301032213102 / 6

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

<a id="canonical-2300030213123000-3320303221232133-0201002110312011-0222122133320032-2002021020003313-3122330112012100-0222221130110120-0321123123323001"></a>

## Next pages — dc_cluster_group_outside_vn / 301032213102 / 7

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1002320332003012-1322332301331332-1321130113113230-0013121301031323-2313001313013132-2130103332211001-1002122032221131-0201233333331021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222301021322121-2012121030220230-1212101013103010-3210032312121021-1211300233002023-1213230213311010-0033133103312300-1103313120103032"></a>

## ingress_egress_gw.forward_proxy_allow_all — forward_proxy_allow_all / 211321313300 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.forward_proxy_allow_all

<a id="canonical-1023330021231200-0102330000330323-2201031231021122-2232100002013320-2302312300303221-1303232310131202-1200310111313232-0231310211231110"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for forward proxy allow all.

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

<a id="canonical-2013223021031120-1110001323223003-0221200030021332-3210112123320220-2002311201211012-0122131220233210-1022222010031131-2200332113223203"></a>

## Direct properties — forward_proxy_allow_all / 211321313300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201100311233323-0113322301102210-3113102002223133-3020132202210203-0012010022111333-1122210311111223-0103212011311203-2210110230003013"></a>

## Next pages — forward_proxy_allow_all / 211321313300 / 4

- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1332122323112132-3002221202203112-1112223122230110-3202203020123013-1011121222220200-2331333002313323-3000011132300012-2020031001021321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213111132332020-3201323303221010-2320231201220220-1321330120301003-1032222000010023-0203030320132011-3000021331110021-3123021022300101"></a>

## ingress_egress_gw.global_network_list — global_network_list / 110003212100 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- ingress_egress_gw.global_network_list

<a id="canonical-1111020122200100-2003231202312120-3213312210123202-1032200103301020-2032302011322030-1112300202322231-0121121231130212-0020300302102022"></a>

Type: `"single"`. Computed.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

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

<a id="canonical-3010031023012310-1113330322330303-2003233200330012-2000132030323113-2022033221231122-3221212201002022-0312132131010321-2320001310320313"></a>

## Direct properties — global_network_list / 110003212100 / 3

- [global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-0330221333031133-3132230213011022-3320213101201032-1131300221021200-0102200333311122-3131332110023213-3123132210231030-3120230223133202): complete subsection reference.

<a id="canonical-1223003133222232-0120212222220203-2210032222101122-2031001003013111-2031111221230301-0110200311332113-1020013300332323-2330303311033223"></a>

## Next pages — global_network_list / 110003212100 / 4

- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-0330221333031133-3132230213011022-3320213101201032-1131300221021200-0102200333311122-3131332110023213-3123132210231030-3120230223133202)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0330221333031133-3132230213011022-3320213101201032-1131300221021200-0102200333311122-3131332110023213-3123132210231030-3120230223133202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302211100333303-2133001102113223-3133112332131210-0231301310031311-3231303330331030-0033032210100221-0030123112032332-1302312130310012"></a>

## ingress_egress_gw.global_network_list.global_network_connections — global_network_connections / 131213022303 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-1332122323112132-3002221202203112-1112223122230110-3202203020123013-1011121222220200-2331333002313323-3000011132300012-2020031001021321)
- ingress_egress_gw.global_network_list.global_network_connections

<a id="canonical-0011201202332022-3301311210123331-0332113333321312-3300001233031103-0133332232032033-0202300102202200-1013111030030021-1231312100013210"></a>

Type: `"list"`. Computed.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

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
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0010201023202110-3000212202320212-3231022222202102-3203333332013032-2130323033130023-1330020123011130-3023312103133230-0101233033010020"></a>

## Direct properties — global_network_connections / 131213022303 / 3

- [sli_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-0233133200231023-3013032013200011-1112321001000313-2020000230200112-0013001223331303-1211032223001223-0130300001213321-1233101020201203): complete subsection reference.

- [slo_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-1123202000000323-1313303100030122-3311011002333321-0313100120131303-0001332330233123-2231030010032301-2311230001233113-3220303121222101): complete subsection reference.

<a id="canonical-0032003222031300-0313203113033131-3222111210023222-3210001331231031-3122210101312321-0002202312103200-0310101011032032-2111021211331313"></a>

## Next pages — global_network_connections / 131213022303 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-0233133200231023-3013032013200011-1112321001000313-2020000230200112-0013001223331303-1211032223001223-0130300001213321-1233101020201203)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-1123202000000323-1313303100030122-3311011002333321-0313100120131303-0001332330233123-2231030010032301-2311230001233113-3220303121222101)
- [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-1332122323112132-3002221202203112-1112223122230110-3202203020123013-1011121222220200-2331333002313323-3000011132300012-2020031001021321)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0233133200231023-3013032013200011-1112321001000313-2020000230200112-0013001223331303-1211032223001223-0130300001213321-1233101020201203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321222203123331-0032002132231133-1102213301210303-2213303102200033-3223322103322222-2121003133313030-0112013333231020-3223020210303332"></a>

## ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr — sli_to_global_dr / 131231321323 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-1332122323112132-3002221202203112-1112223122230110-3202203020123013-1011121222220200-2331333002313323-3000011132300012-2020031001021321)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-0330221333031133-3132230213011022-3320213101201032-1131300221021200-0102200333311122-3131332110023213-3123132210231030-3120230223133202)
- ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-1001202011001323-2023103110103013-1222223112113203-1232323303231022-1001022003112102-1023010303202222-1201321131212221-1331011010312213"></a>

Type: `"single"`. Computed.

Global network reference for direct connection.

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

<a id="canonical-1113123223001110-3223230201110323-3313220213323200-3103110221203031-2232211223103131-2311032200330031-0220031133232333-1202102131023010"></a>

## Direct properties — sli_to_global_dr / 131231321323 / 3

- [global_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-1333133302111020-2231011110323312-0011133033233301-2120320301102202-1331022030011011-2330100332231231-2131030203333012-1223103330120311): complete subsection reference.

<a id="canonical-3010333032301032-2311010132323102-0311221023103313-3123301012221120-0322213222022131-0012011302010001-0302023013303022-3201120303211221"></a>

## Next pages — sli_to_global_dr / 131231321323 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-1333133302111020-2231011110323312-0011133033233301-2120320301102202-1331022030011011-2330100332231231-2131030203333012-1223103330120311)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-0330221333031133-3132230213011022-3320213101201032-1131300221021200-0102200333311122-3131332110023213-3123132210231030-3120230223133202)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1333133302111020-2231011110323312-0011133033233301-2120320301102202-1331022030011011-2330100332231231-2131030203333012-1223103330120311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033320312113302-0221122020022300-1220112023312101-1133300131302103-2031333201213220-2002030332002211-1330113020322311-2213123300133023"></a>

## ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn — global_vn / 320220102313 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-1332122323112132-3002221202203112-1112223122230110-3202203020123013-1011121222220200-2331333002313323-3000011132300012-2020031001021321)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-0330221333031133-3132230213011022-3320213101201032-1131300221021200-0102200333311122-3131332110023213-3123132210231030-3120230223133202)
- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-0233133200231023-3013032013200011-1112321001000313-2020000230200112-0013001223331303-1211032223001223-0130300001213321-1233101020201203)
- ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-0221223303232220-1301220022022311-0331122212002010-3012111100020110-3332002001022321-2203312113213222-3023213132133131-2110231002021303"></a>

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

<a id="canonical-1102311203100121-3303112221101120-1320300233111132-2022101121023132-1021110113132311-0010022033003303-1000231201121210-3101020220222212"></a>

## Direct properties — global_vn / 320220102313 / 3

<a id="canonical-2033113131232000-2022231331110203-2132302230212220-0020312122323303-1332231112002221-3320331222211113-3233202112121100-0233231001223202"></a>

<a id="canonical-2112322130131130-1221223202031020-2321220233221211-2333310301320112-2001032200013131-1221231132322220-0332212032123130-2000321233021130"></a>

## name property — global_vn / 320220102313 / 4

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

<a id="canonical-0322103231002102-3321313112101333-3230111102221001-2220310303212010-0033021212021223-3321321022002131-0333113323302020-0112301320312001"></a>

<a id="canonical-3302212333011123-3131332100311303-2100211022301312-1331101331302302-0320311123103020-1200113122111012-0211102210033210-1222133221023330"></a>

## namespace property — global_vn / 320220102313 / 5

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

<a id="canonical-3021122133211133-0213322101231031-3003020120032002-3231220132300020-0230011230333002-2021313202210100-1012003001211013-2232011023022210"></a>

<a id="canonical-1103200120133120-3331022113131000-0032323031213113-2002021300213032-2103123332201222-0221000003122100-1121222010122220-2322030202013022"></a>

## tenant property — global_vn / 320220102313 / 6

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

<a id="canonical-2101022202102103-2120120102012132-3111010323130120-1100332022313312-1312103312032030-3211211023210020-1200330223231012-3111020230201102"></a>

## Next pages — global_vn / 320220102313 / 7

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-0233133200231023-3013032013200011-1112321001000313-2020000230200112-0013001223331303-1211032223001223-0130300001213321-1233101020201203)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-1123202000000323-1313303100030122-3311011002333321-0313100120131303-0001332330233123-2231030010032301-2311230001233113-3220303121222101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331220323301320-3013033203130313-0001033133110321-1301313112302103-2312221310233123-3122031330213123-3220302301023232-3020312001232321"></a>

## ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr — slo_to_global_dr / 203133313012 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-1332122323112132-3002221202203112-1112223122230110-3202203020123013-1011121222220200-2331333002313323-3000011132300012-2020031001021321)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-0330221333031133-3132230213011022-3320213101201032-1131300221021200-0102200333311122-3131332110023213-3123132210231030-3120230223133202)
- ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-1132131011000122-2332110130022132-0133210332310031-1033331013221213-1223323023203323-0320203211231102-2110312221312023-1133103122102201"></a>

Type: `"single"`. Computed.

Global network reference for direct connection.

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

<a id="canonical-0013311111323302-3232203201232112-0120130320330001-1201030311101231-2312220203320010-1122230300230320-1221320010111130-2221310311132311"></a>

## Direct properties — slo_to_global_dr / 203133313012 / 3

- [global_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-3233310122323322-2120203123001230-2002032123023213-0031030330132021-3013120311331230-2122112321320031-1003312032010003-2321232010303133): complete subsection reference.

<a id="canonical-1221320320121310-1323223031203212-1020331210101023-2320221320313132-1311021331013113-1002322310112323-2103021333110302-3323303310110022"></a>

## Next pages — slo_to_global_dr / 203133313012 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](data-sources--aws_vpc_site--reference--group-002.md#canonical-3233310122323322-2120203123001230-2002032123023213-0031030330132021-3013120311331230-2122112321320031-1003312032010003-2321232010303133)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-0330221333031133-3132230213011022-3320213101201032-1131300221021200-0102200333311122-3131332110023213-3123132210231030-3120230223133202)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-3233310122323322-2120203123001230-2002032123023213-0031030330132021-3013120311331230-2122112321320031-1003312032010003-2321232010303133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213030031023211-0001030032303331-3012110112321200-1322300122013331-3132220221022021-1110310022333323-1200320022311010-1302103031213103"></a>

## ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn — global_vn / 201031232012 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)
- [Property reference](data-sources--aws_vpc_site--reference--group-001.md#canonical-0022321211312211-1012321202211222-1323321322032023-2000003030132313-1113203310310201-1022202211011200-0030121203121003-0320332332121230)
- [ingress_egress_gw](data-sources--aws_vpc_site--reference--group-002.md#canonical-2132213033232212-3321223311011310-1200033213120101-1323311320233131-2220212331110010-0113220233212221-0033000032320222-1232121211332131)
- [ingress_egress_gw.global_network_list](data-sources--aws_vpc_site--reference--group-002.md#canonical-1332122323112132-3002221202203112-1112223122230110-3202203020123013-1011121222220200-2331333002313323-3000011132300012-2020031001021321)
- [ingress_egress_gw.global_network_list.global_network_connections](data-sources--aws_vpc_site--reference--group-002.md#canonical-0330221333031133-3132230213011022-3320213101201032-1131300221021200-0102200333311122-3131332110023213-3123132210231030-3120230223133202)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-1123202000000323-1313303100030122-3311011002333321-0313100120131303-0001332330233123-2231030010032301-2311230001233113-3220303121222101)
- ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-3313320113322121-1320330312310102-3232120123013221-0123030321101221-0331312031012010-0300032013123101-3211213111013132-0130112332130202"></a>

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

<a id="canonical-2000300302032001-3333303111233332-2322102000330233-3001021022120011-0033331313110300-0221231331233300-0021201200230232-1333230032201120"></a>

## Direct properties — global_vn / 201031232012 / 3

<a id="canonical-3211233322330322-3200232330320010-3133032301223102-0302321321231223-2330231123310213-1013100132123202-2321011110220223-0111333322022211"></a>

<a id="canonical-0120313330223333-3020133213310333-3333202130113210-0122102132103023-3221132321312312-2320003100120333-0232022331001132-2303212220212313"></a>

## name property — global_vn / 201031232012 / 4

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

<a id="canonical-1022300321120201-0313112112330211-1232330102201330-1030022201213112-3031112020333201-3211201323303330-0112203110300313-3002103201101233"></a>

<a id="canonical-3310022211002203-1102233132212103-1120200221210330-3020233211320320-3220020130130023-3120121113021022-1200220100213020-3322321222302133"></a>

## namespace property — global_vn / 201031232012 / 5

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

<a id="canonical-0100320133013231-0011213330130211-3300303023313031-0121002303333110-1320321330321121-1332201101000320-1223103023302102-0121312020020232"></a>

<a id="canonical-1221313301100302-3231221122012013-1030200021120113-2103131201111133-3020121001221330-3303122111031320-1323312232121201-0022122033012321"></a>

## tenant property — global_vn / 201031232012 / 6

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

<a id="canonical-1223100212111311-0023013021233023-2231212220000101-3013111122200002-3103203102100212-0031103003322201-0112130031113220-2002232331311022"></a>

## Next pages — global_vn / 201031232012 / 7

- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_vpc_site--reference--group-002.md#canonical-1123202000000323-1313303100030122-3311011002333321-0313100120131303-0001332330233123-2231030010032301-2311230001233113-3220303121222101)
- [xcsh_aws_vpc_site](../data-sources/aws_vpc_site.md#canonical-3200101001132121-0113301212212322-3331232003212322-0100300122200131-2133100121120110-1212232321223202-3330130133031301-1231331023012223)

<a id="canonical-0130311320213320-2133101303120301-1122213123302210-2100223023101230-3113012113213132-2032202222311321-2311201112200200-3003000131210000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
