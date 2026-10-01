---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-3012120331203201-1112000130332131-3202332120003001-0123013103333002-1112013103100001-2021323312110130-1112011111221321-2323233113231232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331131030201020-0120320010011230-2321003333323321-0002313200230202-2221120233033130-3300213102232202-0103113230011112-0013022032202102"></a>

## Property reference — Property reference / 332123200132 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md#canonical-2300201333020020-2232222123331320-0121101102033133-3300123100312103-1100330100011331-1000302303112121-2130232232220122-1003132230103113)
- Property reference

<a id="canonical-1302012313011021-1130223231003123-3113131000001111-3200231133322002-2233123231012220-2201232311010001-2001033021121333-0303032132200322"></a>

## Direct properties — Property reference / 332123200132 / 3

<a id="canonical-3222221300001231-2203333013230013-0222013120023130-3203212313321300-3232303333131121-0223300001122333-2011200012320223-2333212010330013"></a>

<a id="canonical-1123101310001003-0030002110313302-0333011301032000-1212012203013203-3013213302330133-1201113322110220-2211211213132110-0301332213120132"></a>

## address property — Property reference / 332123200132 / 4

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

- [admin_password](resources--azure_vnet_site--reference--group-003.md#canonical-2210301132001310-0233003201110121-0301133333210313-0022133021010133-0010001323111330-3131200010232310-3120111330201220-0103030322130323): complete subsection reference.

<a id="canonical-2230000233013130-3321020332210230-0132133020303130-3032200002301333-3210120032033011-2210023103203212-0112311010301312-2013002301200011"></a>

<a id="canonical-1001132201330202-2331311312202200-2320222033303310-1210100022102210-1120022300023101-3110130132323312-0311132010022113-0111222120332202"></a>

## alternate_region property — Property reference / 332123200132 / 5

Type: `"string"`. Optional, Computed.

\[OneOf: alternate\_region, Azure\_region\] Exclusive with \[Azure\_region\] Name of the Azure
region which does not support availability zones.

Upstream description:

Exclusive with \[Azure\_region\] Name of the Azure region which does not support availability zones.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

OneOf alternatives in this subsection:

- [alternate_region](resources--azure_vnet_site--reference--group-001.md#canonical-2230000233013130-3321020332210230-0132133020303130-3032200002301333-3210120032033011-2210023103203212-0112311010301312-2013002301200011)
- [azure_region](resources--azure_vnet_site--reference--group-001.md#canonical-0012002211331203-3020311310121221-1330031212000031-3110332312003303-2230121133111102-3332222323321212-2221302213123213-0313203002232003)

Select alternatives according to the provider validators above.

<a id="canonical-0320303222021133-3332203220102302-1023321002301112-0013032031111103-0022101312123111-3100231213022330-1013110321113030-0023322023032002"></a>

<a id="canonical-3211031332321012-3312123023111313-1201103330003221-0313321200120301-1232113131021023-3312310332002223-2122221322111203-0033312201331213"></a>

## annotations property — Property reference / 332123200132 / 6

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

- [azure_cred](resources--azure_vnet_site--reference--group-003.md#canonical-3230011000033002-3323231031123210-2202220131333011-2133332202030133-0102111333103212-0313010210131333-2000322000022312-0310231102130011): complete subsection reference.

<a id="canonical-0012002211331203-3020311310121221-1330031212000031-3110332312003303-2230121133111102-3332222323321212-2221302213123213-0313203002232003"></a>

<a id="canonical-1000230112302021-3112300202030000-2322233310322101-3001003212230203-2220311322103022-1323201320012031-2300330210311212-3231101021230012"></a>

## azure_region property — Property reference / 332123200132 / 7

Type: `"string"`. Optional, Computed.

Exclusive with \[alternate\_region\] Name of the Azure region which supports availability zones.

Upstream description:

Exclusive with \[alternate\_region\] Name of the Azure region which supports availability zones.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [block_all_services](resources--azure_vnet_site--reference--group-003.md#canonical-0102302130001001-3200320201310210-2332002131130332-2120323011211313-0122022331010012-2003232011022231-3320202011110231-1313211311312002): complete subsection reference.

- [blocked_services](resources--azure_vnet_site--reference--group-003.md#canonical-3130101112012001-2331300030123213-3211202221212033-2103010203120020-0101230331303200-0000012023202001-3121321331212010-3211233030210112): complete subsection reference.

- [coordinates](resources--azure_vnet_site--reference--group-003.md#canonical-1303220212123123-1311213130112211-1121321020332031-0123033333312122-1332223213121121-0220232322122003-2320320123331331-2130221010313010): complete subsection reference.

- [custom_dns](resources--azure_vnet_site--reference--group-003.md#canonical-3202312201131023-0122231100102303-3111003330022100-3200132320312012-0213031123311311-1102113233231110-1001223221313130-2301331330230212): complete subsection reference.

- [default_blocked_services](resources--azure_vnet_site--reference--group-003.md#canonical-1003130113132011-2303203000130010-2301330110330213-0321202211112213-3303211213321102-3100003110330032-0230021301000022-1000020001020222): complete subsection reference.

<a id="canonical-0133012201312011-1122323200032031-2001232320012002-3120223121121303-2111332113010201-2220101330020221-0202231321311210-0111003110012323"></a>

<a id="canonical-2201101022112121-3231213312312103-2222131020323212-2100310313012101-0113120020110130-2133213102223112-0033202000113132-2020231001030333"></a>

## description property — Property reference / 332123200132 / 8

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

<a id="canonical-0320323022120303-0111333130331222-3221223010102233-0213203201112120-3230310130303301-2221221033312112-0100333020121222-0320111320102220"></a>

<a id="canonical-2032201122023120-3313332302310203-0030112032320221-0223011012320001-0210331221230033-1100221232001033-0023032302013310-2231332112120101"></a>

## disable property — Property reference / 332123200132 / 9

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

- [disable_encryption](resources--azure_vnet_site--reference--group-003.md#canonical-3210310300010211-0031103121230221-3230020213223000-1113033302033332-0012132322012231-0101312023020133-3330211001200322-2112223012021311): complete subsection reference.

<a id="canonical-3200013021013031-3031311031312123-1103113033233200-1013310313012232-2112011220230010-3022012323220011-1021202331002312-0120112110033332"></a>

<a id="canonical-1133230300022101-3110102221212210-2213210211011300-3212300323200333-0310013002111212-1331131330100321-2032213011003320-2131020331002111"></a>

## disk_size property — Property reference / 332123200132 / 10

Type: `"number"`. Optional, Computed.

Disk size to be used for this instance in GiB. 80 is 80 GiB. Server applies default when omitted.

Upstream description:

Disk size to be used for this instance in GiB. 80 is 80 GiB.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(4095),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
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
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

- [enable_encryption](resources--azure_vnet_site--reference--group-003.md#canonical-1000310322001320-2311023113002011-3103023303010213-3012300211031000-1221202202331303-1003131333000322-1213302333210100-2312200331021300): complete subsection reference.

<a id="canonical-2113102222122230-1002002231320011-1113033323220313-1311203131130010-3020102312021300-1001220203131213-2030300300222002-2011233122133323"></a>

<a id="canonical-3021321003303010-1221011023312121-1213210023310231-3222000131312023-2330332222220203-2310133122232133-0311223112223333-2221120011102133"></a>

## ID property — Property reference / 332123200132 / 11

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](resources--azure_vnet_site--reference--group-003.md#canonical-2010130103310022-3113022111002313-1033202302030223-1101200032100112-1013100022301211-0003100230013333-3111003100333222-1021233113020120): complete subsection reference.

- [ingress_egress_gw_ar](resources--azure_vnet_site--reference--group-005.md#canonical-0312012003033102-3210313212120203-0230331020120000-0030121220200120-2223003302123103-2232111103103011-2023112221000212-0212333313012311): complete subsection reference.

- [ingress_gw](resources--azure_vnet_site--reference--group-007.md#canonical-1200122201231222-2333130201131133-2113103023310000-2213221203001333-2113320112011303-3300011003101011-0002313330212230-0130220110020220): complete subsection reference.

- [ingress_gw_ar](resources--azure_vnet_site--reference--group-008.md#canonical-3100320201112313-1011012113210322-1000233330130201-0100100022211103-0200311001112230-3202222013102010-0332322302323212-1211221101203010): complete subsection reference.

- [kubernetes_upgrade_drain](resources--azure_vnet_site--reference--group-008.md#canonical-2311300122330033-0220102210320010-0211200301301230-0303112033111231-3321210103223331-0131201322003031-3210322331300231-2121020201033300): complete subsection reference.

<a id="canonical-3013233333101132-2333332302313201-0022013002301112-3333221122333221-2121131130122211-3030323323133300-2322301333303023-3121131200111031"></a>

<a id="canonical-0222213311002230-1112031131311123-0012322023203203-0023122112131111-0021310232312033-2213221122000033-1222000033210123-0200022220012213"></a>

## labels property — Property reference / 332123200132 / 12

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

- [log_receiver](resources--azure_vnet_site--reference--group-008.md#canonical-1112132000301212-0000210332100022-0311233030332311-1302311011100310-3103210200032102-1221331331333212-3231310120202331-3132101333022222): complete subsection reference.

- [logs_streaming_disabled](resources--azure_vnet_site--reference--group-008.md#canonical-0211021120303002-0301222131100002-1223023031320133-0132301120230322-3232013302201132-2211202011102212-2112131111300021-3332230131331211): complete subsection reference.

<a id="canonical-0330023200122232-3221103231023120-3332133113120313-2300030111331311-0131103131212022-3202022131102321-2200310013030203-3120332123302200"></a>

<a id="canonical-2330020101210332-3122113223130233-2110012100033133-0312201231311302-0332303323320001-0122300330303113-3020311333020220-3130220323012112"></a>

## machine_type property — Property reference / 332123200132 / 13

Type: `"string"`. Required.

Select Instance size based on performance needed. The default setting for Accelerated Networking is
enabled, thus make sure you select a Virtual Machine that supports accelerated networking or disable
the setting under, Select Ingress Gateway or Ingress/Egress Gateway &gt; advanced OPTIONS.

Upstream description:

Select Instance size based on performance needed. The default setting for Accelerated Networking is
enabled, thus make sure you select a Virtual Machine that supports accelerated networking or disable
the setting under, Select Ingress Gateway or Ingress/Egress Gateway &gt; advanced OPTIONS.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-2210221210122003-0123021023012312-1322202122332021-3201312103123030-1003202031333200-0220323122031332-1232132123121111-3032133210030223"></a>

<a id="canonical-1120230220021001-2330201012121021-0123313213111201-1303002031203221-3020033213331202-3333110112031031-1012300332022323-3212113112000123"></a>

## name property — Property reference / 332123200132 / 14

Type: `"string"`. Required.

Name of the Azure VNET Site. Must be unique within the namespace.

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

<a id="canonical-0131021320301300-3313100131001100-1121133230330132-2031310110223020-3110231010013101-2300210202033301-1302223021300031-1220333332113310"></a>

<a id="canonical-1131110111022223-3202122110201323-1203020000001322-3301033201322222-0122221110220212-2220313003112103-0232121331123212-1112100202223313"></a>

## namespace property — Property reference / 332123200132 / 15

Type: `"string"`. Optional, Computed.

Namespace for the Azure VNET Site. The F5 XC API restricts this resource to the system namespace; it
defaults to that value and may be omitted.

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

- [no_worker_nodes](resources--azure_vnet_site--reference--group-008.md#canonical-0102002101312223-3010101332001213-0113101003023111-2131000301330100-2303130300233320-0323301121001101-1330121210330132-2330121312020032): complete subsection reference.

<a id="canonical-3313100110223103-3200233300212030-1130133211032320-1110022012130311-1121331011311311-0123230233300333-0202111001231212-1121123012031022"></a>

<a id="canonical-0310233220203311-1113003333033232-0111211203003202-3201020111322100-3230200021113220-0121223231213002-1202010313100323-3011120120020132"></a>

## nodes_per_az property — Property reference / 332123200132 / 16

Type: `"number"`. Optional, Computed.

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Upstream description:

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 21),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  }
}
```

- [offline_survivability_mode](resources--azure_vnet_site--reference--group-008.md#canonical-1213020322301022-3003203130102301-1221012222003033-2131222331322322-3102132101222032-2310332230000033-3132310321130103-1130223011200320): complete subsection reference.

- [os](resources--azure_vnet_site--reference--group-008.md#canonical-1300313100020023-0200203102221232-3122331022032222-3203003310333312-1333333032323202-1333333020112233-0202302102313001-0332202313002233): complete subsection reference.

<a id="canonical-0130201302020033-2022211130103211-2103003000302130-1230313221003030-3101221002221111-3013233300002100-2233232201011332-0102101111302320"></a>

<a id="canonical-2002001012113121-1201201032012101-3231101013113122-0112201033213233-3211300000120302-3102022012100113-3023001321013321-3123003013223002"></a>

## resource_group property — Property reference / 332123200132 / 17

Type: `"string"`. Required.

Azure resource group for resources that will be created.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3111010033230321-1032001310332020-0033030000101310-0113121031311332-3023301232330102-0221111010111003-1022021010100211-1302210020111033"></a>

<a id="canonical-3120032321323123-3313220221313111-0111000100022300-1321333201331123-3122030032120103-1102211221132131-1110020201130332-3012113302322101"></a>

## ssh_key property — Property reference / 332123200132 / 18

Type: `"string"`. Required.

Public SSH key. Public SSH key for accessing the site.

Upstream description:

Public SSH key for accessing the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [sw](resources--azure_vnet_site--reference--group-008.md#canonical-1010301121111312-1322003232222003-2031100301011123-1331030222133000-1230201012311020-1122001230113333-3021021031220330-3232301313033113): complete subsection reference.

<a id="canonical-0021101211133203-2202112132022330-2031230212223001-0231122102121022-1321100213113333-1122212111001302-0213133301322322-2110101020223030"></a>

<a id="canonical-0122031311103033-3320230301000223-2101003231110023-2100303202210201-0031030022221033-0000003022231330-2123230231123212-3012330320220121"></a>

## tags property — Property reference / 332123200132 / 19

Type: `["map", "string"]`. Optional, Computed.

Azure Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in Azure console. Defaults to \`map\[\]\`. Server applies
default when omitted.

Upstream description:

Azure Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in Azure console.

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

- [timeouts](resources--azure_vnet_site--reference--group-008.md#canonical-1322223300323213-0231222120332102-1102301010221302-3022023203110133-2322311230213221-1332133300210200-2011013021121101-1221221231220201): complete subsection reference.

<a id="canonical-0232321202231010-2303100112332202-0233212003300022-2211103000122103-2113203222313203-0101222333310110-2210033023323033-2233331012111001"></a>

<a id="canonical-2100033330100313-1002030211011323-3133231023132231-1011003011202133-3232223313131003-0201221003030131-1011312110112212-3220011231300030"></a>

## total_nodes property — Property reference / 332123200132 / 20

Type: `"number"`. Optional, Computed.

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Upstream description:

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 61),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 61,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  }
}
```

- [vnet](resources--azure_vnet_site--reference--group-008.md#canonical-2030003203102003-3211120202023310-3332230033311222-2030101330030212-1121003230120332-2223020303131011-1031203132311202-2111303213312103): complete subsection reference.

- [voltstack_cluster](resources--azure_vnet_site--reference--group-008.md#canonical-3101033031331331-3020213110130012-2102010101331003-3121320201023121-3030303022233301-0030021122120010-0330103121002130-3221012320030110): complete subsection reference.

- [voltstack_cluster_ar](resources--azure_vnet_site--reference--group-009.md#canonical-1233112000303130-0322113212321213-0111101101130333-3211122202322203-2100201302230322-2023333331032102-2310112100312223-3021111221211133): complete subsection reference.

- [waf_signatures](resources--azure_vnet_site--reference--group-010.md#canonical-2000101211023100-3303010230231233-0133333212103100-1101121132322201-3121232303021110-3333033122312013-2303300130313111-3332232300233123): complete subsection reference.
