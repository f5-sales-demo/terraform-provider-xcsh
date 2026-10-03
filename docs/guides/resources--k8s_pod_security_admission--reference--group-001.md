---
page_title: "xcsh_k8s_pod_security_admission reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_k8s_pod_security_admission reference."
---

# xcsh_k8s_pod_security_admission reference

<a id="canonical-2222333310123013-1221301313132132-3120203303322302-2023312113233133-3221130101113030-0021200233133103-2011300211002132-0212001303010330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023011220301133-1321202122011110-0110230301313300-2221023120332222-1203100020110020-0222111023013030-0003001211201213-1221011001323003"></a>

## Property reference — Property reference / 222303031023 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)
- Property reference

<a id="canonical-3112121222220032-3032121001313220-2203110121002121-3230312133030011-2203310103230201-3201111130330001-2123330320322232-1213230101013302"></a>

## Direct properties — Property reference / 222303031023 / 3

<a id="canonical-2212331131030013-1000322222331322-3211001032010101-0230330320023220-1111103322131311-1120302132012103-1122132233121122-3323223002331110"></a>

<a id="canonical-1230310123332101-1230113123120031-1223221003322032-2213230300330213-3320113030231220-1313033101023222-0202100310233000-3210213222311130"></a>

## annotations property — Property reference / 222303031023 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

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

<a id="canonical-2222330313303011-2010031001010120-3001211013200020-0002232211123330-1302323020000330-2113313122122301-1101211033200221-1131031321102220"></a>

<a id="canonical-1210312233101300-2203101212031312-2231111233302301-2303330313230120-3110200123133103-0301002111332333-2021101200002300-3311121311130221"></a>

## description property — Property reference / 222303031023 / 5

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2023321200301130-2033222101312322-0221233330002022-0221201213202211-1313323232322001-3222122332023101-0211211031013221-1133013031020112"></a>

<a id="canonical-3000122212001202-0123320231031223-2100223001122102-2022111323333331-2300120002001200-3231013300301121-1212033301121002-0231211112022233"></a>

## disable property — Property reference / 222303031023 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1311130321232032-2030032212201232-0112333131231332-3022212201220032-2133332323322032-3000020031031323-2330113120230033-0311030101020003"></a>

<a id="canonical-0212132202311103-0100301101313203-1220310112220023-3231102302210130-0033320132033123-2023322211103333-2303330001010020-2210233102332021"></a>

## ID property — Property reference / 222303031023 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2233203121100012-3132220032303212-2003133133232002-2101330010133323-3310000321301231-1301211033310032-2232323230123101-2303222121033101"></a>

<a id="canonical-0013200003231220-0311231200203011-2013123112120113-2210120230232132-3002233121320010-2122332330213102-3100103223312212-2312021032132101"></a>

## labels property — Property reference / 222303031023 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-1330220302000200-0010121020101313-1132202112133110-2313001132110122-2223222011030230-0200221330323020-0331211010330012-0221032231032333"></a>

<a id="canonical-2003203102211132-1222302313033212-3002121133213220-3013320330301112-3313130112200020-1010010110222232-0210320121211122-1221301231231332"></a>

## name property — Property reference / 222303031023 / 9

Type: `"string"`. Required.

Name of the K8S Pod Security Admission. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1200113103333203-0330031212100133-0223013113302100-3311133230033333-1211201003332020-1320322333223330-2203011101001222-0302021220132100"></a>

<a id="canonical-3102213112123220-3220022022201103-0223202330012211-3031301222331312-3103330212310222-0201012121200023-1103122120222200-2332123211300321"></a>

## namespace property — Property reference / 222303031023 / 10

Type: `"string"`. Optional, Computed.

Namespace for the K8S Pod Security Admission. The F5 XC API restricts this resource to the system
namespace; it defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
Validators: []validator.String{
  validators.NamespaceValidator(),
  stringvalidator.OneOf("system"),
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

- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032): complete subsection reference.

- [timeouts](resources--k8s_pod_security_admission--reference--group-001.md#canonical-3313222201332230-2320100123321123-3302030323300021-2321323013310010-0201132310033032-0001002200220012-1123310103120131-1130311210122220): complete subsection reference.

<a id="canonical-2011033001022101-0133300220013223-1220130312221310-2100202331110012-3312232330333212-0330313130022133-1032030023032020-3023200130302231"></a>

## All schema paths — Property reference / 222303031023 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2212331131030013-1000322222331322-3211001032010101-0230330320023220-1111103322131311-1120302132012103-1122132233121122-3323223002331110) |
| `description` | [description](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2222330313303011-2010031001010120-3001211013200020-0002232211123330-1302323020000330-2113313122122301-1101211033200221-1131031321102220) |
| `disable` | [disable](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2023321200301130-2033222101312322-0221233330002022-0221201213202211-1313323232322001-3222122332023101-0211211031013221-1133013031020112) |
| `id` | [ID](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1311130321232032-2030032212201232-0112333131231332-3022212201220032-2133332323322032-3000020031031323-2330113120230033-0311030101020003) |
| `labels` | [labels](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2233203121100012-3132220032303212-2003133133232002-2101330010133323-3310000321301231-1301211033310032-2232323230123101-2303222121033101) |
| `name` | [name](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1330220302000200-0010121020101313-1132202112133110-2313001132110122-2223222011030230-0200221330323020-0331211010330012-0221032231032333) |
| `namespace` | [namespace](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1200113103333203-0330031212100133-0223013113302100-3311133230033333-1211201003332020-1320322333223330-2203011101001222-0302021220132100) |
| `pod_security_admission_specs` | [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-3021321113010133-2010302332322132-1012323031013103-2011133321110010-2302333102231200-2202001320010310-0321223102022222-0230313131213231) |
| `pod_security_admission_specs.audit` | [pod_security_admission_specs.audit](resources--k8s_pod_security_admission--reference--group-001.md#canonical-0130302201130231-0010331203301212-2101210203111101-1020311321302123-0203100020303023-1320022111301331-0010331201130201-0320300212322231) |
| `pod_security_admission_specs.baseline` | [pod_security_admission_specs.baseline](resources--k8s_pod_security_admission--reference--group-001.md#canonical-0110302321313111-1111001333320102-0030203000213022-1300010002112322-1122200031313022-2133222012222330-2031233312303311-3302100021322223) |
| `pod_security_admission_specs.enforce` | [pod_security_admission_specs.enforce](resources--k8s_pod_security_admission--reference--group-001.md#canonical-3123211023130113-3210231013013321-3300201212130202-3321301032301102-2332101302102033-1202223221212322-3320300022310111-1102020130203210) |
| `pod_security_admission_specs.privileged` | [pod_security_admission_specs.privileged](resources--k8s_pod_security_admission--reference--group-001.md#canonical-3030013021013202-1211233000303000-1001223103213120-3012302221110223-2232013223131112-1330233212212101-2001221220210133-3123202122203021) |
| `pod_security_admission_specs.restricted` | [pod_security_admission_specs.restricted](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2120312330223003-0300121000031213-1332013101123330-2201301003220001-3000300122112000-0012133012003301-0201330133112113-2303312100113233) |
| `pod_security_admission_specs.warn` | [pod_security_admission_specs.warn](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1312203121202000-3112332320222021-1020220032300012-3120122022303103-3000003020122232-3233023331010112-2302200212331031-0230321233300033) |
| `timeouts` | [timeouts](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1130131200300303-3132131021102103-3333031020121020-3201202103333220-1323323011021010-1112120012111330-3223203101310032-3332213010111200) |
| `timeouts.create` | [timeouts.create](resources--k8s_pod_security_admission--reference--group-001.md#canonical-3001331111302113-1333030322202023-0223032221123131-2031230011010000-3313113102300322-2001330231020023-0322133120100323-3303231203123023) |
| `timeouts.delete` | [timeouts.delete](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1303202031212233-2222130111202003-0312122313303112-0313310330313102-1112311330310030-3332212221312021-2021100020122102-1031002310012032) |
| `timeouts.read` | [timeouts.read](resources--k8s_pod_security_admission--reference--group-001.md#canonical-0201320233132113-2101032233133111-3130223011021011-1101300022311233-1112210230021123-1003201323221122-0103312012100103-3003211203323331) |
| `timeouts.update` | [timeouts.update](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1103030203023323-3132301321301101-2201211202212203-3021232021201131-2221000120222212-2310122003111011-2020123010002031-0120313132211221) |

<a id="canonical-1211211213312023-0111133212203033-3011330233012012-0302131303230311-3230122123002010-2132200122012130-3112200332223001-2001231130002020"></a>

## Next pages — Property reference / 222303031023 / 12

- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032)
- [timeouts](resources--k8s_pod_security_admission--reference--group-001.md#canonical-3313222201332230-2320100123321123-3302030323300021-2321323013310010-0201132310033032-0001002200220012-1123310103120131-1130311210122220)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)

<a id="canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220102103021201-3220120333113331-2000313002213202-3030312103233231-3031211102100113-1130230333213212-1322230330232021-3000323201320030"></a>

## pod_security_admission_specs — pod_security_admission_specs / 333001023333 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2222333310123013-1221301313132132-3120203303322302-2023312113233133-3221130101113030-0021200233133103-2011300211002132-0212001303010330)
- pod_security_admission_specs

<a id="canonical-3021321113010133-2010302332322132-1012323031013103-2011133321110010-2302333102231200-2202001320010310-0321223102022222-0230313131213231"></a>

Type: `"object"`. list nested block, Optional.

K8s Pod Security Admission. Uniform Resource Identifier

Upstream description:

Uniform Resource Identifier

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("audit",
    "enforce"),
  validators.ConflictingListObjectAttributes("audit",
    "warn"),
  validators.ConflictingListObjectAttributes("baseline",
    "privileged"),
  validators.ConflictingListObjectAttributes("baseline",
    "restricted"),
  validators.ConflictingListObjectAttributes("enforce",
    "warn"),
  validators.ConflictingListObjectAttributes("privileged",
    "restricted")}
```

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

Terraform syntax:

```terraform
pod_security_admission_specs {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130032031322021-2100322000222101-3100113032312131-2123100130212320-3212010003123312-3020203220123203-0122220012133320-0133333311102331"></a>

## Direct properties — pod_security_admission_specs / 333001023333 / 3

- [audit](resources--k8s_pod_security_admission--reference--group-001.md#canonical-0111210113021032-1132303221231322-0131113233000101-0310133201123201-3303301122210123-3312133030303320-2322001003103232-1310213321310331): complete subsection reference.

- [baseline](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2133120101221330-1013013031332212-0111021021330212-0122112012131332-2003123222210011-3302203210312121-0002131200013010-1121202031110332): complete subsection reference.

- [enforce](resources--k8s_pod_security_admission--reference--group-001.md#canonical-0310333232313130-2020100002003332-3112203120102203-2303323102221110-1120023211013000-1003021113210131-0210022203233200-3333003010033212): complete subsection reference.

- [privileged](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2221201033113120-3102221030010231-3013121211311033-3310133213032232-2333013322013120-1100103021023312-2222300310113132-1000322132313333): complete subsection reference.

- [restricted](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1102101023232323-3300132013021330-0323320320111131-3211203101102112-2100332332133333-2232222311012203-2112121012100132-2010232022111133): complete subsection reference.

- [warn](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2323230313203131-2003022010302330-3323202121303002-1312133032233021-1300232123120033-0322310312321311-2331001030033123-3221303320030001): complete subsection reference.

<a id="canonical-2111230031133300-1213301222310133-3121331031323301-0321112300322220-3322121021131330-0022231202313101-0033233011121232-2210133132032212"></a>

## Next pages — pod_security_admission_specs / 333001023333 / 4

- [pod_security_admission_specs.audit](resources--k8s_pod_security_admission--reference--group-001.md#canonical-0111210113021032-1132303221231322-0131113233000101-0310133201123201-3303301122210123-3312133030303320-2322001003103232-1310213321310331)
- [pod_security_admission_specs.baseline](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2133120101221330-1013013031332212-0111021021330212-0122112012131332-2003123222210011-3302203210312121-0002131200013010-1121202031110332)
- [pod_security_admission_specs.enforce](resources--k8s_pod_security_admission--reference--group-001.md#canonical-0310333232313130-2020100002003332-3112203120102203-2303323102221110-1120023211013000-1003021113210131-0210022203233200-3333003010033212)
- [pod_security_admission_specs.privileged](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2221201033113120-3102221030010231-3013121211311033-3310133213032232-2333013322013120-1100103021023312-2222300310113132-1000322132313333)
- [pod_security_admission_specs.restricted](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1102101023232323-3300132013021330-0323320320111131-3211203101102112-2100332332133333-2232222311012203-2112121012100132-2010232022111133)
- [pod_security_admission_specs.warn](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2323230313203131-2003022010302330-3323202121303002-1312133032233021-1300232123120033-0322310312321311-2331001030033123-3221303320030001)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2222333310123013-1221301313132132-3120203303322302-2023312113233133-3221130101113030-0021200233133103-2011300211002132-0212001303010330)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)

<a id="canonical-0111210113021032-1132303221231322-0131113233000101-0310133201123201-3303301122210123-3312133030303320-2322001003103232-1310213321310331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133231012311332-2313220300212202-0111011212311211-0030131231003133-0302132320321312-2130122213202213-2000103332331330-2123020110221220"></a>

## pod_security_admission_specs.audit — audit / 123203310213 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2222333310123013-1221301313132132-3120203303322302-2023312113233133-3221130101113030-0021200233133103-2011300211002132-0212001303010330)
- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032)
- pod_security_admission_specs.audit

<a id="canonical-0130302201130231-0010331203301212-2101210203111101-1020311321302123-0203100020303023-1320022111301331-0010331201130201-0320300212322231"></a>

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
audit = {}
```

<a id="canonical-2122012331201102-0330213022323020-1203302202021302-3333321011320222-1203030323030223-2301233131131023-3032120130112032-2213030333200030"></a>

## Direct properties — audit / 123203310213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333311123113003-1130121013031033-1111333133221021-1120100013123320-1033122321310102-3202110313012310-3210000021001133-2321302231001331"></a>

## Next pages — audit / 123203310213 / 4

- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)

<a id="canonical-2133120101221330-1013013031332212-0111021021330212-0122112012131332-2003123222210011-3302203210312121-0002131200013010-1121202031110332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001103103101001-1102011123232303-1301112332312301-2113102323030002-3203213121021132-1300310201313210-3201000222133221-2010110213332001"></a>

## pod_security_admission_specs.baseline — baseline / 022203100010 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2222333310123013-1221301313132132-3120203303322302-2023312113233133-3221130101113030-0021200233133103-2011300211002132-0212001303010330)
- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032)
- pod_security_admission_specs.baseline

<a id="canonical-0110302321313111-1111001333320102-0030203000213022-1300010002112322-1122200031313022-2133222012222330-2031233312303311-3302100021322223"></a>

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
baseline = {}
```

<a id="canonical-0111032222220300-2022023010111110-3210022303121330-0032233012222021-0033303312122133-1201220313033131-1301221203211121-3211102222202332"></a>

## Direct properties — baseline / 022203100010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123121302200222-2201102020221033-0130210223112130-0202312311103130-1203203021000212-0302200321223202-2123232302223003-1201211330231030"></a>

## Next pages — baseline / 022203100010 / 4

- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)

<a id="canonical-0310333232313130-2020100002003332-3112203120102203-2303323102221110-1120023211013000-1003021113210131-0210022203233200-3333003010033212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300103220330111-2221310102000100-3110233311120031-3013032323100230-2133320231223221-1032303110023223-0201311101020013-1001013333321211"></a>

## pod_security_admission_specs.enforce — enforce / 100203113013 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2222333310123013-1221301313132132-3120203303322302-2023312113233133-3221130101113030-0021200233133103-2011300211002132-0212001303010330)
- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032)
- pod_security_admission_specs.enforce

<a id="canonical-3123211023130113-3210231013013321-3300201212130202-3321301032301102-2332101302102033-1202223221212322-3320300022310111-1102020130203210"></a>

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
enforce = {}
```

<a id="canonical-3313321012230222-0102131331020222-0010203120221030-0221301123321102-3321100032302000-1022122111121033-0211321023311131-2112313310032032"></a>

## Direct properties — enforce / 100203113013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113002032231121-1313001232011223-2120132332013130-0021030023321302-1330111301120001-0311230122213032-0320212102010323-2002003332101103"></a>

## Next pages — enforce / 100203113013 / 4

- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)

<a id="canonical-2221201033113120-3102221030010231-3013121211311033-3310133213032232-2333013322013120-1100103021023312-2222300310113132-1000322132313333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130132323312013-2013110203001203-3102002031122333-3103322033233112-2321010333022311-3003323200132133-0233000323003123-3001020132101112"></a>

## pod_security_admission_specs.privileged — privileged / 302321333011 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2222333310123013-1221301313132132-3120203303322302-2023312113233133-3221130101113030-0021200233133103-2011300211002132-0212001303010330)
- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032)
- pod_security_admission_specs.privileged

<a id="canonical-3030013021013202-1211233000303000-1001223103213120-3012302221110223-2232013223131112-1330233212212101-2001221220210133-3123202122203021"></a>

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
privileged = {}
```

<a id="canonical-2133322300323333-2122100013210131-1032023022011200-0200103021113113-2210132303113331-0131230300330011-2122001323113223-2101223332233232"></a>

## Direct properties — privileged / 302321333011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123100000031110-3310111013010320-3222213330221033-3022222030203301-0331030031322102-3313211121202210-2020002023203023-2102002312310113"></a>

## Next pages — privileged / 302321333011 / 4

- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)

<a id="canonical-1102101023232323-3300132013021330-0323320320111131-3211203101102112-2100332332133333-2232222311012203-2112121012100132-2010232022111133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111222032113133-0320311022101322-3000323300330332-2301003011210033-2230000313122110-0111133111032322-2100303321030220-1002022122010011"></a>

## pod_security_admission_specs.restricted — restricted / 220130333010 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2222333310123013-1221301313132132-3120203303322302-2023312113233133-3221130101113030-0021200233133103-2011300211002132-0212001303010330)
- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032)
- pod_security_admission_specs.restricted

<a id="canonical-2120312330223003-0300121000031213-1332013101123330-2201301003220001-3000300122112000-0012133012003301-0201330133112113-2303312100113233"></a>

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
restricted = {}
```

<a id="canonical-0202002331311130-1011230203302012-2032323212332003-3113130113233111-3221131130132000-1011133030133212-0023130321102110-2232011110302131"></a>

## Direct properties — restricted / 220130333010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302122002011221-3220310312033121-1131303002122211-2003012012232102-2132023303323020-0032210333121220-0102300222221202-0013102333013013"></a>

## Next pages — restricted / 220130333010 / 4

- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)

<a id="canonical-2323230313203131-2003022010302330-3323202121303002-1312133032233021-1300232123120033-0322310312321311-2331001030033123-3221303320030001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222002123001311-0223320333001012-3033002320122111-1122003303200321-2012223110022221-0223321002200232-2023010011220132-2013021110112213"></a>

## pod_security_admission_specs.warn — warn / 202320300111 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2222333310123013-1221301313132132-3120203303322302-2023312113233133-3221130101113030-0021200233133103-2011300211002132-0212001303010330)
- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032)
- pod_security_admission_specs.warn

<a id="canonical-1312203121202000-3112332320222021-1020220032300012-3120122022303103-3000003020122232-3233023331010112-2302200212331031-0230321233300033"></a>

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
warn = {}
```

<a id="canonical-1320013013303010-1112321031023001-2020033331033131-2301012033311003-2233333100101033-0211321030232310-3213202212021022-2313312221101122"></a>

## Direct properties — warn / 202320300111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010230233023221-1112133032133103-3232100102311030-2131013210103222-0023000001033303-1013000021002301-2321230010223322-2000100113101132"></a>

## Next pages — warn / 202320300111 / 4

- [pod_security_admission_specs](resources--k8s_pod_security_admission--reference--group-001.md#canonical-1000030230112200-1231331301101110-1303212112010002-0022032303300133-0310013011110210-1012020000003201-2013031303111200-2003222011220032)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)

<a id="canonical-3313222201332230-2320100123321123-3302030323300021-2321323013310010-0201132310033032-0001002200220012-1123310103120131-1130311210122220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222310033031201-0010312000031033-2310012300331202-2022102320312320-2120221302311213-2121021211102023-0301000023021120-0133021012303030"></a>

## timeouts — timeouts / 211331113002 / 2

Breadcrumbs:

- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)
- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2222333310123013-1221301313132132-3120203303322302-2023312113233133-3221130101113030-0021200233133103-2011300211002132-0212001303010330)
- timeouts

<a id="canonical-1130131200300303-3132131021102103-3333031020121020-3201202103333220-1323323011021010-1112120012111330-3223203101310032-3332213010111200"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103331220120020-0301311113022330-3330333012210213-2111323300203003-0120232032211033-1131323232312333-2230000310030200-3122003103320102"></a>

## Direct properties — timeouts / 211331113002 / 3

<a id="canonical-3001331111302113-1333030322202023-0223032221123131-2031230011010000-3313113102300322-2001330231020023-0322133120100323-3303231203123023"></a>

<a id="canonical-3012211110032203-3310300200101022-2031110122233100-1200002121120033-3322332121312310-1013122320220223-1100001200211322-3212020130101002"></a>

## create property — timeouts / 211331113002 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1303202031212233-2222130111202003-0312122313303112-0313310330313102-1112311330310030-3332212221312021-2021100020122102-1031002310012032"></a>

<a id="canonical-3010203213222202-0303331131320000-3110002113131023-1010212121212113-0223032210013333-3333230231300001-0320233310310312-3331330003033033"></a>

## delete property — timeouts / 211331113002 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0201320233132113-2101032233133111-3130223011021011-1101300022311233-1112210230021123-1003201323221122-0103312012100103-3003211203323331"></a>

<a id="canonical-0330320330203021-2323002131122101-1202210131113223-1002232103011332-0230312212323202-0031212311212212-3012332223223023-0023322201320020"></a>

## read property — timeouts / 211331113002 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1103030203023323-3132301321301101-2201211202212203-3021232021201131-2221000120222212-2310122003111011-2020123010002031-0120313132211221"></a>

<a id="canonical-0313002223002331-2032000121301122-1303130220010010-3003212113230310-3211312222020332-1030332300100312-2202230003130311-2322302030021211"></a>

## update property — timeouts / 211331113002 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3020230323231001-3120003201123223-1233131201323000-2133113320233133-1223022121023103-1312110122033010-1211023320202323-3032013313033112"></a>

## Next pages — timeouts / 211331113002 / 8

- [Property reference](resources--k8s_pod_security_admission--reference--group-001.md#canonical-2222333310123013-1221301313132132-3120203303322302-2023312113233133-3221130101113030-0021200233133103-2011300211002132-0212001303010330)
- [xcsh_k8s_pod_security_admission](../resources/k8s_pod_security_admission.md#canonical-0313313203122132-2133320121222233-2103022302032133-2001003113030010-1331332233310223-0122000131122000-1030032313322310-3110123301232023)
