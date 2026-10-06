---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2223012312223021-2322211111100331-2013203110223222-2332033120011030-0213120101002332-0213222103312012-3231222122112222-1200302102302232"></a>

## Direct properties for `trusted_clients.http_header.headers`

<a id="canonical-1010313022332111-1030222321331022-3303110221321031-1031113001100030-2223223220013021-2233103020133103-3111310030032312-3001101323310002"></a>

### `trusted_clients.http_header.headers.exact` property

Type: `"string"`. Computed.

Exclusive with \[presence regular expression\] Header value to match exactly.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-3322322133110020-2020022303210100-1102311320203032-2132212000102321-1010202301102330-3110320121311310-0320220311210313-2101103303102321"></a>

<a id="canonical-3103003101330301-1220223213100323-1312311330212102-1222233103311201-1303000331002323-0030123212131133-2232210010022312-0120320320312231"></a>

### `trusted_clients.http_header.headers.invert_match` property

Type: `"bool"`. Computed.

Invert the result of the match to detect missing header or non-matching value.

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

<a id="canonical-3333003003132321-2000323011212001-0301012200231103-2012300102023223-1022121030022203-1122032131100030-0321222131011331-1010303221330222"></a>

<a id="canonical-3022213101222020-2320311102312222-0110332311002333-2112032103030030-2103113321230033-0212213103111011-1030310120322213-2122233122011231"></a>

### `trusted_clients.http_header.headers.name` property

Type: `"string"`. Computed.

Name. Name of the header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2202332113323133-0332212101103131-1003233211121101-3002033322231310-3130331020330111-2231313011230102-1211131011002123-3311013013333311"></a>

<a id="canonical-3023003222332021-2100121131312223-1011131210102302-2121103123010123-2102311133200321-1310132031303110-3323023302312011-2031203031210101"></a>

### `trusted_clients.http_header.headers.presence` property

Type: `"bool"`. Computed.

Exclusive with \[exact regular expression\] If true, check for presence of header.

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

<a id="canonical-3301100302310222-3303220310322200-0203302103223031-0311011301112321-1001020012011302-1003110103101232-3223110102033112-0203033220012101"></a>

<a id="canonical-3330213003000013-3133020302110302-3122103123331031-2120231302012311-0332213003223123-1103302210211230-2121113001112121-2331030303223113"></a>

### `trusted_clients.http_header.headers.regex` property

Type: `"string"`. Computed.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3202020021333101-0232122300001313-1310212212321110-3311231212132311-2313112132132232-3202311332122013-3132131301122321-3232311311103133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [trusted_clients](data-sources--http_loadbalancer--reference--group-026.md#canonical-3130231101303131-0231102133202313-3123310310222112-3001133311030323-2321323001012002-0302320012032323-1332011122113301-0123003323213020)
- trusted_clients.metadata

<a id="canonical-2323232002303233-0222332000121301-3011201221212132-3200031113032102-0211100311121320-2011213300120302-0331103110001012-3213320211022122"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-2300132332111331-1313130303232101-0330321132103121-1012133202001233-1130202011312110-2031102310102322-2121313000010203-1301001302201011"></a>

### Direct properties for `trusted_clients.metadata`

<a id="canonical-1020211221102321-1010033231023103-1031233022012113-2022211202013310-3023120011311023-0032101021223213-0113212112011100-1010003003323023"></a>

#### `trusted_clients.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2311202330230010-1122133132031121-3010112003313033-0202110222103120-0123233323111200-2303332230103033-3102101010323000-1030030300121111"></a>

<a id="canonical-1011001010231020-2212023003220002-1331210331110111-1111311130300302-1221211212021022-3303222310001211-1013013030110322-2010223211333321"></a>

#### `trusted_clients.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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

<a id="canonical-3112332033131101-3011121131303030-2322233022200023-1033011221031321-1221020032320031-2131032122111013-2221303131132220-2113010210201012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.skip_processing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [trusted_clients](data-sources--http_loadbalancer--reference--group-026.md#canonical-3130231101303131-0231102133202313-3123310310222112-3001133311030323-2321323001012002-0302320012032323-1332011122113301-0123003323213020)
- trusted_clients.skip_processing

<a id="canonical-3112023232220230-2231321030102323-2320211210130111-0132222221100131-0011102011030003-2300021121210313-0322301313001001-0232111032031312"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002202033200000-2031310200300103-2022023232313133-1301002112320111-0320301311221130-1231221131023322-0201122332032300-3310012101133310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `trusted_clients.waf_skip_processing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [trusted_clients](data-sources--http_loadbalancer--reference--group-026.md#canonical-3130231101303131-0231102133202313-3123310310222112-3001133311030323-2321323001012002-0302320012032323-1332011122113301-0123003323213020)
- trusted_clients.waf_skip_processing

<a id="canonical-2200303300333111-1303003110323332-0102320122312033-0000031210102012-3233123112323102-2103032310112310-0333121030220332-0303022110203300"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223203020310031-1320313232132320-1010131100213310-3102302311120331-1313020012310011-2302030032233032-3322311032002211-0310020323003300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_id_client_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- user_id_client_ip

<a id="canonical-1202023213131320-2202333331333320-2202103101012321-0323111200021302-0032300312221111-3022031322320021-1000333220122113-0230322033101131"></a>

Type: `["object", {}]`. Computed.

\[OneOf: user\_id\_client\_ip, user\_identification\] Enable this option. Defaults to \`map\[\]\`.
Server applies default when omitted.

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

OneOf alternatives in this subsection:

- [user_id_client_ip](data-sources--http_loadbalancer--reference--group-027.md#canonical-1202023213131320-2202333331333320-2202103101012321-0323111200021302-0032300312221111-3022031322320021-1000333220122113-0230322033101131)
- [user_identification](data-sources--http_loadbalancer--reference--group-027.md#canonical-1032001011333213-2132203012120130-1022200332020333-2330131011103103-2010020033132133-2210300020022130-2030312122010102-2302132032030200)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300213310030023-1303230033010211-3001312312000330-2020213133231032-1301233011000133-1301323210201312-0220100111121100-0231103130020022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `user_identification` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- user_identification

<a id="canonical-1032001011333213-2132203012120130-1022200332020333-2330131011103103-2010020033132133-2210300020022130-2030312122010102-2302132032030200"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2301021311220110-1000311121023020-3303112320002320-2020303301122013-3132010312012021-0012100032332300-3311222021031101-3010110212132201"></a>

### Direct properties for `user_identification`

<a id="canonical-2211300332010213-2001121011031132-2100212210101111-1211221222021321-1031112021230103-1011332110011000-3231123132232221-0030330130031102"></a>

#### `user_identification.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-1311320100122220-1211110002313000-0013031132330023-2332103331332100-1003113302003211-1113201101223011-3011202210221003-1313012112202012"></a>

<a id="canonical-2021121202301213-2110203201032031-2103333212122213-0020132112210102-2313022210300201-1122112303103320-3110111213000110-3020231300112223"></a>

#### `user_identification.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3231012202300000-1032312201123232-0232320203022111-3033211132320220-1103020320221332-2312002222001322-0112030101311312-1212212333223210"></a>

<a id="canonical-1101030310011203-2132113310230213-0002003102302023-0311011022310011-2002300101301021-1202022203120131-3110210210313003-0232133010120112"></a>

#### `user_identification.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2201211320210323-1131233103012122-1202220210012030-1312122002301333-3032311002210133-3132322100021302-0321312033221231-3303001323312332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- waf_exclusion

<a id="canonical-2111232303203020-2100010322232103-1321111002332231-2330021032002133-1312203032212212-3312022303122103-1132113303310320-0220030201120332"></a>

Type: `"single"`. Computed.

Configuration parameter for waf exclusion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-waf_exclusion_choice": "[\"waf_exclusion_inline_rules\",\"waf_exclusion_policy\"]"
}
```

<a id="canonical-2321230131232133-0212320020333212-1311232101113111-3223220022113003-3310110031110231-0131331321313133-3102102320010132-0213103123303310"></a>

### Direct properties for `waf_exclusion`

- [waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-0230300230321033-0121023132033121-2331203221203221-0213211112222231-0101232320213100-1002301230313112-1010311021120321-3003012233030302): complete subsection reference.

- [waf_exclusion_policy](data-sources--http_loadbalancer--reference--group-027.md#canonical-1321011111113021-0321110131231130-3000230231112102-0230320013311302-1310012003013103-0230023330201311-1121112302211313-3103211000231200): complete subsection reference.

<a id="canonical-0230300230321033-0121023132033121-2331203221203221-0213211112222231-0101232320213100-1002301230313112-1010311021120321-3003012233030302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-027.md#canonical-2201211320210323-1131233103012122-1202220210012030-1312122002301333-3032311002210133-3132322100021302-0321312033221231-3303001323312332)
- waf_exclusion.waf_exclusion_inline_rules

<a id="canonical-1120213200310003-3323211232030302-2201021211000232-3213032333233010-1222321103010123-0212111000202321-2132101212303123-2003013321202021"></a>

Type: `"single"`. Computed.

A list of WAF exclusion rules that will be applied inline.

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

<a id="canonical-2202222120211120-2303020211331332-0312302031220330-1022230201300323-1021321331030131-0333312230203122-2332320330001313-2132323113310013"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules`

- [rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-1130202033312000-2333332311213331-2201123211031102-3000332303130330-0301021313211302-3302232003021022-0000301232023102-2130202013123322): complete subsection reference.

<a id="canonical-1130202033312000-2333332311213331-2201123211031102-3000332303130330-0301021313211302-3302232003021022-0000301232023102-2130202013123322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-027.md#canonical-2201211320210323-1131233103012122-1202220210012030-1312122002301333-3032311002210133-3132322100021302-0321312033221231-3303001323312332)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-0230300230321033-0121023132033121-2331203221203221-0213211112222231-0101232320213100-1002301230313112-1010311021120321-3003012233030302)
- waf_exclusion.waf_exclusion_inline_rules.rules

<a id="canonical-3320200231300021-2231200110111112-1010020321212022-1110133321310021-0320002302121010-1102033033013331-0232323331020313-3320031213201333"></a>

Type: `"list"`. Computed.

An ordered list of WAF Exclusions specific to this Load Balancer.

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
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-2313231300012011-3333120020030121-2312201323112111-1030323021302131-3321131123300020-3112330310232321-1311201233223303-0120202200231113"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules`

- [any_domain](data-sources--http_loadbalancer--reference--group-027.md#canonical-1010312033200003-1112333121232003-3202313133303230-2220222300300230-0102102131212313-0220331031222322-1033312202303301-1322030110021322): complete subsection reference.

- [any_path](data-sources--http_loadbalancer--reference--group-027.md#canonical-0000010011323010-3121012001033120-1030031102121000-3212121000223121-3123031123210230-1103202023132232-1310010323030203-0212230002230130): complete subsection reference.

- [app_firewall_detection_control](data-sources--http_loadbalancer--reference--group-027.md#canonical-2120120301030233-2330103012022230-2001030220203323-0330210200130302-3220010012002212-2313231301321131-2002111100021002-2120322320101131): complete subsection reference.

<a id="canonical-0210213103111011-2203001320231210-1233222023033002-3112100011301330-0310202321331331-1223012332032201-3333121211323322-3031211212211013"></a>

<a id="canonical-3322030312331013-0030013001031203-1100311113231333-0302001202032210-1232133321001330-3323330013213103-0212022003122032-2120111311110302"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3112002003323211-3100013101320130-1123130213130003-0123312213302321-3012011010303130-3113001033321102-3322230122312001-2320031112220203"></a>

<a id="canonical-0123021212001102-2231031330102321-0111021200301000-2300030103330003-3033321132013321-1233131001033121-1211122003013133-3023113230011213"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.expiration_timestamp` property

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Additional upstream details:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

- [metadata](data-sources--http_loadbalancer--reference--group-027.md#canonical-2100023330203021-3230010322023212-3220023033222203-2021031303313221-1312220002111123-2010323212111222-2110103221221223-0331313002320020): complete subsection reference.

<a id="canonical-2233120103321311-2200113003231002-0123203002311000-2022323311213002-0230233133300312-2033132323313000-0132231123131111-1122131301021213"></a>

<a id="canonical-2030303302010103-0111010311103211-0122002200300011-0223323230110113-2111120322121000-3003132200230101-3130122301221020-3023032331131231"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.methods` property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3320232312313223-2332103001330100-3210112000022033-0303201223031002-2132133133212211-1122011113132133-0320021322133032-2022130203120203"></a>

<a id="canonical-3012333120112310-2200130130011130-1222220233332122-1303200233310212-1030122320211302-0333022023110331-1332133333113021-3022112130132321"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.path_prefix` property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1123211331332210-1303201223310231-1003102020112213-3003213120231321-3211001220222230-3122212230100312-1032300023032032-0110330033122222"></a>

<a id="canonical-1300301212001212-2033133030121000-1110001221131333-2012311313200220-3210303003302310-1103302212120331-0313203011330221-0320031212131312"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.path_regex` property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\] Define the regular expression for the path. For example, the regular expression
^/.\*$ will match on all paths.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3101302321033302-1302000302130213-1331333113331013-2133333110011201-2022003203203122-1303030002112123-0311201311102310-2211211213203132"></a>

<a id="canonical-1301101001022233-0101230032011332-1023111112013221-3102023231000212-1231311303002322-0012212231022331-2322120231103223-3201130101302112"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [waf_skip_processing](data-sources--http_loadbalancer--reference--group-027.md#canonical-0220122011122223-2110203320003322-2100333232013333-3110031331021101-3302032121122003-2102223223111322-2120033031232200-1023133130102101): complete subsection reference.

<a id="canonical-1010312033200003-1112333121232003-3202313133303230-2220222300300230-0102102131212313-0220331031222322-1033312202303301-1322030110021322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-027.md#canonical-2201211320210323-1131233103012122-1202220210012030-1312122002301333-3032311002210133-3132322100021302-0321312033221231-3303001323312332)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-0230300230321033-0121023132033121-2331203221203221-0213211112222231-0101232320213100-1002301230313112-1010311021120321-3003012233030302)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-1130202033312000-2333332311213331-2201123211031102-3000332303130330-0301021313211302-3302232003021022-0000301232023102-2130202013123322)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_domain

<a id="canonical-1010002122120213-2023213032312221-3213130323002133-2201321100003013-0132003021311012-0022102320232132-1121113203022010-1113323022203010"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0000010011323010-3121012001033120-1030031102121000-3212121000223121-3123031123210230-1103202023132232-1310010323030203-0212230002230130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.any_path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-027.md#canonical-2201211320210323-1131233103012122-1202220210012030-1312122002301333-3032311002210133-3132322100021302-0321312033221231-3303001323312332)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-0230300230321033-0121023132033121-2331203221203221-0213211112222231-0101232320213100-1002301230313112-1010311021120321-3003012233030302)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-1130202033312000-2333332311213331-2201123211031102-3000332303130330-0301021313211302-3302232003021022-0000301232023102-2130202013123322)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_path

<a id="canonical-0100123002300013-3111133212331220-0020021032113020-3120232233131120-3332023210001030-0032133332332122-2003333120102233-3122232302202130"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120120301030233-2330103012022230-2001030220203323-0330210200130302-3220010012002212-2313231301321131-2002111100021002-2120322320101131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-027.md#canonical-2201211320210323-1131233103012122-1202220210012030-1312122002301333-3032311002210133-3132322100021302-0321312033221231-3303001323312332)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-0230300230321033-0121023132033121-2331203221203221-0213211112222231-0101232320213100-1002301230313112-1010311021120321-3003012233030302)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-1130202033312000-2333332311213331-2201123211031102-3000332303130330-0301021313211302-3302232003021022-0000301232023102-2130202013123322)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control

<a id="canonical-0231021211233202-2200213313332222-1302100233330323-2230122233330100-0132313222130210-0301002303300013-3202133230003032-0012122300012323"></a>

Type: `"single"`. Computed.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

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

<a id="canonical-0013212131332202-0001110133102302-3101302111220300-2322021021222222-0032032110333202-2203112213312233-3011021013301333-1021222121022111"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control`

- [exclude_attack_type_contexts](data-sources--http_loadbalancer--reference--group-027.md#canonical-3013020302032232-1320331102231111-3221100003232321-0013321230120003-3231120310010222-1012312223012222-3131333221100200-2030113300312131): complete subsection reference.

- [exclude_bot_name_contexts](data-sources--http_loadbalancer--reference--group-027.md#canonical-2320111121220130-3110122203313010-3303310033030322-3012002030220101-1112331232002132-0233332113031113-3313313013121122-3230232302011301): complete subsection reference.

- [exclude_signature_contexts](data-sources--http_loadbalancer--reference--group-027.md#canonical-2123320221303120-2222030302111001-0311132210230132-1310021033313133-0110010110030312-3112300210103330-3233310010113033-0011131300120133): complete subsection reference.

- [exclude_violation_contexts](data-sources--http_loadbalancer--reference--group-027.md#canonical-2230133022112101-3100131210013321-0331100330222123-1310310110012222-3300211102222311-2303200323133102-2012320322123210-0233220200002030): complete subsection reference.

<a id="canonical-3013020302032232-1320331102231111-3221100003232321-0013321230120003-3231120310010222-1012312223012222-3131333221100200-2030113300312131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-027.md#canonical-2201211320210323-1131233103012122-1202220210012030-1312122002301333-3032311002210133-3132322100021302-0321312033221231-3303001323312332)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-0230300230321033-0121023132033121-2331203221203221-0213211112222231-0101232320213100-1002301230313112-1010311021120321-3003012233030302)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-1130202033312000-2333332311213331-2201123211031102-3000332303130330-0301021313211302-3302232003021022-0000301232023102-2130202013123322)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--http_loadbalancer--reference--group-027.md#canonical-2120120301030233-2330103012022230-2001030220203323-0330210200130302-3220010012002212-2313231301321131-2002111100021002-2120322320101131)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-3122110010121200-3021312022312103-2203230232121001-0311212201211130-0023012013000203-0213011331120100-2021133020200031-2031200321222301"></a>

Type: `"list"`. Computed.

Exclude an entire attack type only in the named context. For migrated per-parameter exceptions,
prefer this over signature-ID exclusions because one payload can trigger several signatures;
unrelated parameters and attack types remain protected.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3310303010223220-0130310321302332-0131010032310122-3121323121223130-1012302230021131-2332232222022233-0102022031231022-0101110201313311"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts`

<a id="canonical-3012322120033321-1133003322311100-3133332202222003-0111310313303323-1113000022300131-1131312001133030-0101310203013032-2111322102033011"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts.context` property

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Additional upstream details:

Exclusion scope. Use CONTEXT\_PARAMETER with context\_name for one parameter, CONTEXT\_COOKIE for
one cookie, or CONTEXT\_ANY only for an intentionally global scope.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3103211211012013-0302202200333323-3133203222100000-2130001301220120-0231330011022211-3313222203332323-3102301231233012-0322321202303303"></a>

<a id="canonical-1312001201102332-2212101021023232-3032310312110131-1331102102223020-1122012311322000-2130102230011312-0113013010121000-1000301123320112"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts.context_name` property

Type: `"string"`. Computed.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0330101123132322-3100311030020021-2101100013213302-3000000002222100-0323230020011011-1210012100022201-3111111320030013-3202132111103303"></a>

<a id="canonical-3121222011300001-3023320010101231-3303021312200223-3313310200031023-2231313230033131-2132331133310011-3313301331220122-0231112030003313"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts.exclude_attack_type` property

Type: `"string"`. Computed.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of all Attack Types ATTACK\_TYPE\_NONE ATTACK\_TYPE\_NON\_BROWSER\_CLIENT
ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE
ATTACK\_TYPE\_DETECTION\_EVASION ATTACK\_TYPE\_VULNERABILITY\_SCAN
ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS..
Possible values are \`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Additional upstream details:

Attack-type enum excluded in this context, for example ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING. Other
attack types remain enforced.

Receipt-pinned upstream constraints:

```json
{
  "default": "ATTACK_TYPE_NONE",
  "enum": [
    "ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2320111121220130-3110122203313010-3303310033030322-3012002030220101-1112331232002132-0233332113031113-3313313013121122-3230232302011301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-027.md#canonical-2201211320210323-1131233103012122-1202220210012030-1312122002301333-3032311002210133-3132322100021302-0321312033221231-3303001323312332)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-0230300230321033-0121023132033121-2331203221203221-0213211112222231-0101232320213100-1002301230313112-1010311021120321-3003012233030302)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-1130202033312000-2333332311213331-2201123211031102-3000332303130330-0301021313211302-3302232003021022-0000301232023102-2130202013123322)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--http_loadbalancer--reference--group-027.md#canonical-2120120301030233-2330103012022230-2001030220203323-0330210200130302-3220010012002212-2313231301321131-2002111100021002-2120322320101131)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-3123013111010132-0312033123111033-2012112012303231-3030003003202112-1001123300030123-1023203132201220-1033321300312301-0012320121202000"></a>

Type: `"list"`. Computed.

Bot Names to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2030302010020121-0302213322001112-0232320210111213-0030131130002301-0123313332233122-0313020130020022-2232231121130002-0301133230220233"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts`

<a id="canonical-3222110013021331-2121200313322310-3103222001301122-3333333131221033-1311101320020123-1113032130213323-3123103322022300-2011202020221213"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts.bot_name` property

Type: `"string"`. Computed.

Bot Name. Human-readable name for the resource

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

<a id="canonical-2123320221303120-2222030302111001-0311132210230132-1310021033313133-0110010110030312-3112300210103330-3233310010113033-0011131300120133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-027.md#canonical-2201211320210323-1131233103012122-1202220210012030-1312122002301333-3032311002210133-3132322100021302-0321312033221231-3303001323312332)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-0230300230321033-0121023132033121-2331203221203221-0213211112222231-0101232320213100-1002301230313112-1010311021120321-3003012233030302)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-1130202033312000-2333332311213331-2201123211031102-3000332303130330-0301021313211302-3302232003021022-0000301232023102-2130202013123322)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--http_loadbalancer--reference--group-027.md#canonical-2120120301030233-2330103012022230-2001030220203323-0330210200130302-3220010012002212-2313231301321131-2002111100021002-2120322320101131)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-3220311120330111-3003000012231033-1023130330221301-0230130012310113-1200100132102313-2220221023123332-2131003211232203-1030202233121210"></a>

Type: `"list"`. Computed.

Signature IDs to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3122231001213102-1331021003232210-0013230310130332-1131120233112022-2030302012220233-0303222301203010-0021113301210033-1233233022213103"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts`

<a id="canonical-1332301221331020-3130121331112203-0120222221313132-0132212010120302-1303133133330202-3330220122122230-2313321011332223-2202001131113101"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts.context` property

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Additional upstream details:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0312202312030201-3122212031203021-0223112333131001-3303021130022013-0222213021032133-3200133300121220-1030311232323031-0133202030031220"></a>

<a id="canonical-3223210333133233-1003203301221030-3320301232311022-3312012003002213-1011002032022110-0231131302133033-0212302102201101-3111202230220113"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts.context_name` property

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2022101300013120-1023202030323123-0201211313002110-0022332122212131-3321322230313212-1231202031001123-1102311123000133-2132202232211313"></a>

<a id="canonical-3322100300100112-0213030023010200-0012302023223100-3222102202033312-2003321202332201-2021311232033231-2333031122011020-2213110222230311"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts.signature_id` property

Type: `"number"`. Computed.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 299999999,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  }
}
```

<a id="canonical-2230133022112101-3100131210013321-0331100330222123-1310310110012222-3300211102222311-2303200323133102-2012320322123210-0233220200002030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-027.md#canonical-2201211320210323-1131233103012122-1202220210012030-1312122002301333-3032311002210133-3132322100021302-0321312033221231-3303001323312332)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-0230300230321033-0121023132033121-2331203221203221-0213211112222231-0101232320213100-1002301230313112-1010311021120321-3003012233030302)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-1130202033312000-2333332311213331-2201123211031102-3000332303130330-0301021313211302-3302232003021022-0000301232023102-2130202013123322)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--http_loadbalancer--reference--group-027.md#canonical-2120120301030233-2330103012022230-2001030220203323-0330210200130302-3220010012002212-2313231301321131-2002111100021002-2120322320101131)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-2131203122310212-3332003322130012-3302223030102333-3012222010303201-2301013323302002-0123111001303331-3112311221332012-3001102111301132"></a>

Type: `"list"`. Computed.

Violations to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2222001323022011-2110300020311010-0322033030313213-0131112222303121-3110331123221210-0323201202030302-0123012212120101-1323323310002101"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts`

<a id="canonical-1202213133210023-2220223031231233-1210033212212200-2010232332130212-0133310120103301-0032132232310110-1022233000020311-1313013100332013"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts.context` property

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Additional upstream details:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3012100002202133-0212132311201011-1012332300113030-0002212022323103-2220203030203000-2112021332322330-2232101103321211-0313113012101001"></a>

<a id="canonical-3033221332330013-3333203302120021-1003323021001123-1230231002303033-0130123011213002-0220322023002022-1102031233000211-1110311003110202"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts.context_name` property

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-3323221022222313-0010323103201010-2230012122022121-1213212013022133-3313312131221233-1233202330103233-2333201001023102-1120020133002021"></a>

<a id="canonical-2233122121112210-3223332012133202-2030120233200331-1123120220003121-0100330030131010-0033033233101022-3210102221013223-1320321300010210"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts.exclude_violation` property

Type: `"string"`. Computed.

\[Enum:
VIOL\_NONE|VIOL\_FILETYPE|VIOL\_METHOD|VIOL\_MANDATORY\_HEADER|VIOL\_HTTP\_RESPONSE\_STATUS|VIOL\_REQUEST\_MAX\_LENGTH|VIOL\_FILE\_UPLOAD|VIOL\_FILE\_UPLOAD\_IN\_BODY|VIOL\_XML\_MALFORMED|VIOL\_JSON\_MALFORMED|VIOL\_ASM\_COOKIE\_MODIFIED|VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS|VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE|VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT|VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST|VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION|VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS|VIOL\_EVASION\_DIRECTORY\_TRAVERSALS|VIOL\_MALFORMED\_REQUEST|VIOL\_EVASION\_MULTIPLE\_DECODING|VIOL\_DATA\_GUARD|VIOL\_EVASION\_APACHE\_WHITESPACE|VIOL\_COOKIE\_MODIFIED|VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS|VIOL\_EVASION\_IIS\_BACKSLASHES|VIOL\_EVASION\_PERCENT\_U\_DECODING|VIOL\_EVASION\_BARE\_BYTE\_DECODING|VIOL\_EVASION\_BAD\_UNESCAPE|VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST|VIOL\_ENCODING|VIOL\_COOKIE\_MALFORMED|VIOL\_GRAPHQL\_FORMAT|VIOL\_GRAPHQL\_MALFORMED|VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\]
List of all supported Violation Types VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER
VIOL\_HTTP\_RESPONSE\_STATUS VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD
VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED
VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS.. Possible values are \`VIOL\_NONE\`,
\`VIOL\_FILETYPE\`, \`VIOL\_METHOD\`, \`VIOL\_MANDATORY\_HEADER\`, \`VIOL\_HTTP\_RESPONSE\_STATUS\`,
\`VIOL\_REQUEST\_MAX\_LENGTH\`, \`VIOL\_FILE\_UPLOAD\`, \`VIOL\_FILE\_UPLOAD\_IN\_BODY\`,
\`VIOL\_XML\_MALFORMED\`, \`VIOL\_JSON\_MALFORMED\`, \`VIOL\_ASM\_COOKIE\_MODIFIED\`,
\`VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE\`,
\`VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT\`, \`VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION\`,
\`VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS\`,
\`VIOL\_EVASION\_DIRECTORY\_TRAVERSALS\`, \`VIOL\_MALFORMED\_REQUEST\`,
\`VIOL\_EVASION\_MULTIPLE\_DECODING\`, \`VIOL\_DATA\_GUARD\`, \`VIOL\_EVASION\_APACHE\_WHITESPACE\`,
\`VIOL\_COOKIE\_MODIFIED\`, \`VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS\`,
\`VIOL\_EVASION\_IIS\_BACKSLASHES\`, \`VIOL\_EVASION\_PERCENT\_U\_DECODING\`,
\`VIOL\_EVASION\_BARE\_BYTE\_DECODING\`, \`VIOL\_EVASION\_BAD\_UNESCAPE\`,
\`VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST\`, \`VIOL\_ENCODING\`,
\`VIOL\_COOKIE\_MALFORMED\`, \`VIOL\_GRAPHQL\_FORMAT\`, \`VIOL\_GRAPHQL\_MALFORMED\`,
\`VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\`. Defaults to \`VIOL\_NONE\`.

Additional upstream details:

List of all supported Violation Types

VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER VIOL\_HTTP\_RESPONSE\_STATUS
VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED
VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS
VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT
VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION
VIOL\_HTTP\_PROTOCOL\_CRLF\_CHARACTERS\_BEFORE\_REQUEST\_START
VIOL\_HTTP\_PROTOCOL\_NO\_HOST\_HEADER\_IN\_HTTP\_1\_1\_REQUEST
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_PARAMETERS\_PARSING
VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS
VIOL\_HTTP\_PROTOCOL\_CONTENT\_LENGTH\_SHOULD\_BE\_A\_POSITIVE\_NUMBER
VIOL\_EVASION\_DIRECTORY\_TRAVERSALS VIOL\_MALFORMED\_REQUEST VIOL\_EVASION\_MULTIPLE\_DECODING
VIOL\_DATA\_GUARD VIOL\_EVASION\_APACHE\_WHITESPACE VIOL\_COOKIE\_MODIFIED
VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS VIOL\_EVASION\_IIS\_BACKSLASHES
VIOL\_EVASION\_PERCENT\_U\_DECODING VIOL\_EVASION\_BARE\_BYTE\_DECODING VIOL\_EVASION\_BAD\_UNESCAPE
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_FORMDATA\_REQUEST\_PARSING
VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST
VIOL\_HTTP\_PROTOCOL\_HIGH\_ASCII\_CHARACTERS\_IN\_HEADERS VIOL\_ENCODING VIOL\_COOKIE\_MALFORMED
VIOL\_GRAPHQL\_FORMAT VIOL\_GRAPHQL\_MALFORMED VIOL\_GRAPHQL\_INTROSPECTION\_QUERY.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIOL_NONE",
  "enum": [
    "VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START",
    "VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2100023330203021-3230010322023212-3220023033222203-2021031303313221-1312220002111123-2010323212111222-2110103221221223-0331313002320020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-027.md#canonical-2201211320210323-1131233103012122-1202220210012030-1312122002301333-3032311002210133-3132322100021302-0321312033221231-3303001323312332)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-0230300230321033-0121023132033121-2331203221203221-0213211112222231-0101232320213100-1002301230313112-1010311021120321-3003012233030302)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-1130202033312000-2333332311213331-2201123211031102-3000332303130330-0301021313211302-3302232003021022-0000301232023102-2130202013123322)
- waf_exclusion.waf_exclusion_inline_rules.rules.metadata

<a id="canonical-1302110021123112-0030102220331302-0321130311213001-0103020301010221-2103322331110221-1230330020303210-0202332003330003-3013230010100101"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-2221331131133011-2302201013312003-0130321320311322-3212100120130013-3212122303313223-3213301133100211-3213003303023303-2322200223220123"></a>

### Direct properties for `waf_exclusion.waf_exclusion_inline_rules.rules.metadata`

<a id="canonical-2312311020310032-3312031123210201-3113011223001320-1212332210302133-2023003010130123-1200331220111233-0202130302001120-3203000301303033"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3113003322222200-2210010030311101-3110222233113310-3333220201210030-1232103221000130-2330313102101031-1203332120100203-2000120200133001"></a>

<a id="canonical-2300022302330310-0323030132033122-2113131313333111-2121210211102333-3300113213231230-0013121112230223-1113001210133223-0211031110310322"></a>

#### `waf_exclusion.waf_exclusion_inline_rules.rules.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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

<a id="canonical-0220122011122223-2110203320003322-2100333232013333-3110031331021101-3302032121122003-2102223223111322-2120033031232200-1023133130102101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-027.md#canonical-2201211320210323-1131233103012122-1202220210012030-1312122002301333-3032311002210133-3132322100021302-0321312033221231-3303001323312332)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-0230300230321033-0121023132033121-2331203221203221-0213211112222231-0101232320213100-1002301230313112-1010311021120321-3003012233030302)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--http_loadbalancer--reference--group-027.md#canonical-1130202033312000-2333332311213331-2201123211031102-3000332303130330-0301021313211302-3302232003021022-0000301232023102-2130202013123322)
- waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing

<a id="canonical-3322220122032311-3221131312213322-1301023332120300-0113221322323023-3113131122113002-3020130111113030-2021233122331012-0330310120203233"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321011111113021-0321110131231130-3000230231112102-0230320013311302-1310012003013103-0230023330201311-1121112302211313-3103211000231200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_exclusion.waf_exclusion_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [waf_exclusion](data-sources--http_loadbalancer--reference--group-027.md#canonical-2201211320210323-1131233103012122-1202220210012030-1312122002301333-3032311002210133-3132322100021302-0321312033221231-3303001323312332)
- waf_exclusion.waf_exclusion_policy

<a id="canonical-2101013202311201-3002203001201000-3002003321333113-2133002312112011-1120130232032032-3213200312332103-1211201101321031-3021231011133120"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-3021100312300000-0331302130203332-3030303120220330-3303211032222030-3010133222301222-1020011010213132-2233032011103122-0213202231213001"></a>

### Direct properties for `waf_exclusion.waf_exclusion_policy`

<a id="canonical-1112131310112322-2030320020213303-3011021331203213-1131120103312200-3300003301132003-1003023313203321-1102230130330110-0213003232302000"></a>

#### `waf_exclusion.waf_exclusion_policy.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-0021202222110312-1311213122131221-1321331323303001-0032011122100303-0130002311102321-1102130021033202-2020232010103313-3130003320300131"></a>

<a id="canonical-0112000233302112-1232233123111021-3133132302131000-3231201020112220-1133232003233112-0023001323310103-2133121332012333-3101201023032333"></a>

#### `waf_exclusion.waf_exclusion_policy.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-0213100210102021-2203203202223000-1020033313103222-3313231221030031-1320012233330000-3120321202310313-1103133120100211-3221312103023021"></a>

<a id="canonical-0112123232001023-3323011032231131-0221133100121101-2113033332333203-0131330203211310-3132323102212213-3101030112200113-3323200023121322"></a>

#### `waf_exclusion.waf_exclusion_policy.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
