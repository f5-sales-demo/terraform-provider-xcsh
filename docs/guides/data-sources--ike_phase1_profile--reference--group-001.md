---
page_title: "xcsh_ike_phase1_profile reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase1_profile reference."
---

# xcsh_ike_phase1_profile reference

<a id="canonical-1220002200221322-3301212112331131-2221122301111113-1102133331000102-3210331112212011-2311300113222322-0000032230002222-1211320310333123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-1332210000333121-0321333020120302-3121233211013331-1320231330211033-2111031001002303-1300211301302100-1031311212302312-1213122302330023)
- Property reference

<a id="canonical-3021013230300330-0113031230021201-2003301112333232-1321002230123222-3110312102011211-3331223233011222-1013302102022230-2212123003323122"></a>

### Direct properties for `xcsh_ike_phase1_profile`

<a id="canonical-2021213032010312-3222021220210021-3201213130103123-3013203020133313-0201032103203033-2003323100033230-1101133031033223-0201210212331033"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

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

<a id="canonical-2212000103220000-0300233123203100-0011033112310333-3230011110331313-1323013130201112-1301013300201232-0102123331320332-1222331222103030"></a>

<a id="canonical-0301312303232200-1311203123131132-0300000320210211-0112111301033301-0030223013023233-3011032302122220-1122133013232133-0112213211112023"></a>

#### `authentication_algos` property

Type: `["list", "string"]`. Computed.

\[Enum: AUTH\_ALG\_DEFAULT|SHA256\_HMAC|SHA384\_HMAC|SHA512\_HMAC|AUTH\_ALG\_NONE\] Choose one or
more Authentication Algorithm. Use None option when using the aes-gcm or aes-ccm encryption
algorithms. Possible values are \`AUTH\_ALG\_DEFAULT\`, \`SHA256\_HMAC\`, \`SHA384\_HMAC\`,
\`SHA512\_HMAC\`, \`AUTH\_ALG\_NONE\`. Defaults to \`AUTH\_ALG\_DEFAULT\`.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-1021032321020330-1111002000210131-1200131010303212-3332122131101310-1301003123301200-1213212313011310-0203323120312130-1113022332321122"></a>

<a id="canonical-3323110003331123-0222120222231313-1221110320123222-3331303110130231-0103200211222122-2210022000212010-2312302220221133-1301230120133202"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the IKEPhase1Profile.

Additional upstream details:

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0101232332011203-2232101022220130-0311212320000132-1100322030103021-2332003310131021-1022210201013113-0321102010300121-1301220003010311"></a>

<a id="canonical-1233302321331222-2013111333102101-3201221131033231-0323123022001202-3021321202112230-0023001211330211-1321200201110001-1101013031022132"></a>

#### `dh_group` property

Type: `["list", "string"]`. Computed.

\[Enum:
DH\_GROUP\_DEFAULT|DH\_GROUP\_14|DH\_GROUP\_15|DH\_GROUP\_16|DH\_GROUP\_17|DH\_GROUP\_18|DH\_GROUP\_19|DH\_GROUP\_20|DH\_GROUP\_21|DH\_GROUP\_26\]
Choose the acceptable Diffie Hellman (DH) Group or Groups that you are willing to accept as part of
this profile. Possible values are \`DH\_GROUP\_DEFAULT\`, \`DH\_GROUP\_14\`, \`DH\_GROUP\_15\`,
\`DH\_GROUP\_16\`, \`DH\_GROUP\_17\`, \`DH\_GROUP\_18\`, \`DH\_GROUP\_19\`, \`DH\_GROUP\_20\`,
\`DH\_GROUP\_21\`, \`DH\_GROUP\_26\`. Defaults to \`DH\_GROUP\_DEFAULT\`.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-2000101032233001-1020013201003113-1111113021030213-2200222100221123-0113233220101313-0313322200220202-1121220120213031-2113211201000310"></a>

<a id="canonical-0131112023101130-3111121302010012-3131303021031221-0123000003020232-3110022200021320-0331232230210021-2110323123011033-1102203011121120"></a>

#### `encryption_algos` property

Type: `["list", "string"]`. Computed.

\[Enum:
ENC\_ALG\_DEFAULT|AES128\_CBC|AES192\_CBC|AES256\_CBC|TRIPLE\_DES\_CBC|AES128\_GCM|AES192\_GCM|AES256\_GCM\]
Choose one or more encryption algorithms. Possible values are \`ENC\_ALG\_DEFAULT\`,
\`AES128\_CBC\`, \`AES192\_CBC\`, \`AES256\_CBC\`, \`TRIPLE\_DES\_CBC\`, \`AES128\_GCM\`,
\`AES192\_GCM\`, \`AES256\_GCM\`. Defaults to \`ENC\_ALG\_DEFAULT\`.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-1031303233110232-0021200132020200-1130310221322323-1113223200032220-2010003301110233-3030301101013122-2331000021232000-3130130211101311"></a>

<a id="canonical-0131223013013032-2300202022201322-3130202000220011-1322123213301110-3121020221230132-3111232022012133-2220221310201323-1212021330103311"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](data-sources--ike_phase1_profile--reference--group-001.md#canonical-0123031303212103-3332123222221131-2132130102220331-1232032333220210-0103020112011230-2202330010130302-3233013302103330-2202121013020130): complete subsection reference.

- [ike_keylifetime_minutes](data-sources--ike_phase1_profile--reference--group-001.md#canonical-3320123112223131-1032323020301213-2103020130312112-2211232202122103-2212112020000310-1112031210013103-3113110102120010-2003220003331303): complete subsection reference.

<a id="canonical-0301223020032303-1330312011312001-0310323120021002-1203032103101121-0012103022012311-2230231201033101-1022312022123120-2123111122131323"></a>

<a id="canonical-3200022001211203-3201322003123331-3212230112322012-1002213020120211-2121200202023133-2012300302222321-0220313021220110-3210211202021200"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-1313333011133023-0111322203221220-3233220122022301-1113100330012113-2022332131302110-2302303031313312-0211321233323112-3331221122303010"></a>

<a id="canonical-0101123110230221-3033320110101312-2302211302001233-1332111112330011-2301013103312011-3032033030133101-2213011000101033-2132123021211233"></a>

#### `name` property

Type: `"string"`. Required.

Name of the IKEPhase1Profile.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1200312003103201-2233321222021313-1223232020123213-0010110001111301-3022233033002011-0122022012013011-2000232123220101-1232123303133310"></a>

<a id="canonical-0023023101232310-0000013021223312-0301030103123232-3220122120130030-3322012020122100-3021312000013103-0010220010111002-1332311133011231"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the IKEPhase1Profile exists.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0002000302311301-0013232201232130-0003301202313033-0311020131132200-3131232113132200-3200111031033003-0123002331123103-0010012210030200"></a>

<a id="canonical-2002332121133022-1220310100030220-2102121113223113-1120212120120211-3221203121232313-3311312131311100-3220003021030212-2200130030302030"></a>

#### `prf` property

Type: `["list", "string"]`. Computed.

\[Enum: PRF\_DEFAULT|PRFSHA256|PRFSHA384|PRFSHA512\] PseudoRandomFunction. Select
PseudoRandomFunction for IKE SA. Possible values are \`PRF\_DEFAULT\`, \`PRFSHA256\`, \`PRFSHA384\`,
\`PRFSHA512\`. Defaults to \`PRF\_DEFAULT\`.

Receipt-pinned upstream constraints:

```json
{
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

- [reauth_disabled](data-sources--ike_phase1_profile--reference--group-001.md#canonical-2021303222321000-1220123000002301-3212311120122010-0132010011203031-0123002110002320-3213302230100203-0332022003201103-0102110212202302): complete subsection reference.

- [reauth_timeout_days](data-sources--ike_phase1_profile--reference--group-001.md#canonical-3020123101002031-3121311001312002-3210031303110331-3200201030201011-2013303023220210-0212101201023111-3011313020000010-2200110322123000): complete subsection reference.

- [reauth_timeout_hours](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1230220120113100-1331301130230201-1210031102000020-2121110222312123-3010330220203323-0130231331102222-1121303030310122-1212023110220211): complete subsection reference.

- [use_default_keylifetime](data-sources--ike_phase1_profile--reference--group-001.md#canonical-3210132313302233-0211212302003322-0033112003203122-3300302012010012-0303221330212301-1232212200130030-3133011320001032-0213001302222012): complete subsection reference.

<a id="canonical-0200020102102232-1331330312231132-0203303012311020-0200211112212200-0121332330023102-0103011032322320-3310032121101220-1212222311011030"></a>

### All schema paths for `xcsh_ike_phase1_profile`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--ike_phase1_profile--reference--group-001.md#canonical-2021213032010312-3222021220210021-3201213130103123-3013203020133313-0201032103203033-2003323100033230-1101133031033223-0201210212331033) |
| `authentication_algos` | [authentication_algos](data-sources--ike_phase1_profile--reference--group-001.md#canonical-2212000103220000-0300233123203100-0011033112310333-3230011110331313-1323013130201112-1301013300201232-0102123331320332-1222331222103030) |
| `description` | [description](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1021032321020330-1111002000210131-1200131010303212-3332122131101310-1301003123301200-1213212313011310-0203323120312130-1113022332321122) |
| `dh_group` | [dh_group](data-sources--ike_phase1_profile--reference--group-001.md#canonical-0101232332011203-2232101022220130-0311212320000132-1100322030103021-2332003310131021-1022210201013113-0321102010300121-1301220003010311) |
| `encryption_algos` | [encryption_algos](data-sources--ike_phase1_profile--reference--group-001.md#canonical-2000101032233001-1020013201003113-1111113021030213-2200222100221123-0113233220101313-0313322200220202-1121220120213031-2113211201000310) |
| `id` | [ID](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1031303233110232-0021200132020200-1130310221322323-1113223200032220-2010003301110233-3030301101013122-2331000021232000-3130130211101311) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1302003030300022-3213123222230321-0233222121331312-0211223133002130-2213331312030012-2220031311310322-1300133333021321-2131231110303112) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](data-sources--ike_phase1_profile--reference--group-001.md#canonical-3001021232001333-3003313331330010-1010313320312223-1032332230331222-0111222123203132-2312121032213321-0110212011213312-1133313211022032) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1120230113201211-1122102111311031-2002200221020103-0203310001333322-2120023210213011-3212301023312302-2312013010010131-0031301031211133) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](data-sources--ike_phase1_profile--reference--group-001.md#canonical-0210003020021022-0301330030310201-0020202230231033-0130122120110311-1130323022230233-0100013322321200-3001321333222132-3130123311033132) |
| `labels` | [labels](data-sources--ike_phase1_profile--reference--group-001.md#canonical-0301223020032303-1330312011312001-0310323120021002-1203032103101121-0012103022012311-2230231201033101-1022312022123120-2123111122131323) |
| `name` | [name](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1313333011133023-0111322203221220-3233220122022301-1113100330012113-2022332131302110-2302303031313312-0211321233323112-3331221122303010) |
| `namespace` | [namespace](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1200312003103201-2233321222021313-1223232020123213-0010110001111301-3022233033002011-0122022012013011-2000232123220101-1232123303133310) |
| `prf` | [prf](data-sources--ike_phase1_profile--reference--group-001.md#canonical-0002000302311301-0013232201232130-0003301202313033-0311020131132200-3131232113132200-3200111031033003-0123002331123103-0010012210030200) |
| `reauth_disabled` | [reauth_disabled](data-sources--ike_phase1_profile--reference--group-001.md#canonical-0331212221312000-1013121021131231-2113013330103212-0112110231131132-1232201110001012-2332131011333113-1010001233102121-2021130202001122) |
| `reauth_timeout_days` | [reauth_timeout_days](data-sources--ike_phase1_profile--reference--group-001.md#canonical-3212222233133313-2111312211113313-3200113123312321-2220132123001111-1013122010132233-2223130103132202-1230300303000032-1223033133200221) |
| `reauth_timeout_days.duration` | [reauth_timeout_days.duration](data-sources--ike_phase1_profile--reference--group-001.md#canonical-3031222303020033-2322203013300322-3123012102211231-0020120110001013-1111320211102011-0210113111010011-3001103230332020-2130212300001031) |
| `reauth_timeout_hours` | [reauth_timeout_hours](data-sources--ike_phase1_profile--reference--group-001.md#canonical-2213003002033111-3200230020022121-1120303022320321-0323012032133321-0303333030131221-2132031003220223-0230033321031331-2220220200132003) |
| `reauth_timeout_hours.duration` | [reauth_timeout_hours.duration](data-sources--ike_phase1_profile--reference--group-001.md#canonical-3220123332113231-2212223301211232-2100030023122231-1113231002232110-1310102233312323-0102332310333110-2012110012202003-0101233312301200) |
| `use_default_keylifetime` | [use_default_keylifetime](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1203021302201132-1233203130312313-3011011331122032-3121230122033122-0010122201012120-3010122011323312-2121021313002233-2021023000030031) |

<a id="canonical-0123031303212103-3332123222221131-2132130102220331-1232032333220210-0103020112011230-2202330010130302-3233013302103330-2202121013020130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ike_keylifetime_hours` properties

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-1332210000333121-0321333020120302-3121233211013331-1320231330211033-2111031001002303-1300211301302100-1031311212302312-1213122302330023)
- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1220002200221322-3301212112331131-2221122301111113-1102133331000102-3210331112212011-2311300113222322-0000032230002222-1211320310333123)
- ike_keylifetime_hours

<a id="canonical-1302003030300022-3213123222230321-0233222121331312-0211223133002130-2213331312030012-2220031311310322-1300133333021321-2131231110303112"></a>

Type: `"single"`. Computed.

\[OneOf: ike\_keylifetime\_hours, ike\_keylifetime\_minutes, use\_default\_keylifetime; Default:
use\_default\_keylifetime\] Configuration parameter for ike keylifetime hours.

Additional upstream details:

Input Hours.

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

- [ike_keylifetime_hours](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1302003030300022-3213123222230321-0233222121331312-0211223133002130-2213331312030012-2220031311310322-1300133333021321-2131231110303112)
- [ike_keylifetime_minutes](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1120230113201211-1122102111311031-2002200221020103-0203310001333322-2120023210213011-3212301023312302-2312013010010131-0031301031211133)
- [use_default_keylifetime](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1203021302201132-1233203130312313-3011011331122032-3121230122033122-0010122201012120-3010122011323312-2121021313002233-2021023000030031)

Select alternatives according to the provider validators above.

<a id="canonical-0303010311031101-3131232010020233-2202220213210311-1220010102000113-1332131000000013-3032300023030321-1211310212011133-0121030311113111"></a>

### Direct properties for `ike_keylifetime_hours`

<a id="canonical-3001021232001333-3003313331330010-1010313320312223-1032332230331222-0111222123203132-2312121032213321-0110212011213312-1133313211022032"></a>

#### `ike_keylifetime_hours.duration` property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-3320123112223131-1032323020301213-2103020130312112-2211232202122103-2212112020000310-1112031210013103-3113110102120010-2003220003331303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ike_keylifetime_minutes` properties

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-1332210000333121-0321333020120302-3121233211013331-1320231330211033-2111031001002303-1300211301302100-1031311212302312-1213122302330023)
- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1220002200221322-3301212112331131-2221122301111113-1102133331000102-3210331112212011-2311300113222322-0000032230002222-1211320310333123)
- ike_keylifetime_minutes

<a id="canonical-1120230113201211-1122102111311031-2002200221020103-0203310001333322-2120023210213011-3212301023312302-2312013010010131-0031301031211133"></a>

Type: `"single"`. Computed.

Configuration parameter for ike keylifetime minutes.

Additional upstream details:

Set IKE Key Lifetime in minutes.

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

<a id="canonical-3012111231323311-0131121222333131-2030011211300133-2322002212223232-1221322112220230-1121322101231231-0013002201002330-0322023020001322"></a>

### Direct properties for `ike_keylifetime_minutes`

<a id="canonical-0210003020021022-0301330030310201-0020202230231033-0130122120110311-1130323022230233-0100013322321200-3001321333222132-3130123311033132"></a>

#### `ike_keylifetime_minutes.duration` property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 10
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-2021303222321000-1220123000002301-3212311120122010-0132010011203031-0123002110002320-3213302230100203-0332022003201103-0102110212202302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `reauth_disabled` properties

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-1332210000333121-0321333020120302-3121233211013331-1320231330211033-2111031001002303-1300211301302100-1031311212302312-1213122302330023)
- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1220002200221322-3301212112331131-2221122301111113-1102133331000102-3210331112212011-2311300113222322-0000032230002222-1211320310333123)
- reauth_disabled

<a id="canonical-0331212221312000-1013121021131231-2113013330103212-0112110231131132-1232201110001012-2332131011333113-1010001233102121-2021130202001122"></a>

Type: `["object", {}]`. Computed.

\[OneOf: reauth\_disabled, reauth\_timeout\_days, reauth\_timeout\_hours\] Enable this option

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

- [reauth_disabled](data-sources--ike_phase1_profile--reference--group-001.md#canonical-0331212221312000-1013121021131231-2113013330103212-0112110231131132-1232201110001012-2332131011333113-1010001233102121-2021130202001122)
- [reauth_timeout_days](data-sources--ike_phase1_profile--reference--group-001.md#canonical-3212222233133313-2111312211113313-3200113123312321-2220132123001111-1013122010132233-2223130103132202-1230300303000032-1223033133200221)
- [reauth_timeout_hours](data-sources--ike_phase1_profile--reference--group-001.md#canonical-2213003002033111-3200230020022121-1120303022320321-0323012032133321-0303333030131221-2132031003220223-0230033321031331-2220220200132003)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020123101002031-3121311001312002-3210031303110331-3200201030201011-2013303023220210-0212101201023111-3011313020000010-2200110322123000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `reauth_timeout_days` properties

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-1332210000333121-0321333020120302-3121233211013331-1320231330211033-2111031001002303-1300211301302100-1031311212302312-1213122302330023)
- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1220002200221322-3301212112331131-2221122301111113-1102133331000102-3210331112212011-2311300113222322-0000032230002222-1211320310333123)
- reauth_timeout_days

<a id="canonical-3212222233133313-2111312211113313-3200113123312321-2220132123001111-1013122010132233-2223130103132202-1230300303000032-1223033133200221"></a>

Type: `"single"`. Computed.

Configuration parameter for reauth timeout days.

Additional upstream details:

Set Duration in days.

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

<a id="canonical-2101302300102123-2101101031332002-3322102111202011-0201010112121310-1003002013303003-0322023133230303-3221001213022022-0212301030202132"></a>

### Direct properties for `reauth_timeout_days`

<a id="canonical-3031222303020033-2322203013300322-3123012102211231-0020120110001013-1111320211102011-0210113111010011-3001103230332020-2130212300001031"></a>

#### `reauth_timeout_days.duration` property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-1230220120113100-1331301130230201-1210031102000020-2121110222312123-3010330220203323-0130231331102222-1121303030310122-1212023110220211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `reauth_timeout_hours` properties

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-1332210000333121-0321333020120302-3121233211013331-1320231330211033-2111031001002303-1300211301302100-1031311212302312-1213122302330023)
- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1220002200221322-3301212112331131-2221122301111113-1102133331000102-3210331112212011-2311300113222322-0000032230002222-1211320310333123)
- reauth_timeout_hours

<a id="canonical-2213003002033111-3200230020022121-1120303022320321-0323012032133321-0303333030131221-2132031003220223-0230033321031331-2220220200132003"></a>

Type: `"single"`. Computed.

Configuration parameter for reauth timeout hours.

Additional upstream details:

Input Hours.

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

<a id="canonical-0131000112303003-0210012223013133-2331122003321023-0320121133211311-0321302103222003-2013230300232012-0211023232313023-1121301332320303"></a>

### Direct properties for `reauth_timeout_hours`

<a id="canonical-3220123332113231-2212223301211232-2100030023122231-1113231002232110-1310102233312323-0102332310333110-2012110012202003-0101233312301200"></a>

#### `reauth_timeout_hours.duration` property

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-3210132313302233-0211212302003322-0033112003203122-3300302012010012-0303221330212301-1232212200130030-3133011320001032-0213001302222012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_default_keylifetime` properties

Breadcrumbs:

- [xcsh_ike_phase1_profile](../data-sources/ike_phase1_profile.md#canonical-1332210000333121-0321333020120302-3121233211013331-1320231330211033-2111031001002303-1300211301302100-1031311212302312-1213122302330023)
- [Property reference](data-sources--ike_phase1_profile--reference--group-001.md#canonical-1220002200221322-3301212112331131-2221122301111113-1102133331000102-3210331112212011-2311300113222322-0000032230002222-1211320310333123)
- use_default_keylifetime

<a id="canonical-1203021302201132-1233203130312313-3011011331122032-3121230122033122-0010122201012120-3010122011323312-2121021313002233-2021023000030031"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use default keylifetime.

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
