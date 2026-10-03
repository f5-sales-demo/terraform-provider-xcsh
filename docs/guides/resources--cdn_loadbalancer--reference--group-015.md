---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-3000003130123202-0210030102101111-2012032031002233-1132321132321121-0332221132213320-2011300200012031-1132130020213011-0213331233210011"></a>

## Next pages — disable_request_timeout / 332001303003 / 4

- [slow_ddos_mitigation](resources--cdn_loadbalancer--reference--group-014.md#canonical-1233003020202211-1121000220031121-0002002321202011-3200220102103002-3312211001303013-1111001323333210-3232211013001110-3302023302323112)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1302302300220120-1232133331030022-2202232123331001-1001230212313333-2002032030033031-0300122120322121-0023222330210101-0203103301320000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103112130200010-3302223203323302-2001100132013320-3202113211211132-3223030321031232-0302233331210310-2022111131003230-3120320102221202"></a>

## system_default_timeouts — system_default_timeouts / 012310131202 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- system_default_timeouts

<a id="canonical-2003220332131013-2211333310231313-0030002201303021-1102200303220023-3011112311203322-2110130310322231-2112032032222113-2133222202031012"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
system_default_timeouts = {}
```

<a id="canonical-1331213013031321-1231200121102122-2002033010022022-1212011121033113-1102201202210330-3122031312300322-2323201012113110-0110322223302033"></a>

## Direct properties — system_default_timeouts / 012310131202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102121132002223-2233030100122132-1003002321213113-1011121312310202-1301020021130113-1300333011301123-3201210033021103-1332302022011003"></a>

## Next pages — system_default_timeouts / 012310131202 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1120320111222111-0221312232103321-2100002103330133-0121321020113111-3112222032011030-0332323220222220-0312012223333230-3320123022202133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033333132112201-0003122320323232-1321112013033002-1131120120022002-1231101321303211-2020022230131023-1313131031121320-1120032220201212"></a>

## timeouts — timeouts / 012231100002 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- timeouts

<a id="canonical-2023222110021223-0210202233233101-1321110100223020-1120313011212123-0223220120201112-1013311101313021-3002233013330332-1022010010310301"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221222112021312-3120021032330023-0001223200202103-2311012020121133-1312202302030321-3333122232310013-1121002201232132-1130012212120111"></a>

## Direct properties — timeouts / 012231100002 / 3

<a id="canonical-1202010300212021-0220220031031001-1110010200022332-0111212020123101-1310200302213232-0313130023130313-3221133313213320-1002323130220123"></a>

<a id="canonical-3231111223100301-0213013210231131-0001303323013220-1213103130120332-1011332103233202-3020133323012222-1220223320022131-2113213233131103"></a>

## create property — timeouts / 012231100002 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3200032132201311-2302121230021222-0222123300131002-1021013022311310-2020223031212312-0220310230122023-1230230123323011-2002301233101322"></a>

<a id="canonical-2100200303120130-2200211101003321-3102312331001303-2031200233021132-1311322301231333-3312233220232102-1103003311123220-3332310000211330"></a>

## delete property — timeouts / 012231100002 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3233220020230300-2103131210212301-2313232001223301-2133333100111002-3032111213101303-2011130211103030-2013311320311320-3221112002301030"></a>

<a id="canonical-0013230203310130-2333031301320303-2221101322121320-1012302033032330-3300311331123221-1321132102012331-3233330231011222-1323321102302330"></a>

## read property — timeouts / 012231100002 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0021130100120101-2302132113100113-0330220313110200-0100231332332113-1223103101230003-0310121233210103-2112221122132211-3230223131300210"></a>

<a id="canonical-3220120311322211-2203030330330223-3203033231313231-0200121013212231-0121202200021300-0223303332131030-3112302132332131-0222022203332021"></a>

## update property — timeouts / 012231100002 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1303331310230301-1202300310230032-1302121330213022-2103112332003023-1030023120233201-3131213322330323-0031202332301201-1120220101321232"></a>

## Next pages — timeouts / 012231100002 / 8

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203321333110330-3120020012011011-2121301302103230-2010333311132133-1023331230130203-3100202323010002-3222002300130313-3303321232303223"></a>

## trusted_clients — trusted_clients / 132030021002 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- trusted_clients

<a id="canonical-0331013110121213-1103203320113023-3133220211231201-2333202101220103-3100110012031010-3020111011320232-0113200202211221-1322232320320102"></a>

Type: `"object"`. list nested block, Optional.

Define rules to skip processing of one or more features such as WAF, Bot Defense etc.

Upstream description:

Define rules to skip processing of one or more features such as WAF, Bot Defense etc. For clients.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("actions"),
  validators.ConflictingListObjectAttributes("as_number",
    "http_header"),
  validators.ConflictingListObjectAttributes("as_number",
    "ip_prefix"),
  validators.ConflictingListObjectAttributes("as_number",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("as_number",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("bot_skip_processing",
    "skip_processing"),
  validators.ConflictingListObjectAttributes("bot_skip_processing",
    "waf_skip_processing"),
  validators.ConflictingListObjectAttributes("http_header",
    "ip_prefix"),
  validators.ConflictingListObjectAttributes("http_header",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("http_header",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("ip_prefix",
    "ipv6_prefix"),
  validators.ConflictingListObjectAttributes("ip_prefix",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("ipv6_prefix",
    "user_identifier"),
  validators.ConflictingListObjectAttributes("skip_processing",
    "waf_skip_processing")}
```

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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
trusted_clients {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000212210003103-1100310233303223-2110112002302232-3110212312032332-0200310002232003-2200201333332013-0003110131233322-0110031222233010"></a>

## Direct properties — trusted_clients / 132030021002 / 3

<a id="canonical-3112200000232130-3003122100223203-0113201130223100-0300010022020000-0211011120021011-3132100123303120-1110123122012101-0322020210232312"></a>

<a id="canonical-3320002023201121-3203021322233130-3223220321112020-1122233102332030-3110003313323030-1120110101002102-3222010030001033-1103323332333232"></a>

## actions property — trusted_clients / 132030021002 / 4

Type: `["list", "string"]`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(10),
}
```

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

<a id="canonical-0212200203132213-3001331122003020-1100120311121011-2011223102122132-0302001023333313-2220000122122020-2000113213010333-0111101213032303"></a>

<a id="canonical-2331111130030312-2012131022103332-3102211103001212-3112320302121031-0002030112322220-1003133322102221-1330202310212011-3131302320033221"></a>

## as_number property — trusted_clients / 132030021002 / 5

Type: `"number"`. Optional.

Exclusive with \[http\_header ip\_prefix IPv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Upstream description:

Exclusive with \[http\_header ip\_prefix IPv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 401308),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [bot_skip_processing](resources--cdn_loadbalancer--reference--group-015.md#canonical-2023122312000030-3003311220212320-2301130321111110-0030022112103323-1322310121101221-1022311201233033-2333330301223133-0212311012200001): complete subsection reference.

<a id="canonical-1013332032303321-1223021223132000-1230130322032131-0222321312223220-2231332113200131-2323203212300202-1222122120010023-0202012331320020"></a>

<a id="canonical-0313130033231100-2011223101202121-2001222023133301-1231302031010020-2212101310113230-3303301331233211-3302111133122203-3102101322002202"></a>

## expiration_timestamp property — trusted_clients / 132030021002 / 6

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [http_header](resources--cdn_loadbalancer--reference--group-015.md#canonical-1201221302301222-1010320030333121-0113223131202312-2330113130111003-3012011333202310-0101333000203222-3111003100213100-3013310003133311): complete subsection reference.

<a id="canonical-3333000311032033-0131020213223221-0120323022130202-2131131301110212-2111122020132003-2013321321120123-0120320033112233-1312131221102111"></a>

<a id="canonical-0230323302113323-0102232133212033-0101201022210233-0301020301011012-0101201130333100-3310113222302130-0013233102011320-3003213220123312"></a>

## ip_prefix property — trusted_clients / 132030021002 / 7

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-0300100301021032-3101203201231302-1132102303301201-1133111021223302-2113233133033310-0320112302023202-3103102013322100-2201100000231210"></a>

<a id="canonical-0113200213021102-2133312002221203-1323332120223303-2323021033101322-0131221100300321-1320032321023113-0221101120200020-0131321233333313"></a>

## ipv6_prefix property — trusted_clients / 132030021002 / 8

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

- [metadata](resources--cdn_loadbalancer--reference--group-015.md#canonical-3321333031300023-0102112101330300-2231023022110222-0003023011013002-0222123013100321-2220230222132333-1231133330031312-2133323113312100): complete subsection reference.

- [skip_processing](resources--cdn_loadbalancer--reference--group-015.md#canonical-1133331223201300-3010010301001020-1022001210300133-3110022001203122-3310333101330031-3301103022232211-3311232311220231-2323022320130113): complete subsection reference.

<a id="canonical-2123202023131030-3010311112220303-3330112013322203-1323302030032030-1331230311012013-1310113121033130-0311010031033211-0030320210023021"></a>

<a id="canonical-1122022213221120-1132211202110220-3310202011203113-1232322022232131-3213332323210302-1121323232303203-1113031012000001-0231323012122320"></a>

## user_identifier property — trusted_clients / 132030021002 / 9

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix IPv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

Upstream description:

Exclusive with \[as\_number http\_header ip\_prefix IPv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [waf_skip_processing](resources--cdn_loadbalancer--reference--group-015.md#canonical-0033212232020333-0132213332131031-2012121203120303-1333000331032120-1202131120001010-1123301101201001-2130112022103101-1312012221031000): complete subsection reference.

<a id="canonical-0000330331222330-0023311111012123-3201110022202320-2130213322031022-2032020311220301-2102320013002131-2322132310310221-3121300121132221"></a>

## Next pages — trusted_clients / 132030021002 / 10

- [trusted_clients.bot_skip_processing](resources--cdn_loadbalancer--reference--group-015.md#canonical-2023122312000030-3003311220212320-2301130321111110-0030022112103323-1322310121101221-1022311201233033-2333330301223133-0212311012200001)
- [trusted_clients.http_header](resources--cdn_loadbalancer--reference--group-015.md#canonical-1201221302301222-1010320030333121-0113223131202312-2330113130111003-3012011333202310-0101333000203222-3111003100213100-3013310003133311)
- [trusted_clients.metadata](resources--cdn_loadbalancer--reference--group-015.md#canonical-3321333031300023-0102112101330300-2231023022110222-0003023011013002-0222123013100321-2220230222132333-1231133330031312-2133323113312100)
- [trusted_clients.skip_processing](resources--cdn_loadbalancer--reference--group-015.md#canonical-1133331223201300-3010010301001020-1022001210300133-3110022001203122-3310333101330031-3301103022232211-3311232311220231-2323022320130113)
- [trusted_clients.waf_skip_processing](resources--cdn_loadbalancer--reference--group-015.md#canonical-0033212232020333-0132213332131031-2012121203120303-1333000331032120-1202131120001010-1123301101201001-2130112022103101-1312012221031000)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2023122312000030-3003311220212320-2301130321111110-0030022112103323-1322310121101221-1022311201233033-2333330301223133-0212311012200001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133312313131002-0100122301303302-1303321300202002-1202333132321103-0111010011212120-0313111103313121-3123121030111030-0210013210323012"></a>

## trusted_clients.bot_skip_processing — bot_skip_processing / 021001131203 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-015.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- trusted_clients.bot_skip_processing

<a id="canonical-1022012022211010-1210121312322011-1023230031203200-2222210221220032-0120131033201333-0010311121223113-0012130302221330-1011032000131121"></a>

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
bot_skip_processing = {}
```

<a id="canonical-1110121121221011-1020221201101211-2033311330113110-1303202120020132-0322332222321111-1211112332211102-3203023002200133-1331233200112022"></a>

## Direct properties — bot_skip_processing / 021001131203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113330300110000-0222213220130100-3311033323203033-2133032010221000-2132103203132133-2123303010221031-3010030023032212-1232212100301121"></a>

## Next pages — bot_skip_processing / 021001131203 / 4

- [trusted_clients](resources--cdn_loadbalancer--reference--group-015.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1201221302301222-1010320030333121-0113223131202312-2330113130111003-3012011333202310-0101333000203222-3111003100213100-3013310003133311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300021013131333-2023322313210003-1233011311023113-2220121200213030-2103031013133012-3011132003211122-1311333313222023-2320330012100032"></a>

## trusted_clients.http_header — http_header / 113201110032 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-015.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- trusted_clients.http_header

<a id="canonical-0322123120203221-0100300302011021-1121311023232123-3320122012131233-0323003120103312-0110013312300132-0313133021231021-2312311132030230"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("headers")}
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
http_header {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302321112101131-3312330130112322-1100012333112201-3221302301003232-2121333032322102-3202210213013013-1131230232012000-0213213300310302"></a>

## Direct properties — http_header / 113201110032 / 3

- [headers](resources--cdn_loadbalancer--reference--group-015.md#canonical-2103310013311322-2231301123301132-1201212221203032-3131133002313323-2232222322322030-2331020231133012-0310131022023013-2131021111020033): complete subsection reference.

<a id="canonical-0202330033033003-3201101032310033-2110030122303322-3222012021210010-0133003200223232-3033323320212220-2310202102213110-0003333220022320"></a>

## Next pages — http_header / 113201110032 / 4

- [trusted_clients.http_header.headers](resources--cdn_loadbalancer--reference--group-015.md#canonical-2103310013311322-2231301123301132-1201212221203032-3131133002313323-2232222322322030-2331020231133012-0310131022023013-2131021111020033)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-015.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2103310013311322-2231301123301132-1201212221203032-3131133002313323-2232222322322030-2331020231133012-0310131022023013-2131021111020033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330223120330213-1122203211113231-0100102031010001-0133013012333310-2020322231312213-1222103000121321-3103102133210001-3100301011231121"></a>

## trusted_clients.http_header.headers — headers / 123302000001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-015.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- [trusted_clients.http_header](resources--cdn_loadbalancer--reference--group-015.md#canonical-1201221302301222-1010320030333121-0113223131202312-2330113130111003-3012011333202310-0101333000203222-3111003100213100-3013310003133311)
- trusted_clients.http_header.headers

<a id="canonical-2000322121230030-0022220021121112-2201233121221032-0031032311023330-0033033013130303-0330202202333213-1013310233302102-2230001012123033"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP header name and value pairs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "presence"),
  validators.ConflictingListObjectAttributes("exact",
    "regex"),
  validators.ConflictingListObjectAttributes("presence",
    "regex")}
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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301330122201323-1201303330100013-2301302020231032-3222333331031222-3232100301023330-1031113113203310-1330030232223323-2023120212110233"></a>

## Direct properties — headers / 123302000001 / 3

<a id="canonical-3220103332030110-1213001312123313-3210230312321221-1202012303032020-3302302201033210-3032323300023303-3303101030200212-3010332120301130"></a>

<a id="canonical-2102213231101232-2201311200303223-2002212323012321-2012201330113030-0113132321022021-2123001320302212-3310022232333002-2003112111323003"></a>

## exact property — headers / 123302000001 / 4

Type: `"string"`. Optional.

Exclusive with \[presence regular expression\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regular expression\] Header value to match exactly.

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-3313220001133133-1121130121321131-0033032031232200-0110202121223032-1222321101310311-0132131300001123-0310233103113002-0331332101013300"></a>

<a id="canonical-3120232323200023-1001311230101200-1023101321230333-2032221231102220-2021033132131300-0002101200333201-0202031130111123-0021220031001021"></a>

## invert_match property — headers / 123302000001 / 5

Type: `"bool"`. Optional.

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

<a id="canonical-2221010211133323-0323211213102111-1131310122303112-0023000120031133-0013031001122220-1030212330230011-0030103121230000-2203213123210300"></a>

<a id="canonical-2121310023230010-3012120222032111-3330112323210011-0303220131211020-2231321322131003-0032213222203200-3133231001302113-1231220013130012"></a>

## name property — headers / 123302000001 / 6

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

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

<a id="canonical-2031033203233313-3312313022123220-1111001331103320-1133222210201333-3231300230221133-3213321310203120-1300200201233320-2303230020111013"></a>

<a id="canonical-0003002301301101-2121303010120123-1131302120130210-3203010301213320-2233112320322011-1322013312330332-0021300212010231-0201302000002200"></a>

## presence property — headers / 123302000001 / 7

Type: `"bool"`. Optional.

Exclusive with \[exact regular expression\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regular expression\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3010020331200331-1321310000200122-0032230201111120-0101223100111023-2330121200203002-3210023111202221-0233110232221030-1023113131103111"></a>

<a id="canonical-0232100002200331-3211201000013121-0030210233312313-2313130333322102-1130310132112313-3210000211221001-3312200331301212-2212110213002323"></a>

## regular expression property — headers / 123302000001 / 8

Type: `"string"`. Optional.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

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

<a id="canonical-2020322113110212-3210203211103031-3103212302222330-3130221111021033-0111023013013210-0310233000300323-3220120000002010-1332323332210312"></a>

## Next pages — headers / 123302000001 / 9

- [trusted_clients.http_header](resources--cdn_loadbalancer--reference--group-015.md#canonical-1201221302301222-1010320030333121-0113223131202312-2330113130111003-3012011333202310-0101333000203222-3111003100213100-3013310003133311)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3321333031300023-0102112101330300-2231023022110222-0003023011013002-0222123013100321-2220230222132333-1231133330031312-2133323113312100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003323233110122-2101311312331100-0103033121323310-2101213231111233-3300022001220000-3331311133022013-1302231033003210-3100033133113003"></a>

## trusted_clients.metadata — metadata / 013011010210 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-015.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- trusted_clients.metadata

<a id="canonical-1000301111301000-1312320223100112-3200300211222232-2223203202013233-1130130031130303-0221103020230232-1310130022020320-0100132320130311"></a>

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

<a id="canonical-0300313112000203-0112301202311113-0123013103203203-2122032212120301-1301221010212331-0201003300033232-2022313213033033-2021111322211020"></a>

## Direct properties — metadata / 013011010210 / 3

<a id="canonical-0101221020300010-2332030003001321-1132211021032213-3223321111302102-0032303300130112-2032303103033213-2212031012313133-0101000330120211"></a>

<a id="canonical-1120122112013122-2312300321210232-1011231011212100-0232321103012332-2013221021303202-3131232113020320-3103312021131001-3013013233211022"></a>

## description_spec property — metadata / 013011010210 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2212213230022003-3300202232232023-0320131000010103-1132002310223131-0132300101302310-3101202001201002-3321023103032322-2210122212202313"></a>

<a id="canonical-2302123132212301-0130031303221012-0222111313110030-1233020023211330-2000302120021113-3301233203213321-2213202303133121-0233110032030001"></a>

## name property — metadata / 013011010210 / 5

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

<a id="canonical-0121212011330332-3231120331003020-2031020132313233-0202103311220103-0013303011031122-3322210232113213-3331032211022221-2230223010120222"></a>

## Next pages — metadata / 013011010210 / 6

- [trusted_clients](resources--cdn_loadbalancer--reference--group-015.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1133331223201300-3010010301001020-1022001210300133-3110022001203122-3310333101330031-3301103022232211-3311232311220231-2323022320130113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021311222023320-1301313323303301-1123012200132233-1133310103012011-1020103322331022-0301002220030130-1023022203333122-2030210133130102"></a>

## trusted_clients.skip_processing — skip_processing / 211030012030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-015.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- trusted_clients.skip_processing

<a id="canonical-3000003300203230-1300311212133303-1102110023002131-2322002121220233-1000012202023332-1131322103221222-0001101303123131-0003210232023011"></a>

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
skip_processing = {}
```

<a id="canonical-2332001231223302-2210333312232000-0212010203310132-2331220113200320-0222002022211022-2231201103222032-0222001321310123-2021020333001330"></a>

## Direct properties — skip_processing / 211030012030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301212230133200-3021222301101132-0022200113200212-0122331121202122-3322210003230101-3030210101011312-3130122212311320-1000121210132232"></a>

## Next pages — skip_processing / 211030012030 / 4

- [trusted_clients](resources--cdn_loadbalancer--reference--group-015.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0033212232020333-0132213332131031-2012121203120303-1333000331032120-1202131120001010-1123301101201001-2130112022103101-1312012221031000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002030232032320-3101300313221022-0320110012033323-2132201100110032-2020222130023230-3333130122013123-0132213200232212-2231122333013101"></a>

## trusted_clients.waf_skip_processing — waf_skip_processing / 100330321132 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [trusted_clients](resources--cdn_loadbalancer--reference--group-015.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- trusted_clients.waf_skip_processing

<a id="canonical-1122312021132001-2000211210010321-2322101221220113-0023320221030130-2333001313120122-3212313230120213-3013121310120001-3003311332210232"></a>

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
waf_skip_processing = {}
```

<a id="canonical-0013211321102132-2232031232233213-3202231020210010-0020112211202100-3222032031133203-3031203000021000-3330003321201113-3132112212301120"></a>

## Direct properties — waf_skip_processing / 100330321132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030321020122312-0231223132333333-3331222100211313-0003122321321003-0231210001320101-1233113110000020-1300132122231123-3203330320001301"></a>

## Next pages — waf_skip_processing / 100330321132 / 4

- [trusted_clients](resources--cdn_loadbalancer--reference--group-015.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3302200200302032-1033011220200321-0322330033221123-1310331022111310-3130333121221113-1313030021202201-1313311122312131-2320123301120231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010130010011222-1023031200311322-0312311210231231-2021013101102102-1132031020321102-2100031312312031-1310013101320131-0103010312112323"></a>

## user_id_client_ip — user_id_client_ip / 312222003012 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- user_id_client_ip

<a id="canonical-1312331103213230-3110310233011330-2210113221300211-1033233123013102-0122311213033031-2300023231023202-0020301210013322-3331103133021000"></a>

Type: `["object", {}]`. Optional.

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

- [user_id_client_ip](resources--cdn_loadbalancer--reference--group-015.md#canonical-1312331103213230-3110310233011330-2210113221300211-1033233123013102-0122311213033031-2300023231023202-0020301210013322-3331103133021000)
- [user_identification](resources--cdn_loadbalancer--reference--group-015.md#canonical-2112002310201312-1032232321300131-0300121020020022-3132103001332012-2011021103101103-2113302011323011-0303100300100221-2032311102302321)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
user_id_client_ip = {}
```

<a id="canonical-2303313021031103-3102013303123210-1103301233022230-3133013333010311-0203021102230030-1001322030121210-1011301023111033-3202203110000012"></a>

## Direct properties — user_id_client_ip / 312222003012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123123201132221-1100123330211032-0310033303310301-1011303202210331-3312320132033112-3110213132310102-1321230223103112-0200010311303213"></a>

## Next pages — user_id_client_ip / 312222003012 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1112230101301222-3222000100021310-1113222132033220-1211130133022022-2010313030300030-0312321030200003-0211330333102113-1230112232300132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210332013201120-2121101113030333-1130130122023333-3123203130102102-1303320130001321-2203001122022202-2333302032100112-1201021022010101"></a>

## user_identification — user_identification / 033213133333 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- user_identification

<a id="canonical-2112002310201312-1032232321300131-0300121020020022-3132103001332012-2011021103101103-2113302011323011-0303100300100221-2032311102302321"></a>

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
user_identification {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203322031020100-3023002333303210-2333130303122300-2332112211223101-3311320310213132-3333110222210030-0130201230000022-2201122001330303"></a>

## Direct properties — user_identification / 033213133333 / 3

<a id="canonical-0222001330202213-1301303101100023-3010112331222021-1330310122022030-0313213112103330-0212220211132231-0000101100233230-3122222223313303"></a>

<a id="canonical-3320112000011200-0221112323220003-3033320221310013-1331133020203222-3300120122332003-2212233233203210-3331231111130333-3332012201110132"></a>

## name property — user_identification / 033213133333 / 4

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

<a id="canonical-1013111212202101-3003101031301112-1102201131121302-2320023133132131-2310210303312031-2120230303002211-1122011133012300-0213012312013113"></a>

<a id="canonical-2202120010033203-2121322001102133-2110323211030132-2120003003303231-3021223303122132-3121122023301312-1310303233210102-1220022023101000"></a>

## namespace property — user_identification / 033213133333 / 5

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

<a id="canonical-3021200331031030-2111010322222101-1211310300301001-1001001320233213-0201110033122131-1113333003303012-2320113132130302-2020230313330101"></a>

<a id="canonical-0012021322130132-3103130100001113-3203100103102003-1303022133013223-1303233312011033-1233110033201011-2020100212321230-1033110120230203"></a>

## tenant property — user_identification / 033213133333 / 6

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

<a id="canonical-0330033121000222-3133002232003232-1031201101112130-1121100301312233-1130223003323323-3020331202331332-3002302223302231-2103033200323312"></a>

## Next pages — user_identification / 033213133333 / 7

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303321020120023-1201021311220032-1233102031311121-2222330321210323-1210303321303122-3112200213110312-1000123303200131-3222113101231230"></a>

## waf_exclusion — waf_exclusion / 310332313100 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- waf_exclusion

<a id="canonical-3133121121103332-1133013020232202-3130021232011203-2031210311001023-1203011213110230-3113302211313322-0321121230102331-1012313300100101"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for waf exclusion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("waf_exclusion_inline_rules",
    "waf_exclusion_policy")}
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
  "x-ves-oneof-field-waf_exclusion_choice": "[\"waf_exclusion_inline_rules\",\"waf_exclusion_policy\"]"
}
```

Terraform syntax:

```terraform
waf_exclusion {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210002313211222-3221112221311302-3333231133322032-2212032012230121-2020120103320112-0311321101031121-0200001121312202-2332323000310112"></a>

## Direct properties — waf_exclusion / 310332313100 / 3

- [waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201): complete subsection reference.

- [waf_exclusion_policy](resources--cdn_loadbalancer--reference--group-015.md#canonical-1100230020012103-3113212312312201-0333111112211303-3110022123123322-2311123203013303-0310013020333123-1123010130230011-0021210221300211): complete subsection reference.

<a id="canonical-3202203223033203-0222212122123003-3331331320220021-0221130001233122-1032312100011133-0313131020103031-3203121301313131-0033122303223323"></a>

## Next pages — waf_exclusion / 310332313100 / 4

- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_policy](resources--cdn_loadbalancer--reference--group-015.md#canonical-1100230020012103-3113212312312201-0333111112211303-3110022123123322-2311123203013303-0310013020333123-1123010130230011-0021210221300211)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333122002020202-3113122321010023-0222201020100113-2320322102223011-0101202030011003-2222300130101000-0232021303113120-1301133210210032"></a>

## waf_exclusion.waf_exclusion_inline_rules — waf_exclusion_inline_rules / 011011333202 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-015.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- waf_exclusion.waf_exclusion_inline_rules

<a id="canonical-1212001002312000-0012133303123123-2000232211122311-3002102033112112-2110022102222313-3310113010203102-2011310020212000-1223223023100130"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
waf_exclusion_inline_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0102030213200023-3330300020223113-0100023310132331-1022333333120001-3123121102231331-3103322313313101-3223033103120120-0201033110102301"></a>

## Direct properties — waf_exclusion_inline_rules / 011011333202 / 3

- [rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213): complete subsection reference.

<a id="canonical-3311033032232310-2302331313113001-0301003210310030-1121310013231303-2311201001222130-3311233003221033-1021133211313111-0201200001010222"></a>

## Next pages — waf_exclusion_inline_rules / 011011333202 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-015.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302322223323000-0302103222001023-3112101210232113-0211033322132301-1303320223121321-1220202302003202-2023200300330232-2233311133231120"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules — rules / 210223133232 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-015.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- waf_exclusion.waf_exclusion_inline_rules.rules

<a id="canonical-2003101333112322-3001223301200130-0032011223220110-0333103230321002-0220230210300333-3033300003002022-2313331312323233-0332030333122232"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of WAF Exclusions specific to this Load Balancer.

Upstream description:

An ordered list of WAF Exclusions specific to this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "exact_value"),
  validators.ConflictingListObjectAttributes("any_domain",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_prefix"),
  validators.ConflictingListObjectAttributes("any_path",
    "path_regex"),
  validators.ConflictingListObjectAttributes("app_firewall_detection_control",
    "waf_skip_processing"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("path_prefix",
    "path_regex")}
```

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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213100120213233-0202313102031132-3313103313010312-3231010131123322-1300311113001211-0010331233322102-3320111313011011-2112331032012202"></a>

## Direct properties — rules / 210223133232 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-015.md#canonical-3022320322112022-3333012112032330-3300003313032211-2323211231200303-2311023110331323-0032303033101131-3112110001223120-0303310010000223): complete subsection reference.

- [any_path](resources--cdn_loadbalancer--reference--group-015.md#canonical-0001231022023102-0110000230133130-0122331322013010-2103313031223130-3221102322032223-1230200321000030-2121233132312331-1321021200213020): complete subsection reference.

- [app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-015.md#canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021): complete subsection reference.

<a id="canonical-2323233211011212-3232322132321211-0130002201021032-3232001033020210-2310001021011303-1230113120021003-0112033033032313-2010230000203030"></a>

<a id="canonical-3103323002112333-3322021003220202-0321110003300003-1131201100322303-2130323113202023-1310212103013230-2313311021112233-1220131100220230"></a>

## exact_value property — rules / 210223133232 / 4

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0002210212203322-0113010033300122-2212111101011130-2331031200321000-2033210233130110-0333121323322023-3300131013103122-1132031310300101"></a>

<a id="canonical-3231012322330303-1020301310011011-1031201110202101-3013322213202202-0213302232103112-0303312213003013-3111311100302301-2232120230332201"></a>

## expiration_timestamp property — rules / 210223133232 / 5

Type: `"string"`. Optional.

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

- [metadata](resources--cdn_loadbalancer--reference--group-015.md#canonical-1100320320310223-0023211210033123-1113013301230002-3221300313013010-2113301212000110-0030310031330102-0213210202322323-3012112020212112): complete subsection reference.

<a id="canonical-0332332333223133-1223001321022120-3021022310330113-0222230123223311-2232131122000000-1232222120122112-2101333002330123-1130211002100220"></a>

<a id="canonical-2230312302000031-2103330322221023-3312003332311303-0022120330322203-2333032303002221-2202232202012323-2033222121301221-1310100322032021"></a>

## methods property — rules / 210223133232 / 6

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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

<a id="canonical-3011333123211130-1321133220120002-2011220000221302-0222021103301213-0221111331323232-3003033120331121-3313110113002310-2030203033022313"></a>

<a id="canonical-3312103301330032-1003022031022111-1310031022031231-1000201231010123-0231003102311032-3122113222121201-3023332120222023-3032211321120310"></a>

## path_prefix property — rules / 210223133232 / 7

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

Upstream description:

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths)

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0331312212031120-0211003013303102-1123021123013113-2132030210230003-3111002333032230-3122100012121111-2121132113321333-1300130211022023"></a>

<a id="canonical-1201031311302131-1211011132331030-0220000120202030-2201322001303113-3333323220110123-1310113131032212-3302101332332030-0321103330123122"></a>

## path_regex property — rules / 210223133232 / 8

Type: `"string"`. Optional.

Exclusive with \[any\_path path\_prefix\] Define the regular expression for the path. For example, the regular expression
^/.\*$ will match on all paths.

Upstream description:

Exclusive with \[any\_path path\_prefix\] Define the regular expression for the path. For example, the regular expression
^/.\*$ will match on all paths.

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0110233022303222-1300302212001333-3321220222203302-2232202201020030-0131301020031020-1110330322011101-3020102011113323-3010322221130022"></a>

<a id="canonical-3313201121300200-2011331222323301-1203322202221101-3023023221022313-3102021010313311-2001011131323223-2231031012323102-0103223221213213"></a>

## suffix_value property — rules / 210223133232 / 9

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [waf_skip_processing](resources--cdn_loadbalancer--reference--group-015.md#canonical-1112133013100123-1033321203321311-1212031110232013-2131012113311222-2100101110213002-0010002030211101-3313022003322302-1102030212210021): complete subsection reference.

<a id="canonical-3200123300311012-2310033002300033-2210331121020001-1110013101330003-2131031002120223-0030020132200113-0211101111232222-3211202022221322"></a>

## Next pages — rules / 210223133232 / 10

- [waf_exclusion.waf_exclusion_inline_rules.rules.any_domain](resources--cdn_loadbalancer--reference--group-015.md#canonical-3022320322112022-3333012112032330-3300003313032211-2323211231200303-2311023110331323-0032303033101131-3112110001223120-0303310010000223)
- [waf_exclusion.waf_exclusion_inline_rules.rules.any_path](resources--cdn_loadbalancer--reference--group-015.md#canonical-0001231022023102-0110000230133130-0122331322013010-2103313031223130-3221102322032223-1230200321000030-2121233132312331-1321021200213020)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-015.md#canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021)
- [waf_exclusion.waf_exclusion_inline_rules.rules.metadata](resources--cdn_loadbalancer--reference--group-015.md#canonical-1100320320310223-0023211210033123-1113013301230002-3221300313013010-2113301212000110-0030310031330102-0213210202322323-3012112020212112)
- [waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing](resources--cdn_loadbalancer--reference--group-015.md#canonical-1112133013100123-1033321203321311-1212031110232013-2131012113311222-2100101110213002-0010002030211101-3313022003322302-1102030212210021)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3022320322112022-3333012112032330-3300003313032211-2323211231200303-2311023110331323-0032303033101131-3112110001223120-0303310010000223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030132100321033-0311331120332002-0021321011303233-2222221121123230-3323330012111211-1002032330001331-3102221112233032-1113010103012121"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.any_domain — any_domain / 000320131312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-015.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_domain

<a id="canonical-2301113101301021-3123011031200031-2102123013001202-1333320301230313-0012200220031222-0320033033230301-2000112120133211-3023100100032223"></a>

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

<a id="canonical-2201023023323332-1323312232001101-1211023210231222-2320023031110131-3123211132113330-1121102110323103-3120332002010331-2130032301300011"></a>

## Direct properties — any_domain / 000320131312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101020133123131-2220102320223122-2223310231003213-3212002130220013-1231231011231012-0032033131211212-2013000221303303-3001331000221011"></a>

## Next pages — any_domain / 000320131312 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0001231022023102-0110000230133130-0122331322013010-2103313031223130-3221102322032223-1230200321000030-2121233132312331-1321021200213020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210010332032110-1002303231203202-3122231010313133-3032212111202003-3010300211131321-1312112022302231-1133002111120020-2311000201202200"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.any_path — any_path / 330023122303 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-015.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- waf_exclusion.waf_exclusion_inline_rules.rules.any_path

<a id="canonical-0102132222033230-2023322100221231-3001331113212133-3222123031120231-0001131310310221-2032121001020010-1302322002302032-1320131131332322"></a>

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
any_path = {}
```

<a id="canonical-3303221300102331-1323131030213302-1331001203000213-1112021322211201-2132021102222212-1002132231010330-2021111131000030-2112033332022103"></a>

## Direct properties — any_path / 330023122303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332223030223321-3203120211122311-2012122210023130-3103100302300223-1210130333300130-1110303032201220-0211121311220102-3003233033223020"></a>

## Next pages — any_path / 330023122303 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303321323321331-3122320002312322-0200103010233232-1302232222333300-1333133312032012-2320313133002211-2203313313320031-3020020000112022"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control — app_firewall_detection_control / 303130130023 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-015.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control

<a id="canonical-2020012320113230-2201321002210211-0102332322323000-1312222330132010-2123100212213122-0102012203133221-3012131021031130-2232200220300213"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
app_firewall_detection_control {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010121313310120-0130121223303210-2203012313212010-1011310130210120-2223221202201020-3031310221303232-1312313030010100-0322100021000032"></a>

## Direct properties — app_firewall_detection_control / 303130130023 / 3

- [exclude_attack_type_contexts](resources--cdn_loadbalancer--reference--group-015.md#canonical-3333111003113030-0310202220110102-3323000113120032-1120203111100230-2012120032100132-0010001302123233-2312133102313210-3000120031313022): complete subsection reference.

- [exclude_bot_name_contexts](resources--cdn_loadbalancer--reference--group-015.md#canonical-2132132020113210-2232202020110030-2223103032021030-1030133123030210-1100303210200002-1303020232023212-3332002212313210-1023120212221121): complete subsection reference.

- [exclude_signature_contexts](resources--cdn_loadbalancer--reference--group-015.md#canonical-2320222201123010-3220010212313021-0233121121320010-2023232310301102-3032322211311032-1002233201331200-1112001223000030-0022333123133120): complete subsection reference.

- [exclude_violation_contexts](resources--cdn_loadbalancer--reference--group-015.md#canonical-0323332330322210-0121333001000321-2210331102321302-2210302333313133-3302001200121103-3310033132222031-2223210102030211-1022223122131233): complete subsection reference.

<a id="canonical-0022211132013131-3322212230022102-0101310231211010-3023132011120010-1333012212000323-2233201210121001-2002001303112312-0001110022012120"></a>

## Next pages — app_firewall_detection_control / 303130130023 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts](resources--cdn_loadbalancer--reference--group-015.md#canonical-3333111003113030-0310202220110102-3323000113120032-1120203111100230-2012120032100132-0010001302123233-2312133102313210-3000120031313022)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts](resources--cdn_loadbalancer--reference--group-015.md#canonical-2132132020113210-2232202020110030-2223103032021030-1030133123030210-1100303210200002-1303020232023212-3332002212313210-1023120212221121)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts](resources--cdn_loadbalancer--reference--group-015.md#canonical-2320222201123010-3220010212313021-0233121121320010-2023232310301102-3032322211311032-1002233201331200-1112001223000030-0022333123133120)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts](resources--cdn_loadbalancer--reference--group-015.md#canonical-0323332330322210-0121333001000321-2210331102321302-2210302333313133-3302001200121103-3310033132222031-2223210102030211-1022223122131233)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3333111003113030-0310202220110102-3323000113120032-1120203111100230-2012120032100132-0010001302123233-2312133102313210-3000120031313022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103010030331130-3232112201303321-2130101223110101-0111210013303010-3120111123112200-3231330223221202-2001312113022300-3210302331032020"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts — exclude_attack_type_contexts / 013332112330 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-015.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-015.md#canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-1133001132012132-1021021010201032-1130101022303021-3133210221303231-0032130330000011-2330310231200313-2103121322212230-1131221002131222"></a>

Type: `"object"`. list nested block, Optional.

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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_attack_type_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222311011012300-0023012013132131-0221213333312210-2132211033032210-2311033210113133-1032302021130220-2023023012020033-1100202110311201"></a>

## Direct properties — exclude_attack_type_contexts / 013332112330 / 3

<a id="canonical-0313110220123233-3102022222221220-3201123131102112-0321100232213213-1211301321230302-3032132330300210-1010201232021013-3311223100333121"></a>

<a id="canonical-0331222122020000-3230203202313320-3101120313002002-3101233212111131-1112011112100213-0122220201122122-0131231203221032-2131331210131031"></a>

## context property — exclude_attack_type_contexts / 013332112330 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

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

<a id="canonical-2302012033213210-1000322110001233-0212023001333010-3310021220302202-1313021021201110-2000023333332223-0313213001110313-2101302231203330"></a>

<a id="canonical-2132002223210310-1030323322331230-3300001312020331-1301212232020131-3002122332312001-1211303331101332-2223330121310030-3200330030010002"></a>

## context_name property — exclude_attack_type_contexts / 013332112330 / 5

Type: `"string"`. Optional.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

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
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2012000002313331-3313012223100302-1003011311031321-3003323032032300-3232030212220332-2110023022123320-3133311213001112-1031001122223301"></a>

<a id="canonical-2332031311320100-3023333132311131-2202100222323003-3013200221112002-3231233133033031-3203321102003202-3222301301030021-2022021101023303"></a>

## exclude_attack_type property — exclude_attack_type_contexts / 013332112330 / 6

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ATTACK_TYPE_NONE",
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
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"),
}
```

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

<a id="canonical-0323122303120110-2223231333332333-1012033203332133-1322113201211312-3022000120032000-2310322032213123-1212131230333102-1201113323323313"></a>

## Next pages — exclude_attack_type_contexts / 013332112330 / 7

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-015.md#canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2132132020113210-2232202020110030-2223103032021030-1030133123030210-1100303210200002-1303020232023212-3332002212313210-1023120212221121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221023003203323-3311003221121202-3020331123003113-1100013123202111-2332301120002300-0020320033130212-3302001212200021-1022111220112310"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts — exclude_bot_name_contexts / 133030210321 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-015.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-015.md#canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-3312233330320232-3310011332123212-3023221120200112-3323310322303223-1233100323102003-1203032120200132-3310221111130222-2201223301130223"></a>

Type: `"object"`. list nested block, Optional.

Bot Names to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("bot_name")}
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_bot_name_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210000233010122-3201123003022211-3100230310321230-0003221201101333-1120211221300023-3233020100323213-2001211021002331-0322200332312203"></a>

## Direct properties — exclude_bot_name_contexts / 133030210321 / 3

<a id="canonical-0010222300330313-3212121223020132-3211213033123133-3210221121232001-2313120303112113-1022331101000201-3201212012003230-0323233003102000"></a>

<a id="canonical-1231333303333231-1331330302102132-2020332102212132-3301231222223130-0232003323213001-1123320133000030-3201313031010122-3233121131021132"></a>

## bot_name property — exclude_bot_name_contexts / 133030210321 / 4

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2011031333330231-2333102221122021-3021133331220220-2331222321301000-1202331011213332-0112011013121010-1313321211023123-3232001100012220"></a>

## Next pages — exclude_bot_name_contexts / 133030210321 / 5

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-015.md#canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2320222201123010-3220010212313021-0233121121320010-2023232310301102-3032322211311032-1002233201331200-1112001223000030-0022333123133120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122220033121212-0302132231222302-0112230223322333-0220000122331300-3312203333010300-3120000011232202-1003312130231021-0310331221221300"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts — exclude_signature_contexts / 232103332001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-015.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-015.md#canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-2233010310221202-3221123301320321-2213120020213031-3230032022103200-1110301030231333-3113021130030312-3120012123032111-1212131232011112"></a>

Type: `"object"`. list nested block, Optional.

Signature IDs to be excluded for the defined match criteria.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("signature_id")}
```

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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_signature_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130230012131030-2023111111131222-0211100130132123-2301312011011220-2122003223022312-0111120301132333-2101113121112110-1122311210312131"></a>

## Direct properties — exclude_signature_contexts / 232103332001 / 3

<a id="canonical-3313030202332321-2332231320132212-0332303102023311-0302021333133322-1113002213120030-1001200330202330-2123021021233212-2311332032032001"></a>

<a id="canonical-0003021110020203-1212202103323233-1032220331013320-2030200022201233-0200112312302223-2002100012023213-1233331310231202-2330210000000021"></a>

## context property — exclude_signature_contexts / 232103332001 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

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

<a id="canonical-0232002232021233-3300010232131103-0113223212331131-1113220113330332-1020212022333100-0211320120113023-3230101323222110-1230302022032113"></a>

<a id="canonical-1101021230002323-1011103220310303-1211330132113222-0020120303032232-1200331312302300-2213201201203213-3311332003202301-0110201000033300"></a>

## context_name property — exclude_signature_contexts / 232103332001 / 5

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

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

<a id="canonical-0313310131002133-2201322213200330-1100102333233312-0000132231102031-2221033301313010-2012023302332121-1031130132203032-0033113101031322"></a>

<a id="canonical-3022212020020223-1122213132111002-3033212133320333-1121013133320110-1133300100000121-1202232001203321-3300332321102333-1012031322212000"></a>

## signature_id property — exclude_signature_contexts / 232103332001 / 6

Type: `"number"`. Optional.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 299999999),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1310333021300103-3220233200232100-0022023022011122-1332332203320212-3201013102322201-0323323021011022-2103110133023123-0313012013100331"></a>

## Next pages — exclude_signature_contexts / 232103332001 / 7

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-015.md#canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0323332330322210-0121333001000321-2210331102321302-2210302333313133-3302001200121103-3310033132222031-2223210102030211-1022223122131233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131000202301312-1220332230030321-1301103030310132-1032013230230202-2312021202331221-3020212301310112-2200211111301100-2322201330201211"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts — exclude_violation_contexts / 022101000000 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-015.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-015.md#canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021)
- waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-1023100121212323-2112111211310031-2113300003133132-3231022310010220-2303302313203302-1321120301100333-0131222111131223-0123132133013121"></a>

Type: `"object"`. list nested block, Optional.

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

Terraform syntax:

```terraform
exclude_violation_contexts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211202233023122-1232311302210121-0121200303033201-3222323203322210-3223110002101021-0323013012200321-2233021213133231-1123223110202120"></a>

## Direct properties — exclude_violation_contexts / 022101000000 / 3

<a id="canonical-0110331312333112-3313312333321200-3030120233022232-1212121010011212-0121101303110132-3200330032112100-2201301232213310-0101322030323131"></a>

<a id="canonical-3301332010111022-3111031002131223-3132322021033323-3210322013110022-3100313221013232-0201332013303201-3301130312331213-3311333313102323"></a>

## context property — exclude_violation_contexts / 022101000000 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"),
}
```

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

<a id="canonical-3002301230110231-1021302012112210-2113012301302110-3013020013000111-3112011101303010-0122330102103231-2211220311130303-0221120220101202"></a>

<a id="canonical-3132213032330120-1220010210323133-3000221231211203-0133032001023302-0130203210222332-0102003022221123-2113022021232020-0311311111032031"></a>

## context_name property — exclude_violation_contexts / 022101000000 / 5

Type: `"string"`. Optional.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

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

<a id="canonical-1101223313211322-1103122123301102-3001311200232013-0002323102032003-1003132102223203-3001220322231022-3033123022323222-1012032321302022"></a>

<a id="canonical-2223011133001220-1011310000133211-1231103313120201-2312323300312230-3123103330022133-1313021203122122-1312131120213233-2100303201112303"></a>

## exclude_violation property — exclude_violation_contexts / 022101000000 / 6

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIOL_NONE",
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
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
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
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"),
}
```

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

<a id="canonical-0213310230320101-0323030111121222-0311012111132013-2330322221102131-1232023212231033-3131000132113102-3330010001102203-1302210230212131"></a>

## Next pages — exclude_violation_contexts / 022101000000 / 7

- [waf_exclusion.waf_exclusion_inline_rules.rules.app_firewall_detection_control](resources--cdn_loadbalancer--reference--group-015.md#canonical-2330222000211211-3020002030212203-1211032000223311-2020311030201233-3220213012112100-2113101032313203-1021110303020033-0211231233002021)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1100320320310223-0023211210033123-1113013301230002-3221300313013010-2113301212000110-0030310031330102-0213210202322323-3012112020212112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302232220100311-3011001023313313-2301330003221102-0013032323100023-1321100333210222-3330033100123320-2023220113233233-3322211232232331"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.metadata — metadata / 013213132222 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-015.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- waf_exclusion.waf_exclusion_inline_rules.rules.metadata

<a id="canonical-0231032313230033-1203113321100030-2112021202102233-3323022010221333-3322231313111203-2023213102032132-0211223010210210-3321320211322310"></a>

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

<a id="canonical-2233103232311331-1131311310002100-1330312120012320-2102111233122321-1202022020301010-0223333111300202-2210130110111313-0003102331300200"></a>

## Direct properties — metadata / 013213132222 / 3

<a id="canonical-0103031011131123-3201003320322122-0330320301002101-0310110311303000-2112312321022020-1021120121210131-1323120132311320-3333212113013022"></a>

<a id="canonical-0120320300231221-1031033000221223-3210320210200110-3302331101111130-1011321013000210-3331202231032022-2333010230112000-1032303111103212"></a>

## description_spec property — metadata / 013213132222 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1200103221033223-3032021102132112-1000232001302030-1103010302212213-1111102010303033-1202001303221221-0302113231100331-0213120001323103"></a>

<a id="canonical-3213333211033233-0021302311300221-1203202032031213-3203020102110020-3010320323232302-0321120320301212-0333000023331232-2132332220321001"></a>

## name property — metadata / 013213132222 / 5

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

<a id="canonical-1233230201111323-2123303122103232-3031203111020112-3203110232310000-2133113030202002-0321020033001103-3222101112210132-1311100331112132"></a>

## Next pages — metadata / 013213132222 / 6

- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1112133013100123-1033321203321311-1212031110232013-2131012113311222-2100101110213002-0010002030211101-3313022003322302-1102030212210021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033111130231320-2132330320001313-1120002203212000-2310223121003022-0132202022011022-2332301200010021-3331131112111120-0301221330200330"></a>

## waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing — waf_skip_processing / 333211322310 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-015.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [waf_exclusion.waf_exclusion_inline_rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-2201132000312023-0023121113132301-1202230300022001-1202131131121001-3122100011212322-1220203310011301-3303223031031221-0112121013303201)
- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- waf_exclusion.waf_exclusion_inline_rules.rules.waf_skip_processing

<a id="canonical-1111331011223303-2133310001310231-2010100100102101-2310323030123301-3222032002210131-1132102231330131-3112321310031012-2011131131231021"></a>

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
waf_skip_processing = {}
```

<a id="canonical-0132130313323123-3223001103031322-1130330320333032-1121000332021313-3130112201331132-2323333212030103-2121102210311003-0030223102323300"></a>

## Direct properties — waf_skip_processing / 333211322310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102022322211133-0013330000201110-1002101002320233-0132011021021311-2330330030011210-2111103020213210-2330311210233231-0100330322333313"></a>

## Next pages — waf_skip_processing / 333211322310 / 4

- [waf_exclusion.waf_exclusion_inline_rules.rules](resources--cdn_loadbalancer--reference--group-015.md#canonical-1300312300120222-1333020012122221-0223301311222322-3102221302311203-3303210213230301-3233011202331011-0131331122223332-0201033123221213)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1100230020012103-3113212312312201-0333111112211303-3110022123123322-2311123203013303-0310013020333123-1123010130230011-0021210221300211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210202312031011-0301123200111030-3201012112033302-2323123133132011-1212333120322211-0101211303320122-1110203033012121-1000212030002122"></a>

## waf_exclusion.waf_exclusion_policy — waf_exclusion_policy / 210012213303 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [waf_exclusion](resources--cdn_loadbalancer--reference--group-015.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- waf_exclusion.waf_exclusion_policy

<a id="canonical-0010311231300001-2332113020111233-1233233333022231-3012113130331000-3310331232001120-2201203310221101-3002203233111033-0000002320113332"></a>

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
waf_exclusion_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010021223103100-2021211010232131-0010110222012121-2002002212221222-0333032322210223-3233313103222232-0223122121212303-2201223300303111"></a>

## Direct properties — waf_exclusion_policy / 210012213303 / 3

<a id="canonical-3021211313200231-1222121300233100-0122203033101110-0123022312020032-3203303233101301-0020030212011123-3020031332331103-0302000133320132"></a>

<a id="canonical-2322000232320231-1302310003103221-1320131331310122-0301130023322030-2210230303222311-2331100221312300-3131123012101131-1122112223300130"></a>

## name property — waf_exclusion_policy / 210012213303 / 4

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

<a id="canonical-1220020012313013-2030230122330103-3103322002201132-2013033312300002-0222202230021121-2110000202330003-3011203313030031-1231030230120201"></a>

<a id="canonical-2231031113202013-0112221220300330-3220122012301030-2021222100113101-0211201132313333-2031233211003123-3221111333032330-1200011022210301"></a>

## namespace property — waf_exclusion_policy / 210012213303 / 5

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

<a id="canonical-1333113002222110-3220120222310113-0210330203121210-3321102101111301-3202203103202111-2221222032113310-2301022132321130-3022030131202300"></a>

<a id="canonical-3121233022321023-0101220201001120-3110222030333112-0020322122131013-3230321112313002-1100302032010022-0332313102300202-1302031103311003"></a>

## tenant property — waf_exclusion_policy / 210012213303 / 6

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

<a id="canonical-1111233303022230-1101202103002321-3221232030002303-1121302132103131-1012002001023210-2321230323121133-1033003003321002-2331131120211122"></a>

## Next pages — waf_exclusion_policy / 210012213303 / 7

- [waf_exclusion](resources--cdn_loadbalancer--reference--group-015.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
