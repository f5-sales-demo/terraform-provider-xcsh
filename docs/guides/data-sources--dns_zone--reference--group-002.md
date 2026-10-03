---
page_title: "xcsh_dns_zone reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone reference."
---

# xcsh_dns_zone reference

<a id="canonical-3113100112120133-0300203201023121-0032331032121122-2301210320103222-2103133011113100-2122102313213033-1310222302222102-3001102330231221"></a>

## Next pages — values / 320132011321 / 6

- [primary.default_rr_set_group.ds_record.values.sha1_digest](data-sources--dns_zone--reference--group-002.md#canonical-2330211211113323-0133320000331330-3212111101203123-1220330011232233-1300320223102222-3211311200130210-1032213010030230-1222030202212313)
- [primary.default_rr_set_group.ds_record.values.sha256_digest](data-sources--dns_zone--reference--group-002.md#canonical-1020002003030032-1130201032112031-3010331300321220-3320103121300300-0331321313303330-1310231231013013-1320221110333110-2202021210002033)
- [primary.default_rr_set_group.ds_record.values.sha384_digest](data-sources--dns_zone--reference--group-002.md#canonical-2022232333211100-0121202021103121-2302000333012022-1232200321013202-0201321121130123-3310311330202123-2022000220113232-0111002333132313)
- [primary.default_rr_set_group.ds_record](data-sources--dns_zone--reference--group-001.md#canonical-0332332020330012-1003302232123021-0022232311333132-0113122113322112-2123033321110312-3133333133301221-0121330302222103-3212000033322100)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2330211211113323-0133320000331330-3212111101203123-1220330011232233-1300320223102222-3211311200130210-1032213010030230-1222030202212313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211210023030002-0001323210232221-0022113122311023-0311111011300330-0132030102113333-1033011321012203-1023200121323210-2102330033330213"></a>

## primary.default_rr_set_group.ds_record.values.sha1_digest — sha1_digest / 002203110332 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.ds_record](data-sources--dns_zone--reference--group-001.md#canonical-0332332020330012-1003302232123021-0022232311333132-0113122113322112-2123033321110312-3133333133301221-0121330302222103-3212000033322100)
- [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1133333130110222-2223102021122311-2113332033031230-1100201102003302-3122213102212223-1032103101022120-2000002000223200-2023321023200100)
- primary.default_rr_set_group.ds_record.values.sha1_digest

<a id="canonical-1223112030310123-0221210302013131-1302031020223003-0232102210013120-0312301030233230-3011132002111230-2011112321221331-2031230100020003"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 digest.

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

<a id="canonical-0221223221332313-3200123103002112-2123211321133233-3031302333033133-3300020200223300-2010100013110000-1030000212220120-0223232110103310"></a>

## Direct properties — sha1_digest / 002203110332 / 3

<a id="canonical-2023121103322110-3103132303130122-2323231213010203-0220231130100102-3101121223031323-3230022000013212-3230211311003023-0133203133102013"></a>

<a id="canonical-3231322320301203-0203121033233311-2313232202112202-2232003313030231-0302103001311203-1221323203222101-3222211032123130-1033202000312020"></a>

## digest property — sha1_digest / 002203110332 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 40
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-1211121323310331-2011201002010103-1301211323013223-1333032033220231-3211331220112220-2213022123130122-3301332123233123-2321223221223010"></a>

## Next pages — sha1_digest / 002203110332 / 5

- [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1133333130110222-2223102021122311-2113332033031230-1100201102003302-3122213102212223-1032103101022120-2000002000223200-2023321023200100)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1020002003030032-1130201032112031-3010331300321220-3320103121300300-0331321313303330-1310231231013013-1320221110333110-2202021210002033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231220330220023-2313213001223231-3212221130222121-2103300132222102-2100222232312230-1122321330210301-1003321331102312-3201312222113223"></a>

## primary.default_rr_set_group.ds_record.values.sha256_digest — sha256_digest / 230111130322 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.ds_record](data-sources--dns_zone--reference--group-001.md#canonical-0332332020330012-1003302232123021-0022232311333132-0113122113322112-2123033321110312-3133333133301221-0121330302222103-3212000033322100)
- [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1133333130110222-2223102021122311-2113332033031230-1100201102003302-3122213102212223-1032103101022120-2000002000223200-2023321023200100)
- primary.default_rr_set_group.ds_record.values.sha256_digest

<a id="canonical-1211103222131221-0032322120303010-2101302001321032-3231020200110231-3301212111210120-1102020121203303-3220031031021223-3022213213001031"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 digest.

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

<a id="canonical-1122203011002131-3101221233101303-1020303123023012-2030013032132030-3132002210321303-0022220020321222-2033000012020021-3012033102320020"></a>

## Direct properties — sha256_digest / 230111130322 / 3

<a id="canonical-1031223323002303-2111302203323302-0012223231121122-1012003012203033-1120321223130001-0321321302003312-2210122333202000-2000010001121130"></a>

<a id="canonical-2310002331212002-1022103012331010-0020311101333011-1111232033113031-3133201013312120-0123010102233000-3233020111323000-3323022330211302"></a>

## digest property — sha256_digest / 230111130322 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
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
    "minLength": 64
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
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-0213022100011031-0211313332320121-0323112132010011-3102323112222002-2022330311201113-2112230000320121-1012311221231131-0010001021130022"></a>

## Next pages — sha256_digest / 230111130322 / 5

- [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1133333130110222-2223102021122311-2113332033031230-1100201102003302-3122213102212223-1032103101022120-2000002000223200-2023321023200100)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2022232333211100-0121202021103121-2302000333012022-1232200321013202-0201321121130123-3310311330202123-2022000220113232-0111002333132313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101212322133202-1033201010203131-2300200122230023-0010301322322203-3021010201213210-3010021113103031-3220111211033223-2020220313031203"></a>

## primary.default_rr_set_group.ds_record.values.sha384_digest — sha384_digest / 000030333230 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.ds_record](data-sources--dns_zone--reference--group-001.md#canonical-0332332020330012-1003302232123021-0022232311333132-0113122113322112-2123033321110312-3133333133301221-0121330302222103-3212000033322100)
- [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1133333130110222-2223102021122311-2113332033031230-1100201102003302-3122213102212223-1032103101022120-2000002000223200-2023321023200100)
- primary.default_rr_set_group.ds_record.values.sha384_digest

<a id="canonical-1203222223031333-2213313232123133-0132023212213320-2122210011111231-3131223111200131-1123001213031032-3110211201230233-3010301332013201"></a>

Type: `"single"`. Computed.

Configuration parameter for sha384 digest.

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

<a id="canonical-2323010033012132-3123101231023203-3020222231233210-2212121332212020-2120010122331000-2133213301000200-0102112331103020-2111001201003100"></a>

## Direct properties — sha384_digest / 000030333230 / 3

<a id="canonical-2100200132011100-3021232201132311-0110031023111230-0230311130112321-1022312111330101-1000311330010302-0011213030023103-0120313223313312"></a>

<a id="canonical-1011133320101103-2033223320212020-1120203333131101-0112301000032020-1330202100230203-0012322011300123-2232001232023031-2322223022201000"></a>

## digest property — sha384_digest / 000030333230 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-2332133203101121-2100113103110313-3330110222332101-2102230322312133-0320030302022123-3003100233330003-2231012122210220-0310132133302330"></a>

## Next pages — sha384_digest / 000030333230 / 5

- [primary.default_rr_set_group.ds_record.values](data-sources--dns_zone--reference--group-001.md#canonical-1133333130110222-2223102021122311-2113332033031230-1100201102003302-3122213102212223-1032103101022120-2000002000223200-2023321023200100)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-3131300332020230-3331033131000131-1121003101002330-1233102210013032-0213323303121120-0011202223031101-3113322221020122-1031012321320303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002312221010002-2013311313320320-1121202103303210-0122222202210021-1012211132312122-2023222012010000-1112011322022322-0121013032330032"></a>

## primary.default_rr_set_group.eui48_record — eui48_record / 302112011202 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.eui48_record

<a id="canonical-1101003211112233-0020221210212232-2301212122120221-2300232221312301-1030131112022130-3022331321302121-3330012020221302-1003130221133220"></a>

Type: `"single"`. Computed.

Configuration parameter for eui48 record.

Upstream description:

DNS EUI48 Record.

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

<a id="canonical-0203020133112301-0320123312221200-3000211012321333-1222223212202013-1303013332123123-0331203300120212-2203233101331312-1233101013321032"></a>

## Direct properties — eui48_record / 302112011202 / 3

<a id="canonical-3013102300200103-0202323022102033-1200330200122232-1210000233003231-3231021101211002-1232001131313323-2031200010211030-2213112022200230"></a>

<a id="canonical-0321310110001201-0020131023111311-3102203013210031-3210213010113103-0113110113323221-2213012312311200-3312210103110232-0311023132112032"></a>

## name property — eui48_record / 302112011202 / 4

Type: `"string"`. Computed.

EUI48 Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-0001211302020011-2101312333310022-1210110221303312-1000202000302121-2130300202121322-3310030321033312-2100100100120331-3101332333002001"></a>

<a id="canonical-0131233203233101-1021010201202103-3332220010231010-3333112221233202-0232123103100122-0102020221221013-2121232302310033-1232210211120210"></a>

## value property — eui48_record / 302112011202 / 5

Type: `"string"`. Computed.

EUI48 Identifier. A valid eui48 identifier, for example: 01-23-45-67-89-ab.

Upstream description:

A valid eui48 identifier, for example: 01-23-45-67-89-ab.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 17,
  "minLength": 17,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 17,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 17,
    "pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-0111212221102003-1230021332132301-1030021120100222-2210320232033113-0132130021022311-1023032310100331-2023123330300303-1300232120002002"></a>

## Next pages — eui48_record / 302112011202 / 6

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1001031330110020-1133122011223212-1221020310212103-2121030102103310-1031332030211201-3010332112131120-0002121331230022-3122212233332113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023002101212112-1231230202012111-0122021000323331-0330303330321202-2220211301003210-1013302332100203-0112322323020333-1031003323213312"></a>

## primary.default_rr_set_group.eui64_record — eui64_record / 131023330210 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.eui64_record

<a id="canonical-3000330130313332-3332310121201212-2130221001122010-1300301102003012-1013322223111132-2331120122013201-2312123333332202-1012112130103120"></a>

Type: `"single"`. Computed.

Configuration parameter for eui64 record.

Upstream description:

DNS EUI64 Record.

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

<a id="canonical-3101220323313030-2133223122211231-0001021223200323-1203112312303311-1210232131031333-1233131221332013-0313202212021030-0000023120221212"></a>

## Direct properties — eui64_record / 131023330210 / 3

<a id="canonical-1110310131032031-2121201121311221-3220103312311032-2023310333211031-3221112200122201-2212302330023321-1332231122301120-0312310130220002"></a>

<a id="canonical-0202131021022201-2013233022133211-2130122101123011-2222213002202212-0102033011231121-3212031312100321-3213022133110312-1130222012032133"></a>

## name property — eui64_record / 131023330210 / 4

Type: `"string"`. Computed.

EUI64 Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-2222210322121023-2223222212201303-3012320021301102-2131013223033011-3131023001200320-1120233032223220-1021303201121302-3223121031211301"></a>

<a id="canonical-0332002130031020-1000312212111222-2320232003303302-2112113030302030-0021111113311313-3211303010122012-0332231300132320-0213222020013020"></a>

## value property — eui64_record / 131023330210 / 5

Type: `"string"`. Computed.

EUI64 Identifier. A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Upstream description:

A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 23,
  "minLength": 23,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 23,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 23,
    "pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-3121012102322110-2221023100223230-3222110132020133-2211223210022232-3212003231333123-1031230302113322-3103022333103333-0013221130123320"></a>

## Next pages — eui64_record / 131023330210 / 6

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1321110233323023-3033012222210110-1100112210100301-1331110010313330-0202131111030201-0031121030023110-3303012321221102-1202212120312023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101202310310222-3022003313310220-2201003100021022-2123213230222200-2201020122323012-3020103010012233-1323110133130111-3020131000100333"></a>

## primary.default_rr_set_group.lb_record — lb_record / 320133121223 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.lb_record

<a id="canonical-3321013212110302-0023102310032001-3322010023033222-2230130001031320-3001233100201112-1020110302111000-1220132112322010-3203032000220202"></a>

Type: `"single"`. Computed.

DNS Load Balancer Record. DNS Load Balancer Record.

Upstream description:

DNS Load Balancer Record.

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

<a id="canonical-1102011312010130-1111201213013323-1003323130321212-2111331301122121-2211333230212333-1213120202232002-3310331011332111-0323133112131023"></a>

## Direct properties — lb_record / 320133121223 / 3

<a id="canonical-0213222223323021-1232220023221221-2113002203131303-3031102030213313-0231132333033122-2200210100232031-2003322233303023-2010000022313133"></a>

<a id="canonical-0330320213231001-2303330302230330-0211200333333323-1311131032000101-1011010133120103-0333011212332022-1100202002102303-2330102002000202"></a>

## name property — lb_record / 320133121223 / 4

Type: `"string"`. Computed.

Load Balancer record name (except for SRV DNS Load balancer record) should be a simple record name
and not a subdomain of a subdomain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
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
    "maxLength": 255,
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
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

- [value](data-sources--dns_zone--reference--group-002.md#canonical-2311130033030021-2003222210122210-2020122132030012-2213211110331023-0222122311203110-2022333003013121-0100233211220113-2210120022302133): complete subsection reference.

<a id="canonical-3233100101333231-1301123020022311-1220032022122301-2211032310331312-0310112110321231-2123210123122000-1000333101021311-2210031122322033"></a>

## Next pages — lb_record / 320133121223 / 5

- [primary.default_rr_set_group.lb_record.value](data-sources--dns_zone--reference--group-002.md#canonical-2311130033030021-2003222210122210-2020122132030012-2213211110331023-0222122311203110-2022333003013121-0100233211220113-2210120022302133)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2311130033030021-2003222210122210-2020122132030012-2213211110331023-0222122311203110-2022333003013121-0100233211220113-2210120022302133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312011302323310-1220010030202220-2333132131031233-3302032211323222-3220210013320300-2232100012012102-3230000200003200-0110101211202221"></a>

## primary.default_rr_set_group.lb_record.value — value / 013212211303 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.lb_record](data-sources--dns_zone--reference--group-002.md#canonical-1321110233323023-3033012222210110-1100112210100301-1331110010313330-0202131111030201-0031121030023110-3303012321221102-1202212120312023)
- primary.default_rr_set_group.lb_record.value

<a id="canonical-3132203013302103-1312022333333202-3130320101313310-3303130301333122-2113200331320023-0321031213112232-0322210330021013-0121332320212031"></a>

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

<a id="canonical-3100121013300220-0232021120213332-1120033030321332-3221312222201301-1003333323100332-1333203023211031-0320312000030323-3020301223101121"></a>

## Direct properties — value / 013212211303 / 3

<a id="canonical-1131222303231032-1113322200021231-3133023102131231-1200300223133320-3311022322023333-2020231221020003-0011023000010012-3221322012121310"></a>

<a id="canonical-2200023111022330-1013022223322033-1131133103112332-0113200320020300-1233300121333312-2023322120100303-2210011313303131-2203021220113022"></a>

## name property — value / 013212211303 / 4

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

<a id="canonical-1212022211001321-1200111110200222-2303103333002230-1013031323000102-0122222231132301-1333211130211121-0131132321223001-0101122332120212"></a>

<a id="canonical-0233132321113110-3033102112003032-1131113100301130-3031123113302321-1310000020033030-0012110100112210-1220203012013330-3330201201301000"></a>

## namespace property — value / 013212211303 / 5

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

<a id="canonical-3230102333320033-2110000030300102-0300320022222021-3012013123221123-2300002232010302-1112122220233213-0102120012112220-0321103120121233"></a>

<a id="canonical-2021200103122011-3311132212212130-2232200110233313-2131102221302332-1233311211301313-2001312321222312-3011013120022212-3200003103020022"></a>

## tenant property — value / 013212211303 / 6

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

<a id="canonical-2013321132111212-1322202320100200-2200112303102022-1221003311200013-2033021130230312-2020020302200230-1130013110213333-3213110213210000"></a>

## Next pages — value / 013212211303 / 7

- [primary.default_rr_set_group.lb_record](data-sources--dns_zone--reference--group-002.md#canonical-1321110233323023-3033012222210110-1100112210100301-1331110010313330-0202131111030201-0031121030023110-3303012321221102-1202212120312023)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-3310122013002202-2201310111100013-3320132322330303-3003311222123331-1221002132100302-2122021012023311-0330301233113311-0002130121010103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312331122230232-0000202022002122-0311223000110010-2203112121033301-0312011200003323-3102010011320111-3300112303311122-3213302231220223"></a>

## primary.default_rr_set_group.loc_record — loc_record / 310220111111 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.loc_record

<a id="canonical-3302122211311222-3103311102023100-3220103323100201-2123010111112323-3123113301023332-0230002213201013-0330231020002311-2321300011222312"></a>

Type: `"single"`. Computed.

DNS LOC Record. DNS LOC Record.

Upstream description:

DNS LOC Record.

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

<a id="canonical-3020331133313333-3013220202111001-1311321031311123-1212133313200111-2131330012011132-1100211322021013-2003000121232122-2003023110103001"></a>

## Direct properties — loc_record / 310220111111 / 3

<a id="canonical-1103230001020130-3331023302200223-0112220003002223-1000103012121303-0222220111130123-2020333213331303-0120213102211032-0313311132021233"></a>

<a id="canonical-1120301020131031-2010200111223302-2131220102210223-3301122113000100-1300112323013030-1001011122211202-2211131232202132-3133001111022130"></a>

## name property — loc_record / 310220111111 / 4

Type: `"string"`. Computed.

LOC Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-0332210330031101-0030003323333000-0122030130330130-0013312101002233-1233230131230102-0300231310030003-0331020120103023-0133320311021203): complete subsection reference.

<a id="canonical-2112332101102013-1233103321313100-0210132300233003-1333112331201121-3332311021110003-0212131311300011-1233122033311302-1003203002222031"></a>

## Next pages — loc_record / 310220111111 / 5

- [primary.default_rr_set_group.loc_record.values](data-sources--dns_zone--reference--group-002.md#canonical-0332210330031101-0030003323333000-0122030130330130-0013312101002233-1233230131230102-0300231310030003-0331020120103023-0133320311021203)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0332210330031101-0030003323333000-0122030130330130-0013312101002233-1233230131230102-0300231310030003-0331020120103023-0133320311021203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113113223103111-2320222030223202-1020002023000022-2132131222002203-2322301132312211-3033232100122111-2223000313032203-2001002313102011"></a>

## primary.default_rr_set_group.loc_record.values — values / 331131333233 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.loc_record](data-sources--dns_zone--reference--group-002.md#canonical-3310122013002202-2201310111100013-3320132322330303-3003311222123331-1221002132100302-2122021012023311-0330301233113311-0002130121010103)
- primary.default_rr_set_group.loc_record.values

<a id="canonical-3113200210231122-0221222203132301-0212030311303310-0330133133130233-2203302103302001-2132123322012331-2103102010130012-1013320022212120"></a>

Type: `"list"`. Computed.

LOC Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2101003213302130-3123302102002330-2302102222232200-2130033111212133-0130212000331022-0021203202012213-1211301103122101-3012321031233332"></a>

## Direct properties — values / 331131333233 / 3

<a id="canonical-3311200200213133-2021210303021311-2133231022013130-0331120213113330-2022203010312111-1301032302320021-1313021230230100-1121013032231210"></a>

<a id="canonical-1323233021021113-0203113300312320-2012220321321001-3131210312013111-2013121311010012-1320112132031201-0032002023032313-0321031123030022"></a>

## altitude property — values / 331131333233 / 4

Type: `"number"`. Computed.

Altitude. Altitude in meters.

Upstream description:

Altitude in meters.

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
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2300120201211002-2310321321030112-2130230300003332-0231021010202202-3122130320132231-3000013013122001-0222001100232331-3001110311320213"></a>

<a id="canonical-3033333030111201-2313010322321332-0023311130210303-3021001311030110-0002023111233320-1310331112333221-0220013121012202-2311313223320112"></a>

## horizontal_precision property — values / 331131333233 / 5

Type: `"number"`. Computed.

Horizontal Precision. Horizontal Precision in meters.

Upstream description:

Horizontal Precision in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-3110303113003102-1031302331030321-0323110313132003-2131313102230020-0221013012300111-3111333123101320-0010001012103212-0230030110121310"></a>

<a id="canonical-3110000002223213-3020310310222010-1313011232121221-0101321032233021-2230221021122332-1102210100100221-2121313333333200-2221321000113031"></a>

## latitude_degree property — values / 331131333233 / 6

Type: `"number"`. Computed.

Latitude degree, an integer between 0 and 90, including 0 and 90.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 90,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0031230002231000-1210022311012023-0100113101103331-1101303222012222-3232112130033001-0320130022210312-1103201301000220-0102333100100110"></a>

<a id="canonical-0222213320032200-0031210310103201-3122323213320302-0131233032302201-3100121200310223-2323002322301133-0231300123003121-3201332000130113"></a>

## latitude_hemisphere property — values / 331131333233 / 7

Type: `"string"`. Computed.

\[Enum: N|S\] Latitude hemisphere can only be N or S - N: North Hemisphere - S: South Hemisphere.
Possible values are \`N\`, \`S\`. Defaults to \`N\`.

Upstream description:

Latitude hemisphere can only be N or S

&#8203;- N: North Hemisphere

&#8203;- S: South Hemisphere.

Receipt-pinned upstream constraints:

```json
{
  "default": "N",
  "enum": [
    "N",
    "S"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2302110130023131-3121330000331301-1003332213113101-3130231322033301-0111122301311023-1201001100203033-1023313100311211-3131233302311013"></a>

<a id="canonical-1331130020302210-3113111233113203-1103030300310221-3312030023012010-0103333112130032-0012112123002122-3302112200011330-2110201200113100"></a>

## latitude_minute property — values / 331131333233 / 8

Type: `"number"`. Computed.

Latitude minute, an integer between 0 and 59, including 0 and 59.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-1311022203303033-3112312100332301-3310013101031230-3202311111222113-0202330221113131-3320312123311020-2233312002213221-3100303212010301"></a>

<a id="canonical-1301002030313102-0303103030120011-2301211113312302-3232220021232320-0330210002201321-0100301213332122-2320222331112302-0333300010032012"></a>

## latitude_second property — values / 331131333233 / 9

Type: `"number"`. Computed.

Latitude second, an decimal between 0 and 59.999, including 0 and 59.999.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-1332020110030220-0110323332033122-0211221103213133-2132023210320332-3200130321302230-3113133201030012-0201213332203223-1203223113202330"></a>

<a id="canonical-3000232211302310-0230012020101301-0200331102221000-2321023000102023-1211032231201111-3321213133210132-2020130331101033-2123230020101232"></a>

## location_diameter property — values / 331131333233 / 10

Type: `"number"`. Computed.

Diameter of a sphere enclosing the described entity, in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-3120232123033331-0300133111133213-3111211323203213-3310210223031100-1131111302113223-0211120310003212-3032312321123012-0002301303031332"></a>

<a id="canonical-2020110220123102-3003211330323201-0032033230201323-3103231012302112-2333303310213200-3211310120103001-1332113320223032-2333323220312013"></a>

## longitude_degree property — values / 331131333233 / 11

Type: `"number"`. Computed.

Longitude degree, an integer between 0 and 180, including 0 and 180.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2100132230332122-0112003022133101-1111223133113332-1113333203031030-0011231300022013-1210110123230003-3321212000222003-3220101322211323"></a>

<a id="canonical-1311010023102033-0110233231322100-3003111010003111-2100021310302312-0031232320302020-0002302322111333-1130123310312122-3302303230122112"></a>

## longitude_hemisphere property — values / 331131333233 / 12

Type: `"string"`. Computed.

\[Enum: E|W\] Longitude hemisphere can only be E or W - E: East Hemisphere - W: West Hemisphere.
Possible values are \`E\`, \`W\`. Defaults to \`E\`.

Upstream description:

Longitude hemisphere can only be E or W

&#8203;- E: East Hemisphere

&#8203;- W: West Hemisphere.

Receipt-pinned upstream constraints:

```json
{
  "default": "E",
  "enum": [
    "E",
    "W"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2023123133322101-3220201332030302-2223300002022200-1211012232222333-3213230220333023-0101130013312211-3132312322023222-1133300303332100"></a>

<a id="canonical-1200231221033210-1122132231201110-0223321231203031-3222323113303001-2311233122212130-3311012103133203-3201230230222132-0010220022020323"></a>

## longitude_minute property — values / 331131333233 / 13

Type: `"number"`. Computed.

Longitude minute, an integer between 0 and 59, including 0 and 59.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-0130330210323301-2302302203032002-3013230121123223-2012002223231113-3003312111101111-1133012123001123-1001222022213200-3233112113322123"></a>

<a id="canonical-0330023130210233-0312223000203100-1123212320000032-1122120200323210-0301013103222001-0233023021321302-1302113310020003-1221132221120203"></a>

## longitude_second property — values / 331131333233 / 14

Type: `"number"`. Computed.

Longitude second, an decimal between 0 and 59.999, including 0 and 59.999.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-2111033231212233-2011201101302003-2221003230122121-1013131311032202-0001230011101203-3113101020101303-1213112220312030-3331012100003310"></a>

<a id="canonical-3302210313101321-0200211203221002-2121221020110111-0213003100020110-2001310221223313-3320303212012202-2201030200320322-3130003031233133"></a>

## vertical_precision property — values / 331131333233 / 15

Type: `"number"`. Computed.

Vertical Precision. Vertical Precision in meters.

Upstream description:

Vertical Precision in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-2030311233311330-1221202020231113-0222110303102011-3303313222223111-0201013131320001-2100331032200122-3213020200232112-2003003030102012"></a>

## Next pages — values / 331131333233 / 16

- [primary.default_rr_set_group.loc_record](data-sources--dns_zone--reference--group-002.md#canonical-3310122013002202-2201310111100013-3320132322330303-3003311222123331-1221002132100302-2122021012023311-0330301233113311-0002130121010103)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1210121201311120-2131010222033122-0330222031322111-3200022000122302-0333021032123112-2220203201001212-2120030320310132-2322032202111212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210020110023020-1002001331011011-0111123221022300-2021313211130221-3331120003230321-2312100233203021-0021133013213330-1102201022103013"></a>

## primary.default_rr_set_group.mx_record — mx_record / 331322201121 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.mx_record

<a id="canonical-1032311100202302-3103323103112013-1333010002323213-2002133011303012-0230220002213110-1003213032321220-1130322103013231-3303203111301121"></a>

Type: `"single"`. Computed.

DNSMXResourceRecord.

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

<a id="canonical-3233220101310111-2022133102130031-2201102331030132-0331211011333313-2013321121100132-0130131221102001-3030332012012030-2230310132010122"></a>

## Direct properties — mx_record / 331322201121 / 3

<a id="canonical-0333112301302200-3132212312222130-0022323331310010-2303001231202322-0313332131233111-2202322212320001-0012222330120330-2020010020321221"></a>

<a id="canonical-2022003112323101-3201210020113233-2012011020331132-3100101112100030-2312333301312222-2232000332211022-2313231003012313-3003331132210321"></a>

## name property — mx_record / 331322201121 / 4

Type: `"string"`. Computed.

MX Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-1011130322331212-0301101323322001-0320200330211331-2122132002032322-1003202023031203-2021031220230300-1212030120033231-0033332212110320): complete subsection reference.

<a id="canonical-1121300120102102-0221101330120013-2020233102012222-1021010330203111-1110032203032130-3312221301231112-0310111233301013-2010010110211021"></a>

## Next pages — mx_record / 331322201121 / 5

- [primary.default_rr_set_group.mx_record.values](data-sources--dns_zone--reference--group-002.md#canonical-1011130322331212-0301101323322001-0320200330211331-2122132002032322-1003202023031203-2021031220230300-1212030120033231-0033332212110320)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1011130322331212-0301101323322001-0320200330211331-2122132002032322-1003202023031203-2021031220230300-1212030120033231-0033332212110320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312032230300023-1301133310102232-3330311222023331-2131330122122203-1121200223120310-2102012210213112-2022213010211313-0202230122300212"></a>

## primary.default_rr_set_group.mx_record.values — values / 103220201212 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.mx_record](data-sources--dns_zone--reference--group-002.md#canonical-1210121201311120-2131010222033122-0330222031322111-3200022000122302-0333021032123112-2220203201001212-2120030320310132-2322032202111212)
- primary.default_rr_set_group.mx_record.values

<a id="canonical-3233232323213020-1102033211231011-1321202112322120-0333320230101121-0320302103020112-1331300012031212-0022010320131022-3001202130130233"></a>

Type: `"list"`. Computed.

MX Record Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

<a id="canonical-3313322000200103-2310320322133022-2032013302100223-3020210022120123-1310233033320333-3332123231312031-1112211313001313-1223002300203232"></a>

## Direct properties — values / 103220201212 / 3

<a id="canonical-3212233103010021-2112330330101021-3311113123112320-3210112120133231-1123310220120112-1301333133002022-2113311003233020-1203122111002323"></a>

<a id="canonical-2221103313323332-0033102022131233-2003310112133333-3232103112133131-0130233110221213-2133223132300112-1033301030213211-1120100100303000"></a>

## domain property — values / 103220201212 / 4

Type: `"string"`. Computed.

Mail exchanger domain name, please provide the full hostname, for.

Upstream description:

Mail exchanger domain name, please provide the full hostname, for example: mail.example.com.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-3322212023020301-2102010023101223-1120101300130022-3000311103023230-3220132000213212-3210333213332211-3112113323201011-3332223233032202"></a>

<a id="canonical-2010022203111113-0113131332112121-2211030213313303-2332113203100200-3133223132310212-3113301010023332-0311113011231330-1210033000113302"></a>

## priority property — values / 103220201212 / 5

Type: `"number"`. Computed.

Priority. Mail exchanger priority code.

Upstream description:

Mail exchanger priority code.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0213212031312002-2313212313000202-1002000233003210-2010210103013102-2221111321212113-2333320202113133-2311023112021003-1030313003102322"></a>

## Next pages — values / 103220201212 / 6

- [primary.default_rr_set_group.mx_record](data-sources--dns_zone--reference--group-002.md#canonical-1210121201311120-2131010222033122-0330222031322111-3200022000122302-0333021032123112-2220203201001212-2120030320310132-2322032202111212)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2000221213012233-1222212300001003-2201121122210020-3133301203202102-1332032202320301-1003102301122213-2203320000301213-0012113303201103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133220221222203-1303310310332011-1212101231112031-2003032212321302-0312120000302100-0103330001231012-2123300130331330-3030221210220120"></a>

## primary.default_rr_set_group.naptr_record — naptr_record / 001233132231 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.naptr_record

<a id="canonical-1131003222002322-0331310113232231-0120321102333300-2000203102120231-0302103021322220-3232023223211033-2330023121211213-2302003331022033"></a>

Type: `"single"`. Computed.

Configuration parameter for naptr record.

Upstream description:

DNS NAPTR Record.

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

<a id="canonical-1113303322022003-2232211222012212-2212102311300101-1320311232031133-3201200322112122-0200002303202130-3210130322302102-3213330331201001"></a>

## Direct properties — naptr_record / 001233132231 / 3

<a id="canonical-1310020030321333-0101103311331302-2303012331310110-3003213333031333-1013020301321101-1322312211323103-3013310130323002-3013132232123002"></a>

<a id="canonical-1011223110202123-2021001231232313-1200212132020312-2033021121110003-3023100031230113-0210110100102033-0303012213132111-2022021002211213"></a>

## name property — naptr_record / 001233132231 / 4

Type: `"string"`. Computed.

NAPTR Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-0100201121020123-1031113322131133-1011321031132003-3321103133212313-3031002121201312-0312231201330322-1033220103223301-1023211031120113): complete subsection reference.

<a id="canonical-1103332000211132-2230330302212122-1011023000231000-1321110010111013-3110123111331233-2213130333233202-2131131211032203-3220132112001132"></a>

## Next pages — naptr_record / 001233132231 / 5

- [primary.default_rr_set_group.naptr_record.values](data-sources--dns_zone--reference--group-002.md#canonical-0100201121020123-1031113322131133-1011321031132003-3321103133212313-3031002121201312-0312231201330322-1033220103223301-1023211031120113)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0100201121020123-1031113322131133-1011321031132003-3321103133212313-3031002121201312-0312231201330322-1033220103223301-1023211031120113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022202101033110-0001013110121023-1120303322311032-0202332120303103-0132203000012211-3331300212301323-3310331133300120-2323123110230100"></a>

## primary.default_rr_set_group.naptr_record.values — values / 122210021123 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.naptr_record](data-sources--dns_zone--reference--group-002.md#canonical-2000221213012233-1222212300001003-2201121122210020-3133301203202102-1332032202320301-1003102301122213-2203320000301213-0012113303201103)
- primary.default_rr_set_group.naptr_record.values

<a id="canonical-2320200022012033-1033033032110021-0102202303021210-3030031212130021-0300133300102312-2203100103330133-2300012122032212-1100023312311020"></a>

Type: `"list"`. Computed.

NAPTR Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0111332101221002-3021123011122320-2101133120300020-0032102013303332-3321312213302201-1212302301332100-2311012333011323-1030233320012031"></a>

## Direct properties — values / 122210021123 / 3

<a id="canonical-3220320022000021-1233032000023232-0021330031311101-1003210020112310-0111322332300202-2330101220201211-0131110023033311-2302310101333232"></a>

<a id="canonical-0021112211010312-2323330213210023-0332221031300130-0312200331023210-3111013120333220-1111002102103232-2330003231032233-1102212120120001"></a>

## flags property — values / 122210021123 / 4

Type: `"string"`. Computed.

Flag to control aspects of the rewriting and interpretation of the fields in the record. At this
time only four flags, S/A/U/P, are defined.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  }
}
```

<a id="canonical-0300031121323202-3023112232120133-0130222230303332-0302000303203010-0010113133012022-2012120103012321-2323100221100333-2202023212021101"></a>

<a id="canonical-3132223202032331-2002130020223213-2300221210103120-1220020010300100-2032311001330232-0000010200231010-2012331012111231-1223313031312211"></a>

## order property — values / 122210021123 / 5

Type: `"number"`. Computed.

Order in which the NAPTR records must be processed. A lower number indicates a higher preference.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2011021203322230-2231013332301131-3132101230023310-3023000312210330-1333030302112031-2210123203222133-0111202001321110-2302302111210113"></a>

<a id="canonical-3213130012332013-3013232113222102-2021231230120111-1311210220012011-0211332012001201-2211213322031211-3320302130223321-3222331030001100"></a>

## preference property — values / 122210021123 / 6

Type: `"number"`. Computed.

Preference when records have the same order. A lower number indicates a higher preference.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3330121203020203-0221132100100003-2002333210221223-1311220321133311-3323323311333033-2122031012210132-1223102033211211-3032111300330202"></a>

<a id="canonical-1321122000102112-0011000121321122-2121131003320303-0222102032030303-3110131000201322-0122213013313231-2121310100132333-1322213311210122"></a>

## regular expression property — values / 122210021123 / 7

Type: `"string"`. Computed.

Regular expression to construct the next domain name to lookup.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
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
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-3033113003322110-2210300231010000-0003313223021212-0212010220101222-1000113022332113-2210013200131230-1312001131132203-0331103032022122"></a>

<a id="canonical-2231010212321012-0320112121223122-0132122130320121-0110210031213300-3332002120102030-3033300111010223-1120121112133333-2003330313021102"></a>

## replacement property — values / 122210021123 / 8

Type: `"string"`. Computed.

The next NAME to query for NAPTR, SRV, or address records depending on the value of the flags field.

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

<a id="canonical-1103001310312002-2033033120313313-3301110213102322-1200103113111121-0111201223100233-1210011003112100-3331010212033110-2111113103012012"></a>

<a id="canonical-0012013223000133-0031110031100310-1212103120111020-0212202221132021-1122102113133023-0023020122231010-3130113121331311-1133111010120112"></a>

## service property — values / 122210021123 / 9

Type: `"string"`. Computed.

Specifies the service(s) available down this rewrite path.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  }
}
```

<a id="canonical-3313033213323003-0233131333121220-2021020302212011-1200000023030301-0320330110121322-1030000301022111-1131233030003101-2110301131212002"></a>

## Next pages — values / 122210021123 / 10

- [primary.default_rr_set_group.naptr_record](data-sources--dns_zone--reference--group-002.md#canonical-2000221213012233-1222212300001003-2201121122210020-3133301203202102-1332032202320301-1003102301122213-2203320000301213-0012113303201103)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2313231112212011-0201330312333230-1201033232031120-1123102230132211-2311130313033231-0220202113331330-3030032011212201-2022201023220103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021003300030212-2213203212320133-0323010212313030-3312001122120303-0122212011332132-3123203223110001-3111033220210220-0300332301230133"></a>

## primary.default_rr_set_group.ns_record — ns_record / 130301113113 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.ns_record

<a id="canonical-3120011311313020-0111201012211122-2213331000100130-1232101330013230-0030203001332200-2000311000020320-3311002011130312-1322011010013123"></a>

Type: `"single"`. Computed.

DNSNSResourceRecord.

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

<a id="canonical-1031030212322332-2223221201210333-2312121122113212-1121012300312011-0212013110003201-3032312310323221-2200012300303002-1123323013323230"></a>

## Direct properties — ns_record / 130301113113 / 3

<a id="canonical-2213320012022313-1021103303322033-0111031323202112-3312112010303203-3011210120120320-2230203102200111-2211022130202111-3010320210120211"></a>

<a id="canonical-0011300221111202-3130210321012021-2323320211020001-3330321031210102-1030202330210311-3321131323033133-1201312201100233-0023112131212211"></a>

## name property — ns_record / 130301113113 / 4

Type: `"string"`. Computed.

NS Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-0300120203320102-3303321332223332-0323100011233013-3230312112133013-3101333233332231-1033102102210322-3132203032100032-3220031211032023"></a>

<a id="canonical-2320102033233333-0123203033310221-3020210300101102-0322230301021002-1033030001031333-0033211231103000-0311002002213102-3131002202102333"></a>

## values property — ns_record / 130301113113 / 5

Type: `["list", "string"]`. Computed.

Name Servers. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1010331003211210-2132302321131323-3312332131002312-3321223012223221-2132311031013302-2021320200033023-3313301021213203-3300001102111332"></a>

## Next pages — ns_record / 130301113113 / 6

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0032230300312201-2330203102230021-2122122330231113-1232330231302310-3001020022321003-0100000031302303-0320002132012223-3332123310323113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301302133020033-3203320311100113-1222133021121003-1211200322001303-3331203323211132-0032120322123121-3223020011031230-1212031210023330"></a>

## primary.default_rr_set_group.ptr_record — ptr_record / 210323013111 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.ptr_record

<a id="canonical-3303011132132001-3122103213121003-3311211220312210-1111031013000031-3212320011003331-3323012020013020-0120110300202132-1012322330233103"></a>

Type: `"single"`. Computed.

DNSPTRResourceRecord.

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

<a id="canonical-2213012211303220-2002200011123231-0203230000020231-1033122211020031-2102202323332330-1103322212213330-2030022021022031-0330101230232001"></a>

## Direct properties — ptr_record / 210323013111 / 3

<a id="canonical-1023002231030203-1111200200202010-1320211303302021-0033222220111300-1011203222322021-1122221223213310-0210233000330313-0013022333303210"></a>

<a id="canonical-0102321021122113-1000332300021313-2031322231300321-2320330302331002-1333120211212022-0332232210010230-0333211033220003-3102100133112032"></a>

## name property — ptr_record / 210323013111 / 4

Type: `"string"`. Computed.

PTR Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-2313101211020213-0201123002020023-1113010113221221-2301313211102232-3200202311133321-3332021221031222-0321110113201202-0200112033033110"></a>

<a id="canonical-1103023011010212-2103023113012333-1111021231022221-1233212213333011-1301120202130022-1303132222331123-2201110000120032-3201312122230013"></a>

## values property — ptr_record / 210323013111 / 5

Type: `["list", "string"]`. Computed.

Domain Name. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2111200302313031-2111223223203032-3222132301130233-1132220001021301-2112321112110233-0112233133201330-3323320023331023-0030301010333012"></a>

## Next pages — ptr_record / 210323013111 / 6

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-3022002303322101-2133310220222102-1221303323002002-2021320313022100-3022301033201311-3100301120122232-0121032103302112-2011211020313032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201103111222310-2010311122220201-2021132213110313-1120010132103012-2132020321312110-1130203310022111-3300233000000222-3303311220120023"></a>

## primary.default_rr_set_group.srv_record — srv_record / 222033001320 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.srv_record

<a id="canonical-2001102313121111-1031012001223021-1320032301003300-1221012201112033-0210231131333011-1223133113110100-3122001233211300-1233200332020001"></a>

Type: `"single"`. Computed.

DNSSRVResourceRecord.

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

<a id="canonical-3323132310222133-3212300323332020-0102031001332133-0203303333103231-1100022203203303-1002032203231302-2103102130230133-2212203223332221"></a>

## Direct properties — srv_record / 222033001320 / 3

<a id="canonical-0323130100321031-0322103030031330-0332231023332230-2002133130023113-3111213100333233-3133102030110313-3120321101233323-1223230200201300"></a>

<a id="canonical-3120221333312111-0000130313102131-3210230123332011-1330333321132021-2303331101301302-3333303030301213-3311210222000223-3033322033021320"></a>

## name property — srv_record / 222033001320 / 4

Type: `"string"`. Computed.

SRV Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-1320223020133013-1023223333033030-3122310013013203-0033013220230323-1203220300022200-0213310103002010-0022011222300321-0201010301013210): complete subsection reference.

<a id="canonical-1300313223121201-3223121023110133-1021331300011100-3323313030130231-3123300003002003-1200302003012202-0023111100002120-3310100123212002"></a>

## Next pages — srv_record / 222033001320 / 5

- [primary.default_rr_set_group.srv_record.values](data-sources--dns_zone--reference--group-002.md#canonical-1320223020133013-1023223333033030-3122310013013203-0033013220230323-1203220300022200-0213310103002010-0022011222300321-0201010301013210)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1320223020133013-1023223333033030-3122310013013203-0033013220230323-1203220300022200-0213310103002010-0022011222300321-0201010301013210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002213123120032-3313100203112220-0023230300130020-2031132222221233-1123132233301232-1023312200310310-2233013110201020-1013210002033203"></a>

## primary.default_rr_set_group.srv_record.values — values / 221003332320 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.srv_record](data-sources--dns_zone--reference--group-002.md#canonical-3022002303322101-2133310220222102-1221303323002002-2021320313022100-3022301033201311-3100301120122232-0121032103302112-2011211020313032)
- primary.default_rr_set_group.srv_record.values

<a id="canonical-3011201211313333-0132121313130231-1223101003211111-3213221203211211-2320300231031112-1033212012000030-3311101111020022-3113122301120032"></a>

Type: `"list"`. Computed.

SRV Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3222111213011132-3023010320000122-2323232212112103-0113102122013221-2121132311332323-3033012131123213-1102303102212113-1331310002300331"></a>

## Direct properties — values / 221003332320 / 3

<a id="canonical-0001332202013201-2121012121222022-1020110233101232-1311321312220111-0023022022232021-0130000111100131-3221031321133220-0211320210103302"></a>

<a id="canonical-1020222122221103-3130021203331100-2232021221200200-0201012333301312-1023212012233303-2333202013113122-3033033111331222-1212331130012303"></a>

## port property — values / 221003332320 / 4

Type: `"number"`. Computed.

Port. Port on which the service can be found.

Upstream description:

Port on which the service can be found.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0203121331131200-1221202321111221-2331131303003330-3232111021110213-3010231102103002-0003330311113320-1330012010132020-2320212112122011"></a>

<a id="canonical-1021012231230112-1001203333322321-3331223012223120-0332020223323000-0301333332023023-1001330003323032-1012313322121111-1121200332323101"></a>

## priority property — values / 221003332320 / 5

Type: `"number"`. Computed.

Priority of the target. A lower number indicates a higher preference.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0130212300322010-0021212210213323-3301110210322211-2131222211120020-2303212202032322-0322020110111220-2311332113310212-1303202022020203"></a>

<a id="canonical-1321110122212123-3030133330323233-0030100120222011-0302330211312111-0201010213013323-2232113331020211-3201300101223202-2233203012203021"></a>

## target property — values / 221003332320 / 6

Type: `"string"`. Computed.

Hostname of the machine providing the service.

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
    "pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  }
}
```

<a id="canonical-1231123310210022-3110320103030330-1033213211212230-1102120031210331-2123020301223321-2202023200133222-2023030101310323-3302223332123313"></a>

<a id="canonical-3023201113322321-1321002112131011-3313223322113302-0003200331332220-3133111022123201-3021213302302002-0331121300110001-2303102201302333"></a>

## weight property — values / 221003332320 / 7

Type: `"number"`. Computed.

Weight of the target. A higher number indicates a higher preference.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0303222323330320-0112102232221032-2002111130231022-2333001133001312-3221300312221201-0121110012111221-1112322323002220-2111212011133220"></a>

## Next pages — values / 221003332320 / 8

- [primary.default_rr_set_group.srv_record](data-sources--dns_zone--reference--group-002.md#canonical-3022002303322101-2133310220222102-1221303323002002-2021320313022100-3022301033201311-3100301120122232-0121032103302112-2011211020313032)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2131213003100201-0112010012001002-0013220210323230-0322333202002311-1323221222033032-0212121113221221-0222110233100122-1122121333301021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230132232203013-2000030110130320-3021013220020303-3030322231103033-1003222101320223-3120013331302030-0130322130231211-3310320213302213"></a>

## primary.default_rr_set_group.sshfp_record — sshfp_record / 012113323130 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.sshfp_record

<a id="canonical-2013012331332022-3033213302020033-1223002311330010-0102020002302130-2100313302123103-2003323030131310-2232212330323011-3212033020230212"></a>

Type: `"single"`. Computed.

Configuration parameter for sshfp record.

Upstream description:

DNS SSHFP Record.

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

<a id="canonical-0032111021203010-2033300231020001-3021200201233101-0020221311010310-1001031232203310-1230302311013323-0212103320121322-2220330201103012"></a>

## Direct properties — sshfp_record / 012113323130 / 3

<a id="canonical-1023122132223330-2001321301031200-1030312232001233-0201023130202100-2122322002301000-3102310331020333-3320013132232223-3030021323120222"></a>

<a id="canonical-1100003301201110-2100221303013220-0122120011223122-2003210111231123-1011110132020221-3011300002003002-2201112330013112-2310323300232010"></a>

## name property — sshfp_record / 012113323130 / 4

Type: `"string"`. Computed.

SSHFP Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-3202003222120313-3231001131322220-1302303203232200-0333301211221330-1312101031123012-0333320322003303-3011001223203333-1003121222013103): complete subsection reference.

<a id="canonical-3332223300201200-0003100111223202-2120123313030300-1313102332100220-1022030112220103-3103022212202212-3133003220312321-3230320330133310"></a>

## Next pages — sshfp_record / 012113323130 / 5

- [primary.default_rr_set_group.sshfp_record.values](data-sources--dns_zone--reference--group-002.md#canonical-3202003222120313-3231001131322220-1302303203232200-0333301211221330-1312101031123012-0333320322003303-3011001223203333-1003121222013103)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-3202003222120313-3231001131322220-1302303203232200-0333301211221330-1312101031123012-0333320322003303-3011001223203333-1003121222013103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011022012313231-2130311020210332-0312002110330223-3110203022322323-2300323330021322-1223330221233300-0301320120302122-3112123022012011"></a>

## primary.default_rr_set_group.sshfp_record.values — values / 302321132113 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-2131213003100201-0112010012001002-0013220210323230-0322333202002311-1323221222033032-0212121113221221-0222110233100122-1122121333301021)
- primary.default_rr_set_group.sshfp_record.values

<a id="canonical-0002031333331122-1032020130110201-3102202233202031-1002322002200203-2030133203023222-3020310021213033-2031110213101210-2130013011222231"></a>

Type: `"list"`. Computed.

SSHFP Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-1320002312203113-0322201312132320-0223233011331120-0322123013103033-3222203033100212-3333122203211221-3213313132131233-3333321102000111"></a>

## Direct properties — values / 302321132113 / 3

<a id="canonical-0303030022003011-3020231322301030-0330110112122301-0330022000303231-2203332223112112-3201230320313201-1023202102110220-1001333003113030"></a>

<a id="canonical-3112021312222123-3223322322123233-2221231132121113-1232223110003001-2201112233120200-2331232133031212-0202011031220033-0211000103231213"></a>

## algorithm property — values / 302321132113 / 4

Type: `"string"`. Computed.

\[Enum: UNSPECIFIEDALGORITHM|RSA|DSA|ECDSA|Ed25519|Ed448\] SSHFP algorithm value must be compatible
with the specified algorithm. - UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM - RSA: RSA - DSA: DSA -
ECDSA: ECDSA - Ed25519: Ed25519 - Ed448: Ed448. Possible values are \`UNSPECIFIEDALGORITHM\`,
\`RSA\`, \`DSA\`, \`ECDSA\`, \`Ed25519\`, \`Ed448\`. Defaults to \`UNSPECIFIEDALGORITHM\`.

Upstream description:

SSHFP algorithm value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM

&#8203;- RSA: RSA

&#8203;- DSA: DSA

&#8203;- ECDSA: ECDSA

&#8203;- Ed25519: Ed25519

&#8203;- Ed448: Ed448.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIEDALGORITHM",
  "enum": [
    "UNSPECIFIEDALGORITHM",
    "RSA",
    "DSA",
    "ECDSA",
    "Ed25519",
    "Ed448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [sha1_fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-3212233132102313-3020030030330330-1300213312321310-0320003102123111-2030132233231132-3223023333221311-2312331123120223-2012333000021131): complete subsection reference.

- [sha256_fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-1033003133000023-0122010033130000-1001333131130010-2300232320011323-1020123031313230-0023020332032322-0022313002020000-2021012001232012): complete subsection reference.

<a id="canonical-3301002002201233-2221130001211313-1321320230230123-3131110323032220-3332100102033113-2002110222302232-2320221130122333-3311330012233300"></a>

## Next pages — values / 302321132113 / 5

- [primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-3212233132102313-3020030030330330-1300213312321310-0320003102123111-2030132233231132-3223023333221311-2312331123120223-2012333000021131)
- [primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint](data-sources--dns_zone--reference--group-002.md#canonical-1033003133000023-0122010033130000-1001333131130010-2300232320011323-1020123031313230-0023020332032322-0022313002020000-2021012001232012)
- [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-2131213003100201-0112010012001002-0013220210323230-0322333202002311-1323221222033032-0212121113221221-0222110233100122-1122121333301021)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-3212233132102313-3020030030330330-1300213312321310-0320003102123111-2030132233231132-3223023333221311-2312331123120223-2012333000021131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000220233310032-3231230321002033-1113113313322121-2123332320130333-2210233231220220-2320320222033321-1132132132120213-1211321103000130"></a>

## primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint — sha1_fingerprint / 132030202230 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-2131213003100201-0112010012001002-0013220210323230-0322333202002311-1323221222033032-0212121113221221-0222110233100122-1122121333301021)
- [primary.default_rr_set_group.sshfp_record.values](data-sources--dns_zone--reference--group-002.md#canonical-3202003222120313-3231001131322220-1302303203232200-0333301211221330-1312101031123012-0333320322003303-3011001223203333-1003121222013103)
- primary.default_rr_set_group.sshfp_record.values.sha1_fingerprint

<a id="canonical-2301022221023201-1013221230312100-0313330033313231-3210020032122202-2321203113211310-2021331300331312-0010123010032000-2212302103211002"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 fingerprint.

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

<a id="canonical-1102013322323301-1300001220203321-0321100123123311-3310230020110000-0010222003021001-1013001322211212-0312321100011303-1320030233001132"></a>

## Direct properties — sha1_fingerprint / 132030202230 / 3

<a id="canonical-2202012233001101-3131012212320120-2212030130212331-2331310220220331-0001230232101201-1221122111131023-0221201300130200-3003311132232222"></a>

<a id="canonical-1211133212331011-0332003003120220-3201022012230221-3202112322012022-2023330222311311-1321021332200323-0301111311201002-1121302110323201"></a>

## fingerprint property — sha1_fingerprint / 132030202230 / 4

Type: `"string"`. Computed.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 40,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-1031021330112031-3011312120310013-0000101021213311-2010220133213113-2123222002220102-1221032003000002-3030323031102132-2203310301111022"></a>

## Next pages — sha1_fingerprint / 132030202230 / 5

- [primary.default_rr_set_group.sshfp_record.values](data-sources--dns_zone--reference--group-002.md#canonical-3202003222120313-3231001131322220-1302303203232200-0333301211221330-1312101031123012-0333320322003303-3011001223203333-1003121222013103)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1033003133000023-0122010033130000-1001333131130010-2300232320011323-1020123031313230-0023020332032322-0022313002020000-2021012001232012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222213300033130-0333232103332310-1031100110222221-0111133331202103-2032133203313020-3032131020020120-2321232123330001-2301220010233210"></a>

## primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint — sha256_fingerprint / 023222020330 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.sshfp_record](data-sources--dns_zone--reference--group-002.md#canonical-2131213003100201-0112010012001002-0013220210323230-0322333202002311-1323221222033032-0212121113221221-0222110233100122-1122121333301021)
- [primary.default_rr_set_group.sshfp_record.values](data-sources--dns_zone--reference--group-002.md#canonical-3202003222120313-3231001131322220-1302303203232200-0333301211221330-1312101031123012-0333320322003303-3011001223203333-1003121222013103)
- primary.default_rr_set_group.sshfp_record.values.sha256_fingerprint

<a id="canonical-1201210333203213-2223231311222211-0201320320001023-3110022121030130-2302312101110210-0301333010330033-0230022230112212-2210011303031321"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 fingerprint.

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

<a id="canonical-3120323330232022-3212133101303302-2012221133323002-3113133021102130-0312002112223331-2311202131200320-1101101000023002-2030211233110231"></a>

## Direct properties — sha256_fingerprint / 023222020330 / 3

<a id="canonical-1113020303311333-0132332222011212-0232212100333233-3110110303222112-2023213310022210-2300321230330113-2222032010303302-0030102310221231"></a>

<a id="canonical-0031132120231333-2332013332123323-2032113112001310-0330032131000231-0021023301101011-2021102103010332-3301210320133002-3002003203312300"></a>

## fingerprint property — sha256_fingerprint / 023222020330 / 4

Type: `"string"`. Computed.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 64,
    "pattern": "^[0-9a-fA-F]+$"
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
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-3001323011300311-0121021210033302-2013010203002012-1011313032200131-0200120000110320-2010310030110320-1110021300120320-3331101132000122"></a>

## Next pages — sha256_fingerprint / 023222020330 / 5

- [primary.default_rr_set_group.sshfp_record.values](data-sources--dns_zone--reference--group-002.md#canonical-3202003222120313-3231001131322220-1302303203232200-0333301211221330-1312101031123012-0333320322003303-3011001223203333-1003121222013103)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2101121000001100-2301010023202222-0201001332133023-3100101303202310-3001221233210201-2132313122010110-3100111130011132-3022131111200230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303333331303113-3201001233200302-2121030232203133-0311022231001321-3210000013033123-1022311332001210-2233030011330301-3331200023210130"></a>

## primary.default_rr_set_group.tlsa_record — tlsa_record / 230012000331 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.tlsa_record

<a id="canonical-0032202110310221-3032213010113022-0333121311132030-3203010222233313-0303021332131032-3222131113020221-0011020030321203-0301021200311302"></a>

Type: `"single"`. Computed.

Configuration parameter for tlsa record.

Upstream description:

DNS TLSA Record.

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

<a id="canonical-0322113111202010-0033201122122300-3312131103103301-3113010030231202-3003313020113211-3213333133300231-0022333010200220-2001032333130301"></a>

## Direct properties — tlsa_record / 230012000331 / 3

<a id="canonical-1230213033331222-0330232100013210-3213221211033210-3211223220332020-1101010120033211-0233030231000310-0212212300203223-2210211111022110"></a>

<a id="canonical-1323230122203330-3101132111123232-2221232001301011-3133300012213120-0000322020202023-1101112313103112-0120111132230313-2202321131011200"></a>

## name property — tlsa_record / 230012000331 / 4

Type: `"string"`. Computed.

TLSA Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-0120011031211132-3230333303020232-1031031230021310-0113032230023310-3311022011210212-1033010011101331-3112200132321022-2112113100132111): complete subsection reference.

<a id="canonical-3133300332031001-2110003223003301-3203001000220003-2232330132332123-2332231213221233-0131100010213120-0223221113323212-2132221011330033"></a>

## Next pages — tlsa_record / 230012000331 / 5

- [primary.default_rr_set_group.tlsa_record.values](data-sources--dns_zone--reference--group-002.md#canonical-0120011031211132-3230333303020232-1031031230021310-0113032230023310-3311022011210212-1033010011101331-3112200132321022-2112113100132111)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0120011031211132-3230333303020232-1031031230021310-0113032230023310-3311022011210212-1033010011101331-3112200132321022-2112113100132111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320302201211120-1301023202221322-1202121110221232-3330220331331212-1112201032130123-1131203123333011-2111120133321001-2112300120130103"></a>

## primary.default_rr_set_group.tlsa_record.values — values / 233313003332 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [primary.default_rr_set_group.tlsa_record](data-sources--dns_zone--reference--group-002.md#canonical-2101121000001100-2301010023202222-0201001332133023-3100101303202310-3001221233210201-2132313122010110-3100111130011132-3022131111200230)
- primary.default_rr_set_group.tlsa_record.values

<a id="canonical-3002001033031230-0223021112020031-2031230013011312-2212321010023312-2130233001011032-3333311220331311-2220322001301203-3031032102121020"></a>

Type: `"list"`. Computed.

TLSA Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0000232320323323-1200113311011323-0303332103001300-0130021302113312-1000223201020112-0331313131022010-0111133020313032-1231133120032331"></a>

## Direct properties — values / 233313003332 / 3

<a id="canonical-0101102321003212-3113212202302201-2110303021101311-3303012000332311-0311220231312102-0132122221022212-0211013111023301-1111100322111200"></a>

<a id="canonical-2133211030333212-1003222000333211-3012212332201300-0200103312032323-2211110003233223-0002002202033133-2110022213032303-0011220322033122"></a>

## certificate_association_data property — values / 233313003332 / 4

Type: `"string"`. Computed.

The actual data to be matched given the settings of the other fields.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0123210230200102-3213021213222203-2002331032101201-1123001303221233-3323100202113322-3130212113311100-3033330032301210-1231012020213333"></a>

<a id="canonical-3302121233011203-0103223101303320-0201112323002320-1110022101130022-3330213003010202-2111302011133321-1220230221131011-3321123122222001"></a>

## certificate_usage property — values / 233313003332 / 5

Type: `"string"`. Computed.

\[Enum:
CertificateAuthorityConstraint|ServiceCertificateConstraint|TrustAnchorAssertion|DomainIssuedCertificate\]
&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint - ServiceCertificateConstraint:
Service Certificate Constraint - TrustAnchorAssertion: Trust Anchor Assertion -
DomainIssuedCertificate: Domain Issued Certificate. Possible values are
\`CertificateAuthorityConstraint\`, \`ServiceCertificateConstraint\`, \`TrustAnchorAssertion\`,
\`DomainIssuedCertificate\`. Defaults to \`CertificateAuthorityConstraint\`.

Upstream description:

&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint

&#8203;- ServiceCertificateConstraint: Service Certificate Constraint

&#8203;- TrustAnchorAssertion: Trust Anchor Assertion

&#8203;- DomainIssuedCertificate: Domain Issued Certificate.

Receipt-pinned upstream constraints:

```json
{
  "default": "CertificateAuthorityConstraint",
  "enum": [
    "CertificateAuthorityConstraint",
    "ServiceCertificateConstraint",
    "TrustAnchorAssertion",
    "DomainIssuedCertificate"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1123030022332000-1322302232022012-3221001233221210-3233322123131031-2113030131111212-1200000322300232-2221233112232322-1131020233000320"></a>

<a id="canonical-3132220033033130-0001232210011031-2202220123231120-1231303311312221-2102321313230201-1212111311100113-2230033310132312-0100303100331032"></a>

## matching_type property — values / 233313003332 / 6

Type: `"string"`. Computed.

\[Enum: NoHash|SHA256|SHA512\] - NoHash: No Hash - SHA256: SHA-256 - SHA512: SHA-512. Possible
values are \`NoHash\`, \`SHA256\`, \`SHA512\`. Defaults to \`NoHash\`.

Upstream description:

&#8203;- NoHash: No Hash

&#8203;- SHA256: SHA-256

&#8203;- SHA512: SHA-512.

Receipt-pinned upstream constraints:

```json
{
  "default": "NoHash",
  "enum": [
    "NoHash",
    "SHA256",
    "SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1212311331323330-3121101311123002-3322032023130310-1333000010211313-0211023201030232-2111022020011002-3320313123203303-1302133303211202"></a>

<a id="canonical-1103121130100021-2133320313311000-2032321300313111-1301011103031032-1232303201220203-2320213021233132-3213113313102300-3110212323212012"></a>

## selector property — values / 233313003332 / 7

Type: `"string"`. Computed.

\[Enum: FullCertificate|UseSubjectPublicKey\] - FullCertificate: Full Certificate -
UseSubjectPublicKey: Use Subject Public Key. Possible values are \`FullCertificate\`,
\`UseSubjectPublicKey\`. Defaults to \`FullCertificate\`.

Upstream description:

&#8203;- FullCertificate: Full Certificate

&#8203;- UseSubjectPublicKey: Use Subject Public Key.

Receipt-pinned upstream constraints:

```json
{
  "default": "FullCertificate",
  "enum": [
    "FullCertificate",
    "UseSubjectPublicKey"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2031213312322230-3001212130200200-3201113302113132-0133101201111003-2302003310230320-0122223120310022-1203210220302022-3030232111130230"></a>

## Next pages — values / 233313003332 / 8

- [primary.default_rr_set_group.tlsa_record](data-sources--dns_zone--reference--group-002.md#canonical-2101121000001100-2301010023202222-0201001332133023-3100101303202310-3001221233210201-2132313122010110-3100111130011132-3022131111200230)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2023003321332312-0023132123210230-1102330021223031-3112003211020100-3031102022010133-0321133333232212-0313320223232033-0232003031030313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112300033121301-0320131210013123-0213330031301222-3111301311122323-2232023211110233-3202321333221200-3221300132122203-1232022033300100"></a>

## primary.default_rr_set_group.txt_record — txt_record / 332111332011 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- primary.default_rr_set_group.txt_record

<a id="canonical-0113213300213103-1012321032121210-0320323211322311-2113133002212131-2200322100201330-2001332233103233-3303333311332233-0121112232302110"></a>

Type: `"single"`. Computed.

DNSTXTResourceRecord.

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

<a id="canonical-0332013112220110-1323320031233100-1323221302012131-2332101022322011-2311332120130013-1031230333333002-3323002232330101-3121210013231312"></a>

## Direct properties — txt_record / 332111332011 / 3

<a id="canonical-3132322012012100-2213013323323121-2113100213332123-3223330323222110-1300133212023222-0012333100203032-3211111032012332-2333213122021123"></a>

<a id="canonical-1031200032120013-3233130311330202-1021123111223232-2332312020310100-0022312001123010-3002003220322030-1333331110232032-3213001233103002"></a>

## name property — txt_record / 332111332011 / 4

Type: `"string"`. Computed.

TXT Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-0232301130321032-1123011101323311-1103103322303022-1212000212321333-0213003222303132-3211223133203202-1111111310212333-3023300210011323"></a>

<a id="canonical-0133020023030111-3131202220032121-3203022032332122-1321012210312303-1130010021201021-0303330103232230-3112013312323032-3113230312311031"></a>

## values property — txt_record / 332111332011 / 5

Type: `["list", "string"]`. Computed.

Text. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2223213203002202-3012300233003312-3331010232131210-2022210211213222-1131201103111103-2302220020211020-1302133101120303-1301202220203213"></a>

## Next pages — txt_record / 332111332011 / 6

- [primary.default_rr_set_group](data-sources--dns_zone--reference--group-001.md#canonical-2333233213103330-0302202110210122-3330012220203203-2012221330313213-3020122001023320-2112110122003110-2303022223102003-1030220023110111)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1013032121121211-0200310230111302-2302123231122131-1132011211233131-3233303200303110-1001330123313000-0122110323001111-2131223301100220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210221122033313-1310330311103222-3003130012310120-0330020300103113-0233223322231102-0030201303133333-2210200121330220-0202212010202300"></a>

## primary.default_soa_parameters — default_soa_parameters / 323200220333 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- primary.default_soa_parameters

<a id="canonical-1200102321322311-3223202022210202-3010323322310100-2013003011231200-1032331233311132-0021012320100102-3113332223310121-0101330133331122"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default soa parameters.

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

<a id="canonical-0301022000102232-2320332202121010-3033223203223013-3301023231233220-3202033100330311-2112212232130322-3120000111203100-0333331322033332"></a>

## Direct properties — default_soa_parameters / 323200220333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301121103223020-3100321120022013-0220312231301303-2013021023120132-0100101011001303-0201321223130202-0103013121331202-3102202202201032"></a>

## Next pages — default_soa_parameters / 323200220333 / 4

- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0301132310232123-0100020002231220-0220210220113100-2233333133311311-2013203003202010-1111221031021320-3202202231022031-3321200322011102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212101102011133-2120120300102130-3122112021211131-0320020203220010-2300031321111022-3302300322210333-2123232222220302-3110320211233312"></a>

## primary.dnssec_mode — dnssec_mode / 300121103213 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- primary.dnssec_mode

<a id="canonical-1003323000301330-1322020032333013-1121022131303132-2311101012002202-0130133003222100-0133113113112003-2121213120121032-2021301010121112"></a>

Type: `"single"`. Computed.

DNSSEC Mode.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mode": "[\"disable\",\"enable\"]"
}
```

<a id="canonical-1310300202003101-1233302232003101-1202130132210122-2222020011213333-0013130003213103-0003033302303023-0320122010300122-3301133311331122"></a>

## Direct properties — dnssec_mode / 300121103213 / 3

- [disable_spec](data-sources--dns_zone--reference--group-002.md#canonical-0203303200332322-0121103330333203-3311110011232221-2101231000322133-2102222111020223-3132302232302210-1213103200010130-0300321120212112): complete subsection reference.

- [enable](data-sources--dns_zone--reference--group-002.md#canonical-1221121130033221-2123222101202023-0201332012131001-0303022301203123-0122210132133332-2203321123031100-2100301320102120-1203023103200100): complete subsection reference.

<a id="canonical-2111032033200323-3200033303302132-0112211113220010-2210020123213232-0000312230001021-0331222311023000-2332010212112031-0200312121032102"></a>

## Next pages — dnssec_mode / 300121103213 / 4

- [primary.dnssec_mode.disable_spec](data-sources--dns_zone--reference--group-002.md#canonical-0203303200332322-0121103330333203-3311110011232221-2101231000322133-2102222111020223-3132302232302210-1213103200010130-0300321120212112)
- [primary.dnssec_mode.enable](data-sources--dns_zone--reference--group-002.md#canonical-1221121130033221-2123222101202023-0201332012131001-0303022301203123-0122210132133332-2203321123031100-2100301320102120-1203023103200100)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0203303200332322-0121103330333203-3311110011232221-2101231000322133-2102222111020223-3132302232302210-1213103200010130-0300321120212112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331330102020102-1321020103010201-2231022231131213-3113011333221332-1332233020012332-0300310200120030-3320013222302010-2102111112122202"></a>

## primary.dnssec_mode.disable_spec — disable_spec / 033303321233 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.dnssec_mode](data-sources--dns_zone--reference--group-002.md#canonical-0301132310232123-0100020002231220-0220210220113100-2233333133311311-2013203003202010-1111221031021320-3202202231022031-3321200322011102)
- primary.dnssec_mode.disable_spec

<a id="canonical-0133101022110312-1012230221200332-0112223131330022-3122233333232332-0303220302221003-1230223103002022-3112210203012202-3333110013323211"></a>

Type: `["object", {}]`. Computed.

Enable this option

<a id="canonical-1312213120231121-0123331323330300-0211311322110301-1303330033100223-2230112213232310-0011210031112213-1133203132113110-1221201201231010"></a>

## Direct properties — disable_spec / 033303321233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323322111011201-0111031302310021-2311110203332032-1313312022010030-1220211322303210-0233020202020100-3033221103011110-0233213203322210"></a>

## Next pages — disable_spec / 033303321233 / 4

- [primary.dnssec_mode](data-sources--dns_zone--reference--group-002.md#canonical-0301132310232123-0100020002231220-0220210220113100-2233333133311311-2013203003202010-1111221031021320-3202202231022031-3321200322011102)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1221121130033221-2123222101202023-0201332012131001-0303022301203123-0122210132133332-2203321123031100-2100301320102120-1203023103200100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112121133020101-1232210021321310-2300202011330320-2233313033131023-0011022102002221-3320320232300022-2023313132023213-1002103201332011"></a>

## primary.dnssec_mode.enable — enable / 321021323111 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.dnssec_mode](data-sources--dns_zone--reference--group-002.md#canonical-0301132310232123-0100020002231220-0220210220113100-2233333133311311-2013203003202010-1111221031021320-3202202231022031-3321200322011102)
- primary.dnssec_mode.enable

<a id="canonical-1021202202023231-1203202332111212-0003030323332333-2100202023101323-3223302330203000-3231022021132313-1010101012320311-0033321331223221"></a>

Type: `["object", {}]`. Computed.

Enable. DNSSEC enable.

Upstream description:

DNSSEC enable.

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

<a id="canonical-3230313020313201-1023211232301223-1133210232110211-0033303330123202-3022331012300001-3300233223313123-2202223002310001-2231121230003020"></a>

## Direct properties — enable / 321021323111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223322130022220-0103231133200122-1232201130330120-2032003131230100-3112332231211301-0103201220200111-0112220200213122-3001023312030121"></a>

## Next pages — enable / 321021323111 / 4

- [primary.dnssec_mode](data-sources--dns_zone--reference--group-002.md#canonical-0301132310232123-0100020002231220-0220210220113100-2233333133311311-2013203003202010-1111221031021320-3202202231022031-3321200322011102)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321302230002300-0123313012322121-0231023323030011-2300222332033323-0101121133201303-2033102022133111-3232121332113200-1330230102100220"></a>

## primary.rr_set_group — rr_set_group / 101203011331 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- primary.rr_set_group

<a id="canonical-2231202303211010-2021101001111030-2013023101223233-3321213130020213-3332320300011023-2121211132011013-2312031132210200-1131031002133031"></a>

Type: `"list"`. Computed.

Create and manage set groups, and resource record sets within them, x-VES-I/O-managed set is managed
by F5.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
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
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-1000101300212330-1012332012330303-2102120021110222-0211103113002223-2001032111031120-3000233232321111-0132210111301133-3302333023123233"></a>

## Direct properties — rr_set_group / 101203011331 / 3

- [metadata](data-sources--dns_zone--reference--group-002.md#canonical-2300002332200332-3300012320121033-1323312230331320-0222210111132311-1030011331333111-1212002231231222-1321113113133330-0222303310223013): complete subsection reference.

- [rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302): complete subsection reference.

<a id="canonical-1222333131020120-1222203311320201-0213002110002011-1113300101120322-2112032211000333-0202013231213203-3211212012330323-1202130110023320"></a>

## Next pages — rr_set_group / 101203011331 / 4

- [primary.rr_set_group.metadata](data-sources--dns_zone--reference--group-002.md#canonical-2300002332200332-3300012320121033-1323312230331320-0222210111132311-1030011331333111-1212002231231222-1321113113133330-0222303310223013)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2300002332200332-3300012320121033-1323312230331320-0222210111132311-1030011331333111-1212002231231222-1321113113133330-0222303310223013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010113313101123-3113222020321022-2103203033123130-3110301311311222-3211213133003111-2121331132132210-1031223303312000-2111120000221023"></a>

## primary.rr_set_group.metadata — metadata / 110022313102 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- primary.rr_set_group.metadata

<a id="canonical-1330003303230112-2023000311003232-0313021100232031-2130100211300330-1322102200310132-0111211310022101-3331032000001001-1003333223010012"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-0122022112303300-0110123112132030-0321103100312000-0022333130020331-3002132032033213-3103123301231331-3000212332211233-1121023210131032"></a>

## Direct properties — metadata / 110022313102 / 3

<a id="canonical-1322330123203331-2320210002222202-3221301311020012-2330303222321212-1031011322100311-2110232000022102-0032000113233121-0303020233122311"></a>

<a id="canonical-0100032000311022-1013212211130301-2203031200223133-0132203223223202-3230020330121122-3323213010032131-3311223203311112-2121021032113022"></a>

## description_spec property — metadata / 110022313102 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2202023131100022-1103302232221201-3011011233003132-1222131313102331-0013123113203032-1301132213112102-3002233120002323-1130303100213221"></a>

<a id="canonical-2231122221320102-1032332333113300-3033303303101200-3322011021132320-2300221222003121-2333103112110302-0303100330232100-1033130213130130"></a>

## name property — metadata / 110022313102 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0220333030220233-1012130331231031-3231210230023232-2321321020003311-0011210310130002-1103020301233202-1321222202010323-3023113313203132"></a>

## Next pages — metadata / 110022313102 / 6

- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002210200021330-3201320132220312-0332320022210122-2323221213203120-0000310010321223-3311100001121000-2110030310130221-2102031330003013"></a>

## primary.rr_set_group.rr_set — rr_set / 000312230302 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- primary.rr_set_group.rr_set

<a id="canonical-3313212012322332-1132311103230202-1103220203333130-3310221013011303-1033220103302101-1021000000330033-1210313021213300-0031330331100211"></a>

Type: `"list"`. Computed.

Resource Record Sets. Collection of DNS resource record sets.

Upstream description:

Collection of DNS resource record sets.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50000,
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
    "ves.io.schema.rules.repeated.max_items": "50000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "50000"
  }
}
```

<a id="canonical-0223213321020232-1332201301000030-1032123033200312-2123202013032013-0232020231110002-0301330123223302-0302301022330131-0331202222202103"></a>

## Direct properties — rr_set / 000312230302 / 3

- [a_record](data-sources--dns_zone--reference--group-002.md#canonical-2133222331203223-0313012330312212-3123323310001222-2102203030233020-3301203301220002-1031033311120221-0000331011100133-0320223033201232): complete subsection reference.

- [aaaa_record](data-sources--dns_zone--reference--group-002.md#canonical-0300332203033212-2102110000033322-3321232013303310-3001001213000213-0230010313131203-0232112013213231-3212033301122213-3233101333323002): complete subsection reference.

- [afsdb_record](data-sources--dns_zone--reference--group-002.md#canonical-1300113000100303-0103301013032300-2333102222023331-1311010301000230-1312131033332333-3001303320321310-2023130101302221-1203303330012333): complete subsection reference.

- [alias_record](data-sources--dns_zone--reference--group-002.md#canonical-2232212032121023-3022121321313313-0013031312032122-3303023321031011-2033003002102300-0222023203002011-3310121213120113-3120313012020110): complete subsection reference.

- [caa_record](data-sources--dns_zone--reference--group-002.md#canonical-3021102303223021-0311330012113303-1211113233122120-0203302301302020-2223212013012002-2100133123121201-2130010021021103-0021000303323011): complete subsection reference.

- [cds_record](data-sources--dns_zone--reference--group-002.md#canonical-3323310232223321-3112000010212222-2211023330131221-2112123001002122-3232100332311223-1120221212020121-3101311000103133-3320301013333121): complete subsection reference.

- [cert_record](data-sources--dns_zone--reference--group-002.md#canonical-2033130131100231-1310103201133222-0010203201322133-0113003033112230-3113331000213213-0323122330220111-1301022323030020-3123313211203310): complete subsection reference.

- [cname_record](data-sources--dns_zone--reference--group-002.md#canonical-3000110033203111-2312032311212322-2103322303022131-1123102210000211-0312322123020323-2222132212012111-1000102011220113-3123113121110030): complete subsection reference.

<a id="canonical-2000320222210232-1300311103033231-3231232123032330-3001223121001231-2331332212302100-1203320111201203-0121133303013112-1110222233321033"></a>

<a id="canonical-0013023330022121-2221010231221323-2332233220200113-0211132033310030-0303030003022212-2031121000213301-0311323323000112-0021323121321331"></a>

## description_spec property — rr_set / 000312230302 / 4

Type: `"string"`. Computed.

Comment. Human-readable description text

- [ds_record](data-sources--dns_zone--reference--group-003.md#canonical-0003002010030223-3031302330310202-3322320001011330-0003010123133202-3023120301200122-3120120121330033-2203311132332022-2031233200112331): complete subsection reference.

- [eui48_record](data-sources--dns_zone--reference--group-003.md#canonical-1300212020003230-1113311332101220-1132330003331301-3333111020100102-3003313132131220-0122230031303231-2203031223032120-2332332003333102): complete subsection reference.

- [eui64_record](data-sources--dns_zone--reference--group-003.md#canonical-0003303213311201-3003200303313223-1121023300112121-2002211020100233-1112212221032122-1331101211011320-2003332313321333-2110200232323101): complete subsection reference.

- [lb_record](data-sources--dns_zone--reference--group-003.md#canonical-1230020033200103-1330330300123130-3220203300231002-1211332020202310-3103022120112000-2013013012212211-2200123220201232-2330030320121121): complete subsection reference.

- [loc_record](data-sources--dns_zone--reference--group-003.md#canonical-3103002133110011-3300030122231322-1033330212002032-1120030230101113-1221213310123022-1300023322200233-0112012020111212-1320210023313010): complete subsection reference.

- [mx_record](data-sources--dns_zone--reference--group-003.md#canonical-2232020013010222-3011001301213203-0001030100231211-0233002301012011-0102021001102022-2003330212132112-3220310312321133-1330303033320022): complete subsection reference.

- [naptr_record](data-sources--dns_zone--reference--group-003.md#canonical-2130200302222231-0201323023000030-2231013210302021-3320130221100012-2012002232303013-1033323023120320-3201231132301223-2120220121110010): complete subsection reference.

- [ns_record](data-sources--dns_zone--reference--group-003.md#canonical-1301020100313003-2113130112303333-0222200332023331-1011211333020200-3333320312030230-2202333330201232-2032210300303202-3111001112202203): complete subsection reference.

- [ptr_record](data-sources--dns_zone--reference--group-003.md#canonical-2211303020222311-3310323332010203-0232301233301110-3121123122223302-3103033101210322-2020123011002111-3032313210101120-2001120323033330): complete subsection reference.

- [srv_record](data-sources--dns_zone--reference--group-003.md#canonical-0223103023112010-0122311332230213-0020010203033121-2323203031213310-1321013110303211-1310102220022031-3130133320120313-1333011010122312): complete subsection reference.

- [sshfp_record](data-sources--dns_zone--reference--group-003.md#canonical-1320201233022013-2010302020223210-3223212010011202-0302303001331232-1022311200333221-3030011233031331-2201220131002312-2233122232100332): complete subsection reference.

- [tlsa_record](data-sources--dns_zone--reference--group-003.md#canonical-1120111220133012-1033321320120331-3101013303033222-2303001213103323-1230312203202132-1302200000132221-0213201322100310-2310212123010121): complete subsection reference.

<a id="canonical-0113323032231011-0223333222302211-3322300003130030-0323113011131110-0313123301200332-3032113230120122-0323110023212100-0112221033132131"></a>

<a id="canonical-0313312110311002-1211302131233132-2020302331203321-1223211123120103-0113210003010322-0300200223221032-0203013133003331-2002021122102101"></a>

## TTL property — rr_set / 000312230302 / 5

Type: `"number"`. Computed.

Time to live. Time-to-live duration in seconds

Upstream description:

Time-to-live duration in seconds

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 60
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

- [txt_record](data-sources--dns_zone--reference--group-003.md#canonical-2112000100003211-0010122023020330-2302300331333210-3123120333012012-1313313003312213-2120120330311231-1033321003121122-0030131112120013): complete subsection reference.

<a id="canonical-3232031131022122-0301030030103321-0112330320021133-2030103123121110-2220302010303321-3000300302001202-3232323212032113-1211320233311000"></a>

## Next pages — rr_set / 000312230302 / 6

- [primary.rr_set_group.rr_set.a_record](data-sources--dns_zone--reference--group-002.md#canonical-2133222331203223-0313012330312212-3123323310001222-2102203030233020-3301203301220002-1031033311120221-0000331011100133-0320223033201232)
- [primary.rr_set_group.rr_set.aaaa_record](data-sources--dns_zone--reference--group-002.md#canonical-0300332203033212-2102110000033322-3321232013303310-3001001213000213-0230010313131203-0232112013213231-3212033301122213-3233101333323002)
- [primary.rr_set_group.rr_set.afsdb_record](data-sources--dns_zone--reference--group-002.md#canonical-1300113000100303-0103301013032300-2333102222023331-1311010301000230-1312131033332333-3001303320321310-2023130101302221-1203303330012333)
- [primary.rr_set_group.rr_set.alias_record](data-sources--dns_zone--reference--group-002.md#canonical-2232212032121023-3022121321313313-0013031312032122-3303023321031011-2033003002102300-0222023203002011-3310121213120113-3120313012020110)
- [primary.rr_set_group.rr_set.caa_record](data-sources--dns_zone--reference--group-002.md#canonical-3021102303223021-0311330012113303-1211113233122120-0203302301302020-2223212013012002-2100133123121201-2130010021021103-0021000303323011)
- [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-3323310232223321-3112000010212222-2211023330131221-2112123001002122-3232100332311223-1120221212020121-3101311000103133-3320301013333121)
- [primary.rr_set_group.rr_set.cert_record](data-sources--dns_zone--reference--group-002.md#canonical-2033130131100231-1310103201133222-0010203201322133-0113003033112230-3113331000213213-0323122330220111-1301022323030020-3123313211203310)
- [primary.rr_set_group.rr_set.cname_record](data-sources--dns_zone--reference--group-002.md#canonical-3000110033203111-2312032311212322-2103322303022131-1123102210000211-0312322123020323-2222132212012111-1000102011220113-3123113121110030)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-003.md#canonical-0003002010030223-3031302330310202-3322320001011330-0003010123133202-3023120301200122-3120120121330033-2203311132332022-2031233200112331)
- [primary.rr_set_group.rr_set.eui48_record](data-sources--dns_zone--reference--group-003.md#canonical-1300212020003230-1113311332101220-1132330003331301-3333111020100102-3003313132131220-0122230031303231-2203031223032120-2332332003333102)
- [primary.rr_set_group.rr_set.eui64_record](data-sources--dns_zone--reference--group-003.md#canonical-0003303213311201-3003200303313223-1121023300112121-2002211020100233-1112212221032122-1331101211011320-2003332313321333-2110200232323101)
- [primary.rr_set_group.rr_set.lb_record](data-sources--dns_zone--reference--group-003.md#canonical-1230020033200103-1330330300123130-3220203300231002-1211332020202310-3103022120112000-2013013012212211-2200123220201232-2330030320121121)
- [primary.rr_set_group.rr_set.loc_record](data-sources--dns_zone--reference--group-003.md#canonical-3103002133110011-3300030122231322-1033330212002032-1120030230101113-1221213310123022-1300023322200233-0112012020111212-1320210023313010)
- [primary.rr_set_group.rr_set.mx_record](data-sources--dns_zone--reference--group-003.md#canonical-2232020013010222-3011001301213203-0001030100231211-0233002301012011-0102021001102022-2003330212132112-3220310312321133-1330303033320022)
- [primary.rr_set_group.rr_set.naptr_record](data-sources--dns_zone--reference--group-003.md#canonical-2130200302222231-0201323023000030-2231013210302021-3320130221100012-2012002232303013-1033323023120320-3201231132301223-2120220121110010)
- [primary.rr_set_group.rr_set.ns_record](data-sources--dns_zone--reference--group-003.md#canonical-1301020100313003-2113130112303333-0222200332023331-1011211333020200-3333320312030230-2202333330201232-2032210300303202-3111001112202203)
- [primary.rr_set_group.rr_set.ptr_record](data-sources--dns_zone--reference--group-003.md#canonical-2211303020222311-3310323332010203-0232301233301110-3121123122223302-3103033101210322-2020123011002111-3032313210101120-2001120323033330)
- [primary.rr_set_group.rr_set.srv_record](data-sources--dns_zone--reference--group-003.md#canonical-0223103023112010-0122311332230213-0020010203033121-2323203031213310-1321013110303211-1310102220022031-3130133320120313-1333011010122312)
- [primary.rr_set_group.rr_set.sshfp_record](data-sources--dns_zone--reference--group-003.md#canonical-1320201233022013-2010302020223210-3223212010011202-0302303001331232-1022311200333221-3030011233031331-2201220131002312-2233122232100332)
- [primary.rr_set_group.rr_set.tlsa_record](data-sources--dns_zone--reference--group-003.md#canonical-1120111220133012-1033321320120331-3101013303033222-2303001213103323-1230312203202132-1302200000132221-0213201322100310-2310212123010121)
- [primary.rr_set_group.rr_set.txt_record](data-sources--dns_zone--reference--group-003.md#canonical-2112000100003211-0010122023020330-2302300331333210-3123120333012012-1313313003312213-2120120330311231-1033321003121122-0030131112120013)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2133222331203223-0313012330312212-3123323310001222-2102203030233020-3301203301220002-1031033311120221-0000331011100133-0320223033201232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223113110010333-3313310300000303-3330020122323211-1102031233033333-1101223202322132-0321313231322111-1232310302300333-0331003223010311"></a>

## primary.rr_set_group.rr_set.a_record — a_record / 212303201133 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.a_record

<a id="canonical-3011212320111000-3202333013120221-0021212111332023-2032123232032111-0201112310120223-3310200123333112-0212223220022001-2001113312010031"></a>

Type: `"single"`. Computed.

DNSAResourceRecord. A Records

Upstream description:

A Records

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

<a id="canonical-0021213132010121-1322113133032112-1310321223332031-2113220303021331-1320030212312330-2311323332322002-2231031110222010-2011322010203013"></a>

## Direct properties — a_record / 212303201133 / 3

<a id="canonical-3213102231313033-0303331322121002-3300312033112130-0201111312112012-1003032023110211-2223300333323120-1211322302312013-2021322230330230"></a>

<a id="canonical-2321232013202010-1330001131031120-1110123303313111-3031102231322001-3100120320032202-2032110011130230-2012321000130022-0202031031133230"></a>

## name property — a_record / 212303201133 / 4

Type: `"string"`. Computed.

Record name, please provide only the specific subdomain or record name without the base domain.

Upstream description:

A Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-0112323113010023-3023131013113031-3102131201322031-2233113002200212-3130333332313022-3002220000221331-1010312303011011-1023021030113300"></a>

<a id="canonical-3321002300320120-3310212122110111-1322230000123113-2220120033021230-0331311302333001-2230203331303213-3330302301322313-3013021231322322"></a>

## values property — a_record / 212303201133 / 5

Type: `["list", "string"]`. Computed.

IPv4 Addresses. A valid IPv4 address, for example: 192.0.2.242.

Upstream description:

A valid IPv4 address, for example: 192.0.2.242.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0302232202320120-1120011210013232-1212130012110001-1101201331222331-3103210023001230-0001332133320133-1130330112330031-0120023200303203"></a>

## Next pages — a_record / 212303201133 / 6

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0300332203033212-2102110000033322-3321232013303310-3001001213000213-0230010313131203-0232112013213231-3212033301122213-3233101333323002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033200020223202-0301011012120212-2120120220301123-3232303323202301-1110200133113122-2121201000020030-2112202012202103-0133232212231122"></a>

## primary.rr_set_group.rr_set.aaaa_record — aaaa_record / 112301113201 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.aaaa_record

<a id="canonical-2023331012311212-2202230303010330-2220121231320033-0011233220203122-1222030112301320-1131232001323310-2311111111133130-0132120210120321"></a>

Type: `"single"`. Computed.

Configuration parameter for aaaa record.

Upstream description:

RecordSet for AAAA Records.

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

<a id="canonical-2323233333101332-3010023200110002-2200020202111030-1322330321133230-1321332133130221-2310330332221103-0100323322132203-2313102013122321"></a>

## Direct properties — aaaa_record / 112301113201 / 3

<a id="canonical-0132130000323000-0120210230321113-0332300303303101-2131331032001110-0103203311103130-3131232113100322-3103033011222301-3203031100330310"></a>

<a id="canonical-0332123121322020-1103221032231300-0121012112031011-0300113000201311-3221221112002221-3303013133101231-3232201130023212-0030320213320030"></a>

## name property — aaaa_record / 112301113201 / 4

Type: `"string"`. Computed.

AAAA Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-1201210031130232-1121021132332032-0100200001133112-2013220012213033-0321033333113332-2002303321201102-3103030013230103-3032011201322021"></a>

<a id="canonical-1132101131112321-0323231333231300-3030200213111221-2001212032312000-0333003131312102-1212320130123020-2330033312102313-3223013321110113"></a>

## values property — aaaa_record / 112301113201 / 5

Type: `["list", "string"]`. Computed.

IPv6 Addresses. A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Upstream description:

A valid IPv6 address, for example: 2001:0db8:85a3:0000:0000:8a2e:0370:7334.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3212202332102322-1312132100313021-1210110102200313-3330223320020030-0033302130003012-0233001232000012-1101220221302322-2232211202122022"></a>

## Next pages — aaaa_record / 112301113201 / 6

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1300113000100303-0103301013032300-2333102222023331-1311010301000230-1312131033332333-3001303320321310-2023130101302221-1203303330012333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031002120213230-1002230131233010-1222013233102002-3313131333100021-1212331221321202-0322203311032323-0320333002122001-3303010221101123"></a>

## primary.rr_set_group.rr_set.afsdb_record — afsdb_record / 033201113113 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.afsdb_record

<a id="canonical-3033302222102131-3033210012122221-1223220200233121-2000203301300123-2311103123330212-0213212012130122-0230201320011032-0231132102321330"></a>

Type: `"single"`. Computed.

Configuration parameter for afsdb record.

Upstream description:

DNS AFSDB Record.

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

<a id="canonical-2320212132310002-0233100210122303-0310120001132101-3332003121022022-3012332033111233-0122303213313001-2022310213113200-2331301030100112"></a>

## Direct properties — afsdb_record / 033201113113 / 3

<a id="canonical-0032301233120202-0033231113100112-3111330023301212-2200010130112022-0331221002213211-0033300110200210-2322302222122332-1300113202211222"></a>

<a id="canonical-1123233330011033-3230202302313202-2113010200132310-1103333303201110-1003010003333132-1000001301203220-3012330020322311-3113113122323310"></a>

## name property — afsdb_record / 033201113113 / 4

Type: `"string"`. Computed.

AFSDB Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-2301311330323221-2332100303230210-0101102100201033-1203233130220313-2030201110213221-0131201211201033-1101120020202233-3111323103331201): complete subsection reference.

<a id="canonical-2231031322303112-1203103313011110-3021230232233300-0010330120210203-3302031232012013-3200003110222030-2131311023300311-0123113131030012"></a>

## Next pages — afsdb_record / 033201113113 / 5

- [primary.rr_set_group.rr_set.afsdb_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2301311330323221-2332100303230210-0101102100201033-1203233130220313-2030201110213221-0131201211201033-1101120020202233-3111323103331201)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2301311330323221-2332100303230210-0101102100201033-1203233130220313-2030201110213221-0131201211201033-1101120020202233-3111323103331201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233013220311330-1211301002012230-3231232130113233-3212230201101223-3121022211302010-1110010302022321-0010122120233020-2220011313331113"></a>

## primary.rr_set_group.rr_set.afsdb_record.values — values / 030031130313 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.afsdb_record](data-sources--dns_zone--reference--group-002.md#canonical-1300113000100303-0103301013032300-2333102222023331-1311010301000230-1312131033332333-3001303320321310-2023130101302221-1203303330012333)
- primary.rr_set_group.rr_set.afsdb_record.values

<a id="canonical-1023001200301202-0110001232003000-2021300103333001-0123111002000032-1320201020302133-1233211310230010-2233331130101130-1032330312322202"></a>

Type: `"list"`. Computed.

AFSDB Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2020002121001010-1120233120232121-3230033332323302-0223232011311003-3312323300033100-3202010002213132-3300002233300133-1130133032111330"></a>

## Direct properties — values / 030031130313 / 3

<a id="canonical-1110231102320131-3211323223302120-0301301330103020-3102030311203001-1310333221022233-2201120001221211-1031312223130302-2000200323131022"></a>

<a id="canonical-0012022310011120-1100233230212303-2120122220110301-2233132223033220-3130113311031131-0211110223022301-2022320220201123-1121131210200223"></a>

## hostname property — values / 030031130313 / 4

Type: `"string"`. Computed.

Server name of the AFS cell database server or the DCE name server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-3310302212033311-1100003023013330-3132210322002310-0121211123201130-2020120003011121-1210013320110013-2231203210201011-3323222313231230"></a>

<a id="canonical-1010233301122003-0112202322221022-2123322202301321-1233031021101000-2022123221132301-2120232303122303-3031103232210222-0313021122022323"></a>

## subtype property — values / 030031130313 / 5

Type: `"string"`. Computed.

\[Enum: NONE|AFSVolumeLocationServer|DCEAuthenticationServer\] AFS Volume Location Server or DCE
Authentication Server. - NONE: NONE - AFSVolumeLocationServer: AFS Volume Location Server -
DCEAuthenticationServer: DCE Authentication Server. Possible values are \`NONE\`,
\`AFSVolumeLocationServer\`, \`DCEAuthenticationServer\`.

Upstream description:

AFS Volume Location Server or DCE Authentication Server.

&#8203;- NONE: NONE

&#8203;- AFSVolumeLocationServer: AFS Volume Location Server

&#8203;- DCEAuthenticationServer: DCE Authentication Server.

Receipt-pinned upstream constraints:

```json
{
  "default": "NONE",
  "enum": [
    "NONE",
    "AFSVolumeLocationServer",
    "DCEAuthenticationServer"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3103011120023213-1133301110332031-2020313012223232-1100332230012203-3030002122201033-1000012220322202-3031001013333000-0000202212133121"></a>

## Next pages — values / 030031130313 / 6

- [primary.rr_set_group.rr_set.afsdb_record](data-sources--dns_zone--reference--group-002.md#canonical-1300113000100303-0103301013032300-2333102222023331-1311010301000230-1312131033332333-3001303320321310-2023130101302221-1203303330012333)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2232212032121023-3022121321313313-0013031312032122-3303023321031011-2033003002102300-0222023203002011-3310121213120113-3120313012020110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013212011023011-0231303203331311-1220121000211310-3032120321011212-1300301212031132-3033133232010223-3030103212211020-1102223203331031"></a>

## primary.rr_set_group.rr_set.alias_record — alias_record / 210122332123 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.alias_record

<a id="canonical-3330101333113112-1201010022313322-3013003311010121-2021213321311301-1320223332310312-3313023020020102-0333222132313023-1313022000331133"></a>

Type: `"single"`. Computed.

Configuration parameter for alias record.

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

<a id="canonical-0302010322023132-2021113103230112-1010121022220102-0010010033132110-0303010032111102-0301322332203201-2332102122130231-1011310221001320"></a>

## Direct properties — alias_record / 210122332123 / 3

<a id="canonical-0121102012223122-3013201022220312-1033133222301232-0231202321131130-3233312331331022-3330023120303030-0332131132301222-3301101221303330"></a>

<a id="canonical-1202203222321213-2001230211022222-3202120303133011-2132120100021223-1323200103003033-2321013232123230-1211200032301030-0032203330212111"></a>

## value property — alias_record / 210122332123 / 4

Type: `"string"`. Computed.

Domain. A valid domain name, for example: example.com.

Upstream description:

A valid domain name, for example: example.com.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-3301130333230031-2300022010212030-0202230021012313-3133021233101032-2013320202030001-0121123203313212-2312132122223121-3111031302201203"></a>

## Next pages — alias_record / 210122332123 / 5

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-3021102303223021-0311330012113303-1211113233122120-0203302301302020-2223212013012002-2100133123121201-2130010021021103-0021000303323011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030320230121131-0333023012020200-2313023133013230-0031033001121133-3130331133310133-2031021020212233-1113232232103002-2010000112331032"></a>

## primary.rr_set_group.rr_set.caa_record — caa_record / 102121310232 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.caa_record

<a id="canonical-3330112101203321-2003331021202232-1010102323102020-3030102201031133-2032203333030330-1012301022310203-3102020010230303-2121311003313311"></a>

Type: `"single"`. Computed.

DNSCAAResourceRecord.

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

<a id="canonical-0030202233100300-3020113221030003-3310201111011122-2010302210031200-3310211003331332-2322321333212103-2133032312233320-2003001133231221"></a>

## Direct properties — caa_record / 102121310232 / 3

<a id="canonical-3203023010220211-2130302031120031-0012331023301333-3310303222120232-2113103010133130-2311001321201020-0132023102310033-1033300320133310"></a>

<a id="canonical-2120122312102303-1122211013323110-0320130203011313-3000310103210012-1330321311032110-1110210013323101-3231202023233001-0021013000301000"></a>

## name property — caa_record / 102121310232 / 4

Type: `"string"`. Computed.

CAA Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-2033002231100233-2130313021213102-0030032333030102-3301130122220233-0010121203021300-3101133012113020-1332312222131312-2301010021131131): complete subsection reference.

<a id="canonical-3030220303333231-1031030320013301-1233001311131303-3201122201313011-1033132203200022-3001010302232311-2100322030023302-2113020223233131"></a>

## Next pages — caa_record / 102121310232 / 5

- [primary.rr_set_group.rr_set.caa_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2033002231100233-2130313021213102-0030032333030102-3301130122220233-0010121203021300-3101133012113020-1332312222131312-2301010021131131)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2033002231100233-2130313021213102-0030032333030102-3301130122220233-0010121203021300-3101133012113020-1332312222131312-2301010021131131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323100221333222-3202201212110030-2203331001310100-0322000110231232-3021322232103100-0321332112300233-1320213002330301-2100103210332030"></a>

## primary.rr_set_group.rr_set.caa_record.values — values / 010230010201 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.caa_record](data-sources--dns_zone--reference--group-002.md#canonical-3021102303223021-0311330012113303-1211113233122120-0203302301302020-2223212013012002-2100133123121201-2130010021021103-0021000303323011)
- primary.rr_set_group.rr_set.caa_record.values

<a id="canonical-2233002323122031-0323113231103233-3332300130230020-2301012111103101-0202213230312031-0022032223213103-2220221031331222-2222213023101033"></a>

Type: `"list"`. Computed.

CAA Record Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

<a id="canonical-0332010312311033-1000020332313323-0230201222132332-3202100000023200-2022333122233231-1221111033321312-1332112303220010-2230112313311101"></a>

## Direct properties — values / 010230010201 / 3

<a id="canonical-2211321133212110-1031313232033132-3103331013220112-0303210120233121-0301021131013032-3201001113121002-3021012212122312-1232103220103322"></a>

<a id="canonical-1311300032332103-2133313021000221-1102330032003032-3331322031010102-1201021013331231-0323202012210022-0011201013103031-3220133100213330"></a>

## flags property — values / 010230010201 / 4

Type: `"number"`. Computed.

Flag should be an integer between 0 and 255.

Upstream description:

This flag should be an integer between 0 and 255.

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
    "minimum": 0
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

<a id="canonical-1101210323223113-1331223213132323-1322200301010110-2000012101132313-2122322123301211-2023001023323010-2332013021031110-3221032200302220"></a>

<a id="canonical-2110000131032021-2211110312303201-3020322033332033-3201112212032023-3201231033102020-0001023033131313-1321303323323320-1201200111200303"></a>

## tag property — values / 010230010201 / 5

Type: `"string"`. Computed.

\[Enum: issue|issuewild|iodef\] Tag. Tag for categorization and filtering. Possible values are
\`issue\`, \`issuewild\`, \`iodef\`.

Upstream description:

Tag for categorization and filtering

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "issue",
    "issuewild",
    "iodef"
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
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"issue\\\", \\\"issuewild\\\", \\\"iodef\\\"]"
  }
}
```

<a id="canonical-0230313221211321-3323301030103232-1121102112023002-3201100201022111-2333133220013211-3120020021013301-2311131201123333-3220101303303333"></a>

<a id="canonical-0031122311212001-0303321111021331-2130302130121013-2002230022310331-1103123121031213-0200230022302100-3002023001322121-3012320310032112"></a>

## value property — values / 010230010201 / 6

Type: `"string"`. Computed.

Value. Configuration parameter for value

Upstream description:

Configuration parameter for value

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
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
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3001000200100131-1033222230320312-1033300033322110-2003312123320023-1222003011013212-3330331133303131-0110233100301321-2121200220331033"></a>

## Next pages — values / 010230010201 / 7

- [primary.rr_set_group.rr_set.caa_record](data-sources--dns_zone--reference--group-002.md#canonical-3021102303223021-0311330012113303-1211113233122120-0203302301302020-2223212013012002-2100133123121201-2130010021021103-0021000303323011)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-3323310232223321-3112000010212222-2211023330131221-2112123001002122-3232100332311223-1120221212020121-3101311000103133-3320301013333121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213211112232012-1013312031333000-1033331121311101-0300322120011120-2313033003002030-3131210010210221-1210200311332013-0333130101121213"></a>

## primary.rr_set_group.rr_set.cds_record — cds_record / 032031301200 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.cds_record

<a id="canonical-3112102222133222-3003220331112000-1320021313102202-0011300212321131-2113010310310211-3210200013001230-0010310233323110-3012303012123200"></a>

Type: `"single"`. Computed.

DNS CDS Record. DNS CDS Record.

Upstream description:

DNS CDS Record.

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

<a id="canonical-0320111333311013-2220030113110113-0213122212203013-0103330313021330-1121233311133200-0222032232233111-3223011312332031-3123121301232012"></a>

## Direct properties — cds_record / 032031301200 / 3

<a id="canonical-3302113021000030-3032201310101020-1212330002102131-3202003013323323-3123322320212003-0232103121011002-2203201100131123-1031000300221300"></a>

<a id="canonical-1210333213003230-3130001300031200-2000333031321220-2201323323111200-2233333321311013-0112221001012022-0023312212021102-1203110133111211"></a>

## name property — cds_record / 032031301200 / 4

Type: `"string"`. Computed.

CDS Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-2031103001322003-3132033100210321-3313002001022032-3112233112331222-0131013112021202-3032021211200112-1201220102032001-0121023223333123): complete subsection reference.

<a id="canonical-1111300322021121-1113321213200030-2232233221300332-0000323033130032-0100013310122302-3001122322010202-0113223300222223-2100010221301223"></a>

## Next pages — cds_record / 032031301200 / 5

- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2031103001322003-3132033100210321-3313002001022032-3112233112331222-0131013112021202-3032021211200112-1201220102032001-0121023223333123)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2031103001322003-3132033100210321-3313002001022032-3112233112331222-0131013112021202-3032021211200112-1201220102032001-0121023223333123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002232101001302-3220103222032113-3200103222130032-3201130103132311-2102031131020101-1330102330132301-2310120023030201-0212133303322033"></a>

## primary.rr_set_group.rr_set.cds_record.values — values / 123031101000 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-3323310232223321-3112000010212222-2211023330131221-2112123001002122-3232100332311223-1120221212020121-3101311000103133-3320301013333121)
- primary.rr_set_group.rr_set.cds_record.values

<a id="canonical-0323130310222111-2002023011012232-0210113302102211-1233200322033330-0113111301001120-2203020131211220-1322030002033312-3132100113001313"></a>

Type: `"list"`. Computed.

DS Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-1002321031032113-1023300030030312-1310131212101311-1101023302232133-3201132011321010-0221113022001323-3033333002111121-3111031202331033"></a>

## Direct properties — values / 123031101000 / 3

<a id="canonical-1311020133210033-0320312100301033-0200030313011222-2333111231113311-1311030322303311-1101313022202023-2111222020331103-2030022111321310"></a>

<a id="canonical-2233102110202320-3310310200001301-3131113303210112-1131323100133210-0002132012110333-0111013230011203-1222002222300100-0200131220132000"></a>

## ds_key_algorithm property — values / 123031101000 / 4

Type: `"string"`. Computed.

\[Enum:
UNSPECIFIED|RSASHA1|RSASHA1NSEC3SHA1|RSASHA256|RSASHA512|ECDSAP256SHA256|ECDSAP384SHA384|ED25519|ED448\]
DS key-value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1:
RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 -
ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448:
ED448. Possible values are \`UNSPECIFIED\`, \`RSASHA1\`, \`RSASHA1NSEC3SHA1\`, \`RSASHA256\`,
\`RSASHA512\`, \`ECDSAP256SHA256\`, \`ECDSAP384SHA384\`, \`ED25519\`, \`ED448\`.

Upstream description:

DS key-value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIED: UNSPECIFIED

&#8203;- RSASHA1: RSASHA1

&#8203;- RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1

&#8203;- RSASHA256: RSASHA256

&#8203;- RSASHA512: RSASHA512

&#8203;- ECDSAP256SHA256: ECDSAP256SHA256

&#8203;- ECDSAP384SHA384: ECDSAP384SHA384

&#8203;- ED25519: ED25519

&#8203;- ED448: ED448.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIED",
  "enum": [
    "UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3113103231121323-3231030031230133-1300200122020033-3330323030313330-3011112122111203-2003302222012021-2322112322131103-1301213223313032"></a>

<a id="canonical-1021313220022301-1031111023222100-0011333023111232-2010313020133121-2322022320212103-0133023301121301-1010310300303321-3313303200001322"></a>

## key_tag property — values / 123031101000 / 5

Type: `"number"`. Computed.

Short numeric value which can help quickly identify the referenced DNSKEY-record.

Upstream description:

A short numeric value which can help quickly identify the referenced DNSKEY-record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [sha1_digest](data-sources--dns_zone--reference--group-002.md#canonical-1221310323233000-2020211221120032-2020200012022302-0213201022002321-3022012323310230-3023132002312331-0222220232101030-0133303001130101): complete subsection reference.

- [sha256_digest](data-sources--dns_zone--reference--group-002.md#canonical-0300231111232000-3003303231213201-3231031133010120-2113321011000103-1113022200201300-2032213301321031-3203302330300313-3230131031133200): complete subsection reference.

- [sha384_digest](data-sources--dns_zone--reference--group-002.md#canonical-2303011021010130-3223320101310233-2332323322220231-1000113113222221-3321223221113223-3301003120213030-2303202202020301-3111211300210013): complete subsection reference.

<a id="canonical-3010003010112212-0232331312102331-3233022330203321-1020313020201012-1300202311023001-0121201223021002-0021302022321323-3302121020313203"></a>

## Next pages — values / 123031101000 / 6

- [primary.rr_set_group.rr_set.cds_record.values.sha1_digest](data-sources--dns_zone--reference--group-002.md#canonical-1221310323233000-2020211221120032-2020200012022302-0213201022002321-3022012323310230-3023132002312331-0222220232101030-0133303001130101)
- [primary.rr_set_group.rr_set.cds_record.values.sha256_digest](data-sources--dns_zone--reference--group-002.md#canonical-0300231111232000-3003303231213201-3231031133010120-2113321011000103-1113022200201300-2032213301321031-3203302330300313-3230131031133200)
- [primary.rr_set_group.rr_set.cds_record.values.sha384_digest](data-sources--dns_zone--reference--group-002.md#canonical-2303011021010130-3223320101310233-2332323322220231-1000113113222221-3321223221113223-3301003120213030-2303202202020301-3111211300210013)
- [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-3323310232223321-3112000010212222-2211023330131221-2112123001002122-3232100332311223-1120221212020121-3101311000103133-3320301013333121)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1221310323233000-2020211221120032-2020200012022302-0213201022002321-3022012323310230-3023132002312331-0222220232101030-0133303001130101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221103111002021-3201302303013300-0131200133012222-1132232121000220-3332131012223333-3200011103011110-1223210002012133-2203020123321201"></a>

## primary.rr_set_group.rr_set.cds_record.values.sha1_digest — sha1_digest / 110313002101 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-3323310232223321-3112000010212222-2211023330131221-2112123001002122-3232100332311223-1120221212020121-3101311000103133-3320301013333121)
- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2031103001322003-3132033100210321-3313002001022032-3112233112331222-0131013112021202-3032021211200112-1201220102032001-0121023223333123)
- primary.rr_set_group.rr_set.cds_record.values.sha1_digest

<a id="canonical-2321022202123020-0222121220103123-2222131300102200-0103002133131132-0210100012112032-1102002023303223-1012030202000002-1221002330103130"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 digest.

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

<a id="canonical-0132112033012103-1313210303320032-3311200111121220-3230311200023002-0013331030102221-0222021331122323-2313211010003000-3121110100102200"></a>

## Direct properties — sha1_digest / 110313002101 / 3

<a id="canonical-2133322000331020-0030100322011210-0300332323232303-0333110303131021-1210111223031103-0302010333221131-0213212032011220-3201232230330122"></a>

<a id="canonical-2101300211000022-1322012300221123-2122021002202323-3322023303130033-2302313003020010-2031202121022013-0330211331322000-3102123020301121"></a>

## digest property — sha1_digest / 110313002101 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 40
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-1230302231011101-2200101301333122-1213032111100122-2122113102203302-3301231312232113-1231231202323201-2312123223320123-3332033131300302"></a>

## Next pages — sha1_digest / 110313002101 / 5

- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2031103001322003-3132033100210321-3313002001022032-3112233112331222-0131013112021202-3032021211200112-1201220102032001-0121023223333123)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0300231111232000-3003303231213201-3231031133010120-2113321011000103-1113022200201300-2032213301321031-3203302330300313-3230131031133200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302012111223120-0313330031011320-3123302103221132-0310011111321312-0031333333102333-0010330101310111-2321032302200120-2021010032110201"></a>

## primary.rr_set_group.rr_set.cds_record.values.sha256_digest — sha256_digest / 133031232000 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-3323310232223321-3112000010212222-2211023330131221-2112123001002122-3232100332311223-1120221212020121-3101311000103133-3320301013333121)
- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2031103001322003-3132033100210321-3313002001022032-3112233112331222-0131013112021202-3032021211200112-1201220102032001-0121023223333123)
- primary.rr_set_group.rr_set.cds_record.values.sha256_digest

<a id="canonical-2000202202001212-1210220312331330-0110131133032212-1300311323302110-3101211203203101-2321001303100120-3312232210121021-0332220100000133"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 digest.

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

<a id="canonical-1123011121133310-0312231003112330-0101232222320102-0113020012100132-2300111231303021-2000220212302113-1002333013220023-3012030112203200"></a>

## Direct properties — sha256_digest / 133031232000 / 3

<a id="canonical-2202033002002021-1000111021220031-1331033210122233-3313031000000211-3000300233113102-2111230330123332-2230020201302231-0023002103223131"></a>

<a id="canonical-1133013333033321-2132323131003132-0332101133210230-1330132132002330-2031131301110112-2021102231123110-1013101213130130-1033023123212001"></a>

## digest property — sha256_digest / 133031232000 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
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
    "minLength": 64
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
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-0001121231022122-2013103330101123-2033210001010021-1113313332303321-0230202311233121-1011123023231000-2011002301320032-0011303033320233"></a>

## Next pages — sha256_digest / 133031232000 / 5

- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2031103001322003-3132033100210321-3313002001022032-3112233112331222-0131013112021202-3032021211200112-1201220102032001-0121023223333123)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2303011021010130-3223320101310233-2332323322220231-1000113113222221-3321223221113223-3301003120213030-2303202202020301-3111211300210013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220231313102320-1131012213300323-0233313130002103-3303113113213122-2312121331033222-2212233222133230-1310320231302003-3211200201133332"></a>

## primary.rr_set_group.rr_set.cds_record.values.sha384_digest — sha384_digest / 031312310300 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.cds_record](data-sources--dns_zone--reference--group-002.md#canonical-3323310232223321-3112000010212222-2211023330131221-2112123001002122-3232100332311223-1120221212020121-3101311000103133-3320301013333121)
- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2031103001322003-3132033100210321-3313002001022032-3112233112331222-0131013112021202-3032021211200112-1201220102032001-0121023223333123)
- primary.rr_set_group.rr_set.cds_record.values.sha384_digest

<a id="canonical-2012033320333210-1320312201333310-0331302222232311-1003331302102232-1032101231112102-1231021021230101-3333201132021111-0232113010211320"></a>

Type: `"single"`. Computed.

Configuration parameter for sha384 digest.

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

<a id="canonical-1212212313010333-0322110003111030-0032011003313030-2331121323232022-3221223331111303-2000030300310033-2020021022202301-2021103033001203"></a>

## Direct properties — sha384_digest / 031312310300 / 3

<a id="canonical-2211030303020013-0331200320121123-3212311320021110-1031231332012103-2131122222313110-2022313331101021-0200232030320220-1201101232112112"></a>

<a id="canonical-0300311011221122-3003100023233211-3210201211222001-0111132100203011-3021123112201323-2301332321211103-3000320300233221-3221033113330221"></a>

## digest property — sha384_digest / 031312310300 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-0033003113231021-1133201022130320-1131120121122022-1132312330223121-2311323130201100-2110021121101101-3301000310221131-3331310002310013"></a>

## Next pages — sha384_digest / 031312310300 / 5

- [primary.rr_set_group.rr_set.cds_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2031103001322003-3132033100210321-3313002001022032-3112233112331222-0131013112021202-3032021211200112-1201220102032001-0121023223333123)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2033130131100231-1310103201133222-0010203201322133-0113003033112230-3113331000213213-0323122330220111-1301022323030020-3123313211203310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002001213330011-3230320011023200-0330030303301311-0023220233213302-0312013101113023-0333313200112133-1101020132020000-3112330212012302"></a>

## primary.rr_set_group.rr_set.cert_record — cert_record / 200323201101 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.cert_record

<a id="canonical-1333323223212332-2013302100133331-1221312031023331-0321203302021123-3333310322103003-1213310001113303-0211312030311301-3023103012323033"></a>

Type: `"single"`. Computed.

Configuration parameter for cert record.

Upstream description:

DNS CERT Record.

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

<a id="canonical-3103331212322122-0330303320230110-0303300303213302-3001331311331303-3021221020321123-3211320100310100-0023103300011123-1033033222201211"></a>

## Direct properties — cert_record / 200323201101 / 3

<a id="canonical-0012022202310032-3221110201121102-1101331203011021-2020213332011022-1330102310002331-0000003100331121-2100120221011233-0303223210232310"></a>

<a id="canonical-1310001100333332-2303100000230002-3010231121120003-3001231221321232-3110223101232231-1103102212311010-1201133322033233-3110300230100020"></a>

## name property — cert_record / 200323201101 / 4

Type: `"string"`. Computed.

CERT Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-002.md#canonical-2032020132122123-2222213020312321-1332130201131112-3213000100131220-3033100331103133-3230013001320320-1302033020131202-2022220120133330): complete subsection reference.

<a id="canonical-3300313213012302-2310022221030002-3003301031211112-1021322210223030-0331112121010033-1131230131032222-1000132330130333-2122110110010310"></a>

## Next pages — cert_record / 200323201101 / 5

- [primary.rr_set_group.rr_set.cert_record.values](data-sources--dns_zone--reference--group-002.md#canonical-2032020132122123-2222213020312321-1332130201131112-3213000100131220-3033100331103133-3230013001320320-1302033020131202-2022220120133330)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2032020132122123-2222213020312321-1332130201131112-3213000100131220-3033100331103133-3230013001320320-1302033020131202-2022220120133330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313332212101300-0312332033130301-1013201203310220-2232023323110000-1022300331012333-3000132333322010-0103300320020221-0231030312110313"></a>

## primary.rr_set_group.rr_set.cert_record.values — values / 031023120223 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.cert_record](data-sources--dns_zone--reference--group-002.md#canonical-2033130131100231-1310103201133222-0010203201322133-0113003033112230-3113331000213213-0323122330220111-1301022323030020-3123313211203310)
- primary.rr_set_group.rr_set.cert_record.values

<a id="canonical-2130123321101122-3033233202331100-1111013012031301-1313123121032113-3023223120331121-2000332312300303-3000022210322301-2100102202121122"></a>

Type: `"list"`. Computed.

CERT Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2111230230221033-2133111300123213-0112222012211133-3030003122122300-0111232032232210-3302332021013111-2233313130200033-3001202032312012"></a>

## Direct properties — values / 031023120223 / 3

<a id="canonical-3333032312101212-2202131123321222-1020131320202200-3323220001231020-3333330012103133-2031012001301320-2003021211022132-3322000303031121"></a>

<a id="canonical-1102321022103300-0022231312200200-1110233203122203-3230232032310001-1002132031332120-2333020213222222-2333303200123320-0233201031332103"></a>

## algorithm property — values / 031023120223 / 4

Type: `"string"`. Computed.

\[Enum: RESERVEDALGORITHM|RSAMD5|DH|DSASHA1|ECC|RSASHA1ALGORITHM|INDIRECT|PRIVATEDNS|PRIVATEOID\]
CERT algorithm value must be compatible with the specified algorithm. - RESERVEDALGORITHM:
RESERVEDALGORITHM - RSAMD5: RSAMD5 - DH: DH - DSASHA1: DSASHA1 - ECC: ECC - RSASHA1ALGORITHM:
RSA-SHA1 - INDIRECT: INDIRECT - PRIVATEDNS: PRIVATEDNS - PRIVATEOID: PRIVATEOID. Possible values are
\`RESERVEDALGORITHM\`, \`RSAMD5\`, \`DH\`, \`DSASHA1\`, \`ECC\`, \`RSASHA1ALGORITHM\`, \`INDIRECT\`,
\`PRIVATEDNS\`, \`PRIVATEOID\`. Defaults to \`RESERVEDALGORITHM\`.

Upstream description:

CERT algorithm value must be compatible with the specified algorithm.

&#8203;- RESERVEDALGORITHM: RESERVEDALGORITHM

&#8203;- RSAMD5: RSAMD5

&#8203;- DH: DH

&#8203;- DSASHA1: DSASHA1

&#8203;- ECC: ECC

&#8203;- RSASHA1ALGORITHM: RSA-SHA1

&#8203;- INDIRECT: INDIRECT

&#8203;- PRIVATEDNS: PRIVATEDNS

&#8203;- PRIVATEOID: PRIVATEOID.

Receipt-pinned upstream constraints:

```json
{
  "default": "RESERVEDALGORITHM",
  "enum": [
    "RESERVEDALGORITHM",
    "RSAMD5",
    "DH",
    "DSASHA1",
    "ECC",
    "RSASHA1ALGORITHM",
    "INDIRECT",
    "PRIVATEDNS",
    "PRIVATEOID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3100313120323021-3110303122033131-1131312211110202-1133031331222312-3233300031121000-0113332033231221-1003232210002310-2001111102013021"></a>

<a id="canonical-0310232312111103-1222201022212312-1131113211102330-0112133331211232-1002010132201300-0321213221021323-0110202301030111-1021220113111212"></a>

## cert_key_tag property — values / 031023120223 / 5

Type: `"number"`. Computed.

Key Tag. Tag for categorization and filtering

Upstream description:

Tag for categorization and filtering

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1001133220032031-1213300121231022-1001202300323021-1232011122201033-2233112130131312-3300203202310010-2121310221121232-2033001023012003"></a>

<a id="canonical-2020102332202313-3021301202013021-3120211232230312-0220232033022220-3233223111210103-3102121100332020-1310100102323110-2002202313112011"></a>

## cert_type property — values / 031023120223 / 6

Type: `"string"`. Computed.

\[Enum: INVALIDCERTTYPE|PKIX|SPKI|PGP|IPKIX|ISPKI|IPGP|ACPKIX|IACPKIX|URI\_|OID\] CERT type value
must be compatible with the specified types. - INVALIDCERTTYPE: INVALIDCERTTYPE - PKIX: PKIX - SPKI:
SPKI - PGP: PGP - IPKIX: IPKIX - ISPKI: ISPKI - IPGP: IPGP - ACPKIX: ACPKIX - IACPKIX: IACPKIX -
URI\_: URI - OID: OID. Possible values are \`INVALIDCERTTYPE\`, \`PKIX\`, \`SPKI\`, \`PGP\`,
\`IPKIX\`, \`ISPKI\`, \`IPGP\`, \`ACPKIX\`, \`IACPKIX\`, \`URI\_\`, \`OID\`. Defaults to
\`INVALIDCERTTYPE\`.

Upstream description:

CERT type value must be compatible with the specified types.

&#8203;- INVALIDCERTTYPE: INVALIDCERTTYPE

&#8203;- PKIX: PKIX

&#8203;- SPKI: SPKI

&#8203;- PGP: PGP

&#8203;- IPKIX: IPKIX

&#8203;- ISPKI: ISPKI

&#8203;- IPGP: IPGP

&#8203;- ACPKIX: ACPKIX

&#8203;- IACPKIX: IACPKIX

&#8203;- URI\_: URI

&#8203;- OID: OID.

Receipt-pinned upstream constraints:

```json
{
  "default": "INVALIDCERTTYPE",
  "enum": [
    "INVALIDCERTTYPE",
    "PKIX",
    "SPKI",
    "PGP",
    "IPKIX",
    "ISPKI",
    "IPGP",
    "ACPKIX",
    "IACPKIX",
    "URI_",
    "OID"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1022303100312121-2133320033312012-3220332213330112-0202030130002133-3112131122023121-2101032110003213-2222023222321232-3003020310021133"></a>

<a id="canonical-3131333220103201-0132113031130012-3020112203131223-2200110303332103-2332103212020023-0303333233120123-1310122132002131-0203122300213231"></a>

## certificate property — values / 031023120223 / 7

Type: `"string"`. Computed.

Certificate. Certificate in base 64 format.

Upstream description:

Certificate in base 64 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 5242880,
      "min": 100
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "pem",
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
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
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1323311011202101-2202200331310310-2122222323020000-1312111232333103-2300201133121313-0310330211203030-3310002111122322-0131021222312000"></a>

## Next pages — values / 031023120223 / 8

- [primary.rr_set_group.rr_set.cert_record](data-sources--dns_zone--reference--group-002.md#canonical-2033130131100231-1310103201133222-0010203201322133-0113003033112230-3113331000213213-0323122330220111-1301022323030020-3123313211203310)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-3000110033203111-2312032311212322-2103322303022131-1123102210000211-0312322123020323-2222132212012111-1000102011220113-3123113121110030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
