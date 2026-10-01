---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3020313233311232-3003032232302000-1213212013330321-1110123310111132-3021300022033010-3122201103113101-1131022310201232-1102223001002212"></a>

## command property — container / 233033011010 / 5

Type: `["list", "string"]`. Optional.

Command to execute. Overrides the Docker image's ENTRYPOINT.

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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

- [custom_flavor](resources--workload--reference--group-017.md#canonical-1101122001010203-0213022111133112-2022320121302331-0210332313220111-1122010030301033-1213212030220200-2311111310312301-2202210102321311): complete subsection reference.

- [default_flavor](resources--workload--reference--group-017.md#canonical-1021322231332022-0200122020303230-1103121113233003-0303121302202231-0113031111122210-3022123231231321-3331113133012000-1201233031220130): complete subsection reference.

<a id="canonical-0232133120123001-3000001100013103-3321303323200310-1113302213220102-0202212323033013-1312313131211112-3231031133132202-1333123301332020"></a>

<a id="canonical-2331121112213100-2220322210101321-3010210131102030-3231101002003302-3132211231123221-3320332203221111-1310330123313322-1220213320122321"></a>

## flavor property — container / 233033011010 / 6

Type: `"string"`. Optional.

\[Enum:
CONTAINER\_FLAVOR\_TYPE\_TINY|CONTAINER\_FLAVOR\_TYPE\_MEDIUM|CONTAINER\_FLAVOR\_TYPE\_LARGE\]
Container Flavor type - CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny Tiny containers have limit of 0.1 vCPU
and 256 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium Medium containers have limit
of 0.25 vCPU and 512 MiB (mebibyte) memory - CONTAINER\_FLAVOR\_TYPE\_LARGE: Large Large containers
have.. Possible values are \`CONTAINER\_FLAVOR\_TYPE\_TINY\`, \`CONTAINER\_FLAVOR\_TYPE\_MEDIUM\`,
\`CONTAINER\_FLAVOR\_TYPE\_LARGE\`. Defaults to \`CONTAINER\_FLAVOR\_TYPE\_TINY\`.

Upstream description:

Container Flavor type

&#8203;- CONTAINER\_FLAVOR\_TYPE\_TINY: Tiny

Tiny containers have limit of 0.1 vCPU and 256 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_MEDIUM: Medium

Medium containers have limit of 0.25 vCPU and 512 MiB (mebibyte) memory &#8203;-
CONTAINER\_FLAVOR\_TYPE\_LARGE: Large

Large containers have limit of 1 vCPU and 2048 MiB (mebibyte) memory.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTAINER_FLAVOR_TYPE_TINY",
    "CONTAINER_FLAVOR_TYPE_MEDIUM",
    "CONTAINER_FLAVOR_TYPE_LARGE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTAINER_FLAVOR_TYPE_TINY",
  "enum": [
    "CONTAINER_FLAVOR_TYPE_TINY",
    "CONTAINER_FLAVOR_TYPE_MEDIUM",
    "CONTAINER_FLAVOR_TYPE_LARGE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [image](resources--workload--reference--group-017.md#canonical-1031320200322221-3002122102301132-0221210331223313-1010110310231212-1213311212023030-0232310322120233-0311232211231133-3220301303313013): complete subsection reference.

<a id="canonical-2202013202312211-3221210010311123-1320130103202201-1220032201122102-3020201220212220-0023200133230011-0323110323133312-1212022122101330"></a>

<a id="canonical-0220000013021212-0021101101313033-2232212010120120-3131222202233101-0121231011231122-1000121110332002-2113310200132012-1330212333122110"></a>

## init_container property — container / 233033011010 / 7

Type: `"bool"`. Optional.

Specialized container that runs before application container and runs to completion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [liveness_check](resources--workload--reference--group-017.md#canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112): complete subsection reference.

<a id="canonical-2030101133223013-2330003002033321-1023321232213223-0330131033202202-0200310111311312-3031111320011023-0131212102003130-0032033131013033"></a>

<a id="canonical-1210033221130131-0012001200222312-2210312322130233-0233332110211332-3330013332300130-3122112301100022-0220210311030303-0133103320232220"></a>

## name property — container / 233033011010 / 8

Type: `"string"`. Optional.

Name. Name of the container.

Upstream description:

Name of the container.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [readiness_check](resources--workload--reference--group-017.md#canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323): complete subsection reference.

<a id="canonical-3031332231200102-1231332331012132-2303021011112321-1212320232200212-2332000011230103-2211100330020131-2211212011230003-0211023311201232"></a>

## Next pages — container / 233033011010 / 9

- [simple_service.container.custom_flavor](resources--workload--reference--group-017.md#canonical-1101122001010203-0213022111133112-2022320121302331-0210332313220111-1122010030301033-1213212030220200-2311111310312301-2202210102321311)
- [simple_service.container.default_flavor](resources--workload--reference--group-017.md#canonical-1021322231332022-0200122020303230-1103121113233003-0303121302202231-0113031111122210-3022123231231321-3331113133012000-1201233031220130)
- [simple_service.container.image](resources--workload--reference--group-017.md#canonical-1031320200322221-3002122102301132-0221210331223313-1010110310231212-1213311212023030-0232310322120233-0311232211231133-3220301303313013)
- [simple_service.container.liveness_check](resources--workload--reference--group-017.md#canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112)
- [simple_service.container.readiness_check](resources--workload--reference--group-017.md#canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1101122001010203-0213022111133112-2022320121302331-0210332313220111-1122010030301033-1213212030220200-2311111310312301-2202210102321311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310223003001321-0020230331022332-0220333101213221-2220102020002301-1322302312223102-2100231321010331-1330021121121031-3132221010300110"></a>

## simple_service.container.custom_flavor — custom_flavor / 330122012112 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- simple_service.container.custom_flavor

<a id="canonical-0231332000323331-0232001121003230-0213101031233131-2012030002323113-1030023121210121-0110320332210231-2212030133303110-3013310020232113"></a>

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
custom_flavor {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220102012100031-2221320210122223-2000123212103120-1000131130001133-3230032221220223-2131101001231222-1330331120130223-3102121032003201"></a>

## Direct properties — custom_flavor / 330122012112 / 3

<a id="canonical-3302322202313131-2231021121010320-1220003221001332-3010313112322221-2001302123111132-2030231321221222-1313330132202133-1300021131203321"></a>

<a id="canonical-3200123313220312-2331011132313230-3111021301130020-3002200002122231-2112200031313312-0330030113002010-3012023003301031-0010002221311330"></a>

## name property — custom_flavor / 330122012112 / 4

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

<a id="canonical-2101112003302230-0300123230020103-1111322323103302-3203123013110230-3102303322332001-0202203233011232-2010100302013301-0030101323120100"></a>

<a id="canonical-0201112202120023-1031030033220231-0332303022001321-1133222000113021-1023101231202312-0033302211203122-2113111212030003-2111331230221111"></a>

## namespace property — custom_flavor / 330122012112 / 5

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

<a id="canonical-0020220123010030-3201311203200111-3200113033123313-0011201123312003-0323212312231000-0101030011003232-2013001321031012-0123021220123232"></a>

<a id="canonical-1032003010113113-3323122122012312-1022112121222033-1113011221131233-1112100200010300-2110201131230132-0001221331312232-2222033220031212"></a>

## tenant property — custom_flavor / 330122012112 / 6

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

<a id="canonical-2133100230333010-0302012003102010-2113301102011322-3012203110220111-0023021102000303-1211203223320210-3103100202013100-2311330033311010"></a>

## Next pages — custom_flavor / 330122012112 / 7

- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1021322231332022-0200122020303230-1103121113233003-0303121302202231-0113031111122210-3022123231231321-3331113133012000-1201233031220130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030000311333122-3221001233302031-0122013222232202-1233211011001330-3033221001311320-3221302123121210-0302032222211221-2312021323220121"></a>

## simple_service.container.default_flavor — default_flavor / 121112303312 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- simple_service.container.default_flavor

<a id="canonical-3021132000011123-3113222221321033-1121321300311302-0220331332233301-2123010013212010-1333213021222221-3120110033310302-3002132102130130"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default flavor.

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
default_flavor = {}
```

<a id="canonical-3002232031320301-2302100101021100-3221133023230022-0230203303032132-3120102112131331-2230330103320030-2302123102231003-0130131031132120"></a>

## Direct properties — default_flavor / 121112303312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320303020212230-1232231133002210-0311223223230220-1102312203130333-1302132303201223-3330332101120201-0301000330223133-2022113021121231"></a>

## Next pages — default_flavor / 121112303312 / 4

- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1031320200322221-3002122102301132-0221210331223313-1010110310231212-1213311212023030-0232310322120233-0311232211231133-3220301303313013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233101213131113-0121203023320003-1131231203213220-0002121133122000-2000313213112311-2103131013002310-0301223321213101-0103312313303000"></a>

## simple_service.container.image — image / 113122102013 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- simple_service.container.image

<a id="canonical-1112332221131201-2021312133010013-3003121030022100-2101300231213223-2110203022112330-0301010102222201-2113011322302120-1020022311110130"></a>

Type: `"object"`. single nested block, Optional.

ImageType configures the image to use, how to pull the image, and the associated secrets to use if
any.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name"),
  validators.ConflictingObjectAttributes("container_registry",
    "public")}
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
  "x-ves-oneof-field-registry_choice": "[\"container_registry\",\"public\"]"
}
```

Terraform syntax:

```terraform
image {
  # Configure direct properties listed below.
}
```

<a id="canonical-0123123320332333-0221102102211220-1113123321132101-0100112100113322-1310002230100233-2323130011001323-3112103221100331-0003031001230222"></a>

## Direct properties — image / 113122102013 / 3

- [container_registry](resources--workload--reference--group-017.md#canonical-1200331112133231-3310003010123012-1110023010303211-2311012310220121-3021122231012000-1211101013332022-1003223220212311-2101321331222021): complete subsection reference.

<a id="canonical-3312301013333322-2011003320302201-2202132121203111-1022211130322201-2213200101012301-0100322131020203-3021200023033001-2031302323212232"></a>

<a id="canonical-0300223112120320-3201023111231223-0113001132030023-3301223210113332-3010212210123311-3203132100320210-3013010312321301-1120123032310220"></a>

## name property — image / 113122102013 / 4

Type: `"string"`. Optional.

Name is a container image which are usually given a name such as alpine, ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed.

Upstream description:

Name is a container image which are usually given a name such as alpine, ubuntu, or
quay.I/O/etcd:0.13. The format is registry/image:tag or registry/image@image-digest. If registry is
not specified, the Docker public registry is assumed. If tag is not specified, latest is assumed.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [public](resources--workload--reference--group-017.md#canonical-3011103200100232-3221201210120303-2332112223300201-3020111020011211-1313113233020102-3330003212121100-1100202011112101-1131031023231022): complete subsection reference.

<a id="canonical-0312202002101221-0303231031211302-3111213233122332-0222331323111210-2320232220010000-3030000323202111-2233303221123312-3232031201100330"></a>

<a id="canonical-3030113302121030-1112003233010102-0202022300233313-2320031030202031-3013132132231313-3123111201011113-0010133100012110-1031032021102302"></a>

## pull_policy property — image / 113122102013 / 5

Type: `"string"`. Optional.

\[Enum:
IMAGE\_PULL\_POLICY\_DEFAULT|IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT|IMAGE\_PULL\_POLICY\_ALWAYS|IMAGE\_PULL\_POLICY\_NEVER\]
Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload - IMAGE\_PULL\_POLICY\_DEFAULT: Default Default will always pull image if :latest tag
is specified in image name. If :latest tag is not specified in image name, it will pull image only..
Possible values are \`IMAGE\_PULL\_POLICY\_DEFAULT\`, \`IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT\`,
\`IMAGE\_PULL\_POLICY\_ALWAYS\`, \`IMAGE\_PULL\_POLICY\_NEVER\`. Defaults to
\`IMAGE\_PULL\_POLICY\_DEFAULT\`.

Upstream description:

Image pull policy type enumerates the policy choices to use for pulling the image prior to starting
the workload

&#8203;- IMAGE\_PULL\_POLICY\_DEFAULT: Default

Default will always pull image if :latest tag is specified in image name. If :latest tag is not
specified in image name, it will pull image only if it does not already exist on the node &#8203;-
IMAGE\_PULL\_POLICY\_IF\_NOT\_PRESENT: IfNotPresent

Only pull the image if it does not already exist on the node &#8203;- IMAGE\_PULL\_POLICY\_ALWAYS:
Always

Always pull the image &#8203;- IMAGE\_PULL\_POLICY\_NEVER: Never

Never pull the image.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("IMAGE_PULL_POLICY_DEFAULT",
    "IMAGE_PULL_POLICY_IF_NOT_PRESENT",
    "IMAGE_PULL_POLICY_ALWAYS",
    "IMAGE_PULL_POLICY_NEVER"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "IMAGE_PULL_POLICY_DEFAULT",
  "enum": [
    "IMAGE_PULL_POLICY_DEFAULT",
    "IMAGE_PULL_POLICY_IF_NOT_PRESENT",
    "IMAGE_PULL_POLICY_ALWAYS",
    "IMAGE_PULL_POLICY_NEVER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3231332320103003-0010100030100221-0033211223111302-3000230003333010-2320323000131200-2333233033312311-0123122133110012-1133201123223000"></a>

## Next pages — image / 113122102013 / 6

- [simple_service.container.image.container_registry](resources--workload--reference--group-017.md#canonical-1200331112133231-3310003010123012-1110023010303211-2311012310220121-3021122231012000-1211101013332022-1003223220212311-2101321331222021)
- [simple_service.container.image.public](resources--workload--reference--group-017.md#canonical-3011103200100232-3221201210120303-2332112223300201-3020111020011211-1313113233020102-3330003212121100-1100202011112101-1131031023231022)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1200331112133231-3310003010123012-1110023010303211-2311012310220121-3021122231012000-1211101013332022-1003223220212311-2101321331222021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220013320033321-0311300110131301-3113313133222310-0030221101323313-3023133311030220-3232322010210311-3130302210032111-3011010322120020"></a>

## simple_service.container.image.container_registry — container_registry / 003010020330 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.image](resources--workload--reference--group-017.md#canonical-1031320200322221-3002122102301132-0221210331223313-1010110310231212-1213311212023030-0232310322120233-0311232211231133-3220301303313013)
- simple_service.container.image.container_registry

<a id="canonical-1220011200030120-2301130113121230-3030323311330210-2100331133303323-1123111012311002-0000301231001210-1033223332312010-3202103023133201"></a>

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
container_registry {
  # Configure direct properties listed below.
}
```

<a id="canonical-2321220200212011-2020032303010022-3330022113311313-0301330012233211-2330330132312130-1210120233120330-0002310001133133-1230330221231101"></a>

## Direct properties — container_registry / 003010020330 / 3

<a id="canonical-2312213330311131-2123032220303000-1002121223013033-1130202111300013-1130102110111303-0102312223333231-2222110311133221-1110333033011233"></a>

<a id="canonical-0322203132211232-3310311322031130-0303200212203012-1201233032101123-2020010102001021-3122103220320331-1101013331030202-1023312002201132"></a>

## name property — container_registry / 003010020330 / 4

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

<a id="canonical-0132223300020233-2110212110023321-0031022101200313-3131200331111330-3202320232120020-2121112032212022-3110331300130211-2210210023233100"></a>

<a id="canonical-0300230221013012-1012310332012103-2302011323102310-2012211103320223-0122210210203320-1103213333020301-1210121021212211-0010020210110123"></a>

## namespace property — container_registry / 003010020330 / 5

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

<a id="canonical-0330330222010330-3013032201003222-0323210201223313-0130033133300303-1112320330100202-1123112212111330-1021303111222021-3220213233203132"></a>

<a id="canonical-1223302313032321-2301301211131113-1130311300021121-0100113332202311-0102020331112333-2300103033203201-0311012120013100-0211011133230322"></a>

## tenant property — container_registry / 003010020330 / 6

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

<a id="canonical-0023300221303202-3012021031332132-0332310122102202-2321123113003232-1112013022102223-1112302332312300-0132131330030111-1000120001013020"></a>

## Next pages — container_registry / 003010020330 / 7

- [simple_service.container.image](resources--workload--reference--group-017.md#canonical-1031320200322221-3002122102301132-0221210331223313-1010110310231212-1213311212023030-0232310322120233-0311232211231133-3220301303313013)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3011103200100232-3221201210120303-2332112223300201-3020111020011211-1313113233020102-3330003212121100-1100202011112101-1131031023231022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002321123011303-1232300100101230-0111222131310023-1303210312123232-3320123113023133-3200301121313233-1012201221123200-3222103022303223"></a>

## simple_service.container.image.public — public / 123310132120 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.image](resources--workload--reference--group-017.md#canonical-1031320200322221-3002122102301132-0221210331223313-1010110310231212-1213311212023030-0232310322120233-0311232211231133-3220301303313013)
- simple_service.container.image.public

<a id="canonical-1113223313210011-1320221031202012-2030301033121031-2020001320030322-3112001120303322-2332111331011030-3001103202033023-3102231333210111"></a>

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
public = {}
```

<a id="canonical-2013203001331112-3312220003302102-1103331023213323-2023213122131012-0112201230331001-1212310302310123-2013100301113010-0233122020333102"></a>

## Direct properties — public / 123310132120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100231131201000-1210030303302121-2122202300313201-3332122320120121-0221312023303332-0303132100302033-2203133020202201-2233103221333321"></a>

## Next pages — public / 123310132120 / 4

- [simple_service.container.image](resources--workload--reference--group-017.md#canonical-1031320200322221-3002122102301132-0221210331223313-1010110310231212-1213311212023030-0232310322120233-0311232211231133-3220301303313013)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020121122012002-1212121311130133-2011220120101123-0323131120213010-3230212002333302-3220330030101013-3312201132333331-2313100203203131"></a>

## simple_service.container.liveness_check — liveness_check / 210333000311 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- simple_service.container.liveness_check

<a id="canonical-1011300101022230-3003203220333333-0203030211212102-2213201331300233-3000011330103313-0122222001012300-3313122211120231-1132231303103200"></a>

Type: `"object"`. single nested block, Optional.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "http_health_check"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "tcp_health_check"),
  validators.ConflictingObjectAttributes("http_health_check",
    "tcp_health_check")}
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
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

Terraform syntax:

```terraform
liveness_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-1022212312322223-1130123100200322-3303130110010100-3313122303121302-2302102111231331-3330002301012301-3213300232203321-0221122333120112"></a>

## Direct properties — liveness_check / 210333000311 / 3

- [exec_health_check](resources--workload--reference--group-017.md#canonical-1000233212122031-3123100131313301-3002121002223331-3013313223310031-1103132102110303-1102110103010010-1001300002102322-1211131313112020): complete subsection reference.

<a id="canonical-1210303210210322-1210013131201033-2303010122220331-2321132003312303-3311222300000103-2233210021220113-1111113122022100-2332310120301303"></a>

<a id="canonical-3201222121003111-0202121311130220-2310013023120133-3122010222221113-1032032122002113-0230102222012302-2221301201133003-0223221332010102"></a>

## healthy_threshold property — liveness_check / 210333000311 / 4

Type: `"number"`. Optional.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](resources--workload--reference--group-017.md#canonical-0123112211202011-1011003322102323-3231201102302033-2302101301120031-3120122012122032-3212311301223311-2220133131010031-0003200300112301): complete subsection reference.

<a id="canonical-0212212103321330-0332111033203130-2020012210030310-0130221000313120-0030110233232201-1032230203330022-1020032010322312-2031012111002312"></a>

<a id="canonical-0333011311022202-1330103031131001-1313132301223030-1112210201303010-3313311100220000-0210132313122212-2232302001133020-2220112210101122"></a>

## initial_delay property — liveness_check / 210333000311 / 5

Type: `"number"`. Optional.

Number of seconds after the container has started before health checks are initiated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-1210202100003123-1210310101322012-2231303131200113-2331013111122130-1211201200232020-2211200310322312-0332113220310221-1222111232211011"></a>

<a id="canonical-1311221012202023-0212310122211323-1103332220200120-0201113302122002-1121122030100030-3110311123033230-0233333011121100-1212330322321033"></a>

## interval property — liveness_check / 210333000311 / 6

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [tcp_health_check](resources--workload--reference--group-017.md#canonical-1321011121310301-2221231110312101-2203132301320030-3030201110120231-3002002032122030-1130100002311331-3130212221120312-1230122223333111): complete subsection reference.

<a id="canonical-1112203121121123-0000033231010203-1002022132333133-0222102200231323-0033103303131102-3010223111203320-0203022311001212-0303311232321330"></a>

<a id="canonical-2020133331011030-0011201003321111-2110223212310020-0112110020332210-2332301222033301-1312210211312032-0130321002202131-2233022313033130"></a>

## timeout property — liveness_check / 210333000311 / 7

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-3122222111113102-3222330003302223-0121002121200333-2313021032233032-1120221123310202-0103330331203020-3302232012023331-1011220033211212"></a>

<a id="canonical-2223123012232010-1323313132101232-0101121321233120-1023000033302110-3010201222110131-0223330100111211-3132133121333013-0020300132100311"></a>

## unhealthy_threshold property — liveness_check / 210333000311 / 8

Type: `"number"`. Optional.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-0201223202122111-1333220022302003-2100213020312210-3102001001320202-3301022323131230-2200300223102212-0103211320003122-3211321202003102"></a>

## Next pages — liveness_check / 210333000311 / 9

- [simple_service.container.liveness_check.exec_health_check](resources--workload--reference--group-017.md#canonical-1000233212122031-3123100131313301-3002121002223331-3013313223310031-1103132102110303-1102110103010010-1001300002102322-1211131313112020)
- [simple_service.container.liveness_check.http_health_check](resources--workload--reference--group-017.md#canonical-0123112211202011-1011003322102323-3231201102302033-2302101301120031-3120122012122032-3212311301223311-2220133131010031-0003200300112301)
- [simple_service.container.liveness_check.tcp_health_check](resources--workload--reference--group-017.md#canonical-1321011121310301-2221231110312101-2203132301320030-3030201110120231-3002002032122030-1130100002311331-3130212221120312-1230122223333111)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1000233212122031-3123100131313301-3002121002223331-3013313223310031-1103132102110303-1102110103010010-1001300002102322-1211131313112020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223322321320102-2120120321023011-1310323013030103-3031200212032211-1133200311121202-0113211203011033-2332302333220333-1033020022313011"></a>

## simple_service.container.liveness_check.exec_health_check — exec_health_check / 021311330223 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.liveness_check](resources--workload--reference--group-017.md#canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112)
- simple_service.container.liveness_check.exec_health_check

<a id="canonical-0300111010113133-0300210221221301-1011202110132121-0301023133233323-3302132323221031-0223301220111300-2213303211221201-1023200301033330"></a>

Type: `"object"`. single nested block, Optional.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("command")}
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
exec_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230213102220122-3123130013231303-0330131331011123-1020210130021020-2030211000232302-0200303033120202-3232302313102230-0302332021303013"></a>

## Direct properties — exec_health_check / 021311330223 / 3

<a id="canonical-2001330122313001-1120332211030330-0221301012230112-0330333203200332-3331223031100201-1023101321132232-1200300003230221-1323031021103000"></a>

<a id="canonical-2122320030122213-2330322300301113-3100123020302013-3110032111101132-0311311221001002-3220222222213312-2200223000030230-3302333223100213"></a>

## command property — exec_health_check / 021311330223 / 4

Type: `["list", "string"]`. Optional.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3002010132003132-3012331213230120-1131233013330032-2303232222130023-0113201012012203-3203332221103110-1333330111001230-2220211110102110"></a>

## Next pages — exec_health_check / 021311330223 / 5

- [simple_service.container.liveness_check](resources--workload--reference--group-017.md#canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0123112211202011-1011003322102323-3231201102302033-2302101301120031-3120122012122032-3212311301223311-2220133131010031-0003200300112301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103323001211213-0312332031123030-1013103031110020-2033132112033132-0103111013301102-2130130100311333-1110012130131003-1023210202300310"></a>

## simple_service.container.liveness_check.http_health_check — http_health_check / 012020322201 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.liveness_check](resources--workload--reference--group-017.md#canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112)
- simple_service.container.liveness_check.http_health_check

<a id="canonical-1221201200330113-2312001202323230-3010130032113001-3011313001310013-0332221221100000-2233123121313033-0011033311331330-1132300003122133"></a>

Type: `"object"`. single nested block, Optional.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200333303302210-1101011312333201-2221003330332123-3010232020133102-1232001233112201-1313103133310301-1220211130202110-2203200203302013"></a>

## Direct properties — http_health_check / 012020322201 / 3

<a id="canonical-0300331312003222-3021230020002323-2213321020331212-1323112330131020-3201000212021120-0332023300311003-2032332130112101-0222211000103000"></a>

<a id="canonical-3111113121132333-1021311033133323-1110001101122130-0230312312303010-1310302212101021-0331001131212213-3310003011021233-0003332100123330"></a>

## headers property — http_health_check / 012020322201 / 4

Type: `["map", "string"]`. Optional.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

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
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-0021112000311311-1013231323200233-1032331023211203-1313133220220300-0130203311102303-1301122010202112-3331103311121221-0323201320113223"></a>

<a id="canonical-3230002022102233-1002323203113333-1101200233020133-1312303322010231-1301300032300333-1030210223010301-0030223212032311-3111313321000103"></a>

## host_header property — http_health_check / 012020322201 / 5

Type: `"string"`. Optional.

The value of the host header in the HTTP health check request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-3202222300330111-2233031313323300-1030203300110011-3123130123121323-1121312100122110-2021323222231003-2233231133300011-1322111213100231"></a>

<a id="canonical-0220203021333122-3122203020111131-1303013323113033-1322122330311230-3020111301303130-1030202023310333-1302023133101113-1131322010231231"></a>

## path property — http_health_check / 012020322201 / 6

Type: `"string"`. Optional.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

- [port](resources--workload--reference--group-017.md#canonical-0332012012023132-1213232202130103-1201221221101131-2300330101111132-1002201203020001-3203321221133102-3123333133031123-2120302130201231): complete subsection reference.

<a id="canonical-1233011312312032-3132310302101212-2210121020020230-0230201012320101-1003020133301112-2231322232223330-3310101332033031-1223120130102123"></a>

## Next pages — http_health_check / 012020322201 / 7

- [simple_service.container.liveness_check.http_health_check.port](resources--workload--reference--group-017.md#canonical-0332012012023132-1213232202130103-1201221221101131-2300330101111132-1002201203020001-3203321221133102-3123333133031123-2120302130201231)
- [simple_service.container.liveness_check](resources--workload--reference--group-017.md#canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0332012012023132-1213232202130103-1201221221101131-2300330101111132-1002201203020001-3203321221133102-3123333133031123-2120302130201231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021033011303120-3110201001200211-3223322031230321-3321101113213031-0332231233332002-3300230311011001-0211132123303012-2232311132201033"></a>

## simple_service.container.liveness_check.http_health_check.port — port / 323130032130 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.liveness_check](resources--workload--reference--group-017.md#canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112)
- [simple_service.container.liveness_check.http_health_check](resources--workload--reference--group-017.md#canonical-0123112211202011-1011003322102323-3231201102302033-2302101301120031-3120122012122032-3212311301223311-2220133131010031-0003200300112301)
- simple_service.container.liveness_check.http_health_check.port

<a id="canonical-0331301120110322-0223002023311331-0022013120103100-2033123121220211-2212100220332220-0311231223321321-2320121101111123-0003031230311132"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312221121223112-0202220202332303-3130003112132333-1310102102110023-1112321021113013-3100220103211230-3011333000021223-0231322223322312"></a>

## Direct properties — port / 323130032130 / 3

<a id="canonical-0022330232302330-1231202320101031-1312322212222031-2230311111211331-0022102221210300-1223023233102112-2231000113211233-0110302123133132"></a>

<a id="canonical-0033132210210320-0222111013212212-0331301132030121-2320101320203222-2332013313213123-0101133030210133-0133101123012101-0233021100313113"></a>

## name property — port / 323130032130 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-3101220022330013-2002020302213230-2312112000130112-2330332102003102-3230122201212122-3011200012002033-3222122120112132-0013211133010220"></a>

<a id="canonical-3031012133131110-3030210010122213-1231231022132320-1120200311311310-3010013212110023-1321113320031121-2330330012330000-3030311003131013"></a>

## num property — port / 323130032130 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2123300023310312-3130013221100022-2201130122211102-1232312221301000-2212123023212130-1132200322320311-2031013021121012-2222112222203323"></a>

## Next pages — port / 323130032130 / 6

- [simple_service.container.liveness_check.http_health_check](resources--workload--reference--group-017.md#canonical-0123112211202011-1011003322102323-3231201102302033-2302101301120031-3120122012122032-3212311301223311-2220133131010031-0003200300112301)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1321011121310301-2221231110312101-2203132301320030-3030201110120231-3002002032122030-1130100002311331-3130212221120312-1230122223333111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223010033200012-1201001000122031-3022030210320022-3312221313303111-1100213020132231-2302301322000033-3312213321131012-2301322203300121"></a>

## simple_service.container.liveness_check.tcp_health_check — tcp_health_check / 320232222310 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.liveness_check](resources--workload--reference--group-017.md#canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112)
- simple_service.container.liveness_check.tcp_health_check

<a id="canonical-1313220130101003-0202110231310202-0120011102033103-1221331221210011-1320000232303112-3120100002113112-3320002201131230-2102120213301000"></a>

Type: `"object"`. single nested block, Optional.

TCPHealthCheckType describes a health check based on opening a TCP connection.

Receipt-pinned upstream constraints:

```json
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
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223322212313133-1331022322002102-0320332002201132-1211320331110221-1320011012212311-0113221323021011-2020223301203232-0120011101011210"></a>

## Direct properties — tcp_health_check / 320232222310 / 3

- [port](resources--workload--reference--group-017.md#canonical-0032202331212122-1213122130100333-0123002331311201-0021322022122110-0103132013120120-0120333032213021-0231331110003102-3111031301213030): complete subsection reference.

<a id="canonical-3332123311101101-0302320032232022-1000222033021101-3333032000003302-3211020110011013-0131233211333302-1021310313311013-2333000013313212"></a>

## Next pages — tcp_health_check / 320232222310 / 4

- [simple_service.container.liveness_check.tcp_health_check.port](resources--workload--reference--group-017.md#canonical-0032202331212122-1213122130100333-0123002331311201-0021322022122110-0103132013120120-0120333032213021-0231331110003102-3111031301213030)
- [simple_service.container.liveness_check](resources--workload--reference--group-017.md#canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0032202331212122-1213122130100333-0123002331311201-0021322022122110-0103132013120120-0120333032213021-0231331110003102-3111031301213030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203020230221100-1110232000212000-3032023003011233-3031133131112023-3030113330100320-0023102131100312-1200002022020222-1322022233202312"></a>

## simple_service.container.liveness_check.tcp_health_check.port — port / 031231322022 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.liveness_check](resources--workload--reference--group-017.md#canonical-1123003201223020-3031001003003030-2323130113230021-2110112120231030-0331101121132200-2010210232213121-3023312203232202-1112033223321112)
- [simple_service.container.liveness_check.tcp_health_check](resources--workload--reference--group-017.md#canonical-1321011121310301-2221231110312101-2203132301320030-3030201110120231-3002002032122030-1130100002311331-3130212221120312-1230122223333111)
- simple_service.container.liveness_check.tcp_health_check.port

<a id="canonical-1121112132311022-2301302113132321-3120111103101022-1022111331202020-2130201013220133-1003131321121211-3010303011003101-0131321120233222"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-3021222330331003-1030111220223320-2231333110011002-3220121302000202-3201101102131322-1103133113101120-1113022031322231-2013013121312121"></a>

## Direct properties — port / 031231322022 / 3

<a id="canonical-2233010232302311-1013122110312131-3133021330222021-2123223213231002-3323011102132230-0102333110211103-0223231210020222-2023023103302011"></a>

<a id="canonical-2102210301131232-2130320113000301-0132103332210222-0033012300132032-3312312130212232-3221100200222321-3230333032010221-3333330310302213"></a>

## name property — port / 031231322022 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-3013102001023303-3332031023033213-0213203133103320-1130233300221333-1021333102203032-1302011211200110-2321222221331013-0103333221110021"></a>

<a id="canonical-2211132232232012-2331211212000231-0323303312203312-1232200133013203-2310203211312011-3002322012233221-1211002320121131-3102131311222321"></a>

## num property — port / 031231322022 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1011021330310102-2102031022232133-0122112312131130-1010122331210202-2112023113113011-3221223111213222-0200120031111320-3311302130222111"></a>

## Next pages — port / 031231322022 / 6

- [simple_service.container.liveness_check.tcp_health_check](resources--workload--reference--group-017.md#canonical-1321011121310301-2221231110312101-2203132301320030-3030201110120231-3002002032122030-1130100002311331-3130212221120312-1230122223333111)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202222122101112-3110113113210121-3331020111103310-3223130101322223-1333011221120033-3110132103202331-3333011131211100-3200210121210221"></a>

## simple_service.container.readiness_check — readiness_check / 013020100012 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- simple_service.container.readiness_check

<a id="canonical-1312221123021031-3300121033033321-2130230221003102-0330131012201000-1020220321300000-0003030102322323-1300231233133020-2030223313132331"></a>

Type: `"object"`. single nested block, Optional.

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Upstream description:

HealthCheckType describes a health check to be performed against a container to determine whether it
has started up or is alive or ready to receive traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("healthy_threshold",
    "interval",
    "timeout",
    "unhealthy_threshold"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "http_health_check"),
  validators.ConflictingObjectAttributes("exec_health_check",
    "tcp_health_check"),
  validators.ConflictingObjectAttributes("http_health_check",
    "tcp_health_check")}
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
  "x-ves-oneof-field-health_check_choice": "[\"exec_health_check\",\"http_health_check\",\"tcp_health_check\"]"
}
```

Terraform syntax:

```terraform
readiness_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313213030033211-0311213332232332-0112303313100322-3220332120300013-2203030231231211-0320311323213001-1202022220332033-0013211331232200"></a>

## Direct properties — readiness_check / 013020100012 / 3

- [exec_health_check](resources--workload--reference--group-017.md#canonical-2232303312023003-3302232112300100-0133310100100211-3110312312010310-1032110113031121-1212102123203212-0231111010310002-2113113310202331): complete subsection reference.

<a id="canonical-1301011330112211-3010100133113033-1111100020233010-3311233123332000-2001100012312132-1122322221300020-0102030211113011-2320303223132132"></a>

<a id="canonical-1200303003012322-3221300132131303-0110211001022102-0011211113330231-3333333200001312-1303000212222010-1103031010132223-1002022300021102"></a>

## healthy_threshold property — readiness_check / 013020100012 / 4

Type: `"number"`. Optional.

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container..

Upstream description:

Number of consecutive successful responses after having failed before declaring healthy. In other
words, this is the number of healthy health checks required before marking healthy. Note that during
startup and liveliness, only a single successful health check is required to mark a container
healthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](resources--workload--reference--group-017.md#canonical-0013303313233031-1110230113303003-0321332113313013-2131211101210102-2303120323310310-3011231232223303-3031011103321102-1111122103321322): complete subsection reference.

<a id="canonical-3013011102213032-1111323122120123-2303231221302000-1032303101201233-1212003302332322-1310130203122223-0302333131123011-1230021330201020"></a>

<a id="canonical-2330112032203300-1130323333003223-0130203302100331-1302201203222213-1322210233131001-0002100012311011-1221300031332201-3030320210301031"></a>

## initial_delay property — readiness_check / 013020100012 / 5

Type: `"number"`. Optional.

Number of seconds after the container has started before health checks are initiated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-3013313331213202-3003212003301200-2222103320012222-1203303122322132-3321311100210221-0221202030231103-1011230132030332-2333131020312320"></a>

<a id="canonical-2122202331211013-1023311133033113-3012330020312120-0012031131323301-3013200212111323-1310200210323200-2221033023020203-2020101120000012"></a>

## interval property — readiness_check / 013020100012 / 6

Type: `"number"`. Optional.

Time interval in seconds between two health check requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

- [tcp_health_check](resources--workload--reference--group-017.md#canonical-2120132203003331-2020222300211130-1230021011312132-1001012000221222-1131332001122301-3000222103312120-1010001332230132-0210313102002011): complete subsection reference.

<a id="canonical-3101320013012203-2113023323323131-0001201201101031-3332330323013021-2101011121320302-3033203030333100-1033112012312010-2031303321022013"></a>

<a id="canonical-0302312100212213-0031130323232010-0101023121202003-0022123023130220-2113023130020212-3000310111331100-1330231312101013-0131101303231320"></a>

## timeout property — readiness_check / 013020100012 / 7

Type: `"number"`. Optional.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Upstream description:

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 600),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
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
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-1221311032010002-2221310312112211-3233233120330112-1310232020230021-0223212300332022-1211331332323222-0201130033111131-3033202022010301"></a>

<a id="canonical-2111122333103233-2201113333203113-0022222302033313-2220012222200032-2133331013223330-3120330210130132-1120120223130230-2010000020131010"></a>

## unhealthy_threshold property — readiness_check / 013020100012 / 8

Type: `"number"`. Optional.

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Upstream description:

Number of consecutive failed responses before declaring unhealthy. In other words, this is the
number of unhealthy health checks required before a container is marked unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-0132331003120131-3223023220122333-2123322030302232-1312010022333022-2003110321000223-2330001123332131-0331310101311102-2000320230023211"></a>

## Next pages — readiness_check / 013020100012 / 9

- [simple_service.container.readiness_check.exec_health_check](resources--workload--reference--group-017.md#canonical-2232303312023003-3302232112300100-0133310100100211-3110312312010310-1032110113031121-1212102123203212-0231111010310002-2113113310202331)
- [simple_service.container.readiness_check.http_health_check](resources--workload--reference--group-017.md#canonical-0013303313233031-1110230113303003-0321332113313013-2131211101210102-2303120323310310-3011231232223303-3031011103321102-1111122103321322)
- [simple_service.container.readiness_check.tcp_health_check](resources--workload--reference--group-017.md#canonical-2120132203003331-2020222300211130-1230021011312132-1001012000221222-1131332001122301-3000222103312120-1010001332230132-0210313102002011)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2232303312023003-3302232112300100-0133310100100211-3110312312010310-1032110113031121-1212102123203212-0231111010310002-2113113310202331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233021120213121-1110213230322123-3013110212101031-3121230230121030-0030003231312203-1010002210123333-3120312212003001-3133320023332110"></a>

## simple_service.container.readiness_check.exec_health_check — exec_health_check / 333030200321 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.readiness_check](resources--workload--reference--group-017.md#canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323)
- simple_service.container.readiness_check.exec_health_check

<a id="canonical-2111111330333303-1233103131022332-1121010023303033-1001010112123202-0222103033232113-0133000113212202-0023102003123203-2010222111201330"></a>

Type: `"object"`. single nested block, Optional.

ExecHealthCheckType describes a health check based on 'run in container' action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Upstream description:

ExecHealthCheckType describes a health check based on "run in container" action. Exit status of 0 is
treated as live/healthy and non-zero is unhealthy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("command")}
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
exec_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312311013220322-1031122330312333-1231002101102011-0213103211311121-2010200312113312-0221202123122031-3223203010012200-3023202020100033"></a>

## Direct properties — exec_health_check / 333030200321 / 3

<a id="canonical-2031221000000332-3113203013023220-0031312313301203-1300100013132111-1120231322012113-1231020132220122-2303031022313110-1100221332001230"></a>

<a id="canonical-1220013013211101-0310021103223120-0232200112122011-0310111303033001-0210330030313113-2300313320033103-0333131303120132-2000321221131022"></a>

## command property — exec_health_check / 333030200321 / 4

Type: `["list", "string"]`. Optional.

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to..

Upstream description:

Command is the command line to execute inside the container, the working directory for the command
is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a
shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to
explicitly call out to that shell.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0110201121213330-2122022113123330-3301010322021130-1130012111112133-1331032210020202-2112213221000003-2102203030303133-2031223321021211"></a>

## Next pages — exec_health_check / 333030200321 / 5

- [simple_service.container.readiness_check](resources--workload--reference--group-017.md#canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0013303313233031-1110230113303003-0321332113313013-2131211101210102-2303120323310310-3011231232223303-3031011103321102-1111122103321322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110203203020011-3101211111330110-2102330003112230-1230011120213132-1030213230313333-1310301211223022-0021012130111332-2330302333121030"></a>

## simple_service.container.readiness_check.http_health_check — http_health_check / 303131301211 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.readiness_check](resources--workload--reference--group-017.md#canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323)
- simple_service.container.readiness_check.http_health_check

<a id="canonical-3100220130333110-2110122001101133-3010203133332033-1323013203320123-1330222130203300-0123230111230311-1232210010333131-1020233323101023"></a>

Type: `"object"`. single nested block, Optional.

HTTPHealthCheckType describes a health check based on HTTP GET requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
http_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130011003020030-3022020002001323-2101102213112333-1020321300300122-3130213200233022-3213011330120032-1113301330222023-1203111103021213"></a>

## Direct properties — http_health_check / 303131301211 / 3

<a id="canonical-1112111131010302-1111330020230020-2201113233221332-0131310302212231-2132001010023030-1233000202212220-2003333330212121-3210000210300200"></a>

<a id="canonical-0332213133120312-3220022103133222-3300222130132222-2020222232312030-2212332232220222-1313022303322132-2033111230312330-3130010013330210"></a>

## headers property — http_health_check / 303131301211 / 4

Type: `["map", "string"]`. Optional.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

Upstream description:

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked container. This is a list of key-value pairs.

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
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-0220010322303023-3001213011212022-1222122122010300-2232301311113201-0323321111231101-0122332001323230-2333131313013020-0133020333000030"></a>

<a id="canonical-0321023112021012-0113131002100110-2302113000300023-1030300203002333-0300102022013331-3200211012101023-2112230033311030-1113012220022033"></a>

## host_header property — http_health_check / 303131301211 / 5

Type: `"string"`. Optional.

The value of the host header in the HTTP health check request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(262),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-1212032033330102-3023033303000331-2010132133201222-3033322231300202-0111330200213210-0210011011120312-3013212020103112-2300303131033020"></a>

<a id="canonical-1123012132133021-3322132313221132-1032232333023113-2011020110131330-2013121003210313-1110003231110030-3001133002112013-3201020031121100"></a>

## path property — http_health_check / 303131301211 / 6

Type: `"string"`. Optional.

Path. Path to access on the HTTP server.

Upstream description:

Path to access on the HTTP server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

- [port](resources--workload--reference--group-017.md#canonical-2100330323023130-0012320000123300-1222011220120312-0102120120210020-1210331021123232-3230110113032102-3310002103203322-3121033003300200): complete subsection reference.

<a id="canonical-1102213222122121-1113100222112303-1321011230111013-3222100100001122-0023230203121102-1232121002200001-2020112020313123-2213022021320001"></a>

## Next pages — http_health_check / 303131301211 / 7

- [simple_service.container.readiness_check.http_health_check.port](resources--workload--reference--group-017.md#canonical-2100330323023130-0012320000123300-1222011220120312-0102120120210020-1210331021123232-3230110113032102-3310002103203322-3121033003300200)
- [simple_service.container.readiness_check](resources--workload--reference--group-017.md#canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2100330323023130-0012320000123300-1222011220120312-0102120120210020-1210331021123232-3230110113032102-3310002103203322-3121033003300200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112222130002313-2320112112320233-2212232101013302-2121310012011121-2110202110100113-0131030110222123-3330332120203131-3033200000223033"></a>

## simple_service.container.readiness_check.http_health_check.port — port / 133312023003 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.readiness_check](resources--workload--reference--group-017.md#canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323)
- [simple_service.container.readiness_check.http_health_check](resources--workload--reference--group-017.md#canonical-0013303313233031-1110230113303003-0321332113313013-2131211101210102-2303120323310310-3011231232223303-3031011103321102-1111122103321322)
- simple_service.container.readiness_check.http_health_check.port

<a id="canonical-2213021231023100-1232023013022121-1221103122023122-2123022300122310-1000132301101021-1102030110113122-2332103303021200-0330013332312022"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313221320303111-0011121302221120-1133211102013120-0000332123002112-1331303010111321-1131321301310122-1332031310011000-2010223230113011"></a>

## Direct properties — port / 133312023003 / 3

<a id="canonical-1131002211102322-2020132311122320-0030021121330223-1332311230322232-1231031330132221-3021012323133011-1312333322122320-0023202122321012"></a>

<a id="canonical-1213113203122230-2003211111122102-0202102302010200-1333012031320001-2220010122023011-3220122112212101-1000313211333023-3113213023003213"></a>

## name property — port / 133312023003 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-2230013213132122-3330232212211133-1132000322030203-0321320322103200-3333133000312022-3201022201200121-1313023030000210-2332201322102202"></a>

<a id="canonical-0221202333202001-3301333131223323-3301203111311333-2123301031131012-1101113020012201-0200013121011000-3200220013133330-2203332212013221"></a>

## num property — port / 133312023003 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3021032012211011-2022021030121200-1111112223003032-3103110200302213-2332122213202301-1030111210022133-1320311223133002-2020010130223122"></a>

## Next pages — port / 133312023003 / 6

- [simple_service.container.readiness_check.http_health_check](resources--workload--reference--group-017.md#canonical-0013303313233031-1110230113303003-0321332113313013-2131211101210102-2303120323310310-3011231232223303-3031011103321102-1111122103321322)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2120132203003331-2020222300211130-1230021011312132-1001012000221222-1131332001122301-3000222103312120-1010001332230132-0210313102002011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232330222222012-1320313230103110-1013220220311232-2210333100301331-1330020202230010-0030130103112130-1321221002333202-2301003331211111"></a>

## simple_service.container.readiness_check.tcp_health_check — tcp_health_check / 321233113030 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.readiness_check](resources--workload--reference--group-017.md#canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323)
- simple_service.container.readiness_check.tcp_health_check

<a id="canonical-1113003001321000-1133203331120113-1213011131220030-2020202002203111-3133022323311202-3112202030211021-3300312030020020-3023120323321302"></a>

Type: `"object"`. single nested block, Optional.

TCPHealthCheckType describes a health check based on opening a TCP connection.

Receipt-pinned upstream constraints:

```json
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
tcp_health_check {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210230322231310-0203303201111231-1310301112333302-2200122333233311-2301100000322312-3001321123321303-1322211303332230-2021310121300213"></a>

## Direct properties — tcp_health_check / 321233113030 / 3

- [port](resources--workload--reference--group-017.md#canonical-2223220333130313-3102120110210000-3321133301022331-0213020030220300-3000100212221322-0231113001300020-1001031301313020-1113301333232231): complete subsection reference.

<a id="canonical-0233000201000003-1230032101100203-1330331032010230-3020121123222331-3002103321231302-2211002111101132-3131200221010311-2101030010301112"></a>

## Next pages — tcp_health_check / 321233113030 / 4

- [simple_service.container.readiness_check.tcp_health_check.port](resources--workload--reference--group-017.md#canonical-2223220333130313-3102120110210000-3321133301022331-0213020030220300-3000100212221322-0231113001300020-1001031301313020-1113301333232231)
- [simple_service.container.readiness_check](resources--workload--reference--group-017.md#canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2223220333130313-3102120110210000-3321133301022331-0213020030220300-3000100212221322-0231113001300020-1001031301313020-1113301333232231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111102103110230-3001222320311123-0211201131131301-3213201320102101-2223210020133010-0103210233322332-0300132103313023-3230010020102111"></a>

## simple_service.container.readiness_check.tcp_health_check.port — port / 012023102111 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.container](resources--workload--reference--group-016.md#canonical-1011303000022103-1122231110321130-3202121331100121-1301320113212120-3220331111111112-1120022220012313-3313013203321210-0332330312030213)
- [simple_service.container.readiness_check](resources--workload--reference--group-017.md#canonical-3212121030101233-3131221033310023-1233033330130102-3013330332203213-3033003011032322-1323201011111220-2202323123332313-3220300202001323)
- [simple_service.container.readiness_check.tcp_health_check](resources--workload--reference--group-017.md#canonical-2120132203003331-2020222300211130-1230021011312132-1001012000221222-1131332001122301-3000222103312120-1010001332230132-0210313102002011)
- simple_service.container.readiness_check.tcp_health_check.port

<a id="canonical-2032320211203132-2012020332331213-0310203122312311-1121103110322223-1001312222322112-0230112112312213-1212012300132321-1113110002013113"></a>

Type: `"object"`. single nested block, Optional.

Port. Port

Upstream description:

Port

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("name",
    "num")}
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
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

Terraform syntax:

```terraform
port {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130103231113222-0030220130203012-2121021221001333-3101030002120312-1223121303122133-3132321301322012-2201131001122220-1201311230002303"></a>

## Direct properties — port / 012023102111 / 3

<a id="canonical-1311210310003221-2032323310021102-1010232022121010-3001311313330312-3121023331230112-1321031202110123-2231210200022331-1233003313021131"></a>

<a id="canonical-3321321200322231-2121113022221121-3133321022010302-1100133220100222-1231023320020221-2211001321003320-2313000011021023-3020101100233002"></a>

## name property — port / 012023102111 / 4

Type: `"string"`. Optional.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-0103322122023311-3222010022113230-2201332331011213-2232110312223020-1012320203222033-0231303301331323-0332330321120101-2102123122110320"></a>

<a id="canonical-1233311300231102-0022220333230330-2202103120021230-2132122002210002-0232321020312100-2001000032321001-2301300233103332-3310111100223200"></a>

## num property — port / 012023102111 / 5

Type: `"number"`. Optional.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0111131201110310-2131002022031330-3303203021330032-3001010021233003-3201003321311300-2233221112131333-0021212123322001-2011331011113322"></a>

## Next pages — port / 012023102111 / 6

- [simple_service.container.readiness_check.tcp_health_check](resources--workload--reference--group-017.md#canonical-2120132203003331-2020222300211130-1230021011312132-1001012000221222-1131332001122301-3000222103312120-1010001332230132-0210313102002011)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1110023311313011-0202000212230320-0210021130232023-0323322102330130-3223121303212321-2033302130023021-3120001020331002-1031122330013020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201010321220211-2130020322030011-2123301332230021-2100331332012003-1300201221000010-3013132332311211-1302020231323312-1111022030022122"></a>

## simple_service.disabled — disabled / 233320222232 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- simple_service.disabled

<a id="canonical-1331010000111000-1030202022131033-2012020223002100-0003132030122233-3022132001330130-3220121200323313-3101011022232320-0213200112210222"></a>

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
disabled = {}
```

<a id="canonical-0213031331211302-3001222233230330-3221132121213301-1002201100212332-1002211000002322-1300230222213201-3032111103113113-2133321202001130"></a>

## Direct properties — disabled / 233320222232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022232311113203-2003331122223213-1310032010222012-3030333220332320-3233022020001210-0123313100111032-1130120220211220-2331133022320330"></a>

## Next pages — disabled / 233320222232 / 4

- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0003300212123321-0030022111222023-1030022232221021-1132031212101132-3230032012113311-1303232022202331-3111103333310112-3310021021132230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200031022020123-0313200102103210-2001301012000302-3300211220310323-2311012233323012-1303210220021022-2223321210003133-2330113200131123"></a>

## simple_service.do_not_advertise — do_not_advertise / 130011010100 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- simple_service.do_not_advertise

<a id="canonical-3201133201002001-2102111030231333-3131002012103102-0312101212121323-0001330223220013-0030332222032310-1212110332012133-3210100130302133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for do not advertise.

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
do_not_advertise = {}
```

<a id="canonical-3303200132202030-1332122010112212-2003011322110233-0313110030211322-1102211000321031-1222033212201203-0301130223233330-1300211023330301"></a>

## Direct properties — do_not_advertise / 130011010100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231321002012231-1220133320110221-2311300320331003-1321120210200212-2031301113223003-1121101321003301-0222120233133200-2203232133031313"></a>

## Next pages — do_not_advertise / 130011010100 / 4

- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0022120332023020-2312100113020133-0123031032111130-3302333320122221-1232003333212033-1103032302323333-2231001310003211-0310320121030030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331323031320021-0110122023013313-0312123012032100-0313020100000011-2322120102323032-3331231123223202-3013202122213233-0002012220032203"></a>

## simple_service.enabled — enabled / 320222020321 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- simple_service.enabled

<a id="canonical-1013123123330323-1013010313012311-2000301120012101-3202132002222132-1320332112320030-2323023103313302-3012202221213302-3011202333223233"></a>

Type: `"object"`. single nested block, Optional.

Persistent storage volume configuration for the workload.

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
enabled {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210333233010103-2230002300330331-3022132320312211-1203231122000022-0003113030230320-3023003213222300-1113020312021132-3030300110230320"></a>

## Direct properties — enabled / 320222020321 / 3

<a id="canonical-2122211312113320-3320233131201231-0213002312332013-2121231321302322-3122030133100323-1102333211001202-0223330211322232-3102110110110121"></a>

<a id="canonical-2003200212002032-1331022021330130-2232131013001212-1130000323111003-3201302023321231-0010303001233122-0210210021312101-0131011001212320"></a>

## name property — enabled / 320222020321 / 4

Type: `"string"`. Optional.

Name. Name of the volume.

Upstream description:

Name of the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
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
    "ves.io.schema.rules.string.dns_1123_label": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.dns_1123_label": "true"
  }
}
```

- [persistent_volume](resources--workload--reference--group-017.md#canonical-3113200332101222-1011331333203002-1210020131130330-0320221103232321-3010131320311330-1112331013232233-2101133032201232-1113223131002320): complete subsection reference.

<a id="canonical-1202020022110313-0000020333212010-1000221303313330-3312123220000321-3120120030300202-1100113322321112-2320200231022330-0022030321112013"></a>

## Next pages — enabled / 320222020321 / 5

- [simple_service.enabled.persistent_volume](resources--workload--reference--group-017.md#canonical-3113200332101222-1011331333203002-1210020131130330-0320221103232321-3010131320311330-1112331013232233-2101133032201232-1113223131002320)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3113200332101222-1011331333203002-1210020131130330-0320221103232321-3010131320311330-1112331013232233-2101133032201232-1113223131002320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033011223320002-3031303111120031-3132023120233131-0022321330323011-0333203320333332-0130022333110012-0223312333021112-2023031300200303"></a>

## simple_service.enabled.persistent_volume — persistent_volume / 021200102110 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.enabled](resources--workload--reference--group-017.md#canonical-0022120332023020-2312100113020133-0123031032111130-3302333320122221-1232003333212033-1103032302323333-2231001310003211-0310320121030030)
- simple_service.enabled.persistent_volume

<a id="canonical-1121031211201101-0230221011311012-1320210121011002-3010012020313132-3211031101333331-3220101211012313-3113211201110013-3003023001003131"></a>

Type: `"object"`. single nested block, Optional.

Volume containing the Persistent Storage for the workload.

Receipt-pinned upstream constraints:

```json
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
persistent_volume {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122300302213310-3101012313213110-2301331331021001-1313101311010023-2101322120023121-1312212033032203-0031300010030210-3211111110030010"></a>

## Direct properties — persistent_volume / 021200102110 / 3

- [mount](resources--workload--reference--group-017.md#canonical-2111120313310323-2221023331321012-2030221311033320-3003121230321330-3323312203123022-2322013321110122-0100031032103110-0022321200201303): complete subsection reference.

- [storage](resources--workload--reference--group-017.md#canonical-0212031230210030-0110133221302101-1211130030200033-0222110222222113-1311020123323113-1302101310023300-1101100122133002-0221000223123310): complete subsection reference.

<a id="canonical-0121100232200110-3101323310010012-0132201202211020-1331020121333101-0121201332110123-1000031110011131-0223012210300323-0210033222332213"></a>

## Next pages — persistent_volume / 021200102110 / 4

- [simple_service.enabled.persistent_volume.mount](resources--workload--reference--group-017.md#canonical-2111120313310323-2221023331321012-2030221311033320-3003121230321330-3323312203123022-2322013321110122-0100031032103110-0022321200201303)
- [simple_service.enabled.persistent_volume.storage](resources--workload--reference--group-017.md#canonical-0212031230210030-0110133221302101-1211130030200033-0222110222222113-1311020123323113-1302101310023300-1101100122133002-0221000223123310)
- [simple_service.enabled](resources--workload--reference--group-017.md#canonical-0022120332023020-2312100113020133-0123031032111130-3302333320122221-1232003333212033-1103032302323333-2231001310003211-0310320121030030)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2111120313310323-2221023331321012-2030221311033320-3003121230321330-3323312203123022-2322013321110122-0100031032103110-0022321200201303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323222300202021-3212100311320323-1211130303132100-1330133113300202-2102121123303030-2110021010123122-1131311121130020-0100201312003330"></a>

## simple_service.enabled.persistent_volume.mount — mount / 001103230223 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.enabled](resources--workload--reference--group-017.md#canonical-0022120332023020-2312100113020133-0123031032111130-3302333320122221-1232003333212033-1103032302323333-2231001310003211-0310320121030030)
- [simple_service.enabled.persistent_volume](resources--workload--reference--group-017.md#canonical-3113200332101222-1011331333203002-1210020131130330-0320221103232321-3010131320311330-1112331013232233-2101133032201232-1113223131002320)
- simple_service.enabled.persistent_volume.mount

<a id="canonical-0220032131123200-0013033203120011-3331122310210130-2221212231023312-3200220300103100-0323312333100232-2330001202230120-3333102212130211"></a>

Type: `"object"`. single nested block, Optional.

Volume mount describes how volume is mounted inside a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mount_path")}
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
mount {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232213212132313-1330221032021011-3002323300311312-0202112323201012-2303110312320123-3111211332120231-2332222132301122-0313020013320000"></a>

## Direct properties — mount / 001103230223 / 3

<a id="canonical-0132010212030021-3223033023223330-0312131223203131-2122022231032021-0001123311012013-1301011231131303-3020203101330133-3012122111132313"></a>

<a id="canonical-0332300312133220-1003301300012000-3303310210312111-0302221301111002-1313210232321103-2001031101220133-0131000123133031-1010202111132100"></a>

## mode property — mount / 001103230223 / 4

Type: `"string"`. Optional.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3010210323030012-1320021010100201-0221312310322320-2321001122012233-2321213122220211-2113223230010230-0112103132221132-3330222210331232"></a>

<a id="canonical-2002233101012320-0123021130023110-3302010211130301-2233330330112223-3000220021233022-0101101002303322-2313003003203332-3232102323220031"></a>

## mount_path property — mount / 001103230223 / 5

Type: `"string"`. Optional.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[^:]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-0211301020323111-3320220232130113-0110333333231123-2023100310031221-1013330302012201-3213332221331311-3101332203301330-1300120210122111"></a>

<a id="canonical-1100232301220202-2011132322031333-2201103313130022-2311231212000301-0112232332121132-2133332102200112-1323320021201112-0011100123111213"></a>

## sub_path property — mount / 001103230223 / 6

Type: `"string"`. Optional.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3011022332103231-2011133323211322-0200303331210301-1001202333221133-0213312133023031-3301322301113002-2010112313303111-2330022001220111"></a>

## Next pages — mount / 001103230223 / 7

- [simple_service.enabled.persistent_volume](resources--workload--reference--group-017.md#canonical-3113200332101222-1011331333203002-1210020131130330-0320221103232321-3010131320311330-1112331013232233-2101133032201232-1113223131002320)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0212031230210030-0110133221302101-1211130030200033-0222110222222113-1311020123323113-1302101310023300-1101100122133002-0221000223123310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002113023121033-1123313000023001-0330032130132130-3330330231002202-0012221231313212-3310333230301201-1013222322013200-3133220020333222"></a>

## simple_service.enabled.persistent_volume.storage — storage / 223222133010 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.enabled](resources--workload--reference--group-017.md#canonical-0022120332023020-2312100113020133-0123031032111130-3302333320122221-1232003333212033-1103032302323333-2231001310003211-0310320121030030)
- [simple_service.enabled.persistent_volume](resources--workload--reference--group-017.md#canonical-3113200332101222-1011331333203002-1210020131130330-0320221103232321-3010131320311330-1112331013232233-2101133032201232-1113223131002320)
- simple_service.enabled.persistent_volume.storage

<a id="canonical-1000100132031322-3002211033330331-1123302221120222-1312132220011223-3200313220103133-1203203103230310-2111231230102213-0232313112303032"></a>

Type: `"object"`. single nested block, Optional.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

Upstream description:

Persistent storage configuration is used to configure Persistent Volume Claim (PVC)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_size"),
  validators.ConflictingObjectAttributes("class_name",
    "default")}
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
  "x-ves-oneof-field-class_name_choice": "[\"class_name\",\"default\"]"
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201123023210313-2120313303123303-0020312333233123-0013021110302033-0333213130212232-2313033132000330-0300213121230201-3201331011310000"></a>

## Direct properties — storage / 223222133010 / 3

<a id="canonical-1212230222002111-3312131323303032-1331132123020330-1122321213323003-2021231213033221-0123222200311110-2031233002110000-0222022122301301"></a>

<a id="canonical-2301333013123012-1311301220123123-3102312210022220-1032120221312133-2113001112032213-3032330023210122-2331001321020310-0220020032310301"></a>

## access_mode property — storage / 223222133010 / 4

Type: `"string"`. Optional.

\[Enum:
ACCESS\_MODE\_READ\_WRITE\_ONCE|ACCESS\_MODE\_READ\_WRITE\_MANY|ACCESS\_MODE\_READ\_ONLY\_MANY\]
Persistence storage access mode is used to configure access mode for persistent storage -
ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once Read Write Once is used to mount persistent storage
in read/write mode to exactly 1 host - ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many Read Write
Many is used.. Possible values are \`ACCESS\_MODE\_READ\_WRITE\_ONCE\`,
\`ACCESS\_MODE\_READ\_WRITE\_MANY\`, \`ACCESS\_MODE\_READ\_ONLY\_MANY\`. Defaults to
\`ACCESS\_MODE\_READ\_WRITE\_ONCE\`.

Upstream description:

Persistence storage access mode is used to configure access mode for persistent storage

&#8203;- ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once

Read Write Once is used to mount persistent storage in read/write mode to exactly 1 host &#8203;-
ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many

Read Write Many is used to mount persistent storage in read/write mode to many hosts &#8203;-
ACCESS\_MODE\_READ\_ONLY\_MANY: Read Only Many

Read Only Many is used to mount persistent storage in read-only mode to many hosts.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ACCESS_MODE_READ_WRITE_ONCE",
  "enum": [
    "ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3323031332311133-0033013133220223-3303222202031320-0012332200011011-0332221200310020-2313023030201222-2012300030021132-3303212122313003"></a>

<a id="canonical-2130103220131312-3211223212332213-1130221120311221-1323133030220200-1220031121101113-2112132011333302-3012011123101131-0000121103111313"></a>

## class_name property — storage / 223222133010 / 5

Type: `"string"`. Optional.

Exclusive with \[default\] Use the specified class name.

Upstream description:

Exclusive with \[default\] Use the specified class name.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [default](resources--workload--reference--group-017.md#canonical-3012102102201331-1221031112211213-1222123121211333-0030232213203203-0312010203111212-2102322332230021-0201012112121333-3231320112022100): complete subsection reference.

<a id="canonical-3132133011031120-1223310211331330-3333332101012110-0023333121120310-1131301122122131-0133000103100123-2312321331330113-2221322032130001"></a>

<a id="canonical-0012322012310111-2233313211100100-2030010221120301-2002022320330132-3203232311131132-3021202311023013-2021202222223211-0320312333020020"></a>

## storage_size property — storage / 223222133010 / 6

Type: `"number"`. Optional.

Size (in GiB). Size in GiB of the persistent storage.

Upstream description:

Size in GiB of the persistent storage.

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
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0213311013333210-0333120011233113-0032122213100223-1301320313201221-3231102020110111-0020203111021330-2100100011100303-1303230131031132"></a>

## Next pages — storage / 223222133010 / 7

- [simple_service.enabled.persistent_volume.storage.default](resources--workload--reference--group-017.md#canonical-3012102102201331-1221031112211213-1222123121211333-0030232213203203-0312010203111212-2102322332230021-0201012112121333-3231320112022100)
- [simple_service.enabled.persistent_volume](resources--workload--reference--group-017.md#canonical-3113200332101222-1011331333203002-1210020131130330-0320221103232321-3010131320311330-1112331013232233-2101133032201232-1113223131002320)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3012102102201331-1221031112211213-1222123121211333-0030232213203203-0312010203111212-2102322332230021-0201012112121333-3231320112022100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021300331000002-2120101213103000-3032201103323210-1313323020102012-0011111113202333-3333232111322010-2232333213312300-2311313213122003"></a>

## simple_service.enabled.persistent_volume.storage.default — default / 302100032100 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [simple_service.enabled](resources--workload--reference--group-017.md#canonical-0022120332023020-2312100113020133-0123031032111130-3302333320122221-1232003333212033-1103032302323333-2231001310003211-0310320121030030)
- [simple_service.enabled.persistent_volume](resources--workload--reference--group-017.md#canonical-3113200332101222-1011331333203002-1210020131130330-0320221103232321-3010131320311330-1112331013232233-2101133032201232-1113223131002320)
- [simple_service.enabled.persistent_volume.storage](resources--workload--reference--group-017.md#canonical-0212031230210030-0110133221302101-1211130030200033-0222110222222113-1311020123323113-1302101310023300-1101100122133002-0221000223123310)
- simple_service.enabled.persistent_volume.storage.default

<a id="canonical-2222232222023233-3312203132202011-0303300232132200-3300331123133322-2031303031121030-3312031001313302-1231303221321120-3011301303203320"></a>

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
default = {}
```

<a id="canonical-1123212221312031-2300011023232301-2131032103322011-0110103101100120-2323000103131311-1202122112323022-0313311330210112-3133212132232230"></a>

## Direct properties — default / 302100032100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013031320302131-1131300222213000-1003000311212303-0200302300213112-3001202201330011-1311200201133311-2211113010001232-2023103301332030"></a>

## Next pages — default / 302100032100 / 4

- [simple_service.enabled.persistent_volume.storage](resources--workload--reference--group-017.md#canonical-0212031230210030-0110133221302101-1211130030200033-0222110222222113-1311020123323113-1302101310023300-1101100122133002-0221000223123310)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3310302312301012-0213323302330302-2012000032301123-3011230301130001-2021102001021321-2321111011331112-1000322300330133-3212013211230202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221333133112213-0332220003031033-1200010210303310-3333120333132212-2030333221322231-2322211200002001-3002133220011022-0323011132110022"></a>

## simple_service.simple_advertise — simple_advertise / 223111233203 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- simple_service.simple_advertise

<a id="canonical-3130001232212321-0332110202020133-2132201100212101-0000322300221133-1210203232221223-2303320002123131-2312022013122221-1230131020132102"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for simple advertise.

Upstream description:

Advertise OPTIONS for Simple Service.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains",
    "service_port")}
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
simple_advertise {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000100231231103-2022113212320033-0113030001030300-0310231011223320-3130111201120300-2203323013331300-3323121000122301-2331111213211333"></a>

## Direct properties — simple_advertise / 223111233203 / 3

<a id="canonical-0122012001223110-0120001223303221-3102231311113331-1121221211010232-3331100020203300-1112030012302303-2120231002101320-1102032302013231"></a>

<a id="canonical-2031223332221312-0320022130000223-0233123201012203-0223032311000333-0031321320322323-0203220210202200-2012220122300222-0111000023201001"></a>

## domains property — simple_advertise / 223111233203 / 4

Type: `["list", "string"]`. Optional.

List of Domains (host/authority header) that will be matched to Load Balancer. Wildcard hosts are
supported in the suffix or prefix form Supported Domains and search order: 1. Exact Domain names:
www&#46;example.com. 2.

Upstream description:

A list of Domains (host/authority header) that will be matched to Load Balancer. Wildcard hosts are
supported in the suffix or prefix form

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Domains are also used for SNI matching if the Load Balancer type is HTTPS. Domains also indicate the
list of names for which DNS resolution will be automatically resolved to IP addresses by the system.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3000000110013201-2013100122101320-0201322210302332-1031200303002001-0220303321303021-0231110213210333-0011302301233032-0300100221311013"></a>

<a id="canonical-3301112220310020-2301200002223210-2331213011121331-1212313020300101-0032321320333012-0132123322301321-0332030221131132-0132303010320221"></a>

## service_port property — simple_advertise / 223111233203 / 5

Type: `"number"`. Optional.

Service port to advertise on Internet via HTTP loadbalancer using port 80.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1024, 65535),
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1024
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0000022000303313-2111301100221301-2312110332312000-1011322331312201-1332121213332102-3113321321113223-3201303100130210-0312220011231010"></a>

## Next pages — simple_advertise / 223111233203 / 6

- [simple_service](resources--workload--reference--group-016.md#canonical-2103120123000022-0231212311020021-3011030332210231-1230230313313103-1213231032231100-0321113222111303-0220111021013230-2333203212100121)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022323221232223-0223203203302322-1033001223102223-2122131202221323-0013301100022211-0300130230002033-0101220101323013-1232212002123001"></a>

## stateful_service — stateful_service / 121020301123 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- stateful_service

<a id="canonical-3131320200112003-1302100202032202-0133102112232312-1213202222213131-2120333030221010-2233200311233003-3313123001231110-2320312322302201"></a>

Type: `"object"`. single nested block, Optional.

StatefulService maintains per replica state and each replica has its own persistent storage. Each
replica has a unique network identity and stable storage. Stateful service are used for distributed
stateful applications like cassandra, mongodb, redis, etc.

Upstream description:

StatefulService maintains per replica state and each replica has its own persistent storage. Each
replica has a unique network identity and stable storage. Stateful service are used for distributed
stateful applications like cassandra, mongodb, redis, etc.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("containers",
    "persistent_volumes"),
  validators.ConflictingObjectAttributes("num_replicas",
    "scale_to_zero")}
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
  "x-ves-oneof-field-scaling_choice": "[\"num_replicas\",\"scale_to_zero\"]"
}
```

Terraform syntax:

```terraform
stateful_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133203010131332-0300112021200321-3223133101100223-0303002212130201-0120011301102202-2121310303303210-3110201132031020-0321323132012322"></a>

## Direct properties — stateful_service / 121020301123 / 3

- [advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323): complete subsection reference.

- [configuration](resources--workload--reference--group-028.md#canonical-3203012310220003-2210330300331110-2103223003030222-2222201323310103-2300110002113011-0130333021003321-0201020033321220-1102301103203230): complete subsection reference.

- [containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301): complete subsection reference.

- [deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003): complete subsection reference.

<a id="canonical-1303111221221331-1323333220111212-3302320313231221-0200210101220130-2031031300320011-2000123031220021-2233201111020111-3123113032312120"></a>

<a id="canonical-1033010223303323-2131302221003110-2232302020313220-3010123131133302-1332133200011003-1013130232331013-3023130303202130-1201232321330031"></a>

## num_replicas property — stateful_service / 121020301123 / 4

Type: `"number"`. Optional.

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Upstream description:

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  }
}
```

- [persistent_volumes](resources--workload--reference--group-029.md#canonical-2202000332001022-3031133313102112-0003233201110202-2303003230303311-3002213310230220-0131313123131032-0010302220210011-2300030100331332): complete subsection reference.

- [scale_to_zero](resources--workload--reference--group-029.md#canonical-0221221200321221-1133220231313302-3122113332002302-1303223333210013-3310313331033111-2320113332020010-2301111102010030-3110331212031230): complete subsection reference.

- [volumes](resources--workload--reference--group-029.md#canonical-0320321332330022-3200203220020212-3113213300013223-3331000302332100-2331130020201320-0110232101322132-3201120100233223-3122231121020213): complete subsection reference.

<a id="canonical-0103233220221203-1312011331221020-1222323011231300-0120121203311131-1021033221321021-2000311231032302-0110120201020033-1010003132111310"></a>

## Next pages — stateful_service / 121020301123 / 5

- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.configuration](resources--workload--reference--group-028.md#canonical-3203012310220003-2210330300331110-2103223003030222-2222201323310103-2300110002113011-0130333021003321-0201020033321220-1102301103203230)
- [stateful_service.containers](resources--workload--reference--group-028.md#canonical-3221212333032323-1011122002030020-0332100102213333-2230123311320001-3101132000200220-3323302320310311-1200332113313321-1231002100113301)
- [stateful_service.deploy_options](resources--workload--reference--group-029.md#canonical-0211010322210330-0002130101323013-1320121223020011-3332231210010332-1331310201011323-2210300012032000-2333113200230221-2013202000203003)
- [stateful_service.persistent_volumes](resources--workload--reference--group-029.md#canonical-2202000332001022-3031133313102112-0003233201110202-2303003230303311-3002213310230220-0131313123131032-0010302220210011-2300030100331332)
- [stateful_service.scale_to_zero](resources--workload--reference--group-029.md#canonical-0221221200321221-1133220231313302-3122113332002302-1303223333210013-3310313331033111-2320113332020010-2301111102010030-3110331212031230)
- [stateful_service.volumes](resources--workload--reference--group-029.md#canonical-0320321332330022-3200203220020212-3113213300013223-3331000302332100-2331130020201320-0110232101322132-3201120100233223-3122231121020213)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010331033010110-2011311110010103-3021302230332311-2000030002111011-0031320200113003-3321212202030031-3303020112222222-1212013203302131"></a>

## stateful_service.advertise_options — advertise_options / 301320223303 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- stateful_service.advertise_options

<a id="canonical-3223330011020031-0330212002303111-2103131212302003-0310313201230333-1322000211100323-0301031101002100-1200021112333203-3200030221111101"></a>

Type: `"object"`. single nested block, Optional.

Advertise OPTIONS are used to configure how and where to advertise the workload using load
balancers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_in_cluster"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_custom",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_in_cluster",
    "advertise_on_public"),
  validators.ConflictingObjectAttributes("advertise_in_cluster",
    "do_not_advertise"),
  validators.ConflictingObjectAttributes("advertise_on_public",
    "do_not_advertise")}
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
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_in_cluster\",\"advertise_on_public\",\"do_not_advertise\"]"
}
```

Terraform syntax:

```terraform
advertise_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003233010211111-2100011000300330-0303122212020200-0300030202232000-3113020212310232-0301300300230302-0033223302022131-0022310230212030"></a>

## Direct properties — advertise_options / 301320223303 / 3

- [advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002): complete subsection reference.

- [advertise_in_cluster](resources--workload--reference--group-021.md#canonical-2203200033111100-1131311021203132-1232001231010211-3000112112233321-3002013120002033-2111032232010121-0222102100212312-0012310120331313): complete subsection reference.

- [advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202): complete subsection reference.

- [do_not_advertise](resources--workload--reference--group-028.md#canonical-3202321232103323-1000231221221030-0001122222031132-1003133321010023-1002201331103010-1230130022220311-0032100201113003-0333121223233010): complete subsection reference.

<a id="canonical-3121113011013322-1113220120100330-2320202323122123-1132220011233110-3310230303232123-3133032303113323-2022233211223130-2230130023011320"></a>

## Next pages — advertise_options / 301320223303 / 4

- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_in_cluster](resources--workload--reference--group-021.md#canonical-2203200033111100-1131311021203132-1232001231010211-3000112112233321-3002013120002033-2111032232010121-0222102100212312-0012310120331313)
- [stateful_service.advertise_options.advertise_on_public](resources--workload--reference--group-021.md#canonical-3300101333202310-3223231311312310-2030111033033203-1011103231310023-1313110120311121-1221031312313020-1210302113302033-2110231331000202)
- [stateful_service.advertise_options.do_not_advertise](resources--workload--reference--group-028.md#canonical-3202321232103323-1000231221221030-0001122222031132-1003133321010023-1002201331103010-1230130022220311-0032100201113003-0333121223233010)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302231213303220-1200222203031002-1301202003113103-3213030121100331-1120020301311332-1230211103333322-0303100201210333-2132022023233033"></a>

## stateful_service.advertise_options.advertise_custom — advertise_custom / 210020112122 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- stateful_service.advertise_options.advertise_custom

<a id="canonical-1211201231012312-1032003212201201-3213201111203000-0232213003331200-3201332312303101-1200023022122231-3323232333232221-2010023221333112"></a>

Type: `"object"`. single nested block, Optional.

Advertise this workload via loadbalancer on specific sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where",
    "ports")}
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
advertise_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-2011002230313100-3010000221023121-1332030123022103-3301023212003223-2313002122021200-1003022133233230-0102122130023101-1023032221033202"></a>

## Direct properties — advertise_custom / 210020112122 / 3

- [advertise_where](resources--workload--reference--group-017.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310): complete subsection reference.

- [ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231): complete subsection reference.

<a id="canonical-1220213231012331-0311322000332132-3110022300003001-2131111030111221-1311333011220121-2122021333322011-1120223132020011-3012112121213323"></a>

## Next pages — advertise_custom / 210020112122 / 4

- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113022102202132-1101001210202203-0312021310110020-2022031210311222-2121220033220310-1230221002013301-3030331331101112-3331113213131120"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where — advertise_where / 331120301111 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- stateful_service.advertise_options.advertise_custom.advertise_where

<a id="canonical-3002101332033022-1110203201303323-3111222110113121-2013301001130212-3122120231011221-0103031132322131-0202303321122110-3220303302033333"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "vk8s_service")}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3320002112233033-2131211201210211-2103231333102103-3221211032110312-1001112202313100-2331303203001002-3031333013222101-0030223102211230"></a>

## Direct properties — advertise_where / 331120301111 / 3

- [site](resources--workload--reference--group-017.md#canonical-2223211010300233-0033203221303312-1010220021220302-2021010313033010-0231212132123031-3022301303002012-3030332322213232-3221210002300132): complete subsection reference.

- [virtual_site](resources--workload--reference--group-017.md#canonical-0031013123213022-1332033033011221-1030121103331012-1020201303132332-0101213311110012-2300122113201200-2102222023002321-1301121133120003): complete subsection reference.

- [vk8s_service](resources--workload--reference--group-017.md#canonical-3113201100321032-2110322211311302-0001210313012211-3231233213313313-3102200133233230-2000333233123013-1333110321022331-3013320123303013): complete subsection reference.

<a id="canonical-0122212101002203-3011120200110112-1132123123132122-3113023223123023-1230020323131022-0111133110303133-3332303110220323-2203301230100130"></a>

## Next pages — advertise_where / 331120301111 / 4

- [stateful_service.advertise_options.advertise_custom.advertise_where.site](resources--workload--reference--group-017.md#canonical-2223211010300233-0033203221303312-1010220021220302-2021010313033010-0231212132123031-3022301303002012-3030332322213232-3221210002300132)
- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site](resources--workload--reference--group-017.md#canonical-0031013123213022-1332033033011221-1030121103331012-1020201303132332-0101213311110012-2300122113201200-2102222023002321-1301121133120003)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-017.md#canonical-3113201100321032-2110322211311302-0001210313012211-3231233213313313-3102200133233230-2000333233123013-1333110321022331-3013320123303013)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2223211010300233-0033203221303312-1010220021220302-2021010313033010-0231212132123031-3022301303002012-3030332322213232-3221210002300132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331321203112213-2230213110030012-2033313311322330-3030033131121130-2332112003201031-0013300002100312-3213131002302322-2012121011321112"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.site — site / 111322030010 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- stateful_service.advertise_options.advertise_custom.advertise_where.site

<a id="canonical-3030100122133203-0100200333101100-3012210101031123-2121200020330231-0031123232233300-2131032303301313-3101033323000321-2111021033022101"></a>

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

<a id="canonical-1213132200202121-3232011111220210-2221222321212120-1231100003121323-0210112120200023-2323001110012321-2332013222130200-1013113331033011"></a>

## Direct properties — site / 111322030010 / 3

<a id="canonical-2131231302322001-1000330012021011-2003031320011331-1102021013131200-1333130300222303-1220301200313130-3203332012120323-3021030301221010"></a>

<a id="canonical-1200322120300012-1232030230030111-2013003312110032-2022000003230300-1103211300032102-0322310130211003-1311223210203010-1121313300010310"></a>

## ip property — site / 111322030010 / 4

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

<a id="canonical-0323010001001310-2100333320222223-2132021223301130-3302332001030002-1312100023001330-1020123300220113-1132011312033333-1011111302330332"></a>

<a id="canonical-0222203331330320-2223303130333122-0132212303332322-0222323122130211-3213120020103302-1200213232200030-1331102023211010-2303130120020320"></a>

## network property — site / 111322030010 / 5

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

- [site](resources--workload--reference--group-017.md#canonical-2323021101211302-3220332313132300-3003021010013310-2112330231220120-3012002303222131-3001333101022320-3333013030233110-1023103020210111): complete subsection reference.

<a id="canonical-3331320231303320-2012020031101100-2102303122313112-1010110033121030-1132330123300303-1220333201321210-3200333011202130-1100023000130211"></a>

## Next pages — site / 111322030010 / 6

- [stateful_service.advertise_options.advertise_custom.advertise_where.site.site](resources--workload--reference--group-017.md#canonical-2323021101211302-3220332313132300-3003021010013310-2112330231220120-3012002303222131-3001333101022320-3333013030233110-1023103020210111)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2323021101211302-3220332313132300-3003021010013310-2112330231220120-3012002303222131-3001333101022320-3333013030233110-1023103020210111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212033032301112-2213232320101331-3031132233223201-0030221220032110-1221321212022331-1202001232213103-3120301202113300-1013022011332230"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.site.site — site / 330231311033 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- [stateful_service.advertise_options.advertise_custom.advertise_where.site](resources--workload--reference--group-017.md#canonical-2223211010300233-0033203221303312-1010220021220302-2021010313033010-0231212132123031-3022301303002012-3030332322213232-3221210002300132)
- stateful_service.advertise_options.advertise_custom.advertise_where.site.site

<a id="canonical-2330000103132010-1101121120211222-2100133012313021-3220203232302231-2212100033212123-3003033033003203-1233032121131111-3011311322022312"></a>

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

<a id="canonical-1200103011201030-0322102210213013-2130213032002123-0212102232121200-1312020121002000-0121111132232331-0323030323121122-2311110333200023"></a>

## Direct properties — site / 330231311033 / 3

<a id="canonical-0300000300201101-2031002131022313-0301331331333211-2121233122210013-1310003120300123-2002123031222112-1112221323023032-0031220130130301"></a>

<a id="canonical-2032221333033003-2320323120320033-3331321111301021-3133312102110011-1312130031130301-3011131133333002-2311110330202321-0031133220003313"></a>

## name property — site / 330231311033 / 4

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

<a id="canonical-0030032102230132-3330031201221033-2023130231300311-0023103301221233-3232012313222113-2211111330310203-3001003103100333-2030211302211020"></a>

<a id="canonical-1112320202333122-1133003121210113-0310131230233331-1333110013113032-1112312003023100-1201010022010010-1100213012120232-0311130111122033"></a>

## namespace property — site / 330231311033 / 5

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

<a id="canonical-3000010211310313-2100123322311212-2223011023120002-0333030223312110-3332220301013003-0000102202102301-2210001331110200-3000321200102203"></a>

<a id="canonical-1223220022000002-0033131302321131-2313330031233131-1012210103112311-3211203131111113-0121301322231311-1331120001020212-3132000210203012"></a>

## tenant property — site / 330231311033 / 6

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

<a id="canonical-1111131121231013-3320331302131023-1112111310111033-1111312030210333-0133123312202232-2110333112322220-2103232222022130-3200033133223011"></a>

## Next pages — site / 330231311033 / 7

- [stateful_service.advertise_options.advertise_custom.advertise_where.site](resources--workload--reference--group-017.md#canonical-2223211010300233-0033203221303312-1010220021220302-2021010313033010-0231212132123031-3022301303002012-3030332322213232-3221210002300132)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0031013123213022-1332033033011221-1030121103331012-1020201303132332-0101213311110012-2300122113201200-2102222023002321-1301121133120003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330213332303021-0311130132002031-2030310010132322-2221132330210123-0111131121001133-2230133133023332-3231033333203211-0322010313133021"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site — virtual_site / 233133002021 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site

<a id="canonical-0220333133222202-1213221113130000-3210212200321203-3230111233102330-3212122111330023-1211102032122223-0123310322232321-0023033112112100"></a>

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

<a id="canonical-3202300013230301-1020032001020311-2010010003232333-2103232003003301-1330232023201220-0332112031111122-3232302200003202-1223310230131123"></a>

## Direct properties — virtual_site / 233133002021 / 3

<a id="canonical-0001022133021103-0021013300301032-3120332331130301-2123233012201020-1331121012033303-3210212000311220-0023130030001333-1220031211012002"></a>

<a id="canonical-0031330033023001-1133030001320023-1122330200223033-3303102210003100-0012232120330333-0323313333330022-1102110210330301-1211000310010010"></a>

## network property — virtual_site / 233133002021 / 4

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

- [virtual_site](resources--workload--reference--group-017.md#canonical-1333212202330222-1102223021022033-2001133332013310-3013122023221222-1100021323030200-3331213011031030-0333012123031301-3033223122131333): complete subsection reference.

<a id="canonical-2021003311220113-3133113103301133-3133223133003131-3131021223223032-2223213321102323-2231322121322111-2330200030200023-3330023000110301"></a>

## Next pages — virtual_site / 233133002021 / 5

- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site](resources--workload--reference--group-017.md#canonical-1333212202330222-1102223021022033-2001133332013310-3013122023221222-1100021323030200-3331213011031030-0333012123031301-3033223122131333)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1333212202330222-1102223021022033-2001133332013310-3013122023221222-1100021323030200-3331213011031030-0333012123031301-3033223122131333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331130111223210-1123103223012311-3332321121211032-1220011013020222-0112312112220113-1232202210202121-1312332301001003-0132112022100133"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site — virtual_site / 320332122313 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site](resources--workload--reference--group-017.md#canonical-0031013123213022-1332033033011221-1030121103331012-1020201303132332-0101213311110012-2300122113201200-2102222023002321-1301121133120003)
- stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-3112231022331201-0122320233201312-1113011321130010-1330333023030022-0001303222133023-1330200301130133-2321213131030000-2330220130100022"></a>

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

<a id="canonical-1020303012330200-2030331021221003-3101023332012103-0232011331301323-3231130123113312-3000020122020231-1213100122331330-1033010203302101"></a>

## Direct properties — virtual_site / 320332122313 / 3

<a id="canonical-1031222202322033-2232223033002112-2110030221232210-3202010133323211-0212031221000031-3310210320213002-0120310122030010-3032011323013000"></a>

<a id="canonical-1321310202113222-2232323100220102-0232213002332001-1103111111302131-2220220321320100-2310133303023222-1121213201000310-0102100002220022"></a>

## name property — virtual_site / 320332122313 / 4

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

<a id="canonical-3013210031012013-2001020300131230-2001110100330032-0322211202003223-2330321311222002-3321303103231230-1010112013113233-3333302101332113"></a>

<a id="canonical-2322223012200303-1230222200220112-3022300313233110-0103322100233110-2203121101031211-3312102013301001-2012230312032320-0021113333131101"></a>

## namespace property — virtual_site / 320332122313 / 5

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

<a id="canonical-1210320201031133-0112333113202130-3203310003130332-3300303110301213-2000011303110032-1202332013223023-3023212012000133-3113313330112312"></a>

<a id="canonical-3111202332311101-0201012131003312-2331033320131032-0133132030231023-2023131010031332-1212113212002130-1213110011220202-0201001130113110"></a>

## tenant property — virtual_site / 320332122313 / 6

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

<a id="canonical-2122023123223131-1020120332220301-0220223213033033-1023230013132131-2310223113302203-3110200112120230-1003333233202331-3001210213333033"></a>

## Next pages — virtual_site / 320332122313 / 7

- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site](resources--workload--reference--group-017.md#canonical-0031013123213022-1332033033011221-1030121103331012-1020201303132332-0101213311110012-2300122113201200-2102222023002321-1301121133120003)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3113201100321032-2110322211311302-0001210313012211-3231233213313313-3102200133233230-2000333233123013-1333110321022331-3013320123303013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222321311220002-2020222221131120-0213301213121012-2120101113221320-0311022030112202-3033110032233302-1103221103203311-1311113330323300"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service — vk8s_service / 131322010123 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service

<a id="canonical-2103022233101201-0210213120123221-0010323320303022-0121300023000103-3231021132010100-0321001232303102-3202003000111012-0312021331330302"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

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
vk8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322312111012100-2012200103223022-2222320210310021-0002131333231323-3112202301112131-2201133210221221-2031311013031321-0031012232110330"></a>

## Direct properties — vk8s_service / 131322010123 / 3

- [site](resources--workload--reference--group-017.md#canonical-2020100103012113-0010110101130222-1022200321203230-2222130032123120-1013300002223112-3121322123010013-0110323230231133-1310002322321131): complete subsection reference.

- [virtual_site](resources--workload--reference--group-017.md#canonical-3003101031302223-0033311311220102-1003320120221221-2231021023123210-3331032233001201-1033020022011030-3201113002230220-3232213303301221): complete subsection reference.

<a id="canonical-1200231200032002-2322301202322203-2200023130033000-2331210010003233-2022023100102311-3213121112210330-0023111311212221-1122033101031122"></a>

## Next pages — vk8s_service / 131322010123 / 4

- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site](resources--workload--reference--group-017.md#canonical-2020100103012113-0010110101130222-1022200321203230-2222130032123120-1013300002223112-3121322123010013-0110323230231133-1310002322321131)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site](resources--workload--reference--group-017.md#canonical-3003101031302223-0033311311220102-1003320120221221-2231021023123210-3331032233001201-1033020022011030-3201113002230220-3232213303301221)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2020100103012113-0010110101130222-1022200321203230-2222130032123120-1013300002223112-3121322123010013-0110323230231133-1310002322321131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331303133002203-3011301323313132-3102311231020123-0202002300311220-0021211123132200-2011202030322020-2210220313300131-3002300002000013"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site — site / 313313211300 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-017.md#canonical-3113201100321032-2110322211311302-0001210313012211-3231233213313313-3102200133233230-2000333233123013-1333110321022331-3013320123303013)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-2021321331120132-0203203332000011-2302001223300012-2222332130313321-2300012203213100-2301302313233022-3333033332123302-3213132212232131"></a>

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

<a id="canonical-2012220231031011-2303203100131110-2032311101001131-3102112312131223-2010122131332213-0233213113320310-0123023132202310-2202323001230231"></a>

## Direct properties — site / 313313211300 / 3

<a id="canonical-1002303202210120-2231121023110201-0322121002131002-1031301321312122-0033033310111333-3322230311310121-1103102310100121-2021331121312201"></a>

<a id="canonical-0321211030001132-1322302132122323-3300323310230201-3031103210030313-1021310022211123-1312303211203012-3312020103030100-0012000111300020"></a>

## name property — site / 313313211300 / 4

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

<a id="canonical-1032232102111301-0122131032203001-3130323121030202-3311202211001112-2000233232111112-3131232201102202-3132210132232222-0103310012100212"></a>

<a id="canonical-0010331333032211-2301100310210100-2012112003030222-1233231021230211-1030111201232103-0211011101203332-0101312113313123-1333212022323010"></a>

## namespace property — site / 313313211300 / 5

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

<a id="canonical-2103033322011011-3330320231012311-2111103013312122-0001023032012001-0331303200320113-3131213232010230-0300121011303001-1123311030320203"></a>

<a id="canonical-2020230122012023-1301121113213103-2223330310021010-3333111123330122-2320200032221231-2011030220312320-3011122210001002-3331020001220211"></a>

## tenant property — site / 313313211300 / 6

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

<a id="canonical-0222322101200020-0231313323102032-0332121333301000-0023232202232011-3320023232203332-2233223302132311-3223123321010301-0313331332001202"></a>

## Next pages — site / 313313211300 / 7

- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-017.md#canonical-3113201100321032-2110322211311302-0001210313012211-3231233213313313-3102200133233230-2000333233123013-1333110321022331-3013320123303013)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3003101031302223-0033311311220102-1003320120221221-2231021023123210-3331032233001201-1033020022011030-3201113002230220-3232213303301221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333222303302131-0123001210230212-1312022303131323-2201222023030102-2101333121030021-1033300023223132-2322123222301110-3232332201301003"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site — virtual_site / 032331111313 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.advertise_where](resources--workload--reference--group-017.md#canonical-1202202011031000-3220322121323302-0033022130203330-3112100110300123-1310210111123123-3122221103132313-0322021321232201-3122311330133310)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-017.md#canonical-3113201100321032-2110322211311302-0001210313012211-3231233213313313-3102200133233230-2000333233123013-1333110321022331-3013320123303013)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-0123302000200031-3003123302123202-3000303113233003-1212302221222031-1302311201103300-1310330222002303-1120022012031313-0112232303300002"></a>

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

<a id="canonical-2130211221202310-2102322223223003-1002101030212303-0231013030133302-0311113302120312-0230333100032321-2012032002310003-0111030121221323"></a>

## Direct properties — virtual_site / 032331111313 / 3

<a id="canonical-1321330033133012-1233221220030003-0112211103011012-2233332112303100-3011003112131033-2122230302132320-2032030333311032-0322230321002110"></a>

<a id="canonical-2302003220202301-1301221123212301-3303313022232133-3311103031102101-0301010200320300-3323002113013012-0023023122223022-2330102112010113"></a>

## name property — virtual_site / 032331111313 / 4

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

<a id="canonical-1103330123133031-3023132311220333-2102100213003230-3322123031210213-3023332230220231-3322113300321002-0113333331000020-0113320303223223"></a>

<a id="canonical-0111301323110111-0033010012333310-1201230232300232-0311221010110132-3310031301030033-0033220011220123-2220223123131311-3322110123001122"></a>

## namespace property — virtual_site / 032331111313 / 5

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

<a id="canonical-3103000031321332-0231130232010110-2200221101223112-3021012101200031-2313303333110012-0301130023332231-3203211030020322-3023313002020122"></a>

<a id="canonical-3131223331113302-1221323103112323-1032132231221110-3103002133121000-1303022321332313-1313230131110331-3122021303101103-3230011231212033"></a>

## tenant property — virtual_site / 032331111313 / 6

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

<a id="canonical-2221203232012000-0300121200330310-3222311133310102-2131003220211221-3033022322112121-0300331000201003-0010231330231130-1211123011020033"></a>

## Next pages — virtual_site / 032331111313 / 7

- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](resources--workload--reference--group-017.md#canonical-3113201100321032-2110322211311302-0001210313012211-3231233213313313-3102200133233230-2000333233123013-1333110321022331-3013320123303013)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103220210020311-3012023013130001-2000121021330011-0010202323200100-0333113313213212-1313012031310011-0000001200311310-0122103020010202"></a>

## stateful_service.advertise_options.advertise_custom.ports — ports / 213202132323 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- stateful_service.advertise_options.advertise_custom.ports

<a id="canonical-2110313010033020-0020221230302220-2320320022122010-0231201122002231-3313000303103002-2213120300202302-3333311300012020-1312232001022110"></a>

Type: `"object"`. list nested block, Optional.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("http_loadbalancer",
    "tcp_loadbalancer")}
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
ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-2322233030111332-0231313323113233-0113113130123203-2230223132130203-2000202231122331-1310103202130302-3212120311122132-0100003131221333"></a>

## Direct properties — ports / 213202132323 / 3

- [http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002): complete subsection reference.

- [port](resources--workload--reference--group-021.md#canonical-2103101222313132-0030333210220102-2031201232312001-1103110230303213-2123210111312322-2333302111312301-1031201133012330-3303120000023223): complete subsection reference.

- [tcp_loadbalancer](resources--workload--reference--group-021.md#canonical-0020101203031230-1223300211330033-2123113022210223-3310333011133100-3330221301313310-3131021233011311-3221200100200012-0300221100330331): complete subsection reference.

<a id="canonical-1010231113200103-2132313100010123-0333131033313033-1003123220302113-1130033222303123-1110110221210330-0320330302033212-2321121223321330"></a>

## Next pages — ports / 213202132323 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [stateful_service.advertise_options.advertise_custom.ports.port](resources--workload--reference--group-021.md#canonical-2103101222313132-0030333210220102-2031201232312001-1103110230303213-2123210111312322-2333302111312301-1031201133012330-3303120000023223)
- [stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer](resources--workload--reference--group-021.md#canonical-0020101203031230-1223300211330033-2123113022210223-3310333011133100-3330221301313310-3131021233011311-3221200100200012-0300221100330331)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322200313001203-0021221323310002-3013233010031033-3331132203113302-2321110310023333-0203312012321110-1201203211233333-3032323132210023"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer — http_loadbalancer / 021333011310 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer

<a id="canonical-3233013221221131-2012133322101333-2333103311222030-1300332102121212-3213323301000303-2100212310212031-3102012233011233-3232213013123002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http loadbalancer.

Upstream description:

HTTP/HTTPS Load balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains"),
  validators.ConflictingObjectAttributes("default_route",
    "specific_routes"),
  validators.ConflictingObjectAttributes("http",
    "https"),
  validators.ConflictingObjectAttributes("http",
    "https_auto_cert"),
  validators.ConflictingObjectAttributes("https",
    "https_auto_cert")}
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
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]",
  "x-ves-oneof-field-route_choice": "[\"default_route\",\"specific_routes\"]"
}
```

Terraform syntax:

```terraform
http_loadbalancer {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131000232000212-0232120112133222-3122221022102011-3001012201200331-3122022333011332-3010331121001131-3223030121133211-1220031133213213"></a>

## Direct properties — http_loadbalancer / 021333011310 / 3

- [default_route](resources--workload--reference--group-017.md#canonical-2031112003330101-0310120302332023-3120021113320130-3110000310223201-1012000003321130-2312213320212011-0022300121030033-3100123121103201): complete subsection reference.

<a id="canonical-3122002032322020-1000311301022012-2113012331300210-3101110120322332-2103231112330231-2022313201323200-0133313033130311-1212231000203011"></a>

<a id="canonical-2223032222032031-1121210002120231-1333133122033222-1110221120123100-2333220303211100-1112102010122131-3030100210201233-2111022202000222"></a>

## domains property — http_loadbalancer / 021333011310 / 4

Type: `["list", "string"]`. Optional.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\` is invalid
Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the..

Upstream description:

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*.bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*.example.com\`\` will match \`\`bar.example.com\`\`
and \`\`baz-bar.example.com\`\` but not \`\`.example.com\`\`. The longest wildcards match first.
Wildcards must match a whole DNS label. E.g. \`\`\*.example.com\`\` and \*.bar.example.com are
valid, however \`\`\*bar.example.com\`\` or \`\`\*-bar.example.com\`\` is invalid

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [http](resources--workload--reference--group-018.md#canonical-0001001230203210-1010213221120331-1102200030010210-1101322011000303-0020221022232213-1231230310223023-2110212231011313-1321232000313021): complete subsection reference.

- [https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333): complete subsection reference.

- [https_auto_cert](resources--workload--reference--group-019.md#canonical-3032121232310131-1323102031301330-0021130303300211-3312120223313333-2332212203311131-0113312021100231-3031113230222323-1032033323320101): complete subsection reference.

- [specific_routes](resources--workload--reference--group-020.md#canonical-0102002011120221-0321102211211200-1233000013222030-0112300012022023-1023113003101102-0220010103202030-2311210030210102-2200103323322232): complete subsection reference.

<a id="canonical-1300101313320231-0203303330012022-0013322331322312-0213332303122231-1101312021112130-1122002001201122-1030220200000213-1313310000333002"></a>

## Next pages — http_loadbalancer / 021333011310 / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](resources--workload--reference--group-017.md#canonical-2031112003330101-0310120302332023-3120021113320130-3110000310223201-1012000003321130-2312213320212011-0022300121030033-3100123121103201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http](resources--workload--reference--group-018.md#canonical-0001001230203210-1010213221120331-1102200030010210-1101322011000303-0020221022232213-1231230310223023-2110212231011313-1321232000313021)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-018.md#canonical-0112020200111022-0013103013103331-1110301001233101-1300202101312213-3301031210112102-1221232301300211-0231000232130131-3132113303133333)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-019.md#canonical-3032121232310131-1323102031301330-0021130303300211-3312120223313333-2332212203311131-0113312021100231-3031113230222323-1032033323320101)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-020.md#canonical-0102002011120221-0321102211211200-1233000013222030-0112300012022023-1023113003101102-0220010103202030-2311210030210102-2200103323322232)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2031112003330101-0310120302332023-3120021113320130-3110000310223201-1012000003321130-2312213320212011-0022300121030033-3100123121103201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202011303331332-2303132010013121-0013322312232221-2322122030123212-2332300133130110-2002123002223213-2211330102232200-1023222131200021"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route — default_route / 103010212312 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [stateful_service](resources--workload--reference--group-017.md#canonical-2201223210333100-3220200132220221-0002032233312220-1001320301032122-1132131211310022-0213223031122210-3310210030323123-2301312132022333)
- [stateful_service.advertise_options](resources--workload--reference--group-017.md#canonical-1230002303322223-1032323033130121-3020022300122113-0112332301031221-1110131201023233-3032201112132133-0022212031002300-3303231132301323)
- [stateful_service.advertise_options.advertise_custom](resources--workload--reference--group-017.md#canonical-0203102001320123-1120130012130212-0111101220231333-3101331123203121-1110000311003232-1321231201202312-3200200021132112-0200303231330002)
- [stateful_service.advertise_options.advertise_custom.ports](resources--workload--reference--group-017.md#canonical-3111311322332303-1201030020021201-3133021233000313-2022333331131202-3221312032323033-3133120120113232-3330021120310331-2331303302032231)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route

<a id="canonical-2012112020211212-2002133331032331-0301122010221321-1233000233300003-0002212101002131-0210220123322310-1033033101022002-2112000101333121"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Upstream description:

Default route matching all APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_host_rewrite",
    "disable_host_rewrite"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("disable_host_rewrite",
    "host_rewrite")}
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
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

Terraform syntax:

```terraform
default_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212333013311332-1112132230113102-0201211300333021-1030330033100111-3012022301232110-2010322310002131-0330130102030302-0200223333330022"></a>

## Direct properties — default_route / 103010212312 / 3

- [auto_host_rewrite](resources--workload--reference--group-017.md#canonical-3102103313032210-0102301102011232-1303110012101212-3032312111101121-2212320022233233-2110332122303130-0322022331103012-2002002322210003): complete subsection reference.

- [disable_host_rewrite](resources--workload--reference--group-018.md#canonical-2233022211333231-1323213121201103-1312330322030123-0003203301103101-0332011031223232-2322233001313302-0312301223030330-1003213211032213): complete subsection reference.

<a id="canonical-3123011213032330-0310102133232303-2213001223032221-2132211311231011-2011231300311232-1120310313313203-2300202213032233-3011121100110321"></a>

<a id="canonical-1222130100102300-1012223121311032-1210200003120131-0301202000120302-1000202213223112-1022002002131332-2321000130101230-3222203112011133"></a>

## host_rewrite property — default_route / 103010212312 / 4

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-3312130013113202-3223221021303303-2002211321100330-2201113130200132-1100231231203312-3332022212100023-2010220200323103-1300003231333203"></a>

## Next pages — default_route / 103010212312 / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite](resources--workload--reference--group-017.md#canonical-3102103313032210-0102301102011232-1303110012101212-3032312111101121-2212320022233233-2110332122303130-0322022331103012-2002002322210003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite](resources--workload--reference--group-018.md#canonical-2233022211333231-1323213121201103-1312330322030123-0003203301103101-0332011031223232-2322233001313302-0312301223030330-1003213211032213)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-017.md#canonical-0313222022003221-2301103230203023-3120322210033100-2121220010332201-3133231021021220-1313001311100021-0313102021211212-0331322021202002)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3102103313032210-0102301102011232-1303110012101212-3032312111101121-2212320022233233-2110332122303130-0322022331103012-2002002322210003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
