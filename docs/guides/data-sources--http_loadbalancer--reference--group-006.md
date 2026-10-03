---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0033222013203301-2003302103212302-2121301330323030-2200313312320123-0301113310230300-2221032023200111-3303032201023130-2310213101321213"></a>

## Next pages — request_matcher / 201311002301 / 4

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-3203110111213100-1223231012330001-0233203021032310-2131302222300000-3033111300103310-1111123030033200-2101222302022232-3233113133013302)
- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-0232212102313232-2323303111313323-0123031102300131-2201002110331123-3000203131103330-1132020233302030-3202332222023011-3300012211002331)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-1113321233001100-0332320000321103-1123322230302213-3203333011210223-2130120322022232-2212331112131013-0231031211210333-1322300331113100)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-1112311221221333-3231121122102113-0112002031020032-0310001232322203-2123323100231210-2013202102322313-3233202132313111-0302321320330110)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3203110111213100-1223231012330001-0233203021032310-2131302222300000-3033111300103310-1111123030033200-2101222302022232-3233113133013302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232230322110303-0333032221130010-3320320312132011-0010130232033023-0232323010113130-3332002321332300-3110123001220012-2221222123320313"></a>

## api_protection_rules.api_groups_rules.request_matcher.cookie_matchers — cookie_matchers / 332313032031 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers

<a id="canonical-0130012030310310-0013022120113222-1303223113322201-1022120323212111-1032203033002311-3020021010010231-2232133112331223-1013320001220223"></a>

Type: `"list"`. Computed.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2120322320203030-2300021223102301-1232210232123322-1111130000131133-2302300011331001-1211213012013323-3300030303210211-3001130100122032"></a>

## Direct properties — cookie_matchers / 332313032031 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-1102313302033322-0310212323022232-2300220211002031-0332310200331002-0313333222322113-0200200220131232-1122330022112210-0103001111133123): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-0032220033223233-0133311310310112-3012213011302120-0110111311323220-2312012120321131-2130202121100320-1312133003111313-0312033302023023): complete subsection reference.

<a id="canonical-3313031232103303-2330320013001212-3300003102301322-0302330102330101-0033022103232222-0233302201001311-2212210331123300-2313321012230223"></a>

<a id="canonical-1311303232230020-0133021132023103-2121210110210111-1230003233211130-0202221232212201-3322330000210221-3230311121003210-1333103103211220"></a>

## invert_matcher property — cookie_matchers / 332313032031 / 4

Type: `"bool"`. Computed.

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

- [item](data-sources--http_loadbalancer--reference--group-006.md#canonical-2300110132202210-0330210020033233-1012331012012003-2330313121330222-1321220203330300-2322100121000201-0031302132100121-2210333223302001): complete subsection reference.

<a id="canonical-2032211323312000-0323300230302020-0233201032220302-3002210103213101-3121120120021312-1203202030122032-2130333220333100-2320230120333302"></a>

<a id="canonical-2231033311030123-2332102002112212-0121320312033233-3313221210323122-3101020022212013-3030221130301203-3101333012003101-1012211322201220"></a>

## name property — cookie_matchers / 332313032031 / 5

Type: `"string"`. Computed.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3302011122033221-1322003113031220-3101202032021031-2102121303321233-1300101001301222-3201011031013301-3100311310201211-2013310102120301"></a>

## Next pages — cookie_matchers / 332313032031 / 6

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-1102313302033322-0310212323022232-2300220211002031-0332310200331002-0313333222322113-0200200220131232-1122330022112210-0103001111133123)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-0032220033223233-0133311310310112-3012213011302120-0110111311323220-2312012120321131-2130202121100320-1312133003111313-0312033302023023)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item](data-sources--http_loadbalancer--reference--group-006.md#canonical-2300110132202210-0330210020033233-1012331012012003-2330313121330222-1321220203330300-2322100121000201-0031302132100121-2210333223302001)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1102313302033322-0310212323022232-2300220211002031-0332310200331002-0313333222322113-0200200220131232-1122330022112210-0103001111133123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232023103233222-3000100311103102-2333032231102311-3110323130301310-2122223030001011-3112213030111000-3322133002332010-1232300103203103"></a>

## api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_present — check_not_present / 101103030321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-3203110111213100-1223231012330001-0233203021032310-2131302222300000-3033111300103310-1111123030033200-2101222302022232-3233113133013302)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-2033210222013021-3232302212131000-0030033230202213-2223030111300133-3030311301111223-0221020321132020-3332232223212023-2101232100223010"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3331113330011312-3020012313230202-2103010001013100-0002223220332131-2211230222011323-0330111233111202-2302221003010203-0321321030033310"></a>

## Direct properties — check_not_present / 101103030321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202131102023231-0103022203013020-1230001221223011-1101103202030212-2203001002212031-3203003233000120-3312110030223033-1232112210100233"></a>

## Next pages — check_not_present / 101103030321 / 4

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-3203110111213100-1223231012330001-0233203021032310-2131302222300000-3033111300103310-1111123030033200-2101222302022232-3233113133013302)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0032220033223233-0133311310310112-3012213011302120-0110111311323220-2312012120321131-2130202121100320-1312133003111313-0312033302023023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230133313313001-0332111322203232-2003120223202221-0303232010222032-3130130323212320-1332030301003311-3020323321023021-0303013211122120"></a>

## api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_present — check_present / 323101333222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-3203110111213100-1223231012330001-0233203021032310-2131302222300000-3033111300103310-1111123030033200-2101222302022232-3233113133013302)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-0331202322011130-0212020210302221-0212032321301232-0010013013322330-0122112322030101-1221132121021030-1020123302301310-0332012100123031"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3200300331210101-1002320321113031-2023032200322010-0123332330312000-1121212131002122-1112001210002031-3103010230230302-3320322301133112"></a>

## Direct properties — check_present / 323101333222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313211133101320-1232202130211003-2130022010300301-0333210010210110-2020132333010110-0201001111100231-0230212321113310-3011020101030013"></a>

## Next pages — check_present / 323101333222 / 4

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-3203110111213100-1223231012330001-0233203021032310-2131302222300000-3033111300103310-1111123030033200-2101222302022232-3233113133013302)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2300110132202210-0330210020033233-1012331012012003-2330313121330222-1321220203330300-2322100121000201-0031302132100121-2210333223302001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111320010231012-0300133121331200-1010313201103203-3320310200133333-2120311002131123-0123101122103302-3333233330221220-1323131203003223"></a>

## api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item — item / 103120333012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-3203110111213100-1223231012330001-0233203021032310-2131302222300000-3033111300103310-1111123030033200-2101222302022232-3233113133013302)
- api_protection_rules.api_groups_rules.request_matcher.cookie_matchers.item

<a id="canonical-2130022102013322-3211010102003101-1033102101132333-0322132332122320-3002103022131233-2230000322022213-2210103323301300-0023230130033211"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3002311301000103-0232232310322000-2002121203130303-0122120131013132-3221213033203101-3322220130000003-2000011302032322-1011202222103201"></a>

## Direct properties — item / 103120333012 / 3

<a id="canonical-0003002313200212-1132302120001302-2113021110331203-2122321332330101-0322231032000031-0312101203131102-0223310233001223-1302233003011023"></a>

<a id="canonical-3332121003032323-1011302232203133-2121012210121033-3001123000321302-1303322332003110-0230321123130111-2132030313323010-1131102202313021"></a>

## exact_values property — item / 103120333012 / 4

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

<a id="canonical-0103302231220101-3223230203011023-2330110102230023-2303033011131301-1021011020231333-1012302020021310-1123202231130221-1022222311233033"></a>

<a id="canonical-1231130211210022-2023022320203111-1000130311110230-0323030300331032-3102331110300201-0222222303210303-2302330111133033-0101221121213303"></a>

## regex_values property — item / 103120333012 / 5

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

<a id="canonical-0000001233113300-0120013120320001-1023103301223200-0023210322123013-3010003001312323-1102311233012131-1333012201011020-2221233021211332"></a>

<a id="canonical-2132313020310323-2301000211213102-3011221323022122-2203010203122100-0031232102323221-3033302103302020-0001202211212012-2210023132123001"></a>

## transformers property — item / 103120333012 / 6

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3020310110113031-0221030213220020-0030002123102231-2002213102012103-1223222101213220-3100232331231033-0313131013010010-1022103211100101"></a>

## Next pages — item / 103120333012 / 7

- [api_protection_rules.api_groups_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-3203110111213100-1223231012330001-0233203021032310-2131302222300000-3033111300103310-1111123030033200-2101222302022232-3233113133013302)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0232212102313232-2323303111313323-0123031102300131-2201002110331123-3000203131103330-1132020233302030-3202332222023011-3300012211002331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310011103310210-3321100323200033-2031000331222023-1022131130323333-0222223033132222-0020213130133132-0210121310202131-0023312300233110"></a>

## api_protection_rules.api_groups_rules.request_matcher.headers — headers / 332231123130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- api_protection_rules.api_groups_rules.request_matcher.headers

<a id="canonical-2112111133300102-1213212221021223-2330101001232320-3322103110200320-3120122312030021-3302313221103113-3022122130310022-3000122320023331"></a>

Type: `"list"`. Computed.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2113311222322023-1332000102230333-3102100320013312-1303222000120000-0020211210121032-3110211211122012-2011011301033233-3020120031313300"></a>

## Direct properties — headers / 332231123130 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-1010312230120123-0133020133011130-2210132232100331-0112112210333312-2123223121003200-0231322212201131-3302203310021120-0233203303211120): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-1201223320223213-1020212300312111-1201023132213033-3010002321323301-2131301223313012-3003033122000032-0103311123002303-1131303202113112): complete subsection reference.

<a id="canonical-2132110233331201-3022012130012300-0203102113310030-2220202300321213-2320021311213100-2011210021202110-0130202101111020-3333331221001213"></a>

<a id="canonical-2110232222020300-3001011212102210-1010123101010111-0330330312221301-3123330333133233-3023001220033321-3333223331103231-1123111111030103"></a>

## invert_matcher property — headers / 332231123130 / 4

Type: `"bool"`. Computed.

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

- [item](data-sources--http_loadbalancer--reference--group-006.md#canonical-1331221100033320-3012300222231031-2133321111231220-2230310330213012-3232302220102303-1301123200012200-2110103201321232-0221130323121310): complete subsection reference.

<a id="canonical-0232233020221102-3001322001310133-3030132233110222-3130312300032120-1033113321223323-1112211300123331-2301111311010132-3032133020220311"></a>

<a id="canonical-2321122021021000-2233331033030121-1300113030102102-3000211133202200-0020213002212200-3222333133202210-0030123133320323-2312331033232102"></a>

## name property — headers / 332231123130 / 5

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2000100201022022-3133201001331111-2312330001113011-2200323313111120-3012032322311021-1231312133302212-2130310122023032-1112232102013301"></a>

## Next pages — headers / 332231123130 / 6

- [api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-1010312230120123-0133020133011130-2210132232100331-0112112210333312-2123223121003200-0231322212201131-3302203310021120-0233203303211120)
- [api_protection_rules.api_groups_rules.request_matcher.headers.check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-1201223320223213-1020212300312111-1201023132213033-3010002321323301-2131301223313012-3003033122000032-0103311123002303-1131303202113112)
- [api_protection_rules.api_groups_rules.request_matcher.headers.item](data-sources--http_loadbalancer--reference--group-006.md#canonical-1331221100033320-3012300222231031-2133321111231220-2230310330213012-3232302220102303-1301123200012200-2110103201321232-0221130323121310)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1010312230120123-0133020133011130-2210132232100331-0112112210333312-2123223121003200-0231322212201131-3302203310021120-0233203303211120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230013021321222-1312101233300230-1112221123132223-2230223330023212-3300313220330013-1310010212323020-1210201303323313-1221121320201200"></a>

## api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present — check_not_present / 313301222010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-0232212102313232-2323303111313323-0123031102300131-2201002110331123-3000203131103330-1132020233302030-3202332222023011-3300012211002331)
- api_protection_rules.api_groups_rules.request_matcher.headers.check_not_present

<a id="canonical-3132211233331323-3310210020330011-3202333323301332-3323223231312021-0132011212323322-0300202122321313-1200301110023001-1300030012321201"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2002332311211330-3111112211213313-1033101031121223-1121302200202322-2221331023300021-3211123231123131-1110101100030303-1121003300230232"></a>

## Direct properties — check_not_present / 313301222010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132030331133333-0123311133020100-1011323312301233-2332302333230231-0203233233101023-2133033100123321-2112111332233011-2002132333313113"></a>

## Next pages — check_not_present / 313301222010 / 4

- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-0232212102313232-2323303111313323-0123031102300131-2201002110331123-3000203131103330-1132020233302030-3202332222023011-3300012211002331)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1201223320223213-1020212300312111-1201023132213033-3010002321323301-2131301223313012-3003033122000032-0103311123002303-1131303202113112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012120133321300-2030010121001120-2010223021021311-3213032221030330-0223203330210101-3110332111001112-3020002122003023-3031030310331303"></a>

## api_protection_rules.api_groups_rules.request_matcher.headers.check_present — check_present / 231000031003 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-0232212102313232-2323303111313323-0123031102300131-2201002110331123-3000203131103330-1132020233302030-3202332222023011-3300012211002331)
- api_protection_rules.api_groups_rules.request_matcher.headers.check_present

<a id="canonical-0003232111030210-3331303221210320-3300011203101310-2310221031101023-3303023202301000-1333330302111230-0210001103031101-1133220033203323"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1323022323111111-1323310001321200-3031103111321223-3011021100223003-3110101210131003-3001231312220031-3203130101313221-1032321100313000"></a>

## Direct properties — check_present / 231000031003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311122322022330-2022002203322302-1002313321120032-1221012330322113-0011331121121303-0001221303222102-0312111020203222-1223201203300313"></a>

## Next pages — check_present / 231000031003 / 4

- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-0232212102313232-2323303111313323-0123031102300131-2201002110331123-3000203131103330-1132020233302030-3202332222023011-3300012211002331)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1331221100033320-3012300222231031-2133321111231220-2230310330213012-3232302220102303-1301123200012200-2110103201321232-0221130323121310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110002210201333-1233321102132331-0312323330301313-3112120301233330-1332021100031322-1012133022132331-1332123321000103-1231200232233311"></a>

## api_protection_rules.api_groups_rules.request_matcher.headers.item — item / 030231002120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-0232212102313232-2323303111313323-0123031102300131-2201002110331123-3000203131103330-1132020233302030-3202332222023011-3300012211002331)
- api_protection_rules.api_groups_rules.request_matcher.headers.item

<a id="canonical-2133103311130013-2330110020330231-3220031323102121-1232100102030330-1113023111321232-2231002123221131-0300200103120033-1001131132320333"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0320211030231211-2021222033023000-3022311011030000-2011301123330033-3132302121120011-1220121201122133-2120230030021123-3200231311300020"></a>

## Direct properties — item / 030231002120 / 3

<a id="canonical-0221032223003120-1032132333010121-3203232330020122-3300313033201022-3231320132123130-2323031201000200-2011321210121120-3113321000112021"></a>

<a id="canonical-3013021013032023-0133331022121022-2230211021102223-1230012132112213-1303013120102133-1112121011201131-2300202321131201-1203311220002102"></a>

## exact_values property — item / 030231002120 / 4

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

<a id="canonical-3123300103030021-0122223322133300-0101331032333201-0230133000301212-2230130111233003-0133210000323002-3031223333313101-2321023302231122"></a>

<a id="canonical-2110223031200203-1322210233031321-3302111310003333-0130121330210122-1003103023001202-2210331120131233-1123331012322213-1112221221200003"></a>

## regex_values property — item / 030231002120 / 5

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

<a id="canonical-2021000330333220-2303221232133131-3132333102231110-0030203233220010-0321323112321010-0322121031031111-1001132231010332-0203001301310202"></a>

<a id="canonical-2120211112331011-3133012023323000-2031321010310122-1330132200000013-1032230230302110-0123120100130012-1201133032123132-3120210020111221"></a>

## transformers property — item / 030231002120 / 6

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1200211003031202-3223323101200001-3322333321022112-2110103023100102-2321102320213021-2221310022231220-2010211113213222-1210131022023101"></a>

## Next pages — item / 030231002120 / 7

- [api_protection_rules.api_groups_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-0232212102313232-2323303111313323-0123031102300131-2201002110331123-3000203131103330-1132020233302030-3202332222023011-3300012211002331)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1113321233001100-0332320000321103-1123322230302213-3203333011210223-2130120322022232-2212331112131013-0231031211210333-1322300331113100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020212330311223-0211332231303033-3211113323003303-3130221013130231-2313023030102000-2312313200011002-2220332301312103-1022233012101320"></a>

## api_protection_rules.api_groups_rules.request_matcher.jwt_claims — jwt_claims / 121001223132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims

<a id="canonical-3033131130133133-2201212300312123-1220122212012131-0100313021022131-3003230200113303-3231103211202321-1032231101333132-2330130111121212"></a>

Type: `"list"`. Computed.

List of predicates for various JWT claims that need to match. The criteria for matching each JWT
claim are described in individual JWTClaimMatcherType instances. The actual JWT claims values are
extracted from the JWT payload as a list of strings.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-3023212030210311-2221020310013000-2330012200320220-1203213223221222-0321202010302311-2023202123211131-0332331013110221-3310113202122331"></a>

## Direct properties — jwt_claims / 121001223132 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-0330123223201101-1220231000021332-0011300131231300-3032333121231212-2303100103130302-0231013123231313-2221222211021022-2031021233033223): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-3130201230012223-2201333132032111-1321230022132312-0131331211021312-1230222002020122-2131131202001123-3200020103031303-1012231132312120): complete subsection reference.

<a id="canonical-2030311123320131-0103132100031020-1120112322113010-0021131100023232-3123302033331202-0320122233130202-2031213003302312-2213330101120221"></a>

<a id="canonical-3130211023330233-0323330000122103-1233223312231201-1303212023330331-1222002212013300-3122232232233311-0301300232300012-0323111320122300"></a>

## invert_matcher property — jwt_claims / 121001223132 / 4

Type: `"bool"`. Computed.

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

- [item](data-sources--http_loadbalancer--reference--group-006.md#canonical-2330200333121301-2101131203123312-1001003110000223-3010011212110132-2033212313120002-3330303020331012-0013331100021321-0223132332112011): complete subsection reference.

<a id="canonical-0103110310302212-3113303211131322-0020031333132231-2132011112001012-0202100101113030-3111131133211333-0302122130221010-3023021220330102"></a>

<a id="canonical-2111030000233012-1320301113002222-3032333120133112-1331313322221122-0203113321221321-1001030323230233-3300311023102033-3031233121123322"></a>

## name property — jwt_claims / 121001223132 / 5

Type: `"string"`. Computed.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2221032201221102-0031231033210110-2022113132202202-0030323213203002-1133020021000303-1210110121231031-2300321130210221-1311321013110311"></a>

## Next pages — jwt_claims / 121001223132 / 6

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-0330123223201101-1220231000021332-0011300131231300-3032333121231212-2303100103130302-0231013123231313-2221222211021022-2031021233033223)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-3130201230012223-2201333132032111-1321230022132312-0131331211021312-1230222002020122-2131131202001123-3200020103031303-1012231132312120)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item](data-sources--http_loadbalancer--reference--group-006.md#canonical-2330200333121301-2101131203123312-1001003110000223-3010011212110132-2033212313120002-3330303020331012-0013331100021321-0223132332112011)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0330123223201101-1220231000021332-0011300131231300-3032333121231212-2303100103130302-0231013123231313-2221222211021022-2031021233033223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010332221302111-2303110220033331-0232123201211113-1320310220232231-2332120102320201-3310300013103202-1331101322302211-1302320102010112"></a>

## api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present — check_not_present / 022333220221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-1113321233001100-0332320000321103-1123322230302213-3203333011210223-2130120322022232-2212331112131013-0231031211210333-1322300331113100)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-2200302230320021-3230230211022322-2022031103130331-3232110200022331-2001102331011031-1022300030320212-1113312021301032-1212130222303031"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2132321200322121-0133110022231233-0320221322013010-0112230122120323-3101331333201203-3211000322113220-2202202323010303-2301312303013333"></a>

## Direct properties — check_not_present / 022333220221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031323133132111-2113121333333330-0320100200203213-2123133200131001-3001333100013012-0210303010210332-0331122013102203-0113203212312220"></a>

## Next pages — check_not_present / 022333220221 / 4

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-1113321233001100-0332320000321103-1123322230302213-3203333011210223-2130120322022232-2212331112131013-0231031211210333-1322300331113100)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3130201230012223-2201333132032111-1321230022132312-0131331211021312-1230222002020122-2131131202001123-3200020103031303-1012231132312120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210201020013310-0200303102320130-1030010302132202-2030001000312302-0120011202130323-2211313220332322-3033021301120013-2300003010310010"></a>

## api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present — check_present / 121320303301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-1113321233001100-0332320000321103-1123322230302213-3203333011210223-2130120322022232-2212331112131013-0231031211210333-1322300331113100)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.check_present

<a id="canonical-0313321131121022-1110230110300030-2301022200133002-1303323003110222-0303011103222030-0101001013033020-2221003011213322-1011233032312320"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0033301200231013-2120212121000303-0211120222211021-1120202113001201-3123111220332023-1123303221320013-2321233011020132-0000303101102100"></a>

## Direct properties — check_present / 121320303301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233132022303313-2110313123213131-3231332022012132-3230133222232002-0010320302130201-3313022022123322-3323200103021012-1232023201301300"></a>

## Next pages — check_present / 121320303301 / 4

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-1113321233001100-0332320000321103-1123322230302213-3203333011210223-2130120322022232-2212331112131013-0231031211210333-1322300331113100)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2330200333121301-2101131203123312-1001003110000223-3010011212110132-2033212313120002-3330303020331012-0013331100021321-0223132332112011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113101122203333-2232322113110322-0012232121232002-3202011231120300-2221222202130212-0100013210223012-3333131002221321-3302312003320121"></a>

## api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item — item / 112230303233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-1113321233001100-0332320000321103-1123322230302213-3203333011210223-2130120322022232-2212331112131013-0231031211210333-1322300331113100)
- api_protection_rules.api_groups_rules.request_matcher.jwt_claims.item

<a id="canonical-0033212322330101-3033303323100103-0032012111321220-3231320130001130-1012221200000212-3120012220202011-3233102220100130-1223013303330132"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2122323113133101-3113221002003331-1110123102010210-3131022032021011-0213203103030211-1213321211111132-3201002023201113-1221133100331233"></a>

## Direct properties — item / 112230303233 / 3

<a id="canonical-2100013212003103-3213032323003111-2102203012101122-3110302332013000-2010301323002212-1133020330313310-1232133030113100-3021311312111121"></a>

<a id="canonical-3010113021313331-3300102110310011-2033031211002113-0321313102102330-3031131030331023-1021110230031212-0031220022111121-3003202132223221"></a>

## exact_values property — item / 112230303233 / 4

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

<a id="canonical-3020033213123232-0300300303201100-2312313303123211-3022003212331311-3031230222213300-2103111303333302-3033123110313302-2111213120122000"></a>

<a id="canonical-1103103112221301-0102123120200203-0320001223133312-3200030020323012-1101113231123201-3323323010201120-3332012103111003-0000313231210102"></a>

## regex_values property — item / 112230303233 / 5

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

<a id="canonical-3102033122312131-0202311121010032-1100130022101110-0130211023013101-2002311310233232-3311232000122130-0000012301322200-1102202221100230"></a>

<a id="canonical-3203330333210330-0203301132100133-2120301101023302-1112230101213033-2332312300313123-2032202210032221-1102020020102310-0300112221131230"></a>

## transformers property — item / 112230303233 / 6

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1331032031301013-1010322330133230-1013310332302313-2310011021033000-2020313331213021-1223303111200200-2321102012022221-2210012112112321"></a>

## Next pages — item / 112230303233 / 7

- [api_protection_rules.api_groups_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-006.md#canonical-1113321233001100-0332320000321103-1123322230302213-3203333011210223-2130120322022232-2212331112131013-0231031211210333-1322300331113100)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1112311221221333-3231121122102113-0112002031020032-0310001232322203-2123323100231210-2013202102322313-3233202132313111-0302321320330110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023323033032320-1300021213221133-3222322121222123-0221223220200232-2132232220122322-2120223333330010-0003132231002102-1302113130003003"></a>

## api_protection_rules.api_groups_rules.request_matcher.query_params — query_params / 120031123221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- api_protection_rules.api_groups_rules.request_matcher.query_params

<a id="canonical-1313130202021101-2133201310021310-2010010002022223-0330301200332123-3000321031100120-1230020200100323-1133100031333330-2113221012020311"></a>

Type: `"list"`. Computed.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2101221131311000-1230122023023201-0011002202121030-0310302221011002-1112133331332301-3121221112031012-2013211300021030-1323223001303001"></a>

## Direct properties — query_params / 120031123221 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-3203000210000203-1202122322331301-0010022211002113-1031011113300022-3031331210003013-3000100230112131-2030321320112111-0321001123011033): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-1213323101301133-3313330210323102-3301311101333310-2112311123330030-2320211300002033-0202231133110011-3031003301010012-2031110133313322): complete subsection reference.

<a id="canonical-3211331022102222-3202001312230311-2100002100233120-2101122220031102-1123233213310011-3322131110231021-1110313221011122-2023101132220310"></a>

<a id="canonical-2201032102330323-2322231213013111-2322210331301322-0303200000012133-0223112113122011-1102232330310031-1201301023333033-2200203203033302"></a>

## invert_matcher property — query_params / 120031123221 / 4

Type: `"bool"`. Computed.

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

- [item](data-sources--http_loadbalancer--reference--group-006.md#canonical-3230012303111032-1021330130001222-1213001323133232-0310322122020113-0331022003111230-1002032020311330-1201322022301211-1132230213123201): complete subsection reference.

<a id="canonical-1231231332101101-1222103120313211-1322013031102303-2221322231113003-2220232032233312-2202113111311012-1210102322201322-1110100211222030"></a>

<a id="canonical-3113112112231212-1222230302322113-1301013012312211-3311111333311111-2211211030211203-3222130313122013-1301230331003002-0022300202000333"></a>

## key property — query_params / 120031123221 / 5

Type: `"string"`. Computed.

Case-sensitive HTTP query parameter name.

Upstream description:

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2323100023132102-0302210013010221-1230020212121320-3011012223313321-1232211131321212-2133111332033013-0133113333203213-3301031321011331"></a>

## Next pages — query_params / 120031123221 / 6

- [api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-3203000210000203-1202122322331301-0010022211002113-1031011113300022-3031331210003013-3000100230112131-2030321320112111-0321001123011033)
- [api_protection_rules.api_groups_rules.request_matcher.query_params.check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-1213323101301133-3313330210323102-3301311101333310-2112311123330030-2320211300002033-0202231133110011-3031003301010012-2031110133313322)
- [api_protection_rules.api_groups_rules.request_matcher.query_params.item](data-sources--http_loadbalancer--reference--group-006.md#canonical-3230012303111032-1021330130001222-1213001323133232-0310322122020113-0331022003111230-1002032020311330-1201322022301211-1132230213123201)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3203000210000203-1202122322331301-0010022211002113-1031011113300022-3031331210003013-3000100230112131-2030321320112111-0321001123011033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303202000320321-0231332013332023-1203300211003030-1011022220302222-1121300131113302-2010121033322223-2301001103132100-1003333100302330"></a>

## api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_present — check_not_present / 001213033201 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-1112311221221333-3231121122102113-0112002031020032-0310001232322203-2123323100231210-2013202102322313-3233202132313111-0302321320330110)
- api_protection_rules.api_groups_rules.request_matcher.query_params.check_not_present

<a id="canonical-1121123032321030-3301023122002133-1200213022121203-2133022113313212-1010231023102113-1002030023123311-1332110020103112-2320211133110111"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1010012210201013-2312110303312031-0030323020212101-0011223310033121-3211021101001223-1102230120111013-1212113100103321-2322133212111300"></a>

## Direct properties — check_not_present / 001213033201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132121133131231-3121003200030210-1130121323230101-2023333010220130-1112110301031121-0112322013201300-0313211312201032-3021222222303211"></a>

## Next pages — check_not_present / 001213033201 / 4

- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-1112311221221333-3231121122102113-0112002031020032-0310001232322203-2123323100231210-2013202102322313-3233202132313111-0302321320330110)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1213323101301133-3313330210323102-3301311101333310-2112311123330030-2320211300002033-0202231133110011-3031003301010012-2031110133313322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333222313032123-3112303222301033-2312232103132222-0323131300213201-0013132012022201-3332010331111213-0313112222323332-1101211313330200"></a>

## api_protection_rules.api_groups_rules.request_matcher.query_params.check_present — check_present / 323201023033 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-1112311221221333-3231121122102113-0112002031020032-0310001232322203-2123323100231210-2013202102322313-3233202132313111-0302321320330110)
- api_protection_rules.api_groups_rules.request_matcher.query_params.check_present

<a id="canonical-0213310013321223-2130032333310330-0030102011330310-2030213222101210-0010221233213122-1011213100013021-1102023111003021-3230231303232002"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1023011332000220-0221022131221100-1131332230333211-0123333021132303-3210133032131101-1310313002001231-3013231313102233-1122212200300102"></a>

## Direct properties — check_present / 323201023033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302321030211030-2202302200222103-2033030112312333-1212032020131312-0110333020102112-0100010222100332-1233222233201121-1112012023321000"></a>

## Next pages — check_present / 323201023033 / 4

- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-1112311221221333-3231121122102113-0112002031020032-0310001232322203-2123323100231210-2013202102322313-3233202132313111-0302321320330110)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3230012303111032-1021330130001222-1213001323133232-0310322122020113-0331022003111230-1002032020311330-1201322022301211-1132230213123201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212322320300100-3113110231312021-0020032000112321-1322132331331221-3120012013321001-3002003031102111-1113230200013331-1032011130232020"></a>

## api_protection_rules.api_groups_rules.request_matcher.query_params.item — item / 020003021112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_protection_rules](data-sources--http_loadbalancer--reference--group-004.md#canonical-1001113123112200-3021231111112211-1230200132223031-0000322332122021-2130101030201110-1332101110112301-2333122123121201-1302012023030120)
- [api_protection_rules.api_groups_rules](data-sources--http_loadbalancer--reference--group-005.md#canonical-2002232023311130-1033002333311031-0332313310012302-3322211013231312-3200002201220113-3032201213030211-3212031032302200-3011221110322313)
- [api_protection_rules.api_groups_rules.request_matcher](data-sources--http_loadbalancer--reference--group-005.md#canonical-0113212121200323-1203321213131302-3103231301301001-1112300230310120-3123011023222012-1003122223022113-0132012220130010-2031102322101202)
- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-1112311221221333-3231121122102113-0112002031020032-0310001232322203-2123323100231210-2013202102322313-3233202132313111-0302321320330110)
- api_protection_rules.api_groups_rules.request_matcher.query_params.item

<a id="canonical-2031332111223222-3020213003222132-0132331100100110-0201133131032211-3100311133033320-3332201023212023-0232021322130002-0313113103033300"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3130202012200032-3322203100100202-2321111202230030-0232323100012232-1331012310123311-1321002313213021-1002011201131201-3131123213322233"></a>

## Direct properties — item / 020003021112 / 3

<a id="canonical-0221010101010322-3230101203300313-2013211200103211-3233223321001012-2200103112102111-3132021200333322-1233303212003232-2202300220213123"></a>

<a id="canonical-2311111023202232-1031000323212133-2302333121030021-2112322310213121-0213200122213000-0021320331013223-2110123030133300-3123202013221331"></a>

## exact_values property — item / 020003021112 / 4

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

<a id="canonical-1011103332320022-0303221032131031-3210230120001101-1211121011033300-3031311330213100-3010121122123303-1030020000331121-3302002202022030"></a>

<a id="canonical-3101021003213232-2101321112212330-0122320201003101-3312023223121302-1313011023113332-2102121033320123-1001031210122113-1122010131201202"></a>

## regex_values property — item / 020003021112 / 5

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

<a id="canonical-2321233012212221-3203211312331023-2022132230120201-0203113320101100-2212013231131012-1302210003220121-1213132001331311-1331031301111331"></a>

<a id="canonical-0101102201320221-1122032320213031-3230002203323010-2013113021113303-1331121332230003-3030232301332123-0132230131100213-2000222003322013"></a>

## transformers property — item / 020003021112 / 6

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0220110311001123-1221100122113122-2103032012322230-3111012012232221-2023332222230130-2331203120133133-1032010212003103-0232330320221132"></a>

## Next pages — item / 020003021112 / 7

- [api_protection_rules.api_groups_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-006.md#canonical-1112311221221333-3231121122102113-0112002031020032-0310001232322203-2123323100231210-2013202102322313-3233202132313111-0302321320330110)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032010310322011-0131313210312131-1103222132030011-0001223101303121-0320322031001223-3302123311213332-2200311202200231-3121213101031021"></a>

## api_rate_limit — api_rate_limit / 232233220010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- api_rate_limit

<a id="canonical-1002023121010233-2202310032231001-1033211001111301-0230103012300001-1333032021233330-3310221010100120-1020010332012323-3212222311123321"></a>

Type: `"single"`. Computed.

\[OneOf: api\_rate\_limit, disable\_rate\_limit, rate\_limit; Default: disable\_rate\_limit\] Path-
or API-group-scoped rate limiting. Define server\_url\_rules or api\_endpoint\_rules and choose
inline\_rate\_limiter for an inline limit, or ref\_rate\_limiter for a stored rate-limiter
reference.

Upstream description:

Path- or API-group-scoped rate limiting. Define server\_url\_rules or api\_endpoint\_rules and
choose inline\_rate\_limiter for an inline limit, or ref\_rate\_limiter for a stored rate-limiter
reference.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ip_allowed_list_choice": "[\"bypass_rate_limiting_rules\",\"custom_ip_allowed_list\",\"ip_allowed_list\",\"no_ip_allowed_list\"]"
}
```

OneOf alternatives in this subsection:

- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-1002023121010233-2202310032231001-1033211001111301-0230103012300001-1333032021233330-3310221010100120-1020010332012323-3212222311123321)
- [disable_rate_limit](data-sources--http_loadbalancer--reference--group-017.md#canonical-2302102000331233-3112332122210222-1012101130112301-1003310221300013-3103112033310210-3213311023132331-2133321121033023-2003221013233103)
- [rate_limit](data-sources--http_loadbalancer--reference--group-023.md#canonical-0111113122010303-2301222300220321-1211003332003113-1321232202233231-0110133202230113-1313110222123003-1231230132033110-0121210312001113)

Select alternatives according to the provider validators above.

<a id="canonical-0331200030023111-2000212221310223-3003231133110322-0120023303133222-1112000113122111-1121213133112322-3130333001222320-0130202113211020"></a>

## Direct properties — api_rate_limit / 232233220010 / 3

- [api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202): complete subsection reference.

- [bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112): complete subsection reference.

- [custom_ip_allowed_list](data-sources--http_loadbalancer--reference--group-008.md#canonical-3013011132131110-1031303333020113-2201331231223013-0222022331001113-0210133030331301-0013322311301220-2310330013302312-0102233021031233): complete subsection reference.

- [ip_allowed_list](data-sources--http_loadbalancer--reference--group-008.md#canonical-3111133023102333-1120102130112313-3103100331031323-1110010110020202-1210323033311321-0221232303123301-1130212210120030-1332100011133320): complete subsection reference.

- [no_ip_allowed_list](data-sources--http_loadbalancer--reference--group-008.md#canonical-3012211103120212-2201223321132021-2021010120120230-0112133020310233-2120203202111131-3111213210002021-3101102323132030-1220030133230312): complete subsection reference.

- [server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201): complete subsection reference.

<a id="canonical-3331212132313313-0110322212122211-1000201030202212-0320120111032310-2130310022122002-1101333222100101-2012213322110020-3312133212232011"></a>

## Next pages — api_rate_limit / 232233220010 / 4

- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--http_loadbalancer--reference--group-007.md#canonical-3212232000313122-3222113311103233-1033012132010111-0212303312221003-2132033301211122-2220111123302201-1022110100221131-2120310032112112)
- [api_rate_limit.custom_ip_allowed_list](data-sources--http_loadbalancer--reference--group-008.md#canonical-3013011132131110-1031303333020113-2201331231223013-0222022331001113-0210133030331301-0013322311301220-2310330013302312-0102233021031233)
- [api_rate_limit.ip_allowed_list](data-sources--http_loadbalancer--reference--group-008.md#canonical-3111133023102333-1120102130112313-3103100331031323-1110010110020202-1210323033311321-0221232303123301-1130212210120030-1332100011133320)
- [api_rate_limit.no_ip_allowed_list](data-sources--http_loadbalancer--reference--group-008.md#canonical-3012211103120212-2201223321132021-2021010120120230-0112133020310233-2120203202111131-3111213210002021-3101102323132030-1220030133230312)
- [api_rate_limit.server_url_rules](data-sources--http_loadbalancer--reference--group-008.md#canonical-2132122302103230-0002313111033230-3300323123321233-3300121320230303-1222302322301233-3110211313202132-1110112220221330-0103311212321201)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321232103300303-1323313322332101-0210321001032101-0112331132122122-0132222102130201-0122130323331232-2332323131032132-2300200310013230"></a>

## api_rate_limit.api_endpoint_rules — api_endpoint_rules / 233230223313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- api_rate_limit.api_endpoint_rules

<a id="canonical-0122223033122321-0013012013330031-2123131100300323-1131010022012233-2121223010232230-2121101031032212-1302300202123033-1003132003002123"></a>

Type: `"list"`. Computed.

Ordered endpoint-specific rate-limit rules. Each rule must choose exactly one rate\_limiter\_choice:
inline\_rate\_limiter or ref\_rate\_limiter.

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
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

<a id="canonical-3002232122132003-3201231023201332-1320311132132031-1033032233112003-2233021203312022-2120202100001211-1001213213130223-0331013110011012"></a>

## Direct properties — api_endpoint_rules / 233230223313 / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-006.md#canonical-0131032323123312-1311303032130013-1220212230233012-2032221002002001-2210231022021123-3122132310322110-1200311001032320-0310111002030330): complete subsection reference.

- [api_endpoint_method](data-sources--http_loadbalancer--reference--group-006.md#canonical-0120011232003031-3130331002010212-1321200133001303-0121212001033121-3321001302301333-0302203321303322-3121030200303213-0330330012100322): complete subsection reference.

<a id="canonical-0003112212013023-1231232203103132-3133323002210030-3131203123023201-2313203122112213-2300301032010132-1233000231022310-1020322000010113"></a>

<a id="canonical-2313000123203131-0022101101121030-0112202000302320-3323032301223030-1021131213333210-0022231332110030-3300233110133223-0201330210123312"></a>

## api_endpoint_path property — api_endpoint_rules / 233230223313 / 4

Type: `"string"`. Computed.

API Endpoint. The endpoint (path) of the request.

Upstream description:

The endpoint (path) of the request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

- [client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000): complete subsection reference.

- [inline_rate_limiter](data-sources--http_loadbalancer--reference--group-006.md#canonical-1001100222001320-3011330202313020-3310020021000111-2303330102222022-2200220012031211-1333313010120013-1301330312003223-2002111303232321): complete subsection reference.

- [ref_rate_limiter](data-sources--http_loadbalancer--reference--group-006.md#canonical-3312223301223300-0100320231121331-2200103031212013-2133203231122320-3200220223312311-1322011023113302-0321210220333333-0013201011200123): complete subsection reference.

- [request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0231103031100013-0310103221330221-0231022012323121-1213102220211202-0230012101113130-2100003112133013-0213101310021131-3323023321121300): complete subsection reference.

<a id="canonical-3333110001032203-0212003321310202-1010303221222201-1332113301010300-1213302121120310-3021131232122230-0302121101133033-0001310301010312"></a>

<a id="canonical-2010032322320233-2010030230203330-1103220200031300-2033200200100320-2033111001233102-3321003002103303-0331221221213330-3122311201130202"></a>

## specific_domain property — api_endpoint_rules / 233230223313 / 5

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-3120312322310230-2322201332312132-2220023102123032-2020032003323310-0131201230322330-2203313322203021-2033101213023113-2220232102332202"></a>

## Next pages — api_endpoint_rules / 233230223313 / 6

- [api_rate_limit.api_endpoint_rules.any_domain](data-sources--http_loadbalancer--reference--group-006.md#canonical-0131032323123312-1311303032130013-1220212230233012-2032221002002001-2210231022021123-3122132310322110-1200311001032320-0310111002030330)
- [api_rate_limit.api_endpoint_rules.api_endpoint_method](data-sources--http_loadbalancer--reference--group-006.md#canonical-0120011232003031-3130331002010212-1321200133001303-0121212001033121-3321001302301333-0302203321303322-3121030200303213-0330330012100322)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-006.md#canonical-1001100222001320-3011330202313020-3310020021000111-2303330102222022-2200220012031211-1333313010120013-1301330312003223-2002111303232321)
- [api_rate_limit.api_endpoint_rules.ref_rate_limiter](data-sources--http_loadbalancer--reference--group-006.md#canonical-3312223301223300-0100320231121331-2200103031212013-2133203231122320-3200220223312311-1322011023113302-0321210220333333-0013201011200123)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0231103031100013-0310103221330221-0231022012323121-1213102220211202-0230012101113130-2100003112133013-0213101310021131-3323023321121300)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0131032323123312-1311303032130013-1220212230233012-2032221002002001-2210231022021123-3122132310322110-1200311001032320-0310111002030330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220001301103313-2201233120223233-2133322103113031-0031231300320022-3033000132033222-0203303220220030-1303022212201113-2323332321300031"></a>

## api_rate_limit.api_endpoint_rules.any_domain — any_domain / 230303332223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- api_rate_limit.api_endpoint_rules.any_domain

<a id="canonical-2102221102313031-2230022010102313-1202310221023203-3013211330012013-0331121211030221-1201311011021232-0202213030323310-3123323003303031"></a>

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

<a id="canonical-1133022311132100-0133011233211302-1231100222100231-0012132323023200-0123233023333111-3203020212021021-2011330013132003-2212112311302203"></a>

## Direct properties — any_domain / 230303332223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102112332221110-1312031112012102-0121220110132232-3300131330023001-2013210200321233-0003230121302101-1232301121033201-3333130222132001"></a>

## Next pages — any_domain / 230303332223 / 4

- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0120011232003031-3130331002010212-1321200133001303-0121212001033121-3321001302301333-0302203321303322-3121030200303213-0330330012100322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223030303112203-0012323131332033-3220132320230301-0020222032201121-2332312023002100-0010102201200332-3100022110310020-3200122132012000"></a>

## api_rate_limit.api_endpoint_rules.api_endpoint_method — api_endpoint_method / 300323331332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- api_rate_limit.api_endpoint_rules.api_endpoint_method

<a id="canonical-3211233130021311-0302321123031231-3003201222320113-0103021100131332-0320323023301133-1211330023330123-3330233231123031-2001100112130020"></a>

Type: `"single"`. Computed.

HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

Upstream description:

A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

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

<a id="canonical-1312021231323332-1301320221312120-0232000133332332-2011200211201310-2033012000101233-1303102231111311-0113122003003023-2012110233131323"></a>

## Direct properties — api_endpoint_method / 300323331332 / 3

<a id="canonical-2320233213300102-2220031030220010-2002010311033210-0002221331320300-0102101111331300-0020313010323112-2222300303233212-0000101130132220"></a>

<a id="canonical-1103012002220003-2233221301012112-3103301210210132-2103202023020313-0103100311112311-3320311132000131-0113312210310101-3133121200223120"></a>

## invert_matcher property — api_endpoint_method / 300323331332 / 4

Type: `"bool"`. Computed.

Invert Method Matcher. Invert the match result.

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

<a id="canonical-1123112301111222-1323302210322302-2222102122121113-0210213302123023-3233101233002003-2232203330130022-0300222300321013-2001133112100101"></a>

<a id="canonical-2232313001321021-2330302201222121-0020012233333031-0133101112001130-0211300112301302-2133233013122100-1213313021010031-3111003001020032"></a>

## methods property — api_endpoint_method / 300323331332 / 5

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of methods values to match against.

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

<a id="canonical-2000320100130331-0212231321002310-3100102002322001-2010313030032231-1010331210223203-2103030011102021-1230120303221200-3233211202102331"></a>

## Next pages — api_endpoint_method / 300323331332 / 6

- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032011321201021-0010120012202033-1311131133101231-0121033020223213-2303002012331012-2333101330020133-3212231233033223-2131201103131210"></a>

## api_rate_limit.api_endpoint_rules.client_matcher — client_matcher / 310023123333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- api_rate_limit.api_endpoint_rules.client_matcher

<a id="canonical-1231232003220032-0133331130110313-0222302202133231-1111032010102121-0131111032200322-0011323001330200-0003223330200023-1100230031121331"></a>

Type: `"single"`. Computed.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

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

<a id="canonical-2002301133200322-0230202320010113-0321020310011123-0030111302032002-0002311220233012-2133200303020301-2201103122131020-0031121203223002"></a>

## Direct properties — client_matcher / 310023123333 / 3

- [any_client](data-sources--http_loadbalancer--reference--group-006.md#canonical-1032011112101203-2130202330201002-2201131232232301-2321200320332103-1103213000033330-2300212302133203-1201122011331321-3212312133230031): complete subsection reference.

- [any_ip](data-sources--http_loadbalancer--reference--group-006.md#canonical-1123023232132220-2011112031202212-2132112320331111-2212103312332022-2320022210130330-3302313102321231-0320331131001112-2230133001020233): complete subsection reference.

- [asn_list](data-sources--http_loadbalancer--reference--group-006.md#canonical-2220010113322103-2303200312130120-2013201012023200-2132311102321203-3023021232321123-1010000322002131-1300101201102313-3030131102233110): complete subsection reference.

- [asn_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0330233120213000-3202021203001020-2331313312313330-0111303031020112-0200121231020322-2312023231332312-0320031122231211-2313101031013232): complete subsection reference.

- [client_selector](data-sources--http_loadbalancer--reference--group-006.md#canonical-1231020233021330-1100132112121213-3133223330101231-0213212332020011-0210001210131101-2232020012022012-1223013211020113-1302301102101122): complete subsection reference.

- [ip_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0310023002110200-2232121210033223-1101202102102211-2011130300322220-0001201131202011-3123321330013330-1332300103330120-3111213331310031): complete subsection reference.

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-006.md#canonical-1010101230032120-1332030020203310-0310303231312013-0313113302011012-3301123232320113-1222332331202232-3310202312201332-2322133301313021): complete subsection reference.

- [ip_threat_category_list](data-sources--http_loadbalancer--reference--group-006.md#canonical-3202230111330110-1103220101312101-0133122333132123-0132002123120000-1031320313310022-1321201102030001-2022113003101022-0130230103232032): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-3121101201010002-0110102323110222-0122303223121031-2202001231013321-2000132332121310-0111312010022200-3030211033013112-1220311323120003): complete subsection reference.

<a id="canonical-0002332131322323-2311203313012301-0330121302333303-3311200301131313-0210030122313130-0223302300110102-3122200322032021-2330130210133023"></a>

## Next pages — client_matcher / 310023123333 / 4

- [api_rate_limit.api_endpoint_rules.client_matcher.any_client](data-sources--http_loadbalancer--reference--group-006.md#canonical-1032011112101203-2130202330201002-2201131232232301-2321200320332103-1103213000033330-2300212302133203-1201122011331321-3212312133230031)
- [api_rate_limit.api_endpoint_rules.client_matcher.any_ip](data-sources--http_loadbalancer--reference--group-006.md#canonical-1123023232132220-2011112031202212-2132112320331111-2212103312332022-2320022210130330-3302313102321231-0320331131001112-2230133001020233)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_list](data-sources--http_loadbalancer--reference--group-006.md#canonical-2220010113322103-2303200312130120-2013201012023200-2132311102321203-3023021232321123-1010000322002131-1300101201102313-3030131102233110)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0330233120213000-3202021203001020-2331313312313330-0111303031020112-0200121231020322-2312023231332312-0320031122231211-2313101031013232)
- [api_rate_limit.api_endpoint_rules.client_matcher.client_selector](data-sources--http_loadbalancer--reference--group-006.md#canonical-1231020233021330-1100132112121213-3133223330101231-0213212332020011-0210001210131101-2232020012022012-1223013211020113-1302301102101122)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0310023002110200-2232121210033223-1101202102102211-2011130300322220-0001201131202011-3123321330013330-1332300103330120-3111213331310031)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list](data-sources--http_loadbalancer--reference--group-006.md#canonical-1010101230032120-1332030020203310-0310303231312013-0313113302011012-3301123232320113-1222332331202232-3310202312201332-2322133301313021)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list](data-sources--http_loadbalancer--reference--group-006.md#canonical-3202230111330110-1103220101312101-0133122333132123-0132002123120000-1031320313310022-1321201102030001-2022113003101022-0130230103232032)
- [api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-3121101201010002-0110102323110222-0122303223121031-2202001231013321-2000132332121310-0111312010022200-3030211033013112-1220311323120003)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1032011112101203-2130202330201002-2201131232232301-2321200320332103-1103213000033330-2300212302133203-1201122011331321-3212312133230031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112112121213021-0112003210113200-2132323031002332-1120202220313202-3030031203322310-3130123331203020-0102113123131310-3020331301203302"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.any_client — any_client / 100122213031 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.any_client

<a id="canonical-1113032230013132-1303021002300001-3022022132202211-1002032022022032-1203131203210122-3111012301100322-2112321121002120-3030113333333311"></a>

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

<a id="canonical-0130133222300110-2103313221022023-1201110013103101-0300302033301203-1310031322113133-1331013300220322-2220021210320201-2233122002002212"></a>

## Direct properties — any_client / 100122213031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322311030003030-2130002031110120-0032322110212110-2000302011121301-0313211013032210-2003222121121210-0320022230232032-1120131331013003"></a>

## Next pages — any_client / 100122213031 / 4

- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1123023232132220-2011112031202212-2132112320331111-2212103312332022-2320022210130330-3302313102321231-0320331131001112-2230133001020233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121131223232101-1302121132000333-1133311001030113-0122022330023331-2322002123223321-2001333213132121-1033210111232132-0100003101230301"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.any_ip — any_ip / 300011100123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.any_ip

<a id="canonical-2020331312220303-2112010102311200-2223003331213013-1223323233223321-0130003022122330-1333312320313321-3201210201320323-1001111101130301"></a>

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

<a id="canonical-3031020233003111-0133200310133012-2210322122301012-3333123103002203-2320022031201231-0203331311001013-3020310310320103-0311232212023213"></a>

## Direct properties — any_ip / 300011100123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231203211212302-3011200002301002-2130013303131000-3001123233302232-0113021311231123-2303031313001313-1020303100201213-3333311333033330"></a>

## Next pages — any_ip / 300011100123 / 4

- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2220010113322103-2303200312130120-2013201012023200-2132311102321203-3023021232321123-1010000322002131-1300101201102313-3030131102233110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333102203130100-2013132213221123-3103200013313330-0130121013030311-2201203132132001-2202021311111022-0123111200321232-3313110200303213"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.asn_list — asn_list / 213202300021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_list

<a id="canonical-0000230101202332-1213233221102311-2003030230120230-1220131331011221-0121233103200310-1321220302020002-1301023111311203-0201033120332332"></a>

Type: `"single"`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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

<a id="canonical-1313233033003113-0131313330323020-1001023223000323-3011312222222111-0023103020133122-1103211020300012-0302000220003010-1102210113013321"></a>

## Direct properties — asn_list / 213202300021 / 3

<a id="canonical-3302133103233233-1213131003211223-0332233123121300-3311023223310210-1013331131102012-2031023211001022-1301332330222202-1111130001120221"></a>

<a id="canonical-1331113331300200-0320210300320021-3011233010312033-1331202003033333-3221310111002232-0110112210320232-1033320321220030-2311113101022313"></a>

## as_numbers property — asn_list / 213202300021 / 4

Type: `["list", "number"]`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1332003200113331-0210032121033331-2221033230320103-1320003103013300-2303211132222202-2323330111021222-3320033220203021-1302020021122012"></a>

## Next pages — asn_list / 213202300021 / 5

- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0330233120213000-3202021203001020-2331313312313330-0111303031020112-0200121231020322-2312023231332312-0320031122231211-2313101031013232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033300121030322-0323312000131200-3202223122022122-3113130021121201-1311113100302321-0022332130131222-3110221310231100-2212212020231313"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher — asn_matcher / 023223231331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher

<a id="canonical-3323310001133032-2230312322320222-2212321133321332-0002330001321223-3101032300321330-0030200102112322-0102010300322123-2032120211301101"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

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

<a id="canonical-0033203230312121-0003320333003100-0300010033100313-2000130021223232-2113313332223021-1212200221221012-3022331223203120-1112102222220220"></a>

## Direct properties — asn_matcher / 023223231331 / 3

- [asn_sets](data-sources--http_loadbalancer--reference--group-006.md#canonical-3212013222131220-1232133311311113-1331310333133013-2302202001133013-0231000113321201-2203222221320032-0202103223110230-0310003021233122): complete subsection reference.

<a id="canonical-1110212023300100-2223323122211112-1022320333222121-0130110010022222-0232331133203320-1010220020300322-0210313230131010-0013030010003301"></a>

## Next pages — asn_matcher / 023223231331 / 4

- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets](data-sources--http_loadbalancer--reference--group-006.md#canonical-3212013222131220-1232133311311113-1331310333133013-2302202001133013-0231000113321201-2203222221320032-0202103223110230-0310003021233122)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3212013222131220-1232133311311113-1331310333133013-2302202001133013-0231000113321201-2203222221320032-0202103223110230-0310003021233122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102120300200100-3322200202101200-2123212213202023-3023221123203001-0232030120033231-2323301101200023-0203323322123310-3011123003003233"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets — asn_sets / 220222233131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0330233120213000-3202021203001020-2331313312313330-0111303031020112-0200121231020322-2312023231332312-0320031122231211-2313101031013232)
- api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-3233333231210211-1123111220011130-1002022021221321-2311102202230212-0233113323003002-1120131322312102-3201123021222301-2011112030300012"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-2030201132233213-1130332030231330-2111232011113202-2031013331020323-2020010112020130-0203333121020203-3210000231000220-3322222222222220"></a>

## Direct properties — asn_sets / 220222233131 / 3

<a id="canonical-1133323322102120-1331021300311323-2030210110020121-2012010330200032-3333002131211001-1020101311033232-3001333210133222-1132021212022300"></a>

<a id="canonical-2100133110111332-0323232130011113-0230223023313322-3232110001333220-1310302333130012-2212123023101013-2031221002203322-2300220102002101"></a>

## kind property — asn_sets / 220222233131 / 4

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

<a id="canonical-2233032110013301-3130202303300111-0223001120031130-3300203002032321-0120012121233311-3210300322223213-0222331122031102-2130211101010233"></a>

<a id="canonical-3031210330002330-1110001210133331-0110001313212101-3033220332332020-1013201211000010-2302133202333220-1123333311310223-0121212101311323"></a>

## name property — asn_sets / 220222233131 / 5

Type: `"string"`. Computed.

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

<a id="canonical-1322121322311232-0131201001131311-0311031301311121-0103313331231030-2331210303130332-2323031001231103-2321113002331133-2200100233122100"></a>

<a id="canonical-1212212200101011-2132321122032023-3131022300132330-2321311012233132-3211323201102021-0233121300123220-1200233123023220-1332120023001022"></a>

## namespace property — asn_sets / 220222233131 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3033220032321222-3002311030123230-0031002131323300-3022220211333233-0332211210103022-3301031000003033-0131120333301330-0022002320230011"></a>

<a id="canonical-1100233133213200-0331122201102312-2321123112321210-2122321301302000-3011321011032332-2103323022020002-0001000321300113-3121113203310322"></a>

## tenant property — asn_sets / 220222233131 / 7

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

<a id="canonical-2021122013003033-3233032123201110-1031113003133103-0112123212121200-1211333300200112-1310021011003120-1032003002302332-2311320002230112"></a>

<a id="canonical-3032212233313210-2221010213022101-1211230202231123-0002113013123202-1301121021012321-1203220230021312-3101100101113213-3000011023220033"></a>

## uid property — asn_sets / 220222233131 / 8

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

<a id="canonical-2112100311031123-3002120310113020-3200311112303021-2203233322312221-3310320010323220-1301203103120233-2303131203031302-3211220112303133"></a>

## Next pages — asn_sets / 220222233131 / 9

- [api_rate_limit.api_endpoint_rules.client_matcher.asn_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0330233120213000-3202021203001020-2331313312313330-0111303031020112-0200121231020322-2312023231332312-0320031122231211-2313101031013232)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1231020233021330-1100132112121213-3133223330101231-0213212332020011-0210001210131101-2232020012022012-1223013211020113-1302301102101122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310300021232110-0130113320321010-0132210123023232-1010121231312333-1020203103133302-3223200132331201-3010210100313210-1013232130101230"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.client_selector — client_selector / 012120301012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.client_selector

<a id="canonical-2231100330101010-0000012002322332-2020333221320011-2102101203013222-2313133200330033-3023213320103030-1211003233300101-0220310012223303"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2213021212211013-1202210320312301-1203112000200303-3130112003021323-3030202023213000-0120213212303102-2002323003213311-0213231232322303"></a>

## Direct properties — client_selector / 012120301012 / 3

<a id="canonical-2233322313022133-1121102010323132-0221211133011213-1212202203131201-3010322003033101-1312011211331122-3133133020010333-2221311330321333"></a>

<a id="canonical-0313023022120203-3033212101121331-1201303321211301-1301203032000222-2202110111131301-1202223200313033-3032132232222320-0302030131003110"></a>

## expressions property — client_selector / 012120301012 / 4

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

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

<a id="canonical-3032032302122331-2222220302020213-1201320023023022-0200323033021212-1000203332303002-0010013000202202-3022300020202030-0002312013102310"></a>

## Next pages — client_selector / 012120301012 / 5

- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0310023002110200-2232121210033223-1101202102102211-2011130300322220-0001201131202011-3123321330013330-1332300103330120-3111213331310031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302111310321003-1132300023333113-2121231302210332-2303330320012002-2130112000312012-3330310001130322-2010321310202323-3222102023321230"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher — ip_matcher / 103302010220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher

<a id="canonical-1211101301031223-2032223202230130-0111033202222132-3100113100023201-2100303110323023-3120133132303330-0013132323100032-3003021111012030"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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

<a id="canonical-3230000300300320-1103212122111320-1221301332001223-0113322113323210-1211103332332310-3300332122002103-2100120213211132-1321101202100233"></a>

## Direct properties — ip_matcher / 103302010220 / 3

<a id="canonical-1300321330102321-1022203133233312-2100110101000030-3200022220223033-0012330332200120-1311031002232113-2312032131330231-0321120121211130"></a>

<a id="canonical-3311021213332010-0030203202331002-3232131303000030-0220031113113320-0303132231122110-0333220210203110-1132031322021022-2231023020103111"></a>

## invert_matcher property — ip_matcher / 103302010220 / 4

Type: `"bool"`. Computed.

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

- [prefix_sets](data-sources--http_loadbalancer--reference--group-006.md#canonical-1010010231003131-1100312201102222-2330212030011333-2111230333230330-2200130112301013-1132032002120121-2310303121322133-1322310330200023): complete subsection reference.

<a id="canonical-2013010303302033-1221130223011033-2333223212110112-1321003313333133-2232102223100211-2132312123031231-2210323232023332-2032230101110023"></a>

## Next pages — ip_matcher / 103302010220 / 5

- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets](data-sources--http_loadbalancer--reference--group-006.md#canonical-1010010231003131-1100312201102222-2330212030011333-2111230333230330-2200130112301013-1132032002120121-2310303121322133-1322310330200023)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1010010231003131-1100312201102222-2330212030011333-2111230333230330-2200130112301013-1132032002120121-2310303121322133-1322310330200023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023012000003023-0323331032123100-0001000333101103-1222310322333101-1332100330121130-0011010231322101-3211131033101121-0322210210230020"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets — prefix_sets / 222123311230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0310023002110200-2232121210033223-1101202102102211-2011130300322220-0001201131202011-3123321330013330-1332300103330120-3111213331310031)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-2302213321301223-0332321112323200-0113222222022320-1122322223301213-0330120221223202-3320333123103332-3212010210120201-2332022110232233"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-0221330302300333-3300100013330312-3232033310110302-2320310310320221-3120000021121022-1013023310222302-1012312211133010-2213200202023230"></a>

## Direct properties — prefix_sets / 222123311230 / 3

<a id="canonical-2001003113100103-2200230203233133-0311212120320031-2231330301110301-1331222203101221-3223200233131023-2330213301111231-3232011203212031"></a>

<a id="canonical-1000120023120133-3231323213313130-2021232002212130-2320003013322120-1102020202333313-0112233301103023-3112013301023222-1333333232001203"></a>

## kind property — prefix_sets / 222123311230 / 4

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

<a id="canonical-2013100020112322-1133202131332222-3200003230310131-0132101223302232-2322230330022103-1010013311221032-0203223300103222-3010102031121213"></a>

<a id="canonical-2012221120233123-3221020000300233-3003133211303031-3001011130201213-0233112230111022-2202301012000113-1230220020020102-0331302201321321"></a>

## name property — prefix_sets / 222123311230 / 5

Type: `"string"`. Computed.

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

<a id="canonical-2311300100332131-1021311132130222-2210200113323311-3313120200312311-3330330122002333-3333032313201230-3201021211200303-1212010033311121"></a>

<a id="canonical-0031020113031011-1120100031013103-3200021203330101-0212201223122223-2131321221121013-3000010201000030-0020303220001313-1322101101012021"></a>

## namespace property — prefix_sets / 222123311230 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3212312333332212-0303311201230312-2013120132122022-2010021012123323-1021003100213023-0010123212313120-2311213321030300-0210310023331020"></a>

<a id="canonical-3012331100311222-1100131110300123-2201301112030220-2010213022223102-0202233031003200-0333202132120331-3030001201102333-3330331001321002"></a>

## tenant property — prefix_sets / 222123311230 / 7

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

<a id="canonical-0200013001120301-0231320132012330-3003001001203213-3120122121202122-3002311330232201-3231320100202330-2201102220223120-0231013133333112"></a>

<a id="canonical-0310112331300121-1020213311100100-2111300323123102-3232031002130003-3321200201132131-0322112130110312-2303232210312131-3213000100020131"></a>

## uid property — prefix_sets / 222123311230 / 8

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

<a id="canonical-0302300310312201-3133202131212330-3003322230312113-0322100332002333-0322331323032133-2330033221100113-0310322013111033-1330033331320103"></a>

## Next pages — prefix_sets / 222123311230 / 9

- [api_rate_limit.api_endpoint_rules.client_matcher.ip_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0310023002110200-2232121210033223-1101202102102211-2011130300322220-0001201131202011-3123321330013330-1332300103330120-3111213331310031)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1010101230032120-1332030020203310-0310303231312013-0313113302011012-3301123232320113-1222332331202232-3310202312201332-2322133301313021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332321033212312-2333300030012120-3230200112012120-2303222012000321-0002333302223132-1101220300123132-1123200100100022-2113331321103301"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list — ip_prefix_list / 231221233111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_prefix_list

<a id="canonical-3303321331003211-3301103000213110-1332203031022320-3022233230113033-0302223021030313-1320132120222232-0013320000111203-1202021200100031"></a>

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

<a id="canonical-2032320210032333-2212002021022223-0111002321223231-3122321031112202-1201231322322201-2332322003112022-0202301111031320-0221303001012330"></a>

## Direct properties — ip_prefix_list / 231221233111 / 3

<a id="canonical-0113121122013110-1123033320322120-3121001003323110-1210030331333122-1022202210112213-2010111102303012-1120230330232003-0000110113112231"></a>

<a id="canonical-2101011232202331-1032211031301330-3310211223132021-0312023031303233-1001023102020330-0113220102032311-0313113110220111-3131110233302332"></a>

## invert_match property — ip_prefix_list / 231221233111 / 4

Type: `"bool"`. Computed.

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

<a id="canonical-1101223331102312-0231133012312110-0213133113031021-1331232011201213-0011013310322012-1010003023322220-0332333222321220-2322032103201223"></a>

<a id="canonical-3123212033020013-2010031231330013-0110233201222332-2200302211121220-0213222132232323-1011301133221333-3201320323321320-3321231320021123"></a>

## ip_prefixes property — ip_prefix_list / 231221233111 / 5

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

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

<a id="canonical-3220232300330231-0221032203023201-0310221032023012-1100121032131231-3231033333100020-2130312013031011-1301230000011121-2323023103221220"></a>

## Next pages — ip_prefix_list / 231221233111 / 6

- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3202230111330110-1103220101312101-0133122333132123-0132002123120000-1031320313310022-1321201102030001-2022113003101022-0130230103232032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330022000210133-3101003321232333-2303321100312300-2120231022233001-1321033303020233-0312010011312302-1120003221113021-3032220232332033"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list — ip_threat_category_list / 112311023323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.ip_threat_category_list

<a id="canonical-2320313231001003-1210211111021223-1311301030300302-2200102331211333-3013103130310121-3012102001031111-3113212322100010-2200122002331230"></a>

Type: `"single"`. Computed.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

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

<a id="canonical-2001030323121013-2323012231321013-3210012011022321-2322302231003122-2111110103030130-2012012212021001-2121001302101300-3123130213223022"></a>

## Direct properties — ip_threat_category_list / 112311023323 / 3

<a id="canonical-0131201310303322-2102231101012133-3300000122101003-1310032222203023-2302230123301202-3021203021320213-3331223313130203-1002030121310100"></a>

<a id="canonical-0021333131303112-1022003121113021-3223310332012031-2001232212211223-3110110100303312-2311020313113331-1022212110220211-0201200023320231"></a>

## ip_threat_categories property — ip_threat_category_list / 112311023323 / 4

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3230112020031322-0221330001000023-0000222110311230-2213002231112211-2132231301210133-3011200103112001-2213023330012030-1030301132221320"></a>

## Next pages — ip_threat_category_list / 112311023323 / 5

- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3121101201010002-0110102323110222-0122303223121031-2202001231013321-2000132332121310-0111312010022200-3030211033013112-1220311323120003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001123321120332-3211022301211320-3010012302210213-2220010311032021-1132331013133130-1120021000112023-1102230310321002-3002302011322122"></a>

## api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher — tls_fingerprint_matcher / 123333221120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- api_rate_limit.api_endpoint_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-2011123120323103-2331221313112233-1312212201010023-3132021022302100-1202000010021010-2222010010211301-1131310133301332-1203133132000321"></a>

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

<a id="canonical-0320113330011323-0310213112221301-2010022303132110-0313303332303202-1323122021330233-2112003313300203-1220231313123113-2301220010221002"></a>

## Direct properties — tls_fingerprint_matcher / 123333221120 / 3

<a id="canonical-2302202231123311-2232031000213010-0223030101113103-1133220022222030-1321020310132302-0212300100300131-2133102203103201-2312002310022302"></a>

<a id="canonical-2213002103032322-0321202121330030-3123022021331232-0301231332333021-1122313113001222-2201322000113323-1033223201121210-1113201020330322"></a>

## classes property — tls_fingerprint_matcher / 123333221120 / 4

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3210310322303021-0123110013301110-1032001102213020-3302300311222011-3200102020133201-0010203001113320-0301111022033230-0101303012000310"></a>

<a id="canonical-2231012331022030-2221013023002313-1323010302232011-1221020202330112-2003100301130111-3212012113032222-3333330211122011-0100221132330310"></a>

## exact_values property — tls_fingerprint_matcher / 123333221120 / 5

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

<a id="canonical-0132301321203131-1230303202210220-3232322120331000-0201033120013231-3000111212110203-2130120311111203-3321021313323100-0001022320212230"></a>

<a id="canonical-3001121112031320-3203330230010212-2323310130103320-3223200011033021-3233220002302330-2313020032202312-2003110321022102-0020330021221101"></a>

## excluded_values property — tls_fingerprint_matcher / 123333221120 / 6

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

<a id="canonical-3033102030311221-3333320030132203-2201210330023132-0322323001133201-0021231302211101-0301120031032301-2030103230331210-0203100122010101"></a>

## Next pages — tls_fingerprint_matcher / 123333221120 / 7

- [api_rate_limit.api_endpoint_rules.client_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0101232003133323-1221310232031313-0230203131130103-3023111110132221-3023003123302121-1013102123120323-0200222020100312-3103220323103000)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1001100222001320-3011330202313020-3310020021000111-2303330102222022-2200220012031211-1333313010120013-1301330312003223-2002111303232321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012103021322102-2003230212000322-3200112203332102-3013223212332110-1211213130201121-2222131021221003-1322033012221033-1301000321030313"></a>

## api_rate_limit.api_endpoint_rules.inline_rate_limiter — inline_rate_limiter / 221132211131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter

<a id="canonical-3101102021110000-1132112123313200-3231300331121113-1321103033230323-2002333322321111-2231023000313023-1321222103032222-1223103202010320"></a>

Type: `"single"`. Computed.

Configuration parameter for inline rate limiter.

Upstream description:

Inline rate-limiter settings for this domain, base-path, or endpoint rule. Select this field as the
required rate\_limiter\_choice when no stored rate-limiter object is used.

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

<a id="canonical-3212313223112112-2232332033213223-0033032212110301-2311030101022101-2023120211002203-2303111122302322-2020310131022100-1300332023201130"></a>

## Direct properties — inline_rate_limiter / 221132211131 / 3

- [ref_user_id](data-sources--http_loadbalancer--reference--group-006.md#canonical-3101320012233311-3102220101011220-2210033131133032-2112123321312233-0021210032210201-0302211202333120-0021220122312233-2322221330233210): complete subsection reference.

<a id="canonical-1100311300330133-0120331011031011-0310010121001311-2303022133021331-2303200123202012-1323132120212331-0201310133123332-0100000313022302"></a>

<a id="canonical-3103330310111213-1231010130310321-3000321303111132-0311002233101230-2122213033003312-1111202310210103-0123122310112020-0233000312202112"></a>

## threshold property — inline_rate_limiter / 221132211131 / 4

Type: `"number"`. Computed.

The total number of allowed requests for 1 unit (e.g. SECOND/MINUTE/HOUR etc.) of the specified
period.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1032131030202303-1030230003133022-0101332112200100-0303201113133211-1103302102101022-1232102100331310-1133302331203013-0201020033212122"></a>

<a id="canonical-2211101030111300-2301110001313223-2030003023001122-1021311033321200-1013121101000221-0232300010033202-2320103211322221-1133112313203013"></a>

## unit property — inline_rate_limiter / 221132211131 / 5

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

- [use_http_lb_user_id](data-sources--http_loadbalancer--reference--group-006.md#canonical-1021330133002321-3013030230000202-2121113211201030-3223010132103323-3120101212231120-3231122003202032-0000122003312312-0232330201011332): complete subsection reference.

<a id="canonical-3231222201311031-3320220002323222-2232322223013000-1130121202311123-0012222120222130-2101331222103312-3330203003010122-2023100020211303"></a>

## Next pages — inline_rate_limiter / 221132211131 / 6

- [api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id](data-sources--http_loadbalancer--reference--group-006.md#canonical-3101320012233311-3102220101011220-2210033131133032-2112123321312233-0021210032210201-0302211202333120-0021220122312233-2322221330233210)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id](data-sources--http_loadbalancer--reference--group-006.md#canonical-1021330133002321-3013030230000202-2121113211201030-3223010132103323-3120101212231120-3231122003202032-0000122003312312-0232330201011332)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3101320012233311-3102220101011220-2210033131133032-2112123321312233-0021210032210201-0302211202333120-0021220122312233-2322221330233210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022010200211310-1230333002131232-0132011203230203-3002210333012001-2312332110303230-2302221011312302-2002031332100031-2122113012111213"></a>

## api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id — ref_user_id / 221030201020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-006.md#canonical-1001100222001320-3011330202313020-3310020021000111-2303330102222022-2200220012031211-1333313010120013-1301330312003223-2002111303232321)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.ref_user_id

<a id="canonical-1321320201211031-0322221313320302-3232330103131031-2220003023013133-0333210200210110-1100012313132303-0311303013101311-2211030303311033"></a>

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

<a id="canonical-1013103002023302-1302111130011020-3203333111133030-1221211301112013-0022102202132030-1002223010200112-0011211330023220-0010312302203302"></a>

## Direct properties — ref_user_id / 221030201020 / 3

<a id="canonical-1122331202032113-1010001311210321-1121002130200320-0323223202003230-3312022111030003-0130121121131022-1012220321030232-3222131301221002"></a>

<a id="canonical-1231310103330233-3001100211222033-0112210212233303-1113231301200133-1331320200033301-2313332301212211-1333231333333233-1100210103200213"></a>

## name property — ref_user_id / 221030201020 / 4

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

<a id="canonical-1003222231002230-3202013030111130-1021022123020133-3132021212020313-2003210323223322-1223223013210232-0021032120130303-1331202131010221"></a>

<a id="canonical-0313312133012223-0132022202231313-1000012221320133-0331231121101001-1133112030331230-2212320002230210-3200301103232323-2200212031330120"></a>

## namespace property — ref_user_id / 221030201020 / 5

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

<a id="canonical-0311100201121202-0322303021220200-2213320001321210-1012303012331103-3211322131331233-1220130002013100-2330301231213133-2200121103103222"></a>

<a id="canonical-3033020122320101-3300001020030211-0331230012313330-0223110121203020-2300302133013002-2323031130101322-0013330300200320-0312021130011301"></a>

## tenant property — ref_user_id / 221030201020 / 6

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

<a id="canonical-0033303103131222-1022102201003211-1203223021013312-0033031333330022-2231222303322133-0223032222222203-0110230010013101-1113332303220002"></a>

## Next pages — ref_user_id / 221030201020 / 7

- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-006.md#canonical-1001100222001320-3011330202313020-3310020021000111-2303330102222022-2200220012031211-1333313010120013-1301330312003223-2002111303232321)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1021330133002321-3013030230000202-2121113211201030-3223010132103323-3120101212231120-3231122003202032-0000122003312312-0232330201011332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010303300300222-3231103102201112-3301110020303211-0212211222222211-2200220123103213-1211222313221133-3211010230232302-3122113013012212"></a>

## api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id — use_http_lb_user_id / 310311302122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-006.md#canonical-1001100222001320-3011330202313020-3310020021000111-2303330102222022-2200220012031211-1333313010120013-1301330312003223-2002111303232321)
- api_rate_limit.api_endpoint_rules.inline_rate_limiter.use_http_lb_user_id

<a id="canonical-2021332132333023-2003121012232101-1331012133202220-0310010101322211-1110302231021032-2322102111130220-1030201232221312-2133012310230031"></a>

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

<a id="canonical-0103021200213312-2100120133300303-1301333132310000-3131201130333232-0213102203320310-3012300033033203-1312000100220320-2113022021322101"></a>

## Direct properties — use_http_lb_user_id / 310311302122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300223022300002-2233020000101200-0002333322232132-0002302120020001-2003121123021211-1100303000000102-1133233120232120-2210233201213212"></a>

## Next pages — use_http_lb_user_id / 310311302122 / 4

- [api_rate_limit.api_endpoint_rules.inline_rate_limiter](data-sources--http_loadbalancer--reference--group-006.md#canonical-1001100222001320-3011330202313020-3310020021000111-2303330102222022-2200220012031211-1333313010120013-1301330312003223-2002111303232321)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3312223301223300-0100320231121331-2200103031212013-2133203231122320-3200220223312311-1322011023113302-0321210220333333-0013201011200123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103210001111322-1103233212032120-3013110032123000-0303323331201221-0101030213333123-2003212121232223-0301331110321203-3101102200213111"></a>

## api_rate_limit.api_endpoint_rules.ref_rate_limiter — ref_rate_limiter / 103203111232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- api_rate_limit.api_endpoint_rules.ref_rate_limiter

<a id="canonical-0031201023112111-2313112300033000-0021222010333331-3313312020233223-1231321320200220-3020123211200210-3111123021131321-0313301202102011"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

Reference to a stored rate-limiter object for this scoped rule. Select exactly one of
ref\_rate\_limiter and inline\_rate\_limiter.

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

<a id="canonical-2022020011000302-1312301230221111-0322313002122102-0311000122213222-3301021320000130-3110311023032110-0322312020200023-2220103220213232"></a>

## Direct properties — ref_rate_limiter / 103203111232 / 3

<a id="canonical-3112301123233322-0223010032220032-0132323133033123-3011203130321123-2331303332312133-0133133303331101-2002102311212302-0302320120232310"></a>

<a id="canonical-1311113102320010-1002301223200003-0332330110320110-0001313221023132-1200331222221001-2233210220121112-3120330110200133-0131322103002121"></a>

## name property — ref_rate_limiter / 103203111232 / 4

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

<a id="canonical-0230210201322023-1030112313231031-1313312201120320-0322121301120213-2101101200331323-2122333321021002-1012132313322123-0002310100010321"></a>

<a id="canonical-2022131130232010-0203003210321230-0302200213212023-0321332110333122-0200113210303112-1022023122023021-3033000010320131-0303203013000211"></a>

## namespace property — ref_rate_limiter / 103203111232 / 5

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

<a id="canonical-1212000330233321-2210001022010302-0101133331220333-3120332323013121-1030100123321021-0122120211233323-2132313001002103-1220332213311210"></a>

<a id="canonical-1320231101002231-3300102302231310-0031213330203123-0230313111021230-1310123203322132-1003121312222030-2331010220112110-1322120121030202"></a>

## tenant property — ref_rate_limiter / 103203111232 / 6

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

<a id="canonical-2233011210223301-1303310212331030-1302212223331310-3302300201300030-0301220302303211-3312312222331220-2313101212310001-1201300211222133"></a>

## Next pages — ref_rate_limiter / 103203111232 / 7

- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0231103031100013-0310103221330221-0231022012323121-1213102220211202-0230012101113130-2100003112133013-0213101310021131-3323023321121300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222330311010323-1111020133313123-1022232021232001-0013023223203120-1310112112211230-1131202121102103-2031011303111020-0320120021013320"></a>

## api_rate_limit.api_endpoint_rules.request_matcher — request_matcher / 210133313332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- api_rate_limit.api_endpoint_rules.request_matcher

<a id="canonical-2132122232022331-1001301233323123-1001310002013103-1213122310203100-1101112132302013-3032002110000122-1120201130103203-2103032100200022"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1311223312003302-0222212121230100-1232101312202021-1013031033223122-0213021201032131-0231232001030011-1212132130311211-2131201320332033"></a>

## Direct properties — request_matcher / 210133313332 / 3

- [cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-1033221121320333-3213321013002333-1033130300011203-1030233202231103-0130103211211010-2331022223033032-2030011200201112-3100213211210122): complete subsection reference.

- [headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-3311311301133222-2232131331320202-1333103131333221-1011221031032012-1301233320102022-1213232111320323-3233033333121103-2213120210302323): complete subsection reference.

- [jwt_claims](data-sources--http_loadbalancer--reference--group-007.md#canonical-1210330031123323-3321033310123200-3323221033120330-0302113212111112-1231012323312022-2333330011003321-0211002010211131-1233302200112310): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-3102022311222332-0111210032013011-3001010132122001-2311201033202222-0200210203210312-1110200201313233-0102223200212201-3130023112023102): complete subsection reference.

<a id="canonical-1023120003230121-3012130022303332-2133013112020301-1231112202311333-2112323222213102-0301022311033100-1223010031330123-1231222223230011"></a>

## Next pages — request_matcher / 210133313332 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-1033221121320333-3213321013002333-1033130300011203-1030233202231103-0130103211211010-2331022223033032-2030011200201112-3100213211210122)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-3311311301133222-2232131331320202-1333103131333221-1011221031032012-1301233320102022-1213232111320323-3233033333121103-2213120210302323)
- [api_rate_limit.api_endpoint_rules.request_matcher.jwt_claims](data-sources--http_loadbalancer--reference--group-007.md#canonical-1210330031123323-3321033310123200-3323221033120330-0302113212111112-1231012323312022-2333330011003321-0211002010211131-1233302200112310)
- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](data-sources--http_loadbalancer--reference--group-007.md#canonical-3102022311222332-0111210032013011-3001010132122001-2311201033202222-0200210203210312-1110200201313233-0102223200212201-3130023112023102)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1033221121320333-3213321013002333-1033130300011203-1030233202231103-0130103211211010-2331022223033032-2030011200201112-3100213211210122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031111103031300-0012210121113122-2021223122122023-1323030020121020-0311131320011321-3322202200011333-3120130013202133-0203223032213001"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers — cookie_matchers / 312111312212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0231103031100013-0310103221330221-0231022012323121-1213102220211202-0230012101113130-2100003112133013-0213101310021131-3323023321121300)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers

<a id="canonical-3332231002010300-3112202131223003-0111321012323030-3202021002032113-2310103311122133-3120200132110200-1233201313102123-3030013233321333"></a>

Type: `"list"`. Computed.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-0200012101010302-0110323123312330-3200132112022210-0231001132131300-1122030213222212-1223132133020123-0103333111313030-2333112110211233"></a>

## Direct properties — cookie_matchers / 312111312212 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-3310231102302311-0213330221333201-1121101220021221-3201012321331203-0002133303320111-0020312313312312-3030103011020111-2113332232333313): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-3221211201111101-3212123333300020-0002222300300310-2331133003231201-1220031310103213-1320333200313223-0110032121121220-1012003011101301): complete subsection reference.

<a id="canonical-0022302111331023-0231332322220321-3200123122212330-1300101012133132-1121011322231323-0013113200102000-3232033302022330-0210232213000012"></a>

<a id="canonical-3131121303100212-3002233003120033-0002222203203332-1021322013201233-2230133121333231-2322030013101333-2023301332312210-2323011131020003"></a>

## invert_matcher property — cookie_matchers / 312111312212 / 4

Type: `"bool"`. Computed.

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

- [item](data-sources--http_loadbalancer--reference--group-006.md#canonical-1230212311222221-3222232133221330-0210200113023121-1233210133201313-2130200110211021-1013113303110113-2030220201103211-0111200012132332): complete subsection reference.

<a id="canonical-0303012221210330-0332100022123331-0100033002221332-3122202122032022-2331220300211100-3112213120300130-3322322301133330-1312000201002303"></a>

<a id="canonical-2110300001212103-3211001221020311-3100121131302001-1322113330130131-0113223323132212-2201301100030033-0230001012022003-2011012300212213"></a>

## name property — cookie_matchers / 312111312212 / 5

Type: `"string"`. Computed.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0132332111033321-2231330232201102-3312331233321212-0212133003131212-0231003102113120-2023010100233203-3000203003011013-3111130013223001"></a>

## Next pages — cookie_matchers / 312111312212 / 6

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-3310231102302311-0213330221333201-1121101220021221-3201012321331203-0002133303320111-0020312313312312-3030103011020111-2113332232333313)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-3221211201111101-3212123333300020-0002222300300310-2331133003231201-1220031310103213-1320333200313223-0110032121121220-1012003011101301)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item](data-sources--http_loadbalancer--reference--group-006.md#canonical-1230212311222221-3222232133221330-0210200113023121-1233210133201313-2130200110211021-1013113303110113-2030220201103211-0111200012132332)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0231103031100013-0310103221330221-0231022012323121-1213102220211202-0230012101113130-2100003112133013-0213101310021131-3323023321121300)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3310231102302311-0213330221333201-1121101220021221-3201012321331203-0002133303320111-0020312313312312-3030103011020111-2113332232333313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322331300222012-0020021003131132-0123111023313010-1003131002010220-3213232132233132-1320120013030233-2023120123110120-1022332222033100"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present — check_not_present / 322202211203 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0231103031100013-0310103221330221-0231022012323121-1213102220211202-0230012101113130-2100003112133013-0213101310021131-3323023321121300)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-1033221121320333-3213321013002333-1033130300011203-1030233202231103-0130103211211010-2331022223033032-2030011200201112-3100213211210122)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-3201332203111220-3002022202301100-0121103211212032-0021321013003321-3331233103132332-1131001012322012-2310001130313311-0103301112133023"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3312032213020020-3210212312233101-3000023211230233-2310321101200020-1033121311020211-2020013330211201-0112131113313310-3102121110021131"></a>

## Direct properties — check_not_present / 322202211203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131302103021322-2322303331323032-2133220121003330-2202313131203321-3300002031221312-2333010233001101-2220110003321002-1001130332031103"></a>

## Next pages — check_not_present / 322202211203 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-1033221121320333-3213321013002333-1033130300011203-1030233202231103-0130103211211010-2331022223033032-2030011200201112-3100213211210122)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3221211201111101-3212123333300020-0002222300300310-2331133003231201-1220031310103213-1320333200313223-0110032121121220-1012003011101301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122231123230302-1130032023011002-0032232211200302-1020011122321311-2201312131001101-1121221321111102-1122120203220130-0203110321021232"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present — check_present / 323030113310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0231103031100013-0310103221330221-0231022012323121-1213102220211202-0230012101113130-2100003112133013-0213101310021131-3323023321121300)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-1033221121320333-3213321013002333-1033130300011203-1030233202231103-0130103211211010-2331022223033032-2030011200201112-3100213211210122)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-3102321200123330-1000110100003012-0001211023323020-1222230230201001-3031022123022101-0103200110232203-2323233331212330-0210322333332332"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3200101302202301-2223000031203013-3121310003313311-1103122221121021-2321020210000232-1223122200020211-2123032013302332-1110132131131322"></a>

## Direct properties — check_present / 323030113310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231002011202032-0232310332301031-3220201233212230-0001313012212212-1333112332011100-1320311210331001-3012322200311111-0330031313010222"></a>

## Next pages — check_present / 323030113310 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-1033221121320333-3213321013002333-1033130300011203-1030233202231103-0130103211211010-2331022223033032-2030011200201112-3100213211210122)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1230212311222221-3222232133221330-0210200113023121-1233210133201313-2130200110211021-1013113303110113-2030220201103211-0111200012132332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031123021012113-1001122100111122-3222131001101233-0132001001222332-2002032232333300-0133211023322002-2102330301213201-2231133211000012"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item — item / 203123233013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0231103031100013-0310103221330221-0231022012323121-1213102220211202-0230012101113130-2100003112133013-0213101310021131-3323023321121300)
- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-1033221121320333-3213321013002333-1033130300011203-1030233202231103-0130103211211010-2331022223033032-2030011200201112-3100213211210122)
- api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers.item

<a id="canonical-2121021021022331-2120231220302302-2203012123000021-3000323131202323-3203021312101103-0203323330201223-1332011313003313-0022322320323133"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1321320131200310-0010012330133111-3011223221123210-0030121313123311-3303310010321131-0022221301003012-1003023012330030-0221220021212002"></a>

## Direct properties — item / 203123233013 / 3

<a id="canonical-2013022102110212-0100030021232221-0303223103130310-2113213213202300-2331102302233000-3120312303030013-2030303101031211-2022112100210113"></a>

<a id="canonical-0310011200100121-3130121122011022-2122231031210301-3330331301123222-3331203012010122-3310323313233301-0101223320221333-1333021101012013"></a>

## exact_values property — item / 203123233013 / 4

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

<a id="canonical-1331203002011200-3303021230231220-0203221130122200-2320020202000103-1012321032011200-0301230130003332-3310131233031130-2020233032101313"></a>

<a id="canonical-2010131321201012-1010201021113220-1113200013000233-1220233001303220-0110032122202110-2023320312331201-3022333321033102-1233231210132103"></a>

## regex_values property — item / 203123233013 / 5

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

<a id="canonical-1322211031222121-1330221310020121-3013130332202122-0203233122000200-0112110333032211-1122013010111212-1330030313202211-2310000130222023"></a>

<a id="canonical-2301130120301002-0321320102122112-0300111003032333-0320220020323230-0131223230333322-3231110310010320-3130020210102101-3002123322003033"></a>

## transformers property — item / 203123233013 / 6

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0311020223302123-0102132002003030-1211233310102112-0031122200103020-3132311012201020-0200012302311032-2003111221212200-1021033302313233"></a>

## Next pages — item / 203123233013 / 7

- [api_rate_limit.api_endpoint_rules.request_matcher.cookie_matchers](data-sources--http_loadbalancer--reference--group-006.md#canonical-1033221121320333-3213321013002333-1033130300011203-1030233202231103-0130103211211010-2331022223033032-2030011200201112-3100213211210122)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3311311301133222-2232131331320202-1333103131333221-1011221031032012-1301233320102022-1213232111320323-3233033333121103-2213120210302323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122010101133233-1122012133200130-1122230133221031-3001031332122110-1113123113001122-2232012110211200-0222030332220103-3111132003232311"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.headers — headers / 203130222012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0231103031100013-0310103221330221-0231022012323121-1213102220211202-0230012101113130-2100003112133013-0213101310021131-3323023321121300)
- api_rate_limit.api_endpoint_rules.request_matcher.headers

<a id="canonical-0231230333210132-2201220020000001-2210032310133132-0033003220122131-1300002130221121-3302011132213200-0310213033210023-3201011320212002"></a>

Type: `"list"`. Computed.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-3233021220313221-0302023312330221-1133203333131233-2101010022323010-3030202020103110-2231000210330311-0300232313132233-0210211113211013"></a>

## Direct properties — headers / 203130222012 / 3

- [check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-2130012002321002-3110213230113212-3022032033210013-1100333130213202-3320200210112311-1233313203212213-0202202021133322-1230102202100023): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-0122222021302302-0311233331013010-1221313110013030-2022122001331022-1220022313230032-0201000303030101-0322211201012310-3222000002332000): complete subsection reference.

<a id="canonical-3021231312100312-2333213032221021-0333323332301200-0322033110131210-3133223333032330-1100120023013100-3100213220333310-3123332013002321"></a>

<a id="canonical-1103001132120022-0332203323112102-1032131013231220-2112132000232023-1203312113202023-2310213131200332-2002310321011133-3200132030302300"></a>

## invert_matcher property — headers / 203130222012 / 4

Type: `"bool"`. Computed.

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

- [item](data-sources--http_loadbalancer--reference--group-007.md#canonical-1303231222200032-2233231313120011-2221232023112232-2221311221120023-2133102301210010-1320102023310010-3310233003113211-2311211013202212): complete subsection reference.

<a id="canonical-1130120313220233-0230312021223032-3232231002013322-2312032300230103-0023011221002110-2302232311330301-0320312310222212-3333333330231232"></a>

<a id="canonical-3232010211022131-1333320222013002-1221001001121322-3103012332212103-1110120203101020-3012231313200310-3112001210113320-2203311031203303"></a>

## name property — headers / 203130222012 / 5

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0201102313113020-3112232100212102-2032312131100233-2201202210311320-1030131311333012-0012212011313131-1300000312203213-3232213223221122"></a>

## Next pages — headers / 203130222012 / 6

- [api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present](data-sources--http_loadbalancer--reference--group-006.md#canonical-2130012002321002-3110213230113212-3022032033210013-1100333130213202-3320200210112311-1233313203212213-0202202021133322-1230102202100023)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers.check_present](data-sources--http_loadbalancer--reference--group-007.md#canonical-0122222021302302-0311233331013010-1221313110013030-2022122001331022-1220022313230032-0201000303030101-0322211201012310-3222000002332000)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers.item](data-sources--http_loadbalancer--reference--group-007.md#canonical-1303231222200032-2233231313120011-2221232023112232-2221311221120023-2133102301210010-1320102023310010-3310233003113211-2311211013202212)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0231103031100013-0310103221330221-0231022012323121-1213102220211202-0230012101113130-2100003112133013-0213101310021131-3323023321121300)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2130012002321002-3110213230113212-3022032033210013-1100333130213202-3320200210112311-1233313203212213-0202202021133322-1230102202100023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211022111120233-2232303030030102-0111031233200312-3132311130000101-1312312311003301-0101021322131132-2101311001012001-2322132113023111"></a>

## api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present — check_not_present / 130312231202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_rate_limit](data-sources--http_loadbalancer--reference--group-006.md#canonical-2002311232200310-0223223232322212-1310112012110133-2230110010300211-2032113231213203-3002313310330121-0330233231320210-2012011332303311)
- [api_rate_limit.api_endpoint_rules](data-sources--http_loadbalancer--reference--group-006.md#canonical-3301320333023033-0003233031231311-2330213121012000-3232320331311301-1022331023023110-0322310212201223-0131301321030211-2022203020202202)
- [api_rate_limit.api_endpoint_rules.request_matcher](data-sources--http_loadbalancer--reference--group-006.md#canonical-0231103031100013-0310103221330221-0231022012323121-1213102220211202-0230012101113130-2100003112133013-0213101310021131-3323023321121300)
- [api_rate_limit.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-3311311301133222-2232131331320202-1333103131333221-1011221031032012-1301233320102022-1213232111320323-3233033333121103-2213120210302323)
- api_rate_limit.api_endpoint_rules.request_matcher.headers.check_not_present

<a id="canonical-3010132023231020-2113100312310121-0012030130000002-3313200312210313-3102000001033331-2313301010232110-0111003221022100-0011322102212123"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0030101211200012-3220331133112103-3232323201232120-2112111002220302-2221203020003123-1220311100313101-2232012103101100-3222022121320323"></a>

## Direct properties — check_not_present / 130312231202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332133213322313-0320130132121010-0201230300333232-2323310213220031-0332331032313211-2331002313122303-2032101131230133-1111313321130021"></a>

## Next pages — check_not_present / 130312231202 / 4

- [api_rate_limit.api_endpoint_rules.request_matcher.headers](data-sources--http_loadbalancer--reference--group-006.md#canonical-3311311301133222-2232131331320202-1333103131333221-1011221031032012-1301233320102022-1213232111320323-3233033333121103-2213120210302323)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
