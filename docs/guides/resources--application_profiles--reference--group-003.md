---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

<a id="canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.http3

<a id="canonical-2002231130013121-1320001010323121-2331302320313322-2221020113232332-1113213211110201-1022230001113101-1132301132322222-3120032223001131"></a>

Type: `"object"`. single nested block, Optional.

HTTP/3 profiles.

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
http3 {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022032301122103-3322001111102012-1031102102010010-1301331202012031-1111021233011312-1120201231332232-3323231221233210-3002221110212032"></a>

### Direct properties for `virtual_server.http3`

- [client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-1013312311301220-1033133010113103-2312213023210210-1011131133330212-0333233011331131-2323002012312101-3003233002300221-0011220013121002): complete subsection reference.

- [http3_profile](resources--application_profiles--reference--group-003.md#canonical-0330103031220011-2233020132233102-3122223210303012-0113012332332203-2012302030003300-2232201011112130-0001003312201001-1332113123333323): complete subsection reference.

- [http_client_profile](resources--application_profiles--reference--group-003.md#canonical-0110323101223223-0010110032230301-0033333001003323-2301130020131120-1231020331020222-0332330333033103-1020112033001230-3210003311322323): complete subsection reference.

- [http_server_profile](resources--application_profiles--reference--group-003.md#canonical-0220231220211221-1300100113032203-1202030023031212-3101012201312312-0221100301100212-0101010110231132-1003301322301023-0020200203102112): complete subsection reference.

- [quic_profile](resources--application_profiles--reference--group-003.md#canonical-0310002212103133-2323121222001013-1130122010313012-3002210333000020-2312022322312202-0112101012133013-3332302221010313-1300310211101012): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-0100301033103101-0132322221003231-0130012010003121-2100011022032313-1300221101232323-3021101001012203-3000301233212002-2110021113331313): complete subsection reference.

- [tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-1001310203000122-0311102113233113-3302310000201022-0200230311200102-1233203113320233-0132232313301222-3211300233232121-0102200101100012): complete subsection reference.

- [udp_client_profile](resources--application_profiles--reference--group-003.md#canonical-3022120130332131-3313210302001322-3231201032200020-1312101002032320-1200300302023322-3130201220303033-0122013032031230-2112001220310130): complete subsection reference.

- [udp_server_profile](resources--application_profiles--reference--group-003.md#canonical-1001201132330211-3022323222212131-0312122213332223-0022033331023200-1322321300101231-3303303132220203-0323330033312310-0211000013331213): complete subsection reference.

<a id="canonical-1013312311301220-1033133010113103-2312213023210210-1011131133330212-0333233011331131-2323002012312101-3003233002300221-0011220013121002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.client_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-003.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.client_ssl_profile

<a id="canonical-1120110303201101-1301001103033203-2001103223211123-1023320032013330-1112100022313220-3011010013203212-3110122001132223-3111011223120333"></a>

Type: `"object"`. list nested block, Optional.

Client SSL Profile. Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
client_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102200031333231-0120100330223112-1311321230020123-0033113201311210-1310111310311031-3312321020310001-1222030221222000-2232221213021101"></a>

### Direct properties for `virtual_server.http3.client_ssl_profile`

<a id="canonical-2200200203103102-3133013030202230-1021203002003221-2110333022012111-3110302013321330-1331302101000003-1311221130300200-1333203223031111"></a>

#### `virtual_server.http3.client_ssl_profile.kind` property

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

<a id="canonical-3221132330131232-2332303212210212-0110223210030001-1211100200023300-2132000210320310-1133332222003221-0013313032201203-0121233203313220"></a>

<a id="canonical-3011231111220130-3020002213202330-0120010322203020-0132322332101310-1120021112113101-2003023313001002-1203100213120320-0031002330121131"></a>

#### `virtual_server.http3.client_ssl_profile.name` property

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

<a id="canonical-1031323230323132-3303012013032303-3003133221121022-2022000000202030-3203111232200222-1311120211232033-0021311303102132-3021101232132321"></a>

<a id="canonical-3303031121313123-2013330220323000-3011110203233122-1233200323321012-1003122211332030-0000012200120212-0133122102321223-1001031210100210"></a>

#### `virtual_server.http3.client_ssl_profile.namespace` property

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

<a id="canonical-1133102131011033-0311201313100103-3322323331312302-0223000103022232-2320103302021210-0011233002222300-3202321012210232-3221020223320230"></a>

<a id="canonical-1322010022103123-2102121330023001-2301231311231231-3101021302330312-1201023313133003-1232303202030010-1010002001003103-3103113013233102"></a>

#### `virtual_server.http3.client_ssl_profile.tenant` property

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

<a id="canonical-2133030102220021-3331133303233312-0010033013023023-1221011200002020-0232203313301020-0021310222210131-3203120311121221-3010003211132132"></a>

<a id="canonical-3323010031003132-0302213203001200-0210312223133232-1000222122233312-0131330013100123-1211002300231223-2012121012232133-3022332211233122"></a>

#### `virtual_server.http3.client_ssl_profile.uid` property

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

<a id="canonical-0330103031220011-2233020132233102-3122223210303012-0113012332332203-2012302030003300-2232201011112130-0001003312201001-1332113123333323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.http3_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-003.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.http3_profile

<a id="canonical-0322200003130310-2100000313011310-3310110111203133-0322333010302213-2113101012322222-3021021122212110-1320200333300022-1010100310322312"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for http3 profile.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http3_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212123102011211-1002201133001232-3030111313133310-2120002211331302-2022222331203100-1121232123220330-2022332102300110-0321011231033123"></a>

### Direct properties for `virtual_server.http3.http3_profile`

<a id="canonical-0223023331021000-2223031010030023-2311021301121001-1130222000011003-2303133002113130-0102302331300022-2202122200232233-1221133222302313"></a>

#### `virtual_server.http3.http3_profile.kind` property

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

<a id="canonical-1222200231311100-1101302000230030-2231211101101303-2310111330033023-2222021031303132-2212333002031030-0300302330332212-1333010032233210"></a>

<a id="canonical-1110230113202222-0333220012103333-0013213110202321-1131322313020010-2121130223232210-1303101220331011-1022011312330323-1211232131011003"></a>

#### `virtual_server.http3.http3_profile.name` property

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

<a id="canonical-0322112021132133-3021032010020012-2023301002300111-2202223300111312-0010102130331332-3312122322132123-3221222230131233-2231132201033101"></a>

<a id="canonical-3331001321203121-1132123313303301-1113303230212210-2123320333032310-0023102130203312-3322323210321130-3222121102302302-1003203221222110"></a>

#### `virtual_server.http3.http3_profile.namespace` property

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

<a id="canonical-1022330111013223-3102311311132110-1012010030003301-0120212110223111-3012101001002300-3101120312120021-0111103220303012-3131302201321022"></a>

<a id="canonical-0013223132123032-2202230220103131-2121003000330302-1321123231322000-3021001100213111-2023113303103133-3330311303102130-3132220032301311"></a>

#### `virtual_server.http3.http3_profile.tenant` property

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

<a id="canonical-1322323202021031-2012010100131333-2233031213321333-0123121021232122-3331333123011201-2133232131310333-0110012022023220-2030022300001301"></a>

<a id="canonical-1313131121123001-1313322122303010-1203011003302231-3130232030301320-3233111301112203-2212023211202012-2112002301103032-1021110021031331"></a>

#### `virtual_server.http3.http3_profile.uid` property

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

<a id="canonical-0110323101223223-0010110032230301-0033333001003323-2301130020131120-1231020331020222-0332330333033103-1020112033001230-3210003311322323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.http_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-003.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.http_client_profile

<a id="canonical-2111231120012330-1233222013020212-2013003201033331-3302211331120231-1313132022313113-0110012313212012-0131120302222113-2233021322200321"></a>

Type: `"object"`. list nested block, Optional.

HTTP Profile (Client). Client-side configuration

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1123021302300113-2133310210123210-0300331230121202-2210213112211031-2113023220232020-1202202011121012-3311303023221330-2100313212033222"></a>

### Direct properties for `virtual_server.http3.http_client_profile`

<a id="canonical-1232323312312320-0203200313010222-0202101012120212-3320021331120322-3212201230230303-0132321030302112-2332132213012301-1300221230320211"></a>

#### `virtual_server.http3.http_client_profile.kind` property

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

<a id="canonical-3133323121002120-0220100010031021-1132302332013120-0001301021111323-3222323212312302-0331213103333032-0133000103102232-1300003332222123"></a>

<a id="canonical-3131101300112321-3213102120210012-0300102110120301-2111331312302331-3111002201133322-0010220302102232-0010230021311223-3332110103200223"></a>

#### `virtual_server.http3.http_client_profile.name` property

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

<a id="canonical-1123133003010012-1021313333201313-1003311301312113-3133120023301203-2233121221331021-2203201101011233-3120010332132013-3211111122131330"></a>

<a id="canonical-3123002323320231-3212121321302222-2322221222013202-0202112130332130-2303203301233321-1311013221133331-1113013120320203-0311323121002313"></a>

#### `virtual_server.http3.http_client_profile.namespace` property

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

<a id="canonical-0321320130021102-3330121302313022-1001312023200033-3023322222003110-1221000110301300-3123222120321033-3000012022011133-0210012333233313"></a>

<a id="canonical-0233022212100313-0303301100111221-2010103003321003-2121112112010003-0312020310213012-2011310213202020-0031203020032000-3113213202012312"></a>

#### `virtual_server.http3.http_client_profile.tenant` property

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

<a id="canonical-1102013031132323-0002311212321333-0011103210021022-2100323313321233-1202311320131202-1012003310103111-2113202221132212-3030221000201110"></a>

<a id="canonical-0220210332010220-3221103302300221-2302003012330110-3010200221303200-2100220023031101-3122201212012010-3010303330002002-3001200102200221"></a>

#### `virtual_server.http3.http_client_profile.uid` property

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

<a id="canonical-0220231220211221-1300100113032203-1202030023031212-3101012201312312-0221100301100212-0101010110231132-1003301322301023-0020200203102112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.http_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-003.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.http_server_profile

<a id="canonical-2133130302111323-0311332113010102-2203021301212221-3311003113123201-2002013122221310-0203221012012023-0203001021303203-1012221323112022"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for http server profile.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001223031110210-3122301021313210-1033233113113000-3022000312111020-0131111131210121-0211202013202312-1132320333113100-3200111131132123"></a>

### Direct properties for `virtual_server.http3.http_server_profile`

<a id="canonical-2302321123330311-2011003131132333-0323321010202100-0021123221310302-3032012110320221-0331213222200032-3232003103320030-1120123022311111"></a>

#### `virtual_server.http3.http_server_profile.kind` property

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

<a id="canonical-2213310223202332-0121320111231201-3003031213003203-2333233003010103-0313000100201311-0013013223110221-2223330033223010-1002031312311323"></a>

<a id="canonical-0303032201211133-2211323131033100-2120211120331130-0202002023320031-1102331010330111-1120030221330202-1321123333032301-0201103210312213"></a>

#### `virtual_server.http3.http_server_profile.name` property

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

<a id="canonical-1022320011002333-0023330332201230-3113033012211223-1020201320031100-3012002112022130-3031220222302121-0333011313030101-3013021320020321"></a>

<a id="canonical-0211031211110002-0333100203022102-3322322103110202-1001233303110210-3113222132211001-2311121231020023-3130013111101132-3201312010311010"></a>

#### `virtual_server.http3.http_server_profile.namespace` property

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

<a id="canonical-2031231210110123-0203123100120331-3000311221212021-1120222023232111-1200323023213332-3030321303333012-0210333021313222-1031111112113330"></a>

<a id="canonical-1231200030311213-1130121301300201-3310330132221200-1332320303313333-3023021122130013-0223022113321232-0302302001000330-0011023300021212"></a>

#### `virtual_server.http3.http_server_profile.tenant` property

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

<a id="canonical-3330310132210002-0302201100211032-3012312202022322-0211103022100022-0310313332103213-3302222311122223-3312010012100002-0030313103221213"></a>

<a id="canonical-2331210201010000-0020020111122322-2032330112133331-0111302301021313-3020010101121023-1330301232333032-0102223301322133-0213300100231001"></a>

#### `virtual_server.http3.http_server_profile.uid` property

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

<a id="canonical-0310002212103133-2323121222001013-1130122010313012-3002210333000020-2312022322312202-0112101012133013-3332302221010313-1300310211101012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.quic_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-003.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.quic_profile

<a id="canonical-1302232321131103-0021111103200313-2331011300003103-0001032113331103-2123300232330232-0202033332003011-0301212020002233-3203103023000203"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for quic profile.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
quic_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202002123003133-2230221221322032-0333132300112031-1332120123010222-1311002223012230-3321012231031113-2110110232221103-0213222302212021"></a>

### Direct properties for `virtual_server.http3.quic_profile`

<a id="canonical-2103301310301303-2213000303330332-1210031331121212-0211120322231010-2002323220303330-3322331223133220-2021000022032210-0200001200132003"></a>

#### `virtual_server.http3.quic_profile.kind` property

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

<a id="canonical-1011131312312103-2332110030312221-1332211112000033-2122120023202330-0030201011213303-1223101230202303-1102102121010022-1013202121303000"></a>

<a id="canonical-1312311021232211-0102303020122102-3233301021102103-1233120122122030-3213003213030311-2100103003300102-0331001201221102-2232021330212030"></a>

#### `virtual_server.http3.quic_profile.name` property

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

<a id="canonical-3231202020111232-2023212123121022-0013223123210310-1000103102302133-1001211103302122-1233230200223323-3022002231312230-2320032132321003"></a>

<a id="canonical-0233001200122111-2100331222022032-3011323101101300-0201200113202012-1332220222323233-1013223121132013-2231103333122103-3023323330201111"></a>

#### `virtual_server.http3.quic_profile.namespace` property

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

<a id="canonical-0331231031100010-3132103210320110-0022122310201313-3321332303100102-3011332133013120-0320112123112312-0310101003223203-1010123301102220"></a>

<a id="canonical-3231011032113113-0032333200113213-3221212331122203-3321113213100323-1102003312020231-1022033101022320-2032223201222123-2222031121331333"></a>

#### `virtual_server.http3.quic_profile.tenant` property

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

<a id="canonical-1010230110221332-1202213231322301-2211100303001133-1233022232001102-3013100232223000-1131211102001003-3020100002231201-1313212101102330"></a>

<a id="canonical-0211020011332122-1013000321213322-0033223002013321-3133230130111230-3020200001203112-1023100030321331-3200232233211201-3132000030002212"></a>

#### `virtual_server.http3.quic_profile.uid` property

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

<a id="canonical-0100301033103101-0132322221003231-0130012010003121-2100011022032313-1300221101232323-3021101001012203-3000301233212002-2110021113331313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.server_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-003.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.server_ssl_profile

<a id="canonical-3011020032120132-2302133210113203-2231330200220020-2011300122312220-2130310201130222-1203123211131010-0021313120101311-3230333210021311"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for server SSL profile.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
server_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320332011322303-2220033111111320-3303013200232311-0121000001202300-2333232012331213-1021122301231131-2000211100030332-0331312012111013"></a>

### Direct properties for `virtual_server.http3.server_ssl_profile`

<a id="canonical-3123101132120311-1121210230331331-1233123311113011-3322112010310111-0002012322121331-3020203322032132-3113233321330202-1312213011201132"></a>

#### `virtual_server.http3.server_ssl_profile.kind` property

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

<a id="canonical-3303131301110201-1021303100323033-0101202321000101-3210132321223212-3200121332322323-1033002303110310-2002211221031022-2132320312200213"></a>

<a id="canonical-1002202010233121-1201300201130023-0123301300111200-1223011102212302-3131230213321112-3312213223121010-1021322311201320-2130222032331222"></a>

#### `virtual_server.http3.server_ssl_profile.name` property

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

<a id="canonical-3330222313300323-1300303111103222-1212310233133110-2310133303300312-0010323332131012-3320202010313020-1311231111223120-0230231231131333"></a>

<a id="canonical-3123233012331311-0030102220332212-2210230310030320-0201300332132321-0131012112123102-3002330302100333-2121013003330010-1132132203200312"></a>

#### `virtual_server.http3.server_ssl_profile.namespace` property

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

<a id="canonical-2330313112311101-2323331322211132-3032313321111113-2133000113202002-2300220332223110-2230130012233031-2010002010020222-0003033320100202"></a>

<a id="canonical-1130021220131201-3133113103133100-0102322221101301-1021222220203302-3133013133030111-0100231020200300-1112012233303132-3132222320322320"></a>

#### `virtual_server.http3.server_ssl_profile.tenant` property

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

<a id="canonical-0203003212200112-3310300301310211-2323330302211331-3322110132233122-1003021030333111-1012130322132301-0003003311211121-0001103031022320"></a>

<a id="canonical-0200132200320020-2100302022210220-0130100121303213-0023111323121203-3300103211030111-2032330112020302-2212320320101122-3301003121002013"></a>

#### `virtual_server.http3.server_ssl_profile.uid` property

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

<a id="canonical-1001310203000122-0311102113233113-3302310000201022-0200230311200102-1233203113320233-0132232313301222-3211300233232121-0102200101100012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.tcp_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-003.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.tcp_server_profile

<a id="canonical-2112022202030023-0100133022230321-0003322233201012-1233302133322131-1221012120113011-1300301322102201-3220232232012300-2313211223230131"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for tcp server profile.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tcp_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102202103102213-1121101130031331-0100221313223312-1112200213313123-1020011301331012-3331032102111203-1331100011010132-1301101123323021"></a>

### Direct properties for `virtual_server.http3.tcp_server_profile`

<a id="canonical-3233212222020113-2013030122200212-1312021020332233-2213131313303312-2022033212012200-1332231110032222-2330112220031320-1203213110201331"></a>

#### `virtual_server.http3.tcp_server_profile.kind` property

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

<a id="canonical-1313331100013012-1030031223311021-3012301331111302-3102010301010222-2331303203022011-0123303021133010-1311032111222033-1321301300000111"></a>

<a id="canonical-3330333011323111-2110301122132231-3120300312022011-2010210102122031-3113032031300120-2130210201321312-3221013220023021-2023112020320202"></a>

#### `virtual_server.http3.tcp_server_profile.name` property

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

<a id="canonical-1222310012213002-3123101310310103-1303100221030210-3010102203123330-3223033230301311-3312210132221300-0223301222313211-0320132001222301"></a>

<a id="canonical-3112122113332302-1023203210203232-3232000123013121-3331223122002301-3333121301002312-3000333130100012-1121313200012013-2211202233332302"></a>

#### `virtual_server.http3.tcp_server_profile.namespace` property

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

<a id="canonical-3302321003121013-1303320221101203-3330120030100003-3322220212132311-2131210331022232-2123000202110111-1112211223002332-3120001103303301"></a>

<a id="canonical-1030232320110000-1301002331131311-0023120301322332-1201013301130030-1331003222030001-2113201010313223-0011121232233001-1110330303101111"></a>

#### `virtual_server.http3.tcp_server_profile.tenant` property

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

<a id="canonical-2011231123300211-1110331203033123-3133331233211012-3302200221020313-0202100131330100-3311102111013213-2021210000213321-1302330013220212"></a>

<a id="canonical-1220100303321023-0221032211031330-0132313200122313-0011213101120233-3301302103213113-0023003012320022-0232231311313222-3231310012220333"></a>

#### `virtual_server.http3.tcp_server_profile.uid` property

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

<a id="canonical-3022120130332131-3313210302001322-3231201032200020-1312101002032320-1200300302023322-3130201220303033-0122013032031230-2112001220310130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.udp_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-003.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.udp_client_profile

<a id="canonical-1301232231102003-3331013000233123-3210231321202022-1323203122210233-1102103310230133-3132113123220232-2201022010231232-0120001122103011"></a>

Type: `"object"`. list nested block, Optional.

Protocol Profile (Client). Client-side configuration

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
udp_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030120101122012-3312203010212010-2201320033002312-2220002113231021-1033123113103101-0001323102320120-3010103321120313-3121131033233122"></a>

### Direct properties for `virtual_server.http3.udp_client_profile`

<a id="canonical-3013223112100102-3013001111311323-1231002102311223-0011023031023330-3213100330021322-0112130103212321-3230013311130223-1203212231210101"></a>

#### `virtual_server.http3.udp_client_profile.kind` property

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

<a id="canonical-1121320020312332-1110231313033021-3303210021322300-2102203323220310-1021020102232111-3231001101021221-1133231201032011-1322313032310000"></a>

<a id="canonical-0221332313000002-0102212310111011-2313013010203032-1230031303230032-1301030132033012-0120000031212311-2112211010333113-2123012000231002"></a>

#### `virtual_server.http3.udp_client_profile.name` property

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

<a id="canonical-2322102203111312-1033312323301130-2003312130120211-0130002323321213-2010113120023333-3330323300321233-0112113203200133-1030200000111310"></a>

<a id="canonical-3222213110332200-1022113203102023-2002221333330312-0003123302020321-3333120310330003-2203230321122113-2200332231313301-0133122301311303"></a>

#### `virtual_server.http3.udp_client_profile.namespace` property

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

<a id="canonical-1103011102221312-1320032331123122-3213233021233110-0223223312220023-2120020230011321-2031131321131120-0222301222320002-3210233123322330"></a>

<a id="canonical-0012321123322033-0001321311312311-3202221232200230-3013232002122312-3002321132030013-3030211023021211-1332011203103200-1212123002112100"></a>

#### `virtual_server.http3.udp_client_profile.tenant` property

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

<a id="canonical-1011023113133001-0310312311131200-1032221021310230-0213203012203220-0232131313113202-0030201032312003-3311211120310032-3202202133302212"></a>

<a id="canonical-1232021301210313-2132131233111200-2121020103022333-0003310113313231-1023100021123330-1212011100202032-0311131201021330-1201002012200231"></a>

#### `virtual_server.http3.udp_client_profile.uid` property

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

<a id="canonical-1001201132330211-3022323222212131-0312122213332223-0022033331023200-1322321300101231-3303303132220203-0323330033312310-0211000013331213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.http3.udp_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-003.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.udp_server_profile

<a id="canonical-2101303221030311-3333130023032321-2133221203131113-0122010021103001-3012322030021323-1021120122000203-0313111000212222-0032022121330320"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for udp server profile.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
udp_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030023331101110-3112133133031022-0230312012112201-2312310102311310-3232223312002310-0220230101221131-1120211023022223-2032011212001223"></a>

### Direct properties for `virtual_server.http3.udp_server_profile`

<a id="canonical-0000131221231020-1310212223312213-0331013132010110-1311233302130323-3013200313320131-3223120132012132-1120120113023012-1221233222201021"></a>

#### `virtual_server.http3.udp_server_profile.kind` property

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

<a id="canonical-2300230332231111-1013120233103012-0203213202332000-1232000321230111-2212310102023203-2001023000301211-3120203330302323-0330122212330033"></a>

<a id="canonical-2202011320330111-3331211210323001-1211123301010003-3133320012201011-2030320312313233-2020222133132311-2301303120300311-1103302231300311"></a>

#### `virtual_server.http3.udp_server_profile.name` property

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

<a id="canonical-0003000132010202-3210313322011303-2310323320002310-0322300310101032-2030311333313110-2021322120301300-3203311300000122-3021203332112323"></a>

<a id="canonical-0202020332300322-2233113131113203-0013111102310030-2010221323112201-2201130031103121-1031210201113111-2002112300102312-0321221022120203"></a>

#### `virtual_server.http3.udp_server_profile.namespace` property

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

<a id="canonical-2222313230003331-2033032001033011-0320333200333303-1130122013322021-0020211312201313-2000301122120130-3122220220333010-2012302012302320"></a>

<a id="canonical-1111333233112111-3231331103103231-0123033300010103-0222031331020231-2201113311321120-0223213101330310-1023211131202221-1133322213120121"></a>

#### `virtual_server.http3.udp_server_profile.tenant` property

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

<a id="canonical-0203011010032210-2131222022303211-3221221203322223-3230300330201001-1010302121021133-0110123110333002-2232000111331000-0012232023221200"></a>

<a id="canonical-2033301331023021-1322232132132202-3010033210320220-2303032022223121-1021122030332220-3231333030113112-0231323201322232-1330023033123322"></a>

#### `virtual_server.http3.udp_server_profile.uid` property

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

<a id="canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.https

<a id="canonical-3001023313121321-0102311133101033-0023030120103121-2330130013111333-2123113113011130-0202212203232022-1123300210111113-1023120132333222"></a>

Type: `"object"`. single nested block, Optional.

HTTP profiles.

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
https {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222112201011130-0020101021012113-3102201212030313-1333233111132311-0201003113222210-2300011322202221-0221111020233001-3221122303221003"></a>

### Direct properties for `virtual_server.https`

- [client_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-0122200003233301-3311211133001102-3103233332021321-0232000221020331-1000213223212230-3220313123122130-2332212200120320-0313001001103301): complete subsection reference.

- [http2_client_profile](resources--application_profiles--reference--group-003.md#canonical-0203301111203233-3202103203121002-0010112113330300-3100300003203113-3210120202012123-3100113221120002-2200223020221303-0211310322130123): complete subsection reference.

- [http2_server_profile](resources--application_profiles--reference--group-003.md#canonical-1132311103330130-1222222201100031-2313002121311120-3101210200020120-3002120332022011-2300231020301001-0312203101301233-2303010131110231): complete subsection reference.

- [http_client_profile](resources--application_profiles--reference--group-003.md#canonical-3212010031101123-1131101022132221-2211033221011012-0033101123321221-0020131203013333-3333213132110323-0221100301132011-1321322002030113): complete subsection reference.

- [http_server_profile](resources--application_profiles--reference--group-003.md#canonical-1122030322310031-0231311121030033-1302200331211212-0333220201121310-1210223321101021-0222202330113220-0221202102111003-1233230322202223): complete subsection reference.

- [ocsp_profile](resources--application_profiles--reference--group-003.md#canonical-3232210112131231-2320022120332210-3322123003301021-2111110123131030-2101301231333021-2232331231321203-3321222020010312-2320211131133203): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--reference--group-003.md#canonical-3210100200202213-3031331313010323-0132303233310220-2032230221331212-1102102133301311-2123021010322233-1231122120102031-3013320331102233): complete subsection reference.

- [stream_profile](resources--application_profiles--reference--group-003.md#canonical-3212022302301301-3103022201102322-3321120203323221-2023331133030122-0110010301200323-3321303211010033-0202220000313202-2322031102333330): complete subsection reference.

- [tcp_client_profile](resources--application_profiles--reference--group-003.md#canonical-3331122100212103-2103303202311112-2323020330102212-3020211133121232-2333002231201330-3111113212023302-1330211031211310-2112013332011103): complete subsection reference.

- [tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-0022210002333001-3112222323110013-3020033033022233-1111120310313331-1110213213313203-1331130321003201-3100013203212102-1123123203203132): complete subsection reference.

- [websocket_client_profile](resources--application_profiles--reference--group-004.md#canonical-1222221330201310-3101213130310000-2333211331100210-2102023100223033-0030012112211323-2023233213123221-1220103312330003-1013000000133131): complete subsection reference.

- [websocket_server_profile](resources--application_profiles--reference--group-004.md#canonical-0211103332301002-3311220121001320-2202321013331102-0113223312313112-0102220010300023-0213310332210133-3211320101132301-0213110231333311): complete subsection reference.

<a id="canonical-0122200003233301-3311211133001102-3103233332021321-0232000221020331-1000213223212230-3220313123122130-2332212200120320-0313001001103301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.client_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.client_ssl_profile

<a id="canonical-0310121121023300-0202320010121122-2230211332011001-3202032212312211-1102322000011001-0120031112230010-0003210133022323-3330320012220231"></a>

Type: `"object"`. list nested block, Optional.

Client SSL Profile. Client-side configuration

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
client_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020020112330332-1113021322333010-0201032120300010-0212320000232210-1332210213213301-1033022210110230-0030222220202020-0333003310012321"></a>

### Direct properties for `virtual_server.https.client_ssl_profile`

<a id="canonical-3032231110110323-1033321021030322-0123301303001130-3223330130020011-3221120132020132-1223221002031223-2210010320200032-3313013112301010"></a>

#### `virtual_server.https.client_ssl_profile.kind` property

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

<a id="canonical-0003003301000222-0301100210030302-3132310302032101-0220320330301001-3021220010212333-3131332111311002-0220322002310223-3032220023200211"></a>

<a id="canonical-2233222220310103-2110010301333300-2201020021123120-3021110313111220-2203010222031013-1130110030201012-0230213202200131-0303333302232230"></a>

#### `virtual_server.https.client_ssl_profile.name` property

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

<a id="canonical-1311323220110203-2133123021312233-2023032031232113-3203001331011010-0132130330102233-1101300101220331-1031203301132233-3201010223100210"></a>

<a id="canonical-0313023231300030-0101023033320230-0031331210320021-1313123311330032-3112022022332231-0022012131231011-0032223321233312-2333021220212312"></a>

#### `virtual_server.https.client_ssl_profile.namespace` property

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

<a id="canonical-2022202320302123-1133322223320312-1230332000322123-3022112012223322-2212321330310320-1100322000232323-2221111330332110-0021110023012030"></a>

<a id="canonical-1313120233003232-3111322022111100-3011331001301312-2202121130232122-2300012120111310-3002033011001301-2112311113231311-0231303031311013"></a>

#### `virtual_server.https.client_ssl_profile.tenant` property

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

<a id="canonical-2223310102001111-0022200133320100-0231300013011313-0311010133030002-3211032220131323-1313213312231102-3331320220023212-2210302232123231"></a>

<a id="canonical-3131002031210313-2202330323200032-2010221132320121-1000323301231201-0102300020301011-0120132132111120-0331333112031310-2103221310111020"></a>

#### `virtual_server.https.client_ssl_profile.uid` property

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

<a id="canonical-0203301111203233-3202103203121002-0010112113330300-3100300003203113-3210120202012123-3100113221120002-2200223020221303-0211310322130123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.http2_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.http2_client_profile

<a id="canonical-1003123300312211-1200121022121330-1013021323113111-0210100223202000-2320003131022321-2223221010130012-2100302320320122-2133311112211120"></a>

Type: `"object"`. list nested block, Optional.

HTTP/2 Profile Client. Client-side configuration

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http2_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203221322100312-1203320200122331-0032331123000012-3133023020033321-1100320301132332-1122011111011000-0312130022033120-1322000212322011"></a>

### Direct properties for `virtual_server.https.http2_client_profile`

<a id="canonical-0020210321321021-2030112332221211-2120011022223230-0311333010303030-2330123101122300-1220123112032232-1123022003131133-1331113200330021"></a>

#### `virtual_server.https.http2_client_profile.kind` property

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

<a id="canonical-3122002103220133-0131023110022301-0321122302302113-0003311233310333-3332032033033013-1032010113313330-0322012321230310-0300222310200331"></a>

<a id="canonical-1313003010320201-2230320110221331-2200311312010010-3311021021331001-0100011212313200-1332111200322000-2331101130232303-2122020122201020"></a>

#### `virtual_server.https.http2_client_profile.name` property

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

<a id="canonical-1211331331323102-0210322300233311-3033002012102113-1121231111200120-2102220120111131-1221011020303012-2102131130211330-0123321132001200"></a>

<a id="canonical-1331131212020232-0301231110103012-0323131110111330-1222122033132311-0212300113212003-3313333021231303-2021201122031220-3013213001031102"></a>

#### `virtual_server.https.http2_client_profile.namespace` property

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

<a id="canonical-1322300310122122-0003132023032102-1033301113123323-2021133331010012-1131013212021122-2333023300110121-1203302301210210-1302330200320033"></a>

<a id="canonical-3323000210231311-1200113221021112-0322111100011020-2230020203123331-3322132301312220-2131112131220332-2023230313321301-0110202102100010"></a>

#### `virtual_server.https.http2_client_profile.tenant` property

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

<a id="canonical-1103130313031020-3223210221132310-2031233210210213-2200232120203323-3001232212322123-3123211010323320-1302003030312312-1001020323122201"></a>

<a id="canonical-0331120220212311-3011322321131032-2332011232033012-0122213012113131-1313133201131133-1321302113333300-3201220202022001-0112011101300320"></a>

#### `virtual_server.https.http2_client_profile.uid` property

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

<a id="canonical-1132311103330130-1222222201100031-2313002121311120-3101210200020120-3002120332022011-2300231020301001-0312203101301233-2303010131110231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.http2_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.http2_server_profile

<a id="canonical-1301311030022121-3200103330220220-3210233112201320-0010303121231223-0313222113100020-3112130201100122-3322123010223020-1002303231002203"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for http2 server profile.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http2_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0131311230132033-0030223233120330-1000020123131121-3113310100323231-0223130103230201-3310221031102033-0302221330130131-3000022013122120"></a>

### Direct properties for `virtual_server.https.http2_server_profile`

<a id="canonical-1113211231233032-0222313101010330-2331111002231120-0232200310102012-3230200321113121-1320033113303311-1211123020303130-0123033131000123"></a>

#### `virtual_server.https.http2_server_profile.kind` property

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

<a id="canonical-3023233033020131-1321311130111021-2331012103212023-0013301133200031-2211210012330022-3322300200021012-0021231010122232-3021231002013233"></a>

<a id="canonical-0203300221200331-1100130033110313-0012230032321032-1121032321023222-1030331210320302-0333002020230332-2032313130030122-2220031323011323"></a>

#### `virtual_server.https.http2_server_profile.name` property

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

<a id="canonical-2301320332123011-1301313120311322-0000211013023311-0223020320313313-1002031022231321-2133313022023200-2103102102320121-1110311123032110"></a>

<a id="canonical-3302211233203322-1222100030102223-3113320321223110-0012222211313021-1211311112002311-2212133301012120-2133220312300122-1022102012112033"></a>

#### `virtual_server.https.http2_server_profile.namespace` property

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

<a id="canonical-0320221013202203-2010221331110003-3203132200302111-3020303301233130-3202231323130001-2023012112122322-0020320110203112-1020232301110003"></a>

<a id="canonical-3101310121322123-1000222231312312-3331131322123203-1101300120310320-3210333123232103-0303020001003030-0322030123022222-0032201130223230"></a>

#### `virtual_server.https.http2_server_profile.tenant` property

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

<a id="canonical-0103301103002230-0031310120131100-0331022101323301-3211200320000320-0110100321023120-2333130113200321-1310221113101303-3300232120132233"></a>

<a id="canonical-2102020321301010-2233322120130210-3301101223211201-0133010002200020-2131302133012031-3332313311302022-1032331123303320-0101321320220220"></a>

#### `virtual_server.https.http2_server_profile.uid` property

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

<a id="canonical-3212010031101123-1131101022132221-2211033221011012-0033101123321221-0020131203013333-3333213132110323-0221100301132011-1321322002030113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.http_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.http_client_profile

<a id="canonical-0032323313122011-1113013231233012-1211201132213211-1120003003232021-1113331031023332-2003213113121200-0303213030113103-0221312101011203"></a>

Type: `"object"`. list nested block, Optional.

HTTP Profile (Client). Client-side configuration

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132321303310022-0100012322233021-1112122013130313-1322203230311001-1320113312331101-0030111000121232-3022222133210020-0321232221303120"></a>

### Direct properties for `virtual_server.https.http_client_profile`

<a id="canonical-2023113210200000-0200112222213111-0231203130300021-0131103012301303-1233202032321121-2222322012120130-1300000323021303-2103213221211221"></a>

#### `virtual_server.https.http_client_profile.kind` property

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

<a id="canonical-0303323220213020-2000222113103230-0223320012112210-3230021003123010-3131020110011331-3001220030012321-1131211301030010-1331231133232233"></a>

<a id="canonical-0100010221223002-3202330232023011-3323302201320130-3123202133113033-2001122230002332-2030101222300133-1100012313033132-2300230213331233"></a>

#### `virtual_server.https.http_client_profile.name` property

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

<a id="canonical-0022233331201131-2312323210310333-1233301230322012-3231321322022312-3120100003010302-3002200323200103-0220321012113221-0301132013330101"></a>

<a id="canonical-3100231323123321-3022112023030132-1120213313112212-3320330212103002-3011223103102010-3112212003123122-3202333231210111-2221231100331111"></a>

#### `virtual_server.https.http_client_profile.namespace` property

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

<a id="canonical-2021313123201301-3203133211212100-1320023310123102-0211302230010300-0323012020300220-2130133103331020-0313321003222233-1103022003202320"></a>

<a id="canonical-3202303112300003-1003023323201133-3031222032202110-0202200021103102-2221121100221013-2110120120121131-1031122232201020-0002313100320332"></a>

#### `virtual_server.https.http_client_profile.tenant` property

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

<a id="canonical-2233033110310331-2303202013321113-3112300221103022-0223232133221001-0130323302312003-2013100003333210-1103321033320232-0322230010012321"></a>

<a id="canonical-0203330103213220-1311113320123211-0132101032123300-1001312011232001-3111000123301103-1122013310203212-0112231220013201-3213310300311021"></a>

#### `virtual_server.https.http_client_profile.uid` property

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

<a id="canonical-1122030322310031-0231311121030033-1302200331211212-0333220201121310-1210223321101021-0222202330113220-0221202102111003-1233230322202223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.http_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.http_server_profile

<a id="canonical-0022113000123221-1310311301020202-3033300230312220-3112103313120203-3220123203100231-2233231201330130-2013212103023230-2303102003312012"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for http server profile.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3220213223111123-0020133333103220-3133032201212311-0323310220310132-1213213230022132-1223333223122211-3020010210320002-1210213121023322"></a>

### Direct properties for `virtual_server.https.http_server_profile`

<a id="canonical-3203212101112002-2302130332100001-2312322011031023-3123133031311232-0031333131031233-1102130110320233-0002102301013232-2202131022300202"></a>

#### `virtual_server.https.http_server_profile.kind` property

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

<a id="canonical-0203320031122111-0012323323321330-0232021203130212-1333013233130302-3221201132300102-2112331122111320-1320322122120031-1303201303101220"></a>

<a id="canonical-1322231300100322-0203233113210331-3213133323312202-0112300310012230-1202332202310202-0233100131310000-0331233201123002-3330331101321101"></a>

#### `virtual_server.https.http_server_profile.name` property

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

<a id="canonical-0301311010130332-2013223110331003-2302223131330010-1000133101312133-3033012102320203-1313133132032131-2300021120322213-3333203323131311"></a>

<a id="canonical-1011103033130102-2303032133010230-2010323030020002-0022023223330123-1221013000123232-2300333130302323-3121000312033110-2111120200223103"></a>

#### `virtual_server.https.http_server_profile.namespace` property

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

<a id="canonical-1033322130333333-3310212312100302-1222121103220102-3011321303001233-2010212113130103-0210303033332102-2001213320123001-0120101101032032"></a>

<a id="canonical-1030033011102312-1230011310331213-2202223133311321-2123003333211230-2222301221330201-0023230222311201-0230332021120203-3233211230030221"></a>

#### `virtual_server.https.http_server_profile.tenant` property

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

<a id="canonical-0332102231332221-1010303130330101-1123001332332033-2302311132213211-2233102332022310-3332221230333310-1033323303233022-0030111223231310"></a>

<a id="canonical-1221200322001112-2023300130001220-0330033112000221-0320322221013000-2202201323311121-0330032021012213-0202102001100212-1033131322220230"></a>

#### `virtual_server.https.http_server_profile.uid` property

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

<a id="canonical-3232210112131231-2320022120332210-3322123003301021-2111110123131030-2101301231333021-2232331231321203-3321222020010312-2320211131133203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.ocsp_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.ocsp_profile

<a id="canonical-1132012023312201-2221332332121311-0112130333130300-0020032030200233-3313013312303212-1330022213030112-2313113303201211-0103121230002203"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for ocsp profile.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
ocsp_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3331023303222110-2213211022012122-2300312313221210-2322230333231131-0100012022032313-2200300212211033-2010313233122002-3030100031333110"></a>

### Direct properties for `virtual_server.https.ocsp_profile`

<a id="canonical-0331303310021132-3120032200223020-3333010133320103-0333233232201121-2322310013000031-2110223020321130-2201332003010223-2321120231031112"></a>

#### `virtual_server.https.ocsp_profile.kind` property

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

<a id="canonical-3230223322021133-0322233321110023-2103131011132232-0230121003000232-3033332313333333-1023332232331022-2311212200220013-3003223211220211"></a>

<a id="canonical-2102013200201200-0030133102100100-0203302222031113-3011000301121202-2311102301331212-2221200330010213-2132332010333332-1321231103031132"></a>

#### `virtual_server.https.ocsp_profile.name` property

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

<a id="canonical-0200000111002232-0131311002102013-1013002000322012-3301003301302030-2210110001102331-0232022112113211-1320020222210102-0323100232321312"></a>

<a id="canonical-3022330332323032-1012330023323013-1332113320003111-1303233120032221-3233120301330321-3320013200220311-2320112222321122-3001302030012130"></a>

#### `virtual_server.https.ocsp_profile.namespace` property

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

<a id="canonical-2022323331132330-0120331103211023-0233322101011010-0311120122030322-2223033322002110-2002132003022002-2033223321312220-2011311222000203"></a>

<a id="canonical-3133331112131230-1012001031211103-3313223211313122-3001030302321112-1001301013230302-0000202220000112-0111300220233013-0012032023230112"></a>

#### `virtual_server.https.ocsp_profile.tenant` property

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

<a id="canonical-0201011233111201-3003223310323100-0301020223011200-0223333112123322-1120321301331013-0222011321001200-2310330030023103-0321301122013302"></a>

<a id="canonical-1311121011132013-2000323122202120-0210002132221021-2222103213311030-0213110021202203-1110220100131311-1121302020201022-1210222130100200"></a>

#### `virtual_server.https.ocsp_profile.uid` property

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

<a id="canonical-3210100200202213-3031331313010323-0132303233310220-2032230221331212-1102102133301311-2123021010322233-1231122120102031-3013320331102233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.server_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.server_ssl_profile

<a id="canonical-1302131030322013-3023030122211333-1303321201211310-3023221200331023-3301001001112232-3322321222032130-3131022020013032-2011300332200312"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for server SSL profile.

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
server_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033301201123011-1212103310231101-3031311103212023-3310013010120313-2331320331110300-3231020220212333-3203311201022111-1103303020011120"></a>

### Direct properties for `virtual_server.https.server_ssl_profile`

<a id="canonical-0121330131230122-2132203220311032-2203212221212202-1231331211233321-1232022031133022-2010321122311320-3301002322212330-2113103130122030"></a>

#### `virtual_server.https.server_ssl_profile.kind` property

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

<a id="canonical-3011030022032230-1232001313203022-1020123021201331-2022011220011003-2003220100103320-1222302222100331-3302101302312002-0020300312000130"></a>

<a id="canonical-1223320122201211-3221132210331332-0132233011222312-1021213100303133-1002320120112230-0202031021332302-3000222320021021-2013130331003100"></a>

#### `virtual_server.https.server_ssl_profile.name` property

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

<a id="canonical-0011221132001012-1112311330330132-0322031302020303-2002211120002132-0012323003303223-2323003000001203-1211210300223012-0030333133311120"></a>

<a id="canonical-3112102122300033-0103222311213002-0103211132311201-2310102033030312-2330121112223030-1211230020233110-2233033300220311-0011113002303210"></a>

#### `virtual_server.https.server_ssl_profile.namespace` property

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

<a id="canonical-0310330122030031-0101222311320002-1030110132331323-2303313130222120-3111101011303331-0121313301233201-1320011100133302-1013032111302132"></a>

<a id="canonical-2000311303222113-1032222300022311-2303002300032322-1033201212230013-2230130032200201-0301112103113310-2331330321321300-1111023332311231"></a>

#### `virtual_server.https.server_ssl_profile.tenant` property

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

<a id="canonical-0302101201032331-3220323320312022-0010120103003331-0312000331223203-3101231030132310-3113322310313322-3310222102210333-1220222303022232"></a>

<a id="canonical-3313303021203321-1211310030220003-0032113333301022-3222313001131103-1001033110102111-1031221022133301-2122003230011311-0032332222011133"></a>

#### `virtual_server.https.server_ssl_profile.uid` property

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

<a id="canonical-3212022302301301-3103022201102322-3321120203323221-2023331133030122-0110010301200323-3321303211010033-0202220000313202-2322031102333330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.stream_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.stream_profile

<a id="canonical-1323133221230101-1211220021102331-0311111233130122-3331120312301222-3323200113222330-1002010112200121-2320201121033111-1031323222012022"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for stream profile.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
stream_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322312322333301-0110222023211211-2310101210202203-2221130223201121-2200101311101012-0303222122033022-3020121303023123-1131122113213020"></a>

### Direct properties for `virtual_server.https.stream_profile`

<a id="canonical-1100212030322331-1233021212131001-3203103033222110-0312311010323000-3003033133101213-0032013230210332-2011103220222221-3300000123311330"></a>

#### `virtual_server.https.stream_profile.kind` property

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

<a id="canonical-3211331103210323-2221332333220231-3201101220133002-3003330330201213-3332332132020110-0003231230031121-2011321101120303-3032200311112121"></a>

<a id="canonical-0113023032203122-1020100130323020-3210312030002213-1010232202301011-0023312300302310-1011212333122211-3112222220223010-2301201213002103"></a>

#### `virtual_server.https.stream_profile.name` property

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

<a id="canonical-3101131301333110-2002313032232303-3321100022111132-1222010332102301-0102222223130221-0330330012032010-0210332100223200-1030130331021130"></a>

<a id="canonical-0122323021123100-3103331321012020-2020133023331312-2012221111223100-1101021100200022-1013203210033113-2132230313211012-2121101130221213"></a>

#### `virtual_server.https.stream_profile.namespace` property

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

<a id="canonical-2303132231321321-2323232030330122-0033333030001002-3231312310010231-3120002201121233-0313033030003223-3303322230230330-0111332030130232"></a>

<a id="canonical-2011223132030222-1331131020322212-0012033033123303-1313003221302323-3013210213221002-1313322332113022-2121032301211330-0301321032312333"></a>

#### `virtual_server.https.stream_profile.tenant` property

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

<a id="canonical-2300211022301120-1311323011231211-2332020310032031-2023203333203001-2101302313213330-1222121123103302-2310010011232003-2301123322331020"></a>

<a id="canonical-3322203102210202-2221003133023220-3332203013122211-0013333123220312-3133023321203320-3100110213130120-0211131010113111-0222023211223303"></a>

#### `virtual_server.https.stream_profile.uid` property

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

<a id="canonical-3331122100212103-2103303202311112-2323020330102212-3020211133121232-2333002231201330-3111113212023302-1330211031211310-2112013332011103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.tcp_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.tcp_client_profile

<a id="canonical-3303020301320113-0013112032021200-1322222023000302-2120131003333003-0321131323223231-3111231031023213-3310322101031110-3311133000000013"></a>

Type: `"object"`. list nested block, Optional.

Protocol Profile (Client). Client-side configuration

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tcp_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133302302201112-1211223213203001-2302110130221203-0331023311203333-2302030120320000-1020311113021111-1321303121233012-2210221210102001"></a>

### Direct properties for `virtual_server.https.tcp_client_profile`

<a id="canonical-3032300212330101-2020003032001222-2321233223030001-0031220233120212-1033211301113110-0130301220303322-0002023030331023-3232012223031132"></a>

#### `virtual_server.https.tcp_client_profile.kind` property

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

<a id="canonical-0201133010110031-2110033022201102-3032130132113332-0321133230123002-1211121022122113-2230033223023021-2032311302311230-3010331321330103"></a>

<a id="canonical-2212331021202130-3222202121001102-0131302111231230-2331123302032221-2132121030302131-0112113130201100-3110120302101332-2313333102003230"></a>

#### `virtual_server.https.tcp_client_profile.name` property

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

<a id="canonical-2003221130112030-2123320021131023-1001330331110302-2311013101330310-2033020220213132-1102200102312230-0101121120323213-2300121302010132"></a>

<a id="canonical-1202023110113202-3030001302313333-1032110001200102-2132232332013200-0320022230130033-1101012212301233-2001131020110210-2211332202323312"></a>

#### `virtual_server.https.tcp_client_profile.namespace` property

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

<a id="canonical-3022200102020212-2102312321101213-3332122213120300-3030201033213310-1303231031001323-3123211112200022-0103303321333321-3331211000121300"></a>

<a id="canonical-0010113310002032-2022002230220133-0302033110033210-1032030013312113-1123021122223000-1100132113201211-1323311202010132-2223121201132322"></a>

#### `virtual_server.https.tcp_client_profile.tenant` property

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

<a id="canonical-1132330013210000-2001200002021010-2203022121320113-0321020220212100-0230320312221200-3020202033303133-3220122133111023-0031301233003033"></a>

<a id="canonical-2023310102210332-2121002131222102-0132032220011132-3321300232102323-3021211222310313-3332322222230300-0022003010233331-3103020012223123"></a>

#### `virtual_server.https.tcp_client_profile.uid` property

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

<a id="canonical-0022210002333001-3112222323110013-3020033033022233-1111120310313331-1110213213313203-1331130321003201-3100013203212102-1123123203203132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.tcp_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.tcp_server_profile

<a id="canonical-2210332031001101-3032113311222310-0111231003010232-0333222211331022-3130220330300311-0323123333023213-3122230031130333-0322301330020202"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for tcp server profile.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tcp_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303010100100322-3320310030321112-2102203222201121-0320103222330323-2121120322320010-2012211130211321-0301110213203020-0331232122112130"></a>

### Direct properties for `virtual_server.https.tcp_server_profile`

<a id="canonical-2123101320202333-1230313111120003-3321213300233003-0333021020221313-0122301022221021-2011302311110330-2310110030312331-1202330211330231"></a>

#### `virtual_server.https.tcp_server_profile.kind` property

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

<a id="canonical-1211213332131223-3321032031122311-1302230121123021-1110132313232011-3021320321330033-0230000301011221-3023032210133220-0033132122300102"></a>

<a id="canonical-3112301101310031-0301102233313021-2220331310220031-3320122233303223-2002313303203323-2120010100330212-1330112121230312-0300333013331323"></a>

#### `virtual_server.https.tcp_server_profile.name` property

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

<a id="canonical-0223010133022003-0003310213232130-2003103333303231-0330000211011332-1331110033200231-2113210322321010-2011001012221320-2200330310211322"></a>
