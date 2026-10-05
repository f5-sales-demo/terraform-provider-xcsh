---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-3221122122011322-1001003111120311-1022210133021010-2121102231121300-0000032012032030-0300031220200011-2323222232210021-3111232300221103"></a>

## transformers property — item / 012131130203 / 6

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
EnumExtractionComplete: false
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

<a id="canonical-2321310012310130-0301200313030302-0222221032222111-3020010113130332-3313303002311123-3202113223131201-0122300322121000-2022321010021301"></a>

## Next pages — item / 012131130203 / 7

- [api_rate_limit.api_endpoint_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-003.md#canonical-2123010330003131-1200001120312112-0103222130123312-2032231130103322-0030222120220210-0022123103310113-0022203231302333-0133112223211113)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013102210101312-2000332030101023-2003010132312213-0323112120121202-3022013210310302-3322130012131232-3323001211031103-0121300310332131"></a>

## api_rate_limit.bypass_rate_limiting_rules — bypass_rate_limiting_rules / 020301220022 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- api_rate_limit.bypass_rate_limiting_rules

<a id="canonical-1321320021223310-2220311302230202-2012320121211121-1230313313200022-3313001313032332-2320222120211101-0010331023000133-2313001023030123"></a>

Type: `"object"`. single nested block, Optional.

Category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

Upstream description:

This category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

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
bypass_rate_limiting_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310020020310022-3120203000203203-1101023320103320-1200201313331213-2121320323330100-2122011330330022-0310133130131213-0033203330013012"></a>

## Direct properties — bypass_rate_limiting_rules / 020301220022 / 3

- [bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130): complete subsection reference.

<a id="canonical-3211233222132001-0221000323122313-0232012010203012-1222311020220330-3102200022330311-3300102003013210-0223003231131111-1233003021001112"></a>

## Next pages — bypass_rate_limiting_rules / 020301220022 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101312103123313-1233133111101211-1302232331222212-3203102301230300-0311302111121300-1000203000122210-0232212223001320-1120300113232210"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules — bypass_rate_limiting_rules / 032321132231 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules

<a id="canonical-1001331312332113-0011300020312213-1112302310000121-2201312012012310-0130233231321012-0203021210221110-3100223022310011-0233322121321112"></a>

Type: `"object"`. list nested block, Optional.

Category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

Upstream description:

This category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("any_url",
    "api_endpoint"),
  validators.ConflictingListObjectAttributes("any_url",
    "api_groups"),
  validators.ConflictingListObjectAttributes("any_url",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "api_groups"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_groups",
    "base_path")}
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
    "ves.io.schema.rules.repeated.max_items": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "20"
  }
}
```

Terraform syntax:

```terraform
bypass_rate_limiting_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223103222103200-1331132310022331-0301211312112031-3001323303112211-0011022310331031-0230100331103103-1010023102012031-3131230233100110"></a>

## Direct properties — bypass_rate_limiting_rules / 032321132231 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-004.md#canonical-1201320111203120-3332223101112022-2012321103220100-2221223332233033-2010031201113313-2320230303033310-1331101222111303-1201330033323221): complete subsection reference.

- [any_url](resources--cdn_loadbalancer--reference--group-004.md#canonical-2201201000113012-2321221131010120-1120111132033020-1033012110031120-2231131022100311-0331001321033210-3233203000313113-3132222020213012): complete subsection reference.

- [api_endpoint](resources--cdn_loadbalancer--reference--group-004.md#canonical-1220122231011231-2200231221123101-1301123301132322-0312031201012033-1030131101320211-2321000032101302-0201121312022022-2103332113230300): complete subsection reference.

- [api_groups](resources--cdn_loadbalancer--reference--group-004.md#canonical-2113212020210033-3010221101301030-0003101203132101-2111230231331133-2002330113220123-0313120011022233-3031223010200303-2131300133100023): complete subsection reference.

<a id="canonical-2203232033301301-1310232213320032-1311123030012310-0113102211201022-3302332020313130-0120211120021021-0200131311001111-2321020302123300"></a>

<a id="canonical-1201332100111022-1210100231012312-2200223003331022-2010121031210102-3123000200010130-1110320330013212-3122013120223332-2030020332202212"></a>

## base_path property — bypass_rate_limiting_rules / 032321132231 / 4

Type: `"string"`. Optional.

Exclusive with \[any\_url api\_endpoint api\_groups\] The base path which this validation applies
to.

Upstream description:

Exclusive with \[any\_url api\_endpoint api\_groups\] The base path which this validation applies
to.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112): complete subsection reference.

- [request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003): complete subsection reference.

<a id="canonical-3102200022003203-3032333321233132-2203133100300230-2322010133020013-1003233030332133-0332310313031010-0021133312313223-2000300212220332"></a>

<a id="canonical-2010012000023212-3120221113113323-0102123301021223-0101331313032001-0110301323112201-1220111111230010-0223011132010012-1322211301303123"></a>

## specific_domain property — bypass_rate_limiting_rules / 032321132231 / 5

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For example:
api.example.com.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0132311311332021-3133001020301210-2311121023103222-2311200021323232-2212330130102101-3232333100123303-1332313302203103-3122313001033211"></a>

## Next pages — bypass_rate_limiting_rules / 032321132231 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain](resources--cdn_loadbalancer--reference--group-004.md#canonical-1201320111203120-3332223101112022-2012321103220100-2221223332233033-2010031201113313-2320230303033310-1331101222111303-1201330033323221)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url](resources--cdn_loadbalancer--reference--group-004.md#canonical-2201201000113012-2321221131010120-1120111132033020-1033012110031120-2231131022100311-0331001321033210-3233203000313113-3132222020213012)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint](resources--cdn_loadbalancer--reference--group-004.md#canonical-1220122231011231-2200231221123101-1301123301132322-0312031201012033-1030131101320211-2321000032101302-0201121312022022-2103332113230300)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups](resources--cdn_loadbalancer--reference--group-004.md#canonical-2113212020210033-3010221101301030-0003101203132101-2111230231331133-2002330113220123-0313120011022233-3031223010200303-2131300133100023)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1201320111203120-3332223101112022-2012321103220100-2221223332233033-2010031201113313-2320230303033310-1331101222111303-1201330033323221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200323121011121-1033100132101023-3303201101122000-1121112010010331-1233331112020333-3321232022221303-2311300331121110-1020032111113110"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain — any_domain / 112200300001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain

<a id="canonical-0123133031231102-0030112023210223-1203013320332020-3233033220103102-2101200010031021-3002232022233003-1103130300202102-3230030310322010"></a>

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

<a id="canonical-3123221011012233-3031112002330100-2211313212302231-3032031133032220-0312112212200200-3210220102310223-0303100010311230-0231230330033011"></a>

## Direct properties — any_domain / 112200300001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201012001300031-2202222200303333-2133321111330130-3332100230310321-1222332022322023-1132211111301131-2300132032031011-2010333121020122"></a>

## Next pages — any_domain / 112200300001 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2201201000113012-2321221131010120-1120111132033020-1033012110031120-2231131022100311-0331001321033210-3233203000313113-3132222020213012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111033232012313-1332001333313111-3233010200222101-3002310201002100-2220211232110312-2302002120311032-3313313132000113-2233000223222032"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url — any_url / 101213212222 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url

<a id="canonical-2310232231100210-1112212220020102-3330122102320131-1021121330203010-3200130123303223-1122003030232320-0200303001231101-3023003022233101"></a>

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
any_url = {}
```

<a id="canonical-0121110133131303-2122201120301330-0000322022110031-1231301332002113-3023130301222121-3333132222221321-3111022310010103-1211120332203333"></a>

## Direct properties — any_url / 101213212222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301323330323322-3100233020212003-0303222213310332-2331011322113331-2210112220021032-2332322011103021-0321231033222022-0012212300311331"></a>

## Next pages — any_url / 101213212222 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1220122231011231-2200231221123101-1301123301132322-0312031201012033-1030131101320211-2321000032101302-0201121312022022-2103332113230300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130000110220212-0103232020022233-2030203320001321-0101322230310330-3030232102303231-1113203232132021-3033212203010220-1311232302302313"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint — api_endpoint / 011312103021 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint

<a id="canonical-2310212331102203-3103332012002220-1013212220033031-2002222320103202-3213233011021210-2313201200311111-3201120112312112-2321030331200202"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
api_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-2233120020303023-2310310321001331-1301011203002203-2000330123023203-1132220112112130-0101311331100233-2221222032023132-0102100332112031"></a>

## Direct properties — api_endpoint / 011312103021 / 3

<a id="canonical-1223121031132322-2023223001121111-1332322322311300-0223032110321031-1121321312123112-0113302020130210-1300221233311232-3320333232200023"></a>

<a id="canonical-2123312003120303-1100103101331202-2320321121130102-1300301322323330-2120102013203001-1030332131103023-1003331322001202-0223203203311112"></a>

## methods property — api_endpoint / 011312103021 / 4

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1000300312111002-0311132220221020-3233022233333231-3213331030113211-0031200232220001-0231100203200323-3100003012101012-1130212301210331"></a>

<a id="canonical-1133113221331020-0131121020203302-3013231110130122-0100103220022120-0112313223103112-3002121331212301-0303001113033312-2033101111230321"></a>

## path property — api_endpoint / 011312103021 / 5

Type: `"string"`. Optional.

Path. Path to be matched.

Upstream description:

Path to be matched.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

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
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
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

<a id="canonical-3221300230002111-1313101003222300-1000323033020320-3231200203330301-1113300023101033-0221110221112220-0320033222312232-0012031122101120"></a>

## Next pages — api_endpoint / 011312103021 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2113212020210033-3010221101301030-0003101203132101-2111230231331133-2002330113220123-0313120011022233-3031223010200303-2131300133100023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233111003232233-0312302003202121-1211032212322100-0201211221033320-1223121220331212-1110121012322223-2312012222102200-0031021011102001"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups — api_groups / 110120322313 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups

<a id="canonical-3112223002103303-2201313332223213-2133212323023131-1313322201211113-2010020203010132-1103302120110001-3221200323021111-0330311221323322"></a>

Type: `"object"`. single nested block, Optional.

API Groups.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("api_groups")}
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
api_groups {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232231000311312-3120232120313132-3012200013213211-2022023211032211-2331112033103212-1212223123302032-3323321310111321-1003211132001031"></a>

## Direct properties — api_groups / 110120322313 / 3

<a id="canonical-0023112100132021-1201100201031132-1202123000033323-0130331131112331-3333212213132203-1000020231231010-2303123123330111-3001300123021123"></a>

<a id="canonical-3132001110311322-1031000031102302-2330020023230213-3121032000210031-3232130210030111-3203203331212112-2112133031320112-2030330232213303"></a>

## api_groups property — api_groups / 110120322313 / 4

Type: `["list", "string"]`. Optional.

API Groups. Group or collection configuration

Upstream description:

Group or collection configuration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2333331303122110-3202032333120213-1313323130000203-1100303130100331-3020020220323311-1333033231000110-3003223323102323-0200322231301202"></a>

## Next pages — api_groups / 110120322313 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203122303112331-3002021033330121-2003332320132133-2333002120000310-2321320133231113-3322223112102311-2021110312201302-2003113100333321"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher — client_matcher / 132313020003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher

<a id="canonical-0020100232132112-3133012333033112-0220301131221011-3230103330301021-2111031131002221-1002321030221332-1020202131131332-1323220330130223"></a>

Type: `"object"`. single nested block, Optional.

Client Matcher. Client conditions for matching a rule.

Upstream description:

Client conditions for matching a rule.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3212222232122132-3013332201323001-0120110031002032-0211323330112222-0032332011020222-0121232210113200-2313332032300230-2020223203231211"></a>

## Direct properties — client_matcher / 132313020003 / 3

- [any_client](resources--cdn_loadbalancer--reference--group-004.md#canonical-2302331200031202-1210002300221010-2132213322310030-1321220130123132-1201020202232103-1210120201001211-0320002010102133-1031101133103133): complete subsection reference.

- [any_ip](resources--cdn_loadbalancer--reference--group-004.md#canonical-1130320030100033-1211303132000120-0223231303301010-0033211233210222-2311131230013122-0131312331020230-3313303001233130-0312103021202330): complete subsection reference.

- [asn_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-2110333231132032-3211332313121302-2013022022213000-0212200023123101-0103323232220020-3222100300300332-2301122303130302-1303012203121123): complete subsection reference.

- [asn_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3032300131130131-0102330220303332-1020322132022203-0113232203130012-1222313231100203-2032203212320100-1320232230113231-1320333230103021): complete subsection reference.

- [client_selector](resources--cdn_loadbalancer--reference--group-004.md#canonical-2120310313303211-3202011103102020-2010110230302233-0202320101300021-3023311312011032-1221331303122013-0200021332101210-2330330001202330): complete subsection reference.

- [ip_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-1313322031133031-2031123002032102-1031111131332103-0221110302023113-3102021020033133-0333320102312232-2001120103320130-2201222323012320): complete subsection reference.

- [ip_prefix_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-3201021030112331-0103312100311221-2330200122322220-3200133211322310-0303103102303303-0011220101003200-1123122120033131-2031113233010001): complete subsection reference.

- [ip_threat_category_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-2003112110333113-2231021110032122-0123132011102103-1223210331201310-2133130221021232-0212310113131300-1322010011002020-1201310103133011): complete subsection reference.

- [tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2112131213110233-0233201110033223-2013001111300332-3231011123323020-0000330233311011-2131331110201130-0102200023031130-0133223332110331): complete subsection reference.

<a id="canonical-1000203322133132-2233101023101213-2022123321300101-0223302122112302-2031022000020300-1013300330103230-3020111301202123-2012010332320033"></a>

## Next pages — client_matcher / 132313020003 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client](resources--cdn_loadbalancer--reference--group-004.md#canonical-2302331200031202-1210002300221010-2132213322310030-1321220130123132-1201020202232103-1210120201001211-0320002010102133-1031101133103133)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip](resources--cdn_loadbalancer--reference--group-004.md#canonical-1130320030100033-1211303132000120-0223231303301010-0033211233210222-2311131230013122-0131312331020230-3313303001233130-0312103021202330)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-2110333231132032-3211332313121302-2013022022213000-0212200023123101-0103323232220020-3222100300300332-2301122303130302-1303012203121123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3032300131130131-0102330220303332-1020322132022203-0113232203130012-1222313231100203-2032203212320100-1320232230113231-1320333230103021)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector](resources--cdn_loadbalancer--reference--group-004.md#canonical-2120310313303211-3202011103102020-2010110230302233-0202320101300021-3023311312011032-1221331303122013-0200021332101210-2330330001202330)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-1313322031133031-2031123002032102-1031111131332103-0221110302023113-3102021020033133-0333320102312232-2001120103320130-2201222323012320)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-3201021030112331-0103312100311221-2330200122322220-3200133211322310-0303103102303303-0011220101003200-1123122120033131-2031113233010001)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-2003112110333113-2231021110032122-0123132011102103-1223210331201310-2133130221021232-0212310113131300-1322010011002020-1201310103133011)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2112131213110233-0233201110033223-2013001111300332-3231011123323020-0000330233311011-2131331110201130-0102200023031130-0133223332110331)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2302331200031202-1210002300221010-2132213322310030-1321220130123132-1201020202232103-1210120201001211-0320002010102133-1031101133103133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133232331201311-2333101032211001-3012112001103231-0333003232001222-1330013332211120-3212121203312121-0022303223332213-1213011200222000"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client — any_client / 001220022003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client

<a id="canonical-2110223122032030-1121303203030102-1001331220003330-0111331232033030-0002200110020232-1003100102100311-1211302312332002-0313012021330100"></a>

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

<a id="canonical-3320212213310121-1101311301221201-0212301200033320-1002122330202331-0002200022012232-3030030200221110-0223303121002330-2030001311220320"></a>

## Direct properties — any_client / 001220022003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322301002100132-2101122302223331-1311220321100013-3023302213012220-3132101322103021-0232323033121213-1233111133003122-2002221220301221"></a>

## Next pages — any_client / 001220022003 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1130320030100033-1211303132000120-0223231303301010-0033211233210222-2311131230013122-0131312331020230-3313303001233130-0312103021202330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331012302110331-3022133221213223-0122002210022012-1311112210312001-0030021301111300-1220320121013202-3032200113103032-0311330320011033"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip — any_ip / 323221210203 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip

<a id="canonical-0301303133112223-3133300212220223-1313110100002123-1200230220220133-1210100333232111-3302230223010210-0101033111003013-3123130223133310"></a>

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

<a id="canonical-2020133012322311-3222203231003202-3033120121132010-3131013000130213-0110203100003321-0221130133111301-3200021312110301-0203003013332213"></a>

## Direct properties — any_ip / 323221210203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311203303220122-0321122120131333-2303131111222121-0002201231121101-1300022130120310-0023202320030223-2011100221002132-3133133312312332"></a>

## Next pages — any_ip / 323221210203 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2110333231132032-3211332313121302-2013022022213000-0212200023123101-0103323232220020-3222100300300332-2301122303130302-1303012203121123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232303333113023-2321331002103212-3300121323101333-0212312333222003-2201033023013021-1300302003210111-2000102002212302-1112112123201202"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list — asn_list / 112011000213 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list

<a id="canonical-1121120333311012-2121122310320322-0203033132233211-1001201313022130-0001322301331203-0220001231321000-1321000113110130-2031113222320313"></a>

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
EnumExtractionComplete: false
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

<a id="canonical-2312023320303131-3221132121302003-0330031213333020-1231022210233332-2100222312201333-0210303212012320-2303320220023131-0210131232032202"></a>

## Direct properties — asn_list / 112011000213 / 3

<a id="canonical-3303330022202003-0321202022320220-3113021133203321-2200333101102203-0320202000213330-0203123211230123-1011020210222311-1301323331132213"></a>

<a id="canonical-0122212302123223-3131103122022202-2120211200311212-2222133202032303-3131133201331011-3033103323110330-2330233220011312-2310220213301122"></a>

## as_numbers property — asn_list / 112011000213 / 4

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
EnumExtractionComplete: false
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

<a id="canonical-0202332132310221-1312301323210011-1312011231000021-2220300313200321-0212311323032003-1103201222130231-2010230331310200-0223033011313022"></a>

## Next pages — asn_list / 112011000213 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3032300131130131-0102330220303332-1020322132022203-0113232203130012-1222313231100203-2032203212320100-1320232230113231-1320333230103021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300310312302331-0302221001210121-0203002233200231-2232013301112331-0203113321202213-2021230213320030-2322020122001323-3222012202022211"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher — asn_matcher / 030231331012 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher

<a id="canonical-0100031103232123-1211020112133133-1031103311231203-0310133003020213-0013003303101003-0232302112232003-0201111011220200-2320231031021131"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0022101332212220-0012130223200110-0230202230213010-1322310022110220-0130121232033211-0102213002120002-2122001332111123-1100323200331010"></a>

## Direct properties — asn_matcher / 030231331012 / 3

- [asn_sets](resources--cdn_loadbalancer--reference--group-004.md#canonical-0103022321200201-2110322302102221-3310133203120012-1122320300232311-1232131313320010-2210323021230110-3211130100332101-3330030212323320): complete subsection reference.

<a id="canonical-3023100102230010-2332221323312121-1011320132232021-2011230030120321-3111313311223121-2203032000232131-0003001203013112-0322213220201301"></a>

## Next pages — asn_matcher / 030231331012 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets](resources--cdn_loadbalancer--reference--group-004.md#canonical-0103022321200201-2110322302102221-3310133203120012-1122320300232311-1232131313320010-2210323021230110-3211130100332101-3330030212323320)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0103022321200201-2110322302102221-3310133203120012-1122320300232311-1232131313320010-2210323021230110-3211130100332101-3330030212323320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121130130022332-0333302131333200-0231222111123123-3122100010300132-0132020303030231-2200002200230322-2122101331123111-3010113232023211"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets — asn_sets / 223030331123 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3032300131130131-0102330220303332-1020322132022203-0113232203130012-1222313231100203-2032203212320100-1320232230113231-1320333230103021)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-2312110231301001-0022003010110333-0103222000320230-0301200330301003-1031200011202211-1113032230212123-3202211023033223-0012003211112010"></a>

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

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211021302000221-1301320111010112-0320130201221311-1133322023003023-1013120113200033-0010201330103101-3321012111100300-2232113230223303"></a>

## Direct properties — asn_sets / 223030331123 / 3

<a id="canonical-2313110313231101-3000023003133212-1031221220202023-3022102133300303-0310113222033223-1302331100210300-0330030203222022-0320012321122132"></a>

<a id="canonical-1122103032123313-0322303121202113-3113210210211001-0130132112012333-2320100310222022-2021230010102233-2301112132031132-0002011200312002"></a>

## kind property — asn_sets / 223030331123 / 4

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

<a id="canonical-3200120313003321-3320032320322210-2003030232000233-0101010212012101-1013120221232330-1020130013210000-1012132032130300-1302230003222313"></a>

<a id="canonical-0121102302301323-0103212213222123-3311121201113210-2221330122131222-3011012222132301-1113213103211021-2022001020102203-2111300211302001"></a>

## name property — asn_sets / 223030331123 / 5

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

<a id="canonical-1110033101030122-0032032011331020-1020200120133321-1200012213120100-1322101221113000-0133102023121313-0030113100220132-3102321102120103"></a>

<a id="canonical-2300312221022130-0201303132103000-2312300303122130-2303331000233002-0323130012203330-1111231312212100-3332003330333300-2031303013022101"></a>

## namespace property — asn_sets / 223030331123 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0002203330302320-3002121123032222-2033203123020030-2003101302320020-1011212033001003-2322222212120022-0120032102100101-1211030313333213"></a>

<a id="canonical-3312032333100123-2010020010131130-2013030132031233-2130100021110003-1113301313113133-2203120132122020-2202202230101330-0133120121330221"></a>

## tenant property — asn_sets / 223030331123 / 7

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

<a id="canonical-1020102223023003-1031110113201311-2032210310221022-1013310111331131-0200232302130212-1223132010311323-0121023100222101-2200230330012003"></a>

<a id="canonical-3310310002010001-2022030120221322-2333001121113002-1212001123311221-1002313022123320-1011132013002202-3301311133131133-0201030110310200"></a>

## uid property — asn_sets / 223030331123 / 8

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

<a id="canonical-3202111232132212-1103331300321022-2322312033330102-3323003300222210-1321023032110112-3312313101003210-2201313030023212-2011012121200232"></a>

## Next pages — asn_sets / 223030331123 / 9

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3032300131130131-0102330220303332-1020322132022203-0113232203130012-1222313231100203-2032203212320100-1320232230113231-1320333230103021)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2120310313303211-3202011103102020-2010110230302233-0202320101300021-3023311312011032-1221331303122013-0200021332101210-2330330001202330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322011102022023-2311100121332232-0302100002012200-0113133111303000-3111312020130301-1030333200211123-2121232213110203-2213200010023203"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector — client_selector / 133211221313 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector

<a id="canonical-1321311201121333-3303013200000223-1010332232002221-2012003213313331-3333121202023220-1310100200132202-1122012012331301-2003333032313212"></a>

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
EnumExtractionComplete: false
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

<a id="canonical-1031010310100313-0320223232202302-1020230330333121-0102031030230001-0022320011221211-3331331212230320-1122310223321302-3332132303210022"></a>

## Direct properties — client_selector / 133211221313 / 3

<a id="canonical-1203320230110210-0011010033303210-2230330101310200-1332011231030222-1022111202110122-0322030210210020-1322221102301121-0212210223002110"></a>

<a id="canonical-3210132112233201-0103312211101311-2233221302103200-0102122232131302-3021331123121022-0310331320233233-0220330001120030-2000203131200333"></a>

## expressions property — client_selector / 133211221313 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3133300200122223-2201021330211211-1330320030000223-1210013323030032-2231122320032022-3300100301032210-1000132201201201-0223031312320102"></a>

## Next pages — client_selector / 133211221313 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1313322031133031-2031123002032102-1031111131332103-0221110302023113-3102021020033133-0333320102312232-2001120103320130-2201222323012320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001311312332132-1232000211201111-1312233110223000-2110002102332320-1313222001101331-0121211320322030-2003013132203231-1322220033223121"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher — ip_matcher / 023203023332 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher

<a id="canonical-3013033103112012-1230023202321211-3201111020100102-3021330101312200-2003331130213021-0023102321100320-1202322333130220-1331212122232312"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3010120311211323-1301300111101111-2003211012200202-0301121230311121-3031131112322013-3130301230232323-0022331133232020-3012032230213020"></a>

## Direct properties — ip_matcher / 023203023332 / 3

<a id="canonical-3111303003302000-3212321232211000-1232123303012201-0003222121323322-3312032131200022-3131130130222313-1221112322231203-0213011311310313"></a>

<a id="canonical-1110021310232310-0133213323223201-0032111121121231-3203303013103030-0320323123110023-0010313312302010-2310102201033111-2330300232121030"></a>

## invert_matcher property — ip_matcher / 023203023332 / 4

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

- [prefix_sets](resources--cdn_loadbalancer--reference--group-004.md#canonical-1233030103120332-3100201120113232-3300021103020120-1331012131312311-0223233130012210-1303312112102230-1020012002223002-1102313102113313): complete subsection reference.

<a id="canonical-2203110213200123-3301102020313111-0323322231110222-0001001020121210-2130211203333330-2230311110020010-3101120102320312-3130020230022211"></a>

## Next pages — ip_matcher / 023203023332 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets](resources--cdn_loadbalancer--reference--group-004.md#canonical-1233030103120332-3100201120113232-3300021103020120-1331012131312311-0223233130012210-1303312112102230-1020012002223002-1102313102113313)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1233030103120332-3100201120113232-3300021103020120-1331012131312311-0223233130012210-1303312112102230-1020012002223002-1102313102113313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232311230223002-3020111031223121-3010111322001301-3222013120112212-1210303201023223-0000232132031310-0013310303231131-3330002032332003"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets — prefix_sets / 123333231323 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-1313322031133031-2031123002032102-1031111131332103-0221110302023113-3102021020033133-0333320102312232-2001120103320130-2201222323012320)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-2013102030323000-2020011230213322-1202102121213220-1333011130012113-3020033121231320-0113100230110330-3110301321232303-3033001233033211"></a>

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

Terraform syntax:

```terraform
prefix_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102122212022102-3101001031000333-1320231211233320-0031101223132321-1133223310030311-0331221212010221-2211113312020102-2231030113012100"></a>

## Direct properties — prefix_sets / 123333231323 / 3

<a id="canonical-3020300132112321-0131003210301320-1232223133130200-3123032013212120-2223331330332210-0231333220222231-2023123123313023-3100220303210301"></a>

<a id="canonical-1113303112222330-1033310200123200-3333231030132332-0130313132230223-2110111111031011-3001130212203212-0111031233231211-0210122100300220"></a>

## kind property — prefix_sets / 123333231323 / 4

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

<a id="canonical-2213300300103001-2131113223333231-2120100213003220-1313113113103330-3323133331212333-2122122310200213-3020011200202001-0022013222123333"></a>

<a id="canonical-1322301023131320-1313231300001122-3220310302303302-0303220213211300-3120230231001222-1332020222121111-0223013310300220-0022011311022023"></a>

## name property — prefix_sets / 123333231323 / 5

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

<a id="canonical-3121003303131200-3312221033301002-3233012021203100-0130332021032332-1212110102320123-3000010110022033-1322023021303332-3213012200333210"></a>

<a id="canonical-2000030303312130-3231002121133203-0030313220021310-3300022321033222-0022332111323323-2322323103211331-0330130221230021-0202132220030110"></a>

## namespace property — prefix_sets / 123333231323 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2210112102021320-2101032011100222-2311332330023012-2212230000100323-3321330201210010-2323203201002122-2001010111133313-2212202322132100"></a>

<a id="canonical-2301312132230010-3312010101302010-2100120201130122-0223213301121213-1313213113030332-0323131023002303-1010032003111301-0121231322203123"></a>

## tenant property — prefix_sets / 123333231323 / 7

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

<a id="canonical-0322120002012012-1010123321112302-3323200313202322-0312222333121011-3211113202230122-1331311023123120-1012030113321013-0233102302211202"></a>

<a id="canonical-1113213010023010-0221120300022211-2002110312212002-3023331100331233-3323333130032023-3123123210200321-0022232311200112-0213230200322000"></a>

## uid property — prefix_sets / 123333231323 / 8

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

<a id="canonical-0230000333211013-2210201120122103-0301130322001122-1210110020022003-2123000303112321-2133100332103032-2110023132333223-0123313230031300"></a>

## Next pages — prefix_sets / 123333231323 / 9

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-1313322031133031-2031123002032102-1031111131332103-0221110302023113-3102021020033133-0333320102312232-2001120103320130-2201222323012320)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3201021030112331-0103312100311221-2330200122322220-3200133211322310-0303103102303303-0011220101003200-1123122120033131-2031113233010001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201103301200231-3312233201030213-3231233200023303-2133113000123331-1321023133230132-2302203121331010-0130312220223302-2001331231103101"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list — ip_prefix_list / 233010332230 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list

<a id="canonical-3212320033030133-1113311210220320-2000211213100011-0222132313212220-3231130333101213-3100313320021123-3000130220331001-3301332102201102"></a>

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

<a id="canonical-3003012301220323-3322121112322131-0230200012320012-2232021221113023-2321211330023331-2330001121001213-0230332230102133-2202201130200033"></a>

## Direct properties — ip_prefix_list / 233010332230 / 3

<a id="canonical-0111231033312022-1110231030121023-2103200033100323-0333031203212011-0111201220121031-1312223003031110-1000222013301012-3211331223211310"></a>

<a id="canonical-3101320302231300-0001013220121320-2003330100333323-2130211022203033-3011331021301303-0131110230232221-2322223310023201-2031000222302120"></a>

## invert_match property — ip_prefix_list / 233010332230 / 4

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

<a id="canonical-0212001233102113-1002233101310312-2010130010301321-0130201130100130-3313132103202331-1230030033332112-3202020212122310-2213003020201123"></a>

<a id="canonical-0033303300030131-1202102321002223-1020030202011332-2300133223032101-0330323223022011-2331030331202310-3131321030122231-3010121103330310"></a>

## ip_prefixes property — ip_prefix_list / 233010332230 / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2000111021320330-3110132231203220-0012313110333330-3320003201120100-0203200213321013-2112230201233323-1202202321002032-1230203102320322"></a>

## Next pages — ip_prefix_list / 233010332230 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2003112110333113-2231021110032122-0123132011102103-1223210331201310-2133130221021232-0212310113131300-1322010011002020-1201310103133011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233312323111131-0000203131001111-2113100311303103-0131013230002120-3121213021102132-0112202121311010-3020100133212023-3301013033001000"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list — ip_threat_category_list / 312330202000 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list

<a id="canonical-1112310322100010-2123100003000230-1120020013231223-2322200132232301-1013010232120302-1121201212201000-3131233311322010-2322020012001201"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List Type. List of IP threat categories.

Upstream description:

List of IP threat categories.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0121123122030323-3223113010202021-2333131100333221-2223323032331032-2100322133311321-1321021102133110-3102301131211001-1120103312132300"></a>

## Direct properties — ip_threat_category_list / 312330202000 / 3

<a id="canonical-2321013202310213-0123231203013323-3023230320220031-1020330013200310-3112311211033031-3320303010101131-2033301131021213-2133012113311101"></a>

<a id="canonical-3202133120102003-3112230212112221-3332132232313030-2023210133132223-1020232023233313-3130021002100022-1023013213010211-2101101301310202"></a>

## ip_threat_categories property — ip_threat_category_list / 312330202000 / 4

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
EnumExtractionComplete: false
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

<a id="canonical-2002132010320220-0333010012210003-3330032010101021-2212030020331331-3020233223221213-1210231113202200-1130322232322123-1100312102213013"></a>

## Next pages — ip_threat_category_list / 312330202000 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2112131213110233-0233201110033223-2013001111300332-3231011123323020-0000330233311011-2131331110201130-0102200023031130-0133223332110331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300313222212303-2031001223310200-2211211302201021-2330311122221000-1121132101113212-1311020111220312-0133322210120220-3303113233220322"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher — tls_fingerprint_matcher / 103001211202 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-0133002331210023-2111010111223100-3132212210021022-0223033332010122-3003112121012021-3132010331011032-3322011322331120-0013111002132302"></a>

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

<a id="canonical-0020302313231302-2020303030210201-3210221111212103-0013221102100022-2231132201003310-1220300011123012-2332132030333223-0200030321320321"></a>

## Direct properties — tls_fingerprint_matcher / 103001211202 / 3

<a id="canonical-1010321200003212-3020231132210322-2312121013222310-1133021013020101-2120230311331100-3231233203200322-1003303113103312-2022130221032210"></a>

<a id="canonical-0121331110212120-0033120213211321-2012100000113010-3302101303200002-1330020203202201-1322030303113112-0321310133203021-2321331323133131"></a>

## classes property — tls_fingerprint_matcher / 103001211202 / 4

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
EnumExtractionComplete: false
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

<a id="canonical-2110300001012331-2303213312011303-1202232031101231-2333312011200300-2010033002202200-3103023212220030-2212223233202111-1103221021233012"></a>

<a id="canonical-3232320002332213-3212032100313302-1210332230322313-1213210130221221-0302223012113122-1102023311300300-1311230102002011-2220301333320023"></a>

## exact_values property — tls_fingerprint_matcher / 103001211202 / 5

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1130213100332221-0221011111132003-2233100202031121-1022222233100303-3113003131302203-1220233203223032-0230213323300232-0302312110232000"></a>

<a id="canonical-0230002010322121-3323022211030311-3003213022001232-2203312312313120-3033001132113210-2011212330012311-0320113000130211-0100322000101230"></a>

## excluded_values property — tls_fingerprint_matcher / 103001211202 / 6

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
EnumExtractionComplete: false
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

<a id="canonical-1213321003111103-3023202230122212-0120113301300001-0021122331210233-3221121021011303-2121102211030030-0223320233300123-1231100001201113"></a>

## Next pages — tls_fingerprint_matcher / 103001211202 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-2122312220333321-3132332113223003-1220323003033021-0020312102332120-1211122033300122-1113130013013021-3231001301113021-1322301233223112)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213202320033000-1312131020123223-3233131013133312-1131031020200100-3231012023113323-1311100102302233-1001312313220210-0322302122212121"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher — request_matcher / 033300130021 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher

<a id="canonical-2212112322011001-0100010131333222-3313322120003201-3223222102330023-3001221200230021-3210221302220022-0311300310212333-2211010330000102"></a>

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

<a id="canonical-0320331221101321-0210302100112030-2013003001003130-0331212233113132-3221321213211311-3022033210312322-2200322123310021-2201222223022132"></a>

## Direct properties — request_matcher / 033300130021 / 3

- [cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3230212322131111-0220132331111211-2130300110020101-0303031223301012-2233230022030113-3322103001131330-0200101110321222-3313131200332323): complete subsection reference.

- [headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3102122111323102-3333011122330001-3203211020002021-3311333132211201-1301113303101330-0020211323000331-1132020302200002-0132023233120303): complete subsection reference.

- [jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-1222031022000002-2132220200132221-1303333212323111-0002013331302120-3230132222112210-2021112013023033-1202231222232320-0023331200213211): complete subsection reference.

- [query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-1322020301123132-0101232233030022-0123220220230112-1210123232320100-0013323120302112-1130122122233103-3203210132111123-1023001111202323): complete subsection reference.

<a id="canonical-1323203323033021-0332002222102313-0000200011322133-3030231233031313-0322122222331031-1301101301222223-0332131212321303-2300202010113100"></a>

## Next pages — request_matcher / 033300130021 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3230212322131111-0220132331111211-2130300110020101-0303031223301012-2233230022030113-3322103001131330-0200101110321222-3313131200332323)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3102122111323102-3333011122330001-3203211020002021-3311333132211201-1301113303101330-0020211323000331-1132020302200002-0132023233120303)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-1222031022000002-2132220200132221-1303333212323111-0002013331302120-3230132222112210-2021112013023033-1202231222232320-0023331200213211)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-1322020301123132-0101232233030022-0123220220230112-1210123232320100-0013323120302112-1130122122233103-3203210132111123-1023001111202323)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3230212322131111-0220132331111211-2130300110020101-0303031223301012-2233230022030113-3322103001131330-0200101110321222-3313131200332323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211001001301101-0122111003023012-2001213001122210-2122220033021123-3313311220203022-2222132301133230-2103312332202231-1312113022032201"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers — cookie_matchers / 222331203110 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers

<a id="canonical-2101331022310130-2130301011223100-1233003312101032-2010121212212202-2032131101013100-2121223300103222-1213131121302311-1301002123010013"></a>

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
EnumExtractionComplete: false
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

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010312201310303-1201231323202210-2010322311012113-1110202121322120-3033110001312120-3130231212110310-2002311130313020-0103330301030011"></a>

## Direct properties — cookie_matchers / 222331203110 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-3310132110221202-1130122133120123-1130112232332202-0220001000031311-3203001121311010-3320320110000201-1303111331123130-3230313303320023): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-3230210312313201-3213023022132221-3111100123210231-0111203113323011-3210101331232201-1120020303331233-1333022100201332-2222121033130123): complete subsection reference.

<a id="canonical-1311012020113031-0313023010001010-3120230023110032-0202032022010203-2321202133120203-2010001332231003-3122000011302103-3300221200132231"></a>

<a id="canonical-3212030210213120-0221200210323221-1013300022102202-0223133000333031-1231302032110212-3121103301311030-0231122302113233-3011121321313212"></a>

## invert_matcher property — cookie_matchers / 222331203110 / 4

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

- [item](resources--cdn_loadbalancer--reference--group-004.md#canonical-1200030320003030-3220331002321222-3321303320212011-3020133031313003-1332311200321313-3101032020110132-2203213112000321-0020023300001013): complete subsection reference.

<a id="canonical-2300030133313333-0231323123033032-2323203322003121-0212000023021122-2021303033123020-0123021033113300-0113020100032212-0111212131302003"></a>

<a id="canonical-3312003120020022-1311103001212001-2302101302012211-2311213002303013-0320333001000301-1301202223001021-0022220220100202-0110133133303032"></a>

## name property — cookie_matchers / 222331203110 / 5

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2101312303000022-3000321303302020-2332303210312230-0010300201320103-2202012131122012-0333002222221322-0310123310303001-2331000101032202"></a>

## Next pages — cookie_matchers / 222331203110 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-3310132110221202-1130122133120123-1130112232332202-0220001000031311-3203001121311010-3320320110000201-1303111331123130-3230313303320023)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-3230210312313201-3213023022132221-3111100123210231-0111203113323011-3210101331232201-1120020303331233-1333022100201332-2222121033130123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item](resources--cdn_loadbalancer--reference--group-004.md#canonical-1200030320003030-3220331002321222-3321303320212011-3020133031313003-1332311200321313-3101032020110132-2203213112000321-0020023300001013)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3310132110221202-1130122133120123-1130112232332202-0220001000031311-3203001121311010-3320320110000201-1303111331123130-3230313303320023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331112300111132-1221033000123123-1122322113023020-0100003003331101-3130221002021220-3101333102312121-2113223111133110-2031122103022033"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present — check_not_present / 312200021302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3230212322131111-0220132331111211-2130300110020101-0303031223301012-2233230022030113-3322103001131330-0200101110321222-3313131200332323)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-2233033203230111-1001330232200020-0000200011212110-1320010300023222-2222312200031011-3011230333020120-2332033031102323-1131302000131021"></a>

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

<a id="canonical-1220111220323001-3332011100320231-3100031221201132-1210132001003303-2111032320221220-1200001111033130-2002322202320110-1210011210111031"></a>

## Direct properties — check_not_present / 312200021302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010123123333021-0132113021210022-1031101323333022-2333133030100100-1011111020200230-1022103002311201-0013032223320011-2311213011202033"></a>

## Next pages — check_not_present / 312200021302 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3230212322131111-0220132331111211-2130300110020101-0303031223301012-2233230022030113-3322103001131330-0200101110321222-3313131200332323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3230210312313201-3213023022132221-3111100123210231-0111203113323011-3210101331232201-1120020303331233-1333022100201332-2222121033130123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321213212200323-1323311100323203-1330312123110111-3111131032010232-3302301012332020-1032322030200121-2222223133222212-3013221111020232"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present — check_present / 303003222222 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3230212322131111-0220132331111211-2130300110020101-0303031223301012-2233230022030113-3322103001131330-0200101110321222-3313131200332323)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-3133113331112333-1010123303023323-2312000303013001-3110233021113313-0101333102220321-0000023322121233-2001201332200302-3330111111313311"></a>

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

<a id="canonical-1231033000000101-0033233032030112-3002323323003230-0332011321110022-0011322201031301-3222023003332332-1030110120313201-0322313110210211"></a>

## Direct properties — check_present / 303003222222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2132321103231233-1122111030001013-0213131132230100-0332330333322222-3321122032212123-1220102110030012-1132321231011122-3300222101200211"></a>

## Next pages — check_present / 303003222222 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3230212322131111-0220132331111211-2130300110020101-0303031223301012-2233230022030113-3322103001131330-0200101110321222-3313131200332323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1200030320003030-3220331002321222-3321303320212011-3020133031313003-1332311200321313-3101032020110132-2203213112000321-0020023300001013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131210303233232-1313133133010131-0322030321101220-1002100231111101-3221012102120201-0031123132202313-0310310311313002-1122201003310212"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item — item / 131132203322 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3230212322131111-0220132331111211-2130300110020101-0303031223301012-2233230022030113-3322103001131330-0200101110321222-3313131200332323)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item

<a id="canonical-3001232121333231-2202213123123020-0001312000201312-3003003312230112-2311110112213211-0122203002331000-1312322011030232-1300332013203223"></a>

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

<a id="canonical-1001213321013003-3220213221220011-1000030203011323-1010023133002332-1112101331120101-1120121320120102-0012232213322022-0213120313011331"></a>

## Direct properties — item / 131132203322 / 3

<a id="canonical-0123221132313211-2212001013120231-0310110331321031-3022020212333313-1033000133201110-0212233321130230-2013131333203211-0111230101201120"></a>

<a id="canonical-0203110023211320-2220010001300230-3022331321313201-1313031100002133-3310132011123111-2120111111200303-0133012301213001-0003222310122012"></a>

## exact_values property — item / 131132203322 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3132223233332102-0313313300122230-3120313212311313-2101203321330120-1020130103303133-2223003312332232-1313300201221313-0123230013333302"></a>

<a id="canonical-0301303201222010-1102021111310210-1312301312233003-3311303022310313-3131103010022213-1103331013310113-0223213323103000-3012002331132221"></a>

## regex_values property — item / 131132203322 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2313132033100032-0013013111023230-3330012333302322-0232213003230001-2132010010210200-1131130330033010-3302022311311202-3330312131211032"></a>

<a id="canonical-3231331220210110-3100233130230002-2212232203023103-3123231032310300-0011001102310213-1210320202100200-1213122222322032-2013210001102032"></a>

## transformers property — item / 131132203322 / 6

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
EnumExtractionComplete: false
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

<a id="canonical-1232312000330333-1230123322200021-3233131201000101-1131311101032112-2012231231313203-2013013002110323-0220112321113110-3320002123111021"></a>

## Next pages — item / 131132203322 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3230212322131111-0220132331111211-2130300110020101-0303031223301012-2233230022030113-3322103001131330-0200101110321222-3313131200332323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3102122111323102-3333011122330001-3203211020002021-3311333132211201-1301113303101330-0020211323000331-1132020302200002-0132023233120303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000102102220203-2011131232223222-1023223322331233-3303033013331302-1221202100102333-2322111002022111-3213232221113302-1330102001323231"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers — headers / 110210201032 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers

<a id="canonical-2030211013311003-0223231211012002-3130222033312010-1132030121121011-3300332210031033-1122303330300221-2121211303133333-2233323303103212"></a>

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
EnumExtractionComplete: false
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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222031300112101-0032202311322103-2211232321231211-2003221113013213-0212230322312132-3222020113233001-1200023333200323-3310310321310030"></a>

## Direct properties — headers / 110210201032 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-3013330211330200-1332100111313312-3223301013331132-3300323101032322-0103112333233020-1331132121012333-1122211203320022-2133311030333323): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-1332033201222120-0201131233311213-0313333331122123-2101110323101012-0130102222210233-1133132231102330-0231013213010211-2310330212100122): complete subsection reference.

<a id="canonical-3120212231021221-2000312312302110-1101303313332032-2321333021001300-3211330122123232-0233302031211121-1333123000230312-3110031323330220"></a>

<a id="canonical-2220223200313331-0003013212320111-3101320000220223-2310313030002223-1001300321101231-3312331303203103-3103133101001002-2330211130302123"></a>

## invert_matcher property — headers / 110210201032 / 4

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

- [item](resources--cdn_loadbalancer--reference--group-004.md#canonical-1013101010232200-0100312233230123-3223332201121331-1301220113103123-3121303102203310-3300333110302200-1231223233113123-2030301232020332): complete subsection reference.

<a id="canonical-1003100203111231-2122202010011212-2311111313222011-3122302220002332-3210213121111213-2033313112102122-1032100123001231-3232230010032132"></a>

<a id="canonical-3101231312013013-3133312132232301-2303013022012011-1300303300001211-1103332333203201-3313033223103333-3013201012112201-2131103110210022"></a>

## name property — headers / 110210201032 / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1003131201010023-2000131333331131-2123221000303203-0103221003320301-1211103112212312-0122223231031003-2302311031313100-1122202131011120"></a>

## Next pages — headers / 110210201032 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-3013330211330200-1332100111313312-3223301013331132-3300323101032322-0103112333233020-1331132121012333-1122211203320022-2133311030333323)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-1332033201222120-0201131233311213-0313333331122123-2101110323101012-0130102222210233-1133132231102330-0231013213010211-2310330212100122)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item](resources--cdn_loadbalancer--reference--group-004.md#canonical-1013101010232200-0100312233230123-3223332201121331-1301220113103123-3121303102203310-3300333110302200-1231223233113123-2030301232020332)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3013330211330200-1332100111313312-3223301013331132-3300323101032322-0103112333233020-1331132121012333-1122211203320022-2133311030333323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222103310132211-0221210131221033-2011312331022213-1013002201110012-2301300231230213-3200201103011320-1231303111223011-3210131222113010"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present — check_not_present / 212321120111 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3102122111323102-3333011122330001-3203211020002021-3311333132211201-1301113303101330-0020211323000331-1132020302200002-0132023233120303)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present

<a id="canonical-1112032023320000-0130121302301302-0323300221133010-0011121311030321-3120110100211000-1012011031223333-1022302203132203-2223231303202031"></a>

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

<a id="canonical-1111313011103002-1312200310002202-0313231022312203-3003203212011331-3132331303110001-3303102133111230-3213311323302020-3033203311112332"></a>

## Direct properties — check_not_present / 212321120111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230200322200300-3123330201133333-0223232223032213-1013111223320200-1211000120333332-3023120213232210-2031003311002113-0203221021033222"></a>

## Next pages — check_not_present / 212321120111 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3102122111323102-3333011122330001-3203211020002021-3311333132211201-1301113303101330-0020211323000331-1132020302200002-0132023233120303)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1332033201222120-0201131233311213-0313333331122123-2101110323101012-0130102222210233-1133132231102330-0231013213010211-2310330212100122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233213301231313-2131023130200312-3131110200133010-1201132003332111-3101301102320210-1303213110211220-1230211312312123-2203131023211101"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present — check_present / 312330330100 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3102122111323102-3333011122330001-3203211020002021-3311333132211201-1301113303101330-0020211323000331-1132020302200002-0132023233120303)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present

<a id="canonical-0121332033002313-2302220300330132-2022200331031102-2231322003111330-3131203133012100-1020102132203333-0003033133001320-3301230010322300"></a>

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

<a id="canonical-2013230333300311-3022321011332112-1121321101223030-1120321231323023-2311012203030220-0330031223102001-1101312130111310-1200303000001201"></a>

## Direct properties — check_present / 312330330100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010303212111012-0020332203100210-1102220001011313-2121232222031221-3001020313122223-2233011013120100-3323210111003031-1011110223032120"></a>

## Next pages — check_present / 312330330100 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3102122111323102-3333011122330001-3203211020002021-3311333132211201-1301113303101330-0020211323000331-1132020302200002-0132023233120303)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1013101010232200-0100312233230123-3223332201121331-1301220113103123-3121303102203310-3300333110302200-1231223233113123-2030301232020332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310330311110013-2003213101110321-1201222312030202-1131031303213022-2112302120333322-1113202023200120-0031013013220022-1211212131330300"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item — item / 322113000000 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3102122111323102-3333011122330001-3203211020002021-3311333132211201-1301113303101330-0020211323000331-1132020302200002-0132023233120303)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item

<a id="canonical-3200113013101030-0122300101013323-1320022100030011-0132031300121030-0030133111002301-1000311211311000-1011311320102220-0112122101223322"></a>

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

<a id="canonical-0331233302200203-0331321330313030-3233023223021220-3102101332003200-2213033210132033-3313230200030212-2302300210310003-1030311112133013"></a>

## Direct properties — item / 322113000000 / 3

<a id="canonical-2022221321001032-2310211023033112-2102131130310103-2120303130111202-3103102200122000-3013323222000211-2001120222322022-1001131313310231"></a>

<a id="canonical-1332130130030220-1200331031003221-0200331001320100-2322233022011033-0010333121200031-0310230202321320-3121033311333110-0223112323223231"></a>

## exact_values property — item / 322113000000 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2321301030301031-2211003231100032-2230000112210301-2110132001100310-0011302110132210-2231302100020323-3001201313132330-2121331310311013"></a>

<a id="canonical-1122332000223322-3113213202222302-1001222311203200-2121331322201032-1011030310220210-1232332003232132-0113310310131013-3211231003001230"></a>

## regex_values property — item / 322113000000 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0120312213110122-3232233120221122-2032133002100002-2110223100210211-0201201320231022-2202203103013212-1132002112002312-1111122002323203"></a>

<a id="canonical-2230132101201102-3103310313033031-1203201121000333-2132013131213123-3123012212311303-3133201231011221-3321230201310230-0123002310201021"></a>

## transformers property — item / 322113000000 / 6

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
EnumExtractionComplete: false
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

<a id="canonical-0120232223203330-1222322200331102-2332221030132123-0112220110313000-1331013121012321-0101211023133031-0223100201320123-1203310202233110"></a>

## Next pages — item / 322113000000 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](resources--cdn_loadbalancer--reference--group-004.md#canonical-3102122111323102-3333011122330001-3203211020002021-3311333132211201-1301113303101330-0020211323000331-1132020302200002-0132023233120303)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1222031022000002-2132220200132221-1303333212323111-0002013331302120-3230132222112210-2021112013023033-1202231222232320-0023331200213211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212022301133012-1212210331323003-2331130330011301-2301311131033023-1011121231202330-3132312121212131-3312122111320031-0120213333231332"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims — jwt_claims / 221231322210 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims

<a id="canonical-3302313101313332-0333020133021212-3133012201102221-1123022113223210-1203132210101102-2223101310201120-0030222313300220-0322113103331100"></a>

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
EnumExtractionComplete: false
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

Terraform syntax:

```terraform
jwt_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330012103023131-1222130212011200-1030323133032213-0011212110323112-0132000023110231-1320002101212010-3323001210203130-1002011130312020"></a>

## Direct properties — jwt_claims / 221231322210 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-0102331231033001-0300002103030310-2023201023223020-2232020203031302-3232210013332121-3203321111003100-3203312110113331-2133210211021110): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-0023021030200202-1012131123022010-3201200301003111-3310120133311312-0331101122100002-2301003032000133-1013003332112003-0321322031003012): complete subsection reference.

<a id="canonical-2020033230202133-2222203301002023-3313123113200130-0310133010032333-2221102213112213-2001320222011022-1313223332003113-2331122202222010"></a>

<a id="canonical-0300330001123222-0333212311212233-0223112331133300-0311012113100221-0321030022102201-2201033111022130-0200320232332320-1021003223102011"></a>

## invert_matcher property — jwt_claims / 221231322210 / 4

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

- [item](resources--cdn_loadbalancer--reference--group-004.md#canonical-3330013301211313-2101332321331223-2100332033213202-1030122121020322-0231031123011100-0230232322210311-3303022030232331-1231211020022301): complete subsection reference.

<a id="canonical-0110103202003330-1003312120101033-0100332023200123-1330310001022312-3303201130300133-2133322012031111-2200002211213331-0111110232111001"></a>

<a id="canonical-0231231210330332-3013003033102230-3112330112212232-1311132130020320-2222202231332113-3310322131231013-1200321230020221-3030132322102322"></a>

## name property — jwt_claims / 221231322210 / 5

Type: `"string"`. Optional.

JWT Claim Name. JWT claim name.

Upstream description:

JWT claim name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1023122000130030-0001013220010131-1113231112211212-1232210332303020-0131113132220131-1121011213020310-3321201213032113-3301211110001110"></a>

## Next pages — jwt_claims / 221231322210 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-0102331231033001-0300002103030310-2023201023223020-2232020203031302-3232210013332121-3203321111003100-3203312110113331-2133210211021110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-0023021030200202-1012131123022010-3201200301003111-3310120133311312-0331101122100002-2301003032000133-1013003332112003-0321322031003012)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item](resources--cdn_loadbalancer--reference--group-004.md#canonical-3330013301211313-2101332321331223-2100332033213202-1030122121020322-0231031123011100-0230232322210311-3303022030232331-1231211020022301)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0102331231033001-0300002103030310-2023201023223020-2232020203031302-3232210013332121-3203321111003100-3203312110113331-2133210211021110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131220233202102-0310003131302210-0211221333021301-2320003222300030-1311323211000023-1023120332311102-3102021323303320-1311222213313320"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present — check_not_present / 032020201113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-1222031022000002-2132220200132221-1303333212323111-0002013331302120-3230132222112210-2021112013023033-1202231222232320-0023331200213211)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-1022232301111311-3013102223111011-1001202103212222-2311113200223113-0120022011000010-3022311131030223-3210010102001333-1333000310102332"></a>

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

<a id="canonical-1000113132121103-3102020211031113-1320320233111110-3003020100013022-3110210112313013-0330231102312221-1332120100101310-0213033232102133"></a>

## Direct properties — check_not_present / 032020201113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331232132112130-0113110020100002-3301300200100322-1300120111000122-2223322203130130-3210122201021312-1320202310113301-0112032322323100"></a>

## Next pages — check_not_present / 032020201113 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-1222031022000002-2132220200132221-1303333212323111-0002013331302120-3230132222112210-2021112013023033-1202231222232320-0023331200213211)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0023021030200202-1012131123022010-3201200301003111-3310120133311312-0331101122100002-2301003032000133-1013003332112003-0321322031003012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012031031000020-1110233332220030-3232302210331032-0231100023001031-0330301210121331-0100130330020021-1100321300313322-0230113313021020"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present — check_present / 322203303011 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-1222031022000002-2132220200132221-1303333212323111-0002013331302120-3230132222112210-2021112013023033-1202231222232320-0023331200213211)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present

<a id="canonical-0231322232231100-0300001322332101-0003122003233031-0031020333312222-2220200312112031-0131010102321120-3000301332330203-1002112133003022"></a>

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

<a id="canonical-0303200031120302-1332012002030230-1211233102122330-1003000011000201-1030231211213021-0122110031330202-1031222200011202-1300013030312133"></a>

## Direct properties — check_present / 322203303011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203230203130221-2222023331222332-2222220303210200-1133202321332211-3022001323311223-2311330022003123-3231201203233110-2302322231222023"></a>

## Next pages — check_present / 322203303011 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-1222031022000002-2132220200132221-1303333212323111-0002013331302120-3230132222112210-2021112013023033-1202231222232320-0023331200213211)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3330013301211313-2101332321331223-2100332033213202-1030122121020322-0231031123011100-0230232322210311-3303022030232331-1231211020022301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133102320002301-2313120213330220-0111001312220122-3200223130113120-0213303230312003-3311202210021330-0221223311131112-1021022311111013"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item — item / 201012113130 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-1222031022000002-2132220200132221-1303333212323111-0002013331302120-3230132222112210-2021112013023033-1202231222232320-0023331200213211)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item

<a id="canonical-3022302031223331-1001122021032332-0201011001022321-3112002303001302-2212012001321003-1011033012201232-3333201221001010-2011103202322032"></a>

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

<a id="canonical-2302310332310130-2213120212100011-0301023332313112-3312010230332212-0103013322022223-2211303013303110-3132210311022010-0203000333320332"></a>

## Direct properties — item / 201012113130 / 3

<a id="canonical-0223231102231032-1303131213203002-3321311010000031-3200333023213000-3310333123232133-2212202320203020-2013111030123301-3312013000031332"></a>

<a id="canonical-1233111121232302-1210230021202311-1112012202003201-2022330233223122-1323100023001310-1101011003102020-3222003131120023-3300232132212101"></a>

## exact_values property — item / 201012113130 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3221332231033013-2132032031031312-2311223220130021-0202303322230021-1221322103132323-1210113311000111-0222230020311032-3212112201001213"></a>

<a id="canonical-0000022033032302-1011220010031121-0231301000332333-0200022210300302-2000202212301233-2203332100210032-2123202300321032-3113113113311011"></a>

## regex_values property — item / 201012113130 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0302203211131313-0301203332312320-2131200002300232-3311322233301323-0130013332013000-0001332112013313-3121312233310021-2020231020213203"></a>

<a id="canonical-1121300021022321-2211121022221321-3123212213312100-3220002302302131-0222002123121123-0020301121313021-1123010302030113-0130023223033103"></a>

## transformers property — item / 201012113130 / 6

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
EnumExtractionComplete: false
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

<a id="canonical-1202332310013222-1030300233010022-0003122212113303-0032033223021211-1030000002100231-0102232332321302-0033011332013021-0211303230231112"></a>

## Next pages — item / 201012113130 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](resources--cdn_loadbalancer--reference--group-004.md#canonical-1222031022000002-2132220200132221-1303333212323111-0002013331302120-3230132222112210-2021112013023033-1202231222232320-0023331200213211)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1322020301123132-0101232233030022-0123220220230112-1210123232320100-0013323120302112-1130122122233103-3203210132111123-1023001111202323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211120333322311-1321300133331311-0001222301011302-3203113000120103-1020101231013033-3233330311112000-3202002122332320-2002003131330301"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params — query_params / 032030121002 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params

<a id="canonical-0310131303132120-0223020212201222-0100211023220010-1012022230331123-1333123121021222-0111020112032012-0112213012331213-3301101232312310"></a>

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
EnumExtractionComplete: false
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

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3030003331233122-2232101113221121-3331302132233112-0202331033103301-3330320002301200-1032113321020021-0202013012231213-0021003110013031"></a>

## Direct properties — query_params / 032030121002 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-0330321131102030-0033002013122203-0231300323110020-0300231200221103-2110132213310113-3301002100223320-3000013121321011-0311221003032221): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-1133012211221310-1300322030210200-2110320231021302-0311110011322133-1200300132101123-2031120320031013-3212232032103023-0333000101101032): complete subsection reference.

<a id="canonical-1021010001200101-2133022101023222-1230220231010233-0133203131133133-2301301310002020-2102101302230033-3322311013331122-1200302232212221"></a>

<a id="canonical-2221200030200333-2003113112023220-0310100011202330-0013112310212100-1033120120113130-0310023300201202-1001222221323133-1331320100031212"></a>

## invert_matcher property — query_params / 032030121002 / 4

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

- [item](resources--cdn_loadbalancer--reference--group-004.md#canonical-1000202212012101-0103332322113231-2131321203103312-0103211110223131-3031212010322331-1203011223033103-2010010120200132-1232133231001110): complete subsection reference.

<a id="canonical-2230210103020003-0331123032212102-3021013323302123-2113130000303032-1011103103231210-0212111020100100-1232022213311132-1222123001113121"></a>

<a id="canonical-1211212110112102-1133001232221232-2003131303100113-3121221003032002-2122102302013332-3123113000130200-3330223111302223-1030000201110120"></a>

## key property — query_params / 032030121002 / 5

Type: `"string"`. Optional.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1213331231303233-1111301101021031-2101020212321020-3032121233321023-1332032323231022-1202323122320331-2132030113212211-1311321113320333"></a>

## Next pages — query_params / 032030121002 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-0330321131102030-0033002013122203-0231300323110020-0300231200221103-2110132213310113-3301002100223320-3000013121321011-0311221003032221)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present](resources--cdn_loadbalancer--reference--group-004.md#canonical-1133012211221310-1300322030210200-2110320231021302-0311110011322133-1200300132101123-2031120320031013-3212232032103023-0333000101101032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item](resources--cdn_loadbalancer--reference--group-004.md#canonical-1000202212012101-0103332322113231-2131321203103312-0103211110223131-3031212010322331-1203011223033103-2010010120200132-1232133231001110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0330321131102030-0033002013122203-0231300323110020-0300231200221103-2110132213310113-3301002100223320-3000013121321011-0311221003032221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200112200121032-0103211101330212-3223233330223113-1020131032000221-2320010213003321-2022210013213020-2311222132223000-1202333002232031"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present — check_not_present / 121221030301 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-1322020301123132-0101232233030022-0123220220230112-1210123232320100-0013323120302112-1130122122233103-3203210132111123-1023001111202323)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present

<a id="canonical-0122330020113020-0320131011032230-0203203102230101-2002121233023130-0231032031013301-3000002222033323-1010313110133012-2013230210111311"></a>

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

<a id="canonical-0213213203303311-3122303101200202-1233312200210232-2220023121210033-1322000210231100-2103003012123310-3321230302131111-1313231123311033"></a>

## Direct properties — check_not_present / 121221030301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113212222100023-2102212112100110-0203010013113012-2321323122221322-2013023320200212-2330221112133331-1303021231133022-1103111031120113"></a>

## Next pages — check_not_present / 121221030301 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-1322020301123132-0101232233030022-0123220220230112-1210123232320100-0013323120302112-1130122122233103-3203210132111123-1023001111202323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1133012211221310-1300322030210200-2110320231021302-0311110011322133-1200300132101123-2031120320031013-3212232032103023-0333000101101032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112023330313102-1330033333223313-0120031322101030-3222132030310002-1231220211321302-1310233022200323-3011203013210310-0103312000112222"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present — check_present / 031130310002 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-1322020301123132-0101232233030022-0123220220230112-1210123232320100-0013323120302112-1130122122233103-3203210132111123-1023001111202323)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present

<a id="canonical-2210100102223012-0012030000033132-2302002133032332-3000102300131311-1120202233322003-3021130011033010-0213120332103321-1222132011012331"></a>

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

<a id="canonical-1320303200023222-2210001133201321-1322020301213312-2121222221223211-1200300120303031-2321000011313110-3200312101301202-3233002232001212"></a>

## Direct properties — check_present / 031130310002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201300310023031-0031110220022010-3131010100302110-0132311021331000-0300133220112110-0032012210322302-3310222320013101-3003023332232101"></a>

## Next pages — check_present / 031130310002 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-1322020301123132-0101232233030022-0123220220230112-1210123232320100-0013323120302112-1130122122233103-3203210132111123-1023001111202323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1000202212012101-0103332322113231-2131321203103312-0103211110223131-3031212010322331-1203011223033103-2010010120200132-1232133231001110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312001233032211-1130203123020231-1211121113233331-1122100310202023-1333333201010012-0121331213100333-3301211120331300-2020111001312311"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item — item / 021002111033 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-2310121313030302-0320133110312321-2211301230212100-2123220112012012-2113312300221121-1333301030222222-0201000121122112-0321301123033110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-1331023112002203-3331323011320302-2112013023331313-0321301122223200-2300131122233210-2310011100122210-1011121023002131-1321020001002130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](resources--cdn_loadbalancer--reference--group-004.md#canonical-3311111011000200-2211102020113000-0002210230031103-2333110210002132-2011031231101201-2300123301300120-0320101322310221-0000001020312003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-1322020301123132-0101232233030022-0123220220230112-1210123232320100-0013323120302112-1130122122233103-3203210132111123-1023001111202323)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item

<a id="canonical-2201230101103102-3133120103301020-2231222023200132-3022323131311331-0010000222131301-0220023331121230-3233020302130020-1323200202030023"></a>

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

<a id="canonical-3012131230202133-1120202112020223-0030213233201210-2312000110213023-1011203112003323-2103333232302231-2130131123233030-0032031022222131"></a>

## Direct properties — item / 021002111033 / 3

<a id="canonical-0332121001032213-1120101031332030-3001221211123201-3020112113310320-2201200321010110-2102111020122010-1210300333321012-3221131332313122"></a>

<a id="canonical-2110102310102200-1323321220113121-1221122023331011-0110300133110333-2221231001310033-0333122231320200-2022121111331111-0220100311002113"></a>

## exact_values property — item / 021002111033 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1003330012012102-3003123313210233-0021231313002213-0301031023333110-2232121121311101-2132112002300323-3311231203130131-0032023323201222"></a>

<a id="canonical-0222000230333013-2231112332310113-2022313033301233-0202223330230022-3320130301113120-2030200333332211-1310202103002113-0030100002101233"></a>

## regex_values property — item / 021002111033 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0111220311030032-3310031202010333-2331230210210011-2222310002030330-0233010122003123-0302322033022231-1213233013120021-0023313310220010"></a>

<a id="canonical-1120223302011010-1002321101013213-0332113333310232-1202021030213302-2001003132212220-1220103302231311-2333310220002001-0010113100221013"></a>

## transformers property — item / 021002111033 / 6

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
EnumExtractionComplete: false
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

<a id="canonical-0111222002202110-0300101132303202-0202111313201220-2230003002023302-2330220101231211-0002022032213211-0312011010311103-3013032122232320"></a>

## Next pages — item / 021002111033 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](resources--cdn_loadbalancer--reference--group-004.md#canonical-1322020301123132-0101232233030022-0123220220230112-1210123232320100-0013323120302112-1130122122233103-3203210132111123-1023001111202323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2333012122103001-0111232101323001-2333210301300122-3302011003223222-3002023122303203-1103001221000201-1310203202022230-2010003132313013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101130301122132-2100313133112110-2202321301310103-1322230202233133-3132202110332100-0213321302110110-1031112230122030-0330011303232032"></a>

## api_rate_limit.custom_ip_allowed_list — custom_ip_allowed_list / 223022200030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- api_rate_limit.custom_ip_allowed_list

<a id="canonical-0021013103213300-0232231032021332-0120122101130210-2130120003121012-0133222120112322-0100210000321331-2202300023221313-0132022203320221"></a>

Type: `"object"`. single nested block, Optional.

IP Allowed list using existing ip\_prefix\_set objects.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3031310220323122-3230302333113211-0303311100333222-2133202231102300-2013332002322230-1312331023212121-2211200101020012-3012032300311201"></a>

## Direct properties — custom_ip_allowed_list / 223022200030 / 3

- [rate_limiter_allowed_prefixes](resources--cdn_loadbalancer--reference--group-004.md#canonical-3131203222110020-2013332301110203-3003111101121002-1133332023223301-0233231233230211-0222101201000233-3120001000122223-3232112232212123): complete subsection reference.

<a id="canonical-0023321010330020-0010101101012121-0203112221333301-2223011201022110-3011321010300131-0330031121301303-1302021023133323-0301223101131110"></a>

## Next pages — custom_ip_allowed_list / 223022200030 / 4

- [api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](resources--cdn_loadbalancer--reference--group-004.md#canonical-3131203222110020-2013332301110203-3003111101121002-1133332023223301-0233231233230211-0222101201000233-3120001000122223-3232112232212123)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3131203222110020-2013332301110203-3003111101121002-1133332023223301-0233231233230211-0222101201000233-3120001000122223-3232112232212123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033101122132331-2022201130200130-3031230302113023-1032303203213033-2122330311131330-1300302331301111-2011221202000232-3310113203210111"></a>

## api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes — rate_limiter_allowed_prefixes / 011321132322 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.custom_ip_allowed_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-2333012122103001-0111232101323001-2333210301300122-3302011003223222-3002023122303203-1103001221000201-1310203202022230-2010003132313013)
- api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-3312131211310000-2322033231320103-0301201302313002-3321130013222222-3001202233220001-1031331203222012-0302033030102222-3331200121110000"></a>

Type: `"object"`. list nested block, Optional.

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Upstream description:

References to ip\_prefix\_set objects. Requests from source IP addresses that are covered by one of
the allowed IP Prefixes are not subjected to rate limiting.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0012322033312200-0033121210022231-1211112122100202-1333212100022121-1111302101322333-0301220203321032-3120111010021010-1210100021101321"></a>

## Direct properties — rate_limiter_allowed_prefixes / 011321132322 / 3

<a id="canonical-3202130021213100-0332213111312130-3013203132210221-0002322222313003-0101110211113031-2302011301022222-1201330200120031-2311131213330003"></a>

<a id="canonical-1021112023103023-2321100101011301-2303301222320222-1121110322010103-0121311111313110-2323323013001200-1213323230302001-0231330211202133"></a>

## name property — rate_limiter_allowed_prefixes / 011321132322 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3122100202013221-3121323200003002-0300332113222130-2201310203031010-0132330231003111-2302113002130311-0220032032211321-3222012211200220"></a>

<a id="canonical-1031122210302202-1012232030311011-0200232302131331-3222332002001012-1302321020230221-3011230011303103-0301233220310121-3122101232113133"></a>

## namespace property — rate_limiter_allowed_prefixes / 011321132322 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3203220232101010-2222003332210002-3222023000011321-0013112112322330-0223333210201120-3001230033131332-2011220223200130-3012303213102000"></a>

<a id="canonical-0022330221230311-0322222113210223-0321103111310321-3012332200331112-2032311323303312-0233331200123113-0320113013012132-1220120113300332"></a>

## tenant property — rate_limiter_allowed_prefixes / 011321132322 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1330330120132212-0011022300131331-1000122123030232-0313302322201311-1113222002320312-1232222230011302-0031332010302233-2320101211102131"></a>

## Next pages — rate_limiter_allowed_prefixes / 011321132322 / 7

- [api_rate_limit.custom_ip_allowed_list](resources--cdn_loadbalancer--reference--group-004.md#canonical-2333012122103001-0111232101323001-2333210301300122-3302011003223222-3002023122303203-1103001221000201-1310203202022230-2010003132313013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1131100213132110-3310333232222331-2231310112312012-2311333102233132-3301112001213202-3003223102201221-0101320232303301-1030332212123201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033210200301322-3131111102220022-0112222203200212-2230221201322320-0030100111000211-2031020320011122-2213011323213313-1020200033103000"></a>

## api_rate_limit.ip_allowed_list — ip_allowed_list / 033110230332 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- api_rate_limit.ip_allowed_list

<a id="canonical-2113300122210002-2102300000023333-0331023230211021-0223322202311012-3201213122010312-0300301003132313-1030013233203331-1200121302032120"></a>

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

<a id="canonical-2103113232223202-0303213100200333-2232230131132000-2230322031022211-2202303031023031-2023011302100122-0231133311111311-3102100223002303"></a>

## Direct properties — ip_allowed_list / 033110230332 / 3

<a id="canonical-1200223231301323-2332201332013020-0330100332012013-1120102211210003-0122333132211330-1300313213201003-0223231323221120-0011111203200113"></a>

<a id="canonical-1200300221333101-1221123212101000-2121113000023232-0130320311201000-3122322020121013-3331133032003100-3231313322032102-1032332301210023"></a>

## prefixes property — ip_allowed_list / 033110230332 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1013103232101322-3011301120303030-2312212000010011-2320002301221300-0030230012201332-2213132132323320-2133211320121133-2131322122132203"></a>

## Next pages — ip_allowed_list / 033110230332 / 5

- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3313313301303203-1122102000003112-1022231311223221-1302133120330011-1231301301331313-1323332102200203-1110313222013213-0023131100032020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313110330013233-3002233030133013-0231320230020002-2312022131222300-0210030312113220-1332312033320221-3012200112011023-2221122132113030"></a>

## api_rate_limit.no_ip_allowed_list — no_ip_allowed_list / 323033031312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- api_rate_limit.no_ip_allowed_list

<a id="canonical-2211020100021201-2303322033023302-1320120123310132-3100313000103030-2212300132032222-0220322212111323-2300230131130123-0232003231100213"></a>

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

<a id="canonical-2013003100321010-3230321332132123-0202120212310211-3021320010213333-1111221002030332-0223322103031132-2231211213033330-1030133021300300"></a>

## Direct properties — no_ip_allowed_list / 323033031312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213223033030223-1210123220033001-1020131303113312-0310110201212211-3131330133222331-2213031023111110-1220203103221233-1030121022120003"></a>

## Next pages — no_ip_allowed_list / 323033031312 / 4

- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130113223231313-1232032120232210-3120213121111030-1033213120032003-1313000322302322-2002201312020311-2203023112132210-1222213211231232"></a>

## api_rate_limit.server_url_rules — server_url_rules / 012310222022 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- api_rate_limit.server_url_rules

<a id="canonical-3020022312300323-0013213333320033-1033311032012200-3232300302011331-2210202121020323-0211000203231233-2300122221003000-3011012130232013"></a>

Type: `"object"`. list nested block, Optional.

Ordered domain or base-path rules for path-scoped rate limiting. Each rule must choose exactly one
rate\_limiter\_choice: inline\_rate\_limiter or ref\_rate\_limiter.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

Terraform syntax:

```terraform
server_url_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231332210221220-3201032202021320-0000103331201202-1230120303110012-2330311302311111-0110112033310311-2220210231331110-0311023233112132"></a>

## Direct properties — server_url_rules / 012310222022 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-004.md#canonical-1123020333112122-2130131011232202-0220230110200101-0130122132022212-0102201333003233-3220222303131213-1000112121212301-1323210210003012): complete subsection reference.

<a id="canonical-3210002020300133-0133200100302131-1123023033112033-2130113333333101-2223123301032112-3122110100113003-3222303010032033-1101231111203010"></a>

<a id="canonical-0110312120022200-3313033310101122-0002113010322201-1121033301330230-3002121320123213-3211201220132033-0300022130111003-1333101222012001"></a>

## api_group property — server_url_rules / 012310222022 / 4

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
EnumExtractionComplete: true
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

<a id="canonical-3132312122320133-1002321111011313-0103002220322023-1023213311000220-0030301100121213-0313122022221030-0033132322323032-0202131021200230"></a>

<a id="canonical-3021030113100112-2231332230222321-3001210220013123-0133301023031020-1212212201202121-1102010122332002-2310020011232333-2100232221333030"></a>

## base_path property — server_url_rules / 012310222022 / 5

Type: `"string"`. Optional.

Base Path. Prefix of the request path.

Upstream description:

Prefix of the request path.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

- [client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311): complete subsection reference.

- [inline_rate_limiter](resources--cdn_loadbalancer--reference--group-005.md#canonical-0321031211113300-3020232113001120-2211311332112302-3102201332200033-3301102331030312-2110201132013202-3010222101121131-0103331231111021): complete subsection reference.

- [ref_rate_limiter](resources--cdn_loadbalancer--reference--group-005.md#canonical-1210111123202013-2212101212113221-1330130203301223-2331201130033000-0332303011120301-2031002033003132-3203032332303101-2213233012110233): complete subsection reference.

- [request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202): complete subsection reference.

<a id="canonical-0221211213300122-1111333233102103-3021203101113132-0013002031331230-2013313300030231-0200231113011000-1302022231323021-3013302311233211"></a>

<a id="canonical-0331223121230122-2202102230123223-1110232312021212-0022323223201312-3003031010213232-1302300321301111-2030331200210000-0110201320312021"></a>

## specific_domain property — server_url_rules / 012310222022 / 6

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3313202200112200-0110322110122332-3210111002211112-0232233102332220-2202000110013020-1012112112032210-0001211123003332-2032200020231102"></a>

## Next pages — server_url_rules / 012310222022 / 7

- [api_rate_limit.server_url_rules.any_domain](resources--cdn_loadbalancer--reference--group-004.md#canonical-1123020333112122-2130131011232202-0220230110200101-0130122132022212-0102201333003233-3220222303131213-1000112121212301-1323210210003012)
- [api_rate_limit.server_url_rules.client_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-2311131213110020-1302021333201132-0301310233013210-0103312031332200-2021112223001121-1223131112200112-1200311113111123-2113222021021311)
- [api_rate_limit.server_url_rules.inline_rate_limiter](resources--cdn_loadbalancer--reference--group-005.md#canonical-0321031211113300-3020232113001120-2211311332112302-3102201332200033-3301102331030312-2110201132013202-3010222101121131-0103331231111021)
- [api_rate_limit.server_url_rules.ref_rate_limiter](resources--cdn_loadbalancer--reference--group-005.md#canonical-1210111123202013-2212101212113221-1330130203301223-2331201130033000-0332303011120301-2031002033003132-3203032332303101-2213233012110233)
- [api_rate_limit.server_url_rules.request_matcher](resources--cdn_loadbalancer--reference--group-005.md#canonical-3311203210220223-3312032012323203-1232101122033002-2030122303103232-1201231312010323-1323201102111010-1002133333322102-1100230302100202)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1123020333112122-2130131011232202-0220230110200101-0130122132022212-0102201333003233-3220222303131213-1000112121212301-1323210210003012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233332001233300-2210020301213201-0313222311001332-3312102211121322-3003310001013300-3102222122033010-2031301211012230-0202101002232300"></a>

## api_rate_limit.server_url_rules.any_domain — any_domain / 211213123132 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220)
- [api_rate_limit.server_url_rules](resources--cdn_loadbalancer--reference--group-004.md#canonical-0122231131010202-3133321230122230-0101000102220110-0201031022333212-3221130312012310-3321110221321100-0021300310231031-0321221333133033)
- api_rate_limit.server_url_rules.any_domain

<a id="canonical-3323230123333111-2220023302313332-1203322110311002-2123330210032133-0212302020012321-1321011121131313-1233101233012020-2021210012031020"></a>

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

<a id="canonical-1100130302210110-3323113202000133-0003200311001003-0120032121231130-0201310211033102-3123200210000000-0201310130010011-3201321131221210"></a>

## Direct properties — any_domain / 211213123132 / 3

This is an empty object or choice marker. It has no direct properties.
