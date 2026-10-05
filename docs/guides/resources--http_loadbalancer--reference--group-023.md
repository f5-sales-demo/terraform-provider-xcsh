---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2122232220022111-2313330200023223-2002212301312012-0233110300122110-1301102030012312-3332321332320002-3130223121102311-2100013200220030"></a>

## policy_based_challenge.rule_list.rules.spec.asn_list — asn_list / 121310311303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.asn_list

<a id="canonical-3313120122012310-1312303301112131-0330231303213310-2123203000010211-1300220010012002-1202311122313213-0122023202310323-1323200031333331"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213003110022222-3132103211211311-1333322022321322-3022222220012002-1222100332311022-1223103210001131-1020021121220322-1120303302312032"></a>

## Direct properties — asn_list / 121310311303 / 3

<a id="canonical-2011021001231323-2000131220333220-1213030231001021-2222211233332011-3200203011230231-1032303202201010-3033311321013300-3033212211221031"></a>

<a id="canonical-1001200020333010-3113233203232001-2023101320033000-2201333011232113-1020333010032002-0122331200132201-0001021102003001-1303013022132301"></a>

## as_numbers property — asn_list / 121310311303 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

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

<a id="canonical-3103311200320231-3201300033130002-0311110313110013-1210323331320130-0100220332011321-0303303122311020-0032210332011030-0331100330120332"></a>

## Next pages — asn_list / 121310311303 / 5

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2330232220313013-2120000211302223-3012212202222313-1113232301013021-3231202321121233-0332211110312023-1013001211321310-2012322132110120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131123003101103-1001130313231203-0302303201002010-3022100130331111-2133311213231233-1033310021312202-3013012121222013-3332222022130311"></a>

## policy_based_challenge.rule_list.rules.spec.asn_matcher — asn_matcher / 210001223123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.asn_matcher

<a id="canonical-3332020131111132-3123312111131233-0221120323023112-0032020302001230-1123232131010332-2012320311231101-1302113001110312-3223032100201333"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231223333033033-3303102333110111-3033120210100120-2211122020101103-0333122220201113-2030222331322030-0301001201200300-1123002102031022"></a>

## Direct properties — asn_matcher / 210001223123 / 3

- [asn_sets](resources--http_loadbalancer--reference--group-023.md#canonical-0331010323003123-2200312103010033-1210131010130000-2103022000112201-3321121120330130-1022200100002313-0113122201303213-0203132322011021): complete subsection reference.

<a id="canonical-1333123310303111-0211231033000302-3030213023233320-0001120123323322-0211031332012220-0013201330102332-1120020203232000-3320103210212330"></a>

## Next pages — asn_matcher / 210001223123 / 4

- [policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets](resources--http_loadbalancer--reference--group-023.md#canonical-0331010323003123-2200312103010033-1210131010130000-2103022000112201-3321121120330130-1022200100002313-0113122201303213-0203132322011021)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0331010323003123-2200312103010033-1210131010130000-2103022000112201-3321121120330130-1022200100002313-0113122201303213-0203132322011021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332102123323020-3310120000112032-2113123312312323-2101132113131211-0033212110333232-0231330122031101-1111023020011202-2203333010232331"></a>

## policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets — asn_sets / 301131010220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--http_loadbalancer--reference--group-023.md#canonical-2330232220313013-2120000211302223-3012212202222313-1113232301013021-3231202321121233-0332211110312023-1013001211321310-2012322132110120)
- policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets

<a id="canonical-0013021303033210-1310033201100322-3122233210301123-1330330222000221-3003222212323023-1000102221322202-0121222022133110-0330112311003203"></a>

Type: `"object"`. list nested block, Optional.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302113112310220-3121121123332011-3311210022200113-2001102001133302-2233123313303221-1212300101301000-1200200123011322-1222302310220213"></a>

## Direct properties — asn_sets / 301131010220 / 3

<a id="canonical-0320221303031113-1201232302222013-2220013000231333-0231101013122131-1232213201030302-3013222323303001-0021302230012310-0332321322121202"></a>

<a id="canonical-3320122012320132-0202320012122131-0222311000000121-3012332300300322-3010221020002101-3121232001032132-1023113201302313-1321301020212313"></a>

## kind property — asn_sets / 301131010220 / 4

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

<a id="canonical-0230311112202302-2301012301101323-1311131202301211-2200212332113200-1313003010320113-0013223202321213-1112332101330033-0032222300101331"></a>

<a id="canonical-0300213200120212-0123302231201203-3233312031002311-0211102003111120-1311030013032310-2202000211131023-0303111203212021-0300131330131301"></a>

## name property — asn_sets / 301131010220 / 5

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

<a id="canonical-3212100010123103-2201132023213020-1211212303303201-2123031302211020-2013030122021222-1201023120030321-3000022303233000-3323331201010001"></a>

<a id="canonical-2312012303230012-1311213303330030-0211111103303233-1220321213132132-0133023121122111-1211011010311111-1223022030120032-2302032203313023"></a>

## namespace property — asn_sets / 301131010220 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1220333021333210-3132101033202303-3103321130213330-3233000200202023-0123113103103333-1200202320212323-2130233111210101-2330331333021011"></a>

<a id="canonical-2131000323030221-0011320300231300-1213331113030022-2010300102320112-0023200112021311-2013113302012102-0012023000103002-0231212131023330"></a>

## tenant property — asn_sets / 301131010220 / 7

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

<a id="canonical-0212030212030313-1320332202232022-1320023311300203-1131000132123221-1222323212112202-0001200211302323-0233231320133300-0300232321131022"></a>

<a id="canonical-0100113313330001-2303011300320123-0030202232123221-1333301203011201-1131300321313131-3123303232112222-3111310211022222-1202120023022233"></a>

## uid property — asn_sets / 301131010220 / 8

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

<a id="canonical-3231303300303020-0111003331202311-3113303011120021-2031113323121111-1112122312101300-2031121010022022-2123232002122211-1123001313102131"></a>

## Next pages — asn_sets / 301131010220 / 9

- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--http_loadbalancer--reference--group-023.md#canonical-2330232220313013-2120000211302223-3012212202222313-1113232301013021-3231202321121233-0332211110312023-1013001211321310-2012322132110120)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1322031002303310-1111010110113010-3010000120211223-0112031202323033-3213001011000120-0032032230230101-3220223203111233-0023030020301031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300031320131112-3002112103211213-1312300331303122-0333230213013133-3323201332002202-3023103313131333-2330122111012121-2131113320221102"></a>

## policy_based_challenge.rule_list.rules.spec.body_matcher — body_matcher / 302103232210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.body_matcher

<a id="canonical-1120202020002120-0033002221031111-1101231111313331-0021301303031310-2020011020322002-3003230313003111-2203330013030122-0031120312000232"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
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
body_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-1223231322033311-1030001011232201-2121132101000131-2333313100323020-3020002333233102-1302123222012323-0022323023101333-2002300303030131"></a>

## Direct properties — body_matcher / 302103232210 / 3

<a id="canonical-1222110331201200-0100012130110120-0311010032103212-2031202223133001-3231120223330201-0232332132131331-0103113100303210-3212301233203210"></a>

<a id="canonical-0312111021133203-2210301111312123-0101311333122201-0203010133211100-0320301222332031-2302301312223110-3222301232330101-0330212010103302"></a>

## exact_values property — body_matcher / 302103232210 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2101100123321313-3131330133110131-1011331311203301-2211111320222320-3130322230311101-1210230132203313-3132331112012122-2220132303230232"></a>

<a id="canonical-1220331200112020-0212113320310312-0033123200111112-2002300011120023-0212333020231033-2220301201222111-2101310310333200-3232131213310112"></a>

## regex_values property — body_matcher / 302103232210 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1212100332130320-1012101311113323-1300030002112003-2133222013231000-3333013213332003-3033311011131030-0211010201021202-3333323210232100"></a>

<a id="canonical-3131232011122321-3023223123220210-3132313013110113-3331033031212221-2210201030221121-2012121333303110-3221032312130100-0002203001213022"></a>

## transformers property — body_matcher / 302103232210 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1223301133120001-1310033232302000-1032012332033100-3220210113312221-0121113320013313-2312212203301000-1132132013122103-1300101033233210"></a>

## Next pages — body_matcher / 302103232210 / 7

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3301001033202231-2133133323213110-2321233202103301-1310201321100023-1302201121201002-3200232221023023-0110031303001223-0212313222223133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003111203100113-2100031022032131-2303230131013220-2223113232220002-2321003020230210-3003101202312112-2122203331202031-2331101203001132"></a>

## policy_based_challenge.rule_list.rules.spec.client_selector — client_selector / 030212222223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.client_selector

<a id="canonical-0121031210120212-2200023300010320-1303230332213032-2021223200323330-3021322303321000-3022022012331003-0202202311020323-2011303233000013"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121333023321010-1200211301331303-1211011020022210-1313203130113213-1130123011013112-2303023113203011-2021323010302131-2333120231101011"></a>

## Direct properties — client_selector / 030212222223 / 3

<a id="canonical-0302312133102010-0010310201312322-2202231331120321-0223332301122022-1230212203132110-0022023102313323-3201133310001231-1223002021111220"></a>

<a id="canonical-3203102102301011-0210220112120022-0010010323113231-3203323033212230-0210001330221330-1212231011321332-1301321133012013-3220011200203303"></a>

## expressions property — client_selector / 030212222223 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2032121321011322-0211323130102200-0103233011103310-0012000111030012-0001230312233201-2201201310030121-3012022112122211-0323301313322321"></a>

## Next pages — client_selector / 030212222223 / 5

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212230230323221-2130110112023221-0120331300331232-1001103221301133-3010102222223120-3232211311030100-0222312002030101-0111220332200213"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers — cookie_matchers / 102123101020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers

<a id="canonical-2311302301030200-3113330322112312-3320022100223123-0111322103012232-3103312312332003-3310211202231110-2203000200033132-2003313132233310"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-3002210022200332-2222312020011021-0332020211113222-3212033201111133-3213101231230110-3222031333222202-3312333232032233-0022013233230133"></a>

## Direct properties — cookie_matchers / 102123101020 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-023.md#canonical-0103313330202032-0333323222300211-3231010121102231-1022300331232331-1311002322210132-3210201002132031-2200001313310231-1132120002030213): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-023.md#canonical-1011011121103030-1023321311310333-2210310200001020-2031223022321101-0211112103002331-0113300012011220-0213102312210012-3331322211331322): complete subsection reference.

<a id="canonical-2320031132013223-2200212003011211-1122112213300011-3222201101010102-1323020003210220-3100132203131102-2000300332231321-3032220311311221"></a>

<a id="canonical-3230203020331111-3113323233103233-0311121130001332-3320010001133222-3331111331133323-0320012020130333-1100210211221310-2313101122303120"></a>

## invert_matcher property — cookie_matchers / 102123101020 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--http_loadbalancer--reference--group-023.md#canonical-2301120303000033-1120110232011011-0311033303003031-3013101323203110-3103223000001020-2012122203222232-3020031332212020-3132031221321231): complete subsection reference.

<a id="canonical-2320220323220232-0222321220003111-1032113000013020-1302313312110003-3012331330201301-3333220103101322-3232200232133112-1121332003000322"></a>

<a id="canonical-0302003201223330-3332013212110202-0103002320220302-1133301320110101-3101232302312212-1301301003032203-0213111110220022-1323111230023313"></a>

## name property — cookie_matchers / 102123101020 / 5

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3101323021001120-0301133101333312-3321222102312020-2313023331103120-0323220233121210-3231331232131013-0012222122102120-1230133102022200"></a>

## Next pages — cookie_matchers / 102123101020 / 6

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present](resources--http_loadbalancer--reference--group-023.md#canonical-0103313330202032-0333323222300211-3231010121102231-1022300331232331-1311002322210132-3210201002132031-2200001313310231-1132120002030213)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present](resources--http_loadbalancer--reference--group-023.md#canonical-1011011121103030-1023321311310333-2210310200001020-2031223022321101-0211112103002331-0113300012011220-0213102312210012-3331322211331322)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.item](resources--http_loadbalancer--reference--group-023.md#canonical-2301120303000033-1120110232011011-0311033303003031-3013101323203110-3103223000001020-2012122203222232-3020031332212020-3132031221321231)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0103313330202032-0333323222300211-3231010121102231-1022300331232331-1311002322210132-3210201002132031-2200001313310231-1132120002030213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020212322333231-1221131302310302-0010313033110130-3032131222323111-2130021001330120-0102011212001102-0221211132301000-1330212320202332"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present — check_not_present / 022212002032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-023.md#canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present

<a id="canonical-1301310221031133-3232011221111032-2201210210121031-1210020202020322-1203132333021121-0032111003223002-1203222123220112-3111022203232100"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-1213003232320000-1323133322322023-1213030102111221-1033123202110212-3103333310013203-3000300111012113-2233111300013302-0331003220122001"></a>

## Direct properties — check_not_present / 022212002032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212330312300003-3101112110323102-0230110201321122-1013110212332123-1031220013232313-1303211123033330-0023210102011320-1331231300233312"></a>

## Next pages — check_not_present / 022212002032 / 4

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-023.md#canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1011011121103030-1023321311310333-2210310200001020-2031223022321101-0211112103002331-0113300012011220-0213102312210012-3331322211331322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123011203312101-2221100013300001-2110122030110203-2123121100001132-2003312311323033-3021313213122021-0002310020331210-1131122133222002"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present — check_present / 302210103030 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-023.md#canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present

<a id="canonical-2122002112213001-1220221103323200-3311212230131022-3010330110220012-2300122233332000-1300013310123211-1000320321023032-1223110132111310"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-3013032100132013-1220100020101023-0100230112103211-1012022310200230-3120322000231303-0102202302301210-1132302130123210-1001333230210003"></a>

## Direct properties — check_present / 302210103030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110002000120021-0231300210120223-2022300310323031-1031321223122230-0332030011321000-3012200301032120-2332020001133222-0302003131123313"></a>

## Next pages — check_present / 302210103030 / 4

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-023.md#canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2301120303000033-1120110232011011-0311033303003031-3013101323203110-3103223000001020-2012122203222232-3020031332212020-3132031221321231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310230223303231-3123232210103201-2330121313033000-0133213022202002-1023022311100332-1313003031312333-2000231111203212-0203231223131111"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.item — item / 320200133012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-023.md#canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.item

<a id="canonical-0233310200121001-1010031310210020-1002330202312030-0112101200233223-2100130121231000-0033203200213032-0002230032123233-0002301313312010"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-1033301233101110-3223213110013002-2103113302301020-0212212223230303-1333202300223123-2130211133302030-2223313020102131-1210011020233233"></a>

## Direct properties — item / 320200133012 / 3

<a id="canonical-0211321132033130-1330031123321102-1102111320112213-3222023003021013-2333321132233100-3123320101220031-1100321100123210-3210203310010021"></a>

<a id="canonical-1022120132220033-2023302023023331-3011331233033112-2213102232021201-3312223211220020-2031330221131101-1130111201201020-0013021200310033"></a>

## exact_values property — item / 320200133012 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3132130322300322-2211222320230323-0232220200002330-2132033130013202-3320003301322210-2122030300203032-1130233112103121-0133012122111020"></a>

<a id="canonical-0320230123203210-2303301110102023-3200321223301000-1200033333033031-0322033111113002-2111030121211130-3232100310102002-2130120222002121"></a>

## regex_values property — item / 320200133012 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1033022011011132-0030003332130331-2231313220102023-1302201010311232-3121220103132032-2211210132311132-2332233333322033-0321003333232312"></a>

<a id="canonical-1303322122203120-3202101300003121-3320210312211333-2330302303300221-1120102031212230-2012122322002310-3210320233002031-0123111032331212"></a>

## transformers property — item / 320200133012 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3323101312312120-0102123211300310-1011132212031311-1000230211301103-3320102002331211-3122211310022322-0201320213321301-2112210032031220"></a>

## Next pages — item / 320200133012 / 7

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--reference--group-023.md#canonical-1112223030321202-2310131032330100-2302103332230131-0001200330203221-2212303302110122-3033121110232100-1130300130302113-3232330311033030)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0202302323030031-0332332132022130-1301100023003233-2302201010011300-1300320313311232-3332303233323013-2021332220011213-0032313121021322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131320221203300-3100122130133101-2333131303121313-3320003310020311-0112213123221000-1123323323230313-2311223200120213-0113013302320020"></a>

## policy_based_challenge.rule_list.rules.spec.disable_challenge — disable_challenge / 230121232300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.disable_challenge

<a id="canonical-2303302200301023-0230123210232312-3110113233311112-2211230302131132-3222201213233310-3211011001032212-2100113110000000-1021122133032231"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable challenge.

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
disable_challenge = {}
```

<a id="canonical-1202101121123213-1211020203133022-1131132301321002-1212213313233120-0023011201030130-2131331031100133-1211323122231331-3232223100320323"></a>

## Direct properties — disable_challenge / 230121232300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0031132011110112-3222333300032130-2000213032000132-3011231220021303-1132202331333212-3200031033130313-0321023103113202-1131231011203113"></a>

## Next pages — disable_challenge / 230121232300 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2231322321300130-2211003200213110-0230120220232310-0330031202332122-1202330333221203-2022232030123301-0132023201301310-1122331113003030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301200211320232-1120113123120033-3011002003133223-1311033112021233-0312010303020113-1002233031132023-3323002333313301-2001310032111021"></a>

## policy_based_challenge.rule_list.rules.spec.domain_matcher — domain_matcher / 012103203230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.domain_matcher

<a id="canonical-1232323002333033-0112032300221013-1133320131331310-0103302201222000-3001320023012200-2112013212022221-0032101200212311-2212321012220203"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
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
domain_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130000102131321-2203131013303102-2330213000003203-2121011030102202-2312310211311121-0013021303123231-0120023113312203-2011203233211112"></a>

## Direct properties — domain_matcher / 012103203230 / 3

<a id="canonical-0013211133032103-3232013113002302-3230113022221032-1132211331203001-3113023233023301-0103120223333213-1020212133232220-3010130102010320"></a>

<a id="canonical-0322020003220322-1002011230211021-1100000201103020-2111033210131211-3121222302221322-3323330111121300-2012223110031211-3231222110310000"></a>

## exact_values property — domain_matcher / 012103203230 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2202110300322313-2013321002200231-1311231313202100-0132003012202131-1122113020331011-0300000223013202-3213010220210303-3213212130020003"></a>

<a id="canonical-1313212111221310-1032223223132220-0320111233303301-3123213231300312-2211033011332332-1300203002211311-3032211020311013-0023001010103310"></a>

## regex_values property — domain_matcher / 012103203230 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0120033000230122-3332312013122302-2110321233301011-3122222032111323-2302010110003331-1013233202020121-2010110010230122-1330132332120202"></a>

## Next pages — domain_matcher / 012103203230 / 6

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2331103200120201-0301001021002322-0132233020221032-2203102002320022-3001110313102332-0311012113002330-0202002031220211-1023110131120121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323100100320203-2223020323130203-2323113012003320-0213113102210000-0000133332313002-1102021020101210-2112012021323112-3110100001333121"></a>

## policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge — enable_captcha_challenge / 131201203021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge

<a id="canonical-2132331203310320-1222200200202033-0011330110111031-0312132201102202-0013320113013122-3102310210001133-1120212312030103-0100322013013022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable captcha challenge.

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
enable_captcha_challenge = {}
```

<a id="canonical-1220113233210110-1330300032103311-3302331321231012-3023110223222131-3223013121033230-0010123010312001-1031302031100330-0101120323020122"></a>

## Direct properties — enable_captcha_challenge / 131201203021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033303222121132-3032220012033330-1330122223311033-0032103000111312-2203312333322211-2313221333221101-3132230323222022-1312333210320313"></a>

## Next pages — enable_captcha_challenge / 131201203021 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2221112113201102-3022103132211101-1320220303332222-2303330112030301-1020203320020311-2100330302222210-1012001232303023-3130312322023023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012133001131331-1030331020103030-3212313231311202-3003211101331122-0133310023223221-1311030200330331-0010001101300132-0102133121021023"></a>

## policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge — enable_javascript_challenge / 102122221321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge

<a id="canonical-1130202110233103-2131102003313220-3030212130333030-2001123211330000-1012230010300023-2231123011102202-3222120333303323-1012211201202020"></a>

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
enable_javascript_challenge = {}
```

<a id="canonical-3311022320303102-2102032001011211-3101310001100322-1100102201310203-2133213123333322-0320332233231331-3331322023003100-1331223133312303"></a>

## Direct properties — enable_javascript_challenge / 102122221321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030030333001200-0110103211331121-2302223002220013-2231331111333221-3202032323300302-1213232000103133-3121022333100033-3103302331332023"></a>

## Next pages — enable_javascript_challenge / 102122221321 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332133313121102-0030220013123001-3301211213312110-0012131033113322-0101323000121030-1313332032100231-0301213222200002-0223130120112132"></a>

## policy_based_challenge.rule_list.rules.spec.headers — headers / 233003121231 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.headers

<a id="canonical-3213121202022121-0002321330203123-0233123032313100-2330323312031012-3330313002033230-0330302002020002-2201311111210003-0302110231212101"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000021210212322-2221202301011212-3323301220003110-3023311102021310-2210033303112100-0210202222132233-3200121201230330-0101030031012210"></a>

## Direct properties — headers / 233003121231 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-023.md#canonical-2323020220021030-1221122111120010-1000002310111301-3103132212310231-3222000202211323-2223022323033322-3200012012111322-2010121000223113): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-023.md#canonical-1211303211131310-1223121301221223-0331003023331300-3013201222033021-3213232122201131-1201322000101313-3231203021010120-2301213312132013): complete subsection reference.

<a id="canonical-1022001022301210-0301201121032130-1210031220100332-0233100233200020-0212020022312231-2111301031301200-3110000333023133-0130122100020212"></a>

<a id="canonical-2111211221322210-0023111231130122-3032203032030000-1012231210030220-0001011003332301-1112202233312322-1212302021033100-0030311222130323"></a>

## invert_matcher property — headers / 233003121231 / 4

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--http_loadbalancer--reference--group-023.md#canonical-3000212303030230-2323101000310303-2133302302323120-1302012010122311-3230232013302223-2221311321310131-0313233101301123-0123020121000020): complete subsection reference.

<a id="canonical-0223111223131203-3012133001321010-0203303210331200-0332022302323220-2010233130030112-3211000312013102-1331130102132110-2202310222100323"></a>

<a id="canonical-2100013300113302-2232032231020333-3002220100302031-1120213231213321-0310222131303123-3100010012102122-1301001023013000-1310002021312030"></a>

## name property — headers / 233003121231 / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0301023212033303-0321103112122130-1202132212013032-2310031210013001-3120223023030010-0000022313300223-3123133000111131-0212121300002321"></a>

## Next pages — headers / 233003121231 / 6

- [policy_based_challenge.rule_list.rules.spec.headers.check_not_present](resources--http_loadbalancer--reference--group-023.md#canonical-2323020220021030-1221122111120010-1000002310111301-3103132212310231-3222000202211323-2223022323033322-3200012012111322-2010121000223113)
- [policy_based_challenge.rule_list.rules.spec.headers.check_present](resources--http_loadbalancer--reference--group-023.md#canonical-1211303211131310-1223121301221223-0331003023331300-3013201222033021-3213232122201131-1201322000101313-3231203021010120-2301213312132013)
- [policy_based_challenge.rule_list.rules.spec.headers.item](resources--http_loadbalancer--reference--group-023.md#canonical-3000212303030230-2323101000310303-2133302302323120-1302012010122311-3230232013302223-2221311321310131-0313233101301123-0123020121000020)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2323020220021030-1221122111120010-1000002310111301-3103132212310231-3222000202211323-2223022323033322-3200012012111322-2010121000223113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220222033023322-0211011231132333-3203332131332000-1010021013213003-3112202031021013-3212001001222303-0010133020300323-1133213001211020"></a>

## policy_based_challenge.rule_list.rules.spec.headers.check_not_present — check_not_present / 131213031010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-023.md#canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111)
- policy_based_challenge.rule_list.rules.spec.headers.check_not_present

<a id="canonical-0002300313103112-3301102110122300-3301202122012112-1033233202110120-0000213133220003-3202030031123301-2123031230210031-3211322121300300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-3013033223123323-0133103201113321-2020122121001020-2131331111132311-2000030110023200-0222101100303123-1132102012313231-2121322221112300"></a>

## Direct properties — check_not_present / 131213031010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302112221013223-3222323311112222-3213222321030030-3020130111103202-0201010032310011-0301003111330031-1221112312300112-1211232233100031"></a>

## Next pages — check_not_present / 131213031010 / 4

- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-023.md#canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1211303211131310-1223121301221223-0331003023331300-3013201222033021-3213232122201131-1201322000101313-3231203021010120-2301213312132013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003322323221011-1111132102303000-1222102310022013-1103113003121332-0113212032111101-1330022231223111-3330321113301002-2231322310221032"></a>

## policy_based_challenge.rule_list.rules.spec.headers.check_present — check_present / 300003313023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-023.md#canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111)
- policy_based_challenge.rule_list.rules.spec.headers.check_present

<a id="canonical-1321330332113011-2101112110100033-3111113123330110-1121210322220332-3120030323110122-3002022122200332-3323113222112222-0331332222211122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-3333003020131232-1132333002230100-1232130012202022-3301012100332013-1133201331131013-3202332000121311-0312322333231112-2202103213123322"></a>

## Direct properties — check_present / 300003313023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302212010013000-0020233331213023-1130113021013320-0032131333102222-0131230111033330-1221200132223030-0300320313303223-2311200012222030"></a>

## Next pages — check_present / 300003313023 / 4

- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-023.md#canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3000212303030230-2323101000310303-2133302302323120-1302012010122311-3230232013302223-2221311321310131-0313233101301123-0123020121000020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310101323121001-0211110331100322-1312012100130231-2010100321320303-0110122012120201-2230301111233330-0330121003121120-2223212011222032"></a>

## policy_based_challenge.rule_list.rules.spec.headers.item — item / 201011202132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-023.md#canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111)
- policy_based_challenge.rule_list.rules.spec.headers.item

<a id="canonical-0111123330021001-3132230101323211-0013032210303001-1321303020130103-3312300320201001-0213013021220112-1003300220322122-0231003310002013"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303111321112133-0211133322230330-0103333121333212-0130011231113221-0012032233212201-3320300220220032-0332002102113230-1033021332202233"></a>

## Direct properties — item / 201011202132 / 3

<a id="canonical-0333223002223333-1030033220212320-0230212212132010-0113330322320013-3211112120311110-2300121330023322-2213332311201021-3221110110100001"></a>

<a id="canonical-0300100100011322-3323230113201313-2120133313021112-1110201303331312-3013213111332323-3330312010121031-2321113003001333-2022321330121331"></a>

## exact_values property — item / 201011202132 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1121113112100010-1230120030331313-1122012133110102-1033030322321112-2332310033300322-1230323031111201-2123301113133002-2333321220010311"></a>

<a id="canonical-1030330120231211-2110311130000003-1002021013320032-1030301302100120-3211011203011333-2333021311310112-3301022200231231-3133323002010102"></a>

## regex_values property — item / 201011202132 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1211321000113030-1230333002310302-1211233023011231-3322213033113020-0231021023202210-0010303132011122-2021323311133122-2130123112012231"></a>

<a id="canonical-3021133000111022-3022320223223033-0013021121122132-2330320302032001-0111312212220221-3130230001300013-3130102223103201-2021122202310201"></a>

## transformers property — item / 201011202132 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2310232122011213-0113123220002233-1112032112313213-1221022013232101-1322230221112020-2022030113123123-2203233221131020-1112331233301302"></a>

## Next pages — item / 201011202132 / 7

- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--reference--group-023.md#canonical-1313201031011033-0203021101211130-2233102100210202-2231200332022130-0212312030111003-0030330021121330-0011012123121011-2011131311103111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2330011102111203-3200210100131120-2112200113232323-3120021212132232-3203312330132100-2321330223303101-1023321200013301-2303010303011323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302310112031223-3010121020010212-1100133110333102-1033122102232203-0331203313322030-3102331103212303-0323320313002300-2112213220132102"></a>

## policy_based_challenge.rule_list.rules.spec.http_method — http_method / 111223223310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.http_method

<a id="canonical-0320200311032022-0313303221002101-1020100111313001-1023123212203203-2012111201110121-2320030130233021-2202110322203113-3321111323232030"></a>

Type: `"object"`. single nested block, Optional.

HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

Upstream description:

A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

Receipt-pinned upstream constraints:

```json
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
http_method {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003231001311033-1302003133220021-1302010012003120-1033033010210322-0012131211210001-1112111100302113-0311301313311000-3133023001323010"></a>

## Direct properties — http_method / 111223223310 / 3

<a id="canonical-2322231000110332-1013032003311300-2233231232011320-0112133100311110-3121110121213321-3120203021201033-3301000023001232-3123311011303113"></a>

<a id="canonical-1130022312122230-3330131322320310-2312222222230013-1023031032203003-3302231032303303-2202110000233331-0201100232132103-2301313221102003"></a>

## invert_matcher property — http_method / 111223223310 / 4

Type: `"bool"`. Optional.

Invert Method Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0233332131131232-1200032030013311-3130032222102322-1300103210133013-2200101000121211-0020033303302232-3203123111323322-3001030103020221"></a>

<a id="canonical-0202332313022133-1112121022131101-1301303303002003-0212023213132200-0020033133010121-0021201020203013-3121310311020231-2013030210002313"></a>

## methods property — http_method / 111223223310 / 5

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of methods values to match against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2011110332211022-0102320212101012-3001323202032033-0031011123223312-2122022232021132-1033003203033201-2330013311213132-1012022213110102"></a>

## Next pages — http_method / 111223223310 / 6

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3131232312332313-2331122113110111-0131010211100102-0310213132121112-3111200310110021-1010121212033013-3232311000121330-1220231122112110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311232001321213-0102202033130230-3310323102031032-3003233211221223-0200323302030220-3313202302231302-3012232110230012-0331311130301020"></a>

## policy_based_challenge.rule_list.rules.spec.ip_matcher — ip_matcher / 112320201332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.ip_matcher

<a id="canonical-0023210031221301-0320301010030320-2003132110122223-3200103210200311-1321320111103111-2020021032132221-2030003211130130-1331022000121320"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332313230322323-0200010312333020-0101023313022302-1112312032101031-3203030123023223-2200213000003323-0030213232011312-1021032013322111"></a>

## Direct properties — ip_matcher / 112320201332 / 3

<a id="canonical-3132120210121311-3032132212210022-0301222301010110-3322331131000322-0302330031022333-1232013331133202-3003211233330032-2222122322110333"></a>

<a id="canonical-3310102330013012-2210231202222330-3123313030103313-2012033020301100-1231113120201300-2310131332321123-0010322310213231-1221131202322022"></a>

## invert_matcher property — ip_matcher / 112320201332 / 4

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [prefix_sets](resources--http_loadbalancer--reference--group-023.md#canonical-0212332012223100-0032200302201123-2123130003231102-1023020003022121-1133332212200021-0132021010113201-0202231131023203-2100111123101220): complete subsection reference.

<a id="canonical-2132313103300232-1113103202002111-1312001010110301-3102320321013213-2101311221021003-2120123002212032-0232003300001102-0013100003033303"></a>

## Next pages — ip_matcher / 112320201332 / 5

- [policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets](resources--http_loadbalancer--reference--group-023.md#canonical-0212332012223100-0032200302201123-2123130003231102-1023020003022121-1133332212200021-0132021010113201-0202231131023203-2100111123101220)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0212332012223100-0032200302201123-2123130003231102-1023020003022121-1133332212200021-0132021010113201-0202231131023203-2100111123101220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333112213200013-2030313001330212-2131201001101331-0003130210230232-1210331010303300-3331103113222202-1130302102133133-2333201003310213"></a>

## policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets — prefix_sets / 013000233220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](resources--http_loadbalancer--reference--group-023.md#canonical-3131232312332313-2331122113110111-0131010211100102-0310213132121112-3111200310110021-1010121212033013-3232311000121330-1220231122112110)
- policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets

<a id="canonical-0330020020230210-0021231320002110-0010301220333003-3311213313022022-2332322121022201-0002110000300321-1310322230132333-1111103201211230"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011000033323231-3203002033123133-2333333110102202-0100011120123203-2303302203311201-3320301010111030-0030022313211123-0211002210031030"></a>

## Direct properties — prefix_sets / 013000233220 / 3

<a id="canonical-3013310222301300-1303103300311211-0110131313201031-2120010231023233-3133310030321131-0021102231212322-2220021133001110-0213221300022131"></a>

<a id="canonical-3202112310212033-1300000123333020-2203313020100120-3313030332221011-2102312320221222-0133320202101300-2032013010023101-0130033210101023"></a>

## kind property — prefix_sets / 013000233220 / 4

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

<a id="canonical-0133303333302231-0011331323122221-1330200113023302-2003111302011220-2032230022132102-2332320102202322-2130133210111130-0103311111203010"></a>

<a id="canonical-1232212210323111-2212011212002321-0122201331000220-1113113212002302-0303230213303312-1221023113222230-3211130123101331-0220312202300221"></a>

## name property — prefix_sets / 013000233220 / 5

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

<a id="canonical-1202123301203013-3302310302113032-3031301202212123-0201123331033010-0312211221203321-3022113130100123-0303321313302113-0010010011133123"></a>

<a id="canonical-3121303220012122-1101003131010301-3101210121313001-2100003011222222-1213031033110101-3113133110112221-0003122131302002-0002100133133021"></a>

## namespace property — prefix_sets / 013000233220 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2003201002111230-0301232123032323-0220212331012131-1330230313032132-0121223032031130-1201120231210300-2001102010222222-2100311113012312"></a>

<a id="canonical-0221030210231300-2000331201213031-2333133110113001-3213100101021220-0011133331111111-3210203123310013-3322230223223302-0211023303310001"></a>

## tenant property — prefix_sets / 013000233220 / 7

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

<a id="canonical-0332300213330031-3130122101212030-1013222320002312-0313110030003131-3033103122010133-0011311223023230-2113101223322201-1122313320133003"></a>

<a id="canonical-1132030033321000-3131322112302213-3300212101110123-0212211300202210-2202010220111230-3103000212100021-0123002133330323-0033032013111312"></a>

## uid property — prefix_sets / 013000233220 / 8

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

<a id="canonical-3123001331011133-2312321133201331-3232333003131001-0010323012222300-3113320311330300-1000011322312121-0122330202021121-1323201331122112"></a>

## Next pages — prefix_sets / 013000233220 / 9

- [policy_based_challenge.rule_list.rules.spec.ip_matcher](resources--http_loadbalancer--reference--group-023.md#canonical-3131232312332313-2331122113110111-0131010211100102-0310213132121112-3111200310110021-1010121212033013-3232311000121330-1220231122112110)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1102303013213221-2101121232233030-3032103211220121-2031323223113213-2210001011111012-1001212110123010-0120200232130210-3120021232133212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113211320232312-2133312123123223-3201101113230213-3031131321020002-0023001232100112-1110330130202322-0310323211323313-1210232202120311"></a>

## policy_based_challenge.rule_list.rules.spec.ip_prefix_list — ip_prefix_list / 132231121001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.ip_prefix_list

<a id="canonical-3210033200001103-2012311001110030-3021022230112021-3121211323010123-1101313300033203-1101232313320203-1022013101101330-1100220210102312"></a>

Type: `"object"`. single nested block, Optional.

List of IP Prefix strings to match against.

Receipt-pinned upstream constraints:

```json
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
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1112202021011323-1110001102022230-3111011223220212-0321122321203031-3130121012301020-0111313322102010-3302233200213011-0303120130212311"></a>

## Direct properties — ip_prefix_list / 132231121001 / 3

<a id="canonical-0023300312023321-1222112232000303-3312020333323100-3010130101213223-2102133130320000-0111122002211030-2021230230212113-0211012302331033"></a>

<a id="canonical-1203231003002300-0122023203331321-0300030023231213-0012300002333032-3030220110001110-1302132220103123-1211333223332112-2201231021033130"></a>

## invert_match property — ip_prefix_list / 132231121001 / 4

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0133032110130220-2310121223020020-0123310010210101-3003212322113100-0332332333110222-3011012202111312-0122310132302330-3303201322113222"></a>

<a id="canonical-3221320022002203-2030331132321323-2231300112310212-2201311332202331-2220021223222101-0030203102023010-0310222000203102-2203331002103021"></a>

## ip_prefixes property — ip_prefix_list / 132231121001 / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1010213231232202-3323223321031231-0211033331032223-0130313312023220-2023311231021030-1201332301303211-3221020130032300-2311110221010232"></a>

## Next pages — ip_prefix_list / 132231121001 / 6

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2110130302132222-1333323202323202-2101132211132333-3120123331020002-2321313212221101-0010000002320121-1033331331113211-1113310312332201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133220012331120-2122122233102121-3023302110001311-2010200200333100-3010013210303211-1033211101100330-2112133230221113-0323113121313221"></a>

## policy_based_challenge.rule_list.rules.spec.path — path / 100000202132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.path

<a id="canonical-1013203123330210-1012013200102220-0003230000130221-0032221333233303-0232201031022121-1113220202200323-1130210333213332-3201203112101200"></a>

Type: `"object"`. single nested block, Optional.

Path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

Upstream description:

A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
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
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000210211012110-2111230122302121-1132332111210312-3033210212030213-1231121323202202-0111202231121310-3312310232022322-2112011232022002"></a>

## Direct properties — path / 100000202132 / 3

<a id="canonical-2322213223333133-2230111223213113-1101030332232301-2312111231210202-1001000210222102-1102333303220203-2232023020012330-1302120230032331"></a>

<a id="canonical-1200231313123222-3013023230022331-2001023133233103-3331031011132121-2221130302320213-0223213323312022-0312332222113303-2221203110133210"></a>

## encoded_path_matcher property — path / 100000202132 / 4

Type: `"bool"`. Optional.

Match against the encoded, escaped path.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2032022121322200-1132302300113032-1330302213013232-1212021303311003-3221103100300013-2321302312322201-0023000111101321-3101321303201233"></a>

<a id="canonical-1211100300033300-3002132113221111-3303001011132323-2221331013122012-2222002332302033-2010101213123301-2021003012322121-1020110211230030"></a>

## exact_values property — path / 100000202132 / 5

Type: `["list", "string"]`. Optional.

List of exact path values to match the input HTTP path against.

Upstream description:

A list of exact path values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0103011303301111-0020211031023001-0230102012233333-2211311222232032-2133103012030033-1300302330111333-3022200121020112-0013111211022231"></a>

<a id="canonical-3303123032221013-2211030022200201-2200303012133233-0012021213001013-1000310122000311-0132021110212202-1311200002130121-2201302013311111"></a>

## invert_matcher property — path / 100000202132 / 6

Type: `"bool"`. Optional.

Invert Path Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1303031331230310-2001233201022112-0202233230120120-3101303133200212-1032201333023321-2120002230031330-1130220023203132-0023203311310311"></a>

<a id="canonical-3300131311312331-3212331322233102-0312222203003021-2121121222203312-3212013033230102-3000310000203310-2132112112002121-3120220213030102"></a>

## prefix_values property — path / 100000202132 / 7

Type: `["list", "string"]`. Optional.

List of path prefix values to match the input HTTP path against.

Upstream description:

A list of path prefix values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1130102211330003-2222230333300210-3302213111012320-1003323322312130-0222303132322103-0331231303320230-1100120303100111-2321113312020121"></a>

<a id="canonical-2032202003321212-3232011202101102-3001200023110003-2022212211313232-2230213300331013-1011333303202131-1031320133323102-0130112100113010"></a>

## regex_values property — path / 100000202132 / 8

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input HTTP path against.

Upstream description:

A list of regular expressions to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0221232012101112-1313011111020110-2231103332101331-0310313113013220-2000010303033011-3221201010321131-3300132103331132-3222212030302220"></a>

<a id="canonical-3231133310201123-3131203310112230-2011100212030003-0122120230022321-2321321132222021-0123330320103103-2231100323030320-1101010110332120"></a>

## suffix_values property — path / 100000202132 / 9

Type: `["list", "string"]`. Optional.

List of path suffix values to match the input HTTP path against.

Upstream description:

A list of path suffix values to match the input HTTP path against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1312322031021220-0320122303112122-0320230103103300-3102031321330000-3110132213132230-0310331232221303-2132301001233230-0311302031231000"></a>

<a id="canonical-2331021203310021-3332310212320203-0332221320122200-0133120331220211-3202313330021020-1223102023320322-2300303132321322-1330101322100220"></a>

## transformers property — path / 100000202132 / 10

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1230220213222211-1322303310133021-2232312010222330-2330303312320321-0330312003303003-2102121110102123-2120023322013022-2211333023213222"></a>

## Next pages — path / 100000202132 / 11

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102113321203001-1331211222203230-3223232032320001-3120100331313010-2312011112310032-3003003231020022-3212331203001023-3001110023212320"></a>

## policy_based_challenge.rule_list.rules.spec.query_params — query_params / 003320000013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.query_params

<a id="canonical-2203120210113211-2112000210103232-3102220020002112-2020330112232021-2101223003312230-0312303201312023-3031122000103101-3013000312102301"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010301201110232-3113120110021303-0302321233223032-0101020230212321-1122200111012032-1122300121102333-1103321331000102-2210310213003031"></a>

## Direct properties — query_params / 003320000013 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-023.md#canonical-3103300010123120-3200220011010132-3033202200030111-2303133000000131-0311023323212320-0013122300031122-0222001002323220-2332111330223102): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-023.md#canonical-2112003030120300-3231323312231200-0322203333323223-3322103103222131-3012202202123232-0312103310010201-0032132013131023-0001222211300312): complete subsection reference.

<a id="canonical-3021231120131001-1310320220111012-2023102213231133-1100132023023123-1332132102033100-1212122232202300-1123111031233320-3332030310102010"></a>

<a id="canonical-3210231322322222-1003000110123001-2301202033312322-2120033133231212-3230231210230122-1123313001230332-0113220220211021-2001311023131023"></a>

## invert_matcher property — query_params / 003320000013 / 4

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--http_loadbalancer--reference--group-023.md#canonical-3221303202213312-3211202232332112-3101121302023021-3001021033200203-3301011202001122-1221212210100122-1300213130231002-3330311103023100): complete subsection reference.

<a id="canonical-2213321300002311-2112331211233132-3210113330113321-0101011031122232-3022013312103303-3103031011021011-2121230131303312-3223130213312202"></a>

<a id="canonical-3102010330212123-3102101000300113-3203133132321032-1020331120300001-0321103311211321-2012123303122001-0200311100203303-1230132120303233"></a>

## key property — query_params / 003320000013 / 5

Type: `"string"`. Optional.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3230012131311210-1100021210122201-2322102012030332-1101110031001020-1103210021110332-3121300031131033-0020223131011310-2020231012100001"></a>

## Next pages — query_params / 003320000013 / 6

- [policy_based_challenge.rule_list.rules.spec.query_params.check_not_present](resources--http_loadbalancer--reference--group-023.md#canonical-3103300010123120-3200220011010132-3033202200030111-2303133000000131-0311023323212320-0013122300031122-0222001002323220-2332111330223102)
- [policy_based_challenge.rule_list.rules.spec.query_params.check_present](resources--http_loadbalancer--reference--group-023.md#canonical-2112003030120300-3231323312231200-0322203333323223-3322103103222131-3012202202123232-0312103310010201-0032132013131023-0001222211300312)
- [policy_based_challenge.rule_list.rules.spec.query_params.item](resources--http_loadbalancer--reference--group-023.md#canonical-3221303202213312-3211202232332112-3101121302023021-3001021033200203-3301011202001122-1221212210100122-1300213130231002-3330311103023100)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3103300010123120-3200220011010132-3033202200030111-2303133000000131-0311023323212320-0013122300031122-0222001002323220-2332111330223102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303323020032312-2013002212013301-2031003131122121-0222232200212312-2101010330310323-3221213232001233-2102230020033200-3233130013121331"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.check_not_present — check_not_present / 030101023103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-023.md#canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303)
- policy_based_challenge.rule_list.rules.spec.query_params.check_not_present

<a id="canonical-1101303100132312-0332013330032313-0322313012033222-1203311103123331-3301300321311102-1310132011213220-3200210002113130-1333201120100332"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-2122022033223120-2103021330131022-2330310332120133-3033033112310330-1233321303231221-3322133231120111-0303123102010013-3112221032201023"></a>

## Direct properties — check_not_present / 030101023103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312102213303331-1033211223002002-3113223212303021-0201311031330100-3312102320310210-1322323210110320-3300310021033002-3233201211333103"></a>

## Next pages — check_not_present / 030101023103 / 4

- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-023.md#canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2112003030120300-3231323312231200-0322203333323223-3322103103222131-3012202202123232-0312103310010201-0032132013131023-0001222211300312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023203230002023-1133230202000221-1002200012012203-3110131131211321-3030010200200300-2323220312101231-0320331310321201-1220321002201313"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.check_present — check_present / 030121103202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-023.md#canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303)
- policy_based_challenge.rule_list.rules.spec.query_params.check_present

<a id="canonical-2012232230312011-2000002102012122-2312132210110121-2112113203321232-0200231201101010-2323330333301321-2201301222133133-0023212023130201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-0230023023003012-1211201312101120-2113030002333011-1023120103033131-3211032030113200-2313110200011322-2223032321023202-1113210223012310"></a>

## Direct properties — check_present / 030121103202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102313103000030-3020312310130302-2102232332321120-3320002331010213-3102213221001111-3210002200131013-0013200122032232-1102333212222031"></a>

## Next pages — check_present / 030121103202 / 4

- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-023.md#canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3221303202213312-3211202232332112-3101121302023021-3001021033200203-3301011202001122-1221212210100122-1300213130231002-3330311103023100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031103030220132-2331300003310232-1312013111133223-3202032033210111-0332212132212101-0031101013312301-0103000011203012-2310311001030032"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.item — item / 201201202311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-023.md#canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303)
- policy_based_challenge.rule_list.rules.spec.query_params.item

<a id="canonical-2021111131031122-1233033121101103-0232020210133021-3231312303321331-0332101000121310-3203301010032330-3000011302111023-0310102312313100"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131201012300113-0133232213202002-0032112122133311-3121111212323221-1300310103303020-2221003202333120-1010132011321113-1100223233232103"></a>

## Direct properties — item / 201201202311 / 3

<a id="canonical-2000211123332030-2312331230233130-0231333310131100-0233231231311302-3300203021110132-0232132211223001-0021331100330232-0230102120110212"></a>

<a id="canonical-2101011333123332-2023222030321302-3332230321032220-3202211212221111-1333222111030301-0302323001102300-2002021202031201-1310111211201011"></a>

## exact_values property — item / 201201202311 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2231211210223302-3221220321110013-2222131131220002-0002300123031321-2230211202033333-0012031032022033-3013221012101301-3220312010312122"></a>

<a id="canonical-3130122230103123-0312310102002112-2101021200020012-1032322201303331-2321012000332202-0103003302310330-3023111200300032-0311031101303311"></a>

## regex_values property — item / 201201202311 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1130330003220001-2130120332033120-0302103320323222-2323001321301111-3211023222033022-2030212331202010-2222121312012321-0133233211000333"></a>

<a id="canonical-0322213233133300-2212333213212102-3123030210002133-0233212132332000-3000120021313202-3330303022111330-0123221011222211-2331130223112222"></a>

## transformers property — item / 201201202311 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3302130020222331-2220301133131332-2302001303122211-1120211300010303-2202133122330010-3002030133101320-3211030023302301-1032033213220101"></a>

## Next pages — item / 201201202311 / 7

- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--reference--group-023.md#canonical-3301321013233130-3120220022102302-2010030102010301-1203130111033000-2101212222020102-2320031102212001-0133011132121102-0210000220320303)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1131033020202020-3130333330012030-3002001123121030-1323223321202032-1221131132031001-3323021113022011-1203131303033130-3010031021310313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013031323323213-3313121012212012-0112303321001030-1100023031023311-0210001200203111-3132002132022111-2010033301222100-0102210102133120"></a>

## policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher — tls_fingerprint_matcher / 010330112013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-3001323211333112-0203222131002030-0131123030132222-2001132200110102-2323130230010101-3101313130101001-3222101033033013-2001320331003033)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--reference--group-022.md#canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030)
- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher

<a id="canonical-0023121113310101-3000100110211120-0321313013123100-2223023123312331-0021220321232200-0320221203123231-3231231132101211-1112110300100212"></a>

Type: `"object"`. single nested block, Optional.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

Receipt-pinned upstream constraints:

```json
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
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010021303323201-3021331310213001-0121022023311100-3000002131303021-3213113020311211-3210100120233332-1131202302322303-1301020103220021"></a>

## Direct properties — tls_fingerprint_matcher / 010330112013 / 3

<a id="canonical-1021111321203203-3301132133030002-2010322033311100-1303332302003321-3213112020223121-0220221202133210-0303002333032202-1302001231301203"></a>

<a id="canonical-0112323210101301-1331121030333023-3033301212222231-2103032032321022-0110210132003200-2301130123200220-0111301112033112-2133233202130310"></a>

## classes property — tls_fingerprint_matcher / 010330112013 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-2200130213322322-1202121203133021-2311320332102331-2320112203331011-1132223232220312-1301022313313100-0322311023001212-1332110222103030"></a>

<a id="canonical-3033230020200031-1202311330101032-1223133323032300-0200002313333012-0102232113132101-2312220212011202-0101230332230301-0300121123123003"></a>

## exact_values property — tls_fingerprint_matcher / 010330112013 / 5

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1201130212030203-2312220320013322-0320311001231222-2232020111320120-1110313223210301-3210113301312033-2121201200222023-0112222120120130"></a>

<a id="canonical-0003300022322130-0230331323133131-0132001121111103-0300210022301230-3300131303221233-3130231013233212-1220020303322313-1130221000100002"></a>

## excluded_values property — tls_fingerprint_matcher / 010330112013 / 6

Type: `["list", "string"]`. Optional.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3200230332213010-1010022131133103-1303020110202202-0122022223320133-2033112212333032-2231110200312022-0230203023221212-3231231132200313"></a>

## Next pages — tls_fingerprint_matcher / 010330112013 / 7

- [policy_based_challenge.rule_list.rules.spec](resources--http_loadbalancer--reference--group-022.md#canonical-1023220031310023-2010202200302230-1133013211202103-2301013120222102-2322001021302200-0210101320021121-2202031212203213-1230023023223212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0332322110110101-3123223013302023-0321200302320201-2312111021001220-0032301331120201-2310003212230313-2113130213200312-2022222232311001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001233223103102-2321002300333321-1010002332023330-2130131120101022-3320211110300131-0000103020320331-1111331221322023-3112130013330213"></a>

## policy_based_challenge.temporary_user_blocking — temporary_user_blocking / 333300210223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- policy_based_challenge.temporary_user_blocking

<a id="canonical-1213330020020110-0201003231121130-0322032122200333-1222232301221110-1231122130323233-0213300101032113-3211212012313013-1313011303000120"></a>

Type: `"object"`. single nested block, Optional.

Specifies configuration for temporary user blocking resulting from user behavior analysis. When
Malicious User Mitigation is enabled from service policy rules, users' accessing the application
will be analyzed for malicious activity and the configured mitigation actions will be taken on..

Upstream description:

Specifies configuration for temporary user blocking resulting from user behavior analysis.

When Malicious User Mitigation is enabled from service policy rules, users' accessing the
application will be analyzed for malicious activity and the configured mitigation actions will be
taken on identified malicious users. These mitigation actions include setting up temporary blocking
on that user. This configuration specifies settings on how that blocking should be done by the
loadbalancer.

Receipt-pinned upstream constraints:

```json
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
temporary_user_blocking {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223200103123211-2201031001133033-2320302111311102-0210212323311133-3130223121030022-1311321233010231-1230223211221100-1211330123231130"></a>

## Direct properties — temporary_user_blocking / 333300210223 / 3

<a id="canonical-3120200303022222-1320212320330112-2211020312113103-1012301200100332-0131032211233132-2100000001010121-3122131220320112-0210110121030023"></a>

<a id="canonical-3021021321313020-2122221321310002-2332221213231201-3201010022103123-2303013023301211-1123320011313120-1231332102002220-0101233032231003"></a>

## custom_page property — temporary_user_blocking / 333300210223 / 4

Type: `"string"`. Optional.

Custom message is of type . Currently supported URL schemes is . For scheme, message needs to be
encoded in base64 format. You can specify this message as base64 encoded plain text message e.g.
'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '&lt;p&gt;
Blocked..

Upstream description:

Custom message is of type \`uri\_ref\`. Currently supported URL schemes is \`string:///\`. For
\`string:///\` scheme, message needs to be encoded in base64 format. You can specify this message as
base64 encoded plain text message e.g. "Blocked.." or it can be HTML paragraph or a body string
encoded as base64 string E.g. "&lt;p&gt; Blocked &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0331021302321232-1311030023311202-2101033123200021-1032210023321103-3120221211331131-2121010212013211-2222132313232032-0122130121203203"></a>

## Next pages — temporary_user_blocking / 333300210223 / 5

- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313133323322130-1111101311101100-1311201113313021-2211332110031211-0312113310011223-1213211301102330-0220030111133311-1002023200222010"></a>

## protected_cookies — protected_cookies / 331031130323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- protected_cookies

<a id="canonical-0012220312310111-2303132013122322-1110302013133310-3121322300132331-3211323303223113-1121220330010001-3331130300302213-1211121030222300"></a>

Type: `"object"`. list nested block, Optional.

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite.

Upstream description:

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF
(monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note:
We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingListObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingListObjectAttributes("disable_tampering_protection",
    "enable_tampering_protection"),
  validators.ConflictingListObjectAttributes("ignore_max_age",
    "max_age_value"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_none",
    "samesite_strict")}
```

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

Terraform syntax:

```terraform
protected_cookies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320233203100101-0123331113233032-2013002133222330-3323310301303231-1011212011011202-0333210110322202-3220212211231022-3013220001211032"></a>

## Direct properties — protected_cookies / 331031130323 / 3

- [add_httponly](resources--http_loadbalancer--reference--group-023.md#canonical-3300300313122023-0131101133121122-0111302321111221-1320231013003202-0133111311232132-2000031230020032-2303313311132310-0320021312132211): complete subsection reference.

- [add_secure](resources--http_loadbalancer--reference--group-023.md#canonical-2231033302001101-0132231313110213-2031112210223222-3333133100110021-1233300112010200-2023001101013012-2002112312201120-1102200221332013): complete subsection reference.

- [disable_tampering_protection](resources--http_loadbalancer--reference--group-023.md#canonical-3213201012220223-2211332120011102-3023133313022103-0320131213311121-2310000233223030-2021102301111110-1000311211111311-1220303110323223): complete subsection reference.

- [enable_tampering_protection](resources--http_loadbalancer--reference--group-023.md#canonical-2120321222300232-3201213120231321-3002320121130331-3033120022323313-2001001201032333-3213123201002332-1312021122123131-1103100231332211): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-023.md#canonical-0203212110203021-2200220231120210-2230311113031102-3221022331201212-2022331233333222-1023330101000213-2212231131222231-2200000112100101): complete subsection reference.

- [ignore_max_age](resources--http_loadbalancer--reference--group-023.md#canonical-0113232103232213-1220110302131233-2220201300012301-1223330101120121-3102023230122321-2021001312001331-3212303132300032-1121003002111111): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-023.md#canonical-1222213222130010-1030032131320230-1313212131202032-0230100032002121-2200132212221021-2032011033232331-1222130131101301-3033111203030322): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-023.md#canonical-2003122021101102-2103020000110303-1022023202020222-2212321100200322-1113103131312223-1311113032323233-2222101320213033-3231231121303232): complete subsection reference.

<a id="canonical-1323120312300302-3320100011310320-0011202033330112-3301121130120012-0301301111231001-0102330220212010-3133000111121113-1020102133110003"></a>

<a id="canonical-1023311022121202-2100322213130030-0000113020103113-3123001221123122-3013211203111212-3212211322322301-0023203201122112-1011021231131202"></a>

## max_age_value property — protected_cookies / 331031130323 / 4

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(34560000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-1321133322311221-1000131311233203-0132313321131303-0222222031301101-0233031223010201-0331020201130000-1203312001333210-3333222232301112"></a>

<a id="canonical-3211202232231123-2301230321133300-1031030303033022-1003100320213031-0230031321102101-1022212111123121-1033010031213112-1321322020103130"></a>

## name property — protected_cookies / 331031130323 / 5

Type: `"string"`. Optional.

Cookie Name. Name of the Cookie.

Upstream description:

Name of the Cookie.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [samesite_lax](resources--http_loadbalancer--reference--group-023.md#canonical-2000130001101232-3100031031201010-3211203321113321-1000321203033310-2232002312022102-1223200123130111-2321322310212321-2000102201320130): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--reference--group-023.md#canonical-1010202312223010-3101022001131300-0120102210132103-1310330112122232-3020103212011201-0132121012321221-2320211000122301-2023300212031311): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--reference--group-023.md#canonical-3202002233301113-3300122322310232-3222101130120113-3130000312231310-2013103001102003-0122111300211202-0030121200112032-3020100220132200): complete subsection reference.

<a id="canonical-0120311203330012-2033322302100333-2222030032131001-0101031222031310-0111221320222211-1130330223212310-2221102110220200-2320113201002231"></a>

## Next pages — protected_cookies / 331031130323 / 6

- [protected_cookies.add_httponly](resources--http_loadbalancer--reference--group-023.md#canonical-3300300313122023-0131101133121122-0111302321111221-1320231013003202-0133111311232132-2000031230020032-2303313311132310-0320021312132211)
- [protected_cookies.add_secure](resources--http_loadbalancer--reference--group-023.md#canonical-2231033302001101-0132231313110213-2031112210223222-3333133100110021-1233300112010200-2023001101013012-2002112312201120-1102200221332013)
- [protected_cookies.disable_tampering_protection](resources--http_loadbalancer--reference--group-023.md#canonical-3213201012220223-2211332120011102-3023133313022103-0320131213311121-2310000233223030-2021102301111110-1000311211111311-1220303110323223)
- [protected_cookies.enable_tampering_protection](resources--http_loadbalancer--reference--group-023.md#canonical-2120321222300232-3201213120231321-3002320121130331-3033120022323313-2001001201032333-3213123201002332-1312021122123131-1103100231332211)
- [protected_cookies.ignore_httponly](resources--http_loadbalancer--reference--group-023.md#canonical-0203212110203021-2200220231120210-2230311113031102-3221022331201212-2022331233333222-1023330101000213-2212231131222231-2200000112100101)
- [protected_cookies.ignore_max_age](resources--http_loadbalancer--reference--group-023.md#canonical-0113232103232213-1220110302131233-2220201300012301-1223330101120121-3102023230122321-2021001312001331-3212303132300032-1121003002111111)
- [protected_cookies.ignore_samesite](resources--http_loadbalancer--reference--group-023.md#canonical-1222213222130010-1030032131320230-1313212131202032-0230100032002121-2200132212221021-2032011033232331-1222130131101301-3033111203030322)
- [protected_cookies.ignore_secure](resources--http_loadbalancer--reference--group-023.md#canonical-2003122021101102-2103020000110303-1022023202020222-2212321100200322-1113103131312223-1311113032323233-2222101320213033-3231231121303232)
- [protected_cookies.samesite_lax](resources--http_loadbalancer--reference--group-023.md#canonical-2000130001101232-3100031031201010-3211203321113321-1000321203033310-2232002312022102-1223200123130111-2321322310212321-2000102201320130)
- [protected_cookies.samesite_none](resources--http_loadbalancer--reference--group-023.md#canonical-1010202312223010-3101022001131300-0120102210132103-1310330112122232-3020103212011201-0132121012321221-2320211000122301-2023300212031311)
- [protected_cookies.samesite_strict](resources--http_loadbalancer--reference--group-023.md#canonical-3202002233301113-3300122322310232-3222101130120113-3130000312231310-2013103001102003-0122111300211202-0030121200112032-3020100220132200)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3300300313122023-0131101133121122-0111302321111221-1320231013003202-0133111311232132-2000031230020032-2303313311132310-0320021312132211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313302113212300-0110121331202123-0013002031313200-1123232310031030-0123100101213303-1313231002111032-2020331101130011-2333132032102301"></a>

## protected_cookies.add_httponly — add_httponly / 120012110012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.add_httponly

<a id="canonical-3211233200132230-1320201203133210-3222121231102200-3101310200102331-1131033211103203-2210123131100133-0230202031023310-2213013200001220"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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
add_httponly = {}
```

<a id="canonical-2200133100013300-1110131301001103-2202131230211021-0110312031221331-2320212010330233-0012033032120200-2101313021013200-3020300322011033"></a>

## Direct properties — add_httponly / 120012110012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220122021311021-0013220020312113-0001022321002131-2333220111003313-3233102022330232-2211112101103003-0013220301031030-1332011312302011"></a>

## Next pages — add_httponly / 120012110012 / 4

- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2231033302001101-0132231313110213-2031112210223222-3333133100110021-1233300112010200-2023001101013012-2002112312201120-1102200221332013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030221310010232-3202033131332010-1230210122000011-1300112003022222-3300023023202233-2010221012031222-1113100031131110-0011001011321322"></a>

## protected_cookies.add_secure — add_secure / 020213311303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.add_secure

<a id="canonical-0020002302213013-0321120233322011-2303112122210301-0220003000230021-0212331113021030-2020021121211131-3020302100222213-3111320232030233"></a>

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
add_secure = {}
```

<a id="canonical-3311312023202000-0332233330303031-2332111123113211-0201210303002211-3023122102030021-2233110013030221-2121123212130312-0132322013101101"></a>

## Direct properties — add_secure / 020213311303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311232122202021-3031232111232103-1312320030010112-3212020111232113-2103202233200002-2133211211303322-2302212100131300-2021321110300330"></a>

## Next pages — add_secure / 020213311303 / 4

- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3213201012220223-2211332120011102-3023133313022103-0320131213311121-2310000233223030-2021102301111110-1000311211111311-1220303110323223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103200230101300-2201020112013121-1000002023001001-0200220100112120-3113331221010012-1213101200012000-2313123300110332-2013003123310312"></a>

## protected_cookies.disable_tampering_protection — disable_tampering_protection / 031331010033 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.disable_tampering_protection

<a id="canonical-1311001213222122-0311202222022223-0120102331233220-3130333230233133-1110013230032203-3223310200332130-2130333212112111-3111321302113321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable tampering protection.

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
disable_tampering_protection = {}
```

<a id="canonical-3131032303100131-3223112222210132-2011011112321020-0023333231333200-0011111320232231-2031211211132010-2101211301321110-3131021112111210"></a>

## Direct properties — disable_tampering_protection / 031331010033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032330010311220-2302012120011020-0111012010303022-2030012111033201-1330321032213130-2001111300320213-0231123221121330-2223210031132123"></a>

## Next pages — disable_tampering_protection / 031331010033 / 4

- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2120321222300232-3201213120231321-3002320121130331-3033120022323313-2001001201032333-3213123201002332-1312021122123131-1103100231332211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102213001033210-1232112033101032-3211121332121300-0032333000121121-2131232000001022-2100121212011203-3122003202313013-2100133011102122"></a>

## protected_cookies.enable_tampering_protection — enable_tampering_protection / 113012321121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.enable_tampering_protection

<a id="canonical-0002202213011112-0311010233013220-1302123011033320-0012132113121123-2210221023223110-2311311303013232-2010300133122211-3032030323010300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable tampering protection.

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
enable_tampering_protection = {}
```

<a id="canonical-3311210301123311-3320302210132230-0012210210300330-3201210323211210-1323220231031030-3012111222131131-2201003201120212-3002331310030330"></a>

## Direct properties — enable_tampering_protection / 113012321121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310132003330331-1322023013312201-3221210332313120-3122113033002310-0131333233012331-3123322133121013-3131010220301100-1212221301021130"></a>

## Next pages — enable_tampering_protection / 113012321121 / 4

- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0203212110203021-2200220231120210-2230311113031102-3221022331201212-2022331233333222-1023330101000213-2212231131222231-2200000112100101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123321133221220-2211001323022302-3223113220212232-1011011300023312-0322213300211202-1322200333223201-1111301321122322-3321023210202200"></a>

## protected_cookies.ignore_httponly — ignore_httponly / 230202303011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.ignore_httponly

<a id="canonical-2021211012122313-0031031302100232-0001013320130133-0131033130312033-3231021213132201-1230300332232332-0233220101310323-3103110023013223"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

<a id="canonical-1202311011003212-2021001202313210-2033030222210111-0202100011202012-1113323301322333-0231210100323033-0012022311231013-0330010102311022"></a>

## Direct properties — ignore_httponly / 230202303011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103210330300321-0022023203231312-1000203100212111-0013221120131200-3310332000320102-3023233231110001-3313330133121010-3021013010311300"></a>

## Next pages — ignore_httponly / 230202303011 / 4

- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0113232103232213-1220110302131233-2220201300012301-1223330101120121-3102023230122321-2021001312001331-3212303132300032-1121003002111111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123113111323213-2003202001202130-0121210001002030-2221330012031001-3333102112001130-3031312201322121-0002210033222131-2331112103330300"></a>

## protected_cookies.ignore_max_age — ignore_max_age / 133112302322 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.ignore_max_age

<a id="canonical-3102101331001031-1020221120320220-3023120133211031-3232311303122110-2220012122230101-1131002101220123-1201320113013011-3003112300031121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore max age.

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
ignore_max_age = {}
```

<a id="canonical-2300211202100022-2000303021112210-2120003223313230-3000213003332123-3033013222200200-1221032320202131-2021033310333020-1103013022122001"></a>

## Direct properties — ignore_max_age / 133112302322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232113303232323-2133221211101100-3102100200021130-3311200121121320-0320012313323312-2031123312131230-2233122103230130-2201300131231320"></a>

## Next pages — ignore_max_age / 133112302322 / 4

- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1222213222130010-1030032131320230-1313212131202032-0230100032002121-2200132212221021-2032011033232331-1222130131101301-3033111203030322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311113021102110-1331211130133031-0210331202303230-1100031011301021-0012110322013101-2210332132220021-1320120201200223-0222312321103310"></a>

## protected_cookies.ignore_samesite — ignore_samesite / 203111230030 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.ignore_samesite

<a id="canonical-1003303130031103-3220313222211303-1313122113333312-0130130022102001-0312331133011002-2011220003023222-2033310012230110-3132322330103132"></a>

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
ignore_samesite = {}
```

<a id="canonical-0103121330212022-3300302212212332-3000002121003223-2112221232232231-1303023121020331-3313313002303232-1030230122303023-1130303213322210"></a>

## Direct properties — ignore_samesite / 203111230030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020013231023212-2021203200020011-1321312232213212-1312001120223123-3330110313112020-2330130101122211-3011303202232211-3003313000110112"></a>

## Next pages — ignore_samesite / 203111230030 / 4

- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2003122021101102-2103020000110303-1022023202020222-2212321100200322-1113103131312223-1311113032323233-2222101320213033-3231231121303232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123300012020010-3310201201132303-0332110022111023-2323013111301112-1220202202010300-1313303311022123-3232003231223112-1211231003022103"></a>

## protected_cookies.ignore_secure — ignore_secure / 112021213020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.ignore_secure

<a id="canonical-3021010300301100-3321220031113313-3033233222300202-0223123021330123-1120113120322110-0132203013000112-1011020203231313-0203311310020133"></a>

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
ignore_secure = {}
```

<a id="canonical-2330302313123210-2113200221133023-0102103220333320-0030220301010321-1003211002101012-1102220031231110-2112012303032320-2323111211123003"></a>

## Direct properties — ignore_secure / 112021213020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303231031012301-2000203313320113-3120112000312332-1110011222000133-0301312023022230-0022321223322221-1000303232032110-2011032201121020"></a>

## Next pages — ignore_secure / 112021213020 / 4

- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2000130001101232-3100031031201010-3211203321113321-1000321203033310-2232002312022102-1223200123130111-2321322310212321-2000102201320130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112121112303323-2111332101000023-1001302120022232-0311210203310003-0002103311233030-2220033010321223-2303131302300222-1102320001123230"></a>

## protected_cookies.samesite_lax — samesite_lax / 333202321133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.samesite_lax

<a id="canonical-2231000213132033-1031012230003310-1001222221110022-3000311232220212-2310010100220213-3022011033303122-3020122001101120-0322321103313103"></a>

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
samesite_lax = {}
```

<a id="canonical-1133110211021031-1133313301231323-2323102321320033-2110320330112331-2020311130220113-1222212302232020-2302131203303103-2020111322312133"></a>

## Direct properties — samesite_lax / 333202321133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201313020111212-0131123032101132-3020103110023133-0023200132332023-0111032313222301-2203010221312203-1131031122201203-2120122110013121"></a>

## Next pages — samesite_lax / 333202321133 / 4

- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1010202312223010-3101022001131300-0120102210132103-1310330112122232-3020103212011201-0132121012321221-2320211000122301-2023300212031311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023212301101023-0313102230202300-0131100122213312-1133213310333033-2111103201312030-3110202220013223-1332320012203202-3333110013220032"></a>

## protected_cookies.samesite_none — samesite_none / 101320021002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.samesite_none

<a id="canonical-3012022311323223-2331300201331233-0103310031123132-0031221331122013-0203200322000332-2033113110210131-1111332221232013-1130113033101123"></a>

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
samesite_none = {}
```

<a id="canonical-2332131123001201-1001032001223323-2203332110310020-1300032321131330-3210331311013100-1101001133311220-2112123001012022-2220013120013002"></a>

## Direct properties — samesite_none / 101320021002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210121320113101-2231023000112223-1010010320020232-2333011100012100-2321321301101023-1130011323001311-3212200030120232-0332200333323020"></a>

## Next pages — samesite_none / 101320021002 / 4

- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3202002233301113-3300122322310232-3222101130120113-3130000312231310-2013103001102003-0122111300211202-0030121200112032-3020100220132200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112322010023320-0121203001121131-0133012111102300-2133001130210033-3203102303310113-3002132323330032-1032220230102100-2103002002330202"></a>

## protected_cookies.samesite_strict — samesite_strict / 011123312122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- protected_cookies.samesite_strict

<a id="canonical-2121330213211000-2200123230023111-1120021233023311-2222101302323212-2122012322123302-0010323322232231-1011201201130112-0332321022200302"></a>

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
samesite_strict = {}
```

<a id="canonical-2111121232133032-2000120201220100-1130230223102022-0201101010222010-1210031133332003-0310302222022201-2100020233003321-3030102203000131"></a>

## Direct properties — samesite_strict / 011123312122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222330121201221-1110010321113102-1030122123120311-1313210232113201-0210101132110023-2011000021023213-1032033013031221-0222023101212323"></a>

## Next pages — samesite_strict / 011123312122 / 4

- [protected_cookies](resources--http_loadbalancer--reference--group-023.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1312120313213323-2103003010102231-1132113202300010-0103032212233002-0220031323220302-3231002030321130-0010020301002111-1312011221320111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212110122133001-3320022211320112-3020123201230333-1211030223301111-1130110031312102-1202021111030022-2002011023231301-2302013013321111"></a>

## random — random / 331311021231 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- random

<a id="canonical-1021313003120223-3101211023223000-1333330100323102-3033202113022313-2111113220003123-1112002011233323-1312030112001012-1111020111100321"></a>

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
random = {}
```

<a id="canonical-2313100021031103-3301133011121220-3010310223003221-3122201002332213-1310220230200012-3202112220321112-1121022311111303-2203313331313331"></a>

## Direct properties — random / 331311021231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213202220333102-3200303132133213-1300320321103300-0121332213030122-0120003020130120-2210322221213101-2023102100233020-3110103103223232"></a>

## Next pages — random / 331311021231 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321223131033012-2032331230200132-1302111022313203-1000133020113002-2231301330103221-0222223212320201-2322012230101001-3002000100302201"></a>

## rate_limit — rate_limit / 112203102212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- rate_limit

<a id="canonical-1212231232113102-0000210123322230-2321331212231011-1223210321030302-2020032100331231-0213102112123132-3222021323133333-1010333102020321"></a>

Type: `"object"`. single nested block, Optional.

Load-balancer-wide per-client rate limiting. The counter applies across every path; use
api\_rate\_limit rules when only selected paths such as /login should be limited.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "ip_allowed_list"),
  validators.ConflictingObjectAttributes("custom_ip_allowed_list",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("ip_allowed_list",
    "no_ip_allowed_list"),
  validators.ConflictingObjectAttributes("no_policies",
    "policies")}
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
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]",
  "x-ves-oneof-field-policy_choice": "[\"no_policies\",\"policies\"]"
}
```

Terraform syntax:

```terraform
rate_limit {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332200320211133-3201100200332331-3221311312210300-3021111212002232-3020200021001031-3022133221002010-2113133132230311-3123330013110130"></a>

## Direct properties — rate_limit / 112203102212 / 3

- [custom_ip_allowed_list](resources--http_loadbalancer--reference--group-023.md#canonical-3111012121110331-0220030113112232-0310311011223312-3011311300331132-3031132310013331-2312333030201102-2032320032220102-3023013332301231): complete subsection reference.

- [ip_allowed_list](resources--http_loadbalancer--reference--group-023.md#canonical-0232210100232323-3030202021133132-1202300110321011-0302032013001321-3232023111021033-1321311022302323-0310203233111223-3020100011333002): complete subsection reference.

- [no_ip_allowed_list](resources--http_loadbalancer--reference--group-023.md#canonical-3010321303322033-1212332122021131-3013312103113211-2131102032023233-3123302121323033-2320112001033113-3312022230012033-3102230311220002): complete subsection reference.

- [no_policies](resources--http_loadbalancer--reference--group-023.md#canonical-1012222010200213-1013130113011130-1203230321210101-0313310220312312-0311230001310213-1113313130100101-2020020003003112-3312012120123223): complete subsection reference.

- [policies](resources--http_loadbalancer--reference--group-023.md#canonical-1201121223001130-0102301012320112-0332103302021323-3120320310211121-2331102322201132-0130330012330030-3122013001202010-1320120013120101): complete subsection reference.

- [rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031): complete subsection reference.

<a id="canonical-2221201132201300-1021322231332210-2330030030022331-3112330003011321-3313222223300022-1020120200013033-0133113011211323-0211021322333323"></a>

## Next pages — rate_limit / 112203102212 / 4

- [rate_limit.custom_ip_allowed_list](resources--http_loadbalancer--reference--group-023.md#canonical-3111012121110331-0220030113112232-0310311011223312-3011311300331132-3031132310013331-2312333030201102-2032320032220102-3023013332301231)
- [rate_limit.ip_allowed_list](resources--http_loadbalancer--reference--group-023.md#canonical-0232210100232323-3030202021133132-1202300110321011-0302032013001321-3232023111021033-1321311022302323-0310203233111223-3020100011333002)
- [rate_limit.no_ip_allowed_list](resources--http_loadbalancer--reference--group-023.md#canonical-3010321303322033-1212332122021131-3013312103113211-2131102032023233-3123302121323033-2320112001033113-3312022230012033-3102230311220002)
- [rate_limit.no_policies](resources--http_loadbalancer--reference--group-023.md#canonical-1012222010200213-1013130113011130-1203230321210101-0313310220312312-0311230001310213-1113313130100101-2020020003003112-3312012120123223)
- [rate_limit.policies](resources--http_loadbalancer--reference--group-023.md#canonical-1201121223001130-0102301012320112-0332103302021323-3120320310211121-2331102322201132-0130330012330030-3122013001202010-1320120013120101)
- [rate_limit.rate_limiter](resources--http_loadbalancer--reference--group-024.md#canonical-1222330210232133-1221130003231103-2011220111303130-2010113100033230-2233211211121033-0120023333312102-0001213220303301-0323032013201031)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3111012121110331-0220030113112232-0310311011223312-3011311300331132-3031132310013331-2312333030201102-2032320032220102-3023013332301231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132300111322010-0211021003112333-1101000211023210-3133030310113020-3232110100200221-2213032332133220-0010223220111031-2313233311030332"></a>

## rate_limit.custom_ip_allowed_list — custom_ip_allowed_list / 121300222022 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.custom_ip_allowed_list

<a id="canonical-3023030301000222-2020202210131000-2230300203230310-2010122011332100-0212133223233100-0322001102312102-3132333221333020-0012303220013220"></a>

Type: `"object"`. single nested block, Optional.

IP Allowed list using existing ip\_prefix\_set objects.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("rate_limiter_allowed_prefixes")}
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
custom_ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120203110230201-1023123011022123-1113222223000311-1023233001223132-1103021022313013-2312222233023111-1230010132011333-2132222123021120"></a>

## Direct properties — custom_ip_allowed_list / 121300222022 / 3

- [rate_limiter_allowed_prefixes](resources--http_loadbalancer--reference--group-023.md#canonical-0220012023113211-1122013121013033-3230313310101110-1331120330133301-1033312101103121-0131131033321010-0231222121022201-2313210322200033): complete subsection reference.

<a id="canonical-2103032123220110-2311132203021002-1233132202100002-0020230203110211-2120102101201033-2233033013303211-2010013112011311-1001113203100333"></a>

## Next pages — custom_ip_allowed_list / 121300222022 / 4

- [rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](resources--http_loadbalancer--reference--group-023.md#canonical-0220012023113211-1122013121013033-3230313310101110-1331120330133301-1033312101103121-0131131033321010-0231222121022201-2313210322200033)
- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0220012023113211-1122013121013033-3230313310101110-1331120330133301-1033312101103121-0131131033321010-0231222121022201-2313210322200033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133301330101023-0100300220130200-0200013020101221-1300013131303220-1131311221223023-1300210323030321-2133212222122201-0330012120013303"></a>

## rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes — rate_limiter_allowed_prefixes / 330203231113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.custom_ip_allowed_list](resources--http_loadbalancer--reference--group-023.md#canonical-3111012121110331-0220030113112232-0310311011223312-3011311300331132-3031132310013331-2312333030201102-2032320032220102-3023013332301231)
- rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-2000123031021033-3223132012102132-3231013132332332-0130332221003321-1122312113200300-0023001323113301-0133122032332333-1122230203103312"></a>

Type: `"object"`. list nested block, Optional.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Upstream description:

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
rate_limiter_allowed_prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101023303220002-3212131000301113-0001023033100322-0131310233111010-0303322323102012-3222213132033323-2002212031330211-2232030222110013"></a>

## Direct properties — rate_limiter_allowed_prefixes / 330203231113 / 3

<a id="canonical-0130033032121322-2110201010312330-2031012013112332-3203310322133323-2320212210211101-3302112112033131-2331123100110300-3323113302320110"></a>

<a id="canonical-3230021331331230-3331012101002212-2011100000223331-0031131132300111-2202221030103221-0100003033033301-0213203231201331-0203011011321320"></a>

## name property — rate_limiter_allowed_prefixes / 330203231113 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3021320010033301-1203300230030301-3310320321130000-1232331113100101-3321311201212001-1232202002000323-2323223222211221-1321201300211120"></a>

<a id="canonical-1001330000222032-0310031300223322-1212202133202130-2300122332220112-1002111110031202-3121301131022110-1120120202322322-3312312100322111"></a>

## namespace property — rate_limiter_allowed_prefixes / 330203231113 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1230103032221011-1112302112121002-1032213222300133-3121303132302231-0230313112322331-3333332021221320-2110331131320322-2030203223301020"></a>

<a id="canonical-0320233001322221-2330301220302101-0121020211210113-0200230311010312-3332010031113313-0221320233123032-2220001233320232-0320112210201213"></a>

## tenant property — rate_limiter_allowed_prefixes / 330203231113 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0323222010012201-1212332030113003-1110210202002300-3311233130330321-3032232233202013-3312211333120001-2120221131302102-0131001113013020"></a>

## Next pages — rate_limiter_allowed_prefixes / 330203231113 / 7

- [rate_limit.custom_ip_allowed_list](resources--http_loadbalancer--reference--group-023.md#canonical-3111012121110331-0220030113112232-0310311011223312-3011311300331132-3031132310013331-2312333030201102-2032320032220102-3023013332301231)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0232210100232323-3030202021133132-1202300110321011-0302032013001321-3232023111021033-1321311022302323-0310203233111223-3020100011333002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232211132101301-1030100301020331-3233313300013210-2213300100120033-0211101203311312-1021222200303320-1031132103313101-3113023301212200"></a>

## rate_limit.ip_allowed_list — ip_allowed_list / 212032001130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.ip_allowed_list

<a id="canonical-3103033003311002-1200130211131302-2213201303103331-3031320123113330-2021311303122332-1231311010022303-3230112212230320-1002133121102033"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
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
ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103020102322010-1221303131111113-0002031012113002-1202233321000201-2231202130230000-3330223123001323-0013213113331330-1102202001221112"></a>

## Direct properties — ip_allowed_list / 212032001130 / 3

<a id="canonical-2131112302013213-0233210323130101-3310001232032331-3003322021222112-1333002311223101-0130300103011221-2200312331110100-2132302231213031"></a>

<a id="canonical-2231010221022330-0003303113323113-0002122303203033-2321112320031030-1122200210303003-0000300222031003-2322202022130023-0023023202132010"></a>

## prefixes property — ip_allowed_list / 212032001130 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3113301333003100-1302300011130200-0213132120122231-1132113313203121-3130212132031233-3211230311202110-2311223100103013-2032033232201302"></a>

## Next pages — ip_allowed_list / 212032001130 / 5

- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3010321303322033-1212332122021131-3013312103113211-2131102032023233-3123302121323033-2320112001033113-3312022230012033-3102230311220002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131133222132002-0333223231232002-0120121212133003-1202013120023313-0011133021101021-0331201222100300-2102013031003023-3231222021322010"></a>

## rate_limit.no_ip_allowed_list — no_ip_allowed_list / 323133102333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.no_ip_allowed_list

<a id="canonical-0122110212212020-1100201100100210-0122123233122013-0113231212102312-3022132010230013-2210233332010021-2121231233211120-2201203212303220"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
no_ip_allowed_list = {}
```

<a id="canonical-2333303131202202-1023303322033010-2123332010311020-3113213000120321-2110322233330313-3130123122033011-2212302323302213-2222010313210000"></a>

## Direct properties — no_ip_allowed_list / 323133102333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020120030103310-3123233003132013-1120023003113103-3331202310232220-1331132131131000-1302221112310221-1211133202032012-2332133221032232"></a>

## Next pages — no_ip_allowed_list / 323133102333 / 4

- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1012222010200213-1013130113011130-1203230321210101-0313310220312312-0311230001310213-1113313130100101-2020020003003112-3312012120123223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011012132201231-2100131112310310-1321010133023110-2333132101020132-3031210321020213-2010222113022330-2033031333003212-0303330033312032"></a>

## rate_limit.no_policies — no_policies / 330233331001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.no_policies

<a id="canonical-2311111332132033-3203120133203120-3223003203201302-0303213122001132-0320202211030100-0110123003333110-1313212021323313-2201333130013200"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for no policies. Defaults to \`map\[\]\`. Server applies default when
omitted.

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
no_policies = {}
```

<a id="canonical-1221332201013022-3231130103000000-3232123322210332-0012123123332110-1133202020322102-3110233233321023-0123202331111233-3011331010130002"></a>

## Direct properties — no_policies / 330233331001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022311320001123-0200113330333022-1012120121011231-3030223211030233-2022232031032001-1211323030001313-2102113302133122-3203031312201132"></a>

## Next pages — no_policies / 330233331001 / 4

- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1201121223001130-0102301012320112-0332103302021323-3120320310211121-2331102322201132-0130330012330030-3122013001202010-1320120013120101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101003312111121-0030323102020111-2223210103110010-2032320122322332-1013021110103233-1030221133320233-1001220333031121-0221222313111012"></a>

## rate_limit.policies — policies / 322232030112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- rate_limit.policies

<a id="canonical-1311212123020300-1202120201123332-1120101331132123-1310031310211112-3323003212013030-3102221103030333-1321333122323331-2122022321132103"></a>

Type: `"object"`. single nested block, Optional.

List of rate limiter policies to be applied.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("policies")}
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
policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2110123220100301-0113231300201112-0200313220032030-1221220123021120-3010230311031333-2223010022320323-3222311020133233-0323033303002331"></a>

## Direct properties — policies / 322232030112 / 3

- [policies](resources--http_loadbalancer--reference--group-023.md#canonical-3101001310133003-3001131113313331-0130031131012332-3322332320111210-0220323211221110-0202113030011302-0120211202101330-2132032321310302): complete subsection reference.

<a id="canonical-2202012203133110-3231112210322220-1311220012222130-1331320230103023-0221321130012200-2311310222011332-2113212122013202-2200101221131310"></a>

## Next pages — policies / 322232030112 / 4

- [rate_limit.policies.policies](resources--http_loadbalancer--reference--group-023.md#canonical-3101001310133003-3001131113313331-0130031131012332-3322332320111210-0220323211221110-0202113030011302-0120211202101330-2132032321310302)
- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3101001310133003-3001131113313331-0130031131012332-3322332320111210-0220323211221110-0202113030011302-0120211202101330-2132032321310302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303122020201103-0123320031123230-3302222022311100-3213131002331312-1322223033101313-1120030203222033-0232120120023100-2232123221113032"></a>

## rate_limit.policies.policies — policies / 002133321200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [rate_limit](resources--http_loadbalancer--reference--group-023.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023)
- [rate_limit.policies](resources--http_loadbalancer--reference--group-023.md#canonical-1201121223001130-0102301012320112-0332103302021323-3120320310211121-2331102322201132-0130330012330030-3122013001202010-1320120013120101)
- rate_limit.policies.policies

<a id="canonical-1132311333313033-3331110210011121-0223123000200333-3131030320223011-0133301132213021-2311122120312201-2033301003302032-2133001010133022"></a>

Type: `"object"`. list nested block, Optional.

Rate Limiter Policies. Ordered list of rate limiter policies.

Upstream description:

Ordered list of rate limiter policies.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130002201210121-0201200011033220-0100203302213103-1211303202001300-1220211122303201-1312313330311211-1123030110000203-0220332233033331"></a>

## Direct properties — policies / 002133321200 / 3

<a id="canonical-1013132032310330-2010231113122122-1201110122301100-1232130121033210-3122233031300013-1033112103031230-0320233120203001-1211002012202131"></a>

<a id="canonical-0123321323002021-2010033221301223-0323213311132313-1011122102201230-3311031020233322-2012201112123130-2313023230323301-1233020233000001"></a>

## name property — policies / 002133321200 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1033103333211331-3112201311023132-0210310112011103-2030202100012021-1023201033300231-3203213012210211-3230211221302121-3013023022131300"></a>
