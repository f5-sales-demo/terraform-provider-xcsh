---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-1302030201331213-3123230333111232-1202311020110210-0012302221211200-1211001310113003-2003010301033112-2200013112222232-2110223031011332"></a>

## aws_parameters.new_vpc — new_vpc / 122111130223 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- aws_parameters.new_vpc

<a id="canonical-0112231112233111-0033010130321023-2233122310001130-1033203001001231-3331233003331200-0002231332322330-1232333100323122-0230100232233102"></a>

Type: `"single"`. Computed.

AWS VPC Parameters. Parameters to create new AWS VPC.

Upstream description:

Parameters to create new AWS VPC.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-name_choice": "[\"autogenerate\",\"name_tag\"]"
}
```

<a id="canonical-1101213122002103-1301212202002302-2100320211223311-1123022122221230-2123002213111120-0103122110213212-3130222232032330-2301213111220100"></a>

## Direct properties — new_vpc / 122111130223 / 3

- [autogenerate](data-sources--aws_tgw_site--reference--group-002.md#canonical-2301031003212131-1000113311002133-0230303020010213-0221323133302001-2211102101012203-2121302331020030-0312332020221021-3110120213221123): complete subsection reference.

<a id="canonical-2011222113213212-2330310332201313-2113230133003330-2213200021332221-2203111100211203-2302031012110332-1130123002320003-1121133233313131"></a>

<a id="canonical-0211030112230023-2130012021330101-1110131301233310-1210120313301200-2332331010110211-0323012311033013-1111302231233023-2331303113320322"></a>

## name_tag property — new_vpc / 122111130223 / 4

Type: `"string"`. Computed.

Exclusive with \[autogenerate\] Specify the VPC Name.

Upstream description:

Exclusive with \[autogenerate\] Specify the VPC Name.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0001210313003100-0232001313303120-2333231110002223-0121203201203122-1310113210221201-2313231011332102-0123003103203113-2003200302022221"></a>

<a id="canonical-1133303333331022-3103100212300220-1323331303333103-2031330232222132-2112313000310222-0103233130332322-3311100122332120-2111223301000023"></a>

## primary_ipv4 property — new_vpc / 122111130223 / 5

Type: `"string"`. Computed.

IPv4 CIDR block for this VPC. It has to be private address space. The Primary IPv4 block cannot be
modified. All subnets prefixes in this VPC must be part of this CIDR block.

Upstream description:

IPv4 CIDR block for this VPC. It has to be private address space. The Primary IPv4 block cannot be
modified. All subnets prefixes in this VPC must be part of this CIDR block.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "16"
  }
}
```

<a id="canonical-2303320310012213-1120213100001113-1112221302103223-2031333000211301-2111333220030301-2023303222111332-1230322111000112-2130003202210133"></a>

## Next pages — new_vpc / 122111130223 / 6

- [aws_parameters.new_vpc.autogenerate](data-sources--aws_tgw_site--reference--group-002.md#canonical-2301031003212131-1000113311002133-0230303020010213-0221323133302001-2211102101012203-2121302331020030-0312332020221021-3110120213221123)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2301031003212131-1000113311002133-0230303020010213-0221323133302001-2211102101012203-2121302331020030-0312332020221021-3110120213221123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200220220132323-0332232310313331-3232030002322121-3030332032100102-3300301003303230-0031211033131121-2023313003101211-2332313113003113"></a>

## aws_parameters.new_vpc.autogenerate — autogenerate / 021011013303 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [aws_parameters.new_vpc](data-sources--aws_tgw_site--reference--group-001.md#canonical-2203311030300201-1320113131100101-3012322103333133-3003203313213332-2310302023313002-3013210130001312-2033332213210030-0030330322300233)
- aws_parameters.new_vpc.autogenerate

<a id="canonical-2332113211022230-3021032320121113-2023111302300120-0101212303023001-1002012100210020-3132130221131223-1110001201232212-0222232323001203"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for autogenerate.

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

<a id="canonical-3203223133220333-3313231011333100-1332002130120003-3332300130211330-2133202122212333-0310200022031002-0102313022021003-3203001013122222"></a>

## Direct properties — autogenerate / 021011013303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0220302101332303-2203100102232302-2123013123331200-0301233210320032-0303320213230313-2131120320023022-1332033203332122-1321301320013233"></a>

## Next pages — autogenerate / 021011013303 / 4

- [aws_parameters.new_vpc](data-sources--aws_tgw_site--reference--group-001.md#canonical-2203311030300201-1320113131100101-3012322103333133-3003203313213332-2310302023313002-3013210130001312-2033332213210030-0030330322300233)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0231100231122132-2021112330123120-0002300113022030-2301010222102233-1213302232110033-3022113202020003-0322300221102323-2233201022000200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000103330200013-0132013312031311-3120220113011321-0313312303120030-1022313323010120-0111331013300300-2112313013202333-2030211003130321"></a>

## aws_parameters.no_worker_nodes — no_worker_nodes / 023312231033 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- aws_parameters.no_worker_nodes

<a id="canonical-3203003320212301-1211312102132333-0020320012103133-3132112000331121-2132222112010220-2001033122103123-0010230113101110-1001231321221122"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no worker nodes.

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

<a id="canonical-2203310332200110-0132212312002333-2003111002112001-0311003003012320-2133102213022013-1120033002031012-0323313203310310-2110112320303311"></a>

## Direct properties — no_worker_nodes / 023312231033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003001310123120-3301103120133132-0123000103210123-2220231203133120-1001110010111020-0113000203130221-3022111302322131-2103301121011132"></a>

## Next pages — no_worker_nodes / 023312231033 / 4

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3033320011130322-2011303033033001-0310202021301030-3231002233002112-2303103000310110-0332011032300123-0320013032020212-2331110113213020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301012002031300-1312010120101021-2320222132212302-0212100021020112-2131213322310321-0032300032133132-1212003311312312-3132031110322113"></a>

## aws_parameters.reserved_tgw_cidr — reserved_tgw_cidr / 000222213013 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- aws_parameters.reserved_tgw_cidr

<a id="canonical-3003230120221223-2022101210131100-0222321322002132-3202120030331132-2200303120212312-2322211020220110-1203211221020302-1130323112221111"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for reserved tgw cidr.

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

<a id="canonical-3021002211322103-1021102100001323-2320301020323133-2200001301230330-2212123130201311-1321123302201320-3301000322113120-3313020210230120"></a>

## Direct properties — reserved_tgw_cidr / 000222213013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302322322003221-1231313121023212-1011101121312303-2200230101101112-0022333323101313-0330313101311303-1312113210230230-0233020121332112"></a>

## Next pages — reserved_tgw_cidr / 000222213013 / 4

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3202021121111032-1032113112203313-1011031301321120-1031130323101023-1123113100211230-0030303011131111-1220032320323021-3021110302323002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220301210310133-1032133102212321-3210000331332223-2201230010011020-1112000103133310-3112231100231221-2330202322101020-2222232011020101"></a>

## aws_parameters.tgw_cidr — tgw_cidr / 201020320133 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- aws_parameters.tgw_cidr

<a id="canonical-1313132033010001-2200013112203110-3211021301020102-3210132213112210-0123111101232111-1020003223020131-3012021031301030-0203232121220221"></a>

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

<a id="canonical-3113132101033032-2011212302202103-1002302201013223-1121320231311200-1031113312131120-3113233021021101-0202321010012221-3320130233013133"></a>

## Direct properties — tgw_cidr / 201020320133 / 3

<a id="canonical-2233130003332021-2110222203320321-2122001222220000-2031202322030213-0133011133030203-1331003313303003-2103330301130230-3122011001213231"></a>

<a id="canonical-2312103323003031-3323303033220300-0321322020212221-0010021131100233-0123332331230002-1120122102220300-0020210033222102-2210332132112312"></a>

## IPv4 property — tgw_cidr / 201020320133 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3203122033120323-1020111221213232-2303122011213210-3030221322302022-1213302112011132-0112333202033311-1003033313111033-0123011032330203"></a>

## Next pages — tgw_cidr / 201020320133 / 5

- [aws_parameters](data-sources--aws_tgw_site--reference--group-001.md#canonical-3311120101031201-0112330033331111-0330012133130331-2122223112223133-1300102323001230-2222231003110310-3112001230233232-2000103311201302)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3121133133112002-0033332221002003-2013023112031223-1233010230303323-2212120132203032-2310303231211001-0211233302000202-0000211330233312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210002302112013-1000130203033120-2132102033332030-3123231012223032-3133311301322111-2123230220101012-2210301130101211-0022200322313000"></a>

## block_all_services — block_all_services / 330013232132 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- block_all_services

<a id="canonical-0313121231222023-3332011302231030-3001320123301303-3330232232020332-3011112101131021-2020100020112103-1300121100013133-2121033212021222"></a>

Type: `["object", {}]`. Computed.

\[OneOf: block\_all\_services, blocked\_services, default\_blocked\_services; Default:
default\_blocked\_services\] Enable this option

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

- [block_all_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-0313121231222023-3332011302231030-3001320123301303-3330232232020332-3011112101131021-2020100020112103-1300121100013133-2121033212021222)
- [blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-2312023131001001-1001333003132333-2021131112102222-1203313220103300-1203132211301111-1131012000233030-1212011211200331-2002132113103320)
- [default_blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-2312030110333232-3310122233221321-3022221200223202-3223330011103002-2222130122132232-1003220323112212-3013201033332133-3330032030023200)

Select alternatives according to the provider validators above.

<a id="canonical-2302233133313310-0020200130332211-2122231230313022-3300202011013220-1001233000102330-0210332300133003-0201222301330100-1320022233321231"></a>

## Direct properties — block_all_services / 330013232132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231020330032301-1113223101213332-3321112123210220-2333131112023231-2320002331002223-2323320221202232-1333003111332321-1232013132022200"></a>

## Next pages — block_all_services / 330013232132 / 4

- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2100121132300313-3112323000231331-3322303011221212-3103232201110033-2023130132111022-0020113011113210-0001230331021232-0022021332322021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110333121003111-1022333202032122-3102030002113213-1232012211300013-2102011331110211-3002331113313232-3303102031321012-1202010212111022"></a>

## blocked_services — blocked_services / 022330210103 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- blocked_services

<a id="canonical-2312023131001001-1001333003132333-2021131112102222-1203313220103300-1203132211301111-1131012000233030-1212011211200331-2002132113103320"></a>

Type: `"single"`. Computed.

Disable node local services on this site.

Upstream description:

Disable node local services on this site. Note: The chosen services will GET disabled on all nodes
in the site.

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

<a id="canonical-1222312211312031-3001302313012033-3220020302033133-1202021032321223-3102123132122323-1300022021112321-3110021110102201-1010323131221031"></a>

## Direct properties — blocked_services / 022330210103 / 3

- [blocked_service](data-sources--aws_tgw_site--reference--group-002.md#canonical-1231201011232002-3210011200010222-3120201013221220-0123120103201220-3320013321212032-2103320310130211-3102103133333320-2010332301011123): complete subsection reference.

<a id="canonical-1101231000222300-2131133331030013-0032232021220321-1010212232011200-2101023301002323-3312112030011032-0321130121121132-1122001121301331"></a>

## Next pages — blocked_services / 022330210103 / 4

- [blocked_services.blocked_service](data-sources--aws_tgw_site--reference--group-002.md#canonical-1231201011232002-3210011200010222-3120201013221220-0123120103201220-3320013321212032-2103320310130211-3102103133333320-2010332301011123)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1231201011232002-3210011200010222-3120201013221220-0123120103201220-3320013321212032-2103320310130211-3102103133333320-2010332301011123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031302301000210-0231322210332213-2033001310032102-0002313221231223-1121222111201110-3331011311030203-0302130022011113-3103211203312333"></a>

## blocked_services.blocked_service — blocked_service / 323003321022 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-2100121132300313-3112323000231331-3322303011221212-3103232201110033-2023130132111022-0020113011113210-0001230331021232-0022021332322021)
- blocked_services.blocked_service

<a id="canonical-1300022131331211-2220302320131023-3032201012222211-2301330200201311-2010123233323313-1003023132003032-0322313122012100-2320001203332310"></a>

Type: `"list"`. Computed.

Disable Node Local Services. Blocking or denial configuration

Upstream description:

Blocking or denial configuration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1031101113021000-1202311322322021-2011020313311220-1001203102110230-1332223211223033-3303012100311132-0023122310103330-3023312223013023"></a>

## Direct properties — blocked_service / 323003321022 / 3

- [DNS](data-sources--aws_tgw_site--reference--group-002.md#canonical-1210112213223033-0300331211300201-2031100012111022-0110013220123230-3011301303013022-0110221103133033-0113012332211203-1311302323022102): complete subsection reference.

<a id="canonical-1103303203131011-1212122300101020-3102232033110001-2003121321123013-0102121302020101-2132310110001121-1201212132323213-0000223113022322"></a>

<a id="canonical-0203121221323332-0113120130312330-2031020322020203-3100120100022011-2022022103021210-3213311310033102-0332301032211112-0233312020321021"></a>

## network_type property — blocked_service / 323003321022 / 4

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [SSH](data-sources--aws_tgw_site--reference--group-002.md#canonical-0301231123103212-0012222120030111-2131221031023130-2230022312202221-2300301210010032-0321120132020110-3220003111300323-3012201012023301): complete subsection reference.

- [web_user_interface](data-sources--aws_tgw_site--reference--group-002.md#canonical-1103100122133100-1020311012001200-1133112120310023-1213132132110103-3011022110131203-0002113103223130-0323102230130213-1332133331212121): complete subsection reference.

<a id="canonical-3130332311333130-3133131001103131-0132200323300121-0221202003122012-0113211022131203-2230233022031322-3030302201223130-1230300331030310"></a>

## Next pages — blocked_service / 323003321022 / 5

- [blocked_services.blocked_service.dns](data-sources--aws_tgw_site--reference--group-002.md#canonical-1210112213223033-0300331211300201-2031100012111022-0110013220123230-3011301303013022-0110221103133033-0113012332211203-1311302323022102)
- [blocked_services.blocked_service.ssh](data-sources--aws_tgw_site--reference--group-002.md#canonical-0301231123103212-0012222120030111-2131221031023130-2230022312202221-2300301210010032-0321120132020110-3220003111300323-3012201012023301)
- [blocked_services.blocked_service.web_user_interface](data-sources--aws_tgw_site--reference--group-002.md#canonical-1103100122133100-1020311012001200-1133112120310023-1213132132110103-3011022110131203-0002113103223130-0323102230130213-1332133331212121)
- [blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-2100121132300313-3112323000231331-3322303011221212-3103232201110033-2023130132111022-0020113011113210-0001230331021232-0022021332322021)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1210112213223033-0300331211300201-2031100012111022-0110013220123230-3011301303013022-0110221103133033-0113012332211203-1311302323022102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203001002211111-2231323103121331-1021303200121302-1100300123213031-0033222013333300-1021202331113113-0132323211232013-0102132021013332"></a>

## blocked_services.blocked_service.DNS — DNS / 212103331332 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-2100121132300313-3112323000231331-3322303011221212-3103232201110033-2023130132111022-0020113011113210-0001230331021232-0022021332322021)
- [blocked_services.blocked_service](data-sources--aws_tgw_site--reference--group-002.md#canonical-1231201011232002-3210011200010222-3120201013221220-0123120103201220-3320013321212032-2103320310130211-3102103133333320-2010332301011123)
- blocked_services.blocked_service.DNS

<a id="canonical-3323323203231012-0003030211312012-2031323020211133-0230333023000100-0123001312030211-3232103121203222-1220010121100133-0333120000213312"></a>

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

<a id="canonical-2112120222021021-2031112000310223-1220302233321022-1210103310003010-1100030211200202-1323320203311022-3210320313213113-2101130212132033"></a>

## Direct properties — DNS / 212103331332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130031002321023-0200233112210202-0132222302001311-0000312213113220-2022020332023031-0310021331112111-3123121212312011-2212100100010301"></a>

## Next pages — DNS / 212103331332 / 4

- [blocked_services.blocked_service](data-sources--aws_tgw_site--reference--group-002.md#canonical-1231201011232002-3210011200010222-3120201013221220-0123120103201220-3320013321212032-2103320310130211-3102103133333320-2010332301011123)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0301231123103212-0012222120030111-2131221031023130-2230022312202221-2300301210010032-0321120132020110-3220003111300323-3012201012023301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023220311201321-0031300021111233-1333210333200333-3023033111101032-2203013321013023-3012021023023020-2300030112132012-3110210222131211"></a>

## blocked_services.blocked_service.SSH — SSH / 013010031021 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-2100121132300313-3112323000231331-3322303011221212-3103232201110033-2023130132111022-0020113011113210-0001230331021232-0022021332322021)
- [blocked_services.blocked_service](data-sources--aws_tgw_site--reference--group-002.md#canonical-1231201011232002-3210011200010222-3120201013221220-0123120103201220-3320013321212032-2103320310130211-3102103133333320-2010332301011123)
- blocked_services.blocked_service.SSH

<a id="canonical-1120221221301121-2001123210002011-2223032131113030-2121310010133212-2022330233012133-2220210231313120-0112113102333322-0233222222000321"></a>

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

<a id="canonical-2322212212110133-1011123201000000-2132013130121121-3333030212311132-0101220231121112-3011121113213322-0222100200300223-0130221220210013"></a>

## Direct properties — SSH / 013010031021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000223311212312-3320101133133200-1303210033212121-0211332203312323-0320032103211223-1010313323031332-1223213303213300-0223333332202012"></a>

## Next pages — SSH / 013010031021 / 4

- [blocked_services.blocked_service](data-sources--aws_tgw_site--reference--group-002.md#canonical-1231201011232002-3210011200010222-3120201013221220-0123120103201220-3320013321212032-2103320310130211-3102103133333320-2010332301011123)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1103100122133100-1020311012001200-1133112120310023-1213132132110103-3011022110131203-0002113103223130-0323102230130213-1332133331212121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023103133221000-2131132113232112-3233012030023110-1101023131110012-0033101112102331-3012202210332202-3002031322002021-1032120223123031"></a>

## blocked_services.blocked_service.web_user_interface — web_user_interface / 111202330003 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [blocked_services](data-sources--aws_tgw_site--reference--group-002.md#canonical-2100121132300313-3112323000231331-3322303011221212-3103232201110033-2023130132111022-0020113011113210-0001230331021232-0022021332322021)
- [blocked_services.blocked_service](data-sources--aws_tgw_site--reference--group-002.md#canonical-1231201011232002-3210011200010222-3120201013221220-0123120103201220-3320013321212032-2103320310130211-3102103133333320-2010332301011123)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-0120202011113322-0211010203132022-3031213112210133-2002303022203020-1212000212013103-2023221032302232-0311111013301110-1300320111131221"></a>

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

<a id="canonical-3313003322120101-3131330122213112-1302030112210133-1112011110030020-1032112210020031-3031103031113201-0103013011331113-3132003112013003"></a>

## Direct properties — web_user_interface / 111202330003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302231103333202-3300002132302303-3112322213123022-1210021000023132-1120212012311212-2220011312023013-0210300323221331-0003211130332331"></a>

## Next pages — web_user_interface / 111202330003 / 4

- [blocked_services.blocked_service](data-sources--aws_tgw_site--reference--group-002.md#canonical-1231201011232002-3210011200010222-3120201013221220-0123120103201220-3320013321212032-2103320310130211-3102103133333320-2010332301011123)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3320203122001123-1321021122332021-0300001203322000-1032311002033301-2121100322032201-1222120121102201-2210100212312030-1023311021213310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212330311012201-0211120202123321-1232011001013032-3301303301300322-3302331312120323-0103022101232120-2230031121023323-1111011313113022"></a>

## coordinates — coordinates / 001203220312 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- coordinates

<a id="canonical-0101123002002132-3221330333332100-1210130221302010-1300220122000203-0300330031113102-1201320123323231-0023212000321013-3220010023312202"></a>

Type: `"single"`. Computed.

Coordinates of the site which provides the site physical location.

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

<a id="canonical-1322101103033133-3003322122311301-1312110022011321-0312112111120031-3101202131300232-1323323120230300-0013000221201033-1123002030200012"></a>

## Direct properties — coordinates / 001203220312 / 3

<a id="canonical-0130030110301303-2110302010301110-1233130020131323-2300233320133222-0220203023302002-0300313113211002-1321010031120122-0121100022013002"></a>

<a id="canonical-3001032122301132-2301123100333010-2203002132003022-1111202300232322-3003132121110313-2220123020232223-0000001000230011-2003320033203121"></a>

## latitude property — coordinates / 001203220312 / 4

Type: `"number"`. Computed.

Latitude. Latitude of the site location.

Upstream description:

Latitude of the site location.

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
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  }
}
```

<a id="canonical-2120301010302112-2011211300033332-0200233123121233-0200331303011230-1033312002012122-3332202023130103-3032203312331020-1211220011013111"></a>

<a id="canonical-3003312302321212-1313103221011331-2001311100202221-2033221301133132-0332121013231313-0220131223003133-0113320113031212-3233300312111011"></a>

## longitude property — coordinates / 001203220312 / 5

Type: `"number"`. Computed.

Longitude. Longitude of site location.

Upstream description:

Longitude of site location.

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
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  }
}
```

<a id="canonical-2330033200320220-1220310013331203-3302310201311023-3321320003000003-1222322113121122-1210321012313313-1113223033333132-3301232130210112"></a>

## Next pages — coordinates / 001203220312 / 6

- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1011223223020320-0302302203312331-1311132331323301-3220231231123022-3310000100111101-2322211321123200-3220013010110222-2213033012101301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221101322320233-0111001203210100-2132001312332002-3230213302321322-2023313223030111-1113202021213303-3032200000320100-1111010020121130"></a>

## custom_dns — custom_dns / 322110203321 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- custom_dns

<a id="canonical-2221302212322221-1223320221022110-0220332200201012-3312332033230330-3033230112212123-1213011303033320-3203211301032013-0231010203030203"></a>

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

<a id="canonical-3013311231212000-1102301223011023-0322001123220203-0113221023230233-2123222231303332-1301102310031110-1122131222213233-2111122111030031"></a>

## Direct properties — custom_dns / 322110203321 / 3

<a id="canonical-3300133222021223-3221311321202313-0303231233112200-0222200022022301-2132120102232103-1100331202003113-2133011212230202-0013210131012302"></a>

<a id="canonical-3223232011113002-3313003011301310-3212022133203200-3112101333013323-2112202033211020-1031133003301330-0102012320333221-0112032222122312"></a>

## inside_nameserver property — custom_dns / 322110203321 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0302133300333302-3103011221213310-2323220131111231-2220332221001122-1020210303023202-1301031211012131-3211323030132013-3102001233110132"></a>

<a id="canonical-0203301221222332-0211211003223330-0312203123133013-2220020023120101-3022033131202031-1211200302321232-0001321021010101-3311132032103032"></a>

## outside_nameserver property — custom_dns / 322110203321 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2023212300100000-3313100320321332-0312200022132323-1012312000133121-0210200102031020-3211120201301012-1211201332230203-2113313330322310"></a>

## Next pages — custom_dns / 322110203321 / 6

- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1212131321112002-1213201223220020-0101122010302212-1111013310121030-0020210013331103-1030212031001122-2310211012032310-0030321010032123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321000331212311-1301032213033221-0023202223231033-0112210231001131-1322000010223332-1011003102013323-0213010322102003-2221203300012021"></a>

## default_blocked_services — default_blocked_services / 121023200200 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- default_blocked_services

<a id="canonical-2312030110333232-3310122233221321-3022221200223202-3223330011103002-2222130122132232-1003220323112212-3013201033332133-3330032030023200"></a>

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

<a id="canonical-2132123211003031-2320022322023220-3001111123300310-3322202313003102-3113122111321132-2332103132121111-3322011022203300-2113233201310320"></a>

## Direct properties — default_blocked_services / 121023200200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330223013020322-0130223103012330-1003222323111333-1032101202331200-0300303011033132-3001230301230323-3322121322200121-2232223321013223"></a>

## Next pages — default_blocked_services / 121023200200 / 4

- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0122003113200213-2133101112002122-3103022230022203-0333103101220331-3021302000323233-3110133023222210-2033301133212010-3100030200223020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303310311113332-3112022333123002-1232121023313321-2210312303323003-3131132323311030-2122321320232012-1223202212011222-1330103111010223"></a>

## direct_connect_disabled — direct_connect_disabled / 200031123300 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- direct_connect_disabled

<a id="canonical-0213011121022131-0312130230032300-2210110201103011-3233202121120321-3021130001221322-2220003033230012-1230003130100132-2023103022312300"></a>

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

- [direct_connect_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0213011121022131-0312130230032300-2210110201103011-3233202121120321-3021130001221322-2220003033230012-1230003130100132-2023103022312300)
- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-3111131133223020-0022120001021123-3320122232120203-2200102202121330-2100211321231332-3020012303233220-2123122101010000-0003301310111133)
- [private_connectivity](data-sources--aws_tgw_site--reference--group-002.md#canonical-2330311303310000-3201101003320222-1332021212103110-0001203211113110-3100222000231102-3210211331232231-0231310132333032-2201013101212203)

Select alternatives according to the provider validators above.

<a id="canonical-3313032231001031-3111010003123232-3321202202031303-0313322222211301-0220223221233022-2221013120223100-3203310120332101-1213311303332223"></a>

## Direct properties — direct_connect_disabled / 200031123300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023003003012232-2301012121103033-2321121321300130-0331031210101110-2303231330230320-3322321323032000-2331030211110121-3311213030303301"></a>

## Next pages — direct_connect_disabled / 200031123300 / 4

- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033000110031021-3300122313301322-1122110230232031-3103200233321110-0212321312322033-3310121032332302-2311300301232310-1221302003200201"></a>

## direct_connect_enabled — direct_connect_enabled / 332023202212 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- direct_connect_enabled

<a id="canonical-3111131133223020-0022120001021123-3320122232120203-2200102202121330-2100211321231332-3020012303233220-2123122101010000-0003301310111133"></a>

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

<a id="canonical-0200101311210123-0102303321330020-0210030221112332-3202132011300101-1130331231322302-0331012222031303-2003133203100301-1320303113011202"></a>

## Direct properties — direct_connect_enabled / 332023202212 / 3

- [auto_asn](data-sources--aws_tgw_site--reference--group-002.md#canonical-1103323103013220-0232203201323132-2102220120001232-2223303122120212-1030131100102200-1223201330111103-2032313102022010-2112112103121102): complete subsection reference.

<a id="canonical-1320021202311333-0210201310233010-2132232220231003-3333131201223321-2121321202110002-1213032030120102-3031231122210031-2030113013020010"></a>

<a id="canonical-3110100332010001-3133112122331230-2103010011012320-3113011312020302-2300220112020230-0312211232303210-1032300112331222-0103121323311213"></a>

## custom_asn property — direct_connect_enabled / 332023202212 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [hosted_vifs](data-sources--aws_tgw_site--reference--group-002.md#canonical-3122102210321312-3222221303132221-0221221122132303-0120123033113313-3311232031212022-0120001313113130-0321132113330030-2001102330231013): complete subsection reference.

- [standard_vifs](data-sources--aws_tgw_site--reference--group-002.md#canonical-3223331130303203-1032102332230311-1312022330000323-1322110023132332-0130230133222002-2002122012033202-2312021211203030-2002121030100230): complete subsection reference.

<a id="canonical-1230203233302220-0103123113200001-1121301221221330-3232001032112101-1011011222030310-2122210312223233-2120200201123211-0030233110020023"></a>

## Next pages — direct_connect_enabled / 332023202212 / 5

- [direct_connect_enabled.auto_asn](data-sources--aws_tgw_site--reference--group-002.md#canonical-1103323103013220-0232203201323132-2102220120001232-2223303122120212-1030131100102200-1223201330111103-2032313102022010-2112112103121102)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_tgw_site--reference--group-002.md#canonical-3122102210321312-3222221303132221-0221221122132303-0120123033113313-3311232031212022-0120001313113130-0321132113330030-2001102330231013)
- [direct_connect_enabled.standard_vifs](data-sources--aws_tgw_site--reference--group-002.md#canonical-3223331130303203-1032102332230311-1312022330000323-1322110023132332-0130230133222002-2002122012033202-2312021211203030-2002121030100230)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1103323103013220-0232203201323132-2102220120001232-2223303122120212-1030131100102200-1223201330111103-2032313102022010-2112112103121102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022013221001223-1300230231322320-1300012121213011-1310320023012123-1111031133320022-3003012012200010-1312300322211133-0121203131120232"></a>

## direct_connect_enabled.auto_asn — auto_asn / 323203020213 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300)
- direct_connect_enabled.auto_asn

<a id="canonical-3023031033312110-1223100000313302-3230110321012023-0221313032300101-0212313301323023-0000212300303101-1202301222323320-3010201300021231"></a>

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

<a id="canonical-0211222030030133-2312321033331330-3103130123123201-0012131211333222-0313012203033133-3311313212131013-1013120332020230-0023113322323213"></a>

## Direct properties — auto_asn / 323203020213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210303330302200-0322113221211302-3122033012331200-2313112032101002-3200110102233212-2123123332303310-2013113101000101-0211012222212313"></a>

## Next pages — auto_asn / 323203020213 / 4

- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3122102210321312-3222221303132221-0221221122132303-0120123033113313-3311232031212022-0120001313113130-0321132113330030-2001102330231013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113001132312211-3122301001303130-1222132110333320-1221303131312223-1021230102322022-3030013233020030-0001111100210111-1221131201000203"></a>

## direct_connect_enabled.hosted_vifs — hosted_vifs / 223302020110 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300)
- direct_connect_enabled.hosted_vifs

<a id="canonical-3132033233301130-3322100201222101-3201320113032030-0132112022220232-3303113200230030-3133300123322201-2031212231330133-3013032031022032"></a>

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

<a id="canonical-2010201201210213-0322011320002220-2333310303331111-3100211233203132-1011122320223113-1031103330030120-2313031132012222-2100031330311113"></a>

## Direct properties — hosted_vifs / 223302020110 / 3

- [site_registration_over_direct_connect](data-sources--aws_tgw_site--reference--group-002.md#canonical-1323200312211101-3000213103000332-3033030220220301-3321110123121210-3021212000212000-1202001200122110-1121122303001302-3302233022031333): complete subsection reference.

- [site_registration_over_internet](data-sources--aws_tgw_site--reference--group-002.md#canonical-0120213130102001-3131030112310103-0301302322131302-1322310201013033-1031030121332323-2001310120030023-0201320303102323-0033212103321321): complete subsection reference.

- [vif_list](data-sources--aws_tgw_site--reference--group-002.md#canonical-1120303330300310-1302023120223012-0010311013220332-1012220130010311-1123313031102023-3232232001032113-0033323222012230-1233220122203211): complete subsection reference.

<a id="canonical-2231323002130123-1103321330131132-0230201213223222-0131112221032210-2122023122331001-1201233321103311-0131211213320203-3321031001220232"></a>

## Next pages — hosted_vifs / 223302020110 / 4

- [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](data-sources--aws_tgw_site--reference--group-002.md#canonical-1323200312211101-3000213103000332-3033030220220301-3321110123121210-3021212000212000-1202001200122110-1121122303001302-3302233022031333)
- [direct_connect_enabled.hosted_vifs.site_registration_over_internet](data-sources--aws_tgw_site--reference--group-002.md#canonical-0120213130102001-3131030112310103-0301302322131302-1322310201013033-1031030121332323-2001310120030023-0201320303102323-0033212103321321)
- [direct_connect_enabled.hosted_vifs.vif_list](data-sources--aws_tgw_site--reference--group-002.md#canonical-1120303330300310-1302023120223012-0010311013220332-1012220130010311-1123313031102023-3232232001032113-0033323222012230-1233220122203211)
- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1323200312211101-3000213103000332-3033030220220301-3321110123121210-3021212000212000-1202001200122110-1121122303001302-3302233022031333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113323302000112-3023002201101322-1120120132030113-1103330020220121-2201210112102333-2310103231330012-1200003013333303-1231003123331100"></a>

## direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect — site_registration_over_direct_connect / 321123010312 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_tgw_site--reference--group-002.md#canonical-3122102210321312-3222221303132221-0221221122132303-0120123033113313-3311232031212022-0120001313113130-0321132113330030-2001102330231013)
- direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect

<a id="canonical-2321101003301320-0132020133333132-0201310101022302-1303010132031330-3332032102203103-0323220200012211-2312021210220302-0020311330210000"></a>

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

<a id="canonical-0131131222211321-3212220120213030-1003201101031123-0211102312020302-1030112023112013-1231103330220323-2033333313113202-0122033203331023"></a>

## Direct properties — site_registration_over_direct_connect / 321123010312 / 3

<a id="canonical-2323203213210011-3013013213032101-2202113110121330-2312030232233132-3330211231013102-1303223100100012-3233311003322121-3003200011102110"></a>

<a id="canonical-0320211310333313-3002102110300022-3100201320202011-1203021323132223-2103113323321330-0332131320301202-1303333111211313-1003031002101211"></a>

## cloudlink_network_name property — site_registration_over_direct_connect / 321123010312 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2333222212212123-0213212001311232-2312121233200210-2112011333323010-0210003012011010-2320131010020102-2330023130302010-0200130230223211"></a>

## Next pages — site_registration_over_direct_connect / 321123010312 / 5

- [direct_connect_enabled.hosted_vifs](data-sources--aws_tgw_site--reference--group-002.md#canonical-3122102210321312-3222221303132221-0221221122132303-0120123033113313-3311232031212022-0120001313113130-0321132113330030-2001102330231013)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0120213130102001-3131030112310103-0301302322131302-1322310201013033-1031030121332323-2001310120030023-0201320303102323-0033212103321321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111133303120213-2110120120032323-3230222100100102-3110223330002333-3323010222200210-0222030213030020-3210212123332003-2021031111003032"></a>

## direct_connect_enabled.hosted_vifs.site_registration_over_internet — site_registration_over_internet / 000102020123 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_tgw_site--reference--group-002.md#canonical-3122102210321312-3222221303132221-0221221122132303-0120123033113313-3311232031212022-0120001313113130-0321132113330030-2001102330231013)
- direct_connect_enabled.hosted_vifs.site_registration_over_internet

<a id="canonical-0100313110132030-3220332211321231-0222031210222120-0103203001201320-0020122311031132-2110002113011101-3330201203203203-3030312030000013"></a>

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

<a id="canonical-0002300033220023-3320212312102212-1222301113103301-0033031000202102-3122302133231233-2200221011222201-1211100202212212-0030101123021200"></a>

## Direct properties — site_registration_over_internet / 000102020123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003312231132221-1201213301001130-0032013323333020-2331023200303120-1323103202031121-1202211102312100-3113231131332033-2203222130301302"></a>

## Next pages — site_registration_over_internet / 000102020123 / 4

- [direct_connect_enabled.hosted_vifs](data-sources--aws_tgw_site--reference--group-002.md#canonical-3122102210321312-3222221303132221-0221221122132303-0120123033113313-3311232031212022-0120001313113130-0321132113330030-2001102330231013)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1120303330300310-1302023120223012-0010311013220332-1012220130010311-1123313031102023-3232232001032113-0033323222012230-1233220122203211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023203310323013-2033212122320330-3213301132100230-2112321021012111-1230121223201221-2112130321011300-3020221013230230-2331330103000103"></a>

## direct_connect_enabled.hosted_vifs.vif_list — vif_list / 303331032303 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_tgw_site--reference--group-002.md#canonical-3122102210321312-3222221303132221-0221221122132303-0120123033113313-3311232031212022-0120001313113130-0321132113330030-2001102330231013)
- direct_connect_enabled.hosted_vifs.vif_list

<a id="canonical-0333121031232201-1332232200302310-1001210212202202-1123012102310122-1013223000010213-0030021221122333-2233001331333303-2131202220112232"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1002123220021230-3021003102211323-1001010231311121-0303003313203123-1302213303333001-2022232222232231-1230313323022010-3300213201110000"></a>

## Direct properties — vif_list / 303331032303 / 3

<a id="canonical-2231210011203332-1313212310323222-1230100030000122-2231100312221312-2100221131033012-3300232121302330-1322103101023110-2202203222323023"></a>

<a id="canonical-2111223111132000-2231201321033122-1010223110302301-3023111202203130-1231012003001111-1122320122020023-1303220113232011-3200100231022003"></a>

## other_region property — vif_list / 303331032303 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [same_as_site_region](data-sources--aws_tgw_site--reference--group-002.md#canonical-2121021030232021-3233123102000303-3212113213202013-1003220133131200-0213323212123213-2003321223123133-3113101301110002-1123000323030020): complete subsection reference.

<a id="canonical-2002200330330311-2230213211011122-3112203120022233-1102022022011231-2202333301120221-0133133011000032-1230100222110302-3020323211313320"></a>

<a id="canonical-2001131331102012-2331313310222021-2233321301332003-2101032211101231-1331220332201032-2102223333131301-3231312201223133-0010122302233011"></a>

## vif_id property — vif_list / 303331032303 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3110222013231103-0210102013102203-3232000310113231-0020031233132310-2202130313200333-3201103223131303-1120132201211001-1022210132312000"></a>

## Next pages — vif_list / 303331032303 / 6

- [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](data-sources--aws_tgw_site--reference--group-002.md#canonical-2121021030232021-3233123102000303-3212113213202013-1003220133131200-0213323212123213-2003321223123133-3113101301110002-1123000323030020)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_tgw_site--reference--group-002.md#canonical-3122102210321312-3222221303132221-0221221122132303-0120123033113313-3311232031212022-0120001313113130-0321132113330030-2001102330231013)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2121021030232021-3233123102000303-3212113213202013-1003220133131200-0213323212123213-2003321223123133-3113101301110002-1123000323030020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112120030212001-0303302030121011-2310101201132223-2320223010032201-0211222121321121-3303201202313311-0233002320033100-2001231213210122"></a>

## direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region — same_as_site_region / 230100030020 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300)
- [direct_connect_enabled.hosted_vifs](data-sources--aws_tgw_site--reference--group-002.md#canonical-3122102210321312-3222221303132221-0221221122132303-0120123033113313-3311232031212022-0120001313113130-0321132113330030-2001102330231013)
- [direct_connect_enabled.hosted_vifs.vif_list](data-sources--aws_tgw_site--reference--group-002.md#canonical-1120303330300310-1302023120223012-0010311013220332-1012220130010311-1123313031102023-3232232001032113-0033323222012230-1233220122203211)
- direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region

<a id="canonical-1011200330323330-2323320003233213-0131313033300302-2101223133130123-3113300002101103-1130132111212131-2022101233332032-3011110300312020"></a>

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

<a id="canonical-2111301100020230-0033000231332333-3123113033112303-0002112000202212-0133331123131223-3131300112120220-2120213323122100-3201320033130102"></a>

## Direct properties — same_as_site_region / 230100030020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230302320301122-2201112320200312-2203102223230323-2031203321202201-0012003320303111-2122122102003010-2331132002301310-0100233110132101"></a>

## Next pages — same_as_site_region / 230100030020 / 4

- [direct_connect_enabled.hosted_vifs.vif_list](data-sources--aws_tgw_site--reference--group-002.md#canonical-1120303330300310-1302023120223012-0010311013220332-1012220130010311-1123313031102023-3232232001032113-0033323222012230-1233220122203211)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3223331130303203-1032102332230311-1312022330000323-1322110023132332-0130230133222002-2002122012033202-2312021211203030-2002121030100230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231211102031310-2332032220211220-3000321121021210-2231131100321202-3012313123011320-0001320013122302-2231022200232100-0232300011132121"></a>

## direct_connect_enabled.standard_vifs — standard_vifs / 303321120132 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300)
- direct_connect_enabled.standard_vifs

<a id="canonical-0021132032220222-3313030201213213-2113013132001223-1322002331101133-0010201200123023-0210113210223122-0122310002230013-0201233233011023"></a>

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

<a id="canonical-2132132023121332-3212110231203231-3132102013301022-1122220210300020-0031000233230322-1020001200320200-2211302012232312-0202010212310012"></a>

## Direct properties — standard_vifs / 303321120132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023131023231320-0323202320120110-2132303111312130-0122021022010101-2102222320221322-2011312332223133-1213010020031012-2102002222113113"></a>

## Next pages — standard_vifs / 303321120132 / 4

- [direct_connect_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-0312112211011132-0333103021031333-1131302121312230-2113231022311313-1112020011110330-3313230111300202-1031033033320221-1032311103133300)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0322103200110120-1212132221101122-3031220203231011-0212323213322303-1120210312111203-2300021113210022-2033113122011000-2322120103122230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111113001223230-1102121201223223-0022330000330302-0100032203122210-3212001131012200-3131300310222100-2102231320313133-2100113103111020"></a>

## kubernetes_upgrade_drain — kubernetes_upgrade_drain / 031031131222 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- kubernetes_upgrade_drain

<a id="canonical-1122122331212220-3032221130130121-0313003211221012-2230033000102231-2010301113220211-2131223323213320-1032133012300310-1211323311001222"></a>

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

<a id="canonical-3130201310001102-2311312211121223-2331013100323131-2311210033332300-2221232230313310-1012301312300131-2113030110002300-2233030132303002"></a>

## Direct properties — kubernetes_upgrade_drain / 031031131222 / 3

- [disable_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-1330022323302122-2313111301230123-3111011212023010-0220003010013113-0323002001011103-2223311323220103-3222001000222120-0110033132122333): complete subsection reference.

- [enable_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-1202323001201332-2133320121322310-0200022021023023-3332322322123300-1123331201310111-1231333201300312-3311021020331211-1121131303200000): complete subsection reference.

<a id="canonical-2123033321001033-2100131201112132-3001103331122213-0303302220200331-2123010312000302-2012020303312300-3233010221300333-3113210302113103"></a>

## Next pages — kubernetes_upgrade_drain / 031031131222 / 4

- [kubernetes_upgrade_drain.disable_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-1330022323302122-2313111301230123-3111011212023010-0220003010013113-0323002001011103-2223311323220103-3222001000222120-0110033132122333)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-1202323001201332-2133320121322310-0200022021023023-3332322322123300-1123331201310111-1231333201300312-3311021020331211-1121131303200000)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1330022323302122-2313111301230123-3111011212023010-0220003010013113-0323002001011103-2223311323220103-3222001000222120-0110033132122333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221222331101020-3011320021232330-0103303011013312-3033310202103013-3122011002132311-2323012330001311-2232030200223000-3023130133233221"></a>

## kubernetes_upgrade_drain.disable_upgrade_drain — disable_upgrade_drain / 013123100210 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [kubernetes_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-0322103200110120-1212132221101122-3031220203231011-0212323213322303-1120210312111203-2300021113210022-2033113122011000-2322120103122230)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-1021212301023113-1001322031023123-2331000332002122-2210103203322303-0320221312311011-3121211231300203-2313212333102212-2013311031300321"></a>

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

<a id="canonical-1310003301210212-2012212223030222-1012012331311323-0302020112211013-3100222212220012-0021322233132002-2133201311030313-0012323123101001"></a>

## Direct properties — disable_upgrade_drain / 013123100210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021002113110012-2330001121131322-0120221333301212-1023323022121003-3000012333300322-1233322122220111-2313011200312221-1020233322213330"></a>

## Next pages — disable_upgrade_drain / 013123100210 / 4

- [kubernetes_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-0322103200110120-1212132221101122-3031220203231011-0212323213322303-1120210312111203-2300021113210022-2033113122011000-2322120103122230)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1202323001201332-2133320121322310-0200022021023023-3332322322123300-1123331201310111-1231333201300312-3311021020331211-1121131303200000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012000030311223-3021213132112012-0302021132000320-2011023103311323-1310200320303302-2331301000022133-3200311110212001-3230131030022332"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain — enable_upgrade_drain / 303202011121 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [kubernetes_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-0322103200110120-1212132221101122-3031220203231011-0212323213322303-1120210312111203-2300021113210022-2033113122011000-2322120103122230)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-3303003312323121-3230301232021123-0011200130203332-2313202303331010-1221111311231320-0232233001003320-3123030210020223-2112111333111202"></a>

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

<a id="canonical-2321233321301033-2323033330120123-2011212203333030-1321122130101300-1222320330232131-3212320012322212-2322231332110010-2010021312122123"></a>

## Direct properties — enable_upgrade_drain / 303202011121 / 3

- [disable_vega_upgrade_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-2000003330112002-2003330021300013-0000200211203121-3030213033233232-0303121322330021-2303201310233311-2303330002230123-0223110013320030): complete subsection reference.

<a id="canonical-1013312200111212-1302123211112313-0320023312020132-3023231031201202-3222120323002232-1202320102023033-1001211210203322-0111321032033113"></a>

<a id="canonical-1103103321133112-1112111031103200-3130001230213332-1030210013120211-0220001113230311-3312131102302013-3230200131230033-1220103000310231"></a>

## drain_max_unavailable_node_count property — enable_upgrade_drain / 303202011121 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3230133031301232-1222330323310331-1012011001303101-2030010222000001-2100301020321122-1221011021200113-0232202121233220-2220332200220200"></a>

<a id="canonical-0003313322300331-2102102010212033-0221110023223012-1130133110233030-0332101222033313-1230133313231232-3131001303130002-0220113033221333"></a>

## drain_max_unavailable_node_percentage property — enable_upgrade_drain / 303202011121 / 5

Type: `"number"`. Computed.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-0223320123313300-1203200203010303-3033231212310313-0302113102111120-2201020113220113-1211132011032313-1001010011231223-3112011211001032"></a>

<a id="canonical-2030122202003010-0233120112303023-0122201203111120-0233113122012121-2230113233331123-0110122202212323-2301321102321313-1113230011031120"></a>

## drain_node_timeout property — enable_upgrade_drain / 303202011121 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [enable_vega_upgrade_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0331112212013002-2203201322101103-1011202201100112-1003001000011020-0210310013012130-2001021333323133-3131012130123301-1310021333331000): complete subsection reference.

<a id="canonical-3123001321310311-2220233003210203-0321020210230323-3100011331223001-1302310302100031-1033102222131211-1202113123103321-2131013213212203"></a>

## Next pages — enable_upgrade_drain / 303202011121 / 7

- [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-2000003330112002-2003330021300013-0000200211203121-3030213033233232-0303121322330021-2303201310233311-2303330002230123-0223110013320030)
- [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0331112212013002-2203201322101103-1011202201100112-1003001000011020-0210310013012130-2001021333323133-3131012130123301-1310021333331000)
- [kubernetes_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-0322103200110120-1212132221101122-3031220203231011-0212323213322303-1120210312111203-2300021113210022-2033113122011000-2322120103122230)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2000003330112002-2003330021300013-0000200211203121-3030213033233232-0303121322330021-2303201310233311-2303330002230123-0223110013320030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223233133002023-1101033110012202-0313311112200231-2221331010033010-0321101330321020-2112011121310121-0230202330313222-2211102001213123"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — disable_vega_upgrade_mode / 203210030213 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [kubernetes_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-0322103200110120-1212132221101122-3031220203231011-0212323213322303-1120210312111203-2300021113210022-2033113122011000-2322120103122230)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-1202323001201332-2133320121322310-0200022021023023-3332322322123300-1123331201310111-1231333201300312-3311021020331211-1121131303200000)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-2012012130000103-1102333031131303-2001003112030121-2233320123033111-1202331330123223-2222101111321011-1321233100230301-1223003101022221"></a>

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

<a id="canonical-1302323102211003-0323130200123033-0310132032003103-2321321230101112-2233203233331333-1011301332032313-0102123111203003-0123230031113002"></a>

## Direct properties — disable_vega_upgrade_mode / 203210030213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0031202310202031-0320222103023031-0122031131312232-2322213220211331-2200103200120303-0301220003311200-2321202333203002-1012233130312312"></a>

## Next pages — disable_vega_upgrade_mode / 203210030213 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-1202323001201332-2133320121322310-0200022021023023-3332322322123300-1123331201310111-1231333201300312-3311021020331211-1121131303200000)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0331112212013002-2203201322101103-1011202201100112-1003001000011020-0210310013012130-2001021333323133-3131012130123301-1310021333331000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000311013102013-3210110210002231-1203133230321022-1011020311132203-3211313220010133-2312233030220200-0331223312203231-1332313023311010"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — enable_vega_upgrade_mode / 211303322110 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [kubernetes_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-0322103200110120-1212132221101122-3031220203231011-0212323213322303-1120210312111203-2300021113210022-2033113122011000-2322120103122230)
- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-1202323001201332-2133320121322310-0200022021023023-3332322322123300-1123331201310111-1231333201300312-3311021020331211-1121131303200000)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-1012331302012031-2212033001331030-2002301103211302-0031112311231022-2211122111112203-3321210333112302-3023013103021201-3111322130313001"></a>

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

<a id="canonical-0203220220102200-1231103012001310-1002331200231133-0311100230103210-0020210011100133-2100032312120100-0003222230332221-2033112022222322"></a>

## Direct properties — enable_vega_upgrade_mode / 211303322110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232233000000001-2231300333232020-2201322233133200-2021120011101112-3013033112213112-0230122213012120-1123113123130020-3033000222111132"></a>

## Next pages — enable_vega_upgrade_mode / 211303322110 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](data-sources--aws_tgw_site--reference--group-002.md#canonical-1202323001201332-2133320121322310-0200022021023023-3332322322123300-1123331201310111-1231333201300312-3311021020331211-1121131303200000)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2020021301323331-3001033223203323-2332033222202220-3311033003321302-1031013201223320-2333333000113313-1000012203232223-3002210130211211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311302232223202-2002110333203101-1013023313132030-3210031112201212-1333202133210301-1122131111220333-1201212033030022-3212223002223001"></a>

## log_receiver — log_receiver / 203131102020 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- log_receiver

<a id="canonical-3030230100032313-2200010231230310-1212303202312110-0123013031313320-2201210013021220-0221212203303210-3103200302312121-0213011232213332"></a>

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

- [log_receiver](data-sources--aws_tgw_site--reference--group-002.md#canonical-3030230100032313-2200010231230310-1212303202312110-0123013031313320-2201210013021220-0221212203303210-3103200302312121-0213011232213332)
- [logs_streaming_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-1311031232301110-0101211333020012-2300103001323231-0300221120113121-0111103220222123-0222120000232331-1032131220021020-1031000220022031)

Select alternatives according to the provider validators above.

<a id="canonical-0323220223231200-1312201120000102-1030332321221231-0033122320111010-2322222122320323-2231010321203111-2313023220001122-3132021311121001"></a>

## Direct properties — log_receiver / 203131102020 / 3

<a id="canonical-1102033123210212-1211003203000332-0031232321020330-3023210202122331-0122200320332220-2210100123030002-2001202231100220-0000111022122211"></a>

<a id="canonical-1200002010030032-0010131032333031-3033010333030213-2210203010213203-2003100101101231-1202101131102300-3012301102023130-0032221230322033"></a>

## name property — log_receiver / 203131102020 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0232000112231101-2210221322332000-3132210322200313-1233310013122120-2130311213331213-3221031302311310-0231223123302333-1322202323211011"></a>

<a id="canonical-2011330010202101-3100130031200130-0033013333033113-3112102103023101-2020300321102213-2202210111231033-1233002321001001-1302003132112131"></a>

## namespace property — log_receiver / 203131102020 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0020300220133220-0212313132101101-3100200031123220-1101332023221222-2231233311301010-1321012230212012-1332203133031323-1131112032232111"></a>

<a id="canonical-2002102201001230-1313321210032332-1200203012113311-3032201311311230-1103212010113013-0331110112103130-1213120122022121-3123000232110011"></a>

## tenant property — log_receiver / 203131102020 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1333220322031022-0202133332030033-0132123022310113-1122231301322311-3330220200030112-3310303230213122-0302010000113203-3221213231220113"></a>

## Next pages — log_receiver / 203131102020 / 7

- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1113021113220301-2010123223000213-3012113321101301-3220123032012310-0000133322303120-1033130320232331-1300213221112233-1232321020023000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032212333300303-1031222233302031-3020130331020210-0313312100310132-1312021023221100-1322230333122230-2010201200322223-3113120101301010"></a>

## logs_streaming_disabled — logs_streaming_disabled / 311131201322 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- logs_streaming_disabled

<a id="canonical-1311031232301110-0101211333020012-2300103001323231-0300221120113121-0111103220222123-0222120000232331-1032131220021020-1031000220022031"></a>

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

<a id="canonical-0032012033002320-2320130210300020-2310011211113233-3113023220113321-3301200000102201-1202001231230320-1033312012101320-3322302211332030"></a>

## Direct properties — logs_streaming_disabled / 311131201322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211101133220121-1123321011230000-0003120331102333-1320222003322303-0123031131012310-0233120031300112-1220010113031132-2201001322302213"></a>

## Next pages — logs_streaming_disabled / 311131201322 / 4

- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1002101032011221-3313311332231033-1320120021332312-1030110212030332-2020111312332302-2203213132102110-0110133312331031-0330022132023210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211120021232130-2133210022202210-3222100213230230-0033110221333013-0323333032202031-1321313300311120-1311320132102121-3202122002111121"></a>

## offline_survivability_mode — offline_survivability_mode / 332020220020 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- offline_survivability_mode

<a id="canonical-0121322102202200-1331102101131233-1223212212232333-0003220302220012-3222203311222213-0012212203333031-3200021112131013-3230321023213130"></a>

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

<a id="canonical-2022021033002023-1131220011322330-3010323222021113-2000013203122213-2110221302230313-1211031110211321-0002002321331000-1102331332033211"></a>

## Direct properties — offline_survivability_mode / 332020220020 / 3

- [enable_offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-1011032221023211-3332303213033000-2311323210133312-3222031002030013-3021133130302030-1320110013200131-3100232112212222-0133032303031231): complete subsection reference.

- [no_offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-3300233122100233-1231203331022330-2130331002022111-2120312301121301-2100010122020002-2301223120133203-1321213031130210-3003030201201120): complete subsection reference.

<a id="canonical-3100213210030003-1131323022323102-3123123222213321-3201111120323311-0023202331312102-1133313010310030-3102312102133211-2000332000233333"></a>

## Next pages — offline_survivability_mode / 332020220020 / 4

- [offline_survivability_mode.enable_offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-1011032221023211-3332303213033000-2311323210133312-3222031002030013-3021133130302030-1320110013200131-3100232112212222-0133032303031231)
- [offline_survivability_mode.no_offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-3300233122100233-1231203331022330-2130331002022111-2120312301121301-2100010122020002-2301223120133203-1321213031130210-3003030201201120)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1011032221023211-3332303213033000-2311323210133312-3222031002030013-3021133130302030-1320110013200131-3100232112212222-0133032303031231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011223100222101-1202132102130131-2021301311201333-0210323203331300-3003011232331020-3111112221332200-2131113301301102-3312002131030303"></a>

## offline_survivability_mode.enable_offline_survivability_mode — enable_offline_survivability_mode / 023032231230 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-1002101032011221-3313311332231033-1320120021332312-1030110212030332-2020111312332302-2203213132102110-0110133312331031-0330022132023210)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-0102012010001030-3220201221032100-2120133313313132-0123130011110111-3101000221211001-0010312231021102-3100100130032023-0132030020030311"></a>

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

<a id="canonical-0113001320221202-2312010010100300-2032110210232033-1002001130323021-3312330333112230-2302220100320121-2132221201312312-0013003233310022"></a>

## Direct properties — enable_offline_survivability_mode / 023032231230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213232121213203-1303101100332231-3301330332233021-3022022312332133-0101330331102330-2111332322123103-0120002113332213-0021111231000310"></a>

## Next pages — enable_offline_survivability_mode / 023032231230 / 4

- [offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-1002101032011221-3313311332231033-1320120021332312-1030110212030332-2020111312332302-2203213132102110-0110133312331031-0330022132023210)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3300233122100233-1231203331022330-2130331002022111-2120312301121301-2100010122020002-2301223120133203-1321213031130210-3003030201201120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213002313223100-0112131303133101-0130313320323113-1101320001032023-2221330032010320-3032332221030300-1203031101001000-0330220220232212"></a>

## offline_survivability_mode.no_offline_survivability_mode — no_offline_survivability_mode / 003013313220 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-1002101032011221-3313311332231033-1320120021332312-1030110212030332-2020111312332302-2203213132102110-0110133312331031-0330022132023210)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-2231222223113003-3000113203203111-1012301003122221-0000113331323212-2203010210322130-2021323312022301-2103300333001023-0102222231221303"></a>

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

<a id="canonical-3122133212200301-3230030321211313-3030322320201210-1320122031003022-2213303231123012-3210201222231301-3331033100333011-1230220202210003"></a>

## Direct properties — no_offline_survivability_mode / 003013313220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030121310130202-0311311131322031-3022110033223321-0122103001112323-3131333001233011-1002122121210200-3221121120022131-3012230021001120"></a>

## Next pages — no_offline_survivability_mode / 003013313220 / 4

- [offline_survivability_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-1002101032011221-3313311332231033-1320120021332312-1030110212030332-2020111312332302-2203213132102110-0110133312331031-0330022132023210)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1131211032312230-3213210020110320-2302211020330112-2302302311111200-3123111312222202-0112321312131023-1133232031010020-1131203321212130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110321120323232-3303111111221012-1121303201121311-0022013202301131-1220100110301000-0120001021002233-0230222020213330-2120000313110123"></a>

## os — os / 213002331001 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- os

<a id="canonical-1232012323003122-0120303332133123-2220313303021232-3032131333312031-2332331321012213-0130213020300022-2332023032222111-0213321231100301"></a>

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

<a id="canonical-0010303313201221-0000011022231231-2313301112120312-2021000021030221-3023100231102313-1313122112002323-2202323010012211-1033202312033132"></a>

## Direct properties — os / 213002331001 / 3

- [default_os_version](data-sources--aws_tgw_site--reference--group-002.md#canonical-0322333210121233-1023320320022331-3201323132302200-3120102220303122-2323322330133213-0220030221233021-2203002010203002-1303231022302133): complete subsection reference.

<a id="canonical-3000032112303030-1002121020003022-0033031230011321-3111330301032002-2132332130202133-3301201213132032-1012012233313102-3322322232331010"></a>

<a id="canonical-0103112221122013-1110330002322333-1303020010201011-1231003231001310-0322310223031121-3311300200203032-2030311133003222-2102033031301211"></a>

## operating_system_version property — os / 213002331001 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1011210223033221-2021331301232033-1321213010211220-2000311312011323-0100330223303122-3022011332221323-2210001203002021-2212333201232323"></a>

## Next pages — os / 213002331001 / 5

- [os.default_os_version](data-sources--aws_tgw_site--reference--group-002.md#canonical-0322333210121233-1023320320022331-3201323132302200-3120102220303122-2323322330133213-0220030221233021-2203002010203002-1303231022302133)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0322333210121233-1023320320022331-3201323132302200-3120102220303122-2323322330133213-0220030221233021-2203002010203002-1303231022302133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022301130132301-3310003220201332-0102322332011210-0213022102132000-3010101232110122-1230022030221231-0000112331102121-2203221311301130"></a>

## os.default_os_version — default_os_version / 032201123230 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [os](data-sources--aws_tgw_site--reference--group-002.md#canonical-1131211032312230-3213210020110320-2302211020330112-2302302311111200-3123111312222202-0112321312131023-1133232031010020-1131203321212130)
- os.default_os_version

<a id="canonical-2023023303222013-3211003212023212-0030332022103301-1322023030320033-0123212232130033-2331311012021033-1132330010310023-3103310322022200"></a>

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

<a id="canonical-2111002300212210-1003333311302210-1110101300200010-1223230011131331-3221311320322303-2122000213220011-1131213112322013-2031012323030212"></a>

## Direct properties — default_os_version / 032201123230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000331001312232-1100223211011213-0332011111200132-0023031133223023-1311322220333211-3302321321000303-2122330110322103-3310210000031112"></a>

## Next pages — default_os_version / 032201123230 / 4

- [os](data-sources--aws_tgw_site--reference--group-002.md#canonical-1131211032312230-3213210020110320-2302211020330112-2302302311111200-3123111312222202-0112321312131023-1133232031010020-1131203321212130)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332122211010102-1301032302023322-2122203301001001-2002122201230232-2000331320320231-1232233231100332-1222010331022333-2300333232033211"></a>

## performance_enhancement_mode — performance_enhancement_mode / 221203212312 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- performance_enhancement_mode

<a id="canonical-1101020121232223-3210011003013000-0223332002021320-2201131210333333-2013100221112302-2031013202200013-1321130210022200-3220232202332300"></a>

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

<a id="canonical-3010311211222213-0102100123102111-3130311212001303-3302322010203212-1312201131101222-1310300033122133-3222311101321021-3101313320320122"></a>

## Direct properties — performance_enhancement_mode / 221203212312 / 3

- [perf_mode_l3_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1133333100002302-3232021033103203-1230003331113222-3301103200121032-1120223010312210-3333203212032301-0213223330311022-0213100001300301): complete subsection reference.

- [perf_mode_l7_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1030223221012011-2300322110132122-1032100131321021-2323033101312013-0320002202012100-1011233011011013-0130311210123331-3213001310030103): complete subsection reference.

<a id="canonical-0130013312011211-3331300211303222-2131123100113023-1310110320301113-1331100032112020-1013113213303321-2210010032300110-2331203302200130"></a>

## Next pages — performance_enhancement_mode / 221203212312 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1133333100002302-3232021033103203-1230003331113222-3301103200121032-1120223010312210-3333203212032301-0213223330311022-0213100001300301)
- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1030223221012011-2300322110132122-1032100131321021-2323033101312013-0320002202012100-1011233011011013-0130311210123331-3213001310030103)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1133333100002302-3232021033103203-1230003331113222-3301103200121032-1120223010312210-3333203212032301-0213223330311022-0213100001300301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333233301301212-3110222212101122-0003301111100213-2210200213020221-1132231333300021-3323200131312033-0220323331103132-1000331101112202"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced — perf_mode_l3_enhanced / 201131301221 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332)
- performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-3030233020320202-2123120321021312-3301010310332203-0312012120000012-2020033002111011-3130201222210002-3122331110102120-2332221031222311"></a>

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

<a id="canonical-2101120033211331-0230321120132022-3113312100002001-0102131200020121-0002123011302221-2320323000323220-2213330010311222-1323301202111022"></a>

## Direct properties — perf_mode_l3_enhanced / 201131301221 / 3

- [jumbo](data-sources--aws_tgw_site--reference--group-002.md#canonical-1333121320231111-0230331332033100-1133312012012321-3202200102302210-2100311101213232-2013302003103201-1031331333010121-1213201300020230): complete subsection reference.

- [no_jumbo](data-sources--aws_tgw_site--reference--group-002.md#canonical-3110022120220222-2223030213102012-2230110120123112-2110000002233112-2223203023002323-3332312001322101-2313202212321201-3100003211301111): complete subsection reference.

<a id="canonical-0331233120302231-3100332213122301-1331121030030013-1133330002233032-1322003022003333-3002002103202100-1130001321011320-3221321130013311"></a>

## Next pages — perf_mode_l3_enhanced / 201131301221 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--aws_tgw_site--reference--group-002.md#canonical-1333121320231111-0230331332033100-1133312012012321-3202200102302210-2100311101213232-2013302003103201-1031331333010121-1213201300020230)
- [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--aws_tgw_site--reference--group-002.md#canonical-3110022120220222-2223030213102012-2230110120123112-2110000002233112-2223203023002323-3332312001322101-2313202212321201-3100003211301111)
- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1333121320231111-0230331332033100-1133312012012321-3202200102302210-2100311101213232-2013302003103201-1031331333010121-1213201300020230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331132323001230-3223100012110011-3123300221312032-2120021303111032-0020222000022312-1010020001030010-2010223011100030-2223301013331201"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — jumbo / 033002012012 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332)
- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1133333100002302-3232021033103203-1230003331113222-3301103200121032-1120223010312210-3333203212032301-0213223330311022-0213100001300301)
- performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-1120100000332311-0213302031131332-2130020033033212-1321032031113103-2312332111020132-1201312132100100-2131100022230023-3111333101220000"></a>

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

<a id="canonical-2221031030303322-1010321313000301-0300221231310203-3210202300310023-2122233112102332-0021032331021331-2312003010321023-0212300211102202"></a>

## Direct properties — jumbo / 033002012012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131201030121200-3112213100213120-1012110012220032-1031201211331321-2023032002102123-0233031030330310-1011230103233330-0033001130023101"></a>

## Next pages — jumbo / 033002012012 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1133333100002302-3232021033103203-1230003331113222-3301103200121032-1120223010312210-3333203212032301-0213223330311022-0213100001300301)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3110022120220222-2223030213102012-2230110120123112-2110000002233112-2223203023002323-3332312001322101-2313202212321201-3100003211301111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220113102201302-3130122323013133-1300331223031022-0201231123213032-2310130333020202-2330301210331201-2210002311212103-0232223321313322"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — no_jumbo / 223330120223 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332)
- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1133333100002302-3232021033103203-1230003331113222-3301103200121032-1120223010312210-3333203212032301-0213223330311022-0213100001300301)
- performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-2303321201113123-0310112320213100-0022301322230013-2212031111120202-0222112200233211-0323210013003203-3213000123031322-1012033332102033"></a>

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

<a id="canonical-2123020011320221-1232122033332133-2100220003012113-0221000123312131-1201103011020101-3003202310120320-3031123100322200-3113303111303332"></a>

## Direct properties — no_jumbo / 223330120223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113223033311322-0302021001310231-2233212322130012-2000302212123213-2202103211320031-0213032013213000-3212020302103211-3130022300113130"></a>

## Next pages — no_jumbo / 223330120223 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1133333100002302-3232021033103203-1230003331113222-3301103200121032-1120223010312210-3333203212032301-0213223330311022-0213100001300301)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1030223221012011-2300322110132122-1032100131321021-2323033101312013-0320002202012100-1011233011011013-0130311210123331-3213001310030103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310000221002212-2010301022032310-3230011213120223-1000312000300212-2201233030202123-2121310303003202-0131112020132133-0320022120033013"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced — perf_mode_l7_enhanced / 002031023130 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-0313233110130320-1012221112013120-1320201110033212-0310132123111323-2120230311233202-2201201200311223-1310202120033202-3323211112010332"></a>

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

<a id="canonical-2100223112333230-0132012312123302-3330013320231221-1002220231023133-0220101211212220-0020323013131302-0232330221213201-2002201232312000"></a>

## Direct properties — perf_mode_l7_enhanced / 002031023130 / 3

- [jumbo_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-2002033222131000-3223103011213132-3103020002031101-1232311132213111-3330212131311011-1132030000123012-0133332031312303-3132323210032210): complete subsection reference.

- [jumbo_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-3201323033210031-3213130333211311-0133122020313320-1121122310102210-3332100013303322-2123123112111313-2030013320320303-2022212231100112): complete subsection reference.

<a id="canonical-1231333131113123-0013300032023322-2013213022212322-3321233231223030-3020023322000313-3320302223002203-1211131003113213-2100313223030213"></a>

## Next pages — perf_mode_l7_enhanced / 002031023130 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-2002033222131000-3223103011213132-3103020002031101-1232311132213111-3330212131311011-1132030000123012-0133332031312303-3132323210032210)
- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--aws_tgw_site--reference--group-002.md#canonical-3201323033210031-3213130333211311-0133122020313320-1121122310102210-3332100013303322-2123123112111313-2030013320320303-2022212231100112)
- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2002033222131000-3223103011213132-3103020002031101-1232311132213111-3330212131311011-1132030000123012-0133332031312303-3132323210032210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002120232013003-2110130220201020-3011111203320030-0012001001022132-1210322330101301-2312022030221012-0033123030330030-1002021233132100"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — jumbo_disabled / 333210323300 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332)
- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1030223221012011-2300322110132122-1032100131321021-2323033101312013-0320002202012100-1011233011011013-0130311210123331-3213001310030103)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-3203213023101120-2213121223303013-0021110100010223-0312131212003131-2101030021100213-2331301030302033-1210122111011301-0011201003333111"></a>

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

<a id="canonical-3301300313032113-3210310300120321-2012123313223012-3210011101030321-0303322313111321-1112103033300001-0132212223310031-0302122123301331"></a>

## Direct properties — jumbo_disabled / 333210323300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212010003323313-2022001021110201-1131132322230200-3301303023030013-0111311222321110-0222220232103201-1333333333222203-1210320121120021"></a>

## Next pages — jumbo_disabled / 333210323300 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1030223221012011-2300322110132122-1032100131321021-2323033101312013-0320002202012100-1011233011011013-0130311210123331-3213001310030103)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3201323033210031-3213130333211311-0133122020313320-1121122310102210-3332100013303322-2123123112111313-2030013320320303-2022212231100112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022100121020010-3100013030013013-3131301112110120-1032330100302301-3100311202322010-3301233320220112-0000201302103302-3320033220321330"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — jumbo_enabled / 023222333121 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [performance_enhancement_mode](data-sources--aws_tgw_site--reference--group-002.md#canonical-0330322032001013-0000120311203032-3130033011212311-1333301023200011-1202223013022233-0332323331112312-2233100022230001-1003200031300332)
- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1030223221012011-2300322110132122-1032100131321021-2323033101312013-0320002202012100-1011233011011013-0130311210123331-3213001310030103)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-0021300211302203-1220210311311133-0213221020033130-0102020023033133-2131012231303302-2120021000032232-3300233020321001-1112002320220121"></a>

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

<a id="canonical-0123033302203000-3133223133020001-3031300102002121-2323233313303330-1203003131331231-3131212022232313-1011011323000303-2213332121211112"></a>

## Direct properties — jumbo_enabled / 023222333121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301330201331102-3300110000033023-1003112203201233-0331300232322203-0002113211311112-1122221100103323-0103320231220003-0003321031000021"></a>

## Next pages — jumbo_enabled / 023222333121 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_tgw_site--reference--group-002.md#canonical-1030223221012011-2300322110132122-1032100131321021-2323033101312013-0320002202012100-1011233011011013-0130311210123331-3213001310030103)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2301310322233311-3020201320022232-0210231210312320-0210001101002120-0123230230101130-1201233323300232-2231101213013103-0112210022220330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210312021230111-2330001222320121-1313201313013033-2101123201233122-1321220111001132-1101330202131112-3322111222301200-1131200013310221"></a>

## private_connectivity — private_connectivity / 020032120001 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- private_connectivity

<a id="canonical-2330311303310000-3201101003320222-1332021212103110-0001203211113110-3100222000231102-3210211331232231-0231310132333032-2201013101212203"></a>

Type: `"single"`. Computed.

Configuration parameter for private connectivity.

Upstream description:

Private Connect Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_options": "[\"inside\",\"outside\"]"
}
```

<a id="canonical-2130233221330111-1320332230312320-2322312122011020-1013031103321203-3200200302233300-0130021213123030-3213330030102112-1310201221300130"></a>

## Direct properties — private_connectivity / 020032120001 / 3

- [cloud_link](data-sources--aws_tgw_site--reference--group-002.md#canonical-3321322232102332-2112201320310312-0311210203101010-0312213101303202-0232131102223122-3223222232103321-0310321102023012-2213200200022221): complete subsection reference.

- [inside](data-sources--aws_tgw_site--reference--group-002.md#canonical-2010132320123203-0002302222122002-0130011020313231-0202032223312101-1230131023013123-3313020112213310-0303101110111311-2032233030300313): complete subsection reference.

- [outside](data-sources--aws_tgw_site--reference--group-002.md#canonical-1323110210103000-3202031100211003-0332312012313311-2132130311232210-2023023030200222-0030112130333122-0131222013232010-2321112333022001): complete subsection reference.

<a id="canonical-3313230332330202-1001021211131301-0033302213121032-1222113122221332-3100333032100233-0110003001001233-0320203300103312-0030312311100023"></a>

## Next pages — private_connectivity / 020032120001 / 4

- [private_connectivity.cloud_link](data-sources--aws_tgw_site--reference--group-002.md#canonical-3321322232102332-2112201320310312-0311210203101010-0312213101303202-0232131102223122-3223222232103321-0310321102023012-2213200200022221)
- [private_connectivity.inside](data-sources--aws_tgw_site--reference--group-002.md#canonical-2010132320123203-0002302222122002-0130011020313231-0202032223312101-1230131023013123-3313020112213310-0303101110111311-2032233030300313)
- [private_connectivity.outside](data-sources--aws_tgw_site--reference--group-002.md#canonical-1323110210103000-3202031100211003-0332312012313311-2132130311232210-2023023030200222-0030112130333122-0131222013232010-2321112333022001)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3321322232102332-2112201320310312-0311210203101010-0312213101303202-0232131102223122-3223222232103321-0310321102023012-2213200200022221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321232033202202-2202232200031301-1331310222223302-3201201010010221-2130223313010222-1330333232031222-1320132122011013-2033121130000322"></a>

## private_connectivity.cloud_link — cloud_link / 322020020312 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [private_connectivity](data-sources--aws_tgw_site--reference--group-002.md#canonical-2301310322233311-3020201320022232-0210231210312320-0210001101002120-0123230230101130-1201233323300232-2231101213013103-0112210022220330)
- private_connectivity.cloud_link

<a id="canonical-0030010303001102-3211101233101010-1031003113032013-0003010312130110-1313122011002123-0021012031012203-0322032211223232-3000130131121322"></a>

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

<a id="canonical-1223130002012111-0331323003300131-2001321103100022-3330312110230210-1301033231233122-3201022120031202-0103322112120313-3200011323211020"></a>

## Direct properties — cloud_link / 322020020312 / 3

<a id="canonical-1110230310311330-3030213103330201-1132113100030031-3313120002221230-0313103030122001-3100020001330322-1101313133301031-3300231230030200"></a>

<a id="canonical-2003031133033133-2330222230333032-3222330001032200-3100113122320333-1130313023120231-0232201301312002-2011021123320031-3102221211213000"></a>

## name property — cloud_link / 322020020312 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1130022210102202-3121323120301202-3201331203010103-1321223103002313-3110033132022301-1102300212301323-3130230302010323-2311113332301010"></a>

<a id="canonical-0220123111323010-3301312221001131-0220231312201102-3122001232203301-3221121030320232-2202012033030112-2033202112112011-2023333232000200"></a>

## namespace property — cloud_link / 322020020312 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0303001311113201-2100003213330112-2111230133002301-1011212132331303-0131311223322113-3233310020100203-3200033221011232-3211202121021303"></a>

<a id="canonical-2331020213321013-2231301021032022-1213112022111220-2333112113030032-2030312220211132-2110211113121111-2222221201131010-1203232000320132"></a>

## tenant property — cloud_link / 322020020312 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3201310010130021-0111320232030112-0200123221012123-2100223011201203-3122031322210130-0310213222013222-2220123201133101-1231300222302133"></a>

## Next pages — cloud_link / 322020020312 / 7

- [private_connectivity](data-sources--aws_tgw_site--reference--group-002.md#canonical-2301310322233311-3020201320022232-0210231210312320-0210001101002120-0123230230101130-1201233323300232-2231101213013103-0112210022220330)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2010132320123203-0002302222122002-0130011020313231-0202032223312101-1230131023013123-3313020112213310-0303101110111311-2032233030300313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303002300111120-3333331231301120-2033120212110033-1133310213230001-0213012032312023-3223233212322032-2232112310323010-2322002133322211"></a>

## private_connectivity.inside — inside / 203302312033 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [private_connectivity](data-sources--aws_tgw_site--reference--group-002.md#canonical-2301310322233311-3020201320022232-0210231210312320-0210001101002120-0123230230101130-1201233323300232-2231101213013103-0112210022220330)
- private_connectivity.inside

<a id="canonical-2113000200131023-3030031001000010-2031030333032233-1130320311133321-1001132220112220-3112020210200002-2223233211113211-0131103322131020"></a>

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

<a id="canonical-0211003122031010-2012132310300022-1301322323230311-0033321012323112-2030001223121200-1121231332023030-3300131332220031-3003120121002002"></a>

## Direct properties — inside / 203302312033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010123001222001-2210012113322030-2033122012313030-0301231033021003-2223113302332011-3102320120130120-2122301321133103-1322011110221210"></a>

## Next pages — inside / 203302312033 / 4

- [private_connectivity](data-sources--aws_tgw_site--reference--group-002.md#canonical-2301310322233311-3020201320022232-0210231210312320-0210001101002120-0123230230101130-1201233323300232-2231101213013103-0112210022220330)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1323110210103000-3202031100211003-0332312012313311-2132130311232210-2023023030200222-0030112130333122-0131222013232010-2321112333022001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032022231111032-2132121212311010-0301001133122112-0133221230203030-2200311011123203-3021333023003032-1321223011133130-1333230322312011"></a>

## private_connectivity.outside — outside / 320202122302 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [private_connectivity](data-sources--aws_tgw_site--reference--group-002.md#canonical-2301310322233311-3020201320022232-0210231210312320-0210001101002120-0123230230101130-1201233323300232-2231101213013103-0112210022220330)
- private_connectivity.outside

<a id="canonical-1310323331313330-2313213231101102-1002212000230223-0302233232112023-2001011101231112-3011331220000223-1302130030333120-3200112231331203"></a>

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

<a id="canonical-0333221230112200-2010030202223331-1203032212323311-1033210003323330-1122231120211111-0132110221023211-2020033012311121-1202122323312333"></a>

## Direct properties — outside / 320202122302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022131000120320-2100001320322300-1211003322312210-2332333312312102-3101121200213213-2301233110310123-3212332322131032-0231312023300113"></a>

## Next pages — outside / 320202122302 / 4

- [private_connectivity](data-sources--aws_tgw_site--reference--group-002.md#canonical-2301310322233311-3020201320022232-0210231210312320-0210001101002120-0123230230101130-1201233323300232-2231101213013103-0112210022220330)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0202131032330112-1333222111331323-2111010000110002-0310200320103110-0231230201000123-3131020001033213-1033001213101002-0122303031233103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312002012331002-3021133322201011-0230322221121220-2313020213321010-0112312231032330-0100231231220113-0113001001233020-2221120300033222"></a>

## sw — sw / 003101120111 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- sw

<a id="canonical-3121000022210322-1121332102321311-2330332310212001-3020230120010333-3233120322221131-3221030110012133-1312011032321302-1100323303132121"></a>

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

<a id="canonical-0123220333002210-0211303300222022-3121110023132101-1210320223212202-1213313212002313-0102001230230313-1133001213110303-0121113012033021"></a>

## Direct properties — sw / 003101120111 / 3

- [default_sw_version](data-sources--aws_tgw_site--reference--group-002.md#canonical-3120222200220310-2320300020313321-1211200021031130-1233333330232332-0311032120221132-0311303311103321-0210021300213131-3221320102231110): complete subsection reference.

<a id="canonical-0330021302200232-2310211202123212-1331202201113120-2131102000102223-1310033221203321-1303033021202233-2230300133302213-0021203312101100"></a>

<a id="canonical-0310323002001233-0231202330002230-1013221310031120-2111022223113131-2033113232213310-3333312222220031-1332002011021210-1130333101011032"></a>

## volterra_software_version property — sw / 003101120111 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2010302322123133-0331100130331131-2313023310111300-0323331320312130-0020212312211122-1130330222311023-2030232211013101-2223110220303332"></a>

## Next pages — sw / 003101120111 / 5

- [sw.default_sw_version](data-sources--aws_tgw_site--reference--group-002.md#canonical-3120222200220310-2320300020313321-1211200021031130-1233333330232332-0311032120221132-0311303311103321-0210021300213131-3221320102231110)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3120222200220310-2320300020313321-1211200021031130-1233333330232332-0311032120221132-0311303311103321-0210021300213131-3221320102231110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120002001312211-0203033101132120-0003331111203233-1011323300102313-0320230030001313-1120122101212013-1123123113221032-3100032203211311"></a>

## sw.default_sw_version — default_sw_version / 122310201323 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [sw](data-sources--aws_tgw_site--reference--group-002.md#canonical-0202131032330112-1333222111331323-2111010000110002-0310200320103110-0231230201000123-3131020001033213-1033001213101002-0122303031233103)
- sw.default_sw_version

<a id="canonical-0213122220113013-2113220022300013-1200232313120203-2121301030033221-3020132013131221-0310011121112011-3023131312311131-0120230221100220"></a>

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

<a id="canonical-0302232010332001-3230123011202101-2223030210323312-0001323010322033-1000203231213212-3102020103222133-1110303230301032-1022021202110001"></a>

## Direct properties — default_sw_version / 122310201323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011030003300333-0310011320131210-3303013031233131-3212212012020202-0003212002030321-1310123301122130-2210021201121111-0033020100311322"></a>

## Next pages — default_sw_version / 122310201323 / 4

- [sw](data-sources--aws_tgw_site--reference--group-002.md#canonical-0202131032330112-1333222111331323-2111010000110002-0310200320103110-0231230201000123-3131020001033213-1033001213101002-0122303031233103)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121110133011230-2130332033230011-0122103312032313-2201032331232223-1303112030210132-0330232302313111-0231100211302011-1311202103000223"></a>

## tgw_security — tgw_security / 230302200211 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- tgw_security

<a id="canonical-0023100022300302-1300301103130303-3021020032020322-1333132001230213-2103322223003320-2033221013203121-1133002230031233-3102111313311311"></a>

Type: `"single"`. Computed.

Security Configuration for transit gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-east_west_service_policy_choice": "[\"active_east_west_service_policies\",\"east_west_service_policy_allow_all\",\"no_east_west_policy\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]"
}
```

<a id="canonical-0132013302113211-2223321302213033-0300222030301021-0002211211002100-0102320222211120-3101110210211220-0331330321220012-0320322132201331"></a>

## Direct properties — tgw_security / 230302200211 / 3

- [active_east_west_service_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-1223111221121130-0203133113223223-0102132323221213-3231020102031321-3223203211300330-2321100312311320-1222002221033331-0311332230213302): complete subsection reference.

- [active_enhanced_firewall_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-3121012103112123-2223131323010000-2321213100110321-1233031122122303-1003133030001021-2302120313122322-3112033332302000-2113011220102221): complete subsection reference.

- [active_forward_proxy_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-0200013210003100-2111311022132200-0332003130013133-2222310131011233-0212022222023211-2011330121301302-0021202320301030-1120200003133031): complete subsection reference.

- [active_network_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-2130222323310320-2130103102203233-3210213102330312-1113112011302122-3313132312302301-1310322312030322-0332023123122000-2330023121102212): complete subsection reference.

- [east_west_service_policy_allow_all](data-sources--aws_tgw_site--reference--group-002.md#canonical-2111121131102303-1012030112201202-0222131300232320-2032122121023311-0103030113032032-1313101030213320-0300300013311000-3123000002313320): complete subsection reference.

- [forward_proxy_allow_all](data-sources--aws_tgw_site--reference--group-002.md#canonical-0200331020201031-3023221303113223-1202102322123222-0130100010331231-1200303303111211-0110321123133220-3123003002033012-1003002030310003): complete subsection reference.

- [no_east_west_policy](data-sources--aws_tgw_site--reference--group-002.md#canonical-0301100121103102-3121203222311002-1103212021001311-0102032132330311-0300303133231323-0313213123003212-0230002223210331-0211122000111303): complete subsection reference.

- [no_forward_proxy](data-sources--aws_tgw_site--reference--group-002.md#canonical-1301103203102322-0133023022222333-3010101323211220-2222002221021132-2213132332313320-1022312003231022-2300100120031033-2201010011122332): complete subsection reference.

- [no_network_policy](data-sources--aws_tgw_site--reference--group-002.md#canonical-3023032101103211-1222102132121230-0311102331000023-3123210102330121-2232212002002033-0223213203212313-1100003110201222-2220221203033003): complete subsection reference.

<a id="canonical-3033212110231030-2013321301200303-3221130013100022-3123231231233322-2130231113103132-0121101112231132-3100101112322022-0120200003103213"></a>

## Next pages — tgw_security / 230302200211 / 4

- [tgw_security.active_east_west_service_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-1223111221121130-0203133113223223-0102132323221213-3231020102031321-3223203211300330-2321100312311320-1222002221033331-0311332230213302)
- [tgw_security.active_enhanced_firewall_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-3121012103112123-2223131323010000-2321213100110321-1233031122122303-1003133030001021-2302120313122322-3112033332302000-2113011220102221)
- [tgw_security.active_forward_proxy_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-0200013210003100-2111311022132200-0332003130013133-2222310131011233-0212022222023211-2011330121301302-0021202320301030-1120200003133031)
- [tgw_security.active_network_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-2130222323310320-2130103102203233-3210213102330312-1113112011302122-3313132312302301-1310322312030322-0332023123122000-2330023121102212)
- [tgw_security.east_west_service_policy_allow_all](data-sources--aws_tgw_site--reference--group-002.md#canonical-2111121131102303-1012030112201202-0222131300232320-2032122121023311-0103030113032032-1313101030213320-0300300013311000-3123000002313320)
- [tgw_security.forward_proxy_allow_all](data-sources--aws_tgw_site--reference--group-002.md#canonical-0200331020201031-3023221303113223-1202102322123222-0130100010331231-1200303303111211-0110321123133220-3123003002033012-1003002030310003)
- [tgw_security.no_east_west_policy](data-sources--aws_tgw_site--reference--group-002.md#canonical-0301100121103102-3121203222311002-1103212021001311-0102032132330311-0300303133231323-0313213123003212-0230002223210331-0211122000111303)
- [tgw_security.no_forward_proxy](data-sources--aws_tgw_site--reference--group-002.md#canonical-1301103203102322-0133023022222333-3010101323211220-2222002221021132-2213132332313320-1022312003231022-2300100120031033-2201010011122332)
- [tgw_security.no_network_policy](data-sources--aws_tgw_site--reference--group-002.md#canonical-3023032101103211-1222102132121230-0311102331000023-3123210102330121-2232212002002033-0223213203212313-1100003110201222-2220221203033003)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1223111221121130-0203133113223223-0102132323221213-3231020102031321-3223203211300330-2321100312311320-1222002221033331-0311332230213302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322332030213221-3333020312102122-1212200110201130-1033113200301120-1102001301313313-1330320132122221-3320001033323333-2213023022220222"></a>

## tgw_security.active_east_west_service_policies — active_east_west_service_policies / 201101223223 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.active_east_west_service_policies

<a id="canonical-1330121210213301-1310222323032202-3323310311102020-2200111100020032-0130333132111101-2231032131121030-1310201332001110-2121012121032111"></a>

Type: `"single"`. Computed.

Active service policies for the east-west proxy.

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

<a id="canonical-0220312203012133-1102131110132321-0231020131231303-2322330111302313-0222002210310010-0310202322301013-1130012223323102-0331222111323123"></a>

## Direct properties — active_east_west_service_policies / 201101223223 / 3

- [service_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-0210130022232311-1230113111213300-2133210300222022-1322202333000210-3121311122301110-0332200033230223-1132212012321002-0022313111322122): complete subsection reference.

<a id="canonical-0032122322211222-0102012221001202-0101302103111110-2031110211131223-0100212033332113-1122323321302312-0013333201310312-1113222110210220"></a>

## Next pages — active_east_west_service_policies / 201101223223 / 4

- [tgw_security.active_east_west_service_policies.service_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-0210130022232311-1230113111213300-2133210300222022-1322202333000210-3121311122301110-0332200033230223-1132212012321002-0022313111322122)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0210130022232311-1230113111213300-2133210300222022-1322202333000210-3121311122301110-0332200033230223-1132212012321002-0022313111322122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010032102032032-3102230212303231-2230201332003303-2200023221003323-0023013331003130-1220020031131201-1220303111000303-1330310113023213"></a>

## tgw_security.active_east_west_service_policies.service_policies — service_policies / 222121220110 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [tgw_security.active_east_west_service_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-1223111221121130-0203133113223223-0102132323221213-3231020102031321-3223203211300330-2321100312311320-1222002221033331-0311332230213302)
- tgw_security.active_east_west_service_policies.service_policies

<a id="canonical-0023300202230132-1213113002112211-1202311310233132-0312223330202013-3230303223120001-3000130201211111-1130012002321020-3102311111032133"></a>

Type: `"list"`. Computed.

List of references to service\_policy objects.

Upstream description:

A list of references to service\_policy objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3231123231320333-2100001312032100-3102233122301321-1313303023333232-1103033203010201-2231302123020223-2023331320310121-0111131113232112"></a>

## Direct properties — service_policies / 222121220110 / 3

<a id="canonical-2300210223323132-3313002130320201-1030201002213231-0303301201001001-3111300003120320-3020131111300101-0221220122100103-0232320210200113"></a>

<a id="canonical-0223212322001133-1003030222222001-3130331303233032-0322101003000110-2003131103022311-1113102223121200-1311232132022301-2302101222022330"></a>

## name property — service_policies / 222121220110 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0130222230131121-1203201023221202-0010222302100210-2131231312003300-2332303113222313-1312020213310011-1031110113211013-1133032222311030"></a>

<a id="canonical-2013220132102313-2032130313220122-3301020000133330-3021301020313013-2302002331311011-2122133301300013-3032122303132231-0210130122012313"></a>

## namespace property — service_policies / 222121220110 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0022003203333222-2210130213112320-1300320303103010-2021203111322003-0300211231203121-0011012331130010-2122003210222003-1113310111023330"></a>

<a id="canonical-2012321122100031-2120301223212303-3213110211013111-2120302232031313-0222120321202312-2322232211232002-0333030301122030-2320033333000333"></a>

## tenant property — service_policies / 222121220110 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3320002000210202-1020310111200123-2302222321231312-2122013022003123-0020122201200223-3210003303231203-0100321331100121-3010111013112133"></a>

## Next pages — service_policies / 222121220110 / 7

- [tgw_security.active_east_west_service_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-1223111221121130-0203133113223223-0102132323221213-3231020102031321-3223203211300330-2321100312311320-1222002221033331-0311332230213302)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3121012103112123-2223131323010000-2321213100110321-1233031122122303-1003133030001021-2302120313122322-3112033332302000-2113011220102221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230231232031023-2102300130332001-1113213331103203-2103013300030322-0311011022212102-0333231001003123-2313320011131112-3002303303011230"></a>

## tgw_security.active_enhanced_firewall_policies — active_enhanced_firewall_policies / 132333320033 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.active_enhanced_firewall_policies

<a id="canonical-3100122300123201-0233210101101332-2032133310332201-0300320223332330-3200233303333021-1030110010222113-0321103313321202-3330202320330210"></a>

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

<a id="canonical-0332113212222203-1311333122012301-2130122302013323-3321230121001011-0010200000102011-3331232110211001-3311223032112022-3021130300102111"></a>

## Direct properties — active_enhanced_firewall_policies / 132333320033 / 3

- [enhanced_firewall_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-2200310110303210-3110231030130313-3130133122320223-0212101230223100-3221100303233120-1112011311313313-3101302210101311-0011031230101131): complete subsection reference.

<a id="canonical-2223222202221110-2111132230202231-3123023103332331-1320101101023021-0220313230313202-3010121011122311-3000230301101212-1102221102320010"></a>

## Next pages — active_enhanced_firewall_policies / 132333320033 / 4

- [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-2200310110303210-3110231030130313-3130133122320223-0212101230223100-3221100303233120-1112011311313313-3101302210101311-0011031230101131)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2200310110303210-3110231030130313-3130133122320223-0212101230223100-3221100303233120-1112011311313313-3101302210101311-0011031230101131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121200130332033-2111323301220202-2313012011210102-2110201323331322-2313103133123221-2103323302331102-1200010211102230-0211123223031003"></a>

## tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies — enhanced_firewall_policies / 230012231222 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [tgw_security.active_enhanced_firewall_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-3121012103112123-2223131323010000-2321213100110321-1233031122122303-1003133030001021-2302120313122322-3112033332302000-2113011220102221)
- tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-3311213201332131-3213123121330331-1300121311023210-1332223320132301-0222323200222112-0111220013320122-3023300131232031-1000110333232202"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0121311101220203-0032311332320210-0202310122223200-0222023112200222-1331031233001210-2131021330203000-2022311201120223-3201100001323200"></a>

## Direct properties — enhanced_firewall_policies / 230012231222 / 3

<a id="canonical-3010323312003203-1230332212330132-3203200313101303-1013321303231302-1303122010103021-1131201010202322-0330311022312110-0010113221001020"></a>

<a id="canonical-2303200331031011-1201122310311302-2001230030120201-0123310220310032-0002301000110312-2201002301322102-1331332112010102-0010311011031011"></a>

## name property — enhanced_firewall_policies / 230012231222 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0021201300132102-1033001333301030-2100222303012100-2012300120100133-3000300111331202-0212112023323221-2000001110213232-2130121112323313"></a>

<a id="canonical-2133321330220302-1320103323313323-2010302001220331-0023311003132221-3102202133231031-2200323032201132-0200321313203212-2120322202321222"></a>

## namespace property — enhanced_firewall_policies / 230012231222 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2331103130010323-3130111132200330-1220112112021230-1313101012330230-0322123320300202-3322200302213133-0212023012032010-2110201031210310"></a>

<a id="canonical-3333230030332032-2030130303123300-3231003120013323-3331002331222322-1100223232310123-0310032022311231-3012000222303330-2202033301233303"></a>

## tenant property — enhanced_firewall_policies / 230012231222 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2131030231131122-3222033211330100-2121102100101213-2012211201223131-0313300313311331-2132133202032233-3211030003022022-2133301002231332"></a>

## Next pages — enhanced_firewall_policies / 230012231222 / 7

- [tgw_security.active_enhanced_firewall_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-3121012103112123-2223131323010000-2321213100110321-1233031122122303-1003133030001021-2302120313122322-3112033332302000-2113011220102221)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0200013210003100-2111311022132200-0332003130013133-2222310131011233-0212022222023211-2011330121301302-0021202320301030-1120200003133031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110012113111120-0132211331320201-0121230101031032-2121013300233233-1032323211230012-0031212200333310-1231131100021312-1230000210110123"></a>

## tgw_security.active_forward_proxy_policies — active_forward_proxy_policies / 102120221021 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.active_forward_proxy_policies

<a id="canonical-0033102321332031-2332102301121223-0313232211023310-1333321222213222-2323210303021100-1232122101112302-2000133233122212-0213120201022001"></a>

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

<a id="canonical-3220231330300130-3013333132213013-0021121222203332-1200311210131222-2212202303230202-1130302213320331-2230020210002122-2011012211230220"></a>

## Direct properties — active_forward_proxy_policies / 102120221021 / 3

- [forward_proxy_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-3233302333102031-2222230103310213-1022332121222030-2322212132002211-3110113130230121-0003333232013100-1320213103220020-0120103123210220): complete subsection reference.

<a id="canonical-2300211023331221-1200221131310033-1231231332012303-0020113221333010-3312013003212033-3220300332213031-3330303223310121-3220333030230001"></a>

## Next pages — active_forward_proxy_policies / 102120221021 / 4

- [tgw_security.active_forward_proxy_policies.forward_proxy_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-3233302333102031-2222230103310213-1022332121222030-2322212132002211-3110113130230121-0003333232013100-1320213103220020-0120103123210220)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3233302333102031-2222230103310213-1022332121222030-2322212132002211-3110113130230121-0003333232013100-1320213103220020-0120103123210220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123303333320232-0112010200101221-1113231213112002-0322221223233123-2021321012231302-0313120301031321-1230213223332133-1103002211131232"></a>

## tgw_security.active_forward_proxy_policies.forward_proxy_policies — forward_proxy_policies / 030121131120 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [tgw_security.active_forward_proxy_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-0200013210003100-2111311022132200-0332003130013133-2222310131011233-0212022222023211-2011330121301302-0021202320301030-1120200003133031)
- tgw_security.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-2112321233233312-1331021033300013-0121102131110002-0002213231223000-2102111010210002-2210100012200132-0000132223211122-3311302100310103"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2011233130333102-1210212130230321-2110033001003302-1120013112020032-3112222101113001-2321122331201032-2000012123220002-0230020102132030"></a>

## Direct properties — forward_proxy_policies / 030121131120 / 3

<a id="canonical-0101232113031113-3303103321202312-3321311030230001-0210302133011113-3320003223332311-2021222320333133-0201200132023312-2013212133012211"></a>

<a id="canonical-3230313111330020-2331101133003130-1122301323033133-0330303332333103-1023023000122121-2310221110000010-1110121021320321-1120223302220123"></a>

## name property — forward_proxy_policies / 030121131120 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2320121310321232-1111130131203130-1122301313230332-0000310233130122-0230110132321011-0132003031322313-3012313021023200-1130001130313331"></a>

<a id="canonical-0111211311001203-0311101330211213-1003032112103231-3213330000331013-2312133112311320-1211011300301321-3320221210030011-0303232032322310"></a>

## namespace property — forward_proxy_policies / 030121131120 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2022002112133322-0001323210012012-2122131022023033-2030323322033323-3230220030202231-3131031133111101-2131031212213311-1231303211030322"></a>

<a id="canonical-3012311021303332-3220210230000131-1220220010101201-1333100320002212-0111012032102022-0221020100021002-2013110031010233-0123031123000112"></a>

## tenant property — forward_proxy_policies / 030121131120 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0230132102331223-2331133313002221-3120201223010123-1000333211020002-2332102131121132-3103233030013220-3033112331102312-2101123233230132"></a>

## Next pages — forward_proxy_policies / 030121131120 / 7

- [tgw_security.active_forward_proxy_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-0200013210003100-2111311022132200-0332003130013133-2222310131011233-0212022222023211-2011330121301302-0021202320301030-1120200003133031)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2130222323310320-2130103102203233-3210213102330312-1113112011302122-3313132312302301-1310322312030322-0332023123122000-2330023121102212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221222331030210-0321123032023133-3310211201021010-1320312231101121-1312121032003130-2011023100313130-2111120120310033-0113233023331101"></a>

## tgw_security.active_network_policies — active_network_policies / 333303230123 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.active_network_policies

<a id="canonical-1222303002231310-2201021313212000-2230020330321333-0103103301323010-1311200333300131-2003333010222302-1233201021213103-0013202220231320"></a>

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

<a id="canonical-2303013102023111-2030200122210201-0002321220102021-3322201333033223-0113222212010233-3323201202300133-2212202201023210-2210303110022231"></a>

## Direct properties — active_network_policies / 333303230123 / 3

- [network_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-2330203012312330-0323122102031213-0011320200200001-3310301012103122-3021102202001312-3020333312013001-0301102100311301-1021010130300033): complete subsection reference.

<a id="canonical-2330111010333112-1020130221122113-0100333230130230-0033030212002003-1210022202210032-0301103120231300-3301130332321201-1123002032102113"></a>

## Next pages — active_network_policies / 333303230123 / 4

- [tgw_security.active_network_policies.network_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-2330203012312330-0323122102031213-0011320200200001-3310301012103122-3021102202001312-3020333312013001-0301102100311301-1021010130300033)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2330203012312330-0323122102031213-0011320200200001-3310301012103122-3021102202001312-3020333312013001-0301102100311301-1021010130300033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212120230121310-2111222330313201-2222201203003023-2232220320100100-0332100300213312-2133003001223003-3220223132022301-3203113101130013"></a>

## tgw_security.active_network_policies.network_policies — network_policies / 021203111310 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [tgw_security.active_network_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-2130222323310320-2130103102203233-3210213102330312-1113112011302122-3313132312302301-1310322312030322-0332023123122000-2330023121102212)
- tgw_security.active_network_policies.network_policies

<a id="canonical-1121302101300001-0202001231021112-1133232032322033-2212222312131313-0013221133130332-1202112133200120-0110000130321123-2213121112123313"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1133320132132003-3131220232223210-3111302231332002-3000302311111333-3322113111013001-0120320222012112-0301113010211111-0130320003123220"></a>

## Direct properties — network_policies / 021203111310 / 3

<a id="canonical-2222103201031010-2021112100232323-2312023031220012-0012202330003132-3011313120333133-1311231010020103-3023303100100332-2033232032313113"></a>

<a id="canonical-2302023122020312-0223330300213200-2332102022231200-0110111222322111-0312231302231221-1001303321012032-1332232200113231-2312331111031201"></a>

## name property — network_policies / 021203111310 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2130022010310232-0013213301223112-1333023200203321-3300030300102313-3331312033233303-1310202032231221-2311113013201133-0021310010322203"></a>

<a id="canonical-2002310133000332-3003100222331330-1333101132033221-2233322113333311-0111022211311203-3311132312222232-0200110333011020-1011122100003222"></a>

## namespace property — network_policies / 021203111310 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0000101223311230-3112013221302121-3003231312211322-2010132220131122-2103100220123033-1130020230102113-2120333013213320-2120232003320200"></a>

<a id="canonical-1202131300122020-0101123220132001-0023331312200330-3332012313121131-2020210331311320-1212211313133010-2221223022301113-1211221130011002"></a>

## tenant property — network_policies / 021203111310 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1321210303012221-3322001111312330-3233110330222322-0323333131003201-0321322010003200-2301122130103102-1121131233331031-3202133330021210"></a>

## Next pages — network_policies / 021203111310 / 7

- [tgw_security.active_network_policies](data-sources--aws_tgw_site--reference--group-002.md#canonical-2130222323310320-2130103102203233-3210213102330312-1113112011302122-3313132312302301-1310322312030322-0332023123122000-2330023121102212)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2111121131102303-1012030112201202-0222131300232320-2032122121023311-0103030113032032-1313101030213320-0300300013311000-3123000002313320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012210120313133-1130212023120222-3030022321220101-1212012000221110-0331032221020212-0103303333022010-2031232323331210-2123030012033120"></a>

## tgw_security.east_west_service_policy_allow_all — east_west_service_policy_allow_all / 000111022022 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.east_west_service_policy_allow_all

<a id="canonical-3331033230110031-0230010010022130-3101101003311012-0111122010311331-2232323013303200-3321223132332001-1302302302000103-2002310323013031"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for east west service policy allow all.

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

<a id="canonical-2111223000312330-3112021310211102-2332323211122333-1000003332110133-1100033220212023-2302311032021121-1031033032332000-1312301211022021"></a>

## Direct properties — east_west_service_policy_allow_all / 000111022022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3110103032011221-1311301203203120-1131321300123322-1313130101321230-0133001213012303-0330331101322222-0002111320033212-0110230330023300"></a>

## Next pages — east_west_service_policy_allow_all / 000111022022 / 4

- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0200331020201031-3023221303113223-1202102322123222-0130100010331231-1200303303111211-0110321123133220-3123003002033012-1003002030310003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102102231122120-0223300032220312-1100002300331003-3302310231303030-1212212311303100-0010030323020113-0131001022303223-2220103101210110"></a>

## tgw_security.forward_proxy_allow_all — forward_proxy_allow_all / 102021301203 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.forward_proxy_allow_all

<a id="canonical-0112211023013312-3330223220001323-0232131133032102-2222132133003102-2103033103232202-3003213132210012-0120320112213113-3033320133131220"></a>

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

<a id="canonical-1020130202011122-0330203210220113-0211200112003033-2112220230121013-1221112223121302-3102321212303221-2302032211000322-2201221332320221"></a>

## Direct properties — forward_proxy_allow_all / 102021301203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123301002000122-2132131123133303-0021331110220232-3202122020203211-1032112133210322-2300120203323123-0321212121132031-0313221231230001"></a>

## Next pages — forward_proxy_allow_all / 102021301203 / 4

- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0301100121103102-3121203222311002-1103212021001311-0102032132330311-0300303133231323-0313213123003212-0230002223210331-0211122000111303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110002230113003-1222220321100222-2021213102031233-2133001332203202-0103023312333101-0211121112123230-3231102022202312-2203332112321111"></a>

## tgw_security.no_east_west_policy — no_east_west_policy / 211031211213 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.no_east_west_policy

<a id="canonical-2331002212110332-2000100311121000-3311310302011001-1221132123002103-1301013220332310-2100320330222211-3132010312103123-1200002303212020"></a>

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

<a id="canonical-3121131131330331-2212221202100020-3301321131031223-0332220233211130-2221100212012210-2213200322222323-3001012300003330-3203013200313210"></a>

## Direct properties — no_east_west_policy / 211031211213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322331031210230-1311030122000113-0131130102011320-1131131332003001-1012131331020103-3333030230201302-2203201210132013-0121300010131022"></a>

## Next pages — no_east_west_policy / 211031211213 / 4

- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-1301103203102322-0133023022222333-3010101323211220-2222002221021132-2213132332313320-1022312003231022-2300100120031033-2201010011122332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002113032101223-3203012032221123-2313202321021020-1133110333200103-3322212321012330-3021303300010011-2320201333122221-2023022310223022"></a>

## tgw_security.no_forward_proxy — no_forward_proxy / 020020301210 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.no_forward_proxy

<a id="canonical-1023311231031313-0033101023212023-3332201011001102-1101033121233122-3233303123232221-0311331213021211-0210132002102121-2022313001210110"></a>

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

<a id="canonical-0210223020303131-3200201013133132-2120331321331320-0333101300012121-1102110211003203-3033100310001330-1330203230001213-1100301110221301"></a>

## Direct properties — no_forward_proxy / 020020301210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020033110023133-0102031202322023-3003311220332021-3130333213000120-0222023033003300-1100111230133131-1203321301222012-1230101031113220"></a>

## Next pages — no_forward_proxy / 020020301210 / 4

- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-3023032101103211-1222102132121230-0311102331000023-3123210102330121-2232212002002033-0223213203212313-1100003110201222-2220221203033003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200223211100023-1211210013211231-3130332330002000-1203133122120122-1031003333311100-1031313220020230-3312130033213210-0131222222022200"></a>

## tgw_security.no_network_policy — no_network_policy / 301221210303 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.no_network_policy

<a id="canonical-1033230301303120-1031212130330003-1132331221223312-0331311231330122-0322302032233132-1111331031123131-2122130311310110-0232313101113211"></a>

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

<a id="canonical-1102001221320331-0201323213221122-2113212301213013-0032210001000132-0330003132030201-0123001133301210-1310031210203120-1123023130120100"></a>

## Direct properties — no_network_policy / 301221210303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011303310322223-0312031131331000-2002033232003201-2120000330133122-0230110330113200-1132232233331032-2201101112111301-0222002132131222"></a>

## Next pages — no_network_policy / 301221210303 / 4

- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312303002330011-1022032301303131-3211201322120031-2020233303230133-1202333120000310-0203120123113113-0112222233032112-0331231123321020"></a>

## vn_config — vn_config / 220023202221 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- vn_config

<a id="canonical-1132130212133133-1300231313303102-3012311220221030-3002121133320010-3113121310300211-1221101102322223-2121132222023331-2213031113012122"></a>

Type: `"single"`. Computed.

Virtual Network Configuration. Virtual Network Configuration.

Upstream description:

Virtual Network Configuration.

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
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

<a id="canonical-2031032301103013-3233130122200202-3113301133113013-1301112310231320-0122313313003100-2202033020212003-3211320123233122-1232032312212012"></a>

## Direct properties — vn_config / 220023202221 / 3

- [allowed_vip_port](data-sources--aws_tgw_site--reference--group-002.md#canonical-2312123133020312-3232021220010232-1132201301121131-0020101111100121-2103133022200200-1221020013101212-1133122300113202-3303200333320110): complete subsection reference.

- [allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-1320122013130233-1233313130322210-1130202020333203-2131102020311033-1200013112110122-0302000210020203-3032022010322233-0302222222310113): complete subsection reference.

- [dc_cluster_group_inside_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-0123201232133002-3320121122011110-3332021222121331-3133012323330001-3313232012210123-1003323303123012-0333212223201112-3111312232100020): complete subsection reference.

- [dc_cluster_group_outside_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-3012010120110213-0221112133332320-0033033112122032-2103012332320020-1011323100121300-1113003111311133-3123003331123032-3020321300210202): complete subsection reference.

- [global_network_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-1102112012302300-3110320022302020-2101103231133230-2233301202100003-1121103303002200-2133212012023003-1310000112221333-1332021033111203): complete subsection reference.

- [inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320): complete subsection reference.

- [no_dc_cluster_group](data-sources--aws_tgw_site--reference--group-003.md#canonical-3322033323111233-2333300011113211-2303310012333330-1032323330033233-1133313111333130-1303303300213210-1222223033330233-0322021301022133): complete subsection reference.

- [no_global_network](data-sources--aws_tgw_site--reference--group-003.md#canonical-0211231303000231-1332322333303110-2231330333122021-3022013113310302-0331223221023310-0321000310301123-1112130032320132-0012101113120333): complete subsection reference.

- [no_inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-2133103000303110-2322132103102100-0111231000032223-1012202032122131-2132030320322112-2311233020023323-1010100332321223-0212021022322130): complete subsection reference.

- [no_outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-2220121110310322-3310320322232321-2200001213032233-3220332331310131-1101133323003022-1221011103131111-2032233010211203-1023100311031333): complete subsection reference.

- [outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301): complete subsection reference.

- [sm_connection_public_ip](data-sources--aws_tgw_site--reference--group-004.md#canonical-3310333131011322-0330100020123222-2030322120221110-2111221333212323-0130210132002022-1102312332302113-2031133120302121-0210233110231121): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--aws_tgw_site--reference--group-004.md#canonical-0330213031123210-3033120010312021-0023130021033220-2313022211120033-1111033320230122-0232222032233123-2302100330110023-3000000033133323): complete subsection reference.

<a id="canonical-3221103301021113-3031311303031222-2031231010203133-1103323231021011-1110232110121301-2331102203323210-0132032130313001-0230032122112233"></a>

## Next pages — vn_config / 220023202221 / 4

- [vn_config.allowed_vip_port](data-sources--aws_tgw_site--reference--group-002.md#canonical-2312123133020312-3232021220010232-1132201301121131-0020101111100121-2103133022200200-1221020013101212-1133122300113202-3303200333320110)
- [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-1320122013130233-1233313130322210-1130202020333203-2131102020311033-1200013112110122-0302000210020203-3032022010322233-0302222222310113)
- [vn_config.dc_cluster_group_inside_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-0123201232133002-3320121122011110-3332021222121331-3133012323330001-3313232012210123-1003323303123012-0333212223201112-3111312232100020)
- [vn_config.dc_cluster_group_outside_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-3012010120110213-0221112133332320-0033033112122032-2103012332320020-1011323100121300-1113003111311133-3123003331123032-3020321300210202)
- [vn_config.global_network_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-1102112012302300-3110320022302020-2101103231133230-2233301202100003-1121103303002200-2133212012023003-1310000112221333-1332021033111203)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320)
- [vn_config.no_dc_cluster_group](data-sources--aws_tgw_site--reference--group-003.md#canonical-3322033323111233-2333300011113211-2303310012333330-1032323330033233-1133313111333130-1303303300213210-1222223033330233-0322021301022133)
- [vn_config.no_global_network](data-sources--aws_tgw_site--reference--group-003.md#canonical-0211231303000231-1332322333303110-2231330333122021-3022013113310302-0331223221023310-0321000310301123-1112130032320132-0012101113120333)
- [vn_config.no_inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-2133103000303110-2322132103102100-0111231000032223-1012202032122131-2132030320322112-2311233020023323-1010100332321223-0212021022322130)
- [vn_config.no_outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-2220121110310322-3310320322232321-2200001213032233-3220332331310131-1101133323003022-1221011103131111-2032233010211203-1023100311031333)
- [vn_config.outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301)
- [vn_config.sm_connection_public_ip](data-sources--aws_tgw_site--reference--group-004.md#canonical-3310333131011322-0330100020123222-2030322120221110-2111221333212323-0130210132002022-1102312332302113-2031133120302121-0210233110231121)
- [vn_config.sm_connection_pvt_ip](data-sources--aws_tgw_site--reference--group-004.md#canonical-0330213031123210-3033120010312021-0023130021033220-2313022211120033-1111033320230122-0232222032233123-2302100330110023-3000000033133323)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)

<a id="canonical-2312123133020312-3232021220010232-1132201301121131-0020101111100121-2103133022200200-1221020013101212-1133122300113202-3303200333320110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
