---
page_title: "xcsh_secret_management_access reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access reference."
---

# xcsh_secret_management_access reference

<a id="canonical-3322021123121120-1331101321210003-0101221331112313-0323133121220331-1012102110102230-1011002102022021-0102001122103112-1312103231303100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [where](resources--secret_management_access--reference--group-001.md#canonical-3010002013310032-1213220213323213-2123321300111132-1013020133110321-1213123032001233-2000133123322201-2232220303213303-3030000000122211)
- [where.site](resources--secret_management_access--reference--group-001.md#canonical-3033131013101111-0223012002311321-1033203233321312-1220103012131012-1122021131131000-0210211310301120-1113032001320320-1013332231323303)
- where.site.disable_internet_vip

<a id="canonical-1211223111003203-0122202211311203-1201012323200232-2302220322103103-2331123323320300-2312230301322003-0112300111210001-1210201223000020"></a>

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
disable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130032100011033-3031212022022130-0002133032312211-2121030011321011-0030122120011302-3022110113130203-3000010223121222-0222122322010021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [where](resources--secret_management_access--reference--group-001.md#canonical-3010002013310032-1213220213323213-2123321300111132-1013020133110321-1213123032001233-2000133123322201-2232220303213303-3030000000122211)
- [where.site](resources--secret_management_access--reference--group-001.md#canonical-3033131013101111-0223012002311321-1033203233321312-1220103012131012-1122021131131000-0210211310301120-1113032001320320-1013332231323303)
- where.site.enable_internet_vip

<a id="canonical-2100311023212210-1102232230223221-2210002032313010-2131131212330012-2102232111112312-3333211233323222-3310003313132003-1111323132210122"></a>

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
enable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101010002232021-1301001301301102-3131230311301333-1321121222232223-0003103223133301-2102121221031002-3120131021322003-0103202321200321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.ref` properties

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [where](resources--secret_management_access--reference--group-001.md#canonical-3010002013310032-1213220213323213-2123321300111132-1013020133110321-1213123032001233-2000133123322201-2232220303213303-3030000000122211)
- [where.site](resources--secret_management_access--reference--group-001.md#canonical-3033131013101111-0223012002311321-1033203233321312-1220103012131012-1122021131131000-0210211310301120-1113032001320320-1013332231323303)
- where.site.ref

<a id="canonical-1320013132032111-1220203212322220-1331112031220110-0300320231021031-1203011001211222-3313212110133300-1122231110332323-1332203203322223"></a>

Type: `"object"`. list nested block, Optional.

Reference. A site direct reference.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000203132132013-2012023333031030-0230300310332320-2010210210231021-3232012023303230-0210232121020012-3032302222001012-2002002011302313"></a>

### Direct properties for `where.site.ref`

<a id="canonical-2211020011200301-2211232311132221-2120030010323300-0310123121222120-1332321320213232-3320003302001100-2303230013030122-2221310230222012"></a>

#### `where.site.ref.kind` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2322211332302330-1030102303321110-0203030002311110-2003022221303103-0312012330312113-0212021223031212-1200000232022120-2322012030330010"></a>

<a id="canonical-3113130211200200-2212103200300200-0213202103323202-0012003223110312-2101300221111122-0232123321013132-0203313113300100-3031001020201121"></a>

#### `where.site.ref.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1011123023031303-2030122112200030-2322320010131213-1110213011332132-3010120022132030-1103013031003220-3113101030123130-1013222200310101"></a>

<a id="canonical-1110002233312021-0333103120130101-3201023331023000-3303102212122103-3203002000011102-1030330012110232-3232013323231001-2020313313301121"></a>

#### `where.site.ref.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0011113011123031-2221332322331311-0222003120302103-2310121311323002-1301212321011001-3310132210000013-2100221310030300-3131220121310220"></a>

<a id="canonical-1300022320310203-1131220011320022-3203212301102131-3000222113123233-3202220030221203-1121121332020313-1030022211120132-0032220221313312"></a>

#### `where.site.ref.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3230110020233231-1330121120010313-1133311312030003-0321210212131031-1102113121112223-0102310031032313-2132312020022222-1201320030312312"></a>

<a id="canonical-3331321322110022-2030221010301303-2010030111321332-3000001331012113-1300313123130223-0203320210113332-2230103033031300-1002333223111230"></a>

#### `where.site.ref.uid` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0212110313230212-3131103112200033-1001110002003112-0031121031212232-1013030302211120-1133033323213111-3210023102011233-2002113101203130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_network` properties

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [where](resources--secret_management_access--reference--group-001.md#canonical-3010002013310032-1213220213323213-2123321300111132-1013020133110321-1213123032001233-2000133123322201-2232220303213303-3030000000122211)
- where.virtual_network

<a id="canonical-1330333222230132-3013231110212212-1121013333200210-3000203023313023-0232130030113230-1323233331330302-2033333132120333-3231100131122213"></a>

Type: `"object"`. single nested block, Optional.

This specifies a direct reference to a network configuration object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ref")}
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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-2211211102032111-3332102111200231-0100022121211022-0022223313232132-1303222020120133-0332012012131210-3312232001313312-2230333113210333"></a>

### Direct properties for `where.virtual_network`

- [ref](resources--secret_management_access--reference--group-002.md#canonical-1322320232230321-2221100023313112-0322010332122300-2132310232223001-1323302332210030-1120322130000210-3212301112310200-2022220030200321): complete subsection reference.

<a id="canonical-1322320232230321-2221100023313112-0322010332122300-2132310232223001-1323302332210030-1120322130000210-3212301112310200-2022220030200321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_network.ref` properties

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [where](resources--secret_management_access--reference--group-001.md#canonical-3010002013310032-1213220213323213-2123321300111132-1013020133110321-1213123032001233-2000133123322201-2232220303213303-3030000000122211)
- [where.virtual_network](resources--secret_management_access--reference--group-002.md#canonical-0212110313230212-3131103112200033-1001110002003112-0031121031212232-1013030302211120-1133033323213111-3210023102011233-2002113101203130)
- where.virtual_network.ref

<a id="canonical-2201320010010221-3021211320203123-1003031200323322-1031003100323200-2012222223133132-0200331123122332-3212020010232103-3302320211103232"></a>

Type: `"object"`. list nested block, Optional.

Reference. A virtual network direct reference.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-2221222102111233-0003000110010201-1333331123112323-1220132300332212-2133333231333223-3100320222010201-3333021130003030-1131312000132200"></a>

### Direct properties for `where.virtual_network.ref`

<a id="canonical-0313333132220130-2300033102211000-0132001102133210-2221232022100221-0233133312031133-1322301321323333-0011131122322210-2102230313230213"></a>

#### `where.virtual_network.ref.kind` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3220102222301230-1032331331213212-3233213200031012-2303212023112123-0031230222130311-1220132312033200-3000111123223033-2100300102221120"></a>

<a id="canonical-1302202002002310-3111230333303130-3111201133132210-0011123321233221-2000322112023331-2001033120130023-1321033123230300-2120120313232203"></a>

#### `where.virtual_network.ref.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2032310033103211-1330233131321310-0320200000300312-1012100323311012-1303121231202102-0203310022331331-0113120201312130-1000132233303321"></a>

<a id="canonical-1331113130121000-3233323201130123-3101011230103012-0230323303030021-2201012302101322-0033103211331320-0012210301321202-3220230200302001"></a>

#### `where.virtual_network.ref.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2213332131002021-0122113310210030-3313020021222321-3131110313010030-2133231330310222-3031300233331302-3220013222133231-3303120200223222"></a>

<a id="canonical-3033111133102220-1132112102120010-2102001002013013-1322213000302322-1302032300110021-0102312033321230-2232213333123220-3010200121021033"></a>

#### `where.virtual_network.ref.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1112003301003013-3131233133321003-1330233213200302-3022112133102331-1321012113102211-1211213301330000-2110011102313302-3120032232033032"></a>

<a id="canonical-3103103332133212-2111311200311213-0122010211232020-1013122123301311-0121302022321231-0310132320301033-2220200011202312-0021230300003223"></a>

#### `where.virtual_network.ref.uid` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2303221010330132-3011100131031200-3210022223102300-3330131232033111-0133010103230122-1210333210023233-1332211003232113-3332131131021330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site` properties

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [where](resources--secret_management_access--reference--group-001.md#canonical-3010002013310032-1213220213323213-2123321300111132-1013020133110321-1213123032001233-2000133123322201-2232220303213303-3030000000122211)
- where.virtual_site

<a id="canonical-3130113011132233-2331020331120121-3331111101001111-1303101020231310-2222101222113212-2222231102212313-0133310332103031-0310123013100121"></a>

Type: `"object"`. single nested block, Optional.

Virtual Site. A reference to virtual\_site object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ref"),
  validators.ConflictingObjectAttributes("disable_internet_vip",
    "enable_internet_vip")}
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
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303330303213201-1310012123231122-3330300203323330-2013021223310320-1100322110303323-3301230222110111-3220132113211011-1122020022000220"></a>

### Direct properties for `where.virtual_site`

- [disable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-0110211131220030-1312222133333101-0231011323021020-1001032102223230-1331202030220222-2330200030002031-3000333212222131-1023123210231002): complete subsection reference.

- [enable_internet_vip](resources--secret_management_access--reference--group-002.md#canonical-1020013133130332-3230320032032223-0023121331121033-0330301002031212-3133331122203202-0331110223220330-1012212320031023-2221032000223332): complete subsection reference.

<a id="canonical-1021001110312133-0122321112320212-1120030101122101-1021232112012122-0232230311110002-1033311020311232-0123210112201301-3100122232231233"></a>

<a id="canonical-1320321300111012-2030321100120211-1031101023123133-3101122031131000-0121032111130111-2012122323321022-2221313123202210-0210023112233122"></a>

#### `where.virtual_site.network_type` property

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Additional upstream details:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
automatically and present on all sites Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE
is a private network inside site. It is a secure network and is not connected to public network.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
during provisioning of site User defined per-site virtual network. Scope of this virtual network is
limited to the site. This is not yet supported Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC
directly connects to the public internet. Virtual-network of this type is local to every site. Two
virtual networks of this type on different sites are neither related nor connected. Constraints:
There can be atmost one virtual network of this type in a given site. This network type is supported
on RE sites only It is an internally created by the system. They must not be created by user Virtual
Networks with global scope across different sites in F5XC domain. An example global virtual-network
called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VIRTUAL_NETWORK_GLOBAL","VIRTUAL_NETWORK_IP_AUTO","VIRTUAL_NETWORK_IP_FABRIC","VIRTUAL_NETWORK_MANAGEMENT","VIRTUAL_NETWORK_PER_SITE","VIRTUAL_NETWORK_PUBLIC","VIRTUAL_NETWORK_SEGMENT","VIRTUAL_NETWORK_SITE_LOCAL","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE","VIRTUAL_NETWORK_SITE_SERVICE","VIRTUAL_NETWORK_SRV6_NETWORK","VIRTUAL_NETWORK_VER_INTERNAL","VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](resources--secret_management_access--reference--group-002.md#canonical-2001101332220321-1321321021000023-3212321210003212-0322300332331012-3031230013223203-2112332321101033-2321331212311021-0110031110321300): complete subsection reference.

<a id="canonical-0110211131220030-1312222133333101-0231011323021020-1001032102223230-1331202030220222-2330200030002031-3000333212222131-1023123210231002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [where](resources--secret_management_access--reference--group-001.md#canonical-3010002013310032-1213220213323213-2123321300111132-1013020133110321-1213123032001233-2000133123322201-2232220303213303-3030000000122211)
- [where.virtual_site](resources--secret_management_access--reference--group-002.md#canonical-2303221010330132-3011100131031200-3210022223102300-3330131232033111-0133010103230122-1210333210023233-1332211003232113-3332131131021330)
- where.virtual_site.disable_internet_vip

<a id="canonical-0030100332100002-3302121213031121-2230211333321223-1112222211123122-0302330233200103-1302211012221101-3311002202132013-0131033021103322"></a>

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
disable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020013133130332-3230320032032223-0023121331121033-0330301002031212-3133331122203202-0331110223220330-1012212320031023-2221032000223332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [where](resources--secret_management_access--reference--group-001.md#canonical-3010002013310032-1213220213323213-2123321300111132-1013020133110321-1213123032001233-2000133123322201-2232220303213303-3030000000122211)
- [where.virtual_site](resources--secret_management_access--reference--group-002.md#canonical-2303221010330132-3011100131031200-3210022223102300-3330131232033111-0133010103230122-1210333210023233-1332211003232113-3332131131021330)
- where.virtual_site.enable_internet_vip

<a id="canonical-1313221330113203-2000003300122020-2220130013220132-3001310201212211-0113312031311213-0211323101220312-1132212122003130-0123002303203303"></a>

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
enable_internet_vip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001101332220321-1321321021000023-3212321210003212-0322300332331012-3031230013223203-2112332321101033-2321331212311021-0110031110321300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.ref` properties

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md#canonical-1223322133200022-0322132200020012-2200310323001030-1133012320033211-0202131012231312-3333312233130212-3222302310223221-0100000011223020)
- [Property reference](resources--secret_management_access--reference--group-001.md#canonical-0032020332001320-3010332010002310-2231312312333102-1213312333213211-2000323020020311-0210022012013130-0030301110003231-0300312031002123)
- [where](resources--secret_management_access--reference--group-001.md#canonical-3010002013310032-1213220213323213-2123321300111132-1013020133110321-1213123032001233-2000133123322201-2232220303213303-3030000000122211)
- [where.virtual_site](resources--secret_management_access--reference--group-002.md#canonical-2303221010330132-3011100131031200-3210022223102300-3330131232033111-0133010103230122-1210333210023233-1332211003232113-3332131131021330)
- where.virtual_site.ref

<a id="canonical-0223222002022120-3223101332310200-3331011310112033-0331012211101220-2131312031012110-1032232001323322-1013022133310331-0322013223213203"></a>

Type: `"object"`. list nested block, Optional.

Reference. A virtual\_site direct reference.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321110131311330-3302211221030002-1110310203010300-2310113020332112-2232310301123223-3122331313201011-1120021013131113-3113322103013331"></a>

### Direct properties for `where.virtual_site.ref`

<a id="canonical-0202320202032112-0023010100011211-3112303122313333-2122310221100011-3110313300000010-1321311020211002-3123300003123012-2102233113230020"></a>

#### `where.virtual_site.ref.kind` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3023013233232231-1211202113013011-1002323332231322-1102313332103131-0110012213311120-3230301131010332-1133032230111121-1002032020122210"></a>

<a id="canonical-0321130100201210-2230313223121132-2322110232132310-0110322133010232-3100110103122032-2122032211121302-1330312300032302-2223301323122230"></a>

#### `where.virtual_site.ref.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3032002121222000-3022302132233212-2022033313201212-1020030001331302-0213131100201333-2130031210222331-1132322213203131-3130001230032203"></a>

<a id="canonical-0113203012001130-1223203023133123-1033233102101002-0301021121101131-1212320311002122-2111231331011011-0203313202321211-1210133213003313"></a>

#### `where.virtual_site.ref.namespace` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0203301021331121-2022332330221013-2200221003223022-0020130021222232-2222023232001333-2333031110030212-0012111200323223-2122012310011111"></a>

<a id="canonical-1100100112203121-3113313110211110-3212001311332220-3131231003233333-1202321133301223-1301300301301330-2013100210303301-1112011231123232"></a>

#### `where.virtual_site.ref.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1232031202132310-3310132131022220-3110112032033303-1033031103222322-0023001311032032-2203321320332033-3300133310303213-2331333311302303"></a>

<a id="canonical-2221212200312233-1202120021000100-3020010211030323-3301112003012113-1232300001101212-0022220231232332-0331012120313212-3313200223303032"></a>

#### `where.virtual_site.ref.uid` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
