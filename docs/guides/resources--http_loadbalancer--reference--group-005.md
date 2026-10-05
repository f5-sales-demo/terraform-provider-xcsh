---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2020112123211133-1023000111210131-0210323220032233-0320021121233113-0211210123123323-0332121320002033-2123030313023102-2331133203112001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300302130123302-2100321032100221-2232322131131303-0323233001300213-3312231022130231-2331001310333233-2022023011102023-2303310220101111"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.any_client — any_client / 111112213332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.any_client

<a id="canonical-3031103002200100-1311331012331303-0320322210010301-1112303312111201-3210220312110113-1330323033213103-2121013021003130-1233030111311331"></a>

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
any_client = {}
```

<a id="canonical-1000322300013303-3011303001220212-1011323333230231-0000313221331310-0231032031013111-0130021002233211-0022223220111133-1303023033210113"></a>

## Direct properties — any_client / 111112213332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122012002002033-3111100211200221-0332033322322102-2031311130313032-1113203031010120-2331332113210013-1201302131003111-2222220123322202"></a>

## Next pages — any_client / 111112213332 / 4

- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0301101110310301-2100023033303011-3000032120233130-1232013222023103-0003000030002011-2212302221213122-2022311131031000-2012132302012100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331121131332001-3223102020022131-1011001301301133-2200232020230010-0001123330000300-1312110101320231-2201233012332322-0133213013213313"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.any_ip — any_ip / 303000233320 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.any_ip

<a id="canonical-2201121301222003-1010010011102103-1312323212210021-1032112121101111-2012303103202230-1021013302022331-2101310202131003-0213021222022123"></a>

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
any_ip = {}
```

<a id="canonical-0333111111322231-0201013133110023-1023103311000100-0102012321000031-1212233230223003-0322112131013323-2131012031301331-0320130021022310"></a>

## Direct properties — any_ip / 303000233320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302022022012200-0332121313130311-3310201211130131-2130330122133103-1323301203030232-1332120131123310-3133200033203221-0210222132132320"></a>

## Next pages — any_ip / 303000233320 / 4

- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0202132123233123-2323223223222103-3031333211132130-2101222132133323-2221111032321223-3220230132112233-0312302321223103-3110321303011201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302301300121001-0001212233230031-2203300311002331-0310001002310213-2000011020203301-1100232211231101-0022112221131301-2201133032233220"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.asn_list — asn_list / 032313022333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.asn_list

<a id="canonical-2332111321002113-1000030003333131-2132032213233101-1202221212021213-3021313130333212-1201000030203322-0032313220103003-0310122102230133"></a>

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

<a id="canonical-3021321021302010-3131100323331311-1302300013333202-1200323313332202-3110311330013031-2032220121032311-0200103100113130-2312200313021330"></a>

## Direct properties — asn_list / 032313022333 / 3

<a id="canonical-2331031203312321-3102032102322331-1021330303311032-2011110121020133-0130303112332320-3322302011312003-1123232020223333-3303220311231330"></a>

<a id="canonical-2101330013130121-0320300320002230-1110031213221301-1321013002320133-1212302033021201-0301032333122312-1202320303122312-0101111311211133"></a>

## as_numbers property — asn_list / 032313022333 / 4

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

<a id="canonical-2022132331310232-3111310200011320-1300103011111203-0011221123202311-3222003012310233-1032132210210203-2230021221102213-2221333232223113"></a>

## Next pages — asn_list / 032313022333 / 5

- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0031223033323013-3010133103031132-2001331003300213-1201212302233110-3201001010203200-2331202112301022-3212210132133031-0021231310120320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010300312322320-0033102232302012-3112223001222030-2032032030332312-3031012321103200-0120032002031031-3231020211032100-0011210301300133"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher — asn_matcher / 012303200001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher

<a id="canonical-3301031000222302-2231013312113202-2031010211100133-2313020332310001-1321213212032310-0302110000002331-0230012233232201-0103203331012212"></a>

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

<a id="canonical-1101312313011102-1221200232023213-0232331013211201-0032021300201122-0100221121032010-2012032311203100-0331110011130321-1213103333111033"></a>

## Direct properties — asn_matcher / 012303200001 / 3

- [asn_sets](resources--http_loadbalancer--reference--group-005.md#canonical-2112310301111211-2102332321023101-0013312220203213-3003101033322210-2322232103012200-3023221012330200-0321113201301201-2033232032212122): complete subsection reference.

<a id="canonical-3102220203111313-2322320023302011-2203121113203331-0202332133010331-0120033230302233-0332323322113233-0122033203203110-0020033201202201"></a>

## Next pages — asn_matcher / 012303200001 / 4

- [api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets](resources--http_loadbalancer--reference--group-005.md#canonical-2112310301111211-2102332321023101-0013312220203213-3003101033322210-2322232103012200-3023221012330200-0321113201301201-2033232032212122)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2112310301111211-2102332321023101-0013312220203213-3003101033322210-2322232103012200-3023221012330200-0321113201301201-2033232032212122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222300121131021-2020003110231211-1320131311120102-2311300121111031-0121120310102302-1212302311330002-0211220200130032-1012213230010332"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets — asn_sets / 223213121001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- [api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0031223033323013-3010133103031132-2001331003300213-1201212302233110-3201001010203200-2331202112301022-3212210132133031-0021231310120320)
- api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-2313020333100121-3211012003302233-2132200312112133-1232130330000111-2233301100221102-2202012132220231-1110333232230203-2003332133302331"></a>

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

<a id="canonical-1012012232213312-2022110120031000-3012121303220310-3010320331211323-2320310323320302-0313032022012021-1222322132301320-1003021131320331"></a>

## Direct properties — asn_sets / 223213121001 / 3

<a id="canonical-0010023200320321-2131101122013313-2121103032112132-2113310233220210-3322132331211033-3000101302323201-0210210321322023-0121322320101203"></a>

<a id="canonical-2012200230312232-1131232131212201-1133322331103120-1010023032023201-1022230212011123-1332102033301110-1013100321121111-1330101121231202"></a>

## kind property — asn_sets / 223213121001 / 4

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

<a id="canonical-2211100321031222-1232201200120333-1230103102100300-1302031312023331-1120033210231310-3322030311110221-1033222130010003-2002331103101032"></a>

<a id="canonical-1130323000102131-1330013001321021-3122130123212311-2113103133023103-2332022131310021-3023001212330203-3220300311202320-3231013300102000"></a>

## name property — asn_sets / 223213121001 / 5

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

<a id="canonical-2113231310221003-3301302102331133-3110122220022012-1213202330013313-3030031130221013-3320023232323311-3212132122013210-1301331013230012"></a>

<a id="canonical-1112100333312202-3113202011213121-0323012203230302-0113231203103110-1312122000020232-2130200313323231-2302301013110331-2321121000220212"></a>

## namespace property — asn_sets / 223213121001 / 6

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

<a id="canonical-3131013032033312-3300303110313133-0023310333101200-0220010333013332-2101121123132132-3223013211030131-0133100020312210-3000020230232210"></a>

<a id="canonical-1310021212231100-1103301012231331-2200112301302133-1121111010321021-3013233123320303-3223012311301002-0011100212023123-0123331022002030"></a>

## tenant property — asn_sets / 223213121001 / 7

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

<a id="canonical-2330333113232333-3332300200030230-1300010201101301-0332220302230021-2130222202102030-0321022003200213-1213013122031211-0310232202031333"></a>

<a id="canonical-0333021013130001-0322310223021210-1323332011000130-3113012120121110-0202020113110212-0302020020022021-0200112313130232-2113001122101312"></a>

## uid property — asn_sets / 223213121001 / 8

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

<a id="canonical-2132311232300032-0320013201330133-2231223022102003-1111201102021223-3201133013332332-2133121300120013-3031112320223123-2230233330222102"></a>

## Next pages — asn_sets / 223213121001 / 9

- [api_protection_rules.api_endpoint_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0031223033323013-3010133103031132-2001331003300213-1201212302233110-3201001010203200-2331202112301022-3212210132133031-0021231310120320)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1131312220313133-0012332222221011-2103002010311130-0331210300030032-0122301313332302-2111022302232122-3102002331002032-3332302033220321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323221100033233-1320121310002012-3312331231212023-3012212301332031-3123323110202132-3032302121220220-0022313203100323-0312133332333131"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.client_selector — client_selector / 212000200211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.client_selector

<a id="canonical-3022222323110220-3200031330320231-0031211202011120-3321022321012022-3322222133121013-0302000331333332-0310132323203331-1012031020023111"></a>

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

<a id="canonical-1212101211113211-3332333020222112-0100032321333113-2222311333212123-2231103010200020-0200013021013120-3321032331022302-3130001031232022"></a>

## Direct properties — client_selector / 212000200211 / 3

<a id="canonical-2230210121001100-1203130203320233-2223332303013300-1223020103313030-1201011130113231-3132230012001321-1320002002221031-2010221101201021"></a>

<a id="canonical-2022120202223230-3120232232323311-0130131031102302-1133133130032211-2033021201222300-0031220122132321-0133023200311131-3101310032232230"></a>

## expressions property — client_selector / 212000200211 / 4

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

<a id="canonical-2103330312133022-0200020110120313-2120022120002032-0321203210222321-3213303312001132-3322312100213333-2300213221200212-0100213110103221"></a>

## Next pages — client_selector / 212000200211 / 5

- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3210032110232021-3220012112131310-2021133120300231-3321232103230222-1210021330132333-3121002223132100-2222112232130313-1223100210133000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120021123331331-1032323202303133-1022332033331001-2310023202011121-2212230302303001-3321210020120221-1120122320032211-3130202213020102"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher — ip_matcher / 311322212000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher

<a id="canonical-2023223100033321-1201032330310033-0012000112020210-1003100212011221-0203303322223110-3201221130110120-0003323031103013-1210020130032332"></a>

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

<a id="canonical-0333232203313303-0201003333213310-3212100231030111-0120312210330001-0230102121212200-2311203131011311-0101101022130231-2310011322222101"></a>

## Direct properties — ip_matcher / 311322212000 / 3

<a id="canonical-0130301313110102-2103303311102330-3011232322102033-0121102232320102-1302110212212031-1300032223000103-1220330223101003-3000212131120332"></a>

<a id="canonical-1220032010302330-0202322031001102-3201102203323310-0122332311331033-2000013002201212-1321001021123230-2300203322203021-3133002112130221"></a>

## invert_matcher property — ip_matcher / 311322212000 / 4

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

- [prefix_sets](resources--http_loadbalancer--reference--group-005.md#canonical-2221103031301100-2211323222101322-3031312320131102-3030033232122310-1132001032023320-3200003303130130-1130032101123233-2232233310122030): complete subsection reference.

<a id="canonical-3012303213033322-0333323302323230-1030101331323323-3121123101200002-1201302012021113-3231130301130132-2023122230121003-2113031113012112"></a>

## Next pages — ip_matcher / 311322212000 / 5

- [api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets](resources--http_loadbalancer--reference--group-005.md#canonical-2221103031301100-2211323222101322-3031312320131102-3030033232122310-1132001032023320-3200003303130130-1130032101123233-2232233310122030)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2221103031301100-2211323222101322-3031312320131102-3030033232122310-1132001032023320-3200003303130130-1130032101123233-2232233310122030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230230210003030-0233312113200001-3220300211320011-0230103233212022-0020220223103020-1312233012033020-3321331031100313-0122131323201331"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets — prefix_sets / 000033320033 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- [api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-3210032110232021-3220012112131310-2021133120300231-3321232103230222-1210021330132333-3121002223132100-2222112232130313-1223100210133000)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-0023021231110102-2131131232131030-2321220303001211-1023111330132130-1023120123112021-3313331001131312-2201103200301102-1002320303301021"></a>

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

<a id="canonical-1323321100332323-1311031211100111-3220333002300031-2033111301010310-3031221212032311-3012023312021331-1030133101331210-2022301110120331"></a>

## Direct properties — prefix_sets / 000033320033 / 3

<a id="canonical-3210121103230132-1321220101013312-0122121003331233-1200111013002330-0331322111003101-0132321112213112-3312031222301023-3301200122020110"></a>

<a id="canonical-0020301023021132-0202131012022112-3320010022011110-3001200021211311-0003122020023231-0210112321132102-1120311221222030-3233112323201120"></a>

## kind property — prefix_sets / 000033320033 / 4

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

<a id="canonical-0000312320332133-1111221133102211-2203323100132223-1122202223233233-0302023203002321-2100301113013113-3303203132101122-2103221312011333"></a>

<a id="canonical-3101112122002031-2320110231033000-0001111100301111-2113113010212232-0112012302112212-1212022010030021-1032310322033121-1301012023221001"></a>

## name property — prefix_sets / 000033320033 / 5

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

<a id="canonical-3310101210121132-2220022122022322-2021130312100212-2031312033020000-1032010030201323-0101123011102323-0031112330120010-0101320011210222"></a>

<a id="canonical-3301333321133112-0332232031233022-1001020220312022-1210033330230321-1100020311133101-0131130213303112-0112230201133203-2023013332312321"></a>

## namespace property — prefix_sets / 000033320033 / 6

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

<a id="canonical-3232122302211223-1222101221231102-2012130202111310-0313332132230030-0022120202121132-0303312010020123-1311211001122000-1000110233012203"></a>

<a id="canonical-2310031220033103-0201133001220100-3111001001232002-1223121300020303-2220321313333322-2211130331200301-3323311011111033-1310112011032303"></a>

## tenant property — prefix_sets / 000033320033 / 7

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

<a id="canonical-0132300210132013-2131330031030121-2032110132131123-1023223102123220-2133213233320301-0220222211133000-0310003200010200-3123210130130321"></a>

<a id="canonical-3313201202203121-0323213122310220-1120030230122020-0201011133322211-3200110211222030-2230002303213332-0121202102103110-0313003323113332"></a>

## uid property — prefix_sets / 000033320033 / 8

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

<a id="canonical-1202012033313222-1030022133221232-3220002010322033-1230121003013032-2020331120313121-0112302032321102-2113000233300330-2331213212113033"></a>

## Next pages — prefix_sets / 000033320033 / 9

- [api_protection_rules.api_endpoint_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-3210032110232021-3220012112131310-2021133120300231-3321232103230222-1210021330132333-3121002223132100-2222112232130313-1223100210133000)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1133120210023121-0002310111331330-2020101032033222-1121331003220303-1313232312313110-0110223113221122-0321132113112223-1222031231002301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011201011102120-0101310322230101-1332223130120210-2031211032101232-2332011033033323-1302303123230012-3313203003133321-0101020001220113"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list — ip_prefix_list / 330103332032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_prefix_list

<a id="canonical-1213313213202111-2112231220020213-3103032112113201-2120233322230123-0030021110322203-1010112010310210-2121222330232321-2330322223011311"></a>

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

<a id="canonical-2003103023130023-0101001033330012-3103003010002101-3300313103321331-2210212233222302-3333013101030031-0122112212002223-2320201231213001"></a>

## Direct properties — ip_prefix_list / 330103332032 / 3

<a id="canonical-1111011111122311-1223000123011011-0321131212103222-1211103002101113-3111112203313133-3003033201011103-2202331220210131-3301010323101022"></a>

<a id="canonical-1020210110002201-3000311223002302-2130210321133101-0123110032223312-1010322111321310-0131021300111131-0002002032203012-3122330101321323"></a>

## invert_match property — ip_prefix_list / 330103332032 / 4

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

<a id="canonical-0323232012023210-3320333133003221-2123111113030231-0320202200211203-2010002030100003-2103330130031003-2331121012302201-0222012111001121"></a>

<a id="canonical-3212022103020113-1023231213322330-3200202202122322-0113103111031111-2221232200030102-0333232010301301-3131311022211301-1111232202100233"></a>

## ip_prefixes property — ip_prefix_list / 330103332032 / 5

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

<a id="canonical-2302321023121211-2223233313301011-3330303112033231-3111003123301001-0032222010123232-2010131203100202-2122211020323110-2312013330013020"></a>

## Next pages — ip_prefix_list / 330103332032 / 6

- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0123120001322103-2331123203130211-0113111100023101-3012213212000010-2222131223033123-0123212321303310-3101133011113133-1120011212302030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201211212002233-1321123311031032-1321301100302333-1123102110112103-1320201012310233-2132011111332212-3321312112120322-2330321022321220"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list — ip_threat_category_list / 301333133010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.ip_threat_category_list

<a id="canonical-2102001233030121-1220220232203000-0303232312230031-3331033221120031-1112203023120101-2313302302200311-3210103021003211-3323202102310013"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
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
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3120022321203313-1013133300300032-2323333233132302-1200022133302312-2231010321313101-0133223001130300-3113133300210202-1212100322112033"></a>

## Direct properties — ip_threat_category_list / 301333133010 / 3

<a id="canonical-1010123030000012-1321330122302303-1011323201130133-1230011200320330-2233310202131020-1131031232231000-1222123321132030-2131223000131031"></a>

<a id="canonical-1203221201123311-2002121300321023-2132332202302313-1323013032203020-3022221320223100-0323311132022032-1101111331322110-2103031132302303"></a>

## ip_threat_categories property — ip_threat_category_list / 301333133010 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Upstream description:

The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3112002113332321-1300223201011100-1001132312221123-1120022303130300-3330221331111202-2022032132323020-0103003112133233-2002201200312131"></a>

## Next pages — ip_threat_category_list / 301333133010 / 5

- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0130103020102113-3201320201003312-2033202202012021-1303212231222321-2231121221332103-1331011032100122-3110323111300113-2003100223313321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001131032010113-3202300111331203-2001202210323221-0132300313233103-3303230303032320-2000101130021322-0103310022221233-0112102203121322"></a>

## api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher — tls_fingerprint_matcher / 120033321002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- api_protection_rules.api_endpoint_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-3003010103031122-0133303310112333-3110312110333000-2331211002103131-0323233332200311-1211130031302223-1013213003022321-3022213232122010"></a>

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

<a id="canonical-2012111102223212-2131102012030320-0312110223133301-3323210012023122-1033231301020131-3212212310333122-1100332223011012-3302233112003320"></a>

## Direct properties — tls_fingerprint_matcher / 120033321002 / 3

<a id="canonical-0200310333303230-1022300313333130-0321331210333020-0010132113110120-3122130201313332-1021322233201112-2201200322222213-2212200233102022"></a>

<a id="canonical-2203112022032322-1030023221112222-0323012102113301-2113311000322133-1011302011102120-0320122001000322-3300101120202201-1211303110311320"></a>

## classes property — tls_fingerprint_matcher / 120033321002 / 4

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

<a id="canonical-3132331233222303-3031331301121032-0033122333203330-1230202233132132-2130103210231112-1303033220323112-0300110333000232-2201200030123303"></a>

<a id="canonical-1311333200331130-0010301011122211-3110012000022003-3001222201211122-3312302100131111-2020320310023100-0111122031203110-1300301011301010"></a>

## exact_values property — tls_fingerprint_matcher / 120033321002 / 5

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

<a id="canonical-2222302231313031-1013232232133100-3122000023310323-3213132110332130-3230122300011102-1302030330212032-3232031331102311-1113010013203313"></a>

<a id="canonical-1022133130332111-0012113030130033-1033323210011113-3131022032101300-3132123021213033-3102130112031323-1303102303302303-0130001311030201"></a>

## excluded_values property — tls_fingerprint_matcher / 120033321002 / 6

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

<a id="canonical-0203133231101101-0211221112013022-3212310322300002-1113023001310123-1203101123033120-0013301103012311-2101212333003120-2111222313022313"></a>

## Next pages — tls_fingerprint_matcher / 120033321002 / 7

- [api_protection_rules.api_endpoint_rules.client_matcher](resources--http_loadbalancer--reference--group-004.md#canonical-2300110100200312-2211000003111012-0121223331223212-2323033211320332-1011120301332210-2221232302021210-0230132330021023-3112222110211201)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3030001110033000-3131310301001101-1113220222331313-0003111330222130-1113022020310230-0321012302132102-3000133202210303-2332233100122302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113311220033133-1020102310011302-1331320201030323-0033202223011221-2220021130222001-1010312133020012-2230330103130132-2132111033111213"></a>

## api_protection_rules.api_endpoint_rules.metadata — metadata / 022100023330 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- api_protection_rules.api_endpoint_rules.metadata

<a id="canonical-3303332130023002-1210023233020120-0332012032220231-0123022213003113-0211023101212320-0312300211102212-1200033332132120-2210200030110213"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031013033210031-0001222122202203-3011113122233223-1031330321131322-0221330331331333-1312231311330231-0321303030213213-3223000212101112"></a>

## Direct properties — metadata / 022100023330 / 3

<a id="canonical-0211121322110210-1330302113001113-2120220110103121-3132100331121030-1322012123221023-3231113310032222-2232223300313121-2012330233100120"></a>

<a id="canonical-3031302022112111-3131130310303130-0010203332201302-3322213120110232-2333003311120213-1210112103200331-3211131213112333-1231033132223203"></a>

## description_spec property — metadata / 022100023330 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1011322212223232-2220213120232231-1033320200223123-1301130323213233-0103201323320223-1222001021021201-1323303003230010-2212000011000131"></a>

<a id="canonical-0110333013231123-1230211301220302-1323111120110321-2002223213223221-3233010112210300-2222313111132031-0123130032310113-0202323011323231"></a>

## name property — metadata / 022100023330 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-1302223311131200-1220222322223003-3211101002012310-1233000013202332-0111220311300232-0012213222213330-1201112121221023-1211130300222023"></a>

## Next pages — metadata / 022100023330 / 6

- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232103210022201-3202323201131113-3333103330022132-2030312320100232-2223032233002001-1230123321233312-1100313313132213-0112221130101130"></a>

## api_protection_rules.api_endpoint_rules.request_matcher — request_matcher / 000113100000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- api_protection_rules.api_endpoint_rules.request_matcher

<a id="canonical-2233000333303300-0301120300310222-0132201110222011-2202021231010323-3121323112213221-0112231031331121-3112330331221321-0012120203112000"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request matcher.

Upstream description:

Request conditions for matching a rule.

Receipt-pinned upstream constraints:

```json
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
request_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103022213033103-0321133223103203-3323130002123332-3330330122130321-0302020322202100-0031202100231222-1213233111130301-0300110112010302"></a>

## Direct properties — request_matcher / 000113100000 / 3

- [cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-3021011121230231-1331122033331212-0313313111122100-1301200233032301-0221003011133232-1311321302033210-2332120010212101-2001120223300332): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-005.md#canonical-3232223131110320-2312312311330110-3211012310212200-2133330302303310-2030221113011302-2300102101120101-2020022231332232-1313231311122212): complete subsection reference.

- [jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-2222110200102020-3020200322033310-3003111203122333-1132303221211201-0330031100103113-3101011012122301-3302212223030312-1230213330301033): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-005.md#canonical-3312230121022111-1002311110212233-0111012233200211-1310312023102312-1113111113123303-1302201110232311-0310332111123032-3221103322120332): complete subsection reference.

<a id="canonical-2200223203211213-3302320320233201-2313011130121232-1113230013311320-1122112102031232-1020111222332001-0031033012311113-1200133213110013"></a>

## Next pages — request_matcher / 000113100000 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-3021011121230231-1331122033331212-0313313111122100-1301200233032301-0221003011133232-1311321302033210-2332120010212101-2001120223300332)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-3232223131110320-2312312311330110-3211012310212200-2133330302303310-2030221113011302-2300102101120101-2020022231332232-1313231311122212)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-2222110200102020-3020200322033310-3003111203122333-1132303221211201-0330031100103113-3101011012122301-3302212223030312-1230213330301033)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-005.md#canonical-3312230121022111-1002311110212233-0111012233200211-1310312023102312-1113111113123303-1302201110232311-0310332111123032-3221103322120332)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3021011121230231-1331122033331212-0313313111122100-1301200233032301-0221003011133232-1311321302033210-2332120010212101-2001120223300332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313310332033120-2020311333202021-3323231202121121-0333313032010210-2121100132010203-1010131033330012-3201033010313301-1322300130212321"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers — cookie_matchers / 112332021233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers

<a id="canonical-2020113010323330-2202100122302102-0131221301333213-2033211023223023-1012112321003100-0320232123022211-1323313310212102-2211300003132333"></a>

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

<a id="canonical-0023112003031221-0002003002331132-3121332213201301-3321103010332210-2333301202310101-0120103331010211-2331222110323210-2112003222033301"></a>

## Direct properties — cookie_matchers / 112332021233 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-3210232211323311-0023211102021213-1311222011113111-2222132232210022-1212023013010212-1223233112323033-3232131023020133-1300313010311103): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-005.md#canonical-3220223121012213-1031000033003030-1020331010300303-0223030311111023-3123131132203210-3012102330223330-0103203101020022-2023201331302030): complete subsection reference.

<a id="canonical-0011030021121320-1101130201231113-3130320102101323-3031200233302222-3210302310130221-1300132320203303-2002000212330022-0130101020303211"></a>

<a id="canonical-3311021220222210-0013303012300202-1132323122312120-1332123133302022-3210232112102020-3312123013130203-0013110232313321-1030323221231301"></a>

## invert_matcher property — cookie_matchers / 112332021233 / 4

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

- [item](resources--http_loadbalancer--reference--group-005.md#canonical-1002133022121033-0201122130133302-1123103100203310-1030022303002202-1123330300122213-2320111112202230-2113233131230131-3301011321013233): complete subsection reference.

<a id="canonical-0313203100320011-1322210333130301-2310220331110311-0230132203201323-0332031123122033-1122232322000212-1110021231222320-1321121321331130"></a>

<a id="canonical-3202232121131031-1022233213331210-0200122032303020-1113211121223000-2122013220032132-3333200021210212-2003210200003203-2301112200002002"></a>

## name property — cookie_matchers / 112332021233 / 5

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

<a id="canonical-0030010310332130-1020323101302221-1022121131030003-2302232211130222-3301300023231201-1102023121000132-2032001322322312-2101211113231030"></a>

## Next pages — cookie_matchers / 112332021233 / 6

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-3210232211323311-0023211102021213-1311222011113111-2222132232210022-1212023013010212-1223233112323033-3232131023020133-1300313010311103)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_present](resources--http_loadbalancer--reference--group-005.md#canonical-3220223121012213-1031000033003030-1020331010300303-0223030311111023-3123131132203210-3012102330223330-0103203101020022-2023201331302030)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item](resources--http_loadbalancer--reference--group-005.md#canonical-1002133022121033-0201122130133302-1123103100203310-1030022303002202-1123330300122213-2320111112202230-2113233131230131-3301011321013233)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3210232211323311-0023211102021213-1311222011113111-2222132232210022-1212023013010212-1223233112323033-3232131023020133-1300313010311103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031133310312020-1103123011003311-0212122223133032-3000122021103330-1200130200231312-0131230333332031-3233001021002332-1300022302322203"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present — check_not_present / 010032030323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-3021011121230231-1331122033331212-0313313111122100-1301200233032301-0221003011133232-1311321302033210-2332120010212101-2001120223300332)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-1213233301123201-2311033223131311-1002300330110212-2002010133013100-3313113322001203-0033101112110002-0010323313332330-0002011232332003"></a>

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

<a id="canonical-2021111233312331-0121032231131123-3232221322011031-2112023112303223-1031030322113030-1101131311010020-3300133323123200-2112021113312020"></a>

## Direct properties — check_not_present / 010032030323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223120022031301-1323022112202211-3230023310010131-2130101320100202-0001001010222313-0023033212231202-1320321122203221-3001233230331330"></a>

## Next pages — check_not_present / 010032030323 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-3021011121230231-1331122033331212-0313313111122100-1301200233032301-0221003011133232-1311321302033210-2332120010212101-2001120223300332)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3220223121012213-1031000033003030-1020331010300303-0223030311111023-3123131132203210-3012102330223330-0103203101020022-2023201331302030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233322302102203-2213323312020011-2310123121232203-1002010112231022-0010032122003311-1033121101323021-3201333222133200-1302321330223030"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_present — check_present / 111311202200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-3021011121230231-1331122033331212-0313313111122100-1301200233032301-0221003011133232-1311321302033210-2332120010212101-2001120223300332)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-3120102301323031-2231320202031312-2300132322032322-0321002011321202-3320232300323303-3022302311210021-0301113323330013-3131031003312313"></a>

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

<a id="canonical-1300230121021303-0113003321230000-3223321102330220-0301322132120222-3331333003122001-1112100212313013-3303301331110113-0132032130111302"></a>

## Direct properties — check_present / 111311202200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101001130302000-3231131331310031-2023113303001133-0311322323003320-3031032003332031-3101021031002330-0212031332130322-1333122211300223"></a>

## Next pages — check_present / 111311202200 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-3021011121230231-1331122033331212-0313313111122100-1301200233032301-0221003011133232-1311321302033210-2332120010212101-2001120223300332)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1002133022121033-0201122130133302-1123103100203310-1030022303002202-1123330300122213-2320111112202230-2113233131230131-3301011321013233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221112013011231-0021002202210132-2300300110310302-1313022331310021-2222220323321030-0021211023030133-3203321211011112-1113330101121213"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item — item / 011132123101 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-3021011121230231-1331122033331212-0313313111122100-1301200233032301-0221003011133232-1311321302033210-2332120010212101-2001120223300332)
- api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers.item

<a id="canonical-2033221102320210-1031212223002101-0320112131312012-1123322112312011-0011101101212120-2220301101202010-0202212011220303-1230323022203030"></a>

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

<a id="canonical-3013122100110322-0031110030101120-0002222331233103-2003111111303320-2302022300103220-0321032012002032-3130122103010100-2020303013203232"></a>

## Direct properties — item / 011132123101 / 3

<a id="canonical-3301200000132301-0011122023212320-2332003321022223-2201022212112233-3110101331331300-3300302231010201-0223222022201321-2221000202213112"></a>

<a id="canonical-2323122322122233-2101133132110301-3030221033300131-2013300213312330-0322000231001321-2000222321000231-1132020301130113-1033001200323312"></a>

## exact_values property — item / 011132123101 / 4

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

<a id="canonical-0233213212330312-2011021020323311-0222103312021022-0131011222323001-1210023113331100-0023122033313220-2022023210223133-1201212031133233"></a>

<a id="canonical-0002333112213203-2102311323200032-1202020002210123-0302312031033320-1323230300332023-2111003313001333-0321321110323000-0323322102133220"></a>

## regex_values property — item / 011132123101 / 5

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

<a id="canonical-1223030110321032-1300313300322313-3233100132011131-0233013133223121-3210201131001212-2202323221002103-0201212121312231-3300120121133130"></a>

<a id="canonical-0200130103213321-1013333102312303-2100013313222033-3212013203220003-3000002110012221-1103331211233010-0213023102310031-2013212210123311"></a>

## transformers property — item / 011132123101 / 6

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

<a id="canonical-2010330211010220-3231320103212023-1002030003133223-1220223122010133-0102003232022332-3213111031232301-1012202121221020-0331123023233111"></a>

## Next pages — item / 011132123101 / 7

- [api_protection_rules.api_endpoint_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-005.md#canonical-3021011121230231-1331122033331212-0313313111122100-1301200233032301-0221003011133232-1311321302033210-2332120010212101-2001120223300332)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3232223131110320-2312312311330110-3211012310212200-2133330302303310-2030221113011302-2300102101120101-2020022231332232-1313231311122212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103122233202102-3102120221320010-2231133100012333-1102221033102003-2100202113221311-1232331321013232-3130030333111022-3133322132103332"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.headers — headers / 111333330102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- api_protection_rules.api_endpoint_rules.request_matcher.headers

<a id="canonical-0222330302012330-1221032300020231-0003201103223231-1123333330330003-1201212020333332-2122120203003303-0123331203010223-3021131331010113"></a>

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

<a id="canonical-2201110020301231-2100311331231303-3001233132103121-3311133232230002-0300323213113233-2120003021312230-3120021311002001-1330031322131121"></a>

## Direct properties — headers / 111333330102 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-1131212122002031-3301103113313301-3123113013200331-3210203100230333-3101332233223223-0002000322110020-2303130211303320-3113123122202203): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-005.md#canonical-0303123022031230-0210331320202313-3011120333111210-2311131130302031-1322331131013123-0302110202112301-2311331021211312-2102000333231113): complete subsection reference.

<a id="canonical-0130222202023013-1211212211003132-0202203101112313-3310221012023020-3330233333213031-3310031110302212-1302201320131103-0133220110231210"></a>

<a id="canonical-2023111012302212-1230021123220131-3230002021032110-0131131101033322-0331203331030321-1033231320302330-3112031002221310-0110033122121010"></a>

## invert_matcher property — headers / 111333330102 / 4

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

- [item](resources--http_loadbalancer--reference--group-005.md#canonical-0222202333231133-3130233100122132-2001300113002031-1222000230031121-0033300020310102-1102201133102220-2001302332231031-1201030012220110): complete subsection reference.

<a id="canonical-1232102201302331-0011012103220103-3020112300313311-0132332021023102-1211210000101122-3200323101002220-0212122333031212-3113310123333020"></a>

<a id="canonical-1321031213002020-3003330131013013-3323012211132132-1323103132232203-3000132301213032-1203331320002223-3033023003020313-2003303210303230"></a>

## name property — headers / 111333330102 / 5

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

<a id="canonical-3102331122203112-0220113030101032-3230212302320201-2030200211222103-0232312102312230-1321112121131321-1230033210121100-1212103322023110"></a>

## Next pages — headers / 111333330102 / 6

- [api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-1131212122002031-3301103113313301-3123113013200331-3210203100230333-3101332233223223-0002000322110020-2303130211303320-3113123122202203)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present](resources--http_loadbalancer--reference--group-005.md#canonical-0303123022031230-0210331320202313-3011120333111210-2311131130302031-1322331131013123-0302110202112301-2311331021211312-2102000333231113)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers.item](resources--http_loadbalancer--reference--group-005.md#canonical-0222202333231133-3130233100122132-2001300113002031-1222000230031121-0033300020310102-1102201133102220-2001302332231031-1201030012220110)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1131212122002031-3301103113313301-3123113013200331-3210203100230333-3101332233223223-0002000322110020-2303130211303320-3113123122202203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233013032300103-1110003110333233-1313221201001022-0100323230203201-3100332133202112-2133213100030033-1312131312020123-1002221313131011"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_present — check_not_present / 322220112111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-3232223131110320-2312312311330110-3211012310212200-2133330302303310-2030221113011302-2300102101120101-2020022231332232-1313231311122212)
- api_protection_rules.api_endpoint_rules.request_matcher.headers.check_not_present

<a id="canonical-0310313222023133-3322122230030131-2330312120112022-3021131030302312-2301021321332032-0021203230310002-0013013311323000-0120230210310032"></a>

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

<a id="canonical-2133110221101232-0330133002023132-3123212303302111-3223011322130331-0323030010333022-1231220023032112-2032331130211132-3100022022232220"></a>

## Direct properties — check_not_present / 322220112111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212203331033203-1203312210322203-3233333230031323-1232232113332232-3011021132033212-0333033102010113-1201130330313012-3100220032212320"></a>

## Next pages — check_not_present / 322220112111 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-3232223131110320-2312312311330110-3211012310212200-2133330302303310-2030221113011302-2300102101120101-2020022231332232-1313231311122212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0303123022031230-0210331320202313-3011120333111210-2311131130302031-1322331131013123-0302110202112301-2311331021211312-2102000333231113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332011333213233-1301222220310032-1211331133230023-3112302010212313-3023332213302203-2033323111312121-1022022312111313-0333320213322210"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present — check_present / 103313100202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-3232223131110320-2312312311330110-3211012310212200-2133330302303310-2030221113011302-2300102101120101-2020022231332232-1313231311122212)
- api_protection_rules.api_endpoint_rules.request_matcher.headers.check_present

<a id="canonical-3222113130313113-3321021112013222-1100121100200103-2003022010002330-0320120310331310-3023101300121323-2313121011302012-1302223201310100"></a>

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

<a id="canonical-1023010001202011-2310013301332023-0212121022011211-2131322310130003-1132203200233021-0130100103131030-2032312210233022-3032101133003103"></a>

## Direct properties — check_present / 103313100202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231310120000023-1223103320320303-0120301131130001-2122110120120021-2313222001131023-1220231032120322-0131202130320200-2302122201200123"></a>

## Next pages — check_present / 103313100202 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-3232223131110320-2312312311330110-3211012310212200-2133330302303310-2030221113011302-2300102101120101-2020022231332232-1313231311122212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0222202333231133-3130233100122132-2001300113002031-1222000230031121-0033300020310102-1102201133102220-2001302332231031-1201030012220110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232132132331311-1023002212010032-0021221213222223-1322112220213302-3211320123022100-1113301331330312-3030123311332212-1102203021203011"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.headers.item — item / 222033302021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-3232223131110320-2312312311330110-3211012310212200-2133330302303310-2030221113011302-2300102101120101-2020022231332232-1313231311122212)
- api_protection_rules.api_endpoint_rules.request_matcher.headers.item

<a id="canonical-0313310030132200-0201130200111123-1311113023120103-1013313123322033-3123202132222202-2021012221100232-0230220111130102-0013023002020132"></a>

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

<a id="canonical-0033231200101102-3000113331321322-1211113302230223-1201003100122120-2233023320021213-2233122021000330-1132233213301120-3013002211101300"></a>

## Direct properties — item / 222033302021 / 3

<a id="canonical-0222223123231233-1012223010120003-3223210300231133-0211102031021213-0110223310010132-3302303212010312-1101000311311201-3110223320300133"></a>

<a id="canonical-0331023031302332-3130122013033222-2320021011011121-1121233132012030-1031133323310223-0210133031110000-1002200230113222-2320101231022223"></a>

## exact_values property — item / 222033302021 / 4

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

<a id="canonical-2030011133231112-2120331323310222-2201233302022112-2222322021010223-2122321010201233-2030013321102002-1331301102331302-0211303010003010"></a>

<a id="canonical-2302213111310001-0330012012213200-0000213122132203-1223021122212301-2212310233312212-1230202310032112-1001003002313131-1131112110210010"></a>

## regex_values property — item / 222033302021 / 5

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

<a id="canonical-1122001001230103-1010200323210300-0212130123322032-1303230313321022-2312303033202101-3032330133010033-0033202021022330-0113313002020303"></a>

<a id="canonical-2122120101211112-0130113210323230-1212103230320220-0313201210030001-1011233030022312-0010122231222001-1230031322233021-0203312313202112"></a>

## transformers property — item / 222033302021 / 6

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

<a id="canonical-2301310320321300-1311021001110020-1311010033311122-0023302000300022-1222001110013132-2222301330303311-3113310322123201-0132202121120023"></a>

## Next pages — item / 222033302021 / 7

- [api_protection_rules.api_endpoint_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-005.md#canonical-3232223131110320-2312312311330110-3211012310212200-2133330302303310-2030221113011302-2300102101120101-2020022231332232-1313231311122212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2222110200102020-3020200322033310-3003111203122333-1132303221211201-0330031100103113-3101011012122301-3302212223030312-1230213330301033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113110322302302-2000102202222022-1132023032031300-0000102100111122-0112112113103222-1210133210313131-0002012113212020-0133212011021302"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims — jwt_claims / 221312302231 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims

<a id="canonical-0021100202000323-3013102011221022-2231210200230102-0002113030002003-2302111033300120-3322032011213103-3300221111211102-2201313132333023"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

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
jwt_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231221120003111-3331231203222120-3002103033001323-3231230131131333-0021210211020011-0320232001233120-1310233031301221-0100123303010022"></a>

## Direct properties — jwt_claims / 221312302231 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-0003011201021003-0213010021322221-3120231301210002-3021312233122213-1300323113122200-3010112330131222-2332331112213133-2330023000213022): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-005.md#canonical-1013231211303123-1220030002202332-1310300330221320-0313113300201011-2221331103333213-1100201111213201-0312333232202311-0113333203010301): complete subsection reference.

<a id="canonical-0213131013102220-3023310302013302-3202300321321202-0332310103133033-3321110313013133-0110301223313023-0100122303113200-1202011211030323"></a>

<a id="canonical-1010020201011112-0310033230210120-1311310212011221-0022113010021213-1322231111031330-3022133332210303-2203312210102130-1202120102022331"></a>

## invert_matcher property — jwt_claims / 221312302231 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-005.md#canonical-0203322301301102-3102301233101300-1002313100310033-2121102301111000-0001321213001332-2221212100222123-2133100312013320-0300321121203300): complete subsection reference.

<a id="canonical-1202310110320132-0010033112012001-1303212111003101-1003223221131332-1133222300211303-3003123203300011-1123121230231303-3203123101330003"></a>

<a id="canonical-1123311223331320-3102100013010203-2112103232120030-3212302323331123-3033103212321113-0231211313331031-3333220302111003-2300122312320303"></a>

## name property — jwt_claims / 221312302231 / 5

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

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

<a id="canonical-0203221320112133-3301120320113111-0323023120232000-3032212130322313-0312331213331322-1311103011101201-1023033311221120-0302013033333321"></a>

## Next pages — jwt_claims / 221312302231 / 6

- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-0003011201021003-0213010021322221-3120231301210002-3021312233122213-1300323113122200-3010112330131222-2332331112213133-2330023000213022)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present](resources--http_loadbalancer--reference--group-005.md#canonical-1013231211303123-1220030002202332-1310300330221320-0313113300201011-2221331103333213-1100201111213201-0312333232202311-0113333203010301)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item](resources--http_loadbalancer--reference--group-005.md#canonical-0203322301301102-3102301233101300-1002313100310033-2121102301111000-0001321213001332-2221212100222123-2133100312013320-0300321121203300)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0003011201021003-0213010021322221-3120231301210002-3021312233122213-1300323113122200-3010112330131222-2332331112213133-2330023000213022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123022302021322-1003311021003023-0212121330012332-3201011220201312-1321030011302310-3120222032002100-0111300023323120-1312002202201120"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_present — check_not_present / 312322220120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-2222110200102020-3020200322033310-3003111203122333-1132303221211201-0330031100103113-3101011012122301-3302212223030312-1230213330301033)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-3212032301221200-0200100100230321-2002120032231333-3100122231200121-0020323111121200-3030222203220012-1032022111113311-1200103232123020"></a>

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

<a id="canonical-3121333333101322-0132101031200120-2203123130233230-2103112312313330-3002211232102113-0102030310301013-2113321233323211-1302333311213300"></a>

## Direct properties — check_not_present / 312322220120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112320121121310-2321113000132102-3033332222232331-1321302323133123-0230001011012232-2320133122310020-0003313103002223-1031210001022001"></a>

## Next pages — check_not_present / 312322220120 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-2222110200102020-3020200322033310-3003111203122333-1132303221211201-0330031100103113-3101011012122301-3302212223030312-1230213330301033)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1013231211303123-1220030002202332-1310300330221320-0313113300201011-2221331103333213-1100201111213201-0312333232202311-0113333203010301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301023013221320-1302110122331321-3021100001221100-0133221313100013-0333013123120111-0202001110122233-2213323323122231-2021031322121023"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present — check_present / 002100200310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-2222110200102020-3020200322033310-3003111203122333-1132303221211201-0330031100103113-3101011012122301-3302212223030312-1230213330301033)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.check_present

<a id="canonical-2302001030010312-0323221100001213-1331113222231200-3101020203122023-3010332013100113-2311331303132210-0102303103310312-2013321321002111"></a>

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

<a id="canonical-0123201100010132-3332002113220232-2230111313102111-2311013111121333-0231033021300033-0221323111231011-0001213320213313-0320000122222323"></a>

## Direct properties — check_present / 002100200310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221001220222113-0221020330310001-2101121032310313-3201031301030321-1201022132000103-0212013020213232-1331231003003131-0102000123100023"></a>

## Next pages — check_present / 002100200310 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-2222110200102020-3020200322033310-3003111203122333-1132303221211201-0330031100103113-3101011012122301-3302212223030312-1230213330301033)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0203322301301102-3102301233101300-1002313100310033-2121102301111000-0001321213001332-2221212100222123-2133100312013320-0300321121203300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310231202311133-1023132300302102-2003003332003133-2130302231020322-1310330122232123-2023123131320230-1022322331100333-2120221021132200"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item — item / 201013320011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-2222110200102020-3020200322033310-3003111203122333-1132303221211201-0330031100103113-3101011012122301-3302212223030312-1230213330301033)
- api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims.item

<a id="canonical-3112231301211032-0033121231112203-0120100313132221-2113300303103220-1112230303101202-1310300212110202-0000320000102200-1211320223220031"></a>

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

<a id="canonical-2003321211011230-1233322203312331-3032113303122301-1222201121230132-0012221230201320-3212333230032220-2302222302021110-0310011112212212"></a>

## Direct properties — item / 201013320011 / 3

<a id="canonical-3302311203010323-2223013023311300-3031333310310030-1031210122312312-0310112301320333-1220223120111323-3321021232120222-1202021320113001"></a>

<a id="canonical-0222110311021333-3223030312130231-3310021100012103-0312303330033033-0313312020021201-0200303312221202-0212002021213013-3101033321131203"></a>

## exact_values property — item / 201013320011 / 4

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

<a id="canonical-0311133330223130-3302000313202231-2210113113223213-3321001313323300-0013302211011333-1213110220232033-3221111312001010-1132310213233020"></a>

<a id="canonical-3032300213232101-0232100331221022-2221302021203032-0012003030121311-1331232123303013-2020103032202321-3331002213000012-1110211122220121"></a>

## regex_values property — item / 201013320011 / 5

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

<a id="canonical-2301311212300222-2310212312221233-1220032011200030-2203223132331010-3333213210132003-0232101131223010-0303033301010200-0212021322201211"></a>

<a id="canonical-1321313332323322-3220202030122222-0223203303121212-0102323300333211-0102030222233131-1103312222020103-3120103213031230-2112310111010203"></a>

## transformers property — item / 201013320011 / 6

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

<a id="canonical-2302032211130112-0011211003132022-1032100120031111-3232022000302101-1300312130013133-1311003001220102-0103101001220311-2232321213112221"></a>

## Next pages — item / 201013320011 / 7

- [api_protection_rules.api_endpoint_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-005.md#canonical-2222110200102020-3020200322033310-3003111203122333-1132303221211201-0330031100103113-3101011012122301-3302212223030312-1230213330301033)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3312230121022111-1002311110212233-0111012233200211-1310312023102312-1113111113123303-1302201110232311-0310332111123032-3221103322120332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012120032030323-0210213020201220-0123030322030320-3300302311321222-0303233002333020-3222201103322210-2303200022003130-2102201310103113"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.query_params — query_params / 220323330120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params

<a id="canonical-1023132001332122-0113230213001032-0002223023002321-0032132123330233-3022101201132020-3221310211130123-1000320033023210-0222012001132122"></a>

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

<a id="canonical-1103312121312122-0000112102233322-2313333330332132-3133131302213033-1220322112010320-0231021111222303-0012210133103121-3301313012320233"></a>

## Direct properties — query_params / 220323330120 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-3011220122222220-1301122320210331-2213302020311031-1211133011333120-2121203003121002-3022002102301322-0102101101011301-0310222231300312): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-005.md#canonical-2023121032333321-0233223332301213-3330322320230301-2110333020031332-3322122223202121-3133323020311310-3313131120030212-1331030112331012): complete subsection reference.

<a id="canonical-2021210020301220-0101221313121232-0202200311101031-2332120133100313-1121301021202113-0223220033022012-3112332100322013-2302000203211013"></a>

<a id="canonical-1123102120220330-3233003321333200-1012321101130012-0000033311320201-0333211013011121-0332012033332121-0301101000310110-2132310121312131"></a>

## invert_matcher property — query_params / 220323330120 / 4

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

- [item](resources--http_loadbalancer--reference--group-005.md#canonical-1302013023113101-1311012201231101-3200232200130033-0033312220111233-3112201120300322-0010121230021030-0002300313230232-2203312103220012): complete subsection reference.

<a id="canonical-1220012002213031-2021022320033023-3223203221132202-2312312311330312-0311230032331012-3301013132030331-2120111000120331-2223322301230123"></a>

<a id="canonical-2310313120310212-3202102210312111-2220211320330002-2223321332133322-0133020332303112-1223020301312322-0232003120212201-3120211221000321"></a>

## key property — query_params / 220323330120 / 5

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

<a id="canonical-3321021203332130-2232300203221010-2111111233022001-1133200311323201-0011300202212011-2211003033003202-3201332010301323-1320322132022032"></a>

## Next pages — query_params / 220323330120 / 6

- [api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_present](resources--http_loadbalancer--reference--group-005.md#canonical-3011220122222220-1301122320210331-2213302020311031-1211133011333120-2121203003121002-3022002102301322-0102101101011301-0310222231300312)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_present](resources--http_loadbalancer--reference--group-005.md#canonical-2023121032333321-0233223332301213-3330322320230301-2110333020031332-3322122223202121-3133323020311310-3313131120030212-1331030112331012)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params.item](resources--http_loadbalancer--reference--group-005.md#canonical-1302013023113101-1311012201231101-3200232200130033-0033312220111233-3112201120300322-0010121230021030-0002300313230232-2203312103220012)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3011220122222220-1301122320210331-2213302020311031-1211133011333120-2121203003121002-3022002102301322-0102101101011301-0310222231300312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312022023011212-2322323202333030-2123213031131223-2110121230303220-2313130200130021-3300211330103021-1122103332312333-1220032233020021"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_present — check_not_present / 203013230221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-005.md#canonical-3312230121022111-1002311110212233-0111012233200211-1310312023102312-1113111113123303-1302201110232311-0310332111123032-3221103322120332)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_not_present

<a id="canonical-1322203200231100-2322303120212022-1011113111332202-1320133203132101-2230333323231231-2203012123031231-3322020220223020-1131013310122123"></a>

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

<a id="canonical-2330223213301000-0101001133302130-3303332032202003-0232203323323232-1131013201300112-1223002331311310-0110213012122010-1200331112201120"></a>

## Direct properties — check_not_present / 203013230221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002003021200133-1221102320120312-2210232212031011-1211221320030332-1212011333033120-1100022303102321-3010133003223311-3301002030203101"></a>

## Next pages — check_not_present / 203013230221 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-005.md#canonical-3312230121022111-1002311110212233-0111012233200211-1310312023102312-1113111113123303-1302201110232311-0310332111123032-3221103322120332)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2023121032333321-0233223332301213-3330322320230301-2110333020031332-3322122223202121-3133323020311310-3313131120030212-1331030112331012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233002033212121-1133020111013121-1103202233122113-3110033311231003-3102032121222021-3020122232231111-0133300032030113-3313320203320312"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_present — check_present / 101321210011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-005.md#canonical-3312230121022111-1002311110212233-0111012233200211-1310312023102312-1113111113123303-1302201110232311-0310332111123032-3221103322120332)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params.check_present

<a id="canonical-3003233333332201-1303101031131321-3112320010222232-1212312230311223-2221110331013220-0222332311130313-3033313130212101-1131103133122301"></a>

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

<a id="canonical-1201321110311102-0132110100321213-3232203111322010-2020220002010213-1102030302101103-1212323230003032-3213013100133312-2230211220012231"></a>

## Direct properties — check_present / 101321210011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033311003322011-3133313022303102-1322110210300030-2231310130203301-3202020200131230-2323300022202321-3020330312023320-0220231320130323"></a>

## Next pages — check_present / 101321210011 / 4

- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-005.md#canonical-3312230121022111-1002311110212233-0111012233200211-1310312023102312-1113111113123303-1302201110232311-0310332111123032-3221103322120332)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1302013023113101-1311012201231101-3200232200130033-0033312220111233-3112201120300322-0010121230021030-0002300313230232-2203312103220012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301201120113032-1131200221123322-3303103321110212-0130300300032003-1130200321322021-3232230131321210-3022002000002202-1213022121332323"></a>

## api_protection_rules.api_endpoint_rules.request_matcher.query_params.item — item / 313230113211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_endpoint_rules](resources--http_loadbalancer--reference--group-004.md#canonical-1323301303113020-1123110023322233-1100330123101123-3123320111301001-0331011122221301-1030202010103130-1012211221032213-1031312012130313)
- [api_protection_rules.api_endpoint_rules.request_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0112033031323033-1330000120013122-3113223232032211-3031330100112003-3302222121213113-2333200033213011-1131223222320112-2011300332300012)
- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-005.md#canonical-3312230121022111-1002311110212233-0111012233200211-1310312023102312-1113111113123303-1302201110232311-0310332111123032-3221103322120332)
- api_protection_rules.api_endpoint_rules.request_matcher.query_params.item

<a id="canonical-0211222202323310-1030023102123130-0032003101030030-0222020201031103-2133111220333211-2023001223333001-2000102310202331-1301113033222200"></a>

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

<a id="canonical-2301221210020133-3111101002001020-1102023233112031-3222212030202023-3010233231221002-2323122021213213-1212130220321003-0123031031022111"></a>

## Direct properties — item / 313230113211 / 3

<a id="canonical-3211002111132202-2030321212002201-1131301103231230-3222110311203301-1012113210022200-0023332000323200-0033200012030020-3213310010220131"></a>

<a id="canonical-0000110032303102-1312033201301203-2223303113213033-0112123130123320-3331100230031002-3213232121233233-3001312030223111-3210321323113030"></a>

## exact_values property — item / 313230113211 / 4

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

<a id="canonical-1101313110111133-1303313203210231-2201233123010221-2112002332031232-3320103232333222-2212031232210220-2131211203121213-3211001133322023"></a>

<a id="canonical-2031331212212030-3022132033322122-2211313231322110-1230301310102202-1223112013130232-3223301032112231-0030230001031222-3122002330200320"></a>

## regex_values property — item / 313230113211 / 5

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

<a id="canonical-0221000333231023-2122101233111102-3312310133313201-1111102031100310-0001331001111200-0232112032123022-2301333320121012-0332032023133210"></a>

<a id="canonical-3321312112323200-3103213212311220-1003133223230233-0022011232033133-0131322221303132-0032201113023321-0322001023113132-1003010200132310"></a>

## transformers property — item / 313230113211 / 6

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

<a id="canonical-0313033323212201-0000310330332112-2021233120313212-3302011201333220-1333030123003022-2203002113102021-0011232032222323-0023123203003123"></a>

## Next pages — item / 313230113211 / 7

- [api_protection_rules.api_endpoint_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-005.md#canonical-3312230121022111-1002311110212233-0111012233200211-1310312023102312-1113111113123303-1302201110232311-0310332111123032-3221103322120332)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223020322113031-2033210323222222-3210113100222212-3110313221100121-0031212001330120-1031121103200301-2230232322111132-2303100122013320"></a>

## api_protection_rules.api_groups_rules — api_groups_rules / 011001223010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- api_protection_rules.api_groups_rules

<a id="canonical-0301120020103121-2203112221132003-3111131222031312-2103330211133333-2113012021302113-2233022000213032-3212313032122202-0113101022330113"></a>

Type: `"object"`. list nested block, Optional.

Category includes rules per API group or Server URL. For API groups, refer to API Definition which
includes API groups derived from uploaded swaggers.

Upstream description:

This category includes rules per API group or Server URL. For API groups, refer to API Definition
which includes API groups derived from uploaded swaggers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("base_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
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
    "ves.io.schema.rules.repeated.max_items": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

Terraform syntax:

```terraform
api_groups_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3331301033202332-2201223333301220-2000002333032113-2111032010101220-0030122130312033-2313001221202122-0101323130313313-3222023200323300"></a>

## Direct properties — api_groups_rules / 011001223010 / 3

- [action](resources--http_loadbalancer--reference--group-005.md#canonical-0131321221223300-1012113222332130-1102000121201012-2020023111111222-2311102022013011-1013121313200320-2331210002000030-0020022313111311): complete subsection reference.

- [any_domain](resources--http_loadbalancer--reference--group-005.md#canonical-3213023311121323-1113021222023213-1131132223030021-3332000100022232-1320331303331232-1113022133331012-0302020130233200-0320011212213120): complete subsection reference.

<a id="canonical-0011230112122322-1330113212102121-0313133321120331-3211102330221112-1210001220013230-1120002320302103-1031003022002203-0030120033300112"></a>

<a id="canonical-1301220303331332-3201212233303233-2310132020202222-2002213313323011-2321003211333310-3100333333300330-1332032323032202-2302010101032030"></a>

## api_group property — api_groups_rules / 011001223010 / 4

Type: `"string"`. Optional.

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with 'x-F5 Distributed..

Upstream description:

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with "x-F5 Distributed
Cloud-API-group" extensions inside swaggers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1213023212133200-3022323200001200-3323001303100130-0211232000303032-1300101233020312-3200233121313013-2111320303022211-2111000001230110"></a>

<a id="canonical-3312002021321320-1301032212230013-0033002003001000-1101232031100003-2313030011100033-1032111320033321-3233013201221230-0021202031000013"></a>

## base_path property — api_groups_rules / 011001223010 / 5

Type: `"string"`. Optional.

Base Path. Prefix of the request path. For example: /v1.

Upstream description:

Prefix of the request path. For example: /v1.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-006.md#canonical-3122101313113213-0230122323313020-2213032333110001-0122003110111023-2020313332102030-1103032103222301-1211233231220112-0322000110200233): complete subsection reference.

- [request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012): complete subsection reference.

<a id="canonical-0100103332212103-0311112022302222-1323101121010112-1102230310310032-2303301121230013-3313302301220133-0122213311000011-0322010220213011"></a>

<a id="canonical-0311232211333301-3300202313100011-3211222011321301-0122012313030300-3133010320311101-2301232112101020-3123033222230221-0201233321212200"></a>

## specific_domain property — api_groups_rules / 011001223010 / 6

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For example:
api.example.com.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-3022221003103230-3102232123310203-2321222131201030-3030030303330300-3232032113301003-0332301312333321-0121223021031323-1121113221300301"></a>

## Next pages — api_groups_rules / 011001223010 / 7

- [api_protection_rules.api_groups_rules.action](resources--http_loadbalancer--reference--group-005.md#canonical-0131321221223300-1012113222332130-1102000121201012-2020023111111222-2311102022013011-1013121313200320-2331210002000030-0020022313111311)
- [api_protection_rules.api_groups_rules.any_domain](resources--http_loadbalancer--reference--group-005.md#canonical-3213023311121323-1113021222023213-1131132223030021-3332000100022232-1320331303331232-1113022133331012-0302020130233200-0320011212213120)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- [api_protection_rules.api_groups_rules.metadata](resources--http_loadbalancer--reference--group-006.md#canonical-3122101313113213-0230122323313020-2213032333110001-0122003110111023-2020313332102030-1103032103222301-1211233231220112-0322000110200233)
- [api_protection_rules.api_groups_rules.request_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0220022131323212-2000011221210102-1033031110331021-3101230131122312-2122233120233131-1312300011312030-3233210211023020-1012122211111012)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0131321221223300-1012113222332130-1102000121201012-2020023111111222-2311102022013011-1013121313200320-2331210002000030-0020022313111311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310221332122111-0132331210022022-0220323322312133-0132121212102220-1122031000030023-2011303101123223-2100031001110211-2121211322212033"></a>

## api_protection_rules.api_groups_rules.action — action / 010101332311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- api_protection_rules.api_groups_rules.action

<a id="canonical-0103330233213001-3113331100312310-2211003030121312-0312032010131011-3312023331132233-0020011002220121-0303032100133333-3310000030233233"></a>

Type: `"object"`. single nested block, Optional.

The action to take if the input request matches the rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow",
    "deny")}
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
  "x-ves-oneof-field-action": "[\"allow\",\"deny\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313031323322221-3200300123201231-1231222132030321-0120212232132312-0300321000101202-3103102232121112-2323012032022100-0212021101310230"></a>

## Direct properties — action / 010101332311 / 3

- [allow](resources--http_loadbalancer--reference--group-005.md#canonical-2311201103303202-3333101313100311-0010213131331302-0101102133122112-1122303302033132-1102103121303302-0133232333303222-2020001130301001): complete subsection reference.

- [deny](resources--http_loadbalancer--reference--group-005.md#canonical-1011103131332301-0003222000202323-3000303101231300-1321103311102030-1212231223203111-3130301030303131-0320132010233301-3312103032031012): complete subsection reference.

<a id="canonical-3232203320302130-0230211030112121-3101232201110031-3312323000003020-3313203213032100-0133311311123002-1033200112022012-1103002011211211"></a>

## Next pages — action / 010101332311 / 4

- [api_protection_rules.api_groups_rules.action.allow](resources--http_loadbalancer--reference--group-005.md#canonical-2311201103303202-3333101313100311-0010213131331302-0101102133122112-1122303302033132-1102103121303302-0133232333303222-2020001130301001)
- [api_protection_rules.api_groups_rules.action.deny](resources--http_loadbalancer--reference--group-005.md#canonical-1011103131332301-0003222000202323-3000303101231300-1321103311102030-1212231223203111-3130301030303131-0320132010233301-3312103032031012)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2311201103303202-3333101313100311-0010213131331302-0101102133122112-1122303302033132-1102103121303302-0133232333303222-2020001130301001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122220213321021-0303303010320313-3300232200110321-1203100221130011-1221133021222033-0113121210213232-0313010130303012-0013231113223110"></a>

## api_protection_rules.api_groups_rules.action.allow — allow / 202210232120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.action](resources--http_loadbalancer--reference--group-005.md#canonical-0131321221223300-1012113222332130-1102000121201012-2020023111111222-2311102022013011-1013121313200320-2331210002000030-0020022313111311)
- api_protection_rules.api_groups_rules.action.allow

<a id="canonical-2011133102202232-2123011010200010-1221110323001211-2223031320332333-3001122303213313-0020011013003213-3220110210001311-2222201113330301"></a>

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
allow = {}
```

<a id="canonical-0012030201010200-0201012301010121-3011201300110113-0211301210300130-0032011111112022-1331010132103311-3120312103000020-3021331031322213"></a>

## Direct properties — allow / 202210232120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123303001220303-3131122133103000-1021112122200333-0022100200330003-0222220200122031-2213231222313120-2100130331213221-3202131110020013"></a>

## Next pages — allow / 202210232120 / 4

- [api_protection_rules.api_groups_rules.action](resources--http_loadbalancer--reference--group-005.md#canonical-0131321221223300-1012113222332130-1102000121201012-2020023111111222-2311102022013011-1013121313200320-2331210002000030-0020022313111311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1011103131332301-0003222000202323-3000303101231300-1321103311102030-1212231223203111-3130301030303131-0320132010233301-3312103032031012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222102020011112-1202220013102331-1133021220332100-0201301111230302-2021133311032010-1330313011212013-3222101332121020-1330203330220333"></a>

## api_protection_rules.api_groups_rules.action.deny — deny / 031302030222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.action](resources--http_loadbalancer--reference--group-005.md#canonical-0131321221223300-1012113222332130-1102000121201012-2020023111111222-2311102022013011-1013121313200320-2331210002000030-0020022313111311)
- api_protection_rules.api_groups_rules.action.deny

<a id="canonical-2112220020020000-1322322031302330-0320101312301300-3112133132120321-1012100123233202-2211122202021212-3312311300202100-2202003033031101"></a>

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
deny = {}
```

<a id="canonical-3031201022300133-3303303103211012-2131023300321011-1010320000210102-3120001120320212-2101032101012021-1202010120101312-3313123220102033"></a>

## Direct properties — deny / 031302030222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322232110231213-2201130111331210-2132213021003001-3211032010212231-3013221010333131-1023013022302021-2123111121102320-3300232332112303"></a>

## Next pages — deny / 031302030222 / 4

- [api_protection_rules.api_groups_rules.action](resources--http_loadbalancer--reference--group-005.md#canonical-0131321221223300-1012113222332130-1102000121201012-2020023111111222-2311102022013011-1013121313200320-2331210002000030-0020022313111311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3213023311121323-1113021222023213-1131132223030021-3332000100022232-1320331303331232-1113022133331012-0302020130233200-0320011212213120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320321302032133-2032333220122021-2211111000220202-1321230030312031-3011331313321123-0221001032132232-1033003132031113-1013022130100213"></a>

## api_protection_rules.api_groups_rules.any_domain — any_domain / 023302303213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- api_protection_rules.api_groups_rules.any_domain

<a id="canonical-2020023203023022-0310111202201011-0113311230231212-3012303312001121-0303130300030221-2311221300222031-2310120010231323-3001300012332201"></a>

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
any_domain = {}
```

<a id="canonical-0331133132211002-1210102000203223-0012110021303012-0332131013220021-1231200111010321-3113023133311032-1330232322031000-3220330230031001"></a>

## Direct properties — any_domain / 023302303213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312320233102032-3103303101010223-0301002110200122-0200210121300023-0222133203023200-3211313302231322-3121132303003231-2302213131013011"></a>

## Next pages — any_domain / 023302303213 / 4

- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332131301232311-3210100332010211-0330013101122321-1020212231010221-1310130012000320-3113020031011133-1021112332123030-2002203020302020"></a>

## api_protection_rules.api_groups_rules.client_matcher — client_matcher / 302101102012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- api_protection_rules.api_groups_rules.client_matcher

<a id="canonical-2211032301110320-3023013122132232-1121121302211111-3211012100030200-3110320302123030-3002012103201032-2323023231111323-0102211131001211"></a>

Type: `"object"`. single nested block, Optional.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
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
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

Terraform syntax:

```terraform
client_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303203322003121-1302001323020112-0232132001122223-2211031201212110-0223031332212030-2010110330322203-1003201230231010-1332021123031313"></a>

## Direct properties — client_matcher / 302101102012 / 3

- [any_client](resources--http_loadbalancer--reference--group-005.md#canonical-1300013322303103-1331110120302230-0102233212323100-2132323233020331-3002002303113302-1301031100120032-0212000131331212-0220333203032233): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-005.md#canonical-0320011203103032-3303133012331003-1021132330033103-0031233301102301-2322210323303312-2201100111231023-2121210332222222-2112120223122101): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-005.md#canonical-2321121232320223-0232023121030202-0011333130323002-1001123203010011-2102210000232100-1223013002232000-2010332002103322-1112322322201133): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2311000320030232-2331211323320232-3131030121111223-0303223300002310-0212130111331223-1111231131100033-1221200123011211-1312103303113313): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-005.md#canonical-1330201310021011-1330101002323003-2232000112312001-2232012231022311-0101130333300112-0313203200300011-2212333322030032-3122131111200321): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2300303030212121-2323123220320310-3202111012102301-0022011022002333-3222200330223230-2213031232033201-0120232230311322-3102231110323010): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-005.md#canonical-2003011313330030-1202302233223210-0012211101230003-0211300333033030-3320133032010330-0130020033330030-0233023230233212-2003221002130210): complete subsection reference.

- [ip_threat_category_list](resources--http_loadbalancer--reference--group-005.md#canonical-0123200231003121-2230301310320322-1132110230301000-3020301013222211-1010312322020020-1233321132113202-3231120013222233-2312310101110302): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0020130311132331-1120230211123201-2203233322212130-2002331231003300-0300110323300011-3231313303001130-2202322012200333-3212013201120332): complete subsection reference.

<a id="canonical-1102013110330201-0212232021023311-3122233231300003-2101031222002333-1012201230200301-3232321110312000-1033222110003200-1222211232021322"></a>

## Next pages — client_matcher / 302101102012 / 4

- [api_protection_rules.api_groups_rules.client_matcher.any_client](resources--http_loadbalancer--reference--group-005.md#canonical-1300013322303103-1331110120302230-0102233212323100-2132323233020331-3002002303113302-1301031100120032-0212000131331212-0220333203032233)
- [api_protection_rules.api_groups_rules.client_matcher.any_ip](resources--http_loadbalancer--reference--group-005.md#canonical-0320011203103032-3303133012331003-1021132330033103-0031233301102301-2322210323303312-2201100111231023-2121210332222222-2112120223122101)
- [api_protection_rules.api_groups_rules.client_matcher.asn_list](resources--http_loadbalancer--reference--group-005.md#canonical-2321121232320223-0232023121030202-0011333130323002-1001123203010011-2102210000232100-1223013002232000-2010332002103322-1112322322201133)
- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2311000320030232-2331211323320232-3131030121111223-0303223300002310-0212130111331223-1111231131100033-1221200123011211-1312103303113313)
- [api_protection_rules.api_groups_rules.client_matcher.client_selector](resources--http_loadbalancer--reference--group-005.md#canonical-1330201310021011-1330101002323003-2232000112312001-2232012231022311-0101130333300112-0313203200300011-2212333322030032-3122131111200321)
- [api_protection_rules.api_groups_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2300303030212121-2323123220320310-3202111012102301-0022011022002333-3222200330223230-2213031232033201-0120232230311322-3102231110323010)
- [api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list](resources--http_loadbalancer--reference--group-005.md#canonical-2003011313330030-1202302233223210-0012211101230003-0211300333033030-3320133032010330-0130020033330030-0233023230233212-2003221002130210)
- [api_protection_rules.api_groups_rules.client_matcher.ip_threat_category_list](resources--http_loadbalancer--reference--group-005.md#canonical-0123200231003121-2230301310320322-1132110230301000-3020301013222211-1010312322020020-1233321132113202-3231120013222233-2312310101110302)
- [api_protection_rules.api_groups_rules.client_matcher.tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-006.md#canonical-0020130311132331-1120230211123201-2203233322212130-2002331231003300-0300110323300011-3231313303001130-2202322012200333-3212013201120332)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1300013322303103-1331110120302230-0102233212323100-2132323233020331-3002002303113302-1301031100120032-0212000131331212-0220333203032233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101203021301003-2200010320010010-3222100023112021-2323002320102310-3333300321100001-0023201330321120-0320302122110320-1012032332332023"></a>

## api_protection_rules.api_groups_rules.client_matcher.any_client — any_client / 333322113000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.any_client

<a id="canonical-0013223200300220-1000031131103100-3012130322321223-1121103022030213-0122222013221022-0213300033021111-1212330211001223-2122222230133113"></a>

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
any_client = {}
```

<a id="canonical-0131131301220303-3313120003213311-1220212110331233-3100102322210233-1300332331301032-1022310133233022-2102210021022002-0213022231123213"></a>

## Direct properties — any_client / 333322113000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223313230332013-0023020300230221-0120112103312030-3123201101332012-0213311303121223-1222111322012022-1000213200112130-3311322211122121"></a>

## Next pages — any_client / 333322113000 / 4

- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0320011203103032-3303133012331003-1021132330033103-0031233301102301-2322210323303312-2201100111231023-2121210332222222-2112120223122101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222002233121212-2012020033221131-1012010220031013-2131031010213021-0213213021020331-3203032113021131-3323030022030102-1213021311320011"></a>

## api_protection_rules.api_groups_rules.client_matcher.any_ip — any_ip / 200012310331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.any_ip

<a id="canonical-3300010111221233-3221020022321032-1210222002120022-0221020203322032-2311322001320122-0301011130311303-0322103030221033-1211201222200210"></a>

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
any_ip = {}
```

<a id="canonical-2022022111210111-2132222313012102-2032110311121110-2321311300111123-1032222201131130-0122012131100002-0023220120122030-3311321332103111"></a>

## Direct properties — any_ip / 200012310331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100313210031023-3021210221021303-1210201110203112-0220200213001222-0203132020003221-1011233211232310-3220112222303002-3122330120031101"></a>

## Next pages — any_ip / 200012310331 / 4

- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2321121232320223-0232023121030202-0011333130323002-1001123203010011-2102210000232100-1223013002232000-2010332002103322-1112322322201133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010122203110001-2111321202203330-2133200012221000-0031121111101120-2123031220022031-1301332011123330-0321321021202230-3123321230300003"></a>

## api_protection_rules.api_groups_rules.client_matcher.asn_list — asn_list / 201210020212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.asn_list

<a id="canonical-1112011222122203-2102223323033100-0232223132103302-2323101310011231-1212000102110102-2311103213202012-3202030012110031-2330323222300112"></a>

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

<a id="canonical-1311023021120200-0133120211010130-1322110003202132-0030122120330223-2131301103321313-2310111320131020-0320323101213113-2033312022110031"></a>

## Direct properties — asn_list / 201210020212 / 3

<a id="canonical-2122103231223212-1322100213103030-3222230012110202-0112120110221213-0223121331130222-1233203322121203-1230201321331112-0011311300103332"></a>

<a id="canonical-3121321021113020-3120000111133110-0312102312230232-1020302130031320-0000221113033011-3110011313133200-2030200001210301-2000332332002100"></a>

## as_numbers property — asn_list / 201210020212 / 4

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

<a id="canonical-3002012133002010-1220321013103212-0021100033123223-3331023213130000-0113211121122302-3032201101100030-3020311020322021-2330331331001313"></a>

## Next pages — asn_list / 201210020212 / 5

- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2311000320030232-2331211323320232-3131030121111223-0303223300002310-0212130111331223-1111231131100033-1221200123011211-1312103303113313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222103132202300-1102200123301132-0013301320100103-2020312203213331-0213131032002303-3230231222302123-3110333020221333-1202203133122221"></a>

## api_protection_rules.api_groups_rules.client_matcher.asn_matcher — asn_matcher / 101111323331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.asn_matcher

<a id="canonical-0020112121133213-2332121121202111-3303223212122130-1321012303120011-1121301010120302-0320011222211032-0202121031110220-0102032120131313"></a>

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

<a id="canonical-2120202123313102-0311311331010102-3121032300110111-1213320130123111-0022130320032220-2020311210120113-3331333110333200-0300032303221333"></a>

## Direct properties — asn_matcher / 101111323331 / 3

- [asn_sets](resources--http_loadbalancer--reference--group-005.md#canonical-1222201100210311-0021323021012223-1003303023211321-3012200213021021-3220210211223312-2233003310203331-3003010313001312-1332222120003030): complete subsection reference.

<a id="canonical-1001012330020131-1312133033230322-3122202330131321-0120101312330110-1301320312003310-2111133021222021-1120301211111322-0231023021222033"></a>

## Next pages — asn_matcher / 101111323331 / 4

- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets](resources--http_loadbalancer--reference--group-005.md#canonical-1222201100210311-0021323021012223-1003303023211321-3012200213021021-3220210211223312-2233003310203331-3003010313001312-1332222120003030)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1222201100210311-0021323021012223-1003303023211321-3012200213021021-3220210211223312-2233003310203331-3003010313001312-1332222120003030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231132321010201-3321121231120300-1212133222201013-0132123223113202-3031330030103222-0232231302103033-1330211121321222-3030200321010103"></a>

## api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets — asn_sets / 312121023130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2311000320030232-2331211323320232-3131030121111223-0303223300002310-0212130111331223-1111231131100033-1221200123011211-1312103303113313)
- api_protection_rules.api_groups_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-1210321202300300-1011212333331022-3303332223201223-0022001332223300-0131103223112201-0210232310111212-1021233103312002-1121121033031010"></a>

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

<a id="canonical-2003112202000122-0010233033122332-2302213212130232-0232223113231201-1310320123111110-0231013032333113-1322212210231022-2000022312111211"></a>

## Direct properties — asn_sets / 312121023130 / 3

<a id="canonical-3300202012102323-2121213001022000-0112130221213210-0000200032102203-2322132300231133-3000201101222003-0212211100021232-2023013312123032"></a>

<a id="canonical-0320113002001102-3102200131221322-1231222210033003-0111203200301321-1102122232223202-2231321210003113-0102000332032312-1332322230023233"></a>

## kind property — asn_sets / 312121023130 / 4

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

<a id="canonical-2333323023003102-2133200120332232-0102112303023121-0132321321123231-3233210321003113-3230333313003231-1032023121013110-0033201321102022"></a>

<a id="canonical-0103012312203100-1030211221120303-2013012020323031-0013203320302102-3312223132330131-3120303013230000-3002320332320131-2302211322311230"></a>

## name property — asn_sets / 312121023130 / 5

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

<a id="canonical-3310300313010003-2033333003102021-1211122133012210-0201033121100333-1311100212301031-2212310132120300-0222131012332331-3103211201030330"></a>

<a id="canonical-3131202200230003-0311102123220020-0021213233211200-1332313221312013-1231313033302302-3111113200231111-1333320013031202-2132222312120212"></a>

## namespace property — asn_sets / 312121023130 / 6

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

<a id="canonical-1301021213211310-2013331312320003-2123321321312330-3033002233120010-3332030213302133-2332122121010312-2131111003100022-1221111120021202"></a>

<a id="canonical-0312002320320121-3003003013330231-1303001111111013-0130001101010330-0003312233102012-2131222322330203-3230132311212032-2303133031122021"></a>

## tenant property — asn_sets / 312121023130 / 7

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

<a id="canonical-0323301301002000-2131212231332210-3313010232232101-3030200312101103-0002303330032230-0131020111002001-1002322002212100-0331122331112103"></a>

<a id="canonical-2332100122013220-3122200302112332-1030212300120123-2133021112313230-0121202032131320-1101130311210313-2133211213000113-0332132032311312"></a>

## uid property — asn_sets / 312121023130 / 8

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

<a id="canonical-0021121120020320-0202212000233030-3031323222213011-2223122212020002-2020011023312331-0222013333103302-1330121310023303-3120021231313210"></a>

## Next pages — asn_sets / 312121023130 / 9

- [api_protection_rules.api_groups_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2311000320030232-2331211323320232-3131030121111223-0303223300002310-0212130111331223-1111231131100033-1221200123011211-1312103303113313)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1330201310021011-1330101002323003-2232000112312001-2232012231022311-0101130333300112-0313203200300011-2212333322030032-3122131111200321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301223000300111-3210131000323201-2111002322113221-3303312012330201-1001303331231130-3011103331230122-2211033311211121-0131133122103031"></a>

## api_protection_rules.api_groups_rules.client_matcher.client_selector — client_selector / 101103331021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.client_selector

<a id="canonical-2001022122331123-2322333311301203-3200231131103332-3103303133033210-1100032232033211-3010102330313310-3200022332111210-1132102102010122"></a>

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

<a id="canonical-1310223230300320-1011300112303022-2121111121323220-3031202210223233-0113120010122302-2023132321123223-3202202032113123-3022210212101321"></a>

## Direct properties — client_selector / 101103331021 / 3

<a id="canonical-2221202220103133-2120220103230301-3122311222333011-3021302032302032-2112113103122121-0231212012100213-0310232022131001-2013132131212103"></a>

<a id="canonical-1000022123301111-1000320320111002-0332011212300200-2221302302330203-1322103021333113-2321030003011100-2211102201233032-1301231310131131"></a>

## expressions property — client_selector / 101103331021 / 4

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

<a id="canonical-2333223113303032-1000230333200313-1002023001310103-3232203220012302-2002322120021120-0010201013202223-0223300220112330-0103232213000101"></a>

## Next pages — client_selector / 101103331021 / 5

- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2300303030212121-2323123220320310-3202111012102301-0022011022002333-3222200330223230-2213031232033201-0120232230311322-3102231110323010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221323101223300-0230031131230010-0023130101300312-3130302312203001-2212230002321120-0022023100120202-3012023212130233-2313213301102111"></a>

## api_protection_rules.api_groups_rules.client_matcher.ip_matcher — ip_matcher / 300213322322 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.ip_matcher

<a id="canonical-3213320202222233-0021102220021023-1322031121030033-2120222321221131-3332122311201221-2212311223022203-2100322003111110-3132231010310311"></a>

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

<a id="canonical-2202020313113223-0223300301323120-3023122122322313-3001201010020231-3122302331020011-0322132023302320-0213211022200000-3200331031302330"></a>

## Direct properties — ip_matcher / 300213322322 / 3

<a id="canonical-2001211002121330-1231232303023032-1100220220131200-1120302230200313-1013230111202100-0311230122201120-1121330103031321-0321232202002023"></a>

<a id="canonical-0100311122130300-3322111013133331-1111223303120002-2100002021201230-1210210302130013-2121002301232332-0033320301221022-2113320223200010"></a>

## invert_matcher property — ip_matcher / 300213322322 / 4

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

- [prefix_sets](resources--http_loadbalancer--reference--group-005.md#canonical-1323201312113011-1233011112101002-0321030213030330-0003000122121332-2321113010213302-2320200121220211-2313322310332300-0113303230111031): complete subsection reference.

<a id="canonical-1212323131012002-1120210100010011-1211313332002303-1100303031031223-0112331330230121-3233310113021002-0211313023031103-0031203210012220"></a>

## Next pages — ip_matcher / 300213322322 / 5

- [api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets](resources--http_loadbalancer--reference--group-005.md#canonical-1323201312113011-1233011112101002-0321030213030330-0003000122121332-2321113010213302-2320200121220211-2313322310332300-0113303230111031)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1323201312113011-1233011112101002-0321030213030330-0003000122121332-2321113010213302-2320200121220211-2313322310332300-0113303230111031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120123002323131-2002131220020003-0211022113213011-0010021201020333-0010210320301102-0203000202002112-1021331223330211-1222102132132023"></a>

## api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets — prefix_sets / 010323322302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- [api_protection_rules.api_groups_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2300303030212121-2323123220320310-3202111012102301-0022011022002333-3222200330223230-2213031232033201-0120232230311322-3102231110323010)
- api_protection_rules.api_groups_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-2123033322312121-0211001223301300-2310113312020003-2011311031312202-3223032332010033-2033001333330022-1201232320230303-2312120001133013"></a>

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

<a id="canonical-3103302013113120-1111013020320022-1133113203230323-2111200323131311-1201111032333033-0012201203320330-1021002132112313-1220131302013031"></a>

## Direct properties — prefix_sets / 010323322302 / 3

<a id="canonical-1301212202222011-0021022333122002-1320310300003131-1330120213220020-1323233010213213-3002000113330011-0310021213322331-0131033221122000"></a>

<a id="canonical-2023021321330213-0013120200230332-1311113133101003-2202211300120131-2131212132000010-3023123233132003-3321031222333232-2002230311120313"></a>

## kind property — prefix_sets / 010323322302 / 4

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

<a id="canonical-1333122023013112-0233111220113113-0101003231022020-2130223002320313-1011010120303133-1230310312033301-2133322330231112-2013232222032032"></a>

<a id="canonical-1002210233122023-3200101323111310-0322330013002322-1023102320332022-0212031111221001-2300020121232112-1122103303231330-3222213011221121"></a>

## name property — prefix_sets / 010323322302 / 5

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

<a id="canonical-3111112232300310-3203100113013212-0100201321232223-0123023100032300-2210033032131323-1020233130120230-0211112103221201-0033101003212031"></a>

<a id="canonical-1113233030230323-1201330121323222-3302320302133332-3202101013113230-1302113333013333-2013212123031111-1332320122212113-0201011300020012"></a>

## namespace property — prefix_sets / 010323322302 / 6

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

<a id="canonical-3303310321203201-2120202101332032-0113301333233132-1101330123103110-2301222133121321-0102301313022132-0221002111030232-1033032131121022"></a>

<a id="canonical-3303233111031311-2330220131223112-0023021321103110-0032210023321301-2033301023120022-3111002212023031-3030121211102310-3101232222023312"></a>

## tenant property — prefix_sets / 010323322302 / 7

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

<a id="canonical-3003003030220112-0322101320212323-0213001002003120-0030013001032230-2012231202321021-1001202212100123-0002213220232001-3310232021303112"></a>

<a id="canonical-1020300023013230-3132220133111002-1021220002031113-2222221200331121-1220031211321323-1010312211111120-3233113330133130-2000001131322130"></a>

## uid property — prefix_sets / 010323322302 / 8

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

<a id="canonical-1230030101221131-0303201230230302-2302022332132331-1231212332102331-2213310123113112-3131120112123002-0121331210023222-1210210030113213"></a>

## Next pages — prefix_sets / 010323322302 / 9

- [api_protection_rules.api_groups_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-2300303030212121-2323123220320310-3202111012102301-0022011022002333-3222200330223230-2213031232033201-0120232230311322-3102231110323010)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2003011313330030-1202302233223210-0012211101230003-0211300333033030-3320133032010330-0130020033330030-0233023230233212-2003221002130210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210121001210010-1210222011000221-2023132302133303-3331021322201202-2323021030012022-0120233313211212-0001022000112320-0101000202022002"></a>

## api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list — ip_prefix_list / 210310131102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_protection_rules](resources--http_loadbalancer--reference--group-004.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130)
- [api_protection_rules.api_groups_rules](resources--http_loadbalancer--reference--group-005.md#canonical-1302001032222010-1223222102002012-2000003200023112-1322101202200123-3201202122113002-3323202020213011-2230001222123103-3230112203033011)
- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- api_protection_rules.api_groups_rules.client_matcher.ip_prefix_list

<a id="canonical-2021030233002330-2320231023221001-3120303321222301-2123010033331011-2202032222111121-2310002132212231-1022013301133013-0323132212322012"></a>

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

<a id="canonical-0233031320021133-3130033112202013-0300221131032130-2210321132121331-0101123333321220-0220230322203230-3210113223012300-3303033101010131"></a>

## Direct properties — ip_prefix_list / 210310131102 / 3

<a id="canonical-3130232031223300-1010203301320322-3000300231210113-0210330303332210-3022111211033211-0110302112021123-1112113002133101-0122323330332233"></a>

<a id="canonical-1223001202200131-0032121021021112-3330110221010212-1032002022033310-2131102200011313-2200312123133120-2133123322131321-2300221220112311"></a>

## invert_match property — ip_prefix_list / 210310131102 / 4

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

<a id="canonical-3320111021313210-2001002210211002-3323122122321321-2033020333032310-3030232300332032-3101012120132002-3220123100320012-3011212130321223"></a>

<a id="canonical-0133231013211301-2210003120300220-2210003301131023-1223310233233211-1300223233331211-0000022111022002-0303133001131212-2322320210230312"></a>

## ip_prefixes property — ip_prefix_list / 210310131102 / 5

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

<a id="canonical-1230101303330330-3110120221302021-3211301333112113-2021011112213102-0131122323032001-2022310101112303-0032311112203030-3311213302003021"></a>

## Next pages — ip_prefix_list / 210310131102 / 6

- [api_protection_rules.api_groups_rules.client_matcher](resources--http_loadbalancer--reference--group-005.md#canonical-0213211110003101-2312321032300110-2120003201113132-3130200231322230-3133302000013102-2213311222221133-2303103312312013-1103003103201021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0123200231003121-2230301310320322-1132110230301000-3020301013222211-1010312322020020-1233321132113202-3231120013222233-2312310101110302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
