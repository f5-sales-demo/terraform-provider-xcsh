---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-1201120213230200-1101033013012021-2001320303111213-0211233201323112-3021101010112333-0333302322213103-1100332031220001-3211020203322231"></a>

## exact_values property — item / 111100013031 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

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

<a id="canonical-2223211010120301-2300120332000302-0230301230002320-3000213132120121-2233213220012032-3100312133213113-0132223102031021-0220012111033133"></a>

<a id="canonical-2203202312013202-2211320310201222-1220012010132100-3300013002121030-3221012222021023-1031112000220211-1330330232133123-2322120203001011"></a>

## regex_values property — item / 111100013031 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

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

<a id="canonical-0020203000313130-2222210230321312-1322103233121322-2021300301122100-3102331132110100-1022213312320113-1120212211132201-0311312000332332"></a>

<a id="canonical-2202231210212332-0223102232333322-0231021303123112-0011030030011310-0302133220122322-1121133301031331-1212100233321120-1101120002001010"></a>

## transformers property — item / 111100013031 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1013323101322002-3320123122332003-0022132120230101-2011303032030320-2002200213123001-2121113100321332-0221230303000200-2332232232213302"></a>

## Next pages — item / 111100013031 / 7

- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0322033212130133-1113133131322212-0103013001021121-2103031200101123-3011211013203013-2000223231222111-3221121300020221-2012013210300023)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2311200300102223-2200011111023033-1303211002130001-2221330223002003-2133120032110001-1030331310222102-1303231310132020-0311310011103102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222202333203010-3011201321332001-2232323232022131-1203232120022201-1330021002222000-1323002103023200-0331031221032021-0332100331120123"></a>

## policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher — tls_fingerprint_matcher / 033121301011 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher

<a id="canonical-3022121031220100-2201133201220320-0330120123030121-1222220100103110-3132003001313313-0103311221233302-1010131330102023-1111300311332212"></a>

Type: `"single"`. Computed.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

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

<a id="canonical-1312330301113233-0212201211022311-1232213223200131-3200113203230321-1230310020310030-3233012322020200-2222133213302113-2030212213010121"></a>

## Direct properties — tls_fingerprint_matcher / 033121301011 / 3

<a id="canonical-3202302311220131-1222123310023032-2332102110330331-3303332332233323-0203101213132103-3000230120302302-3002031032202231-2111020101110133"></a>

<a id="canonical-1232203210303013-2312300122003100-1210323113112331-0330310222120302-2232102032112003-2200201000330020-3323013021121300-2320202133331031"></a>

## classes property — tls_fingerprint_matcher / 033121301011 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2300333300303121-3210012113132203-0303001032123300-0002221102020323-0213130023313010-0302302130333031-0003011112100213-2213211031000103"></a>

<a id="canonical-2020011100102030-1330312223201003-2130032220010010-1312221012330001-2113310102200233-2032132203331130-0220300312312331-1331130200013211"></a>

## exact_values property — tls_fingerprint_matcher / 033121301011 / 5

Type: `["list", "string"]`. Computed.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

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

<a id="canonical-0212302322211111-0210100330232032-0010101323133230-0321113120101010-2232321323010021-1303133103313213-2113332010022110-1102312101301310"></a>

<a id="canonical-1321121322011131-3023113122020011-2313102331323020-1022320001320322-0311221310122212-2312101313001023-1320103102001201-3102313120101203"></a>

## excluded_values property — tls_fingerprint_matcher / 033121301011 / 6

Type: `["list", "string"]`. Computed.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

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

<a id="canonical-1232103122012120-3020131321201222-2121121211221333-0021232321301202-0213020133022233-2120031131311023-1223120020311012-0220330323102103"></a>

## Next pages — tls_fingerprint_matcher / 033121301011 / 7

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1003222301031030-1210100221301230-3231203311030021-1323133211333010-2213032112313020-0122121033322021-0312113011113021-0102101331131010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203222021032201-2002303301001310-1132000130133310-1303122333133232-0110322021301000-1123001220313322-0002010011312230-3121020312230002"></a>

## policy_based_challenge.temporary_user_blocking — temporary_user_blocking / 023221212223 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.temporary_user_blocking

<a id="canonical-3130201300000331-2022110203333020-2303223020232210-1030201221023110-2222031010102121-2121101111323311-1002212111301102-0120212233321000"></a>

Type: `"single"`. Computed.

Specifies configuration for temporary user blocking resulting from user behavior analysis. When
Malicious User Mitigation is enabled from service policy rules, users' accessing the application
will be analyzed for malicious activity and the configured mitigation actions will be taken on..

Upstream description:

Specifies configuration for temporary user blocking resulting from user behavior analysis.

When Malicious User Mitigation is enabled from service policy rules, users' accessing the
application will be analyzed for malicious activity and the configured mitigation actions will be
taken on identified malicious users. These mitigation actions include setting up temporary blocking
on that user. This configuration specifies settings on how that blocking should be done by the
loadbalancer.

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

<a id="canonical-1330030122221232-3032320021220200-2223123333121131-2320213031112103-2102310303221003-3332331002100332-2023011333123013-1330023333012200"></a>

## Direct properties — temporary_user_blocking / 023221212223 / 3

<a id="canonical-2032310213121230-1132131233303221-2111331000103001-0231302022001202-2023030232301223-2201331203012312-0122111333211102-3011121323302133"></a>

<a id="canonical-0233323310000222-3133111331110310-3101223230201222-2231330233221132-3310022210201220-2213330222300320-2200010313331103-0203212233010311"></a>

## custom_page property — temporary_user_blocking / 023221212223 / 4

Type: `"string"`. Computed.

Custom message is of type . Currently supported URL schemes is . For scheme, message needs to be
encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g.
'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '&lt;p&gt;
Blocked..

Upstream description:

Custom message is of type \`uri\_ref\`. Currently supported URL schemes is \`string:///\`. For
\`string:///\` scheme, message needs to be encoded in Base64 format. You can specify this message as
base64 encoded plain text message e.g. "Blocked.." or it can be HTML paragraph or a body string
encoded as base64 string E.g. "&lt;p&gt; Blocked &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2222222232320330-2231323023203223-0112313003022130-3211103011322201-3000111122311321-3323310121001023-3133032013132202-0112220210300102"></a>

## Next pages — temporary_user_blocking / 023221212223 / 5

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331023020032000-3130010222222110-3233132211031010-0320322211300012-3310133210223023-2013132002122313-3013031023311030-1233221030000020"></a>

## protected_cookies — protected_cookies / 031012032202 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- protected_cookies

<a id="canonical-3322120033102211-3122123112032102-1122311211231022-1323131112100331-2120020030103023-3032101132103323-2311312001013333-3031131030123200"></a>

Type: `"list"`. Computed.

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite.

Upstream description:

Allows setting attributes (SameSite, Secure, and HttpOnly) on cookies in responses. Cookie Tampering
Protection prevents attackers from modifying the value of session cookies. For Cookie Tampering
Protection, enabling a web app firewall (WAF) is a prerequisite. The configured mode of WAF
(monitoring or blocking) will be enforced on the request when cookie tampering is identified. Note:
We recommend enabling Secure and HttpOnly attributes along with cookie tampering protection.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0130221203310133-3110332133313020-3233312022223332-2211102221330203-3100230302301220-1331303311033223-0010331132111300-2120022231321003"></a>

## Direct properties — protected_cookies / 031012032202 / 3

- [add_httponly](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3323020003020112-2002013302330211-0220332211031203-1013000113022221-3222333230121131-1300120212023222-3013101003102201-1111103111220123): complete subsection reference.

- [add_secure](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2332230333100230-3222210000130003-2000331232331123-1300113233003333-2133122012133010-3122212123110331-0221220201020212-2133321231222031): complete subsection reference.

- [disable_tampering_protection](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2123112300330312-3311322211123012-3233130123221132-1133123022332111-0012112231103220-2221131220111321-1132031332210103-1223010030200233): complete subsection reference.

- [enable_tampering_protection](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3212210303231033-3101232000131200-3031202333320000-2323120013002213-3031200300331111-2201102221321322-3212102012021210-1030213131001223): complete subsection reference.

- [ignore_httponly](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3023222000322221-3301111113001313-3001001313302302-2110220011113130-2210332121123222-3203330310123100-2111332003220301-3321332101102310): complete subsection reference.

- [ignore_max_age](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2220113020300131-1310030122001311-1021032023232002-1010000011033313-0211003301232103-1320210122233213-1132320330123030-3311230232100133): complete subsection reference.

- [ignore_samesite](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3100222211313121-1103311102302312-0131002132000200-2230133030233102-2332032012203230-3322022002301330-2220303233200000-0012230103300001): complete subsection reference.

- [ignore_secure](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0121332331101020-3113100032312111-1202213213022300-3230320203102102-0011322323100313-0230013300031020-3311111121313102-0120032101230233): complete subsection reference.

<a id="canonical-3201123130121321-0212322132301123-1000232000303010-3332212000312130-2023111001211333-1203303230233020-2023032002021003-0001133121120111"></a>

<a id="canonical-2233020300213123-2110122300231112-0223012322213113-0030301022332033-3130201113222002-3012021200333100-1103022013011333-3101033330030302"></a>

## max_age_value property — protected_cookies / 031012032202 / 4

Type: `"number"`. Computed.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-3132323013103021-2202330213331231-0121303022123033-0210100120312122-1210130303103321-1230333032212233-0331300102201121-0102230022311131"></a>

<a id="canonical-3332010332123020-1012220313101232-0013201022303032-0003131033133203-3010132122302121-1330000202320232-1122112212131002-2003231020132313"></a>

## name property — protected_cookies / 031012032202 / 5

Type: `"string"`. Computed.

Cookie Name. Name of the Cookie.

Upstream description:

Name of the Cookie.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [samesite_lax](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3223130310002331-3003133311101133-3112312200113123-0022322133120313-2023110033101131-3233030220212132-0321021213210200-3121031220011301): complete subsection reference.

- [samesite_none](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2310103030321232-3230022302331213-1110202221132101-2010001110233330-3310132302212331-1311130130321230-2021130032220032-2133321200020122): complete subsection reference.

- [samesite_strict](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1131123020231133-2212233211031311-2000100023320301-0311233003001103-2021020221202100-3002031000133203-2323213102001301-0101112332030203): complete subsection reference.

<a id="canonical-2230120011103203-0312203333200013-2233212233233132-1331112332303000-0213332233103120-3110303221333312-2300123001203312-0113032310311020"></a>

## Next pages — protected_cookies / 031012032202 / 6

- [protected_cookies.add_httponly](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3323020003020112-2002013302330211-0220332211031203-1013000113022221-3222333230121131-1300120212023222-3013101003102201-1111103111220123)
- [protected_cookies.add_secure](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2332230333100230-3222210000130003-2000331232331123-1300113233003333-2133122012133010-3122212123110331-0221220201020212-2133321231222031)
- [protected_cookies.disable_tampering_protection](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2123112300330312-3311322211123012-3233130123221132-1133123022332111-0012112231103220-2221131220111321-1132031332210103-1223010030200233)
- [protected_cookies.enable_tampering_protection](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3212210303231033-3101232000131200-3031202333320000-2323120013002213-3031200300331111-2201102221321322-3212102012021210-1030213131001223)
- [protected_cookies.ignore_httponly](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3023222000322221-3301111113001313-3001001313302302-2110220011113130-2210332121123222-3203330310123100-2111332003220301-3321332101102310)
- [protected_cookies.ignore_max_age](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2220113020300131-1310030122001311-1021032023232002-1010000011033313-0211003301232103-1320210122233213-1132320330123030-3311230232100133)
- [protected_cookies.ignore_samesite](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3100222211313121-1103311102302312-0131002132000200-2230133030233102-2332032012203230-3322022002301330-2220303233200000-0012230103300001)
- [protected_cookies.ignore_secure](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0121332331101020-3113100032312111-1202213213022300-3230320203102102-0011322323100313-0230013300031020-3311111121313102-0120032101230233)
- [protected_cookies.samesite_lax](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3223130310002331-3003133311101133-3112312200113123-0022322133120313-2023110033101131-3233030220212132-0321021213210200-3121031220011301)
- [protected_cookies.samesite_none](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2310103030321232-3230022302331213-1110202221132101-2010001110233330-3310132302212331-1311130130321230-2021130032220032-2133321200020122)
- [protected_cookies.samesite_strict](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1131123020231133-2212233211031311-2000100023320301-0311233003001103-2021020221202100-3002031000133203-2323213102001301-0101112332030203)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3323020003020112-2002013302330211-0220332211031203-1013000113022221-3222333230121131-1300120212023222-3013101003102201-1111103111220123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132003200002103-3230120121313311-1021320232213000-0202003001032122-2001023311330012-0211333032333011-2011222133232010-3022030033210133"></a>

## protected_cookies.add_httponly — add_httponly / 330100102230 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- protected_cookies.add_httponly

<a id="canonical-1120110301201020-0102210111013210-0023110003201220-2213230102202330-1133330103201120-3220332330100320-0233322033312000-0310131022232123"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for add httponly.

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

<a id="canonical-2332123311033033-3322311323121102-0223213000233120-2303130001112023-2001133231301332-2112121303111123-1131202223313000-2011133000211222"></a>

## Direct properties — add_httponly / 330100102230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030210031312201-2230022020212222-2222311322332202-1301031122203031-2221000033223113-2332311020110201-3332320101011300-0231232200201202"></a>

## Next pages — add_httponly / 330100102230 / 4

- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2332230333100230-3222210000130003-2000331232331123-1300113233003333-2133122012133010-3122212123110331-0221220201020212-2133321231222031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102013331030120-0102303212021330-1212012202000202-3203322023003130-2010030102332113-1321123300102323-0310001031002321-0321321333121332"></a>

## protected_cookies.add_secure — add_secure / 121131111332 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- protected_cookies.add_secure

<a id="canonical-2322010000331122-3303003031023212-1011330023132102-1000212131120303-0132230013313033-1231300133203120-2130012311303013-2321203121212320"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2332023300020102-1001102033220330-1021103113220222-0020001000213231-2200320111201231-1033332010032231-0002223023123222-1112303133331332"></a>

## Direct properties — add_secure / 121131111332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313012232311211-3221020013212302-3113110112203103-1332110133023222-0013200012020210-2233111232113113-0222301112021213-3110013322030133"></a>

## Next pages — add_secure / 121131111332 / 4

- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2123112300330312-3311322211123012-3233130123221132-1133123022332111-0012112231103220-2221131220111321-1132031332210103-1223010030200233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033303311031131-0123300300002003-3103001003310332-1233303020221133-3313122221331131-2102020030310210-2313332032212033-0212000201232030"></a>

## protected_cookies.disable_tampering_protection — disable_tampering_protection / 231213130102 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- protected_cookies.disable_tampering_protection

<a id="canonical-1032301213231220-0001330332322201-0210232100002311-3123311331123013-1330200121300303-1220003032203023-1022123330130201-2220332010000303"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable tampering protection.

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

<a id="canonical-0312303320302102-1330130302023030-3103232003310110-0002301310310212-0221221130100213-0012101322120330-2033313111332110-1233110111012331"></a>

## Direct properties — disable_tampering_protection / 231213130102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033320223013301-0122303120300103-3312330310302310-1300232013222333-3120131230211300-3223330200202012-3010222121130031-2232032301002130"></a>

## Next pages — disable_tampering_protection / 231213130102 / 4

- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3212210303231033-3101232000131200-3031202333320000-2323120013002213-3031200300331111-2201102221321322-3212102012021210-1030213131001223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133023021311311-0011133331000103-3222010310331312-2020023023320033-3212002232011021-2103212130223132-0231213223303203-2031022201111320"></a>

## protected_cookies.enable_tampering_protection — enable_tampering_protection / 000131220210 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- protected_cookies.enable_tampering_protection

<a id="canonical-0221222100121003-2322123002132000-0023321303113020-1300111011032303-3330313000100100-3111120211012101-0230113212001120-2322222213120011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable tampering protection.

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

<a id="canonical-1221112202013131-3021300010100022-2330122210322210-1211321312023203-0101321220020130-3133110230323103-1302121130131221-0032002331220032"></a>

## Direct properties — enable_tampering_protection / 000131220210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030111330331303-2331100031131321-1210023003013012-1102221202210323-2320210030303330-1110202331030030-0201130031122302-3030303032221301"></a>

## Next pages — enable_tampering_protection / 000131220210 / 4

- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3023222000322221-3301111113001313-3001001313302302-2110220011113130-2210332121123222-3203330310123100-2111332003220301-3321332101102310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311013313303032-1032202132021021-2010220033321301-2113000022103021-1232122210003031-2200310331313222-3022001320001010-3120211200003112"></a>

## protected_cookies.ignore_httponly — ignore_httponly / 330003003102 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- protected_cookies.ignore_httponly

<a id="canonical-2233021022120332-0232002323030322-3012332132032001-0303001331203311-0000211000012200-1220123200232313-0210002110320213-3211223312213311"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore httponly.

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

<a id="canonical-1133213202212132-0012103120203123-0201123212030331-3123302130011312-2301110002121031-0110022332003302-3301313021011133-1223102203023133"></a>

## Direct properties — ignore_httponly / 330003003102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223121221120301-0302233122113303-2020213121112003-1112220203222210-2330000230020031-3323010201320211-2202011112000003-3011110233132001"></a>

## Next pages — ignore_httponly / 330003003102 / 4

- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2220113020300131-1310030122001311-1021032023232002-1010000011033313-0211003301232103-1320210122233213-1132320330123030-3311230232100133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312011102303230-1333133203222310-0223131330003032-0203320333132312-3310131102112130-0111032233010120-3332301312221101-2030202220210023"></a>

## protected_cookies.ignore_max_age — ignore_max_age / 010110001030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- protected_cookies.ignore_max_age

<a id="canonical-2120220011022013-0223221202030011-1322212012332131-1002023102201301-2113030333222123-2120220101333001-2011330001311023-3022230231103220"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ignore max age.

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

<a id="canonical-0331303021300221-0113002132202212-0133113331323100-3201013223122123-2313322120133212-1301100130101213-1323301022113011-0221122110231200"></a>

## Direct properties — ignore_max_age / 010110001030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311323323111311-0321113332300302-1003220032002103-1113110103231111-3213213111020330-2201102122112113-3110232020132001-1300111022002103"></a>

## Next pages — ignore_max_age / 010110001030 / 4

- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3100222211313121-1103311102302312-0131002132000200-2230133030233102-2332032012203230-3322022002301330-2220303233200000-0012230103300001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331110003002001-0130213113002333-1303132210111232-1130330232121220-1210302303232322-0332300101000030-1233311113133300-2022120130021030"></a>

## protected_cookies.ignore_samesite — ignore_samesite / 222122321023 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- protected_cookies.ignore_samesite

<a id="canonical-2131003330303320-0321131102311001-1223321232113030-0222001332213000-3333110332022320-0000320133113213-0302111332110232-3301001032331012"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0333120332321203-2223133230210332-2032131321230303-1113230213103220-0202331330130312-1133012033231011-2210320200132211-3213303001123220"></a>

## Direct properties — ignore_samesite / 222122321023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0220010200322331-2302312232011221-3002000231022301-3021222332031131-3312020201011023-0102123000301102-3233220021113001-0023123012222200"></a>

## Next pages — ignore_samesite / 222122321023 / 4

- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0121332331101020-3113100032312111-1202213213022300-3230320203102102-0011322323100313-0230013300031020-3311111121313102-0120032101230233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232132003021013-1120213220012011-3112132310111013-3010112111030023-0213321133113031-0210332001302233-3121330112133131-3003022123232013"></a>

## protected_cookies.ignore_secure — ignore_secure / 112320011023 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- protected_cookies.ignore_secure

<a id="canonical-0320331330112220-2232322012322033-2302233000233032-0000021302223323-0031312330223233-2323012012010011-0302320230302221-3320100322330033"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2231010133213312-3300111101323310-0132232130133113-0201000213201003-2002003210122102-3232312322020121-3203231332122103-3322203230201111"></a>

## Direct properties — ignore_secure / 112320011023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011221021213332-3332120122101233-1211120200220033-2321001100223021-1010003330122323-0002221002123223-0201111230013121-3003330013010123"></a>

## Next pages — ignore_secure / 112320011023 / 4

- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3223130310002331-3003133311101133-3112312200113123-0022322133120313-2023110033101131-3233030220212132-0321021213210200-3121031220011301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120201031202310-3130323221303221-2033320323221112-3212303302031330-3131122202100313-1322030301331013-3121332122030110-1013333221103002"></a>

## protected_cookies.samesite_lax — samesite_lax / 102012023230 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- protected_cookies.samesite_lax

<a id="canonical-1013322203211011-0000010313313120-3020012002320110-3203223031210131-2323000332232001-0032221221033303-0031120112010203-0303010023102112"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1233313302221031-2302312212021133-0132002030021300-0033302300233001-3022111131233320-3233311100121230-1201121002233032-3021333120211300"></a>

## Direct properties — samesite_lax / 102012023230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001221121302333-2303122310002313-2301031300302221-3000113331021210-1212000330201003-1231310001123120-3113120130130220-3102122002233033"></a>

## Next pages — samesite_lax / 102012023230 / 4

- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2310103030321232-3230022302331213-1110202221132101-2010001110233330-3310132302212331-1311130130321230-2021130032220032-2133321200020122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202103222113320-2022022110131311-3031210232230023-3002200021332013-2232210111203113-0202313022101101-2102033332310102-3331222213231313"></a>

## protected_cookies.samesite_none — samesite_none / 323012232030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- protected_cookies.samesite_none

<a id="canonical-2121111111101102-0002300312333210-3021233011121011-2110231322220333-3220322010230312-3212031233200123-1331023202200133-0212221113010313"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1221313331302221-2002313201331322-3211102221012133-3222203210020210-3330222230232111-3202201220332232-3302003001323000-2311230320203202"></a>

## Direct properties — samesite_none / 323012232030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120001101001030-2233101123033201-3231132110313132-1021311223303232-1222003103321123-3033032122003033-3122220133230232-2330011033202320"></a>

## Next pages — samesite_none / 323012232030 / 4

- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1131123020231133-2212233211031311-2000100023320301-0311233003001103-2021020221202100-3002031000133203-2323213102001301-0101112332030203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321131213201100-0131130203301333-0212332032130313-2333113222313001-2200300310230001-3023221210002202-2331022111233023-3323022310223301"></a>

## protected_cookies.samesite_strict — samesite_strict / 203230232003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- protected_cookies.samesite_strict

<a id="canonical-1112230302022331-1023203322310232-2310223133322303-1223200003221233-0120221223020333-3031030210311210-1221301010311230-3213101221223001"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0132132232201331-1122010232223221-1030203032111033-1212113133320100-2202023312011302-1003120121031021-2203103130131120-2010313221111232"></a>

## Direct properties — samesite_strict / 203230232003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301123301131110-0112232110321022-0100330323110033-3112030030330132-2031132101122023-0131203233110223-1123120012201031-2211231012023222"></a>

## Next pages — samesite_strict / 203230232003 / 4

- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213101232101301-1323131001131120-3131203223111320-3102103032331102-1210213002303013-3112103320300201-3201222303331123-2031321230011132"></a>

## rate_limit — rate_limit / 210303223220 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- rate_limit

<a id="canonical-3103123023131201-0322000101103313-0121302000032313-3323033313013312-3121010201201001-2221102102113023-0202100323302113-2113001301020212"></a>

Type: `"single"`. Computed.

RateLimitConfigType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]",
  "x-ves-oneof-field-policy_choice": "[\"no_policies\",\"policies\"]"
}
```

<a id="canonical-1123000202031030-3002123130120000-3231212232231132-3201303030220231-3333013010211013-3201230103323202-1311122112332200-2202201120330123"></a>

## Direct properties — rate_limit / 210303223220 / 3

- [custom_ip_allowed_list](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2212331123312213-2313113110021301-1103033012023210-2113110012322031-0321031332032312-1023121113211323-0211010202011201-0101011300302321): complete subsection reference.

- [ip_allowed_list](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0333030320332113-1220111012132013-0331213200123212-1020313310101321-0130121010303200-3221102003212112-0232220013111220-0023333203201311): complete subsection reference.

- [no_ip_allowed_list](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0022103030001230-1223211301132213-1100000321321011-3202030321102201-2220131303110320-2211123122102332-1002332233231310-0031002132213131): complete subsection reference.

- [no_policies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0210003232302300-3003032111313232-0322001312323312-2202012121312200-0200130130112122-0321211232203001-3301122320312111-3310221030203131): complete subsection reference.

- [policies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2313310002111202-1332131303132010-2111300111031203-3121312022221032-3203333230000233-0300233203011031-2310330201113301-1120323203132331): complete subsection reference.

- [rate_limiter](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303): complete subsection reference.

<a id="canonical-1233301301202130-1001101121321032-2232100202220000-0210220330113301-0200223231323123-3330213011010102-2131201233100003-0131310002300021"></a>

## Next pages — rate_limit / 210303223220 / 4

- [rate_limit.custom_ip_allowed_list](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2212331123312213-2313113110021301-1103033012023210-2113110012322031-0321031332032312-1023121113211323-0211010202011201-0101011300302321)
- [rate_limit.ip_allowed_list](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0333030320332113-1220111012132013-0331213200123212-1020313310101321-0130121010303200-3221102003212112-0232220013111220-0023333203201311)
- [rate_limit.no_ip_allowed_list](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0022103030001230-1223211301132213-1100000321321011-3202030321102201-2220131303110320-2211123122102332-1002332233231310-0031002132213131)
- [rate_limit.no_policies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0210003232302300-3003032111313232-0322001312323312-2202012121312200-0200130130112122-0321211232203001-3301122320312111-3310221030203131)
- [rate_limit.policies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2313310002111202-1332131303132010-2111300111031203-3121312022221032-3203333230000233-0300233203011031-2310330201113301-1120323203132331)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2212331123312213-2313113110021301-1103033012023210-2113110012322031-0321031332032312-1023121113211323-0211010202011201-0101011300302321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110302013320212-3210101033222032-0303232003022020-0010002122103032-3320232023030133-3213130213201133-3000002213220231-0031031233012313"></a>

## rate_limit.custom_ip_allowed_list — custom_ip_allowed_list / 132123121002 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- rate_limit.custom_ip_allowed_list

<a id="canonical-3022130322023032-2203033011030033-3032311210230213-2023023023323032-0033300222332100-0211032132102213-3210313223220330-2313130202322022"></a>

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

<a id="canonical-0302220112023121-3132001211332212-3331121002321212-0210112113013103-1032110223230012-3022023133313032-0110113230230210-1330212031113012"></a>

## Direct properties — custom_ip_allowed_list / 132123121002 / 3

- [rate_limiter_allowed_prefixes](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0020310002131033-0333233032111330-0131212023100212-1303233233032111-1001313123220330-1300033313102210-2300211113211222-0213303230030331): complete subsection reference.

<a id="canonical-1313102331221001-3232020012132003-0100030312123213-3021110231121231-0212200222023330-1032120321000232-1130020103311020-2301001201200210"></a>

## Next pages — custom_ip_allowed_list / 132123121002 / 4

- [rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0020310002131033-0333233032111330-0131212023100212-1303233233032111-1001313123220330-1300033313102210-2300211113211222-0213303230030331)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0020310002131033-0333233032111330-0131212023100212-1303233233032111-1001313123220330-1300033313102210-2300211113211222-0213303230030331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031020003103212-0133032110330031-2032313201111010-1221220000212303-1301131101313011-0332112332230200-3030311003012031-1300310220303032"></a>

## rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes — rate_limiter_allowed_prefixes / 110100203120 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.custom_ip_allowed_list](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2212331123312213-2313113110021301-1103033012023210-2113110012322031-0321031332032312-1023121113211323-0211010202011201-0101011300302321)
- rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-0311311322220122-2231003300033110-3311211330001313-2300221021021301-2211023101120132-1113120313003321-1110001211221333-2331230001332032"></a>

Type: `"list"`. Computed.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3300011320313122-0120221001021011-3133131321220222-1001022032110001-0012221333001131-2332321201330120-1103000313123200-0203333312321212"></a>

## Direct properties — rate_limiter_allowed_prefixes / 110100203120 / 3

<a id="canonical-0021311001001133-2233323010120310-2312311202001102-0022323213012103-0011032010322322-2301223213331103-1002020100002132-3222322212333230"></a>

<a id="canonical-3031112322013202-1221011120020112-1032332303020230-0132102311232212-3012201003322311-3031300020030123-0122001131201331-2032202300210031"></a>

## name property — rate_limiter_allowed_prefixes / 110100203120 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-1201310123012201-0010310033221310-3111323201033330-1313033100020200-2233032330110311-3211033012130111-3220320221321100-3200321020102333"></a>

<a id="canonical-1032332202322213-3320331312022331-1031302233022131-3112210331033012-1032223222113300-3111111331210111-1101001033203202-0332020130320120"></a>

## namespace property — rate_limiter_allowed_prefixes / 110100203120 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3133332121210230-1132002210001020-3211212301000202-2000033031310030-0230330032101330-1033322220031323-2313221300212033-1110030233101331"></a>

<a id="canonical-2102313010221113-0002212232112111-0210221120312122-0033100121210213-2201201012330232-2003313103111022-1033112200313003-2130222030001231"></a>

## tenant property — rate_limiter_allowed_prefixes / 110100203120 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1212003032323220-1103201010233213-3310033020020011-2321212100021202-1202113100103133-2321000020030300-0221333310102020-1003331032301011"></a>

## Next pages — rate_limiter_allowed_prefixes / 110100203120 / 7

- [rate_limit.custom_ip_allowed_list](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2212331123312213-2313113110021301-1103033012023210-2113110012322031-0321031332032312-1023121113211323-0211010202011201-0101011300302321)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0333030320332113-1220111012132013-0331213200123212-1020313310101321-0130121010303200-3221102003212112-0232220013111220-0023333203201311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331012031121012-0221121101030323-3000332313330330-1032201020222230-1120231033131112-3023300200300130-0202313231303231-1131300201011123"></a>

## rate_limit.ip_allowed_list — ip_allowed_list / 102013203323 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- rate_limit.ip_allowed_list

<a id="canonical-0330001223003210-2332230102001321-3310201321231230-2100313130030203-1101300322133000-1123322111221101-2103013103210313-1113032102101330"></a>

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

<a id="canonical-1101300013333220-3030332212100000-1031023330231022-1301232002322003-2020000131303202-2212232233313102-2113312303131320-2110131210003331"></a>

## Direct properties — ip_allowed_list / 102013203323 / 3

<a id="canonical-1010333321232011-0121012300201303-0301301310133003-2000211313133113-0333130030203012-1320303200033333-0130110123201120-2002310221120131"></a>

<a id="canonical-0011102322113313-1120100012233133-3111211210022101-0301330000020320-3132133020023233-3330200231212301-3300111113310013-2122210132021332"></a>

## prefixes property — ip_allowed_list / 102013203323 / 4

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

<a id="canonical-0331133120003102-2001112021001020-3122200231012032-3312132203213200-1021311213202102-3032132122101022-2201310033032033-2030120312223311"></a>

## Next pages — ip_allowed_list / 102013203323 / 5

- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0022103030001230-1223211301132213-1100000321321011-3202030321102201-2220131303110320-2211123122102332-1002332233231310-0031002132213131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021130101010101-3231300013231202-2120032033001101-3321300321000121-2022023202123301-3120300000232012-3131300322310102-2221211230121111"></a>

## rate_limit.no_ip_allowed_list — no_ip_allowed_list / 322320212210 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- rate_limit.no_ip_allowed_list

<a id="canonical-0333312103303200-3313113130200030-3031203033002012-0100033322301230-1022213312310103-1101311011013200-3133231003113312-3131303001110322"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-0221201031200002-1123110220002232-1213312222100213-1221210213323130-1311310201223103-2101301332012113-3022031300021210-0330002200030333"></a>

## Direct properties — no_ip_allowed_list / 322320212210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212320101030311-0331121031311202-3211110231021222-0312223210312332-3010323323100221-3010323230101301-2202212230333000-3231033100102103"></a>

## Next pages — no_ip_allowed_list / 322320212210 / 4

- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0210003232302300-3003032111313232-0322001312323312-2202012121312200-0200130130112122-0321211232203001-3301122320312111-3310221030203131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123112320231202-0021303300021211-0120120023231311-3120031222031323-1130123021223202-2301332223030230-1233313301330132-0102222022310201"></a>

## rate_limit.no_policies — no_policies / 023102121323 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- rate_limit.no_policies

<a id="canonical-0022322201230301-3222103210212231-1113011131021011-1320100310233333-1302021013311020-2123212300212020-1223110223230031-3013011313100001"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no policies. Defaults to \`map\[\]\`. Server applies default when
omitted.

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

<a id="canonical-2103002111021023-0323102002302100-1223010123222313-1200220020222100-1201321020022223-0102220121333010-1021103332120211-1000200212231102"></a>

## Direct properties — no_policies / 023102121323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311001010321233-3011213211103001-0222233301111220-1020331230031223-0132222133003103-2320123131003300-1002123312123321-1110022222013020"></a>

## Next pages — no_policies / 023102121323 / 4

- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2313310002111202-1332131303132010-2111300111031203-3121312022221032-3203333230000233-0300233203011031-2310330201113301-1120323203132331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122213033323300-3013333313111133-0231001223320030-3232000203010213-3232112133323313-1200320212103001-2312313220132032-3202320103103222"></a>

## rate_limit.policies — policies / 122331022001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- rate_limit.policies

<a id="canonical-2002302101102310-2312130131203003-2103203120111012-0000100230221121-3101002120331211-2021303310233010-1131121311001033-1312121022302122"></a>

Type: `"single"`. Computed.

List of rate limiter policies to be applied.

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

<a id="canonical-2033101330203133-1301031101103222-2130030220311130-2013221122002030-0300331010111023-0211330333122311-3322021023212232-1132300212220110"></a>

## Direct properties — policies / 122331022001 / 3

- [policies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2123030230123213-3003201113200321-2102002112113222-1221332132310001-0233022103323232-0111022033303003-2232321202011322-0212021002130321): complete subsection reference.

<a id="canonical-2033003211123132-2232002233202122-2010032223333110-3103120010301230-3322213022132200-2232321123030333-2213203212102300-2010013103120200"></a>

## Next pages — policies / 122331022001 / 4

- [rate_limit.policies.policies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2123030230123213-3003201113200321-2102002112113222-1221332132310001-0233022103323232-0111022033303003-2232321202011322-0212021002130321)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2123030230123213-3003201113200321-2102002112113222-1221332132310001-0233022103323232-0111022033303003-2232321202011322-0212021002130321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323112011010222-0012201000130232-1230003232333100-0223333012120023-2320302220331222-0131212220133131-3200122223320220-3230201331103013"></a>

## rate_limit.policies.policies — policies / 013022323021 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.policies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2313310002111202-1332131303132010-2111300111031203-3121312022221032-3203333230000233-0300233203011031-2310330201113301-1120323203132331)
- rate_limit.policies.policies

<a id="canonical-3312320300123301-3221022003313313-2321331010100021-1311330100023302-2333121231233230-0020202031200230-0033200232301223-3300312302132201"></a>

Type: `"list"`. Computed.

Rate Limiter Policies. Ordered list of rate limiter policies.

Upstream description:

Ordered list of rate limiter policies.

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-1321232203122003-2321333132021202-0100030302333202-2032303020300033-1203310332010232-2112331201211122-1011003133023002-1010322310002111"></a>

## Direct properties — policies / 013022323021 / 3

<a id="canonical-1131121121101021-3232100301213121-2232123111021101-0231221111110333-3133311002312201-1021321002302223-2231021213011231-2332130121233201"></a>

<a id="canonical-3310131232301303-2201200102310123-1201010302032221-3212011111201030-2003021001313211-2011032223210011-2033333003121010-1113000302323033"></a>

## name property — policies / 013022323021 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3113101002001101-0223303202310133-2110010211311021-3012221112122202-2312002022223022-1200200002221321-1300103022003201-2221021330332323"></a>

<a id="canonical-0323203031321102-0113013123000230-2301221003210302-0233020031022130-2002133130232230-3121010010130101-3003313210333101-1133320301113300"></a>

## namespace property — policies / 013022323021 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2132120312032312-1321312232202210-3212210233130310-0003131012201222-0023033211130200-1130231213001110-0332223132301101-0222030032220303"></a>

<a id="canonical-2230113322203012-3222123202032202-3220010000330111-3203310223313013-2330123001321023-2002300003111122-2231132031011333-3020301332023111"></a>

## tenant property — policies / 013022323021 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0200320330032032-0133313113201221-3211211113000330-2200003333120002-0332221220222102-2022202100203312-3323212011120123-3020103122301313"></a>

## Next pages — policies / 013022323021 / 7

- [rate_limit.policies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2313310002111202-1332131303132010-2111300111031203-3121312022221032-3203333230000233-0300233203011031-2310330201113301-1120323203132331)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212132100032320-3232020123033021-0131210303002121-1332013221032002-0011311311300320-3130011222020220-1013202320020123-2003002102222301"></a>

## rate_limit.rate_limiter — rate_limiter / 321121001003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- rate_limit.rate_limiter

<a id="canonical-0323031000020212-1020332030233022-3133010311222331-1230211303131230-3201022033111030-1330201222023320-1030132330000100-2022212301110311"></a>

Type: `"single"`. Computed.

Tuple consisting of a rate limit period unit and the total number of allowed requests for that
period.

Upstream description:

A tuple consisting of a rate limit period unit and the total number of allowed requests for that
period.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"action_block\",\"disabled\"]",
  "x-ves-oneof-field-algorithm": "[\"leaky_bucket\",\"token_bucket\"]"
}
```

<a id="canonical-3022000322213113-1330120131130220-2101121321200022-1101133031210310-0202201000322333-3101222311112323-3221202132313200-2022300022320332"></a>

## Direct properties — rate_limiter / 321121001003 / 3

- [action_block](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3101021031220001-3001323321230112-2231323032031231-3102032211012101-2012121302130120-1303001033213021-3101130232220320-2122303100331133): complete subsection reference.

<a id="canonical-3323232212030120-1203201103011312-1220122323212130-1322200222013110-2103313330130022-2213021000202222-3301232133332321-1330110333132123"></a>

<a id="canonical-1231032002033322-2023233111020222-1320000110221022-3331232220223032-0330330310320211-3322122223033032-0130111021322230-1322113222210220"></a>

## burst_multiplier property — rate_limiter / 321121001003 / 4

Type: `"number"`. Computed.

The maximum burst of requests to accommodate, expressed as a multiple of the rate.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

- [disabled](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0331302311112031-3222133032210112-3302130211000130-1133123210332321-2113023112030121-3311200112202331-0110212213031013-2121312333102320): complete subsection reference.

- [leaky_bucket](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3201330333010112-1212102231030302-0330013131032310-0133101113323222-2230123323311220-1321321310323102-3030011201103331-0122022211303132): complete subsection reference.

<a id="canonical-3311031021010330-3133203130310321-0200200211302221-0211210201031210-1313313003101223-0312331202233300-3203231220221213-0303123033231122"></a>

<a id="canonical-3231111003222023-3212012110212031-2002012011202321-3100113333123330-3033010210120012-2031121123003112-3121312333220311-0301231211103000"></a>

## period_multiplier property — rate_limiter / 321121001003 / 5

Type: `"number"`. Computed.

Setting, combined with Per Period units, provides a duration. Server applies default when omitted.

Upstream description:

This setting, combined with Per Period units, provides a duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
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
    "ves.io.schema.rules.uint32.gte": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0"
  }
}
```

- [token_bucket](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2100233220213310-3331000322232012-1332130230010221-0311210210033233-2323222231303200-1333133201102311-1001122011310233-0321213032103003): complete subsection reference.

<a id="canonical-3012102131022310-2102100031000121-0203030022332112-0012130120231112-3102212220320121-3230020010213300-3011000111113231-0310102103032121"></a>

<a id="canonical-0303001030133201-3023222322310201-3012210010202030-1221120201100131-0011102221113233-0103002131013123-3003323213233223-2003010012303200"></a>

## total_number property — rate_limiter / 321121001003 / 6

Type: `"number"`. Computed.

The total number of allowed requests per rate-limiting period.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8192,
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
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "8192"
  }
}
```

<a id="canonical-1213202012010123-1020212013303230-2033132102130322-3121301212312022-2212203101321131-1211333332310310-3211333200113121-0320020323230232"></a>

<a id="canonical-1031231021221100-2212322010121130-2322021231130232-1221312101211002-1221203122021112-2331012111020230-1013213101030223-3220320230100210"></a>

## unit property — rate_limiter / 321121001003 / 7

Type: `"string"`. Computed.

\[Enum: SECOND|MINUTE|HOUR\] Unit for the period per which the rate limit is applied. - SECOND:
Second Rate limit period unit is seconds - MINUTE: Minute Rate limit period unit is minutes - HOUR:
Hour Rate limit period unit is hours - DAY: Day Rate limit period unit is days. Possible values are
\`SECOND\`, \`MINUTE\`, \`HOUR\`. Defaults to \`SECOND\`.

Upstream description:

Unit for the period per which the rate limit is applied.

&#8203;- SECOND: Second

Rate limit period unit is seconds &#8203;- MINUTE: Minute

Rate limit period unit is minutes &#8203;- HOUR: Hour

Rate limit period unit is hours &#8203;- DAY: Day

Rate limit period unit is days.

Receipt-pinned upstream constraints:

```json
{
  "default": "SECOND",
  "enum": [
    "SECOND",
    "MINUTE",
    "HOUR"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0320121221202010-2322030213211003-1320321032221231-1110013110323203-2012033212131032-3132023020231222-2330002231203013-3003220102320000"></a>

## Next pages — rate_limiter / 321121001003 / 8

- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3101021031220001-3001323321230112-2231323032031231-3102032211012101-2012121302130120-1303001033213021-3101130232220320-2122303100331133)
- [rate_limit.rate_limiter.disabled](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0331302311112031-3222133032210112-3302130211000130-1133123210332321-2113023112030121-3311200112202331-0110212213031013-2121312333102320)
- [rate_limit.rate_limiter.leaky_bucket](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3201330333010112-1212102231030302-0330013131032310-0133101113323222-2230123323311220-1321321310323102-3030011201103331-0122022211303132)
- [rate_limit.rate_limiter.token_bucket](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2100233220213310-3331000322232012-1332130230010221-0311210210033233-2323222231303200-1333133201102311-1001122011310233-0321213032103003)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3101021031220001-3001323321230112-2231323032031231-3102032211012101-2012121302130120-1303001033213021-3101130232220320-2122303100331133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203311320010120-1320312302001211-2103303210030030-1030132312122030-2033230030012212-1121102310011332-1222133031132010-3323201120300321"></a>

## rate_limit.rate_limiter.action_block — action_block / 023231222321 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- rate_limit.rate_limiter.action_block

<a id="canonical-3003333123001223-1001231201202333-0002212202323201-2020312231033021-2113330202001022-2331021312103202-2223002101220330-2010223030202033"></a>

Type: `"single"`. Computed.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

<a id="canonical-0322300220100001-2212121133331030-0233211331312122-0022303110002000-2110222322133130-1302302203323300-1011321132031130-2013132012222321"></a>

## Direct properties — action_block / 023231222321 / 3

- [hours](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3211123323310320-1011311213122201-0203313032010331-3113132222223002-3032123221132011-1313100231111312-2020301103032321-1332011112011111): complete subsection reference.

- [minutes](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2220112301220021-0220033021122302-0201203030311023-3330000301302111-1012133210213233-2121320201330333-2233220312100030-1130330132100023): complete subsection reference.

- [seconds](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3231033012222010-2120232321333112-0312012110302313-2013223121230200-3132113103010121-3220000003010221-1221021310321010-2213212111200310): complete subsection reference.

<a id="canonical-1011001011233312-0302133131111320-0330113111110222-3130130233232000-3122332332320331-3131300111302223-2011203203100020-1311302300021202"></a>

## Next pages — action_block / 023231222321 / 4

- [rate_limit.rate_limiter.action_block.hours](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3211123323310320-1011311213122201-0203313032010331-3113132222223002-3032123221132011-1313100231111312-2020301103032321-1332011112011111)
- [rate_limit.rate_limiter.action_block.minutes](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2220112301220021-0220033021122302-0201203030311023-3330000301302111-1012133210213233-2121320201330333-2233220312100030-1130330132100023)
- [rate_limit.rate_limiter.action_block.seconds](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3231033012222010-2120232321333112-0312012110302313-2013223121230200-3132113103010121-3220000003010221-1221021310321010-2213212111200310)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3211123323310320-1011311213122201-0203313032010331-3113132222223002-3032123221132011-1313100231111312-2020301103032321-1332011112011111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312013113303010-1113001322123211-3321101133310321-1320323312111112-3132010213221010-3101202223300003-1220200013303023-0122020222100031"></a>

## rate_limit.rate_limiter.action_block.hours — hours / 013210032030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3101021031220001-3001323321230112-2231323032031231-3102032211012101-2012121302130120-1303001033213021-3101130232220320-2122303100331133)
- rate_limit.rate_limiter.action_block.hours

<a id="canonical-2202031212331312-1310122211201230-3001313302302132-2130013122111221-0332131012222313-1311121232210300-3030122210233211-2012001232112201"></a>

Type: `"single"`. Computed.

Hours. Input Duration Hours.

Upstream description:

Input Duration Hours.

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

<a id="canonical-1123102030211103-2200210131100301-2323332332200223-2333330201223313-1011223032201122-1010330200031013-0012320230230131-1111131012331030"></a>

## Direct properties — hours / 013210032030 / 3

<a id="canonical-0311001032230122-1231321322233333-3112321100133312-0302130132102101-2301023311113102-1232223020103213-3113313223131223-2333201021220233"></a>

<a id="canonical-0232332011313213-1002110100031000-0012012230222013-2211202231201112-1203231133311333-2130302321112202-0112103202022123-2121231330112011"></a>

## duration property — hours / 013210032030 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 48,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "48"
  }
}
```

<a id="canonical-0122231312222212-2021212303003111-3303322303203322-3311321000233302-2022130112110022-0232221112330020-1210120000210321-3223121302002201"></a>

## Next pages — hours / 013210032030 / 5

- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3101021031220001-3001323321230112-2231323032031231-3102032211012101-2012121302130120-1303001033213021-3101130232220320-2122303100331133)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2220112301220021-0220033021122302-0201203030311023-3330000301302111-1012133210213233-2121320201330333-2233220312100030-1130330132100023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311320210010122-0100102132213000-2101301233131321-1032030211103312-3100332020110330-3230132330201330-2121023222213112-1021322230123312"></a>

## rate_limit.rate_limiter.action_block.minutes — minutes / 130323322012 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3101021031220001-3001323321230112-2231323032031231-3102032211012101-2012121302130120-1303001033213021-3101130232220320-2122303100331133)
- rate_limit.rate_limiter.action_block.minutes

<a id="canonical-2202301202301330-1010021011122321-3232302032013101-2310322213233012-0303030230002032-2122220012110210-0333012212110112-1303011232010103"></a>

Type: `"single"`. Computed.

Minutes. Input Duration Minutes.

Upstream description:

Input Duration Minutes.

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

<a id="canonical-2003103323213010-2200212033300122-2202331213132233-3023322332323221-3200013102303001-1211030022123311-3003233011310131-3112123113302010"></a>

## Direct properties — minutes / 130323322012 / 3

<a id="canonical-1031100223211130-2020212312232323-1231100310320312-2232302102310113-2312333303212312-1032300233203003-1131331233332111-0212311122220000"></a>

<a id="canonical-3303321021201333-0030010212213231-2233132220101123-2130221120000310-1033200223231102-0130002130103003-0132121330013221-1300301021321310"></a>

## duration property — minutes / 130323322012 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "60"
  }
}
```

<a id="canonical-3130001123131122-3131210031023031-2330331211200010-1322211121012003-3121031312223231-2003133102031030-0121031331213131-1331312001213112"></a>

## Next pages — minutes / 130323322012 / 5

- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3101021031220001-3001323321230112-2231323032031231-3102032211012101-2012121302130120-1303001033213021-3101130232220320-2122303100331133)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3231033012222010-2120232321333112-0312012110302313-2013223121230200-3132113103010121-3220000003010221-1221021310321010-2213212111200310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313122001123200-3322200302102333-2311311031302231-3233133330310133-1300311123203210-0222301132030013-0233200021123332-3223012013011313"></a>

## rate_limit.rate_limiter.action_block.seconds — seconds / 333110013001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3101021031220001-3001323321230112-2231323032031231-3102032211012101-2012121302130120-1303001033213021-3101130232220320-2122303100331133)
- rate_limit.rate_limiter.action_block.seconds

<a id="canonical-2312113213323230-1332103132221020-0012133313200220-1010302211210232-3100133312013311-1211203223120032-3131312023220333-2101200031111233"></a>

Type: `"single"`. Computed.

Seconds. Input Duration Seconds.

Upstream description:

Input Duration Seconds.

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

<a id="canonical-0032330003301233-3233011301221032-2100001323013233-0331230200330222-2032000023302202-1002103310011000-1312120030101233-1203221110113130"></a>

## Direct properties — seconds / 333110013001 / 3

<a id="canonical-1302120110211202-3120111030031011-1210103033003300-1221023000322130-3300022010213230-3021313312212220-3001020023212331-3112313013123232"></a>

<a id="canonical-0202201100232003-3221230102201133-1212103323033131-2133103231220022-0222001323213022-1113033310232122-2030223033011100-1320310201222012"></a>

## duration property — seconds / 333110013001 / 4

Type: `"number"`. Computed.

Duration. Configuration parameter for duration

Upstream description:

Configuration parameter for duration

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "300"
  }
}
```

<a id="canonical-0323113130111211-2011331310033221-1221323211230323-3012131012323221-2020331003231300-1310332220031132-3000302200132213-1103312332000031"></a>

## Next pages — seconds / 333110013001 / 5

- [rate_limit.rate_limiter.action_block](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3101021031220001-3001323321230112-2231323032031231-3102032211012101-2012121302130120-1303001033213021-3101130232220320-2122303100331133)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0331302311112031-3222133032210112-3302130211000130-1133123210332321-2113023112030121-3311200112202331-0110212213031013-2121312333102320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132013221013103-1111121020011231-1001023312102302-1001003001020200-0130031032021321-0030112320231300-2130111203031121-1123000202012302"></a>

## rate_limit.rate_limiter.disabled — disabled / 110012010130 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- rate_limit.rate_limiter.disabled

<a id="canonical-3320112322221201-0020020311231301-1313221021301323-0222220213103031-3220323103013010-0130300011201201-2210210013323010-3033203131230032"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1020013323321113-3112231001232112-2003303001120322-2332320333210003-1313332202022223-0132220121322210-2011111133033000-1133210300331023"></a>

## Direct properties — disabled / 110012010130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322210300331102-3131213120233300-1230312303200103-0002023031000022-3333203120321320-1001003132121023-2323000202103030-1320102213011132"></a>

## Next pages — disabled / 110012010130 / 4

- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3201330333010112-1212102231030302-0330013131032310-0133101113323222-2230123323311220-1321321310323102-3030011201103331-0122022211303132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223003110131231-3220220233200233-2233112223002213-3132001331330030-1321021223200132-0313230202032231-3110000113232220-1132112301210320"></a>

## rate_limit.rate_limiter.leaky_bucket — leaky_bucket / 332032022313 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- rate_limit.rate_limiter.leaky_bucket

<a id="canonical-3321030103331122-2103312033303231-1130111200311102-0333220232132302-2100122101103000-0333311300102021-1001331121030232-0203031320312321"></a>

Type: `["object", {}]`. Computed.

Leaky-Bucket is the default rate limiter algorithm for F5.

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

<a id="canonical-3232030120200303-0132331202331033-0122313232322101-1123303111032003-3021112113230012-2303200213133233-0323220000313212-3012123331130120"></a>

## Direct properties — leaky_bucket / 332032022313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201103113132222-2323232021002003-3101013023223131-2311331011331132-2322233300311030-2102311030020330-0322310212213200-3322202323013123"></a>

## Next pages — leaky_bucket / 332032022313 / 4

- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2100233220213310-3331000322232012-1332130230010221-0311210210033233-2323222231303200-1333133201102311-1001122011310233-0321213032103003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312200202320121-3010033210103002-1012223023321332-1222001330301111-0200031232333111-3231330022222011-0213113301332022-0102310312302223"></a>

## rate_limit.rate_limiter.token_bucket — token_bucket / 130120213203 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303)
- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- rate_limit.rate_limiter.token_bucket

<a id="canonical-2121010302310131-2112203311222031-3133302230101321-2022013321131011-0230123133332301-1302011203202122-3031133221101132-2313321133322220"></a>

Type: `["object", {}]`. Computed.

Token-Bucket is a rate limiter algorithm that is stricter with enforcing limits.

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

<a id="canonical-3033231023223031-1231133330323222-1001013233023002-1312133120002003-1120000003212310-3310310200332200-1321332032110230-0130203303000021"></a>

## Direct properties — token_bucket / 130120213203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002110021023112-0330330013011233-3231112303100101-1112202122032211-2222220120132011-2022233233313123-1321031210202020-1220212123013202"></a>

## Next pages — token_bucket / 130120213203 / 4

- [rate_limit.rate_limiter](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0111332033013113-0121220312113321-3022310133003103-2230000121131300-2132231102201010-3310111002303030-3030210010031310-1123313311312303)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2323021210132303-1002202011101212-1233211310323332-0231102230303132-1000132101121121-1023313113211031-3200220110120300-3202302103311313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022321300200121-0232222331022232-3003331120312033-1223001300130112-0323300121123000-2332202102303212-0311020010002332-3211301003321322"></a>

## sensitive_data_policy — sensitive_data_policy / 232031222220 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- sensitive_data_policy

<a id="canonical-2122233223233301-0013212123322233-3031133033203232-2102201333013333-1030211211110231-1012300301301312-3233203101222202-1212331200200122"></a>

Type: `"single"`. Computed.

Policy configuration for this feature.

Upstream description:

Settings for data type policy.

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

<a id="canonical-0222012323232003-1011311232001000-1332202132121120-3102130133202222-2211320102323012-0022132232010213-0311020133220213-0012222321131311"></a>

## Direct properties — sensitive_data_policy / 232031222220 / 3

- [sensitive_data_policy_ref](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0122021200123201-2333022220103031-1003130123300110-2103010222210022-3331133212310133-1210012021332011-0032031301021233-3012221222302111): complete subsection reference.

<a id="canonical-0030001122301101-2121013130132031-0002012310123233-2102030012012123-3231231133303120-3121211211012221-2120113010020123-1201311332331320"></a>

## Next pages — sensitive_data_policy / 232031222220 / 4

- [sensitive_data_policy.sensitive_data_policy_ref](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0122021200123201-2333022220103031-1003130123300110-2103010222210022-3331133212310133-1210012021332011-0032031301021233-3012221222302111)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0122021200123201-2333022220103031-1003130123300110-2103010222210022-3331133212310133-1210012021332011-0032031301021233-3012221222302111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011111013223121-3102020030003002-1332203333133210-1300133333232011-3211310021313212-3301332110013231-3200023001302331-3112031011131200"></a>

## sensitive_data_policy.sensitive_data_policy_ref — sensitive_data_policy_ref / 222200212323 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [sensitive_data_policy](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2323021210132303-1002202011101212-1233211310323332-0231102230303132-1000132101121121-1023313113211031-3200220110120300-3202302103311313)
- sensitive_data_policy.sensitive_data_policy_ref

<a id="canonical-2231132022303230-0121331112312222-1022210213113132-0333323120223110-2233033323330000-1113210210130010-1231210331321103-3220113011113233"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-2200121233100300-1302130302103211-1323321020213320-1102103023232111-0331211133100103-2110312210220231-0111130300332011-1121331023213231"></a>

## Direct properties — sensitive_data_policy_ref / 222200212323 / 3

<a id="canonical-1332020033303331-1313120133321112-3100120000310211-1323013103023310-0310302132333122-3110331233133001-3022212120032101-3212023010130221"></a>

<a id="canonical-2123122001122122-1201302221023130-1123120312001100-3301011121232220-3030003203232000-2211113002332213-2303003212211033-0202221013300300"></a>

## name property — sensitive_data_policy_ref / 222200212323 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2100302322021300-2102102322033121-2233312023012011-2201213301211232-3032230102100322-1210131201311133-3301021233003312-1100031023112332"></a>

<a id="canonical-0233130133003103-2013320021013221-1313023300010223-2223012323321000-0213212310203333-2213233301022130-2220103310023221-3033101122321312"></a>

## namespace property — sensitive_data_policy_ref / 222200212323 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3101220003101333-1101010010313001-2330113323330010-2021220133222230-2000121331002032-2101112102012220-2112322231031000-1003122203101031"></a>

<a id="canonical-2311101100211020-1313201031201001-3021000311330122-3011123011001321-3012112332322210-1030210300102131-0223013320331111-2330030213102013"></a>

## tenant property — sensitive_data_policy_ref / 222200212323 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1231010120033303-1310111030011210-0133113001230010-2201211200120101-2311020001001002-1111213233221121-2033213030222132-0210100313321120"></a>

## Next pages — sensitive_data_policy_ref / 222200212323 / 7

- [sensitive_data_policy](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2323021210132303-1002202011101212-1233211310323332-0231102230303132-1000132101121121-1023313113211031-3200220110120300-3202302103311313)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2013200020130301-1113102330210232-2012033221211111-1333232332201202-0131030031333211-0001020213120320-0310201210113323-0331122232203010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301312220032222-1102120120010310-3230213120212302-0333311001221012-3212231213222013-0333322111313332-3323022311233231-0023132033001030"></a>

## service_policies_from_namespace — service_policies_from_namespace / 213331330312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- service_policies_from_namespace

<a id="canonical-1331300022003002-1003301212201010-2332302332323033-1313300013202011-1100101111101023-0231310011000031-3302000103312122-2320001102213020"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0311110232231200-2331210100220203-1231333222333131-3311202231102201-2121203302301110-2230312120220201-3220010022300201-0311322232120302"></a>

## Direct properties — service_policies_from_namespace / 213331330312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300102232320013-2211131132222323-2303231220221030-1331132132122022-2321032112301130-1022133330312001-1320321201131320-2233111320000201"></a>

## Next pages — service_policies_from_namespace / 213331330312 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0001301213230111-3103221002232321-2201322121023120-2011320222222231-0223121210332333-2311331313122203-3000301231032130-2101022223112121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130220321321010-2313102021123121-0231122302111023-1121230233120231-1210122113023311-1020033112113322-3033110021213300-2003021132011201"></a>

## slow_ddos_mitigation — slow_ddos_mitigation / 013121101311 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- slow_ddos_mitigation

<a id="canonical-0023310130101023-0230201330111310-3031213012121012-3103313113100121-1221222010012130-0111200013330032-3031010222212132-3121322220130212"></a>

Type: `"single"`. Computed.

\[OneOf: slow\_ddos\_mitigation, system\_default\_timeouts; Default: system\_default\_timeouts\]
'Slow and low' attacks tie up server resources, leaving none available for servicing requests from
actual users.

Upstream description:

"Slow and low" attacks tie up server resources, leaving none available for servicing requests from
actual users.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-request_timeout_choice": "[\"disable_request_timeout\",\"request_timeout\"]"
}
```

OneOf alternatives in this subsection:

- [slow_ddos_mitigation](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0023310130101023-0230201330111310-3031213012121012-3103313113100121-1221222010012130-0111200013330032-3031010222212132-3121322220130212)
- [system_default_timeouts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3001200310333022-1331031220203200-0013302003210231-2102121022303323-2101333130010312-0023302113231120-3031313301021300-2120031030013100)

Select alternatives according to the provider validators above.

<a id="canonical-2010123112311213-1033232033232330-0012001323223210-2102100130111002-1201103101323022-0211322312012320-3230212130133132-0001012020132333"></a>

## Direct properties — slow_ddos_mitigation / 013121101311 / 3

- [disable_request_timeout](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1220220212022233-3120201131002122-2132032112200101-0212011221332003-0001313022200102-0013302320021130-3322132230001110-2320330203021223): complete subsection reference.

<a id="canonical-2320000122020013-0210303313220232-2220033122001120-1331123032022023-0332232210100001-3203213030321012-1112011212033001-0322003020021013"></a>

<a id="canonical-0022210033010313-1112002301330223-1012332210302203-2310120113203323-0211001332020221-1002101012303102-3313223002001122-2320300311202110"></a>

## request_headers_timeout property — slow_ddos_mitigation / 013121101311 / 4

Type: `"number"`. Computed.

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The milliseconds. This setting provides protection against Slowloris attacks. Defaults
to \`10000\`.

Upstream description:

The amount of time the client has to send only the headers on the request stream before the stream
is cancelled. The default value is 10000 milliseconds. This setting provides protection against
Slowloris attacks.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "30000"
  }
}
```

<a id="canonical-0223123000120113-1011122122130300-1220021000012131-3310132122302311-2103302032303201-3322232122210201-0112213020101202-2102233313103332"></a>

<a id="canonical-2311202223030323-2032023330222033-3313022323312030-0012202213300333-3330010313303102-2321030231313033-2231221100301322-0022320123300202"></a>

## request_timeout property — slow_ddos_mitigation / 013121101311 / 5

Type: `"number"`. Computed.

Exclusive with \[disable\_request\_timeout\].

Upstream description:

Exclusive with \[disable\_request\_timeout\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 300000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2000
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2000",
    "ves.io.schema.rules.uint32.lte": "300000"
  }
}
```

<a id="canonical-3000132113031111-1113331200101300-0123030111022312-3110113111133330-1303000113023332-2212022212302111-1132011331220321-2221122333310211"></a>

## Next pages — slow_ddos_mitigation / 013121101311 / 6

- [slow_ddos_mitigation.disable_request_timeout](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1220220212022233-3120201131002122-2132032112200101-0212011221332003-0001313022200102-0013302320021130-3322132230001110-2320330203021223)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1220220212022233-3120201131002122-2132032112200101-0212011221332003-0001313022200102-0013302320021130-3322132230001110-2320330203021223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223233111113202-0300122100101001-3112023301300003-0211311313331321-2313102022132121-0202301000101301-3233233030013021-3030121011012213"></a>

## slow_ddos_mitigation.disable_request_timeout — disable_request_timeout / 332021332003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [slow_ddos_mitigation](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0001301213230111-3103221002232321-2201322121023120-2011320222222231-0223121210332333-2311331313122203-3000301231032130-2101022223112121)
- slow_ddos_mitigation.disable_request_timeout

<a id="canonical-2322213132011011-0331103323322123-1000020210313102-3323023222032013-0130023322000023-0110130311321310-3101330333130023-3230013323322310"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable request timeout.

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

<a id="canonical-1233213223103020-2201003103202201-3112202003132231-0232101112202230-3112313210023233-1103101303133113-2312323013003321-3223030220220230"></a>

## Direct properties — disable_request_timeout / 332021332003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321012231311200-2222220203330101-3122033020212033-3200331022031111-0230231000010021-3201120000313022-1202003123322102-1332000020323003"></a>

## Next pages — disable_request_timeout / 332021332003 / 4

- [slow_ddos_mitigation](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0001301213230111-3103221002232321-2201322121023120-2011320222222231-0223121210332333-2311331313122203-3000301231032130-2101022223112121)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3322032032102002-0320311003220131-1102031321312113-0210213301323200-2120310230130103-3321101121233321-0212132023110130-0313002323330033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300210321313112-0333213231320231-0222031330110333-2200303001110121-1011002121120333-1203232303201332-3102030103010233-2231100121020202"></a>

## system_default_timeouts — system_default_timeouts / 301102101332 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- system_default_timeouts

<a id="canonical-3001200310333022-1331031220203200-0013302003210231-2102121022303323-2101333130010312-0023302113231120-3031313301021300-2120031030013100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for system default timeouts.

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

<a id="canonical-2000120310333211-1100300203013323-2132123210312302-1333013033233330-1333232123013000-0311023333031322-3232103212022222-2013331302230010"></a>

## Direct properties — system_default_timeouts / 301102101332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212211323120201-1132123022202333-1122213302002213-3030003132012120-0333021313313023-0202130310312020-3330301131320013-1031113201130111"></a>

## Next pages — system_default_timeouts / 301102101332 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011013103330302-1010132332202310-2230001332023221-0120133211112213-0022232323213110-1312131213102133-2331323121331020-3123002123321130"></a>

## trusted_clients — trusted_clients / 322121221012 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- trusted_clients

<a id="canonical-1231002223203112-2102323012011101-2120210301321000-1212120202121220-1221020100301101-0321313233000031-3032023110221030-3230033232302203"></a>

Type: `"list"`. Computed.

Define rules to skip processing of one or more features such as WAF, Bot Defense etc.

Upstream description:

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. For clients.

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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-3030131311223203-2321003323021201-3130211213032221-1230200220203213-3220130022213133-0311101201223232-2203323300203300-2221221000112312"></a>

## Direct properties — trusted_clients / 322121221012 / 3

<a id="canonical-1103130113231113-2333112300011113-3311010310112210-0211223213022102-3203231223300223-1031303221222231-3120213300323101-0110133020221121"></a>

<a id="canonical-0310113110230023-1221330121223213-0032020313030023-3203311201012020-3031332130120321-1022202203113311-0320222203322010-1030313132113012"></a>

## actions property — trusted_clients / 322121221012 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
SKIP\_PROCESSING\_WAF|SKIP\_PROCESSING\_BOT|SKIP\_PROCESSING\_MUM|SKIP\_PROCESSING\_IP\_REPUTATION|SKIP\_PROCESSING\_API\_PROTECTION|SKIP\_PROCESSING\_OAS\_VALIDATION|SKIP\_PROCESSING\_DDOS\_PROTECTION|SKIP\_PROCESSING\_THREAT\_MESH|SKIP\_PROCESSING\_MALWARE\_PROTECTION\]
Actions that should be taken when client identifier matches the rule. Possible values are
\`SKIP\_PROCESSING\_WAF\`, \`SKIP\_PROCESSING\_BOT\`, \`SKIP\_PROCESSING\_MUM\`,
\`SKIP\_PROCESSING\_IP\_REPUTATION\`, \`SKIP\_PROCESSING\_API\_PROTECTION\`,
\`SKIP\_PROCESSING\_OAS\_VALIDATION\`, \`SKIP\_PROCESSING\_DDOS\_PROTECTION\`,
\`SKIP\_PROCESSING\_THREAT\_MESH\`, \`SKIP\_PROCESSING\_MALWARE\_PROTECTION\`. Defaults to
\`SKIP\_PROCESSING\_WAF\`.

Upstream description:

Actions that should be taken when client identifier matches the rule.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3323210032130133-2102021100011201-1031130311121132-2130110121222130-2032232312320131-1100000200201000-1220203123331330-0210013033021030"></a>

<a id="canonical-1033210023010220-1203121201303103-0331303220013222-2030130312130233-3131231202330001-2020321311222012-3012020100130110-0101230330323110"></a>

## as_number property — trusted_clients / 322121221012 / 5

Type: `"number"`. Computed.

Exclusive with \[http\_header ip\_prefix IPv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Upstream description:

Exclusive with \[http\_header ip\_prefix IPv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 401308,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  }
}
```

- [bot_skip_processing](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1002221023312331-2122202022123303-3203010332010121-1333203310022212-3123210030132320-2032131300232232-2211222212021011-1023102010122303): complete subsection reference.

<a id="canonical-3230320020132232-3120233111333120-0123022133303002-2320331313032212-0003213303123001-0010001212013331-0003312303200233-2202203131022232"></a>

<a id="canonical-2022110310303123-2010211210232131-0330202123111331-2000000322011232-0131221323023123-0103033221222132-3102211022201033-2231021322011320"></a>

## expiration_timestamp property — trusted_clients / 322121221012 / 6

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

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
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [http_header](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3331313011113302-2133211203211001-3331033213021103-3121012200312032-0111302102020100-2110033001021302-3231323300322103-3311322122012121): complete subsection reference.

<a id="canonical-3023211101203210-1123323211202221-0212320321123223-2101132020201111-2300310120133202-3120021312333222-1203322223033103-2013233223023113"></a>

<a id="canonical-2233201111013110-2230331001101030-1320311202110011-2103233011213003-3033121100102213-1320332320121112-3200130001323033-1331201022303233"></a>

## ip_prefix property — trusted_clients / 322121221012 / 7

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header IPv6\_prefix user\_identifier\] IPv4 prefix string.

Upstream description:

Exclusive with \[as\_number http\_header IPv6\_prefix user\_identifier\] IPv4 prefix string.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-2012330230123301-1212221200112112-3120310013001200-3121201102112113-2103020133332301-0002222200331003-2333331121301021-2302002222322223"></a>

<a id="canonical-1133230301321103-1213013022010223-2030033323312022-0100302202133021-0302320133021312-3301203203210130-1322211221021112-2013231133302033"></a>

## ipv6_prefix property — trusted_clients / 322121221012 / 8

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

- [metadata](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3223231021000022-1323110230023303-2001333120313120-2313222332030222-3132312200330212-2013103300312311-0321130232003030-3331121232230020): complete subsection reference.

- [skip_processing](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3121002301000111-3032111303202230-0123013210102222-0022322133120202-2110121100211130-0132320310300120-2321303012320220-1103131131231233): complete subsection reference.

<a id="canonical-1102112132002313-1303333030221120-0020322231011032-3331310122300300-1123121002221021-0322102031022030-0321003220220131-1212000031030200"></a>

<a id="canonical-1230022221130022-0233031313331333-2131011100123100-3030132110311003-2303121000203211-1303332202212112-3022003331301331-0003101012102331"></a>

## user_identifier property — trusted_clients / 322121221012 / 9

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header ip\_prefix IPv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix IPv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

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

- [waf_skip_processing](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3302120220301322-0303132200032323-0103130112023221-1022001202212200-0313022213103220-2013200021311220-1002322311331131-1302331323230003): complete subsection reference.

<a id="canonical-3203213202013233-0023322031110320-2310131010102200-0200010200212120-2100111003233313-1211300320211023-3113123012123123-0030200122301233"></a>

## Next pages — trusted_clients / 322121221012 / 10

- [trusted_clients.bot_skip_processing](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1002221023312331-2122202022123303-3203010332010121-1333203310022212-3123210030132320-2032131300232232-2211222212021011-1023102010122303)
- [trusted_clients.http_header](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3331313011113302-2133211203211001-3331033213021103-3121012200312032-0111302102020100-2110033001021302-3231323300322103-3311322122012121)
- [trusted_clients.metadata](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3223231021000022-1323110230023303-2001333120313120-2313222332030222-3132312200330212-2013103300312311-0321130232003030-3331121232230020)
- [trusted_clients.skip_processing](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3121002301000111-3032111303202230-0123013210102222-0022322133120202-2110121100211130-0132320310300120-2321303012320220-1103131131231233)
- [trusted_clients.waf_skip_processing](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3302120220301322-0303132200032323-0103130112023221-1022001202212200-0313022213103220-2013200021311220-1002322311331131-1302331323230003)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1002221023312331-2122202022123303-3203010332010121-1333203310022212-3123210030132320-2032131300232232-2211222212021011-1023102010122303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221030320312001-2322021311002131-2132321321311220-3011110210030002-0221302111011230-2302021212010013-3003333323332200-1100131301131211"></a>

## trusted_clients.bot_skip_processing — bot_skip_processing / 300212113103 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- trusted_clients.bot_skip_processing

<a id="canonical-1300022011211111-2013333013333230-2101222302130303-2233022201331223-3103312223333030-2121120212122211-3003113100100101-1232320130330003"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1313011032310331-1001312330313203-2300300302111130-1132001101112301-3020201230302220-1211000123320033-0122112211332320-1332212100311212"></a>

## Direct properties — bot_skip_processing / 300212113103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012330211232220-0032133012000220-2033232231320331-1233300203333200-3103220112010320-3310122203223103-2111032110213033-2323320212001022"></a>

## Next pages — bot_skip_processing / 300212113103 / 4

- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3331313011113302-2133211203211001-3331033213021103-3121012200312032-0111302102020100-2110033001021302-3231323300322103-3311322122012121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032032311111011-0013123130312103-0222233123322123-3112230211330110-0133131110310110-1220211230323133-1222101000012103-2030021320213312"></a>

## trusted_clients.http_header — http_header / 211022211321 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- trusted_clients.http_header

<a id="canonical-1033112300123121-2231301220122033-2120223131312202-0021011200120001-3300133130131023-2131223030323223-3323013322200100-0110320123332020"></a>

Type: `"single"`. Computed.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

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

<a id="canonical-2332000023031010-3003032100212121-1102220101023002-0002133131311232-2101112301111212-3131213321110211-3023031223012021-3211121313221323"></a>

## Direct properties — http_header / 211022211321 / 3

- [headers](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3203211213223211-1213100101012313-3103031023200023-0330023211213211-2311233023100013-3232132132220313-1003200120032131-1110023100132320): complete subsection reference.

<a id="canonical-0023012020332321-0023003333201203-2311210130100232-2132310220022232-0212122120013233-2132220021200032-2320322322100322-0321011300111211"></a>

## Next pages — http_header / 211022211321 / 4

- [trusted_clients.http_header.headers](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3203211213223211-1213100101012313-3103031023200023-0330023211213211-2311233023100013-3232132132220313-1003200120032131-1110023100132320)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3203211213223211-1213100101012313-3103031023200023-0330023211213211-2311233023100013-3232132132220313-1003200120032131-1110023100132320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212000013112212-0120212133120111-0303122232000113-0022333012220130-2003001310123112-1211300332031131-0223020030101231-3121203333220222"></a>

## trusted_clients.http_header.headers — headers / 203322020011 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- [trusted_clients.http_header](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3331313011113302-2133211203211001-3331033213021103-3121012200312032-0111302102020100-2110033001021302-3231323300322103-3311322122012121)
- trusted_clients.http_header.headers

<a id="canonical-1111321020000232-2001322220322101-2031322330032332-3002031312121322-0222313323122201-2003302130220203-2312303300021032-1122133302122310"></a>

Type: `"list"`. Computed.

List of HTTP header name and value pairs.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2012120002203313-2302100323003223-1122103121201311-0012013122322210-3322220002222213-2232211100331312-1101013133203113-0223101011203311"></a>

## Direct properties — headers / 203322020011 / 3

<a id="canonical-0032211222003320-2021201213110322-0333333202230233-0103023130302003-2132302021313202-3122030122132130-2302201210122032-0003131322202001"></a>

<a id="canonical-3200212230122130-1110002111112303-3302131202323301-2211123212213332-0123101220202231-0210213230211210-1222333223100130-0001300121312332"></a>

## exact property — headers / 203322020011 / 4

Type: `"string"`. Computed.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-1332311313210200-3210031130301112-1232321123301202-3002211333331200-0302233210100221-2330220231122122-0133133301121102-3210222203332323"></a>

<a id="canonical-3110213130222331-0012101111330022-2133033033323221-0323133222333023-3021211101123212-2332231233300023-0213222103032201-1010122322100010"></a>

## invert_match property — headers / 203322020011 / 5

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

<a id="canonical-0230300301213333-3323021302232030-0202313221222211-2332210321201030-1000230322022210-0323223011211330-3100031332201320-2003122231121223"></a>

<a id="canonical-0332033023232201-3020313002013112-0200100230033312-3202330133013031-0110131201310310-1131011330023320-0102332302033000-2212113001211032"></a>

## name property — headers / 203322020011 / 6

Type: `"string"`. Computed.

Name. Name of the header.

Upstream description:

Name of the header.

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

<a id="canonical-1322332010222223-2210100110130303-3333033331312311-1332100201332133-1003131311321031-2013102132311310-2012212122301232-0033000230020300"></a>

<a id="canonical-0032023333110121-2231031002221021-2111100303323030-0121331011222200-2103111101030101-3011023320002100-1231030203232131-1212113021111333"></a>

## presence property — headers / 203322020011 / 7

Type: `"bool"`. Computed.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

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

<a id="canonical-1023220201111123-0211333301213000-0222202121012220-0001300030101120-3320113320233102-1332322032012330-0222102100330130-3100331013021233"></a>

<a id="canonical-1113121230302203-1101103330232013-3002332230000130-0333311323301011-1111332200310301-0310002222131320-1210231003331311-2223130320013000"></a>

## regex property — headers / 203322020011 / 8

Type: `"string"`. Computed.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1113203010311223-3213302213000002-2301213320220321-3231200121233333-1323033103233033-2131112320210211-3221132331210133-1003023311210122"></a>

## Next pages — headers / 203322020011 / 9

- [trusted_clients.http_header](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3331313011113302-2133211203211001-3331033213021103-3121012200312032-0111302102020100-2110033001021302-3231323300322103-3311322122012121)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3223231021000022-1323110230023303-2001333120313120-2313222332030222-3132312200330212-2013103300312311-0321130232003030-3331121232230020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333333200213302-1331220220303111-2131120200121233-0010313200220030-2201023233311302-2210121000031032-0020232112301310-3033022330033131"></a>

## trusted_clients.metadata — metadata / 203100032102 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- trusted_clients.metadata

<a id="canonical-0011310032132021-3120112111333121-2311311133330201-2132002122331100-2100020011000103-1120001301222033-2320022030202302-3201031223221021"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

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

<a id="canonical-2212133333112312-3310213213110302-1012112031233000-3002113130311323-3012221322113202-2002233302231232-2210320313312122-3113021322033211"></a>

## Direct properties — metadata / 203100032102 / 3

<a id="canonical-1300331013303003-0333303003012113-2200112003122302-3102010000110330-2001221222233231-0112230301203131-0302233302312101-1102222232111112"></a>

<a id="canonical-0032333130231320-0021120002033131-3123012103031332-3120301202133130-1301233111010030-2100201033102213-1222322221032100-3130300330212203"></a>

## description_spec property — metadata / 203100032102 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0122110002222310-1103202023302132-1110032313230323-0022213120132310-2220112011300032-2330213203112120-2332030330030230-1330301310002312"></a>

<a id="canonical-1120010301011231-2031220211202103-3221331303122022-2122232032221112-3221131122110103-1031032230300331-0121302102101023-2011002302301203"></a>

## name property — metadata / 203100032102 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

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

<a id="canonical-2101123312300110-3222031200012202-1130223030121121-3323131132332202-2200302121110130-0302020131123120-1102003332011301-0112302003030333"></a>

## Next pages — metadata / 203100032102 / 6

- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3121002301000111-3032111303202230-0123013210102222-0022322133120202-2110121100211130-0132320310300120-2321303012320220-1103131131231233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200023123313212-1213233020300031-2230331123013203-0201311303121110-2021030133021003-2200230123232101-0122303230213033-3301100201030210"></a>

## trusted_clients.skip_processing — skip_processing / 001220332033 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- trusted_clients.skip_processing

<a id="canonical-2321322132001300-0303103233131111-0321033123120102-2030111330121302-3313013110203133-3313322102211310-0223121012231012-0101333231331332"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0102211333010333-0231102021231120-3102220022331233-0332320033011022-2232030232203113-0102231220323320-2022113213322210-0103301222023032"></a>

## Direct properties — skip_processing / 001220332033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1031001312212312-0213212233330302-0013131332020122-2131303332320103-0001101022313230-1032002332333223-1200330321123030-3032333213203323"></a>

## Next pages — skip_processing / 001220332033 / 4

- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3302120220301322-0303132200032323-0103130112023221-1022001202212200-0313022213103220-2013200021311220-1002322311331131-1302331323230003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201030010000000-2131301210203023-0011023230210201-2022031120332103-1313032010222121-1202121001201333-0311113120322321-3310312312022202"></a>

## trusted_clients.waf_skip_processing — waf_skip_processing / 013322331311 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- trusted_clients.waf_skip_processing

<a id="canonical-3011013220020030-2133013130301101-3020312311223320-3221330002010120-3003001030122130-3221220202200313-1232020333021013-3000310001331330"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1311020101223320-3102302131032213-0300331211213300-0202322303120000-1102123311312000-0302322120122132-1321331121233131-1031221301312130"></a>

## Direct properties — waf_skip_processing / 013322331311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021201311333132-1120113131332101-3321323303220230-2122211113331002-3103232001231201-1023330002132222-1203012130033321-1333331212202120"></a>

## Next pages — waf_skip_processing / 013322331311 / 4

- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1101232232111221-2200012223301301-3121031112233231-2313000312132330-2230301033331100-2213101132232022-2002003032201231-2132233222201102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203331103212032-0103301323202303-3232202212300123-1323302012222323-2220012233303211-0123110221102030-1120303331010110-0131110132312213"></a>

## user_id_client_ip — user_id_client_ip / 203123312201 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- user_id_client_ip

<a id="canonical-1103212001033132-1131132103332211-3110210103311010-2133102010102320-0331222110230320-0110210021332111-1203332023133231-3032231000330123"></a>

Type: `["object", {}]`. Computed.

\[OneOf: user\_id\_client\_ip, user\_identification\] Enable this option

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

- [user_id_client_ip](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1103212001033132-1131132103332211-3110210103311010-2133102010102320-0331222110230320-0110210021332111-1203332023133231-3032231000330123)
- [user_identification](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3023223112130033-0030213022332131-3112101302132332-0010202322122012-0033013021310002-2333323313013112-0330023220021223-0120320020303321)

Select alternatives according to the provider validators above.

<a id="canonical-1023232110213332-3013010201123123-1301102301022210-0132320032133301-1101311111322222-1010011033223221-0313201202033123-3221322323123331"></a>

## Direct properties — user_id_client_ip / 203123312201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111303200231311-0110320020022010-0113232010212332-2113031233133213-2323322231023212-0302131003001233-2101230001101300-1311211022223032"></a>

## Next pages — user_id_client_ip / 203123312201 / 4

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3030211022123120-0133111110122030-3220101023132303-2132132233213000-3210220030203103-2301322311111110-2301202312112311-2202100000202220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013102103113010-3000133200312020-3301101320121011-2211100302131333-3322202023011321-3123022321331031-0220313002002303-0123003031012121"></a>

## user_identification — user_identification / 010002212003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- user_identification

<a id="canonical-3023223112130033-0030213022332131-3112101302132332-0010202322122012-0033013021310002-2333323313013112-0330023220021223-0120320020303321"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-2320121322033010-3130302222330201-2020023112211232-1102300012322301-2130303011120222-0301230300322231-1331133220333222-0111001021112220"></a>

## Direct properties — user_identification / 010002212003 / 3

<a id="canonical-1203000121123121-2003120000320013-3021233213033300-1013232122303212-0111330103100020-2132223322332030-2121132303023223-3022322332030222"></a>

<a id="canonical-1222031123313313-3033230230323122-1113130201010130-0123013032022301-2001232313002022-0023231023232033-2122012130321000-1020011023323331"></a>

## name property — user_identification / 010002212003 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2313102201031103-3220222022031223-3130302320322232-2230320103003021-1323301021313112-1223003313310012-0113220200313232-1032331323002332"></a>

<a id="canonical-0331120133211202-0130133022213102-3133310333020003-3031110032231022-2300020333131031-3200100111200112-3111333103011211-2201031111101132"></a>

## namespace property — user_identification / 010002212003 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1112103033033030-1122313210000230-2232121313322230-3102213211112031-1130310123321233-0010323102010332-1100201121022200-1011000130123230"></a>

<a id="canonical-2331023020000003-1002003232022200-0001100122022233-2102032012033130-3200023221322223-2112332333203302-3022013311333312-2012300212011203"></a>

## tenant property — user_identification / 010002212003 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0202110033303131-3033302010033231-0330200203122302-2103321212212000-1330231012321120-1001311010320231-3013222133222212-2233311030212302"></a>

## Next pages — user_identification / 010002212003 / 7

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012023231221011-1211012110323012-3121032003001110-2232032220122330-3101202033011331-1320300120230120-0020302030211013-3031331111113111"></a>

## waf_exclusion — waf_exclusion / 021212030033 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- waf_exclusion

<a id="canonical-1323231122322113-0110133233002221-2233201022202133-0020023133332330-1021321132112132-3230311133323020-1111032013201301-0030010131101202"></a>

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

<a id="canonical-0000332110003301-1103002231112210-1321130101012312-3321331020122102-2031310113002113-3302033010130303-2023312330310213-0120230320120233"></a>

## Direct properties — waf_exclusion / 021212030033 / 3

- [waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022): complete subsection reference.

- [waf_exclusion_policy](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0321010221131133-3112232201230130-1002303120223013-3111331130200211-2222333313020223-0221201123100030-2000031120023210-0330322333212330): complete subsection reference.

<a id="canonical-1001322033130231-1302110231012122-1001202122332111-2202230323001021-3102220020313131-3201010323020320-1233232021100333-0013020023031022"></a>

## Next pages — waf_exclusion / 021212030033 / 4

- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_policy](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0321010221131133-3112232201230130-1002303120223013-3111331130200211-2222333313020223-0221201123100030-2000031120023210-0330322333212330)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131310000022202-1212201333003331-0023232313321120-3201203231010001-2002000012133300-0032131312231121-2320100101121330-2331010031211030"></a>

## waf_exclusion.waf_exclusion_inline_rules — waf_exclusion_inline_rules / 133302112312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- waf_exclusion.waf_exclusion_inline_rules

<a id="canonical-0001133202332131-0222200331122130-1301233223001210-1132211132002002-1200002313032112-3200313003101013-1231311302133033-1222223031320123"></a>

Type: `"single"`. Computed.

List of WAF exclusion rules that will be applied inline.

Upstream description:

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

<a id="canonical-3131223023023021-0013003020001123-2102132121122331-3130013321003232-1102223312113102-0332213000020313-3132320022330112-3023122022321222"></a>

## Direct properties — waf_exclusion_inline_rules / 133302112312 / 3

- [rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301): complete subsection reference.

<a id="canonical-2131123321313301-0113102303233223-1120222021222310-0212110100123001-3113002321111102-0103203120103211-2320031310010102-3032122032101231"></a>

## Next pages — waf_exclusion_inline_rules / 133302112312 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021333001303322-1112311131212022-1102232010310002-0101130000121020-1220302132301000-1310101000032302-1300202130312011-3222312123212102"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules — rules / 333332122211 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- waf_exclusion.waf_exclusion_inline_rules.rules

<a id="canonical-2233201300313312-2201030313200332-0002310331323323-0302313112320311-3331210110030310-1100121221222031-1101113011111133-3010002111313021"></a>

Type: `"list"`. Computed.

Ordered list of WAF Exclusions specific to this Load Balancer.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1000220311231331-2101111032003211-0202101132102112-1310310211131312-1311332033213311-3021312322011031-1332211020212002-1103010121031222"></a>

## Direct properties — rules / 333332122211 / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3132002120003311-1331211223030211-2012300231100311-0211230102033330-2222002201013123-2001102322312111-2311300203301320-2233011121021303): complete subsection reference.

- [any_path](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3011001323222321-3133223031103001-2112320103312230-1313030102230103-0330110330321030-0022020300333113-3320123012121202-3331333000312211): complete subsection reference.

- [app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200): complete subsection reference.

<a id="canonical-2300330100121302-3001130111101122-2102111023330102-3233011002021310-0102231300330320-3123120222300021-3012000311232210-3200221223102323"></a>

<a id="canonical-1211103320123301-1102103231031300-1102211212111230-2130220323130122-3203220331303323-0330021322300332-1032121321001132-1112002333332032"></a>

## exact_value property — rules / 333332122211 / 4

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

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

<a id="canonical-0031232100112320-0331031013023202-2300001101213222-2202221022112123-1130333110301300-1213333100211333-0121031113012313-1112221111332330"></a>

<a id="canonical-0133300331033223-0201000331230221-1033100131301221-2122011331020133-3332302031230201-1002102332022302-3113020131020111-3312313211302133"></a>

## expiration_timestamp property — rules / 333332122211 / 5

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

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

- [metadata](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1132002310133120-2021321221110121-0021011222300122-1110320101123002-3122011332033011-1133200310300121-2333230123010330-1301000200133122): complete subsection reference.

<a id="canonical-3011102033200223-1003101133031223-2321110000020311-0313230000001031-2021302132330212-3212123013101222-2020101021313233-3130000220031110"></a>

<a id="canonical-1232110332312102-2123023303310103-2012011301223103-1212031013223003-3010103102011023-1232213000232023-0122131213202122-3112200103231213"></a>

## methods property — rules / 333332122211 / 6

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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

<a id="canonical-3123222003201303-1320033333030111-3203321302233033-1232313323232302-3320121112003303-1213032222120013-3222011223232103-2010212021320021"></a>

<a id="canonical-3103213331321120-2030001133332030-0200110001231101-3132101232113231-0100200003332310-3201310323001322-3012212112011331-1200113030022133"></a>

## path_prefix property — rules / 333332122211 / 7

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

Upstream description:

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0031310112211133-2323113111113311-0222031003301323-2130121233320011-0320123313203222-0300001133030331-1322023301323102-2301230311332302"></a>

<a id="canonical-1122133023001211-0202300120202102-2122021123232202-1011010030223300-1310002331333331-2230220023313323-2212221001032313-3101210210021330"></a>

## path_regex property — rules / 333332122211 / 8

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

Upstream description:

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2122012221212130-2110302233201203-2222320331320201-2133212100312211-0330020023211221-3030302000231321-2113320322312113-2320320132012232"></a>

<a id="canonical-0000022131202313-0211022211000121-0132230133200022-1022312031112001-2111130130203113-1212230000131100-3101022313012302-0333200332100122"></a>

## suffix_value property — rules / 333332122211 / 9

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

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

- [waf_skip_processing](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1032010213111131-0112203220133133-1023020203013112-3001302122333112-3332030132000301-3330222033303230-3313221321211113-1023220000033122): complete subsection reference.

<a id="canonical-2033111210031001-0232002210201101-1213020231032121-0231030220310301-2121302002110310-0200311102223110-0112133330033000-2220001130212311"></a>

## Next pages — rules / 333332122211 / 10

- [waf_exclusion.waf_exclusion_inline_rules.rules.any_domain](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3132002120003311-1331211223030211-2012300231100311-0211230102033330-2222002201013123-2001102322312111-2311300203301320-2233011121021303)
- [waf_exclusion.waf_exclusion_inline_rules.rules.any_path](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3011001323222321-3133223031103001-2112320103312230-1313030102230103-0330110330321030-0022020300333113-3320123012121202-3331333000312211)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200)
- [waf_exclusion.waf_exclusion_inline_rules.rules.metadata](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1132002310133120-2021321221110121-0021011222300122-1110320101123002-3122011332033011-1133200310300121-2333230123010330-1301000200133122)
- [waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1032010213111131-0112203220133133-1023020203013112-3001302122333112-3332030132000301-3330222033303230-3313221321211113-1023220000033122)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3132002120003311-1331211223030211-2012300231100311-0211230102033330-2222002201013123-2001102322312111-2311300203301320-2233011121021303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022110302222233-2300222000321233-1031021021013221-2011232232023313-1233322013221302-1211121212030132-3221232013311100-3313131332210330"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.any_domain — any_domain / 330103220001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_domain

<a id="canonical-3132301232012110-3130301331303131-1032033230303200-3021010202203100-1003130123330103-0011020231212310-2132330122320331-0222131130013131"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1013122013212200-2231302033222020-2000323332100100-1321011312132023-0330133022323102-1220121200210222-2310311320211031-2131300231002322"></a>

## Direct properties — any_domain / 330103220001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010210002232200-2321133331322332-0212103032213323-2200310110301001-2222100111101211-2020120321311133-0130102123131030-2320000023233012"></a>

## Next pages — any_domain / 330103220001 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3011001323222321-3133223031103001-2112320103312230-1313030102230103-0330110330321030-0022020300333113-3320123012121202-3331333000312211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133003321021212-3223212112303202-0122303313313011-1310211333011113-1211120322023303-1022313221332231-2021102033312303-1323012032231132"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.any_path — any_path / 032021220022 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_path

<a id="canonical-1130020210123220-3323310212023102-2310213022323101-0201200331332103-2120001110211332-3110203011312221-0323103122321110-3101101230321110"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2320210300003011-1121312031212011-3010020123130303-2232020130203321-1021313212020231-3023123100330200-2300123120000131-1232212300132331"></a>

## Direct properties — any_path / 032021220022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202320123322321-1320030110301203-0300013123001312-2303113123212130-1203102133200220-2020001130232033-1131332013003111-0122111202222123"></a>

## Next pages — any_path / 032021220022 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031313012122132-1312312031023023-0011023222111322-3301133022321020-1232031110103212-1100323010102232-3300000303122113-3331303012233022"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control — app_firewall_detection_control / 031301200221 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control

<a id="canonical-3323021110332122-0032211020313131-3013102333232221-3222122131121012-0002311131010020-2303131310122000-0201322312323200-1120233320010321"></a>

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

<a id="canonical-0101332202302120-0121001233223201-0131111130201031-3312012100330130-2332030101330030-3113201012000303-3123103332230121-3221300320310010"></a>

## Direct properties — app_firewall_detection_control / 031301200221 / 3

- [exclude_attack_type_contexts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2111021103111010-2200232001233022-1313121320012230-2001203300032101-1131321023110033-2113101020222323-1323120233312112-1312211020333331): complete subsection reference.

- [exclude_bot_name_contexts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0311123311100230-3102011311332200-2103010021121131-3320123011302201-2031213133123331-1301212330011221-1100130300223300-3201301012320211): complete subsection reference.

- [exclude_signature_contexts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3123002100101323-3100130021102213-2332233320213022-1120102222231300-2012132220113223-0311001130122230-3120231221010102-3300013312330112): complete subsection reference.

- [exclude_violation_contexts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2011222110302220-2303223101032203-3312222312001320-1110120000023210-0220032212211332-0100000003301221-1001233103030133-0213213231031311): complete subsection reference.

<a id="canonical-1310022231001202-3331213312100011-2012311010312100-3033101333033103-3323102111210201-3120011120101230-1013121113302300-1210133033230121"></a>

## Next pages — app_firewall_detection_control / 031301200221 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2111021103111010-2200232001233022-1313121320012230-2001203300032101-1131321023110033-2113101020222323-1323120233312112-1312211020333331)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0311123311100230-3102011311332200-2103010021121131-3320123011302201-2031213133123331-1301212330011221-1100130300223300-3201301012320211)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3123002100101323-3100130021102213-2332233320213022-1120102222231300-2012132220113223-0311001130122230-3120231221010102-3300013312330112)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2011222110302220-2303223101032203-3312222312001320-1110120000023210-0220032212211332-0100000003301221-1001233103030133-0213213231031311)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2111021103111010-2200232001233022-1313121320012230-2001203300032101-1131321023110033-2113101020222323-1323120233312112-1312211020333331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211021000313033-2301120330330013-0123230221022103-1103122133030021-1203232223021321-3300321031321010-3132132301222001-1301132211120120"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts — exclude_attack_type_contexts / 001103133031 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-0322023303120120-0213122113311232-2133023310330033-2222021301002322-0231212032301101-0033000102000322-1133332323100233-1102322211031233"></a>

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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0321133003222113-1331312000111023-3231131331300201-0110011123220120-2221303032221330-0010021212103110-1000102132020211-3130301223121121"></a>

## Direct properties — exclude_attack_type_contexts / 001103133031 / 3

<a id="canonical-0021013303033312-0222021321102001-0111131221001002-0030221002310002-0312312030213200-2123122011132312-1320222010330322-3000001233133201"></a>

<a id="canonical-2213110202113110-3002123230132023-2000100202212233-2222010111123323-2131113210022200-3210012012333113-3000311131122332-3102201103102310"></a>

## context property — exclude_attack_type_contexts / 001103133031 / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

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

<a id="canonical-1211011320020101-0320221310103310-3102222331300300-2311221121011121-3200331021330033-3313321310022310-1012330010333321-3110203232020133"></a>

<a id="canonical-1010211321103230-3312300202301000-0323301223300012-0002331031313203-0120312300110023-3130221103200203-0301210120032231-1202330003011113"></a>

## context_name property — exclude_attack_type_contexts / 001103133031 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2323021011020210-3221032332312002-2310312313222002-0130310023223131-3120223012310011-3130112201201033-1213000010013221-2012100200220011"></a>

<a id="canonical-1321223331102331-0323120211111120-0222322112321101-0312002121102211-1132100121330321-2223133033230322-0302111332100001-0000220130330110"></a>

## exclude_attack_type property — exclude_attack_type_contexts / 001103133031 / 6

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

Upstream description:

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

<a id="canonical-0321302023031322-0132223021021023-1333111303210312-2311303223110021-2032203112020323-0131000030302231-3111303111011210-3220122321321121"></a>

## Next pages — exclude_attack_type_contexts / 001103133031 / 7

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0311123311100230-3102011311332200-2103010021121131-3320123011302201-2031213133123331-1301212330011221-1100130300223300-3201301012320211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302311122213322-3003120302303221-1023231101002100-1330232030323103-0321023203020331-2223330102212220-1233113211330013-2223323221211103"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts — exclude_bot_name_contexts / 203020111310 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-3020213201322123-1231332101222223-2202000021012131-2101120033211130-1001213121223010-1302333022021123-1331220232122332-1022232212222311"></a>

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2022330303011030-0103220301201232-0213301000113010-2130231223232013-2111222121332022-2221012200202013-0011020100120223-2220111232033010"></a>

## Direct properties — exclude_bot_name_contexts / 203020111310 / 3

<a id="canonical-0212203200131123-1102211311100233-3313220323312010-3222101211313030-1112110203023313-3122110312232100-2302200230203222-2120032012010220"></a>

<a id="canonical-1222211123312022-2202022210232011-1023231220220221-3230303113030301-3111132121122310-3100300023222001-0023012203333230-3021012320301202"></a>

## bot_name property — exclude_bot_name_contexts / 203020111310 / 4

Type: `"string"`. Computed.

Bot Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="canonical-2211022102302312-1002113231130222-2001203313113322-3021200331221213-0003330211300130-3122100210032113-3322212311123023-3020021101231010"></a>

## Next pages — exclude_bot_name_contexts / 203020111310 / 5

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3123002100101323-3100130021102213-2332233320213022-1120102222231300-2012132220113223-0311001130122230-3120231221010102-3300013312330112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031130211201321-2022322313310033-3012113333312122-0210213023202122-2333320112123123-0211002131010233-1122320233133103-0331203100020000"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts — exclude_signature_contexts / 031332033010 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-0010210221303301-2232031113102322-2302223031301333-2310021010232120-3301021333132003-0321020321122033-0321133322313131-1121301200132000"></a>

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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2012322302123013-0133333101110020-0323300000032223-3211332110211113-3320110322001000-1222003002100102-1210031323212222-0300012213221132"></a>

## Direct properties — exclude_signature_contexts / 031332033010 / 3

<a id="canonical-0112022113112112-2211221031320132-2223323022302211-2312202232130132-2102210030111021-1111110320233230-3101001021231120-0113332120201110"></a>

<a id="canonical-1002213333100112-2223110023300230-0033032233232032-0103023113220331-0103131333221101-2210002301101210-3223122122312023-2032213310333032"></a>

## context property — exclude_signature_contexts / 031332033010 / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

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

<a id="canonical-1331323212000121-0013330031220033-2330011001233032-3010211310332011-1033330120333303-0210323133330032-0320012330223022-3233023031110221"></a>

<a id="canonical-0013011002101102-0032332302332320-1330110222323110-0002101230301122-2122101021213231-3131200322031301-2231222121232120-3323333130331112"></a>

## context_name property — exclude_signature_contexts / 031332033010 / 5

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0101200310012012-1330111233203311-3002030200011021-1000220333300110-1011030030322012-3110122312331213-3232120322222031-0111213003323031"></a>

<a id="canonical-2022212101310303-2331002113300112-2122210023313003-2311201221321132-0213321131123311-0331113331311112-2230032132122033-3030223120110323"></a>

## signature_id property — exclude_signature_contexts / 031332033010 / 6

Type: `"number"`. Computed.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3112301323001033-3123331200222232-1003231131123103-0010223311203022-1103323011323321-3213132213332020-0311113030211302-1321020211313120"></a>

## Next pages — exclude_signature_contexts / 031332033010 / 7

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2011222110302220-2303223101032203-3312222312001320-1110120000023210-0220032212211332-0100000003301221-1001233103030133-0213213231031311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032232231101030-2213012221300011-2110000333123033-0021001033013023-1021033200012222-0233222301010000-0120202110221123-0211222331302000"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts — exclude_violation_contexts / 113332321030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-0033311111112212-2021000202001303-3321132031201001-3123031313000122-2111333130013100-2201301321311313-2022021231020132-3031122113221323"></a>

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3020323120101330-0120232222111102-1003323031212103-3220312001030302-1323310120101000-0013320121322303-1003102000122333-1202033012111222"></a>

## Direct properties — exclude_violation_contexts / 113332321030 / 3

<a id="canonical-1331130131310233-1021320033110113-0100223012032303-2330232233230230-1100013002233220-0112332120033311-0000313103202302-2320031301131032"></a>

<a id="canonical-3121102312300201-0032313133133030-3232302111120121-2301301130101310-3103123032112322-1100302222233321-0112022320022321-1101311213313101"></a>

## context property — exclude_violation_contexts / 113332321030 / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

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

<a id="canonical-2001031100201002-2121021123232012-2222031331030120-0103102202010233-1001103230313020-0220213100332201-2332030100013323-2311110021212203"></a>
