---
page_title: "xcsh_securemesh_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site reference."
---

# xcsh_securemesh_site reference

<a id="canonical-3312012311331211-0232301011000030-1003231212110203-0311101202012320-0331323032000001-2210320123331111-0220300200130001-2220021111333332"></a>

## Next pages — list / 032001300000 / 5

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--reference--group-004.md#canonical-0103030001310321-1332113332321113-1032123020022013-2313202333113112-0112303002231022-1232113330310331-3033123313200020-1002002132002220)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-0220020202311120-2111010010120010-2023220300312032-3103011100311321-2323202130230330-0111313132012021-2102201212002112-2210230320100323)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0103030001310321-1332113332321113-1032123020022013-2313202333113112-0112303002231022-1232113330310331-3033123313200020-1002002132002220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132120103123332-2222031311311203-2221321301301120-3220101312220131-1112012311022003-3220301300303031-3001111220113333-1211101023102000"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface — interface / 203003033320 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [custom_network_config.slo_config.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-3223310202111231-3021221111301221-0333113110111223-3213302301023101-2201022203323023-3310013323001202-2032200212120021-1123113201131302)
- [custom_network_config.slo_config.static_routes.static_routes](data-sources--securemesh_site--reference--group-003.md#canonical-0002013202001320-2112331011323332-1120033103023231-1021303001323112-3033200021331231-2321302123110231-0030121101030031-2013000303010220)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-003.md#canonical-0220020202311120-2111010010120010-2023220300312032-3103011100311321-2323202130230330-0111313132012021-2102201212002112-2210230320100323)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-1133032110213012-1133100113111131-2312223322331200-1232010003331131-3132022323200011-1301011220003101-2010302130123313-2103111101020133)
- custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-1211212133113102-0023112010301203-1021010111002333-2313202322132321-3331302133120101-2122022323032132-3010210120101020-0323230102331123"></a>

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

<a id="canonical-1121100013303132-0311330031230003-3002223102132222-2112102201103012-2231123022102100-1000333323101222-0311113103233020-2322020022000312"></a>

## Direct properties — interface / 203003033320 / 3

<a id="canonical-2300102000201103-3310311302122211-1120301132300312-0022110231210002-0000131203000321-0222311312010110-0323011110120200-2111003022210001"></a>

<a id="canonical-3123323203031322-3212302212021201-3002203030003003-3100101020023213-1003200231130212-0110202222030200-1123333023311000-2301133023011212"></a>

## kind property — interface / 203003033320 / 4

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

<a id="canonical-0330201010332002-2121131101233310-3323301103310112-3002320033313210-0312230211000201-2003231012103231-1133032123102033-2201000102031201"></a>

<a id="canonical-1220310022121012-0222122123310331-3200121303112011-2221030101001300-1202133100132303-1121330032030033-0232332323330301-1122231002313233"></a>

## name property — interface / 203003033320 / 5

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

<a id="canonical-0320001221323213-1020213321002300-3231000210023021-3302312010232312-1220021211111000-3321103122110311-2230033020133001-2113100311313020"></a>

<a id="canonical-0220313211010123-1031232003223200-3220021311031332-1022012120320210-3001022233230213-2301132022102230-2312321220330212-2313100303332210"></a>

## namespace property — interface / 203003033320 / 6

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

<a id="canonical-3212202321020010-2301111131030120-3111302133103010-0103032023112300-1023013203023111-3023122131123021-3321313132232323-0133331002332021"></a>

<a id="canonical-3103221021113110-3210221033330020-1303311003211301-1331230100102223-2312213010223333-2231221222001310-0302223122303233-1100023202223003"></a>

## tenant property — interface / 203003033320 / 7

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

<a id="canonical-2203312211320011-3131010323030110-3320011032102223-2332022100101310-2021120023112231-0230122332323313-0231112131210113-3332321233320000"></a>

<a id="canonical-1221122120221203-1203333303112311-1302023320002113-0001133123032120-2233330101313110-1222320310201020-3313132311332231-2211322031101222"></a>

## uid property — interface / 203003033320 / 8

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

<a id="canonical-1220310310010110-1310132110123030-3100213021333333-1023023211122030-3110321322132031-3222102011323200-3312121133221002-3123223232203010"></a>

## Next pages — interface / 203003033320 / 9

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-003.md#canonical-1133032110213012-1133100113111131-2312223322331200-1232010003331131-3132022323200011-1301011220003101-2010302130123313-2103111101020133)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1002121031202323-2322130011221312-3320211122103210-1230301100013321-3102122303301211-0220020310031221-1011131300231103-3112323020303331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002333321110220-0032120010022223-1130021321000032-1032100102321232-3313202222231022-3131211212313301-2230320133001112-2000012110102023"></a>

## custom_network_config.slo_config.static_v6_routes — static_v6_routes / 232302231313 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- custom_network_config.slo_config.static_v6_routes

<a id="canonical-1212021001113201-2122021300221120-1000333120100121-2030130032022213-3213000202020201-2101013021122322-1331310002011210-0211012230310102"></a>

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

<a id="canonical-0301300311111010-1030213221220212-0302212122023010-3200310032321122-2011332203010100-0120321331030203-0303032013323120-1322232332303323"></a>

## Direct properties — static_v6_routes / 232302231313 / 3

- [static_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1230131332010111-3230303013313101-0101120022113311-2100121330202301-1230102320312003-0211321121020102-1212202031221121-2220131233123312): complete subsection reference.

<a id="canonical-2122031302310200-1130200111031112-0332131210330203-1222103002212333-0000001223303030-1030120221001323-2310303000113230-0031131013210331"></a>

## Next pages — static_v6_routes / 232302231313 / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1230131332010111-3230303013313101-0101120022113311-2100121330202301-1230102320312003-0211321121020102-1212202031221121-2220131233123312)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1230131332010111-3230303013313101-0101120022113311-2100121330202301-1230102320312003-0211321121020102-1212202031221121-2220131233123312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002323102030200-1211330220000102-0131001130213010-3313312110020300-0230322302211113-2200210031222112-0000222120232011-1002012123002203"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes — static_routes / 122233012221 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [custom_network_config.slo_config.static_v6_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1002121031202323-2322130011221312-3320211122103210-1230301100013321-3102122303301211-0220020310031221-1011131300231103-3112323020303331)
- custom_network_config.slo_config.static_v6_routes.static_routes

<a id="canonical-1212001302010000-0210013303012030-1313013113110223-2001110232013332-3200210003000111-1320322301313131-3312013333331302-1012131120122010"></a>

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

<a id="canonical-0000310110122332-3201233302233230-3030133302100322-1233133312233120-2212121012112023-0321021021012202-3013223023223031-1322122323322002"></a>

## Direct properties — static_routes / 122233012221 / 3

<a id="canonical-1302110022332131-3211230330032011-2213032303230232-2203000022103110-2031120130231012-3130011133033311-0231201000330003-3222231021201232"></a>

<a id="canonical-0321333310102001-3113203332003022-1233122311221202-1233120131132212-2303010020233313-2003331121103032-1113323311033120-3320130003321211"></a>

## attrs property — static_routes / 122233012221 / 4

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

- [default_gateway](data-sources--securemesh_site--reference--group-004.md#canonical-1233000322331030-0100223322030301-1030121100023100-0330301301331013-1122122230323222-0103120122001021-0131032113103021-1013100110333331): complete subsection reference.

<a id="canonical-3122300222121200-3112223131011333-3212301020303020-2302333111022300-2223023131202210-3223010010220111-1323100302321002-0032312222003012"></a>

<a id="canonical-0000033130203333-2230100222321103-1103113211221222-2312010000223122-2100113233230323-3203021113233103-0030132203111000-0211302332011310"></a>

## ip_address property — static_routes / 122233012221 / 5

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

<a id="canonical-2102202003021330-3330130121313323-0020030033330102-2223321233213100-1331302201120122-0121010132001112-3112331320020321-2300320332222230"></a>

<a id="canonical-3321113311220201-2333323231212013-0131101333121131-1131021001230030-1211211302122223-1212003310030013-1230133002033010-0302120323022002"></a>

## ip_prefixes property — static_routes / 122233012221 / 6

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

- [node_interface](data-sources--securemesh_site--reference--group-004.md#canonical-0331020121130021-2313000123332132-2111010000321313-2213030332020201-0011232013112032-3020212222323010-2131130113331120-2332231331203213): complete subsection reference.

<a id="canonical-3313132201301213-2020233222112122-0210121031320121-1100212110023223-3012013200122221-2123210313023132-1200033210223122-3110003023012332"></a>

## Next pages — static_routes / 122233012221 / 7

- [custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway](data-sources--securemesh_site--reference--group-004.md#canonical-1233000322331030-0100223322030301-1030121100023100-0330301301331013-1122122230323222-0103120122001021-0131032113103021-1013100110333331)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-004.md#canonical-0331020121130021-2313000123332132-2111010000321313-2213030332020201-0011232013112032-3020212222323010-2131130113331120-2332231331203213)
- [custom_network_config.slo_config.static_v6_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1002121031202323-2322130011221312-3320211122103210-1230301100013321-3102122303301211-0220020310031221-1011131300231103-3112323020303331)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1233000322331030-0100223322030301-1030121100023100-0330301301331013-1122122230323222-0103120122001021-0131032113103021-1013100110333331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011132313001220-3013333013221323-0322032311021201-0000022032102333-3320032112210331-3303332130311322-2232300301312013-2321130302331301"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway — default_gateway / 331011222020 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [custom_network_config.slo_config.static_v6_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1002121031202323-2322130011221312-3320211122103210-1230301100013321-3102122303301211-0220020310031221-1011131300231103-3112323020303331)
- [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1230131332010111-3230303013313101-0101120022113311-2100121330202301-1230102320312003-0211321121020102-1212202031221121-2220131233123312)
- custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-3022000203031331-0322233331100123-3301300012132303-0233202332133031-1220332333001003-2033202332121233-0230021230223201-3123332222233001"></a>

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

<a id="canonical-2203300312113323-3211122023312102-3323120033020322-3132203101212233-3102301320132010-1301030210031101-0333013212303033-3103011202203213"></a>

## Direct properties — default_gateway / 331011222020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321302003012101-1031012033013333-3022110333003033-1012110201332132-2020202212201023-1320030132313023-0310130130131323-2112013010001332"></a>

## Next pages — default_gateway / 331011222020 / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1230131332010111-3230303013313101-0101120022113311-2100121330202301-1230102320312003-0211321121020102-1212202031221121-2220131233123312)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0331020121130021-2313000123332132-2111010000321313-2213030332020201-0011232013112032-3020212222323010-2131130113331120-2332231331203213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112302120011133-1310003101211003-0301322121021333-2202313331221111-2012223033213203-3322301202211331-1121010022131002-1113203210033003"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface — node_interface / 003320233013 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [custom_network_config.slo_config.static_v6_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1002121031202323-2322130011221312-3320211122103210-1230301100013321-3102122303301211-0220020310031221-1011131300231103-3112323020303331)
- [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1230131332010111-3230303013313101-0101120022113311-2100121330202301-1230102320312003-0211321121020102-1212202031221121-2220131233123312)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface

<a id="canonical-3132212311001322-1132003312220233-3013032021000113-1200003003320312-0010303022110012-1032330133300302-3230123030301123-0000210330002120"></a>

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

<a id="canonical-0033000202210233-1230301333101212-2010332123031002-1211303120302302-1103120332033130-3022001111112020-1220222330033121-1330211231231200"></a>

## Direct properties — node_interface / 003320233013 / 3

- [list](data-sources--securemesh_site--reference--group-004.md#canonical-2310200033321023-1201202213213122-1120301120222220-0100102330222331-0111032020310130-2102332013111001-3211313013212023-2131101023122000): complete subsection reference.

<a id="canonical-2301331301212122-0201001131330322-1001230213231311-0201310003022223-0200103110003020-3031023231013123-1320120323010210-2031231111120203"></a>

## Next pages — node_interface / 003320233013 / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-004.md#canonical-2310200033321023-1201202213213122-1120301120222220-0100102330222331-0111032020310130-2102332013111001-3211313013212023-2131101023122000)
- [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1230131332010111-3230303013313101-0101120022113311-2100121330202301-1230102320312003-0211321121020102-1212202031221121-2220131233123312)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2310200033321023-1201202213213122-1120301120222220-0100102330222331-0111032020310130-2102332013111001-3211313013212023-2131101023122000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003002022020230-1030020131201102-2323112310223112-3020322102130300-3330021030023232-0211213102022130-0113211233112121-0321210310311221"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list — list / 130222331023 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [custom_network_config.slo_config.static_v6_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1002121031202323-2322130011221312-3320211122103210-1230301100013321-3102122303301211-0220020310031221-1011131300231103-3112323020303331)
- [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1230131332010111-3230303013313101-0101120022113311-2100121330202301-1230102320312003-0211321121020102-1212202031221121-2220131233123312)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-004.md#canonical-0331020121130021-2313000123332132-2111010000321313-2213030332020201-0011232013112032-3020212222323010-2131130113331120-2332231331203213)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-1213312003230022-1002111310010103-0120103123303020-0212323023222022-1212301202313103-3021120103012212-3300220003123131-0133103323030120"></a>

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

<a id="canonical-3102202102302021-1331203023101022-0210231110013223-3230302121312133-1210002031300223-3302013030302310-1120122330130123-2113023333322331"></a>

## Direct properties — list / 130222331023 / 3

- [interface](data-sources--securemesh_site--reference--group-004.md#canonical-1330123323031020-2230030222032003-0132111020220131-2332202301023010-0101330131020102-2101110233303012-3212010312003133-3130222221021300): complete subsection reference.

<a id="canonical-3301230333222303-2132331133123122-0131011002032103-3311332120221110-3022102203200231-1200131111312131-2322000130102311-2002223002330131"></a>

<a id="canonical-2113220211103120-1232010121133011-3011312001331111-3213033221320213-0313022320213010-1013001231332130-0020121332302102-1213200123303231"></a>

## node property — list / 130222331023 / 4

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

<a id="canonical-0332212313233310-1312002022201103-0121333130130322-2300230123310113-0120011112020212-1010302201333231-2313322232200121-1013233222000011"></a>

## Next pages — list / 130222331023 / 5

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--securemesh_site--reference--group-004.md#canonical-1330123323031020-2230030222032003-0132111020220131-2332202301023010-0101330131020102-2101110233303012-3212010312003133-3130222221021300)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-004.md#canonical-0331020121130021-2313000123332132-2111010000321313-2213030332020201-0011232013112032-3020212222323010-2131130113331120-2332231331203213)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1330123323031020-2230030222032003-0132111020220131-2332202301023010-0101330131020102-2101110233303012-3212010312003133-3130222221021300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203032202132332-1212110313310112-3121330222011212-0003133130031122-2100201221132103-1200223102300132-1222030000303021-1002321312331101"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface — interface / 122200011311 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [custom_network_config.slo_config](data-sources--securemesh_site--reference--group-003.md#canonical-2013011311332303-3210000213300110-3103113300033010-3012003112312102-3000211131012021-1200122132320323-0233200120131321-2233212120213001)
- [custom_network_config.slo_config.static_v6_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1002121031202323-2322130011221312-3320211122103210-1230301100013321-3102122303301211-0220020310031221-1011131300231103-3112323020303331)
- [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--securemesh_site--reference--group-004.md#canonical-1230131332010111-3230303013313101-0101120022113311-2100121330202301-1230102320312003-0211321121020102-1212202031221121-2220131233123312)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](data-sources--securemesh_site--reference--group-004.md#canonical-0331020121130021-2313000123332132-2111010000321313-2213030332020201-0011232013112032-3020212222323010-2131130113331120-2332231331203213)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-004.md#canonical-2310200033321023-1201202213213122-1120301120222220-0100102330222331-0111032020310130-2102332013111001-3211313013212023-2131101023122000)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-2302101220311111-1321231311230213-2122322332200133-0312301323331031-2221111233123100-1323323023033233-1320022200223033-3331331132101310"></a>

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

<a id="canonical-2001032211313303-3113321131132102-3021310000131220-2300010331031123-0120313210212200-0000110021120233-0111022110111212-2233302113103203"></a>

## Direct properties — interface / 122200011311 / 3

<a id="canonical-2223021223232303-3203120332003231-0231032133332333-0011023303031223-3030120323301022-1003020013211002-1132002210312120-2302330100213232"></a>

<a id="canonical-3202121220222302-1021313032121310-1300301331313103-2131003103133021-1130101011302300-3111002301302121-2320333103031321-1023220222221103"></a>

## kind property — interface / 122200011311 / 4

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

<a id="canonical-1003201120021112-3331223330322323-3321311220022120-1322221120220303-1201233022130311-3002200012020013-0002110201011232-1232002200232233"></a>

<a id="canonical-0132203301303131-1012333123332000-2202031213021233-0320202222222123-0303300030322201-0313001013123110-3123001110202300-0110131300022013"></a>

## name property — interface / 122200011311 / 5

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

<a id="canonical-3133303020320101-3212030112002310-2302023123031203-1230022010333032-1110023200130133-2113222020220103-0110133022203332-0120321203303112"></a>

<a id="canonical-1021002221111211-2313021110020001-1302031102111210-3311002223233330-0233020003230021-1111202333222131-2130321231032310-1032212030300212"></a>

## namespace property — interface / 122200011311 / 6

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

<a id="canonical-0110230000022231-2101222011022122-1022032310102031-3311231111210313-1010312122130111-1220021022310321-0230021003203222-3332321213201233"></a>

<a id="canonical-1012213232112212-1320122312313101-3003313011023031-3022123112130032-1123332100212112-0310031313221133-0001121201203110-2020111231131100"></a>

## tenant property — interface / 122200011311 / 7

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

<a id="canonical-3313031210110312-2021122121303121-0003102223222120-3312033122301203-2332222113110121-3033021230121121-1232230312112033-3322333330000033"></a>

<a id="canonical-2011131132000211-1001233310121333-2113320111101030-2301011200332220-1003310220202333-1211012012111012-0233232313022201-1313130102331321"></a>

## uid property — interface / 122200011311 / 8

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

<a id="canonical-1101222123211201-2302103303133022-2020001122310321-2121131203200211-0032130022210331-3102121000202033-2232133113222011-1102011102013123"></a>

## Next pages — interface / 122200011311 / 9

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](data-sources--securemesh_site--reference--group-004.md#canonical-2310200033321023-1201202213213122-1120301120222220-0100102330222331-0111032020310130-2102332013111001-3211313013212023-2131101023122000)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3202030221310100-2332312320000312-2311123323010011-3123113311220311-0110233233302131-3330002323231000-1132213111220103-0102232030213320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212211213113322-1232111332203121-1223022233312002-0013012131230213-1322302202012022-2030320131031121-0311322112233233-2031000102203133"></a>

## custom_network_config.sm_connection_public_ip — sm_connection_public_ip / 233232303311 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- custom_network_config.sm_connection_public_ip

<a id="canonical-3332101202300013-1323013311020031-1121022330000133-0132231033113030-3323301332301312-3230201120110002-0123201002112222-0330030212331003"></a>

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

<a id="canonical-3331302003032022-3333230123102213-0333330023310223-1010131323322020-0331032202311221-1002020321320100-2201220222010121-2210311100031330"></a>

## Direct properties — sm_connection_public_ip / 233232303311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321202222100211-0221030323201330-2231022330021211-2321331012110312-1021021020011113-1213313011303110-3133332132222010-1111321332322102"></a>

## Next pages — sm_connection_public_ip / 233232303311 / 4

- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2020320120002130-2233001122130003-0032113133002121-1310301132212321-2221222033132221-1323303312320211-0313232000130213-3301100230311111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332233113230112-3222010132312223-1111031322020331-1101321022210321-1002110130132301-0113031300013003-2020020333303302-3232130232302010"></a>

## custom_network_config.sm_connection_pvt_ip — sm_connection_pvt_ip / 233000022131 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- custom_network_config.sm_connection_pvt_ip

<a id="canonical-3020133301223333-2110301201020200-3003003211233011-1210031202111313-0003012210120331-2120102322321023-1112030033031222-3022323023010201"></a>

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

<a id="canonical-1303130311310332-3310203122013323-3002033213133113-3301232332031221-3031113101130230-3320303120101330-1301313112030012-1233012000220123"></a>

## Direct properties — sm_connection_pvt_ip / 233000022131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220130203130303-3221013131111321-2303002302001130-1120310111311030-1322133303300113-1303112312031003-3312021003223031-0033131310102003"></a>

## Next pages — sm_connection_pvt_ip / 233000022131 / 4

- [custom_network_config](data-sources--securemesh_site--reference--group-001.md#canonical-0111133321011113-2232230331332232-1302221323302032-0122000213231021-3103100213202302-0133021302332222-1211103033203022-2330232123211302)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1332002011330322-3112302133232021-3020303330123333-0121332303210202-2011220212000322-3013133110313113-2300320313113130-0311122221021323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130213100233103-3031013231313003-3322101033012133-3021231202321323-2301213131123023-3220000200102113-3221001111111003-2000201003123021"></a>

## default_blocked_services — default_blocked_services / 012321231221 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- default_blocked_services

<a id="canonical-0222313322110202-3222012232011223-0203132322203021-1303023113330102-3200230321333002-2300333011303130-0231222333102031-0120211113002321"></a>

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

<a id="canonical-1211221303322332-2032200321002122-0220310003321330-2011232120311113-2102300110031011-1010020303203003-1021332313211321-1023122102233201"></a>

## Direct properties — default_blocked_services / 012321231221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030011210330120-0103201201101021-2220131221130321-2102022033233013-1321033023221130-2233133231222123-0310201110122012-2231312003332032"></a>

## Next pages — default_blocked_services / 012321231221 / 4

- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1032132001130300-2232230303121130-2322233123010303-1320223100112113-0002203210303003-1010323303201032-2221222230312013-0301020013000030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111300332332231-1322221130103202-2202301101021023-0301230322120012-0333122313311130-3231333332100212-3310023112122230-0033322021331013"></a>

## default_network_config — default_network_config / 302031211313 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- default_network_config

<a id="canonical-2122210302010011-3212100213220132-2330130311001112-2011021211102302-2103113231321100-3301332130302322-2031303113101321-2302100012130111"></a>

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

<a id="canonical-2103321330302231-0223131111120213-2103332000100022-1031110320220012-1102303022321230-2213312303012023-1131220212221032-2231101013322331"></a>

## Direct properties — default_network_config / 302031211313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021030100113102-1331232323320300-1231320203312133-1120202200133202-2003301321321130-1011002303313301-2200021220300322-2201222333312100"></a>

## Next pages — default_network_config / 302031211313 / 4

- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2003312033033001-0003322123103200-0003210200302311-2323013312221021-1032032302211212-2033100222113033-3311011230311102-0221123310333312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123133302132230-1123001311021232-0010302301332221-3303020011300002-2021133212100032-1033123130211001-0113220112201002-2303331202323220"></a>

## kubernetes_upgrade_drain — kubernetes_upgrade_drain / 202130311031 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- kubernetes_upgrade_drain

<a id="canonical-0210130110302003-1323303232123313-1003012301010312-0110212302012222-3300012002232211-3113113201323101-0312211031230120-0002200102133233"></a>

Type: `"single"`. Computed.

Specify how worker nodes within a site will be upgraded.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

<a id="canonical-2001212220321303-0321111023111322-1101330120323103-3222113033123033-0001112220211101-0301030123322200-0023001023102112-2013100132202212"></a>

## Direct properties — kubernetes_upgrade_drain / 202130311031 / 3

- [disable_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-2301211112030212-0332321231123102-0201133131310332-2013321021112022-2030231200103201-1323120001322233-2021100121213031-1133220222331002): complete subsection reference.

- [enable_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-3303102020032130-3300200133333323-0212103013003010-2100011132030210-0322233033001203-0020030110313020-0002033203300032-3332200001333113): complete subsection reference.

<a id="canonical-2001330232202333-1233100133102123-3231222311323030-0023212233011233-1013030203210110-1313110123232301-1132130230222222-0011220020103010"></a>

## Next pages — kubernetes_upgrade_drain / 202130311031 / 4

- [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-2301211112030212-0332321231123102-0201133131310332-2013321021112022-2030231200103201-1323120001322233-2021100121213031-1133220222331002)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-3303102020032130-3300200133333323-0212103013003010-2100011132030210-0322233033001203-0020030110313020-0002033203300032-3332200001333113)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2301211112030212-0332321231123102-0201133131310332-2013321021112022-2030231200103201-1323120001322233-2021100121213031-1133220222331002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303211312210231-0302311212223001-1200013300320330-2000210101213122-1332002310003023-3130100333132130-1321003221023203-2331200123221033"></a>

## kubernetes_upgrade_drain.disable_upgrade_drain — disable_upgrade_drain / 111233233000 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [kubernetes_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-2003312033033001-0003322123103200-0003210200302311-2323013312221021-1032032302211212-2033100222113033-3311011230311102-0221123310333312)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-0200320320011133-1013200002032332-1203003303021130-2011103233231121-1332033333200103-1203312030322100-1121322211200332-3330222001020323"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable upgrade drain.

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

<a id="canonical-1130110230313011-2302232031300010-2030013133003110-0103003233133122-0233332212223223-1201331002212212-2110223023123333-0223121122221210"></a>

## Direct properties — disable_upgrade_drain / 111233233000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202231132222323-1103233323122010-1023011100102302-0332013210013333-1002000012021110-1102203230202212-2022013010231220-2123200021313202"></a>

## Next pages — disable_upgrade_drain / 111233233000 / 4

- [kubernetes_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-2003312033033001-0003322123103200-0003210200302311-2323013312221021-1032032302211212-2033100222113033-3311011230311102-0221123310333312)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3303102020032130-3300200133333323-0212103013003010-2100011132030210-0322233033001203-0020030110313020-0002033203300032-3332200001333113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000302102230321-1221311322201212-0220131003003212-3010001022211220-1001201213302021-0230203222202033-0221303102332132-0100300321103002"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain — enable_upgrade_drain / 101102113023 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [kubernetes_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-2003312033033001-0003322123103200-0003210200302311-2323013312221021-1032032302211212-2033100222113033-3311011230311102-0221123310333312)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-1303210300033023-0211111222111330-2110102132311233-2313133223010023-2020210203123123-0232321322322133-0222233331023232-3110323110130111"></a>

Type: `"single"`. Computed.

Specify batch upgrade settings for worker nodes within a site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-drain_max_unavailable_choice": "[\"drain_max_unavailable_node_count\", \"drain_max_unavailable_node_percentage\"]",
  "x-ves-oneof-field-vega_upgrade_mode_toggle_choice": "[\"disable_vega_upgrade_mode\",\"enable_vega_upgrade_mode\"]"
}
```

<a id="canonical-2230220011121321-1311023200123133-2120232231101113-1110300130022011-2023133302022133-1123100331233300-1232331233012002-1211233212311202"></a>

## Direct properties — enable_upgrade_drain / 101102113023 / 3

- [disable_vega_upgrade_mode](data-sources--securemesh_site--reference--group-004.md#canonical-0203021312110122-0002101010012222-1132101031310200-0002032312313011-1033223003030302-2210210002301112-2331101112321321-3200220000203012): complete subsection reference.

<a id="canonical-3113111033333320-2313232231333311-1003113213232233-1131021203312231-3133230312211003-3101001132301302-1103111220100231-3011111111221202"></a>

<a id="canonical-0112331002220221-1111002132213333-3320130133020020-2220100010221300-3112100302201313-0232031323020300-2023322131000213-1021311113202111"></a>

## drain_max_unavailable_node_count property — enable_upgrade_drain / 101102113023 / 4

Type: `"number"`. Computed.

Node Batch Size Count. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-1232011112330201-2211013133303221-2211220200030310-1322323032301311-3011133320201200-0022231123100333-0123300212231033-1223200320333202"></a>

<a id="canonical-2011030121030011-1100321332231322-3333210032102211-3001203002111312-0122221111020312-3222212100020011-1212023102202310-2323203212131201"></a>

## drain_max_unavailable_node_percentage property — enable_upgrade_drain / 101102113023 / 5

Type: `"number"`. Computed.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-1302031310102230-2120101033020211-2322111011110011-2220300002120213-2323113001231012-0313033301310112-0221123030010300-0312001322331003"></a>

<a id="canonical-3231203310023333-0021331110331113-2100031033023020-2200021210010012-2113331121132011-2123000031313020-1200322332303221-2033033012010003"></a>

## drain_node_timeout property — enable_upgrade_drain / 101102113023 / 6

Type: `"number"`. Computed.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It
is..

Upstream description:

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 900,
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
    "ves.io.schema.rules.uint32.lte": "900"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  }
}
```

- [enable_vega_upgrade_mode](data-sources--securemesh_site--reference--group-004.md#canonical-2102300201100112-0020122012320112-3000110111313200-3201313013201221-2001023223333101-1021321132133311-1202231323213013-3330322221323032): complete subsection reference.

<a id="canonical-0120222332122312-0200213230012232-1232223122302312-3010001333100122-0333221210102123-3023132012301003-0130003332102230-2112330021200303"></a>

## Next pages — enable_upgrade_drain / 101102113023 / 7

- [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--securemesh_site--reference--group-004.md#canonical-0203021312110122-0002101010012222-1132101031310200-0002032312313011-1033223003030302-2210210002301112-2331101112321321-3200220000203012)
- [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--securemesh_site--reference--group-004.md#canonical-2102300201100112-0020122012320112-3000110111313200-3201313013201221-2001023223333101-1021321132133311-1202231323213013-3330322221323032)
- [kubernetes_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-2003312033033001-0003322123103200-0003210200302311-2323013312221021-1032032302211212-2033100222113033-3311011230311102-0221123310333312)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0203021312110122-0002101010012222-1132101031310200-0002032312313011-1033223003030302-2210210002301112-2331101112321321-3200220000203012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110200320202112-2210020101323212-0230203303111230-1311311321333313-2021301131321200-0310301013003201-2100032110031233-1000010112331303"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — disable_vega_upgrade_mode / 011332302032 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [kubernetes_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-2003312033033001-0003322123103200-0003210200302311-2323013312221021-1032032302211212-2033100222113033-3311011230311102-0221123310333312)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-3303102020032130-3300200133333323-0212103013003010-2100011132030210-0322233033001203-0020030110313020-0002033203300032-3332200001333113)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-2101121203233030-3233211222200201-2333220333213111-1112212323233100-2231311330033112-2303122110300103-3210320201102233-2311030023000013"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable vega upgrade mode.

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

<a id="canonical-3220120110230030-0330103130211030-3312103020131022-1302012212020122-2111111130201022-2130012023101020-1002231212231103-0112311223030021"></a>

## Direct properties — disable_vega_upgrade_mode / 011332302032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033212012300212-1112031331202211-1200333333232123-3001010102213022-3123020001030033-0323231202021132-2003213101133021-2111203112202133"></a>

## Next pages — disable_vega_upgrade_mode / 011332302032 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-3303102020032130-3300200133333323-0212103013003010-2100011132030210-0322233033001203-0020030110313020-0002033203300032-3332200001333113)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2102300201100112-0020122012320112-3000110111313200-3201313013201221-2001023223333101-1021321132133311-1202231323213013-3330322221323032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233231332000000-2101131333031132-1132103130033023-2312302002331331-0213232230121200-2130000111232001-1011000101033002-0330012321012312"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — enable_vega_upgrade_mode / 000230232000 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [kubernetes_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-2003312033033001-0003322123103200-0003210200302311-2323013312221021-1032032302211212-2033100222113033-3311011230311102-0221123310333312)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-3303102020032130-3300200133333323-0212103013003010-2100011132030210-0322233033001203-0020030110313020-0002033203300032-3332200001333113)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-2321121233112220-1303313023012111-2002030102013122-2200120022131031-3010102013121301-3132202300030103-0222222011313312-1001101233300033"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable vega upgrade mode.

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

<a id="canonical-2003000033132231-2123203002002221-3100020032330010-3202331332302213-0110211031332120-0021220310001211-1302333112222123-1021033120212000"></a>

## Direct properties — enable_vega_upgrade_mode / 000230232000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1102030121233211-3013130110011023-1302132101113102-2112332200332330-0303233022321203-2023003333013011-3111033330300113-0113113331011033"></a>

## Next pages — enable_vega_upgrade_mode / 000230232000 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--securemesh_site--reference--group-004.md#canonical-3303102020032130-3300200133333323-0212103013003010-2100011132030210-0322233033001203-0020030110313020-0002033203300032-3332200001333113)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0211220000103000-2211130101211031-0220003100131103-2300012120232311-0013332232122023-1233120111001231-1112013003212021-3201020203022120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223113023003303-0030203101101022-3311100300100130-3201303200203200-0231333031210102-0111101313320323-3133300031330212-1223211102032322"></a>

## log_receiver — log_receiver / 110322212302 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- log_receiver

<a id="canonical-0102131320131032-2101231003111113-2031221232312101-1200101123013210-3000313023310220-2111010331000113-2203120313023102-1030202021202023"></a>

Type: `"single"`. Computed.

\[OneOf: log\_receiver, logs\_streaming\_disabled\] Type establishes a direct reference from one
object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [log_receiver](data-sources--securemesh_site--reference--group-004.md#canonical-0102131320131032-2101231003111113-2031221232312101-1200101123013210-3000313023310220-2111010331000113-2203120313023102-1030202021202023)
- [logs_streaming_disabled](data-sources--securemesh_site--reference--group-004.md#canonical-0332333310211122-2213330231333323-0120012320213233-0023230002030030-2020002103301010-3200213101010323-0233111023320031-1313211021030121)

Select alternatives according to the provider validators above.

<a id="canonical-0300233002131330-0002020110003022-2200111223223002-3103211021133030-1322113013233122-3103133010001120-2113232131333333-2121203122102333"></a>

## Direct properties — log_receiver / 110322212302 / 3

<a id="canonical-2132013330000302-1030232220023313-0231200221002311-2231322013201010-3213332003133010-1100122012101330-3213330003333101-1311102120013010"></a>

<a id="canonical-1203331011221200-2133011100331023-1023231220231320-3013121310213331-0122011221103120-3233300110103032-0101321103220112-0103323031210032"></a>

## name property — log_receiver / 110322212302 / 4

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

<a id="canonical-2213013121110112-1200202331022131-2003020231330032-3332210121001310-1300112233130230-0113110102133212-3221300110322210-2032031033303010"></a>

<a id="canonical-2311032002110031-3110101212233012-0233111230331213-0023201211031223-1233303133320131-0223123100102332-0222113331210200-1211102333131031"></a>

## namespace property — log_receiver / 110322212302 / 5

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

<a id="canonical-1301220211333132-0002202021321333-0120033212203330-1223120212010132-3211310213223011-1020220230121110-3222110131110113-2233211003313220"></a>

<a id="canonical-0230201221132232-1113021002332230-0212331001110012-2000222303220020-0221212002212301-2102112031112023-0221322312001031-2222313201233301"></a>

## tenant property — log_receiver / 110322212302 / 6

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

<a id="canonical-0230133122111213-1100113212123221-3221302311312030-0200312022201200-1131002213303002-2122310100001013-1212310322232312-3232131102133303"></a>

## Next pages — log_receiver / 110322212302 / 7

- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2030232101301321-3022120220023011-2330231031223312-0201232010301230-3330312010210010-0011010130130130-3212310310100011-2332213030110233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121222103301030-2110230333111311-1131311313323013-1312332202230033-2120032120023313-0322330222322113-2111300212010111-1322202132230112"></a>

## logs_streaming_disabled — logs_streaming_disabled / 011310301100 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- logs_streaming_disabled

<a id="canonical-0332333310211122-2213330231333323-0120012320213233-0023230002030030-2020002103301010-3200213101010323-0233111023320031-1313211021030121"></a>

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

<a id="canonical-1012000213011120-0310122110332321-0011202133210030-0313320031211213-0021332300101333-0023200132132233-2220312122213100-3101103202130100"></a>

## Direct properties — logs_streaming_disabled / 011310301100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201020222311311-2120212322020201-3212333100122123-3201110202320311-3302210013331232-3212320232210021-0103121300203131-3230013310311103"></a>

## Next pages — logs_streaming_disabled / 011310301100 / 4

- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3313100220202200-1000302100010131-1301231302121331-1300120222103101-0133123100113323-2120331223001002-2101030221002332-2323223232331200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012212031303303-0222222202321233-3232000011201203-3002023120001021-1333112211012331-0332023222231013-2131230332100001-0231020201033320"></a>

## master_node_configuration — master_node_configuration / 012321031023 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- master_node_configuration

<a id="canonical-2230020233110202-2201330322320011-0303301122202110-1022320121131032-1033112133022203-1001000213112231-1000023203102100-0023321330011213"></a>

Type: `"list"`. Computed.

Master Nodes. Configuration of master nodes.

Upstream description:

Configuration of master nodes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

<a id="canonical-1330310011311312-1232300301013200-1032132212313301-2022201212210301-2102330212031210-3012301331122111-3233012232033002-0331032330011230"></a>

## Direct properties — master_node_configuration / 012321031023 / 3

<a id="canonical-2110013133030011-1023101101020022-0031312320300221-1213102300022012-1332210301302022-2310222132003200-1001221033003030-0031031000122122"></a>

<a id="canonical-2303022100000121-3320331203302212-0100101032010201-1313302031322210-0011121303023313-1321031211133221-3211320310203232-2300231000010200"></a>

## name property — master_node_configuration / 012321031023 / 4

Type: `"string"`. Computed.

Name. Names of master node.

Upstream description:

Names of master node.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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

<a id="canonical-3202110112223130-0030231332133011-3203333301122100-1122013210201310-2310222322030321-1221131010233103-1303012132332310-2300332333021311"></a>

<a id="canonical-0131320232230322-1200331033322013-2313130231130203-0100020013132131-2323210303000023-1330311131222012-1010301220120100-0102120223211332"></a>

## public_ip property — master_node_configuration / 012321031023 / 5

Type: `"string"`. Computed.

IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.

Upstream description:

IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.

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

<a id="canonical-1333130111313203-3330313321101010-3210332123001203-2313132303331010-1111131130222221-0133103332111113-2213202111332102-3302030010323102"></a>

## Next pages — master_node_configuration / 012321031023 / 6

- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1003113100300311-2222313300001010-2021310302002030-1112332132120233-1002033303013130-2231330222100030-0102321101101202-1321023010332110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233002303231123-0033321320331211-0203203131001022-1202132331131310-1003120222233013-0103323331100123-2023023113201330-3133132303010111"></a>

## no_bond_devices — no_bond_devices / 122201311133 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- no_bond_devices

<a id="canonical-2213003202010300-2102012200213123-2322022103020000-1020020123123221-1230331011321230-0331220100231301-0230110030232233-3132321222002122"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no bond devices.

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

<a id="canonical-0213210001213302-2301333133113120-2230023002120003-2132332332113211-1323120130332302-0320032320132230-2202020233220113-1313012033320122"></a>

## Direct properties — no_bond_devices / 122201311133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310123303131320-3103213100120023-0211311202311010-0313303012130023-2130000230122110-3300012111322001-3020213232033131-0011301312202130"></a>

## Next pages — no_bond_devices / 122201311133 / 4

- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2020103213000102-1110301232330123-2322321023321200-2131023113310313-2233322222100111-2032312323132100-0321122012303013-0320311231012313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332100120023333-1012212032122331-3323022222113021-1310003313103312-1131120021313230-0210201012033120-2113210211102022-1212223301121111"></a>

## offline_survivability_mode — offline_survivability_mode / 032302310111 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- offline_survivability_mode

<a id="canonical-2330312322333212-3232020001213131-2111111121000303-2133200333233303-3112331203023332-3112300310000302-3102001112330030-2010132100312213"></a>

Type: `"single"`. Computed.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7..

Upstream description:

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-offline_survivability_mode_choice": "[\"enable_offline_survivability_mode\",\"no_offline_survivability_mode\"]"
}
```

<a id="canonical-3332313131302322-1202320003033103-1331111310120231-3100003211310023-2121320013100012-2020221120021102-1212133023311322-1031131322311202"></a>

## Direct properties — offline_survivability_mode / 032302310111 / 3

- [enable_offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-1223131002232201-2021020233002223-3132311323201122-3332231321300112-1133002331332013-3232332023120300-0103120330131222-2302023333131100): complete subsection reference.

- [no_offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-3211123210223133-2323001233113230-2112010331211222-0121221003121233-1113123110222033-0020130002010132-3113303132003100-2023103030011110): complete subsection reference.

<a id="canonical-1332003322031123-3132120232213011-3113312233023020-3103123300032023-0033331300203002-1022032311023120-3310211232301010-0032212310133332"></a>

## Next pages — offline_survivability_mode / 032302310111 / 4

- [offline_survivability_mode.enable_offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-1223131002232201-2021020233002223-3132311323201122-3332231321300112-1133002331332013-3232332023120300-0103120330131222-2302023333131100)
- [offline_survivability_mode.no_offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-3211123210223133-2323001233113230-2112010331211222-0121221003121233-1113123110222033-0020130002010132-3113303132003100-2023103030011110)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1223131002232201-2021020233002223-3132311323201122-3332231321300112-1133002331332013-3232332023120300-0103120330131222-2302023333131100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220201002121002-3222102100113113-2122021012123230-3011300210102300-0123112230311020-2231131100303311-1033212222322132-3223232232023121"></a>

## offline_survivability_mode.enable_offline_survivability_mode — enable_offline_survivability_mode / 321330010132 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-2020103213000102-1110301232330123-2322321023321200-2131023113310313-2233322222100111-2032312323132100-0321122012303013-0320311231012313)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-3122230113320022-2313310211310210-2203002021223322-2010303101010301-2320303003212113-0333332322102300-1211031132013321-2221201320220020"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable offline survivability mode.

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

<a id="canonical-3123233320211221-3331121021010301-0200300102031303-2213231210033010-0131310113301033-3102310320103131-3120323300233322-0031211323103321"></a>

## Direct properties — enable_offline_survivability_mode / 321330010132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231222010222322-3022010221121132-0201210020102300-2131033001112130-1000330023210001-0003002011021211-0113311121033212-2201021012211003"></a>

## Next pages — enable_offline_survivability_mode / 321330010132 / 4

- [offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-2020103213000102-1110301232330123-2322321023321200-2131023113310313-2233322222100111-2032312323132100-0321122012303013-0320311231012313)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3211123210223133-2323001233113230-2112010331211222-0121221003121233-1113123110222033-0020130002010132-3113303132003100-2023103030011110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123103321313033-0132130001222130-3223300113003130-2010210301312321-3101200310303313-2101311132202120-3001332003021313-0311211221313023"></a>

## offline_survivability_mode.no_offline_survivability_mode — no_offline_survivability_mode / 122021122310 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-2020103213000102-1110301232330123-2322321023321200-2131023113310313-2233322222100111-2032312323132100-0321122012303013-0320311231012313)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-1231133110100020-2211231012200221-1013122123211120-0203330231230011-2303003123010013-0102101310200101-1122130031013231-2203121333222032"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no offline survivability mode.

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

<a id="canonical-2311333320030232-2030010322032310-3320130121202211-2011121013031212-0122233223333201-3223022132021200-1221021021223301-1230200211131132"></a>

## Direct properties — no_offline_survivability_mode / 122021122310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300320131101223-0103222232022122-1233201032230203-1333331333003110-1002113103033230-1211021222012200-0311232132131011-3000122121031310"></a>

## Next pages — no_offline_survivability_mode / 122021122310 / 4

- [offline_survivability_mode](data-sources--securemesh_site--reference--group-004.md#canonical-2020103213000102-1110301232330123-2322321023321200-2131023113310313-2233322222100111-2032312323132100-0321122012303013-0320311231012313)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-3012121031011130-1233001032220331-3130122211200032-2131001233320133-0031212331013332-2032112211110320-0203323010102313-1032213120331101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031111113210200-0212121130233323-3313130122102221-1230203213013200-1301211301233132-0110011303131312-3023021113312122-2311033123311030"></a>

## os — os / 311112131211 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- os

<a id="canonical-0330232302010030-3230130021221032-0310020122222321-1301301322103031-0313211331012012-1313002103013332-3111030003211221-2331333002303002"></a>

Type: `"single"`. Computed.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Upstream description:

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]"
}
```

<a id="canonical-1022231031110310-2030313102320232-3010322120313230-1001010311000000-2103200033121332-2200301013101021-2100311213111120-3200210201000202"></a>

## Direct properties — os / 311112131211 / 3

- [default_os_version](data-sources--securemesh_site--reference--group-004.md#canonical-1021111102010130-1030332033211012-0032002022032230-0300202210332223-0123330113000131-3231113021213333-0033203100310113-3211020123112231): complete subsection reference.

<a id="canonical-1312231230132320-3202210103230302-0100232030221131-2232320102123000-1302312201110233-1232100022333320-0310232121300221-0213112301223133"></a>

<a id="canonical-3111221131012303-2132030313132303-3131212333111103-3203220102122331-3233001213130320-1112201012133212-3311303120011222-1200203021103010"></a>

## operating_system_version property — os / 311112131211 / 4

Type: `"string"`. Computed.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Upstream description:

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-1210213102020300-0112333330213001-3310323001012213-0113303213131213-3312312220122032-0232311113333222-1202233001020030-3210122333223101"></a>

## Next pages — os / 311112131211 / 5

- [os.default_os_version](data-sources--securemesh_site--reference--group-004.md#canonical-1021111102010130-1030332033211012-0032002022032230-0300202210332223-0123330113000131-3231113021213333-0033203100310113-3211020123112231)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1021111102010130-1030332033211012-0032002022032230-0300202210332223-0123330113000131-3231113021213333-0033203100310113-3211020123112231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210030232320313-0331101120312320-2222232031213012-1232303203010313-0310020233021203-1120311333033102-2312013120222011-3113120031030110"></a>

## os.default_os_version — default_os_version / 020203023333 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [os](data-sources--securemesh_site--reference--group-004.md#canonical-3012121031011130-1233001032220331-3130122211200032-2131001233320133-0031212331013332-2032112211110320-0203323010102313-1032213120331101)
- os.default_os_version

<a id="canonical-1122011120133032-1313020121210133-3021011022202333-3212221120001221-3312021202301211-3030213222221122-0113312131102123-0232201112011123"></a>

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

<a id="canonical-0031300221300022-2221022022312331-0322111123212101-3322033330322111-1103011331310000-1002113103323133-2013001010332312-1222033132100002"></a>

## Direct properties — default_os_version / 020203023333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111020122010220-3103012223222301-0003001101300220-1131333301033133-1123210312102212-2223022111010012-1203222121030313-2010120020030030"></a>

## Next pages — default_os_version / 020203023333 / 4

- [os](data-sources--securemesh_site--reference--group-004.md#canonical-3012121031011130-1233001032220331-3130122211200032-2131001233320133-0031212331013332-2032112211110320-0203323010102313-1032213120331101)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0112033030112333-1120313231222032-3130221232312031-2030202013212132-1333301113032000-2010232003132013-0212000321202122-0303200101330331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302212213031102-0210220230111233-1303022321322112-3203130313310001-1211022331023223-1200231131011210-3311302211232210-1103001213303013"></a>

## performance_enhancement_mode — performance_enhancement_mode / 201013111213 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- performance_enhancement_mode

<a id="canonical-2220211101203300-1331123022300032-3023223032213031-1300220112003022-1021132202112331-2103202312112321-1312033132112010-2033330103020020"></a>

Type: `"single"`. Computed.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

<a id="canonical-1312321220333100-2101313231132311-1023311111222002-3211213213303101-0020303312333220-0132021322110203-1231013113000302-3233121310301201"></a>

## Direct properties — performance_enhancement_mode / 201013111213 / 3

- [perf_mode_l3_enhanced](data-sources--securemesh_site--reference--group-004.md#canonical-0331031320320220-2213233310302131-3302110022011211-0123001323002023-0120121013111110-2111321031321213-3102100213021211-1033103213332331): complete subsection reference.

- [perf_mode_l7_enhanced](data-sources--securemesh_site--reference--group-004.md#canonical-1220201103233121-1030112023330001-2010012032003120-3321122032230230-1010000301233113-2201310011333011-0000012231300311-3210300312333013): complete subsection reference.

<a id="canonical-3213022203312222-2123133223231030-0232320210013203-1130120130111012-1011223111322210-0333100313133332-2312102322020211-2213112001003233"></a>

## Next pages — performance_enhancement_mode / 201013111213 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--securemesh_site--reference--group-004.md#canonical-0331031320320220-2213233310302131-3302110022011211-0123001323002023-0120121013111110-2111321031321213-3102100213021211-1033103213332331)
- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--securemesh_site--reference--group-004.md#canonical-1220201103233121-1030112023330001-2010012032003120-3321122032230230-1010000301233113-2201310011333011-0000012231300311-3210300312333013)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0331031320320220-2213233310302131-3302110022011211-0123001323002023-0120121013111110-2111321031321213-3102100213021211-1033103213332331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232302020002122-0233002201130023-2121123102211001-3310231220022322-3033010131232310-2123110123020112-1110222122303133-1302033303010100"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced — perf_mode_l3_enhanced / 311321032012 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [performance_enhancement_mode](data-sources--securemesh_site--reference--group-004.md#canonical-0112033030112333-1120313231222032-3130221232312031-2030202013212132-1333301113032000-2010232003132013-0212000321202122-0303200101330331)
- performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-2032111302222332-2032030030200300-3232230200222223-3301222311012332-3013331233321212-1213113010023311-3201332023002123-3023333132313030"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

<a id="canonical-1120220111223300-3212332200230320-0120323232211212-0212301012200012-0323321000130133-2210102133200120-2103000321032033-3220301301112323"></a>

## Direct properties — perf_mode_l3_enhanced / 311321032012 / 3

- [jumbo](data-sources--securemesh_site--reference--group-004.md#canonical-0013302311332320-2133322230130311-0211330112001103-3120113113300330-1310211120123323-3220120301330100-0311231223232013-2100033101332223): complete subsection reference.

- [no_jumbo](data-sources--securemesh_site--reference--group-004.md#canonical-2121310201220323-2001020220231201-0213303230221011-0221033311202221-0333321130122002-1230202011321321-0311323032210331-2021030302120123): complete subsection reference.

<a id="canonical-0113310322200201-2230311230111200-0200332310002232-3331102200033310-3121332000121210-2133220312100023-2222320200220121-3300232000133003"></a>

## Next pages — perf_mode_l3_enhanced / 311321032012 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--securemesh_site--reference--group-004.md#canonical-0013302311332320-2133322230130311-0211330112001103-3120113113300330-1310211120123323-3220120301330100-0311231223232013-2100033101332223)
- [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--securemesh_site--reference--group-004.md#canonical-2121310201220323-2001020220231201-0213303230221011-0221033311202221-0333321130122002-1230202011321321-0311323032210331-2021030302120123)
- [performance_enhancement_mode](data-sources--securemesh_site--reference--group-004.md#canonical-0112033030112333-1120313231222032-3130221232312031-2030202013212132-1333301113032000-2010232003132013-0212000321202122-0303200101330331)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0013302311332320-2133322230130311-0211330112001103-3120113113300330-1310211120123323-3220120301330100-0311231223232013-2100033101332223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332300130000031-3001231132003010-3102311320211130-3020301230302023-2033332303012100-3330001111231102-2030033303221101-0002001203120102"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — jumbo / 000321111222 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [performance_enhancement_mode](data-sources--securemesh_site--reference--group-004.md#canonical-0112033030112333-1120313231222032-3130221232312031-2030202013212132-1333301113032000-2010232003132013-0212000321202122-0303200101330331)
- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--securemesh_site--reference--group-004.md#canonical-0331031320320220-2213233310302131-3302110022011211-0123001323002023-0120121013111110-2111321031321213-3102100213021211-1033103213332331)
- performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-3102130121131230-2102010201021121-1233223011312322-2321023033300223-3012203310310130-1210030030313120-1012332233021311-1202202201120120"></a>

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

<a id="canonical-1223212322112032-1203123311130022-2033020022330031-3301012110121120-2012330323112012-1320222212312010-0022132113302122-3101220002220012"></a>

## Direct properties — jumbo / 000321111222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102021103322000-2320031133303222-2212101322031312-1011022003212110-1330133332223031-1113002123101133-2111322223220301-0313001312032213"></a>

## Next pages — jumbo / 000321111222 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--securemesh_site--reference--group-004.md#canonical-0331031320320220-2213233310302131-3302110022011211-0123001323002023-0120121013111110-2111321031321213-3102100213021211-1033103213332331)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2121310201220323-2001020220231201-0213303230221011-0221033311202221-0333321130122002-1230202011321321-0311323032210331-2021030302120123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133013203332132-1201033123122130-1320120312220332-3223332202131322-1213211311102112-3101203321000313-3030331102120220-1102312333302012"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — no_jumbo / 003003011001 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [performance_enhancement_mode](data-sources--securemesh_site--reference--group-004.md#canonical-0112033030112333-1120313231222032-3130221232312031-2030202013212132-1333301113032000-2010232003132013-0212000321202122-0303200101330331)
- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--securemesh_site--reference--group-004.md#canonical-0331031320320220-2213233310302131-3302110022011211-0123001323002023-0120121013111110-2111321031321213-3102100213021211-1033103213332331)
- performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-2031200302011233-1311133202120102-1303312300230010-2003002303111032-3300320132000303-2013010223122310-3130100112303211-1223121211111331"></a>

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

<a id="canonical-2001323103130112-0232300301031200-2203003300223301-3221000223113220-2212031112312123-2210332132202010-0300011110220133-3230010321022310"></a>

## Direct properties — no_jumbo / 003003011001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132301122311212-0312213331131220-3233133232010220-3130230031021211-1232123330110333-1032030033312101-3103331122120332-2200021000120013"></a>

## Next pages — no_jumbo / 003003011001 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--securemesh_site--reference--group-004.md#canonical-0331031320320220-2213233310302131-3302110022011211-0123001323002023-0120121013111110-2111321031321213-3102100213021211-1033103213332331)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1220201103233121-1030112023330001-2010012032003120-3321122032230230-1010000301233113-2201310011333011-0000012231300311-3210300312333013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111232210111021-0003213211103332-2133322210300011-1021032221010011-2203323122210321-1033232003331211-3212011133021020-2223032120113222"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced — perf_mode_l7_enhanced / 213012010321 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [performance_enhancement_mode](data-sources--securemesh_site--reference--group-004.md#canonical-0112033030112333-1120313231222032-3130221232312031-2030202013212132-1333301113032000-2010232003132013-0212000321202122-0303200101330331)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-1332200111030101-3110330210101333-2322321022212310-1313130330310013-0300113123221332-1322031023102322-0323213123233003-3302312022210113"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

<a id="canonical-1223323021313232-3220030022101300-1311310211111312-2120111233231223-3133200021323223-0002231002120220-2011002121010230-2322202121110113"></a>

## Direct properties — perf_mode_l7_enhanced / 213012010321 / 3

- [jumbo_disabled](data-sources--securemesh_site--reference--group-004.md#canonical-0113002013210130-3202011031022003-2031223221112001-0231231331203302-0002103320023330-3221321111331320-0230000132112121-1313001032221300): complete subsection reference.

- [jumbo_enabled](data-sources--securemesh_site--reference--group-004.md#canonical-2130031112121132-0102232311333303-1110311211212120-0130231230131321-2202223230231120-2330332131023103-0220203312110222-3011102122331201): complete subsection reference.

<a id="canonical-0211120023122110-2212111001203322-1023103213002120-0211232132111120-0220123113302032-3013200030313323-3133111123233031-0021231312223113"></a>

## Next pages — perf_mode_l7_enhanced / 213012010321 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--securemesh_site--reference--group-004.md#canonical-0113002013210130-3202011031022003-2031223221112001-0231231331203302-0002103320023330-3221321111331320-0230000132112121-1313001032221300)
- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--securemesh_site--reference--group-004.md#canonical-2130031112121132-0102232311333303-1110311211212120-0130231230131321-2202223230231120-2330332131023103-0220203312110222-3011102122331201)
- [performance_enhancement_mode](data-sources--securemesh_site--reference--group-004.md#canonical-0112033030112333-1120313231222032-3130221232312031-2030202013212132-1333301113032000-2010232003132013-0212000321202122-0303200101330331)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0113002013210130-3202011031022003-2031223221112001-0231231331203302-0002103320023330-3221321111331320-0230000132112121-1313001032221300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231001120002012-3000220001023032-2111033033210132-0333221220312011-1303310301012011-2311320120311030-3001220211110200-0021230023031311"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — jumbo_disabled / 133322000202 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [performance_enhancement_mode](data-sources--securemesh_site--reference--group-004.md#canonical-0112033030112333-1120313231222032-3130221232312031-2030202013212132-1333301113032000-2010232003132013-0212000321202122-0303200101330331)
- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--securemesh_site--reference--group-004.md#canonical-1220201103233121-1030112023330001-2010012032003120-3321122032230230-1010000301233113-2201310011333011-0000012231300311-3210300312333013)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-3231222133032010-3311120001311132-2331212131312333-1021331020330201-2112330100230130-0100130212333220-1231001320212332-0011131212102321"></a>

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

<a id="canonical-3300203220011120-1001311312021132-0232221013022201-3131131130202210-0322310221132121-1322132302021312-0111203030331213-2120222303322013"></a>

## Direct properties — jumbo_disabled / 133322000202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333011121000312-1012200313123200-2002232221022230-2222200113001213-0010313213133113-0202131230222311-0100330312103132-2033231300303320"></a>

## Next pages — jumbo_disabled / 133322000202 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--securemesh_site--reference--group-004.md#canonical-1220201103233121-1030112023330001-2010012032003120-3321122032230230-1010000301233113-2201310011333011-0000012231300311-3210300312333013)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2130031112121132-0102232311333303-1110311211212120-0130231230131321-2202223230231120-2330332131023103-0220203312110222-3011102122331201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001200311312302-0220331210033023-3033020120133112-3301030122312100-2231230233312331-1130323111132213-1121212212200330-3012021221033201"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — jumbo_enabled / 332121013313 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [performance_enhancement_mode](data-sources--securemesh_site--reference--group-004.md#canonical-0112033030112333-1120313231222032-3130221232312031-2030202013212132-1333301113032000-2010232003132013-0212000321202122-0303200101330331)
- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--securemesh_site--reference--group-004.md#canonical-1220201103233121-1030112023330001-2010012032003120-3321122032230230-1010000301233113-2201310011333011-0000012231300311-3210300312333013)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-1113021133323323-0013333110230331-1332003111302233-3020113321023003-1230312001332021-0030203032332031-2312223313001223-2022033133231210"></a>

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

<a id="canonical-2010010100230021-0023102323033012-3330220211321303-3010231223110322-3312021011122111-2000010331120232-0333001312221020-1221203301113302"></a>

## Direct properties — jumbo_enabled / 332121013313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210013212113211-1000310321012210-0102302233301022-2223130300221133-3031331100020332-3230102131301123-1331300303330033-0003222102330132"></a>

## Next pages — jumbo_enabled / 332121013313 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--securemesh_site--reference--group-004.md#canonical-1220201103233121-1030112023330001-2010012032003120-3321122032230230-1010000301233113-2201310011333011-0000012231300311-3210300312333013)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1101033231030033-1231112223010232-0121200330001132-1331111222212213-3023302310111332-0331233101000132-0230010020103302-3320133222123131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200101100322212-1031131121212002-0000110020323130-2330312033013220-2212003210212002-1232313130201200-0332321112311313-1232132121030033"></a>

## sw — sw / 301000131200 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- sw

<a id="canonical-0031120033012022-2031213033001211-2103110033322020-2310001210120001-3201013202113103-0102313321033031-1303113132330001-2221201130201031"></a>

Type: `"single"`. Computed.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Upstream description:

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

<a id="canonical-3103133100303210-0132202212101322-1320232020211001-3232110001021321-1201201223130313-1322301331301113-1032223111131313-1203211222120001"></a>

## Direct properties — sw / 301000131200 / 3

- [default_sw_version](data-sources--securemesh_site--reference--group-004.md#canonical-1121023013321203-0211312023202203-0331121120232130-3123111222103110-3313112211323220-1122100330111201-3123021010113002-3000213132333021): complete subsection reference.

<a id="canonical-3112113222112330-3131123000331302-2113101132230321-2030023333030123-2033233230100202-2121003230211133-1023202001023123-1033331210201321"></a>

<a id="canonical-3312031102211301-0301012022232011-0112230301330220-3222203123323101-0023011333232131-1210033002311113-0113032111331011-0230010330323031"></a>

## volterra_software_version property — sw / 301000131200 / 4

Type: `"string"`. Computed.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-3321133220133032-0323330113013230-2112320001101021-2002201300100133-1001230102130102-0200031101101333-1300331202120120-0310003030332133"></a>

## Next pages — sw / 301000131200 / 5

- [sw.default_sw_version](data-sources--securemesh_site--reference--group-004.md#canonical-1121023013321203-0211312023202203-0331121120232130-3123111222103110-3313112211323220-1122100330111201-3123021010113002-3000213132333021)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-1121023013321203-0211312023202203-0331121120232130-3123111222103110-3313112211323220-1122100330111201-3123021010113002-3000213132333021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113032012012020-2121322101111113-1112032220103323-3300103302311312-3310011030212003-2032322031022312-2031121100123002-0311310300132303"></a>

## sw.default_sw_version — default_sw_version / 113201320022 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [sw](data-sources--securemesh_site--reference--group-004.md#canonical-1101033231030033-1231112223010232-0121200330001132-1331111222212213-3023302310111332-0331233101000132-0230010020103302-3320133222123131)
- sw.default_sw_version

<a id="canonical-3210231311203130-3231111231311130-0102211200322000-1311121221011211-2300202303001302-1020122202023221-3221110302231103-1320131031101031"></a>

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

<a id="canonical-1102222211230121-0332312012001333-2322331022220122-2331022312310303-1010111220021300-0230121220331322-3330011201112222-1201332012222000"></a>

## Direct properties — default_sw_version / 113201320022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022013222012020-0121333132311100-1031121232302132-1203131121220033-1322030011321303-1223202130201233-0111332302213030-0121202322001100"></a>

## Next pages — default_sw_version / 113201320022 / 4

- [sw](data-sources--securemesh_site--reference--group-004.md#canonical-1101033231030033-1231112223010232-0121200330001132-1331111222212213-3023302310111332-0331233101000132-0230010020103302-3320133222123131)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0301131103113000-3303232313310330-0122211113001113-3130222002332001-1310200111021022-2330131102030223-3102231232233231-3301133132020110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202120301210111-1102103301233001-2023212323312002-0100002233010303-3231130210033110-3103231233210120-1313200212331312-2302323102203010"></a>

## waf_signatures — waf_signatures / 010322012031 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- waf_signatures

<a id="canonical-2113100111102311-1232323133020231-0002221002002112-3312133210103211-3221032113022331-1121312103302102-1130023311203103-1003312033313130"></a>

Type: `"single"`. Computed.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

<a id="canonical-3232022221311023-0210212132221201-1123302002103220-0223021212120202-0331131311021312-1102320021122120-3200122123132310-3311220310121130"></a>

## Direct properties — waf_signatures / 010322012031 / 3

- [automatic](data-sources--securemesh_site--reference--group-004.md#canonical-2021030302010101-0120102122310303-0310013322010102-2320230333333123-3103002012211312-0202012132002101-2120302222213222-1213211021011210): complete subsection reference.

- [manual](data-sources--securemesh_site--reference--group-004.md#canonical-0300323202223221-1003203323033111-1233012331103121-2132010013032112-1233121202022221-3110020211233113-3303220201302000-0102201010330101): complete subsection reference.

<a id="canonical-1003133330022333-1033020110320122-0012330312020221-2023001010011001-2120011300121321-1103120101000332-2110000311112021-2102032031103333"></a>

## Next pages — waf_signatures / 010322012031 / 4

- [waf_signatures.automatic](data-sources--securemesh_site--reference--group-004.md#canonical-2021030302010101-0120102122310303-0310013322010102-2320230333333123-3103002012211312-0202012132002101-2120302222213222-1213211021011210)
- [waf_signatures.manual](data-sources--securemesh_site--reference--group-004.md#canonical-0300323202223221-1003203323033111-1233012331103121-2132010013032112-1233121202022221-3110020211233113-3303220201302000-0102201010330101)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-2021030302010101-0120102122310303-0310013322010102-2320230333333123-3103002012211312-0202012132002101-2120302222213222-1213211021011210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232221233222213-1131112301322210-0231023311200200-1312002210232332-2202203322212200-0203212103202012-3213302011103332-2223232302201223"></a>

## waf_signatures.automatic — automatic / 132301012302 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [waf_signatures](data-sources--securemesh_site--reference--group-004.md#canonical-0301131103113000-3303232313310330-0122211113001113-3130222002332001-1310200111021022-2330131102030223-3102231232233231-3301133132020110)
- waf_signatures.automatic

<a id="canonical-3123320331231112-0021013113100132-2000323020112012-0020133012031110-1233313003331030-2321322203021022-1212102020320211-0203200300133331"></a>

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

<a id="canonical-2312323320010010-2320232212103133-0301303331221203-3213203210210331-0131313202113322-0020003310332010-0012202300320112-0001230221332030"></a>

## Direct properties — automatic / 132301012302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221102201231320-3231210303323123-3020101131021330-1311303030020112-1110021212230031-0210030110232113-1101222220301303-0221120112222100"></a>

## Next pages — automatic / 132301012302 / 4

- [waf_signatures](data-sources--securemesh_site--reference--group-004.md#canonical-0301131103113000-3303232313310330-0122211113001113-3130222002332001-1310200111021022-2330131102030223-3102231232233231-3301133132020110)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)

<a id="canonical-0300323202223221-1003203323033111-1233012331103121-2132010013032112-1233121202022221-3110020211233113-3303220201302000-0102201010330101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320111111221023-1202211002233300-1202112011001001-3332033300320311-3322213021321313-2103201301001002-2003332202232310-3221023230033120"></a>

## waf_signatures.manual — manual / 031330030120 / 2

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
- [Property reference](data-sources--securemesh_site--reference--group-001.md#canonical-1300330223201312-3320301232310302-2111101121312332-0100322013010302-1020322022330133-3233231332101200-1012210030333230-3121202021111102)
- [waf_signatures](data-sources--securemesh_site--reference--group-004.md#canonical-0301131103113000-3303232313310330-0122211113001113-3130222002332001-1310200111021022-2330131102030223-3102231232233231-3301133132020110)
- waf_signatures.manual

<a id="canonical-0031200001221223-3322031022111201-2112110313030123-1300120012202210-0103222301302202-3133012233033112-2023002130112320-3202011230111201"></a>

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

<a id="canonical-1103012233233022-2213221300331033-2013003031312000-2103223232331021-0313022101233320-3003333321010102-3101122111312311-1123320020312113"></a>

## Direct properties — manual / 031330030120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322331222213112-1100123133100202-1020130311000312-1213321000013332-0203322111101322-1002230100111031-0310221203110330-2112132321322232"></a>

## Next pages — manual / 031330030120 / 4

- [waf_signatures](data-sources--securemesh_site--reference--group-004.md#canonical-0301131103113000-3303232313310330-0122211113001113-3130222002332001-1310200111021022-2330131102030223-3102231232233231-3301133132020110)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md#canonical-1133210123012303-1000103211120131-3120331012033120-2003030111122302-2311132300210120-0232100202010212-3332001312220202-1211011112303230)
