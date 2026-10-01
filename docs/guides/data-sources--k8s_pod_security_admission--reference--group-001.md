---
page_title: "xcsh_k8s_pod_security_admission reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_admission reference."
---

# xcsh_k8s_pod_security_admission reference

<a id="canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023222223300130-2112201003121202-0133301121222033-3003301123232113-3030221231313311-2203200030132301-1131323133312033-1100201321310003"></a>

## Property reference — Property reference / 211100303022 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- Property reference

<a id="canonical-0303332300132322-0110022330333220-1132000032333312-2230111023222222-2022020303001330-3001031330022310-2130323312322220-0022110123131021"></a>

## Direct properties — Property reference / 211100303022 / 3

<a id="canonical-0223022203030111-1000333330123102-1210021202300001-1131221323203000-0321033022222222-0233131330013313-2133222322121000-2100121330332203"></a>

<a id="canonical-2322302000033020-0003220332111011-1222122103213122-2032210002331100-2133133103333331-2302222110203011-3032200233120323-3112020103233300"></a>

## annotations property — Property reference / 211100303022 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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

<a id="canonical-1110311033201221-1133331322301302-2222221102033232-0322031330302021-3110002032321031-1113110110133332-1021213300022012-3222300213220231"></a>

## description property — Property reference / 211100303022 / 5

Type: `"string"`. Computed.

Description of the K8SPodSecurityAdmission.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2332020330232010-1201233031100122-2223302310121331-2121231333203103-0011302322323331-2002203020023210-1013311112201302-2102223022103323"></a>

## ID property — Property reference / 211100303022 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2232122111332133-0033023322333123-2300022121333022-1112023321222232-1133022202232000-2301010022300302-2320330221030310-2123333032103310"></a>

<a id="canonical-0213312100133120-2300122320001221-0113202222331330-3203032002230212-1212302100112202-3233031201300132-3023020130302030-1211203110332032"></a>

## labels property — Property reference / 211100303022 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

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

<a id="canonical-1011120212131310-2103221211323312-1230003100110200-2113212330300001-2212311120033301-3020333032201300-1000123122001211-1232323231201323"></a>

## name property — Property reference / 211100303022 / 8

Type: `"string"`. Required.

Name of the K8SPodSecurityAdmission.

Upstream description:

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2303330312200033-1210101313120001-1222201321120003-0213122233110300-1302121202331011-0112332311120323-2013122102020101-0103111311211033"></a>

<a id="canonical-2023001200010201-3222330222323323-3123220320122332-2303233020133000-0213310330303033-2302111012133010-1131100011221301-3211200132202222"></a>

## namespace property — Property reference / 211100303022 / 9

Type: `"string"`. Optional, Computed.

Namespace where the K8SPodSecurityAdmission exists.

Upstream description:

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

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331): complete subsection reference.

<a id="canonical-1003103310012030-0300113210003102-0200120033001203-2320032131112003-3012330331012113-2101232231111123-3101120133010213-3003000201233300"></a>

## All schema paths — Property reference / 211100303022 / 10

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

<a id="canonical-0131103203313330-0101312033102210-2101321203200013-0221221213212303-1310330221202131-2331012221203133-0301233331230310-0302300231003333"></a>

## Next pages — Property reference / 211100303022 / 11

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)

<a id="canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020100212002203-1020202111003221-1323323001213301-1333101033223000-0232321303332022-2230233202300302-3301110202122212-0111211012113331"></a>

## pod_security_admission_specs — pod_security_admission_specs / 203321220203 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121)
- pod_security_admission_specs

<a id="canonical-0211002001112220-3132211113303220-2021331102232222-1323213130233213-0222121233011200-0203111221032010-3023203012210231-2101300001020100"></a>

Type: `"list"`. Computed.

K8s Pod Security Admission. Uniform Resource Identifier

Upstream description:

Uniform Resource Identifier

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

<a id="canonical-0332012031022132-1231101120003120-0103022202001102-2210122213023301-0003013032303213-3011113033210301-1312320010300032-0013012120311123"></a>

## Direct properties — pod_security_admission_specs / 203321220203 / 3

- [audit](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-3110300103022003-0313333120023000-2311223031222320-0120231220312003-0103012303301311-3311112233131020-0103203211002210-1303202012230311): complete subsection reference.

- [baseline](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-2100003232213130-2020000213311222-3230311300222210-2010323103331223-1000200311202003-3220023112122021-2121112131003121-3011032333333001): complete subsection reference.

- [enforce](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1221031123102221-0011200230132322-3312012022232210-0002033031131313-3302213112112213-1111330210332003-1012022112331233-2013001220031113): complete subsection reference.

- [privileged](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-2302113310211130-2000220110323111-3322132301232100-1012231022321011-3020232220212221-1021330113011130-3221201203110213-0211120310202210): complete subsection reference.

- [restricted](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1021333231221222-0011233232010123-1122100122001131-2000100230223101-1201113222132322-2311223223030020-2202222001123322-1332031302001031): complete subsection reference.

- [warn](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1302301231103202-2002110012323220-2320121233303312-1111123000323231-0103131133223120-1003101113301213-1123112232222201-1011213022110233): complete subsection reference.

<a id="canonical-0120132321322323-0031303010220122-2220231201000032-1302333020010313-0313113013221231-0331033222132001-3122032022103031-0320213303002010"></a>

## Next pages — pod_security_admission_specs / 203321220203 / 4

- [pod_security_admission_specs.audit](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-3110300103022003-0313333120023000-2311223031222320-0120231220312003-0103012303301311-3311112233131020-0103203211002210-1303202012230311)
- [pod_security_admission_specs.baseline](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-2100003232213130-2020000213311222-3230311300222210-2010323103331223-1000200311202003-3220023112122021-2121112131003121-3011032333333001)
- [pod_security_admission_specs.enforce](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1221031123102221-0011200230132322-3312012022232210-0002033031131313-3302213112112213-1111330210332003-1012022112331233-2013001220031113)
- [pod_security_admission_specs.privileged](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-2302113310211130-2000220110323111-3322132301232100-1012231022321011-3020232220212221-1021330113011130-3221201203110213-0211120310202210)
- [pod_security_admission_specs.restricted](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1021333231221222-0011233232010123-1122100122001131-2000100230223101-1201113222132322-2311223223030020-2202222001123322-1332031302001031)
- [pod_security_admission_specs.warn](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1302301231103202-2002110012323220-2320121233303312-1111123000323231-0103131133223120-1003101113301213-1123112232222201-1011213022110233)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)

<a id="canonical-3110300103022003-0313333120023000-2311223031222320-0120231220312003-0103012303301311-3311112233131020-0103203211002210-1303202012230311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331210221022311-0030212322332013-0321111003201333-2101121123213212-1113220200100101-0202312313321012-1120222222030201-2012212330310201"></a>

## pod_security_admission_specs.audit — audit / 312202131122 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- pod_security_admission_specs.audit

<a id="canonical-2333012201021111-1323030032233110-2323121213020313-0133203112120131-3200100232232223-1312211331333211-2321102320311103-1201021121133313"></a>

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

<a id="canonical-0221313031123201-0130133313021322-1120120001331323-3210121320220310-3321201133200010-1222211133131221-3232110201223002-3011032000331320"></a>

## Direct properties — audit / 312202131122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033223001113301-1211020032330232-2021033131211100-3120312200223012-3021232121132203-1303231330111321-0212233210201301-2202022001313022"></a>

## Next pages — audit / 312202131122 / 4

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)

<a id="canonical-2100003232213130-2020000213311222-3230311300222210-2010323103331223-1000200311202003-3220023112122021-2121112131003121-3011032333333001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122122031333013-3230211230220103-2212323311033322-2231110003101011-0211003000312222-3221333111103131-1320112000131120-1321130112001300"></a>

## pod_security_admission_specs.baseline — baseline / 030200120121 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- pod_security_admission_specs.baseline

<a id="canonical-3223101323213000-2100111331312123-3101100201001200-2223133031320302-3122321132202203-3322202131013030-3103011303102210-2031011231311202"></a>

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

<a id="canonical-0222323021213121-3131232021231111-2013301033212113-0202031010111122-0221100330312231-3022330033110120-2022232030322310-2010032131011232"></a>

## Direct properties — baseline / 030200120121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032202202101013-0300303202313333-0010311321133101-3012030111222233-0120123001222323-0230102210031021-2211112123213012-3301231232310111"></a>

## Next pages — baseline / 030200120121 / 4

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)

<a id="canonical-1221031123102221-0011200230132322-3312012022232210-0002033031131313-3302213112112213-1111330210332003-1012022112331233-2013001220031113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111200302301112-2120331232201012-1312230320100031-1332130133200102-0131331331211200-2033110120331033-0020010132332131-3012133230233332"></a>

## pod_security_admission_specs.enforce — enforce / 202013223103 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- pod_security_admission_specs.enforce

<a id="canonical-0012033121230310-0213110312301111-0013223311233013-1030101212121201-0113232002331120-1031202110133222-3031321010100222-0300021132212133"></a>

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

<a id="canonical-1100222011000123-3220220330103002-2010003223031222-3010201022221012-0113312211232010-1233130313302231-2301221002200313-0103011023121321"></a>

## Direct properties — enforce / 202013223103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120311001131203-0100232230100133-2121033110002223-1100122212330210-3023310013010230-3330200121212130-2202000113011111-1100223333122132"></a>

## Next pages — enforce / 202013223103 / 4

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)

<a id="canonical-2302113310211130-2000220110323111-3322132301232100-1012231022321011-3020232220212221-1021330113011130-3221201203110213-0211120310202210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211202000322300-1020101113011020-0223223023102011-0203313110210303-1022112331113203-3102133220222221-0111331023322103-0000010211231000"></a>

## pod_security_admission_specs.privileged — privileged / 300113232020 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- pod_security_admission_specs.privileged

<a id="canonical-3000133323002332-2023102202311032-0003322210223201-2320321121001131-1333001113310301-3211220010220111-2011012301131102-3303101211023203"></a>

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

<a id="canonical-0310210213201222-1210202313313202-0120000001010210-1013032330312221-1100111202302133-0323211113300123-0120323100020320-2332312301022211"></a>

## Direct properties — privileged / 300113232020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201003021001011-1111033013101310-1030001333313130-3100321232123323-2002320331311021-3210223031021112-2323103203330202-1211130021002111"></a>

## Next pages — privileged / 300113232020 / 4

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)

<a id="canonical-1021333231221222-0011233232010123-1122100122001131-2000100230223101-1201113222132322-2311223223030020-2202222001123322-1332031302001031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310130002321231-2123002023130022-3000300020322123-1211102020223120-3203200210302203-1112121331023221-2113001103201201-1000000313210213"></a>

## pod_security_admission_specs.restricted — restricted / 031303312311 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- pod_security_admission_specs.restricted

<a id="canonical-0031321001203203-3020331033011232-2101311010213210-0220010212002232-0230012032230132-3000011012233203-0130020013300230-0213132023210023"></a>

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

<a id="canonical-2002332323122201-1312300300010230-2201101301233100-0011002123300003-1333110021303120-3011330012002001-1021132212302331-1213132132011020"></a>

## Direct properties — restricted / 031303312311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012310121313003-2320332110001123-2133202101302000-2001132322220321-3030301132022320-1301123200010003-1323021003131221-1101113322002101"></a>

## Next pages — restricted / 031303312311 / 4

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)

<a id="canonical-1302301231103202-2002110012323220-2320121233303312-1111123000323231-0103131133223120-1003101113301213-1123112232222201-1011213022110233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003233110321131-2002011331302030-3231313231330003-3303212333210133-0223131333331221-3321200223320222-1303012112210303-3321133330223231"></a>

## pod_security_admission_specs.warn — warn / 031231023331 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
- [Property reference](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1133310000110012-0223210302002202-0012203133212011-3211011320033330-3222003300220023-3233312321311321-3120003300320230-2001100312020121)
- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- pod_security_admission_specs.warn

<a id="canonical-0331312231130020-2230121020112003-0023101233302221-1230213001322023-2233011301212202-1013031111100201-3133020030230032-3013033323100012"></a>

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

<a id="canonical-3013102031202302-1321021311200330-0001321002020323-1200030221303023-1222023112213332-3130232312033101-0003202013120002-3232220210122220"></a>

## Direct properties — warn / 031231023331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130301031223101-0123331310320023-3321112100203133-3133213031032001-0232012132013333-1121110123022200-3020202120320311-0233322110213103"></a>

## Next pages — warn / 031231023331 / 4

- [pod_security_admission_specs](data-sources--k8s_pod_security_admission--reference--group-001.md#canonical-1211112021002321-2221300323023323-3200102331112201-0013023312200103-1031010012002313-3231212323301230-1011020323221331-3120332333030331)
- [xcsh_k8s_pod_security_admission](../data-sources/k8s_pod_security_admission.md#canonical-2020223211310112-2023100102001010-2330201001133013-3032112120032203-0230132011212200-1221110333231223-2011111232310022-0131002111012013)
