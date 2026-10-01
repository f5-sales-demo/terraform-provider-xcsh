---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-1233001102111021-0232012331113311-1203200031110312-2301102003013133-0302102223232222-0321010302310210-3310100000003032-0000001213330101"></a>

## data_lif_dns_name property — netapp_backend_ontap_san / 222032030112 / 5

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1230323320311212-3321000121000220-3202101030230103-3033121310012211-2020103230030033-1100222221032121-1012230321223200-3103132020123000"></a>

<a id="canonical-3320030003032221-0121301020320132-3233102332030223-1231201203213102-1313122333311332-1003311001131303-1221113110100300-3112030002310102"></a>

## data_lif_ip property — netapp_backend_ontap_san / 222032030112 / 6

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Upstream description:

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3113110110230201-1333131201332302-0310000202300211-0221233333122202-1001100112232130-0111000112103022-2032002110300223-2030321313032203"></a>

<a id="canonical-0323012211100210-1023022020112033-2320013002321021-1100321203330013-1131132113132032-2330032031030300-0331202332332101-0012000002312311"></a>

## igroup_name property — netapp_backend_ontap_san / 222032030112 / 7

Type: `"string"`. Optional.

Name of the igroup for SAN volumes to use.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1330022100333103-3101332200331002-0122203133213232-3212332013102331-2011120002322033-1313323201323120-2011102010101103-0100101221131130"></a>

<a id="canonical-3023101030113332-1303213312003201-3020122001232221-2123321231210221-1213111001023222-0332200020101002-1012010211232202-2322230330301023"></a>

## labels property — netapp_backend_ontap_san / 222032030112 / 8

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-3213013313102103-2320201310113222-3202203321021231-1020023013300201-3112120113330011-2133203231003103-1322311122213301-3221222330122202"></a>

<a id="canonical-3011130331211031-2033122020230012-0323233131213020-3202323000133113-2330233130330201-1111212122331021-0013021031102130-2030320000332312"></a>

## limit_aggregate_usage property — netapp_backend_ontap_san / 222032030112 / 9

Type: `"number"`. Optional.

Fail provisioning if usage is above this percentage. Not enforced by default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
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
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-3010313201122300-3202230131233132-2212331211230132-2210332121321313-2032202333320130-3233230123032022-0233332110302303-3300111213032000"></a>

<a id="canonical-1021011101132012-3121311223320232-2302131213100322-2223330131321223-3131120220021323-0311102211211000-1301011332333132-0020210221010121"></a>

## limit_volume_size property — netapp_backend_ontap_san / 222032030112 / 10

Type: `"number"`. Optional.

Fail provisioning if requested volume size in GBi is above this value. Not enforced by default.

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

<a id="canonical-2120301100220321-2023011103303133-1001010232333030-0200232022003231-2032020030223303-2100332123202231-2032220311003033-0231132111111101"></a>

<a id="canonical-2302102203202111-1000300010331211-3101003111201301-3320000110023120-3133323012023223-2100101113220200-2300203121020221-0300303222022202"></a>

## management_lif_dns_name property — netapp_backend_ontap_san / 222032030112 / 11

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1022010312333230-3302111021021323-2102103012311311-3003102132231121-3203131103113013-0110221130021232-3223320231132030-0200320203312011"></a>

<a id="canonical-1002013110232313-2102230311130102-3012112023011003-2310011101120132-0101131310210121-0032000100123033-0111230222132111-1332003202311101"></a>

## management_lif_ip property — netapp_backend_ontap_san / 222032030112 / 12

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Upstream description:

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

- [no_chap](resources--voltstack_site--reference--group-007.md#canonical-3131300300003232-1213113312121110-2101132202112113-2222220200012310-3311313300020010-0323101303011132-1011320300200130-3121210131103020): complete subsection reference.

- [password](resources--voltstack_site--reference--group-007.md#canonical-2132312232120213-3221102302000033-2120222311312033-2100131220202112-0121322011003112-3330331203031231-2222003101331323-2131303022331231): complete subsection reference.

<a id="canonical-1332232022101322-0002211302130121-3101120011120001-0023103031123223-3332212031122100-0332231201210000-2301130330323032-1223232023210323"></a>

<a id="canonical-3301330030200000-2002010112100120-1233030011111211-2323223220010000-1001003113120320-0223322023112031-2111013010323120-1210332323102202"></a>

## region property — netapp_backend_ontap_san / 222032030112 / 13

Type: `"string"`. Optional.

Backend Region. Virtual Pool Region.

Upstream description:

Virtual Pool Region.

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

- [storage](resources--voltstack_site--reference--group-007.md#canonical-2021201323131230-0123032311133332-3132023020103310-1103300113210013-3332210303333201-3223303121220023-3222013232121103-1030113122101222): complete subsection reference.

<a id="canonical-0000301123213300-0113320202011101-3223312111033331-0333132001222031-1202101102230221-0331000211333311-2303033222122323-1030030020231032"></a>

<a id="canonical-0320101033123212-0021303023231333-2122032213300332-3020310302311310-3313232312211110-1001102333101233-2110223101311232-2010210221013122"></a>

## storage_driver_name property — netapp_backend_ontap_san / 222032030112 / 14

Type: `"string"`. Optional.

\[Enum: ontap-san|ontap-san-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-san\`, \`ontap-san-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ontap-san",
    "ontap-san-economy",
    "ontap-nas-flexgroup"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-san",
    "ontap-san-economy",
    "ontap-nas-flexgroup"
  ],
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-san\\\",\\\"ontap-san-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-san\\\",\\\"ontap-san-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="canonical-0333130100010121-0111102210011100-1110123323232111-3332303031021023-1330210112331332-2013332130312300-2320100311030102-0000122003113002"></a>

<a id="canonical-0123022220132122-3023030131110112-0123332212300231-2013032121323110-1021003110003331-3211230333213010-2202331011323230-3223310303330132"></a>

## storage_prefix property — netapp_backend_ontap_san / 222032030112 / 15

Type: `"string"`. Optional.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 80),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 80,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 80,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "80",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "80",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0302222000021211-0011310030033002-3210212233112122-0233211220122123-1302120221131321-3323310103112210-3012311312211133-3000221320210200"></a>

<a id="canonical-2233333102112213-1211131332333303-0010102333033110-2123322312323033-2002231233023220-0321101313023021-1011331112323201-1122322303033233"></a>

## svm property — netapp_backend_ontap_san / 222032030112 / 16

Type: `"string"`. Optional.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1330010033223003-1123301313222110-0001230030230310-2022303020303120-3231113212300012-1233213230313313-2222133320201111-1220122302021110"></a>

<a id="canonical-1303032222111110-1113123333011001-1112102031031110-2130222221110323-3223133203223223-0203030300303322-0322000320312031-3320303302221333"></a>

## trusted_ca_certificate property — netapp_backend_ontap_san / 222032030112 / 17

Type: `"string"`. Optional.

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Upstream description:

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth..

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
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
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [use_chap](resources--voltstack_site--reference--group-007.md#canonical-1302210300022100-0320123323331122-1013313122202331-2023320333011131-3331131330203113-2233112132012323-1330223010303130-1000302021033221): complete subsection reference.

<a id="canonical-0321012133322133-0202113303030020-0011311223331000-2101022300313230-3332030121311022-2331121220220303-2201310102001112-1230103102031231"></a>

<a id="canonical-0112112200110301-0022003301123331-0020331022101331-3021321003330030-3023210322300211-2131233212033033-0200131312201132-1103101323120003"></a>

## username property — netapp_backend_ontap_san / 222032030112 / 18

Type: `"string"`. Optional.

Username. Username to connect to the cluster/SVM.

Upstream description:

Username to connect to the cluster/SVM.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [volume_defaults](resources--voltstack_site--reference--group-007.md#canonical-1113012300123231-0230130213102123-2222113203222300-0231133123010033-2003312131131202-2321303021133030-3102013322103330-2102023110321222): complete subsection reference.

<a id="canonical-3012132100020111-3120211131203322-0211323201022131-1232233131202031-3330110203121013-2103103020123210-3133322003011032-3000001103113313"></a>

## Next pages — netapp_backend_ontap_san / 222032030112 / 19

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--voltstack_site--reference--group-007.md#canonical-2330200031312131-1002311220021131-0032301231013212-1100130321323122-2210200011122022-1300012033011022-1320010332203010-1131002203123322)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap](resources--voltstack_site--reference--group-007.md#canonical-3131300300003232-1213113312121110-2101132202112113-2222220200012310-3311313300020010-0323101303011132-1011320300200130-3121210131103020)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--voltstack_site--reference--group-007.md#canonical-2132312232120213-3221102302000033-2120222311312033-2100131220202112-0121322011003112-3330331203031231-2222003101331323-2131303022331231)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--voltstack_site--reference--group-007.md#canonical-2021201323131230-0123032311133332-3132023020103310-1103300113210013-3332210303333201-3223303121220023-3222013232121103-1030113122101222)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-1302210300022100-0320123323331122-1013313122202331-2023320333011131-3331131330203113-2233112132012323-1330223010303130-1000302021033221)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](resources--voltstack_site--reference--group-007.md#canonical-1113012300123231-0230130213102123-2222113203222300-0231133123010033-2003312131131202-2321303021133030-3102013322103330-2102023110321222)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2330200031312131-1002311220021131-0032301231013212-1100130321323122-2210200011122022-1300012033011022-1320010332203010-1131002203123322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001110030122212-2223312203332122-1310110221212203-1001102220300201-1231102122231210-1023021110101330-3120010303000223-3302203222032210"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key — client_private_key / 013111110222 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key

<a id="canonical-1002003112110332-1102221102302003-1000203221032112-0300221012113200-1010112032002210-1331231022202122-1222302100332310-1003231232003003"></a>

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
client_private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031223130130231-3003020212030202-0133211202200213-2301001332321312-1313321030323201-0103121303020002-2133031111312332-3330111132000201"></a>

## Direct properties — client_private_key / 013111110222 / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-3003212232102112-0131000030131313-0113102210223222-3213332333032132-2320122200033333-2101210311313122-2211102133003223-2303301001301011): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-1232030302212100-1030212110302011-3301330200212002-0031010230032033-3111230231212012-2303131230003220-3320230303103321-2000200221103102): complete subsection reference.

<a id="canonical-1123200113020001-2210331332001221-0100000321232232-1310223022113311-3100203221010132-3022322021011030-2132130001321332-3323102011010032"></a>

## Next pages — client_private_key / 013111110222 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-3003212232102112-0131000030131313-0113102210223222-3213332333032132-2320122200033333-2101210311313122-2211102133003223-2303301001301011)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-1232030302212100-1030212110302011-3301330200212002-0031010230032033-3111230231212012-2303131230003220-3320230303103321-2000200221103102)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3003212232102112-0131000030131313-0113102210223222-3213332333032132-2320122200033333-2101210311313122-2211102133003223-2303301001301011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303013000100030-1131202220112011-2302132130122200-2012210211112112-0331211220020020-0333131331122032-2211302133132123-3002131330030032"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info — blindfold_secret_info / 000103002031 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--voltstack_site--reference--group-007.md#canonical-2330200031312131-1002311220021131-0032301231013212-1100130321323122-2210200011122022-1300012033011022-1320010332203010-1131002203123322)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.blindfold_secret_info

<a id="canonical-0031332132322202-3000212122133001-2133302102311132-3130202013000202-2203332332131012-2120030230132232-0032323120312320-2322223323311122"></a>

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

<a id="canonical-1202103122113102-3300132123323330-1032231031121032-3302323002113011-3033201101213001-2213211133112220-1202132331201013-1213111101103320"></a>

## Direct properties — blindfold_secret_info / 000103002031 / 3

<a id="canonical-2132220020101001-3301131213100330-2310032320321012-1120113313022201-0301210223321321-3012210101011122-3211133122111323-3023323221113201"></a>

<a id="canonical-3131221233013320-3222033101003022-1312312220231231-3213000112230322-1320010200131331-0332201032201332-0121012203020331-0010121112222112"></a>

## decryption_provider property — blindfold_secret_info / 000103002031 / 4

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

<a id="canonical-0300032011100300-0301022123233332-1101320003313100-2300223111013132-2103331202233013-3222312232103012-2320300032220301-3112020130113122"></a>

<a id="canonical-2031212320300210-0110120130310130-2113331102212013-1003110220231111-1101300210321311-1332033212222023-1030223120311323-1020003313111101"></a>

## location property — blindfold_secret_info / 000103002031 / 5

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

<a id="canonical-0133203030113003-0223110233113103-3000323110311001-1322031003121130-2203111103032011-2010003201133111-2303330111101103-0000203201211113"></a>

<a id="canonical-1332222030320322-2123222330203101-2300002223322232-2233000202301303-2121103111311023-1203233122123020-1331232301130022-0200112022011302"></a>

## store_provider property — blindfold_secret_info / 000103002031 / 6

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

<a id="canonical-0113202322032002-3032221102321212-3323211121112302-0012203133202313-1033320130110111-3032301030203013-0012030011121223-0313303100130123"></a>

## Next pages — blindfold_secret_info / 000103002031 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--voltstack_site--reference--group-007.md#canonical-2330200031312131-1002311220021131-0032301231013212-1100130321323122-2210200011122022-1300012033011022-1320010332203010-1131002203123322)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1232030302212100-1030212110302011-3301330200212002-0031010230032033-3111230231212012-2303131230003220-3320230303103321-2000200221103102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113013312221002-3121110302301331-1113103032101223-2330320120201222-2213210232302231-2212120132103203-0130322102113112-1323320120110213"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info — clear_secret_info / 102121332321 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--voltstack_site--reference--group-007.md#canonical-2330200031312131-1002311220021131-0032301231013212-1100130321323122-2210200011122022-1300012033011022-1320010332203010-1131002203123322)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key.clear_secret_info

<a id="canonical-2322312321000103-0131120121330013-3222213010230001-2313122203220111-2012331031131223-0203200213302002-3210322102011031-3323130230220322"></a>

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

<a id="canonical-0323311113001231-2323211310000322-3012312121011211-3321201120300003-0201201221301020-0030300113211212-2002302231032313-3013203102030300"></a>

## Direct properties — clear_secret_info / 102121332321 / 3

<a id="canonical-3123310331211132-3103030020001232-3322303313002010-3120011001323123-1120013313002310-2310023000222123-3233132302132032-1133100310302203"></a>

<a id="canonical-2321220202020301-2331322330212210-1301121002111013-3012102010332220-0122303230110030-3130323320000320-2322113311212233-1333023321212111"></a>

## provider_ref property — clear_secret_info / 102121332321 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0320032311310211-0013010132220120-3010010013021130-3330313231023030-3233023022112112-2322032232133012-3003100220130100-2332120030112120"></a>

<a id="canonical-0223101011122033-2202232020313030-3002013102230131-0332233213330122-2012102101323333-1133023300030002-1302121312321012-3130023330120213"></a>

## URL property — clear_secret_info / 102121332321 / 5

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

<a id="canonical-0013132221001211-0321012001302031-3213120201302030-0221123132301210-1313000310132222-1030023112232013-1012330222020001-2310130012311230"></a>

## Next pages — clear_secret_info / 102121332321 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](resources--voltstack_site--reference--group-007.md#canonical-2330200031312131-1002311220021131-0032301231013212-1100130321323122-2210200011122022-1300012033011022-1320010332203010-1131002203123322)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3131300300003232-1213113312121110-2101132202112113-2222220200012310-3311313300020010-0323101303011132-1011320300200130-3121210131103020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312012130102202-0301301323003010-1023303000133100-2120210332131010-1202320332312211-2320331100202223-1300030232311000-0331101201202232"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap — no_chap / 111302111012 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap

<a id="canonical-0111223113202320-1130010322102301-0323110123031323-1332020211310011-1033002122101332-2030200311023102-1000003003311311-3120213002013212"></a>

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
no_chap = {}
```

<a id="canonical-0121320001222202-0030200202211210-2122311300002001-3110301332300310-2132331123231200-2323300033132201-1202313222313123-0331020321000000"></a>

## Direct properties — no_chap / 111302111012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221012310031211-2020103330120321-2322312033032323-0011301212103020-3311312033212301-1323213133130032-3233133303130301-1111030323232321"></a>

## Next pages — no_chap / 111302111012 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2132312232120213-3221102302000033-2120222311312033-2100131220202112-0121322011003112-3330331203031231-2222003101331323-2131303022331231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010200223131223-1311012012321212-3103110001332223-3100003221032202-1130023013112303-0223121102300102-2022100221201330-2000023021301331"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password — password / 312321312022 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password

<a id="canonical-2321311333130012-3101230112210220-0132032331220113-1303103030333323-0011232002020211-3011021000231310-3011303233032013-0301203101021132"></a>

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
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320213322101230-1120000110322130-3133330213110311-3103010121013201-0313132122100030-1212001112133230-2121031032221332-2232232210323020"></a>

## Direct properties — password / 312321312022 / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-0020210303001211-0000001331310210-2021031023303322-3210203111332000-1111233023131132-3031323100213313-0120303111102312-1032231200012311): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-3231112120030320-0213132310303121-3002120321202210-0232001021112031-2023132231333322-1102031030310302-0312332221210311-0220122230230123): complete subsection reference.

<a id="canonical-3221122002220123-3123010131233132-0122212101112133-1111323120211132-2010000012303201-1121030012020312-3202012232323301-3212300320212221"></a>

## Next pages — password / 312321312022 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-0020210303001211-0000001331310210-2021031023303322-3210203111332000-1111233023131132-3031323100213313-0120303111102312-1032231200012311)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-3231112120030320-0213132310303121-3002120321202210-0232001021112031-2023132231333322-1102031030310302-0312332221210311-0220122230230123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0020210303001211-0000001331310210-2021031023303322-3210203111332000-1111233023131132-3031323100213313-0120303111102312-1032231200012311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031003230023221-3223020213310122-2303300310330010-1122301231112112-0022023123323032-0330322012011303-1130003220111000-2011130020100203"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info — blindfold_secret_info / 113313333200 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--voltstack_site--reference--group-007.md#canonical-2132312232120213-3221102302000033-2120222311312033-2100131220202112-0121322011003112-3330331203031231-2222003101331323-2131303022331231)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.blindfold_secret_info

<a id="canonical-3122103031230312-2022012013010230-0000233030300022-2213333201112021-2333223011101033-2001232020231201-3323130320321320-3012223331333000"></a>

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

<a id="canonical-0001221212013023-3012023023330330-1203011000221300-3211221000110302-1112102010223301-1002220001113022-2222032322011023-1212232020200011"></a>

## Direct properties — blindfold_secret_info / 113313333200 / 3

<a id="canonical-2023030122210032-3132233010203001-1221102330102301-0101320022130113-0202231211222102-2303131023232321-0010103203123131-0103102112020013"></a>

<a id="canonical-2120213122221133-2203002003333312-1320203301213200-1232313332321323-1201021001101312-3303110130031133-0211022212100221-0233000122100132"></a>

## decryption_provider property — blindfold_secret_info / 113313333200 / 4

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

<a id="canonical-1000312300102210-2313201302003201-0213333012331011-1333112210210030-3100213300212123-1231210301003132-1303211232120120-1203120332031023"></a>

<a id="canonical-1013201100030331-2200010033233133-3022003300112000-1222103321113031-3311123031031032-1200112103102003-3231203230320020-1313311231103130"></a>

## location property — blindfold_secret_info / 113313333200 / 5

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

<a id="canonical-1020320211201101-1101331232002223-1021021132102202-1110322122313133-2101120222002202-1202203122210011-0002111013033330-0000001202120220"></a>

<a id="canonical-3232011330313220-1020030230200111-1213320232230121-2101223123112111-3112223030320030-1203002220320110-1033320310030302-3301200100201203"></a>

## store_provider property — blindfold_secret_info / 113313333200 / 6

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

<a id="canonical-3011322010022122-2212212323132101-2332002113301022-3321210233122112-3220111312232100-3311331203011032-2332200322230110-1102123311022113"></a>

## Next pages — blindfold_secret_info / 113313333200 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--voltstack_site--reference--group-007.md#canonical-2132312232120213-3221102302000033-2120222311312033-2100131220202112-0121322011003112-3330331203031231-2222003101331323-2131303022331231)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3231112120030320-0213132310303121-3002120321202210-0232001021112031-2023132231333322-1102031030310302-0312332221210311-0220122230230123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030030233212032-1311320123000313-0122120022110230-2301303320220001-1032132331032112-1033313013031023-1210203132303222-0213010300132202"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info — clear_secret_info / 310100201303 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--voltstack_site--reference--group-007.md#canonical-2132312232120213-3221102302000033-2120222311312033-2100131220202112-0121322011003112-3330331203031231-2222003101331323-2131303022331231)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password.clear_secret_info

<a id="canonical-1000313123030002-1130312230121130-1013133213103020-1023221301121213-2221032102301200-3100201121002320-2023130222131023-2010230212033001"></a>

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

<a id="canonical-1211310010203110-3331133232001212-2300132200011210-0110020233002303-0022003100220102-3323012012032101-3112020221211033-3231323323302223"></a>

## Direct properties — clear_secret_info / 310100201303 / 3

<a id="canonical-2310313233313221-0110031303020000-3331010030112000-3222023223221202-0331032321032223-2202033210100032-1110322033231232-0322332310121233"></a>

<a id="canonical-3133310110132113-3211212203303333-3132030112310011-1203011211323010-2303122111211012-0211230102320301-0310022222032103-3301033020100202"></a>

## provider_ref property — clear_secret_info / 310100201303 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0232021103231121-2023012311130023-0323132112310000-3220223010113133-0021331203133111-3232203120210321-0101023330213230-0211133103110203"></a>

<a id="canonical-1013300030030030-3322022323130303-0003030120122111-2213022222322332-2210021333023202-2323320020012130-0013001010033033-0120010203100202"></a>

## URL property — clear_secret_info / 310100201303 / 5

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

<a id="canonical-3123112221200113-0323300211001013-1213311012022033-1203012310120002-0221201320332320-3012312020131100-2230310000113202-2113002112110220"></a>

## Next pages — clear_secret_info / 310100201303 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](resources--voltstack_site--reference--group-007.md#canonical-2132312232120213-3221102302000033-2120222311312033-2100131220202112-0121322011003112-3330331203031231-2222003101331323-2131303022331231)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2021201323131230-0123032311133332-3132023020103310-1103300113210013-3332210303333201-3223303121220023-3222013232121103-1030113122101222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130102300122302-3300130301123202-3220000102222211-1312011300322031-1032303202111203-1130102322212231-0132000133001001-2221223310210302"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage — storage / 220201031001 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage

<a id="canonical-3023233132101222-3313033133123231-0302213023302223-2021231123131032-2221120112120101-1030121322202222-2233312100230012-2110100331130213"></a>

Type: `"object"`. list nested block, Optional.

List of Virtual Storage Pool definitions which are referred back by Storage Class label match
selection.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010333110210331-0202233023112100-2030332023212112-0112013223130332-2102312202120033-1313200332003320-1101113010103113-1221003111013301"></a>

## Direct properties — storage / 220201031001 / 3

<a id="canonical-3301001211011221-2203233023032230-0333211231122130-3101020131202201-2222021331323211-0001200121323202-2133011000321203-2030203223103301"></a>

<a id="canonical-2011013131011321-2333333222233033-1033211231323301-3130303213103221-3221110130131033-1131132333101330-0022331113311100-2030022011303230"></a>

## labels property — storage / 220201031001 / 4

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class label match
selection.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [volume_defaults](resources--voltstack_site--reference--group-007.md#canonical-1202210112233333-1101003113022321-0303331021102131-1303210211011302-0230210030131133-3102221313032020-3003000301311113-0233310333213002): complete subsection reference.

<a id="canonical-3201123020322111-3300303303302221-2233102200003120-1003313223013032-2102002230101010-2211031111133033-3301223212132212-2033323101002113"></a>

<a id="canonical-3313300330020303-0231202323112000-1100231031000022-3111202012230210-0033001131103323-3112000203120320-2303332131133312-0032113032322023"></a>

## zone property — storage / 220201031001 / 5

Type: `"string"`. Optional.

Virtual Pool Zone. Virtual Storage Pool zone definition.

Upstream description:

Virtual Storage Pool zone definition.

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

<a id="canonical-2232021202133010-3013231102303031-0200202033130322-0111222230201331-0221320100121101-3133111120222012-2002211000212120-0203112300111311"></a>

## Next pages — storage / 220201031001 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](resources--voltstack_site--reference--group-007.md#canonical-1202210112233333-1101003113022321-0303331021102131-1303210211011302-0230210030131133-3102221313032020-3003000301311113-0233310333213002)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1202210112233333-1101003113022321-0303331021102131-1303210211011302-0230210030131133-3102221313032020-3003000301311113-0233310333213002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131013023000113-0330230003033211-0103030213231131-1003332111133020-0030210133100102-3020230223001311-3232331133212332-0010021102300231"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults — volume_defaults / 223210013013 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--voltstack_site--reference--group-007.md#canonical-2021201323131230-0123032311133332-3132023020103310-1103300113210013-3332210303333201-3223303121220023-3222013232121103-1030113122101222)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults

<a id="canonical-2232022223320310-3220303111321302-2231311122231022-2311211221232020-1221031301312202-3311232113213231-2313030323012300-2333310213111320"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
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
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210100113132333-1103113312132020-1331033310212021-3020321212103121-3230220003222301-0100121301031011-0002222202301122-0212231111012003"></a>

## Direct properties — volume_defaults / 223210013013 / 3

<a id="canonical-0312201103010122-2331033122103211-0133311213031012-1212031101011211-0120130313233012-3132002121232322-2231201320030101-0003300023020130"></a>

<a id="canonical-1221032312231213-3101020210301222-3012210331233323-2123110123122130-3310101111011133-3133300103023003-1220221120030230-2303130021203230"></a>

## adaptive_qos_policy property — volume_defaults / 223210013013 / 4

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1100221322210300-0303100100022010-0232301201020131-1110302232320303-1332321333131100-0313002102321023-1132001230333111-3231112231010112"></a>

<a id="canonical-1000022202220003-3221222010112032-3312322102122231-0000221100203222-2230110303311130-1202111023203000-0313213032120122-1230110023313223"></a>

## encryption property — volume_defaults / 223210013013 / 5

Type: `"bool"`. Optional.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

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

<a id="canonical-0030030021222203-0333003321103123-3201210303222101-0010002231300020-3312301003112231-1233101130202201-3331213203232303-3331032310222213"></a>

<a id="canonical-3010212302200013-0123102220301120-3303211331322332-1311221322321320-3301311322231032-1212220001013111-1100122223023031-0312322010330122"></a>

## export_policy property — volume_defaults / 223210013013 / 6

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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

- [no_qos](resources--voltstack_site--reference--group-007.md#canonical-3213130222121212-1003211222103221-3012310103231102-1113023333323033-0200020100003233-2031000311331112-2032321011133213-0322200311122102): complete subsection reference.

<a id="canonical-2211000110232032-2011301300133123-2131031131323212-3300011331303023-3012211120230223-3103101010313133-2122310303331223-1330321222331033"></a>

<a id="canonical-2131211230310223-1033031000131011-3331130303222220-1002122103333213-1313001313033322-2012120212313001-2313221130303311-2302022322103220"></a>

## qos_policy property — volume_defaults / 223210013013 / 7

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0020123221001301-0211303333013011-2202101300021331-1012221032020223-3101110012011221-1021202213100221-2022030301111110-1322011323020301"></a>

<a id="canonical-2131312221132301-1103220111312110-3112100320133302-2131122123320331-3031312030321010-0233102033113333-3113013213300230-3002302133321023"></a>

## security_style property — volume_defaults / 223210013013 / 8

Type: `"string"`. Optional.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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

<a id="canonical-3200033131332303-1210310212100200-0001031011301202-2302030033201230-0032201003000321-3232312102010321-0221123020203103-3213202223021131"></a>

<a id="canonical-2220012223231200-1211132313320312-3031020332032031-0101203323303033-2211201031201120-0310003222121331-2101032230313032-1200022313311322"></a>

## snapshot_dir property — volume_defaults / 223210013013 / 9

Type: `"bool"`. Optional.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

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

<a id="canonical-1112233130302111-3113323220222111-2221212330333010-2132211021032033-2201322212311230-0233200203023133-2130211133121020-1012230120332211"></a>

<a id="canonical-1203222231223213-3201201132312320-2323002022023131-2210203111232301-1231220130212000-0313331320323303-0001133203022101-2223022010230202"></a>

## snapshot_policy property — volume_defaults / 223210013013 / 10

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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

<a id="canonical-1323000231323112-0230001200122212-2230003233011133-3200323313023012-1323211102130121-3013200100001103-3031000022022301-0332223001111302"></a>

<a id="canonical-3201211302112213-0301133220223100-3000200021332102-2033233003120112-1021203032100002-0212120123212221-0113003012021032-0302220210303023"></a>

## snapshot_reserve property — volume_defaults / 223210013013 / 11

Type: `"string"`. Optional.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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

<a id="canonical-2031033002122011-2300200230123012-2300220031102203-1121132332002321-2233031300030310-2331013231230111-3123320331112020-0322203300223330"></a>

<a id="canonical-3333302021133203-1210202330030000-3232110211010022-0333020022132130-3232000120330323-2230020201020202-0103333303203311-0220122000312322"></a>

## space_reserve property — volume_defaults / 223210013013 / 12

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-1300130231100011-2323032203301122-1132010023120201-0211211213123010-2020311220301023-1020112222203322-1032130310202233-2103322030211010"></a>

<a id="canonical-0031231001033223-0011210202133312-1023132320113131-2111332123333310-3333030011233000-3102332122333023-3012333031033300-3122130223332001"></a>

## split_on_clone property — volume_defaults / 223210013013 / 13

Type: `"bool"`. Optional.

Split a clone from its parent upon creation.

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

<a id="canonical-1312132101131002-0220223301023102-2202110102202210-1333220311031120-1312323233222201-2202000011320331-0323222302123302-2320000230122311"></a>

<a id="canonical-2110131311013303-0203212100033311-2210102012031112-0211220130100221-3230031121123030-1212230110333221-2031032103013101-3131101312301332"></a>

## tiering_policy property — volume_defaults / 223210013013 / 14

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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

<a id="canonical-0203203023201311-2113230333221103-1222203000301133-1220210210211012-3312320302002131-3213233220031002-3320032033011313-2301311130203302"></a>

<a id="canonical-2122311133212333-3112013232323022-2333021311123232-2212331220331011-3331023223313121-2113322010220333-2302301211103311-3321302320320120"></a>

## unix_permissions property — volume_defaults / 223210013013 / 15

Type: `"number"`. Optional.

Unix permission mode for new volumes. All allowed 777.

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

<a id="canonical-1111222110132112-3333020221312133-0021102203332333-0221232100131002-1203220310211123-2302021312133122-1011213000313221-1000331213303233"></a>

## Next pages — volume_defaults / 223210013013 / 16

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos](resources--voltstack_site--reference--group-007.md#canonical-3213130222121212-1003211222103221-3012310103231102-1113023333323033-0200020100003233-2031000311331112-2032321011133213-0322200311122102)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--voltstack_site--reference--group-007.md#canonical-2021201323131230-0123032311133332-3132023020103310-1103300113210013-3332210303333201-3223303121220023-3222013232121103-1030113122101222)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3213130222121212-1003211222103221-3012310103231102-1113023333323033-0200020100003233-2031000311331112-2032321011133213-0322200311122102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100132110021210-2313203001133033-2213122221113320-0201131021210101-1030002033022121-2130020103303020-1330101332023103-0331313001323310"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos — no_qos / 033331013212 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](resources--voltstack_site--reference--group-007.md#canonical-2021201323131230-0123032311133332-3132023020103310-1103300113210013-3332210303333201-3223303121220023-3222013232121103-1030113122101222)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](resources--voltstack_site--reference--group-007.md#canonical-1202210112233333-1101003113022321-0303331021102131-1303210211011302-0230210030131133-3102221313032020-3003000301311113-0233310333213002)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults.no_qos

<a id="canonical-0333201031323131-3013003323111312-2100033121222232-3220000031101121-1221323003113021-1223311232230020-0321120021130303-0130022032121331"></a>

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
no_qos = {}
```

<a id="canonical-1222300321300213-0201113120311302-1233332230231202-2313221032233001-3331111102200310-0103213023313310-1102200021303321-3133303220331112"></a>

## Direct properties — no_qos / 033331013212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011001311101032-0031230032222312-3100303221321211-3111032320130233-0031221032331112-0220210212001110-3013102321223320-3201012022032211"></a>

## Next pages — no_qos / 033331013212 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage.volume_defaults](resources--voltstack_site--reference--group-007.md#canonical-1202210112233333-1101003113022321-0303331021102131-1303210211011302-0230210030131133-3102221313032020-3003000301311113-0233310333213002)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1302210300022100-0320123323331122-1013313122202331-2023320333011131-3331131330203113-2233112132012323-1330223010303130-1000302021033221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100033213000113-2312010113010322-0011222112211301-2211123231021300-2333031312330012-0213220231302322-0333203300131210-1300233110203211"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap — use_chap / 220301133221 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap

<a id="canonical-0220011213021123-3111330013002003-0322030220122021-2313003110020200-0000033122332222-2321112113113203-1101220212022020-1310233233333333"></a>

Type: `"object"`. single nested block, Optional.

Device NetApp Backend ONTAP SAN CHAP configuration OPTIONS for enabled CHAP.

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
use_chap {
  # Configure direct properties listed below.
}
```

<a id="canonical-3331231130021222-2220320121203331-2021200302310201-0322111321321212-2302322100010332-0332110200211003-0030000230220213-2122121210302010"></a>

## Direct properties — use_chap / 220301133221 / 3

- [chap_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-1202332322330233-3110112302231232-3321232211312230-3222103310300001-0302213331231203-0311301133312330-2001323300100231-2003233033232332): complete subsection reference.

- [chap_target_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-2000300310112033-3030231233323110-3323023330302310-3233101212130303-0102130230101200-0200331031123300-2321222020002201-0321302101301323): complete subsection reference.

<a id="canonical-0130110133312103-2303201322313102-3123110120223320-2012132001230200-0332002312033320-3312221133130112-1231311222012232-2222130330100212"></a>

<a id="canonical-2032000221222220-0231322200121313-0321203103202333-2330112001323103-2201323312110111-2312311020131033-2232323201313200-0122231000200312"></a>

## chap_target_username property — use_chap / 220301133221 / 4

Type: `"string"`. Optional.

Target username. Required if useCHAP=true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3113010321221230-1103103320223133-0211303110233313-3001222013212222-1003130103300033-3333233003112011-3021220203132310-0131122323212002"></a>

<a id="canonical-2023322020303223-2200333300201313-1100123120212133-2030030313001222-0312012222121121-3030110313332020-2202031313311121-3311220201331300"></a>

## chap_username property — use_chap / 220301133221 / 5

Type: `"string"`. Optional.

Inbound username. Required if useCHAP=true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1231303010003032-3320233020132010-0010111010101001-2210102232113231-2023302133103123-1001220300220011-1302212303121113-0001322313213111"></a>

## Next pages — use_chap / 220301133221 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-1202332322330233-3110112302231232-3321232211312230-3222103310300001-0302213331231203-0311301133312330-2001323300100231-2003233033232332)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-2000300310112033-3030231233323110-3323023330302310-3233101212130303-0102130230101200-0200331031123300-2321222020002201-0321302101301323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1202332322330233-3110112302231232-3321232211312230-3222103310300001-0302213331231203-0311301133312330-2001323300100231-2003233033232332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223320202111110-0013102001002233-2021030001120132-0233102010101120-1031101332330231-0101210203121231-0022011000222232-2020212032011202"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret — chap_initiator_secret / 001103211230 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-1302210300022100-0320123323331122-1013313122202331-2023320333011131-3331131330203113-2233112132012323-1330223010303130-1000302021033221)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret

<a id="canonical-2213333212232001-0210100200330201-0221312333011033-1200021222212122-3311230200010202-2212311100100211-1012120200202331-1101202310301102"></a>

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
chap_initiator_secret {
  # Configure direct properties listed below.
}
```

<a id="canonical-2331230111302200-2220313303331212-0132220202111313-3303112230031000-2232301021001320-2213131320312031-2203332022231221-2202000203013011"></a>

## Direct properties — chap_initiator_secret / 001103211230 / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-2013122211330210-0001003102000003-3233020311210300-2200321132212310-3012102103032333-0033211132113210-2232222203331112-3121223133030100): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-0031320311323220-0120313100221101-3310111313333333-1313330002102022-0313302313203000-3202000123323331-0123223130021101-1031221131312023): complete subsection reference.

<a id="canonical-2213010312202122-2300211302101003-1023120220322202-1221112312320110-3031300201111111-0013320200000311-0102212230232211-0233113333300202"></a>

## Next pages — chap_initiator_secret / 001103211230 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-2013122211330210-0001003102000003-3233020311210300-2200321132212310-3012102103032333-0033211132113210-2232222203331112-3121223133030100)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-0031320311323220-0120313100221101-3310111313333333-1313330002102022-0313302313203000-3202000123323331-0123223130021101-1031221131312023)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-1302210300022100-0320123323331122-1013313122202331-2023320333011131-3331131330203113-2233112132012323-1330223010303130-1000302021033221)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2013122211330210-0001003102000003-3233020311210300-2200321132212310-3012102103032333-0033211132113210-2232222203331112-3121223133030100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200022021310101-2023101012120313-0323023311112032-3231003313321213-0313001012013323-1122212320222100-3002311111330330-3302311212123111"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info — blindfold_secret_info / 230333022233 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-1302210300022100-0320123323331122-1013313122202331-2023320333011131-3331131330203113-2233112132012323-1330223010303130-1000302021033221)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-1202332322330233-3110112302231232-3321232211312230-3222103310300001-0302213331231203-0311301133312330-2001323300100231-2003233033232332)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.blindfold_secret_info

<a id="canonical-3303321313020022-3113201210120211-3203230221022122-3213201112102122-1112103011201201-3120112131110033-2233001231020301-3322101221211332"></a>

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

<a id="canonical-3021002312321101-2021300113001100-1112010131100201-2311122312031333-3111330113021323-3320203011322333-0120030300121001-0102020100112102"></a>

## Direct properties — blindfold_secret_info / 230333022233 / 3

<a id="canonical-3010121031211233-2322303001301131-2203200123321022-0310211012323323-0133133301103302-2123323011312323-2132111030301111-1210020302302300"></a>

<a id="canonical-3023300312003221-0302213113020003-1011122300020132-3322222130220001-3121112221103232-2323201112303202-3313222012310002-0000232300303212"></a>

## decryption_provider property — blindfold_secret_info / 230333022233 / 4

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

<a id="canonical-0132300202131001-0111113032133320-1311120031103012-0101301001030003-1311102223332320-3133021320120131-0010321333112313-1232010032112332"></a>

<a id="canonical-0221232133013330-1201130202302331-0302320303002103-2322312301123033-0030102030101320-1200311202330002-3321222011320213-2300212311110113"></a>

## location property — blindfold_secret_info / 230333022233 / 5

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

<a id="canonical-3203020210302233-2021133220200321-2010301230310020-3303301230031021-3333301030033020-2030133213023020-3013330032033021-0300120313033330"></a>

<a id="canonical-0123113021333202-0133202012323332-0102311213001322-3031321030313313-0312330310310031-0002120233312132-2303220111001221-0031121320233303"></a>

## store_provider property — blindfold_secret_info / 230333022233 / 6

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

<a id="canonical-1011310203210032-2333130100313020-0231123202111332-0320213130200133-1223312330011010-1031230001130231-0130333313120312-1202133232133001"></a>

## Next pages — blindfold_secret_info / 230333022233 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-1202332322330233-3110112302231232-3321232211312230-3222103310300001-0302213331231203-0311301133312330-2001323300100231-2003233033232332)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0031320311323220-0120313100221101-3310111313333333-1313330002102022-0313302313203000-3202000123323331-0123223130021101-1031221131312023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333102101102303-3100022233211300-2232112120001300-1222303131131312-0132011200020313-3203113323101103-1133321003003103-3012212311101130"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info — clear_secret_info / 212032133031 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-1302210300022100-0320123323331122-1013313122202331-2023320333011131-3331131330203113-2233112132012323-1330223010303130-1000302021033221)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-1202332322330233-3110112302231232-3321232211312230-3222103310300001-0302213331231203-0311301133312330-2001323300100231-2003233033232332)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret.clear_secret_info

<a id="canonical-3303310311000200-2201332300312300-1232203011231321-2001001031331001-1022230300133213-0032010123311223-0333121031230230-1100012312120203"></a>

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

<a id="canonical-0213300021133131-3012023022302302-3031112303233320-0201113123223310-1020302233323003-3201011101203331-3102300031131011-0213210033000013"></a>

## Direct properties — clear_secret_info / 212032133031 / 3

<a id="canonical-2131212231211223-2101300032321320-1001213031323122-1212303131202111-3302200020212213-0010022303021321-0322320130121100-0120021211231231"></a>

<a id="canonical-2220022230302331-2021333221300031-3203120031023133-3211203311122233-2323011031333030-3310101113310123-0022213020333020-3111300012021332"></a>

## provider_ref property — clear_secret_info / 212032133031 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3122123201030200-2022123201210221-3003113300330101-3232102322131010-0013312331313310-1112122113002120-3021020332103101-3120312132200313"></a>

<a id="canonical-2221333100132321-2212001330310111-0332333220311112-3211100113230323-3113320130120001-3120323300023110-0022200331110123-1320030310113003"></a>

## URL property — clear_secret_info / 212032133031 / 5

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

<a id="canonical-3311120133210201-2023010221200021-3210321331301112-3122111313011102-1122323012131100-2112230002231223-1331200233120211-2011030231211023"></a>

## Next pages — clear_secret_info / 212032133031 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-1202332322330233-3110112302231232-3321232211312230-3222103310300001-0302213331231203-0311301133312330-2001323300100231-2003233033232332)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2000300310112033-3030231233323110-3323023330302310-3233101212130303-0102130230101200-0200331031123300-2321222020002201-0321302101301323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000212333322001-3223033113212112-2120303221101100-1312321111331121-2200300201310130-3300022033010020-1110303102212131-3100331010031101"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret — chap_target_initiator_secret / 002012331200 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-1302210300022100-0320123323331122-1013313122202331-2023320333011131-3331131330203113-2233112132012323-1330223010303130-1000302021033221)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret

<a id="canonical-3231222102221232-0130212122131220-1102222322321331-1223311321331010-0310013031220312-1002232012232323-1012101020100201-1203022301301003"></a>

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
chap_target_initiator_secret {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210130022313311-1332013200311303-1232321001231120-3221300031013020-0021303121223233-2111003030013230-0223131022130232-1110021332330310"></a>

## Direct properties — chap_target_initiator_secret / 002012331200 / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-2032021202223020-2112131331311121-1121102231121123-1022130102200102-2313323113222001-0020322102301121-0020013331201332-2222303213010210): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-0232333221311231-2320020122312203-3021311321323313-0103112101003221-2200220231020101-1130201022212111-0110212211200031-0221012233302312): complete subsection reference.

<a id="canonical-3230033132201010-2233220210202010-0132022312112011-0232310312230211-1130230012021210-3212131010011123-2102233312123110-0220103133030100"></a>

## Next pages — chap_target_initiator_secret / 002012331200 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-2032021202223020-2112131331311121-1121102231121123-1022130102200102-2313323113222001-0020322102301121-0020013331201332-2222303213010210)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-0232333221311231-2320020122312203-3021311321323313-0103112101003221-2200220231020101-1130201022212111-0110212211200031-0221012233302312)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-1302210300022100-0320123323331122-1013313122202331-2023320333011131-3331131330203113-2233112132012323-1330223010303130-1000302021033221)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2032021202223020-2112131331311121-1121102231121123-1022130102200102-2313323113222001-0020322102301121-0020013331201332-2222303213010210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133001011132323-1232123120132321-2010220111003333-0333232012200230-3221020321222100-2111201023020333-1132201103121333-2303022130102230"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info — blindfold_secret_info / 003001022213 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-1302210300022100-0320123323331122-1013313122202331-2023320333011131-3331131330203113-2233112132012323-1330223010303130-1000302021033221)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-2000300310112033-3030231233323110-3323023330302310-3233101212130303-0102130230101200-0200331031123300-2321222020002201-0321302101301323)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.blindfold_secret_info

<a id="canonical-3203330323101132-1121211323210021-2321302231030222-1212020201210331-3330210322021220-0020101131011312-0130023330111100-0321302021212113"></a>

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

<a id="canonical-3312033321001332-1133123011311120-0232030000023231-3201102121320303-2020313212103303-0102233122221020-3233130132100333-1122311100233302"></a>

## Direct properties — blindfold_secret_info / 003001022213 / 3

<a id="canonical-3303022211100332-1203020123031231-3310210011300330-1022110302323132-2033030310003222-0023113113302023-0133010122301033-2321133023100213"></a>

<a id="canonical-3030332030330203-0223113323010212-2120033213011130-3203203300200303-0030321120203033-0303011130101303-1213222302303302-2111321332223332"></a>

## decryption_provider property — blindfold_secret_info / 003001022213 / 4

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

<a id="canonical-3203231020311120-2202211020012310-3101211113000113-1310302203332302-2330011012012111-0021010011111000-2310211031112221-2101212321331223"></a>

<a id="canonical-2013133113211210-3322303311330322-0312120331233113-0101230311130333-1223221230320101-3303121310011220-1123220321320012-3203300232321113"></a>

## location property — blindfold_secret_info / 003001022213 / 5

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

<a id="canonical-1331320100102013-3331032201103202-3332300221133031-0001222230330312-3030300230300230-3232331130110012-0020013201121021-2312223121333033"></a>

<a id="canonical-2300112102231323-2013123301230233-1110201002032332-1223102202032021-0212300123100130-1010031010202301-3131311331213012-1103330323123202"></a>

## store_provider property — blindfold_secret_info / 003001022213 / 6

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

<a id="canonical-2313210121110112-3112022132332301-3222302111233313-3202300133121321-1222101113022332-1113311222210321-0030121303322030-0222120200031001"></a>

## Next pages — blindfold_secret_info / 003001022213 / 7

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-2000300310112033-3030231233323110-3323023330302310-3233101212130303-0102130230101200-0200331031123300-2321222020002201-0321302101301323)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0232333221311231-2320020122312203-3021311321323313-0103112101003221-2200220231020101-1130201022212111-0110212211200031-0221012233302312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033200111230310-1121011332030013-0233100101322210-0213003101231203-1222130330302023-2031203131320202-2012301232310312-2002310333311131"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info — clear_secret_info / 113223012110 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](resources--voltstack_site--reference--group-007.md#canonical-1302210300022100-0320123323331122-1013313122202331-2023320333011131-3331131330203113-2233112132012323-1330223010303130-1000302021033221)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-2000300310112033-3030231233323110-3323023330302310-3233101212130303-0102130230101200-0200331031123300-2321222020002201-0321302101301323)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret.clear_secret_info

<a id="canonical-1210221101002330-3311311112010213-3211212033321322-2103201231021221-1103002002322303-3123332233010003-2310331021313200-3311211012200230"></a>

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

<a id="canonical-0001322032331110-3013213233223010-2332113100201102-1031311223212032-2330112333002203-0132213330120102-3223021210031221-3202013012001113"></a>

## Direct properties — clear_secret_info / 113223012110 / 3

<a id="canonical-3121222323013210-0211230123112032-0023303010230213-2303230002300130-3133211333302312-2003032103022330-2133113201210000-3203323000102231"></a>

<a id="canonical-0112223031101223-3002210312010313-3331210030322320-0031303020221212-0302231022330100-2111130223131221-2223102211033222-0013301111023312"></a>

## provider_ref property — clear_secret_info / 113223012110 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0300123332032130-2110010321122221-0331110120300113-3312113002320123-0022011321013002-1013132331110132-3022212332220232-0312130313021331"></a>

<a id="canonical-3212033302201101-3010023033223020-2322030001033333-1112232202312023-3121302232231131-1213231223313013-1133020003100302-2320113213110131"></a>

## URL property — clear_secret_info / 113223012110 / 5

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

<a id="canonical-1312021220311103-0303211303313303-3312200002230313-3201110023313210-0031301330302032-0111313210122230-0230112202313203-3001311032113233"></a>

## Next pages — clear_secret_info / 113223012110 / 6

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap.chap_target_initiator_secret](resources--voltstack_site--reference--group-007.md#canonical-2000300310112033-3030231233323110-3323023330302310-3233101212130303-0102130230101200-0200331031123300-2321222020002201-0321302101301323)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1113012300123231-0230130213102123-2222113203222300-0231133123010033-2003312131131202-2321303021133030-3102013322103330-2102023110321222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323323223021000-1303203230331312-3300313312113211-1223021322231101-1230321021213130-3120001202111320-3021031332021012-0121102321013030"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults — volume_defaults / 302311301001 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults

<a id="canonical-3010030102031323-1030221113212023-3010130210102321-3203232322121102-1110230121000113-0101033210023030-1301100300211133-1121111233321020"></a>

Type: `"object"`. single nested block, Optional.

It controls how each volume is provisioned by default using these OPTIONS in a special section of
the configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "no_qos"),
  validators.ConflictingObjectAttributes("adaptive_qos_policy",
    "qos_policy"),
  validators.ConflictingObjectAttributes("no_qos",
    "qos_policy")}
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
  "x-ves-oneof-field-qos_policy_choice": "[\"adaptive_qos_policy\",\"no_qos\",\"qos_policy\"]"
}
```

Terraform syntax:

```terraform
volume_defaults {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330000332203130-2103230222021023-1313223302312222-0023201003320223-2322112201023333-0231103003321212-2230212312113302-0321330100321111"></a>

## Direct properties — volume_defaults / 302311301001 / 3

<a id="canonical-0322012030331102-3211302212012132-1132101303021030-3301010223303100-3001100200202000-0233230021200230-1032130023321301-2121011203223121"></a>

<a id="canonical-0030323320130300-0130030212322200-0221101320120031-1220002111102331-1232020032333312-0113122201101323-0311201321122000-2030330010020222"></a>

## adaptive_qos_policy property — volume_defaults / 302311301001 / 4

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[no\_qos qos\_policy\] Enter Adaptive QoS Policy Name.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0130101311212010-1022102313303302-0101330211233112-0133112230333000-0330221211202310-3001203010212202-2221003033331130-3101212003110312"></a>

<a id="canonical-1303130111030311-0130231133231323-0002201010000120-2221000002113003-3020100111233031-0103012033032200-0022222022201211-3201030122302013"></a>

## encryption property — volume_defaults / 302311301001 / 5

Type: `"bool"`. Optional.

Enable Encryption. Enable NetApp volume encryption.

Upstream description:

Enable NetApp volume encryption.

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

<a id="canonical-3010201131013332-1300023000331231-0122130330113311-0022020203320300-0210223002213002-0322303031101203-0233110200000030-3020133231131031"></a>

<a id="canonical-3102331131003330-1202210203322132-0320211210032202-1201220000211003-0122111310333110-1032311221303110-3203010200131010-3313100032100111"></a>

## export_policy property — volume_defaults / 302311301001 / 6

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Export policy to use.

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

- [no_qos](resources--voltstack_site--reference--group-007.md#canonical-2032231111203001-0101021010111211-3210232311311103-1102113222303221-0132303120120233-0213230122230003-0221102003103200-2132330121202331): complete subsection reference.

<a id="canonical-3321300121323110-0203322132122312-1012110220103223-3211311102120321-0031220023230102-3303131331231322-2112220200303030-2312331300220030"></a>

<a id="canonical-2113133022231222-2121021031321002-1211110120222111-1020333300200233-2321101102331310-1002310123211020-2221033233311012-2212211221321213"></a>

## qos_policy property — volume_defaults / 302311301001 / 7

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Exclusive with \[adaptive\_qos\_policy no\_qos\] Enter QoS Policy Name.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2312332331311231-1312131100112020-2233223201221013-3012121030323123-0320100021302232-1333332000232033-1111310320223031-3311210013032110"></a>

<a id="canonical-0120011102113231-0213110201212323-3232102000221112-1021111120211101-1202022130311011-2022303301302221-2103203131220102-1012001010020120"></a>

## security_style property — volume_defaults / 302311301001 / 8

Type: `"string"`. Optional.

Security Style. Security style for new volumes.

Upstream description:

Security style for new volumes.

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

<a id="canonical-2303321123000313-1220222102231321-2131312132221020-1033201212200113-1203310222012010-2011320203030332-2023031103201010-2031012030123300"></a>

<a id="canonical-3001311202320300-0230031130323223-2222231003013030-3320333231132321-3031101213002330-3133303330010131-2211023230320332-1212013021002000"></a>

## snapshot_dir property — volume_defaults / 302311301001 / 9

Type: `"bool"`. Optional.

Access to Snapshot Directory. Access to the .snapshot directory.

Upstream description:

Access to the .snapshot directory.

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

<a id="canonical-3100320013202021-1120122021120220-0001032120331101-1320200013220103-0110031202033201-1222332000221112-2301000121023222-3310201030001310"></a>

<a id="canonical-3331031202313030-1130113133323233-1331202130100010-1101101231101013-2102200101132312-0121003001300010-0213321101302222-1332101120001112"></a>

## snapshot_policy property — volume_defaults / 302311301001 / 10

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Snapshot policy to use.

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

<a id="canonical-3323130111121300-2320333213313333-1011322031321113-1003120331122310-2310312111222232-3133212001130010-1203100013222220-2310302202231202"></a>

<a id="canonical-3301312000302210-3310023211223221-1222303300213103-2321223031310213-3002303033000320-3120303223330330-3010321033122220-0333220330220023"></a>

## snapshot_reserve property — volume_defaults / 302311301001 / 11

Type: `"string"`. Optional.

Percentage of volume reserved for snapshots. '0' if snapshot policy is 'none', else ''.

Upstream description:

Percentage of volume reserved for snapshots. "0" if snapshot policy is "none", else ""

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

<a id="canonical-0313322110002212-2322310220002323-1131230203332011-3031000011111331-1201110312303022-2100120203111220-2112200131123302-1110211011311123"></a>

<a id="canonical-2101112103201332-1232132310220220-2021331200322333-1013302211212021-2231012113300330-1220313020322111-2331332133323001-1123023012213210"></a>

## space_reserve property — volume_defaults / 302311301001 / 12

Type: `"string"`. Optional.

\[Enum: none|thick\] Space reservation mode; “none” (thin) or “volume” (thick). Possible values are
\`none\`, \`thick\`.

Upstream description:

Space reservation mode; “none” (thin) or “volume” (thick)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("none",
    "thick"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "none",
    "thick"
  ],
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"none\\\",\\\"thick\\\"]"
  }
}
```

<a id="canonical-3010312330311130-3212312021231123-3300312211212211-1312302201320200-0222330330010200-3212313112220203-1103120003020313-2323222211310112"></a>

<a id="canonical-1122112200202202-2012301232231013-0120211032221303-3221222010003323-2300231302311110-0100233011023301-0233211032100231-1211001211312103"></a>

## split_on_clone property — volume_defaults / 302311301001 / 13

Type: `"bool"`. Optional.

Split a clone from its parent upon creation.

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

<a id="canonical-0311333233110202-1101120220021202-0212233122102323-0031111233330201-0330033203000121-0203122031101202-3000101220032032-1030331112132000"></a>

<a id="canonical-0001010311201022-2023133212032311-2221233313112021-1013323212100103-1101013232231223-3002202212203000-2132102002032332-1323101013303320"></a>

## tiering_policy property — volume_defaults / 302311301001 / 14

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Tiering policy to use. "none" is default.

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

<a id="canonical-0011031211301230-1210033212313100-0202300323331231-0013303032221032-2303123311200122-3232133011203010-1202110202203012-0321110033003201"></a>

<a id="canonical-3003310321123131-2023000310002311-1033001230023101-0003320032120130-2000032113031313-3032211031130222-1300130323331310-2003103221230313"></a>

## unix_permissions property — volume_defaults / 302311301001 / 15

Type: `"number"`. Optional.

Unix permission mode for new volumes. All allowed 777.

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

<a id="canonical-0300220201312211-0120202211033032-1101320111201033-3330300213221232-2323212320030103-2032000311212111-1012121123132031-3231010030121132"></a>

## Next pages — volume_defaults / 302311301001 / 16

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos](resources--voltstack_site--reference--group-007.md#canonical-2032231111203001-0101021010111211-3210232311311103-1102113222303221-0132303120120233-0213230122230003-0221102003103200-2132330121202331)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2032231111203001-0101021010111211-3210232311311103-1102113222303221-0132303120120233-0213230122230003-0221102003103200-2132330121202331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223212233310223-3333010311110030-0233112301330232-2121323220100332-3010233332200223-2123212332110300-1201131021000100-0323213110332220"></a>

## custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos — no_qos / 212210002030 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](resources--voltstack_site--reference--group-006.md#canonical-2320103020333102-1021022013301311-0131221310202231-1112221003301021-0300321100322133-0201110031211000-3323031032010313-1110131030020123)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san](resources--voltstack_site--reference--group-006.md#canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](resources--voltstack_site--reference--group-007.md#canonical-1113012300123231-0230130213102123-2222113203222300-0231133123010033-2003312131131202-2321303021133030-3102013322103330-2102023110321222)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults.no_qos

<a id="canonical-2111232033220201-2211121211022000-2001123202300210-1130130023310112-3022001103122231-1022003333230010-2002033010021102-0202301000213010"></a>

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
no_qos = {}
```

<a id="canonical-1031213110223012-1223213120221201-1321001310013122-3230120303121132-0201002002122133-0033023023121313-3210001331133013-0233001303031200"></a>

## Direct properties — no_qos / 212210002030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212201012100130-1133123122301122-0321331022300333-2023130202331022-1321330233322003-1100002001132220-1133133010221102-1022310130213020"></a>

## Next pages — no_qos / 212210002030 / 4

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](resources--voltstack_site--reference--group-007.md#canonical-1113012300123231-0230130213102123-2222113203222300-0231133123010033-2003312131131202-2321303021133030-3102013322103330-2102023110321222)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3132213030333223-1013111232200333-3312303011012223-0311332321032113-1021201312322202-1000002031033130-3100030332232231-2131112101022100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001031111223022-0030331322103310-3021331303101311-0232102333201121-2323213132321211-1013001212102310-2322202003331001-3003033030101232"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator — pure_service_orchestrator / 121332301320 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator

<a id="canonical-2000301003302302-3001003031030011-2321110122321032-2021002212210320-3120112121111203-0133031303330313-1330122101331310-0010213311213023"></a>

Type: `"object"`. single nested block, Optional.

Device configuration for Pure Storage Service Orchestrator.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster_id")}
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
pure_service_orchestrator {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310001312020210-3031032112112230-2213301003031130-0001110303302130-1021333222123101-1300131010130210-2220312202201011-3322033120201011"></a>

## Direct properties — pure_service_orchestrator / 121332301320 / 3

- [arrays](resources--voltstack_site--reference--group-007.md#canonical-0100321132011122-1132132203331030-3001013113022302-3112120333030233-3032003101201120-0222320033322030-3010211220120020-1130332211231211): complete subsection reference.

<a id="canonical-2232330033202102-1233033301212011-2302312001032300-0031232223123003-2221003132030002-1210030003022132-3133300111001220-3300131310332303"></a>

<a id="canonical-3010231222003130-3223313233321331-1230020100010132-0003301230013201-2030130212030130-1221122221230011-3032300221003222-2133012013000321"></a>

## cluster_id property — pure_service_orchestrator / 121332301320 / 4

Type: `"string"`. Optional.

ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also
used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple
K8s clusters running on top of the same storage arrays.

Upstream description:

ClusterID is added as a prefix for all volumes created by this PSO installation. ClusterID is also
used to identify the volumes used by the datastore, pso-db. ClusterID MUST BE UNIQUE for multiple
K8s clusters running on top of the same storage arrays. Characters allowed: alphanumeric and
underscores.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 22),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 22,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 22,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z0-9_]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "22",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9_]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "22",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9_]*$"
  }
}
```

<a id="canonical-1130231233000033-3101310023303332-1100232330003120-3121112301303332-2001213123001323-0213021133011221-1302213223303303-1111211201320103"></a>

<a id="canonical-3021102221021011-2002202020112320-1122132011010220-0231111333202130-1003113202011110-1120302110203020-3003221100221320-2300331322000111"></a>

## enable_storage_topology property — pure_service_orchestrator / 121332301320 / 5

Type: `"bool"`. Optional.

Option is to enable/disable the csi topology feature for pso-csi.

Upstream description:

This option is to enable/disable the csi topology feature for pso-csi.

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

<a id="canonical-1003030112211330-3013322323310321-1013320312033130-2200331021320220-0111303203121113-1232322121333231-0032231201332002-1001310011200022"></a>

<a id="canonical-2101233220330222-1213312330320222-0333130333121032-0021010000320332-1222111032220231-2331100033303221-3323001120213030-3233211331231323"></a>

## enable_strict_topology property — pure_service_orchestrator / 121332301320 / 6

Type: `"bool"`. Optional.

Option is to enable/disable the strict csi topology feature for pso-csi.

Upstream description:

This option is to enable/disable the strict csi topology feature for pso-csi.

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

<a id="canonical-1021010223132122-0321212031301321-3102320022200203-0332220203113213-3212232130232000-1122301112230333-3100301212000223-3201333321310103"></a>

## Next pages — pure_service_orchestrator / 121332301320 / 7

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-0100321132011122-1132132203331030-3001013113022302-3112120333030233-3032003101201120-0222320033322030-3010211220120020-1130332211231211)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0100321132011122-1132132203331030-3001013113022302-3112120333030233-3032003101201120-0222320033322030-3010211220120020-1130332211231211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332000303333230-1332123020323320-1223102100230322-3321213022113322-0010102001200200-3231131010302211-0112231201012200-2011131002320202"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays — arrays / 133322032112 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-3132213030333223-1013111232200333-3312303011012223-0311332321032113-1021201312322202-1000002031033130-3100030332232231-2131112101022100)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays

<a id="canonical-3131311220330213-2000313003000012-1121203310032303-1210200123113223-2303202100003233-2323322213030100-0231101020221000-0211210202231221"></a>

Type: `"object"`. single nested block, Optional.

Arrays Configuration. Device configuration for PSO Arrays.

Upstream description:

Device configuration for PSO Arrays.

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
arrays {
  # Configure direct properties listed below.
}
```

<a id="canonical-1021333132023100-0123030112001000-0203001131212121-2213203220123310-2021131133120312-0023320201110211-2221302121033233-3000111112221321"></a>

## Direct properties — arrays / 133322032112 / 3

- [flash_array](resources--voltstack_site--reference--group-007.md#canonical-3023300311332123-1012132203112032-2312232233323131-1010112010012331-0300012332100033-0012113021100013-0233131020330231-1011122311120321): complete subsection reference.

- [flash_blade](resources--voltstack_site--reference--group-007.md#canonical-3211133003223322-1122130310321330-3231131012111020-0322220301303133-1220201030302232-1301302211320032-3310003301333221-3113222120032113): complete subsection reference.

<a id="canonical-0113030203020111-2212302133231221-3102102212013221-3330302020331311-1322010300033020-1100330110013131-3332131120230131-3202131022222120"></a>

## Next pages — arrays / 133322032112 / 4

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--voltstack_site--reference--group-007.md#canonical-3023300311332123-1012132203112032-2312232233323131-1010112010012331-0300012332100033-0012113021100013-0233131020330231-1011122311120321)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--voltstack_site--reference--group-007.md#canonical-3211133003223322-1122130310321330-3231131012111020-0322220301303133-1220201030302232-1301302211320032-3310003301333221-3113222120032113)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-3132213030333223-1013111232200333-3312303011012223-0311332321032113-1021201312322202-1000002031033130-3100030332232231-2131112101022100)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3023300311332123-1012132203112032-2312232233323131-1010112010012331-0300012332100033-0012113021100013-0233131020330231-1011122311120321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130021012123330-1313123210001121-2331330320033231-1120121030120333-0002133300302023-0100302000332313-0210221233310133-3301020322101301"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array — flash_array / 330033312300 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-3132213030333223-1013111232200333-3312303011012223-0311332321032113-1021201312322202-1000002031033130-3100030332232231-2131112101022100)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-0100321132011122-1132132203331030-3001013113022302-3112120333030233-3032003101201120-0222320033322030-3010211220120020-1130332211231211)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array

<a id="canonical-0003311223232210-0023023002230122-2031313023202030-3013313003133002-0012331123123131-2030222203100023-0310231311120313-3100122231232313"></a>

Type: `"object"`. single nested block, Optional.

Specify what storage flash arrays should be managed the plugin.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("default_fs_type",
    "flash_arrays",
    "iscsi_login_timeout",
    "san_type")}
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
flash_array {
  # Configure direct properties listed below.
}
```

<a id="canonical-3120012012021223-1303103112013111-3301000121003220-2131201013130213-2023022202211012-3131320213121212-3201312032033020-2230231121303011"></a>

## Direct properties — flash_array / 330033312300 / 3

<a id="canonical-3330200013322301-2021303133220013-3021333001120302-1313200130201030-3003331112110220-0202310321233211-0133000100110133-2303030111132302"></a>

<a id="canonical-1230323201311133-3030211023130020-3200113321031201-1113211123231313-3102200223331322-3121033222222201-2211010003032331-1200233203120110"></a>

## default_fs_opt property — flash_array / 330033312300 / 4

Type: `"string"`. Optional.

Block volume default mkfs OPTIONS. Not recommended to change!

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2311332233330333-2302323022012233-2003031100320212-0020122011220033-1310003210133211-2301003232321002-3222133300301231-0220132131022032"></a>

<a id="canonical-3110030110132233-3122021103300130-3022301300033202-0112300223313022-2331110313030330-3113200233231323-1311013222321110-0101000331233123"></a>

## default_fs_type property — flash_array / 330033312300 / 5

Type: `"string"`. Optional.

\[Enum: xfs|ext4\] Block volume default filesystem type. Not recommended to change!. Possible values
are \`xfs\`, \`ext4\`.

Upstream description:

Block volume default filesystem type. Not recommended to change!

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("xfs",
    "ext4"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "xfs",
    "ext4"
  ],
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"xfs\\\",\\\"ext4\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"xfs\\\",\\\"ext4\\\"]"
  }
}
```

<a id="canonical-2111300330001013-2200311321320022-3013213332022332-0001022222112333-1000223131020212-2322000332313020-1113112300311010-2022100103030202"></a>

<a id="canonical-3322302113210021-3102323212202320-0211320322113010-0110032013223121-1002130213100122-2203211230223032-3320121200301022-0100202103122303"></a>

## default_mount_opts property — flash_array / 330033312300 / 6

Type: `["list", "string"]`. Optional.

Block volume default filesystem mount OPTIONS. Not recommended to change!

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3332321200231021-3320203123110311-1211322230333331-3201212011001222-0211011330332113-1023031200322210-2203300011303302-3201131331013112"></a>

<a id="canonical-0111110210330300-1311302230323232-3201030100332301-1013312230202222-2012303303202132-3220322331201022-3211313020131130-0301223303101210"></a>

## disable_preempt_attachments property — flash_array / 330033312300 / 7

Type: `"bool"`. Optional.

Disable Preempt Attachments. Enable/Disable attachment preemption!

Upstream description:

Enable/Disable attachment preemption!

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

- [flash_arrays](resources--voltstack_site--reference--group-007.md#canonical-3330103103130232-2013231201213331-1002021200311320-0233022003110200-1221013311033312-3331300112012321-1213302233230032-2001102003120102): complete subsection reference.

<a id="canonical-0131300300132030-1313010312210210-3332213200130301-1331323320300332-2112102223131132-0332030321010101-2011120113212312-0221122222313212"></a>

<a id="canonical-3131210322132010-2021212033000032-1130330130300131-0020232120321000-2001301100130012-1020010300230033-1321122323212023-1201132303100311"></a>

## iscsi_login_timeout property — flash_array / 330033312300 / 8

Type: `"number"`. Optional.

ISCSI login timeout in seconds. Not recommended to change!

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-0200002030301120-0003232121021030-0132011220021011-1012013130111122-3023010102313120-1010011311223103-3310323313111213-1201022021303301"></a>

<a id="canonical-0330220323201312-3312011230131100-1323020302100110-2221202022333001-0203201012313223-3132130112233203-0221123322331212-2132133111211322"></a>

## san_type property — flash_array / 330033312300 / 9

Type: `"string"`. Optional.

\[Enum: ISCSI|FC\] Block volume access protocol, either ISCSI or FC. Possible values are \`ISCSI\`,
\`FC\`.

Upstream description:

Block volume access protocol, either ISCSI or FC.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ISCSI",
    "FC"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ISCSI",
    "FC"
  ],
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ISCSI\\\",\\\"FC\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ISCSI\\\",\\\"FC\\\"]"
  }
}
```

<a id="canonical-1332031120013002-1010011300302300-0313332113000100-2322311000132033-1101020331302030-2033103312212113-3021333133202310-0201023032111111"></a>

## Next pages — flash_array / 330033312300 / 10

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--voltstack_site--reference--group-007.md#canonical-3330103103130232-2013231201213331-1002021200311320-0233022003110200-1221013311033312-3331300112012321-1213302233230032-2001102003120102)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-0100321132011122-1132132203331030-3001013113022302-3112120333030233-3032003101201120-0222320033322030-3010211220120020-1130332211231211)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3330103103130232-2013231201213331-1002021200311320-0233022003110200-1221013311033312-3331300112012321-1213302233230032-2001102003120102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330322021213311-3221102112230230-0321112130120103-1010330013020133-0033100001301002-2112120210201210-2012031333000201-3020232323000233"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays — flash_arrays / 100211213033 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-3132213030333223-1013111232200333-3312303011012223-0311332321032113-1021201312322202-1000002031033130-3100030332232231-2131112101022100)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-0100321132011122-1132132203331030-3001013113022302-3112120333030233-3032003101201120-0222320033322030-3010211220120020-1130332211231211)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--voltstack_site--reference--group-007.md#canonical-3023300311332123-1012132203112032-2312232233323131-1010112010012331-0300012332100033-0012113021100013-0233131020330231-1011122311120321)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays

<a id="canonical-3302332300300210-2233121030132233-3031330021213201-1010231010321322-1233202332112013-2110000333012030-3333012300112303-2001013113002011"></a>

Type: `"object"`. list nested block, Optional.

For FlashArrays you must set the 'mgmt\_endpoint' and 'api\_token'.

Upstream description:

For FlashArrays you must set the "mgmt\_endpoint" and "api\_token"

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("mgmt_dns_name",
    "mgmt_ip")}
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
flash_arrays {
  # Configure direct properties listed below.
}
```

<a id="canonical-1021333133101020-0222300223232232-0231131032110012-3310312321322323-2223110112001210-0000010130120001-2010131330030112-2113311301233101"></a>

## Direct properties — flash_arrays / 100211213033 / 3

- [api_token](resources--voltstack_site--reference--group-007.md#canonical-1233301000203313-0003330123212311-1222100233332333-2002122001132001-1220302033201121-0202131132133131-0012221200211133-0233223103111332): complete subsection reference.

<a id="canonical-3230030133322221-1320223302121232-1000313111222211-3333330211120302-2220002300100012-3211303023033111-2021221010031230-0023022223103320"></a>

<a id="canonical-1211123123331223-0120013201232102-1023221032002120-1110331301133331-3222303121032221-3111301211313333-2331220101221111-1121003123112013"></a>

## labels property — flash_arrays / 100211213033 / 4

Type: `["map", "string"]`. Optional.

Specifies labels optional, and can be any key-value pair for use with the PSO 'fleet' provisioner.

Upstream description:

The labels are optional, and can be any key-value pair for use with the PSO "fleet" provisioner.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-2312331333303210-3302333310222112-1113300332331300-2303002331033221-1102022313232113-0032202201021303-2332120021323333-2123131323333230"></a>

<a id="canonical-0312110121330213-2223202031011111-1302211300133131-1222113102300131-0231020213021213-3112123000320103-0130130020312201-1201330110030300"></a>

## mgmt_dns_name property — flash_arrays / 100211213033 / 5

Type: `"string"`. Optional.

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2311311333002301-3000303202012202-1233213110133330-1321201002132023-3122020330333331-3232221220212311-2023013020332301-2312122333203022"></a>

<a id="canonical-1020022203122031-3301012013102132-0331223013122320-3231202211102220-2030213301101313-2110002311020002-0023001032333313-2112100130222030"></a>

## mgmt_ip property — flash_arrays / 100211213033 / 6

Type: `"string"`. Optional.

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-1020102131122222-1033332112030212-3101113102332330-1030222103220211-2112303030331233-1130233332031013-0330132312013312-2101022101210112"></a>

## Next pages — flash_arrays / 100211213033 / 7

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--voltstack_site--reference--group-007.md#canonical-1233301000203313-0003330123212311-1222100233332333-2002122001132001-1220302033201121-0202131132133131-0012221200211133-0233223103111332)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--voltstack_site--reference--group-007.md#canonical-3023300311332123-1012132203112032-2312232233323131-1010112010012331-0300012332100033-0012113021100013-0233131020330231-1011122311120321)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1233301000203313-0003330123212311-1222100233332333-2002122001132001-1220302033201121-0202131132133131-0012221200211133-0233223103111332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010030121230220-3211231032013103-1111012303031303-1210212002010003-0031122112000232-2032301330123332-2103102333002111-2300132213320012"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token — api_token / 101332233032 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-3132213030333223-1013111232200333-3312303011012223-0311332321032113-1021201312322202-1000002031033130-3100030332232231-2131112101022100)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-0100321132011122-1132132203331030-3001013113022302-3112120333030233-3032003101201120-0222320033322030-3010211220120020-1130332211231211)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--voltstack_site--reference--group-007.md#canonical-3023300311332123-1012132203112032-2312232233323131-1010112010012331-0300012332100033-0012113021100013-0233131020330231-1011122311120321)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--voltstack_site--reference--group-007.md#canonical-3330103103130232-2013231201213331-1002021200311320-0233022003110200-1221013311033312-3331300112012321-1213302233230032-2001102003120102)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token

<a id="canonical-0230020021232331-3210120010232112-0232321013020301-1002022120130223-0010020123102333-2322030112232122-1011210322032203-0202203322023111"></a>

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
api_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210202210012200-0033001213313230-0212011110122203-0003232303212120-2213023301222320-1120212020303032-2032311323123311-2311310313331323"></a>

## Direct properties — api_token / 101332233032 / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-3323310321002123-1313203322032213-3303321213112032-1321331302212311-0123300211320013-3200220032311121-1012023322312222-3101101022322130): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-1002032210223131-0321210012202130-1130312100131023-2001223221203221-2213133021223222-1200232001033113-3302013310332002-3301322203132020): complete subsection reference.

<a id="canonical-0220121330213213-2123322201132332-1210123130100311-3312332132331332-1030313023231220-1220310331321223-3300321313000121-1103303231330311"></a>

## Next pages — api_token / 101332233032 / 4

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-3323310321002123-1313203322032213-3303321213112032-1321331302212311-0123300211320013-3200220032311121-1012023322312222-3101101022322130)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-1002032210223131-0321210012202130-1130312100131023-2001223221203221-2213133021223222-1200232001033113-3302013310332002-3301322203132020)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--voltstack_site--reference--group-007.md#canonical-3330103103130232-2013231201213331-1002021200311320-0233022003110200-1221013311033312-3331300112012321-1213302233230032-2001102003120102)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3323310321002123-1313203322032213-3303321213112032-1321331302212311-0123300211320013-3200220032311121-1012023322312222-3101101022322130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230000332221031-1000033000302300-0231113213212313-2312303131321211-3000121210300301-2121101130011300-2300121123013112-0333030002221210"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info — blindfold_secret_info / 031323011322 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-3132213030333223-1013111232200333-3312303011012223-0311332321032113-1021201312322202-1000002031033130-3100030332232231-2131112101022100)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-0100321132011122-1132132203331030-3001013113022302-3112120333030233-3032003101201120-0222320033322030-3010211220120020-1130332211231211)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--voltstack_site--reference--group-007.md#canonical-3023300311332123-1012132203112032-2312232233323131-1010112010012331-0300012332100033-0012113021100013-0233131020330231-1011122311120321)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--voltstack_site--reference--group-007.md#canonical-3330103103130232-2013231201213331-1002021200311320-0233022003110200-1221013311033312-3331300112012321-1213302233230032-2001102003120102)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--voltstack_site--reference--group-007.md#canonical-1233301000203313-0003330123212311-1222100233332333-2002122001132001-1220302033201121-0202131132133131-0012221200211133-0233223103111332)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.blindfold_secret_info

<a id="canonical-2102222231312130-0303311120312123-3332302311010130-3312020021211010-3301120223303302-2213330111233113-2020230223011120-0010010210211120"></a>

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

<a id="canonical-0310222113212332-2010330231000013-1221131332302210-3102310202002013-3102130020233003-0020320201330131-0131032311303010-3012300100113102"></a>

## Direct properties — blindfold_secret_info / 031323011322 / 3

<a id="canonical-0302212013122200-0232222332300101-0023222020000111-3100203130222230-3023110103310221-2211213000023333-0213233233001111-0000231210332020"></a>

<a id="canonical-3010101301333020-2201331123022121-2322033202210123-2010302011110230-3311022022120033-2322121101323201-1111303101023201-3022231312232022"></a>

## decryption_provider property — blindfold_secret_info / 031323011322 / 4

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

<a id="canonical-2023112023121331-2323111301220232-1103332300311022-2000323231313313-0323030303302222-2132022233312122-2221113322100313-1003121222313300"></a>

<a id="canonical-1200003013112310-0022223301202233-1121030111222233-3222211010102101-2213120201112310-3003231012220301-0231121132103212-3232000133222233"></a>

## location property — blindfold_secret_info / 031323011322 / 5

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

<a id="canonical-2300101132333220-3031301000203301-2130123133233313-2030222133022101-0110222001030311-0302220033113201-3123301112311200-1033210301210123"></a>

<a id="canonical-3103133202131123-0002110123213031-1131033011300231-1110233312233312-1123330203101030-3103112232332323-1021130330032032-2100313332332113"></a>

## store_provider property — blindfold_secret_info / 031323011322 / 6

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

<a id="canonical-3300133100311132-0111110103020232-2021211111310203-0011222301020002-3223001120011212-0101200013222012-1111332133000132-3012011232111030"></a>

## Next pages — blindfold_secret_info / 031323011322 / 7

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--voltstack_site--reference--group-007.md#canonical-1233301000203313-0003330123212311-1222100233332333-2002122001132001-1220302033201121-0202131132133131-0012221200211133-0233223103111332)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1002032210223131-0321210012202130-1130312100131023-2001223221203221-2213133021223222-1200232001033113-3302013310332002-3301322203132020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020020120322210-2132222222220133-0210311311032300-1221101200213332-0213121213100002-3201213012112200-0111032103210323-2031333333211213"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info — clear_secret_info / 110131201213 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-3132213030333223-1013111232200333-3312303011012223-0311332321032113-1021201312322202-1000002031033130-3100030332232231-2131112101022100)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-0100321132011122-1132132203331030-3001013113022302-3112120333030233-3032003101201120-0222320033322030-3010211220120020-1130332211231211)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array](resources--voltstack_site--reference--group-007.md#canonical-3023300311332123-1012132203112032-2312232233323131-1010112010012331-0300012332100033-0012113021100013-0233131020330231-1011122311120321)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays](resources--voltstack_site--reference--group-007.md#canonical-3330103103130232-2013231201213331-1002021200311320-0233022003110200-1221013311033312-3331300112012321-1213302233230032-2001102003120102)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--voltstack_site--reference--group-007.md#canonical-1233301000203313-0003330123212311-1222100233332333-2002122001132001-1220302033201121-0202131132133131-0012221200211133-0233223103111332)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token.clear_secret_info

<a id="canonical-1011212211333011-2112033200010113-1020021120101102-1032000231132203-1121112233133220-2103322222033131-0132121222012102-2110320210311132"></a>

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

<a id="canonical-0321311302202221-1131301012011011-0301111113000323-2212003012200112-1313320010221231-2032220213311232-0333333332002101-0020201210303022"></a>

## Direct properties — clear_secret_info / 110131201213 / 3

<a id="canonical-2212231202131022-3323322101002203-1203133223022232-2313203330332120-2323321011230223-0132012000100331-2103212220302122-3230332113000311"></a>

<a id="canonical-3312221300121132-1100103132133012-1211212222121031-3232102013332303-0131113133122220-1220330011223032-1322130001201130-2020211322000031"></a>

## provider_ref property — clear_secret_info / 110131201213 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1220022203110210-3020211310233121-0131010210321113-3303332222113311-1233200002103003-2012033330102100-1213222232311031-3131330330322000"></a>

<a id="canonical-0001113103221222-3211302120022002-1220031230201331-0211120132230213-0313203032313310-0220131021321010-0312223311313023-3111001010301033"></a>

## URL property — clear_secret_info / 110131201213 / 5

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

<a id="canonical-1212213021122323-1113003130312011-0213132011210222-3211230331023301-0022131020201301-1330200302310122-2230231123133002-3200023120023221"></a>

## Next pages — clear_secret_info / 110131201213 / 6

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_array.flash_arrays.api_token](resources--voltstack_site--reference--group-007.md#canonical-1233301000203313-0003330123212311-1222100233332333-2002122001132001-1220302033201121-0202131132133131-0012221200211133-0233223103111332)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3211133003223322-1122130310321330-3231131012111020-0322220301303133-1220201030302232-1301302211320032-3310003301333221-3113222120032113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220212231123032-3333112021311030-2021221112000010-1111001103321111-2230302223002001-1332222010231013-1201213221323232-2100303322201301"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade — flash_blade / 233011112310 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-3132213030333223-1013111232200333-3312303011012223-0311332321032113-1021201312322202-1000002031033130-3100030332232231-2131112101022100)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-0100321132011122-1132132203331030-3001013113022302-3112120333030233-3032003101201120-0222320033322030-3010211220120020-1130332211231211)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade

<a id="canonical-3322123213223011-3103113012230321-1210113120030101-0302302320010013-3111110131133332-3111202001300320-0120020322013102-2133233303111120"></a>

Type: `"object"`. single nested block, Optional.

Specify what storage flash blades should be managed the plugin.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("flash_blades")}
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
flash_blade {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222033323301213-1313022302222113-3012330103031322-0013131331321002-0111221130121020-3002111212003232-2220112300022311-1210201020101020"></a>

## Direct properties — flash_blade / 233011112310 / 3

<a id="canonical-0130202331011232-3322333322023001-2221311222120011-3002012012003331-3031022133112312-1000112322213011-0333113003211323-3210233332221120"></a>

<a id="canonical-1302022233030320-1313330231030001-1310300033200032-3021032120031302-0220200230221121-2230033222200022-3211232101001332-1033101123003203"></a>

## enable_snapshot_directory property — flash_blade / 233011112310 / 4

Type: `"bool"`. Optional.

Enable Snapshot Directory. Enable/Disable FlashBlade snapshots.

Upstream description:

Enable/Disable FlashBlade snapshots.

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

<a id="canonical-1002310230011013-2203223323101303-0123320103212023-1332222330130023-3211002232002122-1311121021210321-2323232230131021-2020122221123112"></a>

<a id="canonical-0310323100002212-0022113213201321-1201323132231133-0302220120320213-1203103003121011-3232001233232020-1133031113131320-3303303030101000"></a>

## export_rules property — flash_blade / 233011112310 / 5

Type: `"string"`. Optional.

NFS Export Rules. NFS Export rules.

Upstream description:

NFS Export rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 250),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 250,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 250,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "250",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "250",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [flash_blades](resources--voltstack_site--reference--group-007.md#canonical-2101231320132023-0302202321030132-3003003113231321-3120030023303303-2133200320321303-1321220210022222-3331013302312220-1113220301101121): complete subsection reference.

<a id="canonical-0021220201111202-0121000112332310-1022300123023113-0032123310032203-1000030012210111-3331223011031300-0133331113201022-1133203233030030"></a>

## Next pages — flash_blade / 233011112310 / 6

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--voltstack_site--reference--group-007.md#canonical-2101231320132023-0302202321030132-3003003113231321-3120030023303303-2133200320321303-1321220210022222-3331013302312220-1113220301101121)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-0100321132011122-1132132203331030-3001013113022302-3112120333030233-3032003101201120-0222320033322030-3010211220120020-1130332211231211)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2101231320132023-0302202321030132-3003003113231321-3120030023303303-2133200320321303-1321220210022222-3331013302312220-1113220301101121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203020131301320-3313003113221003-2030100212312113-3311212222000020-1031102001113212-1212320330013320-1333223023202113-3300030130313030"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades — flash_blades / 122230233231 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-3132213030333223-1013111232200333-3312303011012223-0311332321032113-1021201312322202-1000002031033130-3100030332232231-2131112101022100)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-0100321132011122-1132132203331030-3001013113022302-3112120333030233-3032003101201120-0222320033322030-3010211220120020-1130332211231211)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--voltstack_site--reference--group-007.md#canonical-3211133003223322-1122130310321330-3231131012111020-0322220301303133-1220201030302232-1301302211320032-3310003301333221-3113222120032113)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades

<a id="canonical-0101030230011332-2203020001013011-3210322023003130-2122311212022011-2233310200000012-0223123312331203-2032323231020212-2220002033131010"></a>

Type: `"object"`. list nested block, Optional.

For FlashBlades you must set the 'mgmt\_endpoint', 'api\_token' and nfs\_endpoint.

Upstream description:

For FlashBlades you must set the "mgmt\_endpoint", "api\_token" and nfs\_endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("mgmt_dns_name",
    "mgmt_ip"),
  validators.ConflictingListObjectAttributes("nfs_endpoint_dns_name",
    "nfs_endpoint_ip")}
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
flash_blades {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233300220230022-0323201320333223-3303312333131001-1113031013121120-0220222133210033-0003232313133313-3221211110023023-2303013312200110"></a>

## Direct properties — flash_blades / 122230233231 / 3

- [api_token](resources--voltstack_site--reference--group-007.md#canonical-1331201220120101-0231302101001011-1111031302133231-2032331223220320-2233102032132032-2210323011310300-1000000223320012-3000111203123111): complete subsection reference.

<a id="canonical-3303021213001022-3332022333220101-3332221313333313-3323110332033310-0221011313121030-3022030103033111-3003021222031033-3321002323220221"></a>

<a id="canonical-3220212213233100-0122310321002003-3022033322130301-2303331311102223-0112013330020110-3301221312123311-1232202211223032-3033313102012303"></a>

## labels property — flash_blades / 122230233231 / 4

Type: `["map", "string"]`. Optional.

Specifies labels optional, and can be any key-value pair for use with the PSO 'fleet' provisioner.

Upstream description:

The labels are optional, and can be any key-value pair for use with the PSO "fleet" provisioner.

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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-2102031311222330-2200102020322302-2012203300212230-1113233210231231-3123321232311233-3002210302032222-3222203223131222-3313010313102130"></a>

<a id="canonical-0100210201111033-3322230020333303-2133010303213000-1330121311223021-2323211113001120-2012320333023322-3213212022132300-3013301131203003"></a>

## mgmt_dns_name property — flash_blades / 122230233231 / 5

Type: `"string"`. Optional.

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[mgmt\_ip\] Management Endpoint's IP address is discovered using DNS name
resolution. The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1013323101323333-1132222323231200-1123033130303213-3333211202301121-1302333101122023-1120120222202003-0213020113111023-0131000022113330"></a>

<a id="canonical-0032110003233112-3111112232100023-0031001323013311-2333011331032330-0322222321230231-0110201122313110-2323111032111110-1032332220210123"></a>

## mgmt_ip property — flash_blades / 122230233231 / 6

Type: `"string"`. Optional.

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[mgmt\_dns\_name\] Management Endpoint is reachable at the given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2022112021000320-0331320033010132-0000321223131200-1002010321130230-3032110211100031-0220103030211100-3001010200011322-0020010000211120"></a>

<a id="canonical-3002133132002021-3132321013003133-0022302000120001-3112031222332001-1132223201213311-0130032301301112-2030110303012321-1003331103201003"></a>

## nfs_endpoint_dns_name property — flash_blades / 122230233231 / 7

Type: `"string"`. Optional.

Exclusive with \[nfs\_endpoint\_ip\] Endpoint's IP address is discovered using DNS name resolution.
The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[nfs\_endpoint\_ip\] Endpoint's IP address is discovered using DNS name resolution.
The name given here is fully qualified domain name.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3102020012332320-2203130031212210-3200002302023031-1313131321100031-2020223230301211-0111122212101302-0012220320323200-1203023323313330"></a>

<a id="canonical-3201010320012131-0223121211123212-0212033022100103-2222201103111113-1332332131311302-2220131213203221-2303003320000100-2122221323000112"></a>

## nfs_endpoint_ip property — flash_blades / 122230233231 / 8

Type: `"string"`. Optional.

Exclusive with \[nfs\_endpoint\_dns\_name\] Endpoint is reachable at the given IP address.

Upstream description:

Exclusive with \[nfs\_endpoint\_dns\_name\] Endpoint is reachable at the given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-0223320100101021-0133302121223113-2211013200311233-3213232312101303-0230331012222021-2232302201320232-2023011023331331-2023323111102232"></a>

## Next pages — flash_blades / 122230233231 / 9

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--voltstack_site--reference--group-007.md#canonical-1331201220120101-0231302101001011-1111031302133231-2032331223220320-2233102032132032-2210323011310300-1000000223320012-3000111203123111)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--voltstack_site--reference--group-007.md#canonical-3211133003223322-1122130310321330-3231131012111020-0322220301303133-1220201030302232-1301302211320032-3310003301333221-3113222120032113)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1331201220120101-0231302101001011-1111031302133231-2032331223220320-2233102032132032-2210323011310300-1000000223320012-3000111203123111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311112231121230-2210013310222203-3012031302120303-0332222022230321-1130213102112323-2101110223332133-2133121133032031-2311030102010111"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token — api_token / 111200331322 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-3132213030333223-1013111232200333-3312303011012223-0311332321032113-1021201312322202-1000002031033130-3100030332232231-2131112101022100)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-0100321132011122-1132132203331030-3001013113022302-3112120333030233-3032003101201120-0222320033322030-3010211220120020-1130332211231211)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--voltstack_site--reference--group-007.md#canonical-3211133003223322-1122130310321330-3231131012111020-0322220301303133-1220201030302232-1301302211320032-3310003301333221-3113222120032113)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--voltstack_site--reference--group-007.md#canonical-2101231320132023-0302202321030132-3003003113231321-3120030023303303-2133200320321303-1321220210022222-3331013302312220-1113220301101121)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token

<a id="canonical-0212313211322211-2332330133200122-2031230301211020-3032120030302202-3002021233322203-2023122010120111-0102212321330333-2231233301301123"></a>

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
api_token {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310130010330221-0311231322232200-3220022201122301-0111232032102323-3310223002122300-3331302201211103-3333212211033132-0030222000003320"></a>

## Direct properties — api_token / 111200331322 / 3

- [blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-1132222321331001-0310231321100311-2332322201202223-3332222011113111-2211201123202031-0233330122330310-2211303200012121-1312321112012121): complete subsection reference.

- [clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-0323113033023303-1021113121101010-2121302021033213-1200003320123031-2023113230021231-0120100013330230-3100121221312003-3323122222103230): complete subsection reference.

<a id="canonical-3332113010130011-2233103100200320-0312121201310310-1333210333331030-3110022132331130-0030023310121031-0213210113330202-2210323100321322"></a>

## Next pages — api_token / 111200331322 / 4

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info](resources--voltstack_site--reference--group-007.md#canonical-1132222321331001-0310231321100311-2332322201202223-3332222011113111-2211201123202031-0233330122330310-2211303200012121-1312321112012121)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info](resources--voltstack_site--reference--group-007.md#canonical-0323113033023303-1021113121101010-2121302021033213-1200003320123031-2023113230021231-0120100013330230-3100121221312003-3323122222103230)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--voltstack_site--reference--group-007.md#canonical-2101231320132023-0302202321030132-3003003113231321-3120030023303303-2133200320321303-1321220210022222-3331013302312220-1113220301101121)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1132222321331001-0310231321100311-2332322201202223-3332222011113111-2211201123202031-0233330122330310-2211303200012121-1312321112012121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310102331133131-2002321333121223-0100031223020111-0132200032213301-3001111012211122-3320112133023232-1002102023233322-0213130132013332"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info — blindfold_secret_info / 233333133031 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-3132213030333223-1013111232200333-3312303011012223-0311332321032113-1021201312322202-1000002031033130-3100030332232231-2131112101022100)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-0100321132011122-1132132203331030-3001013113022302-3112120333030233-3032003101201120-0222320033322030-3010211220120020-1130332211231211)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--voltstack_site--reference--group-007.md#canonical-3211133003223322-1122130310321330-3231131012111020-0322220301303133-1220201030302232-1301302211320032-3310003301333221-3113222120032113)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--voltstack_site--reference--group-007.md#canonical-2101231320132023-0302202321030132-3003003113231321-3120030023303303-2133200320321303-1321220210022222-3331013302312220-1113220301101121)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--voltstack_site--reference--group-007.md#canonical-1331201220120101-0231302101001011-1111031302133231-2032331223220320-2233102032132032-2210323011310300-1000000223320012-3000111203123111)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.blindfold_secret_info

<a id="canonical-3333221312020202-3330331321203203-2122211120232310-1101100232222203-1102220110323221-0101221131201301-0002113111032033-1023210333303230"></a>

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

<a id="canonical-2323331122002332-2300320033303000-0221113201330031-3121210111001111-0222200220001100-2023012021313022-0012120021000011-0102330111310223"></a>

## Direct properties — blindfold_secret_info / 233333133031 / 3

<a id="canonical-0303312113231020-0322303210132303-3031310303221212-2122202031332211-2331210023030031-3313231213221011-0010230201301022-3320032233133120"></a>

<a id="canonical-3311231301223120-2301020321033300-3303203131230301-2330110300301023-3230312231023110-3332102222310133-0311000012211120-0332300312212011"></a>

## decryption_provider property — blindfold_secret_info / 233333133031 / 4

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

<a id="canonical-3202020302013332-0031202301002303-2303320002111001-3121033332320332-3333332010311030-1021302122131301-3210131333123021-3233310210323331"></a>

<a id="canonical-3010231103023011-2130303212122003-1233122113232021-0113002102303112-1032211022303203-2200012332331323-0121223101001031-1300022332102333"></a>

## location property — blindfold_secret_info / 233333133031 / 5

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

<a id="canonical-1000332123212030-0221211012122223-2131323201311103-3211322032322001-1210312333211323-3320302300110103-0120221210300111-3210113110123323"></a>

<a id="canonical-3020103321212100-0200321033120230-0303130020111302-3213303203201222-0210031223012101-3032102132201332-3031002113030121-0011211031013113"></a>

## store_provider property — blindfold_secret_info / 233333133031 / 6

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

<a id="canonical-1013021022013032-2211112233201122-0212001021030000-1003323220120212-0313213123212211-0301010330112212-3303010013001330-3321001203200130"></a>

## Next pages — blindfold_secret_info / 233333133031 / 7

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--voltstack_site--reference--group-007.md#canonical-1331201220120101-0231302101001011-1111031302133231-2032331223220320-2233102032132032-2210323011310300-1000000223320012-3000111203123111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0323113033023303-1021113121101010-2121302021033213-1200003320123031-2023113230021231-0120100013330230-3100121221312003-3323122222103230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133021223302113-2302310202213300-0311230220320110-1320232322232223-0311310000232310-2322010300321323-0030200130002013-2131131133100220"></a>

## custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info — clear_secret_info / 212132010212 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_device_list](resources--voltstack_site--reference--group-006.md#canonical-3000120020011123-1022230333313213-0001000321031121-1203122203230213-2330200222212203-3333200100302030-1111002032202022-1131133200121321)
- [custom_storage_config.storage_device_list.storage_devices](resources--voltstack_site--reference--group-006.md#canonical-1131100102033112-2001212202303112-3112202123332303-0020133102302002-2003030303020101-2333230030103021-0230002120220201-0311201313132323)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator](resources--voltstack_site--reference--group-007.md#canonical-3132213030333223-1013111232200333-3312303011012223-0311332321032113-1021201312322202-1000002031033130-3100030332232231-2131112101022100)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays](resources--voltstack_site--reference--group-007.md#canonical-0100321132011122-1132132203331030-3001013113022302-3112120333030233-3032003101201120-0222320033322030-3010211220120020-1130332211231211)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade](resources--voltstack_site--reference--group-007.md#canonical-3211133003223322-1122130310321330-3231131012111020-0322220301303133-1220201030302232-1301302211320032-3310003301333221-3113222120032113)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades](resources--voltstack_site--reference--group-007.md#canonical-2101231320132023-0302202321030132-3003003113231321-3120030023303303-2133200320321303-1321220210022222-3331013302312220-1113220301101121)
- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--voltstack_site--reference--group-007.md#canonical-1331201220120101-0231302101001011-1111031302133231-2032331223220320-2233102032132032-2210323011310300-1000000223320012-3000111203123111)
- custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token.clear_secret_info

<a id="canonical-1033010120223013-2000233200133100-1330212112122221-0322311321112120-0230012232131133-1220001103013022-2133122033013012-0312230121332002"></a>

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

<a id="canonical-1311311210100311-1032231333012031-0303201202100112-3331013321110331-1012213011133110-0202320121320023-1031331223303103-3031023210311010"></a>

## Direct properties — clear_secret_info / 212132010212 / 3

<a id="canonical-2103010312212130-3033101321213201-1222110100323210-3203110102032102-3230201002001322-3003230022022310-3121011221202111-3031033221211332"></a>

<a id="canonical-1300132202203000-2012301020010011-3211010200133221-3110123301021131-1130200011112201-2203111103301130-2123002300123022-1030233312132022"></a>

## provider_ref property — clear_secret_info / 212132010212 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1113130301130232-3322201112130001-0322313112211020-3310202010032002-3200033003023300-3221303013021313-0121023002113013-1130300003320232"></a>

<a id="canonical-2120220103012230-0030220111020222-0101232001322030-0323110223010311-2120311021322131-2000321111301022-0322001000221231-0331022112323102"></a>

## URL property — clear_secret_info / 212132010212 / 5

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
