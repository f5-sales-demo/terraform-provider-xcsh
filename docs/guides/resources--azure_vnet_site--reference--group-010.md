---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-0133223003221300-0313023113323130-3102030321120203-2100103231331211-0202000110322133-1310223110201033-2220331332330322-1112112203032020"></a>

## Direct properties — global_network_list / 231303301210 / 3

- [global_network_connections](resources--azure_vnet_site--reference--group-010.md#canonical-3331000333223303-3111011120323011-1321313302331202-3303320321023332-2201100123013303-0011031313121111-1332023122323130-3123301313002312): complete subsection reference.

<a id="canonical-3001002103032303-2233131221030203-3202311132221023-2021112303001123-1013100233123131-0101123022021121-2010200001011213-0101301121113310"></a>

## Next pages — global_network_list / 231303301210 / 4

- [voltstack_cluster_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-010.md#canonical-3331000333223303-3111011120323011-1321313302331202-3303320321023332-2201100123013303-0011031313121111-1332023122323130-3123301313002312)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3331000333223303-3111011120323011-1321313302331202-3303320321023332-2201100123013303-0011031313121111-1332023122323130-3123301313002312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021201000331013-0113110121303213-0131123003300100-1100231133332333-1130132023121003-3122201010111112-3331011221220323-3322103201202301"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections — global_network_connections / 330332330332 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-0302133020210101-2221032211110211-1032111131302013-3223120131021120-3213310203101213-1110303313200321-2030023121010000-1210111333331312)
- voltstack_cluster_ar.global_network_list.global_network_connections

<a id="canonical-1003231031220301-1210131223023230-2321022303320003-2001320001202311-0321022332131331-0321013120032102-3223013211300020-1101323313322020"></a>

Type: `"object"`. list nested block, Optional.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("sli_to_global_dr",
    "slo_to_global_dr")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
global_network_connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133123030332022-3111321022001220-2310033032132211-3112310312012213-3002100331121210-3022023012232102-3303321331122302-3333113222232012"></a>

## Direct properties — global_network_connections / 330332330332 / 3

- [sli_to_global_dr](resources--azure_vnet_site--reference--group-010.md#canonical-1010222223030313-2003012033230120-3110333333001030-1021023123122131-1212001310320322-2331030201233123-0221320013020223-1021333010000321): complete subsection reference.

- [slo_to_global_dr](resources--azure_vnet_site--reference--group-010.md#canonical-2130022003231131-2220112201210232-3303113132120110-0013022132310131-0000330100023211-0333110310003311-1031302221011330-2322220330022330): complete subsection reference.

<a id="canonical-2003211200003322-3032110023002132-0232211110111310-1113320222001130-1121333333232031-2022022312021101-0010313211000123-0031211102220011"></a>

## Next pages — global_network_connections / 330332330332 / 4

- [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr](resources--azure_vnet_site--reference--group-010.md#canonical-1010222223030313-2003012033230120-3110333333001030-1021023123122131-1212001310320322-2331030201233123-0221320013020223-1021333010000321)
- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr](resources--azure_vnet_site--reference--group-010.md#canonical-2130022003231131-2220112201210232-3303113132120110-0013022132310131-0000330100023211-0333110310003311-1031302221011330-2322220330022330)
- [voltstack_cluster_ar.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-0302133020210101-2221032211110211-1032111131302013-3223120131021120-3213310203101213-1110303313200321-2030023121010000-1210111333331312)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1010222223030313-2003012033230120-3110333333001030-1021023123122131-1212001310320322-2331030201233123-0221320013020223-1021333010000321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023000101200103-0311002133122323-2102331102000131-0012010233113223-2002033112332111-1301120302330200-0303313102333222-1000301213233113"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr — sli_to_global_dr / 311202110021 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-0302133020210101-2221032211110211-1032111131302013-3223120131021120-3213310203101213-1110303313200321-2030023121010000-1210111333331312)
- [voltstack_cluster_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-010.md#canonical-3331000333223303-3111011120323011-1321313302331202-3303320321023332-2201100123013303-0011031313121111-1332023122323130-3123301313002312)
- voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-3210021323333303-3312312211121323-3221230221032102-1213101311300112-3003332320021012-2022103200002103-1132111000032230-0300103130023012"></a>

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
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232213203202113-3122313131301112-2123332000103132-1310020210032310-2300311100003033-3322030101131111-2021312000121230-1211321313010131"></a>

## Direct properties — sli_to_global_dr / 311202110021 / 3

- [global_vn](resources--azure_vnet_site--reference--group-010.md#canonical-0200111220020201-1022203231230030-2031030001031311-0231212321300102-3021031230120132-2320311302032130-1133120302201210-2020313322210102): complete subsection reference.

<a id="canonical-0232101102220303-2022333102013310-0023301121021102-1002010023131230-1102323033032203-0111022002200331-3312222131213123-1221030213123222"></a>

## Next pages — sli_to_global_dr / 311202110021 / 4

- [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--azure_vnet_site--reference--group-010.md#canonical-0200111220020201-1022203231230030-2031030001031311-0231212321300102-3021031230120132-2320311302032130-1133120302201210-2020313322210102)
- [voltstack_cluster_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-010.md#canonical-3331000333223303-3111011120323011-1321313302331202-3303320321023332-2201100123013303-0011031313121111-1332023122323130-3123301313002312)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0200111220020201-1022203231230030-2031030001031311-0231212321300102-3021031230120132-2320311302032130-1133120302201210-2020313322210102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033321231223030-2321232220111320-1122110312200010-0233122033122321-0113000332301021-3221121031231313-1332213323112213-2303222300301123"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn — global_vn / 132113133032 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-0302133020210101-2221032211110211-1032111131302013-3223120131021120-3213310203101213-1110303313200321-2030023121010000-1210111333331312)
- [voltstack_cluster_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-010.md#canonical-3331000333223303-3111011120323011-1321313302331202-3303320321023332-2201100123013303-0011031313121111-1332023122323130-3123301313002312)
- [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr](resources--azure_vnet_site--reference--group-010.md#canonical-1010222223030313-2003012033230120-3110333333001030-1021023123122131-1212001310320322-2331030201233123-0221320013020223-1021333010000321)
- voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-0030330033013113-2103101102103210-2212001320313032-2301330032022103-0222121210313022-1310030103000211-0100230002112032-0000121331112202"></a>

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

<a id="canonical-1220121230323001-3031303222220133-3211320012213310-0113201131201031-3221003033101230-1021312002200210-3023223012321121-0032312222202021"></a>

## Direct properties — global_vn / 132113133032 / 3

<a id="canonical-3032020333002213-3012023230131210-0321122102302003-2021013222132322-2202130313131001-2312223330112021-3103233221113013-2200022301212131"></a>

<a id="canonical-2322123030023330-1200221101122330-0333001203203000-1033332032223313-3130123112131230-3011100001013233-3133131321130120-0001211100210220"></a>

## name property — global_vn / 132113133032 / 4

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

<a id="canonical-2312113112100202-0231222122220021-1102002221113220-0220132231131010-2021002001321301-1113320000111330-2301330023012101-3321332213203210"></a>

<a id="canonical-3331000012301212-0103211211303020-1020102222002002-1100303000100030-0322303313320120-3021013213320220-0322331330232103-1210321130300333"></a>

## namespace property — global_vn / 132113133032 / 5

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

<a id="canonical-3132110031312221-1130210222212102-2031112003101223-1322220210012211-2123113110221220-3010213312120303-3120233312310031-2012003131331001"></a>

<a id="canonical-2230032320233121-2010203212032222-3003233103313132-0210320203101103-2102221312012131-1230323000310300-1231001013032301-2233023232333110"></a>

## tenant property — global_vn / 132113133032 / 6

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

<a id="canonical-1312331301112133-1031000200123030-2021102011331112-0022223110333032-2000131022323103-1330032213131322-3133321010030200-1310210112232213"></a>

## Next pages — global_vn / 132113133032 / 7

- [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr](resources--azure_vnet_site--reference--group-010.md#canonical-1010222223030313-2003012033230120-3110333333001030-1021023123122131-1212001310320322-2331030201233123-0221320013020223-1021333010000321)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2130022003231131-2220112201210232-3303113132120110-0013022132310131-0000330100023211-0333110310003311-1031302221011330-2322220330022330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330333312313021-1031003201230302-3110230222103121-3200121222010021-1333132102203121-1001331103103130-0212201132022210-0101003202201020"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr — slo_to_global_dr / 211000013121 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-0302133020210101-2221032211110211-1032111131302013-3223120131021120-3213310203101213-1110303313200321-2030023121010000-1210111333331312)
- [voltstack_cluster_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-010.md#canonical-3331000333223303-3111011120323011-1321313302331202-3303320321023332-2201100123013303-0011031313121111-1332023122323130-3123301313002312)
- voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-3103231222031303-0321021211122102-1323212120102202-3203301303330031-3122020132113231-0032001222333130-1323110003233103-2113320110332100"></a>

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

<a id="canonical-1103032010031033-3031303132110021-3233301310313323-0013333021322300-3300220122132000-2221000131312123-2131030223301301-2300321122021000"></a>

## Direct properties — slo_to_global_dr / 211000013121 / 3

- [global_vn](resources--azure_vnet_site--reference--group-010.md#canonical-0031230023300323-0122211030012203-2002003222301311-3230013231112032-0003003000022320-2022131013323130-2200000033110203-3332022130201202): complete subsection reference.

<a id="canonical-3030123201303331-2031323123123213-2012023033100031-1013321120310320-2122100220100033-3103011230320331-3021112211120311-0121003202122311"></a>

## Next pages — slo_to_global_dr / 211000013121 / 4

- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--azure_vnet_site--reference--group-010.md#canonical-0031230023300323-0122211030012203-2002003222301311-3230013231112032-0003003000022320-2022131013323130-2200000033110203-3332022130201202)
- [voltstack_cluster_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-010.md#canonical-3331000333223303-3111011120323011-1321313302331202-3303320321023332-2201100123013303-0011031313121111-1332023122323130-3123301313002312)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0031230023300323-0122211030012203-2002003222301311-3230013231112032-0003003000022320-2022131013323130-2200000033110203-3332022130201202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323312131030032-0121122020123112-2020231130003233-0012213210131323-1231111133032320-2233332102331200-3100111303000123-1000232330013021"></a>

## voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn — global_vn / 011301203203 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.global_network_list](resources--azure_vnet_site--reference--group-009.md#canonical-0302133020210101-2221032211110211-1032111131302013-3223120131021120-3213310203101213-1110303313200321-2030023121010000-1210111333331312)
- [voltstack_cluster_ar.global_network_list.global_network_connections](resources--azure_vnet_site--reference--group-010.md#canonical-3331000333223303-3111011120323011-1321313302331202-3303320321023332-2201100123013303-0011031313121111-1332023122323130-3123301313002312)
- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr](resources--azure_vnet_site--reference--group-010.md#canonical-2130022003231131-2220112201210232-3303113132120110-0013022132310131-0000330100023211-0333110310003311-1031302221011330-2322220330022330)
- voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-3033321231330100-0221311223321222-3332030012230310-3023132120212001-3011332110123022-3030322033211130-0311211130203101-0000302033201232"></a>

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

<a id="canonical-2223221321010210-1213202200233022-1221302101130101-3330023132233132-2230200322201033-1013200131232203-0121323013203333-2133312022320213"></a>

## Direct properties — global_vn / 011301203203 / 3

<a id="canonical-0322211333020123-3203220213233230-2313312103021130-2011013122112200-2313120231231330-3322110201112223-3031313230032230-3202233320200123"></a>

<a id="canonical-2222221231131300-0012203230122002-3330200200311331-0131212322120233-3103033010201233-3330021330001130-0330220020312310-0330123031031200"></a>

## name property — global_vn / 011301203203 / 4

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

<a id="canonical-1331211200022302-3203320230021010-1230232211201002-1330223302311111-2011312112300213-0300133222233113-2021212211200232-1012320203031102"></a>

<a id="canonical-3310201330231103-2030031012000021-2123312300032011-0312112232133130-2000320302331323-1112001033311101-1122031111202311-2110010331030200"></a>

## namespace property — global_vn / 011301203203 / 5

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

<a id="canonical-3021303033102010-3322030203020232-2133120111121011-1213330003320302-0230100030300012-3012120321310223-1303201221102020-1100231121212233"></a>

<a id="canonical-2113322013301231-1221321001322212-3030002120332022-2011110031113232-1302020230322010-1231003122322300-3300101231313201-3321103003022221"></a>

## tenant property — global_vn / 011301203203 / 6

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

<a id="canonical-1012131310023210-2223030323321111-2103203032311301-0201131130311102-1021213331323123-1103332132220313-0003202033031112-0020112222121110"></a>

## Next pages — global_vn / 011301203203 / 7

- [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr](resources--azure_vnet_site--reference--group-010.md#canonical-2130022003231131-2220112201210232-3303113132120110-0013022132310131-0000330100023211-0333110310003311-1031302221011330-2322220330022330)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0110300210201321-2233321233202030-0203210323012120-2311222321101002-0110210121020330-0320321103302133-0021332210302312-1100012010211331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112333311210012-1300330232310011-0130331003103010-2213312121302200-2220031103332223-3320113110102131-2203221231312201-2302131322100022"></a>

## voltstack_cluster_ar.k8s_cluster — k8s_cluster / 300201232303 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.k8s_cluster

<a id="canonical-3031031302201212-3323331213000103-3013020333030111-1031003122320231-3303213223020021-0330301223322221-0331030012323121-2000211222023102"></a>

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
k8s_cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201033011213021-2131003103110101-2030112002200102-1111301320113231-3003323033320221-2130022012312303-0211200033220112-3000312220232132"></a>

## Direct properties — k8s_cluster / 300201232303 / 3

<a id="canonical-3011103220000323-1211032300332233-0220222320302130-0110020202032110-1220221031102300-3301023201211221-1110321303301001-1223212300123132"></a>

<a id="canonical-3312101131203211-2002113021221221-0222032300003221-2022330120012322-0010120133023102-3311011232033011-1322202120321131-3211223101323311"></a>

## name property — k8s_cluster / 300201232303 / 4

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

<a id="canonical-2321312002333221-1010230200323133-2320211112122332-2232200302012103-2000113103123133-3232212300001300-3312032123333212-1132313323310000"></a>

<a id="canonical-0321122020320102-0332100222221203-3001010100333032-3100012322000213-3102022200100113-3332020033300201-1101300233312013-0003212011023211"></a>

## namespace property — k8s_cluster / 300201232303 / 5

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

<a id="canonical-0220332003333332-2022222033310120-3032300331220101-2300112032112001-0230331123220312-2312012310313322-0232231231322202-1311313301212222"></a>

<a id="canonical-2311320213130203-0301033222112233-0211223312220220-2132311113302130-3330002203121222-2202333212333000-1132302300312220-1120301012033000"></a>

## tenant property — k8s_cluster / 300201232303 / 6

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

<a id="canonical-2333323031110313-0321102013203003-0132331302102020-0010021221100021-3233101323131031-0332332313202211-0010312221201201-2203210013020020"></a>

## Next pages — k8s_cluster / 300201232303 / 7

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1312111013220222-1211300132311131-3210122100203322-3220302230110020-0303331211110022-3331201313233122-1030212312123333-2033112102210032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031211231303100-1312101020321100-3331022203303112-3033102311321100-0203032302331302-3230232233332203-0001123113001022-1002031321012212"></a>

## voltstack_cluster_ar.no_dc_cluster_group — no_dc_cluster_group / 100321012323 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.no_dc_cluster_group

<a id="canonical-0031220013332021-3103313110022200-3031101011121011-1312132300322203-0110002013110103-2220223303101102-2220331221111130-2100323120230131"></a>

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

<a id="canonical-0023310033220301-1110033223100230-3331300003321331-0220030202132020-0130202030202212-2102210313103113-2100211022011000-2111133230131321"></a>

## Direct properties — no_dc_cluster_group / 100321012323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022123223311232-3331131332101311-3201032123122230-1021103000213230-2321311111203201-2020002012110001-1113112013200202-3130110303012130"></a>

## Next pages — no_dc_cluster_group / 100321012323 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3003200021130023-1233323111132230-1011120203303301-0230303301011312-0211003011133211-1000331102100331-3221202211111312-1222010233330313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033333001032112-1311303203112322-3311121100032003-1123102322131003-2122100223312012-0212312310232020-3322021210022113-1103231112111332"></a>

## voltstack_cluster_ar.no_forward_proxy — no_forward_proxy / 033100213200 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.no_forward_proxy

<a id="canonical-1222310233033022-1122012212020201-3101033110001212-0002103130211021-3220103220032011-3320023132110213-3220331030223333-1033311222003122"></a>

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

<a id="canonical-1210032002003132-3101221112033012-1323223003123213-1011112330111132-0331023211100121-3031203111233113-1130211022333023-1102120123131212"></a>

## Direct properties — no_forward_proxy / 033100213200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322122120211333-3111020023022100-1303112100333032-3232212303200323-0021311201201123-1032111201013013-3132023301022102-1020130222012020"></a>

## Next pages — no_forward_proxy / 033100213200 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3303132032223003-0010031311031112-3312103133303103-0222233230331001-3103303201211301-3313021100103310-0030313021132011-0120112203023332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013100223123312-2212113230312312-0020000210313100-1320300310022222-0222201103030002-3031323101010232-1001220013002032-0121103300302032"></a>

## voltstack_cluster_ar.no_global_network — no_global_network / 012132112110 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.no_global_network

<a id="canonical-2121003232211123-2202000132301003-2331102220103000-2023213222122122-2031131110033022-3130111202222323-1011311000210021-1331310121100331"></a>

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

<a id="canonical-1322323003033313-2303120000011103-3101302321113222-0100301013111033-0132031211222100-1033232210033201-1030021210300303-3223220103002233"></a>

## Direct properties — no_global_network / 012132112110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333010231211121-0133212333302021-1321012300220310-0131003232003020-3301110331103312-3212300130022110-0212132110133302-1001201212111313"></a>

## Next pages — no_global_network / 012132112110 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3013312222012123-1001012001300003-1303113001030332-2212223031300022-0012022323110023-1102301123102303-0101303233100120-1213220011032303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220101320330222-2300313111200323-3320302011200232-2302022003110011-3313033030223003-3331000020111010-0132221132313001-0323123031000201"></a>

## voltstack_cluster_ar.no_k8s_cluster — no_k8s_cluster / 300102112313 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.no_k8s_cluster

<a id="canonical-0212323132213332-1113130301103213-3313022113131223-1202032213333121-2030003001333121-3121200233111113-0331021312020000-0230132210201300"></a>

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
no_k8s_cluster = {}
```

<a id="canonical-0211223203222303-1301203312331323-1233333110232233-1033220103301311-1232013303031301-1311223220033131-0023202222330012-1200130211211323"></a>

## Direct properties — no_k8s_cluster / 300102112313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000313010223301-1212210333331100-2303021121201223-2112123202121031-0231213233301321-3010033011223130-1222102101110302-1013332022130131"></a>

## Next pages — no_k8s_cluster / 300102112313 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3210030201312131-2301022101213032-3120333110130331-3221302301101010-1233123102210000-3222133213211020-0222302332321320-0330301231133113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202123110001313-0111231030221001-2122332121322023-1131300000333021-3002322013100122-2230113111232300-3303021312233203-3313213123011130"></a>

## voltstack_cluster_ar.no_network_policy — no_network_policy / 313100110302 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.no_network_policy

<a id="canonical-0322323012021302-2223222201200213-3312021230030022-1303210101233210-1221333030331333-0232323022010103-1210313202123010-1033032002200012"></a>

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

<a id="canonical-0012003133312003-1111121112121132-0321233233211311-2331122131321011-3130110000101112-0133101200212302-3010200030323010-3231311232123131"></a>

## Direct properties — no_network_policy / 313100110302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223330132021220-1133013221300211-2122320312122220-2021230132232110-1132102103211301-1300230230322113-3323021031112110-1132223220123030"></a>

## Next pages — no_network_policy / 313100110302 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2113103013030123-3333113313303300-3313310320000031-1220022002113312-0332300321010030-2330010210211013-3310103010231303-3022211021103012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231200101311313-3330120110100110-2131002001210220-1221100113132232-3132230122031310-3103003233330322-2223303023123110-3033120100121203"></a>

## voltstack_cluster_ar.no_outside_static_routes — no_outside_static_routes / 203030102031 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.no_outside_static_routes

<a id="canonical-0223203301211023-1310101130113222-2120232103003031-0030211111001122-0100323221300310-3103201332013323-3220010132200331-1220130130313231"></a>

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

<a id="canonical-1020211332200131-0000132303310201-2013233202300300-1230003312310032-1002301302123222-2120210122312001-0321000303302300-0102222010313113"></a>

## Direct properties — no_outside_static_routes / 203030102031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112012020120030-1111302132103201-1013010130302201-0132000131311301-1001323033100122-1201221033320101-0020103221031303-0132330200001021"></a>

## Next pages — no_outside_static_routes / 203030102031 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3323133232110233-1133112012201220-0312330113021030-1133010221211103-3133033331010223-3031023302233023-3010112130013103-0032221030002012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112303011210202-2212312012213211-3231220330313303-2000120312223120-2303333200023122-2322111332223102-1100131202102100-1033132132232011"></a>

## voltstack_cluster_ar.node — node / 132100120013 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.node

<a id="canonical-1112203012030321-0020331313331033-0203323121013003-1211221323202320-3122330013330100-0012213323312110-1323023232122131-3311320131121020"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating Single interface Node for Alternate Region.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("fault_domain",
    "node_number",
    "update_domain")}
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
node {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302221303031001-2000133212213232-2013111032013112-3032011210311321-0333020331100313-2201101001201032-2102202011001331-3111330303211221"></a>

## Direct properties — node / 132100120013 / 3

<a id="canonical-0221301223222300-1120230223223003-3200222231111032-2232112220323103-1030302220132121-3220302333100302-3300303121313211-3100120200033101"></a>

<a id="canonical-1030100233000010-2010022112122101-1233002203201101-2102110300020101-2103113011323313-3311121123132333-1312022230323333-1222202222012120"></a>

## fault_domain property — node / 132100120013 / 4

Type: `"number"`. Optional.

Namuber of fault domains to be used while creating the availability set.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 3),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "3"
  }
}
```

- [local_subnet](resources--azure_vnet_site--reference--group-010.md#canonical-2332331131013230-1010013320210113-0313233211220101-0332230021233313-3233123123332223-1123113020101310-0133011123300030-3021120202132223): complete subsection reference.

<a id="canonical-0331110000121221-1212320321012303-0010003122112212-3132012003213302-0132303110230310-1111103311131330-3232200133323130-2103302032332033"></a>

<a id="canonical-2002321131233322-0223003032001303-1000313022003021-2003330233022232-2312122211332330-1120213120320110-1303320321302221-2320231213113322"></a>

## node_number property — node / 132100120013 / 5

Type: `"number"`. Optional.

Number of main nodes to create, either 1 or 3.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

<a id="canonical-0122331212031111-2022021132103111-1221212100003111-0011133303003303-1130132022333132-1133020131030121-3211020001110303-3203012022201213"></a>

<a id="canonical-1201202120022101-1111301111012020-3211321110221303-1101200023030021-3022133310302212-1302130133120103-1320010012101033-1123100233123112"></a>

## update_domain property — node / 132100120013 / 6

Type: `"number"`. Optional.

Namuber of update domains to be used while creating the availability set.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-3112123002310013-3321102103323312-0031000300232113-2020322110320202-2231222012002333-2301023333103131-1201002332313113-1212122211211000"></a>

## Next pages — node / 132100120013 / 7

- [voltstack_cluster_ar.node.local_subnet](resources--azure_vnet_site--reference--group-010.md#canonical-2332331131013230-1010013320210113-0313233211220101-0332230021233313-3233123123332223-1123113020101310-0133011123300030-3021120202132223)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2332331131013230-1010013320210113-0313233211220101-0332230021233313-3233123123332223-1123113020101310-0133011123300030-3021120202132223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331332310100030-2033033310020102-1230021001103112-3032301020201211-1113332333130121-1010103220010321-2113203101030202-0211212133321113"></a>

## voltstack_cluster_ar.node.local_subnet — local_subnet / 332101230032 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.node](resources--azure_vnet_site--reference--group-010.md#canonical-3323133232110233-1133112012201220-0312330113021030-1133010221211103-3133033331010223-3031023302233023-3010112130013103-0032221030002012)
- voltstack_cluster_ar.node.local_subnet

<a id="canonical-3321331211133322-0032310022012110-1101333301101123-1320303121001231-3130221212332200-0031100011233023-3302032131210021-0311313020220211"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for local subnet.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("subnet",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"subnet\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
local_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021130021012112-3020220021302210-0300211000102312-3221133302210301-1133103131133011-0213120001032333-3213121132333000-2133202022311211"></a>

## Direct properties — local_subnet / 332101230032 / 3

- [subnet](resources--azure_vnet_site--reference--group-010.md#canonical-2232031212110232-0202022123030003-3101022300133111-1102322311320033-1300031210010001-0302232230230220-3102231013213222-1010112020131231): complete subsection reference.

- [subnet_param](resources--azure_vnet_site--reference--group-010.md#canonical-0131013322020212-1112332112111022-2033003123333230-3123003023021223-3332130313033111-1122130211131320-2031221133102022-1333333321313302): complete subsection reference.

<a id="canonical-3313020010130011-1101000031103013-2221130203301302-2112221313101203-1230310033012110-3232222212010211-1313032320012310-1113210200101210"></a>

## Next pages — local_subnet / 332101230032 / 4

- [voltstack_cluster_ar.node.local_subnet.subnet](resources--azure_vnet_site--reference--group-010.md#canonical-2232031212110232-0202022123030003-3101022300133111-1102322311320033-1300031210010001-0302232230230220-3102231013213222-1010112020131231)
- [voltstack_cluster_ar.node.local_subnet.subnet_param](resources--azure_vnet_site--reference--group-010.md#canonical-0131013322020212-1112332112111022-2033003123333230-3123003023021223-3332130313033111-1122130211131320-2031221133102022-1333333321313302)
- [voltstack_cluster_ar.node](resources--azure_vnet_site--reference--group-010.md#canonical-3323133232110233-1133112012201220-0312330113021030-1133010221211103-3133033331010223-3031023302233023-3010112130013103-0032221030002012)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2232031212110232-0202022123030003-3101022300133111-1102322311320033-1300031210010001-0302232230230220-3102231013213222-1010112020131231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011312110101332-0000020030111000-2232201020033133-3232200011222223-0133123131033011-3200103331330321-3330100131002200-2100223133232101"></a>

## voltstack_cluster_ar.node.local_subnet.subnet — subnet / 132223231013 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.node](resources--azure_vnet_site--reference--group-010.md#canonical-3323133232110233-1133112012201220-0312330113021030-1133010221211103-3133033331010223-3031023302233023-3010112130013103-0032221030002012)
- [voltstack_cluster_ar.node.local_subnet](resources--azure_vnet_site--reference--group-010.md#canonical-2332331131013230-1010013320210113-0313233211220101-0332230021233313-3233123123332223-1123113020101310-0133011123300030-3021120202132223)
- voltstack_cluster_ar.node.local_subnet.subnet

<a id="canonical-1003230210032020-0320213322032023-3200003003211013-1022020102031202-2001302133313012-0321232022222100-3000001322220323-3131311311223200"></a>

Type: `"object"`. single nested block, Optional.

Subnet specification for network segmentation.

Upstream description:

Parameters for Azure subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("subnet_name"),
  validators.ConflictingObjectAttributes("subnet_resource_grp",
    "vnet_resource_group")}
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
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

Terraform syntax:

```terraform
subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230310301133133-2313211312331332-2233302130310310-2220021303120101-3022031011101333-2121201020130213-2111030200131322-1112322330213013"></a>

## Direct properties — subnet / 132223231013 / 3

<a id="canonical-3323130211311021-0011111121232230-0230110023002223-2130103332330221-3001313200031202-0113022103023230-2003222203013220-2002201321001223"></a>

<a id="canonical-2201303322033110-3111320223302102-1213231133322000-2113212133333201-2213323332232032-2030220102112332-3213320020102322-3013010011211313"></a>

## subnet_name property — subnet / 132223231013 / 4

Type: `"string"`. Optional.

Subnet Name. Name of existing subnet.

Upstream description:

Name of existing subnet.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-3023103010021303-2132012201013322-0203231323210231-2332203213302020-3121133203130222-0023222312033100-0312210323312023-1312301322010220"></a>

<a id="canonical-3232122302121310-2013101231322202-1203233331220321-2122310003112331-1232213200133032-0231310010213320-0311301311121230-0212113230101323"></a>

## subnet_resource_grp property — subnet / 132223231013 / 5

Type: `"string"`. Optional.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [vnet_resource_group](resources--azure_vnet_site--reference--group-010.md#canonical-1033112103320033-2030001301232332-2022303132233323-1323103201212100-0310301322010131-2202321231321100-2213232331212313-3210023102221132): complete subsection reference.

<a id="canonical-3332132213303121-1131333102032112-2202102131321310-2331223202013232-3012220211330212-1001020202330202-3003113202021110-2230232020020233"></a>

## Next pages — subnet / 132223231013 / 6

- [voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group](resources--azure_vnet_site--reference--group-010.md#canonical-1033112103320033-2030001301232332-2022303132233323-1323103201212100-0310301322010131-2202321231321100-2213232331212313-3210023102221132)
- [voltstack_cluster_ar.node.local_subnet](resources--azure_vnet_site--reference--group-010.md#canonical-2332331131013230-1010013320210113-0313233211220101-0332230021233313-3233123123332223-1123113020101310-0133011123300030-3021120202132223)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1033112103320033-2030001301232332-2022303132233323-1323103201212100-0310301322010131-2202321231321100-2213232331212313-3210023102221132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110130130321233-1221202323001101-0011011203223222-0333111330010231-2201231310321120-1100301201310003-0231203012000103-2221101331301123"></a>

## voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group — vnet_resource_group / 311100012131 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.node](resources--azure_vnet_site--reference--group-010.md#canonical-3323133232110233-1133112012201220-0312330113021030-1133010221211103-3133033331010223-3031023302233023-3010112130013103-0032221030002012)
- [voltstack_cluster_ar.node.local_subnet](resources--azure_vnet_site--reference--group-010.md#canonical-2332331131013230-1010013320210113-0313233211220101-0332230021233313-3233123123332223-1123113020101310-0133011123300030-3021120202132223)
- [voltstack_cluster_ar.node.local_subnet.subnet](resources--azure_vnet_site--reference--group-010.md#canonical-2232031212110232-0202022123030003-3101022300133111-1102322311320033-1300031210010001-0302232230230220-3102231013213222-1010112020131231)
- voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group

<a id="canonical-1331210221131333-1230303000112311-3231233311233000-1111222122100032-1203321211232102-2023232312110013-0103011111033021-0001213130012331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vnet resource group.

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
vnet_resource_group = {}
```

<a id="canonical-1002323223032100-2230030210202010-3130300121003303-1102012311203222-2133200003310133-0333120012022101-0312001110333220-0122100211331233"></a>

## Direct properties — vnet_resource_group / 311100012131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311013030100312-2300310132013120-3212301000301120-1020020101212122-3322111222230031-2132201321031312-0122013033211320-1323312132332333"></a>

## Next pages — vnet_resource_group / 311100012131 / 4

- [voltstack_cluster_ar.node.local_subnet.subnet](resources--azure_vnet_site--reference--group-010.md#canonical-2232031212110232-0202022123030003-3101022300133111-1102322311320033-1300031210010001-0302232230230220-3102231013213222-1010112020131231)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0131013322020212-1112332112111022-2033003123333230-3123003023021223-3332130313033111-1122130211131320-2031221133102022-1333333321313302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333130032121311-0333113221321330-2002320020321333-3122321303032003-1100123013303313-1221213110323133-1221000123102311-3301332122302310"></a>

## voltstack_cluster_ar.node.local_subnet.subnet_param — subnet_param / 221132113100 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.node](resources--azure_vnet_site--reference--group-010.md#canonical-3323133232110233-1133112012201220-0312330113021030-1133010221211103-3133033331010223-3031023302233023-3010112130013103-0032221030002012)
- [voltstack_cluster_ar.node.local_subnet](resources--azure_vnet_site--reference--group-010.md#canonical-2332331131013230-1010013320210113-0313233211220101-0332230021233313-3233123123332223-1123113020101310-0133011123300030-3021120202132223)
- voltstack_cluster_ar.node.local_subnet.subnet_param

<a id="canonical-1132200103123120-0133300321220030-3120202100133101-1312002202331032-2131111033320113-1302210023001020-1033223333302321-0123313311112302"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031002032000223-3330002331333102-0310221202111202-1123300030201300-1001130320332123-3110022112123213-3212022010030133-3110122203031330"></a>

## Direct properties — subnet_param / 221132113100 / 3

<a id="canonical-1121002111000033-3020201023031010-0230300010332221-0231301012223210-0110022120001121-1202000303132320-0102103013203102-3003221320203023"></a>

<a id="canonical-2031223001103011-2012300131200132-0332022023320010-1122233011020111-3011303113120001-0200101210323000-3123011112002122-2001310230032231"></a>

## IPv4 property — subnet_param / 221132113100 / 4

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-1320331030020011-1132222222022012-2022213332101322-2212023010330210-1003123023131200-2202213102200022-1320030030121220-0013131010000202"></a>

## Next pages — subnet_param / 221132113100 / 5

- [voltstack_cluster_ar.node.local_subnet](resources--azure_vnet_site--reference--group-010.md#canonical-2332331131013230-1010013320210113-0313233211220101-0332230021233313-3233123123332223-1123113020101310-0133011123300030-3021120202132223)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030133033302201-3102312002033201-3120031022003011-1133121223232120-3032033212001030-0033300333103213-1202202223110023-2333231300323103"></a>

## voltstack_cluster_ar.outside_static_routes — outside_static_routes / 221103313311 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.outside_static_routes

<a id="canonical-1013302131010000-2212023120233013-0223123132101221-1302212002233020-0220112002322213-3302211032322310-1233303123120222-0030130202323212"></a>

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

<a id="canonical-0013000022300332-1110330222003112-2322200121231230-0230320300213013-1033011331013031-0302122202123220-0211300001021030-2201113131023113"></a>

## Direct properties — outside_static_routes / 221103313311 / 3

- [static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132): complete subsection reference.

<a id="canonical-2110111033222302-1322103232002201-1321203100313213-3110112211300002-3113213130133202-1130113313222203-0312302000232032-2001212002231112"></a>

## Next pages — outside_static_routes / 221103313311 / 4

- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033112130132130-1230103000112231-2001132321320023-2022321213322022-1322032300220010-3101211302332031-3130030003201101-2230033012133201"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list — static_route_list / 330222220112 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211)
- voltstack_cluster_ar.outside_static_routes.static_route_list

<a id="canonical-3122032202101033-2010332030103202-1033330200122203-1301001010113330-1200132103220232-3320231021120002-3111122132210112-3302003102100230"></a>

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

<a id="canonical-3122310331030032-3310010202313220-3220331102123220-3233030011331333-2302021310101220-1032301300313122-3122132000020010-2131212202013223"></a>

## Direct properties — static_route_list / 330222220112 / 3

- [custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100): complete subsection reference.

<a id="canonical-3222123021332020-1321210013211203-3102210101220231-0111022332311330-2331301333303230-3030313202321130-0130120232331020-1000000232303311"></a>

<a id="canonical-3333000120032210-3332021130231102-3202013122132301-3321220203202310-3301032233313332-0333113311203103-0000032320103031-0301122223310231"></a>

## simple_static_route property — static_route_list / 330222220112 / 4

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

<a id="canonical-3321130131130000-3200233223232013-1303321213001132-1003023133232000-3020113300301211-3310021003300310-2131321310100020-2320030203120122"></a>

## Next pages — static_route_list / 330222220112 / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131311203222112-1323132323101303-3220122222330112-2302212213030213-2310011111102203-0232313131101202-0133333220213201-2131120122300202"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route — custom_static_route / 121033202230 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-1233100333332301-2233313223110103-3113330101212032-2022020221103013-0232320001223322-1201131130332212-1310113230121210-1310331311110120"></a>

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

<a id="canonical-0102112103320010-1323320301303003-0222103113220023-0313012010002322-0321033003133322-1002212221200212-3201010322330103-2302002203212202"></a>

## Direct properties — custom_static_route / 121033202230 / 3

<a id="canonical-3302001313130301-1213230223223201-0332111021123303-1011003033203333-2221121311101030-0122322330333001-0100223210133222-1112301300002211"></a>

<a id="canonical-0110322312301300-1203303321232002-0132112330022230-0231201231312130-1221012303233313-3333002232023211-2232310313111321-0112211331221022"></a>

## attrs property — custom_static_route / 121033202230 / 4

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

- [labels](resources--azure_vnet_site--reference--group-010.md#canonical-1303330213021213-2111111223010223-0301332330321203-1330013131031302-1303233031203211-1131201130223330-3200220001123113-2130223213120132): complete subsection reference.

- [nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-0121330202211330-1020332010322121-0122321221003033-0131000033202303-2021221020223310-3131032003322133-2330332000201201-3300321020200132): complete subsection reference.

- [subnets](resources--azure_vnet_site--reference--group-010.md#canonical-2010033100330020-1113112112012123-0132201302113233-1001101312221022-1323301100022003-3330220101212312-1303133211300112-0232322223122113): complete subsection reference.

<a id="canonical-2332303312232020-1331303231131331-2021100010310021-0021311223123333-2210112200301003-0302211300331121-0032101321303031-2323101211321120"></a>

## Next pages — custom_static_route / 121033202230 / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.labels](resources--azure_vnet_site--reference--group-010.md#canonical-1303330213021213-2111111223010223-0301332330321203-1330013131031302-1303233031203211-1131201130223330-3200220001123113-2130223213120132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-0121330202211330-1020332010322121-0122321221003033-0131000033202303-2021221020223310-3131032003322133-2330332000201201-3300321020200132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-010.md#canonical-2010033100330020-1113112112012123-0132201302113233-1001101312221022-1323301100022003-3330220101212312-1303133211300112-0232322223122113)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1303330213021213-2111111223010223-0301332330321203-1330013131031302-1303233031203211-1131201130223330-3200220001123113-2130223213120132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111213132310033-1302202032333123-2310032032310201-3111302112100202-0331301130113303-3210120200000032-0113120210310213-2021222331002220"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.labels — labels / 113100330312 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-0232330200021302-3322323201322013-0133101333120030-2222313332321303-0210332311133220-3231312313122223-2213131223300300-2223200300231132"></a>

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

<a id="canonical-3131122323223331-2110020102032123-3302133311110211-3230023233312030-3233203001133111-2130121323013210-3332101030023020-1231032120302212"></a>

## Direct properties — labels / 113100330312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121210002033211-1130333331203210-0121113233023102-3321012233102211-2020013113111021-2131212313032232-3033223123320002-0203222333123131"></a>

## Next pages — labels / 113100330312 / 4

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0121330202211330-1020332010322121-0122321221003033-0131000033202303-2021221020223310-3131032003322133-2330332000201201-3300321020200132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031013113000333-0212122111131122-0300202321333320-2033012320002223-0313030233220122-1323300012032301-3322212010330202-3130120110000232"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 130331223131 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-1322333212302100-3202021120100030-0000112012003213-1101111233213301-1032101023313123-3022001123232123-0233123023311132-1333022322021310"></a>

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

<a id="canonical-2310120033001113-1130230110333320-2321200231021323-0233200223300133-1312113320130012-2321132221102311-1300332100103313-1230230202322303"></a>

## Direct properties — nexthop / 130331223131 / 3

- [interface](resources--azure_vnet_site--reference--group-010.md#canonical-2013003231102233-3123210100232111-1202012331211030-0033312320303001-1032300120020310-0231331010120212-2100210221320333-0012123112233222): complete subsection reference.

- [nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-2332313200022212-3311122022200313-3131310133000231-1320112133033132-2130003011100321-2211113321100132-3123222222123300-1011303320001110): complete subsection reference.

<a id="canonical-2013120100220003-2133122302210202-0120300012112010-0313010232122231-2213122322302000-0001123021232332-0121001210230331-0331123100321322"></a>

<a id="canonical-1210030110310001-1311200031332013-2223201301313001-2222311122330311-0233011120032100-3022203202130322-0100130312001132-1103113121232303"></a>

## type property — nexthop / 130331223131 / 4

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

<a id="canonical-3120102322000030-3330233223220101-2231300132113131-2103031301323103-3320212121310120-3001213011112332-3010011230111202-2201200121113331"></a>

## Next pages — nexthop / 130331223131 / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--azure_vnet_site--reference--group-010.md#canonical-2013003231102233-3123210100232111-1202012331211030-0033312320303001-1032300120020310-0231331010120212-2100210221320333-0012123112233222)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-2332313200022212-3311122022200313-3131310133000231-1320112133033132-2130003011100321-2211113321100132-3123222222123300-1011303320001110)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2013003231102233-3123210100232111-1202012331211030-0033312320303001-1032300120020310-0231331010120212-2100210221320333-0012123112233222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031312230221102-0032013120131210-3220011103011000-3333000033321122-0102123001030102-3031222021231003-2003301301021013-0302200320011200"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 012230100021 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-0121330202211330-1020332010322121-0122321221003033-0131000033202303-2021221020223310-3131032003322133-2330332000201201-3300321020200132)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-1031021223022313-1220231321022013-1010331101203003-0121131220121203-1223303001112303-0013021202132322-0322121033203011-0303003231131300"></a>

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

<a id="canonical-0110211300322300-3301203012112001-0320123012212120-1222011302313323-1210212212302021-1031033320203023-0102100131210001-0032213022320321"></a>

## Direct properties — interface / 012230100021 / 3

<a id="canonical-2222232102021012-2120020012110112-1021101200203202-1011120210012002-0013232310303132-3112330301001230-1111133001230203-1313001312002021"></a>

<a id="canonical-0232130120033201-3031021023031020-2121231323032103-1332111223012220-1223031013300201-0320322010031300-3123122132333013-1030023032122032"></a>

## kind property — interface / 012230100021 / 4

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

<a id="canonical-1221322001330020-2213113311331030-3223013232332232-1120220122332212-1132011012120330-0012223200022212-1133220300120112-3103102111203332"></a>

<a id="canonical-3231202122130330-2312320000300000-3323121331202310-2301330221211230-1131003233232003-0231000113012112-1130013201133130-0232012030310232"></a>

## name property — interface / 012230100021 / 5

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

<a id="canonical-1013311203323013-0312120333023231-2311120110021221-3333101213113201-2002313121132103-1303003120230123-0212113003030202-1311033030311020"></a>

<a id="canonical-2123112320320000-1031100201231213-1233231122300123-0333203332330123-3101131323333310-1010032001020211-2323301303123310-3322013103013312"></a>

## namespace property — interface / 012230100021 / 6

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

<a id="canonical-0223113011123211-0212021200030122-3201010131203101-1211312231230003-2012301121011212-2210212233222111-2130132231120110-1222222210320231"></a>

<a id="canonical-2122130121130121-2003202222321232-3033233133000130-2232330000323121-3201000233022320-0133321133031032-1303001203300122-2312220030232313"></a>

## tenant property — interface / 012230100021 / 7

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

<a id="canonical-1322103013212130-1100021013231332-3222133020201123-3320221310320133-3123012122030130-0310320331312311-0010110210232310-1312110033331133"></a>

<a id="canonical-0031203322033032-1130210322211332-3213220002120302-3031310223332221-2001300200210012-0223233222012230-2010103111123022-0112120201213211"></a>

## uid property — interface / 012230100021 / 8

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

<a id="canonical-0211212211213311-3023212310331202-2311023211113202-3010011002320320-1011030110132213-1313222022230003-1121010311130102-3032210132313333"></a>

## Next pages — interface / 012230100021 / 9

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-0121330202211330-1020332010322121-0122321221003033-0131000033202303-2021221020223310-3131032003322133-2330332000201201-3300321020200132)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2332313200022212-3311122022200313-3131310133000231-1320112133033132-2130003011100321-2211113321100132-3123222222123300-1011303320001110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101010320303310-1030211021322222-1023330020233130-3313023221300203-3232130133333033-1321303133130203-2312101311031200-0023122310222232"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 001201013033 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-0121330202211330-1020332010322121-0122321221003033-0131000033202303-2021221020223310-3131032003322133-2330332000201201-3300321020200132)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-2133132120331031-2321232100200002-0202003333021011-2111310012320322-0013120323103010-0313333122101001-2222330020333312-2032100211020320"></a>

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

<a id="canonical-0220122012310211-0132210113321011-3000131220100233-2331330320200210-2301331130012133-1000111220113002-0311110100333201-0321302322023112"></a>

## Direct properties — nexthop_address / 001201013033 / 3

- [dual_stack](resources--azure_vnet_site--reference--group-010.md#canonical-2333202030223301-3332220211030021-0300021022033120-3011332113212321-3230312002320120-0223001223011001-0033320112320122-1013201323201012): complete subsection reference.

- [IPv4](resources--azure_vnet_site--reference--group-010.md#canonical-0111312122102321-1332211222000000-3302032123032203-3011230333300013-0001122000213200-0212101120232032-1123313010322303-2000003003200012): complete subsection reference.

- [IPv6](resources--azure_vnet_site--reference--group-010.md#canonical-1211130100020110-1230031021112103-1003133013101320-1113013201312121-1332012320032031-0012132321221212-3030100101230021-3212133233122322): complete subsection reference.

<a id="canonical-0330001013213033-3121222230121313-3332133232101221-0022133231312322-3001000223200000-2131222221030102-3222003032032132-2311013323110332"></a>

## Next pages — nexthop_address / 001201013033 / 4

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-010.md#canonical-2333202030223301-3332220211030021-0300021022033120-3011332113212321-3230312002320120-0223001223011001-0033320112320122-1013201323201012)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--azure_vnet_site--reference--group-010.md#canonical-0111312122102321-1332211222000000-3302032123032203-3011230333300013-0001122000213200-0212101120232032-1123313010322303-2000003003200012)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--azure_vnet_site--reference--group-010.md#canonical-1211130100020110-1230031021112103-1003133013101320-1113013201312121-1332012320032031-0012132321221212-3030100101230021-3212133233122322)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-0121330202211330-1020332010322121-0122321221003033-0131000033202303-2021221020223310-3131032003322133-2330332000201201-3300321020200132)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2333202030223301-3332220211030021-0300021022033120-3011332113212321-3230312002320120-0223001223011001-0033320112320122-1013201323201012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031133003110310-3002201312211320-3222210212023233-3132132233223323-1012330332011010-1121223211301122-1122130231122023-3000320023133321"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 033333102000 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-0121330202211330-1020332010322121-0122321221003033-0131000033202303-2021221020223310-3131032003322133-2330332000201201-3300321020200132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-2332313200022212-3311122022200313-3131310133000231-1320112133033132-2130003011100321-2211113321100132-3123222222123300-1011303320001110)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-0021230200023001-0212220202231310-2202322130030312-1221011211101103-0310200212130323-1020021021303312-3030101200313311-3232222032303003"></a>

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

<a id="canonical-0302221221133030-3130133222222122-0101230213031321-0032331103202311-3103013011000030-1300320230233230-0323101012212231-3033201333021001"></a>

## Direct properties — dual_stack / 033333102000 / 3

- [IPv4](resources--azure_vnet_site--reference--group-010.md#canonical-1202130211111322-1213033010322200-3200321000003121-3200113022010230-3122232123200100-1213113302122233-3022233011323012-2313322001311130): complete subsection reference.

- [IPv6](resources--azure_vnet_site--reference--group-010.md#canonical-2033303132133030-3102123030230303-0000202310110233-2110312223130103-3013321123130031-2103213101123223-1132130301233031-3001011331313111): complete subsection reference.

<a id="canonical-0223113000123023-2023311032221221-0110330312301010-1133213023013110-3202303031022020-0311100203101110-1202322212221322-1223013313311200"></a>

## Next pages — dual_stack / 033333102000 / 4

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--azure_vnet_site--reference--group-010.md#canonical-1202130211111322-1213033010322200-3200321000003121-3200113022010230-3122232123200100-1213113302122233-3022233011323012-2313322001311130)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--azure_vnet_site--reference--group-010.md#canonical-2033303132133030-3102123030230303-0000202310110233-2110312223130103-3013321123130031-2103213101123223-1132130301233031-3001011331313111)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-2332313200022212-3311122022200313-3131310133000231-1320112133033132-2130003011100321-2211113321100132-3123222222123300-1011303320001110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1202130211111322-1213033010322200-3200321000003121-3200113022010230-3122232123200100-1213113302122233-3022233011323012-2313322001311130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231122213231033-1120002010320200-3330111022003010-3210031123332322-0201131222333002-3303013033233303-2323012322133331-3122330111201202"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 031333120321 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-0121330202211330-1020332010322121-0122321221003033-0131000033202303-2021221020223310-3131032003322133-2330332000201201-3300321020200132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-2332313200022212-3311122022200313-3131310133000231-1320112133033132-2130003011100321-2211113321100132-3123222222123300-1011303320001110)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-010.md#canonical-2333202030223301-3332220211030021-0300021022033120-3011332113212321-3230312002320120-0223001223011001-0033320112320122-1013201323201012)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-3322322112123133-1002001023020222-3203321133121010-2233312100110023-2121030021013011-3000021220021033-0211332112231333-3021212333231231"></a>

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

<a id="canonical-0221001021323002-2213102202233122-3130331221112220-1232031021313222-2100101002312031-3132112033012131-3110202011233130-1122021212022012"></a>

## Direct properties — IPv4 / 031333120321 / 3

<a id="canonical-2131323313231023-3022320122133331-0013132130123223-1232120030222213-0111003132232323-0301132302020001-1121202001012002-2300001213120031"></a>

<a id="canonical-2202112213110303-3322020011312111-1332311212112002-2230311111211003-3213212223020222-0130222311030312-3212233121200223-2020013223213302"></a>

## addr property — IPv4 / 031333120321 / 4

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

<a id="canonical-2313101213210000-1020201112031223-3133033121130333-1022210313301211-0101230001222123-3032231212133300-2002311231321020-1220330013102301"></a>

## Next pages — IPv4 / 031333120321 / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-010.md#canonical-2333202030223301-3332220211030021-0300021022033120-3011332113212321-3230312002320120-0223001223011001-0033320112320122-1013201323201012)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2033303132133030-3102123030230303-0000202310110233-2110312223130103-3013321123130031-2103213101123223-1132130301233031-3001011331313111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323331202310012-1321220303003313-0223332320220233-0212113130312130-3302310001013033-3232333230121332-0311113300203121-0122321332031323"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 323032331002 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-0121330202211330-1020332010322121-0122321221003033-0131000033202303-2021221020223310-3131032003322133-2330332000201201-3300321020200132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-2332313200022212-3311122022200313-3131310133000231-1320112133033132-2130003011100321-2211113321100132-3123222222123300-1011303320001110)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-010.md#canonical-2333202030223301-3332220211030021-0300021022033120-3011332113212321-3230312002320120-0223001223011001-0033320112320122-1013201323201012)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-1111102302332211-2323112021223323-3022302131321331-0212032003020323-2112022113010222-1023301322310111-0322310220010033-0300111231210333"></a>

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

<a id="canonical-2330300313200310-0011112000201120-1331012333031232-2003223230323122-0112313030310003-1203320032222323-2333112100213001-3132213222011333"></a>

## Direct properties — IPv6 / 323032331002 / 3

<a id="canonical-0212322302020231-2122302333203000-3233000123200020-3032002010101300-2313020121301323-3121020331210012-3320200202032230-1322020032100032"></a>

<a id="canonical-1130331201312000-3012121310321222-1210030313022313-0331333030301132-3321130232220013-3301002100222111-2310011032011231-0202133322330200"></a>

## addr property — IPv6 / 323032331002 / 4

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

<a id="canonical-0223031003112310-2212030222200320-1333320033332132-0201210033313333-1332103211211311-2113231330220313-1020113221303130-3001102320012312"></a>

## Next pages — IPv6 / 323032331002 / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--azure_vnet_site--reference--group-010.md#canonical-2333202030223301-3332220211030021-0300021022033120-3011332113212321-3230312002320120-0223001223011001-0033320112320122-1013201323201012)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0111312122102321-1332211222000000-3302032123032203-3011230333300013-0001122000213200-0212101120232032-1123313010322303-2000003003200012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222313312121332-0200233111232212-3313101202313003-0121023221002202-1323021102100122-1030313022202231-3022231003333133-3002000032311212"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 002320201120 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-0121330202211330-1020332010322121-0122321221003033-0131000033202303-2021221020223310-3131032003322133-2330332000201201-3300321020200132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-2332313200022212-3311122022200313-3131310133000231-1320112133033132-2130003011100321-2211113321100132-3123222222123300-1011303320001110)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-2100131330132101-1032003032020312-2031030222313221-1323203121011013-0213023113231203-2222321100101332-3231113312221323-0331033011002100"></a>

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

<a id="canonical-0030211100330213-3323230230120213-2222113302222211-3120033101123223-3111123002003321-1300202012312212-2000331333022320-3310320303023230"></a>

## Direct properties — IPv4 / 002320201120 / 3

<a id="canonical-0112110301211200-1220212211213110-3333011123131333-3111123210211233-2233322222123220-3223010033323013-0031132022203202-3031013213022123"></a>

<a id="canonical-2233111312031333-2013311003113331-2330210222102221-0202211121012001-0112211120120302-3332232200131031-0023221031313132-3011101303223112"></a>

## addr property — IPv4 / 002320201120 / 4

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

<a id="canonical-1303313323030222-3011103031320101-1332311003101012-3221101001120332-0030033003310220-0113321311303211-2313023333001232-1220113110232202"></a>

## Next pages — IPv4 / 002320201120 / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-2332313200022212-3311122022200313-3131310133000231-1320112133033132-2130003011100321-2211113321100132-3123222222123300-1011303320001110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1211130100020110-1230031021112103-1003133013101320-1113013201312121-1332012320032031-0012132321221212-3030100101230021-3212133233122322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122113221312020-3110010120220121-2201023021111233-1012030212203130-0203212333210210-0223303211120000-0230203132333021-3131230233321120"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 321013011003 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--azure_vnet_site--reference--group-010.md#canonical-0121330202211330-1020332010322121-0122321221003033-0131000033202303-2021221020223310-3131032003322133-2330332000201201-3300321020200132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-2332313200022212-3311122022200313-3131310133000231-1320112133033132-2130003011100321-2211113321100132-3123222222123300-1011303320001110)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-1311131112132002-0111112331330023-3301320111022001-0031301203110112-2330020323300313-2120011110022120-1301121210023031-0111313332233102"></a>

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

<a id="canonical-2022102022320031-3203232132033213-3022211012020010-3310312002122023-1110131022320203-2323101311210333-2023030002121213-0031312313200232"></a>

## Direct properties — IPv6 / 321013011003 / 3

<a id="canonical-2301332100303233-1100333022231111-0131031022133011-2231132023122102-2022331332110023-0222133323311123-2230210031331303-3312120100121310"></a>

<a id="canonical-3331302111023302-3233022202223122-3233003222301320-1202201030220233-1233221230212322-2111302020102331-3111301213000113-0233122202310022"></a>

## addr property — IPv6 / 321013011003 / 4

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

<a id="canonical-3311210001100113-1321330302013210-1201022233112320-3101231133323222-3122202311110030-3112210111113202-1333203302233032-0102222121013111"></a>

## Next pages — IPv6 / 321013011003 / 5

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--azure_vnet_site--reference--group-010.md#canonical-2332313200022212-3311122022200313-3131310133000231-1320112133033132-2130003011100321-2211113321100132-3123222222123300-1011303320001110)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2010033100330020-1113112112012123-0132201302113233-1001101312221022-1323301100022003-3330220101212312-1303133211300112-0232322223122113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123321130300231-0331312011212000-3031301010333222-3233332102001302-2112012311002112-2130231303301020-0202322332021310-3123333333100013"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets — subnets / 121223303310 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-3211313311131123-1113033100201122-2100332022203211-0100000332111110-0300202033310211-1200330200113133-3122020222100313-1210020021200332"></a>

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

<a id="canonical-0332020130011303-1110123020200002-0101120003010023-3001311202333231-1110330221000010-3010102220230132-0231232311300303-1110310332102302"></a>

## Direct properties — subnets / 121223303310 / 3

- [IPv4](resources--azure_vnet_site--reference--group-010.md#canonical-3202202002223321-2332133230313332-0030120210223302-3210111221103122-1321111121013123-0100333022122220-2303311201330003-2033322003311310): complete subsection reference.

- [IPv6](resources--azure_vnet_site--reference--group-010.md#canonical-1010030320330322-3331331222003110-3313333023310123-3223213233012320-0312230202202120-0203100101131132-2201031300020100-2300103233330203): complete subsection reference.

<a id="canonical-3121101331110100-1303330112301232-1123200323301012-1031003102101312-0010023202030332-3300302022313111-0220201310012102-0200020210332033"></a>

## Next pages — subnets / 121223303310 / 4

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--azure_vnet_site--reference--group-010.md#canonical-3202202002223321-2332133230313332-0030120210223302-3210111221103122-1321111121013123-0100333022122220-2303311201330003-2033322003311310)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--azure_vnet_site--reference--group-010.md#canonical-1010030320330322-3331331222003110-3313333023310123-3223213233012320-0312230202202120-0203100101131132-2201031300020100-2300103233330203)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3202202002223321-2332133230313332-0030120210223302-3210111221103122-1321111121013123-0100333022122220-2303311201330003-2033322003311310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323322000123311-0230303330003311-0322322012132003-1213210230302231-0032323222212301-1011133021331111-2202032010312133-3302030320313020"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 300131302212 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-010.md#canonical-2010033100330020-1113112112012123-0132201302113233-1001101312221022-1323301100022003-3330220101212312-1303133211300112-0232322223122113)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-2312232001322202-3011331112221212-0310330133321201-0103220120311213-2102330202230330-0121132002232110-1000101212002211-2233210033303000"></a>

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

<a id="canonical-0310030022322123-0330020132311111-0203030013222202-2120012033333233-2033313233130222-2002302300102031-0003233111202211-2002102111200010"></a>

## Direct properties — IPv4 / 300131302212 / 3

<a id="canonical-0132301121021030-1002120301312132-1310202101203330-1232012032112032-2300333131233033-0101301010100022-0011000223113033-1231200231303212"></a>

<a id="canonical-0321000110113000-0002122213310301-0220311201012121-2330012201213213-2301312301233222-2010332100202230-2132211122333110-0223022132113131"></a>

## plen property — IPv4 / 300131302212 / 4

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

<a id="canonical-2331230211203131-1210113021033132-2331200313221202-3102302332131310-2013121223223002-1300110101311232-1102300223200020-1013113320112322"></a>

<a id="canonical-3032220312030113-1212022300132130-2231302231033320-1131033322230112-0321030222012221-0011130232301001-1212112312012021-3230201302130102"></a>

## prefix property — IPv4 / 300131302212 / 5

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

<a id="canonical-2103303132002103-0130121121113313-3003332203310332-2202220100112113-3323210302313032-1302122221012302-1211221110232232-3301310102010322"></a>

## Next pages — IPv4 / 300131302212 / 6

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-010.md#canonical-2010033100330020-1113112112012123-0132201302113233-1001101312221022-1323301100022003-3330220101212312-1303133211300112-0232322223122113)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1010030320330322-3331331222003110-3313333023310123-3223213233012320-0312230202202120-0203100101131132-2201031300020100-2300103233330203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131310322020233-2221022330320310-1211031123132201-0030000003131121-1030322213103121-3221310222031103-3223230100302121-0201032202030113"></a>

## voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 201230110213 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.outside_static_routes](resources--azure_vnet_site--reference--group-010.md#canonical-1110223330202103-2123101121321230-0023011211322110-0230332011111211-0121222212201220-1011322203210213-0100232102013331-1132331133223211)
- [voltstack_cluster_ar.outside_static_routes.static_route_list](resources--azure_vnet_site--reference--group-010.md#canonical-1212332112111233-0221023211021303-3232222312030210-0331300233323003-1031001222201322-0000333323201212-0133113300033000-3203231331332132)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](resources--azure_vnet_site--reference--group-010.md#canonical-3122332022101031-1300112312211232-1321001301311222-2013033213302110-0321111123203110-0203310321021013-0102000023001330-1020131132112100)
- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-010.md#canonical-2010033100330020-1113112112012123-0132201302113233-1001101312221022-1323301100022003-3330220101212312-1303133211300112-0232322223122113)
- voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-2320022311221003-0320301302130203-2013300212220330-3320011023312112-3033111130202031-0200210302101321-3031103311031003-0213121210220120"></a>

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

<a id="canonical-2223033122323223-3223000332130001-2032331323221333-0002211221320201-0230113310313322-3200233012120210-2000231013031001-2013012221102210"></a>

## Direct properties — IPv6 / 201230110213 / 3

<a id="canonical-3120333301210220-3213100121131021-0003102200032023-3311212232121132-2020122202200011-0210001021222232-0031022212102103-2030120010221233"></a>

<a id="canonical-2111210330203202-1121033032113230-0001200001310323-1100112211103111-3111120020301122-0132222022101310-0021302133302213-0303210313013031"></a>

## plen property — IPv6 / 201230110213 / 4

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

<a id="canonical-3133302331120131-1213321203000023-1131213011031332-1322120202000033-1003123300311231-1210112211111133-1120221202231333-0231102111222201"></a>

<a id="canonical-3300021022112130-3120011020312123-1030002233013323-2311132003332123-0301300210300320-3310130030310032-2022220300211003-1022100001323223"></a>

## prefix property — IPv6 / 201230110213 / 5

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

<a id="canonical-1203131122122310-1030122210032112-3213210002313022-3003103101332313-1111000310012121-2330323333203110-3220020013133333-0132121203130320"></a>

## Next pages — IPv6 / 201230110213 / 6

- [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets](resources--azure_vnet_site--reference--group-010.md#canonical-2010033100330020-1113112112012123-0132201302113233-1001101312221022-1323301100022003-3330220101212312-1303133211300112-0232322223122113)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0201321201012211-3131100121131013-0010232103310121-3330001101211033-1231022213213233-2112001113230320-1301330313102111-1130121102200103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311203022102013-0101113232231031-2220333220211120-2101121033301210-3210320103011310-3203231300023002-0233010212302310-2121210310003333"></a>

## voltstack_cluster_ar.sm_connection_public_ip — sm_connection_public_ip / 023312133031 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.sm_connection_public_ip

<a id="canonical-1231310130323112-3220013200311203-0203113231332101-2233001310222221-1313233230210203-2022210311233333-1102221102023330-3100020221301323"></a>

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

<a id="canonical-1103332123222003-2223313211131231-1003303012032021-3120201112021213-1301031232032233-2321230231133322-2030033300032031-0332020313320220"></a>

## Direct properties — sm_connection_public_ip / 023312133031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1132122012132201-3112300221123213-3023011132220210-3130223222333033-1021220020213033-1221221212012131-2011002200213131-1223323231231132"></a>

## Next pages — sm_connection_public_ip / 023312133031 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-1000100033323233-1201011002000310-1223202013111033-0000100231103002-2022132213031021-3230322203312213-2002023303301213-3213201333110102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230011001022302-0013223211010202-1002322330203210-2222310100122300-0203133221312103-1011132123211233-1201032323232200-3201003002210103"></a>

## voltstack_cluster_ar.sm_connection_pvt_ip — sm_connection_pvt_ip / 033233202212 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.sm_connection_pvt_ip

<a id="canonical-1333003202123133-0133203120220223-1321322212123001-2101133130121132-3013113101212300-1021210132100123-2131033022212013-1123210103311003"></a>

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

<a id="canonical-0033030021033302-1311101013000012-3301331310220133-1111102011333130-3200031101101331-2103330033111222-1021111010213320-1120121003101212"></a>

## Direct properties — sm_connection_pvt_ip / 033233202212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030322321102221-1320121300022000-1323000003103311-0300133210310001-2323032313233121-0220313103130220-3020220222010020-2231233130232103"></a>

## Next pages — sm_connection_pvt_ip / 033233202212 / 4

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-3333220321002101-1012013013201131-3112320212130103-2132312101213111-3011030103300022-0201231001332133-2223211331031221-1312311232113101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103020100001030-3113323213302031-1211112230023023-1120310202331311-3212200203123303-3220302310131032-3103102221220220-2322301320133323"></a>

## voltstack_cluster_ar.storage_class_list — storage_class_list / 203003121003 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- voltstack_cluster_ar.storage_class_list

<a id="canonical-1113113031201321-3030201323211122-2203021233012330-1022201003120301-2101210103301113-1221031313313301-0300123020110332-2232020013120321"></a>

Type: `"object"`. single nested block, Optional.

Add additional custom storage classes in Kubernetes for this site.

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
storage_class_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120201010101003-0313120222313131-0223102332032310-0133302321003132-1313333231201131-2201110111313000-0310033121321331-0333010303202221"></a>

## Direct properties — storage_class_list / 203003121003 / 3

- [storage_classes](resources--azure_vnet_site--reference--group-010.md#canonical-2202310203302333-2232312100213231-3132312310310130-0023223321002322-0332303231312011-0101020212132111-3323212201012201-0013102110003113): complete subsection reference.

<a id="canonical-0001222200110233-3322131312111232-1233233332202120-1113112130012321-2103021031020213-1322101111310031-2132001212303213-2100323211022132"></a>

## Next pages — storage_class_list / 203003121003 / 4

- [voltstack_cluster_ar.storage_class_list.storage_classes](resources--azure_vnet_site--reference--group-010.md#canonical-2202310203302333-2232312100213231-3132312310310130-0023223321002322-0332303231312011-0101020212132111-3323212201012201-0013102110003113)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2202310203302333-2232312100213231-3132312310310130-0023223321002322-0332303231312011-0101020212132111-3323212201012201-0013102110003113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031302320102003-3133013030130301-3213113320233113-1332133233230231-3230232101133133-0133232332232322-2133231313131130-2333112211000002"></a>

## voltstack_cluster_ar.storage_class_list.storage_classes — storage_classes / 330213023233 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133)
- [voltstack_cluster_ar.storage_class_list](resources--azure_vnet_site--reference--group-010.md#canonical-3333220321002101-1012013013201131-3112320212130103-2132312101213111-3011030103300022-0201231001332133-2223211331031221-1312311232113101)
- voltstack_cluster_ar.storage_class_list.storage_classes

<a id="canonical-2120310022321122-2301021232012132-0313012303031213-1213010123132322-0330311111211300-1321102132213013-1203022311031033-1023112211301113"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_class_name")}
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage_classes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222332332232303-2333033220231022-2211310003223322-0022123222003003-2112132221113132-0211203331221203-1000021102211333-3101110000120312"></a>

## Direct properties — storage_classes / 330213023233 / 3

<a id="canonical-3211110131012113-2211023210110223-3220002102331011-1332101331021303-1003301221223220-1331133121233222-3000030203012030-0032333312023320"></a>

<a id="canonical-3232021102310122-3013000320022033-0231101001130203-0210121310211230-2011311022231122-2210022201120101-0211311321010013-0132301303012121"></a>

## default_storage_class property — storage_classes / 330213023233 / 4

Type: `"bool"`. Optional.

Make this storage class default storage class for the K8s cluster.

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

<a id="canonical-2313033111131313-3012020021113232-0322312113021112-2002311323232222-3300012011203021-0022321110021103-3113232202011323-0133311010121300"></a>

<a id="canonical-1222212310001333-2232103112121033-1321133123212332-0220033330121101-0212203121113133-0322000013303000-1022231211300322-1221333013102330"></a>

## storage_class_name property — storage_classes / 330213023233 / 5

Type: `"string"`. Optional.

Name of the storage class as it will appear in K8s.

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
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-2010230100101310-3132011113221232-3310132002002202-3003000121310022-3113102202101203-1211120220013301-1201122312221213-1103122232003203"></a>

## Next pages — storage_classes / 330213023233 / 6

- [voltstack_cluster_ar.storage_class_list](resources--azure_vnet_site--reference--group-010.md#canonical-3333220321002101-1012013013201131-3112320212130103-2132312101213111-3011030103300022-0201231001332133-2223211331031221-1312311232113101)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2000101211023100-3303010230231233-0133333212103100-1101121132322201-3121232303021110-3333033122312013-2303300130313111-3332232300233123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013230102321320-0212213012221111-0223213211220300-0003231223200003-2100210011302202-1032022011311203-3312311201110211-0312213103112203"></a>

## waf_signatures — waf_signatures / 031113023130 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- waf_signatures

<a id="canonical-0330133130002000-1010022012000231-1002113122123222-1101321311200131-2312132331233023-2202113002213212-0100133233110302-0311012102011320"></a>

Type: `"object"`. single nested block, Optional.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("automatic",
    "manual")}
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
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

Terraform syntax:

```terraform
waf_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300203312323113-1213021301032302-2222003312300320-0030023021003122-0302300022122312-3102020310311011-1202122022133203-3031333103133131"></a>

## Direct properties — waf_signatures / 031113023130 / 3

- [automatic](resources--azure_vnet_site--reference--group-010.md#canonical-0023231113011233-1232332000220322-2132212021132033-1303133132122313-3110130022113131-2133003203010312-3331312303221231-0111232012002203): complete subsection reference.

- [manual](resources--azure_vnet_site--reference--group-010.md#canonical-2122230023013322-2131001022032001-3213123331200000-3322202230102220-1320333131333233-1113233222030001-1112010102220112-3103003313232120): complete subsection reference.

<a id="canonical-0102131313132033-1001330201331133-1330133230133303-0322000202300010-2103331303110003-3100011103000332-3003212232021032-0302221102212022"></a>

## Next pages — waf_signatures / 031113023130 / 4

- [waf_signatures.automatic](resources--azure_vnet_site--reference--group-010.md#canonical-0023231113011233-1232332000220322-2132212021132033-1303133132122313-3110130022113131-2133003203010312-3331312303221231-0111232012002203)
- [waf_signatures.manual](resources--azure_vnet_site--reference--group-010.md#canonical-2122230023013322-2131001022032001-3213123331200000-3322202230102220-1320333131333233-1113233222030001-1112010102220112-3103003313232120)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-0023231113011233-1232332000220322-2132212021132033-1303133132122313-3110130022113131-2133003203010312-3331312303221231-0111232012002203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230113321133321-3210013101011323-0211221332113002-2301032031302213-3021112202230331-0102122010200323-3210232112003021-2323320332332033"></a>

## waf_signatures.automatic — automatic / 002311310001 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [waf_signatures](resources--azure_vnet_site--reference--group-010.md#canonical-2000101211023100-3303010230231233-0133333212103100-1101121132322201-3121232303021110-3333033122312013-2303300130313111-3332232300233123)
- waf_signatures.automatic

<a id="canonical-1212232010332132-1023011100223133-1303330132121213-0021013013022101-3103230330200131-3311112110010321-3300031303023202-3120032223212210"></a>

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
automatic = {}
```

<a id="canonical-1010332321020213-0302122311302013-2221200223201212-1033133110222031-0003101213130330-3312302112001010-1031230310312102-1301130123012213"></a>

## Direct properties — automatic / 002311310001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302321320221131-2331032223302123-1310111022000021-3103013012110210-1210022030131223-0210303001012102-1010331230330132-1130002200133100"></a>

## Next pages — automatic / 002311310001 / 4

- [waf_signatures](resources--azure_vnet_site--reference--group-010.md#canonical-2000101211023100-3303010230231233-0133333212103100-1101121132322201-3121232303021110-3333033122312013-2303300130313111-3332232300233123)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)

<a id="canonical-2122230023013322-2131001022032001-3213123331200000-3322202230102220-1320333131333233-1113233222030001-1112010102220112-3103003313232120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203312302331023-1003323110330122-0313010321222330-1233203202313301-3120121332313231-3312312320331312-0231220030201200-2120010131101231"></a>

## waf_signatures.manual — manual / 031120213110 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- [Property reference](resources--azure_vnet_site--reference--group-001.md#canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232)
- [waf_signatures](resources--azure_vnet_site--reference--group-010.md#canonical-2000101211023100-3303010230231233-0133333212103100-1101121132322201-3121232303021110-3333033122312013-2303300130313111-3332232300233123)
- waf_signatures.manual

<a id="canonical-0210013210032323-0201221120100311-0230133121231223-0233332231032023-0223300203212321-1332113230311112-1022130313223031-1020033011021311"></a>

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
manual = {}
```

<a id="canonical-3032332032230100-0011323200321020-2000332330333113-3130013100122000-1300032222132330-0323132113222011-1021020032132003-3223301023100002"></a>

## Direct properties — manual / 031120213110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303312220102331-2331100302103302-2123220320213221-0010123100233311-0002302213103230-3313122123130222-2020113213020332-0212301130123310"></a>

## Next pages — manual / 031120213110 / 4

- [waf_signatures](resources--azure_vnet_site--reference--group-010.md#canonical-2000101211023100-3303010230231233-0133333212103100-1101121132322201-3121232303021110-3333033122312013-2303300130313111-3332232300233123)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
