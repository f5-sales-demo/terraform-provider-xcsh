---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0200002110210112-2223133333230001-0321112311120101-2211233031001021-1023110031102010-2202310132023022-0321311332213331-0211121332111233"></a>

## bot_defense_advanced_protection.mobile_only.mobile — mobile / 233103202110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.mobile_only](resources--http_loadbalancer--reference--group-013.md#canonical-1313221110001232-1000111131110001-2120333202122021-0101332112201311-3303230010022023-1222103103001210-1000210111303323-2003031202312033)
- bot_defense_advanced_protection.mobile_only.mobile

<a id="canonical-0232130003231201-3220210130102201-2122111202221120-1031030312020333-1021022130320200-1201000133223211-1310113312330121-0312213333201122"></a>

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
mobile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220110330332323-0022100313222133-1303023221001132-3322010202210333-0101010013330033-0002200003322313-3323231013320232-0320323020031132"></a>

## Direct properties — mobile / 233103202110 / 3

<a id="canonical-3000211301121102-3000223202232130-2223333303021211-2130332100303313-3322121231130023-0123201133231133-3330230133010232-0133312123020332"></a>

<a id="canonical-3131113222302332-0032113331203101-2012213333232132-3230112002331100-3011232230221311-2312311130022000-1002210113123333-1320221331023222"></a>

## name property — mobile / 233103202110 / 4

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

<a id="canonical-0321303331220312-3111212010130023-1233022303110231-1233123213322333-3031202030320120-1002103220302122-0010232003013311-1210130223312323"></a>

<a id="canonical-3331220013012132-2213303230220101-3023330210002331-3033212131210210-2223223313312120-0300200333212023-2233111222112123-2331003001333121"></a>

## namespace property — mobile / 233103202110 / 5

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

<a id="canonical-2302000202020022-1203001333221302-3113102203210333-1333312313113021-3203001020203121-0033003332103031-2103132311101332-2220133223221332"></a>

<a id="canonical-1020032010333211-2122032321023033-0001123201322333-1201333001021302-3003302210301300-1123013133201123-0232010321003031-0030132013102211"></a>

## tenant property — mobile / 233103202110 / 6

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

<a id="canonical-1023030012020010-3203313231102013-1303103212302312-2032331001312202-1333221121221300-2201213322130020-3300300333101112-0311133321232221"></a>

## Next pages — mobile / 233103202110 / 7

- [bot_defense_advanced_protection.mobile_only](resources--http_loadbalancer--reference--group-013.md#canonical-1313221110001232-1000111131110001-2120333202122021-0101332112201311-3303230010022023-1222103103001210-1000210111303323-2003031202312033)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133133323323102-2211100122320210-2310213330222110-3233011100201300-3021002033223233-1011131000323300-0230103122230112-2003312212221313"></a>

## bot_defense_advanced_protection.web_only — web_only / 201123030201 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- bot_defense_advanced_protection.web_only

<a id="canonical-2320021130120100-0133032320321101-0330301303303031-0110300000121012-1313332000311202-2221223013201220-0203123001320033-3001230100033211"></a>

Type: `"object"`. single nested block, Optional.

Web. Web only configuration.

Upstream description:

Web only configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages_except",
    "js_insertion_rules")}
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]"
}
```

Terraform syntax:

```terraform
web_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233233200001120-0301110012203300-0121223200002213-1232010123203303-1002230000011331-2300322322221003-2201232323122033-3103120133132103"></a>

## Direct properties — web_only / 201123030201 / 3

- [disable_js_insert](resources--http_loadbalancer--reference--group-014.md#canonical-2130220313321211-2102111000221003-0230003023002230-1022110112002202-0210133133123211-2020131103323312-0213032332130032-3121033302113020): complete subsection reference.

- [js_insert_all_pages](resources--http_loadbalancer--reference--group-014.md#canonical-0122213003032102-0300021333330122-1010112231112101-3320330201022123-2110002313133303-0022030322303010-2010202000233202-3012132210200330): complete subsection reference.

- [js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-2113112203330321-3303103230122333-0233222132021112-1213133102232030-1013130211111132-2221210301022210-2332303022031323-0222200102100113): complete subsection reference.

- [js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301): complete subsection reference.

- [web](resources--http_loadbalancer--reference--group-014.md#canonical-0032332312200101-3012322220030112-1320110202102312-0220201232102020-2202130130221013-0301003122332030-0331121213303131-0133333302230210): complete subsection reference.

<a id="canonical-0130020200231202-1202103122333200-2203310013131013-1133100230331113-3202311210001021-1232221203130232-0122232303030132-3001201100312123"></a>

## Next pages — web_only / 201123030201 / 4

- [bot_defense_advanced_protection.web_only.disable_js_insert](resources--http_loadbalancer--reference--group-014.md#canonical-2130220313321211-2102111000221003-0230003023002230-1022110112002202-0210133133123211-2020131103323312-0213032332130032-3121033302113020)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages](resources--http_loadbalancer--reference--group-014.md#canonical-0122213003032102-0300021333330122-1010112231112101-3320330201022123-2110002313133303-0022030322303010-2010202000233202-3012132210200330)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-2113112203330321-3303103230122333-0233222132021112-1213133102232030-1013130211111132-2221210301022210-2332303022031323-0222200102100113)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.web](resources--http_loadbalancer--reference--group-014.md#canonical-0032332312200101-3012322220030112-1320110202102312-0220201232102020-2202130130221013-0301003122332030-0331121213303131-0133333302230210)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2130220313321211-2102111000221003-0230003023002230-1022110112002202-0210133133123211-2020131103323312-0213032332130032-3121033302113020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203201223313020-0102232333113031-3133323132103111-2333200020232232-0131110012312031-2000103321002203-1321200130012203-3330211111300022"></a>

## bot_defense_advanced_protection.web_only.disable_js_insert — disable_js_insert / 011121022200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- bot_defense_advanced_protection.web_only.disable_js_insert

<a id="canonical-1201213102332322-2301233202011023-1131220320302012-3021113231203003-3001303033003301-3300233312120033-1313210111122020-3003131032311201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

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
disable_js_insert = {}
```

<a id="canonical-2202020322210211-2130202333113332-1113302123312200-2131130123101220-0203330332303110-3003002010323103-2002132122102323-2221131232121310"></a>

## Direct properties — disable_js_insert / 011121022200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323303222310223-3202103003002331-1010203312200212-2021201302213313-0122033132011213-0330011021033222-0022133000013031-1020013202202131"></a>

## Next pages — disable_js_insert / 011121022200 / 4

- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0122213003032102-0300021333330122-1010112231112101-3320330201022123-2110002313133303-0022030322303010-2010202000233202-3012132210200330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330013111020313-3212130311213111-0233302100221012-2030212212033123-1122130211000312-2302122322013321-0223332210211122-2103121332220011"></a>

## bot_defense_advanced_protection.web_only.js_insert_all_pages — js_insert_all_pages / 031221101301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- bot_defense_advanced_protection.web_only.js_insert_all_pages

<a id="canonical-0121323202031021-2133322013002231-1212033102002201-3120033301123111-3220232313030223-1220213210310220-0110331010113212-2232131001212033"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages.

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
js_insert_all_pages {
  # Configure direct properties listed below.
}
```

<a id="canonical-2211012120330333-3032221130131302-2202321031033132-0132233121112323-2132011020011332-0330000112321323-0100312200121122-0030022002300133"></a>

## Direct properties — js_insert_all_pages / 031221101301 / 3

<a id="canonical-2020111132120030-3210133222210330-3302002130132101-3021310330223300-2233010300123001-1130033322111131-3313013031230031-1122010001112123"></a>

<a id="canonical-1010003301333311-0231002322332032-1323311220103120-1103033213002302-2320122221101001-1332032012332000-1302110131221022-2223222123110102"></a>

## javascript_location property — js_insert_all_pages / 031221101301 / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2233131023332212-2220311313000212-1011133331131302-3233332332011301-3233323130201000-3012020000311020-0132112331030002-3212120230113221"></a>

## Next pages — js_insert_all_pages / 031221101301 / 5

- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2113112203330321-3303103230122333-0233222132021112-1213133102232030-1013130211111132-2221210301022210-2332303022031323-0222200102100113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132332102301022-2202222321110212-3332121201013230-3332332322200022-0131000002002030-1012321212222311-2021323130132210-3233131012130111"></a>

## bot_defense_advanced_protection.web_only.js_insert_all_pages_except — js_insert_all_pages_except / 220310211213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except

<a id="canonical-0321221331303132-0033121320311213-3101202032321212-2132010312203112-1022330322222131-1233012111122113-0220000103022123-3223212003331101"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages with the exceptions.

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
js_insert_all_pages_except {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333123023033102-2233032221210002-1232030201032211-1313113032103022-1203001132302233-0311103202210100-3010230003212230-0220232200012021"></a>

## Direct properties — js_insert_all_pages_except / 220310211213 / 3

- [exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230): complete subsection reference.

<a id="canonical-2032220013323220-1023102333021201-2321023110122020-2232320221301132-2021303312132103-2231120100322112-1233110223310033-2210213332100232"></a>

<a id="canonical-2332303301102201-0110030133001012-3332020000330333-3002031232120322-0131120210232132-0113323200132213-3011230101300023-1103331230311221"></a>

## javascript_location property — js_insert_all_pages_except / 220310211213 / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0113131330021032-1321222000023202-2032302311312231-1132122222213001-1213112132230313-1330002212212002-3201133300123201-2311010001233102"></a>

## Next pages — js_insert_all_pages_except / 220310211213 / 5

- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320023100301330-3023130300132312-3011313000030223-3023020231201130-2301030331230111-0011322111300003-2322330121333011-1222303322031231"></a>

## bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list — exclude_list / 013320212030 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-2113112203330321-3303103230122333-0233222132021112-1213133102232030-1013130211111132-2221210301022210-2332303022031323-0222200102100113)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list

<a id="canonical-2332000031113210-1032230132322013-0033100112312022-1003203323123102-1301000112222210-2313211233022221-2212012012230000-2122223313330100"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0010313121323310-3323221200311002-2221312113221210-0110310223133233-3132010002133330-3021312103323222-3232221330220000-0030033230133203"></a>

## Direct properties — exclude_list / 013320212030 / 3

- [any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-1111031202211010-3323001121320012-0130021100321300-2230112001313000-2023113210032031-1221312330230001-1022331132320301-0200112130020321): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-014.md#canonical-2231200103232023-1011223122003131-2200002033011013-0023233101132321-0311331112213033-2132103130103322-3301131033013033-3201203300233202): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-014.md#canonical-1110200221023030-3221021100012023-0013102202120010-1223210220031131-3303123230130212-2123133302300120-1331321123220130-0102201222112033): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-014.md#canonical-2111201011110321-1110000313021120-2210300001203111-0301100200223101-1222232001301202-3300100100301021-0322011020201020-2120000131322203): complete subsection reference.

<a id="canonical-1200322000313323-1222103020213321-1312312103032232-2033002332210223-3312000302220220-1211222303133320-1201022300310113-1002321000031000"></a>

## Next pages — exclude_list / 013320212030 / 4

- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-1111031202211010-3323001121320012-0130021100321300-2230112001313000-2023113210032031-1221312330230001-1022331132320301-0200112130020321)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.domain](resources--http_loadbalancer--reference--group-014.md#canonical-2231200103232023-1011223122003131-2200002033011013-0023233101132321-0311331112213033-2132103130103322-3301131033013033-3201203300233202)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.metadata](resources--http_loadbalancer--reference--group-014.md#canonical-1110200221023030-3221021100012023-0013102202120010-1223210220031131-3303123230130212-2123133302300120-1331321123220130-0102201222112033)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path](resources--http_loadbalancer--reference--group-014.md#canonical-2111201011110321-1110000313021120-2210300001203111-0301100200223101-1222232001301202-3300100100301021-0322011020201020-2120000131322203)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-2113112203330321-3303103230122333-0233222132021112-1213133102232030-1013130211111132-2221210301022210-2332303022031323-0222200102100113)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1111031202211010-3323001121320012-0130021100321300-2230112001313000-2023113210032031-1221312330230001-1022331132320301-0200112130020321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202033130220013-2221023130132120-0120113301131233-2300022203220321-1112130131110033-1111130222102330-3303100203210032-3203312220121011"></a>

## bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.any_domain — any_domain / 332303200323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-2113112203330321-3303103230122333-0233222132021112-1213133102232030-1013130211111132-2221210301022210-2332303022031323-0222200102100113)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-1002313001231332-2323013230112022-1023303122303002-2122112131033101-0100211200012212-0200310021230303-3122003023312103-0110230023312112"></a>

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

<a id="canonical-0202012123323211-2101130330003113-1300222311022012-3102330213103321-1112310103212031-2303220233333032-2220220223023332-0130021331222302"></a>

## Direct properties — any_domain / 332303200323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332300102121232-3321101212313023-2220322110130322-3210212212103021-1203101113002132-0130022001001322-1020131133302312-0222211131323001"></a>

## Next pages — any_domain / 332303200323 / 4

- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2231200103232023-1011223122003131-2200002033011013-0023233101132321-0311331112213033-2132103130103322-3301131033013033-3201203300233202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130120131320012-0021212232101303-2011302120101123-2032120110012103-3011300023132120-3303121213320301-0321210301212020-2022200113022203"></a>

## bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.domain — domain / 203113003123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-2113112203330321-3303103230122333-0233222132021112-1213133102232030-1013130211111132-2221210301022210-2332303022031323-0222200102100113)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-0212323120121122-3022103010030120-3103200010222322-3201230020200203-2013010131202132-2331221031031133-3102010103011233-1000011323332203"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020002003121312-2031130220302312-0313233320133001-0133302111211213-1010222300131031-1303012011101023-1110003300003031-3300231131111123"></a>

## Direct properties — domain / 203113003123 / 3

<a id="canonical-0131100322221233-2011301310203020-0133223312030001-3132110133103130-2133121131212230-0031101213132113-2110321133331230-1322213112310030"></a>

<a id="canonical-0232303112300221-3100330311000202-2010011103033102-3132331030333302-3232133111322012-2002330213203230-1112301333211101-1032011201233002"></a>

## exact_value property — domain / 203113003123 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-3112020122022223-2233133132110312-3200322212333133-2211112302100233-3333220201312131-3111210103333130-2212022102203000-3201301302123202"></a>

<a id="canonical-2320003201202100-2310130212033333-1200331102100133-3313301300200233-3331032132223330-3112220301202311-2000322020103302-1113300331010223"></a>

## regex_value property — domain / 203113003123 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3320030033111002-1333302131101121-2112113212321210-2003132210332121-2333332100320203-1002231130100313-3333112010002023-2320122313330022"></a>

<a id="canonical-0120330023221222-1311030200112121-2233320102002011-1332122233003212-2202133003002333-0310032330323222-1321122130022101-1233303322310331"></a>

## suffix_value property — domain / 203113003123 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1221013101021200-0132102032103222-2210330100123320-0230300002322332-0103322313011130-0121320131231020-1303303221310220-2302232211132011"></a>

## Next pages — domain / 203113003123 / 7

- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1110200221023030-3221021100012023-0013102202120010-1223210220031131-3303123230130212-2123133302300120-1331321123220130-0102201222112033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211032210101113-3010220101332011-2322230312303002-2121111230320021-0023102020101111-3121033010022130-1023003002311300-3203111332013331"></a>

## bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.metadata — metadata / 123323113332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-2113112203330321-3303103230122333-0233222132021112-1213133102232030-1013130211111132-2221210301022210-2332303022031323-0222200102100113)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-0033200220033023-3111222310212132-2210331121000220-1010113010103132-3230132320322123-0131012231203112-0002201130202322-2311310312230032"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032332233011320-1332303013013103-0302121100300233-1101031033230002-1212200202220333-0301101113003100-2131033330203102-0331220110233230"></a>

## Direct properties — metadata / 123323113332 / 3

<a id="canonical-3232322131320000-1032203021012030-2012332100030120-1031221113310231-1111030101130130-0001212222200203-2303313122233233-0110131202312032"></a>

<a id="canonical-0332332003303122-2233303032003312-1002033230300303-0113012221110332-3230331202100333-2123311121211021-1033101131121212-2322333011203120"></a>

## description_spec property — metadata / 123323113332 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2302212123230300-0123312332032313-3323101233100201-3110333111330213-0201233312123321-0322111112233011-1123223330232011-0222113301303131"></a>

<a id="canonical-1213020211203022-3222332013131030-3310231021010210-2330132132032231-3222303010110012-1130002132213220-0213310003002010-1201212302133322"></a>

## name property — metadata / 123323113332 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3012312100103011-0101032111213123-2221023320211111-3211013310110001-1002110220010330-3333032232032211-2211021003322320-1002210221130100"></a>

## Next pages — metadata / 123323113332 / 6

- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2111201011110321-1110000313021120-2210300001203111-0301100200223101-1222232001301202-3300100100301021-0322011020201020-2120000131322203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213002011020122-3313213310200010-0300103202103302-3322011210022303-2303033200022231-0322100131332211-0323320131033011-1222322233100222"></a>

## bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path — path / 300230001222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-2113112203330321-3303103230122333-0233222132021112-1213133102232030-1013130211111132-2221210301022210-2332303022031323-0222200102100113)
- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230)
- bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list.path

<a id="canonical-0303132232102310-1220032203332202-2313100201312221-0331001022211011-2002322013222013-3032123022121302-3331230122202002-0120312321002300"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013031020033310-2121331130021121-0230121310332031-1001122122100003-3220302021021323-2333022320110233-0221122203033332-3122111120213011"></a>

## Direct properties — path / 300230001222 / 3

<a id="canonical-0322300023320031-3100020322113101-1213210330133220-0010023132200003-0312200100301311-1200110332133300-3011322031003322-1302121330232030"></a>

<a id="canonical-0213211200132230-2022030113330201-1102100130000300-3210301211221230-3031302000330212-1300133312230112-0103231312300233-1101122220111130"></a>

## path property — path / 300230001222 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-2322302312131112-0310103132310331-3103312002102213-3031031301121210-1330333032132023-3032211000030033-0112030330101122-3000030001002301"></a>

<a id="canonical-2320103020312202-3322320322131221-1103002330032110-0332020220131122-1032222200100032-1211301123312323-0232122230102012-1020331313003331"></a>

## prefix property — path / 300230001222 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-1310212313001131-3122311320332202-2301321310312010-0002333132231211-0213111021232011-0010120310131023-1223132110203003-1130123112003321"></a>

<a id="canonical-3231110201312333-0020303322320012-2012123312222031-1131122323100213-2303021022021321-0112010302200323-1131000233022223-3211233031201102"></a>

## regex property — path / 300230001222 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1023211102311021-2312111130021112-2220231102033301-1302301233133020-3031003232321322-0211200021020220-2031231131311131-1322021312322230"></a>

## Next pages — path / 300230001222 / 7

- [bot_defense_advanced_protection.web_only.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-2300132023031330-2003233122233102-2101333020132232-0001122131020303-2203112203203212-0023312123200011-1202122032022020-3100021210013230)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130321033320020-0220023032313311-3103121131302201-0323110221231222-0303220230023310-3330222012212113-1110230101011212-2230323303220101"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules — js_insertion_rules / 332300231101 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- bot_defense_advanced_protection.web_only.js_insertion_rules

<a id="canonical-2311031300232101-3112032200333120-3211031000121011-1013103222113100-1010311333003000-0212033211010000-1203010210113312-0322312300023132"></a>

Type: `"object"`. single nested block, Optional.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313112100131223-2211133020122033-0103223201300123-2033021030120303-2012132223211132-3123320010111102-1100100133103201-0023312231021001"></a>

## Direct properties — js_insertion_rules / 332300231101 / 3

- [exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320): complete subsection reference.

- [rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303): complete subsection reference.

<a id="canonical-0202020002121312-0101013110212120-1322203310023101-0313221220231033-1222233312001002-3123023020301031-1221310110311103-1011022311200333"></a>

## Next pages — js_insertion_rules / 332300231101 / 4

- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031201022003110-0112021210310102-2021333002312200-3111002113010133-1010221030230211-2331303223211123-0003023120221313-3023112312223333"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list — exclude_list / 020130211122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list

<a id="canonical-3222332222303021-1002220010121332-1201303221300222-1131311220102122-1233010211310233-0123111033120030-0013311323112001-1221113212101013"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3002012021100321-1221310303300011-3101310110111322-0023220312211012-0131213220232112-0222101222130322-2331013203233012-3011331130110220"></a>

## Direct properties — exclude_list / 020130211122 / 3

- [any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-2011103010011123-3322332303233223-2320021013221023-1013330012013130-2223103000200000-1202032102332203-0122322222221113-0010112130122100): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-014.md#canonical-1131122213030023-0310320003132323-2032213312222222-1230011110032223-2111221023221312-3000012001330103-1133212233103221-2120031232010313): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-014.md#canonical-1011020020300202-2332203021211013-0023222120212201-0101123013013310-1311310301020303-3021200102103310-1103312210021013-3102330112113101): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-014.md#canonical-1112201310023113-1302120212211300-1210321011312300-2032300112230112-1003233022203030-2331212302112213-2313130121113332-2211112020112331): complete subsection reference.

<a id="canonical-1232331221030122-0213202202000002-2213133023022032-0220230202012222-0232311121113301-2030100212100132-0321321300331110-2223001231232123"></a>

## Next pages — exclude_list / 020130211122 / 4

- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-2011103010011123-3322332303233223-2320021013221023-1013330012013130-2223103000200000-1202032102332203-0122322222221113-0010112130122100)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain](resources--http_loadbalancer--reference--group-014.md#canonical-1131122213030023-0310320003132323-2032213312222222-1230011110032223-2111221023221312-3000012001330103-1133212233103221-2120031232010313)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata](resources--http_loadbalancer--reference--group-014.md#canonical-1011020020300202-2332203021211013-0023222120212201-0101123013013310-1311310301020303-3021200102103310-1103312210021013-3102330112113101)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path](resources--http_loadbalancer--reference--group-014.md#canonical-1112201310023113-1302120212211300-1210321011312300-2032300112230112-1003233022203030-2331212302112213-2313130121113332-2211112020112331)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2011103010011123-3322332303233223-2320021013221023-1013330012013130-2223103000200000-1202032102332203-0122322222221113-0010112130122100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200103301121013-2132122120212002-2332320003203100-0323233211301212-1020110130311301-0030221110212001-0322211212003321-2120101333111001"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.any_domain — any_domain / 302330113011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.any_domain

<a id="canonical-3322213322100020-1021322200323302-2032112201000133-1212300130032130-1023023112210112-0301000211030120-1032223333021313-3112202220313320"></a>

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

<a id="canonical-1333123332002202-1121233030121100-0120033223312301-3033020031330102-3221201110200013-0130200220222110-0031201133202033-1101120102232110"></a>

## Direct properties — any_domain / 302330113011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022023202332133-1320001121200223-2303032201310110-3323301303300113-1031021321332210-3021200202221031-3223221332203310-0210202233021012"></a>

## Next pages — any_domain / 302330113011 / 4

- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1131122213030023-0310320003132323-2032213312222222-1230011110032223-2111221023221312-3000012001330103-1133212233103221-2120031232010313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133330223010032-1330303222130022-0030133130120230-1020202132121001-1313123303323230-0331132213332323-1020311302223101-3131112223331332"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain — domain / 020220030222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.domain

<a id="canonical-2031320111230100-1101300032133131-0033320130220000-2200030101013121-0010130133122032-2321113211333101-0323023111311133-1333202311202112"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231032203012212-0021211122203032-2032102332002201-2122103313313213-2003011333210021-3311100231011130-3031023123012333-2303112302300212"></a>

## Direct properties — domain / 020220030222 / 3

<a id="canonical-3303003131321232-0222132233200222-3201013203021332-2202101330230213-2011203330312211-0003133233010033-0030030313320330-2221302312200102"></a>

<a id="canonical-0223330212032233-3300032233033321-0202331313110012-1100310211012113-0322001222230232-3201220313113123-0231331220133233-0000332332230101"></a>

## exact_value property — domain / 020220030222 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-0313212120131221-1131320332231033-3112121031021232-1102200113033212-3032100222023121-3223123233031030-0033320011123233-3333122031322033"></a>

<a id="canonical-3302130022130210-2100203303231321-0210122200200010-2033221320022011-3333111011213331-1100312120131322-3032113313202111-1200103013233130"></a>

## regex_value property — domain / 020220030222 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2221100332102103-1332302121220322-3111132323111230-1313023132113003-1333132120030010-3230311203320023-2033021111330210-3130322301001232"></a>

<a id="canonical-2220112201221300-3323213203332033-2213122030031211-3131013000320022-3320113321301002-3133021313113013-0111022132023021-1200030222010102"></a>

## suffix_value property — domain / 020220030222 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1313031021331001-1013121202113113-3200123320331310-3113202130231300-0323202223321103-1133022230033230-0223213003330120-2002131133331100"></a>

## Next pages — domain / 020220030222 / 7

- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1011020020300202-2332203021211013-0023222120212201-0101123013013310-1311310301020303-3021200102103310-1103312210021013-3102330112113101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020222201221131-2102132332211120-0012313021202202-2331233002230222-1131313133212201-3010210112021020-1233011112321213-0201322031131013"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata — metadata / 001303330002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.metadata

<a id="canonical-3001023303310003-0101303210201203-1121210003203311-0310233122100213-3330131201103112-3202220311311110-3010312002212303-1010302311330020"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210031331220231-0003300122101220-1122220232123210-2313022300122320-2103112111233303-1230222210232131-1210313323210111-3033132132213210"></a>

## Direct properties — metadata / 001303330002 / 3

<a id="canonical-1233003133220230-1232110011332210-1102000310132332-3203013112300213-3032100331311232-2222320323210112-1001123212322211-1310212320100131"></a>

<a id="canonical-2222120002103221-0333102123312023-2130133120033033-3310321210120010-1330001210020031-1320033121123033-3213300200221132-2021322122303223"></a>

## description_spec property — metadata / 001303330002 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3032003012021333-1302233030133312-0130013312231312-0112013133313011-2102132130022100-3122032333203322-0132021320033300-2223221233003020"></a>

<a id="canonical-1203131100013112-2131303200003130-1332223312332011-0023021021321000-1013110012121210-1230303231123310-1202303120330230-2210122233213111"></a>

## name property — metadata / 001303330002 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-0132131033132002-3233123331101121-3123121330333212-3213023203110211-1302133321102001-3002010123001012-1230113301110201-2130233212133103"></a>

## Next pages — metadata / 001303330002 / 6

- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1112201310023113-1302120212211300-1210321011312300-2032300112230112-1003233022203030-2331212302112213-2313130121113332-2211112020112331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100102221322223-0031103123330013-3310322212030322-0010031023130230-3332323332103132-2233213330133031-3030013130231313-0011220303330233"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path — path / 100311233312 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320)
- bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list.path

<a id="canonical-3130112331122121-2210112233222001-0010021322203002-2032300030023003-3003120200121130-1022132103133120-1103003231123101-3212330321120102"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013233013221232-0320131332322301-0001032202132021-1030001212121313-2331122311313320-0230022321230101-3012001010330311-1222033302131110"></a>

## Direct properties — path / 100311233312 / 3

<a id="canonical-0202322332130001-2011001222001022-3000021211301210-2000233300302123-2202010222112313-0200133030323000-0211323003230330-0312220313032210"></a>

<a id="canonical-2103200330033233-3020033213302021-2031223020203130-1001203000012303-0113030030313021-2000311332231302-0031200321122130-0022132002323303"></a>

## path property — path / 100311233312 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-1201322320211211-1013010300031201-2210230202332013-3320021211223022-2101103301123033-2032130321303122-3313221032320222-2233101001311100"></a>

<a id="canonical-0113321300310103-3011113330321232-0312323323332131-0230020131330303-0112003031112323-2001203020011313-0323013101321223-0200333000332233"></a>

## prefix property — path / 100311233312 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-1231032203230303-2201013122303103-2302213333210320-3222003121021322-0023110330320133-2332220331013030-0102222102202233-1132122302200113"></a>

<a id="canonical-2303032002131111-2031021120222103-1030121231013303-3320320302012123-0031233021121213-2231020022121323-0231111032021310-3020321003013103"></a>

## regex property — path / 100311233312 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0020233302133333-3132101113130021-1322101302033002-1310233133121010-1332322111233131-3113110312033032-1020300230031113-2310010020033303"></a>

## Next pages — path / 100311233312 / 7

- [bot_defense_advanced_protection.web_only.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-3312310033210221-1301011322030112-1320000130122120-3202221003310031-0031202130222222-3001331023332001-0110230131320030-0110303301130320)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322003031301323-1110311210230010-2332323310012331-1021021201031130-1031031231320201-3333032010030103-3020201000010200-0210200202011322"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.rules — rules / 130212132203 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules

<a id="canonical-0300001222100111-2001221002103222-3303100133210330-2131312101101032-2301131012011013-2331323300033121-1000033001103220-3211301003222132"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Bot Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312132330310303-2223221102111203-0031333332100200-2333001010033023-3220330023330032-2010013200232222-1011130002113023-3112033313203130"></a>

## Direct properties — rules / 130212132203 / 3

- [any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-1200331032112302-0133002312021221-0013103132010212-3331330232303022-0333122302131133-2023201122233322-1101010330031001-0321133001110013): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-014.md#canonical-1301233102201233-0233012311103202-1023210002020011-2233110231333123-1300111222303013-1031033030030211-0220012203013020-3110201010213100): complete subsection reference.

<a id="canonical-0032102021221300-1233311231220310-1211323321213221-0021310100213020-2120330103113022-0201111311133222-3032123133012203-3332233313003212"></a>

<a id="canonical-0210331131113311-3230001303132233-0220031310332320-2213111201123111-3313223323013203-1132311133211101-0000100320011333-2133003302103223"></a>

## javascript_location property — rules / 130212132203 / 4

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](resources--http_loadbalancer--reference--group-014.md#canonical-1002301111101310-2233213103220330-3230330101321010-2132222110030233-2122312323031000-1131103001033230-0320321102312220-3111312232211001): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-014.md#canonical-3303333331111031-2310213001310120-1222102313203113-0003112123311101-2133101020201311-0222300312012013-0230312311102330-1021223130331333): complete subsection reference.

<a id="canonical-2202112002332221-1223220233003313-3333113330030132-1231221201033212-2011322031303330-1301302330223111-3303013230002331-0113221221101110"></a>

## Next pages — rules / 130212132203 / 5

- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-1200331032112302-0133002312021221-0013103132010212-3331330232303022-0333122302131133-2023201122233322-1101010330031001-0321133001110013)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain](resources--http_loadbalancer--reference--group-014.md#canonical-1301233102201233-0233012311103202-1023210002020011-2233110231333123-1300111222303013-1031033030030211-0220012203013020-3110201010213100)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata](resources--http_loadbalancer--reference--group-014.md#canonical-1002301111101310-2233213103220330-3230330101321010-2132222110030233-2122312323031000-1131103001033230-0320321102312220-3111312232211001)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path](resources--http_loadbalancer--reference--group-014.md#canonical-3303333331111031-2310213001310120-1222102313203113-0003112123311101-2133101020201311-0222300312012013-0230312311102330-1021223130331333)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1200331032112302-0133002312021221-0013103132010212-3331330232303022-0333122302131133-2023201122233322-1101010330031001-0321133001110013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333013133123000-3301312130100311-3030110021131201-2102233122113031-3221033131130201-3213333301320030-1203302122301002-0223012100023121"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain — any_domain / 322312011311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.any_domain

<a id="canonical-2101330300232023-2302033033302313-3030300212032320-0232121012022330-2010222000322100-3102110220300003-2332310313311303-2110100012013022"></a>

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

<a id="canonical-3302231321310123-1031213132213020-0112232322132001-3100022001211231-1001323031223320-2332022111111101-2112123102221301-3301223221101023"></a>

## Direct properties — any_domain / 322312011311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010103333320203-3031133132033313-3031233210233111-2000232100133221-1101333010311301-3332000010323221-2333102311132330-2311310102221223"></a>

## Next pages — any_domain / 322312011311 / 4

- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1301233102201233-0233012311103202-1023210002020011-2233110231333123-1300111222303013-1031033030030211-0220012203013020-3110201010213100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210113030312122-1203122332001031-3021311330003230-0323302313010020-3200012323100010-3312203311102132-3302101201100330-2300031323221133"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain — domain / 320311103001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.domain

<a id="canonical-3100222123222003-3131313010022211-2321112222201103-0333010121231000-3213313203021211-3133113113310202-1111223101331011-0231232310301013"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-1312131322200111-3310312210131133-0333332232121003-3211113002032222-0121121231231102-0001030020021232-1110110231210222-1111121132212131"></a>

## Direct properties — domain / 320311103001 / 3

<a id="canonical-2301020021133000-1311331003200221-1021123003121313-0333033000113011-2111332011011201-0120110021230131-3211203331213101-0231202213300132"></a>

<a id="canonical-3132020012213022-2131211202031232-1112111331323103-1112110331301233-1301212203131123-3201333101032203-3103101203332130-0012221020113310"></a>

## exact_value property — domain / 320311103001 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-3120323332312303-0023133121301200-1020312121101302-0231020221323111-0333330222001100-3220023330020011-3212322333101133-0213020023032100"></a>

<a id="canonical-0002030122112203-0323022023123130-0012121311132110-2223033111330213-1030123223212102-3020103332032220-2231110013021012-0220103101310320"></a>

## regex_value property — domain / 320311103001 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0112002321111110-2033300100331322-3101000332013221-1131002222112000-2031311130332331-0321010133331032-3001220220013320-3032111232021333"></a>

<a id="canonical-0113300021013201-0002001110110020-0020100201313320-0101223102100302-3030123333302112-0333121232030213-2322011010001010-0332322223331111"></a>

## suffix_value property — domain / 320311103001 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-2003303013103023-3013020111103110-1012222310031113-1303330221112200-3021333133111232-1320023123123101-0220123210000310-3332331311001132"></a>

## Next pages — domain / 320311103001 / 7

- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1002301111101310-2233213103220330-3230330101321010-2132222110030233-2122312323031000-1131103001033230-0320321102312220-3111312232211001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001032102200011-2203300003332231-1010013123201100-3333023303000120-3022100231113301-1231333332123013-2013030020133311-0322301202032203"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata — metadata / 333222301232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.metadata

<a id="canonical-3000030203002212-0102322000113032-2211112113101102-3100023001122112-1232121300320020-1032100013032332-0011222110103012-0222201313223332"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333302020123222-1011113200122110-3333111310321032-1202301322322230-0022332013222032-2113301002223120-2231333113320030-1131022000031321"></a>

## Direct properties — metadata / 333222301232 / 3

<a id="canonical-2132101122101101-3102021221230003-3303220312132020-0330301312102103-0230302023002022-2102213111203231-1120200323130130-3101030022123201"></a>

<a id="canonical-0232233300212310-3011121323101020-3230120033310110-1202123201201301-0201323323212113-0023312212230303-1122303331313232-2133222210002031"></a>

## description_spec property — metadata / 333222301232 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3030301101112321-0112122101310022-3032010211120313-0022300201021322-1230010023013033-3220130332003222-2122212320321230-0003011032300222"></a>

<a id="canonical-3000111223213132-2010101022233022-2221302032320212-2213220323122133-2210100132320111-3322011213121322-3123110003211322-2222121230203210"></a>

## name property — metadata / 333222301232 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-1101313302223111-1320231312012023-3001012102101113-3232113132323100-3220012113300210-1130001201333312-0312313321131022-3003121021201112"></a>

## Next pages — metadata / 333222301232 / 6

- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3303333331111031-2310213001310120-1222102313203113-0003112123311101-2133101020201311-0222300312012013-0230312311102330-1021223130331333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110000032133001-2001121000331023-1031233200023202-1201321102030221-3330331121310320-1033211212211213-2223331201002211-3011201131230220"></a>

## bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path — path / 002112031103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [bot_defense_advanced_protection.web_only.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0022212011301021-1310030223231023-0311131021123033-1123023031123220-1231313221303033-3231331102130323-2331203032202023-1111013132321301)
- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303)
- bot_defense_advanced_protection.web_only.js_insertion_rules.rules.path

<a id="canonical-0203203313211031-3310330102131011-0330120301012022-0120002330032330-1001100103301023-3300302021222331-2032323313311020-0020113010030201"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-0102333021112113-0310033113333210-3333310021102213-1023013200300122-2231330312313013-0020313002130321-3210002302313121-2313312020022021"></a>

## Direct properties — path / 002112031103 / 3

<a id="canonical-2010231332120320-3231002102112133-1012003033122200-0330102123222131-1120010010110312-2131121201220102-2213133123323301-2113222130000101"></a>

<a id="canonical-2301201222223113-2033201011231021-3023212303212331-0233102112033233-1001123313202322-2113011332231321-1113022130123120-0112113220132202"></a>

## path property — path / 002112031103 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-3302333312213312-3111222031131202-3332110232232030-3302022303222312-3011311133210302-0301223211303120-3112212201221001-1011020011110121"></a>

<a id="canonical-1331110333231100-2221303110331023-2021031120302132-0333131112230130-1212122013011013-0332002302210120-2010133233301011-3122303311312223"></a>

## prefix property — path / 002112031103 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-0321013030320331-2231312200002103-0112102231032021-2330032132322113-2020332320031311-1113311133012010-0323223122101203-2012220101323122"></a>

<a id="canonical-3311031330303323-3230310331022332-3000120121220320-0110130200210310-0111312131103232-1331331113301113-0302002322000331-2311302001322300"></a>

## regex property — path / 002112031103 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3302220023012322-3201221011133021-2112230132001303-2311122111120211-3013220323221121-2032203202033313-2113131111103031-2111123113232102"></a>

## Next pages — path / 002112031103 / 7

- [bot_defense_advanced_protection.web_only.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230132030321021-2011023120233132-0201331333120021-2012102123032302-1313020303332200-2030312102231133-3302101003303122-2122312220132303)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0032332312200101-3012322220030112-1320110202102312-0220201232102020-2202130130221013-0301003122332030-0331121213303131-0133333302230210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213121332130200-3301003331022010-3211120131001333-0211221013230000-3201302311010032-1010210220203001-1011130201223232-0311332132013002"></a>

## bot_defense_advanced_protection.web_only.web — web / 233222202122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-013.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- bot_defense_advanced_protection.web_only.web

<a id="canonical-3202100120323111-2013333220112231-3133232100133331-2320300122131320-2120011112112320-3000233012101323-1022301103101100-0221220221101013"></a>

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
web {
  # Configure direct properties listed below.
}
```

<a id="canonical-2201111202123322-0131133032333012-0031100331012023-3033100213011103-3310110030132220-1010300323323323-3201313022133303-2313122330303222"></a>

## Direct properties — web / 233222202122 / 3

<a id="canonical-2120213121111301-1003213333120030-3023233120030002-0001022120030022-2222033223320000-2131101220003030-1331220330012112-2321100310223113"></a>

<a id="canonical-3322322023201121-2120011233323222-0123022200312333-1132032121103101-0322230011310212-1202323301231320-3200202003010021-2202110313330101"></a>

## name property — web / 233222202122 / 4

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

<a id="canonical-0132330101223222-0211111122322110-2323103203103202-2110133311330101-2302132101030310-3212303210220231-1002133313101122-3323013102220222"></a>

<a id="canonical-2203323323301001-3132113330020331-2023213130303331-1313231201033212-3022102221032021-1200312333010210-3333331101032020-2210303303120233"></a>

## namespace property — web / 233222202122 / 5

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

<a id="canonical-2313221322230111-3031232211012123-3223121130213320-1001233320223133-1220120211001103-3222332102131211-3310202103321321-0112021031323020"></a>

<a id="canonical-2011230303323313-1010210303331312-2130212133101110-3002200313010000-2220330113323222-1202223113231001-1202331320311013-2333333213132233"></a>

## tenant property — web / 233222202122 / 6

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

<a id="canonical-3001022030210013-2021320313310133-2032002232100123-3301112330211132-3313032301131312-3211013112003022-2123001022111331-0001021011103013"></a>

## Next pages — web / 233222202122 / 7

- [bot_defense_advanced_protection.web_only](resources--http_loadbalancer--reference--group-014.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3102000111002330-2330012203002202-2013302232030202-2322011030203033-0013300220121222-2223102330220322-1011323030221103-3323302013003331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331130101300201-2212310330121222-1230313121111023-3202320001101212-3122300113130323-3103333220130132-0013212011102003-2112132133122232"></a>

## caching_policy — caching_policy / 031311212033 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- caching_policy

<a id="canonical-3203201312330020-1323212233312020-3213030330312302-3201133202232001-1012201201233023-0001210000312110-2333103330300332-3113013310222320"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: caching\_policy, disable\_caching; Default: disable\_caching\] Policy configuration for
this feature.

Upstream description:

Caching Policies for the CDN.

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

- [caching_policy](resources--http_loadbalancer--reference--group-014.md#canonical-3203201312330020-1323212233312020-3213030330312302-3201133202232001-1012201201233023-0001210000312110-2333103330300332-3113013310222320)
- [disable_caching](resources--http_loadbalancer--reference--group-018.md#canonical-2310010330303022-3212102030230223-3311321133320300-3300013001030002-1102021010321320-0321310020220321-0011203222103312-1122212300121322)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
caching_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0123012220312020-2113313010002222-3103300323011222-3231230102002013-3020013231131203-1130333231232123-1132001213303012-0322211203302010"></a>

## Direct properties — caching_policy / 031311212033 / 3

- [custom_cache_rule](resources--http_loadbalancer--reference--group-014.md#canonical-1112323230330001-3210133301302130-1001103132120102-1113111200333202-2010211133110132-3203313202133021-1030000033033020-1212102233212101): complete subsection reference.

- [default_cache_action](resources--http_loadbalancer--reference--group-014.md#canonical-2102303123010113-2230132023113110-0301132233211000-0123111331333002-0133122303200233-0021301123001210-2212032330201311-1313032131321121): complete subsection reference.

<a id="canonical-1123210313233130-1223323310010312-2003130022032100-3321133230113030-2202100330012120-1121313321220310-0220201132121232-2231220230313110"></a>

## Next pages — caching_policy / 031311212033 / 4

- [caching_policy.custom_cache_rule](resources--http_loadbalancer--reference--group-014.md#canonical-1112323230330001-3210133301302130-1001103132120102-1113111200333202-2010211133110132-3203313202133021-1030000033033020-1212102233212101)
- [caching_policy.default_cache_action](resources--http_loadbalancer--reference--group-014.md#canonical-2102303123010113-2230132023113110-0301132233211000-0123111331333002-0133122303200233-0021301123001210-2212032330201311-1313032131321121)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1112323230330001-3210133301302130-1001103132120102-1113111200333202-2010211133110132-3203313202133021-1030000033033020-1212102233212101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310310201300323-2201120202321021-3212123023303113-0101221111301332-0121232103001013-0332333221001210-3222323123021233-0230301101022302"></a>

## caching_policy.custom_cache_rule — custom_cache_rule / 210202002323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [caching_policy](resources--http_loadbalancer--reference--group-014.md#canonical-3102000111002330-2330012203002202-2013302232030202-2322011030203033-0013300220121222-2223102330220322-1011323030221103-3323302013003331)
- caching_policy.custom_cache_rule

<a id="canonical-1030232001203231-1120322313102322-1310033221321020-0311213221031110-3031322212233113-0200000033211102-0333322022102132-2020203303122311"></a>

Type: `"object"`. single nested block, Optional.

Custom Cache Rules. Caching policies for CDN.

Upstream description:

Caching policies for CDN.

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
custom_cache_rule {
  # Configure direct properties listed below.
}
```

<a id="canonical-3012120032130300-2133322000323133-1330221103013210-0203013202233002-0113202303112213-2111331301130222-3000201311101122-3302311013210022"></a>

## Direct properties — custom_cache_rule / 210202002323 / 3

- [cdn_cache_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0320220112122003-3200020033122332-1301313302000102-1122133000313102-1020022313003231-1322102210203301-3010230133213232-3003132230022032): complete subsection reference.

<a id="canonical-3011200133203123-0320300202330232-3231011230321332-2021203233122103-1311032112323311-2203003013221203-3130032312002013-1310123213302211"></a>

## Next pages — custom_cache_rule / 210202002323 / 4

- [caching_policy.custom_cache_rule.cdn_cache_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0320220112122003-3200020033122332-1301313302000102-1122133000313102-1020022313003231-1322102210203301-3010230133213232-3003132230022032)
- [caching_policy](resources--http_loadbalancer--reference--group-014.md#canonical-3102000111002330-2330012203002202-2013302232030202-2322011030203033-0013300220121222-2223102330220322-1011323030221103-3323302013003331)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0320220112122003-3200020033122332-1301313302000102-1122133000313102-1020022313003231-1322102210203301-3010230133213232-3003132230022032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112310103300000-2022001001033331-3211030222211120-1210012022023010-0122030321303100-2133110122213101-0102323011030021-2132111013300011"></a>

## caching_policy.custom_cache_rule.cdn_cache_rules — cdn_cache_rules / 102021130222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [caching_policy](resources--http_loadbalancer--reference--group-014.md#canonical-3102000111002330-2330012203002202-2013302232030202-2322011030203033-0013300220121222-2223102330220322-1011323030221103-3323302013003331)
- [caching_policy.custom_cache_rule](resources--http_loadbalancer--reference--group-014.md#canonical-1112323230330001-3210133301302130-1001103132120102-1113111200333202-2010211133110132-3203313202133021-1030000033033020-1212102233212101)
- caching_policy.custom_cache_rule.cdn_cache_rules

<a id="canonical-1322101121113022-0022112212033033-2230311211313132-0323131213121103-0321231323313210-2310033332000210-3222013302001212-3122113031132212"></a>

Type: `"object"`. list nested block, Optional.

Reference to CDN Cache Rule configuration object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
cdn_cache_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232203033002223-0130221312223302-0113323132110030-2030203321223020-1133110221200321-2321301122210113-2321010023031013-3221011112331013"></a>

## Direct properties — cdn_cache_rules / 102021130222 / 3

<a id="canonical-3221120010133022-3232223032202133-2331030122231203-0330331202132213-1322020223203130-3112310002010122-3110022232000222-2331131223302123"></a>

<a id="canonical-3120301020020103-1322111302202210-0033321233323321-2301332221223101-3022311033312323-0113001221132111-0003300021230213-1003131231230131"></a>

## name property — cdn_cache_rules / 102021130222 / 4

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

<a id="canonical-3132201313211200-3032331220011120-1111310031112021-1002001101321202-1030122332302323-1212331000023103-0111311222012303-2222311020223323"></a>

<a id="canonical-1123210002120010-2300200030101233-3020123303212032-2013320330232200-3112200323232011-0212303113310222-2221310111300233-1100123302011333"></a>

## namespace property — cdn_cache_rules / 102021130222 / 5

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

<a id="canonical-0002223230122003-0012203300221322-2113332020033120-0031113232223031-0123110012231033-0121221120313310-0333203012322332-3101222332330011"></a>

<a id="canonical-3322030301012123-0232120120230201-0220322312031003-1210302023003103-3001213031222231-1223303030011021-2122220002210331-1121110000300013"></a>

## tenant property — cdn_cache_rules / 102021130222 / 6

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

<a id="canonical-3302110031023111-3102323300203221-1321313100020321-0321130202311022-3222103202132132-2113120200321330-0100232121231020-2202120201322022"></a>

## Next pages — cdn_cache_rules / 102021130222 / 7

- [caching_policy.custom_cache_rule](resources--http_loadbalancer--reference--group-014.md#canonical-1112323230330001-3210133301302130-1001103132120102-1113111200333202-2010211133110132-3203313202133021-1030000033033020-1212102233212101)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2102303123010113-2230132023113110-0301132233211000-0123111331333002-0133122303200233-0021301123001210-2212032330201311-1313032131321121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123222302110103-1112333021101000-3020321013120002-2103320322230022-3301013202220211-3130020131022323-2202212023310101-3200231212322130"></a>

## caching_policy.default_cache_action — default_cache_action / 132210302201 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [caching_policy](resources--http_loadbalancer--reference--group-014.md#canonical-3102000111002330-2330012203002202-2013302232030202-2322011030203033-0013300220121222-2223102330220322-1011323030221103-3323302013003331)
- caching_policy.default_cache_action

<a id="canonical-0231101131123021-0111200113000033-1303222201222132-0310210233310203-2212213300120232-3002021013033102-1302310322310221-0213303022202231"></a>

Type: `"object"`. single nested block, Optional.

Default Cache Behaviour. This defines a Default Cache Action.

Upstream description:

This defines a Default Cache Action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cache_disabled",
    "cache_ttl_default"),
  validators.ConflictingObjectAttributes("cache_disabled",
    "cache_ttl_override"),
  validators.ConflictingObjectAttributes("cache_ttl_default",
    "cache_ttl_override")}
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
  "x-ves-oneof-field-cache_actions": "[\"cache_disabled\",\"cache_ttl_default\",\"cache_ttl_override\"]"
}
```

Terraform syntax:

```terraform
default_cache_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201012323331313-0212220020030130-0032031323023103-2303202230232130-2322123102232320-0210203202331122-2210210111002130-3112033303312301"></a>

## Direct properties — default_cache_action / 132210302201 / 3

- [cache_disabled](resources--http_loadbalancer--reference--group-014.md#canonical-3210112322012122-2202223201233230-0231212203120000-1333302200310022-3030002313231033-1230310201012301-2221012202100321-3012003200320101): complete subsection reference.

<a id="canonical-2203031003231100-3022333131332021-2221023111303010-1111210311031131-2002200031103213-2100310230301010-0031011003133032-1020222022003100"></a>

<a id="canonical-1213133333201102-0010020300101200-0121100222013200-1331202320221031-1213311332033021-3311332003032322-3021001230111233-2222201320112333"></a>

## cache_ttl_default property — default_cache_action / 132210302201 / 4

Type: `"string"`. Optional.

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

Upstream description:

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

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
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-1103331303022230-2030311020010200-2312003102022223-3331033001011201-1132002101322200-1331310323132103-2000001013111112-0330012000223031"></a>

<a id="canonical-2231123101003122-3123231102302003-3301112230222202-3322120331002203-0320033132123332-2033130101122210-2233203023232323-1310003333201230"></a>

## cache_ttl_override property — default_cache_action / 132210302201 / 5

Type: `"string"`. Optional.

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

Upstream description:

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

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
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-3011122031200233-3211120332032110-0302231333012111-1030310132002203-0313311100320133-0003210010003133-1300300221032302-1003320320123230"></a>

## Next pages — default_cache_action / 132210302201 / 6

- [caching_policy.default_cache_action.cache_disabled](resources--http_loadbalancer--reference--group-014.md#canonical-3210112322012122-2202223201233230-0231212203120000-1333302200310022-3030002313231033-1230310201012301-2221012202100321-3012003200320101)
- [caching_policy](resources--http_loadbalancer--reference--group-014.md#canonical-3102000111002330-2330012203002202-2013302232030202-2322011030203033-0013300220121222-2223102330220322-1011323030221103-3323302013003331)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3210112322012122-2202223201233230-0231212203120000-1333302200310022-3030002313231033-1230310201012301-2221012202100321-3012003200320101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112003332030031-3030100312312131-0103020132100023-1230231011323012-2232301101030010-2131000020031131-0333013311203022-0021301110001130"></a>

## caching_policy.default_cache_action.cache_disabled — cache_disabled / 022320210112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [caching_policy](resources--http_loadbalancer--reference--group-014.md#canonical-3102000111002330-2330012203002202-2013302232030202-2322011030203033-0013300220121222-2223102330220322-1011323030221103-3323302013003331)
- [caching_policy.default_cache_action](resources--http_loadbalancer--reference--group-014.md#canonical-2102303123010113-2230132023113110-0301132233211000-0123111331333002-0133122303200233-0021301123001210-2212032330201311-1313032131321121)
- caching_policy.default_cache_action.cache_disabled

<a id="canonical-2321332220331300-0021013010322012-2122023211002010-0313210333111321-1232132323022201-0002130102301131-1212223230022133-3320332113331132"></a>

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
cache_disabled = {}
```

<a id="canonical-2010233032313012-0110210130110200-0322303211020211-2201031022323133-3122211022033031-2231023100301330-1102000010012220-2012130310112313"></a>

## Direct properties — cache_disabled / 022320210112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323222313123233-3320030111023203-3233130003221203-0131023020113000-0031031030302322-1231102311023110-1131001112212301-3232302210211013"></a>

## Next pages — cache_disabled / 022320210112 / 4

- [caching_policy.default_cache_action](resources--http_loadbalancer--reference--group-014.md#canonical-2102303123010113-2230132023113110-0301132233211000-0123111331333002-0133122303200233-0021301123001210-2212032330201311-1313032131321121)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3331001122010103-1330222312203012-2102022330310300-3121202020223232-3202122212132233-0220222003113103-2301213310311323-1232121323232302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133303122032321-1102031001103232-0130323202211110-0131132123321032-2100221023312112-0033303023201301-0131020332231231-1300311201002102"></a>

## captcha_challenge — captcha_challenge / 103022322221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- captcha_challenge

<a id="canonical-1232032332023102-1010001000122202-0112003231022230-3233031333012222-3331333220233000-3300331300233031-1213033101103311-0120000120323113"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: captcha\_challenge, enable\_challenge, js\_challenge, no\_challenge,
policy\_based\_challenge; Default: no\_challenge\] Enables loadbalancer to perform captcha challenge
Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that
pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is
configured to do Captcha Challenge, it will redirect..

Upstream description:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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

- [captcha_challenge](resources--http_loadbalancer--reference--group-014.md#canonical-1232032332023102-1010001000122202-0112003231022230-3233031333012222-3331333220233000-3300331300233031-1213033101103311-0120000120323113)
- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-0331332220112013-2112002310330311-0000103101120120-1022133131132112-2221133100010123-1020333222003102-1303303322000302-2321231220030231)
- [js_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-1011111331300221-3213330020020313-1033022332302021-1111331213030331-0032003020022230-3000023321331133-2002032233231133-0233011032133100)
- [no_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-3132021103233010-0102003303001330-3110103203321112-0211303310001001-0000320020121320-3033100021101033-2220023223112313-3310221213303023)
- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-0321223103102130-0121132313031223-1220210002011320-2133231330312121-2131132110220002-1313302332130023-1110033030130013-0300230120302232)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
captcha_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-1311102022130030-0332013011230103-1010201013020033-3313023121011013-2320202323123200-2312300211301121-1201210202222000-1102302033132122"></a>

## Direct properties — captcha_challenge / 103022322221 / 3

<a id="canonical-0310033331102021-0301312112130232-0113033213310010-1330111223022001-3300201112031022-2101233333130013-3212321221102111-1232102220300020"></a>

<a id="canonical-2031312221122121-0032210302113123-1121021123132202-0120221220032201-1021322023020213-3200223303321211-0102130001213202-2013303302202300"></a>

## cookie_expiry property — captcha_challenge / 103022322221 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-3223102002230332-3000031121132103-2320100232201302-0121230003321322-0322212333031120-1013133331211202-0030000220132103-3023210211231202"></a>

<a id="canonical-0003032303032103-0330231033303110-2232210123013123-1001112301200023-1211132321022201-0222003021221310-3010012222331222-2201301233302322"></a>

## custom_page property — captcha_challenge / 103022322221 / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

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

<a id="canonical-0032030221020302-2322221000130311-3132213003323232-2321002032112031-3222131213130032-2020210300223102-0032011322031313-0230020101300013"></a>

## Next pages — captcha_challenge / 103022322221 / 6

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032032301023113-3100231103301321-1230030131223003-3203230231233011-2010013330021023-1103010031300331-0113202302310320-3211231113232121"></a>

## client_side_defense — client_side_defense / 113032031213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- client_side_defense

<a id="canonical-0033332020201100-0231102120012232-1311202303221021-2103323133031333-2103310211223013-1103120330033200-3332001302000030-0233201201222132"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: client\_side\_defense, disable\_client\_side\_defense; Default:
disable\_client\_side\_defense\] Defines various configuration OPTIONS for Client-Side Defense
Policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense Policy.

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

- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0033332020201100-0231102120012232-1311202303221021-2103323133031333-2103310211223013-1103120330033200-3332001302000030-0233201201222132)
- [disable_client_side_defense](resources--http_loadbalancer--reference--group-018.md#canonical-1313000133203032-2310331201223313-3311201202200022-2303113031201211-1010211132031221-2313112333213303-2001211021211300-1211030023323302)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
client_side_defense {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033011321203201-0333100010111312-0213203213023123-0300101201322323-3120203322211020-2201031010221311-3232021112321112-3213031110110133"></a>

## Direct properties — client_side_defense / 113032031213 / 3

- [policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310): complete subsection reference.

<a id="canonical-2003210033222110-2033002333312020-3233101033030031-2022020313333313-2312220302033002-2212100132112203-1010103230032000-0032312301112020"></a>

## Next pages — client_side_defense / 113032031213 / 4

- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202032233130320-0223132112300033-1300101303310010-2032110011101312-2320332112120003-2031203333331230-3113131323032323-3023123121131203"></a>

## client_side_defense.policy — policy / 222200102123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- client_side_defense.policy

<a id="canonical-1213221003211122-1132011111321033-2013113123130032-3111322002231211-2103303031223230-0201123200123032-1231100122322021-1031031101000230"></a>

Type: `"object"`. single nested block, Optional.

Defines various configuration OPTIONS for Client-Side Defense policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages_except",
    "js_insertion_rules")}
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]"
}
```

Terraform syntax:

```terraform
policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230313221110233-2202113030001312-1103323211120202-2113022110033300-0102330303112313-0012030111221222-1020311033300333-3020223210113102"></a>

## Direct properties — policy / 222200102123 / 3

- [disable_js_insert](resources--http_loadbalancer--reference--group-014.md#canonical-0113131101113001-3321120322331133-2303222210330300-0300223210012311-3322133012200132-1110000111313120-0313300310231031-0121202103222232): complete subsection reference.

- [js_insert_all_pages](resources--http_loadbalancer--reference--group-014.md#canonical-3321121013323313-0201322201321222-2200102111023133-2312132222201301-2223320202220022-2332133133322110-3112313330120313-2011323100003200): complete subsection reference.

- [js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231): complete subsection reference.

- [js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023): complete subsection reference.

<a id="canonical-1020310011101301-3113000002231002-0212210110011130-0122123202310223-1223233203012301-2203313122032221-0111133123001111-0123311133202201"></a>

## Next pages — policy / 222200102123 / 4

- [client_side_defense.policy.disable_js_insert](resources--http_loadbalancer--reference--group-014.md#canonical-0113131101113001-3321120322331133-2303222210330300-0300223210012311-3322133012200132-1110000111313120-0313300310231031-0121202103222232)
- [client_side_defense.policy.js_insert_all_pages](resources--http_loadbalancer--reference--group-014.md#canonical-3321121013323313-0201322201321222-2200102111023133-2312132222201301-2223320202220022-2332133133322110-3112313330120313-2011323100003200)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0113131101113001-3321120322331133-2303222210330300-0300223210012311-3322133012200132-1110000111313120-0313300310231031-0121202103222232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133200232323313-2132020010120331-2120010202323010-2113322130223303-2220111311110100-3003001121303301-0330233230031010-2131100103122103"></a>

## client_side_defense.policy.disable_js_insert — disable_js_insert / 101020203200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- client_side_defense.policy.disable_js_insert

<a id="canonical-1031311331321213-0203202131310323-0022230000002002-2110313003310132-1123031023301020-3202320021223300-1200010203322230-3133223321133001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

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
disable_js_insert = {}
```

<a id="canonical-3310232231100120-0110201233023322-0000332131133331-0202033201303203-1232312213022132-2233032123331311-2013111003210010-0311212223332001"></a>

## Direct properties — disable_js_insert / 101020203200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311200303223303-1032233233202211-0113130201030121-3101312012321131-0231232120301111-2201032301301112-3301301012121203-3231132022321230"></a>

## Next pages — disable_js_insert / 101020203200 / 4

- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3321121013323313-0201322201321222-2200102111023133-2312132222201301-2223320202220022-2332133133322110-3112313330120313-2011323100003200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331302222113112-1233331030302111-2033011301321312-0102222031023020-1112113021102200-1022332020211022-1010331000120112-3133030333021133"></a>

## client_side_defense.policy.js_insert_all_pages — js_insert_all_pages / 301100120301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- client_side_defense.policy.js_insert_all_pages

<a id="canonical-2032311313000003-2300322321221032-1121303333310212-1022020201223230-1330022331302200-3131001310202301-3301210131010012-3021022311111020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for js insert all pages.

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
js_insert_all_pages = {}
```

<a id="canonical-0111233210132121-0332020131311313-2231333202112233-1122220011020031-1322011020232330-0312210210112131-0132200003010033-3222013312303002"></a>

## Direct properties — js_insert_all_pages / 301100120301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3021003220123311-2300131113233333-1332313330232321-3103313111002223-3102333210201311-2123331232312222-0202300232330302-3113303203112122"></a>

## Next pages — js_insert_all_pages / 301100120301 / 4

- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121222102322232-0131210203230223-2001213203333322-3010312302311321-1113302121200201-3231323223230231-0213212022101131-2323321002331122"></a>

## client_side_defense.policy.js_insert_all_pages_except — js_insert_all_pages_except / 221220332010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- client_side_defense.policy.js_insert_all_pages_except

<a id="canonical-3133230212120011-3032310011031001-3310330000123011-0132000110133330-0231122131212303-1011210322202321-0303132232100232-3202233212302013"></a>

Type: `"object"`. single nested block, Optional.

Insert Client-Side Defense JavaScript in all pages with the exceptions.

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
js_insert_all_pages_except {
  # Configure direct properties listed below.
}
```

<a id="canonical-1101012003333113-2210233213333123-0311323211222101-1012200212321123-0120111331003001-1012102320221032-2020300233201021-3020303102220023"></a>

## Direct properties — js_insert_all_pages_except / 221220332010 / 3

- [exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313): complete subsection reference.

<a id="canonical-2300211013210003-2003020300210132-0013032300112012-1302311330213223-1130033301210020-2312303012203132-3033223010132101-3331030230321300"></a>

## Next pages — js_insert_all_pages_except / 221220332010 / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131100021112022-1201131320231132-3332312203031331-1132312032002331-3331331103311203-2131301221212213-3310333301103100-0031121220113011"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list — exclude_list / 032032301030 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-3131000322120213-2323121100223113-3311003130303011-0321021323110200-3102323022110122-3311132002033203-2210013203313331-0232300310121003"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2011103332033230-0011111133222120-3033230012002020-2131021311223212-0303110212133322-2031003220213231-2101131310302103-1131302013310001"></a>

## Direct properties — exclude_list / 032032301030 / 3

- [any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-3201331123201221-0313230333212000-0132203201121120-2032221301322010-0010132300133000-3301113012010321-3133133012001123-3101022113302312): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-014.md#canonical-2002333202310111-1022200100012001-1223122121131230-2223310031202011-2320100231113213-0302310310331300-3311113032031312-0230213130322133): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-014.md#canonical-2213132032212100-1030112202020323-1032113021332311-3203323121302113-2000021221311113-1231013120322233-0023111100301111-0302322221113211): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-014.md#canonical-0313113200033031-0213201113202133-0212000000101012-1220101330223002-0233201331122321-2012200102121121-0232212210302121-2023222022330020): complete subsection reference.

<a id="canonical-3301232201123311-1032002221310011-3220111301312323-1113201103003000-1132122130121311-0123221131203023-0201000131223000-1213333300103030"></a>

## Next pages — exclude_list / 032032301030 / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-3201331123201221-0313230333212000-0132203201121120-2032221301322010-0010132300133000-3301113012010321-3133133012001123-3101022113302312)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain](resources--http_loadbalancer--reference--group-014.md#canonical-2002333202310111-1022200100012001-1223122121131230-2223310031202011-2320100231113213-0302310310331300-3311113032031312-0230213130322133)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata](resources--http_loadbalancer--reference--group-014.md#canonical-2213132032212100-1030112202020323-1032113021332311-3203323121302113-2000021221311113-1231013120322233-0023111100301111-0302322221113211)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.path](resources--http_loadbalancer--reference--group-014.md#canonical-0313113200033031-0213201113202133-0212000000101012-1220101330223002-0233201331122321-2012200102121121-0232212210302121-2023222022330020)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3201331123201221-0313230333212000-0132203201121120-2032221301322010-0010132300133000-3301113012010321-3133133012001123-3101022113302312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122121332121121-1230312222221310-2131233221330313-2130131111310122-3010200030111100-2303312120223021-3112102001333003-1311121230121010"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain — any_domain / 023121131133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-0121132311002212-3331021223313100-1211230321231113-1122123021020211-2133213130203232-1113132200032031-0101310111131001-1010323311230130"></a>

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

<a id="canonical-0030023130233032-0330200023030122-2202302311021210-3112333313232123-0233331013000331-1113210221111112-0022311310321202-2203202120333203"></a>

## Direct properties — any_domain / 023121131133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033002323033032-1022012212011023-3100120012320112-2200031200321003-0213011203021232-0001231210211310-3301022320020321-3220233201002110"></a>

## Next pages — any_domain / 023121131133 / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2002333202310111-1022200100012001-1223122121131230-2223310031202011-2320100231113213-0302310310331300-3311113032031312-0230213130322133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003310102001121-1231210222031000-3313112100010122-1322003300020023-0033133320231031-3133220312020321-3320010301123320-3133023120113113"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain — domain / 312133233220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-3133220201201202-2020302121212302-0320112130302213-1132132332021112-1031213003110003-0312120111323001-0023031003103323-2223130113110000"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201332333203113-2123021103300030-1310033121010303-3331012231013221-0223310300002132-3200202002321121-1211220100210011-2331232302110321"></a>

## Direct properties — domain / 312133233220 / 3

<a id="canonical-3331123031010323-3201030310221110-2012230332013133-0100013231131123-2330200122113223-1231002220011131-2112133332200201-3222020321302120"></a>

<a id="canonical-3221013111332320-0000232222332301-0220002032303101-1202013212322203-0310021330030030-3302110131300310-1000003213130312-3030130100202123"></a>

## exact_value property — domain / 312133233220 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1120123030013013-2103001112303113-1221220320231112-2113111111032021-2210301322032202-2010232231011300-2131010333111220-3312111222303320"></a>

<a id="canonical-0113032311021211-2212113333320230-2221233310203310-2022133110332211-1123211210020300-0231212103300320-2331300203131133-1222023003202122"></a>

## regex_value property — domain / 312133233220 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3100203031221001-3112130201333213-1021213233002302-3211011131203000-2032231022233231-1101032103321031-1202311211101302-1302111010000122"></a>

<a id="canonical-2122312011032221-3333231130021023-0123131012230020-1330330112231011-3313003300102000-2132032100223030-2132302323031200-0320232211312112"></a>

## suffix_value property — domain / 312133233220 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-2013102312213100-2323013120300121-2313130200331112-0032031303311333-1112332111301300-2311302103310022-1221132221210010-3222333310020221"></a>

## Next pages — domain / 312133233220 / 7

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2213132032212100-1030112202020323-1032113021332311-3203323121302113-2000021221311113-1231013120322233-0023111100301111-0302322221113211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013202322001220-3001222231112132-3211031121303200-1300033303322320-1213200011012021-3112101320231121-0302232233013111-0323210333002112"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata — metadata / 210232203103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-2021001212312321-1010322200001202-0121320212022300-0002321010133133-3301130130120200-0301023031113213-2201101002121233-1321003313310122"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202323330000322-1033101010210302-3210323011122131-1212133001021001-2332112012233312-3212021112013121-1132322010100211-2020220310311121"></a>

## Direct properties — metadata / 210232203103 / 3

<a id="canonical-1112013112202133-1020230302001013-3113111303130133-2301101030010031-3311123001103320-3200121303322103-1231202302202232-0102333010021220"></a>

<a id="canonical-2301102111223013-0312130210322213-1130232220000200-3013231323202111-3223010110203112-2110113103031013-1313103331202313-0333301132002131"></a>

## description_spec property — metadata / 210232203103 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1111032031200311-1322302222301132-0310203000121310-0302313032311031-2011030320112110-2302311110023022-2001100323303232-1202132002110133"></a>

<a id="canonical-3212002101333012-1201122301313120-1121313300010201-0213001333103122-3003313000231022-1030031130332023-1101213310010011-0111103330001023"></a>

## name property — metadata / 210232203103 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3030133203103013-1022023301330030-2313220012230010-1003013322313122-2331333211300012-0002132003320332-2112021102013332-0121011332133313"></a>

## Next pages — metadata / 210232203103 / 6

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0313113200033031-0213201113202133-0212000000101012-1220101330223002-0233201331122321-2012200102121121-0232212210302121-2023222022330020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013013022300321-1010002231031230-3332111021202231-3320210302101230-1330130130132300-3303201122002001-0311301130213223-3303331101011113"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.path — path / 330233231133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-014.md#canonical-3203113232311103-1203301023312113-1103112313002121-2300301213131000-2021012213013110-0200231011312133-0033320012121330-3322223131210231)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-2221000100332222-2303223321131020-0111013311213213-3310323233333123-3000031113100001-2233302121020221-0021123012001131-3212012021021311"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221310312133230-2203133213302202-0020223031132123-0231133100132111-1201120300302223-2202230012333232-1101130013022312-2111221223021233"></a>

## Direct properties — path / 330233231133 / 3

<a id="canonical-1111212032200320-3011301031023323-3221330302210210-0000311030221221-3101013120030130-0131130212220210-2003302020012202-0011111131132110"></a>

<a id="canonical-3021031230112032-2013110101300330-1331023221232211-2002131021310023-1320011032311203-1313111213012323-2013232103211230-3032001113321031"></a>

## path property — path / 330233231133 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-0333303301330320-0312103102302023-0331033022033101-0333203001001100-3131333013111302-2203201101121311-1221102022233120-3101023111201021"></a>

<a id="canonical-0220301032132332-0201233311120311-2110121321320331-3312000233022030-0320213011331310-3022301231122112-1312000202121102-3223022312213212"></a>

## prefix property — path / 330233231133 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-2100031031030130-2012301112332302-0223131222120022-0302012310202000-2333023001003312-2300202331010122-3103322232202213-0221210311001322"></a>

<a id="canonical-0033111132112203-0113210223113211-2032101221221122-0310132330111332-1303303200122201-0131331230322222-0231133001122110-2301311010030102"></a>

## regex property — path / 330233231133 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3012102200233120-1330033020002130-0033320131033003-2233023123102222-2212311211320331-2021120020211320-0223213001131110-0010332013103030"></a>

## Next pages — path / 330233231133 / 7

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0033230223132100-0222132221220121-0013133202323110-3330203023221110-1223312330000301-2201203133133303-1200331111012323-0331103310011313)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213031210023023-3000011212202323-2211002013132210-3303212311102221-2211101103303023-0100112202221020-3300301203030310-1232122001033000"></a>

## client_side_defense.policy.js_insertion_rules — js_insertion_rules / 121200230312 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- client_side_defense.policy.js_insertion_rules

<a id="canonical-2223330120010033-3330220112301033-0301003123001311-3320220321313233-1132332333110202-3003123131132332-1130231203233122-0213223013101130"></a>

Type: `"object"`. single nested block, Optional.

Defines custom JavaScript insertion rules for Client-Side Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Client-Side Defense Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0112210110313230-0003333022032331-3223011000220022-0012233203121100-1202220310111001-1112032132331110-2333231133021332-2031303220203110"></a>

## Direct properties — js_insertion_rules / 121200230312 / 3

- [exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222): complete subsection reference.

- [rules](resources--http_loadbalancer--reference--group-014.md#canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321): complete subsection reference.

<a id="canonical-0210133101220310-0311313123131122-3331131210112031-3323131032032100-2131000012003120-2320313101233330-2202002132312330-1201230110113320"></a>

## Next pages — js_insertion_rules / 121200230312 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- [client_side_defense.policy.js_insertion_rules.rules](resources--http_loadbalancer--reference--group-014.md#canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203022302023012-2230320022001313-1023130211300120-1133002013231222-3222103000111031-0013311233103333-1221030113213231-0113332321000311"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list — exclude_list / 120302212322 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- client_side_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-0302200211013330-3120012120312110-1023212310321322-1101203233113201-1201102222121132-1220100321322210-3122330033303312-2232322300130320"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130021311200300-3231103120310212-2200012322321221-3111300132230123-1202011220023233-3330212331210311-2121020222233032-0300202131221101"></a>

## Direct properties — exclude_list / 120302212322 / 3

- [any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-1211032113213031-1312033220303030-2000131032203030-0210101023300331-1302000301020223-3021320231203002-3032322113012022-3023223132032202): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-014.md#canonical-1302133302131301-0302203010332010-3032202302312331-1322312023023320-0121210331002113-3331322031031103-3212132120101113-3210001000000231): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-014.md#canonical-0323321301322131-3031120031101112-2203233330200212-1222233131100003-0310002222021030-1211312130102001-2320312203103213-1012210013102312): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-014.md#canonical-3211010123012220-1110123231023113-1023333033223033-3122203312012010-3321230002200121-0033210201023331-2132301132120200-3021313130322321): complete subsection reference.

<a id="canonical-0120332313221233-2000321102031133-1123013102303331-2110331122003310-2202111221200310-2133023020011330-3223001112202233-1021332212102313"></a>

## Next pages — exclude_list / 120302212322 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list.any_domain](resources--http_loadbalancer--reference--group-014.md#canonical-1211032113213031-1312033220303030-2000131032203030-0210101023300331-1302000301020223-3021320231203002-3032322113012022-3023223132032202)
- [client_side_defense.policy.js_insertion_rules.exclude_list.domain](resources--http_loadbalancer--reference--group-014.md#canonical-1302133302131301-0302203010332010-3032202302312331-1322312023023320-0121210331002113-3331322031031103-3212132120101113-3210001000000231)
- [client_side_defense.policy.js_insertion_rules.exclude_list.metadata](resources--http_loadbalancer--reference--group-014.md#canonical-0323321301322131-3031120031101112-2203233330200212-1222233131100003-0310002222021030-1211312130102001-2320312203103213-1012210013102312)
- [client_side_defense.policy.js_insertion_rules.exclude_list.path](resources--http_loadbalancer--reference--group-014.md#canonical-3211010123012220-1110123231023113-1023333033223033-3122203312012010-3321230002200121-0033210201023331-2132301132120200-3021313130322321)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1211032113213031-1312033220303030-2000131032203030-0210101023300331-1302000301020223-3021320231203002-3032322113012022-3023223132032202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322230320313223-3000311322313031-0230013321321002-0223203231212100-3201313301322320-2030032113322100-3020310231020231-2032333021301231"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.any_domain — any_domain / 100031221320 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- client_side_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-2303001012113301-0023200320010131-3220013113232303-1133220220310320-1100231331323113-2322033121231203-0333110011112320-1003012332200223"></a>

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

<a id="canonical-2233021010313100-0011121022213300-2112112212323313-1332133212321001-2213302001031302-1113123101301122-2101100112113311-2310303222330303"></a>

## Direct properties — any_domain / 100031221320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013220100330220-2111002113003330-2311220001211012-1000121213011300-0230302230003303-3210000001322132-3002102231120303-1001120233232312"></a>

## Next pages — any_domain / 100031221320 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1302133302131301-0302203010332010-3032202302312331-1322312023023320-0121210331002113-3331322031031103-3212132120101113-3210001000000231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331210213101210-0233313220303113-1312100230102302-1333331111112322-1323331011232333-1222230220312330-1101113220213220-2210233131333303"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.domain — domain / 101101332110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- client_side_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-1221121033332311-2012000301122123-2113031000012302-3012301133313221-1222101310211300-1002231102321103-0332023012300313-1300132120312113"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-1221013130132331-3210012323033230-0130021000201002-0300310311300230-2330321230023201-3013322211111202-2213310002113103-0103301032223331"></a>

## Direct properties — domain / 101101332110 / 3

<a id="canonical-1122100220003232-1022313203011322-1120031301310230-0200223222310232-0313110232203130-3330022323330233-0022210020231223-3220001021202211"></a>

<a id="canonical-2311131310311013-2122030102002211-1012313100101003-3312122013103102-2201322132131320-0022130200231100-1112112332322302-3221032100201013"></a>

## exact_value property — domain / 101101332110 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-0031021330022122-3320020301030213-1210022001202101-1201123202331132-2300123000232203-2310232003013201-0313101321213111-1022210011321013"></a>

<a id="canonical-0110330320011303-2300102002122123-3121012301002112-2031303133313102-3302311122211031-3021301012002001-1310002123212100-1202310221303211"></a>

## regex_value property — domain / 101101332110 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3212121011203331-1120101130313031-3333310303102103-1212023331231311-0330021011331013-2123220030212130-3332100020230211-2020120312022212"></a>

<a id="canonical-0010130212102022-2331032021300331-1003010133211003-3302020131013000-1330323123320210-2220323101130231-0333133003223011-3312211230022131"></a>

## suffix_value property — domain / 101101332110 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1100211031220333-0100301022003030-3312033010013031-0222031313222331-1100230123021002-2321200030033020-3020101003323312-2313310133113032"></a>

## Next pages — domain / 101101332110 / 7

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0323321301322131-3031120031101112-2203233330200212-1222233131100003-0310002222021030-1211312130102001-2320312203103213-1012210013102312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031013303003322-0230323310131112-2001112220310212-2320033213003102-0301232103132312-0301121033231230-3232030010200023-3233312112222321"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.metadata — metadata / 233011023020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- client_side_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-0231121201131001-1310103223323133-2210210103100121-1030332233103330-2000032120122311-0100322132313332-0010323223301031-2200323033223010"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313233020301113-2001110200222000-2101113031131232-2322002321210321-0010010220311130-1330302220211332-2113110311123303-3303100210112323"></a>

## Direct properties — metadata / 233011023020 / 3

<a id="canonical-1331203223102003-2212021133333302-0230023021300001-1002231133223310-1133101023113320-3121010031120313-1303202302033313-0100200130002003"></a>

<a id="canonical-2201220230211323-0211320330123321-1300031203313102-2003230022230330-0322110201232102-0113302001133033-3313320312033213-1030230221220203"></a>

## description_spec property — metadata / 233011023020 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3123123003033212-0003203322121032-0201303233321321-0312022301223301-1113011003002333-1003302330123101-1313333201231231-3311000312231121"></a>

<a id="canonical-3312031332203310-1231012031011230-1123310132113030-0322213303311332-2113023020101233-0210113130022212-0301110111032220-3301132203102020"></a>

## name property — metadata / 233011023020 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-1132012120011333-3322001211010012-0133231111320003-2201223010310112-3121200112133002-3200101032300231-1132132101022312-0300032230121211"></a>

## Next pages — metadata / 233011023020 / 6

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3211010123012220-1110123231023113-1023333033223033-3122203312012010-3321230002200121-0033210201023331-2132301132120200-3021313130322321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321032121132301-0233333031223330-0021233020133222-0030233330201011-3133121213011033-2311211013031233-3231330101202033-2021022022332300"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.path — path / 001333201032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- client_side_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-0213001220012101-0132033101322101-3331002303112111-2033122100033133-1122023211110301-0211031010131320-0213233130013023-0321231231020222"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201203013131010-0133233210121020-0011010032312300-3230201331230000-0330332203200121-2103020231132210-1222010212211122-0201231130223022"></a>

## Direct properties — path / 001333201032 / 3

<a id="canonical-3001213122213233-0302112032230113-0010120333220232-0320330321112320-3112023321132120-1112031030202213-2333300332023000-0013331211122322"></a>

<a id="canonical-2332213022301011-2101022322121012-1222200200113332-1121333003221301-1030210101313110-3212230310101002-0133030222331210-3032112202123212"></a>

## path property — path / 001333201032 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-0220323202003033-3100302023111320-1003020021320121-0033121013322003-0001222002002010-1120300023210333-1011321033310100-0111300031320233"></a>

<a id="canonical-0001302021132111-2321213013211211-2103311133333232-1111311311013003-3032310211033302-2001302123033000-3120121320303231-0023011021123030"></a>

## prefix property — path / 001333201032 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-0211332330211200-2300310012010133-3311012200330220-1213210131320213-1331001012213033-1030113001111020-2331010103230113-2320330323223130"></a>

<a id="canonical-2121001312010323-3301221313333002-2323031123221303-1022203121122202-1202203303323322-1000032121102323-1230010013020122-0323321300001213"></a>

## regex property — path / 001333201032 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2311132132012011-2221012221030301-1221021123310100-3011023010122330-2012100233322202-0031320332100010-1132101330312311-2130112033220003"></a>

## Next pages — path / 001333201032 / 7

- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-014.md#canonical-0212120331010031-2122122110301032-2100131012033113-1321021123133203-2031312013023113-2002012210110200-2001033210320220-2101013013333222)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2003130231131001-2330111303132012-2002331330312302-1320031033110122-3001110220033102-2320021302020322-3323132133020121-0210202032110321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210110201131120-0133121011121321-3012333100023013-0323233200323300-2200003201010223-2230133020323002-3300312002033012-0212201222033300"></a>

## client_side_defense.policy.js_insertion_rules.rules — rules / 031200002200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [client_side_defense](resources--http_loadbalancer--reference--group-014.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222)
- [client_side_defense.policy](resources--http_loadbalancer--reference--group-014.md#canonical-1220301112210323-0023120000113110-1102122320303220-2220012113111020-3122113030000011-2330223210331310-2003312220033111-1123210003321310)
- [client_side_defense.policy.js_insertion_rules](resources--http_loadbalancer--reference--group-014.md#canonical-3230331100022123-1000231223230133-3323322113203301-1302010023311223-2310202002210100-2220300303313123-2323333202310030-1120123301020023)
- client_side_defense.policy.js_insertion_rules.rules

<a id="canonical-0230001303330022-1302211122321222-1112003231233221-1122000331220021-3200220102120021-1333320211222200-3002313112332223-3321013210213102"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Client-Side Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```
