---
page_title: "xcsh_virtual_k8s reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_virtual_k8s reference."
---

# xcsh_virtual_k8s reference

<a id="canonical-1212021120121211-3310201220003221-2122111020012333-1311201311130313-1133211310021231-2232220303133111-3220201300200202-2321121313013020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-2231213311220130-3130200112303103-1233222013120311-0300110031301022-3200010102220211-1202030313103211-3121033330120121-3020002013232100)
- Property reference

<a id="canonical-3033213101132031-1001133011003001-3113001210201103-3220321220110021-3100011011330321-0203212003222211-0212102033231130-0010030120100310"></a>

### Direct properties for `xcsh_virtual_k8s`

<a id="canonical-3223230311122123-0011323002130132-3320032123230132-0220331001303210-2321201131220220-1003303111100101-3310331230011323-2310212110020111"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

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

- [default_flavor_ref](resources--virtual_k8s--reference--group-001.md#canonical-0033111022212320-1223332202022301-1002321320131011-0100203321101010-0032222103132032-1202220133231330-1001222031302011-3102310231213200): complete subsection reference.

<a id="canonical-3033320213211302-3231220033331020-1321231323301333-2332001023100132-2221013330003312-3332320033311320-0120013200010330-1100123033333200"></a>

<a id="canonical-2021033222123002-1332311023001230-2312103030303311-0311200002223302-0001123002322031-2030212310312323-0233120102310201-2123123302112202"></a>

#### `description` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0202302112110313-2110002131230203-2220021031010220-3031200031021112-3302130300001331-3212021101332331-0123000103212010-0131033233223332"></a>

<a id="canonical-0313302201020003-2320233102012031-0031103000003030-1012120133300302-1011333232102202-2013310223112203-3110213210112330-1333100000012130"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

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

- [disabled](resources--virtual_k8s--reference--group-001.md#canonical-1000110211130130-2231200300130313-0012222033232021-3302311001032023-2030022233131330-2020303113200313-0130032100202113-2320003221333033): complete subsection reference.

<a id="canonical-1333130312000012-2212200120300202-1001122021300132-3222310213223000-3212231333122301-2301132033120122-3022020020010021-3220221312332212"></a>

<a id="canonical-3020011231110103-3023110112311110-2122121331223201-2111322132123233-2332223003333021-1003212223332211-3023202322311001-1200020023012323"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [isolated](resources--virtual_k8s--reference--group-001.md#canonical-0231101212002130-3302021000230012-0022123212210330-1033002222133301-0201201030203202-3321311310222323-0323330100103303-0202330120221111): complete subsection reference.

<a id="canonical-1302332130021322-2213113311322321-0000331123013021-1020210010221001-1213312312021233-3332331121002011-0112322322222101-2021133103212111"></a>

<a id="canonical-2001332031000023-2111032321033122-3200222233202300-1210022313101220-3221211313211000-0013300331110033-0101121221112022-1303212202122332"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-2310303302231200-1013313323012201-1031113220001300-0320220023322203-3212032130100130-2222231311013333-2221030312120220-0010110120212003"></a>

<a id="canonical-2023230332012101-2033030001021131-2100230120301200-2001120231023321-2110133300211113-3323112310133002-1202222301321011-1201010002012300"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Virtual K8S. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2112212031001300-2320310211103131-1302013013213011-3132231232333031-2231213002220221-3131201003103322-3320111030221112-0231222010113020"></a>

<a id="canonical-1202023233233333-0303322012311333-1202210222112332-2231130010222002-3012312310023013-0122031222111301-0131011203213200-1301130012110120"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Virtual K8S is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NamespaceValidator(),
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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [timeouts](resources--virtual_k8s--reference--group-001.md#canonical-1330233022100213-3312011233121323-1210121322013011-2110130300020231-3001313220232303-3233300211220313-1221020012032113-2110303232220231): complete subsection reference.

- [vsite_refs](resources--virtual_k8s--reference--group-001.md#canonical-2222010210220313-1320010212121331-1020223131311332-1130300311010222-3311212110121110-1330000211022023-0221333131312110-3010222302012031): complete subsection reference.

<a id="canonical-2302012201110223-0312011211122110-1232101023003102-3213033021102212-2223112022212230-0221023013310332-3221210330230310-1011201101320201"></a>

### All schema paths for `xcsh_virtual_k8s`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--virtual_k8s--reference--group-001.md#canonical-3223230311122123-0011323002130132-3320032123230132-0220331001303210-2321201131220220-1003303111100101-3310331230011323-2310212110020111) |
| `default_flavor_ref` | [default_flavor_ref](resources--virtual_k8s--reference--group-001.md#canonical-2103332020112202-0120001311030321-2333333003333300-1332332133130102-2333312032222121-3321110032313230-1100223113023012-3010000220130230) |
| `default_flavor_ref.name` | [default_flavor_ref.name](resources--virtual_k8s--reference--group-001.md#canonical-3121031013113030-2001211203230002-2013101132111032-0001221201012223-2231232133103320-3201301011123032-0222313030032111-0313103011031211) |
| `default_flavor_ref.namespace` | [default_flavor_ref.namespace](resources--virtual_k8s--reference--group-001.md#canonical-2303001020202231-3002130312032011-1000023031202301-0330133022002300-3010103230010213-0001013123133022-0232033201033000-3311000232223030) |
| `default_flavor_ref.tenant` | [default_flavor_ref.tenant](resources--virtual_k8s--reference--group-001.md#canonical-0222203312313110-2322332203311221-1010232132212123-2033302013000011-3221111230030111-3103011001313231-2330003303210311-0302320331132030) |
| `description` | [description](resources--virtual_k8s--reference--group-001.md#canonical-3033320213211302-3231220033331020-1321231323301333-2332001023100132-2221013330003312-3332320033311320-0120013200010330-1100123033333200) |
| `disable` | [disable](resources--virtual_k8s--reference--group-001.md#canonical-0202302112110313-2110002131230203-2220021031010220-3031200031021112-3302130300001331-3212021101332331-0123000103212010-0131033233223332) |
| `disabled` | [disabled](resources--virtual_k8s--reference--group-001.md#canonical-3233232311320233-3212122201221111-0203231311033302-1202001132200000-2020201133121002-0003131330023101-2120122312101123-2133201321122123) |
| `id` | [ID](resources--virtual_k8s--reference--group-001.md#canonical-1333130312000012-2212200120300202-1001122021300132-3222310213223000-3212231333122301-2301132033120122-3022020020010021-3220221312332212) |
| `isolated` | [isolated](resources--virtual_k8s--reference--group-001.md#canonical-2003112122231002-0313311133032011-2003330033220302-0322133013011223-0012323321200221-1130230020112002-1230011331003022-1301232031312032) |
| `labels` | [labels](resources--virtual_k8s--reference--group-001.md#canonical-1302332130021322-2213113311322321-0000331123013021-1020210010221001-1213312312021233-3332331121002011-0112322322222101-2021133103212111) |
| `name` | [name](resources--virtual_k8s--reference--group-001.md#canonical-2310303302231200-1013313323012201-1031113220001300-0320220023322203-3212032130100130-2222231311013333-2221030312120220-0010110120212003) |
| `namespace` | [namespace](resources--virtual_k8s--reference--group-001.md#canonical-2112212031001300-2320310211103131-1302013013213011-3132231232333031-2231213002220221-3131201003103322-3320111030221112-0231222010113020) |
| `timeouts` | [timeouts](resources--virtual_k8s--reference--group-001.md#canonical-0021221230123331-2310301013121103-1001322201212030-0213333010003301-0000312302300020-0021213211021301-0333202200231103-0033013221332212) |
| `timeouts.create` | [timeouts.create](resources--virtual_k8s--reference--group-001.md#canonical-1323300000032320-3121223232200201-0311330123111311-0031022330211220-3012313103221032-0031200233131333-2002111033111010-2311003130320322) |
| `timeouts.delete` | [timeouts.delete](resources--virtual_k8s--reference--group-001.md#canonical-2222330022223203-0131203330022221-2012223230120221-3230012323032223-1233000113311230-0310332310101101-0332332331231300-0210122231231310) |
| `timeouts.read` | [timeouts.read](resources--virtual_k8s--reference--group-001.md#canonical-1233322220230303-1011222310011023-0133300302112021-2213003211313310-1001002332000121-1313102221022332-1201333031011120-3132001101332223) |
| `timeouts.update` | [timeouts.update](resources--virtual_k8s--reference--group-001.md#canonical-2212013000133311-1230200233012132-2312003100102221-0031333013012301-1020123033211312-1302020233020223-0202320110011210-2133222001323010) |
| `vsite_refs` | [vsite_refs](resources--virtual_k8s--reference--group-001.md#canonical-2331021012111133-2030112331002300-0022331230120002-0320332011112333-0311323211113230-1213322320311203-2121213133322232-1132023020121013) |
| `vsite_refs.kind` | [vsite_refs.kind](resources--virtual_k8s--reference--group-001.md#canonical-1123211320012202-3321130232221231-0200211100031223-3121222203230030-1330310210321210-2120022322022300-2231031133203013-2311120032201101) |
| `vsite_refs.name` | [vsite_refs.name](resources--virtual_k8s--reference--group-001.md#canonical-3102030100022032-3013030130002221-0321132131130113-0023213333103021-1230333010222122-1010300133002113-0301300200023013-1203202200113202) |
| `vsite_refs.namespace` | [vsite_refs.namespace](resources--virtual_k8s--reference--group-001.md#canonical-3302332111212321-1133220331220313-0320111201313302-1323202210030201-3033001233323132-2023301020323113-2031221220322220-1201110311130322) |
| `vsite_refs.tenant` | [vsite_refs.tenant](resources--virtual_k8s--reference--group-001.md#canonical-0322330001231301-3202011102320321-2312200303111022-0022332002033133-0212222023203320-1313023203230003-2100232201233130-1221003232123303) |
| `vsite_refs.uid` | [vsite_refs.uid](resources--virtual_k8s--reference--group-001.md#canonical-2012230331013010-1311023021022313-2211301300101222-0022133230302331-0021130023103231-3001332113303320-1230230311221332-3020321131220111) |

<a id="canonical-0033111022212320-1223332202022301-1002321320131011-0100203321101010-0032222103132032-1202220133231330-1001222031302011-3102310231213200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_flavor_ref` properties

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-2231213311220130-3130200112303103-1233222013120311-0300110031301022-3200010102220211-1202030313103211-3121033330120121-3020002013232100)
- [Property reference](resources--virtual_k8s--reference--group-001.md#canonical-1212021120121211-3310201220003221-2122111020012333-1311201311130313-1133211310021231-2232220303133111-3220201300200202-2321121313013020)
- default_flavor_ref

<a id="canonical-2103332020112202-0120001311030321-2333333003333300-1332332133130102-2333312032222121-3321110032313230-1100223113023012-3010000220130230"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
default_flavor_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103112332221112-1203220020130233-2201003032132033-2020301233003311-1113113302132021-1311130200202020-2330213313300003-1022313132013233"></a>

### Direct properties for `default_flavor_ref`

<a id="canonical-3121031013113030-2001211203230002-2013101132111032-0001221201012223-2231232133103320-3201301011123032-0222313030032111-0313103011031211"></a>

#### `default_flavor_ref.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2303001020202231-3002130312032011-1000023031202301-0330133022002300-3010103230010213-0001013123133022-0232033201033000-3311000232223030"></a>

<a id="canonical-2303133333122312-3321222032302132-0101130132313210-3020213331121033-0012323122120103-2131023301310300-0303131333301313-3023132201332110"></a>

#### `default_flavor_ref.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0222203312313110-2322332203311221-1010232132212123-2033302013000011-3221111230030111-3103011001313231-2330003303210311-0302320331132030"></a>

<a id="canonical-3233232033132233-2331210001300101-0123300132311200-2012222231111200-3033011221021313-1332030003030210-1220113000200321-3221323332211011"></a>

#### `default_flavor_ref.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1000110211130130-2231200300130313-0012222033232021-3302311001032023-2030022233131330-2020303113200313-0130032100202113-2320003221333033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disabled` properties

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-2231213311220130-3130200112303103-1233222013120311-0300110031301022-3200010102220211-1202030313103211-3121033330120121-3020002013232100)
- [Property reference](resources--virtual_k8s--reference--group-001.md#canonical-1212021120121211-3310201220003221-2122111020012333-1311201311130313-1133211310021231-2232220303133111-3220201300200202-2321121313013020)
- disabled

<a id="canonical-3233232311320233-3212122201221111-0203231311033302-1202001132200000-2020201133121002-0003131330023101-2120122312101123-2133201321122123"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disabled, isolated\] Enable this option

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

OneOf alternatives in this subsection:

- [disabled](resources--virtual_k8s--reference--group-001.md#canonical-3233232311320233-3212122201221111-0203231311033302-1202001132200000-2020201133121002-0003131330023101-2120122312101123-2133201321122123)
- [isolated](resources--virtual_k8s--reference--group-001.md#canonical-2003112122231002-0313311133032011-2003330033220302-0322133013011223-0012323321200221-1130230020112002-1230011331003022-1301232031312032)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231101212002130-3302021000230012-0022123212210330-1033002222133301-0201201030203202-3321311310222323-0323330100103303-0202330120221111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `isolated` properties

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-2231213311220130-3130200112303103-1233222013120311-0300110031301022-3200010102220211-1202030313103211-3121033330120121-3020002013232100)
- [Property reference](resources--virtual_k8s--reference--group-001.md#canonical-1212021120121211-3310201220003221-2122111020012333-1311201311130313-1133211310021231-2232220303133111-3220201300200202-2321121313013020)
- isolated

<a id="canonical-2003112122231002-0313311133032011-2003330033220302-0322133013011223-0012323321200221-1130230020112002-1230011331003022-1301232031312032"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
isolated = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330233022100213-3312011233121323-1210121322013011-2110130300020231-3001313220232303-3233300211220313-1221020012032113-2110303232220231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-2231213311220130-3130200112303103-1233222013120311-0300110031301022-3200010102220211-1202030313103211-3121033330120121-3020002013232100)
- [Property reference](resources--virtual_k8s--reference--group-001.md#canonical-1212021120121211-3310201220003221-2122111020012333-1311201311130313-1133211310021231-2232220303133111-3220201300200202-2321121313013020)
- timeouts

<a id="canonical-0021221230123331-2310301013121103-1001322201212030-0213333010003301-0000312302300020-0021213211021301-0333202200231103-0033013221332212"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233101322310320-0210332023200102-3011103030221200-2333120121111031-3030212210022103-2301231333233100-1301303120002001-3213200320032110"></a>

### Direct properties for `timeouts`

<a id="canonical-1323300000032320-3121223232200201-0311330123111311-0031022330211220-3012313103221032-0031200233131333-2002111033111010-2311003130320322"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2222330022223203-0131203330022221-2012223230120221-3230012323032223-1233000113311230-0310332310101101-0332332331231300-0210122231231310"></a>

<a id="canonical-2023013203120111-1301133223202011-3123311220132201-2220032201302111-2301121130112210-0001302310321302-2120000103311000-3013101110211302"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1233322220230303-1011222310011023-0133300302112021-2213003211313310-1001002332000121-1313102221022332-1201333031011120-3132001101332223"></a>

<a id="canonical-3211202002310122-1221302113010033-1010000212203013-1322002002222201-1332102333130003-3201300211121333-1312322023201233-1030133302331223"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2212013000133311-1230200233012132-2312003100102221-0031333013012301-1020123033211312-1302020233020223-0202320110011210-2133222001323010"></a>

<a id="canonical-3001213332223100-1310310222303310-2111110323113302-0203320302232130-0111021330323303-2333033230222222-2122310011020201-0131223023322221"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2222010210220313-1320010212121331-1020223131311332-1130300311010222-3311212110121110-1330000211022023-0221333131312110-3010222302012031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vsite_refs` properties

Breadcrumbs:

- [xcsh_virtual_k8s](../resources/virtual_k8s.md#canonical-2231213311220130-3130200112303103-1233222013120311-0300110031301022-3200010102220211-1202030313103211-3121033330120121-3020002013232100)
- [Property reference](resources--virtual_k8s--reference--group-001.md#canonical-1212021120121211-3310201220003221-2122111020012333-1311201311130313-1133211310021231-2232220303133111-3220201300200202-2321121313013020)
- vsite_refs

<a id="canonical-2331021012111133-2030112331002300-0022331230120002-0320332011112333-0311323211113230-1213322320311203-2121213133322232-1132023020121013"></a>

Type: `"object"`. list nested block, Optional.

Reference to virtual-sites Default virtual-site of the Virtual K8s object. If no virtual-site is
specified in the Kubernetes API resource object annotations via F5 XC/virtual-sites, then this
virtual-site is used select sites on which to instantiate the Kubernetes API resource object.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
vsite_refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110011302201320-0311123332001301-3323111233233110-1002111333300110-1033212033211232-2220321313020333-2302232213210223-3013133203233232"></a>

### Direct properties for `vsite_refs`

<a id="canonical-1123211320012202-3321130232221231-0200211100031223-3121222203230030-1330310210321210-2120022322022300-2231031133203013-2311120032201101"></a>

#### `vsite_refs.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3102030100022032-3013030130002221-0321132131130113-0023213333103021-1230333010222122-1010300133002113-0301300200023013-1203202200113202"></a>

<a id="canonical-2331301012220313-2311022022022310-3023210033232112-1333031022112310-1001130031231020-3310103203332200-3231011323113323-3112310101300332"></a>

#### `vsite_refs.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3302332111212321-1133220331220313-0320111201313302-1323202210030201-3033001233323132-2023301020323113-2031221220322220-1201110311130322"></a>

<a id="canonical-2333322023321003-1122033013320120-2111112211322303-2110321301122303-0203212132202303-0210031313012233-0223100303133011-0303010323111333"></a>

#### `vsite_refs.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0322330001231301-3202011102320321-2312200303111022-0022332002033133-0212222023203320-1313023203230003-2100232201233130-1221003232123303"></a>

<a id="canonical-3010200003221330-0113213311010033-0031333310021312-3221321110010200-2031112211322012-1121031130001313-3003311120022001-1322112221100021"></a>

#### `vsite_refs.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2012230331013010-1311023021022313-2211301300101222-0022133230302331-0021130023103231-3001332113303320-1230230311221332-3020321131220111"></a>

<a id="canonical-0312132211200120-2300103003231320-2013203132021033-3220003030331231-1230032302200323-0203332031032302-2302013033001300-1021000210111221"></a>

#### `vsite_refs.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
