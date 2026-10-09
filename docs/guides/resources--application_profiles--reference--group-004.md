---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

<a id="canonical-2112020122102002-1231023012202001-0232103033111123-1232010112110233-1301332112302121-2002333031303322-0301033220102322-1212220212310303"></a>

## `virtual_server.https.tcp_server_profile.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3332301021033311-2001013131001110-1310310331323133-0333313333111103-1020230100232310-3213023121121103-1233001321330130-2122023113332033"></a>

<a id="canonical-2311131013103311-3032303032230222-2321300002232312-1232223322011211-1223033002030010-0113322001023231-1131220333001213-2111010122213222"></a>

## `virtual_server.https.tcp_server_profile.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2303310310321332-2011003112300000-2223123033322230-0021233113230202-1111111213121203-1000020113011011-1300203323021233-0101033021023022"></a>

<a id="canonical-3003013123030200-0012113031033202-1033313001200033-2210301230030031-2323220202022313-0103222130201133-2032200203211013-0231121102133032"></a>

## `virtual_server.https.tcp_server_profile.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1222221330201310-3101213130310000-2333211331100210-2102023100223033-0030012112211323-2023233213123221-1220103312330003-1013000000133131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.websocket_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.websocket_client_profile

<a id="canonical-0321132123300232-1313031133033202-3120303020320213-2103110113311202-3101300112300221-2302231110003323-3111332212002330-0213011031011312"></a>

Type: `"object"`. list nested block, Optional.

WebSocket Profile Client. Web-related configuration

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
websocket_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2100100101112123-3001211322013031-2012322323122002-0102303301332223-1332022312313022-3312032233312203-1300230233222031-1103131002201102"></a>

### Direct properties for `virtual_server.https.websocket_client_profile`

<a id="canonical-1230101103122213-2220322121321111-3110001101012202-0213003332023021-0001321201121302-0231013122310220-3332231112010002-2203130122311302"></a>

#### `virtual_server.https.websocket_client_profile.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3033102210031010-2220102133122020-0131311020202300-1032310000002131-2000132121113122-1331110023120303-2132220233022210-0312331110020210"></a>

<a id="canonical-0020103032232202-1103222311122021-0103300113123322-3032233200113012-0101023022203223-0130210300032331-2220313230131132-3112032321313303"></a>

#### `virtual_server.https.websocket_client_profile.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1223302120121302-2000212220312303-1013001023021311-1330223230130130-1300112122300003-1020220232230130-3110333110113001-0210310321100201"></a>

<a id="canonical-1322122220300213-3101311103231103-2213200232222221-3103123133111103-3013200101232202-0103023130003003-3211100301001320-0201002221322211"></a>

#### `virtual_server.https.websocket_client_profile.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0031233030311313-3233012110320121-2133020221203023-0321220300100330-3000010113121113-1231133223320003-3131320303322210-3133300213302100"></a>

<a id="canonical-2312300102012222-3222231030233202-0202300112102331-2023100211333122-0020030102200113-1022230010201202-3120130030301222-1100113102132332"></a>

#### `virtual_server.https.websocket_client_profile.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0202100213232022-2112002130110032-1233320311121300-3300312012203332-2322110313201220-0232102103220211-2120313022103033-0302001030300111"></a>

<a id="canonical-2203312001201102-1112132103013231-1010032120030122-0332111130322321-2303001101103332-1220301003031120-0301113120012230-1320022012002332"></a>

#### `virtual_server.https.websocket_client_profile.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0211103332301002-3311220121001320-2202321013331102-0113223312313112-0102220010300023-0213310332210133-3211320101132301-0213110231333311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.https.websocket_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.https](resources--application_profiles--reference--group-003.md#canonical-0112001023212212-3331102203112201-3201210133100333-3300332211003123-1330222323020123-1020313322221200-1000132310132223-1130222102211031)
- virtual_server.https.websocket_server_profile

<a id="canonical-1211103303022013-3011120031202312-3312010310110313-0331213122130111-0103113120321002-0100023210232022-3213113010130310-3303210301033233"></a>

Type: `"object"`. list nested block, Optional.

WebSocket Profile Server. Web-related configuration

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
websocket_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131032303323202-1331101231123212-0101222032233311-3132003121132213-0311321220321320-0103320113213213-3232201131221130-3303322000000203"></a>

### Direct properties for `virtual_server.https.websocket_server_profile`

<a id="canonical-2002132132300230-0201310132212033-1322112023211312-3023130003100103-0032230023013302-3202130322133201-0032303112130113-3220211000320232"></a>

#### `virtual_server.https.websocket_server_profile.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3103101222123011-0220212113302331-0203032132132012-3213123211011220-2011233033232100-2131022112212000-1123201031011012-2303232003210310"></a>

<a id="canonical-1102130133302133-2103222022302233-3023321202203301-3333100332220020-2201013300333202-3021010331313330-2321321321130222-0112023000013332"></a>

#### `virtual_server.https.websocket_server_profile.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0213002233302131-0000211332221202-3301223000223001-3332223012330000-3312002103321233-3033122021300301-0201032310203311-0332030202202133"></a>

<a id="canonical-3022312302000321-2311200131123121-2101111132222313-2230331330322202-0100100112130000-2030312230310031-0002133023202003-3033200103213322"></a>

#### `virtual_server.https.websocket_server_profile.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1023131000113303-0330032322331212-3210211211332231-2003202330121003-3301000110113230-2220122100102220-2330201220000131-3011121012130012"></a>

<a id="canonical-0320301223233300-0233300213323023-1100132030033021-0320111231210022-2210130323122313-1132300312330221-3121212231333312-2222332310002232"></a>

#### `virtual_server.https.websocket_server_profile.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3131223303321333-0012221311011021-3031000011303211-0111313020323322-1030020110220032-3133220121310033-2231220333020022-2312010310032002"></a>

<a id="canonical-1233002210103212-3301212310121222-3313131000233002-1112100130221330-1330023033313231-1213312111011023-2212332013203213-2101333302220222"></a>

#### `virtual_server.https.websocket_server_profile.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2120313011001333-0332132133112211-2032113202213200-1133333320210120-0120003121121103-1110233000210030-2313321130012321-0301012230001201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.immediate_action_on_service_down` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.immediate_action_on_service_down

<a id="canonical-3230213032221222-0013333133210033-1222132013313112-2322100010110313-1330313131112110-0230121231012010-0232033001322133-0310202132120231"></a>

Type: `"object"`. single nested block, Optional.

Specifies the immediate action the BIG-IP system should respond with upon the receipt of the initial
client's SYN packet, if the availability status of the virtual server is Offline or Unavailable.
This is supported for the virtual server of Standard type and TCP protocol. The default is None.
None: Specifies that the system takes no immediate action if the virtual server is reported Offline
or Unavailable. Reset: Specifies that the system resets the connections when the virtual server is
reported Offline or Unavailable. Drop: Specifies that the system drops the connections when the
virtual server is reported Offline or Unavailable.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("immediate_action_on_service_down_drop",
    "immediate_action_on_service_down_none"),
  validators.ConflictingObjectAttributes("immediate_action_on_service_down_drop",
    "immediate_action_on_service_down_reset"),
  validators.ConflictingObjectAttributes("immediate_action_on_service_down_none",
    "immediate_action_on_service_down_reset")}
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
  "x-ves-oneof-field-immediate_action_on_service_down_choice": "[\"immediate_action_on_service_down_drop\",\"immediate_action_on_service_down_none\",\"immediate_action_on_service_down_reset\"]"
}
```

Terraform syntax:

```terraform
immediate_action_on_service_down {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232133220020001-3001001333212220-0331211120033222-0210001311100202-2110213220302011-0321323020021202-2003020020233233-3033202322103301"></a>

### Direct properties for `virtual_server.immediate_action_on_service_down`

- [immediate_action_on_service_down_drop](resources--application_profiles--reference--group-004.md#canonical-3331020103122320-2130223203321331-2022122221212103-2201303102023313-2011113333303021-1010110302202213-3330201123212123-2321310322212122): complete subsection reference.

- [immediate_action_on_service_down_none](resources--application_profiles--reference--group-004.md#canonical-1211320131313223-0300030133102303-1120303202100033-0302033120121202-0211200222132132-2221120100211201-0332212332001102-2320302000100222): complete subsection reference.

- [immediate_action_on_service_down_reset](resources--application_profiles--reference--group-004.md#canonical-1020323002022302-2031200012130221-0130123012221301-0303002023310122-1011303313130223-0213220032010232-0220130012302202-3020330121013201): complete subsection reference.

<a id="canonical-3331020103122320-2130223203321331-2022122221212103-2201303102023313-2011113333303021-1010110302202213-3330201123212123-2321310322212122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-004.md#canonical-2120313011001333-0332132133112211-2032113202213200-1133333320210120-0120003121121103-1110233000210030-2313321130012321-0301012230001201)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop

<a id="canonical-1221013321201031-0111102300222301-3311010131213030-0001200212301332-0200022011103130-1230030322301001-1210102033020203-0310010232233012"></a>

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
immediate_action_on_service_down_drop = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211320131313223-0300030133102303-1120303202100033-0302033120121202-0211200222132132-2221120100211201-0332212332001102-2320302000100222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-004.md#canonical-2120313011001333-0332132133112211-2032113202213200-1133333320210120-0120003121121103-1110233000210030-2313321130012321-0301012230001201)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none

<a id="canonical-2003131000131233-3100223321020233-1201100121133302-3200203320011111-0112132020230001-2020331313201112-0001223210133112-2031122322303203"></a>

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
immediate_action_on_service_down_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020323002022302-2031200012130221-0130123012221301-0303002023310122-1011303313130223-0213220032010232-0220130012302202-3020330121013201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.immediate_action_on_service_down](resources--application_profiles--reference--group-004.md#canonical-2120313011001333-0332132133112211-2032113202213200-1133333320210120-0120003121121103-1110233000210030-2313321130012321-0301012230001201)
- virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset

<a id="canonical-3302211020211201-1110322130010101-0223133020003013-1123322322203021-3312323132031330-3200020132302312-0000211003333112-3033230302112203"></a>

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
immediate_action_on_service_down_reset = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133301311230011-1310300321322210-2110201322213000-0102311121311103-1133120133100303-3022030212232231-3000032201020203-0023300101231100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.last_hop_pool` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.last_hop_pool

<a id="canonical-3130100312210212-0020201222232323-0102223330023013-2103121210010212-2120033023232203-2312233331101003-1200120122133211-3232032330003110"></a>

Type: `"object"`. list nested block, Optional.

Directs reply traffic to the last hop router using the specified pool.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
last_hop_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011023201113313-2123301031202131-1223321221210133-3003210100323300-1111221110011101-2133331303021333-1311311022300313-2031133113320201"></a>

### Direct properties for `virtual_server.last_hop_pool`

<a id="canonical-0133220332112203-2120203331023303-0103103201001100-0020310302031313-1322121321222110-2223230321020200-2323310102331130-2110322010101133"></a>

#### `virtual_server.last_hop_pool.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0131123000130202-2001333301303310-3310133113312212-2211331301120031-3021020201333212-1023003310010321-3102323310010113-1303302321031230"></a>

<a id="canonical-3000033032201312-0003011320211132-3311032003322032-1023122121220132-0210233131233232-2122121003112310-0201121102010112-1002100313122001"></a>

#### `virtual_server.last_hop_pool.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3110220010333333-0132110232200320-2333301323012210-1012002310320022-1210302020303131-0312011232331102-3203201301113230-2213010313021030"></a>

<a id="canonical-1010030312111122-3320032033210110-1012132020032202-3320110132223122-1020220102130123-0303000132201201-3223110221312232-3332202111330113"></a>

#### `virtual_server.last_hop_pool.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1020121201321111-0321222303220311-3031220120003133-1023233230022012-2021133322232211-3002130130133210-0313013101231031-1310120223223031"></a>

<a id="canonical-0030313132323001-3220311020012031-2012322210020211-2002221021200010-3312011032120123-2330031012122113-1123230121212233-3223021202013011"></a>

#### `virtual_server.last_hop_pool.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2003030013011001-2001200312132330-3012223122303112-2103222123001133-3212331021031331-1222010300102330-0130102200233130-0100223322332213"></a>

<a id="canonical-0021110133332301-3221311201132113-2301300311320033-2123022222233220-2321031331123331-0302022030100131-0121222323230220-3313201021023131"></a>

#### `virtual_server.last_hop_pool.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2002031331123221-1331221331321023-3111232122011101-0320122210120313-0212322110001221-0230222031021312-0203110233131210-3013333012331002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.nat64` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.nat64

<a id="canonical-1003101233121130-3210201232023112-3310130202030221-2000220312333131-1103301120112202-2232110110103211-3021002033300032-3230222222320030"></a>

Type: `"object"`. single nested block, Optional.

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if the
system does not have a default route configured and the client is located on a remote network. This
setting is also useful when the system is load balancing transparent devices that do not modify the
source IP address of the packet. Without the last hop option enabled, the system could return
connections to a different transparent node, resulting in asymmetric routing. You can configure this
setting globally and on an object level. You set the global Auto Last Hop value on the System ::
Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the
Default setting. When you configure Auto Last Hop with a value other than Default at the object
level, its setting takes precedence over the global setting. This enables you to configure auto last
hop on a per-virtual server basis. The default is Default, meaning that the system uses the global
auto-lasthop setting to send back the request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("nat64_disable",
    "nat64_enable")}
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
  "x-ves-oneof-field-nat64_choice": "[\"nat64_disable\",\"nat64_enable\"]"
}
```

Terraform syntax:

```terraform
nat64 {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220333102101031-3000200212111212-2211100221012022-1210302133133021-3211123032332112-2103110232302302-0302010303311000-3202033313320233"></a>

### Direct properties for `virtual_server.nat64`

- [nat64_disable](resources--application_profiles--reference--group-004.md#canonical-1101131332320313-1030303021303121-3201103010330112-0200130321003203-3233323021222323-0021303201111130-0000313032310333-0112133213301313): complete subsection reference.

- [nat64_enable](resources--application_profiles--reference--group-004.md#canonical-0132100231000220-3233110220132320-0202230232001221-0301000023030303-3231320123022201-2323010213123020-0132002101033031-3120303123210130): complete subsection reference.

<a id="canonical-1101131332320313-1030303021303121-3201103010330112-0200130321003203-3233323021222323-0021303201111130-0000313032310333-0112133213301313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.nat64.nat64_disable` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.nat64](resources--application_profiles--reference--group-004.md#canonical-2002031331123221-1331221331321023-3111232122011101-0320122210120313-0212322110001221-0230222031021312-0203110233131210-3013333012331002)
- virtual_server.nat64.nat64_disable

<a id="canonical-0032200222113123-3223310301322110-3102323310301321-0323320301120132-3213301222232333-1303323203000301-1230002321301113-0200223100132103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for nat64 disable.

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
nat64_disable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132100231000220-3233110220132320-0202230232001221-0301000023030303-3231320123022201-2323010213123020-0132002101033031-3120303123210130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.nat64.nat64_enable` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.nat64](resources--application_profiles--reference--group-004.md#canonical-2002031331123221-1331221331321023-3111232122011101-0320122210120313-0212322110001221-0230222031021312-0203110233131210-3013333012331002)
- virtual_server.nat64.nat64_enable

<a id="canonical-3213312001302303-0123300102331001-2011233202113212-3023132011020311-0222123030123111-1023322311011003-1003301320323332-1231200122323330"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for nat64 enable.

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
nat64_enable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121121231001031-3330300000133002-1031230222111031-0110330122312033-2031202031313323-3102103200300201-2230013301021233-3121002202231030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.port_translation` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.port_translation

<a id="canonical-3300211131303022-0213330231302003-0323211023223112-0221132220220312-1121333121213210-2133331232122332-0010300301232312-0122121000212121"></a>

Type: `"object"`. single nested block, Optional.

Specifies, when checked (enabled), that the system translates the port of the virtual server. When
cleared (disabled), specifies that the system uses the port without translation. Turning off port
translation for a virtual server is useful if you want to use the virtual server to load balance
connections to any service. The default is enabled.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("port_translation_disable",
    "port_translation_enable")}
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
  "x-ves-oneof-field-port_translation_choice": "[\"port_translation_disable\",\"port_translation_enable\"]"
}
```

Terraform syntax:

```terraform
port_translation {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110110031022023-2213322211011013-2112230023222011-0230111112100113-0111322312202330-2121313223321011-3031301211232210-3202210202113330"></a>

### Direct properties for `virtual_server.port_translation`

- [port_translation_disable](resources--application_profiles--reference--group-004.md#canonical-2322122032001202-0320333323113103-1320111022100231-1122220021233203-0132331132331011-0222313003233312-0311310233221122-0300233002300010): complete subsection reference.

- [port_translation_enable](resources--application_profiles--reference--group-004.md#canonical-2333311333021133-0102003021323220-2320330011130202-0302232011300011-0133021103300303-1030131302202301-2121022231132022-0333012033121023): complete subsection reference.

<a id="canonical-2322122032001202-0320333323113103-1320111022100231-1122220021233203-0132331132331011-0222313003233312-0311310233221122-0300233002300010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.port_translation.port_translation_disable` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.port_translation](resources--application_profiles--reference--group-004.md#canonical-3121121231001031-3330300000133002-1031230222111031-0110330122312033-2031202031313323-3102103200300201-2230013301021233-3121002202231030)
- virtual_server.port_translation.port_translation_disable

<a id="canonical-0331201330112220-3322000201302223-3302323211013310-1011022233300321-0102203301310010-0301110102203200-0222131221201313-0130021112300230"></a>

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
port_translation_disable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333311333021133-0102003021323220-2320330011130202-0302232011300011-0133021103300303-1030131302202301-2121022231132022-0333012033121023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.port_translation.port_translation_enable` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.port_translation](resources--application_profiles--reference--group-004.md#canonical-3121121231001031-3330300000133002-1031230222111031-0110330122312033-2031202031313323-3102103200300201-2230013301021233-3121002202231030)
- virtual_server.port_translation.port_translation_enable

<a id="canonical-2110210322123011-1100230132320211-0101312110330201-0010313202120310-2302300301132231-3303200010330110-1020120100133223-0330011233231021"></a>

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
port_translation_enable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210202300313301-0310203232211033-1321103003021001-0212002312311332-1112331303033331-3323000030313321-0300000231230113-3220231130103020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.request_logging_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.request_logging_profile

<a id="canonical-3021100321112301-1233301012311200-0312330023111200-2001301020303311-2233302130132031-2113023101013111-0133021233213323-3232312330213023"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for request logging profile.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
request_logging_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220123331312013-1233210103033002-0200210010131212-1030102122312233-3313000030023322-1303203310021301-3033033010322301-0232111203321221"></a>

### Direct properties for `virtual_server.request_logging_profile`

<a id="canonical-2233232030023313-3133202323110131-3131322232122022-1020123321200333-2321211132333330-0330210122202000-2232233121020312-1121133312313321"></a>

#### `virtual_server.request_logging_profile.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2330113122030033-3333310220311332-0223021101213303-1323103321221121-3201211312121333-0033032101131310-2332022230133023-3322223132003220"></a>

<a id="canonical-1322333210020223-2120230110312331-1112110332100120-1331221210011010-1233012033321120-1033320220212112-2203021200000120-0123103231200231"></a>

#### `virtual_server.request_logging_profile.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0233120103100302-1311201110110101-1120331232001210-0000121221300121-2120013031003101-0213331233212322-0010322230122333-3123123332012210"></a>

<a id="canonical-3122110221220000-1200001012321230-3012131020313301-0112300320223203-2132303222112321-3312130313222112-0212022200231033-3112010312003011"></a>

#### `virtual_server.request_logging_profile.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0203102211223131-3103102013132000-3130123210303311-2100310013322023-2003321300202302-0101001321232302-0232203321100000-2310013110330001"></a>

<a id="canonical-3310030321013232-0321313131032033-3321220231011032-3023100321202022-3233333330213031-3312213020131021-0220003131103101-1220030031321012"></a>

#### `virtual_server.request_logging_profile.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1031020323010220-1310111021303131-1211130301120110-0120303302330312-0301202230332110-0122211111020020-2323102113133330-0220213033212231"></a>

<a id="canonical-1230033223300210-3233310213203303-0003003213032312-2232300333300232-1230122101210311-3001201021030333-1213132211320002-2022231222231222"></a>

#### `virtual_server.request_logging_profile.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1232230330103113-1011323112020110-2112322333333013-3302112210103221-2002022032123220-0023200230233130-2102030102320103-0013123112003122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.source_port` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.source_port

<a id="canonical-2212010332223131-2111323213003221-3122213303023310-2020333131322323-0103201100101002-1020010232021313-0111203202021022-2223132001320213"></a>

Type: `"object"`. single nested block, Optional.

Specifies whether the system preserves the source port of the connection. The default is Preserve.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("source_port_change",
    "source_port_preserve"),
  validators.ConflictingObjectAttributes("source_port_change",
    "source_port_preserve_strict"),
  validators.ConflictingObjectAttributes("source_port_preserve",
    "source_port_preserve_strict")}
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
  "x-ves-oneof-field-source_port_choice": "[\"source_port_change\",\"source_port_preserve\",\"source_port_preserve_strict\"]"
}
```

Terraform syntax:

```terraform
source_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233131132332312-0200321001011311-3000310222012331-2232102121012220-0223031020202201-3231000001232202-0230103332330300-2130212300001000"></a>

### Direct properties for `virtual_server.source_port`

- [source_port_change](resources--application_profiles--reference--group-004.md#canonical-0001232132032021-0203012123211133-0320111311333103-3333303322101113-3130310123301120-3310300232031132-0113103322101012-3021301033310022): complete subsection reference.

- [source_port_preserve](resources--application_profiles--reference--group-004.md#canonical-0103110100200031-2120012230132120-3221310123003222-0102011031233010-3323321000111202-3210032023300110-1210012201200122-2012133333321303): complete subsection reference.

- [source_port_preserve_strict](resources--application_profiles--reference--group-004.md#canonical-2110002333222312-2100332330303222-0000003122110232-0110132011311230-1123023003303233-2120121032203202-3211123212022100-3030233302310031): complete subsection reference.

<a id="canonical-0001232132032021-0203012123211133-0320111311333103-3333303322101113-3130310123301120-3310300232031132-0113103322101012-3021301033310022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.source_port.source_port_change` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.source_port](resources--application_profiles--reference--group-004.md#canonical-1232230330103113-1011323112020110-2112322333333013-3302112210103221-2002022032123220-0023200230233130-2102030102320103-0013123112003122)
- virtual_server.source_port.source_port_change

<a id="canonical-1311330323021032-2210013333333302-1331301003130203-0003333220201000-3131331012012000-2013032302131030-2303202113122003-3323113201112320"></a>

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
source_port_change = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103110100200031-2120012230132120-3221310123003222-0102011031233010-3323321000111202-3210032023300110-1210012201200122-2012133333321303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.source_port.source_port_preserve` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.source_port](resources--application_profiles--reference--group-004.md#canonical-1232230330103113-1011323112020110-2112322333333013-3302112210103221-2002022032123220-0023200230233130-2102030102320103-0013123112003122)
- virtual_server.source_port.source_port_preserve

<a id="canonical-1022313012110112-0320200311020023-1202203230011030-1020100221232300-1322001023000120-3123313223012212-0303223030221123-1102121132013322"></a>

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
source_port_preserve = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110002333222312-2100332330303222-0000003122110232-0110132011311230-1123023003303233-2120121032203202-3211123212022100-3030233302310031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.source_port.source_port_preserve_strict` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.source_port](resources--application_profiles--reference--group-004.md#canonical-1232230330103113-1011323112020110-2112322333333013-3302112210103221-2002022032123220-0023200230233130-2102030102320103-0013123112003122)
- virtual_server.source_port.source_port_preserve_strict

<a id="canonical-2212111003100111-1020312302333103-1300101311131310-0311112203132102-2032131012231213-1123102122310031-2310031230010301-0130002031012300"></a>

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
source_port_preserve_strict = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230332010120210-3032201031320120-1320321332001001-1230210320213223-3313130100303010-3030221032130120-0211333311033312-3013103322102113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.statistics_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.statistics_profile

<a id="canonical-2033031113031132-0023012103300133-1113311132332021-2001201030102232-3111110013031123-1322303301011332-3021010132203122-3121230300333000"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for statistics profile.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
statistics_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230121030222002-0033031003112020-1222012022220223-2023010123202333-1120310203220321-0213303332301122-0311000112200202-3132212212222123"></a>

### Direct properties for `virtual_server.statistics_profile`

<a id="canonical-1003022020103133-2131111133003331-0330033120200303-2201030200030322-2303321002002211-0202203103201322-2310331120023300-2022202210203031"></a>

#### `virtual_server.statistics_profile.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1300331213131100-1210231223321210-0020103121333312-3101130222010232-3132202211311101-3010011311133122-0031101310031213-1302031013313323"></a>

<a id="canonical-3123312201013233-3012033020121103-2031020102000321-0311023031122023-3111000332220011-0123121020213320-3213102330021031-2310332202131220"></a>

#### `virtual_server.statistics_profile.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3223323302013221-2212221320311113-0100312022002101-0330320332101201-0100233201322021-1030312112310232-0011022110313021-3300023211200223"></a>

<a id="canonical-0102030000302012-3101333231110332-2221303100201220-0312032002223123-2333133220131123-1232122132311030-2210130112332233-2111130010301213"></a>

#### `virtual_server.statistics_profile.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2221133111123231-3110101301023112-3231003013000012-2122230101110303-3212013230220030-2232232033012220-1223233112020031-0002030333311121"></a>

<a id="canonical-2323221120022300-3331121302022011-2000112310323332-2322210131333121-2221022211300321-1101103012331233-0123223300013120-1102302233001111"></a>

#### `virtual_server.statistics_profile.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0132320122003211-0231231122332331-1003010313310213-2123301122301231-0322332221310022-0203123211123220-1222031302133221-1200331122023110"></a>

<a id="canonical-1231020131210033-3321233133300323-2000223022322031-2330000021130100-1130232002033231-2223223313330112-1001013032103320-2023032010111031"></a>

#### `virtual_server.statistics_profile.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.tcp` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.tcp

<a id="canonical-1310210103201212-1320301132201233-1210200013231223-1101230001132123-3302033232011132-2332321312221012-2323132221331212-0130323232233301"></a>

Type: `"object"`. single nested block, Optional.

TCP profiles.

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
tcp {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110020231202222-0303102112022011-0100132321010011-0111033303203232-2122322111020020-2321100121131110-1223303023221233-3230132122201210"></a>

### Direct properties for `virtual_server.tcp`

- [client_ssl_profile](resources--application_profiles--reference--group-004.md#canonical-2132222223303130-1022130203031333-3310131122103110-0212220202301313-3312320232032323-3331223010113331-1322233131031331-3313323020033222): complete subsection reference.

- [ocsp_profile](resources--application_profiles--reference--group-004.md#canonical-1001023112303131-2331323203333210-1232132101202000-2233311002102231-3110312133322122-3322100033331020-3302210231201232-2313103221001212): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--reference--group-004.md#canonical-0121320311113012-3102203220321101-2131030131221121-1312131030132203-2021130233323110-1023233231211102-3231031033113313-3221001212010323): complete subsection reference.

- [tcp_client_profile](resources--application_profiles--reference--group-004.md#canonical-1112301121000320-0032302331230222-3123213322221103-2320310231001320-2232202220023120-0332311010112102-3303010032311201-2001323230021001): complete subsection reference.

- [tcp_server_profile](resources--application_profiles--reference--group-004.md#canonical-0032321113220230-2210221333303233-0123322330001031-0211010312302301-0111111221033222-2310210331211213-3000322022123112-2231210311122112): complete subsection reference.

<a id="canonical-2132222223303130-1022130203031333-3310131122103110-0212220202301313-3312320232032323-3331223010113331-1322233131031331-3313323020033222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.tcp.client_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.tcp](resources--application_profiles--reference--group-004.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231)
- virtual_server.tcp.client_ssl_profile

<a id="canonical-1012331213102311-2311220320111322-0123323233221201-0303203130002202-2000012130110011-0131133020313202-2111122321103312-1332310121103313"></a>

Type: `"object"`. list nested block, Optional.

Client SSL Profile. Client-side configuration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
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

<a id="canonical-3313120332123020-3002311011333001-3300220310220202-0000232231022213-3232232112133322-3323322212023110-2021213021131012-1100021210230213"></a>

### Direct properties for `virtual_server.tcp.client_ssl_profile`

<a id="canonical-1233012303331002-1301300223130020-1210102110333010-1201103032033230-1030230201231331-2330230321233222-2123313200310302-3010200130331201"></a>

#### `virtual_server.tcp.client_ssl_profile.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3312021332221321-2210030022101301-1223102233113103-1310220121202012-0202103103302322-0320203233310212-2322212223322033-1233030323030321"></a>

<a id="canonical-1011322202233302-0303313123002231-1020331302133133-3300333220302320-0110000233011130-1121211202333313-3113032021131103-3200011221001302"></a>

#### `virtual_server.tcp.client_ssl_profile.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2231021120022302-3031120322012220-3123202230023202-1201233022301012-3112121200011123-3331223131023202-1230310323230232-3220010310211111"></a>

<a id="canonical-2000030311312200-1232120320110322-2030132121031121-2113301313120031-2113111302333210-1221131223103211-0200303123321002-3121220121132210"></a>

#### `virtual_server.tcp.client_ssl_profile.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3120212012111032-1013103233222010-1203000331212110-1133102031232311-3121223232313333-0233221303232133-0323231331110220-1201223021000213"></a>

<a id="canonical-1100003222211122-2223020232101113-2311302013013231-1333320210200301-1030231203213330-3211030002032012-2113201312033310-2111033212303023"></a>

#### `virtual_server.tcp.client_ssl_profile.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1121233010211331-3031222020231121-2312000122013110-1220221302100230-3200321013220312-3331200330132012-3001310002311313-3322122323132033"></a>

<a id="canonical-0001111321330321-2102331321121231-0000313211302203-2003131022032031-1302130311202121-3331131213002011-1002003303101021-0022010010122321"></a>

#### `virtual_server.tcp.client_ssl_profile.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1001023112303131-2331323203333210-1232132101202000-2233311002102231-3110312133322122-3322100033331020-3302210231201232-2313103221001212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.tcp.ocsp_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.tcp](resources--application_profiles--reference--group-004.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231)
- virtual_server.tcp.ocsp_profile

<a id="canonical-2201330102321311-3113102120311103-0132000300300123-2232331231112021-0030311232033330-1030333220331133-2233212221203000-2220021013322130"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3311222201333030-2130331003231121-3130130020212123-1032311133233232-2130310230331321-2103021112000300-0101003220112111-2301013311232231"></a>

### Direct properties for `virtual_server.tcp.ocsp_profile`

<a id="canonical-1201102121120222-1022222303212032-0322203300233201-2321122032331211-1312031123202300-3123131021001101-0031312312010221-3032313300130002"></a>

#### `virtual_server.tcp.ocsp_profile.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3220000202123301-2133312020312122-2213001131121321-2300122321203311-1333130233201213-1203023101110133-0311130020302200-2133220203102110"></a>

<a id="canonical-3112123102212310-3233222111133230-2312203232320202-3223101232021122-0211021322212303-2231303012311300-1321300000303300-3130012200332212"></a>

#### `virtual_server.tcp.ocsp_profile.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1030330131020231-3012302212102302-0131222233220221-2313221220330210-1113333230120223-3223121032303210-3022220012221013-3021233212311311"></a>

<a id="canonical-0222211230322302-1232210001301213-2031333302100013-3131312123011211-0021021131313200-1231000201211322-3323031011200110-0133202003122200"></a>

#### `virtual_server.tcp.ocsp_profile.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0032230003100312-1121023211302231-3120031120223200-2212022112333313-2120130201310201-1132102223112330-0303301332002311-0001202023003323"></a>

<a id="canonical-2332132330302112-2331221312002221-3200331001130101-2231323310313332-3302321020012323-3030310120023020-1303003233110132-2013131321122221"></a>

#### `virtual_server.tcp.ocsp_profile.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2221023120123210-3013033120202100-1222203231232011-0210232310230203-3321200030133200-2100022322310200-3100301220121010-1120022013123311"></a>

<a id="canonical-3210322032012302-1322233133021020-3233300002220001-1131330120331331-0113213231223131-2203200220030320-1013311332103110-3113312323100203"></a>

#### `virtual_server.tcp.ocsp_profile.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0121320311113012-3102203220321101-2131030131221121-1312131030132203-2021130233323110-1023233231211102-3231031033113313-3221001212010323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.tcp.server_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.tcp](resources--application_profiles--reference--group-004.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231)
- virtual_server.tcp.server_ssl_profile

<a id="canonical-0313332111330002-2300020230310231-2231001022303321-0233332113220203-3231013103110021-2131313133233102-1133110013032032-3130221031223302"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for server SSL profile.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
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

<a id="canonical-3033111120331301-0321320120303021-3332213021233103-2012213310331202-0023221233130132-0221313122110131-2201113230232132-1310223101130333"></a>

### Direct properties for `virtual_server.tcp.server_ssl_profile`

<a id="canonical-3000301213030020-0232223320210223-1100203023111111-2200123330230123-0220113113122020-0033322302122200-1020203113131031-2033323320221123"></a>

#### `virtual_server.tcp.server_ssl_profile.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0331011311300103-2303020313011023-0323311131001332-2232012123332212-1213101230013110-2203232012211103-0312112021022210-2131202323331112"></a>

<a id="canonical-0011230212110333-2033123220002310-2031213032013313-3013212310023313-1000333101230002-3311011331033022-1300323220210311-3222012221111113"></a>

#### `virtual_server.tcp.server_ssl_profile.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2211313211222321-1200130133311001-1321223203123312-2322013012003030-2101320031032233-3021131130011110-3320231203110033-2110132130210033"></a>

<a id="canonical-0233133311220002-3203023001333333-0333320131031200-2323010222001302-2103022220212232-0032033130221220-2331311013130120-3223130322333020"></a>

#### `virtual_server.tcp.server_ssl_profile.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0320303333101212-1300001123203212-1120030002323031-3032033320113002-2232031333132003-1033022112033322-1300013131312023-1210002101332203"></a>

<a id="canonical-0220113323031231-3232223313213130-0310130032022102-0012012213210022-3030212111232231-3302003132121131-3212031300120203-1030130101133203"></a>

#### `virtual_server.tcp.server_ssl_profile.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2301220023033330-2023131320213003-2131230213320311-1300222113211131-1112331122121130-2211030220030333-2112302322201101-3201301223333221"></a>

<a id="canonical-3112213032132323-0331300201233230-0203200110001103-0030120020112302-2122020300322200-2203101210312021-2310310330333310-1332131313332310"></a>

#### `virtual_server.tcp.server_ssl_profile.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1112301121000320-0032302331230222-3123213322221103-2320310231001320-2232202220023120-0332311010112102-3303010032311201-2001323230021001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.tcp.tcp_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.tcp](resources--application_profiles--reference--group-004.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231)
- virtual_server.tcp.tcp_client_profile

<a id="canonical-2323011223033022-3013123232310303-2200023233330130-2110012032322320-2211000322003323-1132330113030022-1210110101322320-2011132102222300"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1023321102201221-3130121202210200-3313213033131103-2302312001113201-1221013130320232-1102332303000030-0233210213101123-0132110212011113"></a>

### Direct properties for `virtual_server.tcp.tcp_client_profile`

<a id="canonical-1202333133013223-0321300232132300-3322332332223322-3310100210232000-2313000123313022-0133031012031030-0111033001113023-1333023133112133"></a>

#### `virtual_server.tcp.tcp_client_profile.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2313321202202003-2211103320031111-1320213310113001-1231201312123112-0133033013113213-3001313330131133-2200001003013203-2233312331323120"></a>

<a id="canonical-0202233311032000-1132302313231013-2011330323132322-3321320303303230-3131023023313222-3303311013233031-0022330030032331-0001002003222322"></a>

#### `virtual_server.tcp.tcp_client_profile.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2111120120313200-3222101022030010-1301020013112011-2231120323103120-2010320021303330-3022113313332101-0230303203133131-2231323202010312"></a>

<a id="canonical-3130301233233220-1102322131132211-3020303033000032-0302130222010221-3332303021033322-0100331311123132-0032000131331222-1202031230123231"></a>

#### `virtual_server.tcp.tcp_client_profile.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1031111320132203-0200012003211020-3021103333133103-2210212320101013-1000010323023303-2000010111101101-0230111331202123-3231120012033013"></a>

<a id="canonical-3102213111231102-0003310012331131-1301003333010230-1111223020002220-3001313110010302-3123300330312333-2030112021132323-1213010031311330"></a>

#### `virtual_server.tcp.tcp_client_profile.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2031013330200223-0133032220302133-3223210223232202-1110331220121021-0213221001321033-0112020211120203-3232331310200220-1120233212320331"></a>

<a id="canonical-2003120001132130-0203113001333321-1031233231012103-1000211032132300-1021220311223110-2311130310002103-0303203222100123-1002310000222013"></a>

#### `virtual_server.tcp.tcp_client_profile.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0032321113220230-2210221333303233-0123322330001031-0211010312302301-0111111221033222-2310210331211213-3000322022123112-2231210311122112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.tcp.tcp_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.tcp](resources--application_profiles--reference--group-004.md#canonical-0133012121301311-1302322230030102-1233223310322132-2201222201110320-2110222212330303-1221023322200210-2120320110031323-0203220310230231)
- virtual_server.tcp.tcp_server_profile

<a id="canonical-3122102111112221-0210010201323231-2021200021100112-1020010323310032-1011031101220320-1111000012210123-0010330130312223-1231232220122133"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3130111120013310-0313130330202112-3211310020200032-0320100133002301-1332022030012320-2111010302220212-1031013313132330-0113032110230200"></a>

### Direct properties for `virtual_server.tcp.tcp_server_profile`

<a id="canonical-1303000301321302-3323023113031001-2131111011202020-1113011322321101-3101312301111011-2100030020031300-2221222311010002-0030123213213302"></a>

#### `virtual_server.tcp.tcp_server_profile.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1010002221322231-2200211120100222-3223323222032100-0030000033123203-1201012112031311-1201213011010333-0222033300320320-2032120100321330"></a>

<a id="canonical-3233221322132033-2002300132323023-0310021312013323-2023120020021002-1333010222332202-3122103230102303-1311212300320131-0031132332013012"></a>

#### `virtual_server.tcp.tcp_server_profile.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2120132210000323-2102211130312203-1223202021100201-2131133002303321-3203123113221012-2213313132001133-1102132201232002-1021032020110230"></a>

<a id="canonical-2123302123301123-0123212220132132-2222110032113223-2110321132301323-3322221331333232-3210322103013121-2321000123313101-1000101130333321"></a>

#### `virtual_server.tcp.tcp_server_profile.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2030130300331200-1133103021231303-0312010033233120-1332321301230300-2002110230112212-2010022302000312-0122331020010202-3031032010011133"></a>

<a id="canonical-1221302203311332-2202121112201323-1103231021123231-2203212220331002-3011312213113013-3101032010121231-2032322112311320-3021002332112232"></a>

#### `virtual_server.tcp.tcp_server_profile.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3222222311333301-2121222012111010-1002102211212032-2321032311302023-2312132301021221-2013023103122003-2023330233313221-3202312331301211"></a>

<a id="canonical-1102222100000201-1101011030310202-2121210311123033-2022310011010223-3123113102310023-0312210032010020-0213313332103022-0333223103131101"></a>

#### `virtual_server.tcp.tcp_server_profile.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.udp` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.udp

<a id="canonical-0202310031202010-2013230330033220-3203120030300320-3233311000332012-2333000000230213-3011230113132321-2311033300201330-2320323313023303"></a>

Type: `"object"`. single nested block, Optional.

UDP profiles.

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
udp {
  # Configure direct properties listed below.
}
```

<a id="canonical-1100331010201223-0033210310232321-0311112111223232-0102012303303320-0333000203103303-2101220311133200-2111221233001012-0332321233021102"></a>

### Direct properties for `virtual_server.udp`

- [client_ssl_profile](resources--application_profiles--reference--group-004.md#canonical-1122130210022223-2231220222323121-1323211202123333-3030233201003223-3323032210220331-1201311331312321-1032111223231330-2230322133113123): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--reference--group-004.md#canonical-2021322022310220-0032331120333010-1323213313312110-2023101111222200-2330113101321231-2322220013230220-2100331121121300-2000302300221003): complete subsection reference.

- [udp_client_profile](resources--application_profiles--reference--group-004.md#canonical-0303321111021332-3021213130021133-2210022131031212-1110100113102131-1321300011032303-0202032201130003-1000100312002222-0103020212123120): complete subsection reference.

- [udp_server_profile](resources--application_profiles--reference--group-004.md#canonical-1132032320032120-0030311302123110-3033222311022123-2300300010000331-0320011113122103-2012231120102222-1010023023221231-1103333021102012): complete subsection reference.

<a id="canonical-1122130210022223-2231220222323121-1323211202123333-3030233201003223-3323032210220331-1201311331312321-1032111223231330-2230322133113123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.udp.client_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.udp](resources--application_profiles--reference--group-004.md#canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313)
- virtual_server.udp.client_ssl_profile

<a id="canonical-3033031223130232-3113223130200330-2021133322032131-2121022022233233-0110300131213332-1221213010021110-2033230033101033-1301003103210032"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2300201201133210-2231322203211331-2102222212120220-1002310012111301-2303300001302023-1132003333230112-3300000023131023-0030002013232113"></a>

### Direct properties for `virtual_server.udp.client_ssl_profile`

<a id="canonical-0112221330330033-0123100222330113-0000131223033231-3122001033223013-2313131213121103-3032233332110211-0112002133302311-0012332122121212"></a>

#### `virtual_server.udp.client_ssl_profile.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3310003111312011-0110110302200230-1310313220211111-1320110221232332-1232133312013202-1032321212101131-3233002001333230-0003032200000230"></a>

<a id="canonical-1323133131023003-3200122012310202-0003012331012202-0022212322111300-2221133200310132-2212233313203223-1130230010132322-2031032222222031"></a>

#### `virtual_server.udp.client_ssl_profile.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2203021033323210-3131000030120021-3332320122220013-1032233323200031-1123322132130213-3232103210212020-0312111222223001-3103011111101010"></a>

<a id="canonical-0121220313123332-0202322122203202-0223010003321202-2112032101212002-2101112313022213-3200232233203121-3002132213023030-0331313111012321"></a>

#### `virtual_server.udp.client_ssl_profile.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3303230300320023-2312211010220130-3120021233220312-2311103213322001-0320231313302113-3200110203101102-1230220101132132-3111200323330002"></a>

<a id="canonical-3110231011123211-3131021123233121-1230113303131300-1220001220303012-2133300013011223-0112001101211131-2031212031312211-0301313023112313"></a>

#### `virtual_server.udp.client_ssl_profile.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1021230312010001-3323222330013223-0312201330133212-0310300310330323-1130122012220120-3013113132311122-2320000302220020-1022320232202003"></a>

<a id="canonical-0113132233003133-0331000220222321-1033123000211020-0330211130201212-1010112001320023-1300000023023010-0211310311221123-1210100210032133"></a>

#### `virtual_server.udp.client_ssl_profile.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2021322022310220-0032331120333010-1323213313312110-2023101111222200-2330113101321231-2322220013230220-2100331121121300-2000302300221003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.udp.server_ssl_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.udp](resources--application_profiles--reference--group-004.md#canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313)
- virtual_server.udp.server_ssl_profile

<a id="canonical-3003003201122202-0223010232323030-0120011011231320-2210202332111112-1123030111021102-0231120121330001-1203103020331123-0003003303322302"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2131023011023011-0303113103220300-2012001322021300-3110113131233303-0003321311312321-2013133003230231-0121131323302013-2321330303321022"></a>

### Direct properties for `virtual_server.udp.server_ssl_profile`

<a id="canonical-3330011101321333-2102122213332322-3331123203310012-1320120133210330-2131113210033133-0021313322130030-2021220000200022-2211021220032103"></a>

#### `virtual_server.udp.server_ssl_profile.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1313322120200101-0001100011203130-2222023031332012-1313211011110230-2100322022323232-2300311032020022-3131301121032133-3302213120130011"></a>

<a id="canonical-3023033321031201-0331121001221031-3031031131230030-3330030201220300-1112112310301232-0131302000202320-3332022223011110-0133312320330101"></a>

#### `virtual_server.udp.server_ssl_profile.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1221003120222303-2113123323330112-1120232022033123-2211003222320221-1200120022301111-2023021012211121-1322023011013201-0022313301010113"></a>

<a id="canonical-2320122112212011-2203221020100032-3123212211132021-3300203310332312-2120011020123100-1310210132211103-1312312332023212-2233020111022231"></a>

#### `virtual_server.udp.server_ssl_profile.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0222121030223213-1110131020220332-3323322211320002-3131033333022111-0132230222133020-2103303023322220-3022021002231122-3302011210201131"></a>

<a id="canonical-2330313200132023-0320113003202013-3313020003311120-1303112101222110-1012002232332332-3312133322011311-1320201103320331-0111232202013121"></a>

#### `virtual_server.udp.server_ssl_profile.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3013133102112210-1320121112233011-1133130310232133-1032320211021023-3300021022330023-3221103221010123-1102013113013303-3003200012300220"></a>

<a id="canonical-1110103012322130-2030221303232232-1020231032301221-1231320102023333-1332231333132013-2320310133311122-3133202223021312-3012110131121220"></a>

#### `virtual_server.udp.server_ssl_profile.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0303321111021332-3021213130021133-2210022131031212-1110100113102131-1321300011032303-0202032201130003-1000100312002222-0103020212123120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.udp.udp_client_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.udp](resources--application_profiles--reference--group-004.md#canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313)
- virtual_server.udp.udp_client_profile

<a id="canonical-3102232222323031-2323123101221120-2230023032121103-1011120012111311-1103321001112222-0301133311233110-0213100312222021-0122013311111330"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2110100332203322-3033030010213320-1331333312331223-1110222231333020-1121022221101330-1022100032320323-1310223311212303-2232321111120231"></a>

### Direct properties for `virtual_server.udp.udp_client_profile`

<a id="canonical-2030030302310122-3211100003112200-1223203112013030-1113312033320103-2123330331330303-0211100120030331-0231211210212131-1113230110001003"></a>

#### `virtual_server.udp.udp_client_profile.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3102132213113330-1100000322332332-0213231020133021-0013313220102203-3333123202303023-0003131110020201-1103123231120101-3012111333313211"></a>

<a id="canonical-1200023213010122-2210012231123302-3200131312020212-0310323122033322-3213332103020302-3201131110223311-0222231111223312-2122130231201202"></a>

#### `virtual_server.udp.udp_client_profile.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2321111021331113-1232122033013100-3333200333220210-0001322102210200-1331333111123130-2033012002030031-1320102202202203-0323020011312012"></a>

<a id="canonical-0210200333210312-3120023132033233-0101013332332221-0013200301132310-1120211100112021-1122223223320301-3131212333113232-1122210031313332"></a>

#### `virtual_server.udp.udp_client_profile.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0032100333201103-0220001121020133-2221320033230312-2210200323112023-2001120130030213-0302222113003011-0023310301120230-3321030330121022"></a>

<a id="canonical-2102333113301110-0132010310210201-3210032313102032-2122113333010212-1200033001010220-0231132331332033-3103131202132200-1233101013010303"></a>

#### `virtual_server.udp.udp_client_profile.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1300012133023202-0031103230103113-2020122322002101-3100100020132020-2233112012330303-0201332200033202-2130111122123301-0301023110220103"></a>

<a id="canonical-2231113320110122-1331332333313020-1111021111302303-2020323301313023-2322130233230210-0021132333030323-1011133011122001-3321023210122001"></a>

#### `virtual_server.udp.udp_client_profile.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1132032320032120-0030311302123110-3033222311022123-2300300010000331-0320011113122103-2012231120102222-1010023023221231-1103333021102012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.udp.udp_server_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.udp](resources--application_profiles--reference--group-004.md#canonical-0100212231103000-2031131131020010-3003110221032012-0212301320022111-2133230201233310-2233210103212202-3223201320002311-1322101113222313)
- virtual_server.udp.udp_server_profile

<a id="canonical-0230123203221322-0330012221231312-0321110103233220-1311131131303031-0212322030321232-1001002021323123-1333323130323301-0303213012323011"></a>

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1213030232223212-0031331213023232-1312100320113210-3330320332232200-0230032300313013-0213313331131112-3032212123322330-0103210001023032"></a>

### Direct properties for `virtual_server.udp.udp_server_profile`

<a id="canonical-3333223030203310-3022013212003103-1220121302230210-2023313013012312-0112010202020300-0302002003213321-2311120013310122-3023312002011220"></a>

#### `virtual_server.udp.udp_server_profile.kind` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2120003111113001-3231302300033013-2112210312301322-3030223033023233-2100013120213022-3111230331012331-2022020002012113-0033331232310103"></a>

<a id="canonical-0021123122101330-2330203112023022-0023020101230111-2303313011330202-3102212202021300-0013220222200222-1003203212102010-1123211210021120"></a>

#### `virtual_server.udp.udp_server_profile.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0001023012113110-0133300100233311-2203122230012223-0000220021112210-3002112020313031-2322210333210311-2301113100333313-1203301001031003"></a>

<a id="canonical-1002310232223102-0230003322331303-2110131120200231-2332132131031001-3133331230001022-2012131221120231-3223203203022102-3132332333122222"></a>

#### `virtual_server.udp.udp_server_profile.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0231311131102033-1200223220033132-2301230000211130-1321232200332202-0030230310020323-2012031321313002-3113310002201220-2022013320310022"></a>

<a id="canonical-1312113023120331-2311011303121022-1112023203132333-2320111303113121-1221001121211023-0320012302010323-0302112200323023-3020320210030111"></a>

#### `virtual_server.udp.udp_server_profile.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0100303333012021-2200032023233320-2121312121001322-0031313023212230-1202011231110003-3301033311302232-3320333220313223-2012322201011332"></a>

<a id="canonical-3223033011202312-0131201020330121-2310023033000112-2310322013222323-0100003003002030-0233232222111320-1103100232212312-3021001223210211"></a>

#### `virtual_server.udp.udp_server_profile.uid` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0323021232302203-3330212203221312-0021222111110121-2230020130302203-1311230313200231-3310002022201103-2331300211333032-3021322223300011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.virtual_server_state` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.virtual_server_state

<a id="canonical-0122303122010231-1321031133131130-1202302030103331-3111033233131023-3101020302031303-3022112323200233-2003020011121311-2213000201203322"></a>

Type: `"object"`. single nested block, Optional.

Displays the current state on the object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("state_disabled",
    "state_enabled")}
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
  "x-ves-oneof-field-state_choice": "[\"state_disabled\",\"state_enabled\"]"
}
```

Terraform syntax:

```terraform
virtual_server_state {
  # Configure direct properties listed below.
}
```

<a id="canonical-1230011010102021-0031120032011202-0321223001323110-3230202322333122-3133233203121302-3310323331232310-2130332232132033-0221223322310221"></a>

### Direct properties for `virtual_server.virtual_server_state`

- [state_disabled](resources--application_profiles--reference--group-004.md#canonical-0121331000132220-1313223010021111-1300122021210133-2121010231300102-3330300311303212-0031200013322333-1123123322332330-2010212310333330): complete subsection reference.

- [state_enabled](resources--application_profiles--reference--group-004.md#canonical-2223230303100333-0102311020102121-3202121131020232-2012301223110233-3111211321210113-1033030133000223-2033131020033201-1222211122320011): complete subsection reference.

<a id="canonical-0121331000132220-1313223010021111-1300122021210133-2121010231300102-3330300311303212-0031200013322333-1123123322332330-2010212310333330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.virtual_server_state.state_disabled` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.virtual_server_state](resources--application_profiles--reference--group-004.md#canonical-0323021232302203-3330212203221312-0021222111110121-2230020130302203-1311230313200231-3310002022201103-2331300211333032-3021322223300011)
- virtual_server.virtual_server_state.state_disabled

<a id="canonical-0120301130000121-0203210122211133-1021132101231303-0101332022233021-2133030330213310-3301330232222331-0202031022230300-1121322301311323"></a>

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
state_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223230303100333-0102311020102121-3202121131020232-2012301223110233-3111211321210113-1033030133000223-2033131020033201-1222211122320011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.virtual_server_state.state_enabled` properties

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.virtual_server_state](resources--application_profiles--reference--group-004.md#canonical-0323021232302203-3330212203221312-0021222111110121-2230020130302203-1311230313200231-3310002022201103-2331300211333032-3021322223300011)
- virtual_server.virtual_server_state.state_enabled

<a id="canonical-2030200221303101-1201013133000300-1132311030000301-3132012300330200-2133023103032200-2032232030013012-1332111122103212-3201113121302101"></a>

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
state_enabled = {}
```

This is an empty object or choice marker. It has no direct properties.
