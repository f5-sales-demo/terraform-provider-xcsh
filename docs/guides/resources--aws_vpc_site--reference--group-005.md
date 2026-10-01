---
page_title: "xcsh_aws_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site reference."
---

# xcsh_aws_vpc_site reference

<a id="canonical-2213133200200310-2032131322313120-2231122031000211-1303221113212003-2033213022221203-2201223212001112-1212030332232210-3230202000230122"></a>

## Direct properties — global_network_connections / 310010331231 / 3

- [sli_to_global_dr](resources--aws_vpc_site--reference--group-005.md#canonical-0310010333332202-0321210313233001-0112121111323111-0012131121122302-0212001301032231-1020312133132001-3130032023223332-3320020131330300): complete subsection reference.

- [slo_to_global_dr](resources--aws_vpc_site--reference--group-005.md#canonical-0033110331310310-2001202301213112-0023212201012132-0001300122001233-3103311030210222-0133230002231330-2222010032302010-0103013103223300): complete subsection reference.

<a id="canonical-2322000322132301-0322212101301021-3330212320031002-1313010110022132-3331300323002013-1211330111022220-2211133332313023-1333120323101101"></a>

## Next pages — global_network_connections / 310010331231 / 4

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--reference--group-005.md#canonical-0310010333332202-0321210313233001-0112121111323111-0012131121122302-0212001301032231-1020312133132001-3130032023223332-3320020131330300)
- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--reference--group-005.md#canonical-0033110331310310-2001202301213112-0023212201012132-0001300122001233-3103311030210222-0133230002231330-2222010032302010-0103013103223300)
- [voltstack_cluster.global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-2203021112020001-3133120221323131-3110312113100112-2200221102023130-1333113333312321-3332230222220203-3213111121221120-3032201310331111)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0310010333332202-0321210313233001-0112121111323111-0012131121122302-0212001301032231-1020312133132001-3130032023223332-3320020131330300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022000002102313-0223213222312312-2032010333031203-1121203132303000-2003223002032111-3002201120101130-3233230000110210-2032233223300010"></a>

## voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr — sli_to_global_dr / 133222010120 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-2203021112020001-3133120221323131-3110312113100112-2200221102023130-1333113333312321-3332230222220203-3213111121221120-3032201310331111)
- [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-2220100231202200-1000303320332323-0223113001021233-1013001212100200-0001230320000202-0110123212001021-2233111031310210-0023303011021311)
- voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-3103233002033212-3220130333020022-2302013102032221-3133320202121031-0103212130213211-2121331232032312-2320222123123113-2130201320000202"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-2123213003032100-1330320322201022-0233312200330321-0111031213213031-0101200201101302-1303201110112130-3002333231222313-0310221212123221"></a>

## Direct properties — sli_to_global_dr / 133222010120 / 3

- [global_vn](resources--aws_vpc_site--reference--group-005.md#canonical-0122123132212120-0302001002300303-3313133203033013-1223201302110213-3301101201133002-0202013231233002-0110201012210312-2102323203232011): complete subsection reference.

<a id="canonical-3000000320333023-0302030213201223-0011000100120001-2121231333312031-0312133013131100-3133103232033330-1030110320102313-2133012100222123"></a>

## Next pages — sli_to_global_dr / 133222010120 / 4

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--aws_vpc_site--reference--group-005.md#canonical-0122123132212120-0302001002300303-3313133203033013-1223201302110213-3301101201133002-0202013231233002-0110201012210312-2102323203232011)
- [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-2220100231202200-1000303320332323-0223113001021233-1013001212100200-0001230320000202-0110123212001021-2233111031310210-0023303011021311)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0122123132212120-0302001002300303-3313133203033013-1223201302110213-3301101201133002-0202013231233002-0110201012210312-2102323203232011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330113021001311-3333001112221122-0311133210113300-2100133121322212-2201231320331102-3031112230121303-0023222001203011-3320103330122323"></a>

## voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn — global_vn / 320210301220 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-2203021112020001-3133120221323131-3110312113100112-2200221102023130-1333113333312321-3332230222220203-3213111121221120-3032201310331111)
- [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-2220100231202200-1000303320332323-0223113001021233-1013001212100200-0001230320000202-0110123212001021-2233111031310210-0023303011021311)
- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--reference--group-005.md#canonical-0310010333332202-0321210313233001-0112121111323111-0012131121122302-0212001301032231-1020312133132001-3130032023223332-3320020131330300)
- voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-3310102320113003-0320200231202130-1103003101113301-3023001332123103-1003211020313101-0103311221303320-3300023032121000-0301013210220331"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222210012012110-0220133210023000-1133003130330111-1321233133132000-1030321000111203-2232231302202312-0231022033322230-3300122022003031"></a>

## Direct properties — global_vn / 320210301220 / 3

<a id="canonical-2201320031333101-0323301000011122-1111110013120300-2302212021200210-2013201033230110-1320123011223303-1331031133212121-0022200110213313"></a>

<a id="canonical-1132331130230203-2121210320210220-2032022121133033-2011210003313001-0102233220132100-0020120122121120-3022220230111021-1123200101011330"></a>

## name property — global_vn / 320210301220 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-3010100231033211-0033112022320232-3122020101311100-1310122230312322-3303300333321122-1112312311020023-3021312000033321-3302212113021312"></a>

<a id="canonical-3233200333322221-3120123010231330-0033010213213121-0333033220122321-2232303230220110-0002002110120333-2330311213231013-3123322020111033"></a>

## namespace property — global_vn / 320210301220 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-1300323210200231-2000033010322033-0203230220320030-1021313330230023-2301122001223221-2321300002001222-2221121310233120-2311302223313231"></a>

<a id="canonical-2113333131102003-3220022001103021-1302331213013010-3330222312020301-2203103123100131-0102002311320110-0322222022103233-1023221033212110"></a>

## tenant property — global_vn / 320210301220 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-3311222033131023-2022232113321130-0200320000302303-3102123100230231-3313121012321010-3303132011313123-0131230200311313-1131101112311222"></a>

## Next pages — global_vn / 320210301220 / 7

- [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--reference--group-005.md#canonical-0310010333332202-0321210313233001-0112121111323111-0012131121122302-0212001301032231-1020312133132001-3130032023223332-3320020131330300)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0033110331310310-2001202301213112-0023212201012132-0001300122001233-3103311030210222-0133230002231330-2222010032302010-0103013103223300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303332012312322-2002022013200220-0222001122332121-3211323233332233-0211301013030303-2103003023202112-0132113032012033-1233012332230303"></a>

## voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr — slo_to_global_dr / 322321120120 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-2203021112020001-3133120221323131-3110312113100112-2200221102023130-1333113333312321-3332230222220203-3213111121221120-3032201310331111)
- [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-2220100231202200-1000303320332323-0223113001021233-1013001212100200-0001230320000202-0110123212001021-2233111031310210-0023303011021311)
- voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-1021210001323230-0333331033200001-3101102200212011-2320031320102113-0031330133320013-1132230013203103-2020020200032013-3330122212023111"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332002201013102-3031123233100012-3330022301223313-0221133101320210-2131030110320012-2020303211330130-2110330221231303-0131321101002120"></a>

## Direct properties — slo_to_global_dr / 322321120120 / 3

- [global_vn](resources--aws_vpc_site--reference--group-005.md#canonical-3210200331100022-0031103221102003-0110213212102133-1323300221233023-1000310230031322-0120012330122111-0132000222112132-0031131331020030): complete subsection reference.

<a id="canonical-1011130302210013-3333313033002013-3100123013003300-2030311310013222-0313033022320203-3203023011132312-3301000110033002-1121303322313102"></a>

## Next pages — slo_to_global_dr / 322321120120 / 4

- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--aws_vpc_site--reference--group-005.md#canonical-3210200331100022-0031103221102003-0110213212102133-1323300221233023-1000310230031322-0120012330122111-0132000222112132-0031131331020030)
- [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-2220100231202200-1000303320332323-0223113001021233-1013001212100200-0001230320000202-0110123212001021-2233111031310210-0023303011021311)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3210200331100022-0031103221102003-0110213212102133-1323300221233023-1000310230031322-0120012330122111-0132000222112132-0031131331020030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200103202303100-3320210201013330-1211012132121311-0312122100230132-0133231301211332-2201213211203123-1213332201111200-2302022211122202"></a>

## voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn — global_vn / 110130113130 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-2203021112020001-3133120221323131-3110312113100112-2200221102023130-1333113333312321-3332230222220203-3213111121221120-3032201310331111)
- [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-2220100231202200-1000303320332323-0223113001021233-1013001212100200-0001230320000202-0110123212001021-2233111031310210-0023303011021311)
- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--reference--group-005.md#canonical-0033110331310310-2001202301213112-0023212201012132-0001300122001233-3103311030210222-0133230002231330-2222010032302010-0103013103223300)
- voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-0100120013312331-3303301321011232-1232031032303230-1200113033302200-2312233103133122-1222313102031323-1303302131022311-2222331032003222"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-3001000012231133-0020303221011322-3201122122202112-1113212111020131-0333111120303111-1312010032023032-2122121201211003-1332202031200330"></a>

## Direct properties — global_vn / 110130113130 / 3

<a id="canonical-3001023330221210-1331320010020032-2102321332033312-1033111130212013-2203112233331203-2030312210323121-0333111323312021-3110100022003012"></a>

<a id="canonical-1302213311313103-2330210302111210-0102032013311322-1231122121230312-3111221103113003-1131310213002203-0130321333213103-0300310031300201"></a>

## name property — global_vn / 110130113130 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-0010101331103001-3201023330003312-0100122333201200-1220022312111322-0110233101131213-2221133102332213-3020003221130202-2020022032100010"></a>

<a id="canonical-2203330031210123-1322322233131310-1103201000122101-2113131031022231-1312212001102110-1132230302332321-3313220033212332-0310220323131111"></a>

## namespace property — global_vn / 110130113130 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-0212301031230321-1003322123220133-3300303030023302-2123000120103000-2023100102301221-3002013213300101-0022332232332011-0120221033013120"></a>

<a id="canonical-0122103221301013-2031231202233032-2213132323221231-3003022300230202-3133330211331213-2020031021323231-3003202101102100-0121132033303322"></a>

## tenant property — global_vn / 110130113130 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-2031333100200112-3323330211203110-1313123132002322-1200330021210001-1220022121220132-2311010302301333-1101212311030312-1031320220320102"></a>

## Next pages — global_vn / 110130113130 / 7

- [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--reference--group-005.md#canonical-0033110331310310-2001202301213112-0023212201012132-0001300122001233-3103311030210222-0133230002231330-2222010032302010-0103013103223300)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1002223031333221-1312133120110331-2000030232301122-0320112303200023-0220312221030033-0212231012320032-1310132121121023-2031111103310323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003000110003113-2332132130033103-1320300023321231-3212200232210302-1001131210332201-1031032021200023-0010122103210303-3023131010132220"></a>

## voltstack_cluster.k8s_cluster — k8s_cluster / 321032320211 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.k8s_cluster

<a id="canonical-2123202230030310-0123012200333320-1031321230333201-3011110030313003-1302113123302221-3300132303102333-3013033123130313-0320013002203123"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
k8s_cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101123123100312-0233330100220132-2013112022000103-0322232110323223-2320211231110331-1212202031222233-3011030130311211-0222011111111010"></a>

## Direct properties — k8s_cluster / 321032320211 / 3

<a id="canonical-0201002032002311-2231031322322032-1010332201113031-1333120203331000-0202232021100011-0301313313131121-3210113222010113-0120123110023200"></a>

<a id="canonical-1020310333332323-0203030310021123-2311321203112021-1331332331133133-2003213232020111-1112131322023101-3300033210303000-3111010222333001"></a>

## name property — k8s_cluster / 321032320211 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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

<a id="canonical-0010300201200211-1212121011011312-0231031333110121-0221232133120323-0321311220111331-2110300022300222-3300112013003020-3001310322220230"></a>

<a id="canonical-3013031222030222-0303101312133033-2223033231102220-1212001122103032-3202123120233300-2202001111031223-3010013131033312-1232330120303113"></a>

## namespace property — k8s_cluster / 321032320211 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3111001302231333-3213313233011100-2122003201033123-1200211112010330-1033310221302102-2323012231201230-0220000211123130-1310111323223213"></a>

<a id="canonical-3102320001001001-2120020230210111-3112132102033001-0203003001301100-0123022310310102-2021013002203231-2022031120331100-3212030011133000"></a>

## tenant property — k8s_cluster / 321032320211 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-2210011023313013-2012223320123331-2230113133322122-2300001233222002-3313123302022130-3133312321220123-3221200133220012-1022130323013321"></a>

## Next pages — k8s_cluster / 321032320211 / 7

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1201130103133211-0130322100310130-1030002231330321-3012320113211113-3320121023220131-3320101131300223-1033121201321000-1102000310302133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310101213222103-1103310333231230-1110200300230120-3221123101022000-0323111300201330-1030131231102312-2302223322022110-2323201112231312"></a>

## voltstack_cluster.no_dc_cluster_group — no_dc_cluster_group / 130322133123 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.no_dc_cluster_group

<a id="canonical-0012000103221302-3212103230310311-2330220222213012-2121020331202232-3001013212031233-3111021333303221-1321113321132102-2222331302210033"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_dc_cluster_group = {}
```

<a id="canonical-2023031302111003-0010120011330102-2203120003200311-3300232330022300-3100311210201231-1032013010201333-0322213220213103-3032020230112030"></a>

## Direct properties — no_dc_cluster_group / 130322133123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030223131321033-1032022003100130-1202002212210100-0311223023330011-2311201012100322-3100202232333222-3332110003111030-2231032030221031"></a>

## Next pages — no_dc_cluster_group / 130322133123 / 4

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2313231002032112-2201321312013223-0112230330312033-2003333122232032-3133311113001101-0011131103220332-0023311001021021-2311133330312311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113022131322202-3320203300102112-1031121322213100-3133101323312202-0202332013201231-3203331222201122-1101303030210100-0003030202011122"></a>

## voltstack_cluster.no_forward_proxy — no_forward_proxy / 122133310011 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.no_forward_proxy

<a id="canonical-0213132213020101-1033031001202033-1230320011210002-1101022103113013-0001102122221201-1302321021003203-3112333223133200-2102220233012101"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_forward_proxy = {}
```

<a id="canonical-0212200021012001-0333200131133002-0130100001303022-1331013333101230-3123111001112230-3123310022330201-1103330232322023-0212212033221223"></a>

## Direct properties — no_forward_proxy / 122133310011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321130323013120-2101301200132022-0022223232211131-3012332102102022-3311101120121001-2322330111203312-2000003331032112-0113012201202300"></a>

## Next pages — no_forward_proxy / 122133310011 / 4

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2211203113330213-2300232001010302-0011302301022233-2311112202113222-0013300032330123-3003323320102322-0211311232102023-2101230102121220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202303323110200-2013112011320200-0002030313313101-1331102303213310-2301023130122123-0332210220203020-2232211322120321-3223302310221202"></a>

## voltstack_cluster.no_global_network — no_global_network / 011030212021 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.no_global_network

<a id="canonical-1203022323003133-1102132303213003-1032022321301002-2030023213113032-1332320012330003-2330333301123132-3311122121021010-1112003130133332"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no global network.

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

Terraform syntax:

```terraform
no_global_network = {}
```

<a id="canonical-0331001213122002-0223020232121132-2330123013021333-3010033020323031-2010022210320112-0133013333330310-3010331011332221-3331320013302101"></a>

## Direct properties — no_global_network / 011030212021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232132021320302-2123201010310131-3100012303100023-2310320021113322-0122331121023201-3301301221001030-0310010301013020-2300003000320013"></a>

## Next pages — no_global_network / 011030212021 / 4

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2211323200131213-1211213333312303-1102030302103231-2031030230212212-0333102113012100-2211131133011220-2333203202301123-0200302011331231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120233030311230-1313210233330102-2121033010210220-0020000220221120-3322222013032131-3211332013102111-0211301212311221-2001100111002131"></a>

## voltstack_cluster.no_k8s_cluster — no_k8s_cluster / 032310101302 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.no_k8s_cluster

<a id="canonical-3230112232022112-2322012221221023-0100100210310002-3222233302100133-3321232111021103-1000302302032220-2231321211231032-3332312112323002"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_k8s_cluster = {}
```

<a id="canonical-2311002323232032-1111100123033111-3113122232313330-0130231013033221-0122120213233212-1331022332010000-2003220121103112-0020122013123101"></a>

## Direct properties — no_k8s_cluster / 032310101302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223220310012301-3233133203303102-0002103220333112-0211103002111013-3123102211320232-1211110112112231-0302230003001321-0133223320300033"></a>

## Next pages — no_k8s_cluster / 032310101302 / 4

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3003302203010233-2003212201230112-1311312031311112-3222121001020130-1000333010121301-1003303232131221-3201010031333013-2032210121203021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010333210200132-3033003133312201-3303132223121331-2330332322121232-3313031121211223-3231103010313232-1120030322021101-3120112233013033"></a>

## voltstack_cluster.no_network_policy — no_network_policy / 213202101130 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.no_network_policy

<a id="canonical-1201331233331310-1013313333302311-2213331000000230-1232323003022202-2200032231322221-2011032321031310-1201303202023231-1020111302013202"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_network_policy = {}
```

<a id="canonical-3001300120200032-0021110331100031-0023122033232222-0132303301011103-2201110332101231-0023211322132003-0203301321102100-2000031223001132"></a>

## Direct properties — no_network_policy / 213202101130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212113122023130-1130021323331123-2102000322013021-0130202113032233-1110013122022023-2011101221323000-1300310030301221-0333030230111002"></a>

## Next pages — no_network_policy / 213202101130 / 4

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1310102322003113-1232131320210012-1220231032232010-1320213211120222-3133213231120020-2233132133313031-3300233101003110-2130330303133321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003011133123000-0331033012203323-1202321212322230-0303211131233221-1011032221000330-0211321132230311-1302322202223133-3021002103223101"></a>

## voltstack_cluster.no_outside_static_routes — no_outside_static_routes / 211331021030 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.no_outside_static_routes

<a id="canonical-3301220100020221-0310000332113223-1000211331311223-1233010230321200-0310010223230122-1310333101121002-0011013210220323-3031023323101021"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no outside static routes.

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

Terraform syntax:

```terraform
no_outside_static_routes = {}
```

<a id="canonical-2321130033123331-2132332030013310-0110132210103232-2003211212232222-1230110200132233-3221201100330310-1332101110211312-0311303013121301"></a>

## Direct properties — no_outside_static_routes / 211331021030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010102323002311-2230301002210011-1333330112130221-3030123333210231-1102313023231033-0003122023223113-1133032030311113-3113013103211221"></a>

## Next pages — no_outside_static_routes / 211331021030 / 4

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103111021011102-3332102102202100-0102101220110311-2031312132320102-0110202232232112-3221232303231031-3103023032110222-0232013133310220"></a>

## voltstack_cluster.outside_static_routes — outside_static_routes / 211300311003 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.outside_static_routes

<a id="canonical-3211321313320003-1033012003231211-1123003010011310-3300012121203221-3230132320220010-1310300312022331-2123311220121233-0112120010320202"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside static routes.

Upstream description:

List of static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_route_list")}
```

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

Terraform syntax:

```terraform
outside_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2030102213113331-3313201123133323-0120303233201130-2102321032123122-3330101213020231-0030231120101012-3121111213110100-2000010001101312"></a>

## Direct properties — outside_static_routes / 211300311003 / 3

- [static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320): complete subsection reference.

<a id="canonical-1210201120331110-2202030102003101-0021311000200033-2002331110320101-1113313010213322-0021021333302113-0110222201023203-2221003022320213"></a>

## Next pages — outside_static_routes / 211300311003 / 4

- [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120321002332023-2001022221110200-2002321310332222-2011101031010331-3311122021323320-2031111330110320-3131330321102301-0133010023233113"></a>

## voltstack_cluster.outside_static_routes.static_route_list — static_route_list / 001121200133 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211)
- voltstack_cluster.outside_static_routes.static_route_list

<a id="canonical-2202311120130012-0200203031312311-3322113232111202-2312133323331232-1122023330233233-1333022113100110-3221213303310133-3031002130131131"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_static_route",
    "simple_static_route")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
static_route_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000320203020312-3023013022021310-1010012221033103-3223310031003110-1112012123222120-2321232302202032-0000022301310000-2103031330132003"></a>

## Direct properties — static_route_list / 001121200133 / 3

- [custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033): complete subsection reference.

<a id="canonical-0213022122200132-2023132032130110-2332313100023102-3103330211232122-3022231303102323-0103013213302322-2113002332331120-2222311113312132"></a>

<a id="canonical-1031321300120301-2300113013303103-1223100131122121-1310312320133233-3330302102202323-3122011031232020-0130100033010301-2333101222302203"></a>

## simple_static_route property — static_route_list / 001121200133 / 4

Type: `"string"`. Optional.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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

<a id="canonical-3313013123331101-0013210102333033-0203022011320202-3122312221111233-3013111333323322-2220330121110230-2321102100303011-1303331312022230"></a>

## Next pages — static_route_list / 001121200133 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101021203113333-3221211203223300-0002201103223321-3223212033101311-0131030333213110-3010311110021331-1330303120213301-2301201221131322"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route — custom_static_route / 001012020030 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-2202131110133103-3313033200133300-2201120022232313-2030202000103121-2101110121031112-2111133232003312-0313223102133201-0111022003131331"></a>

Type: `"object"`. single nested block, Optional.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnets")}
```

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

Terraform syntax:

```terraform
custom_static_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200300323201110-0231010002332332-1003202110022031-3130010330302001-3011232101303021-0223310333221331-0310133230132132-3111131223313123"></a>

## Direct properties — custom_static_route / 001012020030 / 3

<a id="canonical-1103003000210321-3010233320323133-3301000113001122-3123201112230330-1012100002122023-0002031332022200-0010012121002120-3123122313223112"></a>

<a id="canonical-2313012103200230-1003203332101302-0113221303123222-1010030130020212-3201110300000111-2101002020321031-2321031323121120-0212022202110310"></a>

## attrs property — custom_static_route / 001012020030 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](resources--aws_vpc_site--reference--group-005.md#canonical-3100133022001303-1033112230323203-3032132113220002-2130303030210001-0233133210032320-1331230231030200-0122202311212110-3320112002121023): complete subsection reference.

- [nexthop](resources--aws_vpc_site--reference--group-005.md#canonical-1003022201011203-3010322212000002-1202203120131010-2233232232031000-1323321011221123-0010311020110322-3210032321002232-3011120323113110): complete subsection reference.

- [subnets](resources--aws_vpc_site--reference--group-005.md#canonical-0203332230232223-0203331231222032-1323312021032100-1301001321301033-1231112012010132-0010211300210001-1001131223113330-3131232122211023): complete subsection reference.

<a id="canonical-2233332331222123-3313203032020010-0303313020213333-1303122102111013-3330221310330321-2130200312300123-0000301200231213-0210230213031321"></a>

## Next pages — custom_static_route / 001012020030 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](resources--aws_vpc_site--reference--group-005.md#canonical-3100133022001303-1033112230323203-3032132113220002-2130303030210001-0233133210032320-1331230231030200-0122202311212110-3320112002121023)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-005.md#canonical-1003022201011203-3010322212000002-1202203120131010-2233232232031000-1323321011221123-0010311020110322-3210032321002232-3011120323113110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-005.md#canonical-0203332230232223-0203331231222032-1323312021032100-1301001321301033-1231112012010132-0010211300210001-1001131223113330-3131232122211023)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3100133022001303-1033112230323203-3032132113220002-2130303030210001-0233133210032320-1331230231030200-0122202311212110-3320112002121023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230120212323233-2113223003112123-0222022321033012-2023132223330311-1303120201201101-3130101312331322-2121303130111113-2212100220232131"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels — labels / 320110012212 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-3022202212133321-0002022300202111-3112100230112300-0023203330323303-1212202321232012-2032031230231120-1012200321110313-0202121310202302"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Static Route, these labels can be used in network policy.

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

Terraform syntax:

```terraform
labels {}
```

<a id="canonical-0123031333320122-3231201022313011-3221110003211123-1212113100031031-1030232330102311-1031231000101010-1103221012131031-2330300112332133"></a>

## Direct properties — labels / 320110012212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202321032213322-3201113113110003-1112203300332020-2013130121102031-3133002130113022-2230130212113301-3311133311223110-3133131030110130"></a>

## Next pages — labels / 320110012212 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1003022201011203-3010322212000002-1202203120131010-2233232232031000-1323321011221123-0010311020110322-3210032321002232-3011120323113110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001221033331113-1321012212013303-2332000012032222-1232113210310312-3303333132010233-0011222230101032-2111032022312103-0031122332020332"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 100013001020 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-0010020131221232-0220310222122012-1123012121123022-0102101100302200-2132113003103312-2221220120031023-1013130001111221-2223003330101003"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

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

Terraform syntax:

```terraform
nexthop {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001101132331201-1130111132100021-2010111121001013-2122302103301331-2031330000330001-0312000130332310-3102102122121022-2032222132131011"></a>

## Direct properties — nexthop / 100013001020 / 3

- [interface](resources--aws_vpc_site--reference--group-005.md#canonical-2232223330003211-2000210011021301-2010313122210120-3013022212322020-1330111223031311-0030231113111320-3312202013321232-0131131220121010): complete subsection reference.

- [nexthop_address](resources--aws_vpc_site--reference--group-005.md#canonical-1122313131110100-1133033022213222-1221213123112011-0332120120301120-2000112121220222-2130301232011222-2133330210000322-3232033032010100): complete subsection reference.

<a id="canonical-0002103221323031-2031121012132200-0132211210330030-1313223213131320-2232011121010011-3300132231022130-0212120300211301-0303003210200032"></a>

<a id="canonical-0201320323130022-3012320103121021-2313030331222023-0113211122123120-3133220303220121-3302302102113311-2223100201113201-2120331210311011"></a>

## type property — nexthop / 100013001020 / 4

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1111220312132012-1322020311021112-1322001130323312-1233123202113332-0013121213100021-0000330133003013-3333102111021301-0133230003101133"></a>

## Next pages — nexthop / 100013001020 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_vpc_site--reference--group-005.md#canonical-2232223330003211-2000210011021301-2010313122210120-3013022212322020-1330111223031311-0030231113111320-3312202013321232-0131131220121010)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-005.md#canonical-1122313131110100-1133033022213222-1221213123112011-0332120120301120-2000112121220222-2130301232011222-2133330210000322-3232033032010100)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2232223330003211-2000210011021301-2010313122210120-3013022212322020-1330111223031311-0030231113111320-3312202013321232-0131131220121010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331200301032220-3133122202321022-2310312231230203-2112322110311330-0112013102133230-1013130211031130-3130132003303010-1010312231111231"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 110313322323 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-005.md#canonical-1003022201011203-3010322212000002-1202203120131010-2233232232031000-1323321011221123-0010311020110322-3210032321002232-3011120323113110)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-3100032023013033-2210010102230001-3132321221031222-1011233200232000-3311302101310330-0330100223100311-2201032212333231-0110230230020210"></a>

Type: `"object"`. list nested block, Optional.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0131133001301011-2310122130322312-0320010001320202-0031331203123121-3320233230120120-1331310312201221-3321310321011031-2213232310321001"></a>

## Direct properties — interface / 110313322323 / 3

<a id="canonical-3032331120122000-0111203213220331-0312331333103211-3133120303222112-3132000020021013-3323303223032120-0303330033231121-0330133301002103"></a>

<a id="canonical-3213112221320131-3112231322213312-3300121012213322-2012003301031213-3132030221203301-1213313232132332-0011210012222223-2112132313212303"></a>

## kind property — interface / 110313322323 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2202030200323120-0103321200213320-2230033233122112-0112210220132230-2223123212233133-0033221111103021-2012332202032203-2231001020232322"></a>

<a id="canonical-2321112112320130-3313210230113100-2310010031103222-1302230211321332-1300022311211303-1311333032311101-3022223223210001-0201101110012332"></a>

## name property — interface / 110313322323 / 5

Type: `"string"`. Optional.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0131310132233110-3113002320220331-2021121011111302-0200211210020102-0001303022201322-2230313323000310-3220022111131000-2010012132313201"></a>

<a id="canonical-1230312011021101-3222012122000020-0120122303312202-1333303101102322-2030331321303022-2131023330303021-0203112322123111-2313122133111331"></a>

## namespace property — interface / 110313322323 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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
  }
}
```

<a id="canonical-1011203121000022-2300300022320122-3031311232100023-3320220200101310-0311233123103111-2130102001012322-1302201131032323-1233113132233121"></a>

<a id="canonical-3023321233322312-3323302131303322-3110310230202302-0313311202212201-2121201222131320-0321111320001313-3130013031232020-3102101310023233"></a>

## tenant property — interface / 110313322323 / 7

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0331023011302013-2323120212031000-0112223210122230-2032001332312203-0021233300131333-1230102021021020-0101113102330023-0102212113020202"></a>

<a id="canonical-1011323103020001-0311201301021001-1003033010033323-1331311033221000-1020323220021231-0033333012231211-2221130032023003-3310023201003013"></a>

## uid property — interface / 110313322323 / 8

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3103003323122031-3221332231312321-1323011030003222-1031333132211310-0103032210231331-2022332321031023-0332010112110031-0332212031331102"></a>

## Next pages — interface / 110313322323 / 9

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-005.md#canonical-1003022201011203-3010322212000002-1202203120131010-2233232232031000-1323321011221123-0010311020110322-3210032321002232-3011120323113110)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1122313131110100-1133033022213222-1221213123112011-0332120120301120-2000112121220222-2130301232011222-2133330210000322-3232033032010100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133120300002231-2111100321023002-2312200332222100-1200133010302202-0133010000220311-1332220330003202-0300033030331110-2121100012030030"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 223200310003 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-005.md#canonical-1003022201011203-3010322212000002-1202203120131010-2233232232031000-1323321011221123-0010311020110322-3210032321002232-3011120323113110)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-1331232020010233-1000110301303320-2320223002033300-0333113231331103-2300311002031013-0033301202221321-0113102130031103-3210002001301133"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232203113112030-3013313313223120-2203330310033322-1102201302230321-1303000232330303-0031220022223122-3010011302202130-3002020333102203"></a>

## Direct properties — nexthop_address / 223200310003 / 3

- [dual_stack](resources--aws_vpc_site--reference--group-005.md#canonical-1123200012132332-0021001113032031-0300301213210333-2031111120021111-0331000123221000-1001121013210303-0020003232232122-2211010330230232): complete subsection reference.

- [IPv4](resources--aws_vpc_site--reference--group-005.md#canonical-3030301321312030-1233010333223213-3321301200113113-1120302112021232-3203211110112210-0012331003012033-2030132032210033-2333032131302110): complete subsection reference.

- [IPv6](resources--aws_vpc_site--reference--group-005.md#canonical-3022033223033013-0301123102221030-3313331321110020-0010231322330331-1210100100213023-3022232032222311-1031133131121301-2003223312212321): complete subsection reference.

<a id="canonical-1323010032230033-1303313021012300-1310302332311312-2032212310102331-3302303231230313-0012030333221131-1002131332302130-0221233133013320"></a>

## Next pages — nexthop_address / 223200310003 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-005.md#canonical-1123200012132332-0021001113032031-0300301213210333-2031111120021111-0331000123221000-1001121013210303-0020003232232122-2211010330230232)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_vpc_site--reference--group-005.md#canonical-3030301321312030-1233010333223213-3321301200113113-1120302112021232-3203211110112210-0012331003012033-2030132032210033-2333032131302110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_vpc_site--reference--group-005.md#canonical-3022033223033013-0301123102221030-3313331321110020-0010231322330331-1210100100213023-3022232032222311-1031133131121301-2003223312212321)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-005.md#canonical-1003022201011203-3010322212000002-1202203120131010-2233232232031000-1323321011221123-0010311020110322-3210032321002232-3011120323113110)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1123200012132332-0021001113032031-0300301213210333-2031111120021111-0331000123221000-1001121013210303-0020003232232122-2211010330230232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321013202202322-0202100322000123-2302301200030320-3111103113333033-0200210211231033-1012211300131013-1223330010333112-0330210133120233"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 221123103320 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-005.md#canonical-1003022201011203-3010322212000002-1202203120131010-2233232232031000-1323321011221123-0010311020110322-3210032321002232-3011120323113110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-005.md#canonical-1122313131110100-1133033022213222-1221213123112011-0332120120301120-2000112121220222-2130301232011222-2133330210000322-3232033032010100)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-3230220222300301-3301010232323303-1232211001310123-2333331033001212-3211031002123211-3220313301011133-1202112000110021-1131022220101212"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

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

Terraform syntax:

```terraform
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-0010131202121200-0133211220000021-3231302333111111-1101223321133033-2003032001331130-3113330122221012-1030102033202120-3003200211011011"></a>

## Direct properties — dual_stack / 221123103320 / 3

- [IPv4](resources--aws_vpc_site--reference--group-005.md#canonical-0030303230000323-2321130300000123-1033010300111000-1203313300221330-1301230332202323-0113130003101200-2130310231001313-2322313212223110): complete subsection reference.

- [IPv6](resources--aws_vpc_site--reference--group-005.md#canonical-2330202202123120-0113120000303023-0230321033333121-0322330023332301-2223212031200221-3313132232330022-0011120330310010-0133311232300121): complete subsection reference.

<a id="canonical-0330213133112021-0120332213221210-1011023300323233-3131020123223123-1211101010100000-1230203123033033-0121031023232112-1130111213313120"></a>

## Next pages — dual_stack / 221123103320 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_vpc_site--reference--group-005.md#canonical-0030303230000323-2321130300000123-1033010300111000-1203313300221330-1301230332202323-0113130003101200-2130310231001313-2322313212223110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_vpc_site--reference--group-005.md#canonical-2330202202123120-0113120000303023-0230321033333121-0322330023332301-2223212031200221-3313132232330022-0011120330310010-0133311232300121)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-005.md#canonical-1122313131110100-1133033022213222-1221213123112011-0332120120301120-2000112121220222-2130301232011222-2133330210000322-3232033032010100)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0030303230000323-2321130300000123-1033010300111000-1203313300221330-1301230332202323-0113130003101200-2130310231001313-2322313212223110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103323331102000-0213011003021103-0311113013110210-3002021132332201-1300333100222231-2031211311100131-1331333103011311-0022033211112020"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 303233332001 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-005.md#canonical-1003022201011203-3010322212000002-1202203120131010-2233232232031000-1323321011221123-0010311020110322-3210032321002232-3011120323113110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-005.md#canonical-1122313131110100-1133033022213222-1221213123112011-0332120120301120-2000112121220222-2130301232011222-2133330210000322-3232033032010100)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-005.md#canonical-1123200012132332-0021001113032031-0300301213210333-2031111120021111-0331000123221000-1001121013210303-0020003232232122-2211010330230232)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-1312100002230131-3332322102022120-2033130020033230-1332011011113131-2111201000310330-0122322013102023-2210120200002231-3002210100310111"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310211222202011-0131101010300031-3101222022313331-0322323123223111-0233322123030013-0210003211113200-1102023201311101-0323110110233122"></a>

## Direct properties — IPv4 / 303233332001 / 3

<a id="canonical-0201001023100021-2300300233122211-3302111020210333-2200120303033011-3100303323203301-0121123113023220-3110303122223003-1000323332120120"></a>

<a id="canonical-3110130100311121-2033020231302221-3132032103202323-2332202010313013-2100230310100330-3233233332320123-0322133103123013-2030333232231013"></a>

## addr property — IPv4 / 303233332001 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-3233031202130333-3312030332311300-1212322030012021-1113332320302131-0100123212112203-0000113220330003-0332123031211030-0310130312001322"></a>

## Next pages — IPv4 / 303233332001 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-005.md#canonical-1123200012132332-0021001113032031-0300301213210333-2031111120021111-0331000123221000-1001121013210303-0020003232232122-2211010330230232)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2330202202123120-0113120000303023-0230321033333121-0322330023332301-2223212031200221-3313132232330022-0011120330310010-0133311232300121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030131113122301-1333300220102122-2023301221033330-0201320010002221-1130120023010212-0031110201320111-2103030123123103-2100133103100313"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 112110113231 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-005.md#canonical-1003022201011203-3010322212000002-1202203120131010-2233232232031000-1323321011221123-0010311020110322-3210032321002232-3011120323113110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-005.md#canonical-1122313131110100-1133033022213222-1221213123112011-0332120120301120-2000112121220222-2130301232011222-2133330210000322-3232033032010100)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-005.md#canonical-1123200012132332-0021001113032031-0300301213210333-2031111120021111-0331000123221000-1001121013210303-0020003232232122-2211010330230232)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-3333033013031030-1213222202321321-3013010323032303-0310033302013202-2311303012311210-0003102102022113-0222213012131102-0210332230031020"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

Terraform syntax:

```terraform
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-1100301131301132-2012300121000131-1311011320211200-2212111030033020-3022200130100102-1233000120322123-1033111112212121-3222303131012231"></a>

## Direct properties — IPv6 / 112110113231 / 3

<a id="canonical-2201330320130021-0100013123333003-3313102102032213-0002331133300120-0333310222100133-0122132302311312-0211210100312230-0111310233312033"></a>

<a id="canonical-1213322233223330-3120211120123011-3013033130230231-3333221312010303-2112030131313301-3213110222233322-2200331113311231-2021032312212223"></a>

## addr property — IPv6 / 112110113231 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3031010212212103-0203033101020222-0233132222013230-1000302101201011-1201131132330210-2133002233323131-1331133120211012-1000201001021203"></a>

## Next pages — IPv6 / 112110113231 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-005.md#canonical-1123200012132332-0021001113032031-0300301213210333-2031111120021111-0331000123221000-1001121013210303-0020003232232122-2211010330230232)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3030301321312030-1233010333223213-3321301200113113-1120302112021232-3203211110112210-0012331003012033-2030132032210033-2333032131302110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320013001023100-0110212222231103-2230330013322310-0213201131233032-3103001031023323-2333213303220221-2000032322031031-0121323233021212"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 032120322222 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-005.md#canonical-1003022201011203-3010322212000002-1202203120131010-2233232232031000-1323321011221123-0010311020110322-3210032321002232-3011120323113110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-005.md#canonical-1122313131110100-1133033022213222-1221213123112011-0332120120301120-2000112121220222-2130301232011222-2133330210000322-3232033032010100)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-0212333220001010-1320323311000022-2120203303112100-2110033102321320-0213132113311122-1013210020220211-1213112103300122-0010002300200112"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312122230122212-0233323031012201-3232222033313211-3123010122132211-2113332022113012-1013022111113010-3213202330330132-1213210002130330"></a>

## Direct properties — IPv4 / 032120322222 / 3

<a id="canonical-1302010223032020-3022330320031113-1133122021300020-0221211303331303-3220032120222002-1131331231321010-2313020213203013-2030032130021231"></a>

<a id="canonical-3112110102322023-3111223013201123-0103100212032211-2102120100201210-1230132332110221-2023002101310102-0200121110123001-0201220200011201"></a>

## addr property — IPv4 / 032120322222 / 4

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-3100332311300222-0022010012211212-2131203321101100-0133223300301121-2103211203221030-1131211100102023-2113320333230133-0013323011313003"></a>

## Next pages — IPv4 / 032120322222 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-005.md#canonical-1122313131110100-1133033022213222-1221213123112011-0332120120301120-2000112121220222-2130301232011222-2133330210000322-3232033032010100)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3022033223033013-0301123102221030-3313331321110020-0010231322330331-1210100100213023-3022232032222311-1031133131121301-2003223312212321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221120201113011-3321133030002201-1111232222312022-0211123301223022-3120101102131210-1110203110102223-3000300100033233-0323230233220333"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 001300220110 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-005.md#canonical-1003022201011203-3010322212000002-1202203120131010-2233232232031000-1323321011221123-0010311020110322-3210032321002232-3011120323113110)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-005.md#canonical-1122313131110100-1133033022213222-1221213123112011-0332120120301120-2000112121220222-2130301232011222-2133330210000322-3232033032010100)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-2120023203201311-0133112132031300-2021103003201011-1013202321001211-0122000213330320-1321011302212120-2110022030203122-1212200302321310"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

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

Terraform syntax:

```terraform
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2110132211032101-1121000331102013-1212230132102331-1132330220322233-0100221303313301-3231230111202222-1111213131130120-0203303212011131"></a>

## Direct properties — IPv6 / 001300220110 / 3

<a id="canonical-0110131231012110-3031010122321333-2103102031103023-3023033121130201-2210130021013231-0110010202133323-3113000331102110-0213121222122300"></a>

<a id="canonical-3132320232110002-3122021013330321-3201133103032110-2023101133202030-2222012103123101-3001122323320122-3013132113030030-0131122322213311"></a>

## addr property — IPv6 / 001300220110 / 4

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3223221121113212-2322001021322010-2211312200203232-3312032031320211-3110232221203301-3132031030322200-0313021313023321-1122100031033122"></a>

## Next pages — IPv6 / 001300220110 / 5

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-005.md#canonical-1122313131110100-1133033022213222-1221213123112011-0332120120301120-2000112121220222-2130301232011222-2133330210000322-3232033032010100)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0203332230232223-0203331231222032-1323312021032100-1301001321301033-1231112012010132-0010211300210001-1001131223113330-3131232122211023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232323111210111-2330300011203101-2331211312100301-3311221023310111-1030032232113233-0002133100301010-1323231110032010-1231003030102213"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets — subnets / 331022132001 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-1120023002032200-3013212021321121-3022202233210111-3310122213020011-3212011201302213-3313200220300321-3302323133313102-2022133333331121"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201320232220302-3121001230210032-3033002123302313-1121331310123030-1130012320211122-2221221211330333-2303121221010120-3213131102223012"></a>

## Direct properties — subnets / 331022132001 / 3

- [IPv4](resources--aws_vpc_site--reference--group-005.md#canonical-2211200012001023-2020211021302120-2221020032002022-0310000322133103-2123310021323320-0302301121100302-3021113201233223-0013031011022002): complete subsection reference.

- [IPv6](resources--aws_vpc_site--reference--group-005.md#canonical-0113013202031213-2021211113033303-3133323331031313-2033232022212032-2122012031133301-3301121333210312-1003300020012322-2211300330331130): complete subsection reference.

<a id="canonical-3011001231101332-0111212003312122-2132203030031121-1000030331123232-0110230022301023-1003330000232102-2103232120021333-3301000023030220"></a>

## Next pages — subnets / 331022132001 / 4

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_vpc_site--reference--group-005.md#canonical-2211200012001023-2020211021302120-2221020032002022-0310000322133103-2123310021323320-0302301121100302-3021113201233223-0013031011022002)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_vpc_site--reference--group-005.md#canonical-0113013202031213-2021211113033303-3133323331031313-2033232022212032-2122012031133301-3301121333210312-1003300020012322-2211300330331130)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2211200012001023-2020211021302120-2221020032002022-0310000322133103-2123310021323320-0302301121100302-3021113201233223-0013031011022002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202101203110022-0000303233213303-3300011203031003-1210233331313023-1203231111133212-0023203320123331-3303213132221200-2312100330103331"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 330001310103 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-005.md#canonical-0203332230232223-0203331231222032-1323312021032100-1301001321301033-1231112012010132-0010211300210001-1001131223113330-3131232122211023)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-3213111220201203-0220320100333312-2120113102120302-1203100330223202-2210321100132202-0331221200010230-1113210113322123-0223032000133233"></a>

Type: `"object"`. single nested block, Optional.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

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

Terraform syntax:

```terraform
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-1022130030123130-3123312203113121-2000311121003322-0021300101301323-1111222132103122-3332012122301333-3210333332302233-0132111332231121"></a>

## Direct properties — IPv4 / 330001310103 / 3

<a id="canonical-0330212212013130-1131033112130203-1100121230211311-2123010300023310-1212321311303330-3220301001002322-1013102223301033-1101033201010101"></a>

<a id="canonical-0130203020010133-2130021101210033-0210012111131231-1130330311030020-0001103303120212-2332203031333303-2122201302312102-2303300013333331"></a>

## plen property — IPv4 / 330001310103 / 4

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-3302330000200201-3313223201222120-3310222222212230-3200323100113021-0221333330211032-0302022001120011-1101113211320000-2213020202320311"></a>

<a id="canonical-1000231021322123-1313011211112033-3013100233201023-1312311310133013-2301010110210131-3331301301301203-2321002332330111-3021310222232101"></a>

## prefix property — IPv4 / 330001310103 / 5

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-0002220000131303-1103100330030122-1022201131012232-3010331101323221-1013200211023211-2213000011122233-0120211031120302-0323103021023233"></a>

## Next pages — IPv4 / 330001310103 / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-005.md#canonical-0203332230232223-0203331231222032-1323312021032100-1301001321301033-1231112012010132-0010211300210001-1001131223113330-3131232122211023)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0113013202031213-2021211113033303-3133323331031313-2033232022212032-2122012031133301-3301121333210312-1003300020012322-2211300330331130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103330302110131-0222230100210303-0131100302112100-1032112131030103-2031132102312300-2210313103131030-3310212023222213-1102011323002332"></a>

## voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 301222302213 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211)
- [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-0023131221002122-2031312121313000-3020303103232311-3210201322122103-3021301133132302-2233232131320320-1213033331100313-1202013210111320)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-1233002131021222-1201333013120231-2032331102133013-0012302320213230-0211210003113321-3311100102200222-3200322203312311-2130100201013033)
- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-005.md#canonical-0203332230232223-0203331231222032-1323312021032100-1301001321301033-1231112012010132-0010211300210001-1001131223113330-3131232122211023)
- voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-3200302132120023-3113033101203011-1032320300321223-3323333202021022-0220012101123002-2300200213201122-3333000332130012-2003200011321033"></a>

Type: `"object"`. single nested block, Optional.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

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

Terraform syntax:

```terraform
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111113100033312-1030022313333301-3231213233232203-2121212032322220-0302312011211202-0331113202321201-2102233302320012-2300130312101001"></a>

## Direct properties — IPv6 / 301222302213 / 3

<a id="canonical-0210112233310320-3010303123133221-1112321003331031-2323200201003032-3132203230023130-1230201103132020-3323202021123323-2323210232321011"></a>

<a id="canonical-1232012230213333-1220311120003112-0000333001301302-1031201133110002-1113322202113230-2030102200332101-1310002023011231-2332120323100302"></a>

## plen property — IPv6 / 301222302213 / 4

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-0300132103112012-2121211311021302-0311310023323200-1202133110321031-2313220121312112-0121000021202331-3020232011030221-0222001121032110"></a>

<a id="canonical-2322310330200332-3020033321231223-2233011313020303-1023120322332300-1323200302200312-0300111133220000-2112021321100310-0112303000110012"></a>

## prefix property — IPv6 / 301222302213 / 5

Type: `"string"`. Optional.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-1013330223000222-3310303103031010-0320011231221223-2201212120222131-0310111021332321-0221123010332331-3332020203030021-0232000130012321"></a>

## Next pages — IPv6 / 301222302213 / 6

- [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-005.md#canonical-0203332230232223-0203331231222032-1323312021032100-1301001321301033-1231112012010132-0010211300210001-1001131223113330-3131232122211023)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3330300220001130-0310110133023222-3100121233311020-3003033323103312-0012202211032230-0000213020231331-2311200222213300-0312320110100312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233030010023122-3100202311013203-2302021331332212-0100103323330112-3221223023330022-0322330222231000-2301022120130223-1111333200300301"></a>

## voltstack_cluster.sm_connection_public_ip — sm_connection_public_ip / 131302302332 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.sm_connection_public_ip

<a id="canonical-2233320320000123-2022222122310320-1301101311033300-2230110312012030-2032012003111101-2223313113002032-0202111312212132-0323111130231213"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
sm_connection_public_ip = {}
```

<a id="canonical-1211011202100232-1113032121100311-2312021332320230-0211202202100032-1110111002100133-0312013222011212-2133121102311301-0023132103012021"></a>

## Direct properties — sm_connection_public_ip / 131302302332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000301222200210-3120033300203301-2032231223021103-0002020322321132-2311302212101222-2202111203121121-3030132333212201-0223311030120302"></a>

## Next pages — sm_connection_public_ip / 131302302332 / 4

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2232300111003333-1310011333201021-2203131132102031-1203110020323312-0113010123223312-0013231001331122-2202311013111223-2312133012222332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213222100213202-2031001122030120-1303003131203003-3032233000322032-3103121102212221-2223223103303331-2112131312231320-0313203202013133"></a>

## voltstack_cluster.sm_connection_pvt_ip — sm_connection_pvt_ip / 120001113022 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.sm_connection_pvt_ip

<a id="canonical-2010321100122113-0021303010321202-2302133200012312-0311200330003020-0320132112333130-0011013133021020-3310021230323012-2202113232321313"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
sm_connection_pvt_ip = {}
```

<a id="canonical-2322030001112200-0002113212220201-2310211202111312-1111000223132301-3123100131131232-2202221020332230-0233002311131212-1120030211112101"></a>

## Direct properties — sm_connection_pvt_ip / 120001113022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3021012313100022-1132131010212333-1231311333001121-0123033312121211-1300002131111300-2302010203203112-1222310123302010-2113202210233300"></a>

## Next pages — sm_connection_pvt_ip / 120001113022 / 4

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3300031200330010-1112010200123032-2113122132333302-0123210112231022-1332012112011310-0001003120130110-2310221002131212-3231202032201211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113330213111111-1113102303210032-0311302220212120-0310312103003331-3200333333132110-1020122123321311-2303321312132320-1130332130321323"></a>

## voltstack_cluster.storage_class_list — storage_class_list / 311311202233 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.storage_class_list

<a id="canonical-1202200302021313-3001210103330001-1311222112020303-0323321233012300-2332031131311111-1131212300210002-2103223120103110-1103023021221013"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this site.

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

Terraform syntax:

```terraform
storage_class_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3120211213333100-1033012313220030-1121001103322201-0201312330122312-3200310120333323-1331101231230202-3300330230210100-0331210323002100"></a>

## Direct properties — storage_class_list / 311311202233 / 3

- [storage_classes](resources--aws_vpc_site--reference--group-005.md#canonical-3321023131321101-0300112112330321-3032003030011321-1321031311221313-2210213012102020-3023330332210003-1312020230332033-0012031221113203): complete subsection reference.

<a id="canonical-0331030000303231-1023222222101123-1033011332102301-0120320113102233-1121320203011022-2002002322220030-0010010120131023-3010323123310213"></a>

## Next pages — storage_class_list / 311311202233 / 4

- [voltstack_cluster.storage_class_list.storage_classes](resources--aws_vpc_site--reference--group-005.md#canonical-3321023131321101-0300112112330321-3032003030011321-1321031311221313-2210213012102020-3023330332210003-1312020230332033-0012031221113203)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3321023131321101-0300112112330321-3032003030011321-1321031311221313-2210213012102020-3023330332210003-1312020230332033-0012031221113203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302023122310111-2001123320321012-0303310213210230-1310021332321230-0031122111123121-0120031113213003-1130103101221021-3012001132120003"></a>

## voltstack_cluster.storage_class_list.storage_classes — storage_classes / 011022113001 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.storage_class_list](resources--aws_vpc_site--reference--group-005.md#canonical-3300031200330010-1112010200123032-2113122132333302-0123210112231022-1332012112011310-0001003120130110-2310221002131212-3231202032201211)
- voltstack_cluster.storage_class_list.storage_classes

<a id="canonical-1113001233300000-3313130222322033-0102222333223002-1331332310231000-0110101100212303-1201003301100030-3322130313230102-0010311010211301"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_class_name")}
```

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_classes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0101133021322033-2102110232022133-2211030022032021-1021001220001021-3300211221020331-1230122231330201-2203031212313030-1112231230300320"></a>

## Direct properties — storage_classes / 011022113001 / 3

<a id="canonical-1332321120120123-0232331322012121-2102111220032133-3103000123211222-0310203011102003-2230011223102132-0303001232323100-0031311133222030"></a>

<a id="canonical-3213012230211012-0113323010030131-1300200022211202-2002323310001132-0112103120122111-0313020031231033-2302002222222122-3031112310123323"></a>

## default_storage_class property — storage_classes / 011022113001 / 4

Type: `"bool"`. Optional.

Make this storage class default storage class for the K8s cluster.

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

<a id="canonical-1231010113012131-2120112221012212-0322022000311331-0001233133333333-1300022012132321-2122131300000100-1313003331122003-0230320112013003"></a>

<a id="canonical-2000220222203130-1033130310001113-0310310101123120-0003013212323133-2122301130133030-2222221202022022-1230203230300301-0322321133311231"></a>

## storage_class_name property — storage_classes / 011022113001 / 5

Type: `"string"`. Optional.

Name of the storage class as it will appear in K8s.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0011132120321303-0320300132220122-1212232011103303-1213202311311220-0301203232023100-0322321001202013-2013230122023311-1311201312011123"></a>

## Next pages — storage_classes / 011022113001 / 6

- [voltstack_cluster.storage_class_list](resources--aws_vpc_site--reference--group-005.md#canonical-3300031200330010-1112010200123032-2113122132333302-0123210112231022-1332012112011310-0001003120130110-2310221002131212-3231202032201211)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2003301201220030-2310111030110321-1023333120123110-3020131220223010-3230003013133033-1023220013310123-3332200230333020-2220002022320010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310320021110233-1010320010302321-0020312333223223-0032023231011303-1230101103000122-3230332222333011-3000021111232201-3010130212231300"></a>

## vpc — vpc / 100110002211 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- vpc

<a id="canonical-3001132322310212-1213202111313031-3011022301100112-2022102222011033-1030003311232103-1100203221121201-0321330223103211-1022312112121100"></a>

Type: `"object"`. single nested block, Optional.

Defines choice about AWS VPC for a view.

Upstream description:

This defines choice about AWS VPC for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("new_vpc",
    "vpc_id")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"new_vpc\",\"vpc_id\"]"
}
```

Terraform syntax:

```terraform
vpc {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200030113200222-2323012203101230-1232000233310210-2303323011302010-1001103213203231-3123323211220300-3210103321323103-2333330200332232"></a>

## Direct properties — vpc / 100110002211 / 3

- [new_vpc](resources--aws_vpc_site--reference--group-005.md#canonical-3210012313332220-3200021210032131-1023033122320120-2203033212230103-3110133102212020-3300021333120300-0321300131112110-0222200222031002): complete subsection reference.

<a id="canonical-0233333131322132-0213020303330231-2122332110003311-0030211233222303-0011332322101002-0021133112211132-0303030321223101-1300100200201020"></a>

<a id="canonical-3233222012003011-3011203121210232-2202320010313221-1311111021230103-0021320100031001-1122002300010123-3002213313020312-2310210020122111"></a>

## vpc_id property — vpc / 100110002211 / 4

Type: `"string"`. Optional.

Exclusive with \[new\_vpc\] Information about existing VPC ID.

Upstream description:

Exclusive with \[new\_vpc\] Information about existing VPC ID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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
    },
    "pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-2221312312302123-3001120101122301-2222231033111022-1312020333120000-2233123202123212-0002011021232112-0233332130133330-2200322010011112"></a>

## Next pages — vpc / 100110002211 / 5

- [vpc.new_vpc](resources--aws_vpc_site--reference--group-005.md#canonical-3210012313332220-3200021210032131-1023033122320120-2203033212230103-3110133102212020-3300021333120300-0321300131112110-0222200222031002)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3210012313332220-3200021210032131-1023033122320120-2203033212230103-3110133102212020-3300021333120300-0321300131112110-0222200222031002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023023210322312-2111120221202230-1302032322230012-0010333301023020-3213023231310111-0222201223101101-1302101003120323-3000221223303123"></a>

## vpc.new_vpc — new_vpc / 131033103302 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [vpc](resources--aws_vpc_site--reference--group-005.md#canonical-2003301201220030-2310111030110321-1023333120123110-3020131220223010-3230003013133033-1023220013310123-3332200230333020-2220002022320010)
- vpc.new_vpc

<a id="canonical-1000101300233111-3223231002330120-1123102332102220-0003311210200011-2013002230201230-2101020023011232-1132302230012211-2131122032031112"></a>

Type: `"object"`. single nested block, Optional.

AWS VPC Parameters. Parameters to create new AWS VPC.

Upstream description:

Parameters to create new AWS VPC.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("primary_ipv4"),
  validators.ConflictingObjectAttributes("autogenerate",
    "name_tag")}
```

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

Terraform syntax:

```terraform
new_vpc {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220300301033030-0112021122312031-2322013310333302-3010103200332003-2011023313200200-2303230123122110-3131203012211003-0002213332031331"></a>

## Direct properties — new_vpc / 131033103302 / 3

- [autogenerate](resources--aws_vpc_site--reference--group-005.md#canonical-0012033022123102-2121021231222110-2321213213231030-1332331023323221-2130200312033232-1121023101122301-3231133210103233-0132002213112210): complete subsection reference.

<a id="canonical-3323300211112332-0220003000010331-0003032311232003-0201122313321222-2213012213012101-3032020303120012-3330300202101112-2332003311130032"></a>

<a id="canonical-2200321033213223-3211223003012023-3031011123020001-3310013100311022-0311202203301021-0133031113333213-3123030001322331-2013013022110113"></a>

## name_tag property — new_vpc / 131033103302 / 4

Type: `"string"`. Optional.

Exclusive with \[autogenerate\] Specify the VPC Name.

Upstream description:

Exclusive with \[autogenerate\] Specify the VPC Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-1231200203213232-3200330111111322-1000213312300220-0212200101222003-1022231223322212-0210230001100100-1201011130100202-0001300021000123"></a>

<a id="canonical-2001010333232211-0201122002021332-0133232032310333-1101201213223032-3303000100013102-2003232110212330-3013230313300021-0301221030112333"></a>

## primary_ipv4 property — new_vpc / 131033103302 / 5

Type: `"string"`. Optional.

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

<a id="canonical-3223121131301013-3220021230303213-2200230333122123-0310303001001201-0122111020010001-2121233322032122-0301100231002222-0312110202311323"></a>

## Next pages — new_vpc / 131033103302 / 6

- [vpc.new_vpc.autogenerate](resources--aws_vpc_site--reference--group-005.md#canonical-0012033022123102-2121021231222110-2321213213231030-1332331023323221-2130200312033232-1121023101122301-3231133210103233-0132002213112210)
- [vpc](resources--aws_vpc_site--reference--group-005.md#canonical-2003301201220030-2310111030110321-1023333120123110-3020131220223010-3230003013133033-1023220013310123-3332200230333020-2220002022320010)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0012033022123102-2121021231222110-2321213213231030-1332331023323221-2130200312033232-1121023101122301-3231133210103233-0132002213112210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110311230202211-0130001100101120-0111000312012212-0123032130101003-2222000033132310-2133132110131022-0113233011222223-1033220332222333"></a>

## vpc.new_vpc.autogenerate — autogenerate / 111311121320 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [vpc](resources--aws_vpc_site--reference--group-005.md#canonical-2003301201220030-2310111030110321-1023333120123110-3020131220223010-3230003013133033-1023220013310123-3332200230333020-2220002022320010)
- [vpc.new_vpc](resources--aws_vpc_site--reference--group-005.md#canonical-3210012313332220-3200021210032131-1023033122320120-2203033212230103-3110133102212020-3300021333120300-0321300131112110-0222200222031002)
- vpc.new_vpc.autogenerate

<a id="canonical-3222022302202232-2122210101110310-0322020222301021-1111013212220322-0020312013122220-1310013221321110-0022131311122000-2130021131131001"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
autogenerate = {}
```

<a id="canonical-3231002001120003-3202030232022000-3131021310002110-0020123331200202-1300133333221302-2333020310132021-3132222312200322-0333010132003000"></a>

## Direct properties — autogenerate / 111311121320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322330310301330-1001311010222311-0102121221010211-3203022130031110-2130030210321210-1233203203320230-2010102122200123-2313323012312110"></a>

## Next pages — autogenerate / 111311121320 / 4

- [vpc.new_vpc](resources--aws_vpc_site--reference--group-005.md#canonical-3210012313332220-3200021210032131-1023033122320120-2203033212230103-3110133102212020-3300021333120300-0321300131112110-0222200222031002)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2110022231300120-3130210123131302-1331212201020032-2333303013221203-0003211002203112-2311312323322011-1330000301331202-1312022031112120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030322321100002-2310212111213012-1331122102313032-1221201211001103-1021202033032111-3320222220022332-1231023220222020-3232322110223303"></a>

## waf_signatures — waf_signatures / 131312212011 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- waf_signatures

<a id="canonical-2011112213031031-1333203331230322-0303120322210030-0022200112003000-1013232101022021-1222320213122133-0303033102333320-1001301201100120"></a>

Type: `"object"`. single nested block, Optional.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("automatic",
    "manual")}
```

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

Terraform syntax:

```terraform
waf_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111201111313102-1220210320212120-0110000020201322-1010322213320220-3000023202202201-3033122033332023-1122230131301221-1033122320303000"></a>

## Direct properties — waf_signatures / 131312212011 / 3

- [automatic](resources--aws_vpc_site--reference--group-005.md#canonical-3000122211330130-0232103310301002-0330120331330302-0332111000301002-3030001023200010-0022231131001331-0012133232020021-0133200312303331): complete subsection reference.

- [manual](resources--aws_vpc_site--reference--group-005.md#canonical-1013210112222211-0330311012221023-1003133001312121-2313213103022021-0212220212213010-3000203331223323-0233123230030322-0031000113332320): complete subsection reference.

<a id="canonical-0323232023132222-3010223320202323-0313230022103021-3310232011002230-0322331103111321-3201131233121133-2303323331201133-3301020132030133"></a>

## Next pages — waf_signatures / 131312212011 / 4

- [waf_signatures.automatic](resources--aws_vpc_site--reference--group-005.md#canonical-3000122211330130-0232103310301002-0330120331330302-0332111000301002-3030001023200010-0022231131001331-0012133232020021-0133200312303331)
- [waf_signatures.manual](resources--aws_vpc_site--reference--group-005.md#canonical-1013210112222211-0330311012221023-1003133001312121-2313213103022021-0212220212213010-3000203331223323-0233123230030322-0031000113332320)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3000122211330130-0232103310301002-0330120331330302-0332111000301002-3030001023200010-0022231131001331-0012133232020021-0133200312303331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102102213013002-1323131131110321-3312021002022032-2100302103010121-3310330333113203-2330123333203133-2203121332233302-3001321331113230"></a>

## waf_signatures.automatic — automatic / 020232022331 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [waf_signatures](resources--aws_vpc_site--reference--group-005.md#canonical-2110022231300120-3130210123131302-1331212201020032-2333303013221203-0003211002203112-2311312323322011-1330000301331202-1312022031112120)
- waf_signatures.automatic

<a id="canonical-0021300122232232-3232320311330200-2021330203000220-0301120101002231-3113032301220302-1301001301210003-0012333220110100-1211022203031201"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
automatic = {}
```

<a id="canonical-1033101021022011-0130031202231113-2122123230221323-0320120132103003-0010131110321003-3013310113020223-0201320210202113-3203010201011013"></a>

## Direct properties — automatic / 020232022331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311330022003202-2201332331200101-0013223320321211-1023233222122033-3010131230003033-0103101011212130-1302103233133300-3303323313202311"></a>

## Next pages — automatic / 020232022331 / 4

- [waf_signatures](resources--aws_vpc_site--reference--group-005.md#canonical-2110022231300120-3130210123131302-1331212201020032-2333303013221203-0003211002203112-2311312323322011-1330000301331202-1312022031112120)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1013210112222211-0330311012221023-1003133001312121-2313213103022021-0212220212213010-3000203331223323-0233123230030322-0031000113332320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200001212022233-2310101001113122-0012113112120321-2302301313133233-2203300001213310-1012330010223313-0003300301312031-2000103330133310"></a>

## waf_signatures.manual — manual / 231123133133 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [waf_signatures](resources--aws_vpc_site--reference--group-005.md#canonical-2110022231300120-3130210123131302-1331212201020032-2333303013221203-0003211002203112-2311312323322011-1330000301331202-1312022031112120)
- waf_signatures.manual

<a id="canonical-3022031212312110-0313023132123022-3102200212221131-3102102201012013-2303233212002130-1103022123321133-0102222020000202-3220333123032123"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
manual = {}
```

<a id="canonical-0033010123001231-0102311210010110-2310131321031123-0002312200112201-3133320023300030-3232120201132232-1131322112123103-2001013033313101"></a>

## Direct properties — manual / 231123133133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203202232330020-2012210300133002-3332212122101232-2201210110231020-0203023331301222-1102132031313220-0101000010113321-1130121201213230"></a>

## Next pages — manual / 231123133133 / 4

- [waf_signatures](resources--aws_vpc_site--reference--group-005.md#canonical-2110022231300120-3130210123131302-1331212201020032-2333303013221203-0003211002203112-2311312323322011-1330000301331202-1312022031112120)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
