---
page_title: "xcsh_ike_phase1_profile reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike_phase1_profile reference."
---

# xcsh_ike_phase1_profile reference

<a id="canonical-3320222211011213-1313031020110201-1001212000222033-3130231131000323-0312231100011132-2003002022233112-3032113203121302-3122030323211321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333000221230032-3020202102133332-3213013230202330-0211110110030200-3201212000201201-3113012311321123-0203013012211133-2331202000123220"></a>

## Property reference — Property reference / 223103023012 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220)
- Property reference

<a id="canonical-1013122012202211-3223311210120230-0012120000103130-0130001121112103-0131203033312030-3111311221101313-1221320201020100-1200200101220211"></a>

## Direct properties — Property reference / 223103023012 / 3

<a id="canonical-0122232131030113-0002122101200111-3213000312300313-3212122233222302-0122322003112001-2010301120120021-3312100310012322-1212312230222323"></a>

<a id="canonical-1200013231301030-1232211011303131-2302222313030012-1302300212031003-2021132203022120-2200213302322212-3123131230131011-3310222320232333"></a>

## annotations property — Property reference / 223103023012 / 4

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

<a id="canonical-3132020012213011-0230120103032003-2221230321331321-3231300010222122-0013002321221332-3130022301103020-0123203012033231-1031213001302220"></a>

<a id="canonical-3011021000122211-3330012001323102-0322120223031000-1123222010102321-3013100200221210-0111020321332013-1013123130201030-3120333310213031"></a>

## authentication_algos property — Property reference / 223103023012 / 5

Type: `["list", "string"]`. Required.

\[Enum: AUTH\_ALG\_DEFAULT|SHA256\_HMAC|SHA384\_HMAC|SHA512\_HMAC|AUTH\_ALG\_NONE\] Choose one or
more Authentication Algorithm. Use None option when using the aes-gcm or aes-ccm encryption
algorithms. Possible values are \`AUTH\_ALG\_DEFAULT\`, \`SHA256\_HMAC\`, \`SHA384\_HMAC\`,
\`SHA512\_HMAC\`, \`AUTH\_ALG\_NONE\`. Defaults to \`AUTH\_ALG\_DEFAULT\`.

Upstream description:

Choose one or more Authentication Algorithm. Use None option when using the aes-gcm or aes-ccm
encryption algorithms.

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

<a id="canonical-2132003120111312-1200210222330032-3021332030332002-0110020120123102-1112301302312312-0223301330211232-2100013331311223-1000310220111232"></a>

<a id="canonical-3303222203200020-2311111032321131-3332302023111221-2031013310130031-1003022013033331-3202200113210030-3330322012320111-0202213130212032"></a>

## description property — Property reference / 223103023012 / 6

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

<a id="canonical-3321300303221212-2210221211213221-1003010311122313-3002213330022120-2101031011021000-2223203311012121-0313223211030220-3120002001201100"></a>

<a id="canonical-2221201212133333-3302013000211200-2211001102321230-3210200123033313-0230332231110311-0311300221300301-3320201013001030-3210132332132323"></a>

## dh_group property — Property reference / 223103023012 / 7

Type: `["list", "string"]`. Required.

\[Enum:
DH\_GROUP\_DEFAULT|DH\_GROUP\_14|DH\_GROUP\_15|DH\_GROUP\_16|DH\_GROUP\_17|DH\_GROUP\_18|DH\_GROUP\_19|DH\_GROUP\_20|DH\_GROUP\_21|DH\_GROUP\_26\]
Choose the acceptable Diffie Hellman (DH) Group or Groups that you are willing to accept as part of
this profile. Possible values are \`DH\_GROUP\_DEFAULT\`, \`DH\_GROUP\_14\`, \`DH\_GROUP\_15\`,
\`DH\_GROUP\_16\`, \`DH\_GROUP\_17\`, \`DH\_GROUP\_18\`, \`DH\_GROUP\_19\`, \`DH\_GROUP\_20\`,
\`DH\_GROUP\_21\`, \`DH\_GROUP\_26\`. Defaults to \`DH\_GROUP\_DEFAULT\`.

Upstream description:

Choose the acceptable Diffie Hellman (DH) Group or Groups that you are willing to accept as part of
this profile.

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

<a id="canonical-0321333132323333-1131131232120301-3021300333201103-1202212123320120-1213232122101022-3320303033120223-2310221002102032-3131013003203110"></a>

<a id="canonical-3310033232021113-1121110113200320-0020310211132101-2313033031122001-1323121311222021-3101101103212310-3020232233321223-0231302002200230"></a>

## disable property — Property reference / 223103023012 / 8

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

<a id="canonical-3211223230210201-0000313011132002-2113100300123133-2131112312031132-0210312031111102-2333332012221011-3033000013331132-2002312322332223"></a>

<a id="canonical-1212121120310003-0122231232313220-3210322032210310-3331333202332033-1230321331103110-2133233022320212-2333303100312330-2123210201300331"></a>

## encryption_algos property — Property reference / 223103023012 / 9

Type: `["list", "string"]`. Required.

\[Enum:
ENC\_ALG\_DEFAULT|AES128\_CBC|AES192\_CBC|AES256\_CBC|TRIPLE\_DES\_CBC|AES128\_GCM|AES192\_GCM|AES256\_GCM\]
Choose one or more encryption algorithms. Possible values are \`ENC\_ALG\_DEFAULT\`,
\`AES128\_CBC\`, \`AES192\_CBC\`, \`AES256\_CBC\`, \`TRIPLE\_DES\_CBC\`, \`AES128\_GCM\`,
\`AES192\_GCM\`, \`AES256\_GCM\`. Defaults to \`ENC\_ALG\_DEFAULT\`.

Upstream description:

Choose one or more encryption algorithms.

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

<a id="canonical-0322203311113312-2302312232121101-1003032101323003-1213101110132020-2010321113030300-1320210013111321-2323130012302231-2303111310210321"></a>

<a id="canonical-3022000201120233-2333111032031200-3120110131022130-3200033312002202-1033232100220103-3331313022231003-3133123013133333-3132200013300222"></a>

## ID property — Property reference / 223103023012 / 10

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ike_keylifetime_hours](resources--ike_phase1_profile--reference--group-001.md#canonical-0202020210020100-0010323001013223-3102230231101231-0332020330300301-3131031101323222-2122211233201101-2032303103300212-2333331323213101): complete subsection reference.

- [ike_keylifetime_minutes](resources--ike_phase1_profile--reference--group-001.md#canonical-1233332313021110-1030011230320011-0013110302032120-0111103012113333-0302233032003013-2121303331212323-0211011132210010-2131100332012121): complete subsection reference.

<a id="canonical-0212312121031103-3101231133332032-2203333222320000-0320131302003020-1111330010220330-1201300203301021-2101302211222332-0223103201033030"></a>

<a id="canonical-3003120001021003-3200303111230122-0230220120012110-0012200221322331-1103032333333220-2130303231030131-1222232302102210-2102322310232030"></a>

## labels property — Property reference / 223103023012 / 11

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

<a id="canonical-1031210332311210-0031110333303013-3001121011100302-1132300311200103-2032332201220032-3230201202231212-0120322203311022-1230021310200033"></a>

<a id="canonical-1123321303102212-2221222303333222-2022213132210100-3321133131233120-0203131220223201-2333101103200120-0001323333000132-3300031130211303"></a>

## name property — Property reference / 223103023012 / 12

Type: `"string"`. Required.

Name of the IKE Phase1 Profile. Must be unique within the namespace.

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

<a id="canonical-3012211203323003-0312032302123320-3302013333002202-2331113111110301-0021201311112113-0132221103013201-3102113003332010-1323002102331311"></a>

<a id="canonical-1200121122300003-3210200221000102-0331101031133133-1320213023123133-0233033112231310-0221030021123331-3112120030311113-2032222212112022"></a>

## namespace property — Property reference / 223103023012 / 13

Type: `"string"`. Required.

Namespace where the IKE Phase1 Profile is created.

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

<a id="canonical-0211003112023203-1300221332102300-3320012221132333-0222020321111113-3333022022212113-2102300031110111-1330022323320010-1010133102032302"></a>

<a id="canonical-1133100222101330-2201330211320033-3021203113231012-1130002122033022-0220021013123012-0232032000013031-2210100221021012-0013310001102112"></a>

## prf property — Property reference / 223103023012 / 14

Type: `["list", "string"]`. Required.

\[Enum: PRF\_DEFAULT|PRFSHA256|PRFSHA384|PRFSHA512\] PseudoRandomFunction. Select
PseudoRandomFunction for IKE SA. Possible values are \`PRF\_DEFAULT\`, \`PRFSHA256\`, \`PRFSHA384\`,
\`PRFSHA512\`. Defaults to \`PRF\_DEFAULT\`.

Upstream description:

Select PseudoRandomFunction for IKE SA.

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

- [reauth_disabled](resources--ike_phase1_profile--reference--group-001.md#canonical-0023133132012211-2211123022033213-3230133023222202-0123231122020310-3020222100002002-1300333302303323-2022131100112010-2222302020100330): complete subsection reference.

- [reauth_timeout_days](resources--ike_phase1_profile--reference--group-001.md#canonical-3022232010212211-3311312110222112-2212001323132301-3322123113101312-3021320311332100-3233231203303320-2011010223222322-1200322130012222): complete subsection reference.

- [reauth_timeout_hours](resources--ike_phase1_profile--reference--group-001.md#canonical-2320021330102133-0033231100001111-0111003111200113-0122201001320210-1130201031310221-1203230001112102-0211030232222222-1230133333012201): complete subsection reference.

- [timeouts](resources--ike_phase1_profile--reference--group-001.md#canonical-0102202232133320-3203011001112213-3121102113021033-2120301213211033-0312121132102320-2013020020320110-0110320311230222-2331122103023010): complete subsection reference.

- [use_default_keylifetime](resources--ike_phase1_profile--reference--group-001.md#canonical-3032311121311133-3013223020103033-3312300313101003-0321121110012103-1223120221131333-1100223111301200-1322032200210113-3333013330021011): complete subsection reference.

<a id="canonical-1021111232201122-3130311101213220-3230133123333101-1020033103210022-1121301331212230-0003030013100131-2121301233221112-1121030003330301"></a>

## All schema paths — Property reference / 223103023012 / 15

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--ike_phase1_profile--reference--group-001.md#canonical-0122232131030113-0002122101200111-3213000312300313-3212122233222302-0122322003112001-2010301120120021-3312100310012322-1212312230222323) |
| `authentication_algos` | [authentication_algos](resources--ike_phase1_profile--reference--group-001.md#canonical-3132020012213011-0230120103032003-2221230321331321-3231300010222122-0013002321221332-3130022301103020-0123203012033231-1031213001302220) |
| `description` | [description](resources--ike_phase1_profile--reference--group-001.md#canonical-2132003120111312-1200210222330032-3021332030332002-0110020120123102-1112301302312312-0223301330211232-2100013331311223-1000310220111232) |
| `dh_group` | [dh_group](resources--ike_phase1_profile--reference--group-001.md#canonical-3321300303221212-2210221211213221-1003010311122313-3002213330022120-2101031011021000-2223203311012121-0313223211030220-3120002001201100) |
| `disable` | [disable](resources--ike_phase1_profile--reference--group-001.md#canonical-0321333132323333-1131131232120301-3021300333201103-1202212123320120-1213232122101022-3320303033120223-2310221002102032-3131013003203110) |
| `encryption_algos` | [encryption_algos](resources--ike_phase1_profile--reference--group-001.md#canonical-3211223230210201-0000313011132002-2113100300123133-2131112312031132-0210312031111102-2333332012221011-3033000013331132-2002312322332223) |
| `id` | [ID](resources--ike_phase1_profile--reference--group-001.md#canonical-0322203311113312-2302312232121101-1003032101323003-1213101110132020-2010321113030300-1320210013111321-2323130012302231-2303111310210321) |
| `ike_keylifetime_hours` | [ike_keylifetime_hours](resources--ike_phase1_profile--reference--group-001.md#canonical-3323122320311203-1322200313223121-1012000022202132-3212023313221132-2302210312132223-1221210103030120-2033030303021000-2012102112021031) |
| `ike_keylifetime_hours.duration` | [ike_keylifetime_hours.duration](resources--ike_phase1_profile--reference--group-001.md#canonical-2002031220013332-0020023001111112-2033301112331113-0303203322020032-2122233013302232-1300203101232201-2221033031103212-3120030310012300) |
| `ike_keylifetime_minutes` | [ike_keylifetime_minutes](resources--ike_phase1_profile--reference--group-001.md#canonical-2130201001133332-3220233230021220-3012221233113333-3120132003000230-3230130020220131-0311212112112011-2320010202111003-2231230230120020) |
| `ike_keylifetime_minutes.duration` | [ike_keylifetime_minutes.duration](resources--ike_phase1_profile--reference--group-001.md#canonical-0113001200120011-2013231300002133-2221321301021222-3000123331331113-0122221103021233-1112131211100212-3202032301200032-3020012123213131) |
| `labels` | [labels](resources--ike_phase1_profile--reference--group-001.md#canonical-0212312121031103-3101231133332032-2203333222320000-0320131302003020-1111330010220330-1201300203301021-2101302211222332-0223103201033030) |
| `name` | [name](resources--ike_phase1_profile--reference--group-001.md#canonical-1031210332311210-0031110333303013-3001121011100302-1132300311200103-2032332201220032-3230201202231212-0120322203311022-1230021310200033) |
| `namespace` | [namespace](resources--ike_phase1_profile--reference--group-001.md#canonical-3012211203323003-0312032302123320-3302013333002202-2331113111110301-0021201311112113-0132221103013201-3102113003332010-1323002102331311) |
| `prf` | [prf](resources--ike_phase1_profile--reference--group-001.md#canonical-0211003112023203-1300221332102300-3320012221132333-0222020321111113-3333022022212113-2102300031110111-1330022323320010-1010133102032302) |
| `reauth_disabled` | [reauth_disabled](resources--ike_phase1_profile--reference--group-001.md#canonical-0232123231303000-1010110202312000-3010332303002023-2100332210033202-1102323103223213-3013303331332210-1130203300311122-0121112330310231) |
| `reauth_timeout_days` | [reauth_timeout_days](resources--ike_phase1_profile--reference--group-001.md#canonical-0311110201230321-3103202032001301-1000102032200012-3222303110233303-2011313011330121-3101312033321223-2203123222020230-2210011303123210) |
| `reauth_timeout_days.duration` | [reauth_timeout_days.duration](resources--ike_phase1_profile--reference--group-001.md#canonical-2323111102122310-2311103321201222-0020013200030122-3102122230201131-0312121110213320-1013022120212031-0100333331301003-3021200113210022) |
| `reauth_timeout_hours` | [reauth_timeout_hours](resources--ike_phase1_profile--reference--group-001.md#canonical-3220112000103122-1132010010111211-3020321110313102-3102120210010203-1023300210332101-1000212201211230-0132303000312010-1033113020132320) |
| `reauth_timeout_hours.duration` | [reauth_timeout_hours.duration](resources--ike_phase1_profile--reference--group-001.md#canonical-1110122321110030-2022211332002311-2233131212201313-0200121130231333-0203033222121000-0033130110312131-0203022330112103-0032000213110132) |
| `timeouts` | [timeouts](resources--ike_phase1_profile--reference--group-001.md#canonical-3301102113321031-2330303210300023-3122130330312030-2123123202323003-2032233101031220-1020002021212321-0331320213303332-2121130321332101) |
| `timeouts.create` | [timeouts.create](resources--ike_phase1_profile--reference--group-001.md#canonical-0113321303310120-0311331130101002-3221120221201300-2012221120032202-2221211023303211-0302320222010211-0221031310300221-1230233013333130) |
| `timeouts.delete` | [timeouts.delete](resources--ike_phase1_profile--reference--group-001.md#canonical-0000020030331203-2113230012303232-2113133231230210-3321212130131210-2101130202331023-2200301013120010-3213120232012131-1213302132210012) |
| `timeouts.read` | [timeouts.read](resources--ike_phase1_profile--reference--group-001.md#canonical-0033123013312202-0213123303203201-1333111021121213-3223333200100301-1003012023321100-0312133000113100-2030321023113122-2101313313333032) |
| `timeouts.update` | [timeouts.update](resources--ike_phase1_profile--reference--group-001.md#canonical-2001100320323320-2000011001000202-2310212321221000-2110321203111301-0331311331331333-2101111030320131-2012332132031113-3130121111001223) |
| `use_default_keylifetime` | [use_default_keylifetime](resources--ike_phase1_profile--reference--group-001.md#canonical-2331331133331311-2202012033123012-0010121012101213-2202102322233022-3110023000103031-0122220103102013-2332302031102333-1002312223032113) |

<a id="canonical-3111222033010031-1132203030030233-1320233022203031-0310100212033033-1330222003313203-1322233020013332-1300121103213333-3310022211202123"></a>

## Next pages — Property reference / 223103023012 / 16

- [ike_keylifetime_hours](resources--ike_phase1_profile--reference--group-001.md#canonical-0202020210020100-0010323001013223-3102230231101231-0332020330300301-3131031101323222-2122211233201101-2032303103300212-2333331323213101)
- [ike_keylifetime_minutes](resources--ike_phase1_profile--reference--group-001.md#canonical-1233332313021110-1030011230320011-0013110302032120-0111103012113333-0302233032003013-2121303331212323-0211011132210010-2131100332012121)
- [reauth_disabled](resources--ike_phase1_profile--reference--group-001.md#canonical-0023133132012211-2211123022033213-3230133023222202-0123231122020310-3020222100002002-1300333302303323-2022131100112010-2222302020100330)
- [reauth_timeout_days](resources--ike_phase1_profile--reference--group-001.md#canonical-3022232010212211-3311312110222112-2212001323132301-3322123113101312-3021320311332100-3233231203303320-2011010223222322-1200322130012222)
- [reauth_timeout_hours](resources--ike_phase1_profile--reference--group-001.md#canonical-2320021330102133-0033231100001111-0111003111200113-0122201001320210-1130201031310221-1203230001112102-0211030232222222-1230133333012201)
- [timeouts](resources--ike_phase1_profile--reference--group-001.md#canonical-0102202232133320-3203011001112213-3121102113021033-2120301213211033-0312121132102320-2013020020320110-0110320311230222-2331122103023010)
- [use_default_keylifetime](resources--ike_phase1_profile--reference--group-001.md#canonical-3032311121311133-3013223020103033-3312300313101003-0321121110012103-1223120221131333-1100223111301200-1322032200210113-3333013330021011)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220)

<a id="canonical-0202020210020100-0010323001013223-3102230231101231-0332020330300301-3131031101323222-2122211233201101-2032303103300212-2333331323213101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013303112033222-3331100323013011-0033003033113121-0112213220320321-1003322033232000-1002330002223013-3312131321012121-0222102012220200"></a>

## ike_keylifetime_hours — ike_keylifetime_hours / 202232230222 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220)
- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-3320222211011213-1313031020110201-1001212000222033-3130231131000323-0312231100011132-2003002022233112-3032113203121302-3122030323211321)
- ike_keylifetime_hours

<a id="canonical-3323122320311203-1322200313223121-1012000022202132-3212023313221132-2302210312132223-1221210103030120-2033030303021000-2012102112021031"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ike\_keylifetime\_hours, ike\_keylifetime\_minutes, use\_default\_keylifetime; Default:
use\_default\_keylifetime\] Configuration parameter for ike keylifetime hours.

Upstream description:

Input Hours.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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

OneOf alternatives in this subsection:

- [ike_keylifetime_hours](resources--ike_phase1_profile--reference--group-001.md#canonical-3323122320311203-1322200313223121-1012000022202132-3212023313221132-2302210312132223-1221210103030120-2033030303021000-2012102112021031)
- [ike_keylifetime_minutes](resources--ike_phase1_profile--reference--group-001.md#canonical-2130201001133332-3220233230021220-3012221233113333-3120132003000230-3230130020220131-0311212112112011-2320010202111003-2231230230120020)
- [use_default_keylifetime](resources--ike_phase1_profile--reference--group-001.md#canonical-2331331133331311-2202012033123012-0010121012101213-2202102322233022-3110023000103031-0122220103102013-2332302031102333-1002312223032113)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ike_keylifetime_hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300002030121000-2120311322301022-1333031003211000-2111301330113213-0311213003331030-2301203020301333-0112300120121030-0021021133332313"></a>

## Direct properties — ike_keylifetime_hours / 202232230222 / 3

<a id="canonical-2002031220013332-0020023001111112-2033301112331113-0303203322020032-2122233013302232-1300203101232201-2221033031103212-3120030310012300"></a>

<a id="canonical-0300320001333120-2023101023320332-1030022310021311-0120213223101100-3203013330031222-3223203023333002-3232033200123303-1201102332131132"></a>

## duration property — ike_keylifetime_hours / 202232230222 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3132331102133211-2122222330222323-2321011213012000-2203002033003232-0130132203033211-1103232202331133-2122103131031311-1322110220213301"></a>

## Next pages — ike_keylifetime_hours / 202232230222 / 5

- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-3320222211011213-1313031020110201-1001212000222033-3130231131000323-0312231100011132-2003002022233112-3032113203121302-3122030323211321)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220)

<a id="canonical-1233332313021110-1030011230320011-0013110302032120-0111103012113333-0302233032003013-2121303331212323-0211011132210010-2131100332012121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100332211122002-0030322202131010-3231331232001031-3123201022312001-2221310220213113-1132311213033001-3030212220230130-0131221300211112"></a>

## ike_keylifetime_minutes — ike_keylifetime_minutes / 032131222013 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220)
- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-3320222211011213-1313031020110201-1001212000222033-3130231131000323-0312231100011132-2003002022233112-3032113203121302-3122030323211321)
- ike_keylifetime_minutes

<a id="canonical-2130201001133332-3220233230021220-3012221233113333-3120132003000230-3230130020220131-0311212112112011-2320010202111003-2231230230120020"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ike keylifetime minutes.

Upstream description:

Set IKE Key Lifetime in minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
ike_keylifetime_minutes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302112100023221-0231231302023101-0321232003300202-0103210121323302-1033013020121132-0233310012303032-0121023222322230-1031323031000020"></a>

## Direct properties — ike_keylifetime_minutes / 032131222013 / 3

<a id="canonical-0113001200120011-2013231300002133-2221321301021222-3000123331331113-0122221103021233-1112131211100212-3202032301200032-3020012123213131"></a>

<a id="canonical-0213313232032300-2011103303331113-0202303332333312-2002232013231113-0232322203202313-3301131102012011-1311300210200131-1010211333130131"></a>

## duration property — ike_keylifetime_minutes / 032131222013 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(10, 300),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0230313311233022-3001133202101333-2111212110312030-2311302201212331-3121023133032301-3021230003300122-2131210023100222-3323013233323211"></a>

## Next pages — ike_keylifetime_minutes / 032131222013 / 5

- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-3320222211011213-1313031020110201-1001212000222033-3130231131000323-0312231100011132-2003002022233112-3032113203121302-3122030323211321)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220)

<a id="canonical-0023133132012211-2211123022033213-3230133023222202-0123231122020310-3020222100002002-1300333302303323-2022131100112010-2222302020100330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322321111210031-3313100213111211-1011312100232333-3231032033101112-0012201312311131-1131223133302232-0330101103032003-0211321220323100"></a>

## reauth_disabled — reauth_disabled / 121033221230 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220)
- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-3320222211011213-1313031020110201-1001212000222033-3130231131000323-0312231100011132-2003002022233112-3032113203121302-3122030323211321)
- reauth_disabled

<a id="canonical-0232123231303000-1010110202312000-3010332303002023-2100332210033202-1102323103223213-3013303331332210-1130203300311122-0121112330310231"></a>

Type: `["object", {}]`. Optional.

\[OneOf: reauth\_disabled, reauth\_timeout\_days, reauth\_timeout\_hours\] Enable this option

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

OneOf alternatives in this subsection:

- [reauth_disabled](resources--ike_phase1_profile--reference--group-001.md#canonical-0232123231303000-1010110202312000-3010332303002023-2100332210033202-1102323103223213-3013303331332210-1130203300311122-0121112330310231)
- [reauth_timeout_days](resources--ike_phase1_profile--reference--group-001.md#canonical-0311110201230321-3103202032001301-1000102032200012-3222303110233303-2011313011330121-3101312033321223-2203123222020230-2210011303123210)
- [reauth_timeout_hours](resources--ike_phase1_profile--reference--group-001.md#canonical-3220112000103122-1132010010111211-3020321110313102-3102120210010203-1023300210332101-1000212201211230-0132303000312010-1033113020132320)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
reauth_disabled = {}
```

<a id="canonical-3320331222110200-2331220220210300-1020200002223022-3102321211310230-2311220011020023-2211222031011232-2012131203032313-2231120323313133"></a>

## Direct properties — reauth_disabled / 121033221230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323301322300200-0221103212101311-2123003322321131-3310100320103123-3323301032233333-2133300031310022-0331313010100112-1200011210210013"></a>

## Next pages — reauth_disabled / 121033221230 / 4

- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-3320222211011213-1313031020110201-1001212000222033-3130231131000323-0312231100011132-2003002022233112-3032113203121302-3122030323211321)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220)

<a id="canonical-3022232010212211-3311312110222112-2212001323132301-3322123113101312-3021320311332100-3233231203303320-2011010223222322-1200322130012222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221122002232221-2330121130013133-1033221121232222-2101232222201232-2113121013200200-1321031303212003-2223013313311330-1302323300310320"></a>

## reauth_timeout_days — reauth_timeout_days / 132302312201 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220)
- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-3320222211011213-1313031020110201-1001212000222033-3130231131000323-0312231100011132-2003002022233112-3032113203121302-3122030323211321)
- reauth_timeout_days

<a id="canonical-0311110201230321-3103202032001301-1000102032200012-3222303110233303-2011313011330121-3101312033321223-2203123222020230-2210011303123210"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for reauth timeout days.

Upstream description:

Set Duration in days.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
reauth_timeout_days {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201223021330022-1032021030303003-1101030301231122-1133321131123010-2102323221110221-1223111113020312-3000220112330111-0312311210211303"></a>

## Direct properties — reauth_timeout_days / 132302312201 / 3

<a id="canonical-2323111102122310-2311103321201222-0020013200030122-3102122230201131-0312121110213320-1013022120212031-0100333331301003-3021200113210022"></a>

<a id="canonical-0110332132103331-2212222011213102-2231210021231220-2301022133231010-3122030121330331-1300113333213003-2131131323013202-0123102211111310"></a>

## duration property — reauth_timeout_days / 132302312201 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2002010032101332-3223320311200123-3111012020102100-1000303220002332-1320010312303302-3233303210302300-2012113010213332-2203311213110013"></a>

## Next pages — reauth_timeout_days / 132302312201 / 5

- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-3320222211011213-1313031020110201-1001212000222033-3130231131000323-0312231100011132-2003002022233112-3032113203121302-3122030323211321)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220)

<a id="canonical-2320021330102133-0033231100001111-0111003111200113-0122201001320210-1130201031310221-1203230001112102-0211030232222222-1230133333012201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203201020010221-1103030220030213-3202230230031110-3303011312033103-0333322033332100-3111321332003111-2122201022233230-0012231013101130"></a>

## reauth_timeout_hours — reauth_timeout_hours / 022032233231 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220)
- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-3320222211011213-1313031020110201-1001212000222033-3130231131000323-0312231100011132-2003002022233112-3032113203121302-3122030323211321)
- reauth_timeout_hours

<a id="canonical-3220112000103122-1132010010111211-3020321110313102-3102120210010203-1023300210332101-1000212201211230-0132303000312010-1033113020132320"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for reauth timeout hours.

Upstream description:

Input Hours.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("duration")}
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
reauth_timeout_hours {
  # Configure direct properties listed below.
}
```

<a id="canonical-2211211221210322-0033111011212123-1301300013000200-1200122203211210-2201203013131230-2103110203103002-3311113130311100-1033021112330221"></a>

## Direct properties — reauth_timeout_hours / 022032233231 / 3

<a id="canonical-1110122321110030-2022211332002311-2233131212201313-0200121130231333-0203033222121000-0033130110312131-0203022330112103-0032000213110132"></a>

<a id="canonical-3000130013130200-2023031313210031-0021123332031231-1322222232011022-3010213123313233-3011200031330220-2011330300333203-3023133310120231"></a>

## duration property — reauth_timeout_hours / 022032233231 / 4

Type: `"number"`. Optional.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0200210321000232-0013132313010122-3223232022130133-2212030203131023-1131101312303312-1102302112202303-1002132333202112-1110012111103101"></a>

## Next pages — reauth_timeout_hours / 022032233231 / 5

- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-3320222211011213-1313031020110201-1001212000222033-3130231131000323-0312231100011132-2003002022233112-3032113203121302-3122030323211321)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220)

<a id="canonical-0102202232133320-3203011001112213-3121102113021033-2120301213211033-0312121132102320-2013020020320110-0110320311230222-2331122103023010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221000220020033-3233112202303201-2230222213022223-3010032323310123-3010332333011232-1320100200000302-3333221311121211-3022233110213012"></a>

## timeouts — timeouts / 132220022132 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220)
- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-3320222211011213-1313031020110201-1001212000222033-3130231131000323-0312231100011132-2003002022233112-3032113203121302-3122030323211321)
- timeouts

<a id="canonical-3301102113321031-2330303210300023-3122130330312030-2123123202323003-2032233101031220-1020002021212321-0331320213303332-2121130321332101"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211111321002210-2211120031120020-2113013102012320-1032303233000210-0202210011112121-0310133021202223-1200101001323232-0003010030200201"></a>

## Direct properties — timeouts / 132220022132 / 3

<a id="canonical-0113321303310120-0311331130101002-3221120221201300-2012221120032202-2221211023303211-0302320222010211-0221031310300221-1230233013333130"></a>

<a id="canonical-0022300030220222-1203002110101330-1132120212131321-0031201021223113-0030233111322010-0310313123211312-2320113103121311-0223313202321222"></a>

## create property — timeouts / 132220022132 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0000020030331203-2113230012303232-2113133231230210-3321212130131210-2101130202331023-2200301013120010-3213120232012131-1213302132210012"></a>

<a id="canonical-0132301012332321-3312033112012211-3210012320011333-3110122303032101-1200113321313100-0032333232122012-0332030133331200-1031333201210302"></a>

## delete property — timeouts / 132220022132 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0033123013312202-0213123303203201-1333111021121213-3223333200100301-1003012023321100-0312133000113100-2030321023113122-2101313313333032"></a>

<a id="canonical-0002021321122132-1011022000123303-1012202102111131-1313013303201131-3012323101020023-2310121332022023-1121000333002113-1012111212103201"></a>

## read property — timeouts / 132220022132 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2001100320323320-2000011001000202-2310212321221000-2110321203111301-0331311331331333-2101111030320131-2012332132031113-3130121111001223"></a>

<a id="canonical-0203103313331222-2010313210112311-0022233211213012-3230013113002031-0213130030131030-3110021310100121-1112321011003311-2100122001301003"></a>

## update property — timeouts / 132220022132 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1000310233321022-0201103132211031-1200110300213123-3321203103002001-0302122122120301-0333302001203120-2230321202100302-2232023033321032"></a>

## Next pages — timeouts / 132220022132 / 8

- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-3320222211011213-1313031020110201-1001212000222033-3130231131000323-0312231100011132-2003002022233112-3032113203121302-3122030323211321)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220)

<a id="canonical-3032311121311133-3013223020103033-3312300313101003-0321121110012103-1223120221131333-1100223111301200-1322032200210113-3333013330021011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300011321002312-1313033100010012-2131321113321012-2023022232223011-1330231100123220-2203313000131030-3311001133303100-2130301332202030"></a>

## use_default_keylifetime — use_default_keylifetime / 023020330323 / 2

Breadcrumbs:

- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220)
- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-3320222211011213-1313031020110201-1001212000222033-3130231131000323-0312231100011132-2003002022233112-3032113203121302-3122030323211321)
- use_default_keylifetime

<a id="canonical-2331331133331311-2202012033123012-0010121012101213-2202102322233022-3110023000103031-0122220103102013-2332302031102333-1002312223032113"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use default keylifetime.

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
use_default_keylifetime = {}
```

<a id="canonical-0212131020133121-3321110112203210-3322100022101221-0231011221233002-0032321132002111-1222231330330012-2221120012132221-2010210302321100"></a>

## Direct properties — use_default_keylifetime / 023020330323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013021031222202-1313301223211211-1010022313323132-1303220332123122-2223302130120023-2131221223301303-0333001302330031-1112131122222330"></a>

## Next pages — use_default_keylifetime / 023020330323 / 4

- [Property reference](resources--ike_phase1_profile--reference--group-001.md#canonical-3320222211011213-1313031020110201-1001212000222033-3130231131000323-0312231100011132-2003002022233112-3032113203121302-3122030323211321)
- [xcsh_ike_phase1_profile](../resources/ike_phase1_profile.md#canonical-2212131331003031-3002012130000333-1310000022332011-2212300010303012-2020303310210313-3200311302002200-0312110113121102-3232332103331220)
