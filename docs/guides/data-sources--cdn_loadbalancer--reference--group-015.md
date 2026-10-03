---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

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

<a id="canonical-3212312110032222-3201223330333221-1302331120033212-1021200200211322-3330330221333102-1010300231131221-2323122231013020-2320330300232320"></a>

## context_name property — exclude_violation_contexts / 113332321030 / 5

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

<a id="canonical-1101031011311112-2203022211320311-3322211133230212-0122120233331121-2312112222131303-2021013202131300-0132220102312222-1130230101000231"></a>

<a id="canonical-0021231113123013-3100002211212020-2310313110200120-3210010203131103-3031320120023001-3032121013220332-3133303012013221-3302220222220222"></a>

## exclude_violation property — exclude_violation_contexts / 113332321030 / 6

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

Upstream description:

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

<a id="canonical-1201002003110033-1230113111203000-1022210201122011-0312131132323130-2133221210221130-2102012230220132-0203113132222312-2031301102113213"></a>

## Next pages — exclude_violation_contexts / 113332321030 / 7

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0213303020130303-0230331320122323-0211120020113320-0022010303223201-0130203102101122-1331301121101010-1231310013111130-0321202303330200)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1132002310133120-2021321221110121-0021011222300122-1110320101123002-3122011332033011-1133200310300121-2333230123010330-1301000200133122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323322103201013-2330121130013321-3221013221330122-0031012133203123-2321203002023033-2232131311221213-3302121201220230-0000130213121333"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.metadata — metadata / 333331231232 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- waf_exclusion.waf_exclusion_inline_rules.rules.metadata

<a id="canonical-1112313103131012-2112101110312021-0321020020031132-1200111000132031-1210201222121232-0021031103221321-3200003201232303-0021102021222032"></a>

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

<a id="canonical-0032313113031010-2132120012111112-0111120231231322-2013312013330313-3112100020333021-2010233203310120-2313230031201230-3332303010203220"></a>

## Direct properties — metadata / 333331231232 / 3

<a id="canonical-0133210121121331-3013130210221303-1112032001023310-2303033220100311-0312123301200111-1212201310103200-0311332112213021-3220213332030120"></a>

<a id="canonical-0200011312112132-0302011321110200-0000123223213103-2313133100000023-0020222210211220-3213331011230030-3012222233001331-2211102201002020"></a>

## description_spec property — metadata / 333331231232 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1223331011121000-3111120222022021-1130031122101212-0222021323311103-3333322313033211-2013032321201321-3132102032033121-3333313012323301"></a>

<a id="canonical-1332131200013211-3202110101313131-1012330330201331-3111011123003310-2220130312011010-2133023120001110-3032231221322132-0333310032313101"></a>

## name property — metadata / 333331231232 / 5

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

<a id="canonical-3313211032312311-2122233021113113-0302100310301022-0030130010233133-2321201301132232-0321333221213000-1131133202211233-1021330020103202"></a>

## Next pages — metadata / 333331231232 / 6

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1032010213111131-0112203220133133-1023020203013112-3001302122333112-3332030132000301-3330222033303230-3313221321211113-1023220000033122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211211211210103-2131302133322212-2111213220233110-0113131121001021-3111211101001311-3000332321312132-0322210210110222-2113311112301101"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing — waf_skip_processing / 202003112220 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [waf_exclusion.waf_exclusion_inline_rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3121012110330020-0020001232122133-1322111211322231-0130222233012232-1331331230031313-2202201020101221-1002110001321110-0121223013222022)
- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing

<a id="canonical-3210100213303332-3313031222303332-3310223203100312-1023230030213122-3210202320201131-2230331321231112-2332303131112313-2013002030203021"></a>

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

<a id="canonical-2032221303220101-3203021301310012-2311320011001223-3110223300103103-2322201232021312-3101130002111010-2112111330032331-1112103233020130"></a>

## Direct properties — waf_skip_processing / 202003112220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231203130120022-3030013201030230-0303222033132221-0210000201232112-0221111103122012-1231001323322122-0132013230203111-2022221303330132"></a>

## Next pages — waf_skip_processing / 202003112220 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1130121322212301-0233202331310122-0223000222302322-3000313230122221-0311201233220303-2322123203200100-2311131032230100-2131233213010301)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0321010221131133-3112232201230130-1002303120223013-3111331130200211-2222333313020223-0221201123100030-2000031120023210-0330322333212330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312010112022031-0032313023023211-1303000020012231-2310000223233130-1303322332302310-3221210122133113-3031323103230023-0332032002213322"></a>

## waf_exclusion.waf_exclusion_policy — waf_exclusion_policy / 331331112302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- waf_exclusion.waf_exclusion_policy

<a id="canonical-2121321300001002-1213202131231302-2131220321001302-3130122001220200-2122301223312011-3022130111233110-3303110320100010-3301002120212311"></a>

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

<a id="canonical-3221232010110323-1331101001322310-1130313113323003-1113002201302101-3311001313031111-0031203111330120-2200011301211022-1112302022112131"></a>

## Direct properties — waf_exclusion_policy / 331331112302 / 3

<a id="canonical-3331030213310010-1231030210010012-0031012033321011-1030320101202313-3001012303332300-3033013100300011-0200131301122221-2210130230023003"></a>

<a id="canonical-1312022010203231-1211222310031203-1011011212211020-1031210103132230-2123331303111222-0132323220322032-1223000221312212-0113022322113301"></a>

## name property — waf_exclusion_policy / 331331112302 / 4

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

<a id="canonical-1332221010021312-0222232233031033-0131123321312032-2122303022221110-3131012123123211-2211103030111332-1102311301203233-2231031230331212"></a>

<a id="canonical-3201232321332212-2010120031201303-3311210303031233-3220003112132323-0111311233333131-0322112102123130-3231310202321223-1130202220223110"></a>

## namespace property — waf_exclusion_policy / 331331112302 / 5

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

<a id="canonical-3131130330100212-2202311003120320-0000110223120001-3222022130320120-2122102212303310-2132303013012032-1011001321210223-0313112100200133"></a>

<a id="canonical-1121231101013312-1021303320001101-0223010133312311-0311322133001010-0023313301330032-1320032332030202-3323211112002102-1303131021211332"></a>

## tenant property — waf_exclusion_policy / 331331112302 / 6

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

<a id="canonical-2110312032321133-1013102202221100-1231012232012101-3312221101010203-2323310020223031-3223212020203103-2330123123230030-2020012203232320"></a>

## Next pages — waf_exclusion_policy / 331331112302 / 7

- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
