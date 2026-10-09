---
page_title: "xcsh_k8s_pod_security_admission reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_admission reference."
---

# xcsh_k8s_pod_security_admission reference

<a id="canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- Property reference

<a id="canonical-1023222223300130-2112201003121202-0133301121222033-3003301123232113-3030221231313311-2203200030132301-1131323133312033-1100201321310003"></a>

### Direct properties for `xcsh_k8s_pod_security_admission`

<a id="canonical-0223022203030111-1000333330123102-1210021202300001-1131221323203000-0321033022222222-0233131330013313-2133222322121000-2100121330332203"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-3223210300200210-0112333322010321-1212012022002233-1233102312331301-1220013132120000-2332331022321332-3301121203020301-3033312220013320"></a>

<a id="canonical-0303332300132322-0110022330333220-1132000032333312-2230111023222222-2022020303001330-3001031330022310-2130323312322220-0022110123131021"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the K8SPodSecurityAdmission.

Additional upstream details:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-0031110210332303-0331223330322231-0101120133203112-0101213033020003-1100033333222202-3013002010102012-2202001330203223-1221332122112033"></a>

<a id="canonical-2322302000033020-0003220332111011-1222122103213122-2032210002331100-2133133103333331-2302222110203011-3032200233120323-3112020103233300"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2232122111332133-0033023322333123-2300022121333022-1112023321222232-1133022202232000-2301010022300302-2320330221030310-2123333032103310"></a>

<a id="canonical-1110311033201221-1133331322301302-2222221102033232-0322031330302021-3110002032321031-1113110110133332-1021213300022012-3222300213220231"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2031122033300322-2120312310311313-0310013003031033-0333122000311101-0302231023332333-1331201213012101-3221120332103220-2130210221102103"></a>

<a id="canonical-2332020330232010-1201233031100122-2223302310121331-2121231333203103-0011302322323331-2002203020023210-1013311112201302-2102223022103323"></a>

#### `name` property

Type: `"string"`. Required.

Name of the K8SPodSecurityAdmission.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2303330312200033-1210101313120001-1222201321120003-0213122233110300-1302121202331011-0112332311120323-2013122102020101-0103111311211033"></a>

<a id="canonical-0213312100133120-2300122320001221-0113202222331330-3203032002230212-1212302100112202-3233031201300132-3023020130302030-1211203110332032"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace where the K8SPodSecurityAdmission exists.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331): complete subsection reference.

<a id="canonical-1011120212131310-2103221211323312-1230003100110200-2113212330300001-2212311120033301-3020333032201300-1000123122001211-1232323231201323"></a>

### All schema paths for `xcsh_k8s_pod_security_admission`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-0223022203030111-1000333330123102-1210021202300001-1131221323203000-0321033022222222-0233131330013313-2133222322121000-2100121330332203) |
| `description` | [description](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-3223210300200210-0112333322010321-1212012022002233-1233102312331301-1220013132120000-2332331022321332-3301121203020301-3033312220013320) |
| `id` | [ID](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-0031110210332303-0331223330322231-0101120133203112-0101213033020003-1100033333222202-3013002010102012-2202001330203223-1221332122112033) |
| `labels` | [labels](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-2232122111332133-0033023322333123-2300022121333022-1112023321222232-1133022202232000-2301010022300302-2320330221030310-2123333032103310) |
| `name` | [name](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-2031122033300322-2120312310311313-0310013003031033-0333122000311101-0302231023332333-1331201213012101-3221120332103220-2130210221102103) |
| `namespace` | [namespace](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-2303330312200033-1210101313120001-1222201321120003-0213122233110300-1302121202331011-0112332311120323-2013122102020101-0103111311211033) |
| `pod_security_admission_specs` | [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-0211002001112220-3132211113303220-2021331102232222-1323213130233213-0222121233011200-0203111221032010-3023203012210231-2101300001020100) |
| `pod_security_admission_specs.audit` | [pod_security_admission_specs.audit](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-2333012201021111-1323030032233110-2323121213020313-0133203112120131-3200100232232223-1312211331333211-2321102320311103-1201021121133313) |
| `pod_security_admission_specs.baseline` | [pod_security_admission_specs.baseline](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-3223101323213000-2100111331312123-3101100201001200-2223133031320302-3122321132202203-3322202131013030-3103011303102210-2031011231311202) |
| `pod_security_admission_specs.enforce` | [pod_security_admission_specs.enforce](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-0012033121230310-0213110312301111-0013223311233013-1030101212121201-0113232002331120-1031202110133222-3031321010100222-0300021132212133) |
| `pod_security_admission_specs.privileged` | [pod_security_admission_specs.privileged](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-3000133323002332-2023102202311032-0003322210223201-2320321121001131-1333001113310301-3211220010220111-2011012301131102-3303101211023203) |
| `pod_security_admission_specs.restricted` | [pod_security_admission_specs.restricted](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-0031321001203203-3020331033011232-2101311010213210-0220010212002232-0230012032230132-3000011012233203-0130020013300230-0213132023210023) |
| `pod_security_admission_specs.warn` | [pod_security_admission_specs.warn](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-0331312231130020-2230121020112003-0023101233302221-1230213001322023-2233011301212202-1013031111100201-3133020030230032-3013033323100012) |

<a id="canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `pod_security_admission_specs` properties

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121)
- pod_security_admission_specs

<a id="canonical-0211002001112220-3132211113303220-2021331102232222-1323213130233213-0222121233011200-0203111221032010-3023203012210231-2101300001020100"></a>

Type: `"list"`. Computed.

K8s Pod Security Admission. Uniform Resource Identifier

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0020100212002203-1020202111003221-1323323001213301-1333101033223000-0232321303332022-2230233202300302-3301110202122212-0111211012113331"></a>

### Direct properties for `pod_security_admission_specs`

- [audit](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-3110300103022003-0313333120023000-2311223031222320-0120231220312003-0103012303301311-3311112233131020-0103203211002210-1303202012230311): complete subsection reference.

- [baseline](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-2100003232213130-2020000213311222-3230311300222210-2010323103331223-1000200311202003-3220023112122021-2121112131003121-3011032333333001): complete subsection reference.

- [enforce](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1221031123102221-0011200230132322-3312012022232210-0002033031131313-3302213112112213-1111330210332003-1012022112331233-2013001220031113): complete subsection reference.

- [privileged](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-2302113310211130-2000220110323111-3322132301232100-1012231022321011-3020232220212221-1021330113011130-3221201203110213-0211120310202210): complete subsection reference.

- [restricted](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1021333231221222-0011233232010123-1122100122001131-2000100230223101-1201113222132322-2311223223030020-2202222001123322-1332031302001031): complete subsection reference.

- [warn](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1302301231103202-2002110012323220-2320121233303312-1111123000323231-0103131133223120-1003101113301213-1123112232222201-1011213022110233): complete subsection reference.

<a id="canonical-3110300103022003-0313333120023000-2311223031222320-0120231220312003-0103012303301311-3311112233131020-0103203211002210-1303202012230311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `pod_security_admission_specs.audit` properties

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- pod_security_admission_specs.audit

<a id="canonical-2333012201021111-1323030032233110-2323121213020313-0133203112120131-3200100232232223-1312211331333211-2321102320311103-1201021121133313"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100003232213130-2020000213311222-3230311300222210-2010323103331223-1000200311202003-3220023112122021-2121112131003121-3011032333333001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `pod_security_admission_specs.baseline` properties

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- pod_security_admission_specs.baseline

<a id="canonical-3223101323213000-2100111331312123-3101100201001200-2223133031320302-3122321132202203-3322202131013030-3103011303102210-2031011231311202"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221031123102221-0011200230132322-3312012022232210-0002033031131313-3302213112112213-1111330210332003-1012022112331233-2013001220031113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `pod_security_admission_specs.enforce` properties

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- pod_security_admission_specs.enforce

<a id="canonical-0012033121230310-0213110312301111-0013223311233013-1030101212121201-0113232002331120-1031202110133222-3031321010100222-0300021132212133"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302113310211130-2000220110323111-3322132301232100-1012231022321011-3020232220212221-1021330113011130-3221201203110213-0211120310202210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `pod_security_admission_specs.privileged` properties

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- pod_security_admission_specs.privileged

<a id="canonical-3000133323002332-2023102202311032-0003322210223201-2320321121001131-1333001113310301-3211220010220111-2011012301131102-3303101211023203"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021333231221222-0011233232010123-1122100122001131-2000100230223101-1201113222132322-2311223223030020-2202222001123322-1332031302001031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `pod_security_admission_specs.restricted` properties

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- pod_security_admission_specs.restricted

<a id="canonical-0031321001203203-3020331033011232-2101311010213210-0220010212002232-0230012032230132-3000011012233203-0130020013300230-0213132023210023"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302301231103202-2002110012323220-2320121233303312-1111123000323231-0103131133223120-1003101113301213-1123112232222201-1011213022110233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `pod_security_admission_specs.warn` properties

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- pod_security_admission_specs.warn

<a id="canonical-0331312231130020-2230121020112003-0023101233302221-1230213001322023-2233011301212202-1013031111100201-3133020030230032-3013033323100012"></a>

Type: `["object", {}]`. Computed.

Enable this option

Additional upstream details:

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

This is an empty object or choice marker. It has no direct properties.
