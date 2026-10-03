---
page_title: "xcsh_cloud_elastic_ip reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_elastic_ip reference."
---

# xcsh_cloud_elastic_ip reference

<a id="canonical-3000332101332101-3033231112032333-3203132203302322-1221122003220232-1112221212310323-1330110330201000-0101323211230023-3313311003132210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302333010022113-1311321121002210-3031033213313310-3122322211120133-1201102312303312-0023211131230020-0021230310332101-3322313311212330"></a>

## Property reference — Property reference / 112322111102 / 2

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-1012233313021212-0003222033112111-1303213323303313-1020020302213310-1211110033020313-2122312122003031-2320202301023301-3102202323122130)
- Property reference

<a id="canonical-1312120213000200-0122122030201303-3131133231202232-0302101321310111-1302223212002011-0332122001313120-0231031313332200-0013030031300323"></a>

## Direct properties — Property reference / 112322111102 / 3

<a id="canonical-2031223121012120-3032230123200321-1221213102032230-0113332132022010-1211211333020221-0310221321013301-1011121221103333-3030103213013121"></a>

<a id="canonical-2030231000013112-3200200333310313-0201210331101333-2100231300221032-2331131100230123-1321202111210011-3011303001321020-0022220022313321"></a>

## annotations property — Property reference / 112322111102 / 4

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

<a id="canonical-2130003022201202-2230110200332020-1333201311112111-0010222230023321-3302221032002320-1202120012311131-3122311300212031-3022213200012130"></a>

<a id="canonical-1003230013222010-3330113122013011-2322012112220012-2320123311113023-3111003203100231-1321103301322220-0033133111122312-2020301121211113"></a>

## description property — Property reference / 112322111102 / 5

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

<a id="canonical-2321200011033103-1003321010231113-3032113331323220-3022031222311133-0213032021212130-3000132001301013-2120321213010323-3012222011233130"></a>

<a id="canonical-3213311113122212-3223102323133222-2012233031311300-1320300313323012-0030113102033012-0212110020123211-2121003000312131-0233223112110323"></a>

## disable property — Property reference / 112322111102 / 6

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

<a id="canonical-1313022001231131-0132330013031122-1332100030323322-3232221321100001-1111111320313322-1123132222110223-3303311232020020-1030331302331300"></a>

<a id="canonical-0020223031130102-2320122233203212-0233111333302012-1210133100013113-0123213220222232-0030021301303102-3101020212200213-2011121113132232"></a>

## ID property — Property reference / 112322111102 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3033303323013233-2332203130012203-0012121332123211-2223202101230001-3023212222320123-1030321010032101-1232232201303103-2310111123303031"></a>

<a id="canonical-2233001011123120-2231022012312322-0010230132323222-2202100223212122-0310133210222131-0022013231103111-2312031101000212-1300111303323003"></a>

## item_count property — Property reference / 112322111102 / 8

Type: `"number"`. Required.

Number of Elastic Ips / Public Ips associated with this object per Node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 8),
}
```

<a id="canonical-2103312231132200-2100033032331201-3011212132130301-0200033130231001-3333322213032000-1333211313002313-2131311002100201-2133320101320212"></a>

<a id="canonical-1211200333113102-1110232010000222-3330002330201200-0212002233032302-0231201210011001-0232220101333131-3011013022111202-2001301231112020"></a>

## labels property — Property reference / 112322111102 / 9

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

<a id="canonical-0222033330010213-0200331330133231-0022300310033030-2330211022322000-3332112131323313-1032110010331110-3231223230003030-2022003111322131"></a>

<a id="canonical-0032232323323120-1210221333022013-2312212311202231-0111231301022003-2302303121220300-1232202110023332-3033201223303331-1013133010132102"></a>

## name property — Property reference / 112322111102 / 10

Type: `"string"`. Required.

Name of the Cloud Elastic IP. Must be unique within the namespace.

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

<a id="canonical-2322123102102113-1203323331230132-3122321302223003-2313222330202131-2012202223203102-2020332213130212-1211000030021112-1010003013223230"></a>

<a id="canonical-2113100331133110-3102201103011031-1033200133002111-1022113230131010-0112333321333310-3300101222301103-3122211120202302-2020012313223102"></a>

## namespace property — Property reference / 112322111102 / 11

Type: `"string"`. Required.

Namespace where the Cloud Elastic IP is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
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

- [site_ref](resources--cloud_elastic_ip--reference--group-001.md#canonical-2102121212113212-1112222023302311-1111213230011002-2022232101311132-3321202123233120-2020113123321112-1200200220301033-3332330302121120): complete subsection reference.

- [timeouts](resources--cloud_elastic_ip--reference--group-001.md#canonical-0333031312331000-3300132311020030-3221133003000211-3031332022113301-0202021002223002-1103112311021221-2310003223220123-0300302133023103): complete subsection reference.

<a id="canonical-0033113000022102-0331323102313213-2021323313233221-2132333203130002-1102200102330322-3013233312101133-3302011320033013-0001100301311022"></a>

## All schema paths — Property reference / 112322111102 / 12

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cloud_elastic_ip--reference--group-001.md#canonical-2031223121012120-3032230123200321-1221213102032230-0113332132022010-1211211333020221-0310221321013301-1011121221103333-3030103213013121) |
| `description` | [description](resources--cloud_elastic_ip--reference--group-001.md#canonical-2130003022201202-2230110200332020-1333201311112111-0010222230023321-3302221032002320-1202120012311131-3122311300212031-3022213200012130) |
| `disable` | [disable](resources--cloud_elastic_ip--reference--group-001.md#canonical-2321200011033103-1003321010231113-3032113331323220-3022031222311133-0213032021212130-3000132001301013-2120321213010323-3012222011233130) |
| `id` | [ID](resources--cloud_elastic_ip--reference--group-001.md#canonical-1313022001231131-0132330013031122-1332100030323322-3232221321100001-1111111320313322-1123132222110223-3303311232020020-1030331302331300) |
| `item_count` | [item_count](resources--cloud_elastic_ip--reference--group-001.md#canonical-3033303323013233-2332203130012203-0012121332123211-2223202101230001-3023212222320123-1030321010032101-1232232201303103-2310111123303031) |
| `labels` | [labels](resources--cloud_elastic_ip--reference--group-001.md#canonical-2103312231132200-2100033032331201-3011212132130301-0200033130231001-3333322213032000-1333211313002313-2131311002100201-2133320101320212) |
| `name` | [name](resources--cloud_elastic_ip--reference--group-001.md#canonical-0222033330010213-0200331330133231-0022300310033030-2330211022322000-3332112131323313-1032110010331110-3231223230003030-2022003111322131) |
| `namespace` | [namespace](resources--cloud_elastic_ip--reference--group-001.md#canonical-2322123102102113-1203323331230132-3122321302223003-2313222330202131-2012202223203102-2020332213130212-1211000030021112-1010003013223230) |
| `site_ref` | [site_ref](resources--cloud_elastic_ip--reference--group-001.md#canonical-1020033011000202-1330111121203323-2021233213203212-1331212212230202-0020301323232232-2200313100030121-2233031301033101-3001113220122210) |
| `site_ref.kind` | [site_ref.kind](resources--cloud_elastic_ip--reference--group-001.md#canonical-1332103000232232-1220113113300011-2201213112011032-1313021102030333-3222032331103013-2233210023103102-3132023121032213-3231100013313133) |
| `site_ref.name` | [site_ref.name](resources--cloud_elastic_ip--reference--group-001.md#canonical-1203123212021101-2232112301310010-2113010301003121-3321210333201023-3030022320032112-3010312331323233-2230302123032032-2001202020312100) |
| `site_ref.namespace` | [site_ref.namespace](resources--cloud_elastic_ip--reference--group-001.md#canonical-2132003130311101-1032211202223330-3131000010201302-1113202231213110-0232100311032133-1310100200110133-1102322102201000-3213200320010001) |
| `site_ref.tenant` | [site_ref.tenant](resources--cloud_elastic_ip--reference--group-001.md#canonical-3203221303220200-0100132100020032-0221332223121000-2200300111220011-0331231021323222-3222213300221323-2202023330220120-1133132321333302) |
| `site_ref.uid` | [site_ref.uid](resources--cloud_elastic_ip--reference--group-001.md#canonical-2210302101100223-1130221233102211-2323102203311132-1132113121202301-0033231203111011-2223132222132303-0312231332032013-0110131100201000) |
| `timeouts` | [timeouts](resources--cloud_elastic_ip--reference--group-001.md#canonical-2301201212121311-2032220013323030-0021111003332100-0002202333331231-3321321123312121-1031000032133330-1210201012011122-3130210301130110) |
| `timeouts.create` | [timeouts.create](resources--cloud_elastic_ip--reference--group-001.md#canonical-2300100110121120-3113202001113210-1213320000333011-2300202103002012-2130013110201020-0331203323303310-2002000303021032-3131132010231021) |
| `timeouts.delete` | [timeouts.delete](resources--cloud_elastic_ip--reference--group-001.md#canonical-1102013233032320-2033132213003121-0202003012212132-3303032233001203-0222330003102010-2231320012011210-1313001211231133-0232300202231220) |
| `timeouts.read` | [timeouts.read](resources--cloud_elastic_ip--reference--group-001.md#canonical-1211230300011123-2223030332002131-2031230303111101-3121023111113112-2131102130100123-3011331332122121-0302223121202201-1201131213300023) |
| `timeouts.update` | [timeouts.update](resources--cloud_elastic_ip--reference--group-001.md#canonical-2210213112201302-3211132302232311-1213200321112103-1321203300010223-3123312132000120-0133302032021031-0120030103203223-2113100120221100) |

<a id="canonical-2111212233013122-0203213202333103-1310022022310130-0321121200102033-2231120322112132-2223300221130110-2103212313330203-2221122002132212"></a>

## Next pages — Property reference / 112322111102 / 13

- [site_ref](resources--cloud_elastic_ip--reference--group-001.md#canonical-2102121212113212-1112222023302311-1111213230011002-2022232101311132-3321202123233120-2020113123321112-1200200220301033-3332330302121120)
- [timeouts](resources--cloud_elastic_ip--reference--group-001.md#canonical-0333031312331000-3300132311020030-3221133003000211-3031332022113301-0202021002223002-1103112311021221-2310003223220123-0300302133023103)
- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-1012233313021212-0003222033112111-1303213323303313-1020020302213310-1211110033020313-2122312122003031-2320202301023301-3102202323122130)

<a id="canonical-2102121212113212-1112222023302311-1111213230011002-2022232101311132-3321202123233120-2020113123321112-1200200220301033-3332330302121120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023121120131300-2201102010003332-0332003221032103-1212122120021013-1223222322003303-0130333023212233-3321202222223023-3001302100113033"></a>

## site_ref — site_ref / 131301000320 / 2

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-1012233313021212-0003222033112111-1303213323303313-1020020302213310-1211110033020313-2122312122003031-2320202301023301-3102202323122130)
- [Property reference](resources--cloud_elastic_ip--reference--group-001.md#canonical-3000332101332101-3033231112032333-3203132203302322-1221122003220232-1112221212310323-1330110330201000-0101323211230023-3313311003132210)
- site_ref

<a id="canonical-1020033011000202-1330111121203323-2021233213203212-1331212212230202-0020301323232232-2200313100030121-2233031301033101-3001113220122210"></a>

Type: `"object"`. list nested block, Optional.

Site to which this cloud elastic IP object is attached.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
site_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103113111321212-0332313303332032-2012130101231222-0000203301132210-3023023222322031-2320022111320222-3322330022033020-1101000010303002"></a>

## Direct properties — site_ref / 131301000320 / 3

<a id="canonical-1332103000232232-1220113113300011-2201213112011032-1313021102030333-3222032331103013-2233210023103102-3132023121032213-3231100013313133"></a>

<a id="canonical-3303002000211203-1001233010003001-2110333302211113-2122231303130121-2303103213232020-1303211130230021-1011333032301203-2321313133002101"></a>

## kind property — site_ref / 131301000320 / 4

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

<a id="canonical-1203123212021101-2232112301310010-2113010301003121-3321210333201023-3030022320032112-3010312331323233-2230302123032032-2001202020312100"></a>

<a id="canonical-1323222303333100-0122123111030101-0031003213313012-0230033031102210-0033313131012322-1000112030232110-2202320012221033-1233212202213311"></a>

## name property — site_ref / 131301000320 / 5

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

<a id="canonical-2132003130311101-1032211202223330-3131000010201302-1113202231213110-0232100311032133-1310100200110133-1102322102201000-3213200320010001"></a>

<a id="canonical-1003013221221130-2032203011032312-2211311313300223-3101022131113100-1130103330021203-1130210133113010-2222023100002102-2322131213022302"></a>

## namespace property — site_ref / 131301000320 / 6

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

<a id="canonical-3203221303220200-0100132100020032-0221332223121000-2200300111220011-0331231021323222-3222213300221323-2202023330220120-1133132321333302"></a>

<a id="canonical-1311232012102113-0010223200122022-0332002110320002-2201313203100233-2031233201302012-0101203221323122-3311203112020331-0310202130312133"></a>

## tenant property — site_ref / 131301000320 / 7

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

<a id="canonical-2210302101100223-1130221233102211-2323102203311132-1132113121202301-0033231203111011-2223132222132303-0312231332032013-0110131100201000"></a>

<a id="canonical-0010331330323031-1111223221113010-0301122221100101-3220200120123113-1211132032310130-3202003233133213-3032131113332310-2100210210331133"></a>

## uid property — site_ref / 131301000320 / 8

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

<a id="canonical-0210220322103320-3112311123312031-1113130002032330-1322113330333001-3223323133222233-2321111323213010-3311112010322022-2213131311200320"></a>

## Next pages — site_ref / 131301000320 / 9

- [Property reference](resources--cloud_elastic_ip--reference--group-001.md#canonical-3000332101332101-3033231112032333-3203132203302322-1221122003220232-1112221212310323-1330110330201000-0101323211230023-3313311003132210)
- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-1012233313021212-0003222033112111-1303213323303313-1020020302213310-1211110033020313-2122312122003031-2320202301023301-3102202323122130)

<a id="canonical-0333031312331000-3300132311020030-3221133003000211-3031332022113301-0202021002223002-1103112311021221-2310003223220123-0300302133023103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130231211131120-2321301233331210-2020111200001132-3301000112230030-0113311312310001-1130120123211320-1033222300331101-0102012310233112"></a>

## timeouts — timeouts / 110011312031 / 2

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-1012233313021212-0003222033112111-1303213323303313-1020020302213310-1211110033020313-2122312122003031-2320202301023301-3102202323122130)
- [Property reference](resources--cloud_elastic_ip--reference--group-001.md#canonical-3000332101332101-3033231112032333-3203132203302322-1221122003220232-1112221212310323-1330110330201000-0101323211230023-3313311003132210)
- timeouts

<a id="canonical-2301201212121311-2032220013323030-0021111003332100-0002202333331231-3321321123312121-1031000032133330-1210201012011122-3130210301130110"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1033000330133212-0022000030132012-1300121202223033-3323132232100232-2113302002211203-3331303002211201-1020122321133133-0221000223101023"></a>

## Direct properties — timeouts / 110011312031 / 3

<a id="canonical-2300100110121120-3113202001113210-1213320000333011-2300202103002012-2130013110201020-0331203323303310-2002000303021032-3131132010231021"></a>

<a id="canonical-2012230030032221-3230031003223030-1000330123002121-1203102220303230-0200132220002003-1213032231030331-2122332330221110-0023201021101332"></a>

## create property — timeouts / 110011312031 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1102013233032320-2033132213003121-0202003012212132-3303032233001203-0222330003102010-2231320012011210-1313001211231133-0232300202231220"></a>

<a id="canonical-3001202111013021-3212210321102320-3102232201313122-1110320022211302-3031333131000300-1121222130133331-2001301103103033-3330010132313022"></a>

## delete property — timeouts / 110011312031 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1211230300011123-2223030332002131-2031230303111101-3121023111113112-2131102130100123-3011331332122121-0302223121202201-1201131213300023"></a>

<a id="canonical-2130130203131111-3120230212032101-2203012323230201-1000220201221020-2002302231323000-3013203121002112-1132233110211113-1132022001113130"></a>

## read property — timeouts / 110011312031 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2210213112201302-3211132302232311-1213200321112103-1321203300010223-3123312132000120-0133302032021031-0120030103203223-2113100120221100"></a>

<a id="canonical-1003113000333123-2220120133331211-3202102210021230-1223210211132122-2220132230312332-1010131123231332-0323221231123212-2221012122001202"></a>

## update property — timeouts / 110011312031 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2313002333231201-1232013103133102-1020121033122322-2022223020031312-0122322111010212-0113211323113012-0212012322103032-1312320320233131"></a>

## Next pages — timeouts / 110011312031 / 8

- [Property reference](resources--cloud_elastic_ip--reference--group-001.md#canonical-3000332101332101-3033231112032333-3203132203302322-1221122003220232-1112221212310323-1330110330201000-0101323211230023-3313311003132210)
- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md#canonical-1012233313021212-0003222033112111-1303213323303313-1020020302213310-1211110033020313-2122312122003031-2320202301023301-3102202323122130)
