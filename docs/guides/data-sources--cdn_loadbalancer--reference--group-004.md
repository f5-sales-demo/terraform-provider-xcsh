---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-2231003001231333-0002302120031333-3010030113331003-1310223130312101-0311101323011311-1132021123232222-3123301012013111-3131113210133333"></a>

## base_path property — bypass_rate_limiting_rules / 333311131311 / 4

Type: `"string"`. Computed.

Exclusive with \[any\_url api\_endpoint api\_groups\] The base path which this validation applies
to.

Upstream description:

Exclusive with \[any\_url api\_endpoint api\_groups\] The base path which this validation applies
to.

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

- [client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012): complete subsection reference.

- [request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000): complete subsection reference.

<a id="canonical-1331321011203123-0030211030001321-0311000002203122-0231003231333000-3133322220131302-1232001212002310-3132103202331122-3202222110102110"></a>

<a id="canonical-2133220213002120-0023021102012320-2313300113201210-1322321101202001-0322111011311003-1132132130310021-2020110002211221-2103230113123223"></a>

## specific_domain property — bypass_rate_limiting_rules / 333311131311 / 5

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain. For example:
api.example.com.

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

<a id="canonical-1111103113022113-2012232030300131-3312302303100321-1021232111301020-0100000323211320-1001203113323123-1111312311320321-3002033120101233"></a>

## Next pages — bypass_rate_limiting_rules / 333311131311 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3321121013132130-1213221130003332-3132332031332321-3030121301313201-3310201312213013-2320322002030310-3223013230021221-2000211011310222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2112011030023333-2010102213021002-2013231110322121-2100021322333200-2130332323333210-3311012302103113-3201301331331002-3131032323231113)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0020130032311000-3331131321332303-1211000310310012-3010010323103020-0232321003213200-2310003222112333-1210300310011220-3121321111322302)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2222203313123212-1123011311332100-2321223202123103-2110131302103033-1121001011130200-3233010301203301-2203332022113020-2221210123121211)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3321121013132130-1213221130003332-3132332031332321-3030121301313201-3310201312213013-2320322002030310-3223013230021221-2000211011310222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331130010111012-3131200330222103-3202131302001301-3313222133323320-1232231110320112-2110112300310310-0332121330133123-1300003333213022"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain — any_domain / 212202030011 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_domain

<a id="canonical-1010011013103020-2120131103132211-1333133113030210-3120113300031221-0123231013323002-3013102210322132-3121203302021202-3121022002210103"></a>

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

<a id="canonical-2300001000000012-1222213002131010-3003121313112011-2230013132222300-3103132010102300-2121332130103113-2013003230321203-1022012323223030"></a>

## Direct properties — any_domain / 212202030011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000023003023212-2210131232203122-2201233002011100-0230321110231133-1023200131312123-0123220323103013-1102103211003321-0222102310001022"></a>

## Next pages — any_domain / 212202030011 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2112011030023333-2010102213021002-2013231110322121-2100021322333200-2130332323333210-3311012302103113-3201301331331002-3131032323231113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321023100233021-2211303003222112-2231002323123100-0010210102200321-1101323011033212-2031100102302300-1232021331310210-1120102032302311"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url — any_url / 220123200102 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.any_url

<a id="canonical-1313213110203331-2312332332111310-3330220022232233-2231012232201303-1023121333222100-2311310330031102-3022131232031213-1022303322213330"></a>

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

<a id="canonical-0302302311011113-1222131002103200-3301332203010011-3132100211031203-2233220223202212-1030323313120221-2220001113303002-0302202312130320"></a>

## Direct properties — any_url / 220123200102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030310132311323-2123312130320011-2301321220013112-1230132123011301-3001012103201121-1211222222032312-1333010221121201-2220321233322233"></a>

## Next pages — any_url / 220123200102 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0020130032311000-3331131321332303-1211000310310012-3010010323103020-0232321003213200-2310003222112333-1210300310011220-3121321111322302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231002131131020-2000032011301130-0223133113001231-3310012120012330-3001201000113121-3120020022330201-3322300012203301-3011320321010320"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint — api_endpoint / 123122313020 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_endpoint

<a id="canonical-2012103223210031-2123101203203000-3201002002303321-0120200320302013-2312230103123131-0111320213001112-3233030331002201-0121103102301222"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

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

<a id="canonical-2233220011213110-2032013113332012-0003013101001233-3220030222100130-3302101230031311-3301130103101321-3102131013302032-3322020133003232"></a>

## Direct properties — api_endpoint / 123122313020 / 3

<a id="canonical-3110210313022331-2330032101222221-2313102013030100-3321033312230012-2130120231230013-0303000220131033-3103123123231100-2110111301103322"></a>

<a id="canonical-2323011123123321-3102012321230200-2023100003123201-2010111020231013-2211030010300121-0331313320103020-3000211010322321-1222212230312211"></a>

## methods property — api_endpoint / 123122313020 / 4

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

<a id="canonical-2211023212011211-2003233130302000-1221230021122003-1130102031201300-3032331210312031-2221301203323031-0211211201131113-2232110030212203"></a>

<a id="canonical-0010320203202002-1320222331021120-2020302030333133-0201112212300001-1000100232102323-1102130131100202-0223133111112031-0110110331212210"></a>

## path property — api_endpoint / 123122313020 / 5

Type: `"string"`. Computed.

Path. Path to be matched.

Upstream description:

Path to be matched.

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

<a id="canonical-2100110000210002-3033201102212130-0200221010302012-1211003030323320-3222303012211123-1102120203132330-1313032012121112-3232320331201013"></a>

## Next pages — api_endpoint / 123122313020 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2222203313123212-1123011311332100-2321223202123103-2110131302103033-1121001011130200-3233010301203301-2203332022113020-2221210123121211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023121232110003-3032322330133301-0221302013112031-0103201113223312-1302031010202221-0030221220123121-0102103220201303-2102112123212103"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups — api_groups / 210313111102 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.api_groups

<a id="canonical-1320100011223330-0002230231003300-0233021230032031-0322203111030131-2333103333300132-1120323110020031-3210331030230112-3130130311010212"></a>

Type: `"single"`. Computed.

API Groups.

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

<a id="canonical-0320312121000133-1113121110010130-3322031332331332-2311201001111110-0110313230021213-0212122011021303-2311130012103120-2302202203333233"></a>

## Direct properties — api_groups / 210313111102 / 3

<a id="canonical-1120130232331122-1331231203000313-3113213032021313-1313200113203000-3103301302213233-0212100220130330-2222200230212002-3212202303311322"></a>

<a id="canonical-0132012223211223-2331022300122003-3332100311132111-1220122102220323-3123102032120013-2302312133300122-3333012132033011-1012203201203120"></a>

## api_groups property — api_groups / 210313111102 / 4

Type: `["list", "string"]`. Computed.

API Groups. Group or collection configuration

Upstream description:

Group or collection configuration

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

<a id="canonical-1111131131123302-2333230110111130-3022333002320021-2013232100000332-3111332213110002-1002203320131102-2211102200231321-2232123333300212"></a>

## Next pages — api_groups / 210313111102 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010212311032011-1112222022221003-1001102001312321-0002231122310220-1311301120230310-0012120023213232-0203013002000230-2321332203130322"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher — client_matcher / 202322332212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher

<a id="canonical-3213033212232321-3100013331100133-0031021310233101-1223321003331102-1122230013103313-2011202133303313-0333100333220231-3031320110322212"></a>

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

<a id="canonical-2233113202032331-3310120021000212-3101201011302021-1030311231023000-0220102220112003-2311132122001132-2011122023223303-2332121120033101"></a>

## Direct properties — client_matcher / 202322332212 / 3

- [any_client](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0031222203132322-0100311112100122-0003203132031032-1230102023011333-3101131120132201-1013033112332321-0221310211113031-2333231022203311): complete subsection reference.

- [any_ip](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1132313332201100-3112212013201130-2203222222201121-0310000002111201-0032023232331320-1222321202332301-3221003033211130-1210203323231010): complete subsection reference.

- [asn_list](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2022101230231311-2121010323233332-0203323120330223-3022231222301222-0021201021312120-2023102022321210-0110221203222113-3222130122121003): complete subsection reference.

- [asn_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1333200011002113-0233213222022212-3323220113012323-2213310200031023-3111321201011232-1302132320112112-0031232323010120-3200211312121210): complete subsection reference.

- [client_selector](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0000033330030200-1010322233110101-1321001132031022-0013123112001112-0001212030000100-1230110321012110-1211232311132111-3211111331000033): complete subsection reference.

- [ip_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2202331131330221-2212120223133022-0100322203100030-0021303233302010-0330023230220212-1110020020013012-3012131232322003-0003212032001100): complete subsection reference.

- [ip_prefix_list](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1120001120322001-3321220222101333-1330023031322303-3000200112013302-0230023333030311-3111101332102123-0120021200202202-0203103303212010): complete subsection reference.

- [ip_threat_category_list](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0022023031203231-2012122201023333-3223331120301200-2012011002220023-1323133303010133-3303110010022101-3230112312330010-0000322210223110): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2113111133303102-0322023132020130-3120100001102013-2313123132103121-0322202000002220-1202030121031200-2203330012312000-2302133332021030): complete subsection reference.

<a id="canonical-2101212301331133-0013310132122221-3203223302003130-1201203220121111-1232223213303103-1121033233100201-3111310223233320-3132203213322023"></a>

## Next pages — client_matcher / 202322332212 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0031222203132322-0100311112100122-0003203132031032-1230102023011333-3101131120132201-1013033112332321-0221310211113031-2333231022203311)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1132313332201100-3112212013201130-2203222222201121-0310000002111201-0032023232331320-1222321202332301-3221003033211130-1210203323231010)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2022101230231311-2121010323233332-0203323120330223-3022231222301222-0021201021312120-2023102022321210-0110221203222113-3222130122121003)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1333200011002113-0233213222022212-3323220113012323-2213310200031023-3111321201011232-1302132320112112-0031232323010120-3200211312121210)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0000033330030200-1010322233110101-1321001132031022-0013123112001112-0001212030000100-1230110321012110-1211232311132111-3211111331000033)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2202331131330221-2212120223133022-0100322203100030-0021303233302010-0330023230220212-1110020020013012-3012131232322003-0003212032001100)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1120001120322001-3321220222101333-1330023031322303-3000200112013302-0230023333030311-3111101332102123-0120021200202202-0203103303212010)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0022023031203231-2012122201023333-3223331120301200-2012011002220023-1323133303010133-3303110010022101-3230112312330010-0000322210223110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2113111133303102-0322023132020130-3120100001102013-2313123132103121-0322202000002220-1202030121031200-2203330012312000-2302133332021030)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0031222203132322-0100311112100122-0003203132031032-1230102023011333-3101131120132201-1013033112332321-0221310211113031-2333231022203311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203312120232321-3322200013222122-3331032133302103-0200213320013330-2110230323021101-0001301000222022-1030002102121110-1222131022323232"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client — any_client / 102130323032 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_client

<a id="canonical-3323330213030223-0232012003300133-1321121301033220-2031032010102133-2332012221301001-0031210030333113-1231032203221301-1023312212312223"></a>

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

<a id="canonical-0231321303232300-2121131233200010-0221213132013202-1323333211202100-2232232201230000-2123122202232322-3103223222213121-3322003022211323"></a>

## Direct properties — any_client / 102130323032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313231032202031-3103232033323211-0132312003323133-3012101101133313-2303330103032130-1210330302333131-3302103113110131-1121331031322213"></a>

## Next pages — any_client / 102130323032 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1132313332201100-3112212013201130-2203222222201121-0310000002111201-0032023232331320-1222321202332301-3221003033211130-1210203323231010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031021200200212-0332023223202021-1231001010121002-1301323100333010-3012113133120202-2323203212111120-0212212233202331-0213232010020311"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip — any_ip / 303321003201 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.any_ip

<a id="canonical-2310222102130121-1031321311001123-3010310212202302-1201330011002222-2121201103103012-2123121100120001-2030001211033100-3300323323203222"></a>

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

<a id="canonical-0211312101303122-1132000223302213-0102233231211330-3131212010321123-0003331322300030-0312031321033123-1122312300132113-0323313113103321"></a>

## Direct properties — any_ip / 303321003201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101000213003232-3131133332331003-0312201013120012-3001333032222300-1022012322211320-3021300312011020-2213202121211323-2032220032020121"></a>

## Next pages — any_ip / 303321003201 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2022101230231311-2121010323233332-0203323120330223-3022231222301222-0021201021312120-2023102022321210-0110221203222113-3222130122121003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221331113222030-1222031111301202-1300031202100222-0113301231030013-0133312330102211-3000232312022231-2303311120121321-3131212300110132"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list — asn_list / 030110323030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_list

<a id="canonical-0101001231002230-0301031111121302-0300022200201112-2013112210011233-3213333232202222-2031031020221133-2100113230123231-0102323101032221"></a>

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

<a id="canonical-1033310210333332-3212122313231301-1312320300002031-1232000201102330-1302122220120000-2021332100131210-2011332001101233-0020110130200230"></a>

## Direct properties — asn_list / 030110323030 / 3

<a id="canonical-1301233310021123-0310120221030202-3212110101331312-3020211321121023-3011000232331120-1123110202000001-3233030232103331-1233323301121102"></a>

<a id="canonical-2001212330200210-3203131331003001-0111112333012330-2332012313102203-1103100010011322-0232101030230013-1231101031121223-3222012230220111"></a>

## as_numbers property — asn_list / 030110323030 / 4

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

<a id="canonical-0020221001113213-1012021100113331-2331023210021300-1321212003303020-3103030211332203-2230323220310300-0130330133022020-2123100013301103"></a>

## Next pages — asn_list / 030110323030 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1333200011002113-0233213222022212-3323220113012323-2213310200031023-3111321201011232-1302132320112112-0031232323010120-3200211312121210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022202013120311-3300013130021030-2330103200031023-1313211123101222-3311123230122112-1011323012032120-3131301333012110-3020012131122003"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher — asn_matcher / 310020323032 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher

<a id="canonical-1200102113031321-2000133133010102-3310230013121322-0233033222130312-3210211011030332-3111210011021222-1013333331123011-1111010021110031"></a>

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

<a id="canonical-1203020010321323-2301233311200303-1003222132121012-1132230201231100-3103230023032020-3202030120321301-1100020333013330-1001332210210020"></a>

## Direct properties — asn_matcher / 310020323032 / 3

- [asn_sets](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3010022013322130-0222323310323303-3023231022313220-0033032020321231-1132102230001333-1230211211003032-0201330220320212-0122121221220302): complete subsection reference.

<a id="canonical-2333103103200222-2232322023120131-0133031020103020-0232301030313013-1222321113311312-2330330022013013-2303112120002222-2330320020300233"></a>

## Next pages — asn_matcher / 310020323032 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3010022013322130-0222323310323303-3023231022313220-0033032020321231-1132102230001333-1230211211003032-0201330220320212-0122121221220302)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3010022013322130-0222323310323303-3023231022313220-0033032020321231-1132102230001333-1230211211003032-0201330220320212-0122121221220302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322233000333011-0312331123102311-1321001223200001-0112233310130030-0210110032322110-0310112002123302-0211001322113110-3012033211233332"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets — asn_sets / 100321221132 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1333200011002113-0233213222022212-3323220113012323-2213310200031023-3111321201011232-1302132320112112-0031232323010120-3200211312121210)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets

<a id="canonical-2110000122131101-0013223012013102-1010301000231132-3312210100323102-1113201210230303-2231323222203303-1131300002202222-0332323330013132"></a>

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

<a id="canonical-2302130232202020-2221111011222230-0211130311012212-1010131021231020-2121231030122203-0201010323031100-2202013320110202-3212131300101302"></a>

## Direct properties — asn_sets / 100321221132 / 3

<a id="canonical-1021311211222002-0113311310011121-0203001021301301-0202300032123031-2000133221021200-0302221022233010-2233021011010112-0220033333211232"></a>

<a id="canonical-1131301100332312-1112300233233331-0331001102300230-0210320000032020-0301230011131223-0132021113023122-2131003322122320-3103033001132130"></a>

## kind property — asn_sets / 100321221132 / 4

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

<a id="canonical-0001331313100033-0020213103313033-1202110100031010-3303031130211322-3321113012300123-3302121101021320-3121221110331111-1032001030012232"></a>

<a id="canonical-1332121102103131-1300210311213320-2222132032320132-1030333301113200-2021123233011300-1222101010233121-3130320003221200-1010023311033031"></a>

## name property — asn_sets / 100321221132 / 5

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

<a id="canonical-3202313021023011-0231212022231013-2212302210032112-1320023120111222-2011323102001300-3312203100312002-2000310221000313-0000012210313323"></a>

<a id="canonical-0010121130311323-0211020303021233-3013202222121013-3332101302102213-3122201203100302-1203211113002201-3303313321201030-2000020210003232"></a>

## namespace property — asn_sets / 100321221132 / 6

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

<a id="canonical-2210222000120001-1322320332200130-0110131233310022-3003000201023023-0230001101233023-0202310120321100-1012021333320300-1332202030130212"></a>

<a id="canonical-1302000303121313-0303220331222033-2112000120310302-3322013112020000-2321102020202112-0022223011323011-0032012102210000-0020113010211031"></a>

## tenant property — asn_sets / 100321221132 / 7

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

<a id="canonical-3301230122110333-2032002322232121-3000210301121130-0132103102012310-3122001130011133-0120032332112312-3201310103032130-0210221033103231"></a>

<a id="canonical-3331220210102131-3203322233223000-3311132231121333-1320120310210020-1103101010312310-1012003021111110-0100233212202012-2310032112231120"></a>

## uid property — asn_sets / 100321221132 / 8

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

<a id="canonical-1101002212201313-1030222320112100-1013030330333020-1333333203031122-0021331233333213-0100123201020011-1230110023332211-2210300000321312"></a>

## Next pages — asn_sets / 100321221132 / 9

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1333200011002113-0233213222022212-3323220113012323-2213310200031023-3111321201011232-1302132320112112-0031232323010120-3200211312121210)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0000033330030200-1010322233110101-1321001132031022-0013123112001112-0001212030000100-1230110321012110-1211232311132111-3211111331000033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010313320020021-0332112332122311-0001112012122210-1323103021223032-0030030100221002-0200232120012101-0103032312101112-2221100311121031"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector — client_selector / 311122211003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.client_selector

<a id="canonical-3330012131132130-1221223320233113-3230110321011330-2320221330102003-3120332113101212-0232313020110003-3123230013221321-2021131333200113"></a>

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

<a id="canonical-3102231130300023-2213032110230112-1210030213320230-2313223123303123-3310220131310130-2200032121221323-0211102331123111-3222311202012021"></a>

## Direct properties — client_selector / 311122211003 / 3

<a id="canonical-2102123303020122-1030333031322133-0323211201212110-3122132012232113-1321202210133323-0020230202220003-1300201311111002-0100232211320103"></a>

<a id="canonical-0220330211132212-3211113232310323-2201130113023021-2211211122000122-3002132130011002-0123012222112300-3132232311233210-3011322020002102"></a>

## expressions property — client_selector / 311122211003 / 4

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

<a id="canonical-3222223211013120-1101220100323332-3011200101202131-3211231103131322-0101013332231331-0010033110111002-2223012333023132-1301020200203023"></a>

## Next pages — client_selector / 311122211003 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2202331131330221-2212120223133022-0100322203100030-0021303233302010-0330023230220212-1110020020013012-3012131232322003-0003212032001100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321212310033022-3030030121103220-1313101201223303-0333323111323032-2102132333230020-2123013222133222-0121103201310331-0001331202012230"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher — ip_matcher / 312033321212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher

<a id="canonical-1331102320300210-3122031203201200-1123033231112102-1222202122111331-3203123031030110-3333233021110012-1013223032211202-2202031133312213"></a>

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

<a id="canonical-2203013321212232-2002223030110200-0011023202332313-2022033001002232-0112132323200201-0331312122233102-3221102321120121-3002300102103222"></a>

## Direct properties — ip_matcher / 312033321212 / 3

<a id="canonical-1010132321131300-0132311323321021-3321332302012101-2012220001223300-3111133031100200-2211220230232321-1013200300231313-1302023201322113"></a>

<a id="canonical-0123130032323222-1322232033023131-0232111201211033-1133112323011311-2303130132302323-3020301031233231-2122301012222203-0003301211212010"></a>

## invert_matcher property — ip_matcher / 312033321212 / 4

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

- [prefix_sets](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0021320303030223-2332313001111203-3210223333001103-3212311111033123-0322130322233001-2330322330020232-3300020101122001-2000311333023110): complete subsection reference.

<a id="canonical-1312030111232333-2001031003223111-2232012022000023-3220110300133221-3031333123211103-2211310202311220-3331311301311303-2020201231323123"></a>

## Next pages — ip_matcher / 312033321212 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0021320303030223-2332313001111203-3210223333001103-3212311111033123-0322130322233001-2330322330020232-3300020101122001-2000311333023110)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0021320303030223-2332313001111203-3210223333001103-3212311111033123-0322130322233001-2330322330020232-3300020101122001-2000311333023110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131213321223003-1201320320311222-3022002010111321-0120320232233031-3130321020100300-0212221312230211-2132111322110020-2022030221020203"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets — prefix_sets / 000320113201 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2202331131330221-2212120223133022-0100322203100030-0021303233302010-0330023230220212-1110020020013012-3012131232322003-0003212032001100)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher.prefix_sets

<a id="canonical-0233020100330310-0200212130130011-2220320001111000-0002011113021333-2012302230113311-2332313132301333-2021210100301033-0222020102233311"></a>

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

<a id="canonical-1201321130330332-0220012120212230-2121100113310210-3203001220131213-2023311123122223-3131102201200311-2200201230011013-1021131102001120"></a>

## Direct properties — prefix_sets / 000320113201 / 3

<a id="canonical-2003230200300202-0021131023003201-2001123112011333-1121323231021021-3330231120200100-1321030020120331-2203201332120131-2130322312311132"></a>

<a id="canonical-1330233123231311-0233332302310300-1123021223222303-2133203111122211-0013032123220310-1321223010000123-2301312000223120-3033203132213120"></a>

## kind property — prefix_sets / 000320113201 / 4

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

<a id="canonical-1312300303330320-1203201310111033-1332122123322022-3223210233002113-3223122011131312-2100121101213113-1323203311103301-3001012322103003"></a>

<a id="canonical-0110303010002130-0032330031000133-3211120333000023-3233313202220232-2231302101100132-3200110302001120-0112333000003000-0120001033303302"></a>

## name property — prefix_sets / 000320113201 / 5

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

<a id="canonical-0123221302011211-0010023032112303-2222102231123302-1132213113311223-3203112231111200-0331031301002330-2103111131102120-3313130010211010"></a>

<a id="canonical-0002320311323022-3132131101213120-2010101301222311-3203132331211320-3313233231122203-0100021033122320-1033120131300330-1233321130032121"></a>

## namespace property — prefix_sets / 000320113201 / 6

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

<a id="canonical-0333323032030111-1121230012012131-0113032133013111-1132330321213202-0113111132301000-3022200230132020-3331320130112233-2130031332222302"></a>

<a id="canonical-3101200303001203-0122130022122010-2030323302103000-2132013013220320-3102302311330102-2301312320300102-2200033332300321-0130232321302200"></a>

## tenant property — prefix_sets / 000320113201 / 7

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

<a id="canonical-1200122112112111-1311022310323320-3122320130110102-3330300223100221-1010031111022131-1022131311022131-0312031033301111-1322122221332321"></a>

<a id="canonical-0032213302033212-2011033011222200-0003022012012332-0233022222031213-1132031210033321-2130303311211101-2023233222200111-1302021101002013"></a>

## uid property — prefix_sets / 000320113201 / 8

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

<a id="canonical-2202002332302013-0223201002211232-0303220211222033-1131310323200112-1122223333310120-1120313003221123-0123130111003011-3010022212200002"></a>

## Next pages — prefix_sets / 000320113201 / 9

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2202331131330221-2212120223133022-0100322203100030-0021303233302010-0330023230220212-1110020020013012-3012131232322003-0003212032001100)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1120001120322001-3321220222101333-1330023031322303-3000200112013302-0230023333030311-3111101332102123-0120021200202202-0203103303212010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023312321302023-3302321210022321-0100130302021332-0222213000203210-1111031000331020-3030121223130333-1331203032113302-0233101313012320"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list — ip_prefix_list / 312110012103 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_prefix_list

<a id="canonical-3213130230332033-1301110122010313-3320111231200202-2022100220101031-3211000122202223-3102200322321332-3333022120003323-3223321133212131"></a>

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

<a id="canonical-2101010330132333-1121021120013313-3221320003021123-3300010011020211-1113222210311101-3002222010212302-1230001233301012-2212011210123001"></a>

## Direct properties — ip_prefix_list / 312110012103 / 3

<a id="canonical-2330313202122303-1310021212122010-1332001013300310-1201011322010033-3013222300001012-2132323233103120-0031232221101200-3321321111302232"></a>

<a id="canonical-0113113103012031-0101300202332301-2102101021023233-3323103012222330-0202212312211302-0100102033012203-3001132132002300-2313303333303000"></a>

## invert_match property — ip_prefix_list / 312110012103 / 4

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

<a id="canonical-0002312202200203-3120211111322010-0223231113230111-3233331021120310-0221213031103012-3302100321310233-3021103331120331-1323001223100201"></a>

<a id="canonical-1003112233200331-3122202033330031-0311333001110102-2103000221001130-3330311331331110-1303120223032230-2202211201101102-0032312033020221"></a>

## ip_prefixes property — ip_prefix_list / 312110012103 / 5

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

<a id="canonical-3003222321020002-1101333322111120-1323331330323310-2212122002100002-3122010312010211-3132211212110121-2123300110332310-0120333131333332"></a>

## Next pages — ip_prefix_list / 312110012103 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0022023031203231-2012122201023333-3223331120301200-2012011002220023-1323133303010133-3303110010022101-3230112312330010-0000322210223110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033013323002230-0003231313113321-3301101123021210-3122231321222331-0321200001100132-0220001202103100-0113202001303210-2231011130330203"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list — ip_threat_category_list / 122112231333 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.ip_threat_category_list

<a id="canonical-0131221110213010-3000333210303332-3302230330112210-0232122201120300-3011213300313110-0223113122030231-1001313213033033-3302011301021211"></a>

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

<a id="canonical-0023333232102201-0320310323321131-1131032302201301-3302113003220333-3330232310203101-3020021131022112-2203332333130121-1200301100333123"></a>

## Direct properties — ip_threat_category_list / 122112231333 / 3

<a id="canonical-2111200223223022-2320002112101313-3020323130103332-3113210203013123-1133032033110103-1312113013030013-3320031311111232-2001302013013120"></a>

<a id="canonical-0103323012311103-3303311220220301-0002303222032301-0211200103332230-1000202032121023-3211130011130300-0310032112310020-1102102232120300"></a>

## ip_threat_categories property — ip_threat_category_list / 122112231333 / 4

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

<a id="canonical-0221213233213302-3220110331211032-2300102332022233-1111222212023212-0313303122002101-1332213103331131-3113103203102200-2031300333310313"></a>

## Next pages — ip_threat_category_list / 122112231333 / 5

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2113111133303102-0322023132020130-3120100001102013-2313123132103121-0322202000002220-1202030121031200-2203330012312000-2302133332021030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320230031310202-3200300223023030-3010022013112221-1032312001310202-2200132322023311-1103022132222122-0231303212111233-0013100123231030"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher — tls_fingerprint_matcher / 022030030312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.tls_fingerprint_matcher

<a id="canonical-1120112002023110-3311032103321230-0203201022010112-3320321201322211-0200200332232130-1321102203010131-0302213331222322-3323122323102303"></a>

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

<a id="canonical-1031311330332200-3010211033133212-0111011212311313-1022302011312013-0013310221122131-3100011333223112-0132222013013012-3021133211233010"></a>

## Direct properties — tls_fingerprint_matcher / 022030030312 / 3

<a id="canonical-3200023222030231-3112312302021310-1302223132223203-2333122323033031-0123230020020223-3101013101220320-3030330322311202-0021000133112221"></a>

<a id="canonical-3101000331222232-1101021222021033-1101000131312333-0132122110001223-1103333001222101-0303213231201020-1200132101321221-0200122332302010"></a>

## classes property — tls_fingerprint_matcher / 022030030312 / 4

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

<a id="canonical-3133112111322201-0111133120121100-2213030032120212-3100210213321232-1331202121333002-1233031123302330-0333030131002312-2230322302302220"></a>

<a id="canonical-3113311322022231-3133133023121031-0203231000220131-3100322310122302-1121323201100302-1321313013220210-1201330202000101-2303112320031322"></a>

## exact_values property — tls_fingerprint_matcher / 022030030312 / 5

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

<a id="canonical-0111302001130302-2300230131200302-1033031332131112-2113210223132130-2210001203231222-0322221323231323-1022210300101033-3131022331032110"></a>

<a id="canonical-3212202231132200-1330113211111312-1132031201001102-3301133303133330-3323022321302120-2132013302203301-2032323032032222-3331100031301120"></a>

## excluded_values property — tls_fingerprint_matcher / 022030030312 / 6

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

<a id="canonical-3012000233212111-0102112021203031-3300110212222023-2333112113230230-2131122330212021-3331330323130311-2210123021300213-1303020213020022"></a>

## Next pages — tls_fingerprint_matcher / 022030030312 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3000112213313202-3331230002113213-1332230002131201-0201311011232233-3100031202210300-2200123230210230-2010321201010313-1201202113222012)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233131112200121-1110000210233122-0132133103002122-0101211330031133-2120002202232022-1322001320022303-1113022230303300-1021313120103322"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher — request_matcher / 121112320002 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher

<a id="canonical-3210202033313123-3300332300121002-0232320022110300-1010220211113000-0110201201122122-2030010010120002-2112032101123122-0202322022221223"></a>

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

<a id="canonical-2112003230213303-2113232201102130-3130113021022000-0232300303122100-2330210030330133-1301112122031003-2012030333003121-2220031010120323"></a>

## Direct properties — request_matcher / 121112320002 / 3

- [cookie_matchers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0321331013031320-2212012300101022-1203220330032033-1131222102203212-0110202013101100-3020100231130310-3302223201312112-3302022203300122): complete subsection reference.

- [headers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1030120301301200-2113032102213202-0320122122232132-1221322101011121-1221132030000123-0220331000221300-2330122322111200-1103202111301111): complete subsection reference.

- [jwt_claims](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2122302200123210-2300130301121320-2331322133200302-0312300323131112-3323113030031332-2233322230123010-3221321203202200-0333211022031121): complete subsection reference.

- [query_params](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3132310103201013-2010300221013120-3202331131232333-0032003011032022-3331302103311103-3022013200332312-3001021021131212-2011100333232220): complete subsection reference.

<a id="canonical-1331100331132312-2300333120100123-2200001221313101-0113203221233102-0100302333023001-1210022112013210-0321220131320210-1120011212000130"></a>

## Next pages — request_matcher / 121112320002 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0321331013031320-2212012300101022-1203220330032033-1131222102203212-0110202013101100-3020100231130310-3302223201312112-3302022203300122)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1030120301301200-2113032102213202-0320122122232132-1221322101011121-1221132030000123-0220331000221300-2330122322111200-1103202111301111)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2122302200123210-2300130301121320-2331322133200302-0312300323131112-3323113030031332-2233322230123010-3221321203202200-0333211022031121)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3132310103201013-2010300221013120-3202331131232333-0032003011032022-3331302103311103-3022013200332312-3001021021131212-2011100333232220)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0321331013031320-2212012300101022-1203220330032033-1131222102203212-0110202013101100-3020100231130310-3302223201312112-3302022203300122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032120311031323-0213310022130303-3111130203131333-3323013303033331-3032210310110321-1310121333323022-3132201322122023-1103033213010100"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers — cookie_matchers / 211132030100 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers

<a id="canonical-1113310013103131-0233210031303003-2010330230130103-1021212000101033-0231331112032101-2102332220102230-3011330203100311-0312102131111220"></a>

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

<a id="canonical-1131010023323120-1101322121101112-3211210302332231-3010201212133222-2310311100110023-2100030021100303-0223202003013203-2021110303010012"></a>

## Direct properties — cookie_matchers / 211132030100 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0303313001011331-3013111201302210-3221001030022100-0212032220002003-0010322213222320-2011212120232211-0111122010022133-0033222102323100): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1230001031221321-1231110213302320-3321302302332212-3212032331131000-2000311220013321-1223122322220331-2321103132210203-3211210131103000): complete subsection reference.

<a id="canonical-2203331223203210-3033120303000312-1223302300020000-2013200220313033-2001020013013002-2313033231333322-1212030132011011-0312022011222230"></a>

<a id="canonical-2031200022322320-0111110300202221-2311101330101313-0222302102213232-1223013112201232-0033201121133311-0310203121200211-0311303003322200"></a>

## invert_matcher property — cookie_matchers / 211132030100 / 4

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

- [item](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0211311231132120-3223313321010303-0011030101021130-1013312222223210-2321310020110312-0112110012000111-3312231300232311-1112130220332332): complete subsection reference.

<a id="canonical-0112131120122222-0102131031113330-3230033231120111-2202232213111210-2210300001322210-2031103202011323-1320321132001212-0322022133121230"></a>

<a id="canonical-0033130031230312-1201032233332221-1002010002133023-1230213100111033-2212122120213113-3233330332310310-2130222333213122-2033033313212213"></a>

## name property — cookie_matchers / 211132030100 / 5

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

<a id="canonical-3311000321301131-3323032132111103-2031002223230113-2130213303023001-1201313332122232-0022213311212030-2021101303300200-0322032332020103"></a>

## Next pages — cookie_matchers / 211132030100 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0303313001011331-3013111201302210-3221001030022100-0212032220002003-0010322213222320-2011212120232211-0111122010022133-0033222102323100)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1230001031221321-1231110213302320-3321302302332212-3212032331131000-2000311220013321-1223122322220331-2321103132210203-3211210131103000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0211311231132120-3223313321010303-0011030101021130-1013312222223210-2321310020110312-0112110012000111-3312231300232311-1112130220332332)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0303313001011331-3013111201302210-3221001030022100-0212032220002003-0010322213222320-2011212120232211-0111122010022133-0033222102323100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321330202210311-0101300323122020-2121322001203220-3212300233212310-0120321120112101-2133121302322131-3332210211332012-3230310210130103"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present — check_not_present / 233230013313 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0321331013031320-2212012300101022-1203220330032033-1131222102203212-0110202013101100-3020100231130310-3302223201312112-3302022203300122)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_not_present

<a id="canonical-0232103200323230-0302112110100212-1233103001313000-0200120001003020-1301023203301332-2201312102112323-0121303213332200-1001101202310332"></a>

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

<a id="canonical-2300221102033231-2101312223202212-3122220002120231-1000012101113000-0303301113111020-2212210010123321-1001220321213001-3311020222100300"></a>

## Direct properties — check_not_present / 233230013313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032232313210000-1320022301203031-1100233011321222-0120123220210033-1313210320311013-1012000133321320-2103212302230102-0032123202303230"></a>

## Next pages — check_not_present / 233230013313 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0321331013031320-2212012300101022-1203220330032033-1131222102203212-0110202013101100-3020100231130310-3302223201312112-3302022203300122)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1230001031221321-1231110213302320-3321302302332212-3212032331131000-2000311220013321-1223122322220331-2321103132210203-3211210131103000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310002132003313-2133303203203233-0122311212030332-3330020032302223-3010101003001203-3031310101322012-1131301000011303-3022210101102111"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present — check_present / 232012213000 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0321331013031320-2212012300101022-1203220330032033-1131222102203212-0110202013101100-3020100231130310-3302223201312112-3302022203300122)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.check_present

<a id="canonical-0233310330203330-1133122301000202-2103302330330022-3302320202333222-2302131131300333-1103233222323012-0100103201123311-3012102322310010"></a>

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

<a id="canonical-2202122120213322-2211110211303030-1110032320320211-3313011202302003-1322012032203103-3122100011001230-2322320210033301-1001010220121200"></a>

## Direct properties — check_present / 232012213000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210312023230001-1200121312110012-2133213022310120-3020313323202323-2302021010130131-2300103321301112-1013121233122033-2313303302220103"></a>

## Next pages — check_present / 232012213000 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0321331013031320-2212012300101022-1203220330032033-1131222102203212-0110202013101100-3020100231130310-3302223201312112-3302022203300122)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0211311231132120-3223313321010303-0011030101021130-1013312222223210-2321310020110312-0112110012000111-3312231300232311-1112130220332332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211033102310210-2032223020300032-0111130223313002-0303100230102320-2103322303032102-1010001333230022-3202200333132310-0012213000132122"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item — item / 022122102213 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0321331013031320-2212012300101022-1203220330032033-1131222102203212-0110202013101100-3020100231130310-3302223201312112-3302022203300122)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers.item

<a id="canonical-1023331123231202-1331222130023320-2210020101303122-2320212013223322-3031133012013202-0230030000202230-0122100130000331-0210113331322203"></a>

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

<a id="canonical-3330323120113222-0011210023232330-1320000313112031-1111310013112111-2132101233103200-3101312302033321-0300302333123301-2320233100133100"></a>

## Direct properties — item / 022122102213 / 3

<a id="canonical-0113100210213320-3010111212311230-0032101123132213-1123102131323120-3212101301300233-2201022101130210-1223333212313001-3031331230131221"></a>

<a id="canonical-1303220020133013-3330000110233300-3332012102221301-0232103100001000-3300022033000301-2312102002023301-3021313200021010-1313230130130002"></a>

## exact_values property — item / 022122102213 / 4

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

<a id="canonical-0313121021210021-2122203332113102-1232223113310000-0220333113123022-2202011023111222-2201221111232220-2320103101230232-0013300121322310"></a>

<a id="canonical-0313110300332202-3230131310031332-0230202221302303-2212311301330322-3132010333023000-3321323320302033-2132331321033302-0221323123120322"></a>

## regex_values property — item / 022122102213 / 5

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

<a id="canonical-2320133020331232-1023103330302333-3103033212013010-0100100010013132-2113321001002221-1031002021010302-2203321200320020-2213223303011332"></a>

<a id="canonical-1120031021103002-1221012312301113-0021010010031302-1312010211103131-1001022101122002-0121100010221231-0003112132120133-3202222123120123"></a>

## transformers property — item / 022122102213 / 6

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

<a id="canonical-3022220102113100-3022012011331013-3212301310012320-2302221003030210-0221230001002303-2330221203221202-2322321231301332-3231230202021213"></a>

## Next pages — item / 022122102213 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0321331013031320-2212012300101022-1203220330032033-1131222102203212-0110202013101100-3020100231130310-3302223201312112-3302022203300122)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1030120301301200-2113032102213202-0320122122232132-1221322101011121-1221132030000123-0220331000221300-2330122322111200-1103202111301111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130330330133003-1010323212130000-3322212101111122-2312220013332230-1313330002212101-2313031222231132-3133211322020103-1110031310323221"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers — headers / 133100231303 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers

<a id="canonical-1033001220002130-0030113313132331-0221122033122312-0020111230201223-2102112233223303-0013211101111022-0010220032002132-0003102122221302"></a>

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

<a id="canonical-0012322031223223-3100133013222002-3132320120202000-0102123230110230-2212321232332130-3230123111000023-3032200313031120-1322222201011212"></a>

## Direct properties — headers / 133100231303 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2300203302330000-1113331120312103-3012203011332233-0231220202132013-0332000211122220-1032323002122312-3121232010322011-3322202120222012): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2000300230213113-2310133032200222-1032320100112111-1121202103203001-1310131122302031-2200001230003320-3132213322103310-1222303302121103): complete subsection reference.

<a id="canonical-3100220203110231-3331220031211202-0101212000300122-0203101102101100-1221231112332231-3312333033020122-2003022000223101-3020210013031131"></a>

<a id="canonical-2001100012222130-0331123222331330-2222323203103020-2113231300322111-3020120331021310-0320103113232231-2203032212132100-0333113133232321"></a>

## invert_matcher property — headers / 133100231303 / 4

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

- [item](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1010221010113231-2210103210122133-0230121223002302-0113001011213211-1320032211031102-1003011332233332-3312131133133110-0202123112123232): complete subsection reference.

<a id="canonical-1233303031123223-2233002012221312-1211033012011330-2003322100002322-1010133222010301-2013003221312300-0312123331322033-2131033231231021"></a>

<a id="canonical-0020200013001202-2120303001312023-3112200300300321-2031110302233330-1333021203222023-2101202101320333-0222311322230001-1132300131120201"></a>

## name property — headers / 133100231303 / 5

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

<a id="canonical-1312213002000313-2332330322133211-3032300300302010-2203023033222310-0301110300210312-0203031120213120-0323333233310331-2213331321030130"></a>

## Next pages — headers / 133100231303 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2300203302330000-1113331120312103-3012203011332233-0231220202132013-0332000211122220-1032323002122312-3121232010322011-3322202120222012)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2000300230213113-2310133032200222-1032320100112111-1121202103203001-1310131122302031-2200001230003320-3132213322103310-1222303302121103)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1010221010113231-2210103210122133-0230121223002302-0113001011213211-1320032211031102-1003011332233332-3312131133133110-0202123112123232)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2300203302330000-1113331120312103-3012203011332233-0231220202132013-0332000211122220-1032323002122312-3121232010322011-3322202120222012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220203301333313-1321032223321301-3131023030100013-0101223033311031-1201111113102333-1330330120312230-1031120021020330-3232213022121120"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present — check_not_present / 313012223011 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1030120301301200-2113032102213202-0320122122232132-1221322101011121-1221132030000123-0220331000221300-2330122322111200-1103202111301111)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_not_present

<a id="canonical-0033232331300122-2300112210123211-0022302132021320-3332033301102230-3103012112232132-0102331332002301-2201332320032212-2210231302002021"></a>

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

<a id="canonical-3113203000302332-2231212030230230-2021020210010012-0021232202211111-3130120021032120-1121333303331321-3103212201000020-1021022322021131"></a>

## Direct properties — check_not_present / 313012223011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322010033030031-3230131211332013-2202300231002111-0112230222223303-3232020201022012-3332321012231103-2300033320123031-3210301202323210"></a>

## Next pages — check_not_present / 313012223011 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1030120301301200-2113032102213202-0320122122232132-1221322101011121-1221132030000123-0220331000221300-2330122322111200-1103202111301111)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2000300230213113-2310133032200222-1032320100112111-1121202103203001-1310131122302031-2200001230003320-3132213322103310-1222303302121103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012233200102013-0300023210320110-0011323000322302-1112120111330022-3123022121001120-1221001233000031-3002221201001133-1033212030300021"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present — check_present / 110032133330 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1030120301301200-2113032102213202-0320122122232132-1221322101011121-1221132030000123-0220331000221300-2330122322111200-1103202111301111)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.check_present

<a id="canonical-0032310121110333-0013011213330013-2302100012330203-2230021011111121-2332032033002020-2033313113213313-0110203010013211-3103321032203021"></a>

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

<a id="canonical-3202212320212032-0100131032233031-2002222233123223-0020020102021120-0120333213001013-2203123000002012-3331213132131001-2002112132023201"></a>

## Direct properties — check_present / 110032133330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221133323130120-3101112203230310-0330222123101031-0030323310000212-0002333023130030-3313221130222220-2202233100003212-3202303121203131"></a>

## Next pages — check_present / 110032133330 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1030120301301200-2113032102213202-0320122122232132-1221322101011121-1221132030000123-0220331000221300-2330122322111200-1103202111301111)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1010221010113231-2210103210122133-0230121223002302-0113001011213211-1320032211031102-1003011332233332-3312131133133110-0202123112123232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002313111110121-1112112100312300-3321012213120100-0023110133222201-2311303132301203-0200100030011220-1223131031112211-0121220000122332"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item — item / 332032211301 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1030120301301200-2113032102213202-0320122122232132-1221322101011121-1221132030000123-0220331000221300-2330122322111200-1103202111301111)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers.item

<a id="canonical-1311003221332203-2003220232312013-3223210020322103-3031220133021211-0122230020003023-0310203301132303-3132323133100312-3122012211000123"></a>

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

<a id="canonical-3021010211130310-1230120332201330-1322313011302003-1201121003233202-3123330321131122-3302312200033130-2303231313011012-3002230101030211"></a>

## Direct properties — item / 332032211301 / 3

<a id="canonical-1323021200231111-1112021002112330-1202121112212011-3113110010231212-0000221233111300-1200330201220012-2033103101333000-0033331201002103"></a>

<a id="canonical-3320131321033330-1000012301032123-1332100323322322-2211210323113300-3211233102021333-3222121133132320-1022100003011331-3233102110202101"></a>

## exact_values property — item / 332032211301 / 4

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

<a id="canonical-2013232022312213-1133022203211113-1332231023032100-1020001131102123-2031302330221321-1121131331020203-1221333013221021-3221121122120012"></a>

<a id="canonical-0333122300033210-2110033102302331-1103221211031233-1201301132321102-1001103333311000-0203122000310330-3112030322130302-3201130112222130"></a>

## regex_values property — item / 332032211301 / 5

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

<a id="canonical-2323212101000101-2333130001001100-1332213221023123-3300330303320112-1233023102103332-2321302321331230-3102131332033230-1210321223132031"></a>

<a id="canonical-3032100233013310-1121301303110212-0121013331101003-0130012213120320-1232210031131211-2322202332230123-1123132202103222-2320010131120321"></a>

## transformers property — item / 332032211301 / 6

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

<a id="canonical-0012133003210300-1103033300312100-0313100312022300-2113002301330000-2023132010002232-1211202302312311-1223232220012001-1011031311103300"></a>

## Next pages — item / 332032211301 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.headers](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1030120301301200-2113032102213202-0320122122232132-1221322101011121-1221132030000123-0220331000221300-2330122322111200-1103202111301111)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2122302200123210-2300130301121320-2331322133200302-0312300323131112-3323113030031332-2233322230123010-3221321203202200-0333211022031121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311230003200211-3233321323101330-2101212301013133-2030022003213112-3013132030312101-1111213203202202-3132022010232311-2031032222332232"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims — jwt_claims / 300032101330 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims

<a id="canonical-3323231002312000-3230323202122303-2201113003033321-2011203322020130-2313222323322102-1133200213011021-0022330031021132-2022130002030022"></a>

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

<a id="canonical-1113201312130011-3020212021133333-2301020331312122-3231011022030222-3122222020303311-2303321030322201-0103320032100023-1110120322311013"></a>

## Direct properties — jwt_claims / 300032101330 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0101331301111133-2211223010200301-1321223220203210-1332103201031321-1012201102102133-3110220001003103-2201130211223200-2133123213002233): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2100123100123021-1121003300111102-0002302310312001-2021003002133103-2321010011320231-0123300032230232-0133221211300032-1031120221003303): complete subsection reference.

<a id="canonical-2101003222332100-0031321112200020-3320001030233220-1302113313222220-0111023102112132-2010213131202022-3133301210023322-3133012303213132"></a>

<a id="canonical-2020230221022311-1012123030132112-1320122112030031-0323032031023022-0101100002223233-0132302003023303-3111130232310020-3021320111023120"></a>

## invert_matcher property — jwt_claims / 300032101330 / 4

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

- [item](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0022301003133103-0302212001200223-0333231303101300-3000012113221001-3202311202030310-0230123023200131-3221021200300203-0331210010313022): complete subsection reference.

<a id="canonical-2322032333200201-0201033031301300-3320002223122201-3332231330231232-1121322100203202-2000102223310012-0323101110031312-0130233222030103"></a>

<a id="canonical-3202231303331233-2230202213121203-3123332102202330-2132311001100210-3223321321112332-3100320032201101-3300211102332001-3122000313212032"></a>

## name property — jwt_claims / 300032101330 / 5

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

<a id="canonical-3203200121031232-0131111131223103-0132230333103102-2120020033131010-3010020020101330-3022310300013103-1110022112023032-0120211003120321"></a>

## Next pages — jwt_claims / 300032101330 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0101331301111133-2211223010200301-1321223220203210-1332103201031321-1012201102102133-3110220001003103-2201130211223200-2133123213002233)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2100123100123021-1121003300111102-0002302310312001-2021003002133103-2321010011320231-0123300032230232-0133221211300032-1031120221003303)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0022301003133103-0302212001200223-0333231303101300-3000012113221001-3202311202030310-0230123023200131-3221021200300203-0331210010313022)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0101331301111133-2211223010200301-1321223220203210-1332103201031321-1012201102102133-3110220001003103-2201130211223200-2133123213002233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333321101310221-0310202000330030-2013013003031022-2010103131312100-3120313212120032-1213112211333003-0132022113011303-1111323223332101"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present — check_not_present / 301130130130 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2122302200123210-2300130301121320-2331322133200302-0312300323131112-3323113030031332-2233322230123010-3221321203202200-0333211022031121)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_not_present

<a id="canonical-3211231002103021-3330213223121221-2113010031100310-3023312120002322-1022321321201123-0132232110223220-2032302130312311-2301203113011221"></a>

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

<a id="canonical-3022012133300120-1313030133321103-3110001202001123-3033103011301102-3033031130032323-3111302113323231-0220200303302210-0200222323222133"></a>

## Direct properties — check_not_present / 301130130130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030333021100333-0311113312320201-3330311321331131-2022312332200301-1013211213103003-2310003231300310-0031111100103101-0223202132331121"></a>

## Next pages — check_not_present / 301130130130 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2122302200123210-2300130301121320-2331322133200302-0312300323131112-3323113030031332-2233322230123010-3221321203202200-0333211022031121)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2100123100123021-1121003300111102-0002302310312001-2021003002133103-2321010011320231-0123300032230232-0133221211300032-1031120221003303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331013122110332-0030111122020033-2201212032120021-2010223132220101-3132132120330230-1120001231312231-1103122203222333-0112000120233110"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present — check_present / 131322002231 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2122302200123210-2300130301121320-2331322133200302-0312300323131112-3323113030031332-2233322230123010-3221321203202200-0333211022031121)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.check_present

<a id="canonical-3233231322312223-2201321222103230-3222112112100200-1233111032313200-2113231100311012-2310102303222022-2231300112001310-1010301121313113"></a>

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

<a id="canonical-1232222130211003-0123211312112022-0221203210123300-1000312102310213-3301303332033003-1102001011130333-1023103322312210-1233300330021012"></a>

## Direct properties — check_present / 131322002231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313320220322313-2231023311322030-0232001312322332-2300323000123011-3301331030333110-1221110113220323-0023013233131221-2103201223300210"></a>

## Next pages — check_present / 131322002231 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2122302200123210-2300130301121320-2331322133200302-0312300323131112-3323113030031332-2233322230123010-3221321203202200-0333211022031121)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0022301003133103-0302212001200223-0333231303101300-3000012113221001-3202311202030310-0230123023200131-3221021200300203-0331210010313022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131121021122030-1102332322003210-3032202031020320-2320031123233223-3323221231033320-1213333313330101-1212020223030023-1123030122312231"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item — item / 000203110310 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2122302200123210-2300130301121320-2331322133200302-0312300323131112-3323113030031332-2233322230123010-3221321203202200-0333211022031121)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims.item

<a id="canonical-0211133122220023-2022213103003030-1332210010202120-1331031321002102-2021332103030201-0213030013122310-0212212012113231-0300333020302033"></a>

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

<a id="canonical-2323322013002112-1130320212012223-1211202101312131-0103001312323010-2322020313111021-3101301211011113-3000331013333221-1323323220132123"></a>

## Direct properties — item / 000203110310 / 3

<a id="canonical-3311332301313122-2231202301203333-1133203231002300-2102020003303300-1122110333120221-1321203103231320-2333212023113320-3231020321011033"></a>

<a id="canonical-2122001212331111-0120302321020033-0221323010010203-0100211033331013-1103010003003300-2223213112300131-0330331102113211-2012321213013010"></a>

## exact_values property — item / 000203110310 / 4

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

<a id="canonical-3223210223313001-0000212122102102-0213213132331003-1231020012302122-0222113211123202-3130113110100030-3313002112002121-1130130321133101"></a>

<a id="canonical-2213102132023233-2030033033320330-1232321313220102-3200003021132021-1203323200231220-2010331203113111-3131303331303113-2312213133212233"></a>

## regex_values property — item / 000203110310 / 5

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

<a id="canonical-0212010111122313-1233023211233202-0132331320130311-1323201223010100-3221321200203221-0122311331000322-0002312023012301-3221132022201321"></a>

<a id="canonical-0301333213332002-3323312020300010-2211120221110033-1013321000203013-1202323211033121-3331010202010202-2201312301002112-1222030133320122"></a>

## transformers property — item / 000203110310 / 6

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

<a id="canonical-1300311233021230-0103221202011103-2023210120103322-0130101032120200-0312120330102211-2332112032122321-1221132031231323-2120002001201031"></a>

## Next pages — item / 000203110310 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.jwt_claims](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2122302200123210-2300130301121320-2331322133200302-0312300323131112-3323113030031332-2233322230123010-3221321203202200-0333211022031121)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3132310103201013-2010300221013120-3202331131232333-0032003011032022-3331302103311103-3022013200332312-3001021021131212-2011100333232220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110021323003230-3023022112332121-1011120221013322-0301231023211313-1112223133121033-0223233301121120-3313203322313022-1300012101101012"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params — query_params / 220202111222 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params

<a id="canonical-3031133220321110-1000121122002031-1212210332313202-2232211212230210-0311011133222332-0031022332322222-3030202223131202-1121101111232232"></a>

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

<a id="canonical-2021111230330202-2103110110011220-2201222320132002-0023223121022003-1213023012211311-3113333301103013-1212232331302212-2231220332231013"></a>

## Direct properties — query_params / 220202111222 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1320011110310313-1320210001131301-3300232211101031-3223111133202120-1221023313222120-2220211003102312-0332031233021320-0002000012132203): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1001312323012320-0021230103230122-1011100033123200-1211332213132112-0022112331233202-2322010222123030-0001201312013211-2220031301021222): complete subsection reference.

<a id="canonical-2000323012011030-1321200133122300-0111013003332133-3220030132230220-1130200211302010-0113012021232100-1202330223211212-1231320230101222"></a>

<a id="canonical-0231010312100123-2132333202033232-3230312000303312-2031323230110131-1221131221000020-1202110200013220-3330122000211312-2223203012130100"></a>

## invert_matcher property — query_params / 220202111222 / 4

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

- [item](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2032111031231023-3323001331310000-0030332112320021-0120212231100033-0033032000130020-3203213300111123-1300033023231223-1221233000323130): complete subsection reference.

<a id="canonical-1302013032302213-1320213311100003-1313312113103020-3223133003300123-3023221130010113-3221010320222121-0002333303320230-0202211202302111"></a>

<a id="canonical-1000300323332001-3300002033331301-1200200100203123-1001111311333223-3230111130011130-0030300232231310-3003200012010002-1201322330023011"></a>

## key property — query_params / 220202111222 / 5

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

<a id="canonical-2323022323312232-3112122333123331-3302330213303323-1233003210011302-1001221030310112-0332031123101310-0302020031132212-2111202300102203"></a>

## Next pages — query_params / 220202111222 / 6

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1320011110310313-1320210001131301-3300232211101031-3223111133202120-1221023313222120-2220211003102312-0332031233021320-0002000012132203)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1001312323012320-0021230103230122-1011100033123200-1211332213132112-0022112331233202-2322010222123030-0001201312013211-2220031301021222)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2032111031231023-3323001331310000-0030332112320021-0120212231100033-0033032000130020-3203213300111123-1300033023231223-1221233000323130)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1320011110310313-1320210001131301-3300232211101031-3223111133202120-1221023313222120-2220211003102312-0332031233021320-0002000012132203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002221130030201-2033232211030333-1131020112020222-1121303031213230-3223230030300311-1213201101012313-0102300123202321-2102232303100320"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present — check_not_present / 200011103320 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3132310103201013-2010300221013120-3202331131232333-0032003011032022-3331302103311103-3022013200332312-3001021021131212-2011100333232220)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_not_present

<a id="canonical-1131011122020001-2332223030311222-3213233201231013-1213131221121203-3012320010132221-0020212010101213-1332011222010030-1022133112032310"></a>

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

<a id="canonical-1011002333232112-2032333123120100-1332111102320211-2320322013222301-1301002312311301-1230313230321203-1222032313310321-3003023000310012"></a>

## Direct properties — check_not_present / 200011103320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003313031231131-1133301211133222-2012233011222123-3013313130130200-2130303201331230-1031322331233131-2021300212213031-3201013101230020"></a>

## Next pages — check_not_present / 200011103320 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3132310103201013-2010300221013120-3202331131232333-0032003011032022-3331302103311103-3022013200332312-3001021021131212-2011100333232220)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1001312323012320-0021230103230122-1011100033123200-1211332213132112-0022112331233202-2322010222123030-0001201312013211-2220031301021222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131311100011203-0111301202002112-2233131230300222-3023232232220001-0211222120202212-0120123223121302-2011123221302330-0131012202023111"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present — check_present / 133322322321 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3132310103201013-2010300221013120-3202331131232333-0032003011032022-3331302103311103-3022013200332312-3001021021131212-2011100333232220)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.check_present

<a id="canonical-3331102021310021-2021201000103021-1223020233313332-2232001232132131-0313233223300213-3211132233132220-0221211231200012-2132213023020131"></a>

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

<a id="canonical-3221302013312210-0110132021003210-1330320311113203-1111102310310222-2213232200331110-2310032231333211-0112321020231331-0300322213121032"></a>

## Direct properties — check_present / 133322322321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231300322230122-3320103203133021-2231220022223102-3203011320212322-3011333221203230-2200123212330031-2111223313031301-0310331321201320"></a>

## Next pages — check_present / 133322322321 / 4

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3132310103201013-2010300221013120-3202331131232333-0032003011032022-3331302103311103-3022013200332312-3001021021131212-2011100333232220)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2032111031231023-3323001331310000-0030332112320021-0120212231100033-0033032000130020-3203213300111123-1300033023231223-1221233000323130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200101033200030-2033223003200230-1131220213102130-3321310302030032-0102001021313302-2120312322011212-0012013210331112-2133011310013001"></a>

## api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item — item / 102220310213 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0100023232303101-1013100112321303-1010322303132332-1312231013101231-1110330320213101-2031213320022030-0332023111132012-0031323303211123)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1011220030033133-0223111231130332-2201221331001300-0001331033133001-3012313023111310-3003121310300211-0010132030131200-0302301122020032)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0202100101231000-3230101220203330-0003112213211020-2001132032033111-3120201331321120-3110100030022321-2003030213230122-0303002320322000)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3132310103201013-2010300221013120-3202331131232333-0032003011032022-3331302103311103-3022013200332312-3001021021131212-2011100333232220)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params.item

<a id="canonical-1103113333002122-0033131231303131-2013033213101021-2212211012001300-0133302012233300-2131132313312332-1311320232330113-3233031233221023"></a>

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

<a id="canonical-1032110332022210-2032200032201331-2112202101111130-2132102031103110-2012321203322222-2212010012230300-1310331302321313-3103202012100200"></a>

## Direct properties — item / 102220310213 / 3

<a id="canonical-2202332320121033-1032211110330233-1020011132230001-3120301223000122-1023003231203203-3112113103003232-1223321022123100-3033223200231022"></a>

<a id="canonical-1101230000100013-0322232301310122-0113303230113321-2313003211212100-0010020213132222-2003030000033330-0230013233121122-0202332101111312"></a>

## exact_values property — item / 102220310213 / 4

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

<a id="canonical-2000011100033310-3201032133033201-1321010003330220-3233312101113121-2332303102120132-1221000301203213-3200322302031323-0101302021032031"></a>

<a id="canonical-0321011030033212-1021002233333311-2021202313023213-3101013110102230-1303201010023002-2222120131110020-0222322212032223-3223131331021221"></a>

## regex_values property — item / 102220310213 / 5

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

<a id="canonical-1300133213310201-0221203322300210-1000301200130013-3002231001101221-2323201231013323-0200133113230021-3013032010030310-3132302101323220"></a>

<a id="canonical-3222032103003202-2101310223033003-2230103220311231-2200311330313032-1313101201102121-1203122032010332-1320103013133031-3102210020001313"></a>

## transformers property — item / 102220310213 / 6

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

<a id="canonical-1000111321331333-3101121323332022-2330230100210222-2321132000312113-2321200231311220-1112323122200222-3123033021122302-2102022212202221"></a>

## Next pages — item / 102220310213 / 7

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.request_matcher.query_params](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3132310103201013-2010300221013120-3202331131232333-0032003011032022-3331302103311103-3022013200332312-3001021021131212-2011100333232220)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1010130033000023-0003001221210201-0012121132223103-0120030303211233-1033132223023030-1321021102203331-2100133012021110-2303033132131131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003332131330322-1331120120110312-1022233130020232-2023013213203122-3122100313001320-2023301300120103-2221330100020100-1300203322110131"></a>

## api_rate_limit.custom_ip_allowed_list — custom_ip_allowed_list / 000103310330 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- api_rate_limit.custom_ip_allowed_list

<a id="canonical-3203200313103020-3121213313310212-3000230120202011-1230112232000303-0310233221120122-2210100231320222-0202311113121122-2132332023101132"></a>

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

<a id="canonical-2210223010200001-1132121303000000-2220003330131013-3313320213230323-2323212212211102-1112031332222123-1021031312303020-3000201323012222"></a>

## Direct properties — custom_ip_allowed_list / 000103310330 / 3

- [rate_limiter_allowed_prefixes](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2313001303220112-0111323121220310-0310320000120011-0220012232331323-2100303120130213-1201010102123002-0323100303001333-0310210030330301): complete subsection reference.

<a id="canonical-2321003231202123-3100203112022202-3002220322000131-2013331130333330-0231132122333130-1131313301000012-3331121320233013-2022302203310230"></a>

## Next pages — custom_ip_allowed_list / 000103310330 / 4

- [api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2313001303220112-0111323121220310-0310320000120011-0220012232331323-2100303120130213-1201010102123002-0323100303001333-0310210030330301)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2313001303220112-0111323121220310-0310320000120011-0220012232331323-2100303120130213-1201010102123002-0323100303001333-0310210030330301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013033013111213-3212211322021201-1200201223210333-2320020023212332-3213210132232020-2210331211103210-3233323333330211-3321322122213022"></a>

## api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes — rate_limiter_allowed_prefixes / 120031301131 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.custom_ip_allowed_list](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1010130033000023-0003001221210201-0012121132223103-0120030303211233-1033132223023030-1321021102203331-2100133012021110-2303033132131131)
- api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes

<a id="canonical-2113320301310110-3231010300201120-0032211333111301-0221220333200101-1233020330211303-1301212333303212-1123213201032011-1221023313300221"></a>

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

<a id="canonical-3131103113202113-0222201211031300-3010031132032100-1021330020331211-0113212022322303-2203133301222321-2212211311101203-1230301003330200"></a>

## Direct properties — rate_limiter_allowed_prefixes / 120031301131 / 3

<a id="canonical-0023330331323203-2320023113002103-0310223330322022-0132100311001110-3333220202322332-1312103010121210-3021002100130012-0031020101330020"></a>

<a id="canonical-2103300230013203-1123110032100311-2122230021320303-2202133100112230-2120113013222012-1331022113221030-3212012303231000-3322222200100312"></a>

## name property — rate_limiter_allowed_prefixes / 120031301131 / 4

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

<a id="canonical-3031013110031200-0001000331120131-3233122221000132-0111333110330102-3023202032232301-1213222313311312-1301200110331300-0323311030000022"></a>

<a id="canonical-1032201211221221-2221203131100231-1331221320312212-0211002331131022-3000122022030010-1333200230032221-2000002311020020-2322130300220320"></a>

## namespace property — rate_limiter_allowed_prefixes / 120031301131 / 5

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

<a id="canonical-1300101201230302-0011032222101122-1002332000233213-3220101333213110-1233300010030310-2113013021112200-2102123012311103-1233013233212023"></a>

<a id="canonical-2222223112132200-3303032030020301-3312213312121331-2131121200203111-3312222030202331-3021331221322010-3210323130330002-1032310020310220"></a>

## tenant property — rate_limiter_allowed_prefixes / 120031301131 / 6

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

<a id="canonical-2201102332301131-0101231133022121-3213133003311020-1210302113023203-2211212300020233-0103001331222012-1213201121003122-1103111220003331"></a>

## Next pages — rate_limiter_allowed_prefixes / 120031301131 / 7

- [api_rate_limit.custom_ip_allowed_list](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-1010130033000023-0003001221210201-0012121132223103-0120030303211233-1033132223023030-1321021102203331-2100133012021110-2303033132131131)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2303033302312111-0231012102313300-0030302012321311-0331312000002132-3021100202110103-2120133133300232-0230222110332003-1102203220130300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220222010122323-0333202130303321-1022323021032201-3132323113101111-1103131113200013-1102333212210020-0011101003333111-0121323021101233"></a>

## api_rate_limit.ip_allowed_list — ip_allowed_list / 233032332102 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- api_rate_limit.ip_allowed_list

<a id="canonical-0012001300022000-2130001230313211-2221323102232232-3100103032303302-1302032023330213-3233010022022320-2312023101300013-3001220003103223"></a>

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

<a id="canonical-3033133333010102-0031200201023133-0021120211013300-0003122310030130-3321101021001021-1113313303123210-2111013223303002-1003023113120000"></a>

## Direct properties — ip_allowed_list / 233032332102 / 3

<a id="canonical-2103220021203201-3223222011013010-2012232212022023-1311133030132321-0333113013221322-1203133012123213-1003233310301121-1233312131303112"></a>

<a id="canonical-1323033000210023-0202303113212311-0203110201110033-3011103000201012-0132030032021011-0130111232233203-2301311332112320-1133122302221201"></a>

## prefixes property — ip_allowed_list / 233032332102 / 4

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

<a id="canonical-0233211321021322-3101333233003121-0011333012310313-3130111023031123-1330012002323011-0203233210030220-0100211201122011-1332013032311322"></a>

## Next pages — ip_allowed_list / 233032332102 / 5

- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0231102032303012-3001232002310032-0303213312332003-1321010302001031-0232222203033013-0102231202310231-2103133110112032-0131120102211320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321031031101312-3131032330212112-2233103030313231-2010000012213001-3110113310123232-1132201111322330-3310331002100311-2000020120131231"></a>

## api_rate_limit.no_ip_allowed_list — no_ip_allowed_list / 023120322201 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- api_rate_limit.no_ip_allowed_list

<a id="canonical-3122331200202103-1200222131323233-3310310030110112-3020000131210012-3300223210233110-1011232103103033-3213102211303303-3220000312300200"></a>

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

<a id="canonical-3213231003330232-3331313312010200-0301102323100220-0033230023132110-3023001331221110-0131231321102301-3131311012223021-1212032303003233"></a>

## Direct properties — no_ip_allowed_list / 023120322201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113203031331012-1232230122022110-3310231330113131-2021033301112012-0232210213122312-0222000201013033-3011312001133303-1032320010103110"></a>

## Next pages — no_ip_allowed_list / 023120322201 / 4

- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000020213111220-1132112310312231-2223322222000202-1132013203132303-2102301310220010-1121011001111303-2213002002012232-3321302203210232"></a>

## api_rate_limit.server_url_rules — server_url_rules / 331303300102 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- api_rate_limit.server_url_rules

<a id="canonical-3320322303200332-3123220122301100-3321100303231010-1300133230333231-0222222022110002-0203301103020131-3331211322021012-1231102222110123"></a>

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

<a id="canonical-0023133200323320-2123321000201113-0312120113103313-2131022231002111-2231003300231030-0302110121130002-1302212322322231-0032320221331201"></a>

## Direct properties — server_url_rules / 331303300102 / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0300301130330202-0121200322320021-2203230032110133-1203021000102121-3232212110321103-1112100323312302-2302203002331130-0301101333130320): complete subsection reference.

<a id="canonical-0031311333032223-2102122201202111-1332210000000121-2211311121121202-0130232233313301-0110131202033032-1311130312121330-1021330230003213"></a>

<a id="canonical-2100011202010303-0203121103030312-0011323130223311-2213112012321330-1101023301303302-3322103031333111-3101331232010203-3110223132000301"></a>

## api_group property — server_url_rules / 331303300102 / 4

Type: `"string"`. Computed.

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with 'x-F5 Distributed..

Upstream description:

API groups derived from API Definition swaggers. For example oas-all-operations including all paths
and methods from the swaggers, oas-base-URLs covering all requests under base-paths from the
swaggers. Custom groups can be created if user tags paths or operations with "x-F5 Distributed
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

<a id="canonical-2000102231332121-1311221013212031-0011323303230113-1201233332330311-1231301312101330-1121313002000331-3031333130220211-2230013223031320"></a>

<a id="canonical-3020300130301330-2321133331212000-3003012122010323-1112022312100333-0313333132300120-1213333103131031-3210012331203203-2310303130023223"></a>

## base_path property — server_url_rules / 331303300102 / 5

Type: `"string"`. Computed.

Base Path. Prefix of the request path.

Upstream description:

Prefix of the request path.

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

- [client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022): complete subsection reference.

- [inline_rate_limiter](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3020203131201121-0321103000110021-0000001321301203-1220123001023020-3230213101001103-0331101023220030-3223330201312212-3232321223223320): complete subsection reference.

- [ref_rate_limiter](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1121003122321003-1030113023222001-1120323002113122-1103033302313332-0220011200302103-2120112030110322-1320212132033201-2023201113201010): complete subsection reference.

- [request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303): complete subsection reference.

<a id="canonical-2110202032232313-2103023100111131-0322013102221130-3110322123012222-1001200013032203-0022010211100023-1120321101333120-3203332220103103"></a>

<a id="canonical-1230321211323231-3120302331211110-0221001331123012-2322001110223321-3303131013302313-0131200032200311-1102312000230131-1131312333030220"></a>

## specific_domain property — server_url_rules / 331303300102 / 6

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

<a id="canonical-2213230333013221-2102001013122013-2300313101012220-1030232332021220-3230120002312213-3023100001302211-3002232313023302-0132332313212012"></a>

## Next pages — server_url_rules / 331303300102 / 7

- [api_rate_limit.server_url_rules.any_domain](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0300301130330202-0121200322320021-2203230032110133-1203021000102121-3232212110321103-1112100323312302-2302203002331130-0301101333130320)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- [api_rate_limit.server_url_rules.inline_rate_limiter](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3020203131201121-0321103000110021-0000001321301203-1220123001023020-3230213101001103-0331101023220030-3223330201312212-3232321223223320)
- [api_rate_limit.server_url_rules.ref_rate_limiter](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1121003122321003-1030113023222001-1120323002113122-1103033302313332-0220011200302103-2120112030110322-1320212132033201-2023201113201010)
- [api_rate_limit.server_url_rules.request_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3331120332232001-2230203100310132-0131103131322321-2203311031000322-0032002310010331-1232313002103322-1022322111212203-3123023100100303)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0300301130330202-0121200322320021-2203230032110133-1203021000102121-3232212110321103-1112100323312302-2302203002331130-0301101333130320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211220120002120-1110331122330122-0322032203221011-0212111010302222-0301121102303211-3221122011021032-2000012103133032-2320130011302110"></a>

## api_rate_limit.server_url_rules.any_domain — any_domain / 002203112122 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- api_rate_limit.server_url_rules.any_domain

<a id="canonical-2323131230113001-2202010331013111-1110011100033121-3322022113132020-2003222301111323-2312220130120020-3321132210303311-1221312113130322"></a>

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

<a id="canonical-0101022102113311-2133231231201301-0033013101113030-2311121231012003-2220201010123121-2021302012113211-2023312101202232-1230112000223032"></a>

## Direct properties — any_domain / 002203112122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130230311313201-2133203221311021-3111001312202102-1333132022311311-0023133111320131-3011201223013110-2313132232020320-0210323020320123"></a>

## Next pages — any_domain / 002203112122 / 4

- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213121330302002-0230030223210332-2331311020203301-2221202311301131-2222130212012300-1233021123013221-3211221132232231-1202111300323301"></a>

## api_rate_limit.server_url_rules.client_matcher — client_matcher / 022331031232 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- api_rate_limit.server_url_rules.client_matcher

<a id="canonical-0210221010103002-2210223002321202-1112210121213213-2031231330131202-0201133332312213-3201113322130130-3220000213211132-3321202102233222"></a>

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

<a id="canonical-3300231223111013-1220111303130213-3232212130302011-2330221210100133-1211033203303102-2310132000130213-0132020231012232-3020223320330130"></a>

## Direct properties — client_matcher / 022331031232 / 3

- [any_client](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3100333201213100-0233210101312033-0013330320301030-1032013322202320-2221113000023020-2003231333113203-0103321211101312-2012020331330101): complete subsection reference.

- [any_ip](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2123231201203203-0210221320012000-1301302130201233-2112110130230231-0323131223030222-0231313311211331-3123323031110031-2121213111032222): complete subsection reference.

- [asn_list](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0212313230132222-1100122132011210-2331330302223220-0311333300303323-3221111020031331-1033002130113000-1133100323232332-1312122232023323): complete subsection reference.

- [asn_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0133022212232102-1112100110032100-0132032331300111-0012313212321133-2331322300300022-0030113112200233-2020233032030313-1100320203220131): complete subsection reference.

- [client_selector](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1102311331212020-3211200320200222-2332222313232022-1303200130323321-0201321202321212-0130211201023332-0132232331302303-2110212200321220): complete subsection reference.

- [ip_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2232310130330032-0230113113110122-3311201323331201-1123310322203310-3122013132022210-0203031021101102-0303110220332330-1131012111133130): complete subsection reference.

- [ip_prefix_list](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1111210112232110-3133000020031222-3110321003330330-2102221330302033-1120133301333103-1120330011032222-1323100133211013-1203121330010323): complete subsection reference.

- [ip_threat_category_list](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3122022200133313-0102331311012322-3021212000002223-2000203333200202-1322331230332033-2221130003120212-3032002303233211-2303213201111032): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1331332223312102-1132112313212300-1023223213002222-2333213112000033-3131223010112201-3111213233133211-0202302112202003-2021131013002322): complete subsection reference.

<a id="canonical-0013113010220332-1310013133110133-0202001132102020-3111220033112021-2030231333220210-1331102003212120-3210223210222031-1010001131321012"></a>

## Next pages — client_matcher / 022331031232 / 4

- [api_rate_limit.server_url_rules.client_matcher.any_client](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-3100333201213100-0233210101312033-0013330320301030-1032013322202320-2221113000023020-2003231333113203-0103321211101312-2012020331330101)
- [api_rate_limit.server_url_rules.client_matcher.any_ip](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2123231201203203-0210221320012000-1301302130201233-2112110130230231-0323131223030222-0231313311211331-3123323031110031-2121213111032222)
- [api_rate_limit.server_url_rules.client_matcher.asn_list](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0212313230132222-1100122132011210-2331330302223220-0311333300303323-3221111020031331-1033002130113000-1133100323232332-1312122232023323)
- [api_rate_limit.server_url_rules.client_matcher.asn_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-0133022212232102-1112100110032100-0132032331300111-0012313212321133-2331322300300022-0030113112200233-2020233032030313-1100320203220131)
- [api_rate_limit.server_url_rules.client_matcher.client_selector](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1102311331212020-3211200320200222-2332222313232022-1303200130323321-0201321202321212-0130211201023332-0132232331302303-2110212200321220)
- [api_rate_limit.server_url_rules.client_matcher.ip_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-2232310130330032-0230113113110122-3311201323331201-1123310322203310-3122013132022210-0203031021101102-0303110220332330-1131012111133130)
- [api_rate_limit.server_url_rules.client_matcher.ip_prefix_list](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1111210112232110-3133000020031222-3110321003330330-2102221330302033-1120133301333103-1120330011032222-1323100133211013-1203121330010323)
- [api_rate_limit.server_url_rules.client_matcher.ip_threat_category_list](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-3122022200133313-0102331311012322-3021212000002223-2000203333200202-1322331230332033-2221130003120212-3032002303233211-2303213201111032)
- [api_rate_limit.server_url_rules.client_matcher.tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-005.md#canonical-1331332223312102-1132112313212300-1023223213002222-2333213112000033-3131223010112201-3111213233133211-0202302112202003-2021131013002322)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3100333201213100-0233210101312033-0013330320301030-1032013322202320-2221113000023020-2003231333113203-0103321211101312-2012020331330101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322133313320011-0333220001321032-1210113302010302-1211132101310200-2301202202012030-2321111010110213-0010213201201030-3332023120201002"></a>

## api_rate_limit.server_url_rules.client_matcher.any_client — any_client / 023210223231 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.any_client

<a id="canonical-0011230331302121-2302111030123220-2010100332311213-0012100330331120-1231212303201331-1303210223100113-1003000021301312-0232323112100130"></a>

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

<a id="canonical-2331230212003211-2100103221013001-1302302122312303-1221033122032201-3210003231212321-0003322323110313-3230230331332333-2232300021321110"></a>

## Direct properties — any_client / 023210223231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320000112021012-2110100333203022-1322222020022123-2210110122213313-1121103010133102-1023323121123000-1122201120032023-2320232030233000"></a>

## Next pages — any_client / 023210223231 / 4

- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2123231201203203-0210221320012000-1301302130201233-2112110130230231-0323131223030222-0231313311211331-3123323031110031-2121213111032222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321303030320311-0232023331110120-3312200133123012-0322202231303231-1033223330020223-3120020233322030-1011000110010120-0100322100211223"></a>

## api_rate_limit.server_url_rules.client_matcher.any_ip — any_ip / 311303010022 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023)
- [api_rate_limit.server_url_rules](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-2331201322132003-3203102221020331-2302223231102210-0032223031203321-3200020211103200-1013212320332121-1000222312132322-3232212123303011)
- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- api_rate_limit.server_url_rules.client_matcher.any_ip

<a id="canonical-0302120331313321-3012112201120331-3223333112121331-2302022013322013-0210111200002000-2011221123112030-0213002323000323-1032031331303011"></a>

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

<a id="canonical-0113121100323000-1021221033203333-2213031012313312-3323323021212011-0322332110230110-0023213111022331-3131103302112321-2030312021213213"></a>

## Direct properties — any_ip / 311303010022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302111101112223-0232001231233103-1002000320233213-3231031223312322-1131311133133101-0130012122222120-0100023230002033-0302201231222201"></a>

## Next pages — any_ip / 311303010022 / 4

- [api_rate_limit.server_url_rules.client_matcher](data-sources--cdn_loadbalancer--reference--group-004.md#canonical-0110130123300101-0033103311332111-2333002321111332-0330032201023332-1220001232132302-3233022300003101-1002101202212312-0131313230230022)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0212313230132222-1100122132011210-2331330302223220-0311333300303323-3221111020031331-1033002130113000-1133100323232332-1312122232023323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
