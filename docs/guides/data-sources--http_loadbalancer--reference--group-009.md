---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0010002031230121-3002230210030330-1320202310122032-1313232333300310-3133320300331020-0132202223120032-2011213231102030-3032113321330013"></a>

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets.tenant` property

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

<a id="canonical-3012122130033112-0131010311331012-2003120230113331-1100102312133003-2331200033202323-0003201133110301-2123033231212332-1313113213330333"></a>

<a id="canonical-1123203020002202-1003330012012113-1330111122230203-0033012020311321-0020322020221111-0132330032212010-2000202110231321-3302102023312002"></a>

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets.uid` property

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

<a id="canonical-3102132222212321-1112001130122301-3120001322031323-0201001213103202-0013311222212313-1233032022030211-3221110233031122-1012212110002223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-0333123001021322-0031021001132201-1030222331200120-2230112232000133-3202200211103310-0020332133231010-0113331200301000-0333112121202121)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list

<a id="canonical-3130302011302012-0133222122312233-0322113321012131-2330310211200330-3202033133220333-2323112113122200-0123222220333123-1002203213222232"></a>

Type: `"single"`. Computed.

List of IP Prefix strings to match against.

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

<a id="canonical-3331032000303300-2321300013122331-0120331201321002-1013211132111022-0301131132212232-1101232133021103-2001101322232120-1120321203130332"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list`

<a id="canonical-0223010011132131-3112010033333301-3202003120303222-2311303113302110-3111022330312110-0030210310131021-1012110200203302-2010200130000301"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list.invert_match` property

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

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

<a id="canonical-0122121310132303-0332032332130302-2233002233322210-2112122222110000-3210303012330132-1321022013023211-2212310101330011-1300112232033132"></a>

<a id="canonical-3303303012332212-3033031010113101-3033322223010223-0101333222022232-3212201010003001-1101111222213212-2211112332311113-0120211030031223"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list.ip_prefixes` property

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0321130332301111-3002100300010220-0302131010311030-1200013203101121-1002202220023010-1330312231132203-1230110112321030-1011331120021232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-0333123001021322-0031021001132201-1030222331200120-2230112232000133-3202200211103310-0020332133231010-0113331200301000-0333112121202121)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list

<a id="canonical-3311203231123203-1030331330001020-1123211023111003-1103323032122311-1203122333033001-1102023200112211-3323030320302203-3130120230031303"></a>

Type: `"single"`. Computed.

IP Threat Category List Type. List of IP threat categories.

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

<a id="canonical-2332320022223110-2212023002201230-0113101021002033-2103220233233203-1023201012030201-1012212213103313-3210110322102210-1202121102123103"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list`

<a id="canonical-0132101312202300-2032301011203332-2231001310021232-1010100001311133-3201033231320333-0000002311322131-3112023331121312-0200003100231222"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list.ip_threat_categories` property

Type: `["list", "string"]`. Computed.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2013323111223303-2223232322331103-1022110201231112-2030012111221230-3312320332323230-3010131021312311-1202332312100320-3211031210133201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--http_loadbalancer--reference--group-008.md#canonical-0333123001021322-0031021001132201-1030222331200120-2230112232000133-3202200211103310-0020332133231010-0113331200301000-0333112121202121)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-2012302311313212-2303331212312222-3101301223023323-2333121203321030-1212033323223012-3321012031332133-0032212111232231-1020011310233331"></a>

Type: `"single"`. Computed.

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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

<a id="canonical-2230010012021210-0320130231203202-0001120222020323-1202130310102302-2132221302120120-3101032121233013-2132233032131222-1022003230211000"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher`

<a id="canonical-3012120103002002-3201011210021331-2120032200022202-3122132110221031-0200233102131132-3203033211113000-3121210132212221-1022103123023233"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher.classes` property

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Additional upstream details:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2101111331011013-0020303211323102-0212102110001202-1302101013201032-3303012300000100-2303302333302313-0212112120233023-2032213001221220"></a>

<a id="canonical-2102001002222001-1321301300121023-2320123002100201-1022202102102200-0333333330231022-3311223313033001-1122132012200303-2011300122013323"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0023231333211213-2131210111002221-1312333001230222-0132232013010023-2032213232221211-0003003233002103-3122120210221010-1222003203111203"></a>

<a id="canonical-3112202110023131-3010233322132023-3332331130033021-2103103122103223-3133300030333002-1302131233221102-2002120310331222-1310312301323321"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher.excluded_values` property

Type: `["list", "string"]`. Computed.

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher

<a id="canonical-0313320123220333-2113330303012001-2002302022131123-1131100111303221-2131033132121312-3222133003103332-3213323010003302-0113030032201332"></a>

Type: `"single"`. Computed.

Configuration parameter for request matcher.

Additional upstream details:

Request conditions for matching a rule.

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

<a id="canonical-2132320212110222-3123020300211233-0231111332100023-0303220113132311-0310032330321012-0020100223200011-1232322221321223-3013322020322212"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher`

- [cookie_matchers](data-sources--http_loadbalancer--reference--group-009.md#canonical-1233311132012023-3023230203131203-1312032221213100-0322132303110222-1133221301001113-1023222132211012-0210111032301211-0010312202122233): complete subsection reference.

- [headers](data-sources--http_loadbalancer--reference--group-009.md#canonical-0222031210020210-2212032113312333-2231312312332100-2211300300033212-0022311310212113-2312131223031001-3120003103121133-2320220022311002): complete subsection reference.

- [jwt_claims](data-sources--http_loadbalancer--reference--group-009.md#canonical-3032123012130201-0213303103022310-2023011220211321-3201131220332210-0110223213022032-2330120233313113-1301103120300010-3112312101021020): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--reference--group-009.md#canonical-2322110002103002-3112212213120000-2331021212233301-3001000321212203-1032030231210311-3111130112120331-1123302031201303-3210000020121310): complete subsection reference.

<a id="canonical-1233311132012023-3023230203131203-1312032221213100-0322132303110222-1133221301001113-1023222132211012-0210111032301211-0010312202122233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers

<a id="canonical-1100300001301313-0320323230022303-3003021132232311-0212211001233333-2100220223122222-0133203103102123-2310033232301020-1100113311021030"></a>

Type: `"list"`. Computed.

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-1203121123320313-1210120333203122-0130223333301112-2232103032002132-3213303111311202-3032020203220321-3011201001233230-1002311332303130"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers`

- [check_not_present](data-sources--http_loadbalancer--reference--group-009.md#canonical-3100322133001101-1332031103121103-1311222330220211-1220333023133000-3111023313110321-3121313022302310-0032302110220320-1001303313031001): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-009.md#canonical-0010130203020310-3011320013001332-3013221123331313-0223331303031103-1211230233001000-3333100120033200-1113023213132211-2121030333011123): complete subsection reference.

<a id="canonical-3013311332111303-0122320311222000-2110223311333210-2222233111033102-0323033000200320-3101032113332313-2203102311300130-1131110022110321"></a>

<a id="canonical-1211200123110323-3012213221122321-0333022213331022-3121211323133010-1130333033301203-2032113231020013-3231001321320323-1332312302110232"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.invert_matcher` property

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

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

- [item](data-sources--http_loadbalancer--reference--group-009.md#canonical-2110313101223112-1210011000011010-2311331120301330-2132111133210030-0302111222331203-3022031303111121-0022300000013333-0102111330113200): complete subsection reference.

<a id="canonical-1310111233120022-2330231133223030-0303313220121113-1321211303101311-3213323131230203-1231111330202112-1203111133110330-2201111112101101"></a>

<a id="canonical-2220212311023313-0031323011102203-3322023222232002-1213221333110220-0311033022032112-1333122023002131-2101032322103330-3123213130202232"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.name` property

Type: `"string"`. Computed.

Cookie Name. A case-sensitive cookie name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3100322133001101-1332031103121103-1311222330220211-1220333023133000-3111023313110321-3121313022302310-0032302110220320-1001303313031001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-009.md#canonical-1233311132012023-3023230203131203-1312032221213100-0322132303110222-1133221301001113-1023222132211012-0210111032301211-0010312202122233)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-3032220033000331-2111122301203210-0120201300221231-2000322300022230-3100302002232201-0012311112312310-1001122021312110-3032103001330110"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-0010130203020310-3011320013001332-3013221123331313-0223331303031103-1211230233001000-3333100120033200-1113023213132211-2121030333011123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-009.md#canonical-1233311132012023-3023230203131203-1312032221213100-0322132303110222-1133221301001113-1023222132211012-0210111032301211-0010312202122233)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-2311210110200123-1213123011013020-2313300222003010-3230030210221300-1220200123131233-3130330113101313-3211302200221013-2333232001331102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-2110313101223112-1210011000011010-2311331120301330-2132111133210030-0302111222331203-3022031303111121-0022300000013333-0102111330113200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-009.md#canonical-1233311132012023-3023230203131203-1312032221213100-0322132303110222-1133221301001113-1023222132211012-0210111032301211-0010312202122233)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item

<a id="canonical-1121101112310022-1300302121313022-0020100002111013-3310303311103020-0231300002310232-1320201131131103-0311100130000213-2300130203030320"></a>

Type: `"single"`. Computed.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-0121010102330232-2102303312120210-2132301030313233-2223001233102203-0300112330011012-0030102031113232-3201111130101121-3133031332232033"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item`

<a id="canonical-3120230232330323-2231303030201000-1213012212310022-3120321302122110-2031331000120002-3020001200012021-1131331300301212-3222233310020010"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0221100110301000-0303100113023201-1111233001231023-0120303203012211-3230131211030330-3130020323223302-3123322230333033-1322133312112203"></a>

<a id="canonical-3111003201303311-1022323202012122-0112121110221333-2130113221330320-3120021303032321-3301311022101310-1203031130131031-0013010103313213"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3320230221121031-3320320000102321-3310322231022212-3102131320001002-2010311032230223-2211221132033201-0323002011113300-2023211220303131"></a>

<a id="canonical-2313000211021120-0303030320130312-1011122311122022-3032022333000230-3122021021302002-3100313003220003-0210103103103012-3200101310021110"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0222031210020210-2212032113312333-2231312312332100-2211300300033212-0022311310212113-2312131223031001-3120003103121133-2320220022311002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers

<a id="canonical-1313002303333002-0310232001012100-0021331231022111-0121100002003102-2123313120212033-0301131213321320-3233332010200132-2310221013203110"></a>

Type: `"list"`. Computed.

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-0333323332121202-2201311302112002-2223133112302033-3101000202101331-3323030131110130-1113223233312213-1210122332122210-0320332133121023"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers`

- [check_not_present](data-sources--http_loadbalancer--reference--group-009.md#canonical-0201231122023333-0113200120000102-2320211223120111-1030001100000200-0211210212132111-0130032100111031-2312130313200331-3231202120333232): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-009.md#canonical-0330331131221221-3121233012200302-3200132021112213-1113310210123331-1122212303321022-1010221202011132-3221032102111300-1032011222002331): complete subsection reference.

<a id="canonical-2032310112123310-0000113230023221-1322001131012033-0102003222211030-0301012213002212-3310311300130030-1131033122223333-0102111323323300"></a>

<a id="canonical-2010201202330320-3333320232303101-1100321001113011-3303030203113210-0122121032012112-2303231102003210-3020323020011322-2013033031002322"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.invert_matcher` property

Type: `"bool"`. Computed.

Invert Header Matcher. Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-009.md#canonical-2322010022333000-0322311001200031-0322202300302210-0312111103110301-1312022112300110-0133203312032010-2303100031333123-3033123213310010): complete subsection reference.

<a id="canonical-2012031310132003-3011301203122022-0013013311030033-1232300112233112-2221231032233330-3313133312211233-1332222112200123-3023330212120133"></a>

<a id="canonical-3331023132330332-3221312223301323-2130332332322010-1012200200022102-2123233121133320-1121010312001221-2201120103202330-1011110301232223"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.name` property

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0201231122023333-0113200120000102-2320211223120111-1030001100000200-0211210212132111-0130032100111031-2312130313200331-3231202120333232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-009.md#canonical-0222031210020210-2212032113312333-2231312312332100-2211300300033212-0022311310212113-2312131223031001-3120003103121133-2320220022311002)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present

<a id="canonical-0101113133020223-0120110230012102-3201131232113331-2011313120033020-1202001212112003-3012030311201101-2113012203030220-1213123101022321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-0330331131221221-3121233012200302-3200132021112213-1113310210123331-1122212303321022-1010221202011132-3221032102111300-1032011222002331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-009.md#canonical-0222031210020210-2212032113312333-2231312312332100-2211300300033212-0022311310212113-2312131223031001-3120003103121133-2320220022311002)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present

<a id="canonical-3222102020202001-0103312220120303-2121100131102032-1321323101000230-0110301221123312-3220311112203303-3330030310021302-0213311210302011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-2322010022333000-0322311001200031-0322202300302210-0312111103110301-1312022112300110-0133203312032010-2303100031333123-3033123213310010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-009.md#canonical-0222031210020210-2212032113312333-2231312312332100-2211300300033212-0022311310212113-2312131223031001-3120003103121133-2320220022311002)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item

<a id="canonical-0231323101301102-0023103300302012-0101013003121112-1001111020220233-0012230032221321-1230230323023211-1332221003032121-3321001303103320"></a>

Type: `"single"`. Computed.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-2312331331121323-2013201113303001-1022000133220023-2311122223002012-0120032011231123-3100002300032010-3303000112200001-0100030102022030"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item`

<a id="canonical-2012223130110223-1003011320323111-0033120203032111-1201322310131210-0200100231133333-0221100301112130-1002033331022002-3133023023233333"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1013100210022231-2021101322003223-1320023330222000-0131111311302220-3123033033101313-1221312323103110-0120211203102121-3312011030011021"></a>

<a id="canonical-0313320321313123-0021212222313220-2031210322131230-2111023222321333-0123213001301233-0113110323313310-0030221332332200-2323302231333032"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2332303311222120-0033202303332322-0203101113322301-1120032312031200-2232020021131121-3301132231020213-1310300300030313-1302231110203202"></a>

<a id="canonical-3300021222131113-0223203012110233-3022323231223213-0113302020323112-1012013302030011-0303233201102221-1213133201221233-1320000210021212"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3032123012130201-0213303103022310-2023011220211321-3201131220332210-0110223213022032-2330120233313113-1301103120300010-3112312101021020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims

<a id="canonical-2310112100033122-0332030012013332-3203321020121110-1302001213010220-2322013131231030-3311313331111133-0011000332130000-0210033201130010"></a>

Type: `"list"`. Computed.

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-0100113030201031-3132120002323112-2012220201303100-0100321300220010-0330221220320031-0112200030212311-3201231030103120-0132212203033313"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims`

- [check_not_present](data-sources--http_loadbalancer--reference--group-009.md#canonical-3320000110003312-0102312223112203-3010202113313023-0203012130030201-0113132223310222-3331020310012212-3321031303133330-0020012111302332): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-009.md#canonical-0012302102101323-0321132312310033-0032002223113200-1101223332021201-0233033031011033-1231302321112232-0100033110130110-0032313320010331): complete subsection reference.

<a id="canonical-3300200120210113-3103113132121330-0121330011210112-3201021010330021-1200212201210011-0322120021002333-1233102301002112-3001330231031231"></a>

<a id="canonical-1123330322103200-0031321133012032-1022020211303220-1113003313222122-3300332131021031-3102030112330131-3230033011221030-2223311001000000"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.invert_matcher` property

Type: `"bool"`. Computed.

Invert Matcher. Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-009.md#canonical-2310200200311320-0200331233120320-1202203323221310-1003000312312233-0223122302200120-2132232121323311-1111003222203131-2223113202032321): complete subsection reference.

<a id="canonical-1111221010002203-3013013303032302-3001222021203222-2212303023330113-2012021300201130-3313213011233011-1131003202211113-1331331313330010"></a>

<a id="canonical-3213220000113322-1110010031201201-0221331023220000-1121301230131120-1130203030332023-1233223130131312-3001312333113212-0220202322111130"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.name` property

Type: `"string"`. Computed.

JWT Claim Name. JWT claim name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3320000110003312-0102312223112203-3010202113313023-0203012130030201-0113132223310222-3331020310012212-3321031303133330-0020012111302332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-009.md#canonical-3032123012130201-0213303103022310-2023011220211321-3201131220332210-0110223213022032-2330120233313113-1301103120300010-3112312101021020)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-2300201300023112-1120303021031211-0102012332321113-2032032233312201-0333313313303320-0202132110103003-0123202132230202-0210122201033330"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-0012302102101323-0321132312310033-0032002223113200-1101223332021201-0233033031011033-1231302321112232-0100033110130110-0032313320010331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-009.md#canonical-3032123012130201-0213303103022310-2023011220211321-3201131220332210-0110223213022032-2330120233313113-1301103120300010-3112312101021020)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present

<a id="canonical-3210310021211113-2220003203011112-0021320031130331-0323000322310100-0331010223310113-2103113022223023-0102033032021222-0101113203302013"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-2310200200311320-0200331233120320-1202203323221310-1003000312312233-0223122302200120-2132232121323311-1111003222203131-2223113202032321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-009.md#canonical-3032123012130201-0213303103022310-2023011220211321-3201131220332210-0110223213022032-2330120233313113-1301103120300010-3112312101021020)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item

<a id="canonical-3310030323200302-3010021132132110-2120103203201000-1031130112111003-2202002330121123-1330101211033020-2313333111200013-1233022211320333"></a>

Type: `"single"`. Computed.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-1130110220333021-3200300220123312-1011302230100121-1102121213000100-0111200212001223-0101301200230231-1101212020330031-2011000220312113"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item`

<a id="canonical-2020332232120010-1032233221122111-0311133323311102-2133222200310230-1132001020310022-3101220222200110-2110311001131301-2332013210122300"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1111321013020112-2123013200322033-1013121333132133-2101212321321201-0100103233300301-2223233301210320-1202203203101320-0313333000330212"></a>

<a id="canonical-0133032111123331-0001213023200021-0303023001113202-2003200023101203-1213203010303021-0322130023012113-3320110203110210-2121211113203013"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0113330123232223-1011020310312200-0132213300220223-2302223101033331-2220022132312202-1232333120211223-1200010113013322-0113212231211201"></a>

<a id="canonical-1223312101113230-1133012110211301-0213303221003122-3212130330020321-2133313133320322-0212312201000330-0223202032202110-0112131301130211"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2322110002103002-3112212213120000-2331021212233301-3001000321212203-1032030231210311-3111130112120331-1123302031201303-3210000020121310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params

<a id="canonical-3130020023322131-2200031321231033-0220201233310310-1233122331103321-3012211300323211-0312023223102301-1321033000102311-3212221213203110"></a>

Type: `"list"`. Computed.

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-3032322103220322-2131232313330312-1232000302022122-1023303003101201-0002113313230302-0200330231331022-3013100332130330-2332130033020313"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params`

- [check_not_present](data-sources--http_loadbalancer--reference--group-009.md#canonical-3113213230330101-0120302223131123-1030230211203231-0201130131201000-0231232311223210-1232232001002121-1323000323011331-1332131223021331): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-009.md#canonical-1131111310101233-1333102022220222-0022200023113101-1112311133223220-1100121202320231-2020331230323310-3220200121030031-3101111323313022): complete subsection reference.

<a id="canonical-2133332003320222-2102233222220023-2312030131213332-3332201033222330-3211020021213120-2111012322031102-3323110130320111-0003221123223003"></a>

<a id="canonical-3123122020211210-2032223203232101-1013300200002233-3023202112030302-3312131222222312-1102121302033011-1123210030321111-3323312321303210"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.invert_matcher` property

Type: `"bool"`. Computed.

Invert Query Parameter Matcher. Invert the match result.

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

- [item](data-sources--http_loadbalancer--reference--group-009.md#canonical-1323212230331200-3220000230311323-2010021033200132-0033320220030133-3021203032120310-3202220023213101-1312133100030100-3121301013302020): complete subsection reference.

<a id="canonical-3121311322203103-3313302201221233-3011303110121321-3223013120323302-0300132003022001-2232010310030101-1032220022121121-0222313231330101"></a>

<a id="canonical-1101123213310332-0210231203313133-3221003110133110-1330213223103031-0003022122110302-0120002221313102-0101200121010131-1330313312100311"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.key` property

Type: `"string"`. Computed.

A case-sensitive HTTP query parameter name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3113213230330101-0120302223131123-1030230211203231-0201130131201000-0231232311223210-1232232001002121-1323000323011331-1332131223021331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-009.md#canonical-2322110002103002-3112212213120000-2331021212233301-3001000321212203-1032030231210311-3111130112120331-1123302031201303-3210000020121310)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present

<a id="canonical-3123121002311033-1301003002003013-3223102211220003-3002133310131203-1022201032312120-0230322131113103-1210333031023033-1033232103111132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-1131111310101233-1333102022220222-0022200023113101-1112311133223220-1100121202320231-2020331230323310-3220200121030031-3101111323313022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-009.md#canonical-2322110002103002-3112212213120000-2331021212233301-3001000321212203-1032030231210311-3111130112120331-1123302031201303-3210000020121310)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present

<a id="canonical-3300321120130312-0233230322023033-2322300012303110-3210113013112211-0202233110103313-2200003330322033-2331012221123121-2310301322120032"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-1323212230331200-3220000230311323-2010021033200132-0033320220030133-3021203032120310-3202220023213101-1312133100030100-3121301013302020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2003011020000210-0302230010310310-2033210013221201-3201303121203221-3200211201302232-2230120323002132-3313331000020011-1111000122032011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2332112133103200-2233222300201211-3030022100212033-0222323032312201-1120011202323000-2321323021131210-0230333310331310-2213011030302122)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-009.md#canonical-2322110002103002-3112212213120000-2331021212233301-3001000321212203-1032030231210311-3111130112120331-1123302031201303-3210000020121310)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item

<a id="canonical-3231233122200023-2130132023131221-2330210012223131-1213233202210323-0333020132120032-3331301101113213-0132023323000220-3331210120020000"></a>

Type: `"single"`. Computed.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-3112111010111321-3321030212012330-0320210210003030-3123122030320110-3120120201111312-3221201220203211-0100230103132002-1002232012013022"></a>

### Direct properties for `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item`

<a id="canonical-3021223100002210-3123031021012011-2330313310232210-3110320103031213-1132110023021132-0300212012032030-2320202033222200-1033200333010223"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3203130231020212-1233301003201311-2322321223312231-3103112120211121-1213113313021321-2211201023031130-2022321211030230-3121110001000212"></a>

<a id="canonical-0332030330133132-0033020113121130-0230233312122300-0303023013102201-2330302211221000-3211233320312210-3131221110123132-3202213100100120"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2222311121012320-3300001232002212-1330113321301331-3030330230131332-2222232322310013-2312220202202032-0121203230033223-3220130301000201"></a>

<a id="canonical-1332100110022031-3210310311201112-0000131330222230-2033203102033102-2030023021232321-3133230130332130-2030021111202203-3030212332312002"></a>

#### `api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3013011132131110-1031303333020113-2201331231223013-0222022331001113-0210133030331301-0013322311301220-2310330013302312-0102233021031233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.custom_ip_allowed_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- api_rate_limit.custom_ip_allowed_list

<a id="canonical-3021102231212311-3022121213121211-3122323001032113-1033022132320223-1121230101030001-2233200210001333-2100031330330210-3030330122020033"></a>

Type: `"single"`. Computed.

IP Allowed list using existing ip\_prefix\_set objects.

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

<a id="canonical-0311320130130110-1011102230333002-0320033102203232-2311030121203113-1202023133133212-0232010321020300-3222012123110000-2313023003013013"></a>

### Direct properties for `api_rate_limit.custom_ip_allowed_list`

- [rate_limiter_allowed_prefixes](data-sources--http_loadbalancer--reference--group-009.md#canonical-0031211130331220-0201102312320102-0103301032011103-0201120122101002-2331320231130032-0220233202103320-2021233133202132-0111322232332111): complete subsection reference.

<a id="canonical-0031211130331220-0201102312320102-0103301032011103-0201120122101002-2331320231130032-0220233202103320-2021233133202132-0111322232332111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.custom_ip_allowed_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3013011132131110-1031303333020113-2201331231223013-0222022331001113-0210133030331301-0013322311301220-2310330013302312-0102233021031233)
- api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-2123011121112130-2131102103020002-2013331321221200-3110122020122123-1232013001021132-0322022212301322-1230103020031102-1122113311210233"></a>

Type: `"list"`. Computed.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-1211132003023113-3301220022213303-1322122330203320-3321222022203110-3232331302310213-2121122311210120-0211110110031030-2122311130303333"></a>

### Direct properties for `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes`

<a id="canonical-1332212302213121-3122001000211202-3222110210031131-0201003223021223-0203102130322011-0332033322121310-1312231312110121-2313013000131311"></a>

#### `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.name` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3030011030111032-1022212222010113-1011001013133000-1320003103331000-0013132131133021-2101023101201300-2133221323312010-3202132102301100"></a>

<a id="canonical-0213102320311312-1312210203300201-0113310020120001-2203020122211212-0331120002323212-0002320021300122-3210202003322010-3003302322203021"></a>

#### `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.namespace` property

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2220022103030032-3211021120123320-3020211021201213-1031030100003300-0301221111323300-0101311031100022-3132202300033012-3332201013031000"></a>

<a id="canonical-0320111332022003-0212132101202132-3202012300030030-3122133010113012-2322201002132220-2021011120230131-2111221212131301-2010120323130311"></a>

#### `api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3111133023102333-1120102130112313-3103100331031323-1110010110020202-1210323033311321-0221232303123301-1130212210120030-1332100011133320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.ip_allowed_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- api_rate_limit.ip_allowed_list

<a id="canonical-2000000202033300-3100211321132331-0121113100101021-0332232030220322-3012231120330000-0231023032331003-2303031030321100-3223120032121000"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-0233313123012110-2102303001232011-2330220023321202-2110032201223230-2222310020001120-3121331311232332-3313312100320122-1022001300111230"></a>

### Direct properties for `api_rate_limit.ip_allowed_list`

<a id="canonical-3023030102200122-2102321322112203-1133310303201032-0032210020332131-2033031120112003-3233133320030130-0200212011333310-0300021330033223"></a>

#### `api_rate_limit.ip_allowed_list.prefixes` property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3012211103120212-2201223321132021-2021010120120230-0112133020310233-2120203202111131-3111213210002021-3101102323132030-1220030133230312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.no_ip_allowed_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- api_rate_limit.no_ip_allowed_list

<a id="canonical-2130013322033233-0312000013002102-2231332331101102-0012202232013332-3313003312102322-3332032131322001-2132302322000230-0022321321031013"></a>

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

<a id="canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- api_rate_limit.server_url_rules

<a id="canonical-0211301223013032-3331233013321023-2213301311302210-0112211232133122-2300212310302200-0122131102013013-2331313333223030-0122203231120031"></a>

Type: `"list"`. Computed.

Ordered domain or base-path rules for path-scoped rate limiting. Each rule must choose exactly one
rate\_limiter\_choice: inline\_rate\_limiter or ref\_rate\_limiter.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

<a id="canonical-2001211232333310-2301210202103330-1122212021211021-1233212031221123-3102032031010121-3011221312102002-3033102131122321-0130123210130113"></a>

### Direct properties for `api_rate_limit.server_url_rules`

- [any_domain](data-sources--http_loadbalancer--reference--group-009.md#canonical-0333032203131322-0000303222311212-0111021203021102-0000100122003233-0131310130201312-1200323300031110-0200001033221122-3011331021203101): complete subsection reference.

<a id="canonical-2032031131033332-1311123201023113-3232321012013032-2231031101200100-0322103211322233-0232023022210201-2323322133000313-3301231103303312"></a>

<a id="canonical-1310133112200231-2113223033321103-0301312230021103-0332102102001100-2023020123030330-2030223022013222-3201110030302001-1331021032022220"></a>

#### `api_rate_limit.server_url_rules.api_group` property

Type: `"string"`. Computed.

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with 'x-F5 Distributed..

Additional upstream details:

Custom groups can be created if user tags paths or operations with "x-F5 Distributed
Cloud-API-group" extensions inside swaggers.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2300302323031321-0021212132202003-0313310012003320-0323210322020020-0200333220110032-3022003120212312-0120103300121023-2133000013102322"></a>

<a id="canonical-2222230030111331-3332310003101122-1220000332110123-1312323133000332-0023221223110313-3310031133213101-2322132330311303-1230220230313100"></a>

#### `api_rate_limit.server_url_rules.base_path` property

Type: `"string"`. Computed.

Base Path. Prefix of the request path.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2000220132112331-2102033200110211-2121131321320133-2112200020102113-3133031031230213-3310131033122100-2220321100032210-2010000303202310): complete subsection reference.

- [inline_rate_limiter](data-sources--http_loadbalancer--reference--group-010.md#canonical-0322033310030223-3022000203103231-3203020110202213-2221301013202210-3110312310203011-0023332220310100-1311220130202000-3201021231210030): complete subsection reference.

- [ref_rate_limiter](data-sources--http_loadbalancer--reference--group-010.md#canonical-0230200231320022-0022211103313020-3313102213001013-2131023001201221-2123213222111222-2111231003133322-2133231102123132-3302111101201122): complete subsection reference.

- [request_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-2301102203323100-1330001120321122-2111331031130013-3130111131000001-0120310131303133-3302220111303313-1333103021131032-2321211212031313): complete subsection reference.

<a id="canonical-3322312020303103-0300133011120031-0310330321233003-1000002032332210-0120003301033332-2202021111320020-3001303321312320-1100022233010010"></a>

<a id="canonical-1012233100213001-1021320331133201-3102310210031210-3212111213323210-1321233121312003-3103133121001302-2311132021201312-0121031030200203"></a>

#### `api_rate_limit.server_url_rules.specific_domain` property

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-0333032203131322-0000303222311212-0111021203021102-0000100122003233-0131310130201312-1200323300031110-0200001033221122-3011331021203101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- api_rate_limit.server_url_rules.any_domain

<a id="canonical-0131123102113322-2311232121200002-1022211202110222-1203032220321021-3212312210013033-2020220233321211-1110120320020303-3023112010123101"></a>

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

<a id="canonical-2000220132112331-2102033200110211-2121131321320133-2112200020102113-3133031031230213-3310131033122100-2220321100032210-2010000303202310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- api_rate_limit.server_url_rules.client_matcher

<a id="canonical-1022032233013203-1000122320233201-3231213203120131-2011131223012010-1000331130022210-2331233322001131-1011003201231020-2031311312201331"></a>

Type: `"single"`. Computed.

Client Matcher. Client conditions for matching a rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

<a id="canonical-1313303112130122-2322220332120123-2301101212332020-0303330210301022-2232030303230133-0023220101313033-1000320030110233-0301021021033003"></a>

### Direct properties for `api_rate_limit.server_url_rules.client_matcher`

- [any_client](data-sources--http_loadbalancer--reference--group-009.md#canonical-2322122103030121-0023320221023133-2202033233010323-1321310102210200-0121030030213122-3110202233310023-1030333210130023-1203223210233021): complete subsection reference.

- [any_ip](data-sources--http_loadbalancer--reference--group-009.md#canonical-0130000021130101-1111303321131101-3031031120103220-2202021003231111-0023213012322303-1013321210313200-0213021300333113-2301001122102312): complete subsection reference.

- [asn_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-1021123312221201-1312222301111302-0233031121333333-1112332031331023-2310013123003021-2321211002000112-1113110312032301-0122123101303130): complete subsection reference.

- [asn_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-0001122010010121-1201010123310202-2120312113113333-0023210203203012-1233023110101200-1222311231232122-0001033100311021-1031221311311303): complete subsection reference.

- [client_selector](data-sources--http_loadbalancer--reference--group-010.md#canonical-3111210223132323-0121012002312020-0220110001201123-3002220100131232-0100103213330022-1330010130033300-2313010211122202-1311133120322302): complete subsection reference.

- [ip_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-3122322002000200-3120320003232323-0302010210201232-1233232103213311-3111211031112111-0220232001101101-2331023211232022-0213001120023010): complete subsection reference.

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-0311222103112303-2230331030010021-0313331233320111-3132301213210203-1021012331121103-1101310110000222-3112130120323031-2300123230221020): complete subsection reference.

- [ip_threat_category_list](data-sources--http_loadbalancer--reference--group-010.md#canonical-3120113231122011-1303331120313130-0331123002200313-1103111321312013-3130331030123131-3022203021213101-0031030311322232-1321332310031012): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-010.md#canonical-0210212023001300-2110302130121011-3111122332320113-1020310333033231-2233103310202212-0200131132000300-0101221231331202-2112211000022313): complete subsection reference.

<a id="canonical-2322122103030121-0023320221023133-2202033233010323-1321310102210200-0121030030213122-3110202233310023-1030333210130023-1203223210233021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.any_client` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2000220132112331-2102033200110211-2121131321320133-2112200020102113-3133031031230213-3310131033122100-2220321100032210-2010000303202310)
- api_rate_limit.server_url_rules.client_matcher.any_client

<a id="canonical-3320003030323320-1212203230231212-3212112310322232-3310033111133113-2232300320113120-0330100023010011-2303213132013100-3230212000010231"></a>

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

<a id="canonical-0130000021130101-1111303321131101-3031031120103220-2202021003231111-0023213012322303-1013321210313200-0213021300333113-2301001122102312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_rate_limit.server_url_rules.client_matcher.any_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-007.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--http_loadbalancer--reference--group-009.md#canonical-2000220132112331-2102033200110211-2121131321320133-2112200020102113-3133031031230213-3310131033122100-2220321100032210-2010000303202310)
- api_rate_limit.server_url_rules.client_matcher.any_ip

<a id="canonical-1003220211233213-3031233033230231-0210122120120322-0303311303023320-0323033322301132-1211012100012003-0302320310120002-0122320203223131"></a>

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
