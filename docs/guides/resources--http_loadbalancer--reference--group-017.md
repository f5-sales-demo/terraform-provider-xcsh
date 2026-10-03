---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3030133112010223-0303113001203112-0212310322112010-2120030300200000-3120130312202330-0233121223012320-1200103300003300-3312200313320230"></a>

## dns_name property — private_name / 330303000231 / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [inside_network](resources--http_loadbalancer--reference--group-017.md#canonical-1012330102231300-3300211120313123-1033000300003311-2033220022003310-3110013230300333-3111000023130032-1032303200312332-0223310111113132): complete subsection reference.

- [outside_network](resources--http_loadbalancer--reference--group-017.md#canonical-0311301122230130-3202322330221100-3010223302023213-3000100231023333-1000112032103002-1132120213013212-3131232033220233-0012222212112021): complete subsection reference.

<a id="canonical-1221000301200322-3330011231323302-0231220220131132-2213333022101012-1301322203313333-2333200033123022-1032031303110323-3101231123320210"></a>

<a id="canonical-3220013130232200-1320200203100003-3321111103230233-1222030220200100-0231131220310012-3123221103013101-3030122322313202-3033320313021200"></a>

## refresh_interval property — private_name / 330303000231 / 5

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

- [segment](resources--http_loadbalancer--reference--group-017.md#canonical-3322321310112123-3331203210113301-2130130132133122-2032103130113112-2102312033030203-0213230232130021-0013232233102032-3321020322030233): complete subsection reference.

- [site_locator](resources--http_loadbalancer--reference--group-017.md#canonical-0211232012223111-1001231003113023-1212212001033231-0031131311031121-2101212231112023-0332221013322020-3210130213310212-0001122020323330): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-017.md#canonical-0312130022121113-1033003212331120-0023101102200212-1312013113201020-0203311020213123-0223031101003330-3232310110011201-2312012230001330): complete subsection reference.

<a id="canonical-1212212220330123-0110110302330100-1013102313122032-1211013002013130-1123022133202301-2301030013211013-0103002110301002-3331302112133220"></a>

## Next pages — private_name / 330303000231 / 6

- [default_pool.origin_servers.private_name.inside_network](resources--http_loadbalancer--reference--group-017.md#canonical-1012330102231300-3300211120313123-1033000300003311-2033220022003310-3110013230300333-3111000023130032-1032303200312332-0223310111113132)
- [default_pool.origin_servers.private_name.outside_network](resources--http_loadbalancer--reference--group-017.md#canonical-0311301122230130-3202322330221100-3010223302023213-3000100231023333-1000112032103002-1132120213013212-3131232033220233-0012222212112021)
- [default_pool.origin_servers.private_name.segment](resources--http_loadbalancer--reference--group-017.md#canonical-3322321310112123-3331203210113301-2130130132133122-2032103130113112-2102312033030203-0213230232130021-0013232233102032-3321020322030233)
- [default_pool.origin_servers.private_name.site_locator](resources--http_loadbalancer--reference--group-017.md#canonical-0211232012223111-1001231003113023-1212212001033231-0031131311031121-2101212231112023-0332221013322020-3210130213310212-0001122020323330)
- [default_pool.origin_servers.private_name.snat_pool](resources--http_loadbalancer--reference--group-017.md#canonical-0312130022121113-1033003212331120-0023101102200212-1312013113201020-0203311020213123-0223031101003330-3232310110011201-2312012230001330)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1012330102231300-3300211120313123-1033000300003311-2033220022003310-3110013230300333-3111000023130032-1032303200312332-0223310111113132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231221013022011-0130010311102231-2100330311010021-2333013113200323-0311023122211100-0121312101002123-3122201221213312-2111213100000100"></a>

## default_pool.origin_servers.private_name.inside_network — inside_network / 220300230000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- default_pool.origin_servers.private_name.inside_network

<a id="canonical-0201300111332021-1302021030321323-3333021213202233-1200021003223030-1022212320212310-3313023211131022-0133111023222130-3313012200323102"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

<a id="canonical-3121211221233000-0303222003301000-2010121021220303-3113213323010233-3122302123200203-1100233323313033-1122123230001323-3111333213331332"></a>

## Direct properties — inside_network / 220300230000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012130233101111-1032103102231202-3313131011131301-1023000003000122-1200121103122113-2203311113123102-3230322313102003-1021312012222013"></a>

## Next pages — inside_network / 220300230000 / 4

- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0311301122230130-3202322330221100-3010223302023213-3000100231023333-1000112032103002-1132120213013212-3131232033220233-0012222212112021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123223330331203-1320031122301221-2021322332132330-1121001103211301-0213320232102212-3212200220222323-0233123003230000-1013223312121232"></a>

## default_pool.origin_servers.private_name.outside_network — outside_network / 220110203113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- default_pool.origin_servers.private_name.outside_network

<a id="canonical-1303232330031210-1303120031000032-3121101321212101-3011030103331201-0300021021232131-1323122223113303-0302002200012013-2130332303030222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

<a id="canonical-2102200311323130-3112210220012212-0210123331130110-1231132213001222-1233001223123033-2321131210123302-1233221211320202-1232321033230023"></a>

## Direct properties — outside_network / 220110203113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123230212213322-1002222233212033-2322111000133202-1200232001032221-0212231032201211-3020310023300213-3010221030020322-3201131323221302"></a>

## Next pages — outside_network / 220110203113 / 4

- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3322321310112123-3331203210113301-2130130132133122-2032103130113112-2102312033030203-0213230232130021-0013232233102032-3321020322030233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211333233120030-2211100212213312-3332210233112313-3131322331213002-0003303012312310-0020121310121332-3013013332323330-1223120222220121"></a>

## default_pool.origin_servers.private_name.segment — segment / 302321001221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- default_pool.origin_servers.private_name.segment

<a id="canonical-3022021303003031-2123233030302002-2011220010333131-3022300113220132-0221301023133002-3320202022222133-2032023213100133-1133213011213000"></a>

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
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-1101302301003120-2223131202302022-1321121221232021-1313100303320312-3003202101303221-0012112131330203-1113100133030032-1031311033033121"></a>

## Direct properties — segment / 302321001221 / 3

<a id="canonical-3133220022030122-2001112013033120-0301332133201113-0221312333132323-3003011133211330-2201313321201230-0030200011220110-2201102232303312"></a>

<a id="canonical-0210030033230132-0320212110023020-3231131002210033-0220211331011113-2031223312333322-3321031323030230-3103131233001213-2003131130303302"></a>

## name property — segment / 302321001221 / 4

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

<a id="canonical-3131032023101000-1101313122023232-3202002233300101-3212221023331011-0323111231002113-0301232010020211-2223302303302233-2022212302102102"></a>

<a id="canonical-1213102010232200-3322302122313021-1001022230010210-3212030030122123-0010130120222333-3311020000101233-0320223231132301-0100230323000012"></a>

## namespace property — segment / 302321001221 / 5

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

<a id="canonical-3113201210120312-0120312220002321-3103203002233112-0010313123233212-3321122300201121-2031201120003100-2303300222102011-1233211012301031"></a>

<a id="canonical-3232211001230310-1302210032102102-2320230310331302-0123310102232323-0132111301222003-2121331023202120-3013103321310303-3030222221313133"></a>

## tenant property — segment / 302321001221 / 6

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

<a id="canonical-1213001203100003-2223313111312003-3333330022333023-3013232113201203-2313111231302300-0011200101013020-3322023333030323-1222213320200123"></a>

## Next pages — segment / 302321001221 / 7

- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0211232012223111-1001231003113023-1212212001033231-0031131311031121-2101212231112023-0332221013322020-3210130213310212-0001122020323330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121100023202223-3112023213031201-1100112121113300-2230320101001233-0032230233231231-2203221123030003-1222202121303021-0213021322321011"></a>

## default_pool.origin_servers.private_name.site_locator — site_locator / 132300303011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- default_pool.origin_servers.private_name.site_locator

<a id="canonical-0302202221122101-1020001003021233-0032023122222011-1221020112122103-2132010320021102-2120111200010321-0300002230031221-2230323332022001"></a>

Type: `"object"`. single nested block, Optional.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131302300321220-3301112303003203-3202203112233332-3203231113130313-3120322221300300-0223230003120212-3202010111011203-2102111013132102"></a>

## Direct properties — site_locator / 132300303011 / 3

- [site](resources--http_loadbalancer--reference--group-017.md#canonical-1332313011120202-1010002020031311-2303320320310102-1231133210103131-2100011110331232-2120331112022210-3101220212213003-1023322200303033): complete subsection reference.

- [virtual_site](resources--http_loadbalancer--reference--group-017.md#canonical-3130200003012212-2200212123310111-3112101032311330-2012320313331221-2010131132322310-0221123031111332-1110002013020130-0321020223010200): complete subsection reference.

<a id="canonical-3010211231002011-3202230220221003-3312132113012313-2023020332301032-0133223113021100-2321011011331030-2020220013031131-2032223331201123"></a>

## Next pages — site_locator / 132300303011 / 4

- [default_pool.origin_servers.private_name.site_locator.site](resources--http_loadbalancer--reference--group-017.md#canonical-1332313011120202-1010002020031311-2303320320310102-1231133210103131-2100011110331232-2120331112022210-3101220212213003-1023322200303033)
- [default_pool.origin_servers.private_name.site_locator.virtual_site](resources--http_loadbalancer--reference--group-017.md#canonical-3130200003012212-2200212123310111-3112101032311330-2012320313331221-2010131132322310-0221123031111332-1110002013020130-0321020223010200)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1332313011120202-1010002020031311-2303320320310102-1231133210103131-2100011110331232-2120331112022210-3101220212213003-1023322200303033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212303123122001-1023332133231033-2300333001022031-3111000211232303-1200231012031331-1130102201120013-2230132022303333-3223201012001201"></a>

## default_pool.origin_servers.private_name.site_locator.site — site / 132001122103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [default_pool.origin_servers.private_name.site_locator](resources--http_loadbalancer--reference--group-017.md#canonical-0211232012223111-1001231003113023-1212212001033231-0031131311031121-2101212231112023-0332221013322020-3210130213310212-0001122020323330)
- default_pool.origin_servers.private_name.site_locator.site

<a id="canonical-3011232212001132-0012120012320321-3203323001312230-1332203313101030-0122122310023210-1113102111312203-1013001220102211-1123030312100101"></a>

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

<a id="canonical-1212202213030213-3313100131131111-1122202323000010-3210121332000221-1300331302200220-1100100300300230-0333002012222133-3201203231003103"></a>

## Direct properties — site / 132001122103 / 3

<a id="canonical-1112010021112123-3201102202302321-0220032131233133-0310202121331023-0100030103201033-0203230023013110-0122301130111133-2332022121003203"></a>

<a id="canonical-1101211011012322-1311333112031322-0102123032221300-2203333012000131-1002230111211332-1300103213202231-2030303320210301-1301102311101322"></a>

## name property — site / 132001122103 / 4

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

<a id="canonical-1202300301222100-2012022221333013-0130102030301210-3211110031202011-1112130231303231-3301233000303100-3021130001013120-0101123122323123"></a>

<a id="canonical-3100032123321020-1121221202330320-0333101130021202-3300132202000033-1032122203300012-2010032112031202-3131100101132331-2001020100331103"></a>

## namespace property — site / 132001122103 / 5

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

<a id="canonical-3321233211033221-1112031222131010-0111133332112121-1020200133103212-0010332012112320-0010120322322131-1312102121302022-0321131212103032"></a>

<a id="canonical-2110230232330322-2112202202003302-1032220301311230-3111312311103011-3131131103311131-2021213123211322-3211213210123033-0221131013023103"></a>

## tenant property — site / 132001122103 / 6

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

<a id="canonical-2020121202312020-3333211020030311-1010221003331103-2011020301200213-0311003001103001-2222332231110111-1021121120001300-3032332330101003"></a>

## Next pages — site / 132001122103 / 7

- [default_pool.origin_servers.private_name.site_locator](resources--http_loadbalancer--reference--group-017.md#canonical-0211232012223111-1001231003113023-1212212001033231-0031131311031121-2101212231112023-0332221013322020-3210130213310212-0001122020323330)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3130200003012212-2200212123310111-3112101032311330-2012320313331221-2010131132322310-0221123031111332-1110002013020130-0321020223010200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223221322210212-3111322032302000-2121113112333322-0132022222200010-0213301012301120-2113011120303210-3302201210311012-3310102023110102"></a>

## default_pool.origin_servers.private_name.site_locator.virtual_site — virtual_site / 222322231211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [default_pool.origin_servers.private_name.site_locator](resources--http_loadbalancer--reference--group-017.md#canonical-0211232012223111-1001231003113023-1212212001033231-0031131311031121-2101212231112023-0332221013322020-3210130213310212-0001122020323330)
- default_pool.origin_servers.private_name.site_locator.virtual_site

<a id="canonical-0302210233131221-3232120302303310-1133033103111201-3300011212221102-0101010023112321-2321222002132013-0223002103201020-2213012313131001"></a>

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

<a id="canonical-3101303321313233-1332132103232211-2020231332130212-3222100133131020-1123001013303132-1003210112323001-3120322212112013-0123321000123321"></a>

## Direct properties — virtual_site / 222322231211 / 3

<a id="canonical-3032210012121211-0020022121323101-3002112200130113-3321113021123233-0032313032222222-1311120011003331-2323100232302320-3122123332213031"></a>

<a id="canonical-0023221331103231-0221002200302031-3112321310313101-1332310112001310-0333200331123113-3132133033100020-3113313120332223-2300031200310101"></a>

## name property — virtual_site / 222322231211 / 4

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

<a id="canonical-3200302122000310-3220202232022232-0021023220310201-3133232302030220-1132131001032020-1221023113200231-2133233212113203-2322133101011120"></a>

<a id="canonical-1101330232220301-0011201230331132-3120133033211323-1103223203012001-3112232002001320-0310220101133003-0110200132130202-0113002113212101"></a>

## namespace property — virtual_site / 222322231211 / 5

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

<a id="canonical-2112200130122133-3333330303023123-2323223000030303-3321121222100313-3331232330003120-1300230212232230-1103221010002100-3031003233322001"></a>

<a id="canonical-1331303003020333-2003301103323222-2030100102111233-3021232233230220-2321203001302103-3322211011013001-1200121112120113-1203000023031021"></a>

## tenant property — virtual_site / 222322231211 / 6

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

<a id="canonical-3332133023022203-2313333001322232-3032112012101002-2322313132131201-2220132220210132-3301011313302221-0231201130313122-3100301033122132"></a>

## Next pages — virtual_site / 222322231211 / 7

- [default_pool.origin_servers.private_name.site_locator](resources--http_loadbalancer--reference--group-017.md#canonical-0211232012223111-1001231003113023-1212212001033231-0031131311031121-2101212231112023-0332221013322020-3210130213310212-0001122020323330)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0312130022121113-1033003212331120-0023101102200212-1312013113201020-0203311020213123-0223031101003330-3232310110011201-2312012230001330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201320020130103-2220032303002212-3303002110102321-1112300021103311-2112232130102210-3300020310113320-3222323021311201-1233023333320310"></a>

## default_pool.origin_servers.private_name.snat_pool — snat_pool / 122231003133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- default_pool.origin_servers.private_name.snat_pool

<a id="canonical-2320301120230101-1321101213212110-0231221300332001-3023123200110103-1221133312100032-3322323102022310-0230013131131311-3113211013332331"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122111102010122-1222300101013301-0003211300223300-0112302003003322-3131111103212013-1001300312301200-3202223300111112-1021213231101300"></a>

## Direct properties — snat_pool / 122231003133 / 3

- [no_snat_pool](resources--http_loadbalancer--reference--group-017.md#canonical-1221311033021110-0022231221001322-3333312321200100-2220310323023132-2323323222132122-3101233013033220-0133113303032131-3110200300020302): complete subsection reference.

- [snat_pool](resources--http_loadbalancer--reference--group-017.md#canonical-0023201330313033-3012311220001023-3020110330231121-1123002133300333-0230303300110332-2203000201021020-1110113003100210-2301230331311013): complete subsection reference.

<a id="canonical-3002033011231010-2301022330023131-3113201100120000-0112102101220223-3230021030013230-2312101121001220-1323020003202232-3223000000132130"></a>

## Next pages — snat_pool / 122231003133 / 4

- [default_pool.origin_servers.private_name.snat_pool.no_snat_pool](resources--http_loadbalancer--reference--group-017.md#canonical-1221311033021110-0022231221001322-3333312321200100-2220310323023132-2323323222132122-3101233013033220-0133113303032131-3110200300020302)
- [default_pool.origin_servers.private_name.snat_pool.snat_pool](resources--http_loadbalancer--reference--group-017.md#canonical-0023201330313033-3012311220001023-3020110330231121-1123002133300333-0230303300110332-2203000201021020-1110113003100210-2301230331311013)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1221311033021110-0022231221001322-3333312321200100-2220310323023132-2323323222132122-3101233013033220-0133113303032131-3110200300020302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002321211130231-0032101320333212-2012032100120021-1311132021202303-0322113303001003-0022202323301201-0011003121303100-2013212012133001"></a>

## default_pool.origin_servers.private_name.snat_pool.no_snat_pool — no_snat_pool / 231121100203 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [default_pool.origin_servers.private_name.snat_pool](resources--http_loadbalancer--reference--group-017.md#canonical-0312130022121113-1033003212331120-0023101102200212-1312013113201020-0203311020213123-0223031101003330-3232310110011201-2312012230001330)
- default_pool.origin_servers.private_name.snat_pool.no_snat_pool

<a id="canonical-3010011311112200-0330331122233210-0203123023122111-3231232220230010-1333302213032132-3013220010023111-1320001113233330-1301233110031332"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

<a id="canonical-2100033010303301-3113123120223020-1103122302302230-3233020210212212-3322332111330002-1022112223001023-1032102201333232-2223203011210212"></a>

## Direct properties — no_snat_pool / 231121100203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332212231330003-3302033121200122-1121310101300100-1210313312332133-0103132312122302-3211321201320131-0101000000232303-0013123211032103"></a>

## Next pages — no_snat_pool / 231121100203 / 4

- [default_pool.origin_servers.private_name.snat_pool](resources--http_loadbalancer--reference--group-017.md#canonical-0312130022121113-1033003212331120-0023101102200212-1312013113201020-0203311020213123-0223031101003330-3232310110011201-2312012230001330)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0023201330313033-3012311220001023-3020110330231121-1123002133300333-0230303300110332-2203000201021020-1110113003100210-2301230331311013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232102210211131-0030103333320023-2110231230311201-1300310202020011-0330120131132113-3302131020301023-3220233333011021-2220002123003111"></a>

## default_pool.origin_servers.private_name.snat_pool.snat_pool — snat_pool / 221230103120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.private_name](resources--http_loadbalancer--reference--group-016.md#canonical-3331022330001133-3221230332331003-1230030223132033-3303010102000230-1013131023131203-1232013221202321-3102123331031201-2211013220321032)
- [default_pool.origin_servers.private_name.snat_pool](resources--http_loadbalancer--reference--group-017.md#canonical-0312130022121113-1033003212331120-0023101102200212-1312013113201020-0203311020213123-0223031101003330-3232310110011201-2312012230001330)
- default_pool.origin_servers.private_name.snat_pool.snat_pool

<a id="canonical-3103122222102210-1012332212123003-1330010023111300-1221023101111210-0011332012303200-1113323033323333-0303222033220103-2302113302130333"></a>

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311332320112313-3120223101203002-2012333311312220-1232023331131233-3323323013333330-2023110122311303-0320101132133300-1202033033221013"></a>

## Direct properties — snat_pool / 221230103120 / 3

<a id="canonical-1012311200321331-1320031233130221-2302010222201100-0020311210310333-3002011300322112-0010103000322223-0323312202323330-0123310223012013"></a>

<a id="canonical-1112021000232322-1303321302200102-1231102320230031-2132300120201012-0313012201011210-1032213112001330-0000322133203201-2300302001233213"></a>

## prefixes property — snat_pool / 221230103120 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2032030111112220-2330221310023331-3033232332010132-1013100313212221-1320010303331013-2101032101220232-3313101220331011-1210102201000121"></a>

## Next pages — snat_pool / 221230103120 / 5

- [default_pool.origin_servers.private_name.snat_pool](resources--http_loadbalancer--reference--group-017.md#canonical-0312130022121113-1033003212331120-0023101102200212-1312013113201020-0203311020213123-0223031101003330-3232310110011201-2312012230001330)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1123323200112132-2011301210333320-0230333321233220-3120200223212323-0221012302113321-1102011330332121-0221103113230111-0011130112303131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320121003222032-3310311231032211-2212031120013103-1010112113121012-0202231230232310-1231023111303210-2212313121302320-3313002301322323"></a>

## default_pool.origin_servers.public_ip — public_ip / 102333031310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.public_ip

<a id="canonical-2233020311003121-3121333321021131-3001122003132113-1301222300022020-1120113331012303-0312322233233120-0102201110013213-2020312300030233"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0202231230210022-2100333022113220-1111311312231121-0312010203322231-2201221300203121-3223303233313232-0232312001101020-0223011003002222"></a>

## Direct properties — public_ip / 102333031310 / 3

<a id="canonical-1000013300233110-0012120220113212-2010132312303303-3231031102201032-2131001003003211-1012030320130230-1011300130213021-1200312020010213"></a>

<a id="canonical-3011211012231330-2322111303030223-3001101010111303-1101101300033032-0212303031320023-3221231332302323-3323230212031000-2220131212320120"></a>

## ip property — public_ip / 102333031310 / 4

Type: `"string"`. Optional.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Upstream description:

Exclusive with \[\] Public IPv4 address.

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

<a id="canonical-0103203131123002-0111022312212030-0332300310010232-1201233321020012-0301013000222312-3020012320321231-1131000032301103-0100231220001003"></a>

## Next pages — public_ip / 102333031310 / 5

- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3302320302223102-1023033333230010-2222202133303121-2101111202221230-2113030322031221-1332101010000231-1202031111300230-1003313332002300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302123012033012-3230112200032102-1222323103130102-1002102230210231-1301002131301222-2033203301323333-1103323003201233-1100223202233321"></a>

## default_pool.origin_servers.public_name — public_name / 221310230221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.public_name

<a id="canonical-1001132011011000-0202223032003200-2201310223322201-2301022130002302-2310300302010222-3101213112032003-2323032031122031-2013322101233301"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public DNS name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
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
public_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320331003032222-1030022331322123-1013130302320223-3130122222212223-3232002013020003-0311202330031201-1110201231220202-3032103030211312"></a>

## Direct properties — public_name / 221310230221 / 3

<a id="canonical-0333302200003211-1210112202300202-0220113030021222-2000123111031031-3110333222010230-3000302321313102-2001003200333210-3202302130002112"></a>

<a id="canonical-3101211310012233-0200200320110333-0230113332333022-1230231110233133-3213232112122021-3310123103021213-2301101113022331-2330300310312313"></a>

## dns_name property — public_name / 221310230221 / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

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
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0023123303332312-3122011222213000-0211231320220230-2331322222311001-1232100123220010-3223120300221100-2221002333213110-0212123331231301"></a>

<a id="canonical-2003220030321030-2111130330100203-1222322223221221-3333012010332001-1311123223322330-3102332212102010-1311230000321103-1202302032321232"></a>

## refresh_interval property — public_name / 221310230221 / 5

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-0332212300311332-2033330331230131-2221222021012200-0101023032001111-2112313322313102-0112010311321303-0331121221213030-1211212221012103"></a>

## Next pages — public_name / 221310230221 / 6

- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0022310013130030-2111303220210213-2020201222032233-1120103131300012-1111203033323012-1021113330333233-3121232112020030-0200301203202122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203231103120212-1232312001313221-2223333132133133-2303213133032222-0300130320322121-0221221001112112-2220123220301330-2221120101331320"></a>

## default_pool.origin_servers.vn_private_ip — vn_private_ip / 301113130110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.vn_private_ip

<a id="canonical-2210121123312130-3103200230220021-0013001213201023-0331001313021131-0203202312200302-2012112310012022-1230221032133132-3123200303232133"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with IP on Virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-virtual_network_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
vn_private_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003200211320312-1320032311101033-2330113223011223-0033103112002321-2113112102120033-1120132032013330-1332203231313123-2031321333100332"></a>

## Direct properties — vn_private_ip / 301113130110 / 3

<a id="canonical-1101003013130331-3001011200313212-0300323212213213-2311110210321110-3100023320302121-0303100212220002-3112330322123123-0310113311002322"></a>

<a id="canonical-3113323002320003-1031022101101333-1211321021223100-0331302302322301-1220030333203133-0101033330012212-0223333332103022-2302233010323011"></a>

## ip property — vn_private_ip / 301113130110 / 4

Type: `"string"`. Optional.

IPv4. Exclusive with \[\] IPv4 address.

Upstream description:

Exclusive with \[\] IPv4 address.

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

- [virtual_network](resources--http_loadbalancer--reference--group-017.md#canonical-0022111032132322-0332003132203223-2021223302101220-2331130303203021-3023123313013130-3330101201022111-1223020213011322-1133103323103121): complete subsection reference.

<a id="canonical-3302002211111212-3121302223120121-2132010103112222-1213121130322013-0330100310033202-2011322300000223-3133232003213111-3213200033320232"></a>

## Next pages — vn_private_ip / 301113130110 / 5

- [default_pool.origin_servers.vn_private_ip.virtual_network](resources--http_loadbalancer--reference--group-017.md#canonical-0022111032132322-0332003132203223-2021223302101220-2331130303203021-3023123313013130-3330101201022111-1223020213011322-1133103323103121)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0022111032132322-0332003132203223-2021223302101220-2331130303203021-3023123313013130-3330101201022111-1223020213011322-1133103323103121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031020010100311-3021333303030010-0211302300122020-0020232100111200-3233111202131312-0102002210003132-3132231212300130-2313212300331323"></a>

## default_pool.origin_servers.vn_private_ip.virtual_network — virtual_network / 300002032012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.vn_private_ip](resources--http_loadbalancer--reference--group-017.md#canonical-0022310013130030-2111303220210213-2020201222032233-1120103131300012-1111203033323012-1021113330333233-3121232112020030-0200301203202122)
- default_pool.origin_servers.vn_private_ip.virtual_network

<a id="canonical-2222011203123233-1231300123133010-1302300210111000-2103002203232133-2031123112002200-3321113323303113-0321013120333312-0303331013203100"></a>

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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231330321103330-3222021321030102-1031333231321010-0213132030313201-0012322010000322-3213133123030021-2100331330221210-0333212213212213"></a>

## Direct properties — virtual_network / 300002032012 / 3

<a id="canonical-0111132101010223-0301022232122322-1311111222032122-1031021203103310-1302233303013320-3100120111011302-1323033103102332-0030131231201332"></a>

<a id="canonical-1022211103320111-1101210012120010-2013130133302110-3323121231031320-3131300111221203-2213220101023020-2322023222200012-3312312101032232"></a>

## name property — virtual_network / 300002032012 / 4

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

<a id="canonical-3321123122101032-3011112312021012-3121100022223300-3000220113012023-2223003010233133-0012211113000110-2101303132201031-2012321022013003"></a>

<a id="canonical-2320122320002332-0032230312211200-3300022120210333-2031000203332200-3021011113130331-0013130222131310-3010203120022230-2202311331310011"></a>

## namespace property — virtual_network / 300002032012 / 5

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

<a id="canonical-3113310013103302-1211100331111300-2013123333022101-0202332000302211-3333112003322321-0131111103312020-3021011210211103-3130003323100213"></a>

<a id="canonical-3311002102232213-3022211001031001-1001013213132212-1032202201302120-2322300210021311-2323231323200313-2213032301002222-2123032032133021"></a>

## tenant property — virtual_network / 300002032012 / 6

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

<a id="canonical-1003110203301232-0010012301112110-3322312200222213-2312031030112000-3333321210210302-1033003031120100-0230101022030101-0210120010302213"></a>

## Next pages — virtual_network / 300002032012 / 7

- [default_pool.origin_servers.vn_private_ip](resources--http_loadbalancer--reference--group-017.md#canonical-0022310013130030-2111303220210213-2020201222032233-1120103131300012-1111203033323012-1021113330333233-3121232112020030-0200301203202122)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2101033313001230-3100221301111002-2011003233203011-1100130312221121-1002231322001133-0002023210303033-3312032213131120-2332031232122312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223011300100011-2120100033322020-0001033103022102-0201100331322113-2113030033113200-2212021322300120-2332020021230023-2131100310313012"></a>

## default_pool.origin_servers.vn_private_name — vn_private_name / 120123303202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- default_pool.origin_servers.vn_private_name

<a id="canonical-2220132331000203-3110331213121213-1230131221203200-3213321231111321-1213102122203002-2020001233110101-1102020200321030-2220133333332113"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with DNS name on Virtual Network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
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
vn_private_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123212113221032-2301230320033132-2012301202122100-1231013020331221-3101110233302322-0113231203133023-1100023220323210-0223032100012033"></a>

## Direct properties — vn_private_name / 120123303202 / 3

<a id="canonical-1311303101122130-3300123300122221-3301210122222203-1012132302312301-2112320111001300-2230213032023133-0132233212221123-1131023230213030"></a>

<a id="canonical-2210113132132022-1011311030212003-1122233213100100-3230233101113231-3011310112100022-3122120312121302-0311211203133212-1120322001001121"></a>

## dns_name property — vn_private_name / 120123303202 / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [private_network](resources--http_loadbalancer--reference--group-017.md#canonical-0203020320000112-0102231033232010-1223102323000102-2100120021102020-0320320120202210-2013203021122133-2011000000301030-2011300203301111): complete subsection reference.

<a id="canonical-3333311001013132-3333303032210221-1301123321301321-2021100012303323-0203300200113122-0303231103230103-2301222202213013-1133013303222223"></a>

## Next pages — vn_private_name / 120123303202 / 5

- [default_pool.origin_servers.vn_private_name.private_network](resources--http_loadbalancer--reference--group-017.md#canonical-0203020320000112-0102231033232010-1223102323000102-2100120021102020-0320320120202210-2013203021122133-2011000000301030-2011300203301111)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0203020320000112-0102231033232010-1223102323000102-2100120021102020-0320320120202210-2013203021122133-2011000000301030-2011300203301111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023200301022111-0203023133002333-2011021330101000-2021131020301130-1112032312131203-3013023122103020-1012320322113233-1231122120110233"></a>

## default_pool.origin_servers.vn_private_name.private_network — private_network / 122301020010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.origin_servers](resources--http_loadbalancer--reference--group-016.md#canonical-3032210112122003-3001300202101031-2123213322103223-0023311332223220-3323333210222101-3133123322012320-1210021330031123-1312333220120031)
- [default_pool.origin_servers.vn_private_name](resources--http_loadbalancer--reference--group-017.md#canonical-2101033313001230-3100221301111002-2011003233203011-1100130312221121-1002231322001133-0002023210303033-3312032213131120-2332031232122312)
- default_pool.origin_servers.vn_private_name.private_network

<a id="canonical-2331213320121203-1220232222333230-1322222102013220-1112010332020110-1020313232110022-0133002120120112-0112031123001302-2113033022000322"></a>

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
private_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-1123033330210120-1101010110111212-1021210130110330-0302210113130130-0003312031222022-1101021300312200-3221222121213200-2130003330022120"></a>

## Direct properties — private_network / 122301020010 / 3

<a id="canonical-0011022111002210-1100033313003213-0132120333103120-2203231211201231-3233131011300303-1022121012131113-3101000020303213-0111213023203223"></a>

<a id="canonical-3311033330330213-0100132332301200-2322122221133201-3320022301101332-3022210133122233-2101033233323100-3022102030201123-2332322302000102"></a>

## name property — private_network / 122301020010 / 4

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

<a id="canonical-1001300112001202-2021132333113100-0200010200333013-3113203220203232-3013330013000200-2221122112103213-3100220033332133-3312211111203200"></a>

<a id="canonical-3100133213001200-3032032011113311-3213012330120111-1001323012213021-1032103220322011-1001032113033210-1230113102032222-3202102330200020"></a>

## namespace property — private_network / 122301020010 / 5

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

<a id="canonical-3223331230003012-0130313330232130-0230233311212200-3022210100300310-1320323311020300-0120030112030103-2030321023330030-1033123110000302"></a>

<a id="canonical-1313312213020023-0233002003201033-0320302003301122-2113102123123103-3120330023330100-1001010222232131-1313310303000310-3222203331123003"></a>

## tenant property — private_network / 122301020010 / 6

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

<a id="canonical-1223031103312211-2301310330230333-3010303330300201-3321301213233023-3310232123112011-3331222030012200-1323320110220010-1103221030131133"></a>

## Next pages — private_network / 122301020010 / 7

- [default_pool.origin_servers.vn_private_name](resources--http_loadbalancer--reference--group-017.md#canonical-2101033313001230-3100221301111002-2011003233203011-1100130312221121-1002231322001133-0002023210303033-3312032213131120-2332031232122312)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1232100333033333-0221010121032103-2010033203301010-1223212021203320-2001102032013210-0333301222330003-1231200013113223-2010211132110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232121322112020-3312310121200210-3000230022220021-1331212103302202-0023012230110202-3001202023103203-0320130021311011-0201212130322313"></a>

## default_pool.same_as_endpoint_port — same_as_endpoint_port / 031002331012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.same_as_endpoint_port

<a id="canonical-2033213221230001-2133023101322313-1230200120100220-1012323122020200-2222200231023032-1012202202131121-3102330121032032-1030231322320113"></a>

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
same_as_endpoint_port = {}
```

<a id="canonical-3202331212130020-2311010111101111-1023002121121133-1333201013130120-0202312200132012-3101100230033310-1213203110322111-2023212310131032"></a>

## Direct properties — same_as_endpoint_port / 031002331012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2013010031000300-3010013021331203-1100120112311012-1312213223302230-0230311110113331-0331303101002310-0332330321333201-2213200131331233"></a>

## Next pages — same_as_endpoint_port / 031002331012 / 4

- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1112003010300102-0001210211031011-0021210310211031-2120013323123301-1203112022230211-2311310013202331-2323031211322202-1321003211023230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302310033303032-3323330200011311-3000120333200003-0010033311321113-1221202201023033-0110230231300033-0121123233023220-3122102133100002"></a>

## default_pool.upstream_conn_pool_reuse_type — upstream_conn_pool_reuse_type / 220201022223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.upstream_conn_pool_reuse_type

<a id="canonical-0123322202101000-1321100330232001-2300223133032122-2100000323122112-0020213022331320-3213210102203003-0120120302223030-2210302332110313"></a>

Type: `"object"`. single nested block, Optional.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_conn_pool_reuse",
    "enable_conn_pool_reuse")}
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
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

Terraform syntax:

```terraform
upstream_conn_pool_reuse_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103133012031013-0323133110000313-0313313231001021-1311321012123202-2103323023232212-0320033101031230-0112223030020203-3031302230332301"></a>

## Direct properties — upstream_conn_pool_reuse_type / 220201022223 / 3

- [disable_conn_pool_reuse](resources--http_loadbalancer--reference--group-017.md#canonical-3022003332112203-3313121000200322-1020012101323212-1011010321303313-0220213231201133-3123213123103002-0213302001100112-0013020212002000): complete subsection reference.

- [enable_conn_pool_reuse](resources--http_loadbalancer--reference--group-017.md#canonical-0122233311301230-2033102130212202-1303001033113312-2303130331113333-0121303320101230-1213010322331113-0123121210000310-3310332303220133): complete subsection reference.

<a id="canonical-2123313022002011-2133113031031100-1103010222110333-3122032313301130-2101111332000113-2310033313013333-0103121113302331-3201011313111230"></a>

## Next pages — upstream_conn_pool_reuse_type / 220201022223 / 4

- [default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse](resources--http_loadbalancer--reference--group-017.md#canonical-3022003332112203-3313121000200322-1020012101323212-1011010321303313-0220213231201133-3123213123103002-0213302001100112-0013020212002000)
- [default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse](resources--http_loadbalancer--reference--group-017.md#canonical-0122233311301230-2033102130212202-1303001033113312-2303130331113333-0121303320101230-1213010322331113-0123121210000310-3310332303220133)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3022003332112203-3313121000200322-1020012101323212-1011010321303313-0220213231201133-3123213123103002-0213302001100112-0013020212002000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131312001200003-2021023320220111-2032132200010131-0220211001300331-3312203013130002-2111213030200210-2333023121002031-1333020030232022"></a>

## default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse — disable_conn_pool_reuse / 212221320000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.upstream_conn_pool_reuse_type](resources--http_loadbalancer--reference--group-017.md#canonical-1112003010300102-0001210211031011-0021210310211031-2120013323123301-1203112022230211-2311310013202331-2323031211322202-1321003211023230)
- default_pool.upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-1033210212033123-0322132311231131-2003001201100033-1222031301130021-3020120221133033-0133021221203001-1300033210001320-3322032323211021"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable conn pool reuse.

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
disable_conn_pool_reuse = {}
```

<a id="canonical-2220211310332111-1213112231331111-0001232212332122-0120132320032133-1210002333110033-2000010100033312-2323021312231110-1212121022301023"></a>

## Direct properties — disable_conn_pool_reuse / 212221320000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321013030300203-3032223201120313-2131233310130211-1003223230331330-2211031220012033-3210123023211202-1133321022021323-0113011030220302"></a>

## Next pages — disable_conn_pool_reuse / 212221320000 / 4

- [default_pool.upstream_conn_pool_reuse_type](resources--http_loadbalancer--reference--group-017.md#canonical-1112003010300102-0001210211031011-0021210310211031-2120013323123301-1203112022230211-2311310013202331-2323031211322202-1321003211023230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0122233311301230-2033102130212202-1303001033113312-2303130331113333-0121303320101230-1213010322331113-0123121210000310-3310332303220133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331123111101231-1101203323231233-3221031333103203-0123100032212313-1310030330200103-2020311201303222-2202022113301210-1023032003102212"></a>

## default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse — enable_conn_pool_reuse / 320123323232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.upstream_conn_pool_reuse_type](resources--http_loadbalancer--reference--group-017.md#canonical-1112003010300102-0001210211031011-0021210310211031-2120013323123301-1203112022230211-2311310013202331-2323031211322202-1321003211023230)
- default_pool.upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-3113313020022133-0330300231330231-3032212110201123-0231101332023132-3121312031113003-3020230103323212-3303122210032002-3213330300132320"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable conn pool reuse.

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
enable_conn_pool_reuse = {}
```

<a id="canonical-1300103312310220-0323313012202222-1123322212033302-0220321100210021-2000312210103312-0123113012012003-1302202330120012-1233103233232101"></a>

## Direct properties — enable_conn_pool_reuse / 320123323232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301113222330113-0222323120211133-3032210021212113-1230000131220031-3322300300130201-3131201322300122-1200212320212031-0300011011301010"></a>

## Next pages — enable_conn_pool_reuse / 320123323232 / 4

- [default_pool.upstream_conn_pool_reuse_type](resources--http_loadbalancer--reference--group-017.md#canonical-1112003010300102-0001210211031011-0021210310211031-2120013323123301-1203112022230211-2311310013202331-2323031211322202-1321003211023230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103122312033133-0232020122221200-3322322232313322-3012123330332032-0000121120031111-1021333002233122-0311211200123330-1203031110312330"></a>

## default_pool.use_tls — use_tls / 102333211323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.use_tls

<a id="canonical-2332213023031332-1013323323031210-0100123112302011-2330013230103010-0102132311330320-2010003302003033-2313120303330331-0312102122123221"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Upstream description:

Upstream TLS Parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_session_key_caching",
    "disable_session_key_caching"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("disable_sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls_obj"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "use_server_verification"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "volterra_trusted_ca"),
  validators.ConflictingObjectAttributes("sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("use_mtls",
    "use_mtls_obj"),
  validators.ConflictingObjectAttributes("use_server_verification",
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
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333032221123011-1312320102032030-0222130012112100-0300313100213112-3310300033310112-3300102003222320-0201120203223113-2323001120313230"></a>

## Direct properties — use_tls / 102333211323 / 3

- [default_session_key_caching](resources--http_loadbalancer--reference--group-017.md#canonical-2202223033321103-1023313022220032-1100110331023310-2000233222211323-0101221310312021-3102332013202020-0101210122012110-2103210131213301): complete subsection reference.

- [disable_session_key_caching](resources--http_loadbalancer--reference--group-017.md#canonical-1013331020231333-0202203320300021-1232330310100320-2300031020020203-1123330312131323-3001113200312303-0122030002201222-3233033231132121): complete subsection reference.

- [disable_sni](resources--http_loadbalancer--reference--group-017.md#canonical-2300213333200132-0123123213203223-0200321201003332-1133131212130322-0120022102301322-3132311102333002-1200211200102323-3233332031012120): complete subsection reference.

<a id="canonical-1213330013021023-0213323002023222-2310113003121231-0010330232301332-1211121220311200-0331022332003023-0330221202302201-1302013110331311"></a>

<a id="canonical-0001021102223102-1002303213321120-2321132020021113-3001122012301201-3312032111232212-2323012221313122-2233031230223221-1213112022113001"></a>

## max_session_keys property — use_tls / 102333211323 / 4

Type: `"number"`. Optional.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

- [no_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-0303331032002210-0312010001012120-1330112001013310-0010320010222322-2132130013000332-1231131311012023-3021102123223331-2330222230010032): complete subsection reference.

- [skip_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-0303332132110311-1223303331210202-2030002213310231-2101021031103002-3222212230031331-2213230130330002-3303131031031310-2010310111301030): complete subsection reference.

<a id="canonical-3203103222000233-2001032322130211-2232220032202333-3332301033030323-2013113331033102-3033013031113121-0111203103331320-1122023110231123"></a>

<a id="canonical-2033212320030022-0100221130321013-0200000132130310-1211030202320102-2320322011110331-1323010121221010-2223110212233212-3233311221020220"></a>

## sni property — use_tls / 102333211323 / 5

Type: `"string"`. Optional.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112): complete subsection reference.

- [use_host_header_as_sni](resources--http_loadbalancer--reference--group-017.md#canonical-1322020121333230-3003002212033100-2332300001021121-1102021010223121-3131123312322320-1020132011323230-0202022322031312-3313322321233013): complete subsection reference.

- [use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031): complete subsection reference.

- [use_mtls_obj](resources--http_loadbalancer--reference--group-017.md#canonical-2011310303210021-1222202221103110-1130033232102230-1131030100322101-0032220031011132-0202012301223101-3013101110032002-0102022212323012): complete subsection reference.

- [use_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-3302122303122133-0323233203310323-3012112030002300-2003211111010032-2031330012231010-1011303111010020-0113222311123003-1002223100220030): complete subsection reference.

- [volterra_trusted_ca](resources--http_loadbalancer--reference--group-017.md#canonical-3321010323221300-3221232231012333-1000122333111322-3013010233100302-2311021323120212-1121301113331022-3222031002320022-3101011201003023): complete subsection reference.

<a id="canonical-1102101021232323-1331201231200030-3023001020222321-0200101111220020-3103210322012131-0310331032302110-1021100000012101-0303322211001312"></a>

## Next pages — use_tls / 102333211323 / 6

- [default_pool.use_tls.default_session_key_caching](resources--http_loadbalancer--reference--group-017.md#canonical-2202223033321103-1023313022220032-1100110331023310-2000233222211323-0101221310312021-3102332013202020-0101210122012110-2103210131213301)
- [default_pool.use_tls.disable_session_key_caching](resources--http_loadbalancer--reference--group-017.md#canonical-1013331020231333-0202203320300021-1232330310100320-2300031020020203-1123330312131323-3001113200312303-0122030002201222-3233033231132121)
- [default_pool.use_tls.disable_sni](resources--http_loadbalancer--reference--group-017.md#canonical-2300213333200132-0123123213203223-0200321201003332-1133131212130322-0120022102301322-3132311102333002-1200211200102323-3233332031012120)
- [default_pool.use_tls.no_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-0303331032002210-0312010001012120-1330112001013310-0010320010222322-2132130013000332-1231131311012023-3021102123223331-2330222230010032)
- [default_pool.use_tls.skip_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-0303332132110311-1223303331210202-2030002213310231-2101021031103002-3222212230031331-2213230130330002-3303131031031310-2010310111301030)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- [default_pool.use_tls.use_host_header_as_sni](resources--http_loadbalancer--reference--group-017.md#canonical-1322020121333230-3003002212033100-2332300001021121-1102021010223121-3131123312322320-1020132011323230-0202022322031312-3313322321233013)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls_obj](resources--http_loadbalancer--reference--group-017.md#canonical-2011310303210021-1222202221103110-1130033232102230-1131030100322101-0032220031011132-0202012301223101-3013101110032002-0102022212323012)
- [default_pool.use_tls.use_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-3302122303122133-0323233203310323-3012112030002300-2003211111010032-2031330012231010-1011303111010020-0113222311123003-1002223100220030)
- [default_pool.use_tls.volterra_trusted_ca](resources--http_loadbalancer--reference--group-017.md#canonical-3321010323221300-3221232231012333-1000122333111322-3013010233100302-2311021323120212-1121301113331022-3222031002320022-3101011201003023)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2202223033321103-1023313022220032-1100110331023310-2000233222211323-0101221310312021-3102332013202020-0101210122012110-2103210131213301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101113021000230-1223103212112113-0133210030320311-2333300211121023-0330023122210312-0101111210331011-3223031231102021-3121301323112032"></a>

## default_pool.use_tls.default_session_key_caching — default_session_key_caching / 211303111121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.default_session_key_caching

<a id="canonical-0121200021230013-1210110220011333-2010002013003130-0332331033121331-3210220321103133-0003331123200131-0031332131222032-3333001233022131"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
default when omitted.

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
default_session_key_caching = {}
```

<a id="canonical-1021013031000210-3022020000202210-3011120001232210-0320332103330032-2311133123132202-2211013321030012-2123301121131332-3200321000303333"></a>

## Direct properties — default_session_key_caching / 211303111121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210311200331130-2010330332123323-2111002332202033-0000232111132331-0201331312133303-3110200113002313-0011332323021133-0103022332032303"></a>

## Next pages — default_session_key_caching / 211303111121 / 4

- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1013331020231333-0202203320300021-1232330310100320-2300031020020203-1123330312131323-3001113200312303-0122030002201222-3233033231132121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233113233010221-1123312011010223-0130301030011223-0023012202112333-1102212333023102-3312310211311010-3322120231333133-2322320210033213"></a>

## default_pool.use_tls.disable_session_key_caching — disable_session_key_caching / 101102200313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.disable_session_key_caching

<a id="canonical-3133021332013103-1320013333102132-1221130111322301-2103212230210211-0003102122023023-0211011210211020-1221123121131000-3233100211221103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable session key caching.

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
disable_session_key_caching = {}
```

<a id="canonical-0232012112113300-0031002010330223-0300231302312212-2202333130320323-1012111000033212-1333001110101210-0200120331012202-1331312003032321"></a>

## Direct properties — disable_session_key_caching / 101102200313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212320000231003-1033020220120121-3122312223221130-0003213301123333-3003002203313130-2230132032211333-2120200333212331-3321211323312300"></a>

## Next pages — disable_session_key_caching / 101102200313 / 4

- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2300213333200132-0123123213203223-0200321201003332-1133131212130322-0120022102301322-3132311102333002-1200211200102323-3233332031012120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003031121321013-1200012110011211-2031322211000313-0103220012303122-3203012232100323-2220212031201000-2302201202001212-0032321331003111"></a>

## default_pool.use_tls.disable_sni — disable_sni / 221221323111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.disable_sni

<a id="canonical-2220212132201313-2030112331313111-2010033223000013-0212133231032022-3210020311230301-3211012122013333-3113202200333123-2122133012213030"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable sni.

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
disable_sni = {}
```

<a id="canonical-3033031121023332-1330311012101110-3331122021003320-0132003030322111-1312102230113122-0301000013132221-0031122332221102-1323013310210323"></a>

## Direct properties — disable_sni / 221221323111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303233120333222-1222301032233123-2322110103230312-1102003200023132-0102322213200022-0211101120110122-1021011312312011-2210021123221121"></a>

## Next pages — disable_sni / 221221323111 / 4

- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0303331032002210-0312010001012120-1330112001013310-0010320010222322-2132130013000332-1231131311012023-3021102123223331-2330222230010032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103203230022223-0322333321322102-3203221011111331-0313232321100123-1220231213020122-0200010311323301-1103001132212032-3302211010102333"></a>

## default_pool.use_tls.no_mtls — no_mtls / 301302110010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.no_mtls

<a id="canonical-1332101130102002-2013003030331203-2321313313211121-1313313112103031-3132102123231203-1113032031122220-1011112032010133-3230233321301131"></a>

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
no_mtls = {}
```

<a id="canonical-2033011212120230-0223013230312133-0200030012023020-1002332211301320-0023300003322302-3012330022033022-3002232021003121-0333030013132331"></a>

## Direct properties — no_mtls / 301302110010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202103131213303-1300102001220100-2300321011131022-0311033211330303-0131122100233013-2002203320312332-3130330201003333-3332330232210322"></a>

## Next pages — no_mtls / 301302110010 / 4

- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0303332132110311-1223303331210202-2030002213310231-2101021031103002-3222212230031331-2213230130330002-3303131031031310-2010310111301030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300021332103301-1331230112113311-1133323103212332-1200020330122120-0021303321101001-2110002302210330-3131211030330003-1230231211103210"></a>

## default_pool.use_tls.skip_server_verification — skip_server_verification / 130231231031 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.skip_server_verification

<a id="canonical-1212002212011012-3301133312303113-1013031220110323-3303223000111220-3012313232202033-2132132223120222-0322311011312202-0002032201302312"></a>

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
skip_server_verification = {}
```

<a id="canonical-3302322230320323-0323033330110233-1222323210333311-0012231203120021-2222120210121230-0123222033120110-1230012331110120-0232302003030100"></a>

## Direct properties — skip_server_verification / 130231231031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0300202013123332-3003002111231012-1120111001233000-3013130320112220-3313121022000300-0003023003101010-2212003131321212-2101023122320212"></a>

## Next pages — skip_server_verification / 130231231031 / 4

- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231013032332023-2323012321201032-0113233032310002-0121100300001020-2201301333031211-0032310301110012-0233111121030003-2131203211113222"></a>

## default_pool.use_tls.tls_config — tls_config / 101021302212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.tls_config

<a id="canonical-3200322132320010-0012131232103001-1322321033121031-0031033101302232-1313330021103112-3102301211222010-1001032231012121-2320123131013313"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130122231121302-1331310133031001-0222331111220202-3130213300111030-2101001321313122-1111121222013220-0121231011001103-1111221123113313"></a>

## Direct properties — tls_config / 101021302212 / 3

- [custom_security](resources--http_loadbalancer--reference--group-017.md#canonical-2200333031323122-3012011201002322-3130330223013331-1233310103131031-2030012001303132-0303212231120001-1201102011200321-1010332223231233): complete subsection reference.

- [default_security](resources--http_loadbalancer--reference--group-017.md#canonical-0100203223333331-3222113310331031-3023321131313312-3310033221020010-1313130223033303-1002200231311002-3020333203131231-3220021002103110): complete subsection reference.

- [low_security](resources--http_loadbalancer--reference--group-017.md#canonical-0303100332303222-3233213233330122-3131223311022000-1321203230112312-0031221232100122-2031312301300013-2212123210100030-0103022302311131): complete subsection reference.

- [medium_security](resources--http_loadbalancer--reference--group-017.md#canonical-0222121100303332-0023323120332332-0330001011121000-1210020031030130-2220031231123120-0120333002033001-2221322033313322-0131003300131300): complete subsection reference.

<a id="canonical-3211112112213111-3032032023203123-1320231010222200-3231222200331022-1212131023111003-0002212033211231-0122023332002320-0323212120033320"></a>

## Next pages — tls_config / 101021302212 / 4

- [default_pool.use_tls.tls_config.custom_security](resources--http_loadbalancer--reference--group-017.md#canonical-2200333031323122-3012011201002322-3130330223013331-1233310103131031-2030012001303132-0303212231120001-1201102011200321-1010332223231233)
- [default_pool.use_tls.tls_config.default_security](resources--http_loadbalancer--reference--group-017.md#canonical-0100203223333331-3222113310331031-3023321131313312-3310033221020010-1313130223033303-1002200231311002-3020333203131231-3220021002103110)
- [default_pool.use_tls.tls_config.low_security](resources--http_loadbalancer--reference--group-017.md#canonical-0303100332303222-3233213233330122-3131223311022000-1321203230112312-0031221232100122-2031312301300013-2212123210100030-0103022302311131)
- [default_pool.use_tls.tls_config.medium_security](resources--http_loadbalancer--reference--group-017.md#canonical-0222121100303332-0023323120332332-0330001011121000-1210020031030130-2220031231123120-0120333002033001-2221322033313322-0131003300131300)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2200333031323122-3012011201002322-3130330223013331-1233310103131031-2030012001303132-0303212231120001-1201102011200321-1010332223231233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102220012202123-3321223213022213-3301212212010123-0333200120320230-1033011213022010-3321101213333321-1332021300021123-1310232010223020"></a>

## default_pool.use_tls.tls_config.custom_security — custom_security / 102203312333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- default_pool.use_tls.tls_config.custom_security

<a id="canonical-1220133112231103-3133202012022230-0331123031020012-3112010330311112-0202101003110102-1213221301210002-0300333021103032-1121322030113003"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112003212202120-3021113011313011-0233233221011133-1011020231013310-1231200131020013-2220221123130033-0303332321101210-0210131310221023"></a>

## Direct properties — custom_security / 102203312333 / 3

<a id="canonical-2302131311321310-2313111213002113-0100103000220313-2320303332300231-2303022003032033-3132133303200303-2222210003100101-0321130323100122"></a>

<a id="canonical-2002032131210120-1031100130223013-3020001330000232-0222311122301100-0323212221310101-3002323113221332-3102230320032000-2210233221322203"></a>

## cipher_suites property — custom_security / 102203312333 / 4

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3021112011223320-0021132033313301-1023310221110301-3311031233211110-1301132310331000-2010122332103110-2133232211213330-1203020030100311"></a>

<a id="canonical-1012211102113221-1112112333300031-3021000311033212-3322112122320311-2301333320312223-3131200001211131-0313022302333131-0331111021200012"></a>

## max_version property — custom_security / 102203312333 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1301111312203121-3301313013110313-3101201213000210-3230133310320030-0031223213103012-1130123011223330-1010032130303022-3332012300313323"></a>

<a id="canonical-3132001120112311-1202000033020013-2303112330022223-0332011032222230-3031132011232211-3223232222123301-0101300301220222-3323121333211023"></a>

## min_version property — custom_security / 102203312333 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1301032321110203-3131330001202021-2210121212212030-3130103211121123-2322203113233121-3010001330303231-2321210130333120-2102120023123312"></a>

## Next pages — custom_security / 102203312333 / 7

- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0100203223333331-3222113310331031-3023321131313312-3310033221020010-1313130223033303-1002200231311002-3020333203131231-3220021002103110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112232113330222-1100331100313020-1102202000230311-1133301231002231-3223033102101213-0100331130221322-1120331310323310-3301012133120332"></a>

## default_pool.use_tls.tls_config.default_security — default_security / 113002220311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- default_pool.use_tls.tls_config.default_security

<a id="canonical-2001121231113323-2202302022312013-1331012232331221-2323103111030031-1130222031133131-0113313212323213-0300232212322333-3320000103001121"></a>

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
default_security = {}
```

<a id="canonical-1200111210323200-2023031231222301-2312322212310220-2321200323113133-1010310322032023-3322111300203332-0222101010002233-1110033032303101"></a>

## Direct properties — default_security / 113002220311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301103321013210-1212200203011022-1231020123220021-0123130213131211-2333123322300023-2033013201321333-2021312301231120-0310212210332320"></a>

## Next pages — default_security / 113002220311 / 4

- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0303100332303222-3233213233330122-3131223311022000-1321203230112312-0031221232100122-2031312301300013-2212123210100030-0103022302311131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033332233100200-3021333000103000-3220112333333200-3132203022130222-0033021033133310-3300002121030122-3323000332031233-2012001333023323"></a>

## default_pool.use_tls.tls_config.low_security — low_security / 033321232321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- default_pool.use_tls.tls_config.low_security

<a id="canonical-0130332222331002-0011221102013332-0301322031321302-3022120020233011-2302221311111002-3302121031031122-3230332233003002-0003201330112311"></a>

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
low_security = {}
```

<a id="canonical-2133222200110122-2032131202331100-1320032131100001-0132002222221002-3110022310022231-1033030103221113-1032210303312113-0021223310112033"></a>

## Direct properties — low_security / 033321232321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032221101103203-1331200021201223-2311022120031001-0000130331320310-2132333030213311-0120023133320121-0220313113313132-1110203132230132"></a>

## Next pages — low_security / 033321232321 / 4

- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0222121100303332-0023323120332332-0330001011121000-1210020031030130-2220031231123120-0120333002033001-2221322033313322-0131003300131300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030110011010013-1121313010300012-3121231312011030-0223300201120113-0322221110113133-1031301310331112-2312232310020222-2123001212211203"></a>

## default_pool.use_tls.tls_config.medium_security — medium_security / 231000311021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- default_pool.use_tls.tls_config.medium_security

<a id="canonical-1000302303231022-3322010013221311-0123300200102313-2202220003120212-3111311222313132-2222122030223012-1312020130221101-3112232313213203"></a>

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
medium_security = {}
```

<a id="canonical-3110022313131222-1320323131022310-0132323121221303-1030012003231322-0222110132231210-2332302011132300-3232330001332333-0012122222313231"></a>

## Direct properties — medium_security / 231000311021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331310033301021-1312100230012031-1003231301111322-3003212312222103-2031031230113232-0032223130211230-2313203231013332-3021311131313332"></a>

## Next pages — medium_security / 231000311021 / 4

- [default_pool.use_tls.tls_config](resources--http_loadbalancer--reference--group-017.md#canonical-3022001312010311-2033321103031000-0230211202123211-3032221333033333-2023331303010222-3103112113331032-3303302020113030-0113122302220112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1322020121333230-3003002212033100-2332300001021121-1102021010223121-3131123312322320-1020132011323230-0202022322031312-3313322321233013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311233101303002-1313113113032101-0230313020123230-3323231002213300-0210030132033133-3010021031231322-1323113312130011-3221233210032211"></a>

## default_pool.use_tls.use_host_header_as_sni — use_host_header_as_sni / 211033030311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.use_host_header_as_sni

<a id="canonical-0313031012231211-1002132313113203-0121133132332010-3323033303133001-2320223110111303-0231130011133120-1311031200103022-3301103201002302"></a>

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
use_host_header_as_sni = {}
```

<a id="canonical-3100213120231122-0301311330030201-2310323312111233-1303100022220230-1021100232332030-1120120022233222-0133303210232202-2220331213130331"></a>

## Direct properties — use_host_header_as_sni / 211033030311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232312321211323-1223003321312033-3123210313013230-3021133113202002-2312122230303113-0231210311021233-0313303003120020-3221311012031031"></a>

## Next pages — use_host_header_as_sni / 211033030311 / 4

- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302122203211220-2121332323211221-0102313012221101-0033133321201112-1003001100120130-1000100220210213-2313103302202011-2133103030320110"></a>

## default_pool.use_tls.use_mtls — use_mtls / 312102200302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.use_mtls

<a id="canonical-0012033332103102-2000112211213221-2202123011211022-1112211212300213-3323011221300011-1203333113303113-3102003221010000-1301333232022310"></a>

Type: `"object"`. single nested block, Optional.

MTLS Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates")}
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
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333303030012123-3330302113220213-1210023032020223-3100313001201123-1102101032030301-2231312231203230-2021302323323013-2333122223012210"></a>

## Direct properties — use_mtls / 312102200302 / 3

- [tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023): complete subsection reference.

<a id="canonical-1112232110303100-2332313320331333-1322120030202002-0211323002201021-2332322011320012-2112231302321032-0102332121121100-1311200110102030"></a>

## Next pages — use_mtls / 312102200302 / 4

- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322332012030213-2101321010323011-0323230010300033-0311312032020030-1220103001011021-3023011030213022-2013210133320030-3112131223100023"></a>

## default_pool.use_tls.use_mtls.tls_certificates — tls_certificates / 121232130122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- default_pool.use_tls.use_mtls.tls_certificates

<a id="canonical-2212133130313131-0212312110312103-1001211323113313-0122031120232010-2021312001202331-3111320320020322-2211331321031122-1112301000111222"></a>

Type: `"object"`. list nested block, Optional.

MTLS Client Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
```

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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103022201121202-1302132303010331-0023323221331221-0322223330213331-3011130023133202-0113120122130133-0233013302312332-3333021001221030"></a>

## Direct properties — tls_certificates / 121232130122 / 3

<a id="canonical-0020200123311232-3102103202333331-0002123021121203-3031120233301302-2100323232320100-2020032022301301-3100213212010200-3211120101203113"></a>

<a id="canonical-2033010230112102-3130213100112022-0220203003032213-0310023023321023-0103311120133130-2300033321031122-3303320110231211-0300102213322313"></a>

## certificate_url property — tls_certificates / 121232130122 / 4

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

- [custom_hash_algorithms](resources--http_loadbalancer--reference--group-017.md#canonical-2020213113300220-2333110233301203-3300222130030231-2130303232320030-2032111112212201-2300222323013202-0110222032000202-1320111112111031): complete subsection reference.

<a id="canonical-2221103223010020-3100323022012230-3032130320011322-2113230031211331-1222220132311223-0211100112030330-0302212101133010-0201131020300123"></a>

<a id="canonical-2223031313130023-0031110330030203-2212103202020222-1132302231030203-2221011320000233-1001131031321213-2221013213030300-0030130203332312"></a>

## description_spec property — tls_certificates / 121232130122 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--http_loadbalancer--reference--group-017.md#canonical-1322333013020002-3320101320322330-1230010223210130-0332001200121133-0020320123030210-0012032023300221-0031101222021121-2210211232313232): complete subsection reference.

- [private_key](resources--http_loadbalancer--reference--group-017.md#canonical-2131202211301100-0331121201212013-2133312320113023-3033313101313002-3322302102012100-0123222110310223-0321031132101002-3021002201003223): complete subsection reference.

- [use_system_defaults](resources--http_loadbalancer--reference--group-017.md#canonical-1112030331010001-1223303203031330-0333010110213321-3202130323302212-0212221311313102-1030203333133001-1312232202322201-0223303001130330): complete subsection reference.

<a id="canonical-2010103022013320-3130122112200022-3223201221303002-2303323111212001-0032201203212121-0233023131002300-2230112232120203-2103203221023123"></a>

## Next pages — tls_certificates / 121232130122 / 6

- [default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms](resources--http_loadbalancer--reference--group-017.md#canonical-2020213113300220-2333110233301203-3300222130030231-2130303232320030-2032111112212201-2300222323013202-0110222032000202-1320111112111031)
- [default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](resources--http_loadbalancer--reference--group-017.md#canonical-1322333013020002-3320101320322330-1230010223210130-0332001200121133-0020320123030210-0012032023300221-0031101222021121-2210211232313232)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key](resources--http_loadbalancer--reference--group-017.md#canonical-2131202211301100-0331121201212013-2133312320113023-3033313101313002-3322302102012100-0123222110310223-0321031132101002-3021002201003223)
- [default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults](resources--http_loadbalancer--reference--group-017.md#canonical-1112030331010001-1223303203031330-0333010110213321-3202130323302212-0212221311313102-1030203333133001-1312232202322201-0223303001130330)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2020213113300220-2333110233301203-3300222130030231-2130303232320030-2032111112212201-2300222323013202-0110222032000202-1320111112111031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223031031211033-0320311121020102-1310011323203221-2211230320301032-1332022133211022-3133010133211020-1321002202133110-2002301230121103"></a>

## default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 322230331103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- default_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms

<a id="canonical-1030021013012233-1132023103020003-0322031310212302-1102323132200031-1330231033232303-1101231202333203-0200121200223332-0323333130022223"></a>

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

<a id="canonical-0033213203301303-0003023003003013-0223320212202100-3232321032232212-2022121133022332-3303223322022311-3302103003110002-1011320011100023"></a>

## Direct properties — custom_hash_algorithms / 322230331103 / 3

<a id="canonical-3331002031032101-2001130333223101-0302021313121000-3033323313301110-2233313222022113-0310102131332223-3010311101032113-2133332103300131"></a>

<a id="canonical-1110221230201200-0211230013322322-3132322013331032-0210223130020031-1032201033001110-2021223020010002-3100203203232130-3033313102233033"></a>

## hash_algorithms property — custom_hash_algorithms / 322230331103 / 4

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

<a id="canonical-0212113320021220-0020230221211211-0003120020300011-2310312200000012-3232220221232003-0123223220112133-2210311132132203-3222011000133031"></a>

## Next pages — custom_hash_algorithms / 322230331103 / 5

- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1322333013020002-3320101320322330-1230010223210130-0332001200121133-0020320123030210-0012032023300221-0031101222021121-2210211232313232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002312330110032-2123321122331302-3010102032231031-0023130001122013-0032330200113213-1102203303112300-3302121222332102-3110120213023301"></a>

## default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 332330221000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- default_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="canonical-1300032021202110-0011031032313012-0300212012032212-3333302112232001-0331202031331003-3330010131203303-1201000113210300-3211001103103331"></a>

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

<a id="canonical-0002232013320123-2033210103321330-2032223031300032-2320113030123000-2113130313202132-1220323030311320-1320230213302132-1312121013332133"></a>

## Direct properties — disable_ocsp_stapling / 332330221000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022032302031301-2030313103030123-0201112332213231-2021123201331201-0133322310301033-2211210301013200-1023002130001001-2112230032322322"></a>

## Next pages — disable_ocsp_stapling / 332330221000 / 4

- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2131202211301100-0331121201212013-2133312320113023-3033313101313002-3322302102012100-0123222110310223-0321031132101002-3021002201003223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000020210331122-1132010321001303-2131303013231130-2312030102130212-2113301020003331-2331220300303232-2023111000331011-1122212223201201"></a>

## default_pool.use_tls.use_mtls.tls_certificates.private_key — private_key / 111322332331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- default_pool.use_tls.use_mtls.tls_certificates.private_key

<a id="canonical-2203210231332111-3031132312132123-1130302233133303-0110113131033132-2310332200311320-0330113002033032-3000202300132332-3111130320320313"></a>

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

<a id="canonical-3012030300111112-0203220223322022-3003201202100300-0000202112111001-2123030233112103-3213231320221022-2330101210301002-0332212331213003"></a>

## Direct properties — private_key / 111322332331 / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-3202131212220022-3310331222003333-0132103303101313-3122213000000300-2332013301033012-3213003113213002-3322023211002302-2102301302313033): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-3303001112302110-2223022023002212-0233233102211101-2011031233223201-2222213211201220-2121132300103222-3002120121011222-3323020013000102): complete subsection reference.

<a id="canonical-2021011033023003-1213101122230021-3330221103113332-3230323313020030-0202122321320202-1023221011122020-2330332123023132-1012300111032113"></a>

## Next pages — private_key / 111322332331 / 4

- [default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-3202131212220022-3310331222003333-0132103303101313-3122213000000300-2332013301033012-3213003113213002-3322023211002302-2102301302313033)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](resources--http_loadbalancer--reference--group-017.md#canonical-3303001112302110-2223022023002212-0233233102211101-2011031233223201-2222213211201220-2121132300103222-3002120121011222-3323020013000102)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3202131212220022-3310331222003333-0132103303101313-3122213000000300-2332013301033012-3213003113213002-3322023211002302-2102301302313033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022123013132032-1211200012123101-3203133100311113-1330210213232022-1131312032133021-3230003113201313-2023110023022120-1131222231323102"></a>

## default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 000332322300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key](resources--http_loadbalancer--reference--group-017.md#canonical-2131202211301100-0331121201212013-2133312320113023-3033313101313002-3322302102012100-0123222110310223-0321031132101002-3021002201003223)
- default_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0022023210023202-2110310333003320-1012230102222210-2233113231311001-0133022031032111-3201010222100011-1032010111103100-3001101003123221"></a>

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

<a id="canonical-3231112100121132-2030200201021300-2003030123133322-2003200330323323-2003312223220220-0103112113321223-2312121112000310-0333120312220303"></a>

## Direct properties — blindfold_secret_info / 000332322300 / 3

<a id="canonical-2333112213321311-2001302311302120-3231303003201322-3022311002323320-3030210223320032-1001332012312303-0112023031001221-0011200221322313"></a>

<a id="canonical-0032031131221303-0132210003012202-3333101212203313-0300302322331020-0130222203322333-0130102020222133-2303100032333132-1101022210223210"></a>

## decryption_provider property — blindfold_secret_info / 000332322300 / 4

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

<a id="canonical-2111132120203332-1311130030030123-3213013212012201-2031333230020311-3233301023312330-3113330012020303-2132001123033300-2113231333020013"></a>

<a id="canonical-2211302133122200-0133321100110020-2033313312213002-2133211011313210-0231310000023213-3102010300122303-3230301130332200-1210212330202222"></a>

## location property — blindfold_secret_info / 000332322300 / 5

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

<a id="canonical-0200303010213013-1132000102112230-2003002331021311-3213033131202221-3310031302201130-2213111100122120-0022023220201031-2111201322113232"></a>

<a id="canonical-2213131020303011-3023030030312021-1333013011323220-0232222130203232-3231332122221002-0213000222031312-1100120231312302-3220231031310232"></a>

## store_provider property — blindfold_secret_info / 000332322300 / 6

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

<a id="canonical-3232312233202233-3213302131313120-1203331321331031-0303322200033211-1132100131321131-2310010301123212-0311002120021023-0122013212112123"></a>

## Next pages — blindfold_secret_info / 000332322300 / 7

- [default_pool.use_tls.use_mtls.tls_certificates.private_key](resources--http_loadbalancer--reference--group-017.md#canonical-2131202211301100-0331121201212013-2133312320113023-3033313101313002-3322302102012100-0123222110310223-0321031132101002-3021002201003223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3303001112302110-2223022023002212-0233233102211101-2011031233223201-2222213211201220-2121132300103222-3002120121011222-3323020013000102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120230331230010-1032221231322231-2230320202013233-3123322312010201-1003002323310101-1310011332310131-1303020203311212-2003033002230001"></a>

## default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info — clear_secret_info / 132100222032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- [default_pool.use_tls.use_mtls.tls_certificates.private_key](resources--http_loadbalancer--reference--group-017.md#canonical-2131202211301100-0331121201212013-2133312320113023-3033313101313002-3322302102012100-0123222110310223-0321031132101002-3021002201003223)
- default_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

<a id="canonical-3300000121331033-2030010003221332-0033231020032330-2023223021123323-1001021121030230-1033001102331233-2321102003300023-0012201032120330"></a>

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

<a id="canonical-2032120031030031-3103133010001123-2131200301121033-1203103231032202-2003133303003033-0232030202032332-3332331321300132-3022221221202222"></a>

## Direct properties — clear_secret_info / 132100222032 / 3

<a id="canonical-0200323133202323-3011021311113011-1303110212331210-1112101223331200-2031320312033333-0221222200323121-1013130021123101-0100323201211333"></a>

<a id="canonical-0330302210130210-3333010122030320-1133013222230302-1303223221223232-1331223322133221-0130011033031331-1332102201113010-2131321113011031"></a>

## provider_ref property — clear_secret_info / 132100222032 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3202120300223332-3022201000103233-1010230210303111-2022212302032133-3133120013120033-1032031221023220-1302011331020303-1232333130333211"></a>

<a id="canonical-3201031021323233-3322103222131133-1200233103320031-2000333031221211-2111012211203331-0213212020010033-3002001203112221-0133013122302222"></a>

## URL property — clear_secret_info / 132100222032 / 5

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

<a id="canonical-1203123212123122-0030233230333320-2101201031132101-3312020021103302-0202120012100213-1301311003031333-0023020003311101-0220033011300132"></a>

## Next pages — clear_secret_info / 132100222032 / 6

- [default_pool.use_tls.use_mtls.tls_certificates.private_key](resources--http_loadbalancer--reference--group-017.md#canonical-2131202211301100-0331121201212013-2133312320113023-3033313101313002-3322302102012100-0123222110310223-0321031132101002-3021002201003223)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1112030331010001-1223303203031330-0333010110213321-3202130323302212-0212221311313102-1030203333133001-1312232202322201-0223303001130330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120031011010203-3212201101311130-1122310220211232-2020111122323330-1000331333003311-0001310002010220-3332323012311323-2022311310200313"></a>

## default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults — use_system_defaults / 132113111100 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_mtls](resources--http_loadbalancer--reference--group-017.md#canonical-3310213332012300-2321311313133101-3233031103331120-0031100323302331-0112100121100221-0230032322001311-1212023322222130-3013331133011031)
- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- default_pool.use_tls.use_mtls.tls_certificates.use_system_defaults

<a id="canonical-0133000112321230-1110222120112232-0333021301203130-3323020310312330-0023121000131003-1102101110023131-3330122333222230-0223021310123001"></a>

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

<a id="canonical-3321333333102221-1310223200310132-3232021220112120-0001011010100321-2022122021133130-2031033320102013-0123111030010100-0333333321122312"></a>

## Direct properties — use_system_defaults / 132113111100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012223213310322-0232100333231100-0100301231011202-0321110023133222-2311310110120312-3232223311202313-2003000031213323-2220030031121213"></a>

## Next pages — use_system_defaults / 132113111100 / 4

- [default_pool.use_tls.use_mtls.tls_certificates](resources--http_loadbalancer--reference--group-017.md#canonical-2221123201123120-2301211320101011-1131200100122133-3112200312220332-3213103203332321-1022000031031022-3201330020110311-1121022130030023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2011310303210021-1222202221103110-1130033232102230-1131030100322101-0032220031011132-0202012301223101-3013101110032002-0102022212323012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231122113302332-2002303233330021-0332111332303312-3023103303132330-2332103233301201-3220121103302022-1120330231012032-0320120231010122"></a>

## default_pool.use_tls.use_mtls_obj — use_mtls_obj / 331311223232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.use_mtls_obj

<a id="canonical-0022112112130111-3011023102110131-0333031030310303-2030022132112002-3133330123223031-1132202201000033-0130031010113333-0033302321210101"></a>

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
use_mtls_obj {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302202001310110-1101211010303021-1032211223322000-1232121312203003-3332330010311011-0321132021331220-0003133013313231-2013033123223330"></a>

## Direct properties — use_mtls_obj / 331311223232 / 3

<a id="canonical-0121211103102032-0330313001322120-1112233233332320-0130021120223001-0203302301023132-1303301312221132-1221121130032311-2110223112313130"></a>

<a id="canonical-3102322103321232-0122203111103223-2102232232331100-2032233303011301-2022033303310230-3002023012311031-3011112002202211-1301201333322133"></a>

## name property — use_mtls_obj / 331311223232 / 4

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

<a id="canonical-0123332020332101-0232133120232030-2000000233320132-1320110222321033-0220020120231321-0321213213001001-3211301002323303-1330020100320110"></a>

<a id="canonical-1100311230011010-3020301300323222-0003012022321222-2131002031130032-0030223302012113-3030330031021100-0211323112213033-2223201132010021"></a>

## namespace property — use_mtls_obj / 331311223232 / 5

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

<a id="canonical-2020133233123002-2322012220031112-1200111211102032-2113130302302200-2010102122302000-0310102312220002-2030202200222210-3221003330003323"></a>

<a id="canonical-1013130110131111-0211212212331133-2111123302213010-3302032023200113-0123310320132033-3202232312333010-1223300323311330-0212022102003313"></a>

## tenant property — use_mtls_obj / 331311223232 / 6

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

<a id="canonical-1233312201101120-2031320110333102-3101213201132011-3301130111210122-1120302230333102-0133213233330330-3331103332103023-1312002030310002"></a>

## Next pages — use_mtls_obj / 331311223232 / 7

- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3302122303122133-0323233203310323-3012112030002300-2003211111010032-2031330012231010-1011303111010020-0113222311123003-1002223100220030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332002310031003-3001331101220003-3212223023322323-0121333111101322-0000211211232121-2311233003220121-1032333303122010-2031313303131031"></a>

## default_pool.use_tls.use_server_verification — use_server_verification / 111330102123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.use_server_verification

<a id="canonical-2130103020232302-1031121000213331-1200111111332320-1123033212132121-1232002101101333-3211110303301331-3210133303122301-3102331023033332"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for use server verification.

Upstream description:

Upstream TLS Validation Context.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
use_server_verification {
  # Configure direct properties listed below.
}
```

<a id="canonical-3321330102033021-2122301031113312-0311012031222230-2121230320311030-1222222133230230-1300113301122133-0311122210322212-2203322332333311"></a>

## Direct properties — use_server_verification / 111330102123 / 3

- [trusted_ca](resources--http_loadbalancer--reference--group-017.md#canonical-3231101200310233-1023123032133030-2222223213120303-3302312032100301-0101231330003333-1230030230003100-2231003210132301-2012033101130000): complete subsection reference.

<a id="canonical-3302121201002223-1132313201012120-1122021012131202-2321102221122231-2311231033221310-1033030303311121-2313113220020103-2230223112333331"></a>

<a id="canonical-2231033101112222-2322222301122210-0003001203213231-3322131031100322-0232012232131013-2111101212011100-1202203210133311-3211322303001321"></a>

## trusted_ca_url property — use_server_verification / 111330102123 / 4

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-1010123221233022-2012301120302201-1122130021121123-1201313203112310-0332112033011032-0230022030320230-3130300130110232-1200012311323103"></a>

## Next pages — use_server_verification / 111330102123 / 5

- [default_pool.use_tls.use_server_verification.trusted_ca](resources--http_loadbalancer--reference--group-017.md#canonical-3231101200310233-1023123032133030-2222223213120303-3302312032100301-0101231330003333-1230030230003100-2231003210132301-2012033101130000)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3231101200310233-1023123032133030-2222223213120303-3302312032100301-0101231330003333-1230030230003100-2231003210132301-2012033101130000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231001132223201-3000000230221201-3230301222103130-0111320221302321-0030221130133001-3003331233013102-2011232310231012-3302310102230223"></a>

## default_pool.use_tls.use_server_verification.trusted_ca — trusted_ca / 311213002203 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [default_pool.use_tls.use_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-3302122303122133-0323233203310323-3012112030002300-2003211111010032-2031330012231010-1011303111010020-0113222311123003-1002223100220030)
- default_pool.use_tls.use_server_verification.trusted_ca

<a id="canonical-1231200013223113-3313130001203221-0300003313303101-2321020102203333-3002301211023301-0111230320300103-1111010320123121-3333222302202002"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222321332201113-1102203230332110-0233103200220002-0001321230302221-1032233212332323-3021332200322330-1000210222110302-3200312033033032"></a>

## Direct properties — trusted_ca / 311213002203 / 3

<a id="canonical-3213210320020231-0003132323301200-3232213232232323-3302123312201310-2032212230132002-1333320233001021-1313033203123313-3203312333330102"></a>

<a id="canonical-0133100203011322-1032013323121132-2012222331222131-1332113030212203-1100112231213000-3020003322203031-2223200131311022-3122323211010102"></a>

## name property — trusted_ca / 311213002203 / 4

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

<a id="canonical-3202102011303222-0022021021120200-3313011233333120-0022133020230332-3023001112110013-3123131200213303-0211321023130233-1331322321231010"></a>

<a id="canonical-0032003232312033-1231311332131332-2023210223330033-0212131201011331-1002133120222033-0012320332012010-2031312000100012-3032221202330213"></a>

## namespace property — trusted_ca / 311213002203 / 5

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

<a id="canonical-1211213032201123-1101110211320311-3002013002111020-2023201011122120-3122022312332313-0231013113112212-0332301133201321-2203322031021012"></a>

<a id="canonical-2202012323210122-0233231110232332-3123002321303330-1210030132102313-0000130211332220-2132333332211220-1031312310213303-1021232332023302"></a>

## tenant property — trusted_ca / 311213002203 / 6

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

<a id="canonical-0130122231300031-0012320322302213-1332002322123010-1310002123302310-3311130110332122-3320303302331130-0100120110330331-0011212021131101"></a>

## Next pages — trusted_ca / 311213002203 / 7

- [default_pool.use_tls.use_server_verification](resources--http_loadbalancer--reference--group-017.md#canonical-3302122303122133-0323233203310323-3012112030002300-2003211111010032-2031330012231010-1011303111010020-0113222311123003-1002223100220030)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3321010323221300-3221232231012333-1000122333111322-3013010233100302-2311021323120212-1121301113331022-3222031002320022-3101011201003023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330123212133231-3320212210020212-2010101013332300-1001111301312133-2220103132332210-3203300312331001-3122230303220212-2010200300232333"></a>

## default_pool.use_tls.volterra_trusted_ca — volterra_trusted_ca / 221333113310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- default_pool.use_tls.volterra_trusted_ca

<a id="canonical-1123131102231031-3303002311113322-2223231122012003-1221030332331211-3223111220030330-1032310301010110-3231233101113000-2110012102333313"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for volterra trusted ca. Defaults to \`map\[\]\`. Server applies default
when omitted.

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

<a id="canonical-2322320203313011-1011122130301011-3013320103123232-0322103021132021-3123223131030102-1023123212132133-1123333101022223-1321130332131103"></a>

## Direct properties — volterra_trusted_ca / 221333113310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010003013322003-3231000201001212-2010000123201211-2030033133200121-0221020110201131-3102010312223331-0310113131210323-1121002333320231"></a>

## Next pages — volterra_trusted_ca / 221333113310 / 4

- [default_pool.use_tls](resources--http_loadbalancer--reference--group-017.md#canonical-1220102102233330-1130323131232203-1010221220003123-0231201012101300-2222312303123110-1100032001232130-3230220301231310-1122123221111022)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0210031322021130-2010122303312010-3232211331331322-0132231002001122-1111001310002033-1200111313232223-3023311310020010-2112111310030332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332302222233002-1022201332120301-1130021200330033-2321122302112121-2312313112303302-2001301003102032-1032000212211130-3031303000121202"></a>

## default_pool.view_internal — view_internal / 000323221110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- default_pool.view_internal

<a id="canonical-3111311300001101-1221223112021201-0221010001311121-2322133311013113-0202010001021132-0303200203011113-2102003220222010-2213223113000030"></a>

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
view_internal {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321012230120312-3030331101102322-0213222012012323-3121111131332033-0211023011120020-1101001302000322-3030132201202212-0132023333312133"></a>

## Direct properties — view_internal / 000323221110 / 3

<a id="canonical-2330201021132320-2202103302033221-1223110301030212-1210130012230213-3210132102120002-1001323103002030-2220033103133223-2131111310322330"></a>

<a id="canonical-2133223222112030-1320112311020112-1322022131230120-1221330231310223-1330010330210302-1133120002210301-3120230310123131-2110031003121202"></a>

## name property — view_internal / 000323221110 / 4

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

<a id="canonical-1021013023031201-3020331300312210-1023213331112123-0032310012033303-0212302203133210-0123200110130131-2001122112230101-1023220200023222"></a>

<a id="canonical-3113222023220233-0103011332103031-3233312010120300-1130120100103220-1332231013133210-0303222133001003-1023001103312101-2133321121100120"></a>

## namespace property — view_internal / 000323221110 / 5

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

<a id="canonical-3300212223020031-1202212011022000-2031210322333130-2033121132021000-3230132232123100-2211220301323322-0321113101221012-1130213312120220"></a>

<a id="canonical-1122003313031213-1003331003110111-1120130333331130-3320110002331310-2231321023000023-1133130202201120-0032203220031121-2301313323120303"></a>

## tenant property — view_internal / 000323221110 / 6

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

<a id="canonical-0100100312320103-2111300111221033-1120230002330010-0103213320003110-1013132032223010-2031123123230230-3203013313223130-2021331012032332"></a>

## Next pages — view_internal / 000323221110 / 7

- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333223003212333-2032110013333332-0132202330213022-0021102200222000-1122122323130323-2112303013332332-0001213202032131-0210212101121222"></a>

## default_pool_list — default_pool_list / 332311023112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- default_pool_list

<a id="canonical-0130032320003112-3013231322132011-2010113111012220-3312013033110232-0201300301333010-2200230132012313-2122033102033333-3101112132302330"></a>

Type: `"object"`. single nested block, Optional.

Origin Pool List Type. List of Origin Pools.

Upstream description:

List of Origin Pools.

Receipt-pinned upstream constraints:

```json
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
default_pool_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203331021311323-2012330000303112-2103102131301223-1032001313011221-1010233320021101-0223123132333020-3231110222003101-1310201222111331"></a>

## Direct properties — default_pool_list / 332311023112 / 3

- [pools](resources--http_loadbalancer--reference--group-017.md#canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112): complete subsection reference.

<a id="canonical-3010221321311110-3012311202011320-2122212211232331-0013121030131001-3232033033100013-2020210121122031-3020032230113032-3010100230013332"></a>

## Next pages — default_pool_list / 332311023112 / 4

- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320223320331321-3221011123300323-3302000000033320-3113130111031300-0323033222220323-1321310012301331-3231033022120221-3220032030010132"></a>

## default_pool_list.pools — pools / 220231103333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311)
- default_pool_list.pools

<a id="canonical-1130133122303033-3113103332212302-2223111002322122-3312003023012013-1321033110333130-2213233123210322-2330020110210033-0103212032210123"></a>

Type: `"object"`. list nested block, Optional.

Origin Pools. List of Origin Pools.

Upstream description:

List of Origin Pools.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cluster",
    "pool")}
```

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-0113203302233112-1001122031303022-2313321023320320-1032313021123112-2012022312020033-1310112020012032-3232321102213121-0001122123222010"></a>

## Direct properties — pools / 220231103333 / 3

- [cluster](resources--http_loadbalancer--reference--group-017.md#canonical-0203211033012202-0330031110233023-3003123220203232-0332013233023211-3220312232333201-2131113120302313-1301222332311031-0011201322300131): complete subsection reference.

- [endpoint_subsets](resources--http_loadbalancer--reference--group-017.md#canonical-3210332310321130-0113222213001001-2113013221322330-3212101123033102-3001002230002303-2132202333322300-1201201130002212-1120123101133323): complete subsection reference.

- [pool](resources--http_loadbalancer--reference--group-017.md#canonical-3013121320100123-0230112023312222-2313223300220001-0112310021033000-1001121323133302-2112022210113221-1313103310012302-3212310030321130): complete subsection reference.

<a id="canonical-3220013120332110-2330332100201203-3201012203133303-2311130120312132-1233130333001130-1120103003303110-1120202112311201-0132333330222330"></a>

<a id="canonical-3300133133230030-1101000330222113-2132032031211202-3110133033312000-2120232030002232-0201312300133321-2303201000032001-1231230012322112"></a>

## priority property — pools / 220231103333 / 4

Type: `"number"`. Optional.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the..

Upstream description:

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool Priority of 1 means highest priority and is considered active. When
active origin pool is not available, lower priority origin pools are made active as per the
increasing priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-0213013110313133-0003022203030213-3033000010320023-2113231020313113-0020330133323230-3203123222100322-3031010223100220-0020323202001132"></a>

<a id="canonical-0002221220123210-3103302211222312-2313323220222322-0000110300113110-3130103123302203-3122230312223110-1131002122120011-3022322130010023"></a>

## weight property — pools / 220231103333 / 5

Type: `"number"`. Optional.

Weight of this origin pool, valid only with multiple origin pool. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
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
  }
}
```

<a id="canonical-0113013011011013-3123321022223303-3233131202101103-0323202021200211-0202230330131103-3301213030100313-2201321102000132-1231030230331023"></a>

## Next pages — pools / 220231103333 / 6

- [default_pool_list.pools.cluster](resources--http_loadbalancer--reference--group-017.md#canonical-0203211033012202-0330031110233023-3003123220203232-0332013233023211-3220312232333201-2131113120302313-1301222332311031-0011201322300131)
- [default_pool_list.pools.endpoint_subsets](resources--http_loadbalancer--reference--group-017.md#canonical-3210332310321130-0113222213001001-2113013221322330-3212101123033102-3001002230002303-2132202333322300-1201201130002212-1120123101133323)
- [default_pool_list.pools.pool](resources--http_loadbalancer--reference--group-017.md#canonical-3013121320100123-0230112023312222-2313223300220001-0112310021033000-1001121323133302-2112022210113221-1313103310012302-3212310030321130)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0203211033012202-0330031110233023-3003123220203232-0332013233023211-3220312232333201-2131113120302313-1301222332311031-0011201322300131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032300223120013-0121112210213021-1031132233222331-2200211321102301-0102313223032300-2022010200130031-2322220132210302-0102100311111312"></a>

## default_pool_list.pools.cluster — cluster / 002113121112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311)
- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112)
- default_pool_list.pools.cluster

<a id="canonical-2022032001300210-0100113302101222-3222032123312312-2103201330210120-2101032102101110-2020300132231333-1131210122301102-1320322013203231"></a>

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
cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-0010331000011002-2223303223033332-1233332311032113-2323323113320301-1303111002111311-2301302300012322-2221203212303231-2000331110020212"></a>

## Direct properties — cluster / 002113121112 / 3

<a id="canonical-0311122122232003-0001330212231032-0323030213332321-0012221211302122-2003033001030210-3023310203222202-1121310000101322-0002111102232213"></a>

<a id="canonical-1310302033320232-0021130020221332-2222331200333301-0323033233332122-2100000030123033-1012012110201121-1103233123033122-3320211021012331"></a>

## name property — cluster / 002113121112 / 4

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

<a id="canonical-0031320031332122-3030010133313031-0033331201001223-2330020223303233-3211232012020303-1100302133331102-2221200102013223-1333202223010221"></a>

<a id="canonical-3101020203111321-1000313023232121-3320323132221110-0221023100331223-2030013310132231-0023330323033332-3312233021221103-2022303030101231"></a>

## namespace property — cluster / 002113121112 / 5

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

<a id="canonical-2203132230103200-3030001133021013-0313332320330322-3221131102121132-2203131102233131-3211013012332223-2010111123133012-0311030233332021"></a>

<a id="canonical-2330121133222223-1122330232122302-3130213320311233-3000033201131012-0203201303313032-3000113202000333-1302212130323301-1331321231321212"></a>

## tenant property — cluster / 002113121112 / 6

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

<a id="canonical-3313120232322001-0202110102232201-0010121121213130-0233031102032031-3100000221033330-2033231211330021-1113011230230223-0230113032031231"></a>

## Next pages — cluster / 002113121112 / 7

- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3210332310321130-0113222213001001-2113013221322330-3212101123033102-3001002230002303-2132202333322300-1201201130002212-1120123101133323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121230003011013-2302001301223202-2101200023023322-1330120021200132-0121120210021113-0011301312203033-3321100322313003-1312133031120231"></a>

## default_pool_list.pools.endpoint_subsets — endpoint_subsets / 310102320020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311)
- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112)
- default_pool_list.pools.endpoint_subsets

<a id="canonical-1323010032123123-0013103233303002-1223231222330131-1111223120220111-2332113311112011-3202211000212220-3321213312023021-0302111200223122"></a>

Type: `"object"`. single nested block, Optional.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer For origin servers which are discovered in K8s or Consul..

Upstream description:

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

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
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

Terraform syntax:

```terraform
endpoint_subsets {}
```

<a id="canonical-3323220301133233-0101020302010321-1212001132300233-0131321013020010-0120030302332032-0211120232010203-2121323101320010-3122011233323213"></a>

## Direct properties — endpoint_subsets / 310102320020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200221001011003-1200030332122031-0030110011232211-1300112112332303-2033323311313323-2232011201030003-2033002011033003-0111202233221103"></a>

## Next pages — endpoint_subsets / 310102320020 / 4

- [default_pool_list.pools](resources--http_loadbalancer--reference--group-017.md#canonical-3001002133320230-2120231222300222-0321102003221001-3110121302022323-0220111101122202-3330111333232120-3211021111120102-3212311302101112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3013121320100123-0230112023312222-2313223300220001-0112310021033000-1001121323133302-2112022210113221-1313103310012302-3212310030321130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
