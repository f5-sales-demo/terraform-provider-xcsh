---
page_title: "xcsh_cloud_user_account reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_user_account reference."
---

# xcsh_cloud_user_account reference

<a id="canonical-3120211002101200-1032113201021233-1220320323313131-0212322201031232-0233112321202013-1303123010311023-0111001022220010-3320211201030303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323302222111012-0112103021033010-2010322102230021-3232011320331021-1102323210021000-1231123033002001-3211210300113322-3133301123023322"></a>

## Property reference — Property reference / 100022000032 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)
- Property reference

<a id="canonical-3321112102001333-2330111331102103-1303100302332333-3110312000200131-2012320212201031-1230022103333021-1313333002221021-1231002113223133"></a>

## Direct properties — Property reference / 100022000032 / 3

<a id="canonical-0322223011331031-2210002230202113-1232123032021020-3003203320022100-0110321330312102-1202132121021023-1032022000322312-1110211212213210"></a>

<a id="canonical-1023332121100001-3302022022323300-2333003210323302-1121220322000230-2112220132112213-1023023021032002-0323132110203300-3122002333223123"></a>

## annotations property — Property reference / 100022000032 / 4

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

- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-0012012233031021-3100322321221333-2032113331301133-3312231010011311-1132101332122101-0022323230302012-2330201031132100-3032303003121110): complete subsection reference.

<a id="canonical-3022001002201330-1302011013113020-2312100101113302-0031033212022001-3023121200010002-1212013103021002-2221123003303133-3230330231310201"></a>

<a id="canonical-1203113201300011-0213120122313300-2203233230323102-1210031100030120-0010112212120030-2332111203100012-0110021020232103-1000121301202010"></a>

## description property — Property reference / 100022000032 / 5

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

<a id="canonical-0331201110301101-1201322102033113-2133130023120001-0010232301212002-2131323111303110-3123200123231211-2021330113030100-0031230231300123"></a>

<a id="canonical-1013012302331202-2322011203330211-1333222123333013-0320030312001013-2213011001033302-3131311201233302-1231200110210201-3010231333223303"></a>

## disable property — Property reference / 100022000032 / 6

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

<a id="canonical-3133011021212223-0332323300212222-2320123003122031-0112201320323212-2101121133313013-3113002211202231-0320323332012103-0303123332001002"></a>

<a id="canonical-2131300003212320-2312311330112211-2033230210321002-3210011022011030-0230220123211320-0133231003200222-0131312220302000-0111321102331220"></a>

## ID property — Property reference / 100022000032 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3103322020011102-3113331300131131-3232201111302020-0231131110333131-2033311230123020-3330120003213102-1200121002010033-3031312203212003"></a>

<a id="canonical-3212213222130020-3320033210201301-0322002320032310-3220011011113131-2011321003132011-2332203323300222-1031112213122022-0210330301001230"></a>

## labels property — Property reference / 100022000032 / 8

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

<a id="canonical-3333110221220123-1201120230030332-1123203112312023-1031331103221021-2003210030020233-0010000003112132-0300333031201103-3331121220122222"></a>

<a id="canonical-2031000133221301-2013331323122320-3122011301100210-3203321332032320-2312222101032120-1202010212002233-3210002010022101-3011113221201011"></a>

## name property — Property reference / 100022000032 / 9

Type: `"string"`. Required.

Name of the Cloud User Account. Must be unique within the namespace.

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

<a id="canonical-1100001113230132-3122120303201320-2031220221321103-3213000303110322-3132013332303220-1301331010020123-3131123000233210-2110212033330332"></a>

<a id="canonical-2331103020000013-3102213313232233-3233303233110332-2021320132323131-2133200000200132-1032220001201321-2021233231021222-3101310103100131"></a>

## namespace property — Property reference / 100022000032 / 10

Type: `"string"`. Required.

Namespace where the Cloud User Account is created.

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

- [timeouts](resources--cloud_user_account--reference--group-001.md#canonical-0221113313211300-2231231022211031-3221203221120320-0031001120313020-3123202310330231-1322011032313320-0320332131030022-2003021023203233): complete subsection reference.

<a id="canonical-0133313232302002-0122001233232313-0203112131111301-3032331301102332-2210332130303020-1312321013202132-3232113033332200-2130021310032031"></a>

## All schema paths — Property reference / 100022000032 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cloud_user_account--reference--group-001.md#canonical-0322223011331031-2210002230202113-1232123032021020-3003203320022100-0110321330312102-1202132121021023-1032022000322312-1110211212213210) |
| `aws_provider` | [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-2033103103101013-3230002333301102-2313100123113120-0102033023201203-2033313312213023-3003112220321321-3002202321301010-3310000310212322) |
| `aws_provider.aws_account_number` | [aws_provider.aws_account_number](resources--cloud_user_account--reference--group-001.md#canonical-1010120002313303-0111313233330312-0130131223302121-2211021001101322-0203101022001002-0103131131123111-0231232013212231-0000332022010221) |
| `aws_provider.aws_assume_role` | [aws_provider.aws_assume_role](resources--cloud_user_account--reference--group-001.md#canonical-0323023331003323-2022000102113130-2332302003032223-1211322000310023-2201322233010023-1310023312030010-2331013030303000-0000020320321113) |
| `aws_provider.aws_assume_role.custom_external_id` | [aws_provider.aws_assume_role.custom_external_id](resources--cloud_user_account--reference--group-001.md#canonical-2112012333101330-0013312222031323-0302320033300100-3303021020113120-2133330311212110-2221313221320221-0131103010221320-2130132303323310) |
| `aws_provider.aws_assume_role.duration_seconds` | [aws_provider.aws_assume_role.duration_seconds](resources--cloud_user_account--reference--group-001.md#canonical-2222300232202220-0213100321121032-0220303013001323-3230202032202123-1133131200120312-1320003103022030-0012131012233023-1230302200212201) |
| `aws_provider.aws_assume_role.external_id_is_optional` | [aws_provider.aws_assume_role.external_id_is_optional](resources--cloud_user_account--reference--group-001.md#canonical-2232131023122233-3010303123033033-1323133032022111-2231012312323110-0011321222301021-3130120001230200-0120103210122100-3301122213030121) |
| `aws_provider.aws_assume_role.external_id_is_tenant_id` | [aws_provider.aws_assume_role.external_id_is_tenant_id](resources--cloud_user_account--reference--group-001.md#canonical-2033330000123012-1302313220101223-0222000122310103-2310121103323033-3111123000132023-1202322211002303-2021330032222332-3233033022203002) |
| `aws_provider.aws_assume_role.role_arn` | [aws_provider.aws_assume_role.role_arn](resources--cloud_user_account--reference--group-001.md#canonical-1213223311223023-0300300003313312-0032121213020101-1301123033322032-3133232233321210-0220321230123202-1030011112300102-2113230002031000) |
| `aws_provider.aws_assume_role.session_name` | [aws_provider.aws_assume_role.session_name](resources--cloud_user_account--reference--group-001.md#canonical-1001222233130000-0132322332133002-3013011030301233-2020101220100313-1022000333210031-1010302023021313-0322101113132220-2213211011210310) |
| `aws_provider.aws_assume_role.session_tags` | [aws_provider.aws_assume_role.session_tags](resources--cloud_user_account--reference--group-001.md#canonical-0101310313110020-2213013103222002-3030103013131123-0111212010311303-1332203101113103-1211013013131333-2110221010302101-0030111000030220) |
| `aws_provider.aws_secret_key` | [aws_provider.aws_secret_key](resources--cloud_user_account--reference--group-001.md#canonical-2123223001321213-2201122000302313-1123331330233303-1021232032331120-3200101110302303-3021300210101233-0131120013230212-1322301303321300) |
| `aws_provider.aws_secret_key.access_key` | [aws_provider.aws_secret_key.access_key](resources--cloud_user_account--reference--group-001.md#canonical-0230101231110132-3111023233213330-1200202210232020-3111322122001301-3130211023212301-0030010003322230-3101011113230120-0321323333311200) |
| `aws_provider.aws_secret_key.secret_key` | [aws_provider.aws_secret_key.secret_key](resources--cloud_user_account--reference--group-001.md#canonical-1322220312033111-3123302003303010-2233123201023200-0323123122322321-2110213313012131-3013000012323011-1231233132113303-3023113321333103) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info](resources--cloud_user_account--reference--group-001.md#canonical-0321233302201020-2012320120213231-0222031200221010-1133310023322321-1122030312201202-1013310322003313-0212132200002123-3213210033121123) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.decryption_provider` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info.decryption_provider](resources--cloud_user_account--reference--group-001.md#canonical-1121100010132212-3320031100301301-3212210012233031-0331023031033122-1130202231311233-2210110210320133-2202322011201212-3333233110222233) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.location` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info.location](resources--cloud_user_account--reference--group-001.md#canonical-2033011331222000-1001031022010301-3220101031312212-2212032111102200-1001101020302323-1232123011013212-0112312121133310-2021012332110022) |
| `aws_provider.aws_secret_key.secret_key.blindfold_secret_info.store_provider` | [aws_provider.aws_secret_key.secret_key.blindfold_secret_info.store_provider](resources--cloud_user_account--reference--group-001.md#canonical-0102201112321320-1000302000213232-0012311320333110-1231232022330123-2331202202123002-1211110312220133-2303231031122323-2202020233132313) |
| `aws_provider.aws_secret_key.secret_key.clear_secret_info` | [aws_provider.aws_secret_key.secret_key.clear_secret_info](resources--cloud_user_account--reference--group-001.md#canonical-3012300012031122-1222333213111011-1011331300033231-0113220023020331-2131022210232031-0021011230333300-3330310310211233-1003303210112102) |
| `aws_provider.aws_secret_key.secret_key.clear_secret_info.provider_ref` | [aws_provider.aws_secret_key.secret_key.clear_secret_info.provider_ref](resources--cloud_user_account--reference--group-001.md#canonical-1122112320130033-1213113202303103-1130220301122103-3233022230002333-1322003003220121-0031031032021312-3312202300233331-0113030333123222) |
| `aws_provider.aws_secret_key.secret_key.clear_secret_info.url` | [aws_provider.aws_secret_key.secret_key.clear_secret_info.url](resources--cloud_user_account--reference--group-001.md#canonical-3031010311102103-1320030330121332-0203112222230202-1300330020212331-3200031203131123-3232330010001032-1101002002220133-0313300010002002) |
| `description` | [description](resources--cloud_user_account--reference--group-001.md#canonical-3022001002201330-1302011013113020-2312100101113302-0031033212022001-3023121200010002-1212013103021002-2221123003303133-3230330231310201) |
| `disable` | [disable](resources--cloud_user_account--reference--group-001.md#canonical-0331201110301101-1201322102033113-2133130023120001-0010232301212002-2131323111303110-3123200123231211-2021330113030100-0031230231300123) |
| `id` | [id](resources--cloud_user_account--reference--group-001.md#canonical-3133011021212223-0332323300212222-2320123003122031-0112201320323212-2101121133313013-3113002211202231-0320323332012103-0303123332001002) |
| `labels` | [labels](resources--cloud_user_account--reference--group-001.md#canonical-3103322020011102-3113331300131131-3232201111302020-0231131110333131-2033311230123020-3330120003213102-1200121002010033-3031312203212003) |
| `name` | [name](resources--cloud_user_account--reference--group-001.md#canonical-3333110221220123-1201120230030332-1123203112312023-1031331103221021-2003210030020233-0010000003112132-0300333031201103-3331121220122222) |
| `namespace` | [namespace](resources--cloud_user_account--reference--group-001.md#canonical-1100001113230132-3122120303201320-2031220221321103-3213000303110322-3132013332303220-1301331010020123-3131123000233210-2110212033330332) |
| `timeouts` | [timeouts](resources--cloud_user_account--reference--group-001.md#canonical-0212030212210223-1100020003213223-2330311020332021-0222030230121231-1003311302003302-2202211110330103-1230023021223310-2231300123301110) |
| `timeouts.create` | [timeouts.create](resources--cloud_user_account--reference--group-001.md#canonical-3021330023230202-1220013333021320-0322202233223223-1223130033121122-1111120231002200-3232102021333210-1032213102123021-3313021211102023) |
| `timeouts.delete` | [timeouts.delete](resources--cloud_user_account--reference--group-001.md#canonical-0122031313010310-0333120221230123-0333120302120012-3011001122030121-0201132110313131-3313233112132012-0321303101003022-3132302102110131) |
| `timeouts.read` | [timeouts.read](resources--cloud_user_account--reference--group-001.md#canonical-1120321303333300-3100033332303211-1100311123112330-2022013031101332-0223301103302313-1032210020302020-3232111311023311-2102010032330010) |
| `timeouts.update` | [timeouts.update](resources--cloud_user_account--reference--group-001.md#canonical-3002122231331221-2311123203320101-3011100313311321-2113202222333312-0331102111103031-1133331010302322-1031313112332233-3303300003121112) |

<a id="canonical-2323301312332011-2201122332212001-0103213121131312-2211123010233320-0010231013303031-1313302031131100-3012123103313103-2210012320002132"></a>

## Next pages — Property reference / 100022000032 / 12

- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-0012012233031021-3100322321221333-2032113331301133-3312231010011311-1132101332122101-0022323230302012-2330201031132100-3032303003121110)
- [timeouts](resources--cloud_user_account--reference--group-001.md#canonical-0221113313211300-2231231022211031-3221203221120320-0031001120313020-3123202310330231-1322011032313320-0320332131030022-2003021023203233)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)

<a id="canonical-0012012233031021-3100322321221333-2032113331301133-3312231010011311-1132101332122101-0022323230302012-2330201031132100-3032303003121110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320200222111221-1300011122110033-2231110013221110-0010200302201001-2220213323313223-3031113323321230-1302033312210121-1201011010313021"></a>

## aws_provider — aws_provider / 030310301310 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-3120211002101200-1032113201021233-1220320323313131-0212322201031232-0233112321202013-1303123010311023-0111001022220010-3320211201030303)
- aws_provider

<a id="canonical-2033103103101013-3230002333301102-2313100123113120-0102033023201203-2033313312213023-3003112220321321-3002202321301010-3310000310212322"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for aws provider.

Upstream description:

Create AWS Provider Type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_account_number"),
  validators.ConflictingObjectAttributes("aws_assume_role",
    "aws_secret_key")}
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
  "x-ves-oneof-field-aws_authentication_type": "[\"aws_assume_role\",\"aws_secret_key\"]"
}
```

Terraform syntax:

```terraform
aws_provider {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200330110001203-2113322223310011-3120220113031011-3332301103210012-2223202100010310-3123132222212001-0300013012330011-1012122131310030"></a>

## Direct properties — aws_provider / 030310301310 / 3

<a id="canonical-1010120002313303-0111313233330312-0130131223302121-2211021001101322-0203101022001002-0103131131123111-0231232013212231-0000332022010221"></a>

<a id="canonical-0003232200023133-0121201232102103-2233130330220100-0223032320122222-3120222031322132-2233230101302023-0003322310200133-2323020121231312"></a>

## aws_account_number property — aws_provider / 030310301310 / 4

Type: `"string"`. Optional.

Account Number. 12 Digit Account Number.

Upstream description:

12 Digit Account Number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.uint64.gte": "100000000000",
    "ves.io.schema.rules.uint64.lte": "999999999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "100000000000",
    "ves.io.schema.rules.uint64.lte": "999999999999"
  }
}
```

- [aws_assume_role](resources--cloud_user_account--reference--group-001.md#canonical-0203122032220210-0210012000031000-3100110332321333-1031003211333210-3230132231332023-0000321022123333-3332133100122113-0033301221200031): complete subsection reference.

- [aws_secret_key](resources--cloud_user_account--reference--group-001.md#canonical-3210120310120211-1033322231223001-3203001032300111-3111213102301102-0100132230013100-1012330113030302-1102203301011011-1232122300323130): complete subsection reference.

<a id="canonical-0232331210013032-2303222202132131-3132300230323023-3203313103100032-1121231003123000-1133322133303232-3221311302022301-1202320320023310"></a>

## Next pages — aws_provider / 030310301310 / 5

- [aws_provider.aws_assume_role](resources--cloud_user_account--reference--group-001.md#canonical-0203122032220210-0210012000031000-3100110332321333-1031003211333210-3230132231332023-0000321022123333-3332133100122113-0033301221200031)
- [aws_provider.aws_secret_key](resources--cloud_user_account--reference--group-001.md#canonical-3210120310120211-1033322231223001-3203001032300111-3111213102301102-0100132230013100-1012330113030302-1102203301011011-1232122300323130)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-3120211002101200-1032113201021233-1220320323313131-0212322201031232-0233112321202013-1303123010311023-0111001022220010-3320211201030303)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)

<a id="canonical-0203122032220210-0210012000031000-3100110332321333-1031003211333210-3230132231332023-0000321022123333-3332133100122113-0033301221200031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322200022020210-3223121323112203-3032233330321111-1103110303222123-0333002021122232-1303220031120112-2312012100221312-3110301300203100"></a>

## aws_provider.aws_assume_role — aws_assume_role / 332211310303 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-3120211002101200-1032113201021233-1220320323313131-0212322201031232-0233112321202013-1303123010311023-0111001022220010-3320211201030303)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-0012012233031021-3100322321221333-2032113331301133-3312231010011311-1132101332122101-0022323230302012-2330201031132100-3032303003121110)
- aws_provider.aws_assume_role

<a id="canonical-0323023331003323-2022000102113130-2332302003032223-1211322000310023-2201322233010023-1310023312030010-2331013030303000-0000020320321113"></a>

Type: `"object"`. single nested block, Optional.

AWS Assume Role to Handle Delegated Access.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration_seconds",
    "role_arn",
    "session_name"),
  validators.ConflictingObjectAttributes("custom_external_id",
    "external_id_is_optional"),
  validators.ConflictingObjectAttributes("custom_external_id",
    "external_id_is_tenant_id"),
  validators.ConflictingObjectAttributes("external_id_is_optional",
    "external_id_is_tenant_id")}
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
  "x-ves-oneof-field-external_id": "[\"custom_external_id\",\"external_id_is_optional\",\"external_id_is_tenant_id\"]"
}
```

Terraform syntax:

```terraform
aws_assume_role {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100330213213311-2232000000001133-3123003222300121-0220302230011032-1233000201000320-2330003210200101-2131312203331000-3331200322332032"></a>

## Direct properties — aws_assume_role / 332211310303 / 3

<a id="canonical-2112012333101330-0013312222031323-0302320033300100-3303021020113120-2133330311212110-2221313221320221-0131103010221320-2130132303323310"></a>

<a id="canonical-3032221213210123-1031313320202321-3200010320101112-0300322313012320-3010301130211000-1311230000313123-2113330302010113-2023311132020033"></a>

## custom_external_id property — aws_assume_role / 332211310303 / 4

Type: `"string"`. Optional.

Exclusive with \[external\_id\_is\_optional external\_id\_is\_tenant\_id\] External ID is Custom ID.

Upstream description:

Exclusive with \[external\_id\_is\_optional external\_id\_is\_tenant\_id\] External ID is Custom ID.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2"
  }
}
```

<a id="canonical-2222300232202220-0213100321121032-0220303013001323-3230202032202123-1133131200120312-1320003103022030-0012131012233023-1230302200212201"></a>

<a id="canonical-3011123013221022-3002310332330222-2120323322023133-0213102031300320-0303300222033221-2202210220111102-3211210332001231-1110320322333132"></a>

## duration_seconds property — aws_assume_role / 332211310303 / 5

Type: `"number"`. Optional.

The duration, in seconds of the role session.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(3600, 43200),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 43200,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 3600
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "43200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "43200"
  }
}
```

- [external_id_is_optional](resources--cloud_user_account--reference--group-001.md#canonical-0111331230012032-3300101331303302-0102302032121321-1132003103313323-0231031210303310-0003123003123011-2221210012032123-1000330030113332): complete subsection reference.

- [external_id_is_tenant_id](resources--cloud_user_account--reference--group-001.md#canonical-2333223212133110-1003021331020013-2212033203133211-0121301122012121-2230110212200112-2033303011223300-0203022312322123-0100133101311120): complete subsection reference.

<a id="canonical-1213223311223023-0300300003313312-0032121213020101-1301123033322032-3133232233321210-0220321230123202-1030011112300102-2113230002031000"></a>

<a id="canonical-3233323100213033-2311030223011330-1300212321131301-0233121323123010-2113022323202320-0333313302222313-2003120010323300-2230101300133000"></a>

## role_arn property — aws_assume_role / 332211310303 / 6

Type: `"string"`. Optional.

IAM Role ARN. IAM Role ARN to assume the role.

Upstream description:

IAM Role ARN to assume the role.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(20, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 20,
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
    "minLength": 20,
    "pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "20",
    "ves.io.schema.rules.string.pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048",
    "ves.io.schema.rules.string.min_len": "20",
    "ves.io.schema.rules.string.pattern": "^(arn:aws:iam::)([0-9]{12}:role/.*)$"
  }
}
```

<a id="canonical-1001222233130000-0132322332133002-3013011030301233-2020101220100313-1022000333210031-1010302023021313-0322101113132220-2213211011210310"></a>

<a id="canonical-2323033310321111-3113010300113030-2213123231103311-2111113230230020-0320031011321031-3313022130332212-2203012020103303-1232310123230323"></a>

## session_name property — aws_assume_role / 332211310303 / 7

Type: `"string"`. Optional.

Use the role session name to uniquely identify a session, which will be used for deploy, monitor
from F5XC console.

Upstream description:

Use the role session name to uniquely identify a session, which will be used for deploy, monitor
from F5XC console.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 2,
    "pattern": "[\\\\w+=,.@-]*"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2",
    "ves.io.schema.rules.string.pattern": "[\\\\w+=,.@-]*"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "2",
    "ves.io.schema.rules.string.pattern": "[\\\\w+=,.@-]*"
  }
}
```

<a id="canonical-0101310313110020-2213013103222002-3030103013131123-0111212010311303-1332203101113103-1211013013131333-2110221010302101-0030111000030220"></a>

<a id="canonical-0011012111131100-2333011030230332-1313111130330331-2200030120300223-1202222131022133-2020321002233100-3001321333222122-3023213131232110"></a>

## session_tags property — aws_assume_role / 332211310303 / 8

Type: `["map", "string"]`. Optional.

Session tags are key-value pair attributes that you pass when you assume an IAM role.

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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

<a id="canonical-0032030121031332-2122331211011320-1321100102130310-3232021102112323-3013313130330030-1323013013200013-1312012312022223-1133123301312131"></a>

## Next pages — aws_assume_role / 332211310303 / 9

- [aws_provider.aws_assume_role.external_id_is_optional](resources--cloud_user_account--reference--group-001.md#canonical-0111331230012032-3300101331303302-0102302032121321-1132003103313323-0231031210303310-0003123003123011-2221210012032123-1000330030113332)
- [aws_provider.aws_assume_role.external_id_is_tenant_id](resources--cloud_user_account--reference--group-001.md#canonical-2333223212133110-1003021331020013-2212033203133211-0121301122012121-2230110212200112-2033303011223300-0203022312322123-0100133101311120)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-0012012233031021-3100322321221333-2032113331301133-3312231010011311-1132101332122101-0022323230302012-2330201031132100-3032303003121110)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)

<a id="canonical-0111331230012032-3300101331303302-0102302032121321-1132003103313323-0231031210303310-0003123003123011-2221210012032123-1000330030113332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202230002223313-2313232233133210-0030113021300030-3001321003101332-1020122103310330-1200200320031121-3330123003200032-1130311310321031"></a>

## aws_provider.aws_assume_role.external_id_is_optional — external_id_is_optional / 131211213002 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-3120211002101200-1032113201021233-1220320323313131-0212322201031232-0233112321202013-1303123010311023-0111001022220010-3320211201030303)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-0012012233031021-3100322321221333-2032113331301133-3312231010011311-1132101332122101-0022323230302012-2330201031132100-3032303003121110)
- [aws_provider.aws_assume_role](resources--cloud_user_account--reference--group-001.md#canonical-0203122032220210-0210012000031000-3100110332321333-1031003211333210-3230132231332023-0000321022123333-3332133100122113-0033301221200031)
- aws_provider.aws_assume_role.external_id_is_optional

<a id="canonical-2232131023122233-3010303123033033-1323133032022111-2231012312323110-0011321222301021-3130120001230200-0120103210122100-3301122213030121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for external ID is optional.

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
external_id_is_optional = {}
```

<a id="canonical-2221211231232030-3323223301203212-0232300310311112-0123003303020102-3132123033023100-2333301121010133-0311000201100323-3133300021320003"></a>

## Direct properties — external_id_is_optional / 131211213002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231132102023102-0330013220311200-3103200130311010-3201230021301121-1323232023313313-1110023310102213-0123113013320001-2231322230323131"></a>

## Next pages — external_id_is_optional / 131211213002 / 4

- [aws_provider.aws_assume_role](resources--cloud_user_account--reference--group-001.md#canonical-0203122032220210-0210012000031000-3100110332321333-1031003211333210-3230132231332023-0000321022123333-3332133100122113-0033301221200031)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)

<a id="canonical-2333223212133110-1003021331020013-2212033203133211-0121301122012121-2230110212200112-2033303011223300-0203022312322123-0100133101311120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000303032032303-2002123203102313-3003213331003000-3001123123220333-1333123112330011-0233133310222323-0101133122002201-1330313222120032"></a>

## aws_provider.aws_assume_role.external_id_is_tenant_id — external_id_is_tenant_id / 213033013110 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-3120211002101200-1032113201021233-1220320323313131-0212322201031232-0233112321202013-1303123010311023-0111001022220010-3320211201030303)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-0012012233031021-3100322321221333-2032113331301133-3312231010011311-1132101332122101-0022323230302012-2330201031132100-3032303003121110)
- [aws_provider.aws_assume_role](resources--cloud_user_account--reference--group-001.md#canonical-0203122032220210-0210012000031000-3100110332321333-1031003211333210-3230132231332023-0000321022123333-3332133100122113-0033301221200031)
- aws_provider.aws_assume_role.external_id_is_tenant_id

<a id="canonical-2033330000123012-1302313220101223-0222000122310103-2310121103323033-3111123000132023-1202322211002303-2021330032222332-3233033022203002"></a>

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
external_id_is_tenant_id = {}
```

<a id="canonical-3103002330101332-1132001331203303-2313222003001033-0010323320213311-0301202131000302-3120022202202022-0330110022222011-3202120031313220"></a>

## Direct properties — external_id_is_tenant_id / 213033013110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030011011103102-1320101323213013-3303100203003121-0223223121330113-3322303332120032-0033132021120003-0002003301133232-2210301022201323"></a>

## Next pages — external_id_is_tenant_id / 213033013110 / 4

- [aws_provider.aws_assume_role](resources--cloud_user_account--reference--group-001.md#canonical-0203122032220210-0210012000031000-3100110332321333-1031003211333210-3230132231332023-0000321022123333-3332133100122113-0033301221200031)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)

<a id="canonical-3210120310120211-1033322231223001-3203001032300111-3111213102301102-0100132230013100-1012330113030302-1102203301011011-1232122300323130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121033302032323-2202210002001013-3122310001233111-0200230011103003-1202322303110020-2313332110321331-0323332210310212-3310022333312020"></a>

## aws_provider.aws_secret_key — aws_secret_key / 010331111211 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-3120211002101200-1032113201021233-1220320323313131-0212322201031232-0233112321202013-1303123010311023-0111001022220010-3320211201030303)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-0012012233031021-3100322321221333-2032113331301133-3312231010011311-1132101332122101-0022323230302012-2330201031132100-3032303003121110)
- aws_provider.aws_secret_key

<a id="canonical-2123223001321213-2201122000302313-1123331330233303-1021232032331120-3200101110302303-3021300210101233-0131120013230212-1322301303321300"></a>

Type: `"object"`. single nested block, Optional.

AWS Programmatic Access Credentials type.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("access_key")}
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
aws_secret_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-3331031001213213-2103320313032102-3213323100321000-3202130233320320-3303221222322311-0100122132212010-1020200113120131-0323002102103201"></a>

## Direct properties — aws_secret_key / 010331111211 / 3

<a id="canonical-0230101231110132-3111023233213330-1200202210232020-3111322122001301-3130211023212301-0030010003322230-3101011113230120-0321323333311200"></a>

<a id="canonical-0220021102121221-1221103210231010-1313302313200122-1301120100310001-1301320323232301-1231320010222323-3331121331210010-0321031110312232"></a>

## access_key property — aws_secret_key / 010331111211 / 4

Type: `"string"`. Optional.

Access Key ID. Access key ID for your AWS account.

Upstream description:

Access key ID for your AWS account.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [secret_key](resources--cloud_user_account--reference--group-001.md#canonical-2001221302332211-2120100113112113-0332202130220011-0122302022011012-2013320011211203-2112222032103022-3201233201320212-1222131110323002): complete subsection reference.

<a id="canonical-0332122303032112-2311130013311301-2031000330320213-3100001112113302-2302201232020112-2020230213012132-1020302312213113-0100110003322120"></a>

## Next pages — aws_secret_key / 010331111211 / 5

- [aws_provider.aws_secret_key.secret_key](resources--cloud_user_account--reference--group-001.md#canonical-2001221302332211-2120100113112113-0332202130220011-0122302022011012-2013320011211203-2112222032103022-3201233201320212-1222131110323002)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-0012012233031021-3100322321221333-2032113331301133-3312231010011311-1132101332122101-0022323230302012-2330201031132100-3032303003121110)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)

<a id="canonical-2001221302332211-2120100113112113-0332202130220011-0122302022011012-2013320011211203-2112222032103022-3201233201320212-1222131110323002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321003122312131-2001333013010233-0200322131233230-1121111010233323-1120033031300321-0220013001222110-1301200333312002-3122120213312133"></a>

## aws_provider.aws_secret_key.secret_key — secret_key / 131233333002 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-3120211002101200-1032113201021233-1220320323313131-0212322201031232-0233112321202013-1303123010311023-0111001022220010-3320211201030303)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-0012012233031021-3100322321221333-2032113331301133-3312231010011311-1132101332122101-0022323230302012-2330201031132100-3032303003121110)
- [aws_provider.aws_secret_key](resources--cloud_user_account--reference--group-001.md#canonical-3210120310120211-1033322231223001-3203001032300111-3111213102301102-0100132230013100-1012330113030302-1102203301011011-1232122300323130)
- aws_provider.aws_secret_key.secret_key

<a id="canonical-1322220312033111-3123302003303010-2233123201023200-0323123122322321-2110213313012131-3013000012323011-1231233132113303-3023113321333103"></a>

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
secret_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-0013211223032200-2120320022211132-0232332221012321-1130123220100202-1013132301002200-0203112020030320-1102020003032103-1132110300131333"></a>

## Direct properties — secret_key / 131233333002 / 3

- [blindfold_secret_info](resources--cloud_user_account--reference--group-001.md#canonical-0031333121202102-1230320232132103-2300023311220330-3111331100113333-3130032020332200-2300012222130123-2301302130002310-1120023001231000): complete subsection reference.

- [clear_secret_info](resources--cloud_user_account--reference--group-001.md#canonical-1133313101003120-3013100103310303-2231312031111233-0222222011322110-3233320133321320-2211032113123301-1002331322322331-2320323030303113): complete subsection reference.

<a id="canonical-1123133110310121-3300111000233233-1312322231321212-2223103331000202-0030012312323212-0300101132102202-2133032310213303-2111023030030221"></a>

## Next pages — secret_key / 131233333002 / 4

- [aws_provider.aws_secret_key.secret_key.blindfold_secret_info](resources--cloud_user_account--reference--group-001.md#canonical-0031333121202102-1230320232132103-2300023311220330-3111331100113333-3130032020332200-2300012222130123-2301302130002310-1120023001231000)
- [aws_provider.aws_secret_key.secret_key.clear_secret_info](resources--cloud_user_account--reference--group-001.md#canonical-1133313101003120-3013100103310303-2231312031111233-0222222011322110-3233320133321320-2211032113123301-1002331322322331-2320323030303113)
- [aws_provider.aws_secret_key](resources--cloud_user_account--reference--group-001.md#canonical-3210120310120211-1033322231223001-3203001032300111-3111213102301102-0100132230013100-1012330113030302-1102203301011011-1232122300323130)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)

<a id="canonical-0031333121202102-1230320232132103-2300023311220330-3111331100113333-3130032020332200-2300012222130123-2301302130002310-1120023001231000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302331212013212-3102011110120001-0333230103302033-2211302102211100-3223322032320021-1010320001000210-1330233113212103-0223231332103310"></a>

## aws_provider.aws_secret_key.secret_key.blindfold_secret_info — blindfold_secret_info / 123300232322 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-3120211002101200-1032113201021233-1220320323313131-0212322201031232-0233112321202013-1303123010311023-0111001022220010-3320211201030303)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-0012012233031021-3100322321221333-2032113331301133-3312231010011311-1132101332122101-0022323230302012-2330201031132100-3032303003121110)
- [aws_provider.aws_secret_key](resources--cloud_user_account--reference--group-001.md#canonical-3210120310120211-1033322231223001-3203001032300111-3111213102301102-0100132230013100-1012330113030302-1102203301011011-1232122300323130)
- [aws_provider.aws_secret_key.secret_key](resources--cloud_user_account--reference--group-001.md#canonical-2001221302332211-2120100113112113-0332202130220011-0122302022011012-2013320011211203-2112222032103022-3201233201320212-1222131110323002)
- aws_provider.aws_secret_key.secret_key.blindfold_secret_info

<a id="canonical-0321233302201020-2012320120213231-0222031200221010-1133310023322321-1122030312201202-1013310322003313-0212132200002123-3213210033121123"></a>

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

<a id="canonical-3113101133023322-3211032002110112-1303121013333302-0210010123121222-3221000211332300-0211001331023100-0210012230311022-0212313020033131"></a>

## Direct properties — blindfold_secret_info / 123300232322 / 3

<a id="canonical-1121100010132212-3320031100301301-3212210012233031-0331023031033122-1130202231311233-2210110210320133-2202322011201212-3333233110222233"></a>

<a id="canonical-0101213201300122-2013032201301210-0111321213102133-0331133313222320-0202132303131222-0210113033231102-2001021231133021-2210333010232312"></a>

## decryption_provider property — blindfold_secret_info / 123300232322 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2033011331222000-1001031022010301-3220101031312212-2212032111102200-1001101020302323-1232123011013212-0112312121133310-2021012332110022"></a>

<a id="canonical-2002020232103312-2231301303033122-3003331121302102-1230130002000101-0300320123321133-0002231112011113-0201233323130122-1113123323210202"></a>

## location property — blindfold_secret_info / 123300232322 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0102201112321320-1000302000213232-0012311320333110-1231232022330123-2331202202123002-1211110312220133-2303231031122323-2202020233132313"></a>

<a id="canonical-3210120102321330-0313121132311003-0210310103300001-0132222213323033-1221021320110111-1122323333303120-2013123121033110-1100023232323000"></a>

## store_provider property — blindfold_secret_info / 123300232322 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3011321110130032-3303033002000212-2033122301111133-2012111200020230-0133233120131212-2223312222311220-0010030321133001-0020222221003013"></a>

## Next pages — blindfold_secret_info / 123300232322 / 7

- [aws_provider.aws_secret_key.secret_key](resources--cloud_user_account--reference--group-001.md#canonical-2001221302332211-2120100113112113-0332202130220011-0122302022011012-2013320011211203-2112222032103022-3201233201320212-1222131110323002)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)

<a id="canonical-1133313101003120-3013100103310303-2231312031111233-0222222011322110-3233320133321320-2211032113123301-1002331322322331-2320323030303113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221230202123211-3132122320012321-2300123002210131-2332103323000003-1201301313122032-3002122023210322-2101230121211223-3003101322330203"></a>

## aws_provider.aws_secret_key.secret_key.clear_secret_info — clear_secret_info / 313332332032 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-3120211002101200-1032113201021233-1220320323313131-0212322201031232-0233112321202013-1303123010311023-0111001022220010-3320211201030303)
- [aws_provider](resources--cloud_user_account--reference--group-001.md#canonical-0012012233031021-3100322321221333-2032113331301133-3312231010011311-1132101332122101-0022323230302012-2330201031132100-3032303003121110)
- [aws_provider.aws_secret_key](resources--cloud_user_account--reference--group-001.md#canonical-3210120310120211-1033322231223001-3203001032300111-3111213102301102-0100132230013100-1012330113030302-1102203301011011-1232122300323130)
- [aws_provider.aws_secret_key.secret_key](resources--cloud_user_account--reference--group-001.md#canonical-2001221302332211-2120100113112113-0332202130220011-0122302022011012-2013320011211203-2112222032103022-3201233201320212-1222131110323002)
- aws_provider.aws_secret_key.secret_key.clear_secret_info

<a id="canonical-3012300012031122-1222333213111011-1011331300033231-0113220023020331-2131022210232031-0021011230333300-3330310310211233-1003303210112102"></a>

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

<a id="canonical-0021303111031201-2012333122233033-3110212333032312-1322302310021021-2103030311001131-1102213002100020-3333230101322210-1112001323122122"></a>

## Direct properties — clear_secret_info / 313332332032 / 3

<a id="canonical-1122112320130033-1213113202303103-1130220301122103-3233022230002333-1322003003220121-0031031032021312-3312202300233331-0113030333123222"></a>

<a id="canonical-0122303121221132-0320220201021330-3011203020000100-0212210000333022-0000111001311131-0120231302310022-3232002330001320-3131031020303301"></a>

## provider_ref property — clear_secret_info / 313332332032 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3031010311102103-1320030330121332-0203112222230202-1300330020212331-3200031203131123-3232330010001032-1101002002220133-0313300010002002"></a>

<a id="canonical-3231101220023213-1222233112031110-0113113111313322-1112223312312221-2220322212030330-0123012311210103-1023011200202312-3211030200112003"></a>

## URL property — clear_secret_info / 313332332032 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2222132332333200-3223113033020110-0130022001331301-1211302111332131-0020101102001123-2023133221301003-1110211112201320-2320202332113000"></a>

## Next pages — clear_secret_info / 313332332032 / 6

- [aws_provider.aws_secret_key.secret_key](resources--cloud_user_account--reference--group-001.md#canonical-2001221302332211-2120100113112113-0332202130220011-0122302022011012-2013320011211203-2112222032103022-3201233201320212-1222131110323002)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)

<a id="canonical-0221113313211300-2231231022211031-3221203221120320-0031001120313020-3123202310330231-1322011032313320-0320332131030022-2003021023203233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020221232110100-0303023111221013-1330311301103101-2111330321021333-2202313030210302-0110330102130021-3211031023111033-0002222121200010"></a>

## timeouts — timeouts / 302230301212 / 2

Breadcrumbs:

- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)
- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-3120211002101200-1032113201021233-1220320323313131-0212322201031232-0233112321202013-1303123010311023-0111001022220010-3320211201030303)
- timeouts

<a id="canonical-0212030212210223-1100020003213223-2330311020332021-0222030230121231-1003311302003302-2202211110330103-1230023021223310-2231300123301110"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0023113332011003-2223232203133132-2312030113201032-0322312211220330-0113210301211022-0123313311230112-2010003002223323-2010001122330133"></a>

## Direct properties — timeouts / 302230301212 / 3

<a id="canonical-3021330023230202-1220013333021320-0322202233223223-1223130033121122-1111120231002200-3232102021333210-1032213102123021-3313021211102023"></a>

<a id="canonical-2000011102222023-1311101210210013-0322103312023302-3310301230120100-3230223333032201-2312021010301233-2322010210332220-2112211322303232"></a>

## create property — timeouts / 302230301212 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0122031313010310-0333120221230123-0333120302120012-3011001122030121-0201132110313131-3313233112132012-0321303101003022-3132302102110131"></a>

<a id="canonical-3122132331230301-3303311312230203-0020301331332023-0312310021131302-0120200113313131-1303100222032211-1113323111213221-2031312031121113"></a>

## delete property — timeouts / 302230301212 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1120321303333300-3100033332303211-1100311123112330-2022013031101332-0223301103302313-1032210020302020-3232111311023311-2102010032330010"></a>

<a id="canonical-2211032122013130-2300231313032130-1022320301311033-0122303000332313-3322232232111233-2012121033220233-1321201033331312-2033213202101022"></a>

## read property — timeouts / 302230301212 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3002122231331221-2311123203320101-3011100313311321-2113202222333312-0331102111103031-1133331010302322-1031313112332233-3303300003121112"></a>

<a id="canonical-1313202003233222-2032331311331201-2131003210021012-1331212330332000-3203222001102132-3013220232033012-1232330100203320-0312331003001320"></a>

## update property — timeouts / 302230301212 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1223301032300320-2131113223202020-1211201111321301-0303010200120231-0033210002220333-3222013213103103-0000130101113133-3312101000213111"></a>

## Next pages — timeouts / 302230301212 / 8

- [Property reference](resources--cloud_user_account--reference--group-001.md#canonical-3120211002101200-1032113201021233-1220320323313131-0212322201031232-0233112321202013-1303123010311023-0111001022220010-3320211201030303)
- [xcsh_cloud_user_account](../resources/cloud_user_account.md#canonical-1031132021211202-3000130330113323-2203111122011131-3322201313323003-2113322122001212-2031120103110031-3122331032002130-0200201122101321)
