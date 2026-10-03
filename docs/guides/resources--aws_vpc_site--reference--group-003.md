---
page_title: "xcsh_aws_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site reference."
---

# xcsh_aws_vpc_site reference

<a id="canonical-1200202221222111-0231121031022132-0000122000031231-3011200311331201-2213001031220020-3133203222002101-3212232031322331-2120330003021130"></a>

## name property — global_vn / 321213021022 / 4

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

<a id="canonical-2112103002012102-1131101103202123-3101103131102101-3213323223130321-2313100032312030-2030132113112331-2200100010210110-3202203230031302"></a>

<a id="canonical-3211301123013230-1203100031000122-0010122022200233-1330333120230032-1003331230202031-1023231102032131-1033223103313012-0011222111120033"></a>

## namespace property — global_vn / 321213021022 / 5

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

<a id="canonical-1113222021111113-3102303102131233-0200300331120030-1113030321333030-1303031101130202-0120301021322020-1331320020320230-1030100330211311"></a>

<a id="canonical-3010313113132302-3212112121311212-2133332101310002-1212301030200303-1212131023010110-3230302122211212-0322301003131213-0312022100223122"></a>

## tenant property — global_vn / 321213021022 / 6

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

<a id="canonical-3013112312221313-3033033301313310-3233111030102003-2202322003222310-3310200333301121-1020330201210230-2213001331330032-3222232010210111"></a>

## Next pages — global_vn / 321213021022 / 7

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--reference--group-002.md#canonical-2210203120222300-0302312202231103-0301231011130312-3010133212000303-0030212033332212-3133000122020130-2302222112203213-3111200230232232)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3333112200022311-2133223212010101-3230211221110023-3121110323000232-2133113222312032-0100112123230121-0030203222120223-0322323111333131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123012011321320-3100032012010232-0320022111130202-2320000101200210-1201203110130310-1231311210002320-0020330001333302-0232103010033202"></a>

## ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr — slo_to_global_dr / 330301203122 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-0110112001213332-2233201102333212-1321232102201102-3213031203300330-1210310211301200-3133010302300211-1313222203301301-0033103200013221)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-3333302222033232-1233003332200333-3010013003100320-2301133331212231-2202321233313221-1031311123033011-1031232031331103-3022132100300121)
- ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-1201302112113130-3321213332021112-3031103301023000-1023320011311021-2230331230203120-3211322130132001-2131101312002300-3113001033110121"></a>

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

<a id="canonical-3313132300033222-1202312133123313-3001321220102132-0313121201212220-0021101133313030-0331100123131232-3331220330231322-2322333311020232"></a>

## Direct properties — slo_to_global_dr / 330301203122 / 3

- [global_vn](resources--aws_vpc_site--reference--group-003.md#canonical-3122332101100020-0213310100310111-3213313231301122-3131221313111020-0133212322101312-2332202230332300-1312231011003131-0331030111023031): complete subsection reference.

<a id="canonical-0300022131020003-0223223032312032-0313131301321110-2130020002321103-1221032333201003-3021312103121120-2100300120301323-3310333333021332"></a>

## Next pages — slo_to_global_dr / 330301203122 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--aws_vpc_site--reference--group-003.md#canonical-3122332101100020-0213310100310111-3213313231301122-3131221313111020-0133212322101312-2332202230332300-1312231011003131-0331030111023031)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-3333302222033232-1233003332200333-3010013003100320-2301133331212231-2202321233313221-1031311123033011-1031232031331103-3022132100300121)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3122332101100020-0213310100310111-3213313231301122-3131221313111020-0133212322101312-2332202230332300-1312231011003131-0331030111023031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130022012011030-1123203321000303-1121303133221302-2201021223123023-0122110323323130-1113112230230102-2300003012100022-3011022103330221"></a>

## ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn — global_vn / 133323211201 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-0110112001213332-2233201102333212-1321232102201102-3213031203300330-1210310211301200-3133010302300211-1313222203301301-0033103200013221)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-3333302222033232-1233003332200333-3010013003100320-2301133331212231-2202321233313221-1031311123033011-1031232031331103-3022132100300121)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--reference--group-003.md#canonical-3333112200022311-2133223212010101-3230211221110023-3121110323000232-2133113222312032-0100112123230121-0030203222120223-0322323111333131)
- ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-2013321103011120-0232123132221100-1031303132213121-0201221230302330-1220012010030213-2313222101213222-0022010113110103-2203301210122010"></a>

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

<a id="canonical-1102121310032112-2321302203213330-2213030332212213-2230323213212230-2122321130100200-2231213111122023-1323030121001323-1000033130103300"></a>

## Direct properties — global_vn / 133323211201 / 3

<a id="canonical-0210120203320230-0020321300330230-1130300331033100-0001211131201301-1200122231320012-1313302031102031-3000203331210132-3021101230103011"></a>

<a id="canonical-1333022322000223-2022101313130302-0111310000303211-2003200303120310-0201203203230331-0120121102332320-0332001302023330-0302000022013302"></a>

## name property — global_vn / 133323211201 / 4

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

<a id="canonical-3222202120220213-1221312102222323-3000201222323101-0022221000023001-3210300002003233-2020200020112231-3312313222200133-3010312321001001"></a>

<a id="canonical-3131322200213313-1213130131333001-3033030200013123-0232322211011222-0200203231130012-3332322122030302-0333013223313132-0303200331302123"></a>

## namespace property — global_vn / 133323211201 / 5

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

<a id="canonical-1011110031000132-1201333210033132-1103332233301032-0020232033232313-0221323212331132-3102133323230032-2103310323203233-0103310102202012"></a>

<a id="canonical-2213303231220303-2001233000102032-1022220133330331-1113001122020331-2022210000120303-1220033131323120-2123032321211310-1110110333132013"></a>

## tenant property — global_vn / 133323211201 / 6

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

<a id="canonical-0333122220121230-3330211330223301-3101233233233200-1032322213101213-0303330321331100-3131122211122133-0120113031112101-0031111021220320"></a>

## Next pages — global_vn / 133323211201 / 7

- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--reference--group-003.md#canonical-3333112200022311-2133223212010101-3230211221110023-3121110323000232-2133113222312032-0100112123230121-0030203222120223-0322323111333131)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320103303220030-1111210001021022-1330301103213220-1303310213003101-3300020012330210-2120320213100220-2302222300302233-2020122132310320"></a>

## ingress_egress_gw.inside_static_routes — inside_static_routes / 023022031031 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.inside_static_routes

<a id="canonical-2211120012230130-2021312110103100-1232302222233020-3310010021022220-0101012133111113-1033020011230112-1130311223020211-3002320203331133"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside static routes.

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
inside_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001232202212220-3323130121301230-0201320123121312-1310222101202121-2012220220302202-0103003302320110-0033313022000311-0300121221101220"></a>

## Direct properties — inside_static_routes / 023022031031 / 3

- [static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331): complete subsection reference.

<a id="canonical-3321120331132310-2313211300222312-1131332121230011-0003111203132102-0012110303211122-0311023010201230-2023310213331213-0210002110022102"></a>

## Next pages — inside_static_routes / 023022031031 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030201110000113-0232112332330110-3132310123002030-2000230012112302-3203101331122012-2212100332022101-0302320122010202-2120012102131231"></a>

## ingress_egress_gw.inside_static_routes.static_route_list — static_route_list / 203021013131 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032)
- ingress_egress_gw.inside_static_routes.static_route_list

<a id="canonical-3011003101000011-0121321013230302-0220200302231211-1032221011110130-0111201322300111-3023130020130203-3210301020012221-2201011022210032"></a>

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

<a id="canonical-2013321203313133-3002022022230322-1131123100212200-3001030133122011-2320302111313300-0021133223330112-2020123212112122-1002300313030103"></a>

## Direct properties — static_route_list / 203021013131 / 3

- [custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002): complete subsection reference.

<a id="canonical-2323201333001302-2212031010110322-0002013222301323-2313201331300003-0322312330032200-2010133032010031-0003013321123030-1132121110122133"></a>

<a id="canonical-0323120123001322-3302231323301323-1321222012123010-3320333030320113-2002000222011113-3113231123332202-1201131002213020-2123310323310033"></a>

## simple_static_route property — static_route_list / 203021013131 / 4

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

<a id="canonical-2233111030013330-0102013332202230-1202132202111312-0133231312312100-3233122230233003-3022322013030030-3110220033321020-2033132100213001"></a>

## Next pages — static_route_list / 203021013131 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213113133333110-3102123310020032-3101121301233103-1033130021213001-3132133120210033-1303222132113201-0131223022232122-1003301002222132"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route — custom_static_route / 032231131130 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-1230301210303223-3233223003233321-2230133222323333-0222322010230102-0323102110130210-2002301020021330-1022220323201032-0033211330233231"></a>

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

<a id="canonical-2100210331023010-3000000130110232-3222022222212032-0302200020200130-3313323201121020-3323120130313312-1130311301111002-1003203101112032"></a>

## Direct properties — custom_static_route / 032231131130 / 3

<a id="canonical-1301322101003303-2013220201332201-3032212202300123-3130323302013312-3013130233121010-3333310121032030-1323122230232001-2133101200121332"></a>

<a id="canonical-1212121100320112-3231231231003023-3133123321122123-3330023001201201-3033020033321302-2033100030001330-2103232331210113-2011300002120302"></a>

## attrs property — custom_static_route / 032231131130 / 4

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](resources--aws_vpc_site--reference--group-003.md#canonical-3312023300112132-2032132330333010-3122210023332212-3321030131022003-2203310302011201-1003320302211323-2001201033133123-1231322112200113): complete subsection reference.

- [nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2010310223032010-1203110102332232-0210302230012312-0303123300211302-3110210111202302-1033132101032012-1200312331310222-0222123212211033): complete subsection reference.

- [subnets](resources--aws_vpc_site--reference--group-003.md#canonical-1203330202311220-2211112031100000-3032121030223223-0101223300032312-1201112223301330-0130230100011130-1113212311120301-3313331032231122): complete subsection reference.

<a id="canonical-2330213231322123-0003120112233001-0331332303010100-0321100120322330-1010212011213032-3131211112121222-0202322332233113-2111110311020133"></a>

## Next pages — custom_static_route / 032231131130 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](resources--aws_vpc_site--reference--group-003.md#canonical-3312023300112132-2032132330333010-3122210023332212-3321030131022003-2203310302011201-1003320302211323-2001201033133123-1231322112200113)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2010310223032010-1203110102332232-0210302230012312-0303123300211302-3110210111202302-1033132101032012-1200312331310222-0222123212211033)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-1203330202311220-2211112031100000-3032121030223223-0101223300032312-1201112223301330-0130230100011130-1113212311120301-3313331032231122)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3312023300112132-2032132330333010-3122210023332212-3321030131022003-2203310302011201-1003320302211323-2001201033133123-1231322112200113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131321331131001-2201313031220132-0123102213120223-0200310313221230-3133230311003123-0300121200023222-0213213121100330-2313303012331203"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels — labels / 122023202223 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-2121223100111322-0013010220233301-2221023003001313-2123330002310132-3202121312321333-3313120311122112-0102320303003110-0032331010003323"></a>

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

<a id="canonical-0101011213233110-0212123032102231-0003022001023120-3020302033001302-1222213122023022-1002131132110000-1231103221312103-0210310020300201"></a>

## Direct properties — labels / 122023202223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213133221303213-3312210302022012-3010332003211203-2201021313100321-1130103000233200-2203120010211323-1322210023200010-0221230003221330"></a>

## Next pages — labels / 122023202223 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2010310223032010-1203110102332232-0210302230012312-0303123300211302-3110210111202302-1033132101032012-1200312331310222-0222123212211033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210112211301222-3132210301020232-3013112311102101-1302000031000011-1233303101301013-3220012001302103-3031120120013202-3230003122300022"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 300301133010 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-0120022331002332-3200102312222132-1301301323130221-1011200020103211-2122120121022333-2111121220123113-2111131233311030-2332011101211201"></a>

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

<a id="canonical-2333122201022023-1111202033113301-1223003210222033-1320332230120220-1103313233202112-1232013123121333-1033010202121213-1331112203223022"></a>

## Direct properties — nexthop / 300301133010 / 3

- [interface](resources--aws_vpc_site--reference--group-003.md#canonical-3313131120103000-0103333033313103-2131203102202111-0002222200122020-3221010031233033-2111223013221222-0230203001321002-0100331113031023): complete subsection reference.

- [nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-1111210130332211-0133323233110021-3201033120130102-0031320231020232-0132110031133333-0301312022302130-2020001130131033-1312012120211312): complete subsection reference.

<a id="canonical-1323111222111230-2102202323021313-1101132131232320-2301223022213223-0122012313030132-2102110233322303-0013111210113332-0310212303000131"></a>

<a id="canonical-2301112223222011-1013022302210321-3003303213022231-3321203120103001-1201302002102002-1121220110100013-2220330113232013-2032121021213310"></a>

## type property — nexthop / 300301133010 / 4

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

<a id="canonical-3122100100201031-2321210310321321-3322113201210201-1022010031130123-1312121100323312-2213122003001130-1111031200313223-0200322123213000"></a>

## Next pages — nexthop / 300301133010 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_vpc_site--reference--group-003.md#canonical-3313131120103000-0103333033313103-2131203102202111-0002222200122020-3221010031233033-2111223013221222-0230203001321002-0100331113031023)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-1111210130332211-0133323233110021-3201033120130102-0031320231020232-0132110031133333-0301312022302130-2020001130131033-1312012120211312)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3313131120103000-0103333033313103-2131203102202111-0002222200122020-3221010031233033-2111223013221222-0230203001321002-0100331113031023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211131222323003-3112211013203332-0131320003003133-0011310001122132-2312112223213020-1322210323023331-1101022021133313-1112113012303202"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 133320300203 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2010310223032010-1203110102332232-0210302230012312-0303123300211302-3110210111202302-1033132101032012-1200312331310222-0222123212211033)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-3232102000110311-0122111333100300-2230311101320131-0120000121202121-1300213330311310-2010032322201010-1231313021310212-0031323112021333"></a>

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

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220301032300333-2031332001211201-1202103331220320-2221132133123033-1101212002330120-2212212213213303-0222220333330001-0320001031121022"></a>

## Direct properties — interface / 133320300203 / 3

<a id="canonical-1213332030311030-2111212221332031-2200202003233100-1131121100213120-0100010120113312-1022132131212221-0130023220330200-2011133233210100"></a>

<a id="canonical-2111322022112000-2311330222232003-3201003321231110-0231322312001223-2303023130330012-1021311210231011-2202130300232101-1222201202232002"></a>

## kind property — interface / 133320300203 / 4

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

<a id="canonical-2331001321100220-3033030020113302-1123320010113300-1032330313120123-1101020200302112-2132003121133113-2011031310110211-2220331311320013"></a>

<a id="canonical-1233332311301123-0312100330310303-1330211003320300-2130333301201302-2033322302110000-1310001230322111-1101311320001102-0131100032132331"></a>

## name property — interface / 133320300203 / 5

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

<a id="canonical-2330313321301323-3023313310223013-3201223021300330-0100210330221310-0323030133302132-3211111020121001-2012210021300031-3232022122202113"></a>

<a id="canonical-0230022230103312-1302032130030031-2321300203310301-1022212113203210-1013100030030220-1330203111301113-3221122031303002-3311002013221232"></a>

## namespace property — interface / 133320300203 / 6

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

<a id="canonical-0112021300202200-2213021330010311-0102101031100131-0112333322200010-2113213021203232-2211002021100203-1220130023302233-0321110221013111"></a>

<a id="canonical-1031230212123210-0332000313221003-0120002330212200-0033321230322310-3201322331130310-1003120113210233-3103132320112333-0013103220201232"></a>

## tenant property — interface / 133320300203 / 7

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

<a id="canonical-0301032210323221-3132231001113302-1022000101231031-2003220301203310-1203000223203200-0201221210102320-2231111013033201-1233031002232000"></a>

<a id="canonical-2212320213223302-3310212032113020-1333003220120223-2122232112010223-0301222000023203-3001220011220033-2203133302300100-3132102310132012"></a>

## uid property — interface / 133320300203 / 8

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

<a id="canonical-0100133310223200-1001112021221200-2323002300022133-1223301030230010-1230212233332021-2233301121101330-0012133212033113-3001200113002030"></a>

## Next pages — interface / 133320300203 / 9

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2010310223032010-1203110102332232-0210302230012312-0303123300211302-3110210111202302-1033132101032012-1200312331310222-0222123212211033)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1111210130332211-0133323233110021-3201033120130102-0031320231020232-0132110031133333-0301312022302130-2020001130131033-1312012120211312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010032321103313-2311100233222312-3213201302113310-0203001112232211-0313133311322101-3312002022003011-3222302031101102-0212232020020230"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 231233020120 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2010310223032010-1203110102332232-0210302230012312-0303123300211302-3110210111202302-1033132101032012-1200312331310222-0222123212211033)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-2332300033323031-2322332230002220-2002113303330123-1221010301211222-1310112133210021-3220121102331012-0001131103120233-3100112223001200"></a>

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

<a id="canonical-1322320230013222-3033213202103010-2230031031002330-2200331212220313-2012032011010230-1210311331322002-1330002102313320-0223122303333323"></a>

## Direct properties — nexthop_address / 231233020120 / 3

- [dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-0200113313023211-1012131333123203-2120333120031320-2021022221132002-2331300233330021-2310223120233032-3120212201223102-0113011130131202): complete subsection reference.

- [IPv4](resources--aws_vpc_site--reference--group-003.md#canonical-3213022120302001-3213323302302020-2001202103110213-3113313022201012-2222012220113320-2310100010022322-1203330103102130-2200330130031333): complete subsection reference.

- [IPv6](resources--aws_vpc_site--reference--group-003.md#canonical-2230023330001312-3332300131113013-1111103211310233-1120212030200211-0313132221020233-1111000111121223-3223211221211032-1120033213120022): complete subsection reference.

<a id="canonical-0000021330311323-1100203222100211-1113101212201000-3222200131033023-0133333103313132-2100321112110320-0131122100003230-0010312110030313"></a>

## Next pages — nexthop_address / 231233020120 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-0200113313023211-1012131333123203-2120333120031320-2021022221132002-2331300233330021-2310223120233032-3120212201223102-0113011130131202)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-3213022120302001-3213323302302020-2001202103110213-3113313022201012-2222012220113320-2310100010022322-1203330103102130-2200330130031333)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-2230023330001312-3332300131113013-1111103211310233-1120212030200211-0313132221020233-1111000111121223-3223211221211032-1120033213120022)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2010310223032010-1203110102332232-0210302230012312-0303123300211302-3110210111202302-1033132101032012-1200312331310222-0222123212211033)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0200113313023211-1012131333123203-2120333120031320-2021022221132002-2331300233330021-2310223120233032-3120212201223102-0113011130131202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113221203302213-1223120233022030-3321332102310102-1211133310232101-0010033233012002-0030202321130010-0012223201222201-1323303203222301"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 321322031023 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2010310223032010-1203110102332232-0210302230012312-0303123300211302-3110210111202302-1033132101032012-1200312331310222-0222123212211033)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-1111210130332211-0133323233110021-3201033120130102-0031320231020232-0132110031133333-0301312022302130-2020001130131033-1312012120211312)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-0311312110101002-1321223301311121-0230322023211212-0100310233301210-1021003111122201-2302213201303002-2333123212001113-3202310100003313"></a>

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

<a id="canonical-1113212133123113-2010100021023231-3331322200111030-3333311030300110-3031222030130321-1023233102010111-3303030333033032-1300111320231022"></a>

## Direct properties — dual_stack / 321322031023 / 3

- [IPv4](resources--aws_vpc_site--reference--group-003.md#canonical-0301332023111030-1333232000300002-2203232012103123-3010212222312030-2013032221210300-1301312113321230-2300123202221030-1320203230310130): complete subsection reference.

- [IPv6](resources--aws_vpc_site--reference--group-003.md#canonical-0022003031012020-1111231032103220-2223011202002212-1321000000120310-0122132032020033-0201031213300301-3112311310130230-2323301113011011): complete subsection reference.

<a id="canonical-1310213313310322-0310010122121322-2230102230002313-2031320333301200-0001031232232003-3322302300100312-0301000231101111-0003003201221320"></a>

## Next pages — dual_stack / 321322031023 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-0301332023111030-1333232000300002-2203232012103123-3010212222312030-2013032221210300-1301312113321230-2300123202221030-1320203230310130)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-0022003031012020-1111231032103220-2223011202002212-1321000000120310-0122132032020033-0201031213300301-3112311310130230-2323301113011011)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-1111210130332211-0133323233110021-3201033120130102-0031320231020232-0132110031133333-0301312022302130-2020001130131033-1312012120211312)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0301332023111030-1333232000300002-2203232012103123-3010212222312030-2013032221210300-1301312113321230-2300123202221030-1320203230310130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201323313213110-2301221003231323-0103031003032331-3333010330123110-1330013211122011-0110033023010313-3100322321020213-3121111030310223"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 313021210011 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2010310223032010-1203110102332232-0210302230012312-0303123300211302-3110210111202302-1033132101032012-1200312331310222-0222123212211033)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-1111210130332211-0133323233110021-3201033120130102-0031320231020232-0132110031133333-0301312022302130-2020001130131033-1312012120211312)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-0200113313023211-1012131333123203-2120333120031320-2021022221132002-2331300233330021-2310223120233032-3120212201223102-0113011130131202)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-2221021012202230-3222222003100203-3231030121221132-0012200001020202-1010322003030322-3232221131121031-3003201032231033-2112230330033122"></a>

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

<a id="canonical-2303222331310002-2111311122120131-1232130200112013-3000220100003300-0132302120233230-1213332321132110-2301333310031023-1320322032232100"></a>

## Direct properties — IPv4 / 313021210011 / 3

<a id="canonical-2201233021121013-3202322313011130-3123101230300331-0001311112022233-3121032302210012-1102012313030330-0331013003302000-1323202222322312"></a>

<a id="canonical-1313021120032202-3010322003120322-1302211002003202-2012132232130210-2121123120003300-1022203102021310-0101111312031231-2010311103012101"></a>

## addr property — IPv4 / 313021210011 / 4

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

<a id="canonical-0223202233220103-1033320132113110-1132221022221211-3130020311103100-0211120300210310-0303112200131213-3011013312121303-2320113210311323"></a>

## Next pages — IPv4 / 313021210011 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-0200113313023211-1012131333123203-2120333120031320-2021022221132002-2331300233330021-2310223120233032-3120212201223102-0113011130131202)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0022003031012020-1111231032103220-2223011202002212-1321000000120310-0122132032020033-0201031213300301-3112311310130230-2323301113011011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110023232100200-1211102301003322-2320102232013130-2321022320331310-0221321232023313-0021330221302200-3222311102220313-0021133001333330"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 103131003220 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2010310223032010-1203110102332232-0210302230012312-0303123300211302-3110210111202302-1033132101032012-1200312331310222-0222123212211033)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-1111210130332211-0133323233110021-3201033120130102-0031320231020232-0132110031133333-0301312022302130-2020001130131033-1312012120211312)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-0200113313023211-1012131333123203-2120333120031320-2021022221132002-2331300233330021-2310223120233032-3120212201223102-0113011130131202)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-0320013231233130-3230333203011101-1013023120110212-3031331222103332-1302002103130113-1221210013031231-3022030312220223-0220113331213003"></a>

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

<a id="canonical-0213110122223220-2002320011002022-2033021010310303-1302210231010020-1021133222211113-1331213212112202-0022033121002023-1031300203101103"></a>

## Direct properties — IPv6 / 103131003220 / 3

<a id="canonical-2201223221321132-2101022133023131-2310032312131022-1103200300201010-0032101001301021-0021331123031212-1103222220012210-1210002333121020"></a>

<a id="canonical-0031223101222103-2112103220113303-1331312020302213-1131133030222000-0210023221313323-0312211310000023-2200003023000030-2333333230202232"></a>

## addr property — IPv6 / 103131003220 / 4

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

<a id="canonical-2231120200320230-2022123021203113-1201121111232302-1232301333020313-0322321000201300-1000113200331333-1230211200220231-2332222232001132"></a>

## Next pages — IPv6 / 103131003220 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-0200113313023211-1012131333123203-2120333120031320-2021022221132002-2331300233330021-2310223120233032-3120212201223102-0113011130131202)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3213022120302001-3213323302302020-2001202103110213-3113313022201012-2222012220113320-2310100010022322-1203330103102130-2200330130031333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022220232011223-1302232003022021-2133130323133101-1131022233233110-1022110102233100-1023303130103131-3110223111213132-2313333223301331"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 332333200301 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2010310223032010-1203110102332232-0210302230012312-0303123300211302-3110210111202302-1033132101032012-1200312331310222-0222123212211033)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-1111210130332211-0133323233110021-3201033120130102-0031320231020232-0132110031133333-0301312022302130-2020001130131033-1312012120211312)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-0301230120002131-3011033331002333-3332112332131221-0123202310010033-3220102121033323-2231132231032322-2112112323203020-0103233112002131"></a>

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

<a id="canonical-3221222331222110-2321011001003310-0202331323220002-1130120221032330-0113120133303031-2131223110102310-2330022223102030-2012320020023032"></a>

## Direct properties — IPv4 / 332333200301 / 3

<a id="canonical-1312132320100300-2232231131331120-1321001310322210-0333130330301233-0311011333312321-0232200012032030-0010202301203232-0322311310022333"></a>

<a id="canonical-1233120000301333-0121203012201121-1033123322323330-3132213031021313-2300321120132001-2312102133210021-0130011331203303-2012321312002112"></a>

## addr property — IPv4 / 332333200301 / 4

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

<a id="canonical-1211231303002221-2331122313100210-3232120031230301-2002201333033332-3013203130302101-1221123010031122-1320212330111221-1031131001121131"></a>

## Next pages — IPv4 / 332333200301 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-1111210130332211-0133323233110021-3201033120130102-0031320231020232-0132110031133333-0301312022302130-2020001130131033-1312012120211312)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2230023330001312-3332300131113013-1111103211310233-1120212030200211-0313132221020233-1111000111121223-3223211221211032-1120033213120022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320321000112202-3121003312213001-1000232302312100-0111330221303211-2213023103330322-1322013201131213-1211102210203123-3110323023310300"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 311222033200 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2010310223032010-1203110102332232-0210302230012312-0303123300211302-3110210111202302-1033132101032012-1200312331310222-0222123212211033)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-1111210130332211-0133323233110021-3201033120130102-0031320231020232-0132110031133333-0301312022302130-2020001130131033-1312012120211312)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-1332101123202301-1023133322222210-1122013333233013-2321333220223101-1120232023021113-3012101201233003-1020011130111220-2112331321130110"></a>

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

<a id="canonical-0023331221302133-0303103130003030-1011200001030331-1213333210232122-3223120112301332-1223311121010232-3130103203001103-0230200122102220"></a>

## Direct properties — IPv6 / 311222033200 / 3

<a id="canonical-0331103323112320-3102232013003231-0303323323310322-3002031103130100-0203230013031113-2302021023100130-3120013102230210-3310101301330012"></a>

<a id="canonical-1320112332012032-1210021033322322-0201321222111121-1313312323010013-1012221110102001-3320200322133330-3222032201312333-0012012111301212"></a>

## addr property — IPv6 / 311222033200 / 4

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

<a id="canonical-3010221033130012-1123232000033302-3012223020311301-2313033022231112-0221311011022130-3121333132202021-0021003330020212-2132031012031203"></a>

## Next pages — IPv6 / 311222033200 / 5

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-1111210130332211-0133323233110021-3201033120130102-0031320231020232-0132110031133333-0301312022302130-2020001130131033-1312012120211312)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1203330202311220-2211112031100000-3032121030223223-0101223300032312-1201112223301330-0130230100011130-1113212311120301-3313331032231122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220230322330221-2323312002130020-2232112311131200-3121321020002112-1233031233032011-2131123112010220-1001030021002022-3000303312312012"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets — subnets / 001311301032 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-2110030202203003-1323213220311123-1310303103222132-3032020233111320-2231220012100201-3232100211302102-2110020200213022-0002211000113123"></a>

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

<a id="canonical-2103100011202000-3032011223303011-1000030231033303-1101030112110102-3130110310311331-1310323003223322-0223323301301122-2223323131030122"></a>

## Direct properties — subnets / 001311301032 / 3

- [IPv4](resources--aws_vpc_site--reference--group-003.md#canonical-1002210123121010-1010122130102023-3102110213331233-1011200312300012-2202033220010133-1202301312333221-0012031210012302-1011010300212011): complete subsection reference.

- [IPv6](resources--aws_vpc_site--reference--group-003.md#canonical-0112200313131000-2021132020223031-2212202022022000-0222103222201112-3330021031102003-3033121311222121-2312113101220130-3103301210122000): complete subsection reference.

<a id="canonical-1112123323120130-0130013202110232-0330010230213310-0233011310112010-3202311312310321-3112221212110322-3101333132013123-2332221303321232"></a>

## Next pages — subnets / 001311301032 / 4

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-1002210123121010-1010122130102023-3102110213331233-1011200312300012-2202033220010133-1202301312333221-0012031210012302-1011010300212011)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-0112200313131000-2021132020223031-2212202022022000-0222103222201112-3330021031102003-3033121311222121-2312113101220130-3103301210122000)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1002210123121010-1010122130102023-3102110213331233-1011200312300012-2202033220010133-1202301312333221-0012031210012302-1011010300212011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033230123330032-2113131002030230-3320003103102203-3013031113010020-0133211313212023-2023120210002210-0011322211202111-0220212023313111"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 202211301231 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-1203330202311220-2211112031100000-3032121030223223-0101223300032312-1201112223301330-0130230100011130-1113212311120301-3313331032231122)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-2031110202033201-0221220122323003-1203213021021031-1220222202221333-1310111011010103-3302331233303031-1023211313310122-3333030101212231"></a>

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

<a id="canonical-1012200101203120-1230313233020201-3330120320011323-2102002132022131-1111222322202030-2323313323201201-1323210110311131-1201113332131203"></a>

## Direct properties — IPv4 / 202211301231 / 3

<a id="canonical-0001100101112133-2130222202031113-2030320333103312-1131212032033113-0121030302133112-1301111213303010-0111310011333232-0100320103233113"></a>

<a id="canonical-3121302103331222-1212223210112120-2100200200210111-1210032111010012-2220130220303002-0222333110202010-1002000302220230-1131130202321203"></a>

## plen property — IPv4 / 202211301231 / 4

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-0110003333001000-0003033333220021-2332130233202311-1323030023110220-3032213332330211-0033230030010030-1202121313032133-0321113012331022"></a>

<a id="canonical-0322310333021222-0101133200210221-2033122323111113-1330113202212333-1211333331302201-2220320210003000-0101013031311030-0103031331212120"></a>

## prefix property — IPv4 / 202211301231 / 5

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

<a id="canonical-0032130310123102-3112122322211322-2233132231302110-3310203220301311-0231333321022213-3221311200012020-0332031021330111-2000333123320223"></a>

## Next pages — IPv4 / 202211301231 / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-1203330202311220-2211112031100000-3032121030223223-0101223300032312-1201112223301330-0130230100011130-1113212311120301-3313331032231122)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0112200313131000-2021132020223031-2212202022022000-0222103222201112-3330021031102003-3033121311222121-2312113101220130-3103301210122000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200131002023202-0231002303212013-2330323113303231-0302122302302310-0013231121011110-2112000021221103-3011230312232001-3121122311321320"></a>

## ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 232302230201 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032)
- [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-2122022330331131-1130130320201110-2233030222101031-0310012323130223-1102312212303231-0231300210100211-3233112222023132-3223211102222331)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2210101200332312-1210220032322100-1210300323230303-0223011001233032-3131220221302020-1023212022210300-0130333031323231-1112301102332002)
- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-1203330202311220-2211112031100000-3032121030223223-0101223300032312-1201112223301330-0130230100011130-1113212311120301-3313331032231122)
- ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-2022201002000120-0103323202020331-2210133202313032-3113110011211101-2010100332122221-1132113203231210-2100221121331212-0200021131301311"></a>

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

<a id="canonical-1111111122030211-0321310033221201-0000011220022210-0102332301321200-0322001310302001-2220302003133123-0310023101021013-1200333312013123"></a>

## Direct properties — IPv6 / 232302230201 / 3

<a id="canonical-2012131202330201-3030000032313310-1233311233102203-1020001320021110-2101001222103001-1020101010122231-3130330303121311-1121200331211221"></a>

<a id="canonical-2101110021022230-2311003030310332-2221232203230112-1010112001113333-2322031123230300-3300030003023302-1333311200213003-3002123231222210"></a>

## plen property — IPv6 / 232302230201 / 4

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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-1121122111011021-1213022132303303-1130100221032312-1333121202210133-0023203122323100-2020103121012222-1122101230211231-1022213220113222"></a>

<a id="canonical-3302220003000033-0201213321001021-2211102000213003-3132131021211213-1300202232223131-2031330220311313-0300323010023330-1330010021323331"></a>

## prefix property — IPv6 / 232302230201 / 5

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

<a id="canonical-0313130112211022-2212201030002102-1023212033012323-0130303102001231-2303332132132232-3202100223002121-2321013001121221-3311002311132313"></a>

## Next pages — IPv6 / 232302230201 / 6

- [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-1203330202311220-2211112031100000-3032121030223223-0101223300032312-1201112223301330-0130230100011130-1113212311120301-3313331032231122)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2102113123102003-2220310101303013-0300012200101230-2331120032031002-2020320223320023-1023022111122321-3001131033200032-2103330230232000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002010223331133-3000020103211303-0021022331111022-1121202123133032-1212133013322320-2323301232320033-2313013312203112-1131200330120122"></a>

## ingress_egress_gw.no_dc_cluster_group — no_dc_cluster_group / 322213320301 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.no_dc_cluster_group

<a id="canonical-3112131303332121-3110023232023132-0223013220102031-2200102333233130-3011020103012223-1201302310333023-1200122123032023-2010130221213023"></a>

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

<a id="canonical-3212312313112203-0030212021000223-0321311012132101-2132231012313131-2113223102221023-0323300100212300-3210122122013102-2103020023212022"></a>

## Direct properties — no_dc_cluster_group / 322213320301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102320101303003-1000332202022123-0312331010013223-3302210310323030-0003312033023111-2010120032322222-2311303222233300-1212010321212222"></a>

## Next pages — no_dc_cluster_group / 322213320301 / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0300100013213131-2213313111311121-2320032011010033-1121103120121321-3332103121230003-0311021312130233-1200313321213133-3212230332110001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012333220320332-3031102203202330-1213313101330332-3300231302322131-0122313201122323-0001222310233200-2331301203313110-0202321001131013"></a>

## ingress_egress_gw.no_forward_proxy — no_forward_proxy / 323102321102 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.no_forward_proxy

<a id="canonical-0212303303301021-1222112100211312-2002330312303032-2123232200130232-2300020112303122-0011310221101123-1100030130023220-3101011022113121"></a>

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

<a id="canonical-1023101323110131-0231231221110012-1332313101020020-2333133310232101-2023001203012022-3322201312130303-1130301202110033-1322101330102113"></a>

## Direct properties — no_forward_proxy / 323102321102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030023120131221-0030221320010103-3301210032311333-1022112031210023-3310111300233303-1121231130001223-2102033122323212-2113313122131232"></a>

## Next pages — no_forward_proxy / 323102321102 / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2303130302120032-0011211313313230-1011003323201320-2200003100222230-1300022023002331-3030200233303333-3232203312323132-0003221101001120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001012213003131-3330311020021300-2023221312121102-2120113012222210-0211102101320123-3220111020103333-0131011110101131-1133033102132003"></a>

## ingress_egress_gw.no_global_network — no_global_network / 102300333203 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.no_global_network

<a id="canonical-1320200012001122-1020212303120121-3032021211011331-3101301100210233-3322013002112303-0222103200103220-0323000223111212-1120233120330133"></a>

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

<a id="canonical-1322133103030001-0110313203221133-0031013332211320-3231221100332202-0213133010103000-0322010312313033-2300130001031031-3030101200021101"></a>

## Direct properties — no_global_network / 102300333203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210102122023000-2223311123301010-3231300030220213-3110011130311032-1332210230320202-2222022033111201-3202222122201200-2332011220132031"></a>

## Next pages — no_global_network / 102300333203 / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1333002030033131-2021313011200121-0222321022222323-0132313232113010-1221003222232002-0113220201003323-2012030231320232-2301303310222211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030301222131010-0113303022011003-0211133002013323-3330320112210200-3021110210231020-3112233223031133-0302022213012333-1013302210122010"></a>

## ingress_egress_gw.no_inside_static_routes — no_inside_static_routes / 000100302102 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.no_inside_static_routes

<a id="canonical-0121003122231131-2302013320231322-0111013002111230-0011012203311312-1232331203110333-2312203211022231-2103220301311321-3010333030123303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no inside static routes.

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
no_inside_static_routes = {}
```

<a id="canonical-1313001031122311-2330123133133100-2020213023001330-0201320221011002-1012222023320222-1222130032233022-0222023002103130-3223111201030023"></a>

## Direct properties — no_inside_static_routes / 000100302102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321302213313020-0200231210311302-3001330312202203-1023121001130220-1121230022021312-2023321220213221-2022030230130301-2133010132023313"></a>

## Next pages — no_inside_static_routes / 000100302102 / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0333203101111010-1010131002112103-3313120201301303-3330022000222010-2110330132312202-3231120103313130-2213103003331033-2202002022021232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131302222110002-1312213220001322-1113202200311301-3320211110033223-0323131102023333-1013211333313101-1110231220233201-2331121101310210"></a>

## ingress_egress_gw.no_network_policy — no_network_policy / 323100131320 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.no_network_policy

<a id="canonical-2010231202123310-0010000131231313-2233120101013100-0323233133201011-3003322133003211-3203223103103303-1002110212201231-2312003311000302"></a>

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

<a id="canonical-2221033321112033-1113003132330322-2101111230230132-0223121223223322-3122200020102323-2202001301133131-0021131131220030-0130113130320122"></a>

## Direct properties — no_network_policy / 323100131320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202101332331323-3120220123220121-2310303203113211-1121301113121201-2303131013012210-0022120103332303-3103323233133213-2010022302220223"></a>

## Next pages — no_network_policy / 323100131320 / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3212221300032003-1100311223312301-0322102032310120-2122130010202011-0310231012101011-1233013031023021-3230010032221120-2221220331311101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202202020201321-2100010202130320-3210032202301002-3030323220113212-1111112313210333-0213212033333211-2021003022113131-3221003120110311"></a>

## ingress_egress_gw.no_outside_static_routes — no_outside_static_routes / 332311300331 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.no_outside_static_routes

<a id="canonical-2202010003222333-2313221121130112-3313213333100131-2312102021020010-0321212002033020-1113100230101001-0220121023312132-2130202210111323"></a>

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

<a id="canonical-0230212220300212-3222020101331030-1120223123120223-2032123001032003-0223300203312233-3012323011000113-2100321233303332-0322112012211132"></a>

## Direct properties — no_outside_static_routes / 332311300331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210233020021121-3322302200021121-0232233313033023-0010031302333321-2021102223310320-1130201030033201-1032122213321013-0301033203031212"></a>

## Next pages — no_outside_static_routes / 332311300331 / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131002121231000-1012233222103132-3210033233013132-0103032300322001-1133101023133102-3002312013210013-3023330302230112-2011323201102032"></a>

## ingress_egress_gw.outside_static_routes — outside_static_routes / 132332020300 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.outside_static_routes

<a id="canonical-3333311300120330-0302331112121013-1203321331012102-2100011313302000-2110203231220211-2302113002302233-1323022133113331-1111032201223120"></a>

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

<a id="canonical-3130200203300212-0231202101020203-3131103010120122-0231111331223012-1221103003201020-2033230202200110-1111130120212131-2012113312123202"></a>

## Direct properties — outside_static_routes / 132332020300 / 3

- [static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033): complete subsection reference.

<a id="canonical-3323032123203003-1231013021313201-0132010211010031-3303013233332231-1313221111121022-1002302011233101-1122033101223303-0113020310113212"></a>

## Next pages — outside_static_routes / 132332020300 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123020023313000-3001330330130023-3221232101330131-1000103113101220-0100030003231310-3233201320321212-2011031323032021-0132130310210202"></a>

## ingress_egress_gw.outside_static_routes.static_route_list — static_route_list / 102202022032 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002)
- ingress_egress_gw.outside_static_routes.static_route_list

<a id="canonical-0331330322000311-2112233202303212-2023221320003110-2331323220210310-0030333002212030-2202210332021033-0000223113101303-0232301321230102"></a>

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

<a id="canonical-0300311011121311-3311033131010220-3122000121221222-3231013320232221-1113031112303030-0203312330102303-3232220301302131-1310321132220232"></a>

## Direct properties — static_route_list / 102202022032 / 3

- [custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111): complete subsection reference.

<a id="canonical-3300202110200021-3131200222231130-0003120031103123-1233000102231313-3011110121110131-0331023323203310-2302000313203113-0311000021101032"></a>

<a id="canonical-1203231323132021-2213322302211232-1122103121113022-0022112313220211-0102220222223310-3003331302110221-1333002011120133-3320311211200020"></a>

## simple_static_route property — static_route_list / 102202022032 / 4

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

<a id="canonical-2313111132031221-2202220133022331-0021031221112030-1020301121213112-2130313200230100-3313322102312220-2100330233300310-0132331100012311"></a>

## Next pages — static_route_list / 102202022032 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322101320001212-0320222122012202-2222211112111210-0231003203010200-3303211213000333-2303101100112300-1020301302202110-1123123301032302"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route — custom_static_route / 301212112322 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-3102101030021011-1113230232321010-0000310212303300-2330330200211202-1213132110112301-1331121312331313-3111013100233002-0302321220220032"></a>

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

<a id="canonical-2220003201311300-2002130312011132-3211230131230010-3011001013301021-1310221330231320-3313302030103311-3312200322202022-0012100112320223"></a>

## Direct properties — custom_static_route / 301212112322 / 3

<a id="canonical-1211021121310010-2121100222033130-1222212313130223-0122332320010110-0010231231233032-0201320002013133-3020003111122121-1122333230203320"></a>

<a id="canonical-2003001211110020-2020003303131222-3112200233111312-1212202113020210-1030001123011202-3013322021213322-2122120023330222-0130133030112200"></a>

## attrs property — custom_static_route / 301212112322 / 4

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](resources--aws_vpc_site--reference--group-003.md#canonical-3003103023102111-0122301020000023-2133110220132321-3331133120130333-2210030023123031-3120010333002002-3301002333002323-2323023113113202): complete subsection reference.

- [nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2012311123033121-3000003102022201-1113232202321231-2303313002101021-2230111010111230-0303121032103110-2121131301022000-2311222333303122): complete subsection reference.

- [subnets](resources--aws_vpc_site--reference--group-003.md#canonical-2031222123323301-1210132111220001-2303002131221011-3201311130213300-0310120003111322-2000132101221223-0221333321123131-0001103102302331): complete subsection reference.

<a id="canonical-3320211023211003-1200001010122213-1230023303031011-0132221112033001-2032031202001023-3132303131022231-3300202122000022-3332012022012210"></a>

## Next pages — custom_static_route / 301212112322 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](resources--aws_vpc_site--reference--group-003.md#canonical-3003103023102111-0122301020000023-2133110220132321-3331133120130333-2210030023123031-3120010333002002-3301002333002323-2323023113113202)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2012311123033121-3000003102022201-1113232202321231-2303313002101021-2230111010111230-0303121032103110-2121131301022000-2311222333303122)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-2031222123323301-1210132111220001-2303002131221011-3201311130213300-0310120003111322-2000132101221223-0221333321123131-0001103102302331)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3003103023102111-0122301020000023-2133110220132321-3331133120130333-2210030023123031-3120010333002002-3301002333002323-2323023113113202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313310002310320-2313332323233001-3310020311103302-2320012332021033-1301332030123300-0032101022223032-1103322003303230-3202023111120300"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels — labels / 133311312213 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-3132222113232031-1210030102001021-1313120230303110-3030133323122233-2332030221200303-3223022112203023-2032032003302230-3130312020000303"></a>

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

<a id="canonical-1322112132331332-0333111211032231-3322121300300313-0313331210233303-0221233023302120-3011203000221333-0101332020100122-2232210133201301"></a>

## Direct properties — labels / 133311312213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121032222302012-0202302112333310-3300303203113221-2220322322233331-0133233020102311-0121120123133213-1332122233133111-3300233022210313"></a>

## Next pages — labels / 133311312213 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2012311123033121-3000003102022201-1113232202321231-2303313002101021-2230111010111230-0303121032103110-2121131301022000-2311222333303122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331110200202321-2311321331132030-3200010022322012-3001331312030233-0021023222321003-2123003122210030-1233322122232013-1121100132001012"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 211321221323 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-3000123131123132-3002031021000133-1011310032302130-3321012112012102-0202200022112222-0013321311001032-1211021021003030-3110131230202112"></a>

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

<a id="canonical-1220200212013132-0213030201121100-2101321210100003-0023000133002030-1001113132001331-0021033113331033-2231310131332320-2210112130113111"></a>

## Direct properties — nexthop / 211321221323 / 3

- [interface](resources--aws_vpc_site--reference--group-003.md#canonical-2020001122122200-3100012231121110-2103323300000321-0112111333110021-3300300100202111-2113031223232333-1022032112311333-3223010020203210): complete subsection reference.

- [nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-3133333330200130-3302121123133330-3230212021100321-0131332302310000-0033200020103321-2223300111230003-0320310132003001-2003002021322003): complete subsection reference.

<a id="canonical-2030123031033313-1321311310312330-1111301123202302-1303312013123202-2312300101330001-2100310020232302-1333210032210020-2202100210130132"></a>

<a id="canonical-3120230301201110-2032332100100022-2211302010132022-3313123312022100-2010122110212311-3000110033133202-2022012120102222-3023201130002031"></a>

## type property — nexthop / 211321221323 / 4

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

<a id="canonical-1212102221031313-3131000302013233-0330322211223003-2330222220211010-3100203010103111-2311112123231301-2233203113320123-2030121101311023"></a>

## Next pages — nexthop / 211321221323 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_vpc_site--reference--group-003.md#canonical-2020001122122200-3100012231121110-2103323300000321-0112111333110021-3300300100202111-2113031223232333-1022032112311333-3223010020203210)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-3133333330200130-3302121123133330-3230212021100321-0131332302310000-0033200020103321-2223300111230003-0320310132003001-2003002021322003)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2020001122122200-3100012231121110-2103323300000321-0112111333110021-3300300100202111-2113031223232333-1022032112311333-3223010020203210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313223133231013-2313022121111202-0321031131310333-3331130020212333-3310331011232022-3331212223031113-1020112020302102-1111211121201100"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 130230010113 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2012311123033121-3000003102022201-1113232202321231-2303313002101021-2230111010111230-0303121032103110-2121131301022000-2311222333303122)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-1011332320112322-1203212303031311-2003321201132223-0201311223121113-0123122231133321-0111312202103203-3002113303132003-2321222100132310"></a>

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

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-1002031100111320-0033112213303232-1211030311130021-0312332120322331-1332111232003300-3030221331021212-1333212223110111-0222200123231311"></a>

## Direct properties — interface / 130230010113 / 3

<a id="canonical-2112113130013313-0313021002133122-0232302322103203-2233212110103122-0213210303023200-0030010133100210-2033220321320203-0032332002311012"></a>

<a id="canonical-1213001200022031-2223231321100322-0123312322013330-3121321002113011-1212333101231213-0212120131302322-3130213211010022-1022202100022023"></a>

## kind property — interface / 130230010113 / 4

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

<a id="canonical-0202320000302332-0021032221310211-1122033300330333-1231020102331203-3321123122033002-1111221130221102-1010320200121313-1220011100210211"></a>

<a id="canonical-2212132231301002-2121230222211101-2122023111311320-2103011033212121-0313123033303121-3132032000000032-1132002111323233-2111023133120002"></a>

## name property — interface / 130230010113 / 5

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

<a id="canonical-2121012332110301-3223313022123323-1013130101101211-3213003221201231-3332230220232002-1011310101012002-3202003030202132-0313302033021301"></a>

<a id="canonical-0011230023001330-3203211333030303-0112121313231001-2130101333231013-3011100132300213-0001223033133330-3131223012231021-1310311003012031"></a>

## namespace property — interface / 130230010113 / 6

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

<a id="canonical-1031002120323231-3133200013132312-1000130003230021-1310012113221201-3131030033300220-1220120032220112-0102021301012230-3223003301012101"></a>

<a id="canonical-2313132113333333-2222112213102011-3230000322312221-1332331330212231-1211313113103032-0001312213212210-0201223000110220-3001302232123102"></a>

## tenant property — interface / 130230010113 / 7

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

<a id="canonical-1023200020333112-3130330100300003-3122111112312102-3222233010120012-1321010010303033-3030011331010011-2030110301300200-2222023323211113"></a>

<a id="canonical-1131223121120130-3323310102212021-2120221000033133-3301310233202301-1320011020212301-3213220332303320-1030013213121330-1303022020032100"></a>

## uid property — interface / 130230010113 / 8

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

<a id="canonical-1311223333000313-2311203002311032-0331132110000200-0113213311011221-2012133012020313-3103133330313121-2210112220011212-1220011001200133"></a>

## Next pages — interface / 130230010113 / 9

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2012311123033121-3000003102022201-1113232202321231-2303313002101021-2230111010111230-0303121032103110-2121131301022000-2311222333303122)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3133333330200130-3302121123133330-3230212021100321-0131332302310000-0033200020103321-2223300111230003-0320310132003001-2003002021322003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030303213212103-0222003011120002-2300013022303313-2212100130203322-1232221211013203-1211321202312322-0110320313031200-0003121130302110"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 322013012131 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2012311123033121-3000003102022201-1113232202321231-2303313002101021-2230111010111230-0303121032103110-2121131301022000-2311222333303122)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-0201113313030200-1033011223223020-1331012321231122-2330322322200012-2022013010020103-1332031030311011-2102211121311202-0113331310320101"></a>

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

<a id="canonical-1333311121023020-2001003331300213-3022213130221012-2101230123222131-0200310223133320-1201332122032200-2132011123120323-2120223211203202"></a>

## Direct properties — nexthop_address / 322013012131 / 3

- [dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-2020112032103330-2110113312300033-2132300012010100-1010312101120120-2221113333301232-1110103101030120-0222320031201031-1023300201100310): complete subsection reference.

- [IPv4](resources--aws_vpc_site--reference--group-003.md#canonical-3223013211222101-1021133121322012-1101122123200200-0031131101222231-1210132232301101-3323123132123302-2231330220210311-0021201221110102): complete subsection reference.

- [IPv6](resources--aws_vpc_site--reference--group-003.md#canonical-3103020221101203-0300311020000022-1311233110130021-3333230312203110-2011201223203233-3233300131031133-1123213110210000-0023001301322300): complete subsection reference.

<a id="canonical-2123021132123300-3230023031211132-0200320220322032-3230113330312232-3132332231130101-2303322120103123-0313322001231302-1310331111133220"></a>

## Next pages — nexthop_address / 322013012131 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-2020112032103330-2110113312300033-2132300012010100-1010312101120120-2221113333301232-1110103101030120-0222320031201031-1023300201100310)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-3223013211222101-1021133121322012-1101122123200200-0031131101222231-1210132232301101-3323123132123302-2231330220210311-0021201221110102)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-3103020221101203-0300311020000022-1311233110130021-3333230312203110-2011201223203233-3233300131031133-1123213110210000-0023001301322300)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2012311123033121-3000003102022201-1113232202321231-2303313002101021-2230111010111230-0303121032103110-2121131301022000-2311222333303122)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2020112032103330-2110113312300033-2132300012010100-1010312101120120-2221113333301232-1110103101030120-0222320031201031-1023300201100310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302030022031201-1013333021211132-3222122001313022-2000220110221223-3221323322133010-3201200030122001-3132120111333023-0001313132101002"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 130030001223 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2012311123033121-3000003102022201-1113232202321231-2303313002101021-2230111010111230-0303121032103110-2121131301022000-2311222333303122)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-3133333330200130-3302121123133330-3230212021100321-0131332302310000-0033200020103321-2223300111230003-0320310132003001-2003002021322003)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-2202331123211333-0332132330001021-3121222030022010-1120200002120110-3210121112211231-2202312122311333-2023310300223010-0021303210233220"></a>

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

<a id="canonical-0111102232113320-2332011022033310-1212210223130211-3330100323333311-3211310031331021-1313212120200200-0230301131321113-3210003113002132"></a>

## Direct properties — dual_stack / 130030001223 / 3

- [IPv4](resources--aws_vpc_site--reference--group-003.md#canonical-1011002211010033-2132003222232313-1111131100312200-3121320200311120-0001112011121310-2102013020021220-2330223330202002-3221230210000130): complete subsection reference.

- [IPv6](resources--aws_vpc_site--reference--group-003.md#canonical-0011011030211222-1221031013031123-0210121311013332-3100011023223001-0200223110011000-0103302102332200-0100002212123303-0012210031221130): complete subsection reference.

<a id="canonical-2132223130201211-2210101100033223-2113102010131233-2033332331312112-0301131210321023-3023230032032221-1320321122110123-1011203331233331"></a>

## Next pages — dual_stack / 130030001223 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-1011002211010033-2132003222232313-1111131100312200-3121320200311120-0001112011121310-2102013020021220-2330223330202002-3221230210000130)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-0011011030211222-1221031013031123-0210121311013332-3100011023223001-0200223110011000-0103302102332200-0100002212123303-0012210031221130)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-3133333330200130-3302121123133330-3230212021100321-0131332302310000-0033200020103321-2223300111230003-0320310132003001-2003002021322003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1011002211010033-2132003222232313-1111131100312200-3121320200311120-0001112011121310-2102013020021220-2330223330202002-3221230210000130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331032210210032-1330002232333120-3230033303001311-2331322231010233-0121222302331112-1220211232302221-0121001233231033-3103230120103213"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 122310330002 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2012311123033121-3000003102022201-1113232202321231-2303313002101021-2230111010111230-0303121032103110-2121131301022000-2311222333303122)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-3133333330200130-3302121123133330-3230212021100321-0131332302310000-0033200020103321-2223300111230003-0320310132003001-2003002021322003)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-2020112032103330-2110113312300033-2132300012010100-1010312101120120-2221113333301232-1110103101030120-0222320031201031-1023300201100310)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-0323311131213332-1030220001123033-1211330122332301-3030220013020003-3200030131122320-3121313310233120-3300220200332230-1133020213230030"></a>

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

<a id="canonical-1323131021231031-0010201222121001-2112113103032331-1212113021333322-0031132001030103-1020323332322223-3032333113011330-0232013321300313"></a>

## Direct properties — IPv4 / 122310330002 / 3

<a id="canonical-3301303210301203-0102122031111322-2113320230333130-1213112221321321-2002320313133031-2130220232031033-0201231012110113-2110211332100200"></a>

<a id="canonical-2300131323132202-0223002222321301-1100312033300222-2312220331011302-3323310110003121-2203123130232330-0313132210120322-0122320012020320"></a>

## addr property — IPv4 / 122310330002 / 4

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

<a id="canonical-0312232123303202-0012320020203021-2112220112302303-3120001030103120-0302203131022331-1201033311330013-0222323330000212-1220233233112212"></a>

## Next pages — IPv4 / 122310330002 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-2020112032103330-2110113312300033-2132300012010100-1010312101120120-2221113333301232-1110103101030120-0222320031201031-1023300201100310)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0011011030211222-1221031013031123-0210121311013332-3100011023223001-0200223110011000-0103302102332200-0100002212123303-0012210031221130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231232202303100-3103203000210332-2202230201032311-0323230333210300-3011202331100211-3200103333323001-0011303221300230-3301003110100322"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 323023213120 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2012311123033121-3000003102022201-1113232202321231-2303313002101021-2230111010111230-0303121032103110-2121131301022000-2311222333303122)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-3133333330200130-3302121123133330-3230212021100321-0131332302310000-0033200020103321-2223300111230003-0320310132003001-2003002021322003)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-2020112032103330-2110113312300033-2132300012010100-1010312101120120-2221113333301232-1110103101030120-0222320031201031-1023300201100310)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-1121133120210221-2233213332213230-1131331010000013-0122013300321223-0213213012012031-3320120333002333-0233100031130131-1133000330332233"></a>

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

<a id="canonical-0200002323303022-2300211021110031-0212312322033230-1221021012203133-0120112200002202-1310120202010020-3110133011122330-1312312032021121"></a>

## Direct properties — IPv6 / 323023213120 / 3

<a id="canonical-1330030313303230-1311322202131231-0020120232130123-2211200202130022-0023311321231301-3321211110213202-1020110311300320-2232030210210233"></a>

<a id="canonical-3201210331223133-3130311130122301-0112033100211033-0212031130100203-2232011112333231-1312311320101302-3120301021210333-2302232101110032"></a>

## addr property — IPv6 / 323023213120 / 4

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

<a id="canonical-1321020132013330-0330311131211201-3311210020302303-3110001133010110-1103031130311021-2312202201001300-1302220321030330-0121021002031213"></a>

## Next pages — IPv6 / 323023213120 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-2020112032103330-2110113312300033-2132300012010100-1010312101120120-2221113333301232-1110103101030120-0222320031201031-1023300201100310)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3223013211222101-1021133121322012-1101122123200200-0031131101222231-1210132232301101-3323123132123302-2231330220210311-0021201221110102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130030122310322-1022033200233211-0321001101023331-2013302201230210-2222033131300032-2310110220002221-3211030323031033-2003320223211223"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 221331033021 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2012311123033121-3000003102022201-1113232202321231-2303313002101021-2230111010111230-0303121032103110-2121131301022000-2311222333303122)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-3133333330200130-3302121123133330-3230212021100321-0131332302310000-0033200020103321-2223300111230003-0320310132003001-2003002021322003)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-0131202201220212-1301201111303122-2332030302113233-2103222011012122-2233022213120001-1330200232033313-3002213130303133-3200303301221130"></a>

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

<a id="canonical-2333222021003022-2001113132031130-3132213120300312-0100131130321120-2210202312112302-3300220000332201-0021133131320310-1210312300001201"></a>

## Direct properties — IPv4 / 221331033021 / 3

<a id="canonical-3312021213233231-2332300102222011-3003302302033023-2030331003202001-0122110030332000-2123221203121330-0112031203111003-2100200232130300"></a>

<a id="canonical-1333010112311332-0112130101112122-2123231130021003-0300122210002223-1000101202302021-3012313200010032-1331033031003132-1131311200230332"></a>

## addr property — IPv4 / 221331033021 / 4

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

<a id="canonical-0231111013221223-2333122110221220-1230311232202122-0123310131330000-3100132301233310-1033203012001230-1301333332121331-1022321003210233"></a>

## Next pages — IPv4 / 221331033021 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-3133333330200130-3302121123133330-3230212021100321-0131332302310000-0033200020103321-2223300111230003-0320310132003001-2003002021322003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3103020221101203-0300311020000022-1311233110130021-3333230312203110-2011201223203233-3233300131031133-1123213110210000-0023001301322300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130031120331113-3333213311101320-2100013110321201-3102311222111001-0300101023113111-2100132103003313-3220100012010331-0221222222030123"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 230000031020 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-2012311123033121-3000003102022201-1113232202321231-2303313002101021-2230111010111230-0303121032103110-2121131301022000-2311222333303122)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-3133333330200130-3302121123133330-3230212021100321-0131332302310000-0033200020103321-2223300111230003-0320310132003001-2003002021322003)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-2030103132332031-0320022111203313-3022011021122330-3030221121312322-3001331022331223-0103123301330111-0112211330311113-1300113310300223"></a>

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

<a id="canonical-0122110001230133-1220100103131310-2330131132120101-2311131031013310-0320031101122231-0202103103030000-2202122302010131-2231323300221323"></a>

## Direct properties — IPv6 / 230000031020 / 3

<a id="canonical-1212322011011013-0011111101200010-2120000130020023-2310002032320113-1322332013110220-2213223320320112-3301103012212211-3232303123130202"></a>

<a id="canonical-0301212112233313-3220102110210312-3330023020113023-2033200223231331-0330310203003213-2302121103223322-1223300030202033-1201012122203310"></a>

## addr property — IPv6 / 230000031020 / 4

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

<a id="canonical-3332120002302003-3123320012221311-3221303221210003-3001221222101303-1201222022132101-1213333010332013-2002312120231120-2311023312011312"></a>

## Next pages — IPv6 / 230000031020 / 5

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-3133333330200130-3302121123133330-3230212021100321-0131332302310000-0033200020103321-2223300111230003-0320310132003001-2003002021322003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2031222123323301-1210132111220001-2303002131221011-3201311130213300-0310120003111322-2000132101221223-0221333321123131-0001103102302331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101201001303311-3310210020123303-2013113301203101-0132323101210330-1000101333221111-2300213021330211-3223211000112301-0030122203030102"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets — subnets / 111112032323 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-0332122021213301-3100200102310321-3132232123202321-1102302002332213-1310320212230110-1311222122031103-1202130002210331-3022010333231133"></a>

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

<a id="canonical-2010102213000213-0113132223200101-2020110220133310-1323301311212110-0100002020323322-3202101103302022-1023222310121302-0102200123321020"></a>

## Direct properties — subnets / 111112032323 / 3

- [IPv4](resources--aws_vpc_site--reference--group-003.md#canonical-2222021321221212-0320022113122031-3232321102001201-3302213320110213-2122202130102222-0313231310132102-3220123222100323-0322111200211233): complete subsection reference.

- [IPv6](resources--aws_vpc_site--reference--group-003.md#canonical-1000320312313211-0121100300010032-3103123231301121-1023132031201201-1323200021202103-2123313132231121-3222102000320222-1122032220021322): complete subsection reference.

<a id="canonical-0320200230202102-2133223131021311-1022022020332100-0022202302000322-3032011223201201-1010003231131201-0110322320111131-1303313021313132"></a>

## Next pages — subnets / 111112032323 / 4

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-2222021321221212-0320022113122031-3232321102001201-3302213320110213-2122202130102222-0313231310132102-3220123222100323-0322111200211233)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-1000320312313211-0121100300010032-3103123231301121-1023132031201201-1323200021202103-2123313132231121-3222102000320222-1122032220021322)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2222021321221212-0320022113122031-3232321102001201-3302213320110213-2122202130102222-0313231310132102-3220123222100323-0322111200211233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013310302321331-3333321311123202-2101130323011013-0110313210302223-0212202220333000-2201112030020302-3232333131312320-1023321212202130"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 321300130102 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-2031222123323301-1210132111220001-2303002131221011-3201311130213300-0310120003111322-2000132101221223-0221333321123131-0001103102302331)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-1302203031321013-1022131022033022-1312122333232011-3003131320003130-1123030000120301-1222000210312312-3321121302133030-1222033013301312"></a>

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

<a id="canonical-0011003012101332-0233001120232213-0213111220313302-3132110003312300-0011121121200201-3212310021222322-1103023232123321-1121031030330000"></a>

## Direct properties — IPv4 / 321300130102 / 3

<a id="canonical-2103321022033000-2321031012112012-2130111330101002-0201003230220331-3022202033012301-1111112220320122-3120122331333311-0020221320000123"></a>

<a id="canonical-1000310131003231-1233113022320231-3231020223230020-2120032333021213-1131232221100333-2012233030212332-2003022102300201-1200320121000322"></a>

## plen property — IPv4 / 321300130102 / 4

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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-3312111302001003-2013200330200122-3202013210013312-0131223331113210-1113222333113131-3002302132201232-3012030101302031-0200331023223113"></a>

<a id="canonical-3232121033313232-1033033312211122-2302201312000203-0032032321101223-3133113333003233-1232321320220130-0230120100021211-1112223222121123"></a>

## prefix property — IPv4 / 321300130102 / 5

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

<a id="canonical-0232133221310101-1122310222211212-0111103102111302-3133200323110232-0303111020111002-3031122233330233-2130120030211213-1133302012120203"></a>

## Next pages — IPv4 / 321300130102 / 6

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-2031222123323301-1210132111220001-2303002131221011-3201311130213300-0310120003111322-2000132101221223-0221333321123131-0001103102302331)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1000320312313211-0121100300010032-3103123231301121-1023132031201201-1323200021202103-2123313132231121-3222102000320222-1122032220021322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233022113322121-2313123013033003-2001132201303323-1023010333002011-0302310321011002-0012023213332121-0102211110032030-2131232102030230"></a>

## ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 133033112213 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002)
- [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3013223032333132-3102113310201230-3011312112122300-3303003032003031-1320103110013132-2110233121130201-3223223302303033-2311021333231033)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1012231132212302-2322112332122321-3213213313123320-1312302323032121-0222222322132211-0230311213221222-0131321031023310-2102322032212111)
- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-2031222123323301-1210132111220001-2303002131221011-3201311130213300-0310120003111322-2000132101221223-0221333321123131-0001103102302331)
- ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-3200323122000220-0313030100020102-3320232013110222-2220213020132221-3322220223011220-3321213230112022-2323232312311321-1313302023222103"></a>

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

<a id="canonical-1102203312020123-0101232011102203-2200232113320300-2001123000113311-0132320103213131-1332211001302202-2311211130301023-1201323231202322"></a>

## Direct properties — IPv6 / 133033112213 / 3

<a id="canonical-0300112001113023-1221110321202330-2032000300211230-1111332323132302-2020300311113332-3032131311332133-2030330211300100-2201323300201221"></a>

<a id="canonical-2210332220213222-0311221300302233-0301002133202200-0212120021013303-2333032013020000-3031322220020112-3022313121003212-1313003232231331"></a>

## plen property — IPv6 / 133033112213 / 4

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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-1120212231232132-0323122221123132-2210020301232122-1312031201201212-3120210333201310-2332322203321312-3300323211232202-1132201311233011"></a>

<a id="canonical-3303330011022210-1232310222210232-3001320101122021-2232310022301110-0110111300312021-3301222302222132-3110211203230023-2200200311332103"></a>

## prefix property — IPv6 / 133033112213 / 5

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

<a id="canonical-1102020120330311-2301311123301111-2203013230212333-2113133003010132-1032322213001012-2021020330311231-2303111112223303-3320332223313331"></a>

## Next pages — IPv6 / 133033112213 / 6

- [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-2031222123323301-1210132111220001-2303002131221011-3201311130213300-0310120003111322-2000132101221223-0221333321123131-0001103102302331)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2133010232202133-3021221010321021-2013032211210330-3301312300002132-3000011030113011-2311302233333003-2311302030320122-0211301220311201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310311213302030-2331002223120303-0212321022220131-3021330220221011-3330111213330231-1330002021310011-0330200200122212-3010013001323300"></a>

## ingress_egress_gw.performance_enhancement_mode — performance_enhancement_mode / 020230102200 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.performance_enhancement_mode

<a id="canonical-2300123131010232-1123113031113220-1233211112103020-3311002221012221-0201023123110022-3123121012330220-2220002222220033-3022131101101332"></a>

Type: `"object"`. single nested block, Optional.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("perf_mode_l3_enhanced",
    "perf_mode_l7_enhanced")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

Terraform syntax:

```terraform
performance_enhancement_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032330131023003-3211212312302210-2010020233000330-2211003130211131-3230200230212313-1201122210333133-2100023002332123-2113310223232222"></a>

## Direct properties — performance_enhancement_mode / 020230102200 / 3

- [perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-3120100220023123-3100100110001110-0011202001301022-2003210211212220-2002013312323022-1023132022021110-0322300003112103-3021100103202211): complete subsection reference.

- [perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-0232321100103301-3203312303221031-2122111132131303-0203300310312230-0221103131221221-2313301330021232-3200210201223130-2221111333111213): complete subsection reference.

<a id="canonical-1231113201211220-2333111332312323-0133222232012130-3203313322312323-0102111001313330-2121011313113300-3103121332222212-2310331122211130"></a>

## Next pages — performance_enhancement_mode / 020230102200 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-3120100220023123-3100100110001110-0011202001301022-2003210211212220-2002013312323022-1023132022021110-0322300003112103-3021100103202211)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-0232321100103301-3203312303221031-2122111132131303-0203300310312230-0221103131221221-2313301330021232-3200210201223130-2221111333111213)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3120100220023123-3100100110001110-0011202001301022-2003210211212220-2002013312323022-1023132022021110-0322300003112103-3021100103202211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212031130023231-1223121121210133-0233031331110202-3112310000013301-2323222001021303-2121032301102002-2023221122033231-1210021011023111"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced — perf_mode_l3_enhanced / 211303201213 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-2133010232202133-3021221010321021-2013032211210330-3301312300002132-3000011030113011-2311302233333003-2311302030320122-0211301220311201)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-3110202022021311-3033310031210102-2331222030010020-0301031332221131-1122120032322022-0332121321232210-1222122031023221-0300323213133323"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo",
    "no_jumbo")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l3_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-1321233311031301-1222230123003331-2321320233023013-3011012221021221-1022002322313132-2123013023033003-3012000103011200-1211123102301130"></a>

## Direct properties — perf_mode_l3_enhanced / 211303201213 / 3

- [jumbo](resources--aws_vpc_site--reference--group-003.md#canonical-2311120200212032-0110032223032010-0312130102100222-0111103310033201-3000130302120220-2230122332130032-2102330001223112-1011330031202310): complete subsection reference.

- [no_jumbo](resources--aws_vpc_site--reference--group-003.md#canonical-2120022122200330-0123110211322022-0130032000012333-3331233320202011-0013320201223001-1220310112021303-1132230203210200-0213030022222322): complete subsection reference.

<a id="canonical-0202212022233101-3020203002003102-3033210122101313-1222322301003210-2112123013301331-0302321223220121-0300203112313122-2312332300203033"></a>

## Next pages — perf_mode_l3_enhanced / 211303201213 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--aws_vpc_site--reference--group-003.md#canonical-2311120200212032-0110032223032010-0312130102100222-0111103310033201-3000130302120220-2230122332130032-2102330001223112-1011330031202310)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--aws_vpc_site--reference--group-003.md#canonical-2120022122200330-0123110211322022-0130032000012333-3331233320202011-0013320201223001-1220310112021303-1132230203210200-0213030022222322)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-2133010232202133-3021221010321021-2013032211210330-3301312300002132-3000011030113011-2311302233333003-2311302030320122-0211301220311201)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2311120200212032-0110032223032010-0312130102100222-0111103310033201-3000130302120220-2230122332130032-2102330001223112-1011330031202310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202321223102201-1223013233333123-3020131120132302-2122223032230013-0312222002331231-2221120300200313-2100023013321311-2222221002201311"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — jumbo / 230220223011 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-2133010232202133-3021221010321021-2013032211210330-3301312300002132-3000011030113011-2311302233333003-2311302030320122-0211301220311201)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-3120100220023123-3100100110001110-0011202001301022-2003210211212220-2002013312323022-1023132022021110-0322300003112103-3021100103202211)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-1212203211201222-1013313312303320-0000122330233310-1320333103201021-3010301132131223-1013020101230122-3333300100212011-1300301230012120"></a>

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
jumbo = {}
```

<a id="canonical-0013213301132032-0200102221312321-0330230303220120-1022213331021202-3202213102231003-0303032312201232-0313313020202112-3020213310002222"></a>

## Direct properties — jumbo / 230220223011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321221313130121-1332203112102103-1300312131322212-0222301312132313-3331220020103000-1132133210220031-0002112123212310-2003033333103211"></a>

## Next pages — jumbo / 230220223011 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-3120100220023123-3100100110001110-0011202001301022-2003210211212220-2002013312323022-1023132022021110-0322300003112103-3021100103202211)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2120022122200330-0123110211322022-0130032000012333-3331233320202011-0013320201223001-1220310112021303-1132230203210200-0213030022222322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122220130220001-3321332300230211-0032113301003201-3012133330100302-3323121101031031-2202300013120211-2111323231122033-1310320210033211"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — no_jumbo / 222111311222 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-2133010232202133-3021221010321021-2013032211210330-3301312300002132-3000011030113011-2311302233333003-2311302030320122-0211301220311201)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-3120100220023123-3100100110001110-0011202001301022-2003210211212220-2002013312323022-1023132022021110-0322300003112103-3021100103202211)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-3023310121033201-2221111012031201-1112231102320022-0110211123032131-0231112332330003-1202022013132023-0302102221212220-1130100323120222"></a>

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
no_jumbo = {}
```

<a id="canonical-0131303010130012-1121001330220303-1002131210121323-1131030130222131-1213310121332101-1212221013012230-2011100013031000-2030303300113232"></a>

## Direct properties — no_jumbo / 222111311222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321203022200023-3313320130123211-3112300031310012-0311130111100300-0203030133320133-0102033032012103-0021013111021331-2310322100110102"></a>

## Next pages — no_jumbo / 222111311222 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-3120100220023123-3100100110001110-0011202001301022-2003210211212220-2002013312323022-1023132022021110-0322300003112103-3021100103202211)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0232321100103301-3203312303221031-2122111132131303-0203300310312230-0221103131221221-2313301330021232-3200210201223130-2221111333111213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110321200012100-2233233002003320-0312311120213313-1010212331003020-1122103122121121-0033230202312231-0010201002023303-1322210221001313"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced — perf_mode_l7_enhanced / 221231212101 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-2133010232202133-3021221010321021-2013032211210330-3301312300002132-3000011030113011-2311302233333003-2311302030320122-0211301220311201)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-2101012121300333-3100303213101213-3113232130023123-0203322012022211-1302132211112201-1320200212312002-1012021211031323-1100302030313321"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo_disabled",
    "jumbo_enabled")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l7_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200101113310003-3212112233321023-2302100312222030-1333013223110121-1231023112101113-0012033023321322-3020332313032031-2122202301002211"></a>

## Direct properties — perf_mode_l7_enhanced / 221231212101 / 3

- [jumbo_disabled](resources--aws_vpc_site--reference--group-003.md#canonical-3210233113203322-1203010323202101-3322233310031233-1203323202110320-1122130101123021-0313332300331120-3300222231030002-1333013031113303): complete subsection reference.

- [jumbo_enabled](resources--aws_vpc_site--reference--group-003.md#canonical-2230130301112000-0201021023012010-3030031302330311-1221220120013321-0031311231301333-3323102130112122-2020120332211213-0032111232033312): complete subsection reference.

<a id="canonical-3320011332301010-3022211012320312-0213012202102202-0021210311222032-0222303023102131-2223112211330330-1333233333302310-1000130010312211"></a>

## Next pages — perf_mode_l7_enhanced / 221231212101 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--aws_vpc_site--reference--group-003.md#canonical-3210233113203322-1203010323202101-3322233310031233-1203323202110320-1122130101123021-0313332300331120-3300222231030002-1333013031113303)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--aws_vpc_site--reference--group-003.md#canonical-2230130301112000-0201021023012010-3030031302330311-1221220120013321-0031311231301333-3323102130112122-2020120332211213-0032111232033312)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-2133010232202133-3021221010321021-2013032211210330-3301312300002132-3000011030113011-2311302233333003-2311302030320122-0211301220311201)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3210233113203322-1203010323202101-3322233310031233-1203323202110320-1122130101123021-0313332300331120-3300222231030002-1333013031113303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310202031220331-2210301011230023-0121303101323012-3210333013032310-2213030033212100-3320311203212202-1333303300233131-0103311200033223"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — jumbo_disabled / 100201211002 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-2133010232202133-3021221010321021-2013032211210330-3301312300002132-3000011030113011-2311302233333003-2311302030320122-0211301220311201)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-0232321100103301-3203312303221031-2122111132131303-0203300310312230-0221103131221221-2313301330021232-3200210201223130-2221111333111213)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-2012131113022302-1130301202212012-2323221120112033-0331232311123332-1203112121033212-1033223032010223-3112112210031302-1102231203102003"></a>

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
jumbo_disabled = {}
```

<a id="canonical-0303211123002222-3313212211000230-0333230200233230-1313312120222110-2122132303012123-0023003303320130-3323213011001313-3122223212121323"></a>

## Direct properties — jumbo_disabled / 100201211002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322203122030330-3321303003312103-2131201300132133-0122022201013333-2003300320032230-1231030220221220-1300130302021002-0113121200133131"></a>

## Next pages — jumbo_disabled / 100201211002 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-0232321100103301-3203312303221031-2122111132131303-0203300310312230-0221103131221221-2313301330021232-3200210201223130-2221111333111213)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2230130301112000-0201021023012010-3030031302330311-1221220120013321-0031311231301333-3323102130112122-2020120332211213-0032111232033312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231122210112311-2121002030210130-2300130221010203-1223230300203232-3013133221320230-2103112101202202-1031202321120122-2221310330213133"></a>

## ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — jumbo_enabled / 113022322030 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-2133010232202133-3021221010321021-2013032211210330-3301312300002132-3000011030113011-2311302233333003-2311302030320122-0211301220311201)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-0232321100103301-3203312303221031-2122111132131303-0203300310312230-0221103131221221-2313301330021232-3200210201223130-2221111333111213)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-0303232330123302-0210311322222022-2331223230312301-0301000333001230-0001122032100212-2112232233023211-0301123201210322-2110323333120210"></a>

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
jumbo_enabled = {}
```

<a id="canonical-0003232131330203-2100110222200322-0012320213101121-1121012332022010-2131220013300212-2320230013232003-1102121210030230-2100300313302310"></a>

## Direct properties — jumbo_enabled / 113022322030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212123303312113-3022110133030213-3123013310012103-0021233222330031-1332131023213100-0322233233120133-2330001213223232-2121103121120121"></a>

## Next pages — jumbo_enabled / 113022322030 / 4

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-0232321100103301-3203312303221031-2122111132131303-0203300310312230-0221103131221221-2313301330021232-3200210201223130-2221111333111213)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2002101132330123-2132301310023301-2122003302320220-0032010113032131-2000321312020231-2331331002032220-0312033010032022-3001123020212102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311113113232122-2210321310101101-0112300310222210-3331333212122231-3221103022311002-0112112100101103-1013130010033003-3000211313310132"></a>

## ingress_egress_gw.sm_connection_public_ip — sm_connection_public_ip / 111102131001 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.sm_connection_public_ip

<a id="canonical-1033100121131301-2310222101030100-1221302231012301-2311201210222023-3221213230003131-0222223011023203-1220133201200000-3222201031320003"></a>

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

<a id="canonical-0312322202133120-3113230110233010-3322010201021223-2113213320101110-1110023103333321-1211002201223221-3200011110003013-3230111221332301"></a>

## Direct properties — sm_connection_public_ip / 111102131001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323121100223231-0130000221012330-3003323122302112-2332220303313121-1310201000202030-1323123121101111-1132120032033033-0113013103212302"></a>

## Next pages — sm_connection_public_ip / 111102131001 / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2022002132013302-3133212203112321-3201130121023311-1321200003012130-3121321120003222-0120032003022210-3110123100000102-3112320312011021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233032212121001-0200301030101321-0032012033323021-3023212300031021-2200102313220221-2322220020223133-1012030210013110-2213001112021111"></a>

## ingress_egress_gw.sm_connection_pvt_ip — sm_connection_pvt_ip / 333102011101 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.sm_connection_pvt_ip

<a id="canonical-3320332030031303-2001133233031333-2102032102000021-0023131223031122-3020232030130230-2132212100330232-2330200201030012-0222230112003110"></a>

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

<a id="canonical-1020310020201220-0233113221113330-0020031001312303-3222313213121100-3101013022030310-3212112131111200-3020220003000301-2300310013021220"></a>

## Direct properties — sm_connection_pvt_ip / 333102011101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323110322100100-2202010122312302-0102200022220033-3001210130122100-2211310113212130-1221333323221101-3213222131231200-0021101331113032"></a>

## Next pages — sm_connection_pvt_ip / 333102011101 / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222312122033202-0112031033220310-2201202111331032-3013222321210121-1001313323320323-3001033022002330-0333213333303313-1200223211333220"></a>

## ingress_gw — ingress_gw / 132021220030 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- ingress_gw

<a id="canonical-1213030012102200-2132122202121220-2010120201212311-1321031001321231-3302202232203113-2220103102300303-3322131013213033-1113012313110132"></a>

Type: `"object"`. single nested block, Optional.

AWS Ingress Gateway. Single interface AWS ingress site.

Upstream description:

Single interface AWS ingress site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_certified_hw",
    "az_nodes")}
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
ingress_gw {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103012101131211-3100103312110010-3011323102302321-1232211103322301-0322321210303332-3020132333030013-1201221321002220-0303020103231003"></a>

## Direct properties — ingress_gw / 132021220030 / 3

- [allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-2202312332212100-1031132303122122-1232023011302122-2222323223002320-2301112220203003-1311230220303301-1223300032132002-0201022012133120): complete subsection reference.

<a id="canonical-0121232011233123-1311031102022212-3110300113222212-1122233132213201-0031110231313021-0103223102300312-0011300112310320-2130001131311123"></a>

<a id="canonical-1202110112333221-1232022223220002-2202230312333013-2220313022122231-3220220033201212-0311322131321212-1121132231012332-3323321111321210"></a>

## aws_certified_hw property — ingress_gw / 132021220030 / 4

Type: `"string"`. Optional.

\[Enum: aws-byol-voltmesh\] AWS Certified Hardware. Name for AWS certified hardware. The only
possible value is \`aws-byol-voltmesh\`.

Upstream description:

Name for AWS certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("aws-byol-voltmesh"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "aws-byol-voltmesh"
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
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-2330123030021300-0333222331002031-2122221322133023-1323322003320020-1210003321202232-0102331330021201-0022103012213020-0103331133211210): complete subsection reference.

- [performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-3113111133121313-2001031200031030-3223100231113023-1120200011200033-1112201030331232-0320220033031131-0201132212302222-3102332030132220): complete subsection reference.
