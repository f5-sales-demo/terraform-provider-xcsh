---
page_title: "xcsh_alert_policy reference"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_alert_policy reference."
---

# xcsh_alert_policy reference

<a id="canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- Property reference

<a id="canonical-1220220110201120-3130200121101302-1300211001023320-0111211023031020-2233133211311130-2323202112312311-1012311320323210-2000010220121203"></a>

### Direct properties for `xcsh_alert_policy`

<a id="canonical-3220103313110311-3300203033121000-0030022220111223-2312301022201312-2120113211130331-0232233112021321-2210000002123001-3112110010320323"></a>

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

<a id="canonical-0310122130132301-3012132131132311-2320002230220222-1313233133113320-2101132220330133-1013333102110213-1323022120120311-0302001011212121"></a>

<a id="canonical-2032101123021001-1300311123013010-3002202220010321-0030022032310011-1120123100311303-1122032130120310-1222322111330200-3310021120212233"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2111112322122311-2213133223201233-3021110331021333-1212133022100113-0332203220312223-0133331332323232-0212030100202023-0301002221200001"></a>

<a id="canonical-0300002012333120-1000132101332032-3300322311303023-3003001312320211-1021022112222300-0221102330121013-3302200021223101-1211312213220232"></a>

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

<a id="canonical-3200112312330000-0003012111333133-1300301313212330-0023201332211312-2021102012103211-3122131322111101-3131131301002301-3232103330033011"></a>

<a id="canonical-0111221312023210-3213312133023210-2033230333011122-1111303013023323-0003021013013303-0322323011031321-2222323212312220-3030013010123312"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1003121230321203-0321231120331013-2201221200301110-3211131022302331-0122321201103302-0330133200013222-0202313001010113-3203013320212310"></a>

<a id="canonical-2322233001021021-3322230331232312-1333021010021232-0101130030001100-3100202331333303-0213213003001330-0223033013311010-0203022112000100"></a>

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

<a id="canonical-0312021202033111-2221111322133101-1303201302122012-3030022010123020-0013030231010312-0132010202221203-1211232121202011-2110200103101300"></a>

<a id="canonical-0321113332033130-1321332221221130-2322221021330303-0103111022112333-1322010103232233-0120030031133202-3301101300211031-3312132320212300"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Alert Policy. Must be unique within the namespace.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1333001031121103-2110311012233332-0303333002211000-2213113312130302-1202033212020002-1212322201130300-0120120003131012-3112122012001033"></a>

<a id="canonical-1313322303102101-1333021103032123-0022333032233313-2200132233232211-0100002312311323-3021120101110222-2232302330122000-3220231231233020"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Alert Policy is created.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-2302223211220010-3133333223123323-0320311202331210-3230201211033133-0220001322010300-2323033221331022-2321110220123020-0201210012212021): complete subsection reference.

- [receivers](resources--alert_policy--reference--group-001.md#canonical-3131322330100230-2131122213031210-1231122112211222-3020300103001122-2130202110300030-0001002122021021-2210020203303312-2111210332131202): complete subsection reference.

- [routes](resources--alert_policy--reference--group-001.md#canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033): complete subsection reference.

- [timeouts](resources--alert_policy--reference--group-001.md#canonical-0232311332023021-1302331123211130-1333333130122230-1313103331111032-3302231011300011-1002113030003001-1130133123122213-2330301331133121): complete subsection reference.

<a id="canonical-3320211221112232-1102210201010120-1313120113010103-0320222212211230-3221220122002022-1112102111132103-3022110130230102-1330000100033301"></a>

### All schema paths for `xcsh_alert_policy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--alert_policy--reference--group-001.md#canonical-3220103313110311-3300203033121000-0030022220111223-2312301022201312-2120113211130331-0232233112021321-2210000002123001-3112110010320323) |
| `description` | [description](resources--alert_policy--reference--group-001.md#canonical-0310122130132301-3012132131132311-2320002230220222-1313233133113320-2101132220330133-1013333102110213-1323022120120311-0302001011212121) |
| `disable` | [disable](resources--alert_policy--reference--group-001.md#canonical-2111112322122311-2213133223201233-3021110331021333-1212133022100113-0332203220312223-0133331332323232-0212030100202023-0301002221200001) |
| `id` | [ID](resources--alert_policy--reference--group-001.md#canonical-3200112312330000-0003012111333133-1300301313212330-0023201332211312-2021102012103211-3122131322111101-3131131301002301-3232103330033011) |
| `labels` | [labels](resources--alert_policy--reference--group-001.md#canonical-1003121230321203-0321231120331013-2201221200301110-3211131022302331-0122321201103302-0330133200013222-0202313001010113-3203013320212310) |
| `name` | [name](resources--alert_policy--reference--group-001.md#canonical-0312021202033111-2221111322133101-1303201302122012-3030022010123020-0013030231010312-0132010202221203-1211232121202011-2110200103101300) |
| `namespace` | [namespace](resources--alert_policy--reference--group-001.md#canonical-1333001031121103-2110311012233332-0303333002211000-2213113312130302-1202033212020002-1212322201130300-0120120003131012-3112122012001033) |
| `notification_parameters` | [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-3033201321301122-3130300012120220-0031332223003032-3012033312301023-3013321112233003-2222300321130002-0113001322222222-2033113223220323) |
| `notification_parameters.custom` | [notification_parameters.custom](resources--alert_policy--reference--group-001.md#canonical-3212310323101102-3011211212213201-0022030231102102-0323321030132203-0301122030302121-2200111031312102-2213112032223011-0013133311123102) |
| `notification_parameters.custom.labels` | [notification_parameters.custom.labels](resources--alert_policy--reference--group-001.md#canonical-3213003110221023-2112020210032302-3103202103022020-3032111220213130-0100103021312312-1310231311120101-3221323221022102-3300331010120222) |
| `notification_parameters.default` | [notification_parameters.default](resources--alert_policy--reference--group-001.md#canonical-3322030211113131-1320130211022033-0122133112203220-1310122220131031-2032211333323110-1332112131132232-2311212101320121-2210200132313022) |
| `notification_parameters.group_interval` | [notification_parameters.group_interval](resources--alert_policy--reference--group-001.md#canonical-3100302232300130-1122312101210211-1212223312213111-0211303103221131-1202313302010311-2130323120233321-2301110233022121-2000221112031110) |
| `notification_parameters.group_wait` | [notification_parameters.group_wait](resources--alert_policy--reference--group-001.md#canonical-2013001323222332-2213113320120010-1120032323133023-1110212302122303-2322000210320111-1030131111313011-0001233022223002-3112132322012231) |
| `notification_parameters.individual` | [notification_parameters.individual](resources--alert_policy--reference--group-001.md#canonical-3032110111022333-1023320221223100-2322232212110210-2122120323231231-0022132301302231-0031223333131111-2010323320302231-1022023110320203) |
| `notification_parameters.repeat_interval` | [notification_parameters.repeat_interval](resources--alert_policy--reference--group-001.md#canonical-2112212112212311-3232132201223023-2303320011313033-3323020003320110-2002210311320000-1300230213120313-0320313210303203-2113213000000113) |
| `notification_parameters.ves_io_group` | [notification_parameters.ves_io_group](resources--alert_policy--reference--group-001.md#canonical-3003231023000303-3101012031310332-0331131332322122-0133220311200033-0110323322112221-1312023103212330-2231210001313032-1112030322023222) |
| `receivers` | [receivers](resources--alert_policy--reference--group-001.md#canonical-3022110002301212-1313313133003301-1020133033012100-3133320230233311-1231023311020322-0003303012321123-0222110301003112-2233210232323232) |
| `receivers.kind` | [receivers.kind](resources--alert_policy--reference--group-001.md#canonical-1011123210332213-2232103011230200-0200120323001101-1013103001002303-2122013020201133-1312033012203333-2022201123211031-1310101211220110) |
| `receivers.name` | [receivers.name](resources--alert_policy--reference--group-001.md#canonical-1332030101103131-1203120111012330-3120213112210210-3303012233033021-0233001322331320-1002030131200000-1203132031110020-2033100321103321) |
| `receivers.namespace` | [receivers.namespace](resources--alert_policy--reference--group-001.md#canonical-1203331031330301-2322102131223212-2322203001201230-2230102333033023-0232212120221123-1031230202212113-0310223122300030-2200021123020231) |
| `receivers.tenant` | [receivers.tenant](resources--alert_policy--reference--group-001.md#canonical-1021332013020120-0002201332022030-2201310303002012-1031300102110003-3010030201220213-3110212023220312-2210110123122100-1323002323122132) |
| `receivers.uid` | [receivers.uid](resources--alert_policy--reference--group-001.md#canonical-1303321010332121-2132221021302011-3233001313023022-0201233322113022-3212213313313202-0130320100212312-2211010003032313-0001023333130210) |
| `routes` | [routes](resources--alert_policy--reference--group-001.md#canonical-0301321223132032-0312112021303120-1330003130200103-1103103033301201-3310110301100033-3332202222133303-2302002211113211-1011000100332200) |
| `routes.alertname` | [routes.alertname](resources--alert_policy--reference--group-001.md#canonical-0320031112311211-3032022332222133-3112112223101303-2203113213223122-3211022301202333-1132122001111330-1330131020031112-1322230232322302) |
| `routes.alertname_regex` | [routes.alertname_regex](resources--alert_policy--reference--group-001.md#canonical-1102100033130312-3133031111301310-3230213122120311-3233123203000111-0331010001103202-2020113232301132-3131013122313120-0012313011010303) |
| `routes.any` | [routes.any](resources--alert_policy--reference--group-001.md#canonical-2002001320200300-3331123220030113-0202020323323303-3210003013230023-3222002131211112-3332332122112320-0201021211200101-3010201201103321) |
| `routes.custom` | [routes.custom](resources--alert_policy--reference--group-001.md#canonical-0121300113230032-3310201302322311-0121210012133301-1212220132331122-3202020021033300-2301011200023313-3230300323121130-2032121100310003) |
| `routes.custom.alertlabel` | [routes.custom.alertlabel](resources--alert_policy--reference--group-001.md#canonical-3031023300000023-0311031110020223-0033332130333333-2111333210123030-3313021123333220-2330011223112223-3113321131132032-1100110231112003) |
| `routes.custom.alertname` | [routes.custom.alertname](resources--alert_policy--reference--group-001.md#canonical-2333101113300213-1203031230130213-2312213020023020-2223320312322321-0000322323012310-2010021111113131-2311210230301232-1233031302021212) |
| `routes.custom.alertname.exact_match` | [routes.custom.alertname.exact_match](resources--alert_policy--reference--group-001.md#canonical-2033112023012032-3002002323231011-0103130120301301-3213001323111113-0102230233012233-2003022122101230-2321221222313102-3011300231021200) |
| `routes.custom.alertname.regex_match` | [routes.custom.alertname.regex_match](resources--alert_policy--reference--group-001.md#canonical-1100223213310010-0022212331312101-0003320003032112-0111232233011123-1103213321013021-3110000111133123-2233003100023222-0322210222321030) |
| `routes.custom.group` | [routes.custom.group](resources--alert_policy--reference--group-001.md#canonical-1311123122031003-2131030113300323-2022131331331213-1302233302300000-1312113131110003-0320302222122203-3230233311113323-0320221230231020) |
| `routes.custom.group.exact_match` | [routes.custom.group.exact_match](resources--alert_policy--reference--group-001.md#canonical-3230131122130112-3212031220331022-1233233122303303-3203131031222233-2112211232201031-1203013333001113-2112301102001300-2300113220203332) |
| `routes.custom.group.regex_match` | [routes.custom.group.regex_match](resources--alert_policy--reference--group-001.md#canonical-1103310132120201-2112133133220303-3113130230012110-0032330131012333-3130010222102012-1033220303332231-3333323132323310-3112120201131102) |
| `routes.custom.severity` | [routes.custom.severity](resources--alert_policy--reference--group-001.md#canonical-3022233222123223-2232020231131132-3231121130111201-0311120200200320-2300130131323101-1220122133133230-2332123132223222-1021103203021122) |
| `routes.custom.severity.exact_match` | [routes.custom.severity.exact_match](resources--alert_policy--reference--group-001.md#canonical-2100131220212303-1133201031112120-3001030132011321-1310313031303211-0112330110111100-3323333212230003-1022223123313332-2130232122032102) |
| `routes.custom.severity.regex_match` | [routes.custom.severity.regex_match](resources--alert_policy--reference--group-001.md#canonical-3230132132111101-3012212020120123-0010131002023230-1132013120303012-1111322123211101-1213131103222030-2133322013331110-1331223000301311) |
| `routes.dont_send` | [routes.dont_send](resources--alert_policy--reference--group-001.md#canonical-0013000301100321-0011222023223213-2022201211223021-3221202312232300-2311121111122231-2111103321013012-3222123001012113-3303100311210312) |
| `routes.group` | [routes.group](resources--alert_policy--reference--group-001.md#canonical-0002332032210032-3332310302223100-0011332020220023-2131332031231031-3222011011013113-2021011311302333-3030000121220112-1210302102010320) |
| `routes.group.groups` | [routes.group.groups](resources--alert_policy--reference--group-001.md#canonical-1110112322112201-2222032020123113-2010321312333113-2031311012200310-0232332322222200-0120113010300200-3110323111320122-3020130200222311) |
| `routes.notification_parameters` | [routes.notification_parameters](resources--alert_policy--reference--group-001.md#canonical-3333002302120121-2031222300320321-3233232221223000-0022310320320300-0103212310032030-0203030302032132-1301332130131331-3310121331110103) |
| `routes.notification_parameters.custom` | [routes.notification_parameters.custom](resources--alert_policy--reference--group-001.md#canonical-0233120213322102-0003102023201102-1333202033231221-0131211230002101-3202013230321232-1100220200130023-0313120313330323-2011203130113302) |
| `routes.notification_parameters.custom.labels` | [routes.notification_parameters.custom.labels](resources--alert_policy--reference--group-001.md#canonical-1220121301322112-0302033000321132-1100321201312113-2120022311313110-1120212123012000-0022320310121111-1113300020301122-2321322220132313) |
| `routes.notification_parameters.default` | [routes.notification_parameters.default](resources--alert_policy--reference--group-001.md#canonical-1012002322102222-1021310100002322-0103212012032300-2033221123020001-0322100321200302-2220232113033230-0311322111203022-2321200111233202) |
| `routes.notification_parameters.group_interval` | [routes.notification_parameters.group_interval](resources--alert_policy--reference--group-001.md#canonical-3000103032212100-2000230200223211-3022330121211312-0031200332323301-2333310101330002-2302122202313111-3320110301023222-3330110211233123) |
| `routes.notification_parameters.group_wait` | [routes.notification_parameters.group_wait](resources--alert_policy--reference--group-001.md#canonical-1013233311300100-3113221230112323-2101212010121302-3203013220221202-2312022333122213-0332231132322001-3001021002312031-2232001221120101) |
| `routes.notification_parameters.individual` | [routes.notification_parameters.individual](resources--alert_policy--reference--group-001.md#canonical-1331001302112113-0310033031221300-0000200232232301-1101102331002122-3322322222230213-0322213110311021-0013001222022121-0320001110123110) |
| `routes.notification_parameters.repeat_interval` | [routes.notification_parameters.repeat_interval](resources--alert_policy--reference--group-001.md#canonical-3013203321120321-1220210303131202-1303321213113202-0132301031332113-3121102303201303-3112112130010212-0310032031313020-1121323103211212) |
| `routes.notification_parameters.ves_io_group` | [routes.notification_parameters.ves_io_group](resources--alert_policy--reference--group-001.md#canonical-2313033020011230-0112033221100303-3213302121333213-3213313123222101-2010020113231000-3001012023221330-0301213221333313-2101332032030300) |
| `routes.send` | [routes.send](resources--alert_policy--reference--group-001.md#canonical-0223313000203232-1002323001123120-3022021200313210-0023220102133233-2031323030103120-0333212113310030-1011303002220222-0002012022331020) |
| `routes.severity` | [routes.severity](resources--alert_policy--reference--group-001.md#canonical-0300112321332330-2132332211333020-1201031122110322-2211110220000301-2233211312230123-0320133130023123-3023122123231102-0031313211300301) |
| `routes.severity.severities` | [routes.severity.severities](resources--alert_policy--reference--group-001.md#canonical-3033002121033331-2000032210312310-1220021313000111-0311120001002312-2212220233302210-3010313202110302-1301102022331002-3112130030333021) |
| `timeouts` | [timeouts](resources--alert_policy--reference--group-001.md#canonical-3110210011201330-1023303000322123-2233012122331302-3011120021201100-3201233010001130-0222330320023110-2120313123203202-3103112230112111) |
| `timeouts.create` | [timeouts.create](resources--alert_policy--reference--group-001.md#canonical-3312203233110313-1113131310030333-2232022021001323-2211322003001322-2202311130121333-2121021210200130-0131133300113011-0001213323210013) |
| `timeouts.delete` | [timeouts.delete](resources--alert_policy--reference--group-001.md#canonical-0303030220321111-0033323211310102-3230333033011120-3001320021313302-3001302033211110-2121232332013203-3320103112313303-3120320203120200) |
| `timeouts.read` | [timeouts.read](resources--alert_policy--reference--group-001.md#canonical-1202323023332032-3331321301333110-1200011131311211-1233323132202202-1111210321223000-2202221302032302-1230200300212032-3002010113031030) |
| `timeouts.update` | [timeouts.update](resources--alert_policy--reference--group-001.md#canonical-3002033133200122-1031220210030021-1100000010230101-2033112322211020-2031223231010323-2322003100303102-2013123320111022-3122101011011013) |

<a id="canonical-2302223211220010-3133333223123323-0320311202331210-3230201211033133-0220001322010300-2323033221331022-2321110220123020-0201210012212021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `notification_parameters` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- notification_parameters

<a id="canonical-3033201321301122-3130300012120220-0031332223003032-3012033312301023-3013321112233003-2222300321130002-0113001322222222-2033113223220323"></a>

Type: `"object"`. single nested block, Optional.

Set of notification parameters to decide how and when the alert notifications should be sent to the
receivers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom",
    "default"),
  validators.ConflictingObjectAttributes("custom",
    "individual"),
  validators.ConflictingObjectAttributes("custom",
    "ves_io_group"),
  validators.ConflictingObjectAttributes("default",
    "individual"),
  validators.ConflictingObjectAttributes("default",
    "ves_io_group"),
  validators.ConflictingObjectAttributes("individual",
    "ves_io_group")}
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
  "x-ves-oneof-field-group_by": "[\"custom\",\"default\",\"individual\",\"ves_io_group\"]"
}
```

Terraform syntax:

```terraform
notification_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313300323021130-3010032011011133-0023011102103100-2232000012103202-0110202233230220-0230202212122312-0003332220300112-3301210213132203"></a>

### Direct properties for `notification_parameters`

- [custom](resources--alert_policy--reference--group-001.md#canonical-3113333100001232-1201212313303031-0121102133332311-2231102000302312-2033011323111100-3332002323232120-1021202100212120-3331103121103022): complete subsection reference.

- [default](resources--alert_policy--reference--group-001.md#canonical-2123120301301120-3113010213203313-2220210202320102-0302201203112102-1333203130012323-3112010320233112-0102122320020203-3003231132010032): complete subsection reference.

<a id="canonical-3100302232300130-1122312101210211-1212223312213111-0211303103221131-1202313302010311-2130323120233321-2301110233022121-2000221112031110"></a>

<a id="canonical-0011133021330112-0222121010010302-2020003233130011-0011022013212023-0333123100112021-2222223111021121-0213332101020210-1322301012023202"></a>

#### `notification_parameters.group_interval` property

Type: `"string"`. Optional.

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval defaults to "1m"

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-2013001323222332-2213113320120010-1120032323133023-1110212302122303-2322000210320111-1030131111313011-0001233022223002-3112132322012231"></a>

<a id="canonical-3023303311301023-0233120202200311-3110321210320210-1202313120203313-0313021101310212-2313322030312101-3102021010021120-1103211200103003"></a>

#### `notification_parameters.group_wait` property

Type: `"string"`. Optional.

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to '30s'.

Additional upstream details:

Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_wait defaults to "30s"

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [individual](resources--alert_policy--reference--group-001.md#canonical-2211120302112020-2121111212221001-0323321313030302-2110113300311111-1103130100012331-2210312223322113-2010232200332201-3131330131202133): complete subsection reference.

<a id="canonical-2112212112212311-3232132201223023-2303320011313033-3323020003320110-2002210311320000-1300230213120313-0320313210303203-2113213000000113"></a>

<a id="canonical-2333100221100313-0202231302011322-2232113211133033-2220201121133002-2023130010332122-0112123133113332-1320221011021130-3201311210132013"></a>

#### `notification_parameters.repeat_interval` property

Type: `"string"`. Optional.

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to '4h'.

Additional upstream details:

Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval defaults to "4h"

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [ves_io_group](resources--alert_policy--reference--group-001.md#canonical-2110212131113322-0322002230203220-2301312221121132-2300121212231030-1121333202131201-1120031310210111-1333012212113113-2133032123031100): complete subsection reference.

<a id="canonical-3113333100001232-1201212313303031-0121102133332311-2231102000302312-2033011323111100-3332002323232120-1021202100212120-3331103121103022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `notification_parameters.custom` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-2302223211220010-3133333223123323-0320311202331210-3230201211033133-0220001322010300-2323033221331022-2321110220123020-0201210012212021)
- notification_parameters.custom

<a id="canonical-3212310323101102-3011211212213201-0022030231102102-0323321030132203-0301122030302121-2200111031312102-2213112032223011-0013133311123102"></a>

Type: `"object"`. single nested block, Optional.

Specify list of custom labels to group/aggregate the alerts.

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
custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211013330303101-2133000221220010-3030002033310233-2330001021231310-0012000303203003-2201110233312321-1003333131233312-3331221101121002"></a>

### Direct properties for `notification_parameters.custom`

<a id="canonical-3213003110221023-2112020210032302-3103202103022020-3032111220213130-0100103021312312-1310231311120101-3221323221022102-3300331010120222"></a>

#### `notification_parameters.custom.labels` property

Type: `["list", "string"]`. Optional.

Name of labels to group/aggregate the alerts.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2123120301301120-3113010213203313-2220210202320102-0302201203112102-1333203130012323-3112010320233112-0102122320020203-3003231132010032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `notification_parameters.default` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-2302223211220010-3133333223123323-0320311202331210-3230201211033133-0220001322010300-2323033221331022-2321110220123020-0201210012212021)
- notification_parameters.default

<a id="canonical-3322030211113131-1320130211022033-0122133112203220-1310122220131031-2032211333323110-1332112131132232-2311212101320121-2210200132313022"></a>

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
default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211120302112020-2121111212221001-0323321313030302-2110113300311111-1103130100012331-2210312223322113-2010232200332201-3131330131202133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `notification_parameters.individual` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-2302223211220010-3133333223123323-0320311202331210-3230201211033133-0220001322010300-2323033221331022-2321110220123020-0201210012212021)
- notification_parameters.individual

<a id="canonical-3032110111022333-1023320221223100-2322232212110210-2122120323231231-0022132301302231-0031223333131111-2010323320302231-1022023110320203"></a>

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
individual = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110212131113322-0322002230203220-2301312221121132-2300121212231030-1121333202131201-1120031310210111-1333012212113113-2133032123031100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `notification_parameters.ves_io_group` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-2302223211220010-3133333223123323-0320311202331210-3230201211033133-0220001322010300-2323033221331022-2321110220123020-0201210012212021)
- notification_parameters.ves_io_group

<a id="canonical-3003231023000303-3101012031310332-0331131332322122-0133220311200033-0110323322112221-1312023103212330-2231210001313032-1112030322023222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ves io group.

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
ves_io_group = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131322330100230-2131122213031210-1231122112211222-3020300103001122-2130202110300030-0001002122021021-2210020203303312-2111210332131202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `receivers` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- receivers

<a id="canonical-3022110002301212-1313313133003301-1020133033012100-3133320230233311-1231023311020322-0003303012321123-0222110301003112-2233210232323232"></a>

Type: `"object"`. list nested block, Optional.

List of Alert Receivers where the alerts will be sent.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
receivers {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312313010031023-2321333103021122-2321323030031320-0030031221000303-3332002310000111-0221012203013221-1130312130213020-2033103030001113"></a>

### Direct properties for `receivers`

<a id="canonical-1011123210332213-2232103011230200-0200120323001101-1013103001002303-2122013020201133-1312033012203333-2022201123211031-1310101211220110"></a>

#### `receivers.kind` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1332030101103131-1203120111012330-3120213112210210-3303012233033021-0233001322331320-1002030131200000-1203132031110020-2033100321103321"></a>

<a id="canonical-1130322032021020-3110122111333330-2303132200322030-3220213320212032-1013010302231100-2301300031123032-2000222011310003-3022302021221201"></a>

#### `receivers.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1203331031330301-2322102131223212-2322203001201230-2230102333033023-0232212120221123-1031230202212113-0310223122300030-2200021123020231"></a>

<a id="canonical-3030112333230321-1231310022021131-3303022123210001-3102322213323211-2020031320303322-3033200222311222-2033330311303213-0231232020003200"></a>

#### `receivers.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1021332013020120-0002201332022030-2201310303002012-1031300102110003-3010030201220213-3110212023220312-2210110123122100-1323002323122132"></a>

<a id="canonical-0000111030001112-3032301112201201-0331322320302103-3032333100222220-2303031322123132-3000021221323300-3231231023213101-3031111002033222"></a>

#### `receivers.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1303321010332121-2132221021302011-3233001313023022-0201233322113022-3212213313313202-0130320100212312-2211010003032313-0001023333130210"></a>

<a id="canonical-2011310210113020-3112312112011200-0312132333201212-0012012110112133-2312000120311100-0012313203030303-2201030231212121-0131301103311303"></a>

#### `receivers.uid` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- routes

<a id="canonical-0301321223132032-0312112021303120-1330003130200103-1103103033301201-3310110301100033-3332202222133303-2302002211113211-1011000100332200"></a>

Type: `"object"`. list nested block, Optional.

Set of routes to match the incoming alert. The routes are evaluated in the specified order and
terminates on the first match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("alertname",
    "alertname_regex"),
  validators.ConflictingListObjectAttributes("alertname",
    "any"),
  validators.ConflictingListObjectAttributes("alertname",
    "custom"),
  validators.ConflictingListObjectAttributes("alertname",
    "group"),
  validators.ConflictingListObjectAttributes("alertname",
    "severity"),
  validators.ConflictingListObjectAttributes("alertname_regex",
    "any"),
  validators.ConflictingListObjectAttributes("alertname_regex",
    "custom"),
  validators.ConflictingListObjectAttributes("alertname_regex",
    "group"),
  validators.ConflictingListObjectAttributes("alertname_regex",
    "severity"),
  validators.ConflictingListObjectAttributes("any",
    "custom"),
  validators.ConflictingListObjectAttributes("any",
    "group"),
  validators.ConflictingListObjectAttributes("any",
    "severity"),
  validators.ConflictingListObjectAttributes("custom",
    "group"),
  validators.ConflictingListObjectAttributes("custom",
    "severity"),
  validators.ConflictingListObjectAttributes("dont_send",
    "send"),
  validators.ConflictingListObjectAttributes("group",
    "severity")}
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1101202030312021-1333203333300201-2213113230023320-1102220330222230-1010322333102330-1131320110113101-3313001311102132-3330032031302302"></a>

### Direct properties for `routes`

<a id="canonical-0320031112311211-3032022332222133-3112112223101303-2203113213223122-3211022301202333-1132122001111330-1330131020031112-1322230232322302"></a>

#### `routes.alertname` property

Type: `"string"`. Optional.

\[Enum:
SITE\_CUSTOMER\_TUNNEL\_INTERFACE\_DOWN|SITE\_PHYSICAL\_INTERFACE\_DOWN|TUNNELS\_TO\_CUSTOMER\_SITE\_DOWN|SERVICE\_SERVER\_ERROR|SERVICE\_CLIENT\_ERROR|SERVICE\_HEALTH\_LOW|SERVICE\_UNAVAILABLE|SERVICE\_SERVER\_ERROR\_PER\_SOURCE\_SITE|SERVICE\_CLIENT\_ERROR\_PER\_SOURCE\_SITE|SERVICE\_ENDPOINT\_HEALTHCHECK\_FAILURE|SYNTHETIC\_MONITOR\_HEALTH\_CRITICAL|MALICIOUS\_USER\_DETECTED|WAF\_TOO\_MANY\_ATTACKS|API\_SECURITY\_TOO\_MANY\_ATTACKS|SERVICE\_POLICY\_TOO\_MANY\_ATTACKS|WAF\_TOO\_MANY\_MALICIOUS\_BOTS|BOT\_DEFENSE\_TOO\_MANY\_SECURITY\_EVENTS|THREAT\_CAMPAIGN|VES\_CLIENT\_SIDE\_DEFENSE\_SUSPICIOUS\_DOMAIN|VES\_CLIENT\_SIDE\_DEFENSE\_SENSITIVE\_FIELD\_READ|TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_FAILURE|TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_STILL\_FAILING|TLS\_AUTOMATIC\_CERTIFICATE\_EXPIRED|TLS\_CUSTOM\_CERTIFICATE\_EXPIRING|TLS\_CUSTOM\_CERTIFICATE\_EXPIRING\_SOON|TLS\_CUSTOM\_CERTIFICATE\_EXPIRED|L7DDOS|DNS\_ZONE\_IGNORED\_DUPLICATE\_RECORD|API\_SECURITY\_UNUSED\_API\_DETECTED|API\_SECURITY\_SHADOW\_API\_DETECTED|API\_SECURITY\_SENSITIVE\_DATA\_IN\_RESPONSE\_DETECTED|API\_SECURITY\_RISK\_SCORE\_HIGH\_DETECTED|ROUTED\_DDOS\_ALERT\_NOTIFICATION|ROUTED\_DDOS\_MITIGATION\_NOTIFICATION|ROUTED\_DDOS\_TUNNEL\_STATUS\_UPDATE\_NOTIFICATION|L7\_DDOS\_AUTO\_MITIGATION\]
List of Alert Names Customer tunnel interface down Physical Interface down Tunnel Interfaces to
Customer Site Down Virtual Host server error Virtual Host client error Service Health Low Service
Unavailable Virtual Host server error Virtual Host client error Endpoint Healthcheck failure
Synthetic.. Possible values are \`SITE\_CUSTOMER\_TUNNEL\_INTERFACE\_DOWN\`,
\`SITE\_PHYSICAL\_INTERFACE\_DOWN\`, \`TUNNELS\_TO\_CUSTOMER\_SITE\_DOWN\`,
\`SERVICE\_SERVER\_ERROR\`, \`SERVICE\_CLIENT\_ERROR\`, \`SERVICE\_HEALTH\_LOW\`,
\`SERVICE\_UNAVAILABLE\`, \`SERVICE\_SERVER\_ERROR\_PER\_SOURCE\_SITE\`,
\`SERVICE\_CLIENT\_ERROR\_PER\_SOURCE\_SITE\`, \`SERVICE\_ENDPOINT\_HEALTHCHECK\_FAILURE\`,
\`SYNTHETIC\_MONITOR\_HEALTH\_CRITICAL\`, \`MALICIOUS\_USER\_DETECTED\`,
\`WAF\_TOO\_MANY\_ATTACKS\`, \`API\_SECURITY\_TOO\_MANY\_ATTACKS\`,
\`SERVICE\_POLICY\_TOO\_MANY\_ATTACKS\`, \`WAF\_TOO\_MANY\_MALICIOUS\_BOTS\`,
\`BOT\_DEFENSE\_TOO\_MANY\_SECURITY\_EVENTS\`, \`THREAT\_CAMPAIGN\`,
\`VES\_CLIENT\_SIDE\_DEFENSE\_SUSPICIOUS\_DOMAIN\`,
\`VES\_CLIENT\_SIDE\_DEFENSE\_SENSITIVE\_FIELD\_READ\`,
\`TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_FAILURE\`,
\`TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_STILL\_FAILING\`, \`TLS\_AUTOMATIC\_CERTIFICATE\_EXPIRED\`,
\`TLS\_CUSTOM\_CERTIFICATE\_EXPIRING\`, \`TLS\_CUSTOM\_CERTIFICATE\_EXPIRING\_SOON\`,
\`TLS\_CUSTOM\_CERTIFICATE\_EXPIRED\`, \`L7DDOS\`, \`DNS\_ZONE\_IGNORED\_DUPLICATE\_RECORD\`,
\`API\_SECURITY\_UNUSED\_API\_DETECTED\`, \`API\_SECURITY\_SHADOW\_API\_DETECTED\`,
\`API\_SECURITY\_SENSITIVE\_DATA\_IN\_RESPONSE\_DETECTED\`,
\`API\_SECURITY\_RISK\_SCORE\_HIGH\_DETECTED\`, \`ROUTED\_DDOS\_ALERT\_NOTIFICATION\`,
\`ROUTED\_DDOS\_MITIGATION\_NOTIFICATION\`, \`ROUTED\_DDOS\_TUNNEL\_STATUS\_UPDATE\_NOTIFICATION\`,
\`L7\_DDOS\_AUTO\_MITIGATION\`. Defaults to \`SITE\_CUSTOMER\_TUNNEL\_INTERFACE\_DOWN\`.

Additional upstream details:

List of Alert Names

Customer tunnel interface down Physical Interface down Tunnel Interfaces to Customer Site Down
Virtual Host server error Virtual Host client error Service Health Low Service Unavailable Virtual
Host server error Virtual Host client error Endpoint Healthcheck failure Synthetic monitor health
critical Malicious user detected Virtual Host WAF security events detected Virtual Host API security
events detected Virtual Host Service Policy security events detected Virtual Host Many Malicious
Bots based WAF security events detected Virtual Host Many Malicious Bots based Bot Defense security
events detected Virtual Host Many Threat campaign based WAF security events detected Suspicious
domain identified by Client-Side Defense service Client-Side Defense has identified a suspicious
script that is reading sensitive form field TLS Automatic Certificate renewal is failing TLS
Automatic Certificate renewal is still failing after multiple retries TLS Automatic Certificate has
expired TLS Custom Certificate will expire in less than 28 days TLS Custom Certificate will expire
in less than 15 days TLS Custom Certificate has expired DDoS security event detected DNS Zone
Ignored a Duplicate Record Create Request Unused APIs Detected Shadow APIs Detected Endpoints With
Sensitive Data In Response Detected High Risk Score Endpoints Detected A routed DDoS traffic anomaly
has been detected A routed DDoS mitigation has been implemented to block malicious traffic A routed
DDoS tunnel status has been changed L7 DDoS attack was detected, automatic mitigation is taking
place.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["API_SECURITY_RISK_SCORE_HIGH_DETECTED","API_SECURITY_SENSITIVE_DATA_IN_RESPONSE_DETECTED","API_SECURITY_SHADOW_API_DETECTED","API_SECURITY_TOO_MANY_ATTACKS","API_SECURITY_UNUSED_API_DETECTED","BOT_DEFENSE_TOO_MANY_SECURITY_EVENTS","DNS_ZONE_IGNORED_DUPLICATE_RECORD","L7DDOS","L7_DDOS_AUTO_MITIGATION","MALICIOUS_USER_DETECTED","ROUTED_DDOS_ALERT_NOTIFICATION","ROUTED_DDOS_MITIGATION_NOTIFICATION","ROUTED_DDOS_TUNNEL_STATUS_UPDATE_NOTIFICATION","SERVICE_CLIENT_ERROR","SERVICE_CLIENT_ERROR_PER_SOURCE_SITE","SERVICE_ENDPOINT_HEALTHCHECK_FAILURE","SERVICE_HEALTH_LOW","SERVICE_POLICY_TOO_MANY_ATTACKS","SERVICE_SERVER_ERROR","SERVICE_SERVER_ERROR_PER_SOURCE_SITE","SERVICE_UNAVAILABLE","SITE_CUSTOMER_TUNNEL_INTERFACE_DOWN","SITE_PHYSICAL_INTERFACE_DOWN","SYNTHETIC_MONITOR_HEALTH_CRITICAL","THREAT_CAMPAIGN","TLS_AUTOMATIC_CERTIFICATE_EXPIRED","TLS_AUTOMATIC_CERTIFICATE_RENEWAL_FAILURE","TLS_AUTOMATIC_CERTIFICATE_RENEWAL_STILL_FAILING","TLS_CUSTOM_CERTIFICATE_EXPIRED","TLS_CUSTOM_CERTIFICATE_EXPIRING","TLS_CUSTOM_CERTIFICATE_EXPIRING_SOON","TUNNELS_TO_CUSTOMER_SITE_DOWN","VES_CLIENT_SIDE_DEFENSE_SENSITIVE_FIELD_READ","VES_CLIENT_SIDE_DEFENSE_SUSPICIOUS_DOMAIN","WAF_TOO_MANY_ATTACKS","WAF_TOO_MANY_MALICIOUS_BOTS"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SITE_CUSTOMER_TUNNEL_INTERFACE_DOWN",
    "SITE_PHYSICAL_INTERFACE_DOWN",
    "TUNNELS_TO_CUSTOMER_SITE_DOWN",
    "SERVICE_SERVER_ERROR",
    "SERVICE_CLIENT_ERROR",
    "SERVICE_HEALTH_LOW",
    "SERVICE_UNAVAILABLE",
    "SERVICE_SERVER_ERROR_PER_SOURCE_SITE",
    "SERVICE_CLIENT_ERROR_PER_SOURCE_SITE",
    "SERVICE_ENDPOINT_HEALTHCHECK_FAILURE",
    "SYNTHETIC_MONITOR_HEALTH_CRITICAL",
    "MALICIOUS_USER_DETECTED",
    "WAF_TOO_MANY_ATTACKS",
    "API_SECURITY_TOO_MANY_ATTACKS",
    "SERVICE_POLICY_TOO_MANY_ATTACKS",
    "WAF_TOO_MANY_MALICIOUS_BOTS",
    "BOT_DEFENSE_TOO_MANY_SECURITY_EVENTS",
    "THREAT_CAMPAIGN",
    "VES_CLIENT_SIDE_DEFENSE_SUSPICIOUS_DOMAIN",
    "VES_CLIENT_SIDE_DEFENSE_SENSITIVE_FIELD_READ",
    "TLS_AUTOMATIC_CERTIFICATE_RENEWAL_FAILURE",
    "TLS_AUTOMATIC_CERTIFICATE_RENEWAL_STILL_FAILING",
    "TLS_AUTOMATIC_CERTIFICATE_EXPIRED",
    "TLS_CUSTOM_CERTIFICATE_EXPIRING",
    "TLS_CUSTOM_CERTIFICATE_EXPIRING_SOON",
    "TLS_CUSTOM_CERTIFICATE_EXPIRED",
    "L7DDOS",
    "DNS_ZONE_IGNORED_DUPLICATE_RECORD",
    "API_SECURITY_UNUSED_API_DETECTED",
    "API_SECURITY_SHADOW_API_DETECTED",
    "API_SECURITY_SENSITIVE_DATA_IN_RESPONSE_DETECTED",
    "API_SECURITY_RISK_SCORE_HIGH_DETECTED",
    "ROUTED_DDOS_ALERT_NOTIFICATION",
    "ROUTED_DDOS_MITIGATION_NOTIFICATION",
    "ROUTED_DDOS_TUNNEL_STATUS_UPDATE_NOTIFICATION",
    "L7_DDOS_AUTO_MITIGATION"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_CUSTOMER_TUNNEL_INTERFACE_DOWN",
  "enum": [
    "SITE_CUSTOMER_TUNNEL_INTERFACE_DOWN",
    "SITE_PHYSICAL_INTERFACE_DOWN",
    "TUNNELS_TO_CUSTOMER_SITE_DOWN",
    "SERVICE_SERVER_ERROR",
    "SERVICE_CLIENT_ERROR",
    "SERVICE_HEALTH_LOW",
    "SERVICE_UNAVAILABLE",
    "SERVICE_SERVER_ERROR_PER_SOURCE_SITE",
    "SERVICE_CLIENT_ERROR_PER_SOURCE_SITE",
    "SERVICE_ENDPOINT_HEALTHCHECK_FAILURE",
    "SYNTHETIC_MONITOR_HEALTH_CRITICAL",
    "MALICIOUS_USER_DETECTED",
    "WAF_TOO_MANY_ATTACKS",
    "API_SECURITY_TOO_MANY_ATTACKS",
    "SERVICE_POLICY_TOO_MANY_ATTACKS",
    "WAF_TOO_MANY_MALICIOUS_BOTS",
    "BOT_DEFENSE_TOO_MANY_SECURITY_EVENTS",
    "THREAT_CAMPAIGN",
    "VES_CLIENT_SIDE_DEFENSE_SUSPICIOUS_DOMAIN",
    "VES_CLIENT_SIDE_DEFENSE_SENSITIVE_FIELD_READ",
    "TLS_AUTOMATIC_CERTIFICATE_RENEWAL_FAILURE",
    "TLS_AUTOMATIC_CERTIFICATE_RENEWAL_STILL_FAILING",
    "TLS_AUTOMATIC_CERTIFICATE_EXPIRED",
    "TLS_CUSTOM_CERTIFICATE_EXPIRING",
    "TLS_CUSTOM_CERTIFICATE_EXPIRING_SOON",
    "TLS_CUSTOM_CERTIFICATE_EXPIRED",
    "L7DDOS",
    "DNS_ZONE_IGNORED_DUPLICATE_RECORD",
    "API_SECURITY_UNUSED_API_DETECTED",
    "API_SECURITY_SHADOW_API_DETECTED",
    "API_SECURITY_SENSITIVE_DATA_IN_RESPONSE_DETECTED",
    "API_SECURITY_RISK_SCORE_HIGH_DETECTED",
    "ROUTED_DDOS_ALERT_NOTIFICATION",
    "ROUTED_DDOS_MITIGATION_NOTIFICATION",
    "ROUTED_DDOS_TUNNEL_STATUS_UPDATE_NOTIFICATION",
    "L7_DDOS_AUTO_MITIGATION"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1102100033130312-3133031111301310-3230213122120311-3233123203000111-0331010001103202-2020113232301132-3131013122313120-0012313011010303"></a>

<a id="canonical-2102200032322022-3321033022211331-2302121320031022-1223112133310123-1112323320322301-0103311010221010-0023110210222012-2211022023122212"></a>

#### `routes.alertname_regex` property

Type: `"string"`. Optional.

Exclusive with \[alertname any custom group severity\] Regular Expression match for the alertname.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [any](resources--alert_policy--reference--group-001.md#canonical-1132320103011322-1331023221321120-3300230121112023-2131300020333333-0230023203002101-3132021113323113-2113330320002322-3133220322301333): complete subsection reference.

- [custom](resources--alert_policy--reference--group-001.md#canonical-3230201200300110-3010211222100103-3201230231102031-1213011103213132-0310132223101002-0010233123332210-2030021202212202-3031020303123101): complete subsection reference.

- [dont_send](resources--alert_policy--reference--group-001.md#canonical-2033132130120130-2013330110222211-2312223032330102-0121311331303313-2000211203133311-3311213221211212-0311012112221233-1331210210232012): complete subsection reference.

- [group](resources--alert_policy--reference--group-001.md#canonical-3121013101212101-0131111102203233-0331001313303122-3302030133103131-2133310232203033-0331232102021310-3300003021111333-2222330132002220): complete subsection reference.

- [notification_parameters](resources--alert_policy--reference--group-001.md#canonical-1132102222000120-1312023210212021-2212103031202003-2303010133130331-2001130111201011-0311320332230320-2230201020303103-2220332311321331): complete subsection reference.

- [send](resources--alert_policy--reference--group-001.md#canonical-2110310313011121-1112331112232321-2332311130112102-0101233330303222-2032322323223301-3100321320112112-0230232022322233-1202023210320223): complete subsection reference.

- [severity](resources--alert_policy--reference--group-001.md#canonical-2033001201101022-2221131331322322-1113222133113330-3211223030232202-3033130223101322-3133000201133203-3000131300220330-1023330031001313): complete subsection reference.

<a id="canonical-1132320103011322-1331023221321120-3300230121112023-2131300020333333-0230023203002101-3132021113323113-2113330320002322-3133220322301333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.any` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [routes](resources--alert_policy--reference--group-001.md#canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033)
- routes.any

<a id="canonical-2002001320200300-3331123220030113-0202020323323303-3210003013230023-3222002131211112-3332332122112320-0201021211200101-3010201201103321"></a>

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
any = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230201200300110-3010211222100103-3201230231102031-1213011103213132-0310132223101002-0010233123332210-2030021202212202-3031020303123101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [routes](resources--alert_policy--reference--group-001.md#canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033)
- routes.custom

<a id="canonical-0121300113230032-3310201302322311-0121210012133301-1212220132331122-3202020021033300-2301011200023313-3230300323121130-2032121100310003"></a>

Type: `"object"`. single nested block, Optional.

A set of matchers an alert has to fulfill to match the route.

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
custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132301023223003-2220131310301130-1201122201012320-0133111310311223-3031333212303211-0201131203110201-3333113233210010-1130001133323312"></a>

### Direct properties for `routes.custom`

- [alertlabel](resources--alert_policy--reference--group-001.md#canonical-3213201222132211-0321212210121103-3111320003000323-2223331023233323-1201003002012201-2103021113203132-3312110213031123-0203211132311101): complete subsection reference.

- [alertname](resources--alert_policy--reference--group-001.md#canonical-2020012332312022-1230112312111312-3111233110312222-2330002312011311-1201320303221211-1321231121320122-1301033130033020-3233100021113303): complete subsection reference.

- [group](resources--alert_policy--reference--group-001.md#canonical-2011010113021203-2333000210101230-3032311201110303-1013211322323222-1113220300312112-0201232013320033-3012310221102211-0031001203003130): complete subsection reference.

- [severity](resources--alert_policy--reference--group-001.md#canonical-3010333121102303-1101131221000001-0131322123023133-0130331133201303-2021210233033132-3200213201213312-0002312032003303-1232113323213003): complete subsection reference.

<a id="canonical-3213201222132211-0321212210121103-3111320003000323-2223331023233323-1201003002012201-2103021113203132-3312110213031123-0203211132311101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom.alertlabel` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [routes](resources--alert_policy--reference--group-001.md#canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033)
- [routes.custom](resources--alert_policy--reference--group-001.md#canonical-3230201200300110-3010211222100103-3201230231102031-1213011103213132-0310132223101002-0010233123332210-2030021202212202-3031020303123101)
- routes.custom.alertlabel

<a id="canonical-3031023300000023-0311031110020223-0033332130333333-2111333210123030-3313021123333220-2330011223112223-3113321131132032-1100110231112003"></a>

Type: `"object"`. single nested block, Optional.

AlertLabel to configure the alert policy rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 3
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.keys.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
      "ves.io.schema.rules.map.max_pairs": "3"
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
    "ves.io.schema.rules.map.keys.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.map.max_pairs": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.keys.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.map.max_pairs": "3"
  }
}
```

Terraform syntax:

```terraform
alertlabel {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020012332312022-1230112312111312-3111233110312222-2330002312011311-1201320303221211-1321231121320122-1301033130033020-3233100021113303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom.alertname` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [routes](resources--alert_policy--reference--group-001.md#canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033)
- [routes.custom](resources--alert_policy--reference--group-001.md#canonical-3230201200300110-3010211222100103-3201230231102031-1213011103213132-0310132223101002-0010233123332210-2030021202212202-3031020303123101)
- routes.custom.alertname

<a id="canonical-2333101113300213-1203031230130213-2312213020023020-2223320312322321-0000322323012310-2010021111113131-2311210230301232-1233031302021212"></a>

Type: `"object"`. single nested block, Optional.

Label Matcher.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_match",
    "regex_match")}
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
  "x-ves-oneof-field-matcher_type": "[\"exact_match\",\"regex_match\"]"
}
```

Terraform syntax:

```terraform
alertname {
  # Configure direct properties listed below.
}
```

<a id="canonical-3012121023320332-1033000032003001-3032111022110203-3211102001002013-3200202001032333-3022300013102102-2200230231320021-3101003132320010"></a>

### Direct properties for `routes.custom.alertname`

<a id="canonical-2033112023012032-3002002323231011-0103130120301301-3213001323111113-0102230233012233-2003022122101230-2321221222313102-3011300231021200"></a>

#### `routes.custom.alertname.exact_match` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_match\] Equality match value for the label.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1100223213310010-0022212331312101-0003320003032112-0111232233011123-1103213321013021-3110000111133123-2233003100023222-0322210222321030"></a>

<a id="canonical-1100122110011321-2222230130322002-0333112132332302-0021233202332303-0121102110010032-3123333321110222-0012210221102012-2122313112321010"></a>

#### `routes.custom.alertname.regex_match` property

Type: `"string"`. Optional.

Exclusive with \[exact\_match\] Regular expression match value for the label.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2011010113021203-2333000210101230-3032311201110303-1013211322323222-1113220300312112-0201232013320033-3012310221102211-0031001203003130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom.group` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [routes](resources--alert_policy--reference--group-001.md#canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033)
- [routes.custom](resources--alert_policy--reference--group-001.md#canonical-3230201200300110-3010211222100103-3201230231102031-1213011103213132-0310132223101002-0010233123332210-2030021202212202-3031020303123101)
- routes.custom.group

<a id="canonical-1311123122031003-2131030113300323-2022131331331213-1302233302300000-1312113131110003-0320302222122203-3230233311113323-0320221230231020"></a>

Type: `"object"`. single nested block, Optional.

Label Matcher.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_match",
    "regex_match")}
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
  "x-ves-oneof-field-matcher_type": "[\"exact_match\",\"regex_match\"]"
}
```

Terraform syntax:

```terraform
group {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213220300011031-1230001033132330-1122311322310231-2023110322232001-1132023111230313-1010010220031130-1311102200111030-3223203110110103"></a>

### Direct properties for `routes.custom.group`

<a id="canonical-3230131122130112-3212031220331022-1233233122303303-3203131031222233-2112211232201031-1203013333001113-2112301102001300-2300113220203332"></a>

#### `routes.custom.group.exact_match` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_match\] Equality match value for the label.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1103310132120201-2112133133220303-3113130230012110-0032330131012333-3130010222102012-1033220303332231-3333323132323310-3112120201131102"></a>

<a id="canonical-1001303231101303-0122233231311223-1101100333321203-0020201121320111-1002032321211300-2010232301130112-3321201222223221-0013113033312100"></a>

#### `routes.custom.group.regex_match` property

Type: `"string"`. Optional.

Exclusive with \[exact\_match\] Regular expression match value for the label.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3010333121102303-1101131221000001-0131322123023133-0130331133201303-2021210233033132-3200213201213312-0002312032003303-1232113323213003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.custom.severity` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [routes](resources--alert_policy--reference--group-001.md#canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033)
- [routes.custom](resources--alert_policy--reference--group-001.md#canonical-3230201200300110-3010211222100103-3201230231102031-1213011103213132-0310132223101002-0010233123332210-2030021202212202-3031020303123101)
- routes.custom.severity

<a id="canonical-3022233222123223-2232020231131132-3231121130111201-0311120200200320-2300130131323101-1220122133133230-2332123132223222-1021103203021122"></a>

Type: `"object"`. single nested block, Optional.

Label Matcher.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_match",
    "regex_match")}
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
  "x-ves-oneof-field-matcher_type": "[\"exact_match\",\"regex_match\"]"
}
```

Terraform syntax:

```terraform
severity {
  # Configure direct properties listed below.
}
```

<a id="canonical-2023211303011212-3230230032210220-2002310331031232-0321123312211023-1022032100120311-2200013002032133-1111210122300230-2210011230320231"></a>

### Direct properties for `routes.custom.severity`

<a id="canonical-2100131220212303-1133201031112120-3001030132011321-1310313031303211-0112330110111100-3323333212230003-1022223123313332-2130232122032102"></a>

#### `routes.custom.severity.exact_match` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_match\] Equality match value for the label.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3230132132111101-3012212020120123-0010131002023230-1132013120303012-1111322123211101-1213131103222030-2133322013331110-1331223000301311"></a>

<a id="canonical-1231121121013332-2021122320320120-2311222212323311-3210021333111311-2202233033023222-0200323120221120-2232323222232232-1102323232110310"></a>

#### `routes.custom.severity.regex_match` property

Type: `"string"`. Optional.

Exclusive with \[exact\_match\] Regular expression match value for the label.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2033132130120130-2013330110222211-2312223032330102-0121311331303313-2000211203133311-3311213221211212-0311012112221233-1331210210232012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.dont_send` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [routes](resources--alert_policy--reference--group-001.md#canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033)
- routes.dont_send

<a id="canonical-0013000301100321-0011222023223213-2022201211223021-3221202312232300-2311121111122231-2111103321013012-3222123001012113-3303100311210312"></a>

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
dont_send = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121013101212101-0131111102203233-0331001313303122-3302030133103131-2133310232203033-0331232102021310-3300003021111333-2222330132002220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.group` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [routes](resources--alert_policy--reference--group-001.md#canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033)
- routes.group

<a id="canonical-0002332032210032-3332310302223100-0011332020220023-2131332031231031-3222011011013113-2021011311302333-3030000121220112-1210302102010320"></a>

Type: `"object"`. single nested block, Optional.

Select one or more known group names to match the incoming alert.

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
group {
  # Configure direct properties listed below.
}
```

<a id="canonical-2030122122133321-2223132002200001-1303100333313322-2022221133012022-0123231122123322-0010010111130002-1022032112033003-2311012003331220"></a>

### Direct properties for `routes.group`

<a id="canonical-1110112322112201-2222032020123113-2010321312333113-2031311012200310-0232332322222200-0120113010300200-3110323111320122-3020130200222311"></a>

#### `routes.group.groups` property

Type: `["list", "string"]`. Optional.

\[Enum:
INFRASTRUCTURE|IAAS\_CAAS|VIRTUAL\_HOST|VOLT\_SHARE|UAM|SECURITY|TIMESERIES\_ANOMALY|SHAPE\_SECURITY|SECURITY\_CSD|CDN|SYNTHETIC\_MONITORS|TLS|SECURITY\_BOT\_DEFENSE|CLOUD\_LINK|DNS|ROUTED\_DDOS\]
Groups. Name of groups to match the alert. Possible values are \`INFRASTRUCTURE\`, \`IAAS\_CAAS\`,
\`VIRTUAL\_HOST\`, \`VOLT\_SHARE\`, \`UAM\`, \`SECURITY\`, \`TIMESERIES\_ANOMALY\`,
\`SHAPE\_SECURITY\`, \`SECURITY\_CSD\`, \`CDN\`, \`SYNTHETIC\_MONITORS\`, \`TLS\`,
\`SECURITY\_BOT\_DEFENSE\`, \`CLOUD\_LINK\`, \`DNS\`, \`ROUTED\_DDOS\`. Defaults to
\`INFRASTRUCTURE\`.

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

<a id="canonical-1132102222000120-1312023210212021-2212103031202003-2303010133130331-2001130111201011-0311320332230320-2230201020303103-2220332311321331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.notification_parameters` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [routes](resources--alert_policy--reference--group-001.md#canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033)
- routes.notification_parameters

<a id="canonical-3333002302120121-2031222300320321-3233232221223000-0022310320320300-0103212310032030-0203030302032132-1301332130131331-3310121331110103"></a>

Type: `"object"`. single nested block, Optional.

Set of notification parameters to decide how and when the alert notifications should be sent to the
receivers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom",
    "default"),
  validators.ConflictingObjectAttributes("custom",
    "individual"),
  validators.ConflictingObjectAttributes("custom",
    "ves_io_group"),
  validators.ConflictingObjectAttributes("default",
    "individual"),
  validators.ConflictingObjectAttributes("default",
    "ves_io_group"),
  validators.ConflictingObjectAttributes("individual",
    "ves_io_group")}
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
  "x-ves-oneof-field-group_by": "[\"custom\",\"default\",\"individual\",\"ves_io_group\"]"
}
```

Terraform syntax:

```terraform
notification_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-1223212302222232-0132002220321231-3131211311322130-0131330012110111-1233002130121133-0030301211012320-2220231001032102-2023032130123321"></a>

### Direct properties for `routes.notification_parameters`

- [custom](resources--alert_policy--reference--group-001.md#canonical-3213323001131022-3011101233110210-2011023022003012-3000121111232231-3223321212203220-3211102101212011-0200222201223020-3300020113113011): complete subsection reference.

- [default](resources--alert_policy--reference--group-001.md#canonical-1300112330323001-1320223033102300-0210310100122033-0112133021103330-2031320131313330-3000220230303320-2012320132123012-0123033320021302): complete subsection reference.

<a id="canonical-3000103032212100-2000230200223211-3022330121211312-0031200332323301-2333310101330002-2302122202313111-3320110301023222-3330110211233123"></a>

<a id="canonical-3021110112201102-0130102200102221-2321011130033020-3202210212121322-2100030300311022-1202210331222012-2203102010113000-1231222333331333"></a>

#### `routes.notification_parameters.group_interval` property

Type: `"string"`. Optional.

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval defaults to "1m"

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-1013233311300100-3113221230112323-2101212010121302-3203013220221202-2312022333122213-0332231132322001-3001021002312031-2232001221120101"></a>

<a id="canonical-3123132112303010-2222213032312210-3000310323101120-2201331221302323-1200021102031102-3012303310231120-0011213213103321-1313111322022131"></a>

#### `routes.notification_parameters.group_wait` property

Type: `"string"`. Optional.

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to '30s'.

Additional upstream details:

Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_wait defaults to "30s"

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [individual](resources--alert_policy--reference--group-001.md#canonical-2313311003320030-3113122200311303-1211111233221231-2213112322222321-1333302222030020-1013230333231121-1013222121332023-2310330223302223): complete subsection reference.

<a id="canonical-3013203321120321-1220210303131202-1303321213113202-0132301031332113-3121102303201303-3112112130010212-0310032031313020-1121323103211212"></a>

<a id="canonical-2000211002332130-0011030312313330-3000000330013111-2111110201222001-0113213130220021-1133310221303310-2131023230133321-1211121321212132"></a>

#### `routes.notification_parameters.repeat_interval` property

Type: `"string"`. Optional.

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to '4h'.

Additional upstream details:

Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval defaults to "4h"

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [ves_io_group](resources--alert_policy--reference--group-001.md#canonical-2012331301013322-0212101002121031-0301201032012321-3210000221311300-1303301323302001-0021101321010231-0123013223131200-3013213003320023): complete subsection reference.

<a id="canonical-3213323001131022-3011101233110210-2011023022003012-3000121111232231-3223321212203220-3211102101212011-0200222201223020-3300020113113011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.notification_parameters.custom` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [routes](resources--alert_policy--reference--group-001.md#canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033)
- [routes.notification_parameters](resources--alert_policy--reference--group-001.md#canonical-1132102222000120-1312023210212021-2212103031202003-2303010133130331-2001130111201011-0311320332230320-2230201020303103-2220332311321331)
- routes.notification_parameters.custom

<a id="canonical-0233120213322102-0003102023201102-1333202033231221-0131211230002101-3202013230321232-1100220200130023-0313120313330323-2011203130113302"></a>

Type: `"object"`. single nested block, Optional.

Specify list of custom labels to group/aggregate the alerts.

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
custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-1110011210001301-1322133301131030-2223302313012111-1020110301033300-3123033001322212-0302302020213303-0302232110013112-3232130121100222"></a>

### Direct properties for `routes.notification_parameters.custom`

<a id="canonical-1220121301322112-0302033000321132-1100321201312113-2120022311313110-1120212123012000-0022320310121111-1113300020301122-2321322220132313"></a>

#### `routes.notification_parameters.custom.labels` property

Type: `["list", "string"]`. Optional.

Name of labels to group/aggregate the alerts.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1300112330323001-1320223033102300-0210310100122033-0112133021103330-2031320131313330-3000220230303320-2012320132123012-0123033320021302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.notification_parameters.default` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [routes](resources--alert_policy--reference--group-001.md#canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033)
- [routes.notification_parameters](resources--alert_policy--reference--group-001.md#canonical-1132102222000120-1312023210212021-2212103031202003-2303010133130331-2001130111201011-0311320332230320-2230201020303103-2220332311321331)
- routes.notification_parameters.default

<a id="canonical-1012002322102222-1021310100002322-0103212012032300-2033221123020001-0322100321200302-2220232113033230-0311322111203022-2321200111233202"></a>

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
default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313311003320030-3113122200311303-1211111233221231-2213112322222321-1333302222030020-1013230333231121-1013222121332023-2310330223302223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.notification_parameters.individual` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [routes](resources--alert_policy--reference--group-001.md#canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033)
- [routes.notification_parameters](resources--alert_policy--reference--group-001.md#canonical-1132102222000120-1312023210212021-2212103031202003-2303010133130331-2001130111201011-0311320332230320-2230201020303103-2220332311321331)
- routes.notification_parameters.individual

<a id="canonical-1331001302112113-0310033031221300-0000200232232301-1101102331002122-3322322222230213-0322213110311021-0013001222022121-0320001110123110"></a>

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
individual = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012331301013322-0212101002121031-0301201032012321-3210000221311300-1303301323302001-0021101321010231-0123013223131200-3013213003320023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.notification_parameters.ves_io_group` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [routes](resources--alert_policy--reference--group-001.md#canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033)
- [routes.notification_parameters](resources--alert_policy--reference--group-001.md#canonical-1132102222000120-1312023210212021-2212103031202003-2303010133130331-2001130111201011-0311320332230320-2230201020303103-2220332311321331)
- routes.notification_parameters.ves_io_group

<a id="canonical-2313033020011230-0112033221100303-3213302121333213-3213313123222101-2010020113231000-3001012023221330-0301213221333313-2101332032030300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ves io group.

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
ves_io_group = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110310313011121-1112331112232321-2332311130112102-0101233330303222-2032322323223301-3100321320112112-0230232022322233-1202023210320223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.send` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [routes](resources--alert_policy--reference--group-001.md#canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033)
- routes.send

<a id="canonical-0223313000203232-1002323001123120-3022021200313210-0023220102133233-2031323030103120-0333212113310030-1011303002220222-0002012022331020"></a>

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
send = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033001201101022-2221131331322322-1113222133113330-3211223030232202-3033130223101322-3133000201133203-3000131300220330-1023330031001313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.severity` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- [routes](resources--alert_policy--reference--group-001.md#canonical-2130320122111321-3211321113022113-2302033210311013-2320311313020000-0321211210311022-2301210032020133-2200131112233201-3001032100200033)
- routes.severity

<a id="canonical-0300112321332330-2132332211333020-1201031122110322-2211110220000301-2233211312230123-0320133130023123-3023122123231102-0031313211300301"></a>

Type: `"object"`. single nested block, Optional.

Select one or more severity levels to match the incoming alert.

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
severity {
  # Configure direct properties listed below.
}
```

<a id="canonical-1312220100323101-1301231000101220-1032022100022311-0023220313230003-3300003131301232-2312113220203013-0122131303230220-2112111022022101"></a>

### Direct properties for `routes.severity`

<a id="canonical-3033002121033331-2000032210312310-1220021313000111-0311120001002312-2212220233302210-3010313202110302-1301102022331002-3112130030333021"></a>

#### `routes.severity.severities` property

Type: `["list", "string"]`. Optional.

\[Enum: MINOR|MAJOR|CRITICAL\] Severities. List of severity levels. Possible values are \`MINOR\`,
\`MAJOR\`, \`CRITICAL\`. Defaults to \`MINOR\`.

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

<a id="canonical-0232311332023021-1302331123211130-1333333130122230-1313103331111032-3302231011300011-1002113030003001-1130133123122213-2330301331133121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_alert_policy](../resources/alert_policy.md#canonical-1131112113200013-1111133122303310-3323002103120120-3332003320311212-2230020100302130-3333100100222023-1233003312313110-1031321030030100)
- [Property reference](resources--alert_policy--reference--group-001.md#canonical-0313231030203220-0020221223300020-0330020213312031-2003313201130301-0220021112010310-2120231231002201-0120232133233211-2002302212201122)
- timeouts

<a id="canonical-3110210011201330-1023303000322123-2233012122331302-3011120021201100-3201233010001130-0222330320023110-2120313123203202-3103112230112111"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2133301311323133-1303330103100203-2101102303321312-1013220101303221-1013003211232020-2122111130002110-3123032331310033-3001203312031300"></a>

### Direct properties for `timeouts`

<a id="canonical-3312203233110313-1113131310030333-2232022021001323-2211322003001322-2202311130121333-2121021210200130-0131133300113011-0001213323210013"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0303030220321111-0033323211310102-3230333033011120-3001320021313302-3001302033211110-2121232332013203-3320103112313303-3120320203120200"></a>

<a id="canonical-2221023211232012-1220122130033233-2113231233131323-3112332323111210-3331123301311030-3130012323211021-2220231122221002-1313012131030213"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1202323023332032-3331321301333110-1200011131311211-1233323132202202-1111210321223000-2202221302032302-1230200300212032-3002010113031030"></a>

<a id="canonical-3111333230011210-3332300220123130-2323202201223033-2003213313303113-1113101102022033-0033013133111001-3311321002012011-2023021022131223"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3002033133200122-1031220210030021-1100000010230101-2033112322211020-2031223231010323-2322003100303102-2013123320111022-3122101011011013"></a>

<a id="canonical-1213012003121000-2021200001210002-0222123200200332-0100032121032110-0222302312320310-1211120302312022-2121222223000121-3212120010020100"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
