---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-2032021010012021-0132232212131233-0032332301133102-1232110323201303-2330101102112103-1130013221101300-0002333311110200-1013110031022033"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site — virtual_site / 223300313000 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1031102212303210-2323303003313103-0120222000000003-2311032230222212-1102020333320302-0300003221102303-1013001233102200-2230333210301223)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323012200112321-0233320302012321-1013033013130212-2201201222211233-3100022112132300-0003111311310033-3131310012001112-0013303233230303)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1023212332121021-1203300333013112-0132223101030032-3311122322201033-2123023230101031-1310123320320032-3210221320220021-1232101110313030)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-2300231330123212-3103010303300101-0213100331032030-3311201123122011-3110131301020021-1313311120013120-2113123033032310-0000233300122230"></a>

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

<a id="canonical-0311310021022033-0200123103122222-0220003200003323-0100003203321033-3202202103023011-2120312132002321-3321112312212303-0112301023231021"></a>

## Direct properties — virtual_site / 223300313000 / 3

<a id="canonical-3131132331120122-3223303301123301-2101232000130133-0110213331001212-1023113003302020-2323320003000001-1113122310320223-3113313331132212"></a>

<a id="canonical-0111231020132322-1023310111011201-2133021232121020-0101030322213310-3201131122202211-1330022032123023-0101222123213020-2210210331232031"></a>

## name property — virtual_site / 223300313000 / 4

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

<a id="canonical-2320302213133330-2003233320333010-3120322322022011-3100302220212012-0032213221203032-0321231123031230-2003331231121121-1310331311310020"></a>

<a id="canonical-0300131101013311-1133203222102332-1210011101022213-0033011201212020-0001330221201212-2212313223231301-0333203301110110-0212223320211211"></a>

## namespace property — virtual_site / 223300313000 / 5

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

<a id="canonical-3103322221002010-3112120013320302-0112122121211133-0303332320303230-2220021313200122-2301030113111032-3133020232120111-2212201211220030"></a>

<a id="canonical-3103223112020022-2333201220312202-1222303000022313-3332010100322311-1010331302233323-2002130200012022-0311011220302013-3012223110002230"></a>

## tenant property — virtual_site / 223300313000 / 6

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

<a id="canonical-2002323030003010-1000012121031230-1123300003133000-1200022030222133-1131011323220211-1320112210103231-0333032321310330-1123310213331032"></a>

## Next pages — virtual_site / 223300313000 / 7

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1023212332121021-1203300333013112-0132223101030032-3311122322201033-2123023230101031-1310123320320032-3210221320220021-1232101110313030)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3312032130021330-3212102123320232-0213011310323122-0022113030030231-3323210230320221-1030023032302123-3132212013230101-0110213332222021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030210210201311-2221102301330222-1303130300103111-2131132021211213-3310321110000200-0132112200100311-0000300033011120-2313201320230103"></a>

## proxy_advertisement.do_not_advertise — do_not_advertise / 020202201200 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- proxy_advertisement.do_not_advertise

<a id="canonical-3031312323231020-0113200313213210-0033130220231123-1200131232031233-2103303311133022-0312121011122032-3321311301333313-0022033223202221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for do not advertise.

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

<a id="canonical-0123323033232010-0300011320310302-0222233320013020-3333322230111212-0202002212010002-1203133210101023-2030001322121213-0131122110221300"></a>

## Direct properties — do_not_advertise / 020202201200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120001112220010-3222330212300010-0313220303313031-2122020213300223-0300021222333131-3122102033223131-3202201000133023-3212123120302102"></a>

## Next pages — do_not_advertise / 020202201200 / 4

- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132311111123002-0110211303010120-2303113300123310-3230212312202323-2212222311120330-3322023103333321-1331011231101233-2222303220233013"></a>

## proxy_config — proxy_config / 321032121221 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- proxy_config

<a id="canonical-1231210333102133-2313202213322022-2310312222033130-2121022032332201-0122120101110200-0120012230031221-1133223213031110-3132012221102322"></a>

Type: `"single"`. Computed.

HTTP/HTTPS Load Balancer. HTTP/HTTPS Load balancer.

Upstream description:

HTTP/HTTPS Load balancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]"
}
```

<a id="canonical-1231312332312012-3200032122022331-0033132102112222-0303221331130132-2302223211032013-2200133122002133-1002031210201131-0010221311033330"></a>

## Direct properties — proxy_config / 321032121221 / 3

<a id="canonical-3230213210313020-2020303033023122-2323101130120200-1313312022031113-1112113100102300-0022132002300312-3311322232313120-3133303303320000"></a>

<a id="canonical-2002311203021123-2330131021001111-1022230112230133-1022233303220112-1031202220002033-2110121002131103-1202122323020100-3011031130300001"></a>

## domains property — proxy_config / 321032121221 / 4

Type: `["list", "string"]`. Computed.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\`.

Upstream description:

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*-bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*-bar.example.com\`\` will match
\`\`baz-bar.example.com\`\` but not \`\`-bar.example.com\`\`. The longest wildcards match first.

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [http](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0011021330202210-0002020013211203-3031010020133013-3110000232312130-0001313101300321-0312001331300310-1333323101313213-2101212130332011): complete subsection reference.

- [https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321): complete subsection reference.

- [https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330): complete subsection reference.

<a id="canonical-3200203110121213-2120011330031013-1223312031101020-3331322123330210-2102000301130313-3231100003222213-2113012200323033-1212322202330233"></a>

## Next pages — proxy_config / 321032121221 / 5

- [proxy_config.http](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0011021330202210-0002020013211203-3031010020133013-3110000232312130-0001313101300321-0312001331300310-1333323101313213-2101212130332011)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3103011203132331-3110230112302332-0302323312123030-3232232001220112-3303213211033012-0221300221133102-0310301320311230-1123110002321330)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0011021330202210-0002020013211203-3031010020133013-3110000232312130-0001313101300321-0312001331300310-1333323101313213-2101212130332011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233100031122321-3111000023333301-0023103030131332-0223023033301310-0233020220010023-1311000112303031-1203010130210231-0003303131313211"></a>

## proxy_config.http — http / 020323111323 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- proxy_config.http

<a id="canonical-0201322302313330-2111332301120033-2311111103100321-2223001033000122-2303030300200010-0123002001300030-1231120123232113-3331133131313001"></a>

Type: `"single"`. Computed.

HTTP Choice. Choice for selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

<a id="canonical-1021002122123330-1203312000222311-1131333121033030-2332111013113332-2002111232120000-2221321020121033-1103322223213011-2210303300112320"></a>

## Direct properties — http / 020323111323 / 3

<a id="canonical-1210201310033332-1003311010320030-3133110110022310-0002213030132210-0202003123220012-1131033011023200-1333202133031030-0022132022202320"></a>

<a id="canonical-3032103032031330-0011013321011301-0212301333201302-0330020211321112-1221230112210113-1110110210020233-3130312201232032-0101212231032102"></a>

## dns_volterra_managed property — http / 020323111323 / 4

Type: `"bool"`. Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Upstream description:

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1023000321232100-3033123133210311-0033302131012302-3330311301123220-2133120132310030-0233033313132110-2221113122320312-2110022001032332"></a>

<a id="canonical-0332223223211323-3100311223031110-3101011113012202-0133003103021120-3032131220322102-1112212133231031-3031211233330112-0010322330313000"></a>

## port property — http / 020323111323 / 5

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3232003033211322-0311310003213102-2003220133301100-3120331330312313-1010033232123131-0211002112220210-0311000132030102-2222131220211202"></a>

<a id="canonical-2003230103013000-0321202122313030-3221330120312203-3231231002322012-0321213223221223-1032321301103233-0002133333230222-2210000123201223"></a>

## port_ranges property — http / 020323111323 / 6

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-2011312010331130-3010222200110231-1303132213311230-1002123010002123-3200032310320210-2330031203101231-3333031010100220-2011112101111332"></a>

## Next pages — http / 020323111323 / 7

- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203002313200223-2213210311200230-0010201133212121-0223031023320031-2123231312303132-1133013303102010-1113100030003200-3033002033010301"></a>

## proxy_config.https — https / 103100003210 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- proxy_config.https

<a id="canonical-3131210001331111-1202123021322110-2023023311130330-0230100302301332-3210232200200233-3231011021323203-0223301322310112-2131200112121210"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

<a id="canonical-1312303230003233-3131233303110123-2230230232102003-1310131211222103-0131212330232320-3223033102123123-1211022032313000-3213003301120033"></a>

## Direct properties — https / 103100003210 / 3

<a id="canonical-1033332002033311-0310122232131012-3000123133001122-1133033130213120-1012202021300310-1113210103231132-3120110302221133-1230033211020032"></a>

<a id="canonical-1022130211002112-2231000333031021-1123221102300101-3132201222103101-1300300230211323-2203101013022302-1301302201211201-0110003033110132"></a>

## add_hsts property — https / 103100003210 / 4

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0223203231112032-2102010131131012-0010222220012032-2320121200110231-0201231312120210-3313312210212133-2103233013121220-2322213331032110"></a>

<a id="canonical-3123003111211330-1002230333231322-0302301022021110-2103133300112110-1213131321303011-2333211330100313-1130101003331032-3000032322233302"></a>

## append_server_name property — https / 103100003210 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1010220232111020-0012112302000310-0323101031023100-1131011320103303-2213211001212310-2123110123021133-0300131020331201-1303223232212322): complete subsection reference.

<a id="canonical-2233330131102310-0213111020102010-3103212210033110-1323001302233332-3030301321022101-1100131331000023-2301102112310332-2120310102200102"></a>

<a id="canonical-0230110113130032-2331030033021102-3133132022312133-3312220020013301-1022201010111211-2121102313121130-2110133212231003-0223233002101113"></a>

## connection_idle_timeout property — https / 103100003210 / 6

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0331121010100121-1223320123102002-2331110023123211-2303112313223333-2133310113300122-2002110011321130-1233020121210300-1001201113203300): complete subsection reference.

- [default_loadbalancer](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1313301030300323-0311033002233211-1100331312211011-1102320003312100-3100332020323101-0103312113202220-2231202312103223-3032221212311332): complete subsection reference.

- [disable_path_normalize](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3011231311323101-2323203002201101-2000230001120021-2313313201321132-0313030103323133-0120113203223211-0223200330233310-0301312223212022): complete subsection reference.

- [enable_path_normalize](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0222102223110101-2311000212121200-1022311003231221-3301110103020210-0322312322100322-1023011012132101-1330012032323331-1102031231122003): complete subsection reference.

- [http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010): complete subsection reference.

<a id="canonical-0200210020210213-3213310120131313-3221320311112110-2100001202031232-2311022013033230-0211323002212202-2131212311012203-1311323331322023"></a>

<a id="canonical-3020231202321033-1030011322102032-2202213003002322-3313202000133211-2122011332231122-0121332211201023-3200021330323112-1122000112113212"></a>

## http_redirect property — https / 103100003210 / 7

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [non_default_loadbalancer](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0020110300033033-3031303222203120-0333011030102032-1223202132032201-0021211131323220-2101023111332012-1233131232230211-3202203330302201): complete subsection reference.

- [pass_through](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0202323330310033-1233203111130031-0320101100321323-0002103202032223-1323222010212022-1111323110203232-0232022132031303-1112203333301222): complete subsection reference.

<a id="canonical-1101220131221202-1212023323213220-0230133033002030-1330101003300333-2310032112031011-1213110221132113-1113030113321013-0102010112230020"></a>

<a id="canonical-0232213003000213-1001332133210110-1022211200232103-2010321031331302-0120213111003101-3132111120033210-3133003003231301-1033123201001230"></a>

## port property — https / 103100003210 / 8

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1022101111013211-1231330120003213-2212133223010003-0230311311112201-1121323200233003-2302011111230103-0121010203203021-0301122101221311"></a>

<a id="canonical-0032103132133031-3130333302233001-3210132120131032-1333300233320130-1223213021200323-3220321120011233-0200031021231022-1132122000121220"></a>

## port_ranges property — https / 103100003210 / 9

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-1332220233223222-1031003231203133-0201002030200123-0320021221312111-3331330023012203-0332310332313210-2011132312021320-2210220112111222"></a>

<a id="canonical-0133033331012022-1321131203211132-0101103030232220-3130123001100322-2003202123013112-0201020333131023-1110310020001220-2122103220000320"></a>

## server_name property — https / 103100003210 / 10

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320): complete subsection reference.

- [tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100): complete subsection reference.

<a id="canonical-0021201331133201-1231100321111233-3231110131032320-1112301211222230-3200302012100211-0032323321112100-2130022032231302-2023133210322013"></a>

## Next pages — https / 103100003210 / 11

- [proxy_config.https.coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1010220232111020-0012112302000310-0323101031023100-1131011320103303-2213211001212310-2123110123021133-0300131020331201-1303223232212322)
- [proxy_config.https.default_header](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0331121010100121-1223320123102002-2331110023123211-2303112313223333-2133310113300122-2002110011321130-1233020121210300-1001201113203300)
- [proxy_config.https.default_loadbalancer](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1313301030300323-0311033002233211-1100331312211011-1102320003312100-3100332020323101-0103312113202220-2231202312103223-3032221212311332)
- [proxy_config.https.disable_path_normalize](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3011231311323101-2323203002201101-2000230001120021-2313313201321132-0313030103323133-0120113203223211-0223200330233310-0301312223212022)
- [proxy_config.https.enable_path_normalize](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0222102223110101-2311000212121200-1022311003231221-3301110103020210-0322312322100322-1023011012132101-1330012032323331-1102031231122003)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- [proxy_config.https.non_default_loadbalancer](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0020110300033033-3031303222203120-0333011030102032-1223202132032201-0021211131323220-2101023111332012-1233131232230211-3202203330302201)
- [proxy_config.https.pass_through](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0202323330310033-1233203111130031-0320101100321323-0002103202032223-1323222010212022-1111323110203232-0232022132031303-1112203333301222)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1010220232111020-0012112302000310-0323101031023100-1131011320103303-2213211001212310-2123110123021133-0300131020331201-1303223232212322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022210122321122-2001000232033321-1113013130201233-3101030001230002-3033113001213332-1220013220023122-0213023120202301-2002321031003121"></a>

## proxy_config.https.coalescing_options — coalescing_options / 121001001131 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.coalescing_options

<a id="canonical-0223231011323112-3000232123323222-0332233310120310-3001030300330231-3333200112321222-2333101032331223-3233012033000113-1013110120100322"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

<a id="canonical-0222333121023101-0003222023322122-1021122330320001-3113311301210000-1330113300101130-0123023221022023-3323110220032321-3022021200201002"></a>

## Direct properties — coalescing_options / 121001001131 / 3

- [default_coalescing](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0201301100311301-1230220231000123-2233210031232121-1103320120113103-3000203121201132-3030130233012232-2310100331002020-3321010322121200): complete subsection reference.

- [strict_coalescing](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3000101003000001-2010013302122010-3010031131312123-1001210321223233-2012120213223202-3220021213011122-2032131333230223-0220112231020330): complete subsection reference.

<a id="canonical-3310210002120330-2312123130301022-1300320231320201-1132002121321012-2300113303321133-1323023211010312-0200210302012031-2213321200012323"></a>

## Next pages — coalescing_options / 121001001131 / 4

- [proxy_config.https.coalescing_options.default_coalescing](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0201301100311301-1230220231000123-2233210031232121-1103320120113103-3000203121201132-3030130233012232-2310100331002020-3321010322121200)
- [proxy_config.https.coalescing_options.strict_coalescing](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3000101003000001-2010013302122010-3010031131312123-1001210321223233-2012120213223202-3220021213011122-2032131333230223-0220112231020330)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0201301100311301-1230220231000123-2233210031232121-1103320120113103-3000203121201132-3030130233012232-2310100331002020-3321010322121200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210312300213023-0320111330201302-2022022223030032-0323201022303330-3130303303201232-0122302100011112-2233001203323020-1202312113121012"></a>

## proxy_config.https.coalescing_options.default_coalescing — default_coalescing / 003200133231 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1010220232111020-0012112302000310-0323101031023100-1131011320103303-2213211001212310-2123110123021133-0300131020331201-1303223232212322)
- proxy_config.https.coalescing_options.default_coalescing

<a id="canonical-3132130121020200-1300333001210012-2311212022321233-3322322020003031-1330023133132101-0212313130211000-3131112022021112-2110313130020002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-0313232010132002-2103021110203011-2101111113323121-1212231023122212-3323001133223003-1012030213302010-0203031220003323-2333223130123223"></a>

## Direct properties — default_coalescing / 003200133231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302002312013003-1222011333131120-1321300001330132-3000130310322232-2231110312031320-3311310032111022-2101232103112322-0320333321131013"></a>

## Next pages — default_coalescing / 003200133231 / 4

- [proxy_config.https.coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1010220232111020-0012112302000310-0323101031023100-1131011320103303-2213211001212310-2123110123021133-0300131020331201-1303223232212322)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3000101003000001-2010013302122010-3010031131312123-1001210321223233-2012120213223202-3220021213011122-2032131333230223-0220112231020330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000320010103331-2130113211220131-0131111010120332-1000332320200210-1101032131110021-2310213203130222-1330013330002100-3213012122300302"></a>

## proxy_config.https.coalescing_options.strict_coalescing — strict_coalescing / 132120103111 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1010220232111020-0012112302000310-0323101031023100-1131011320103303-2213211001212310-2123110123021133-0300131020331201-1303223232212322)
- proxy_config.https.coalescing_options.strict_coalescing

<a id="canonical-3102333313000021-0230033211232111-3120020123023330-3113203323133023-1311301301013100-3203302311230302-0300221320311203-0322000303303103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-1210220220230013-3001030210212112-3123323201031220-3021231311301032-3033321210332121-3303110220012232-0300331021130132-2213300022331222"></a>

## Direct properties — strict_coalescing / 132120103111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3021021120220313-0310313210200012-3213300100033132-1030113021313112-3303012011231322-3011203022300030-2130233200110130-1010313120131003"></a>

## Next pages — strict_coalescing / 132120103111 / 4

- [proxy_config.https.coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1010220232111020-0012112302000310-0323101031023100-1131011320103303-2213211001212310-2123110123021133-0300131020331201-1303223232212322)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0331121010100121-1223320123102002-2331110023123211-2303112313223333-2133310113300122-2002110011321130-1233020121210300-1001201113203300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200200203132031-0203230110002033-1331220220223011-0311132211002311-1322321212012031-3322221132332130-3212113211212213-0102023132001110"></a>

## proxy_config.https.default_header — default_header / 012331202213 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.default_header

<a id="canonical-3311022003010330-0320020132023301-2110032022032111-0123300032111330-3101032303020332-0330110132002333-3130312030302322-3303222033330030"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

<a id="canonical-1022022033213103-3130013332211133-3300211303100330-2100032023033032-1330220303122102-1030030023111013-3102300033313312-3332111000030312"></a>

## Direct properties — default_header / 012331202213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331311033220211-1221020300300312-2212230203121321-3130201030331201-0220210013022322-1002231202322003-1333210111201223-1120121020121023"></a>

## Next pages — default_header / 012331202213 / 4

- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1313301030300323-0311033002233211-1100331312211011-1102320003312100-3100332020323101-0103312113202220-2231202312103223-3032221212311332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133031200202110-1330223322320212-0012323102302011-2123021232021223-1203102103310002-3100131132101022-3202220111232233-2210033313112021"></a>

## proxy_config.https.default_loadbalancer — default_loadbalancer / 233101023322 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.default_loadbalancer

<a id="canonical-2303231011310021-2121212100000201-2231210333303212-0101131301121120-3300202303301302-2030222012131020-3003202301032210-0111002003101313"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default loadbalancer.

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

<a id="canonical-0033301311233132-1310131212300323-2120210203320002-0230111202030123-2001121320011001-3232221320312031-0231203213230003-2122132331202103"></a>

## Direct properties — default_loadbalancer / 233101023322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102321112121212-2303012213023021-0301023200322221-0230321002020232-1302030201201001-1110221330130131-3032221300223121-0131123023100022"></a>

## Next pages — default_loadbalancer / 233101023322 / 4

- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3011231311323101-2323203002201101-2000230001120021-2313313201321132-0313030103323133-0120113203223211-0223200330233310-0301312223212022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302213233103312-0011231020033300-1123200133212013-0313230230331022-2020203112211001-0133103032133111-2231022002333311-3213200000012110"></a>

## proxy_config.https.disable_path_normalize — disable_path_normalize / 320011012130 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.disable_path_normalize

<a id="canonical-1111030201013103-3030332220030332-1303232300120201-1230223320032332-2013002313202013-1020001333213210-3333010100231210-2001123300333111"></a>

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

<a id="canonical-2112223333320011-1033332031131010-1030102121020032-2200123123321002-3100332003130131-1220100313010203-2312200210310221-1123323123120313"></a>

## Direct properties — disable_path_normalize / 320011012130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313111332232012-0013321211103132-2210323302032223-1201201122210100-0132210033302310-2331002200032203-3032332300320031-2301330221333122"></a>

## Next pages — disable_path_normalize / 320011012130 / 4

- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0222102223110101-2311000212121200-1022311003231221-3301110103020210-0322312322100322-1023011012132101-1330012032323331-1102031231122003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321323323122001-1123031300223233-2002012010120221-0212100020320120-1323113203230303-1312021323310032-2323121011222112-1111203213103210"></a>

## proxy_config.https.enable_path_normalize — enable_path_normalize / 313212020302 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.enable_path_normalize

<a id="canonical-2231231313011132-3322020313011102-2121111311033203-1203133323122223-2312030121223312-3300012011023103-1100122221111221-3331121110132233"></a>

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

<a id="canonical-1013010331312123-2120202203031032-0132211003132221-3220323003020101-1100220302322210-2221203103200233-3021012333310102-2213332133030023"></a>

## Direct properties — enable_path_normalize / 313212020302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111210223010313-1003323030112330-2300211232112003-1110103323020111-3231011333131332-1323221112223030-2131211203201002-1203310300233103"></a>

## Next pages — enable_path_normalize / 313212020302 / 4

- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012113330000301-3202132202111113-1120323233232213-0332312012000310-1112130101010001-2210100110321100-1200310221200112-1332300021210223"></a>

## proxy_config.https.http_protocol_options — http_protocol_options / 211311302122 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.http_protocol_options

<a id="canonical-2020232033203000-0122333001010013-1013233013122011-2123211103311130-0331303022202031-2321020020100120-3032203031221001-1332313132211202"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

<a id="canonical-1003002331103033-0121023320233022-2200110131113313-1023003123321302-3003021112130312-3330012023123131-1030112211121321-3031101311032020"></a>

## Direct properties — http_protocol_options / 211311302122 / 3

- [http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3233023112203213-1011001323132233-1220102230110322-2210300302000131-3011103312112021-0113211223102301-2222303003302201-0100031132101201): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0222230210022000-1200302133300220-1010303223330123-0200323010132223-3333331033303302-1331011023131010-1212131032200301-2233303112321313): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3231020003213122-0011232310032202-2031100100021022-2023321111122223-0312313330210300-2100301102131223-3013030123222221-0110333213303222): complete subsection reference.

<a id="canonical-2321023201322323-3021122131112211-2033233333332201-3030212321022031-3100102000223132-2100312232030310-0111301222202133-0121212030331322"></a>

## Next pages — http_protocol_options / 211311302122 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3233023112203213-1011001323132233-1220102230110322-2210300302000131-3011103312112021-0113211223102301-2222303003302201-0100031132101201)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0222230210022000-1200302133300220-1010303223330123-0200323010132223-3333331033303302-1331011023131010-1212131032200301-2233303112321313)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v2_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3231020003213122-0011232310032202-2031100100021022-2023321111122223-0312313330210300-2100301102131223-3013030123222221-0110333213303222)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3233023112203213-1011001323132233-1220102230110322-2210300302000131-3011103312112021-0113211223102301-2222303003302201-0100031132101201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011033301302123-1202323101210310-0102103110110233-0023011030020101-2222201032022213-2030210122113101-2332213021302120-2332120222210221"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 112312332103 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-0312330010131033-2113021130013333-3311222321101023-3301101332031012-0223120330321002-0013222321131022-1003303300133320-1110030102223220"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3332320031130201-2203112012033120-3003012001020012-1003210313200321-0101231103223231-2133213233233101-2013012130132312-2231233111303312"></a>

## Direct properties — http_protocol_enable_v1_only / 112312332103 / 3

- [header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1331100001023220-0230220013312022-0302023012213121-1230022100210132-2201121001111331-2300313313231221-3013332233322120-0212010103100201): complete subsection reference.

<a id="canonical-0132312001123003-1231223020113103-0303300023201212-1002310111231021-1011030013310302-0333311002022002-3002211012103331-1310202310331200"></a>

## Next pages — http_protocol_enable_v1_only / 112312332103 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1331100001023220-0230220013312022-0302023012213121-1230022100210132-2201121001111331-2300313313231221-3013332233322120-0212010103100201)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1331100001023220-0230220013312022-0302023012213121-1230022100210132-2201121001111331-2300313313231221-3013332233322120-0212010103100201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022120212111303-3230012102302101-1000301032212300-0231200321203030-1301331301100020-0213033321333112-3122003121323210-2123002012201033"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 223303202012 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3233023112203213-1011001323132233-1220102230110322-2210300302000131-3011103312112021-0113211223102301-2222303003302201-0100031132101201)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-1310032333303232-2300012323233112-2132020103112102-0122213120020001-3331112213011302-0103110211202313-0100020123221012-1031031121033320"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

<a id="canonical-3201211232133030-3331201332212021-2031000012222030-0022020031233230-0223111020111022-3302020123102011-2210132230011202-1102321211321302"></a>

## Direct properties — header_transformation / 223303202012 / 3

- [default_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0000310123233001-2021232010200113-2232111321012130-3112323131323233-3122210010023211-2031102203113003-1023201223022102-2133113211220020): complete subsection reference.

- [preserve_case_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0223202012222003-3210230110011202-0012302020011112-2233313012231223-1021031322131211-1323300231032100-0321233030133330-1323303313331012): complete subsection reference.

- [proper_case_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0123231201002122-1311233120132130-3022220110013000-2111212311131102-2001213202222011-2230201120301321-1101130022130200-0210121121300130): complete subsection reference.

<a id="canonical-1220333203222130-0232122320313102-0303002302232211-0300222213202003-2002230220311002-3032023120102030-0001023120203300-0011000101000111"></a>

## Next pages — header_transformation / 223303202012 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0000310123233001-2021232010200113-2232111321012130-3112323131323233-3122210010023211-2031102203113003-1023201223022102-2133113211220020)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0223202012222003-3210230110011202-0012302020011112-2233313012231223-1021031322131211-1323300231032100-0321233030133330-1323303313331012)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0123231201002122-1311233120132130-3022220110013000-2111212311131102-2001213202222011-2230201120301321-1101130022130200-0210121121300130)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3233023112203213-1011001323132233-1220102230110322-2210300302000131-3011103312112021-0113211223102301-2222303003302201-0100031132101201)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0000310123233001-2021232010200113-2232111321012130-3112323131323233-3122210010023211-2031102203113003-1023201223022102-2133113211220020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222222111032131-0310202323102232-0321311310203013-0230300003301221-1012033100232313-2230030322200302-3231012023120212-0021000222211220"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 222122022220 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3233023112203213-1011001323132233-1220102230110322-2210300302000131-3011103312112021-0113211223102301-2222303003302201-0100031132101201)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1331100001023220-0230220013312022-0302023012213121-1230022100210132-2201121001111331-2300313313231221-3013332233322120-0212010103100201)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-3002202220301331-1023230122223303-1123331023010202-2311332132123212-0301202300121011-2312320323210203-2312122310011313-3000112322000020"></a>

Type: `["object", {}]`. Computed.

Use the platform's current default HTTP header transformation behavior.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3213030011200131-0131112303212333-1311022030303333-1120311322101220-1013023000102311-1320233331120313-1320121310003130-3200202013203313"></a>

## Direct properties — default_header_transformation / 222122022220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303122220010332-0022021032011130-1231320011111333-0211301213032020-3030330330113323-1133121311121231-1112322210013003-3323031010023023"></a>

## Next pages — default_header_transformation / 222122022220 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1331100001023220-0230220013312022-0302023012213121-1230022100210132-2201121001111331-2300313313231221-3013332233322120-0212010103100201)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0223202012222003-3210230110011202-0012302020011112-2233313012231223-1021031322131211-1323300231032100-0321233030133330-1323303313331012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002221113022220-0220022331131121-1313121021302231-0011010211222330-2013122312000230-1310210313022133-2200302133100212-2110112132212323"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 132310202312 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3233023112203213-1011001323132233-1220102230110322-2210300302000131-3011103312112021-0113211223102301-2222303003302201-0100031132101201)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1331100001023220-0230220013312022-0302023012213121-1230022100210132-2201121001111331-2300313313231221-3013332233322120-0212010103100201)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-2132310332210102-2230333101302022-1301321113101011-0303313201231311-2232010310031303-2301212233322300-3112211303013012-3031333211021333"></a>

Type: `["object", {}]`. Computed.

Preserve HTTP header-name case when upstream case must remain unchanged.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2103333013311013-1310220331333010-1302000201222222-2133221103031020-2220002211131012-3222133103310031-0303131232001010-3112311223003132"></a>

## Direct properties — preserve_case_header_transformation / 132310202312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312130100111012-3130330303211210-3233200111021221-3023120232011032-1323221312032201-0023020232223002-0202222312331033-3122021303000101"></a>

## Next pages — preserve_case_header_transformation / 132310202312 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1331100001023220-0230220013312022-0302023012213121-1230022100210132-2201121001111331-2300313313231221-3013332233322120-0212010103100201)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0123231201002122-1311233120132130-3022220110013000-2111212311131102-2001213202222011-2230201120301321-1101130022130200-0210121121300130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322302211222310-3221011102302230-0030000220003323-2133302011320221-1313333221002322-0132222232013232-0313202121023011-2023112112330020"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 202100022120 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3233023112203213-1011001323132233-1220102230110322-2210300302000131-3011103312112021-0113211223102301-2222303003302201-0100031132101201)
- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1331100001023220-0230220013312022-0302023012213121-1230022100210132-2201121001111331-2300313313231221-3013332233322120-0212010103100201)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-0022233102323321-2202130121021321-3022113221120012-2133100300302121-3312300100323020-2331301011203112-2210113103100301-3331220311022212"></a>

Type: `["object", {}]`. Computed.

Transform HTTP header names to proper case when explicit transformation is required.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1012331230222012-2030201133333010-0323313211202300-3112010022330232-3110113033113020-1222223121313112-2103012112011230-2213031223230313"></a>

## Direct properties — proper_case_header_transformation / 202100022120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122230330221320-1020011201213222-3102130213113202-0300100033013000-0002202022302321-2201021130231201-0322210313011202-1110203101320302"></a>

## Next pages — proper_case_header_transformation / 202100022120 / 4

- [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1331100001023220-0230220013312022-0302023012213121-1230022100210132-2201121001111331-2300313313231221-3013332233322120-0212010103100201)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0222230210022000-1200302133300220-1010303223330123-0200323010132223-3333331033303302-1331011023131010-1212131032200301-2233303112321313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133003302102220-1001021301322100-0333303320001321-0321231003233130-0133020232021121-2121120231230131-1033230222202223-1333022130332200"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 222120313132 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2213131012001100-0113012110132301-2032101031231233-0300102231313021-0033312202320102-0032110130132220-0320333001100211-2312221321320313"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

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

<a id="canonical-0121101101333112-0101333110211300-0302311001113221-3332303032031021-0222333020321102-1023202320332320-3033200311101231-0111031121202320"></a>

## Direct properties — http_protocol_enable_v1_v2 / 222120313132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332303101030100-2233313133302122-2123102321023002-3303102032220320-0122230301220000-1002313123323312-3202120323132130-2230131312010210"></a>

## Next pages — http_protocol_enable_v1_v2 / 222120313132 / 4

- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3231020003213122-0011232310032202-2031100100021022-2023321111122223-0312313330210300-2100301102131223-3013030123222221-0110333213303222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202230210033222-2112302002022212-1310223200101303-1213322313013330-2010022000201012-2232311002030033-0203301200203221-1230120021213231"></a>

## proxy_config.https.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 331322230303 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- proxy_config.https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-2232132303313112-0200000020322011-2312131010032230-1133220100302323-2313203231132131-1100312111222313-2020013221221200-3212233222133223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

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

<a id="canonical-2101233003303233-3002020321323100-0322323230110001-3031003032110020-1011221211010302-0000122022220331-2320221310200112-1123000322221200"></a>

## Direct properties — http_protocol_enable_v2_only / 331322230303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233222011012013-2110032302322023-0220333200101221-1022111222132200-3203000012103232-0302111322012312-2122100210123111-1022311233331211"></a>

## Next pages — http_protocol_enable_v2_only / 331322230303 / 4

- [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1213013322013023-0100011121032010-2111132031101200-3012330211223233-2233012332331311-3212123322120213-2301033120301220-3002002101110010)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0020110300033033-3031303222203120-0333011030102032-1223202132032201-0021211131323220-2101023111332012-1233131232230211-3202203330302201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211003103123202-1302010023200323-1031331200123120-1211022013031131-2203322202032231-3322012323220312-3120313200012313-2020321002210133"></a>

## proxy_config.https.non_default_loadbalancer — non_default_loadbalancer / 010133103222 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.non_default_loadbalancer

<a id="canonical-0210223233031120-1000222222333132-2232011202202123-1323202312132210-3220221210321110-2121233222203100-1002020110113313-3320223323010312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

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

<a id="canonical-3010311133202031-2302223020113001-3231112030011023-2003132200221102-0030131330101103-3110201312230333-3023302303211231-2000033131211323"></a>

## Direct properties — non_default_loadbalancer / 010133103222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002100122333232-1322201102210012-3031233301223321-1003110120001313-2021330322000110-2210013101320130-0323101212322021-2110313303221332"></a>

## Next pages — non_default_loadbalancer / 010133103222 / 4

- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0202323330310033-1233203111130031-0320101100321323-0002103202032223-1323222010212022-1111323110203232-0232022132031303-1112203333301222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101010322301203-0302211321110320-1001021232222020-3211112310011123-2200001333230321-1333013123103233-1030033203133331-2200200313201302"></a>

## proxy_config.https.pass_through — pass_through / 010221311023 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.pass_through

<a id="canonical-0332022220001232-1320033131013103-3110333312203221-3201111230332103-1020312003332212-2100300020002312-2332002321120110-3312010012132321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

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

<a id="canonical-0220212123302222-1300112330200102-2022101001312003-3312213320133332-1103030000310310-1231320310122131-3301213222231023-0102111033021031"></a>

## Direct properties — pass_through / 010221311023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301001231231123-2312211210120110-3011121201000012-0213000011203020-0132030323132101-1000320100233313-0002331001323223-1132212203002030"></a>

## Next pages — pass_through / 010221311023 / 4

- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203321133030233-2132010102222010-0112331103023222-2111203330133210-1300302321303010-2002201210301120-3100323100222122-2033210021013133"></a>

## proxy_config.https.tls_cert_params — tls_cert_params / 013013310023 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.tls_cert_params

<a id="canonical-3122310321332130-3023333001301322-0022230033301010-0131301212333330-1122301332230332-1223002012120301-2023012112110013-2230001312331233"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-3313322020032312-3021032032033230-3031302121311110-2330213031011031-3331020203322032-3030033031030011-2320210222102121-3022002101222320"></a>

## Direct properties — tls_cert_params / 013013310023 / 3

- [certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0022323232333203-1012222123002210-1310333231033311-0212213221333303-1120213321301313-0002212330132130-0133000303310211-3232133333203200): complete subsection reference.

- [no_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0120000123102023-2110310333100130-1323131321110301-3112223211301323-1113211100021123-2200111000330010-1013210122222003-0100101303223011): complete subsection reference.

- [tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023): complete subsection reference.

- [use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303): complete subsection reference.

<a id="canonical-0112210203211210-3332213222230332-3110121320203230-1033220233232011-0310023311123333-1310131133100013-0210012003111123-0133200230221023"></a>

## Next pages — tls_cert_params / 013013310023 / 4

- [proxy_config.https.tls_cert_params.certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0022323232333203-1012222123002210-1310333231033311-0212213221333303-1120213321301313-0002212330132130-0133000303310211-3232133333203200)
- [proxy_config.https.tls_cert_params.no_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0120000123102023-2110310333100130-1323131321110301-3112223211301323-1113211100021123-2200111000330010-1013210122222003-0100101303223011)
- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0022323232333203-1012222123002210-1310333231033311-0212213221333303-1120213321301313-0002212330132130-0133000303310211-3232133333203200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203021211333022-2023010311020013-3030000322320030-1033012213120322-0011003123132330-0231313331330120-1311110323333221-1221102113030230"></a>

## proxy_config.https.tls_cert_params.certificates — certificates / 320011321123 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- proxy_config.https.tls_cert_params.certificates

<a id="canonical-0012010113200032-1223311313032332-0112311213223312-3313231133311231-1011023100130310-2232130112302203-1220310022321221-3013131330301021"></a>

Type: `"list"`. Computed.

Select one or more certificates with any domain names.

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

<a id="canonical-2312033213200023-0323123211030200-3333030320211131-3333322000133102-2321203122113002-1013201023100202-0113303332210203-2310301311221232"></a>

## Direct properties — certificates / 320011321123 / 3

<a id="canonical-3113212031330030-0033131021032003-3121120032313332-3130033230202101-1222231220312122-2012331330023023-1003121031121310-1012133132010220"></a>

<a id="canonical-0202310212111023-2122110332331301-0301000212033212-1031010223302133-2101103321320000-1333132133333031-2321232030300120-0112131222120320"></a>

## name property — certificates / 320011321123 / 4

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

<a id="canonical-2300202210323110-2321000122221301-2331100300223320-2032010101023033-1103011233022101-2333131112110021-1302333322130131-1111332030123303"></a>

<a id="canonical-2000333232203212-2123122213321012-1330111222220023-0312010312021201-2010102210031032-0021022322231023-1102203010231231-0012323023123223"></a>

## namespace property — certificates / 320011321123 / 5

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

<a id="canonical-0210320223023102-3312023001031210-1030223113322213-2200022333222200-1202302303102002-2222121323122113-1100203230102033-2323120132021133"></a>

<a id="canonical-0011201311303000-0210203203030333-1331313032010303-2133323221110123-3000323310023323-3320312112120221-0103023032133302-2323222203320231"></a>

## tenant property — certificates / 320011321123 / 6

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

<a id="canonical-3311331302031303-3001112223010100-1113210300321120-1320210133321221-3333132132001310-0300320001121130-3330320212022303-3101102000232222"></a>

## Next pages — certificates / 320011321123 / 7

- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0120000123102023-2110310333100130-1323131321110301-3112223211301323-1113211100021123-2200111000330010-1013210122222003-0100101303223011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132030221202122-1130333210211202-0330032121203021-3130020333101231-1033223131310221-2020323020120200-1130232111221030-3023200023202332"></a>

## proxy_config.https.tls_cert_params.no_mtls — no_mtls / 000033013223 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- proxy_config.https.tls_cert_params.no_mtls

<a id="canonical-2102303322033200-1222032222111111-3222021312212332-1203120020123222-0203120132001203-0133032000211022-1012212201022013-2302332220120212"></a>

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

<a id="canonical-3131023030000110-0231131230010130-3202232000112030-3323102032131021-0112200203011233-1120331031333332-3312323332202310-1211122213213012"></a>

## Direct properties — no_mtls / 000033013223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201032302222302-2332201012122323-3013302001332111-3102230223100232-0132120001010101-3002322121002113-3200010133332232-3010002011110122"></a>

## Next pages — no_mtls / 000033013223 / 4

- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213013132331203-3110222133302300-3012232230011221-1211331333332332-1011303023312031-2332213011130130-3133302123203223-3231022011312300"></a>

## proxy_config.https.tls_cert_params.tls_config — tls_config / 323222331323 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- proxy_config.https.tls_cert_params.tls_config

<a id="canonical-3331311330100132-3321222311203013-0322132311102103-3111101333133032-3022112121202201-1133332222201330-2310300210303020-2210033103123021"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-1213210113310121-0233001133300020-2220202300323323-1212323330231202-0100300012030021-2111103222320031-2310302321112103-0031233323012212"></a>

## Direct properties — tls_config / 323222331323 / 3

- [custom_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0200200113202333-1213031102323202-1233223223020000-2211020331232110-1320131001110031-1133120131212233-0322033200310121-3031313033323210): complete subsection reference.

- [default_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2222233232033312-1112222110113100-3000132201001332-3220221331132200-3201022133222321-0221203300200212-3112322232100030-0231022321133012): complete subsection reference.

- [low_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1202300221110131-1303212030233202-0013311003011331-0022322101032012-2123221213133331-0230113231000313-3313111110102130-1313033003113301): complete subsection reference.

- [medium_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1131213330232321-0201220033010230-0203023031001232-0123131112131011-2111101132132020-3302211322031211-2320003213100202-2133312312020011): complete subsection reference.

<a id="canonical-3331220022133212-1113033020002002-1323330010231303-0130033011003331-2302003333031212-2221202320302333-1122233220003020-3310333322031101"></a>

## Next pages — tls_config / 323222331323 / 4

- [proxy_config.https.tls_cert_params.tls_config.custom_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0200200113202333-1213031102323202-1233223223020000-2211020331232110-1320131001110031-1133120131212233-0322033200310121-3031313033323210)
- [proxy_config.https.tls_cert_params.tls_config.default_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2222233232033312-1112222110113100-3000132201001332-3220221331132200-3201022133222321-0221203300200212-3112322232100030-0231022321133012)
- [proxy_config.https.tls_cert_params.tls_config.low_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1202300221110131-1303212030233202-0013311003011331-0022322101032012-2123221213133331-0230113231000313-3313111110102130-1313033003113301)
- [proxy_config.https.tls_cert_params.tls_config.medium_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1131213330232321-0201220033010230-0203023031001232-0123131112131011-2111101132132020-3302211322031211-2320003213100202-2133312312020011)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0200200113202333-1213031102323202-1233223223020000-2211020331232110-1320131001110031-1133120131212233-0322033200310121-3031313033323210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110121132131233-2321322331102212-3222313011220230-0212330322210333-1211233212103320-2233303030223022-1230022222100210-1031213101103203"></a>

## proxy_config.https.tls_cert_params.tls_config.custom_security — custom_security / 123030023301 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023)
- proxy_config.https.tls_cert_params.tls_config.custom_security

<a id="canonical-0202303133113310-3000002010303332-0013312210203002-2121222322131221-2122320102211002-0022222132233030-3322310113211333-1303120103013122"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1103023310301100-2310212001111331-0231011021312223-2022022120222221-2012330300031013-0000030110101022-2312121213303133-0032013122232330"></a>

## Direct properties — custom_security / 123030023301 / 3

<a id="canonical-2103322232331000-1130011110013132-2210120302100021-0220231022100110-2331011130223100-3102113300330002-1312221131033221-1110230120232001"></a>

<a id="canonical-3200020030022030-2233102200323021-1012301321300230-2020012022200022-1323212102023101-1220233101230221-2031320022330121-2113330120321323"></a>

## cipher_suites property — custom_security / 123030023301 / 4

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1230012312302133-3033300223120113-2233333230333120-3211212030211321-1131010311300111-0032213312220111-2033033300033311-3230113110323301"></a>

<a id="canonical-0020110223313303-3200110312332101-0223103333320003-0110102130320113-2032103331333102-0310023310213212-3302032210312131-1102211303133112"></a>

## max_version property — custom_security / 123030023301 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3233021320101312-3322002332302312-1103332310202200-0300002232300102-0112232333131330-2302333312133123-1032201012301110-3332200322023011"></a>

<a id="canonical-3031122021312020-3313020323131033-2022111222013303-0001020303202023-1302220322102101-1103020231120123-1223201031303320-1111331222311313"></a>

## min_version property — custom_security / 123030023301 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0131331231130013-1212030103200123-3121302100123221-3232330221000221-1321033031010311-0030013023330303-3033203102202332-2213103332002000"></a>

## Next pages — custom_security / 123030023301 / 7

- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2222233232033312-1112222110113100-3000132201001332-3220221331132200-3201022133222321-0221203300200212-3112322232100030-0231022321133012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100323202321020-3210200021123002-0230111102302003-3022232003131331-0112220001313311-2211223303300312-1302000330131302-2330211203003133"></a>

## proxy_config.https.tls_cert_params.tls_config.default_security — default_security / 011031212102 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023)
- proxy_config.https.tls_cert_params.tls_config.default_security

<a id="canonical-2011131020303231-0123323201330332-0100220102332203-2121210113201232-0320230333022103-0303310022103201-1003323223010120-0022110033210301"></a>

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

<a id="canonical-1023031122311201-2011122113112221-1011102113031320-3111001223113002-1110232320132132-0120121112202013-2122230332303231-1121200333103011"></a>

## Direct properties — default_security / 011031212102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3110003212023232-2112313212232310-3330230100311122-0223002132303002-0322131232223300-0220213310333103-1031110203020230-1220202221101122"></a>

## Next pages — default_security / 011031212102 / 4

- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1202300221110131-1303212030233202-0013311003011331-0022322101032012-2123221213133331-0230113231000313-3313111110102130-1313033003113301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213001020021112-0031032202133320-0323000212332101-3333121322231130-3133131111230113-3332312223333030-3322112120100333-3333203033201313"></a>

## proxy_config.https.tls_cert_params.tls_config.low_security — low_security / 232233033000 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023)
- proxy_config.https.tls_cert_params.tls_config.low_security

<a id="canonical-0100201211030102-1220202111320022-2130213102221113-1230132000300222-2231111201020112-1001223011203132-3321022013321111-2213020030331301"></a>

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

<a id="canonical-1321022013200110-2012000110332303-1212201212211012-3310221033322222-0111233211013333-1023313312310321-0023201101221110-2121003322001012"></a>

## Direct properties — low_security / 232233033000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333011203202112-2112200323130300-0220110100003023-2103003032123212-2100132301100313-2021031003230031-3333030203122102-2003001323311120"></a>

## Next pages — low_security / 232233033000 / 4

- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1131213330232321-0201220033010230-0203023031001232-0123131112131011-2111101132132020-3302211322031211-2320003213100202-2133312312020011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201321032010123-1331113311210110-0203112112210211-2100112013202221-1220321131100013-1223110120021100-0032232213233300-1220230222223011"></a>

## proxy_config.https.tls_cert_params.tls_config.medium_security — medium_security / 120200021330 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023)
- proxy_config.https.tls_cert_params.tls_config.medium_security

<a id="canonical-2210202220021222-0010200321103210-3100212230303022-0212233233102310-2211021233303230-0101011113203123-1213302300221002-2300113312231300"></a>

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

<a id="canonical-2030001113131020-2323022322121123-1002130320033203-3232333202333331-0203133003210023-0331133133013100-3202322013322031-0203102223231130"></a>

## Direct properties — medium_security / 120200021330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0330233101230320-2333002223021212-0301233001001303-2332233302203312-0111032032212323-3000131203303300-2110203123022101-2102022222330133"></a>

## Next pages — medium_security / 120200021330 / 4

- [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2232021111122322-3121310113320000-2022320323113131-3102020123100221-2102120230003123-3201310200100311-1303000332121010-3331011112120023)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033031112330320-1122111033102122-1121331301230123-2311011200032212-2311023320212331-3002312213112330-3222033123302311-1231210210120021"></a>

## proxy_config.https.tls_cert_params.use_mtls — use_mtls / 220102333033 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- proxy_config.https.tls_cert_params.use_mtls

<a id="canonical-3323333331121130-0101330023022013-2001210012100000-2022000100121200-0310131122310202-0101031010201313-1203001013330301-3133120100203002"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-3210203111020110-3001120023323002-0230000321033202-2321112213221303-1031030001112211-2202020303311220-1102111111333313-3113112101221201"></a>

## Direct properties — use_mtls / 220102333033 / 3

<a id="canonical-0223021201013221-2100222321101211-0212033301233020-3300102302231001-3020001133101233-2200013120303212-0212310202313002-1311213003002001"></a>

<a id="canonical-3330010001201300-1022102212102021-0003023230122123-0023311233330111-0231310203210333-0323030111100230-0111320032323110-1212033012202213"></a>

## client_certificate_optional property — use_mtls / 220102333033 / 4

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [crl](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2310221313023321-0120200010103302-0013013300313200-2100011310201311-0213220213000213-2212033212203220-1331210103023001-0322000010301330): complete subsection reference.

- [no_crl](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0012222132331101-2312210321320022-1002303101213002-2320123030211333-3332021112023303-2203130332133011-1100332132002313-2332111000132322): complete subsection reference.

- [trusted_ca](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1323013112333303-2220322000200100-1313103200210002-1101023103223001-0130031033110100-1132301003102203-3222321102111210-2303322212012013): complete subsection reference.

<a id="canonical-3110001310323220-3130221312113321-0010222103233203-2221133032231331-2013020333121313-2033130232023133-2203102323220113-0211111022323322"></a>

<a id="canonical-2010210331120311-2030112012202013-2330010303211332-0232130010332233-3312001310032313-1222301222210212-1122222212310012-1321312302021002"></a>

## trusted_ca_url property — use_mtls / 220102333033 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1121132130223121-1103202113102313-2312023321211323-3130331013222023-2133111021003202-1100011201021312-1200110012001001-0103033211131010): complete subsection reference.

- [xfcc_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0220112023033102-0311310031331012-1221203323231030-2202230310001320-0232220222120003-3233120103310003-3302301023113023-0203232112121210): complete subsection reference.

<a id="canonical-1310230300202121-0010003330320212-1023021011223300-2022223323032221-2320333122311020-3312130220312130-2331201222321312-2310231332202010"></a>

## Next pages — use_mtls / 220102333033 / 6

- [proxy_config.https.tls_cert_params.use_mtls.crl](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2310221313023321-0120200010103302-0013013300313200-2100011310201311-0213220213000213-2212033212203220-1331210103023001-0322000010301330)
- [proxy_config.https.tls_cert_params.use_mtls.no_crl](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0012222132331101-2312210321320022-1002303101213002-2320123030211333-3332021112023303-2203130332133011-1100332132002313-2332111000132322)
- [proxy_config.https.tls_cert_params.use_mtls.trusted_ca](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1323013112333303-2220322000200100-1313103200210002-1101023103223001-0130031033110100-1132301003102203-3222321102111210-2303322212012013)
- [proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1121132130223121-1103202113102313-2312023321211323-3130331013222023-2133111021003202-1100011201021312-1200110012001001-0103033211131010)
- [proxy_config.https.tls_cert_params.use_mtls.xfcc_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0220112023033102-0311310031331012-1221203323231030-2202230310001320-0232220222120003-3233120103310003-3302301023113023-0203232112121210)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2310221313023321-0120200010103302-0013013300313200-2100011310201311-0213220213000213-2212033212203220-1331210103023001-0322000010301330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212302202210233-3011113031111031-0220310213203033-0310302300012232-1223301222202302-1123033113023333-3010121022111032-1302123301021213"></a>

## proxy_config.https.tls_cert_params.use_mtls.crl — crl / 320223132100 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303)
- proxy_config.https.tls_cert_params.use_mtls.crl

<a id="canonical-2013230130231331-3111213220120122-1313330120031121-3122113321103211-3032213332212102-2211011113012213-2011330023021230-1231311103321013"></a>

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

<a id="canonical-2220210103332210-0333100332323313-2232112310302323-3103320303100222-3113133102011123-3330032223023332-0112302102021013-0033230322212322"></a>

## Direct properties — crl / 320223132100 / 3

<a id="canonical-1000220020200012-1003223023321320-0103022332311222-2032202232021003-3222030213011123-0020020213120103-3110002330030210-1302002311000023"></a>

<a id="canonical-1102110020021113-1110233321221213-1221302131122231-2310322322110312-2031101032010202-2211323321300123-0331030201103222-1320230022113113"></a>

## name property — crl / 320223132100 / 4

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

<a id="canonical-2031100130301331-3300031331121033-2200220010131121-3122323222123120-1303201210332213-3121202201032201-2300200301030033-0021221312101001"></a>

<a id="canonical-0013220213021113-3113313302003220-3013302133201323-0220220001113311-1030132333301302-0013031113303123-3210023333132030-3233323333202021"></a>

## namespace property — crl / 320223132100 / 5

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

<a id="canonical-3303213323110123-3232230130022133-0322003010110203-1213330313232020-2121013100320312-3313122301211103-1031003131311232-1003100213220001"></a>

<a id="canonical-3233201031223023-1003011013232213-1321130201333231-1213130310030200-0033030320020322-0220003113230101-1230310012001032-3233022012232020"></a>

## tenant property — crl / 320223132100 / 6

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

<a id="canonical-1013123133131322-3113012222311221-0102120022312012-3133201213133221-1123023202131032-3023200010232121-1222012303010221-2331121222121233"></a>

## Next pages — crl / 320223132100 / 7

- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0012222132331101-2312210321320022-1002303101213002-2320123030211333-3332021112023303-2203130332133011-1100332132002313-2332111000132322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331232112012310-0023200032223230-2202210221112031-2210213021312020-3213301200032101-1012310031333213-2321333000230121-2300103132030013"></a>

## proxy_config.https.tls_cert_params.use_mtls.no_crl — no_crl / 230301313032 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303)
- proxy_config.https.tls_cert_params.use_mtls.no_crl

<a id="canonical-3220230133002012-3023132332222313-3220100102330013-2300120223232302-0020132312130331-1013233322312203-1313301012222031-2021220011013031"></a>

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

<a id="canonical-0232020233220132-3320221202230003-1202331232200121-2021023020003213-0113033003133230-0210021122010222-2100001223001223-2132332131112303"></a>

## Direct properties — no_crl / 230301313032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333133013130032-1331222132113320-2312013031100313-2202102123131223-1130223323102202-1103223312000130-0300231222310303-0001123030332202"></a>

## Next pages — no_crl / 230301313032 / 4

- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1323013112333303-2220322000200100-1313103200210002-1101023103223001-0130031033110100-1132301003102203-3222321102111210-2303322212012013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212312131222023-3200210102122010-1202123203321210-1300212201300103-2131113220133311-0331100020322220-2012001303320311-3003211331333311"></a>

## proxy_config.https.tls_cert_params.use_mtls.trusted_ca — trusted_ca / 330111032123 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303)
- proxy_config.https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-1320322012231103-1013213030132322-0101012012332221-3102031120102301-0033003121300312-1021302301331323-2311302232320021-2320011202013231"></a>

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

<a id="canonical-2323001011321313-3121220021030012-0333232230231310-3303100310112203-2021213303323032-3011232333022132-2031122011113030-0232230331323220"></a>

## Direct properties — trusted_ca / 330111032123 / 3

<a id="canonical-3321312102123103-1121310202312232-3032023022101320-3300103312022211-3112323132131321-0023200103112102-3121113030303321-1120101230302221"></a>

<a id="canonical-0031002232233030-0310203302002333-3331202000121012-1303233123232310-1120300333202112-1121312301000013-2123303003101102-2221212200003030"></a>

## name property — trusted_ca / 330111032123 / 4

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

<a id="canonical-0133201231011233-1110333033023312-0221110313013023-0003201011032021-1112232010112123-1312221202310113-3213020213100002-0331201030101132"></a>

<a id="canonical-2203210223102313-3313331100212231-2230101302023210-1203110031022101-3202032331211202-3323303002003100-2112322113220312-1312300132110203"></a>

## namespace property — trusted_ca / 330111032123 / 5

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

<a id="canonical-1110030112223112-0310110120121200-3330030101000020-1130311213301203-0213323213001013-1131212121112123-1032231210320302-1023111232323032"></a>

<a id="canonical-2030132311133003-3121021300030132-1211222112120030-3130222311010123-2102033311113331-0130312012110310-3110032033212100-2222203300301333"></a>

## tenant property — trusted_ca / 330111032123 / 6

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

<a id="canonical-2002001123120003-1320231013232121-0012200220211133-0122111232213011-2120332020230201-0213223233311221-2010333110102000-0320020010222001"></a>

## Next pages — trusted_ca / 330111032123 / 7

- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1121132130223121-1103202113102313-2312023321211323-3130331013222023-2133111021003202-1100011201021312-1200110012001001-0103033211131010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323010123200001-3331103223211103-3000033203102232-3122100101123021-3202000120012100-2230211301103010-2022011230013321-1212300313330313"></a>

## proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled — xfcc_disabled / 210120231010 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303)
- proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-2330113221112113-1320323320130300-3313132313032302-3312003333210001-2220001302302011-3003022122212333-2121133122130332-0210201001032031"></a>

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

<a id="canonical-1203013301121012-1200300302011303-0302111110103130-0231311100303212-3331012322030203-0330133110132000-3222210331221332-2002330103201032"></a>

## Direct properties — xfcc_disabled / 210120231010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230311221322030-3003321111121323-1112110211112211-2203323020313333-0211302332321120-1203233311121003-1312130023030013-3003220022100011"></a>

## Next pages — xfcc_disabled / 210120231010 / 4

- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0220112023033102-0311310031331012-1221203323231030-2202230310001320-0232220222120003-3233120103310003-3302301023113023-0203232112121210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031331302121302-1333230232123000-1112120331213001-2001122121203233-3121132021012212-2331330331323233-1231132202212212-2331012230220222"></a>

## proxy_config.https.tls_cert_params.use_mtls.xfcc_options — xfcc_options / 300033331133 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130100101203213-3100002030123311-2020313003130132-0210033233130112-0010010231331130-1320120201010003-1003131102003301-2330013030000320)
- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303)
- proxy_config.https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-0333001012300132-2333031203202200-2011203013312100-1320310020101130-3201013330100303-0311122213123113-3101332233200123-2223302310111222"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0133022122222302-2333023332133223-3122231300203101-3011033020130111-2110202030310211-2103023003212032-3121103202322022-0002022121120120"></a>

## Direct properties — xfcc_options / 300033331133 / 3

<a id="canonical-0131321223303031-0133121131121120-2230320210110233-2202112121130211-3310323101220121-0231121311121001-0200301133333020-0122310320232103"></a>

<a id="canonical-2030233103002112-1032103232330332-0031110210300210-3030002223003220-3013212121032011-3121100303223323-3201021301132031-1321202332121210"></a>

## xfcc_header_elements property — xfcc_options / 300033331133 / 4

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-1131310322232203-0111101121123120-3111333022333013-2103021320311222-0011310010321213-2020031321020330-1033221031132323-0211302113133211"></a>

## Next pages — xfcc_options / 300033331133 / 5

- [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1313122123223121-3103032123203302-3100213201112311-3302012232300010-3313302003131122-3311111111320220-1101321031312110-2032222030331303)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011102022310221-2002332311211202-3311000130113232-1311222201130112-2010322330000302-2313321131132132-3001122202230212-2331212011300303"></a>

## proxy_config.https.tls_parameters — tls_parameters / 121300303212 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- proxy_config.https.tls_parameters

<a id="canonical-1201123003111312-3320313130233200-2321331213113100-0003203220311322-1320311011310223-1033012213322111-1232231132122133-0030211221111110"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

Upstream description:

Inline TLS parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-2102312112320033-3003111123300021-2213200001312100-2322312322302101-1310132012013202-1011110132322120-3012131113030122-0022303033323322"></a>

## Direct properties — tls_parameters / 121300303212 / 3

- [no_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3103211131002331-2002013001220022-2320310201010130-1010003223300220-3103100013213202-3013320310020003-1001102312233132-3232203210132002): complete subsection reference.

- [tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231): complete subsection reference.

- [tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011): complete subsection reference.

- [use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100): complete subsection reference.

<a id="canonical-1132202330303313-2223021302030023-3220220003201223-3122301103133303-1321030313011223-0332223030333211-2033203013001223-1000320111310102"></a>

## Next pages — tls_parameters / 121300303212 / 4

- [proxy_config.https.tls_parameters.no_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3103211131002331-2002013001220022-2320310201010130-1010003223300220-3103100013213202-3013320310020003-1001102312233132-3232203210132002)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011)
- [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012031302302213-3323300232020033-1301011321131202-2032103232231301-1310102102120123-1203130131123330-0310133011222322-2103323133012100)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3103211131002331-2002013001220022-2320310201010130-1010003223300220-3103100013213202-3013320310020003-1001102312233132-3232203210132002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032002303310122-0220210033200330-0210121003311323-3131121311023010-3322233300123222-2201013301231022-0323220233011133-1233212102302101"></a>

## proxy_config.https.tls_parameters.no_mtls — no_mtls / 113002320131 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- proxy_config.https.tls_parameters.no_mtls

<a id="canonical-1100013033220213-2233310102033130-2110202321202203-0212132030023300-3002210011323221-1030212023013000-1233010010213232-3022110103121021"></a>

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

<a id="canonical-1333012320100001-3010321203323213-3033333012301032-1210123210302332-2313211032223332-1220130201103031-2233333131313331-2011323002213310"></a>

## Direct properties — no_mtls / 113002320131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300332312123323-1313030021231022-2130231112000331-1231220022100021-0033320330211031-1213232221223200-1333213012330120-1312111213121101"></a>

## Next pages — no_mtls / 113002320131 / 4

- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202021203001012-1203120300122312-3302233231321321-0301102130331111-3113211221322130-3213312230110103-0110321233200210-3132333101000303"></a>

## proxy_config.https.tls_parameters.tls_certificates — tls_certificates / 032321312320 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- proxy_config.https.tls_parameters.tls_certificates

<a id="canonical-3320323100113312-2220312130011303-0123030113323211-0031222103100213-1032213130131310-1222222303020021-3221203123122013-0031331013033103"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

Upstream description:

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3110131132313113-0222303310031113-1333302313130011-0332023130213211-0232333001220203-3323003212330311-3103010313303230-0021232002232230"></a>

## Direct properties — tls_certificates / 032321312320 / 3

<a id="canonical-0303030323131121-1333000132000120-1111313300303203-1033022133313022-3300212333011333-0323130003312211-1023202110212233-0133202001000132"></a>

<a id="canonical-3021320300032001-0111310020203033-1321031003011322-0012333130101100-0310220113200002-1110112200211212-0230002213121201-1001133021331213"></a>

## certificate_url property — tls_certificates / 032321312320 / 4

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2011203310232100-2112322003133230-3222313123011110-0213232303232213-3010110112021031-2112233010302131-2222021030010301-1231111230333102): complete subsection reference.

<a id="canonical-0100230010000020-3010103313122030-3131230202203330-1131102001303213-1320021022133103-2232121013030232-0132212131231212-2100003332323300"></a>

<a id="canonical-1332320131030311-3321300221133000-2010220202303231-3311302020201100-3330311133013120-2023000232200111-3203120230332233-0312231211311230"></a>

## description_spec property — tls_certificates / 032321312320 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1330011020332031-1220021211113001-0032321202221101-1303313333203300-1233202302310313-2202123030110222-0202120123130330-3230310122210000): complete subsection reference.

- [private_key](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0000220132002111-1232110221010003-3013333010111212-0131023200022123-1322312110000313-0111300211131013-3312000132210230-2023010020111232): complete subsection reference.

- [use_system_defaults](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3132031010020200-0021303130021321-1210330310132010-1311331201300313-2111210101232023-3213113222211121-3303231303200301-1322122102032001): complete subsection reference.

<a id="canonical-1123023100302110-3211123001331120-1210000231231212-3203032200220223-0203121120203112-1322003303131302-0122323021331322-0121210313322313"></a>

## Next pages — tls_certificates / 032321312320 / 6

- [proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2011203310232100-2112322003133230-3222313123011110-0213232303232213-3010110112021031-2112233010302131-2222021030010301-1231111230333102)
- [proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1330011020332031-1220021211113001-0032321202221101-1303313333203300-1233202302310313-2202123030110222-0202120123130330-3230310122210000)
- [proxy_config.https.tls_parameters.tls_certificates.private_key](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0000220132002111-1232110221010003-3013333010111212-0131023200022123-1322312110000313-0111300211131013-3312000132210230-2023010020111232)
- [proxy_config.https.tls_parameters.tls_certificates.use_system_defaults](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3132031010020200-0021303130021321-1210330310132010-1311331201300313-2111210101232023-3213113222211121-3303231303200301-1322122102032001)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2011203310232100-2112322003133230-3222313123011110-0213232303232213-3010110112021031-2112233010302131-2222021030010301-1231111230333102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103202013030330-1301231013112130-1300112132202122-2312220002012332-1002132201022010-2022232310000101-1302313222230022-2320312310233320"></a>

## proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 030132332023 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-1121011121201310-3022001113102121-1313021322231002-0023013031111212-1122011302123011-1131302130233032-2222311120022111-2013003023101212"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0200211230203112-0213320130201333-1031123120312231-0001322210000123-0323222120232001-0210111002331133-2011032030023133-1101323021000003"></a>

## Direct properties — custom_hash_algorithms / 030132332023 / 3

<a id="canonical-2200020112131100-2031330330022003-1030303013233020-1222233033031010-2232303233121010-1132220102013113-3023303112330023-2231011330123321"></a>

<a id="canonical-2011200111201120-1321233200031322-1213303022001100-2311230230011032-3330210122300313-3313200302231003-1331210113121200-0211313112303120"></a>

## hash_algorithms property — custom_hash_algorithms / 030132332023 / 4

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1230302021002010-1213230223213321-1303320213012222-3032030033112203-0032300223011131-0322303021023010-3030331210001210-3110002210002321"></a>

## Next pages — custom_hash_algorithms / 030132332023 / 5

- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1330011020332031-1220021211113001-0032321202221101-1303313333203300-1233202302310313-2202123030110222-0202120123130330-3230310122210000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000110200222103-1202231101102233-1120333002220111-1323013232112220-3300302013312303-0322101011310333-0231121000330103-0223303220131112"></a>

## proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 030013020321 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-3112011131032311-0111312110200132-2202003321313111-2101210131233203-1213033313000333-2132032011010300-1131132021300230-1113230332332201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

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

<a id="canonical-3320301311123212-2133121232312131-2203233101013211-0030000320101210-3132102010301303-1202131220330330-3311213330012001-1031030313120311"></a>

## Direct properties — disable_ocsp_stapling / 030013020321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1230211133021232-1333330032323331-3203120201331300-1031221222113130-1020010230033132-1023203023300121-1300001030320200-3030100303210301"></a>

## Next pages — disable_ocsp_stapling / 030013020321 / 4

- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0000220132002111-1232110221010003-3013333010111212-0131023200022123-1322312110000313-0111300211131013-3312000132210230-2023010020111232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102200302110003-2231033003230300-2312332022003301-1130232003320332-3132131223233102-3123200320101113-3113310020132310-3233213001112320"></a>

## proxy_config.https.tls_parameters.tls_certificates.private_key — private_key / 131113102211 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- proxy_config.https.tls_parameters.tls_certificates.private_key

<a id="canonical-0032223323211231-2113320120213000-0111230213022012-1220010202312311-0313123311113020-3013021103123033-1111221300031323-3232030331233301"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-1123013003231312-3011032212032221-2301101000033031-3212300233223021-0333112313103301-2133022013322120-1110220100131302-2000101313213021"></a>

## Direct properties — private_key / 131113102211 / 3

- [blindfold_secret_info](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0131001320232323-3330311030331133-2031230303203103-2112210013223332-2300330322032011-2110201333113021-3131320322203220-1010222132020223): complete subsection reference.

- [clear_secret_info](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1312333132310200-0103210210300030-1222330122312223-0132012200110323-1131102112223003-2130201033122030-1120331213303233-1313223330302311): complete subsection reference.

<a id="canonical-0213020100110332-2331321033111322-2112111111202002-0112121312022103-0103110300213113-3320320112130330-3110222232101230-2111322321331302"></a>

## Next pages — private_key / 131113102211 / 4

- [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0131001320232323-3330311030331133-2031230303203103-2112210013223332-2300330322032011-2110201333113021-3131320322203220-1010222132020223)
- [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1312333132310200-0103210210300030-1222330122312223-0132012200110323-1131102112223003-2130201033122030-1120331213303233-1313223330302311)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0131001320232323-3330311030331133-2031230303203103-2112210013223332-2300330322032011-2110201333113021-3131320322203220-1010222132020223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000310201232212-3223102000221220-3023301221022222-0001230002122331-0331010102210002-2232301102103230-0220103212233130-0102223031231022"></a>

## proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 232230001033 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- [proxy_config.https.tls_parameters.tls_certificates.private_key](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0000220132002111-1232110221010003-3013333010111212-0131023200022123-1322312110000313-0111300211131013-3312000132210230-2023010020111232)
- proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-2200023321001131-3311220312221020-1111210312231103-2011132202220311-3120213113122313-2020312010101211-2223220022021331-1300312002130220"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3122332330020232-0322032033332330-0100320011121113-1231331010001100-0133013223201121-0030333120031022-3103103333210032-2330221212101212"></a>

## Direct properties — blindfold_secret_info / 232230001033 / 3

<a id="canonical-3203210032121330-0230123312002310-3132321111223331-2132203301333223-0000322102001232-2333100313111131-2110320130312331-0230331103121012"></a>

<a id="canonical-1311213222321313-3120201321333233-2220100101200231-1233110122102223-3310023020032020-3003210301033330-3030233032121301-0223322301103102"></a>

## decryption_provider property — blindfold_secret_info / 232230001033 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-3133203300202213-2010110021031120-2310111103022013-1323330032110333-2212011002120023-3032333113111010-1331223321113011-1100132330230011"></a>

<a id="canonical-3031211112032013-3331032133230132-1302312200120202-2112103030333302-3002332312231102-1120201230232323-3223323230323023-0332220120131012"></a>

## location property — blindfold_secret_info / 232230001033 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0003033313021112-1122233222001103-2123023031212310-2312121130212102-0112122323100230-3100021033101012-2022331230120120-2323320112301022"></a>

<a id="canonical-1230320100303303-0203200021011033-0111220010133012-2001033103031232-2322331122303233-2031201132123002-0303103112222233-1331122230023233"></a>

## store_provider property — blindfold_secret_info / 232230001033 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-3003301113111310-3120111113103303-2302203123210333-1202011011211100-0023212301231103-3201013313031300-2230333100003301-0331032231232111"></a>

## Next pages — blindfold_secret_info / 232230001033 / 7

- [proxy_config.https.tls_parameters.tls_certificates.private_key](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0000220132002111-1232110221010003-3013333010111212-0131023200022123-1322312110000313-0111300211131013-3312000132210230-2023010020111232)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-1312333132310200-0103210210300030-1222330122312223-0132012200110323-1131102112223003-2130201033122030-1120331213303233-1313223330302311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010333303310130-3101130313203213-0330213100300301-2322031120331202-0220221123213133-0022320000213200-3323330333301101-0103023030021313"></a>

## proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info — clear_secret_info / 300203033123 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- [proxy_config.https.tls_parameters.tls_certificates.private_key](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0000220132002111-1232110221010003-3013333010111212-0131023200022123-1322312110000313-0111300211131013-3312000132210230-2023010020111232)
- proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-2203023003102211-1300330303212210-1122003002201130-3121000201231202-1130012303321003-0313301130211310-2232130102233210-3023203302023122"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0131300131331032-2002231023123023-3011123212220302-0310201202001103-2330102232000222-2133133112001131-0300112312331021-0020223213011030"></a>

## Direct properties — clear_secret_info / 300203033123 / 3

<a id="canonical-0212003223201121-3213222022213233-3101233023311001-3230213033130132-2300211102020031-3130000231303130-1201202331123231-1131133303033012"></a>

<a id="canonical-3130012330013031-1100001101133302-1031311313311221-1212210112113130-0302130220212333-0020012021022320-3301220203110311-1231322321103222"></a>

## provider_ref property — clear_secret_info / 300203033123 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2133231200012200-0102101010102020-3221030200001102-2323110103000120-1010100031101002-0332010003130021-1220011022133310-0020200023312210"></a>

<a id="canonical-2012220213131020-2110223331001001-2113003231102131-0112223330202120-1012320102220301-3331332232221220-1311311021010322-2211322001310113"></a>

## URL property — clear_secret_info / 300203033123 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0310002311120303-3011111332112231-3311301110131011-3013223331032023-0122023210131233-3021330101222033-3133223213313333-0022133311320110"></a>

## Next pages — clear_secret_info / 300203033123 / 6

- [proxy_config.https.tls_parameters.tls_certificates.private_key](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0000220132002111-1232110221010003-3013333010111212-0131023200022123-1322312110000313-0111300211131013-3312000132210230-2023010020111232)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3132031010020200-0021303130021321-1210330310132010-1311331201300313-2111210101232023-3213113222211121-3303231303200301-1322122102032001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000021110031133-2210013323322121-3312303220010012-0130121321323301-3220313300032101-1311303230001201-2332131020023212-2102332303220233"></a>

## proxy_config.https.tls_parameters.tls_certificates.use_system_defaults — use_system_defaults / 233220001222 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- proxy_config.https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-2131012012330020-2201200032231133-2222102001211331-0131002102002102-3033021312133232-2013003331001100-3033000320101203-3030003312221022"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

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

<a id="canonical-0132321110323030-3011010101233220-0210203023320302-1111033331002103-3332103012300011-0303020301102233-3310100122021311-1210121120331112"></a>

## Direct properties — use_system_defaults / 233220001222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210233300131110-2030301300103102-3223320332101133-0003033001330010-2023310333221112-0133333101312313-2003320313201133-2030310333231203"></a>

## Next pages — use_system_defaults / 233220001222 / 4

- [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3230332130201103-1120101011302320-0032131332012022-3211131312023303-0103120111213333-1102323013320312-1310023331312020-2333133011221231)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102013203132030-3312103200202123-2311312301302032-2211320101212302-2303312123200113-0221013221133210-3102331001121310-0032202113032210"></a>

## proxy_config.https.tls_parameters.tls_config — tls_config / 121123003113 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- proxy_config.https.tls_parameters.tls_config

<a id="canonical-2121230122113200-0212020330310131-0310333011030132-2120020003233100-2213122010011030-3120303122121221-3023023123230012-1233211002133223"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-0212210302010000-2012000033302123-0313121110021101-1220130122220210-2231233111021322-0020130203311201-0303000223323001-1100023303220110"></a>

## Direct properties — tls_config / 121123003113 / 3

- [custom_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0330120033123122-2112211002012033-2303001101101321-3031102232122003-2132013101320230-1032121303230203-2320031322320132-0232011311021220): complete subsection reference.

- [default_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0332311203233100-1201333122301330-3002021230212111-3303122103203312-1201220101000032-2202002330111331-3221302003313312-2111021121310130): complete subsection reference.

- [low_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0310021031211322-3303021001013103-2113023211023302-2321110031011133-2012021320132323-2321232102013100-0102200213333221-1200311321111231): complete subsection reference.

- [medium_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2311030031330313-2102313002123011-3302012221210203-1303232331330002-2102003202300333-2103311033100210-2032330201201000-3203222122103120): complete subsection reference.

<a id="canonical-3200211321201303-2211210100001333-0020003310201221-2312002302002212-3203023010223111-0021021321322021-0333312311213320-0013133023012101"></a>

## Next pages — tls_config / 121123003113 / 4

- [proxy_config.https.tls_parameters.tls_config.custom_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0330120033123122-2112211002012033-2303001101101321-3031102232122003-2132013101320230-1032121303230203-2320031322320132-0232011311021220)
- [proxy_config.https.tls_parameters.tls_config.default_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0332311203233100-1201333122301330-3002021230212111-3303122103203312-1201220101000032-2202002330111331-3221302003313312-2111021121310130)
- [proxy_config.https.tls_parameters.tls_config.low_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0310021031211322-3303021001013103-2113023211023302-2321110031011133-2012021320132323-2321232102013100-0102200213333221-1200311321111231)
- [proxy_config.https.tls_parameters.tls_config.medium_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2311030031330313-2102313002123011-3302012221210203-1303232331330002-2102003202300333-2103311033100210-2032330201201000-3203222122103120)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0330120033123122-2112211002012033-2303001101101321-3031102232122003-2132013101320230-1032121303230203-2320031322320132-0232011311021220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301312131031120-1223011232200301-2201122103331212-1320232133331032-0330322331001033-1203130312310123-0322002331221301-2311130233311330"></a>

## proxy_config.https.tls_parameters.tls_config.custom_security — custom_security / 000333001113 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011)
- proxy_config.https.tls_parameters.tls_config.custom_security

<a id="canonical-3122231310302001-2030100011103000-3010312033223302-1100113123132312-1011022121100210-1221222110213221-3020301332233102-0130122203120331"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0033101120133222-3232002303323122-0123303011000201-0211103112013232-0022000002130022-3020101221002020-0201201001202100-1330011021323201"></a>

## Direct properties — custom_security / 000333001113 / 3

<a id="canonical-3031001212333300-2013213133112030-0332131201101222-1333333211220120-3121010210303133-0003220000111020-0120322013113321-2302123221231201"></a>

<a id="canonical-3022321102212332-3322032330002311-3030203012232300-1220002332110303-3200333112133012-2332313133133120-3032021131231001-1202132120322310"></a>

## cipher_suites property — custom_security / 000333001113 / 4

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0020212002313020-2303323230131023-1331203132000210-0001322003331313-1320101023232123-3130033320301020-2332303322303022-3101112300302220"></a>

<a id="canonical-1110231103320003-1201120230331230-0333000121011032-0311202031100011-2111300113033133-3311013231201122-0211123000222012-3301330320223020"></a>

## max_version property — custom_security / 000333001113 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3321200203313332-3032003010003302-0322331010323230-3212201330231131-0131230201223001-3320021113230012-1202101021031310-2122132232102012"></a>

<a id="canonical-0013200302331020-0032102323102132-2310330221203102-3231202130001212-2110222131033211-0221301113102122-3121130313213320-2203333111003231"></a>

## min_version property — custom_security / 000333001113 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0233301101323222-0332300112110123-2332201021031103-1110200011112202-1031020111312031-3110202210032200-2200103013321010-1032130102103102"></a>

## Next pages — custom_security / 000333001113 / 7

- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0332311203233100-1201333122301330-3002021230212111-3303122103203312-1201220101000032-2202002330111331-3221302003313312-2111021121310130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300112030313330-0312210010220101-1133220010103122-2110022301301120-1230212122312130-2232230230111212-3022302031313112-2320013031222310"></a>

## proxy_config.https.tls_parameters.tls_config.default_security — default_security / 310121322003 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011)
- proxy_config.https.tls_parameters.tls_config.default_security

<a id="canonical-2012022031013200-3131321000000311-0031322211310113-0010013021013330-3331320322210131-0320123231201133-3312120321001331-2301230233202320"></a>

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

<a id="canonical-0212122310021110-3021233321311101-3301231100001332-2020023300033200-1021013312111321-1031221323301230-1301120131123000-3001121102030331"></a>

## Direct properties — default_security / 310121322003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030011112102310-0012302210000223-2301220231211210-3232303221222032-0302231123122013-2313220221121032-1212011013121232-0223032333013112"></a>

## Next pages — default_security / 310121322003 / 4

- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-0310021031211322-3303021001013103-2113023211023302-2321110031011133-2012021320132323-2321232102013100-0102200213333221-1200311321111231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232021213002033-0210022003322333-0220110120200333-1212301210233233-2303023223202300-0330103113130300-3312202210212303-1121302210313300"></a>

## proxy_config.https.tls_parameters.tls_config.low_security — low_security / 310200311332 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011)
- proxy_config.https.tls_parameters.tls_config.low_security

<a id="canonical-1220231112001133-3313001100211310-1222032222212112-1310102202331033-0112231212113212-0212203020113330-0202202232232102-3110311223112322"></a>

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

<a id="canonical-0232200102113320-2331111031231101-3022100013201312-0102020220312123-2232132133030013-2321320320301333-0120302132002330-0221321223033031"></a>

## Direct properties — low_security / 310200311332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003320220311202-1122221333032133-2232002213102231-1010013230010231-2230020313033231-0330223223101233-3012031333102212-3331333312002020"></a>

## Next pages — low_security / 310200311332 / 4

- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011)
- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)

<a id="canonical-2311030031330313-2102313002123011-3302012221210203-1303232331330002-2102003202300333-2103311033100210-2032330201201000-3203222122103120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332212123313203-3122102020002031-0310310123313011-1213321000220212-0313111320221120-1031000110303010-2323223303220333-2032131011022223"></a>

## proxy_config.https.tls_parameters.tls_config.medium_security — medium_security / 131002133222 / 2

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322)
- [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3110001102111133-2000223133210133-0322102231212333-3100320000212223-0120130310121013-2120131112130020-3313012000111032-1013000111001321)
- [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0102230112122011-1210231113333231-0320030100023200-1211113211300233-3102020333202111-2331033321101322-0111000031020032-2020233110232100)
- [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3320113112122021-2132120121121301-2110033111312202-1310332120320201-1103121122122221-3031212121013121-3200130303000021-3131213312102011)
- proxy_config.https.tls_parameters.tls_config.medium_security

<a id="canonical-3011212221001302-2002211233102332-3221321301030222-0320210030023333-3203313003301312-2111330333032013-0120230310133201-0230331011312221"></a>

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

<a id="canonical-3023213120302112-1222233003130003-3333230310302213-1003200120311110-1022132331332012-1000132030032100-0113312110102023-3211212211322312"></a>

## Direct properties — medium_security / 131002133222 / 3

This is an empty object or choice marker. It has no direct properties.
