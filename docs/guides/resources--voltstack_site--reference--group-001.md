---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202220133032313-3020221112000130-3113013103330332-0131310123330100-1232113221202131-3010210223213121-2321222030123220-1132031331112013"></a>

## Property reference — Property reference / 022302131011 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- Property reference

<a id="canonical-2202112100032112-2333310203232113-3120212131113223-0322133132003320-2313222102020100-1013010323132121-1123310003112312-0020020102222330"></a>

## Direct properties — Property reference / 022302131011 / 3

<a id="canonical-2003330122310210-0102211010002223-2331113313101233-3133130322121311-1012031123130031-0120332032220120-2020103322131210-1320210230233130"></a>

<a id="canonical-0010220012102100-1023123310131100-1100031120013302-3132202030213030-1330213230101121-0130112113020103-3203300012011100-1202000120301033"></a>

## address property — Property reference / 022302131011 / 4

Type: `"string"`. Optional, Computed.

Site's geographical address that can be used to determine its latitude and longitude.

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

- [allow_all_usb](resources--voltstack_site--reference--group-003.md#canonical-0030120021301113-2123132320331013-1120102200032123-3222210012001222-1020233230011133-1323031022303022-0311233200203212-0012310132210301): complete subsection reference.

<a id="canonical-3033223121112222-1000010321233210-2202302002232220-1312001101200111-1111310123100010-2221103030200020-1311011333012210-1232211320311320"></a>

<a id="canonical-0113210001322011-1231233303130031-3333132023311032-0030001220223313-2100032030033111-0123022332333323-0033021032220022-1111132131321213"></a>

## annotations property — Property reference / 022302131011 / 5

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

- [blocked_services](resources--voltstack_site--reference--group-003.md#canonical-2020132101313030-3113220123311212-0222132220232001-0131232001123101-0331312303313212-1102212023321301-2111310003032332-0112233122301303): complete subsection reference.

- [bond_device_list](resources--voltstack_site--reference--group-003.md#canonical-3030103000231013-1102103111323200-0331203022323213-3221021030201010-2113303311303120-3003303202322133-2100131233333213-2032320213020130): complete subsection reference.

- [coordinates](resources--voltstack_site--reference--group-003.md#canonical-2030333302021000-3122102200321211-1013103203020030-2010332120131323-3332113230202332-0212120100201220-0310323121132012-2123232100022023): complete subsection reference.

- [custom_dns](resources--voltstack_site--reference--group-003.md#canonical-3010120121022300-0332021013031021-0010322332231022-0021103021100111-1313020321320300-0120102001111321-2020002020113003-2000110203023302): complete subsection reference.

- [custom_network_config](resources--voltstack_site--reference--group-003.md#canonical-2013113330012020-2023003101322103-1112120002210200-1003203232123221-1132002321132310-1033003311221232-2113331202032203-0303020131203001): complete subsection reference.

- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303): complete subsection reference.

- [default_blocked_services](resources--voltstack_site--reference--group-009.md#canonical-2130132222030011-3302231322323332-2022003322201021-3320010312002031-2110201300121320-3303303203301232-1023233112330131-2030030202101202): complete subsection reference.

- [default_network_config](resources--voltstack_site--reference--group-009.md#canonical-0112131200332302-1221213220313221-0032220113131303-1021302330212020-3203213003303121-2130001313123223-2021030302131113-1223131022132013): complete subsection reference.

- [default_sriov_interface](resources--voltstack_site--reference--group-009.md#canonical-0333212220131232-2232130221010210-1121201121210113-3310332003011111-3130113303011212-3232313102100302-3300112310202231-0003110030333311): complete subsection reference.

- [default_storage_config](resources--voltstack_site--reference--group-009.md#canonical-0231300133322100-2202330301302132-0012003013121222-1003333120321002-3123333301332301-1111202112020313-2132220133031330-1332323333333102): complete subsection reference.

- [deny_all_usb](resources--voltstack_site--reference--group-009.md#canonical-0011333310220211-1323201302012101-1220001201131311-3020322100132123-3310230132323001-3232033103012131-0111122033312000-1002013231313120): complete subsection reference.

<a id="canonical-2210132212010110-0112030000300132-2103112012103302-0010030101020001-2021323132203100-3012213002002012-1220021001302223-3130111010311230"></a>

<a id="canonical-2133030000232301-0100323330131113-2000313112103120-1123330122113333-1012122231021330-0331120332030011-3131221020010003-0022020011222231"></a>

## description property — Property reference / 022302131011 / 6

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

<a id="canonical-0320300101010332-2103220322302220-0212133112313130-0100302020023101-2330132203302302-2133330222130113-1003101111312323-3202031330230122"></a>

<a id="canonical-0232102330022001-1200013101030023-1301332220222003-0021321302132002-2312033002212001-1331303300122121-1010231012002030-1001122031210101"></a>

## disable property — Property reference / 022302131011 / 7

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

- [disable_gpu](resources--voltstack_site--reference--group-009.md#canonical-2001102123123301-1032022332333230-2033203001333102-2301030133321222-0231303123012321-2313313000331121-1121001130310123-2211322122100120): complete subsection reference.

- [disable_vm](resources--voltstack_site--reference--group-009.md#canonical-2303200333033000-1103112001331020-2330122121000001-0212131223103310-1320003332230302-2031132100202331-0212021000103231-3033113123220331): complete subsection reference.

- [enable_gpu](resources--voltstack_site--reference--group-009.md#canonical-3110030200022202-0211030001002221-3101311323333202-0033031300310321-2111030230000111-0011021012012333-2022321231103223-0201201200013113): complete subsection reference.

- [enable_vgpu](resources--voltstack_site--reference--group-009.md#canonical-0221232202222003-3012320211333202-3133320111220022-3322020133223311-3301131103333231-2212300202332303-0001010232100011-3200231310210210): complete subsection reference.

- [enable_vm](resources--voltstack_site--reference--group-009.md#canonical-2203102322033333-2320020222223323-3002100113202302-3211002201121201-3222121033010011-1012332312100221-0000322233121131-0123131131200320): complete subsection reference.

<a id="canonical-3131312313302313-3133100033313123-1220332311323033-1210332113333021-2101211012322311-3310210201031003-0012320111310223-2202333123212031"></a>

<a id="canonical-1333000303101121-0220023103121302-1022123222200111-3103310311222121-2211220131011001-1033132023100330-3031313003033332-0112211003200131"></a>

## ID property — Property reference / 022302131011 / 8

Type: `"string"`. Computed.

Unique identifier for the resource.

- [k8s_cluster](resources--voltstack_site--reference--group-009.md#canonical-0033231220232021-3311331322200101-1011323101020101-2311312312310112-1201020132031333-0332133211003211-2230203202320011-2111013112002230): complete subsection reference.

- [kubernetes_upgrade_drain](resources--voltstack_site--reference--group-009.md#canonical-0321231022303102-3013003323223300-2001323323123222-1211001221020013-3022220032230101-3000222101031212-1101323201122001-3031320013200111): complete subsection reference.

<a id="canonical-0322320102122113-0230213023323231-2331101010110320-0223232120310113-2200123320021030-2223013002212330-1211202312001300-0313103130012330"></a>

<a id="canonical-1012200323003231-0232332032100121-3301302230003122-1100110202021211-1212131231222303-3023203202213230-3021031021300300-1323330112130023"></a>

## labels property — Property reference / 022302131011 / 9

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

- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-1132320332333230-2221113332303320-2330321010200300-2023122310102023-1022200001333011-3303122211221301-1130130210003310-1133211030011101): complete subsection reference.

- [log_receiver](resources--voltstack_site--reference--group-010.md#canonical-3310021323320200-2131332321230123-0021023332302120-2312333000133000-0003300021330330-3031120120202103-0332000001202021-2012232322101300): complete subsection reference.

- [logs_streaming_disabled](resources--voltstack_site--reference--group-010.md#canonical-2013003210030303-0313301312303101-1102333300333132-2302212310311100-1312130101123323-1332121302130312-0122131212310021-2013303121303232): complete subsection reference.

- [master_node_configuration](resources--voltstack_site--reference--group-010.md#canonical-2231013002202321-1310031001031301-2022032301131330-1102123002203131-0101322113031031-0102323000220120-2121200313301302-1032303230301033): complete subsection reference.

<a id="canonical-0010310313213120-2102000201210300-1232100100231133-2011323322203303-0013011022223103-2033121131131303-0322130203110230-1320030202230010"></a>

<a id="canonical-3321031302313323-0200121111221001-3230133212310210-2200113001103033-0123111210122202-3033002010320120-3323012021331230-1232312321013011"></a>

## name property — Property reference / 022302131011 / 10

Type: `"string"`. Required.

Name of the Voltstack Site. Must be unique within the namespace.

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

<a id="canonical-2312133331333333-1113211321330102-1201132323310112-0030231322110331-2100121123102310-1030222020120311-2020210221230333-3121103103102012"></a>

<a id="canonical-2122130121301020-3201213002011013-3020033311311021-0300210003120100-0031020121002032-0230311221211001-0200031220123102-3112310131221320"></a>

## namespace property — Property reference / 022302131011 / 11

Type: `"string"`. Required.

Namespace where the Voltstack Site is created.

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

- [no_bond_devices](resources--voltstack_site--reference--group-010.md#canonical-0111122002330203-3213023011101331-1202010121032023-2310021300012322-3231300020223000-0312320110132201-3233220021110031-1221330021232331): complete subsection reference.

- [no_k8s_cluster](resources--voltstack_site--reference--group-010.md#canonical-1110311231110111-0010333301003210-2303111023211302-3123121321000123-1332030003230203-2130132012333122-1121202021001232-1223130230213223): complete subsection reference.

- [no_local_control_plane](resources--voltstack_site--reference--group-010.md#canonical-1212130122323321-1021203223212003-2213000310220323-3003323121130210-3312333213031112-1210302211021100-1311032332331321-1322130033211213): complete subsection reference.

- [offline_survivability_mode](resources--voltstack_site--reference--group-010.md#canonical-0120123201103100-3121123323333132-2323310020001010-3232000323111313-0230313201222013-3202130322013321-2133303000333212-0112033111011301): complete subsection reference.

- [os](resources--voltstack_site--reference--group-010.md#canonical-2230223331213210-2003111102030202-0101123130331303-1321202222013223-3333311120130122-1231202131312311-1202210112301113-3211213301122003): complete subsection reference.

- [sriov_interfaces](resources--voltstack_site--reference--group-010.md#canonical-3202202012201132-3231030210000322-0123032022001103-3031211120312222-3332121322033010-1220331122103013-1123232303130120-1200100021331233): complete subsection reference.

- [sw](resources--voltstack_site--reference--group-010.md#canonical-1310110011022322-0333211000331321-3000220132023131-1232213111323122-3332210311120230-2120231033120210-3213230302031001-1112101021020320): complete subsection reference.

- [timeouts](resources--voltstack_site--reference--group-010.md#canonical-1130122203230021-0301012333023233-1330223323223133-3120132222110331-2110010321001222-2011322100322230-0323000302330311-0022221031222233): complete subsection reference.

- [usb_policy](resources--voltstack_site--reference--group-010.md#canonical-3332312032102223-1310111330103012-1312102310030023-3213120011021310-2232310220220331-3331131211023113-3023103301303211-1013210310330303): complete subsection reference.

<a id="canonical-0333221321233013-0201000001122033-3121130222300322-3011233303230223-0221233123132301-3301313112013032-2102113121213021-0020011101032121"></a>

<a id="canonical-0022302300101311-0002231130210202-0211002022020012-1031102132002302-3023020033003310-1301111312323022-3131313020332310-1330323303133332"></a>

## volterra_certified_hw property — Property reference / 022302131011 / 12

Type: `"string"`. Required.

Name for generic server certified hardware to form this App Stack site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

- [waf_signatures](resources--voltstack_site--reference--group-010.md#canonical-3001210003032302-1322031333003130-3131300313220213-0112200202131003-3012101010211301-1323112023133110-3121212221011012-2031210320020132): complete subsection reference.

<a id="canonical-3322221213013313-0100022212013303-2031302221203200-3023031022230022-2011231210321310-2300200101213332-3131111310003220-0102311222023330"></a>

<a id="canonical-0011020301332200-2313212102301122-0013101212113302-3332332312121221-0113213330322122-3010112021200010-3002113221230001-3300202310221220"></a>

## worker_nodes property — Property reference / 022302131011 / 13

Type: `["list", "string"]`. Optional.

Worker Nodes. Names of worker nodes.

Upstream description:

Names of worker nodes.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```
