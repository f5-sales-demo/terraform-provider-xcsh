---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3301120021012012-1020313000003012-1123203210313221-1322230031210222-1111012123310020-3223010332201220-0212332222112100-3133231233032301"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item — item / 320330300221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item

<a id="canonical-2232311113030211-2030310331202002-0203131102122220-2030002002320103-3021122002102223-1301202233212122-1201132221203200-2023020310223133"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-2211211310332120-2321110301321203-1311230133021333-3002130210210013-1023321310310103-2113213220210230-0200211103000113-1110013101022030"></a>

## Direct properties — item / 320330300221 / 3

<a id="canonical-2101222210121330-3321320211222010-0301012111212121-0323102121212032-0100032232223200-3312212032310210-3303013310213121-0030021301321013"></a>

<a id="canonical-3102333232120331-1231301131302132-0123331230102323-3223331332323213-2211310203110102-2211220023323210-1301131123031103-1001111311210032"></a>

## exact_values property — item / 320330300221 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-3131002031111121-3023100021131011-3012133331202130-2210211001202010-3311031111123023-1303022333203021-2230222003313201-2031323012323120"></a>

<a id="canonical-2303333213120021-1133030033013300-3333030303011212-2123230020033230-1010010120022330-0021213121330032-3132301322302232-1302013222212223"></a>

## regex_values property — item / 320330300221 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-0022220002020230-1003010030131333-3303110321122313-2213133011022103-2003221333303101-2020202103302021-0022330130301031-1203130102230330"></a>

<a id="canonical-1322333132222102-2221010011201203-0323031230220232-2000111220222300-2003202011221311-1023102112012332-1313120323200320-0031231302033023"></a>

## transformers property — item / 320330300221 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-1123203021331313-0102231220120310-3312310211321121-0313122111122223-1321312203101021-2003221210203223-1130102220013223-3131111222201221"></a>

## Next pages — item / 320330300221 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-007.md#canonical-3003020030130111-3002032223011033-2322123210022323-1312013322321102-3310300031001223-0201333302120213-2103332320332020-0020220233221311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330032300200001-2002231213322231-0030021030102230-2112000130023222-0013100201203103-1131221031131203-0130212103111212-1033111102230313"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims — jwt_claims / 010112330220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims

<a id="canonical-2131222033203211-3130123231030020-0222211332333330-0010122200133101-1331332022230023-2333001203131011-3033303223200320-1110210131130032"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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

Terraform syntax:

```terraform
jwt_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322120111321202-2121302011200123-0123131332021231-1200233133101300-0130011111110222-1102113021300133-0033302031012311-1112303322220013"></a>

## Direct properties — jwt_claims / 010112330220 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-2030312121020031-0233023120302331-1311101330233222-2010021032222120-2101010233312112-3122310011100011-0101001332310321-2111122212212302): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-008.md#canonical-3310131023313222-0323311302301330-0330121123220032-1103002110303210-1132022103030201-1222113010022210-2003300020022001-0030231132332212): complete subsection reference.

<a id="canonical-2033320033213331-2132313032212122-0011231123030200-3122010210022023-2221203322113032-0211322231302033-0132022102300020-2133133311232320"></a>

<a id="canonical-1031212221000130-0310123030200210-2333103331002002-3210010222021202-2123023023022213-1010221121111033-3302012320133131-0322020132323300"></a>

## invert_matcher property — jwt_claims / 010112330220 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-008.md#canonical-0312023210333011-3220122310300130-1332132131103031-3011312321333103-3130103313320121-3331322220220322-3332331211132303-1113203221103111): complete subsection reference.

<a id="canonical-2020323023320221-1300312210331021-2122123102102202-3132220030330233-1232223303211210-3223000310130101-0030201010223033-2320313003112130"></a>

<a id="canonical-3002320222212132-2030211111332232-2301321233203311-0200010123101302-2333100131111101-0300130221101101-2321010110033102-3101332302003213"></a>

## name property — jwt_claims / 010112330220 / 5

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3010121122030021-1200000300300120-2020221330312122-3211330012210303-0122132231003321-0230101200322012-3222112021100010-3103301132313020"></a>

## Next pages — jwt_claims / 010112330220 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-2030312121020031-0233023120302331-1311101330233222-2010021032222120-2101010233312112-3122310011100011-0101001332310321-2111122212212302)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present](resources--http_loadbalancer--reference--group-008.md#canonical-3310131023313222-0323311302301330-0330121123220032-1103002110303210-1132022103030201-1222113010022210-2003300020022001-0030231132332212)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item](resources--http_loadbalancer--reference--group-008.md#canonical-0312023210333011-3220122310300130-1332132131103031-3011312321333103-3130103313320121-3331322220220322-3332331211132303-1113203221103111)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2030312121020031-0233023120302331-1311101330233222-2010021032222120-2101010233312112-3122310011100011-0101001332310321-2111122212212302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330031112001223-2133032213310011-1312303131230210-3003110123011333-1213332211012211-0323323331203220-3232231013100301-1022201332320302"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present — check_not_present / 020132202302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-0201321003202122-3122222320310313-1111310222320002-0031310012211013-1013321002032321-1122222133300300-0001113130031013-3011120231100022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-2211003203321220-1123333311233003-0102222110333203-0032012103103003-1021111123021033-3321003312330101-3010203332222101-2000123312121010"></a>

## Direct properties — check_not_present / 020132202302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223323001033300-0221101323331233-3132100021022323-0300332033230121-3031011311133112-3130112001130122-1232020331230323-3212130003202302"></a>

## Next pages — check_not_present / 020132202302 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3310131023313222-0323311302301330-0330121123220032-1103002110303210-1132022103030201-1222113010022210-2003300020022001-0030231132332212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312023002230000-1301212123013213-2213332131002130-2330332330133321-0130312211010301-1200032222311223-0201112000110133-2310211133200211"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present — check_present / 332310213231 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present

<a id="canonical-2312301311303133-2211312233203211-3010313022023001-3121302021100211-2323033122003020-2012101303322103-2320011133203121-0100102020103101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-3201010122232121-3132020203230013-3131333110331303-2010013222231100-3203332323203323-0332110332013221-1222011000102210-3301010023302032"></a>

## Direct properties — check_present / 332310213231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101132012000232-1122013303200331-2102131103012221-0113020221222120-2311122330201220-1133000320021330-3023000123013211-2111112203223122"></a>

## Next pages — check_present / 332310213231 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0312023210333011-3220122310300130-1332132131103031-3011312321333103-3130103313320121-3331322220220322-3332331211132303-1113203221103111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033313212210320-3203333020320132-2030010122300033-1213110333302203-0023203122322321-0031130001302001-2010212201310330-2032031300003000"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item — item / 111010230133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item

<a id="canonical-3121021123020123-3210011130010302-1122120113033220-1312301031022020-2332111201220030-2121223000331212-1012233213313300-0132121301130312"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230033021230313-0220033013321113-1030230031020103-3013020331031013-1312301130332113-3230313203221012-2020032303210032-0003331113313121"></a>

## Direct properties — item / 111010230133 / 3

<a id="canonical-2013301233100303-0321202212312013-1202223313200212-1222111221231233-3310110013213021-3331213101013320-0021311120332032-1133303022312011"></a>

<a id="canonical-0332010231333211-0222222232323201-0230211110323133-1033203020300233-3033330021321221-0001300001230310-1011101031022131-0121113021123213"></a>

## exact_values property — item / 111010230133 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-2212032231020022-2303223201233120-1031310000201211-2213322002033031-2022113313123000-1003213232222112-3010133311023011-1011203311131303"></a>

<a id="canonical-2311110023231331-0321133131322233-1320233100101100-2011322010132032-3111131230300230-3323003013200120-0221122022200101-2220111120132203"></a>

## regex_values property — item / 111010230133 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-3010001313020230-0221212322202202-3010122100322121-1023112131312231-2311003203201013-2101231232310303-3000102301030302-3330021103030330"></a>

<a id="canonical-0221210311002322-0130230011130122-1211030111301002-0000122013212331-2311313321202121-0010000100122301-1212223223012333-1202231000030310"></a>

## transformers property — item / 111010230133 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-3030313310221232-0122133301010102-0323112133323121-0002300131231132-3121023020120011-1222002232010312-0112322112100103-1231312100130202"></a>

## Next pages — item / 111010230133 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-2211012103110202-0023203012003023-1022312311230222-1202232300332322-2031112221003312-2023202113001130-3203320111022120-2300103112320200)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310010123120233-3023330101330132-0111331213223032-3311202203113023-2000300003202100-2231031203103010-3330313123203300-2311132311210120"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params — query_params / 211202203231 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params

<a id="canonical-0230100302023013-2001113023300110-1232120030000011-2001231222132202-0233333320233231-2023000211212022-0333210033200211-1221213002030203"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332231333212033-2001302303103330-2020030131130021-1000322021113331-0213223123013001-0303311133113323-0203033000112011-2330033110300222"></a>

## Direct properties — query_params / 211202203231 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-2133331220212010-1131020000300212-0000202201222330-2102033122223221-3002330212312012-0110323033001201-2323121022032323-2121221021033131): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-008.md#canonical-1030312120322111-0033230321102320-1211023133031131-2303322232000221-3210112012331103-2210022021220122-0302033212333022-2010323221130123): complete subsection reference.

<a id="canonical-3203003323010322-1020002033023023-2321322202113113-3132100022200211-1321211123002000-3101111110133110-1111222031022300-1211310133320010"></a>

<a id="canonical-0121300013213013-0011113230031110-0100301211121223-0020323232002300-0320123003002232-1211213120331231-2002331012023103-2220033021021313"></a>

## invert_matcher property — query_params / 211202203231 / 4

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-008.md#canonical-3202300301312031-0020030211222211-0313121300121311-2031212131321320-1102131113201100-3312312301133303-3130013111111322-1113322333210333): complete subsection reference.

<a id="canonical-2032101200220213-2122331222030110-1123103312122312-1333223010320302-3211102230312002-2212102212312202-3211033301232031-2133220300000023"></a>

<a id="canonical-2231022323323303-1310101220010220-3001023303203033-1320211312321012-0000210012110121-2022230223111022-0131121202032320-3002230301001203"></a>

## key property — query_params / 211202203231 / 5

Type: `"string"`. Optional.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

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

<a id="canonical-2022202103122120-1102301212012321-0322021020122012-0020103232130103-0003313223330200-2323202113300322-1221122200120110-2332110320322133"></a>

## Next pages — query_params / 211202203231 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-2133331220212010-1131020000300212-0000202201222330-2102033122223221-3002330212312012-0110323033001201-2323121022032323-2121221021033131)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present](resources--http_loadbalancer--reference--group-008.md#canonical-1030312120322111-0033230321102320-1211023133031131-2303322232000221-3210112012331103-2210022021220122-0302033212333022-2010323221130123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item](resources--http_loadbalancer--reference--group-008.md#canonical-3202300301312031-0020030211222211-0313121300121311-2031212131321320-1102131113201100-3312312301133303-3130013111111322-1113322333210333)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2133331220212010-1131020000300212-0000202201222330-2102033122223221-3002330212312012-0110323033001201-2323121022032323-2121221021033131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113023223202123-3202101232310302-1210010223110322-0030323332332332-0233233312031020-0013122111003321-3000313213212033-0031311330022222"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present — check_not_present / 213233303121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present

<a id="canonical-3302111023300033-0212112200122102-2331123213320131-1211020210333302-1022122103313112-3230133020301312-2130201310233312-2002010033201122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-3322200111211320-2302113322123220-1320021330103223-1120231231100102-1130203330213220-0311310020213021-3013011332012301-0321132203210030"></a>

## Direct properties — check_not_present / 213233303121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023132031303320-1313022100132130-3132013113330110-1030121103112332-2310132312320321-1003222321123203-0122311310220133-0121310222130322"></a>

## Next pages — check_not_present / 213233303121 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1030312120322111-0033230321102320-1211023133031131-2303322232000221-3210112012331103-2210022021220122-0302033212333022-2010323221130123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203011203201010-3300301301320000-3032021103122302-3000302131100202-1123021121010000-0120200333231012-0321211320233100-0220220323223220"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present — check_present / 111303321311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present

<a id="canonical-3111020013202112-2320301232313300-1200333333023323-3022000210133201-0110212233122203-2010113333211221-2110333211103223-2232330120300320"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-0003233120130010-2201112011001003-1213102003231103-3210030032022133-1031323002111122-2230323113222302-1122303000022233-3133223103022212"></a>

## Direct properties — check_present / 111303321311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333200310202333-1311100023302111-0313110322003131-1131303031000322-2200122231001301-3102201232003033-3310121201301302-1101333032332103"></a>

## Next pages — check_present / 111303321311 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3202300301312031-0020030211222211-0313121300121311-2031212131321320-1102131113201100-3312312301133303-3130013111111322-1113322333210333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121031322321202-2020022110230321-3323120220132322-2300003030213010-3303330121203212-3233223002010012-1320101221101203-3132333211200310"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item — item / 320321203330 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-2003100320230200-1222231121132023-1330022221010220-0032310102121102-1300132222001312-3112123230322102-2111130022101322-0323232112212001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--http_loadbalancer--reference--group-007.md#canonical-0023003003220312-2002000013100233-3120123201133133-3202121322223032-2303010332220231-1110212100033110-0000331020111133-0211012111112033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--http_loadbalancer--reference--group-007.md#canonical-0112000320010011-0312310321032033-2100031030022022-1221130000201122-1022021202102202-2010210102011302-2203031111100133-2033330320121222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item

<a id="canonical-0222223133123003-1130123233020020-0000023113032012-3200313130032333-0111201113110210-2122212011132210-1121311003111010-1310032001200111"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-2110311333311323-0102033303323233-3200332323202023-3010020003332231-1013110000223032-1101213303023103-1032331021232033-0300200012333030"></a>

## Direct properties — item / 320321203330 / 3

<a id="canonical-0230121300110213-3132030212223321-1303011002120332-1000321232312021-0013102122030120-1231031100020223-1012222021033202-0303212031012103"></a>

<a id="canonical-0301101032323210-1320003021232211-0113232110321330-0302010333322201-1201213001210220-2321003331313230-2000232223010012-0102322123133201"></a>

## exact_values property — item / 320321203330 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-0203301102003303-2222233313330110-2121020303100303-3022031231132300-2233222131101121-3333320232033210-0302001130321323-1013210112202012"></a>

<a id="canonical-1303200102003020-2331203123100132-3013113300212233-0230000130311010-3130102222010032-3313223311210210-3330301330030323-0302022233202131"></a>

## regex_values property — item / 320321203330 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-3123010223133230-1300122323133003-0013230020002003-0023233312300103-2031102121030031-1320122032010100-1201131211333310-0112332100120122"></a>

<a id="canonical-1222111130032131-1310231120233210-1320230010020000-3230201021102101-2003321312112220-1111232103232322-2100213213002310-1312320020233031"></a>

## transformers property — item / 320321203330 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-2102311202010300-1203003001221233-1310131303132303-3023303031213122-3100110220330102-0301122002320101-0303010313223210-0121232313033233"></a>

## Next pages — item / 320321203330 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-008.md#canonical-1310331230222321-3231200203303030-2001032333200123-0022233322301002-1113012002022322-2112001023312000-2231010030301023-2121101033110011)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2313230220001103-3212301300010212-3013123133132033-1130203100112210-1121313032010300-0313200120132330-2230131221332112-1011333232233322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200320100220220-3000020002333123-0132311322310213-3023310102223002-1020112202001332-2023201203201221-3302322111021111-2203222133120231"></a>

## api_rate_limit.custom_ip_allowed_list — custom_ip_allowed_list / 212302032130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- api_rate_limit.custom_ip_allowed_list

<a id="canonical-0223321212212223-3323033301230012-1310121200013311-0223303033003103-0110303301020330-1213220233123013-2211330312003211-1321022110131220"></a>

Type: `"object"`. single nested block, Optional.

IP Allowed list using existing ip\_prefix\_set objects.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate_limiter_allowed_prefixes")}
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
custom_ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330102323320121-2331100320203021-2312303231110233-3102323312002321-3210311001121221-1010312202302130-1212002102021133-1030022102100320"></a>

## Direct properties — custom_ip_allowed_list / 212302032130 / 3

- [rate_limiter_allowed_prefixes](resources--http_loadbalancer--reference--group-008.md#canonical-1010211010113121-0331310121113033-2030031032022330-1021322001300232-3231010303030310-2212033322222311-3023322011123320-0021202201302200): complete subsection reference.

<a id="canonical-2230220232231212-1200032100211122-3300033211012201-3100123032100201-2133231222012102-0030303102020122-2102331303313112-2212112313100332"></a>

## Next pages — custom_ip_allowed_list / 212302032130 / 4

- [api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](resources--http_loadbalancer--reference--group-008.md#canonical-1010211010113121-0331310121113033-2030031032022330-1021322001300232-3231010303030310-2212033322222311-3023322011123320-0021202201302200)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1010211010113121-0331310121113033-2030031032022330-1021322001300232-3231010303030310-2212033322222311-3023322011123320-0021202201302200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303103133331003-2133200100121031-0002330000023200-2112110331233230-2011012221202320-3123311201102021-1223101003010222-2003332032332033"></a>

## api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes — rate_limiter_allowed_prefixes / 033130130023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.custom_ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-2313230220001103-3212301300010212-3013123133132033-1130203100112210-1121313032010300-0313200120132330-2230131221332112-1011333232233322)
- api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-3312110313030210-2131230011132231-1010202033133332-2100210220332002-1030303211221312-3131213103012332-1332122223133213-3111011131200210"></a>

Type: `"object"`. list nested block, Optional.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Upstream description:

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
rate_limiter_allowed_prefixes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1221012113222303-3201020122310122-3333102002321300-2303303323120221-1031311000013001-3120033023031112-2031132100013322-2200032030320031"></a>

## Direct properties — rate_limiter_allowed_prefixes / 033130130023 / 3

<a id="canonical-2322113020233302-1013001312301010-2113201112300003-3203222031231200-1322011123103303-2131223202002312-1023332313300013-3121021301010222"></a>

<a id="canonical-3122103133201110-3210122001202300-2321200032112321-2022211323123200-1330210023330131-1200300213122123-0232222220102110-0203330122211312"></a>

## name property — rate_limiter_allowed_prefixes / 033130130023 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-2010002123333030-2112221213223230-3301202303322002-3202012203231101-0301230203330010-1130110122130113-1133213201211232-0100321010101033"></a>

<a id="canonical-2031111122300230-1312010030112330-3121222222233120-1202100202122322-2010023320333011-3013333212300221-3212110211021213-0113100131003101"></a>

## namespace property — rate_limiter_allowed_prefixes / 033130130023 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-2003321330223331-1021132032302121-0212011123323201-0301312210112130-2223032022102233-0103120320131011-1233100223022130-0021131220121023"></a>

<a id="canonical-3312201001210233-2302031021131311-3313133032231221-0020330022101122-3020321313200300-0102002100303013-3010132111103310-0102131200222010"></a>

## tenant property — rate_limiter_allowed_prefixes / 033130130023 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-2000030103323100-3220022300310323-2113333023220022-0130000213010023-0000303130110122-3112012011130200-2232230113113223-1033213220111122"></a>

## Next pages — rate_limiter_allowed_prefixes / 033130130023 / 7

- [api_rate_limit.custom_ip_allowed_list](resources--http_loadbalancer--reference--group-008.md#canonical-2313230220001103-3212301300010212-3013123133132033-1130203100112210-1121313032010300-0313200120132330-2230131221332112-1011333232233322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1201210302313203-2211300232121320-0323321131002323-0331122322021300-2230032123313000-3330020132002300-3030111322130233-3130310012323300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031302332013223-3020113112310220-2112211011333322-2031221020230233-1200212113131013-0102011211101103-0021130332223120-0312011101300302"></a>

## api_rate_limit.ip_allowed_list — ip_allowed_list / 223223321133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- api_rate_limit.ip_allowed_list

<a id="canonical-3012333312322230-1212020221132021-3111130330021031-2202011012313020-0113002210323312-0332123301131333-3223100320121322-0320222232023001"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_allowed_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013003001133102-2002133130131310-0300101130231320-1321011120311320-3122320223032322-1210303020300112-2110301102322213-2313100020331001"></a>

## Direct properties — ip_allowed_list / 223223321133 / 3

<a id="canonical-3132130030310123-2122120100331231-0012220001232330-2003022220321111-2010222123112213-1211203331333302-0232313102301213-1210200030330213"></a>

<a id="canonical-0213031202230222-1031221031010212-0330202202331221-0323021323222333-2000221130320232-2123321133011323-3303220022012220-0133121022232330"></a>

## prefixes property — ip_allowed_list / 223223321133 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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

<a id="canonical-3300311331010200-2112302303232210-0131232122001322-0203021302033202-2210313202132311-3103103001120232-2120200201121330-3122202200023102"></a>

## Next pages — ip_allowed_list / 223223321133 / 5

- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1030302110202203-3022230330030230-2123232230310122-2013030321120121-3102121110303121-1120312201020213-1322321131223223-1130233033033321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023021023310123-2302210131321010-0232200332221230-2123111213200113-3111220030332102-1311121333202211-3031200322320111-1323001322102113"></a>

## api_rate_limit.no_ip_allowed_list — no_ip_allowed_list / 301213010223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- api_rate_limit.no_ip_allowed_list

<a id="canonical-2001212101100010-1031331233322320-3300321322323011-1313221031220113-1111030210022001-0302201211302001-3330310202332003-0222110203332211"></a>

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
no_ip_allowed_list = {}
```

<a id="canonical-3021301003021102-0031110333310113-0200110110210202-3320121200013212-1012232033112220-2131202322021132-0220202023101111-2003031310222221"></a>

## Direct properties — no_ip_allowed_list / 301213010223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002310123303212-3031030033230222-3012213213020213-3320203010301021-2300322121333102-2110210010202111-3002210223003013-0200112320123203"></a>

## Next pages — no_ip_allowed_list / 301213010223 / 4

- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112131322133112-0311110030232200-1021300200112010-2330333122100033-3031312222202220-3022030033223130-2132011301000303-1220331020023103"></a>

## api_rate_limit.server_url_rules — server_url_rules / 201120223113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- api_rate_limit.server_url_rules

<a id="canonical-1022322002012112-1333230230322312-3130222323122203-1230312210223231-0311012301032012-0233333131213112-2003111123210302-2110230030012100"></a>

Type: `"object"`. list nested block, Optional.

Ordered domain or base-path rules for path-scoped rate limiting. Each rule must choose exactly one
rate\_limiter\_choice: inline\_rate\_limiter or ref\_rate\_limiter.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("base_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("inline_rate_limiter",
    "ref_rate_limiter")}
```

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
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

Terraform syntax:

```terraform
server_url_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102221301233332-3122111213333202-1301333233031233-2111301223312233-2300023220300223-0112103100313132-3331333130233303-1010031203023010"></a>

## Direct properties — server_url_rules / 201120223113 / 3

- [any_domain](resources--http_loadbalancer--reference--group-008.md#canonical-0212023132210021-3231013310023330-3332013202121221-0002003101003132-3121122302111200-3213110110213330-0033110231020100-3113033032103031): complete subsection reference.

<a id="canonical-1020332131331132-2130131300220002-1101010231213131-3310130032300300-0323220100032300-1313130213020322-1323302210203012-0031102311221132"></a>

<a id="canonical-0310203201000021-1213211033312020-3301322020001121-2012013213032102-3221101113112011-1332102311331001-2033233132301313-3301200032002312"></a>

## api_group property — server_url_rules / 201120223113 / 4

Type: `"string"`. Optional.

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with 'x-F5 Distributed..

Upstream description:

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with "x-F5 Distributed
Cloud-API-group" extensions inside swaggers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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

<a id="canonical-3313310122333011-3123110332133223-2120011013302313-0322101133033012-3311101223103210-0232223202233321-0003213132330233-1233102131312033"></a>

<a id="canonical-2233202221011322-2102200013200112-0323012231111112-2010031032003213-2011023111300112-1013111223110233-3323322232010310-2310321333231130"></a>

## base_path property — server_url_rules / 201120223113 / 5

Type: `"string"`. Optional.

Base Path. Prefix of the request path.

Upstream description:

Prefix of the request path.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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

- [client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311): complete subsection reference.

- [inline_rate_limiter](resources--http_loadbalancer--reference--group-008.md#canonical-3230011100300201-3021322232013132-2120130113220102-3000312330300303-2132212200020332-2301101031230211-2123121301333223-1211301231000212): complete subsection reference.

- [ref_rate_limiter](resources--http_loadbalancer--reference--group-008.md#canonical-1003023331222033-2221012320323333-0303300312100321-0001000033001020-3133220003001112-2131311032211230-1003011113023110-0101320100203112): complete subsection reference.

- [request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031): complete subsection reference.

<a id="canonical-3123333301133020-2200200331133310-1311231301003320-2131122103021110-3333210222022322-1122103200201113-1303003120101101-1203120330230103"></a>

<a id="canonical-1013112223331003-3203201123011313-2222212101221200-3002312300111130-3112133301230113-2000322113301331-0122233011213113-3311122232201312"></a>

## specific_domain property — server_url_rules / 201120223113 / 6

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-2201301131022023-2332021323222312-0311011223201213-0301013210330331-1303011022132132-1323002121231211-0233032110032113-1300312123022202"></a>

## Next pages — server_url_rules / 201120223113 / 7

- [api_rate_limit.server_url_rules.any_domain](resources--http_loadbalancer--reference--group-008.md#canonical-0212023132210021-3231013310023330-3332013202121221-0002003101003132-3121122302111200-3213110110213330-0033110231020100-3113033032103031)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-008.md#canonical-3230011100300201-3021322232013132-2120130113220102-3000312330300303-2132212200020332-2301101031230211-2123121301333223-1211301231000212)
- [api_rate_limit.server_url_rules.ref_rate_limiter](resources--http_loadbalancer--reference--group-008.md#canonical-1003023331222033-2221012320323333-0303300312100321-0001000033001020-3133220003001112-2131311032211230-1003011113023110-0101320100203112)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0212023132210021-3231013310023330-3332013202121221-0002003101003132-3121122302111200-3213110110213330-0033110231020100-3113033032103031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212230023223110-3031030210213132-1011113010323310-2011313213101300-2203201033220013-2300231000300112-0120010023222311-2202002031211010"></a>

## api_rate_limit.server_url_rules.any_domain — any_domain / 222233002120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- api_rate_limit.server_url_rules.any_domain

<a id="canonical-3203321320012003-0011220012131130-0320133200230300-2002312213101131-2231110302021100-0213113302100120-1311033322302232-2120300033223120"></a>

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
any_domain = {}
```

<a id="canonical-3130122300321200-0312303022200131-3131113230231002-2222113202310302-3221120031200220-0310233321030321-0311023123032132-3233211212200332"></a>

## Direct properties — any_domain / 222233002120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133113111123311-1020311132220312-0002322110033310-2003001100003110-1233020201103332-0001113311032330-3221312220231231-0201213311023122"></a>

## Next pages — any_domain / 222233002120 / 4

- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312332203323103-0332211132333313-3332102231123213-3201300332223031-3210202311012210-0103132102010221-1331333000301313-2031100203122203"></a>

## api_rate_limit.server_url_rules.client_matcher — client_matcher / 130310120231 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- api_rate_limit.server_url_rules.client_matcher

<a id="canonical-1213112212110323-3222300213131031-0221222122121033-0112211103103323-0320120313200300-3330311131333331-1211102100003010-3002213121331023"></a>

Type: `"object"`. single nested block, Optional.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_client",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_list",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("asn_matcher",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("client_selector",
    "ip_threat_category_list"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
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
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\",\"ip_threat_category_list\"]",
  "x-ves-oneof-field-ip_asn_choice": "[\"any_ip\",\"asn_list\",\"asn_matcher\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

Terraform syntax:

```terraform
client_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2111221032120213-0132031310212323-1200033313203222-3313212211330101-2033212230010301-0323013202102131-2032112113200002-2020232012131312"></a>

## Direct properties — client_matcher / 130310120231 / 3

- [any_client](resources--http_loadbalancer--reference--group-008.md#canonical-3201313120222312-3232003221302110-0311130020311312-0301122020312122-3101232211013300-2031210131210003-3200321013030133-0011010102222102): complete subsection reference.

- [any_ip](resources--http_loadbalancer--reference--group-008.md#canonical-0022011202103021-0333123310221021-2132010333112323-3312330330102330-3013211101220331-1222231010222202-2200120320010011-0210121103221302): complete subsection reference.

- [asn_list](resources--http_loadbalancer--reference--group-008.md#canonical-1133033031101131-3302230123123022-1122323011022101-0320331123222100-0301010130230202-3013302033301123-3330213033221301-0203302213331121): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-2000002131130203-0313323223312023-2132220302130203-3322322103201031-1223203023121310-2021010311321311-0130130232033020-2200211201222213): complete subsection reference.

- [client_selector](resources--http_loadbalancer--reference--group-008.md#canonical-3221232222331013-2013333101232303-1211121213320301-0100113132313033-2331203203133003-2031021202220201-1321001103020332-3123212023212212): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1201312112023220-3110102133111300-3200210011301312-1102221323231012-3201330200312020-0221113222213323-0212313312200010-1101310120100232): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--reference--group-008.md#canonical-2203121131301233-2203201100322000-2220212301120232-2023030221110012-2001211200121321-1311112302313220-1333303302223032-3033032303332202): complete subsection reference.

- [ip_threat_category_list](resources--http_loadbalancer--reference--group-008.md#canonical-1231330211322320-1122320030111011-1330301312012202-2310122321121110-0212323111010212-3313202200230031-0222223021333031-2130231112220022): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-2111102223312220-3103020330310103-3100313013131110-0322313121022222-0030312201003130-0023132031031331-1201122310213220-3210012002132112): complete subsection reference.

<a id="canonical-3133012310321132-3302313101000120-0221002101133323-2303000002113301-3200030131111201-1321202133333320-1202033321220010-0113300311213121"></a>

## Next pages — client_matcher / 130310120231 / 4

- [api_rate_limit.server_url_rules.client_matcher.any_client](resources--http_loadbalancer--reference--group-008.md#canonical-3201313120222312-3232003221302110-0311130020311312-0301122020312122-3101232211013300-2031210131210003-3200321013030133-0011010102222102)
- [api_rate_limit.server_url_rules.client_matcher.any_ip](resources--http_loadbalancer--reference--group-008.md#canonical-0022011202103021-0333123310221021-2132010333112323-3312330330102330-3013211101220331-1222231010222202-2200120320010011-0210121103221302)
- [api_rate_limit.server_url_rules.client_matcher.asn_list](resources--http_loadbalancer--reference--group-008.md#canonical-1133033031101131-3302230123123022-1122323011022101-0320331123222100-0301010130230202-3013302033301123-3330213033221301-0203302213331121)
- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-2000002131130203-0313323223312023-2132220302130203-3322322103201031-1223203023121310-2021010311321311-0130130232033020-2200211201222213)
- [api_rate_limit.server_url_rules.client_matcher.client_selector](resources--http_loadbalancer--reference--group-008.md#canonical-3221232222331013-2013333101232303-1211121213320301-0100113132313033-2331203203133003-2031021202220201-1321001103020332-3123212023212212)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1201312112023220-3110102133111300-3200210011301312-1102221323231012-3201330200312020-0221113222213323-0212313312200010-1101310120100232)
- [api_rate_limit.server_url_rules.client_matcher.ip_prefix_list](resources--http_loadbalancer--reference--group-008.md#canonical-2203121131301233-2203201100322000-2220212301120232-2023030221110012-2001211200121321-1311112302313220-1333303302223032-3033032303332202)
- [api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list](resources--http_loadbalancer--reference--group-008.md#canonical-1231330211322320-1122320030111011-1330301312012202-2310122321121110-0212323111010212-3313202200230031-0222223021333031-2130231112220022)
- [api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-2111102223312220-3103020330310103-3100313013131110-0322313121022222-0030312201003130-0023132031031331-1201122310213220-3210012002132112)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3201313120222312-3232003221302110-0311130020311312-0301122020312122-3101232211013300-2031210131210003-3200321013030133-0011010102222102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232130203122232-2221111002231223-0130101321312033-3210232022110330-3012001000032231-0311011233000021-0202223231213100-0010030221212030"></a>

## api_rate_limit.server_url_rules.client_matcher.any_client — any_client / 102033301333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.any_client

<a id="canonical-0113332230032013-2132020102112333-0000330033003122-1201230011312323-3230013130322012-1222202113231122-1311330130123322-1303302233213011"></a>

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
any_client = {}
```

<a id="canonical-2111020111201131-1001021101023221-1230120020302002-2111111321220112-3122333022131002-0031211232310330-3212122232102003-0223300111321212"></a>

## Direct properties — any_client / 102033301333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222000333320321-0003130210230310-3301033300321012-2330211323310021-3201313202303233-2100131030332321-1110222320120321-1211001323033211"></a>

## Next pages — any_client / 102033301333 / 4

- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0022011202103021-0333123310221021-2132010333112323-3312330330102330-3013211101220331-1222231010222202-2200120320010011-0210121103221302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201100011021320-2313330331011100-1230122101223311-2000020101103322-1122221203130322-3110230331121303-3103301031210113-1230221223313233"></a>

## api_rate_limit.server_url_rules.client_matcher.any_ip — any_ip / 102133330300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.any_ip

<a id="canonical-3033220100120313-1021031312011213-3132233223122020-1220231332221022-1001301131003111-1203111020100123-2131230203133231-0001313213220130"></a>

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
any_ip = {}
```

<a id="canonical-3033121132132233-1223023210031102-0021221222031312-2130121202330203-0320122132120221-2113310131333211-0321212032230331-3311003323300321"></a>

## Direct properties — any_ip / 102133330300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230021230021211-3220110323112011-0121213221100213-2311032102213103-2303213010003102-2202020013031122-0301000201320132-2331132100203102"></a>

## Next pages — any_ip / 102133330300 / 4

- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1133033031101131-3302230123123022-1122323011022101-0320331123222100-0301010130230202-3013302033301123-3330213033221301-0203302213331121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232321231030302-0300011322223102-2330033130000221-0101333330123131-0133303020223223-2210313203030020-3331321230312230-0330333232220230"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_list — asn_list / 023111223133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.asn_list

<a id="canonical-0320200130002221-3232133203102002-1012020302001133-1333331003232230-0110100103332022-2330320001203010-3020023220211123-3312001310221220"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1321301003210322-1031101322223013-1022120220113010-1012301232330022-2001212003131120-2220330330123032-0302120030310230-3003221002011031"></a>

## Direct properties — asn_list / 023111223133 / 3

<a id="canonical-2122323111300010-1312122213330010-3300301111320010-1000100002313330-1201302113302133-0200312333303313-1020320330220310-2200220111222020"></a>

<a id="canonical-2103301031032302-3333122331031121-2301221032133121-1023132311312303-2020132112320022-2101213210100323-3120232302001032-2130120133133033"></a>

## as_numbers property — asn_list / 023111223133 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3301331101012132-3122321313101322-1120233011203010-1230033202303001-3231012203130122-1210032310122333-3301213323121023-1210033110310332"></a>

## Next pages — asn_list / 023111223133 / 5

- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2000002131130203-0313323223312023-2132220302130203-3322322103201031-1223203023121310-2021010311321311-0130130232033020-2200211201222213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000011323330333-1121210113221031-2222223223221101-1032212333300023-1333333032202000-0322130220212330-3022023110131303-2302222332012321"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_matcher — asn_matcher / 120230101333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher

<a id="canonical-0133300223222310-3001330111103003-2202313203133321-1332000101230212-3101102232021031-3220121021201032-1112033233233122-2202132031131131"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131231000212032-3120313321211131-2313011000103001-0310230320202300-1323122131230322-3312111011000301-1022000311103031-3222131233223022"></a>

## Direct properties — asn_matcher / 120230101333 / 3

- [asn_sets](resources--http_loadbalancer--reference--group-008.md#canonical-3222322102130103-0201110333302231-0120331323202112-3203033200323230-2003032231301100-3103302210231300-0200020130200120-0223211300331131): complete subsection reference.

<a id="canonical-0022020123301213-2130230002102231-1033222022033030-3203331322022320-2000320300131002-1321023322221223-0122003122232222-3033022030200030"></a>

## Next pages — asn_matcher / 120230101333 / 4

- [api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets](resources--http_loadbalancer--reference--group-008.md#canonical-3222322102130103-0201110333302231-0120331323202112-3203033200323230-2003032231301100-3103302210231300-0200020130200120-0223211300331131)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3222322102130103-0201110333302231-0120331323202112-3203033200323230-2003032231301100-3103302210231300-0200020130200120-0223211300331131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011320231333032-0322001012130200-3113121331213133-0022213310201132-2130320033012030-3103200013303131-3302303031131322-2020111321233221"></a>

## api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets — asn_sets / 332132223332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-2000002131130203-0313323223312023-2132220302130203-3322322103201031-1223203023121310-2021010311321311-0130130232033020-2200211201222213)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-1230331022312210-0322033313310333-0310022113022113-3222023131210233-0022323332330200-1303320233330010-2300221321301212-3021020232013302"></a>

Type: `"object"`. list nested block, Optional.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-1321133330232221-0331131131021212-3033121200333221-2302322020131131-2020031203230210-3222201021331233-0123111220321033-3310233332220003"></a>

## Direct properties — asn_sets / 332132223332 / 3

<a id="canonical-1312212220211210-2331111212222132-0201201003111211-1322312000131232-0332230131013301-2331033231023012-2023121231012201-1031232001133332"></a>

<a id="canonical-2012320300023101-1021101100300122-2133220113203011-0203100311011201-2001020013010020-1321002020112101-3033023003210023-1123020010021301"></a>

## kind property — asn_sets / 332132223332 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-1131123223322103-0230331321222133-2101213231123120-3233001012320012-0221032020203101-2011323011300003-3032323121203321-2131330110002122"></a>

<a id="canonical-3331023322211023-0323322231222001-0301123113033333-1012233002221322-0201122002300310-0121033133331211-0020013331031010-3311103011001201"></a>

## name property — asn_sets / 332132223332 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2310232113002202-2120301332133300-0020303002231010-2130230020021000-1311012102110110-0220300333302331-3321121021320101-2033331020322103"></a>

<a id="canonical-3300202002212300-2123100002302102-2011031032031120-2022303022010120-2131311213101310-1220103300131133-2301110322111313-3023320212230303"></a>

## namespace property — asn_sets / 332132223332 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-2333302222000021-0231032110002113-1102013023100000-0023021221031310-0323000033122333-0120132020013130-3103302231012013-2011113003112300"></a>

<a id="canonical-1132001221213311-0332130332303031-0133221012101221-0213103002101330-2103121310130031-0031321210032222-2132323100213100-3223300131020002"></a>

## tenant property — asn_sets / 332132223332 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3313113223130113-3121020310333021-0322110302112030-0323203123222233-1321202113222010-2222201231232123-1023211111301301-3220202221123013"></a>

<a id="canonical-0013313121311330-1202322330130320-1321223222223122-1121223101301122-2312113323020011-2302223023222003-3210030031011220-3112021112132103"></a>

## uid property — asn_sets / 332132223332 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-3110120222132233-3120311021202313-2302123320001331-0121011231030000-2213302313130211-2131313303320310-2212203003220113-1032100032210101"></a>

## Next pages — asn_sets / 332132223332 / 9

- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-2000002131130203-0313323223312023-2132220302130203-3322322103201031-1223203023121310-2021010311321311-0130130232033020-2200211201222213)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3221232222331013-2013333101232303-1211121213320301-0100113132313033-2331203203133003-2031021202220201-1321001103020332-3123212023212212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322123231023210-1321011332213300-0002011320313012-0002110113003200-0000120132122012-1100212313113201-3121032131113220-2123302200002120"></a>

## api_rate_limit.server_url_rules.client_matcher.client_selector — client_selector / 132130201320 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.client_selector

<a id="canonical-3232230102200303-3123112230222022-1130013331131021-3210102111103232-0021031112300232-0100303303233310-1210132012003233-2011202330320132"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-2101200233130223-0132030123212120-3011011333131023-2220303230313302-0303001133102113-0203202120112333-1212313310311022-1203130232130112"></a>

## Direct properties — client_selector / 132130201320 / 3

<a id="canonical-1030130202023322-0133220102202222-0220221223023321-2333223131220231-0332123100301103-2001322201032123-1201023111321001-2032201213123303"></a>

<a id="canonical-3130331230300203-2300313011000130-1221331312203033-0013003331123122-0011123123230230-1300002230102301-3231303110121013-3221232103210302"></a>

## expressions property — client_selector / 132130201320 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2222021201131201-0202331311201231-2132010232231310-3132120110123030-0223202121101101-2101231012110002-1210031202211220-0131231022111100"></a>

## Next pages — client_selector / 132130201320 / 5

- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1201312112023220-3110102133111300-3200210011301312-1102221323231012-3201330200312020-0221113222213323-0212313312200010-1101310120100232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032322121131322-3302103111100123-3120030131100202-0103200232100232-3131313022111322-3330103203123000-2221013202030110-0212033203323033"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_matcher — ip_matcher / 302333222102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher

<a id="canonical-3111212012331101-2220200112330133-2012303132011311-1222132200202101-1013223222220023-0031312030321003-0101320132032022-0111000212330212"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022220310331302-2211300103233110-0122013011302200-1001331201230333-0213321310021032-2000211232220010-1331122120031211-0212112130322231"></a>

## Direct properties — ip_matcher / 302333222102 / 3

<a id="canonical-0221222320013222-3110220123202203-1212023310030212-3133331231101111-3101013201222001-1120310200301123-2111000022311123-3222222132110312"></a>

<a id="canonical-3121233022302023-1003133111022130-1301303222023312-2011013322323112-1211301212323031-0000030313231010-2320210203133102-2213312131113233"></a>

## invert_matcher property — ip_matcher / 302333222102 / 4

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](resources--http_loadbalancer--reference--group-008.md#canonical-1111220221122301-1130033211133132-2103221030323230-1001320030112000-3032323021002323-0003122231221330-3000302230021030-3032112232201030): complete subsection reference.

<a id="canonical-2230031102020333-2331210223231020-3312012110211101-0132201202333032-1012033112302232-0122032020132011-0003122301120233-3202213312111333"></a>

## Next pages — ip_matcher / 302333222102 / 5

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets](resources--http_loadbalancer--reference--group-008.md#canonical-1111220221122301-1130033211133132-2103221030323230-1001320030112000-3032323021002323-0003122231221330-3000302230021030-3032112232201030)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1111220221122301-1130033211133132-2103221030323230-1001320030112000-3032323021002323-0003122231221330-3000302230021030-3032112232201030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121000101331203-0131123312032203-2102201113313023-0022110333311010-0323233010031032-2031130103312123-1310300011002231-2023110202212012"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets — prefix_sets / 013222022131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1201312112023220-3110102133111300-3200210011301312-1102221323231012-3201330200312020-0221113222213323-0212313312200010-1101310120100232)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-3011123330223303-0302302202320101-1023202031313100-0312132113311323-0330333030121113-2333130210002230-1101130122301220-1232233030202022"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0101221021002103-0001013112301113-0010133120212101-2302111331020012-2012311313231221-2222301002020331-0020132303213213-1102220332030201"></a>

## Direct properties — prefix_sets / 013222022131 / 3

<a id="canonical-0012121313331232-1300010123121220-2231120121223123-2302000213120323-0101112202221232-2332013310030320-2130110123102233-1020322233121220"></a>

<a id="canonical-2100220100301221-2011201321030010-3023232013222023-3011012203310032-1213332213103101-2231313220020332-1303233222102103-0210031113303003"></a>

## kind property — prefix_sets / 013222022131 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-3331201001110013-3002200012002313-2213010030211232-1322310321303213-1100311233100121-3103130210210100-1121033300002123-2102123110311012"></a>

<a id="canonical-3303323222213301-1033122120133303-0320023002111032-0332101223320021-2102020011330120-2331103023200133-3131202313021121-0022023110120003"></a>

## name property — prefix_sets / 013222022131 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3323313213312033-1002303132312030-0201231001111203-2311012333100110-1030223301113331-0333231101113133-1132210203121030-1003032133000231"></a>

<a id="canonical-1311122232303323-0303200001033020-3101300202110011-3102102200231220-1232131233200033-3212103331320100-0032210030103220-3012321311321231"></a>

## namespace property — prefix_sets / 013222022131 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-1032001322113122-3303111022131033-3221112002332213-1123213132211313-3222022323020032-2130321010121122-0130010201201230-0032001330113112"></a>

<a id="canonical-2321313131210233-1103221013202220-3232200133222201-3311032231001331-3102010023130231-3133312313131032-2032330000311103-1123012102030332"></a>

## tenant property — prefix_sets / 013222022131 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-0203123320301120-0103103103012013-3112131320213021-3221011133210131-3333323333222322-0333301101321332-2233331231322212-1130333311100112"></a>

<a id="canonical-1202111212212313-1112212201033333-3111033113130103-0312203031331133-3230000313030332-1222130021031200-3130201133212332-0123023033213012"></a>

## uid property — prefix_sets / 013222022131 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-2221222032001113-1032121212211001-0022121131333102-1111301010111200-0300211200102031-2212320121000322-2112020000202212-2123212032332330"></a>

## Next pages — prefix_sets / 013222022131 / 9

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1201312112023220-3110102133111300-3200210011301312-1102221323231012-3201330200312020-0221113222213323-0212313312200010-1101310120100232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2203121131301233-2203201100322000-2220212301120232-2023030221110012-2001211200121321-1311112302313220-1333303302223032-3033032303332202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332012031332212-3030233312221321-1323031000220002-3313320321220303-3130212303101132-2002311203020203-1002021311101032-0302223111022013"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_prefix_list — ip_prefix_list / 322000001303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.ip_prefix_list

<a id="canonical-3023122011132231-2303000011023013-2311321220323022-3103122213212101-2021301031003213-0122333213301311-2012003121001202-1123210003123220"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302023211303330-0032300230321223-1121312232023221-0210022113001112-2331002223233302-1011030033012210-2223323113231130-1111000320010331"></a>

## Direct properties — ip_prefix_list / 322000001303 / 3

<a id="canonical-1103023001203200-0231123000031001-0323130230222031-3021323012201231-1221011203223221-1010020220203222-3321130022131313-0221332312123003"></a>

<a id="canonical-2122321022230021-1110130200300320-0223333030021022-0322120132233032-3202121231100202-0022101120031020-0220002102001300-3312201331323200"></a>

## invert_match property — ip_prefix_list / 322000001303 / 4

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-3331332102210030-1200100321010211-2210332202220101-2311011212332013-3330210112130232-3220203022110102-3033231210213312-3303031032122001"></a>

<a id="canonical-0213133001031213-2200131023110302-2231021220200022-2202022003100131-3203213023030113-3222123202331113-2323021322120023-0223021301123020"></a>

## ip_prefixes property — ip_prefix_list / 322000001303 / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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

<a id="canonical-0312133323221331-1231320320131212-1021320122122311-1132103201021202-0323203222130322-0223130312302322-1012103222300302-0220030312132130"></a>

## Next pages — ip_prefix_list / 322000001303 / 6

- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1231330211322320-1122320030111011-1330301312012202-2310122321121110-0212323111010212-3313202200230031-0222223021333031-2130231112220022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310003031232313-3231303112323000-3300220122223001-1022310131313012-0332000020030320-0322331003131210-2113033310220320-2331030210012011"></a>

## api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list — ip_threat_category_list / 220020321000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list

<a id="canonical-1300113312000132-0101322103330301-1102000231030003-1022001211030211-3332113032100303-1311330112310013-2010113001021300-3103022002020330"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
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
ip_threat_category_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020301032321123-1211033210011322-1030330123130032-1330031031130120-1301101110330313-3231100100001223-0330002110320232-3131122011132113"></a>

## Direct properties — ip_threat_category_list / 220020321000 / 3

<a id="canonical-0022330011031210-2212330012212230-0322000330131301-1331201102133223-2233131200310020-0001323103333010-1000203222100212-1300330023011011"></a>

<a id="canonical-0202102031233013-2002020121303001-1323221232133303-3303111033121021-0012202023300123-3031312202033321-1121121220021133-0233100112302202"></a>

## ip_threat_categories property — ip_threat_category_list / 220020321000 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`,
\`WEB\_ATTACKS\`, \`BOTNETS\`, \`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`,
\`MOBILE\_THREATS\`, \`TOR\_PROXY\`, \`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to
\`SPAM\_SOURCES\`.

Upstream description:

The IP threat categories is obtained from the list and is used to auto-generate equivalent label
selection expressions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

<a id="canonical-2201023321011323-3023321022112132-3122230321002123-3031033110010322-3220031112112011-3120303010230131-1312233133020333-2113013110313212"></a>

## Next pages — ip_threat_category_list / 220020321000 / 5

- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2111102223312220-3103020330310103-3100313013131110-0322313121022222-0030312201003130-0023132031031331-1201122310213220-3210012002132112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312213303122012-3031000210201033-1102100302011122-2022133330011131-3333333220033302-3123122131330112-2133021213122133-0302302203212311"></a>

## api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher — tls_fingerprint_matcher / 102310200302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-1302113200032023-0310000120211211-3202002030113213-3323222032222031-2012001310021102-2210022030013010-3311133201210033-2133100121200210"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3021013001202121-0021013213301130-3302301021101200-3123230220321022-0132103131232323-0031011132200212-2011331303011020-2220220232033133"></a>

## Direct properties — tls_fingerprint_matcher / 102310200302 / 3

<a id="canonical-3100012232313232-2211133211020333-1200020222202210-1033300020300213-0032230321201101-0310010131022221-0310122323103213-1110023132233002"></a>

<a id="canonical-3130130200013031-1212303033021112-0203313120011130-0220312100333030-0011120203031032-2020202320002122-2112120203123313-1021001203313320"></a>

## classes property — tls_fingerprint_matcher / 102310200302 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-0222313131000022-2221120112303230-3230121022013332-2212220031101230-1221111210211221-1220020211211101-2212222033120021-3230002031011301"></a>

<a id="canonical-2111303112233012-2203221120230320-0203033112011321-3001313031032221-2233001232032003-0303130231303302-1103102230311130-0323200031021221"></a>

## exact_values property — tls_fingerprint_matcher / 102310200302 / 5

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-0131112101211123-3130302033220120-2322002312003322-3133203201133322-1331102123212020-3201021301220113-0131200211120222-0310220022200310"></a>

<a id="canonical-2031113313230123-1333113131220303-3022103321332010-3213103301032333-0133220001231020-0030103213231110-2211132210020113-1232113303312032"></a>

## excluded_values property — tls_fingerprint_matcher / 102310200302 / 6

Type: `["list", "string"]`. Optional.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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

<a id="canonical-3211020323322313-1131212010202221-3233303120131011-1000203002231110-1201330311120302-1021021020031121-0133200103303213-2101223331002332"></a>

## Next pages — tls_fingerprint_matcher / 102310200302 / 7

- [api_rate_limit.server_url_rules.client_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-1301312330333201-2203020312210332-1031301330312303-2030313133000330-0233022211030020-0023110102221033-1313001333300213-3210113302213311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3230011100300201-3021322232013132-2120130113220102-3000312330300303-2132212200020332-2301101031230211-2123121301333223-1211301231000212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001033333231212-0022330112101212-1302110131212132-0102132200210120-0032211232313221-0203111003220000-2101202313233323-2333201010221300"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter — inline_rate_limiter / 302113133021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- api_rate_limit.server_url_rules.inline_rate_limiter

<a id="canonical-1313000212311313-1331312031123311-3011103110321300-3001311203110301-0021012110133321-3032013102022222-3212102320003202-1131330031221133"></a>

Type: `"object"`. single nested block, Optional.

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("threshold"),
  validators.ConflictingObjectAttributes("ref_user_id",
    "use_http_lb_user_id")}
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
  "x-ves-oneof-field-count_by_choice": "[\"ref_user_id\",\"use_http_lb_user_id\"]"
}
```

Terraform syntax:

```terraform
inline_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-1301011110030100-0323320230201111-1122210311122133-2101230013221001-0133012223300211-0011013301210232-3332020200132310-2223322120121031"></a>

## Direct properties — inline_rate_limiter / 302113133021 / 3

- [ref_user_id](resources--http_loadbalancer--reference--group-008.md#canonical-3122110102021131-0210011132022301-2023111000333232-2031133321123202-0211131222132230-3013231033012330-0331030100133220-0330233330001203): complete subsection reference.

<a id="canonical-2122101101102120-3010332123211200-0332120012001022-2123300212313133-2013233021223022-0030300101001223-2023201313103003-0033200000211302"></a>

<a id="canonical-3212212233202330-3020312121312300-3330211112033003-0121233220003230-3310303002320310-3103320131030210-0333031020102302-0001103202112103"></a>

## threshold property — inline_rate_limiter / 302113133021 / 4

Type: `"number"`. Optional.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 8192),
}
```

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
    "minimum": 1,
    "multipleOf": 1
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

<a id="canonical-2030211122303320-2101132213201232-0013311001223210-2013002320202231-1300311110122313-1120003210310113-0003232001232300-3121202322203103"></a>

<a id="canonical-0232213111131113-3000130301033113-3130321032323322-0310322113200032-0111031032031203-2300010232230121-1123022132320213-3102130321011310"></a>

## unit property — inline_rate_limiter / 302113133021 / 5

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("SECOND",
    "MINUTE",
    "HOUR"),
}
```

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

- [use_http_lb_user_id](resources--http_loadbalancer--reference--group-008.md#canonical-0111112213112131-0132010230220112-2313023320220012-2321101131120111-3313220011300212-1010223212133203-3012222122321311-2331010313001322): complete subsection reference.

<a id="canonical-0311030031101220-0123331031322333-1023323133231013-1033323002011123-1200203122233230-2003110033123030-3100120322331133-3330120301132331"></a>

## Next pages — inline_rate_limiter / 302113133021 / 6

- [api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id](resources--http_loadbalancer--reference--group-008.md#canonical-3122110102021131-0210011132022301-2023111000333232-2031133321123202-0211131222132230-3013231033012330-0331030100133220-0330233330001203)
- [api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id](resources--http_loadbalancer--reference--group-008.md#canonical-0111112213112131-0132010230220112-2313023320220012-2321101131120111-3313220011300212-1010223212133203-3012222122321311-2331010313001322)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3122110102021131-0210011132022301-2023111000333232-2031133321123202-0211131222132230-3013231033012330-0331030100133220-0330233330001203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311333222131121-0302230321200102-2201233121311030-2233021120030020-2322012003232333-2002103131302002-2022303101002003-0213200032130211"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id — ref_user_id / 212101012032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-008.md#canonical-3230011100300201-3021322232013132-2120130113220102-3000312330300303-2132212200020332-2301101031230211-2123121301333223-1211301231000212)
- api_rate_limit.server_url_rules.inline_rate_limiter.ref_user_id

<a id="canonical-0312321121332201-3113300202001302-3222322030131120-1213033113110320-3001121022122211-1032333330302202-3032130333220323-0003023001120332"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
ref_user_id {
  # Configure direct properties listed below.
}
```

<a id="canonical-2322112331221333-0232012223102310-3132031030122111-3220213112321312-0110031220331330-1023333010303233-1010231320303030-2001102312033310"></a>

## Direct properties — ref_user_id / 212101012032 / 3

<a id="canonical-3021002031022030-2031303233320112-1013032001120230-3131221232301001-2032310231322012-0031130332120332-2210020301120311-1200102032033332"></a>

<a id="canonical-2121212012101002-1303122213213002-2131333223221301-3303233022112322-2233223300030132-0232120322010102-2222223012023320-1310302133010010"></a>

## name property — ref_user_id / 212101012032 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-1202103030033230-3013100312012120-2120112330021110-0230000023021331-1120300333311200-0022010010331002-3031013312310001-2203011120032130"></a>

<a id="canonical-2201011122133130-1302330333020013-1221330202301312-3110030300223131-2100012023320323-1300321232111013-1011113102101133-3112230123302320"></a>

## namespace property — ref_user_id / 212101012032 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-2330331132022020-2122310211300012-1030000302220332-2011010301101020-0211312121123013-0011100303322303-2133320202011232-3032213021332120"></a>

<a id="canonical-0102210033223233-1310211132111021-1002113233332213-1313320031120102-2213302130122231-3120210301120202-3131201310331333-0232203133300301"></a>

## tenant property — ref_user_id / 212101012032 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-3322033203222101-1330123331113010-3001313001331301-1031111211221211-2112312100302320-2033033233212022-2013232103331202-3313302121303233"></a>

## Next pages — ref_user_id / 212101012032 / 7

- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-008.md#canonical-3230011100300201-3021322232013132-2120130113220102-3000312330300303-2132212200020332-2301101031230211-2123121301333223-1211301231000212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0111112213112131-0132010230220112-2313023320220012-2321101131120111-3313220011300212-1010223212133203-3012222122321311-2331010313001322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032121023323313-1011131300102011-0312102310201230-3130301323211012-2012032131021231-3231311311312321-3033310322210123-3222321202200122"></a>

## api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id — use_http_lb_user_id / 011200113233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-008.md#canonical-3230011100300201-3021322232013132-2120130113220102-3000312330300303-2132212200020332-2301101031230211-2123121301333223-1211301231000212)
- api_rate_limit.server_url_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-2321003232233303-0032003201123232-1223011301311110-2232231301312220-1133230212302020-1120312313312200-1112311333202303-1113021212223001"></a>

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
use_http_lb_user_id = {}
```

<a id="canonical-3321303022222020-2022032113211200-0221101223213222-3001121211010010-3111130223301320-0033302210332133-2313033300111220-3131322121033313"></a>

## Direct properties — use_http_lb_user_id / 011200113233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323001113022110-0001022122002123-2331012120102321-0211200221032233-0302310320113302-3210212311321231-0222101232121122-0022322023200003"></a>

## Next pages — use_http_lb_user_id / 011200113233 / 4

- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--http_loadbalancer--reference--group-008.md#canonical-3230011100300201-3021322232013132-2120130113220102-3000312330300303-2132212200020332-2301101031230211-2123121301333223-1211301231000212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1003023331222033-2221012320323333-0303300312100321-0001000033001020-3133220003001112-2131311032211230-1003011113023110-0101320100203112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300013011223122-0300320310133221-0111102233032322-3113202033131201-3312030133330232-3000220303223113-3303213202222202-3333331132013012"></a>

## api_rate_limit.server_url_rules.ref_rate_limiter — ref_rate_limiter / 110322033212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- api_rate_limit.server_url_rules.ref_rate_limiter

<a id="canonical-1330031203223123-3121121020100021-1113012332021303-0121331213121122-2320211100220000-0122130101021303-3300023020221223-0001002303233203"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

Reference to a stored rate-limiter object for this scoped rule. Select exactly one of
ref\_rate\_limiter and inline\_rate\_limiter.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
ref_rate_limiter {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212210330200323-1231001321213201-0210123200031010-3133201331010120-3110112321313210-3300003103222221-1203220003212032-0333231030210332"></a>

## Direct properties — ref_rate_limiter / 110322033212 / 3

<a id="canonical-1310012201311023-0201332130332211-2211011303100323-2021212021311023-1133322120000201-3013023012100100-1123112131003310-2130223123011331"></a>

<a id="canonical-2221301230011110-1010131002312320-2011311201232321-1322120321211301-1223100001331122-3302310223230203-3013100002230202-0231121110223013"></a>

## name property — ref_rate_limiter / 110322033212 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-0300103000212111-2001320101200213-0321013210111113-2312303311300012-1112212023100112-1110023321203031-3113111210013311-0021212212331001"></a>

<a id="canonical-1003011003000122-3001113022321310-1121132032330233-1313220102020333-0030102113330201-1100010122133201-0202033322201020-0010311210330111"></a>

## namespace property — ref_rate_limiter / 110322033212 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-1211130120212002-2302203130300323-0103001320133332-2121031311132003-1201312100322313-1033312001211223-2322000021331000-3221013001122213"></a>

<a id="canonical-3312313232133300-2302002211320100-1033313033213202-1330201003031000-2101300122200112-1112232121120031-1320210102101013-1312003202321321"></a>

## tenant property — ref_rate_limiter / 110322033212 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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

<a id="canonical-2012213211331101-3103230112012103-2011212201122222-3001310322021232-1011300030001311-1131331010123020-0011322210333120-2123313332101203"></a>

## Next pages — ref_rate_limiter / 110322033212 / 7

- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312011021331002-0022203030222121-3010132210103331-3220322121311030-3111132200231031-0102132301013323-3211312102201301-2001112120013031"></a>

## api_rate_limit.server_url_rules.request_matcher — request_matcher / 113030311230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- api_rate_limit.server_url_rules.request_matcher

<a id="canonical-1113130322101231-2331120311211001-1030011233331110-1231131322213301-2013021322301312-1213232230202310-0122321003212010-0130330221223301"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for request matcher.

Upstream description:

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

Terraform syntax:

```terraform
request_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313022333202222-1112102103003012-2301013300012013-1232023023031133-0133020212123302-1113103321120203-0333103330301101-3311232030320022"></a>

## Direct properties — request_matcher / 113030311230 / 3

- [cookie_matchers](resources--http_loadbalancer--reference--group-008.md#canonical-3030321010032312-3113020200210321-1011003032003303-2302331212000220-1302011033022103-2110231223003002-2030101213023133-0200310211021311): complete subsection reference.

- [headers](resources--http_loadbalancer--reference--group-008.md#canonical-2110010003100323-3222013231131210-0210010210203221-1303021011022201-2201021022202310-0030031033220203-3110332013112333-0102002230320002): complete subsection reference.

- [jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-0302102023233130-2111123023213203-1201203120102000-3202032313032233-2203031230230022-2201032230331021-3001131132230201-1323020131132011): complete subsection reference.

- [query_params](resources--http_loadbalancer--reference--group-009.md#canonical-3132321313100031-2233313300131223-0233020023131212-1210213032330332-3320001101033110-1331311133030131-1220323302313300-0001103011302011): complete subsection reference.

<a id="canonical-0102200221223033-0003022321010220-1300320000331001-2121320123303222-0001231122001211-2331203321011210-3002232100302031-2210123331201032"></a>

## Next pages — request_matcher / 113030311230 / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-008.md#canonical-3030321010032312-3113020200210321-1011003032003303-2302331212000220-1302011033022103-2110231223003002-2030101213023133-0200310211021311)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-2110010003100323-3222013231131210-0210010210203221-1303021011022201-2201021022202310-0030031033220203-3110332013112333-0102002230320002)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-0302102023233130-2111123023213203-1201203120102000-3202032313032233-2203031230230022-2201032230331021-3001131132230201-1323020131132011)
- [api_rate_limit.server_url_rules.request_matcher.query_params](resources--http_loadbalancer--reference--group-009.md#canonical-3132321313100031-2233313300131223-0233020023131212-1210213032330332-3320001101033110-1331311133030131-1220323302313300-0001103011302011)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3030321010032312-3113020200210321-1011003032003303-2302331212000220-1302011033022103-2110231223003002-2030101213023133-0200310211021311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312131010113303-0132320133203210-0002031030103223-2311332120132203-0122320122030120-2122231101231230-1330223011112303-2300321302123210"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers — cookie_matchers / 020121001030 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers

<a id="canonical-3133312200321003-0312332230100021-1223101200313030-2300010323313330-3030212021203021-0021033123112312-3112222303011000-3030210112313320"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333232233331121-0122023331000122-3320210222103003-0102022133123121-3313300121033232-0303330003221320-3222322311011222-3302233212001122"></a>

## Direct properties — cookie_matchers / 020121001030 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-2202121022331120-2021033031311003-2201002000201210-3113213321210313-2330120211331220-1302132220300001-3230021023012033-1220211313332010): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-008.md#canonical-1103012203322132-3332330132312221-3003020232323030-3010222032121131-2012301020000312-2230020212030331-2033011030222122-3330231013010230): complete subsection reference.

<a id="canonical-1303301023020022-3330102033230130-3203033103230002-2013010223333202-2133202202000210-2001231113213322-1001300002021210-0200300000023020"></a>

<a id="canonical-1031021121302222-2023011001021023-0031230212132113-2320020323202333-2132001001222031-3321203301133320-1021122012313120-3310232010011030"></a>

## invert_matcher property — cookie_matchers / 020121001030 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](resources--http_loadbalancer--reference--group-008.md#canonical-0233132330222300-0322203213003123-1110332230220311-2001222013301113-0113101031000130-2020223223231233-3121331032013112-3333332311132003): complete subsection reference.

<a id="canonical-1022103000322023-1231231223300022-2022023210002031-1310023320023333-1201232203211000-0131031012113030-0310121103202331-1021320113130302"></a>

<a id="canonical-1330310210031022-0010220330111132-3120222113313313-1010031010100203-3000102323131100-0332233223123023-1213233213003021-1130230010132220"></a>

## name property — cookie_matchers / 020121001030 / 5

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0332122030123133-2331133312200310-1012033230323000-0303321203010123-0113220121002210-2121310223333312-1230323121320323-0330223100300322"></a>

## Next pages — cookie_matchers / 020121001030 / 6

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-2202121022331120-2021033031311003-2201002000201210-3113213321210313-2330120211331220-1302132220300001-3230021023012033-1220211313332010)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present](resources--http_loadbalancer--reference--group-008.md#canonical-1103012203322132-3332330132312221-3003020232323030-3010222032121131-2012301020000312-2230020212030331-2033011030222122-3330231013010230)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item](resources--http_loadbalancer--reference--group-008.md#canonical-0233132330222300-0322203213003123-1110332230220311-2001222013301113-0113101031000130-2020223223231233-3121331032013112-3333332311132003)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2202121022331120-2021033031311003-2201002000201210-3113213321210313-2330120211331220-1302132220300001-3230021023012033-1220211313332010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223122001030033-2122220322312123-1012033022123002-3022313023021031-0212021301311021-0311123300233033-2203123023022112-1030303221222232"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present — check_not_present / 221121100211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-008.md#canonical-3030321010032312-3113020200210321-1011003032003303-2302331212000220-1302011033022103-2110231223003002-2030101213023133-0200310211021311)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-3232131220012232-3032000200111030-2211230300030212-3123302200313202-3131003200202301-3303320300020232-3112123003202221-0321330103100020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-3111120302312201-2111133300001230-3320123322230221-1303033032110021-3113110103130110-2003121300330201-0022002111320012-0032011311201131"></a>

## Direct properties — check_not_present / 221121100211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213201031100322-0123113030021010-0132010133121020-1010133112313333-2202121102330100-0211321313122120-2021302123013213-0101310220311332"></a>

## Next pages — check_not_present / 221121100211 / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-008.md#canonical-3030321010032312-3113020200210321-1011003032003303-2302331212000220-1302011033022103-2110231223003002-2030101213023133-0200310211021311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1103012203322132-3332330132312221-3003020232323030-3010222032121131-2012301020000312-2230020212030331-2033011030222122-3330231013010230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110300122223223-0321230022310010-1120010003203223-0232310303231010-1301103232220303-3333210012011113-1220113212111231-0213112220033131"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present — check_present / 003002222322 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-008.md#canonical-3030321010032312-3113020200210321-1011003032003303-2302331212000220-1302011033022103-2110231223003002-2030101213023133-0200310211021311)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-3030031032203012-0202211030221232-0011100001330331-0131031332113201-1233332332111113-3110113103020231-0331231101011011-1003330023212011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-2213203033222313-1302120232020321-3003020331313331-3020013110110012-1102002221230100-3122322222133330-3231300110023001-1303320001302002"></a>

## Direct properties — check_present / 003002222322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112302333320023-0023313320000333-2323211303313200-0121300200012202-0302023021012220-0001333312113122-1202221020323320-3202221230101321"></a>

## Next pages — check_present / 003002222322 / 4

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-008.md#canonical-3030321010032312-3113020200210321-1011003032003303-2302331212000220-1302011033022103-2110231223003002-2030101213023133-0200310211021311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0233132330222300-0322203213003123-1110332230220311-2001222013301113-0113101031000130-2020223223231233-3121331032013112-3333332311132003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132000212312223-3212232123202332-1313033130133033-2123331202323302-2100130233200103-2130013023132013-1233002112201103-1233033003313000"></a>

## api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item — item / 303012230213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-008.md#canonical-3030321010032312-3113020200210321-1011003032003303-2302331212000220-1302011033022103-2110231223003002-2030101213023133-0200310211021311)
- api_rate_limit.server_url_rules.request_matcher.cookie_matchers.item

<a id="canonical-3311221321023323-3102310303101333-1323020331231111-0120233002132100-3130200031212220-0010202333130131-1203120112130320-3333223012212010"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-1033113103033101-2222121011313113-0230131220231202-0030032001012031-3223122132212103-3303132301101200-0131300102303320-1012113010001311"></a>

## Direct properties — item / 303012230213 / 3

<a id="canonical-3101113213333002-1210002131212311-0130131211313233-1221110202231111-2200132111113332-0100322133121201-2232011130001133-1132002301031320"></a>

<a id="canonical-3312303202211310-0200132330211130-2321212000232021-3133030123200303-2330131213022001-1103231222101330-3003321330130020-1103331201303312"></a>

## exact_values property — item / 303012230213 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-0032100323230130-0030031020103202-0231333320220330-3122030103333210-2023010122032001-0332131230321001-3032032110110221-3111021200232220"></a>

<a id="canonical-1330220113300332-3302311100010210-3220233210010313-1230133003103133-2010101200213131-0333033002112222-2120132011312330-0330211200120033"></a>

## regex_values property — item / 303012230213 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-1302222113313131-3230012313112003-0101123100002011-2310322231032333-0013301321133310-3202001321233111-2320132113302032-3302101303211333"></a>

<a id="canonical-0213131130121313-1231312211132223-2233223232210103-0323103010032003-2032220310113222-1120320123003032-3320010302003320-0201103300200030"></a>

## transformers property — item / 303012230213 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-0330301323323203-3301232311322022-2302230011200232-3133310020112011-2310033321133203-0123033000200330-1230332020233213-0131121301102011"></a>

## Next pages — item / 303012230213 / 7

- [api_rate_limit.server_url_rules.request_matcher.cookie_matchers](resources--http_loadbalancer--reference--group-008.md#canonical-3030321010032312-3113020200210321-1011003032003303-2302331212000220-1302011033022103-2110231223003002-2030101213023133-0200310211021311)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2110010003100323-3222013231131210-0210010210203221-1303021011022201-2201021022202310-0030031033220203-3110332013112333-0102002230320002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302111001100221-3231223313112012-3022203311111021-1203131232000320-0232232123321321-1221230001110031-3320301203010121-3310233320332231"></a>

## api_rate_limit.server_url_rules.request_matcher.headers — headers / 321212011122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- api_rate_limit.server_url_rules.request_matcher.headers

<a id="canonical-2331332131320001-3221232321332120-0302023121000121-2101001302102210-0313232113101231-3030111310231000-3101203101032120-3101300013023210"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103112302302132-2310111100230321-2032112302110201-2002002100311330-1231121303320232-3230120110032323-2231212110220301-2310030213102032"></a>

## Direct properties — headers / 321212011122 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-1003032032032203-3312021311031000-2133021323020023-2222002033001332-2030331122213212-2030222230021231-3332011312110303-3230331132003103): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-008.md#canonical-1002220323233121-3113030212220310-0132310320312122-1021323123032112-1222202130201301-3001012131022113-0222331333230300-0303300110210303): complete subsection reference.

<a id="canonical-0121122020321032-2200200100101002-1002131033330033-0223021300132323-0211222331331322-1333000221333111-0221122202300013-3230100222122111"></a>

<a id="canonical-0212000330311002-0112011023130131-0031223113001032-1311322312112211-2022101020201021-0310133230313132-0313302000102021-1212011302321002"></a>

## invert_matcher property — headers / 321212011122 / 4

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-008.md#canonical-3101031232033010-3033203020010231-1313301131133032-0301312311312211-0311113321321321-3213333333130130-2311300320121103-0332011012130010): complete subsection reference.

<a id="canonical-0023023131112020-1120112011232333-2321132223221301-2023231001321310-3231213223313200-0103120133331101-2133331333033213-2300211323113002"></a>

<a id="canonical-0222311221311321-2230112121231022-0230103230333133-0202220123211321-0212122103120121-1200203213032222-1211321003300302-2320033221303330"></a>

## name property — headers / 321212011122 / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2320032222131220-1022202130122231-0212303013102021-2013120131330303-0201131301133033-0232310200300022-1222233313103222-2111332112322301"></a>

## Next pages — headers / 321212011122 / 6

- [api_rate_limit.server_url_rules.request_matcher.headers.check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-1003032032032203-3312021311031000-2133021323020023-2222002033001332-2030331122213212-2030222230021231-3332011312110303-3230331132003103)
- [api_rate_limit.server_url_rules.request_matcher.headers.check_present](resources--http_loadbalancer--reference--group-008.md#canonical-1002220323233121-3113030212220310-0132310320312122-1021323123032112-1222202130201301-3001012131022113-0222331333230300-0303300110210303)
- [api_rate_limit.server_url_rules.request_matcher.headers.item](resources--http_loadbalancer--reference--group-008.md#canonical-3101031232033010-3033203020010231-1313301131133032-0301312311312211-0311113321321321-3213333333130130-2311300320121103-0332011012130010)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1003032032032203-3312021311031000-2133021323020023-2222002033001332-2030331122213212-2030222230021231-3332011312110303-3230331132003103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321323220202103-3131310230302301-1123131001010321-2102013220113203-3010302110030313-1112123220331011-2231331021311133-2122311130212223"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.check_not_present — check_not_present / 320011102032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-2110010003100323-3222013231131210-0210010210203221-1303021011022201-2201021022202310-0030031033220203-3110332013112333-0102002230320002)
- api_rate_limit.server_url_rules.request_matcher.headers.check_not_present

<a id="canonical-0133332013231331-0213232131333002-3111113123031100-0323200120023013-1022320112312122-3333030233320002-0233120022310202-2011123201120112"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-0102011223223011-0032231320211203-3133322022202223-2200001103310133-3201232012010033-2311223112310020-1210311322303123-1021110133103202"></a>

## Direct properties — check_not_present / 320011102032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012012110333333-3001331322030022-0333121100120332-0301030033232110-1011020300003111-3233321202303103-0202002301321133-3303323013213001"></a>

## Next pages — check_not_present / 320011102032 / 4

- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-2110010003100323-3222013231131210-0210010210203221-1303021011022201-2201021022202310-0030031033220203-3110332013112333-0102002230320002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1002220323233121-3113030212220310-0132310320312122-1021323123032112-1222202130201301-3001012131022113-0222331333230300-0303300110210303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301203131003202-2031332312032002-1332332000112331-0200332013133013-0132032002001303-0202313312120310-3201100100331010-0130201103203232"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.check_present — check_present / 102202001302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-2110010003100323-3222013231131210-0210010210203221-1303021011022201-2201021022202310-0030031033220203-3110332013112333-0102002230320002)
- api_rate_limit.server_url_rules.request_matcher.headers.check_present

<a id="canonical-2021210303201032-2223313010010001-3323202301000333-2321110112102111-2132011111133130-1011212321020022-1313003201100031-0033332221300100"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-1031212033111130-1302313010223120-0320231210222010-1321200020310020-3313300223100310-2012201102203213-0020122023333311-1233331002120222"></a>

## Direct properties — check_present / 102202001302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1001203311131032-3102102103230302-0301223213322322-0210110313103232-1120031210002330-1211112310012113-0210021110330001-0102011323103202"></a>

## Next pages — check_present / 102202001302 / 4

- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-2110010003100323-3222013231131210-0210010210203221-1303021011022201-2201021022202310-0030031033220203-3110332013112333-0102002230320002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3101031232033010-3033203020010231-1313301131133032-0301312311312211-0311113321321321-3213333333130130-2311300320121103-0332011012130010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220203212002311-1331020300122102-1010212230201022-0230113010232210-0103132303021122-2001130302131111-3030032231033111-2330121302301333"></a>

## api_rate_limit.server_url_rules.request_matcher.headers.item — item / 200101302012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-2110010003100323-3222013231131210-0210010210203221-1303021011022201-2201021022202310-0030031033220203-3110332013112333-0102002230320002)
- api_rate_limit.server_url_rules.request_matcher.headers.item

<a id="canonical-2230212202003002-1000331030232302-2113023323013203-1013010102313322-1223120322012310-0231232323030022-2232102001130222-1123230200331001"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210300033001112-2103221031232230-3020312223232010-0130011111110113-3301211200031022-0022232131032031-1322221103330213-3312223233020233"></a>

## Direct properties — item / 200101302012 / 3

<a id="canonical-1210021013011130-2022020310322033-3033102111013002-2032122332010113-0232331111333030-0032021212010022-3013020021211320-3232103201311101"></a>

<a id="canonical-2022023202010233-0121301222303210-2212000100112031-1111303010333003-2203133233203033-0003012033112130-0012121103010313-3302231300032320"></a>

## exact_values property — item / 200101302012 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-1111213222012212-0132320211211101-0221230130311122-1233310123322301-2201031300213301-1132023221033310-3300210100120033-0110022001113032"></a>

<a id="canonical-3103012302300131-1133131231123103-1203313201200211-2123223232011013-2033323232032212-1312030220222112-0000101113103120-0003002220203222"></a>

## regex_values property — item / 200101302012 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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

<a id="canonical-0203203032333323-0203120202012213-1232122032320100-3013231111110103-2202331310332301-1132123323232321-3203203132101201-3312112302302111"></a>

<a id="canonical-0103222231003123-1202033013232110-1012110020100202-2230022103020333-0013012103033101-3020320300230222-2121030300131022-3322100212331030"></a>

## transformers property — item / 200101302012 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-3111001231133330-3111033331132011-2103212322122210-0223312210303233-1323230210030010-1010123213102222-2223100003031032-0012121320311022"></a>

## Next pages — item / 200101302012 / 7

- [api_rate_limit.server_url_rules.request_matcher.headers](resources--http_loadbalancer--reference--group-008.md#canonical-2110010003100323-3222013231131210-0210010210203221-1303021011022201-2201021022202310-0030031033220203-3110332013112333-0102002230320002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0302102023233130-2111123023213203-1201203120102000-3202032313032233-2203031230230022-2201032230331021-3001131132230201-1323020131132011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221211230332001-2132012232220123-2320213021210132-0230312010020300-1122200330300031-0122221110112320-3203113203132310-1002032301132001"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims — jwt_claims / 133021131311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims

<a id="canonical-2122202123120102-0330310022311130-0010321120131103-1200312123023100-0112300133321012-0102113203002022-1313203213021003-1002200020100112"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

A list of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings. Note that all specified JWT claim predicates
must evaluate to true. Note that this feature only works on LBs with JWT Validation feature enabled.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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

Terraform syntax:

```terraform
jwt_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203022122011310-0323310032112023-0000021212221032-1100300020233301-1030003023222312-2300213301321021-1201302033202112-3313020023202112"></a>

## Direct properties — jwt_claims / 133021131311 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-3110212003110301-0012213201332323-2231013003321122-3222231313211112-3123231322300031-3033313331110022-3011033203220223-1322230321101012): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-008.md#canonical-2111123000110232-0111130230112022-1302021303100330-0102030300233330-0230202021202330-1130333223323133-2112213212031212-1333112300103233): complete subsection reference.

<a id="canonical-1211011020203222-3102202230120200-0332210000002113-2313320310201211-1233022013020110-3311110301100202-1212120301111123-0201112213220100"></a>

<a id="canonical-3130331031121130-1002331222302303-0303210032121001-1012020310112031-2111320300032020-0311303102300311-1233123310312033-3133223313000222"></a>

## invert_matcher property — jwt_claims / 133021131311 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [item](resources--http_loadbalancer--reference--group-009.md#canonical-2111202232213110-2023210223012321-3201222211232212-0122021020133110-2033032012020003-1320232321213322-1330131003233010-0000311123130312): complete subsection reference.

<a id="canonical-1310131200301010-3023303003301133-0000101000113021-1331020231123031-3313311111032233-1320123002201321-2112332202322132-1110312013120021"></a>

<a id="canonical-2212331211001133-1002212331122120-2112322130133212-1021110012131032-0321000210301320-2131113111303301-3330331012122133-1131013022211011"></a>

## name property — jwt_claims / 133021131311 / 5

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0021022130112312-3223300211121220-3122133033331021-1313102333030011-0122200331031123-0022200310210010-3010231313031211-1130110000222131"></a>

## Next pages — jwt_claims / 133021131311 / 6

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present](resources--http_loadbalancer--reference--group-008.md#canonical-3110212003110301-0012213201332323-2231013003321122-3222231313211112-3123231322300031-3033313331110022-3011033203220223-1322230321101012)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_present](resources--http_loadbalancer--reference--group-008.md#canonical-2111123000110232-0111130230112022-1302021303100330-0102030300233330-0230202021202330-1130333223323133-2112213212031212-1333112300103233)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims.item](resources--http_loadbalancer--reference--group-009.md#canonical-2111202232213110-2023210223012321-3201222211232212-0122021020133110-2033032012020003-1320232321213322-1330131003233010-0000311123130312)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3110212003110301-0012213201332323-2231013003321122-3222231313211112-3123231322300031-3033313331110022-3011033203220223-1322230321101012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203321320031123-3202310333022311-0012320321332332-1220032303301003-3323300021113310-1023110333111203-2033231333123331-2131113133332130"></a>

## api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present — check_not_present / 003000322221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102)
- [api_rate_limit.server_url_rules](resources--http_loadbalancer--reference--group-008.md#canonical-2303300100212022-3213100021113021-1233032210023103-1020002130210222-0033123132222033-3101320131312033-0111030113132230-1111201011032103)
- [api_rate_limit.server_url_rules.request_matcher](resources--http_loadbalancer--reference--group-008.md#canonical-0211111213012012-3333321222303331-3301323031020331-3232322030111023-2311220012221132-3202030011210213-3202021130321313-3102030333320031)
- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-0302102023233130-2111123023213203-1201203120102000-3202032313032233-2203031230230022-2201032230331021-3001131132230201-1323020131132011)
- api_rate_limit.server_url_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-1302130020323013-3311031003310020-0002202003112033-3311020103210013-1230310013233031-1133201120020232-2111123023103002-0302033200012301"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-1221201230330310-2330131231230333-1231310231311322-0133010032103213-3121133123003202-3123032012031310-0310311333212123-3312333320032023"></a>

## Direct properties — check_not_present / 003000322221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011112121102300-0222031032333231-3311120301123322-1320302131120320-2310311300101222-2323033021332020-3020020110202223-1123322203020310"></a>

## Next pages — check_not_present / 003000322221 / 4

- [api_rate_limit.server_url_rules.request_matcher.jwt_claims](resources--http_loadbalancer--reference--group-008.md#canonical-0302102023233130-2111123023213203-1201203120102000-3202032313032233-2203031230230022-2201032230331021-3001131132230201-1323020131132011)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2111123000110232-0111130230112022-1302021303100330-0102030300233330-0230202021202330-1130333223323133-2112213212031212-1333112300103233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
