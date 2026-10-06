---
page_title: "xcsh_usb_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_usb_policy reference."
---

# xcsh_usb_policy reference

<a id="canonical-0203202331120213-3031223102221310-3303300333312020-2233220001321002-3220003200013332-1202131202021320-0033221202012231-3201003230011131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_usb_policy](../resources/usb_policy.md#canonical-3313033212003300-1002332031103020-1323001313000200-3003230033322232-3111031201110211-0323302100321032-3212002010220013-3201110212312232)
- Property reference

<a id="canonical-1301232320303033-3113322203202303-2103203210023321-2311013011203210-1221322221020020-2101011310222033-2213113303131130-3221011000010103"></a>

### Direct properties for `xcsh_usb_policy`

- [allowed_devices](resources--usb_policy--reference--group-001.md#canonical-3320231330031332-2111103103322211-1212323033212013-0110213130300113-0021012312002300-1222212221322113-0023133330222231-2322130331230220): complete subsection reference.

<a id="canonical-1122102023101312-1022003301120223-1033302313232201-1201123031102311-2130330321022100-3230313131320130-0320100122201001-0311301230331220"></a>

<a id="canonical-0230201010010021-3003300233202303-2020032102021300-0030123103322202-3131213320113020-2121321233122212-0302001201002122-3022233323111212"></a>

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

<a id="canonical-0203133011323021-1231220000310323-0111221001310331-3213312103301201-0312210003003102-2300022102031132-1001012131011100-2110200123332331"></a>

<a id="canonical-3200111213320303-3223203222033213-2123301103130212-0123120302220022-2130200112303231-0220303031211131-2310202220133200-2323121111112211"></a>

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

<a id="canonical-3323010022011110-1021032231213332-2222110321203210-0010132101222113-3311002111100303-2011311122000111-2103020212203302-2333112101232113"></a>

<a id="canonical-3223320300200211-0310201000102221-2102100011213033-2332220123302210-0100112213121110-3031120311211133-2332021300110332-2031131311321030"></a>

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

<a id="canonical-3010111203311023-1032022311112223-3033213023113200-3021302130000301-1322232201010110-2230033021131332-3232103132312022-1032113301001310"></a>

<a id="canonical-3103301231312012-3231000220200333-1103220332030211-0031132002131232-0023323221022230-1300322113122222-0103111312321003-0120212230121123"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1231231130221030-0003311231032013-1331221113122301-2122002002032022-1330331123233011-2333312323103310-2311222100121310-1002001001132100"></a>

<a id="canonical-2231113221322112-3210321010203023-2013103012313100-3333031220332030-2121020013021321-3232320103132100-0113122112102012-1231331020332001"></a>

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

<a id="canonical-2132331303022203-3230111201012312-2022003230123333-3122321221110212-3221013311311011-0210001121112330-3320301203322023-0111102200322101"></a>

<a id="canonical-2110013001000300-3021302230012022-3210200303333130-1332321013233111-0103210222221321-0300012100323331-3333000313033001-0320023003101200"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Usb Policy. Must be unique within the namespace.

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

<a id="canonical-2201210320111231-0230112330321111-0010330212213230-2301331332131220-1232233223232023-0001310111223131-3030222013103011-1210031133331301"></a>

<a id="canonical-2110320213132022-1221021021302002-1013022201210133-2311110301311131-0333201202111213-2323033100200133-2031020323012200-1303001213130222"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Usb Policy is created.

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

- [timeouts](resources--usb_policy--reference--group-001.md#canonical-3300131102022123-1002203003021200-1000033132210032-2030322333112113-1100033310120313-2201312001021030-1003100222211132-0003010110030112): complete subsection reference.

<a id="canonical-0120311310131320-0013311122302221-2023331331301001-3010212021132233-0102231222312300-0020012130120103-0100320200130303-2010212210132302"></a>

### All schema paths for `xcsh_usb_policy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `allowed_devices` | [allowed_devices](resources--usb_policy--reference--group-001.md#canonical-3210100222012221-0330131020230030-0010013320320333-2112303031121013-0031130002000321-1232220210203333-0001022223110222-1100313203012323) |
| `allowed_devices.b_device_class` | [allowed_devices.b_device_class](resources--usb_policy--reference--group-001.md#canonical-1311133221000032-2033230222333022-1131322220111112-3100213100020133-3303223311311013-3032231102302033-0130310133230223-3313112321333330) |
| `allowed_devices.b_device_protocol` | [allowed_devices.b_device_protocol](resources--usb_policy--reference--group-001.md#canonical-2020002221200103-3231211200010212-1300222021221102-2333310333302310-1202020001031123-1121010332300200-1031222031230231-3212112013131022) |
| `allowed_devices.b_device_sub_class` | [allowed_devices.b_device_sub_class](resources--usb_policy--reference--group-001.md#canonical-0030110320100010-3102301200331303-2212202232213302-1011001003213001-2213213131033310-0220002203122032-1331211030223102-0132310303201313) |
| `allowed_devices.i_serial` | [allowed_devices.i_serial](resources--usb_policy--reference--group-001.md#canonical-1303322100302331-1022110013223202-2223011133312020-2003231333220310-3000233130211213-2233212310012320-3213113212301030-1033001012323013) |
| `allowed_devices.id_product` | [allowed_devices.id_product](resources--usb_policy--reference--group-001.md#canonical-3001102310103302-3013320202301033-2003303300212330-0203220313300112-1210000220210033-0231022012023103-2203332232331202-2022133210002332) |
| `allowed_devices.id_vendor` | [allowed_devices.id_vendor](resources--usb_policy--reference--group-001.md#canonical-0313130031103102-0300223100013113-2020012223021203-1321020230232121-3202123312310001-1133103120030130-3123311330031323-1303300203201121) |
| `annotations` | [annotations](resources--usb_policy--reference--group-001.md#canonical-1122102023101312-1022003301120223-1033302313232201-1201123031102311-2130330321022100-3230313131320130-0320100122201001-0311301230331220) |
| `description` | [description](resources--usb_policy--reference--group-001.md#canonical-0203133011323021-1231220000310323-0111221001310331-3213312103301201-0312210003003102-2300022102031132-1001012131011100-2110200123332331) |
| `disable` | [disable](resources--usb_policy--reference--group-001.md#canonical-3323010022011110-1021032231213332-2222110321203210-0010132101222113-3311002111100303-2011311122000111-2103020212203302-2333112101232113) |
| `id` | [ID](resources--usb_policy--reference--group-001.md#canonical-3010111203311023-1032022311112223-3033213023113200-3021302130000301-1322232201010110-2230033021131332-3232103132312022-1032113301001310) |
| `labels` | [labels](resources--usb_policy--reference--group-001.md#canonical-1231231130221030-0003311231032013-1331221113122301-2122002002032022-1330331123233011-2333312323103310-2311222100121310-1002001001132100) |
| `name` | [name](resources--usb_policy--reference--group-001.md#canonical-2132331303022203-3230111201012312-2022003230123333-3122321221110212-3221013311311011-0210001121112330-3320301203322023-0111102200322101) |
| `namespace` | [namespace](resources--usb_policy--reference--group-001.md#canonical-2201210320111231-0230112330321111-0010330212213230-2301331332131220-1232233223232023-0001310111223131-3030222013103011-1210031133331301) |
| `timeouts` | [timeouts](resources--usb_policy--reference--group-001.md#canonical-1010111211223301-1301202101022033-2021132031123233-3232021101221010-1022020211322303-3133333132322321-1221302030310302-3311213033131132) |
| `timeouts.create` | [timeouts.create](resources--usb_policy--reference--group-001.md#canonical-3022120022132223-2130020302133213-0020101330103013-3113232203000323-3302001333220013-3113312021312220-1011032323122221-0323310100123212) |
| `timeouts.delete` | [timeouts.delete](resources--usb_policy--reference--group-001.md#canonical-0212103203120000-0101123122023211-0212200311331212-1133023001323330-1222012122313022-3311023222003002-2013021133003302-1033333233113301) |
| `timeouts.read` | [timeouts.read](resources--usb_policy--reference--group-001.md#canonical-3031112302130223-2011320310231102-3113101011002201-2333010010010020-2012102133121311-1112100010310122-2332212310301213-2122223121023020) |
| `timeouts.update` | [timeouts.update](resources--usb_policy--reference--group-001.md#canonical-3101312231130102-3301102000322112-1333333013132023-1313001220220131-2210311200133331-3123200011330111-1210213222332313-0202002031201201) |

<a id="canonical-3320231330031332-2111103103322211-1212323033212013-0110213130300113-0021012312002300-1222212221322113-0023133330222231-2322130331230220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `allowed_devices` properties

Breadcrumbs:

- [xcsh_usb_policy](../resources/usb_policy.md#canonical-3313033212003300-1002332031103020-1323001313000200-3003230033322232-3111031201110211-0323302100321032-3212002010220013-3201110212312232)
- [Property reference](resources--usb_policy--reference--group-001.md#canonical-0203202331120213-3031223102221310-3303300333312020-2233220001321002-3220003200013332-1202131202021320-0033221202012231-3201003230011131)
- allowed_devices

<a id="canonical-3210100222012221-0330131020230030-0010013320320333-2112303031121013-0031130002000321-1232220210203333-0001022223110222-1100313203012323"></a>

Type: `"object"`. list nested block, Optional.

Allowed USB devices. List of allowed USB devices.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.message.required_one_nonzero_field": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.message.required_one_nonzero_field": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
allowed_devices {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202330221213010-2302331231201113-1122310003133322-0213220112210220-2212203001130233-0013320021003303-1010320032131233-2033010122133301"></a>

### Direct properties for `allowed_devices`

<a id="canonical-1311133221000032-2033230222333022-1131322220111112-3100213100020133-3303223311311013-3032231102302033-0130310133230223-3313112321333330"></a>

#### `allowed_devices.b_device_class` property

Type: `"string"`. Optional.

Class. The class of this device.

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

<a id="canonical-2020002221200103-3231211200010212-1300222021221102-2333310333302310-1202020001031123-1121010332300200-1031222031230231-3212112013131022"></a>

<a id="canonical-0210112320132033-0322231120122003-0101133112020100-0331010322301123-0030130232103322-3111001011012130-0123132212103002-1203311122023302"></a>

#### `allowed_devices.b_device_protocol` property

Type: `"string"`. Optional.

The protocol (within the subclass) of this device.

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

<a id="canonical-0030110320100010-3102301200331303-2212202232213302-1011001003213001-2213213131033310-0220002203122032-1331211030223102-0132310303201313"></a>

<a id="canonical-3133311310131022-1303111323020011-1202020033300311-3012323313303320-0032211013133303-2231313333020321-2312003233221202-0031333330133130"></a>

#### `allowed_devices.b_device_sub_class` property

Type: `"string"`. Optional.

The subclass (within the class) of this device.

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

<a id="canonical-1303322100302331-1022110013223202-2223011133312020-2003231333220310-3000233130211213-2233212310012320-3213113212301030-1033001012323013"></a>

<a id="canonical-3122312312230112-1032022303022313-3123310313001322-3011100132312031-1322113030301121-0112010230021012-2131202110022002-3021302130103210"></a>

#### `allowed_devices.i_serial` property

Type: `"string"`. Optional.

Index of Serial Number String Descriptor.

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

<a id="canonical-3001102310103302-3013320202301033-2003303300212330-0203220313300112-1210000220210033-0231022012023103-2203332232331202-2022133210002332"></a>

<a id="canonical-1010132331010000-0023201301131111-3322021220023103-3312232302300032-2333213102132203-2002321303321102-1123122202310212-3323100133333212"></a>

#### `allowed_devices.id_product` property

Type: `"string"`. Optional.

Product ID (Assigned by Manufacturer) in hex.

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

<a id="canonical-0313130031103102-0300223100013113-2020012223021203-1321020230232121-3202123312310001-1133103120030130-3123311330031323-1303300203201121"></a>

<a id="canonical-1011030311101120-3100322101003232-1311310103133130-0310323323201300-3123111323110200-1222113021121231-2112302112033010-1223322003033312"></a>

#### `allowed_devices.id_vendor` property

Type: `"string"`. Optional.

Vendor ID. Vendor ID (Assigned by USB Org) in hex.

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

<a id="canonical-3300131102022123-1002203003021200-1000033132210032-2030322333112113-1100033310120313-2201312001021030-1003100222211132-0003010110030112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_usb_policy](../resources/usb_policy.md#canonical-3313033212003300-1002332031103020-1323001313000200-3003230033322232-3111031201110211-0323302100321032-3212002010220013-3201110212312232)
- [Property reference](resources--usb_policy--reference--group-001.md#canonical-0203202331120213-3031223102221310-3303300333312020-2233220001321002-3220003200013332-1202131202021320-0033221202012231-3201003230011131)
- timeouts

<a id="canonical-1010111211223301-1301202101022033-2021132031123233-3232021101221010-1022020211322303-3133333132322321-1221302030310302-3311213033131132"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323023001132202-3330111301033023-0002120131103032-1033232021213221-3213130323032312-0310333312110012-3120333011130032-2020202102201032"></a>

### Direct properties for `timeouts`

<a id="canonical-3022120022132223-2130020302133213-0020101330103013-3113232203000323-3302001333220013-3113312021312220-1011032323122221-0323310100123212"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0212103203120000-0101123122023211-0212200311331212-1133023001323330-1222012122313022-3311023222003002-2013021133003302-1033333233113301"></a>

<a id="canonical-3321022203321221-1311331300202001-1303022121111000-0110222210133213-0002232301032122-0103113133211303-2230210321312001-1330030012002203"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3031112302130223-2011320310231102-3113101011002201-2333010010010020-2012102133121311-1112100010310122-2332212310301213-2122223121023020"></a>

<a id="canonical-2221312321103123-3121211311131033-3100021123101120-2031003103221122-2312103011100212-1123312323130103-1133301332223012-2132211101331222"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3101312231130102-3301102000322112-1333333013132023-1313001220220131-2210311200133331-3123200011330111-1210213222332313-0202002031201201"></a>

<a id="canonical-1321332310300231-3331032120021132-1003303320222133-0030320303102020-0103231322010230-0202131003200212-1233013312212203-2111202322113030"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
