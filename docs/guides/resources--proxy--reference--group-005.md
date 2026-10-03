---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-0132300030132321-2203101323301222-1020033031122021-1333132213110120-3001233330013223-1211033230000233-2212103020332222-2301331132121311"></a>

## Next pages — site_local_network / 023111302111 / 4

- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323022202332022-3031023120231021-0003033031001021-1122020301302033-3320310220112312-0111323200133230-1231202211130133-3213300132201102"></a>

## site_virtual_sites — site_virtual_sites / 221203320002 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- site_virtual_sites

<a id="canonical-0333011221000001-0111132130112111-0231020012230322-3301122013323311-0123011330302122-0120213032111302-3301221321211312-3112303021033110"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a VIP on specific sites.

Upstream description:

This defines a way to advertise a VIP on specific sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where")}
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
site_virtual_sites {
  # Configure direct properties listed below.
}
```

<a id="canonical-1033110312101213-1213100233021332-2313023211012313-1322122221031010-1031031301102023-1202211213112121-3003223013232200-2100202001322131"></a>

## Direct properties — site_virtual_sites / 221203320002 / 3

- [advertise_where](resources--proxy--reference--group-005.md#canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002): complete subsection reference.

<a id="canonical-3210203132302032-3321200123200233-3012302210322231-3301030213121311-0011121113332300-0030020023101022-0200010023013000-0212033312102120"></a>

## Next pages — site_virtual_sites / 221203320002 / 4

- [site_virtual_sites.advertise_where](resources--proxy--reference--group-005.md#canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130122003113103-3330210333021220-0221223003112203-0011321211203331-2301023222021300-1330312121133000-1002312021111102-3130010310133003"></a>

## site_virtual_sites.advertise_where — advertise_where / 112121120111 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [site_virtual_sites](resources--proxy--reference--group-005.md#canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302)
- site_virtual_sites.advertise_where

<a id="canonical-2100110201320221-3330213111022300-3111200010313311-0331110213021031-2123233022023023-3121012211012322-3222213131320121-0203133310233020"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("port",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
advertise_where {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220033321311330-0022311010131231-1101001132202130-2032131222003131-0213122203321333-2000030320123212-2212201232302001-3121012222203221"></a>

## Direct properties — advertise_where / 112121120111 / 3

<a id="canonical-2313331113002231-0223313032111320-2121123130113111-2112002101331020-3013211030130110-0320211001020010-0312303030130123-3100331002223113"></a>

<a id="canonical-0033030111201120-0000310031033231-3323220001211221-3230031103111011-1300212200133302-1313220331220101-3310111011130130-1213011201330323"></a>

## port property — advertise_where / 112121120111 / 4

Type: `"number"`. Optional.

Exclusive with \[use\_default\_port\] TCP port to Listen.

Upstream description:

Exclusive with \[use\_default\_port\] TCP port to Listen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [site](resources--proxy--reference--group-005.md#canonical-2110203130033231-3030200320001133-0300102313110100-3302313102131100-1110020000222223-1123310133220200-1021130110103102-1123010013130013): complete subsection reference.

- [use_default_port](resources--proxy--reference--group-005.md#canonical-2012100001011113-3312200301101231-3022131230331303-2320221302012122-0221222000320123-2101221112033311-1111331100113220-1312222100321310): complete subsection reference.

- [virtual_site](resources--proxy--reference--group-005.md#canonical-1112302320303012-3313113022211002-1011302211200120-2133320023002120-0212331100221302-2013011030032231-0303300001131201-0121103010123110): complete subsection reference.

<a id="canonical-2121332021001011-0032231221303121-1220133200311301-1312210211223201-3023003121322332-3231312123013322-3313100110320221-3002330133212023"></a>

## Next pages — advertise_where / 112121120111 / 5

- [site_virtual_sites.advertise_where.site](resources--proxy--reference--group-005.md#canonical-2110203130033231-3030200320001133-0300102313110100-3302313102131100-1110020000222223-1123310133220200-1021130110103102-1123010013130013)
- [site_virtual_sites.advertise_where.use_default_port](resources--proxy--reference--group-005.md#canonical-2012100001011113-3312200301101231-3022131230331303-2320221302012122-0221222000320123-2101221112033311-1111331100113220-1312222100321310)
- [site_virtual_sites.advertise_where.virtual_site](resources--proxy--reference--group-005.md#canonical-1112302320303012-3313113022211002-1011302211200120-2133320023002120-0212331100221302-2013011030032231-0303300001131201-0121103010123110)
- [site_virtual_sites](resources--proxy--reference--group-005.md#canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2110203130033231-3030200320001133-0300102313110100-3302313102131100-1110020000222223-1123310133220200-1021130110103102-1123010013130013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322021000330210-0332101121133223-0202232222321221-2321220300101001-2000013121313121-2232012231203230-0223102323121202-1000321310331133"></a>

## site_virtual_sites.advertise_where.site — site / 220020300122 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [site_virtual_sites](resources--proxy--reference--group-005.md#canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-005.md#canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002)
- site_virtual_sites.advertise_where.site

<a id="canonical-3322321212133033-3111230213233223-2022230112131332-3202102321032031-0110133333110320-3302032210313332-0303331212200032-3310010110232220"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a CE site along with network type and an optional IP address where a load
balancer could be advertised.

Upstream description:

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1103323002332032-3023030130113210-3120330320203102-2301210030010120-2102110101012332-0302110120312332-0321000120213102-2101321121301320"></a>

## Direct properties — site / 220020300122 / 3

<a id="canonical-3020123232013323-3210201103102201-2103012332120020-3002111323022032-2000010312231311-3232103100233123-1120022203113223-0110122221133311"></a>

<a id="canonical-2023022213331122-1220233312033103-0300131002330212-2020303302121101-0111103031331012-0022102202312123-0332313332130123-0331103223222031"></a>

## ip property — site / 220020300122 / 4

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

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

<a id="canonical-3332301132001121-2330200212011001-1133312132310230-2123020113231313-3210101012031113-1322322130033302-3210222023131233-1023120311110100"></a>

<a id="canonical-3001201232100123-1221122231110223-0102112013212021-2032322130113311-1101201130032100-3330222003222110-3022233103301230-2010001111011301"></a>

## network property — site / 220020300122 / 5

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](resources--proxy--reference--group-005.md#canonical-0022310220020012-1031101221022300-3210031221310121-3032030331230003-0221201323303013-3013333303211123-3130211030221323-3233002033123330): complete subsection reference.

<a id="canonical-0100001120021012-2033121121303003-0031332231311233-0111130013031232-2013002201233023-0133122131331000-0010132202332220-0131222202123132"></a>

## Next pages — site / 220020300122 / 6

- [site_virtual_sites.advertise_where.site.site](resources--proxy--reference--group-005.md#canonical-0022310220020012-1031101221022300-3210031221310121-3032030331230003-0221201323303013-3013333303211123-3130211030221323-3233002033123330)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-005.md#canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0022310220020012-1031101221022300-3210031221310121-3032030331230003-0221201323303013-3013333303211123-3130211030221323-3233002033123330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030122203100232-2300022313100313-3222113130211301-3100102203110113-1311333003132032-2013023032103133-1320203013201213-0300300000302113"></a>

## site_virtual_sites.advertise_where.site.site — site / 103030112131 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [site_virtual_sites](resources--proxy--reference--group-005.md#canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-005.md#canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002)
- [site_virtual_sites.advertise_where.site](resources--proxy--reference--group-005.md#canonical-2110203130033231-3030200320001133-0300102313110100-3302313102131100-1110020000222223-1123310133220200-1021130110103102-1123010013130013)
- site_virtual_sites.advertise_where.site.site

<a id="canonical-2103321233003120-2201200120122210-3131312013011000-2031220232012110-0201110232321202-3133001211302030-3132303203030301-1123130121013213"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012022313033300-1111013112123120-2302001002020101-3120321122130310-3010203201021311-3330332113110322-0211002022230222-1331120201121233"></a>

## Direct properties — site / 103030112131 / 3

<a id="canonical-3121310320130220-2331301102112020-1220113322032121-0031302332011122-0020033121333001-0212302110201103-2231100123111333-0232333022132113"></a>

<a id="canonical-1311201010012323-3002233233212211-3101211023210033-3120011123012122-3213303002303212-2111101100210103-3131213021122301-1233101103323130"></a>

## name property — site / 103030112131 / 4

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

<a id="canonical-1010221100200201-2213001313323102-2121021201110331-3202330100130002-3311023221213023-3111003032100031-2330030133310310-1111130201031233"></a>

<a id="canonical-0203300322320233-2131231300121310-3330112220122320-0200100113131100-1102220103203220-1113210101300111-0210212012303210-1001032031020101"></a>

## namespace property — site / 103030112131 / 5

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

<a id="canonical-3133131122003012-1302032100213122-0310302322101230-1222213220330320-1110231123323312-2230211231012031-0120313203331100-2112002210310032"></a>

<a id="canonical-2001002102321003-2100013213002001-0333100302302113-3121020313213110-0313022330011013-3230013303111002-2322230300133321-2220222111232323"></a>

## tenant property — site / 103030112131 / 6

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

<a id="canonical-1121022033310301-2012110211232211-0201222100003012-1102120300313112-0200211000211120-0212013032002301-2012122200121010-2200122202022233"></a>

## Next pages — site / 103030112131 / 7

- [site_virtual_sites.advertise_where.site](resources--proxy--reference--group-005.md#canonical-2110203130033231-3030200320001133-0300102313110100-3302313102131100-1110020000222223-1123310133220200-1021130110103102-1123010013130013)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2012100001011113-3312200301101231-3022131230331303-2320221302012122-0221222000320123-2101221112033311-1111331100113220-1312222100321310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202113121200230-0121220301332133-0012221002101201-0100000130121123-2130312220032302-2110133322202202-0322120311010333-3112012312130130"></a>

## site_virtual_sites.advertise_where.use_default_port — use_default_port / 222211132311 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [site_virtual_sites](resources--proxy--reference--group-005.md#canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-005.md#canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002)
- site_virtual_sites.advertise_where.use_default_port

<a id="canonical-0122120302113102-2312133200030003-0320023221321311-1211202100300300-3232310132131300-3003202022122110-0012312121103301-1103211230011310"></a>

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
use_default_port = {}
```

<a id="canonical-0023030023303102-0302011222300000-0233112111221033-2323330312200203-0032100103021121-1222003310123020-2121210320100132-3333132033010021"></a>

## Direct properties — use_default_port / 222211132311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2132030011303031-0222322010333112-2013023002203110-3003113231321311-0010320101210213-1223031230033211-1302131322113110-2212201332123113"></a>

## Next pages — use_default_port / 222211132311 / 4

- [site_virtual_sites.advertise_where](resources--proxy--reference--group-005.md#canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1112302320303012-3313113022211002-1011302211200120-2133320023002120-0212331100221302-2013011030032231-0303300001131201-0121103010123110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320033202000200-0213131232021113-2022132213111101-3021333222321133-3312012000101013-2023102113310203-0302101230032110-2211000302231212"></a>

## site_virtual_sites.advertise_where.virtual_site — virtual_site / 312113110011 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [site_virtual_sites](resources--proxy--reference--group-005.md#canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-005.md#canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002)
- site_virtual_sites.advertise_where.virtual_site

<a id="canonical-0212121332300203-2110123100212012-1102113201011032-2333332232213032-3333021310230323-0012103113020323-1322021231120310-2011323023331022"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a customer site virtual site along with network type where a load balancer
could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0023111320330111-2332210203111010-3211031322233220-0213212300330203-2333301021111330-3213012023023322-0032111032231123-1220230000022230"></a>

## Direct properties — virtual_site / 312113110011 / 3

<a id="canonical-1023312001111300-1100323013130010-3132301100320232-2101021300021111-2000311030233303-3122230221012130-3023003123201320-2330313321010010"></a>

<a id="canonical-1321133212133233-1322133003023123-0131011003130112-1112113331303013-0322031000022202-2232033333120112-1333131333303000-1210013101001200"></a>

## network property — virtual_site / 312113110011 / 4

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](resources--proxy--reference--group-005.md#canonical-0210101233323000-1323230321122112-2322203003111110-0221010023212220-2231220113112322-2333133212033103-1212013232103323-3020232130133321): complete subsection reference.

<a id="canonical-1110120331310023-1033120211221130-3132030323131323-2302232312320312-1022303122232232-0233232220120302-2221300023030010-2233133003330203"></a>

## Next pages — virtual_site / 312113110011 / 5

- [site_virtual_sites.advertise_where.virtual_site.virtual_site](resources--proxy--reference--group-005.md#canonical-0210101233323000-1323230321122112-2322203003111110-0221010023212220-2231220113112322-2333133212033103-1212013232103323-3020232130133321)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-005.md#canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0210101233323000-1323230321122112-2322203003111110-0221010023212220-2231220113112322-2333133212033103-1212013232103323-3020232130133321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031230022203202-3112013130301000-1100332231132320-0032021210200032-1032031033020322-3300303030103233-1203023222030313-1322320320212122"></a>

## site_virtual_sites.advertise_where.virtual_site.virtual_site — virtual_site / 220133031333 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [site_virtual_sites](resources--proxy--reference--group-005.md#canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302)
- [site_virtual_sites.advertise_where](resources--proxy--reference--group-005.md#canonical-0021313233203001-3231101003103230-3120120101021312-1113111010303102-2001002322132330-2330301333313313-1012201310232331-3232210300100002)
- [site_virtual_sites.advertise_where.virtual_site](resources--proxy--reference--group-005.md#canonical-1112302320303012-3313113022211002-1011302211200120-2133320023002120-0212331100221302-2013011030032231-0303300001131201-0121103010123110)
- site_virtual_sites.advertise_where.virtual_site.virtual_site

<a id="canonical-0221220331332212-2030311010233103-3111323303330033-1303122020131120-3012203103332302-1021100230031222-2303032031301320-0012210232123021"></a>

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020112021313310-1102220221011011-1201101300223113-0220033132300200-3132020320131312-0012000020200231-3203211020321121-1321322322031031"></a>

## Direct properties — virtual_site / 220133031333 / 3

<a id="canonical-0333011010230103-0210213231232323-3232310311111101-0013203301102123-2231332122102002-1113123133120112-3300230110310210-3311023102322321"></a>

<a id="canonical-2002333210113331-2201330213320022-2230301313323013-3202013001103001-0220010000222032-2322221122011312-3313010012212001-3001013230232312"></a>

## name property — virtual_site / 220133031333 / 4

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

<a id="canonical-3010310031323201-1202312201220132-3102011121220122-1000310331002102-2111221323121221-3200312222231333-3230132003111322-1213110232113303"></a>

<a id="canonical-3110311300101020-1000120312103201-3221101211101313-1212112121202020-3110033123030123-3101301201012133-0133012222033301-0220302320312300"></a>

## namespace property — virtual_site / 220133031333 / 5

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

<a id="canonical-1232220200003130-1331033010203112-1233030100221122-0301122111200133-0110300231213002-1203020023003011-0203111220123131-3223201233230230"></a>

<a id="canonical-0200133323002310-1233320331133032-2112201202202220-0322013233023223-2220222302130213-1012021122213222-1021220200023123-1133103213321031"></a>

## tenant property — virtual_site / 220133031333 / 6

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

<a id="canonical-2310121302222103-1223311133212200-2123003230112001-2002130122111130-1013210313020332-0231213332022322-0213121133101322-2210200302001233"></a>

## Next pages — virtual_site / 220133031333 / 7

- [site_virtual_sites.advertise_where.virtual_site](resources--proxy--reference--group-005.md#canonical-1112302320303012-3313113022211002-1011302211200120-2133320023002120-0212331100221302-2013011030032231-0303300001131201-0121103010123110)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2113110122012210-0300103202002221-2012320121222331-2230022312003312-3203002101213002-1312002203112032-1202231032011220-3312233322001132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233031011202310-1331021310220132-3012010211322222-0223220230321230-0023220211022201-0130000102121021-2331130200221211-0133133132322132"></a>

## timeouts — timeouts / 200200002200 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- timeouts

<a id="canonical-3100221132113223-0221031321312001-3112300110111021-3322203020203310-3012010232232321-0222122130021011-3230012320013131-0303023223020122"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313212000201331-1111011133332023-1303233222330112-1130031311333111-2112002010200010-1122221323011333-0223303131112203-1022132033130212"></a>

## Direct properties — timeouts / 200200002200 / 3

<a id="canonical-1103120221321111-3002033010233002-3110210222233021-0231311012103102-1123033223001123-3001110112012100-2021320110011201-0300330321102300"></a>

<a id="canonical-0022020130320303-1101330230000231-0331231030313333-1021120232233330-3013120300012322-2232321333131222-0000220120203332-2302133310211323"></a>

## create property — timeouts / 200200002200 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1221303323102020-1110010213210303-2331222111130233-2013220310022102-2323220301220031-0022100223230213-3211022101003231-2032033220301300"></a>

<a id="canonical-0213210023313011-0132111030310010-1201021210010220-3020003321001320-2302330010313232-2312033003000201-2333320210022301-0201333101021332"></a>

## delete property — timeouts / 200200002200 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2320022221231000-1231021201033201-1023021213332301-1230312111300001-0323033312020123-1023201001033331-1133021011021300-2032110033212110"></a>

<a id="canonical-1210012013120322-0220120222212330-0300231202300310-3220111313221031-2030231012302221-1212232020333001-1300212221120113-0320303130300101"></a>

## read property — timeouts / 200200002200 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2232312313013221-0311031132223112-0030011111200223-1103230231301131-3200311300003133-1212112321110313-3121322030010203-3133001203330223"></a>

<a id="canonical-1330103101232200-0103300010233013-0301023310003220-2302322130022001-1000111332131002-1103323130303331-2123003322200333-3333230222023023"></a>

## update property — timeouts / 200200002200 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2312003112230211-2232003233321211-3312313223032322-1201111220300301-3321210031011101-3123101010111331-1323302332120303-0300011000313330"></a>

## Next pages — timeouts / 200200002200 / 8

- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322103130100103-1122010031120010-0023311333322203-1110310301011131-0120132113000101-0310311023130313-0123230021000031-1200113310101122"></a>

## tls_intercept — tls_intercept / 130211021303 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- tls_intercept

<a id="canonical-2300212110133011-2303221131021333-1221330301300022-0231220012001111-0213123123321222-0000030203023101-1203222021230313-0321231013011230"></a>

Type: `"object"`. single nested block, Optional.

Configuration to enable TLS interception.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_certificate",
    "volterra_certificate"),
  validators.ConflictingObjectAttributes("enable_for_all_domains",
    "policy"),
  validators.ConflictingObjectAttributes("trusted_ca_url",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-interception_policy_choice": "[\"enable_for_all_domains\",\"policy\"]",
  "x-ves-oneof-field-signing_cert_choice": "[\"custom_certificate\",\"volterra_certificate\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca_url\",\"volterra_trusted_ca\"]"
}
```

Terraform syntax:

```terraform
tls_intercept {
  # Configure direct properties listed below.
}
```

<a id="canonical-1113133301102202-1322030002301202-0110331202111330-1130121212232022-2012101332330023-2213233113121110-1223212300230213-1012130321222120"></a>

## Direct properties — tls_intercept / 130211021303 / 3

- [custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223): complete subsection reference.

- [enable_for_all_domains](resources--proxy--reference--group-005.md#canonical-0010312210200030-1032102030200221-1013231023111111-0213001030001301-3302211232102022-2010221321321011-1320131130020100-1230311302022102): complete subsection reference.

- [policy](resources--proxy--reference--group-005.md#canonical-2113311103000032-0223311313330133-0101130032201233-1131032300210211-2301121231203212-0112333210002303-3232333300200313-0031201013101103): complete subsection reference.

<a id="canonical-2302310120010303-0110132203322000-0001021231100100-2330023331030131-0023031130300232-3200023320232110-1113013233133323-1310310320311202"></a>

<a id="canonical-3300021212022120-2120310231330200-1213201211132322-2320311232001322-2223322331313032-2331303100201122-3233000002211213-3030313110030101"></a>

## trusted_ca_url property — tls_intercept / 130211021303 / 4

Type: `"string"`. Optional.

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

Upstream description:

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [volterra_certificate](resources--proxy--reference--group-005.md#canonical-0322201212300210-3122231010311203-0031222013232003-3121000333001303-2020203321212103-0110333232012202-2000012101030312-2210122003113011): complete subsection reference.

- [volterra_trusted_ca](resources--proxy--reference--group-005.md#canonical-0221120012112232-3222102003023211-2332123231332121-1022321130233300-3132233121322211-0210112032300212-1100120323003230-3333123211101001): complete subsection reference.

<a id="canonical-1021213312332131-2113132331011000-0323232103120323-0103032213010133-2011132210311213-1230220231100013-0010130331221100-0312112313120212"></a>

## Next pages — tls_intercept / 130211021303 / 5

- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- [tls_intercept.enable_for_all_domains](resources--proxy--reference--group-005.md#canonical-0010312210200030-1032102030200221-1013231023111111-0213001030001301-3302211232102022-2010221321321011-1320131130020100-1230311302022102)
- [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-2113311103000032-0223311313330133-0101130032201233-1131032300210211-2301121231203212-0112333210002303-3232333300200313-0031201013101103)
- [tls_intercept.volterra_certificate](resources--proxy--reference--group-005.md#canonical-0322201212300210-3122231010311203-0031222013232003-3121000333001303-2020203321212103-0110333232012202-2000012101030312-2210122003113011)
- [tls_intercept.volterra_trusted_ca](resources--proxy--reference--group-005.md#canonical-0221120012112232-3222102003023211-2332123231332121-1022321130233300-3132233121322211-0210112032300212-1100120323003230-3333123211101001)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232021000201100-0301220310032113-0121222121203222-1322133101200030-3022123233002220-3121220120330031-3300313122313133-3332032020113320"></a>

## tls_intercept.custom_certificate — custom_certificate / 021203132013 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- tls_intercept.custom_certificate

<a id="canonical-0130022112311203-0002331003312132-1033021013303233-2012221301021003-0123002202003302-3320330131212202-3310321213331132-0333330300202030"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for custom certificate.

Upstream description:

Handle to fetch certificate and key.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("certificate_url"),
  validators.ConflictingObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
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
  "x-ves-oneof-field-ocsp_stapling_choice": "[\"custom_hash_algorithms\",\"disable_ocsp_stapling\",\"use_system_defaults\"]"
}
```

Terraform syntax:

```terraform
custom_certificate {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033001111101210-3133133011132120-1220301201301230-1210011132101300-1021023102233231-2010303122130231-2232213230223212-2201223120312311"></a>

## Direct properties — custom_certificate / 021203132013 / 3

<a id="canonical-2201022013301311-0213100320031220-1002032113200202-1310102202002013-2330221000130002-3323111100231222-0301212033232303-1222032322333202"></a>

<a id="canonical-3000202023111201-3221230001233100-3031011122330131-3003230332131020-0032233132130022-1202100033031311-2001301110323233-1200202022103020"></a>

## certificate_url property — custom_certificate / 021203132013 / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--proxy--reference--group-005.md#canonical-3032013211022031-2003020322332212-2232310313202130-1212131023200112-2311110231233322-0213332121331132-2131203121303020-3333330323131203): complete subsection reference.

<a id="canonical-3133130011332203-2030002102133223-2312210333013330-3000211323231223-2002111120220221-1101123210002112-0201001310000103-1232113321303100"></a>

<a id="canonical-1232131220032203-2021001331112230-1011032203200323-3213330313210211-0003110212002202-0031201032300222-1031011131110122-3312313120231302"></a>

## description_spec property — custom_certificate / 021203132013 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--proxy--reference--group-005.md#canonical-3313223130002100-2133323221002213-2313132023211030-0111231131322301-3031302232031101-1030231322202333-2022231030132030-1201002010230031): complete subsection reference.

- [private_key](resources--proxy--reference--group-005.md#canonical-3321230201331111-0132320001302220-0020330020122133-1310210212012031-2232300311202031-3302001322330103-0220322112132033-2212201013212221): complete subsection reference.

- [use_system_defaults](resources--proxy--reference--group-005.md#canonical-0203212312012200-0022331233322101-3312232012110231-2001030223033320-3211223032133022-0222102213210232-0213123221113002-2200213003332120): complete subsection reference.

<a id="canonical-0232311320332233-1122002002221211-2310122333333200-3100230322110103-2012020230133033-2203013123130020-0100321100220000-1022031201203232"></a>

## Next pages — custom_certificate / 021203132013 / 6

- [tls_intercept.custom_certificate.custom_hash_algorithms](resources--proxy--reference--group-005.md#canonical-3032013211022031-2003020322332212-2232310313202130-1212131023200112-2311110231233322-0213332121331132-2131203121303020-3333330323131203)
- [tls_intercept.custom_certificate.disable_ocsp_stapling](resources--proxy--reference--group-005.md#canonical-3313223130002100-2133323221002213-2313132023211030-0111231131322301-3031302232031101-1030231322202333-2022231030132030-1201002010230031)
- [tls_intercept.custom_certificate.private_key](resources--proxy--reference--group-005.md#canonical-3321230201331111-0132320001302220-0020330020122133-1310210212012031-2232300311202031-3302001322330103-0220322112132033-2212201013212221)
- [tls_intercept.custom_certificate.use_system_defaults](resources--proxy--reference--group-005.md#canonical-0203212312012200-0022331233322101-3312232012110231-2001030223033320-3211223032133022-0222102213210232-0213123221113002-2200213003332120)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3032013211022031-2003020322332212-2232310313202130-1212131023200112-2311110231233322-0213332121331132-2131203121303020-3333330323131203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330010231031133-3020122122020232-2213032121233103-1211113302001121-0103311203211321-0311310122320212-3022031023331031-0032020101130111"></a>

## tls_intercept.custom_certificate.custom_hash_algorithms — custom_hash_algorithms / 112130000222 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- tls_intercept.custom_certificate.custom_hash_algorithms

<a id="canonical-0010221002113203-1320123333021111-2131113310302131-2223121311002101-0211030313032012-1213232231303030-0332013011332201-2012220031313001"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-1301031103210132-2212230022120230-2302221322203200-0110033231302211-2111200022021231-3020022003122230-1322201012201300-1133013002223323"></a>

## Direct properties — custom_hash_algorithms / 112130000222 / 3

<a id="canonical-1123210200321201-0103130010111311-3011202013331200-2330210113233011-2232220100323300-0310023113112211-2121012020332202-2231231231203112"></a>

<a id="canonical-0320110200323303-2301112300331202-1233131213220302-3112332021123020-1322002121121310-1303333101102311-3203011300100000-1301323202012320"></a>

## hash_algorithms property — custom_hash_algorithms / 112130000222 / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2111120212311231-3331233333333121-1221013113200323-2022032203011332-0023012301112120-3033323323213231-2330031202022312-3213000031331112"></a>

## Next pages — custom_hash_algorithms / 112130000222 / 5

- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3313223130002100-2133323221002213-2313132023211030-0111231131322301-3031302232031101-1030231322202333-2022231030132030-1201002010230031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322313100103003-1221003113210302-0123132322223230-1033032322022323-1030312213130230-0033130013211013-3310011323122333-3322222210231231"></a>

## tls_intercept.custom_certificate.disable_ocsp_stapling — disable_ocsp_stapling / 033011022031 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- tls_intercept.custom_certificate.disable_ocsp_stapling

<a id="canonical-0211213312302113-3322133023012121-0200203231303133-3323110331113331-3302203033022223-1013333112320101-2000123033001203-1130200012111122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

<a id="canonical-2331333300200301-2230222111113220-2330022332131112-2230232310000021-1201032022310312-3220132330113221-1032020332030112-3212021331033212"></a>

## Direct properties — disable_ocsp_stapling / 033011022031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220330311203133-0030003023031003-1321223023011032-2212020023222033-1021021130113020-3002132313130123-1321003223011131-0032210023320212"></a>

## Next pages — disable_ocsp_stapling / 033011022031 / 4

- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3321230201331111-0132320001302220-0020330020122133-1310210212012031-2232300311202031-3302001322330103-0220322112132033-2212201013212221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232220032100121-2121001102003033-2022103331120100-2203303113232300-2331232110121232-2111302101303122-2231000002310112-3311100201100313"></a>

## tls_intercept.custom_certificate.private_key — private_key / 322212101231 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- tls_intercept.custom_certificate.private_key

<a id="canonical-2200200320331121-1333131312130012-3322323323023132-0203231323032321-3111332333001312-1230102301032332-3202021313322200-0330202120220030"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330213211110102-2022331332011202-2330132221332003-3113130122210000-1020222221320333-0223133033000012-0212330100112312-1020323021331131"></a>

## Direct properties — private_key / 322212101231 / 3

- [blindfold_secret_info](resources--proxy--reference--group-005.md#canonical-3023221303101103-1211031223130023-0132333022302321-2023220133310312-1021103200000333-2310003100323222-2121103122121012-0230102020230211): complete subsection reference.

- [clear_secret_info](resources--proxy--reference--group-005.md#canonical-3201000030300312-0322013220130313-0013123021000321-0030321111111122-2212302321121233-2302202131031011-1103212303112102-1223020013122013): complete subsection reference.

<a id="canonical-0220000332333302-3333032022203232-2330001100302123-2023302212020122-2222201230010101-1123233230131121-2332033313331100-3331333311211302"></a>

## Next pages — private_key / 322212101231 / 4

- [tls_intercept.custom_certificate.private_key.blindfold_secret_info](resources--proxy--reference--group-005.md#canonical-3023221303101103-1211031223130023-0132333022302321-2023220133310312-1021103200000333-2310003100323222-2121103122121012-0230102020230211)
- [tls_intercept.custom_certificate.private_key.clear_secret_info](resources--proxy--reference--group-005.md#canonical-3201000030300312-0322013220130313-0013123021000321-0030321111111122-2212302321121233-2302202131031011-1103212303112102-1223020013122013)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3023221303101103-1211031223130023-0132333022302321-2023220133310312-1021103200000333-2310003100323222-2121103122121012-0230102020230211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021112131033110-3322211010302012-3103112003333101-1121322010022012-0203201200133020-1022323130220003-1022020023300322-1022332302113012"></a>

## tls_intercept.custom_certificate.private_key.blindfold_secret_info — blindfold_secret_info / 303111333312 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- [tls_intercept.custom_certificate.private_key](resources--proxy--reference--group-005.md#canonical-3321230201331111-0132320001302220-0020330020122133-1310210212012031-2232300311202031-3302001322330103-0220322112132033-2212201013212221)
- tls_intercept.custom_certificate.private_key.blindfold_secret_info

<a id="canonical-1303000210122130-0033031020131202-2223232022111303-0101023310321310-0103201102201023-2200101021211100-3303302032331312-2300211330320103"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310120021211111-1013030301032023-1320323133302213-1230133300223330-1102202330303312-0202233032300322-0130020023230010-3233201033331011"></a>

## Direct properties — blindfold_secret_info / 303111333312 / 3

<a id="canonical-2311302231311013-1313022212223011-3320300010200122-2103232330000310-0211000232310332-0220232233231302-2110310213001311-0310203102110013"></a>

<a id="canonical-3233122022310003-3200332122312303-1001232213212100-3011223033320132-0110021111210002-1101012102022201-3133302210220220-1131201332021133"></a>

## decryption_provider property — blindfold_secret_info / 303111333312 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-2302122000031322-1230230313132022-1102003222122101-2002320132312312-2002103012330122-3310002312111221-1031233102023120-1020121323132022"></a>

<a id="canonical-2111131113030132-1110013211023002-2223220210111032-1321310220302023-1201301313002233-1130001032233102-0100003210323232-2003122132201022"></a>

## location property — blindfold_secret_info / 303111333312 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3310333120321022-1021132212211032-2331030010132110-0030113122233032-3102312133000121-1111213102300120-0012303330102203-2333331113101201"></a>

<a id="canonical-2320213223210102-0110113002001111-0312111323020021-3011201222330022-0220223230211300-1012222121223131-1121321303010032-0123200332021330"></a>

## store_provider property — blindfold_secret_info / 303111333312 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-0111020310213030-0133213023322030-0122212322212010-2313320222130330-3102221020232322-2112330200231111-3311311010311230-1213011103032130"></a>

## Next pages — blindfold_secret_info / 303111333312 / 7

- [tls_intercept.custom_certificate.private_key](resources--proxy--reference--group-005.md#canonical-3321230201331111-0132320001302220-0020330020122133-1310210212012031-2232300311202031-3302001322330103-0220322112132033-2212201013212221)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3201000030300312-0322013220130313-0013123021000321-0030321111111122-2212302321121233-2302202131031011-1103212303112102-1223020013122013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013001211110031-0203231301133233-2000311033112102-3333001310220031-1320122213221120-3313033131011312-3012031100033210-3311310230032213"></a>

## tls_intercept.custom_certificate.private_key.clear_secret_info — clear_secret_info / 221023000003 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- [tls_intercept.custom_certificate.private_key](resources--proxy--reference--group-005.md#canonical-3321230201331111-0132320001302220-0020330020122133-1310210212012031-2232300311202031-3302001322330103-0220322112132033-2212201013212221)
- tls_intercept.custom_certificate.private_key.clear_secret_info

<a id="canonical-3211200113130130-2323212132131212-2010223130002213-1210211300333223-2300220012101110-2313201212010302-2010300002320333-3221010210032210"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321302311203201-1013301023122300-2332322103323301-0130003033202201-2221031331030220-3103030311310301-2023320131333321-1211120322012202"></a>

## Direct properties — clear_secret_info / 221023000003 / 3

<a id="canonical-3220233012210213-0003323133131222-2021100300232131-3201323313110031-0330213312331311-0113212001233312-3233213220032102-3010222023312203"></a>

<a id="canonical-0001322112120003-1101320322103301-0201312200022133-0112230323133130-1101330110230020-3010212030100131-0313130222310011-1232213210011002"></a>

## provider_ref property — clear_secret_info / 221023000003 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2213022103112120-1123211032221101-0033212213001120-1021332101023210-2031003323031010-1220310103320001-0333000210100022-0231120000011000"></a>

<a id="canonical-0303101200020111-1232201010123202-2133100220331230-2312212110003103-3303120011223030-1213312211130323-1321302030232120-1212320300212031"></a>

## URL property — clear_secret_info / 221023000003 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1111232320030203-0313211111332031-3231320030031332-0222001332210001-0110002330203112-1030011211000122-2333330122332012-2313123123311322"></a>

## Next pages — clear_secret_info / 221023000003 / 6

- [tls_intercept.custom_certificate.private_key](resources--proxy--reference--group-005.md#canonical-3321230201331111-0132320001302220-0020330020122133-1310210212012031-2232300311202031-3302001322330103-0220322112132033-2212201013212221)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0203212312012200-0022331233322101-3312232012110231-2001030223033320-3211223032133022-0222102213210232-0213123221113002-2200213003332120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120123113001213-0012300123300100-0302331003202221-0213110003132223-2003230032103321-0010130100322120-2110123032221213-3311330200230232"></a>

## tls_intercept.custom_certificate.use_system_defaults — use_system_defaults / 320120113112 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- tls_intercept.custom_certificate.use_system_defaults

<a id="canonical-1110203332211003-3020032332111020-3013002222220323-1032201121213320-0022011103023313-3231220113303122-2211023121012311-2000111032121001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

<a id="canonical-0001201300311100-3133003032112211-1200123301122321-3023032132001310-2212233300231100-0013003230323233-1120213012103202-2120000123030123"></a>

## Direct properties — use_system_defaults / 320120113112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010212231113102-3132122223011220-0132331223020223-2232021003212203-2233110022033012-3013200221012023-1221031310310110-0023030330331100"></a>

## Next pages — use_system_defaults / 320120113112 / 4

- [tls_intercept.custom_certificate](resources--proxy--reference--group-005.md#canonical-2222323212312221-3303111113221022-1023133211123101-3301233201020312-0112032313303222-2100131321310321-3100110101112031-0201012311323223)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0010312210200030-1032102030200221-1013231023111111-0213001030001301-3302211232102022-2010221321321011-1320131130020100-1230311302022102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302333301332312-3330233312321122-2103113122320333-1032033030021333-3012222003322103-3100031222112112-1212231321300313-3200202010203202"></a>

## tls_intercept.enable_for_all_domains — enable_for_all_domains / 013231102210 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- tls_intercept.enable_for_all_domains

<a id="canonical-3303033331133302-0302031220133100-3101211302231333-0133120110321330-2112331230232232-2002001213311121-2122333311103001-3031221331320202"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable for all domains.

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
enable_for_all_domains = {}
```

<a id="canonical-2121020132300300-2033112010011110-3111130002310112-1223123233303102-0011010333310212-3223021321013111-3123001233023010-2210303201302103"></a>

## Direct properties — enable_for_all_domains / 013231102210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333221333333300-1313102302323222-0303201200132132-1030321020032111-1132300012132121-0030123222233201-0201321120210123-3203021002122302"></a>

## Next pages — enable_for_all_domains / 013231102210 / 4

- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-2113311103000032-0223311313330133-0101130032201233-1131032300210211-2301121231203212-0112333210002303-3232333300200313-0031201013101103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131003201201102-3112222303112120-1212132210233302-3010031011230320-1201303022222201-2010002111000013-0031022123002230-0101232313013320"></a>

## tls_intercept.policy — policy / 333102310113 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- tls_intercept.policy

<a id="canonical-3011100031221211-1203111210222303-3201101010103122-1012213130310031-1131102203013211-1200222033133233-3123020100000232-3001100100320300"></a>

Type: `"object"`. single nested block, Optional.

Policy to enable or disable TLS interception.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interception_rules")}
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
policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111123302010013-3203123211021223-1320120003013022-2102233023111102-0010301033303312-2022223131313212-3001303110010103-3303120131220130"></a>

## Direct properties — policy / 333102310113 / 3

- [interception_rules](resources--proxy--reference--group-005.md#canonical-3311202020033131-3223212112332321-1220131201010111-1332333322321321-1111201132321203-3320013320122230-3211000232330200-2310032120132032): complete subsection reference.

<a id="canonical-1100212133131002-2322020231310002-3210233133001112-2022003223030332-2120011233132002-0232332301311320-0113001122021032-2210023231300231"></a>

## Next pages — policy / 333102310113 / 4

- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-3311202020033131-3223212112332321-1220131201010111-1332333322321321-1111201132321203-3320013320122230-3211000232330200-2310032120132032)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-3311202020033131-3223212112332321-1220131201010111-1332333322321321-1111201132321203-3320013320122230-3211000232330200-2310032120132032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023100312203222-0200212322012032-3203020002013330-2110120210300132-3323000231322233-0200021222033322-3122221311011103-1102102000103211"></a>

## tls_intercept.policy.interception_rules — interception_rules / 010021010033 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-2113311103000032-0223311313330133-0101130032201233-1131032300210211-2301121231203212-0112333210002303-3232333300200313-0031201013101103)
- tls_intercept.policy.interception_rules

<a id="canonical-3312120023110110-1021033230023313-3012032301110231-0312332203030120-1302020330230020-2013110232000331-3202313011323212-0202011032302323"></a>

Type: `"object"`. list nested block, Optional.

List of ordered rules to enable or disable for TLS interception.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("disable_interception",
    "enable_interception")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interception_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1033020320013100-0110201103121333-0301200101313122-1021213122301012-0111031102323100-1331221222003110-0323120313303332-0210200312031302"></a>

## Direct properties — interception_rules / 010021010033 / 3

- [disable_interception](resources--proxy--reference--group-005.md#canonical-1000112331102113-3212211312313230-0302331033011221-1311123131101303-0211200322130033-1332120220320010-2313223123033301-0312113131110321): complete subsection reference.

- [domain_match](resources--proxy--reference--group-005.md#canonical-1023120230321222-1033211013013300-0023103003301010-2100102002020300-0332210232003222-1301032200113301-3222220323213330-2221221100213320): complete subsection reference.

- [enable_interception](resources--proxy--reference--group-005.md#canonical-1123312222230112-2211100331321323-3022231331031311-1131132221023111-2301031013132333-3201220202231311-2310211212210221-2330130331323033): complete subsection reference.

<a id="canonical-3310012032313220-3222033201220322-0030120032313323-1312230222331222-1320301313202021-0003321311002202-3302301032232022-3120021300102232"></a>

## Next pages — interception_rules / 010021010033 / 4

- [tls_intercept.policy.interception_rules.disable_interception](resources--proxy--reference--group-005.md#canonical-1000112331102113-3212211312313230-0302331033011221-1311123131101303-0211200322130033-1332120220320010-2313223123033301-0312113131110321)
- [tls_intercept.policy.interception_rules.domain_match](resources--proxy--reference--group-005.md#canonical-1023120230321222-1033211013013300-0023103003301010-2100102002020300-0332210232003222-1301032200113301-3222220323213330-2221221100213320)
- [tls_intercept.policy.interception_rules.enable_interception](resources--proxy--reference--group-005.md#canonical-1123312222230112-2211100331321323-3022231331031311-1131132221023111-2301031013132333-3201220202231311-2310211212210221-2330130331323033)
- [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-2113311103000032-0223311313330133-0101130032201233-1131032300210211-2301121231203212-0112333210002303-3232333300200313-0031201013101103)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1000112331102113-3212211312313230-0302331033011221-1311123131101303-0211200322130033-1332120220320010-2313223123033301-0312113131110321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202231113113033-2320012011131011-2212102200220121-2121230221321211-2023113320232003-2011122233321233-2111112200122330-2002100303131220"></a>

## tls_intercept.policy.interception_rules.disable_interception — disable_interception / 000112000313 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-2113311103000032-0223311313330133-0101130032201233-1131032300210211-2301121231203212-0112333210002303-3232333300200313-0031201013101103)
- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-3311202020033131-3223212112332321-1220131201010111-1332333322321321-1111201132321203-3320013320122230-3211000232330200-2310032120132032)
- tls_intercept.policy.interception_rules.disable_interception

<a id="canonical-0131220000203202-2012032110322310-1130312113203001-3202200031223021-2111112001303320-3031213121103212-2212320133201212-1022020211033011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable interception.

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
disable_interception = {}
```

<a id="canonical-2233020231001103-2021323330101313-1111330332220213-0310110203112223-1302003332022333-0230010123022100-3222222201102210-2320000301101203"></a>

## Direct properties — disable_interception / 000112000313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103121231122033-2131311123322333-3330101021223303-0020233221330310-2222322010013012-3132301332331233-0131332301110003-2210022011001001"></a>

## Next pages — disable_interception / 000112000313 / 4

- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-3311202020033131-3223212112332321-1220131201010111-1332333322321321-1111201132321203-3320013320122230-3211000232330200-2310032120132032)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1023120230321222-1033211013013300-0023103003301010-2100102002020300-0332210232003222-1301032200113301-3222220323213330-2221221100213320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322000310032110-2031000322100331-2031010121030003-2301330011203100-3011003321132131-3303212123311203-1023322233103222-2303300300113101"></a>

## tls_intercept.policy.interception_rules.domain_match — domain_match / 232011203031 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-2113311103000032-0223311313330133-0101130032201233-1131032300210211-2301121231203212-0112333210002303-3232333300200313-0031201013101103)
- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-3311202020033131-3223212112332321-1220131201010111-1332333322321321-1111201132321203-3320013320122230-3211000232330200-2310032120132032)
- tls_intercept.policy.interception_rules.domain_match

<a id="canonical-0222232012012321-1012022110323022-2130020212210131-3222232100323101-0110033332333101-1213131322200231-2331333113132111-1121000230312002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for domain match.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain_match {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303012203023011-3201120223112203-1132301213232311-3231231130332222-3001023012202130-2230302230133022-0123021111333020-1303330221123201"></a>

## Direct properties — domain_match / 232011203031 / 3

<a id="canonical-3201000310002110-3111121001031331-1220221031121301-1313022123030033-0020000223300020-3023033002213233-3030301112233131-0123033223222203"></a>

<a id="canonical-2322121323313103-2221321022213030-3001101323311212-0203322133223123-0031301231022103-0103310132102313-1222110211130322-3312121230231002"></a>

## exact_value property — domain_match / 232011203031 / 4

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1232322302030023-2132213220131130-0032303303133001-2311010030202200-1301230211311323-2222010132121103-1023231021323013-2031332131031030"></a>

<a id="canonical-1201033200031220-3010011301222223-2010203200022003-1111200201010133-3131203232123223-2002012222123222-1231312131103130-0312021332210330"></a>

## regex_value property — domain_match / 232011203031 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0230330021220112-1120022222202102-1300102121033322-0231001100213030-1111320121100221-3120221233230210-0011102121321221-3002121303032233"></a>

<a id="canonical-3102133100013100-1030230103131121-1001320210102310-2222300101302310-2212000232031132-2203001003303033-1030023231010130-0123123023222303"></a>

## suffix_value property — domain_match / 232011203031 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3103332222002001-0112022303120323-1133102211322033-3112010333003301-3220011310231031-0321312332201201-1011311333110030-3210121232321223"></a>

## Next pages — domain_match / 232011203031 / 7

- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-3311202020033131-3223212112332321-1220131201010111-1332333322321321-1111201132321203-3320013320122230-3211000232330200-2310032120132032)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-1123312222230112-2211100331321323-3022231331031311-1131132221023111-2301031013132333-3201220202231311-2310211212210221-2330130331323033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113333033021222-2012332232013130-0210322110011201-2132323320312233-2233302101031103-0233332101203210-1003012320011003-2100302001331133"></a>

## tls_intercept.policy.interception_rules.enable_interception — enable_interception / 021000212121 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [tls_intercept.policy](resources--proxy--reference--group-005.md#canonical-2113311103000032-0223311313330133-0101130032201233-1131032300210211-2301121231203212-0112333210002303-3232333300200313-0031201013101103)
- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-3311202020033131-3223212112332321-1220131201010111-1332333322321321-1111201132321203-3320013320122230-3211000232330200-2310032120132032)
- tls_intercept.policy.interception_rules.enable_interception

<a id="canonical-3013000203331130-3003021302221020-2112333323201001-0011311321113103-0002221001030022-2312213230213010-0312201033202321-3231302301120002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable interception.

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
enable_interception = {}
```

<a id="canonical-0022131120333001-1002333203023002-3020001120223202-3021022322010133-1031101013220202-0212010113123210-2203201303230020-3220033100313111"></a>

## Direct properties — enable_interception / 021000212121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010322321302222-0013100203132102-2021001200201310-0322231001120133-0100220030021120-3212233322130130-2021300110130103-2310122311230122"></a>

## Next pages — enable_interception / 021000212121 / 4

- [tls_intercept.policy.interception_rules](resources--proxy--reference--group-005.md#canonical-3311202020033131-3223212112332321-1220131201010111-1332333322321321-1111201132321203-3320013320122230-3211000232330200-2310032120132032)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0322201212300210-3122231010311203-0031222013232003-3121000333001303-2020203321212103-0110333232012202-2000012101030312-2210122003113011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131032133000003-0023201311001023-2000301002210000-3020223001323032-2332320002112203-3033020213020212-1223110023031132-0321010323122201"></a>

## tls_intercept.volterra_certificate — volterra_certificate / 111033133012 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- tls_intercept.volterra_certificate

<a id="canonical-2121122002311101-2001321331222012-1332131220210100-0323202103032100-2011030332222001-3112331133023222-1302120010030002-0311103030022013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra certificate.

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
volterra_certificate = {}
```

<a id="canonical-2113101113213311-0320203022312333-0032102000211100-1303222213302112-0130221232110201-3322030102221022-1312113202201133-3012032131322313"></a>

## Direct properties — volterra_certificate / 111033133012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331220122103301-1232203211132131-3331313230033200-2311100003113211-0103002202213233-2110031231311220-3033121002303333-0133312301130323"></a>

## Next pages — volterra_certificate / 111033133012 / 4

- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)

<a id="canonical-0221120012112232-3222102003023211-2332123231332121-1022321130233300-3132233121322211-0210112032300212-1100120323003230-3333123211101001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000132020202213-2301321322310003-1101112033133101-3013130022001212-1323130133230301-0113222220332010-3231012201103113-3212122102121121"></a>

## tls_intercept.volterra_trusted_ca — volterra_trusted_ca / 303323101032 / 2

Breadcrumbs:

- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
- [Property reference](resources--proxy--reference--group-001.md#canonical-1332312330112202-0302233320331232-2000013322011121-0003103202230221-3121132031010303-2222021012203301-0203131033201011-3322301300211003)
- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- tls_intercept.volterra_trusted_ca

<a id="canonical-1013120230302223-0033003022013000-2223010222321303-2021123331333122-1122330221121323-1302122111221013-3302102203012022-0003002030322113"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra trusted ca.

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
volterra_trusted_ca = {}
```

<a id="canonical-3222010011310313-1223000200103023-0231011121201021-2321001002201020-2203112131100300-1301301311030221-0000330313301303-0000132233330222"></a>

## Direct properties — volterra_trusted_ca / 303323101032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210123100130101-0121013320033223-3233123222033032-0301022122122221-3023122012320132-2033112102211331-3322212130122130-0330313230331232"></a>

## Next pages — volterra_trusted_ca / 303323101032 / 4

- [tls_intercept](resources--proxy--reference--group-005.md#canonical-1012032132320210-0010122100332331-0000321130320013-3133303110113120-0200133111003231-2120031011211003-1123012303313320-0010122102231031)
- [xcsh_proxy](../resources/proxy.md#canonical-0132300111013030-3233132200202310-2233332113220102-2031222100321320-0203021302013323-3111102110130123-3333011301331210-2310111332313332)
